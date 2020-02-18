// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/transport_sockets/alts/v3/alts.proto

package envoy_extensions_transport_sockets_alts_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Configuration for ALTS transport socket. This provides Google's ALTS protocol to Envoy.
// https://cloud.google.com/security/encryption-in-transit/application-layer-transport-security/
type Alts struct {
	// The location of a handshaker service, this is usually 169.254.169.254:8080
	// on GCE.
	HandshakerService string `protobuf:"bytes,1,opt,name=handshaker_service,json=handshakerService,proto3" json:"handshaker_service,omitempty"`
	// The acceptable service accounts from peer, peers not in the list will be rejected in the
	// handshake validation step. If empty, no validation will be performed.
	PeerServiceAccounts  []string `protobuf:"bytes,2,rep,name=peer_service_accounts,json=peerServiceAccounts,proto3" json:"peer_service_accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Alts) Reset()         { *m = Alts{} }
func (m *Alts) String() string { return proto.CompactTextString(m) }
func (*Alts) ProtoMessage()    {}
func (*Alts) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b780b78c099503e, []int{0}
}

func (m *Alts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Alts.Unmarshal(m, b)
}
func (m *Alts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Alts.Marshal(b, m, deterministic)
}
func (m *Alts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Alts.Merge(m, src)
}
func (m *Alts) XXX_Size() int {
	return xxx_messageInfo_Alts.Size(m)
}
func (m *Alts) XXX_DiscardUnknown() {
	xxx_messageInfo_Alts.DiscardUnknown(m)
}

var xxx_messageInfo_Alts proto.InternalMessageInfo

func (m *Alts) GetHandshakerService() string {
	if m != nil {
		return m.HandshakerService
	}
	return ""
}

func (m *Alts) GetPeerServiceAccounts() []string {
	if m != nil {
		return m.PeerServiceAccounts
	}
	return nil
}

func init() {
	proto.RegisterType((*Alts)(nil), "envoy.extensions.transport_sockets.alts.v3.Alts")
}

func init() {
	proto.RegisterFile("envoy/extensions/transport_sockets/alts/v3/alts.proto", fileDescriptor_3b780b78c099503e)
}

var fileDescriptor_3b780b78c099503e = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x95, 0x52, 0x81, 0xe2, 0x8d, 0x20, 0x44, 0xd5, 0x01, 0x05, 0xa6, 0x8a, 0xc1, 0x16,
	0x8d, 0xa8, 0x10, 0x5b, 0xb3, 0xb0, 0x56, 0xe5, 0x00, 0xd1, 0x23, 0x31, 0x8d, 0xd5, 0xc8, 0xcf,
	0xb2, 0x5f, 0xad, 0xf6, 0x06, 0x9c, 0x81, 0x23, 0x70, 0x0f, 0x2e, 0xc5, 0x84, 0xe2, 0x04, 0x65,
	0xe8, 0xd2, 0xc9, 0x96, 0xfe, 0xff, 0x7b, 0xf6, 0xf7, 0xd8, 0x93, 0xd4, 0x1e, 0x0f, 0x42, 0xee,
	0x49, 0x6a, 0xa7, 0x50, 0x3b, 0x41, 0x16, 0xb4, 0x33, 0x68, 0xa9, 0x70, 0x58, 0x6e, 0x25, 0x39,
	0x01, 0x0d, 0x39, 0xe1, 0xb3, 0x70, 0x72, 0x63, 0x91, 0x30, 0x79, 0x08, 0x18, 0x1f, 0x30, 0x7e,
	0x84, 0xf1, 0x50, 0xf7, 0xd9, 0xf4, 0x6e, 0x57, 0x19, 0x10, 0xa0, 0x35, 0x12, 0x50, 0x78, 0xc2,
	0x4b, 0xdb, 0x42, 0x4a, 0x6f, 0xba, 0x71, 0xd3, 0x1b, 0x0f, 0x8d, 0xaa, 0x80, 0xa4, 0xf8, 0xbf,
	0x74, 0xc1, 0xfd, 0x77, 0xc4, 0xc6, 0xcb, 0x86, 0x5c, 0xb2, 0x60, 0x49, 0x0d, 0xba, 0x72, 0x35,
	0x6c, 0xa5, 0x2d, 0x9c, 0xb4, 0x5e, 0x95, 0x72, 0x12, 0xa5, 0xd1, 0x2c, 0xce, 0x2f, 0x7e, 0xf3,
	0xb1, 0x1d, 0xa5, 0xd1, 0xfa, 0x72, 0xa8, 0xbc, 0x75, 0x8d, 0x64, 0xce, 0xae, 0x8d, 0x1c, 0x88,
	0x02, 0xca, 0x12, 0x77, 0x9a, 0xdc, 0x64, 0x94, 0x9e, 0xcd, 0xe2, 0xf5, 0x55, 0x1b, 0xf6, 0xdd,
	0x65, 0x1f, 0xbd, 0x2c, 0xbe, 0x7e, 0x3e, 0x6f, 0x1f, 0x99, 0xe8, 0x1c, 0x4b, 0xd4, 0x1f, 0x6a,
	0x73, 0xe4, 0xd7, 0xeb, 0xcd, 0xa1, 0x31, 0x35, 0xf0, 0xf6, 0x8f, 0xf9, 0x2b, 0x7b, 0x56, 0xc8,
	0x03, 0x65, 0x2c, 0xee, 0x0f, 0xfc, 0xf4, 0x25, 0xe5, 0x71, 0x3b, 0x61, 0xd5, 0x3a, 0xaf, 0xa2,
	0xf7, 0xf3, 0x20, 0x9f, 0xfd, 0x05, 0x00, 0x00, 0xff, 0xff, 0x78, 0x69, 0xef, 0x30, 0x9d, 0x01,
	0x00, 0x00,
}
