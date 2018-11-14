#include "mux_socket.h"

#include "common/common/assert.h"
#include "common/common/empty_string.h"
#include "common/network/address_impl.h"
#include "common/network/listen_socket_impl.h"
#include "envoy/registry/registry.h"

namespace Envoy {
namespace Cilium {

std::string MuxSocketName("cilium.transport_sockets.mux");

typedef std::function<void()> ReadCB;

// Data for a mux keyed by the fd
// Not sure yet if access from multiple threads is actually needed
class MuxData {
public:
  MuxData(Mux& mux, const ShimTuple& id) : mux_(mux), id_(id) {}

  void setReadCallback(ReadCB cb) {
    ENVOY_LOG_MISC(trace, "MUX setting read callback");
    readCallback_ = cb;
  }

  Mux& mux_;
  const ShimTuple id_;
  ReadCB readCallback_{};

  mutable Thread::MutexBasicLockable lock_;
  Buffer::OwnedImpl read_buffer_ GUARDED_BY(lock_);  // Input frames ready to be read
  bool end_stream_{false} GUARDED_BY(lock_);
};

// Multiplexed read buffers by fd
static Thread::MutexBasicLockable mux_lock;
static std::map<int, MuxData*> muxed_buffers GUARDED_BY(mux_lock);

void MuxSocket::setTransportSocketCallbacks(Network::TransportSocketCallbacks& callbacks) {
  callbacks_ = &callbacks;

  // Find mux_data_ based on the fd
  fd_ = callbacks_->fd();

  Thread::LockGuard guard(mux_lock);
  auto it = muxed_buffers.find(fd_);
  if (it != muxed_buffers.end()) {
    ENVOY_LOG_MISC(trace, "MUX found muxed read buffer for fd {}", fd_);
    mux_data_ = it->second;
    mux_data_->setReadCallback([this]() {
	ENVOY_LOG_MISC(trace, "MUX SETTING read buffer ready");
	callbacks_->setReadBufferReady();
      });
  } else {
    ENVOY_LOG_MISC(trace, "MUX DID NOT find muxed read buffer for fd {}", fd_);
  }
}

void MuxSocket::closeSocket(Network::ConnectionEvent event) {
  ENVOY_LOG_MISC(trace, "MuxSocket::closeSocket({})", int(event));
  if (event == Network::ConnectionEvent::RemoteClose) {
    if (mux_data_ != nullptr) {
      mux_data_->mux_.removeBuffer(fd_);
      mux_data_ = nullptr;
    }
  }
}

MuxSocket::~MuxSocket() {
  closeSocket(Network::ConnectionEvent::RemoteClose);
}

Network::IoResult MuxSocket::doRead(Buffer::Instance& buffer) {
  if (mux_data_ == nullptr) {
    ENVOY_LOG_MISC(trace, "No MUX data!");
    return {Network::PostIoAction::Close, 0, true};
  }

  // Kick the mux transport to read data
  mux_data_->mux_.readAndDemux();

  // Move all available data to the caller's buffer
  mux_data_->lock_.lock();
  bool end_stream = mux_data_->end_stream_;
  uint64_t bytes_read = mux_data_->read_buffer_.length();
  buffer.move(mux_data_->read_buffer_);
  mux_data_->lock_.unlock();

  ENVOY_LOG_MISC(trace, "[{}] doRead read {} bytes", callbacks_->connection().id(), bytes_read);

  if (callbacks_->shouldDrainReadBuffer()) {
    ENVOY_LOG_MISC(trace, "[{}] doRead calling setReadBufferReady()", callbacks_->connection().id());
    callbacks_->setReadBufferReady();
  }

  return {Network::PostIoAction::KeepOpen, bytes_read, end_stream};
}

Network::IoResult MuxSocket::doWrite(Buffer::Instance& buffer, bool end_stream) {
  if (mux_data_ == nullptr) {
    ENVOY_LOG_MISC(trace, "No MUX data!");
    return {Network::PostIoAction::Close, 0, false};
  }

  Network::PostIoAction action;
  uint64_t bytes_written = 0;
  ASSERT(!shutdown_ || buffer.length() == 0);
  do {
    if (buffer.length() == 0) {
      if (end_stream && !shutdown_) {
        // Ignore the result. This can only fail if the connection failed. In that case, the
        // error will be detected on the next read, and dealt with appropriately.
        ::shutdown(fd_, SHUT_WR);
        shutdown_ = true;
      }
      action = Network::PostIoAction::KeepOpen;
      break;
    }
    Api::SysCallIntResult result = mux_data_->mux_.prependAndWrite(mux_data_->id_, buffer);
    ENVOY_CONN_LOG(trace, "write returns: {}", callbacks_->connection(), result.rc_);

    if (result.rc_ == -1) {
      ENVOY_CONN_LOG(trace, "write error: {} ({})", callbacks_->connection(), result.errno_,
                     strerror(result.errno_));
      if (result.errno_ == EAGAIN) {
        action = Network::PostIoAction::KeepOpen;
      } else {
        action = Network::PostIoAction::Close;
      }
      break;
    } else {
      bytes_written += result.rc_;
    }
  } while (true);

  return {action, bytes_written, false};
}

std::string MuxSocket::protocol() const { return EMPTY_STRING; }

void MuxSocket::onConnected() { callbacks_->raiseEvent(Network::ConnectionEvent::Connected); }

Mux::Mux(Event::Dispatcher& dispatcher, Network::ConnectionSocket& socket, NewConnectionCB addNewConnetion, CloseMuxCB closeMux)
  : socket_(socket), addNewConnection_(addNewConnetion), closeMux_(closeMux) {
  file_event_ =
    dispatcher.createFileEvent(socket_.fd(),
			       [this](uint32_t events) {
				 if (events & Event::FileReadyType::Read) {
				   onRead();
				 }
				 if (events & Event::FileReadyType::Write) {
				   onWrite();
				 }
				 if (events & Event::FileReadyType::Closed) {
				   onClose();
				 }
			       },
			       Event::FileTriggerType::Edge,
			       Event::FileReadyType::Read |
			       Event::FileReadyType::Write |
			       Event::FileReadyType::Closed);

  timer_ = dispatcher.createTimer([this]() -> void { onTimeout(); });
  timer_->enableTimer(std::chrono::milliseconds(1));

#if 1
  // TCP proxy does not connected in the test unless we short-circuit the connection set-up.
  ShimTuple id{};
  auto ip = socket_.localAddress()->ip();
  auto ip_src = socket_.remoteAddress()->ip();
  if (ip->ipv4()) {
    id.tuple_v4.daddr_ = ip->ipv4()->address();
    id.tuple_v4.dport_ = htons(ip->port());
    id.tuple_v4.saddr_ = ip_src->ipv4()->address();
    id.tuple_v4.sport_ = htons(ip_src->port());
  } else if (ip->ipv6()) {
    id.tuple_v6.daddr_ = ip->ipv6()->address();
    id.tuple_v6.dport_ = htons(ip->port());
    id.tuple_v6.saddr_ = ip_src->ipv6()->address();
    id.tuple_v6.sport_ = htons(ip_src->port());
  } else {
    ENVOY_LOG_MISC(trace, "MUX unknown IP address format!");
    return;
  }
  addBuffer(id);
#endif
}

Mux::~Mux() {
  ENVOY_LOG_MISC(trace, "MUX checking for buffers to clean up...");
  Thread::LockGuard guard(lock_);
  Thread::LockGuard mux_guard(mux_lock);
  for (auto it = muxed_buffers.begin(); it != muxed_buffers.end(); ) {
    MuxData* mux_data = it->second;
    if (&mux_data->mux_ == this) {
      ENVOY_LOG_MISC(trace, "MUX found buffer to clean up");
      std::vector<uint32_t> key(mux_data->id_._tuple_, std::end(mux_data->id_._tuple_));
      it = muxed_buffers.erase(it);
      if (current_reader_ == mux_data) {
	current_reader_ = nullptr;
      }
      buffers_.erase(key); // Frees the MuxData object
    } else {
      it++;
    }
  }
  ASSERT(buffers_.empty());
}

void Mux::onRead() {
  ENVOY_LOG_MISC(trace, "MUX test: onRead()");
  readAndDemux();
}

void Mux::onWrite() {
  ENVOY_LOG_MISC(trace, "MUX test: onWrite({})", write_buffer_.length());
  if (write_buffer_.length() > 0) {
    Api::SysCallIntResult result = write_buffer_.write(socket_.fd());
    ENVOY_LOG_MISC(trace, "MUX write returns: {}", result.rc_);
  }
}

void Mux::onTimeout() {
  ENVOY_LOG_MISC(trace, "MUX test: timeout");
  timer_.reset();
#if 0
  // Create a copy of the socket and pass it to newConnection callback.
  int fd2 = dup(socket_.fd());
  ASSERT(fd2 >= 0, "dup() failed");

  Network::ConnectionSocketPtr sock = std::make_unique<Network::ConnectionSocketImpl>(fd2, socket_.localAddress(), socket_.remoteAddress());
  sock->addOptions(socket_.options()); // copy a referene to the options on the original socket.
  if (socket_.localAddressRestored()) {
    sock->setLocalAddress(socket_.localAddress(), true);
  }
  ENVOY_LOG_MISC(trace, "MUX test: newConnection on dupped fd {}", fd2);

  addNewConnection_(std::move(sock));
#endif
}

void Mux::onClose() {
  ENVOY_LOG_MISC(debug, "MUX test: Closing socket");
  // Try flushing the data one more time
  onRead();
  onWrite();
  timer_.reset();
  file_event_.reset();
}

// called with 'lock_' held!
MuxData* Mux::addBuffer(const ShimTuple& id) {
  // Create a copy of the socket and pass it to addNewConnection callback.
  int fd2 = dup(socket_.fd());
  ASSERT(fd2 >= 0, "dup() failed");

  std::vector<uint32_t> key(id._tuple_, std::end(id._tuple_));
  // 'buffers_' owns the MuxData objects!
  auto pair = buffers_.emplace(key, std::make_unique<MuxData>(*this, id));
  ASSERT(pair.second == true); // inserted
  MuxData* mux_data = pair.first->second.get();
			       
  // Add to the static index by fd as well
  {
    Thread::LockGuard guard(mux_lock);
    muxed_buffers.emplace(fd2, mux_data);
  }

  struct sockaddr_storage ss_src{}, ss_dst{};
  size_t ss_size;
  // Call the addNewConnection callback
  if (id.tuple_v6.dport_ != 0) {
    struct sockaddr_in6* sin_src = reinterpret_cast<struct sockaddr_in6*>(&ss_src);
    ss_size = sizeof(*sin_src);
    sin_src->sin6_family = AF_INET6;
    memcpy(static_cast<void*>(&sin_src->sin6_addr.s6_addr), static_cast<const void*>(&id.tuple_v6.saddr_), 
	   sizeof(absl::uint128));
    sin_src->sin6_port = id.tuple_v6.sport_;

    struct sockaddr_in6* sin_dst = reinterpret_cast<struct sockaddr_in6*>(&ss_dst);
    sin_dst->sin6_family = AF_INET6;
    memcpy(static_cast<void*>(&sin_dst->sin6_addr.s6_addr), static_cast<const void*>(&id.tuple_v6.daddr_), 
	   sizeof(absl::uint128));
    sin_dst->sin6_port = id.tuple_v6.dport_;
  } else {
    struct sockaddr_in* sin_src = reinterpret_cast<struct sockaddr_in*>(&ss_src);
    ss_size = sizeof(*sin_src);
    sin_src->sin_family = AF_INET;
    sin_src->sin_addr.s_addr = id.tuple_v4.saddr_;
    sin_src->sin_port = id.tuple_v4.sport_;

    struct sockaddr_in* sin_dst = reinterpret_cast<struct sockaddr_in*>(&ss_dst);
    sin_dst->sin_family = AF_INET;
    sin_dst->sin_addr.s_addr = id.tuple_v4.daddr_;
    sin_dst->sin_port = id.tuple_v4.dport_;
  }
  Network::ConnectionSocketPtr sock =
    std::make_unique<Network::ConnectionSocketImpl>(fd2, Network::Address::addressFromSockAddr(ss_dst, ss_size, false), Network::Address::addressFromSockAddr(ss_src, ss_size, false));
  sock->addOptions(socket_.options()); // copy a reference to the options on the original socket.
  if (socket_.localAddressRestored()) {
    sock->setLocalAddress(sock->localAddress(), true);
  }
  sock->setDetectedTransportProtocol(MuxSocketName);

  ENVOY_LOG_MISC(trace, "MUX test: newConnection on dupped fd {}", fd2);

  addNewConnection_(std::move(sock));
  return mux_data;
}

void Mux::removeBuffer(int fd) {
  Thread::LockGuard guard(lock_);
  Thread::LockGuard mux_guard(mux_lock);
  auto it = muxed_buffers.find(fd);
  if (it != muxed_buffers.end()) {
    ENVOY_LOG_MISC(trace, "MUX found buffer to delete");
    std::vector<uint32_t> key(it->second->id_._tuple_, std::end(it->second->id_._tuple_));
    muxed_buffers.erase(it);
    if (current_reader_ == it->second) {
      current_reader_ = nullptr;
    }
    buffers_.erase(key); // Frees the MuxData
    if (muxed_buffers.empty()) {
      ENVOY_LOG_MISC(trace, "MUX no muxed sockets left, closing the mux");
      closeMux_();
    }
  }
}

void Mux::readAndDemux() {
  Thread::LockGuard guard(lock_);
  do {
    // 16K read is arbitrary. TODO(mattklein123) PERF: Tune the read size.
    Api::SysCallIntResult result = read_buffer_.read(socket_.fd(), 16384);
    ENVOY_LOG_MISC(trace, "MUX read returns: {}, read buffer length: {}", result.rc_, read_buffer_.length());

    if (result.rc_ == 0) {
      // Remote close.
      ENVOY_LOG_MISC(trace, "MUX returned zero bytes, ending streams");
      for (auto it = buffers_.begin(); it != buffers_.end(); it++) {
	Thread::LockGuard buffer_guard(it->second->lock_);
	it->second->end_stream_ = true;
      }
      break;
    } else if (result.rc_ == -1) {
      // Remote error (might be no data).
      if (result.errno_ != EAGAIN) {
	ENVOY_LOG_MISC(trace, "MUX read error: {}", result.errno_);
	for (auto it = buffers_.begin(); it != buffers_.end(); it++) {
	  Thread::LockGuard buffer_guard(it->second->lock_);
	  it->second->end_stream_ = true;
	}
      }
      break;
    } else {
      // distribute the read bytes to demuxed read buffers
      do {
	// Are we in a middle of a partial frame?
	if (remaining_read_length_ > 0) {
	  auto len = std::min(remaining_read_length_, read_buffer_.length());
	  if (current_reader_ != nullptr) {
	    // Move input to the demuxed socket
	    Thread::LockGuard buffer_guard(current_reader_->lock_);
	    ENVOY_LOG_MISC(trace, "MUX transfering {} bytes", len);
	    current_reader_->read_buffer_.move(read_buffer_, len);
	    // Wake the dupped fd (not sure if necessary)
	    if (current_reader_->readCallback_) {
	      current_reader_->readCallback_();
	    }
	  } else {
	    // Demuxed socket has been closed before all input was received. Drain the bytes.
	    ENVOY_LOG_MISC(trace, "MUX draining {} bytes", len);
	    read_buffer_.drain(len);
	  }
	  remaining_read_length_ -= len;
	}
	if (remaining_read_length_ == 0) {
	  current_reader_ = nullptr;

	  // Do we have enough data for the next shim header?
	  if (read_buffer_.length() >= sizeof(ShimHeader)) {
#if 0
	    ShimHeader hdr{};
	    read_buffer_.copyOut(0, sizeof(ShimHeader), &hdr);
	    read_buffer_.drain(sizeof(ShimHeader));
	    remaining_read_length_ = hdr.length_;
	    std::vector<uint32_t> key(hdr.id_._tuple_, std::end(hdr.id_._tuple_));
#else
	    remaining_read_length_ = read_buffer_.length();
	    ShimTuple id{};
	    auto ip = socket_.localAddress()->ip();
	    auto ip_src = socket_.remoteAddress()->ip();
	    if (ip->ipv4()) {
	      id.tuple_v4.daddr_ = ip->ipv4()->address();
	      id.tuple_v4.dport_ = htons(ip->port());
	      id.tuple_v4.saddr_ = ip_src->ipv4()->address();
	      id.tuple_v4.sport_ = htons(ip_src->port());
	    } else if (ip->ipv6()) {
	      id.tuple_v6.daddr_ = ip->ipv6()->address();
	      id.tuple_v6.dport_ = htons(ip->port());
	      id.tuple_v6.saddr_ = ip_src->ipv6()->address();
	      id.tuple_v6.sport_ = htons(ip_src->port());
	    } else {
	      ENVOY_LOG_MISC(trace, "MUX unknown IP address format!");
	      return;
	    }
	    std::vector<uint32_t> key(id._tuple_, std::end(id._tuple_));
#endif
	    auto it = buffers_.find(key);
	    if (it != buffers_.end()) {
	      if (remaining_read_length_ > 0 ) {
		ENVOY_LOG_MISC(trace, "MUX found buffer for frame length {}", remaining_read_length_);
		current_reader_ = it->second.get();
	      } else {
		ENVOY_LOG_MISC(trace, "MUX found buffer, but closing for 0-length frame");
		Thread::LockGuard buffer_guard(it->second->lock_);
		it->second->end_stream_ = true;
		// current_reader_ remoins nullptr
	      }
	    } else if (remaining_read_length_ > 0) {
	      // New connection?
	      ENVOY_LOG_MISC(trace, "MUX did NOT find a buffer, creating a new one for frame length {}", remaining_read_length_);
#if 0
	      current_reader_ = addBuffer(hdr.id_);
#else
	      current_reader_ = addBuffer(id);
#endif
	    }
	  }
	}
      } while (remaining_read_length_ > 0 && read_buffer_.length() > 0);
    }
  } while (true);
}

Api::SysCallIntResult Mux::prependAndWrite(const ShimTuple& /*id*/, Buffer::Instance& buffer) {
  if (buffer.length() == 0) {
    return {0, 0};
  }
  Thread::LockGuard guard(lock_);

  int len = std::min(int(buffer.length()), 16384); // Limit for fearness?
#if 0
  // Prepend a shim header for the data
  ShimHeader shim;
  shim.length_ = len;
  shim.id_ = id;
  write_buffer_.add(&shim, sizeof(shim));

  Api::SysCallIntResult result = write_buffer_.write(socket_.fd());
  ENVOY_LOG_MISC(trace, "SHIM write returns: {}", result.rc_);

  if (result.rc_ == -1) {
    if (result.errno_ != EAGAIN) {
      ENVOY_LOG_MISC(trace, "write error: {} ({})", result.errno_, strerror(result.errno_));
    }
    // Remove the added Shim from the buffer on error
    Buffer::OwnedImpl temp;
    temp.move(write_buffer_, write_buffer_.length() - sizeof(shim));
    write_buffer_.drain(sizeof(shim));
    write_buffer_.move(temp);

    return result;
  }
#endif

  write_buffer_.move(buffer, len);
  Api::SysCallIntResult result = write_buffer_.write(socket_.fd());
  ENVOY_LOG_MISC(trace, "MUX write returns: {}", result.rc_);
  // We keep the data in the buffer, so pretend we sent it.
  return {len, 0};
}

Network::TransportSocketPtr MuxSocketFactory::createTransportSocket() const {
  return std::make_unique<MuxSocket>();
}

bool MuxSocketFactory::implementsSecureTransport() const { return false; }

Network::TransportSocketFactoryPtr UpstreamMuxSocketConfigFactory::createTransportSocketFactory(
    const Protobuf::Message&, Server::Configuration::TransportSocketFactoryContext&) {
  return std::make_unique<MuxSocketFactory>();
}

Network::TransportSocketFactoryPtr DownstreamMuxSocketConfigFactory::createTransportSocketFactory(
    const Protobuf::Message&, Server::Configuration::TransportSocketFactoryContext&,
    const std::vector<std::string>&) {
  return std::make_unique<MuxSocketFactory>();
}

ProtobufTypes::MessagePtr MuxSocketConfigFactory::createEmptyConfigProto() {
  return std::make_unique<ProtobufWkt::Empty>();
}

static Registry::RegisterFactory<UpstreamMuxSocketConfigFactory,
                                 Server::Configuration::UpstreamTransportSocketConfigFactory>
    upstream_registered_;

static Registry::RegisterFactory<DownstreamMuxSocketConfigFactory,
                                 Server::Configuration::DownstreamTransportSocketConfigFactory>
    downstream_registered_;

} // namespace Cilium
} // namespace Envoy
