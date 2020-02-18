// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/network/local_ratelimit/v3/local_rate_limit.proto

package envoy_extensions_filters_network_local_ratelimit_v3

import (
	fmt "fmt"
	v31 "github.com/cilium/proxy/go/envoy/config/core/v3"
	v3 "github.com/cilium/proxy/go/envoy/type/v3"
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

type LocalRateLimit struct {
	// The prefix to use when emitting :ref:`statistics
	// <config_network_filters_local_rate_limit_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The token bucket configuration to use for rate limiting connections that are processed by the
	// filter's filter chain. Each incoming connection processed by the filter consumes a single
	// token. If the token is available, the connection will be allowed. If no tokens are available,
	// the connection will be immediately closed.
	//
	// .. note::
	//   In the current implementation each filter and filter chain has an independent rate limit.
	//
	// .. note::
	//   In the current implementation the token bucket's :ref:`fill_interval
	//   <envoy_api_field_type.v3.TokenBucket.fill_interval>` must be >= 50ms to avoid too aggressive
	//   refills.
	TokenBucket *v3.TokenBucket `protobuf:"bytes,2,opt,name=token_bucket,json=tokenBucket,proto3" json:"token_bucket,omitempty"`
	// Runtime flag that controls whether the filter is enabled or not. If not specified, defaults
	// to enabled.
	RuntimeEnabled       *v31.RuntimeFeatureFlag `protobuf:"bytes,3,opt,name=runtime_enabled,json=runtimeEnabled,proto3" json:"runtime_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *LocalRateLimit) Reset()         { *m = LocalRateLimit{} }
func (m *LocalRateLimit) String() string { return proto.CompactTextString(m) }
func (*LocalRateLimit) ProtoMessage()    {}
func (*LocalRateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffceb0cea724f411, []int{0}
}

func (m *LocalRateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalRateLimit.Unmarshal(m, b)
}
func (m *LocalRateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalRateLimit.Marshal(b, m, deterministic)
}
func (m *LocalRateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalRateLimit.Merge(m, src)
}
func (m *LocalRateLimit) XXX_Size() int {
	return xxx_messageInfo_LocalRateLimit.Size(m)
}
func (m *LocalRateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalRateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_LocalRateLimit proto.InternalMessageInfo

func (m *LocalRateLimit) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *LocalRateLimit) GetTokenBucket() *v3.TokenBucket {
	if m != nil {
		return m.TokenBucket
	}
	return nil
}

func (m *LocalRateLimit) GetRuntimeEnabled() *v31.RuntimeFeatureFlag {
	if m != nil {
		return m.RuntimeEnabled
	}
	return nil
}

func init() {
	proto.RegisterType((*LocalRateLimit)(nil), "envoy.extensions.filters.network.local_ratelimit.v3.LocalRateLimit")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/network/local_ratelimit/v3/local_rate_limit.proto", fileDescriptor_ffceb0cea724f411)
}

var fileDescriptor_ffceb0cea724f411 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xbf, 0xae, 0xd3, 0x30,
	0x14, 0xc6, 0x95, 0xf0, 0xdf, 0x45, 0x05, 0x85, 0x81, 0x2a, 0x03, 0x04, 0xa6, 0x4c, 0xb6, 0xd4,
	0x6c, 0x6c, 0x04, 0xb5, 0x48, 0x55, 0x87, 0x12, 0x31, 0xb1, 0x44, 0x4e, 0x7a, 0x5a, 0xac, 0xba,
	0x76, 0xe4, 0x9c, 0x84, 0xf6, 0x0d, 0x10, 0x8f, 0xc0, 0xc8, 0xbb, 0xf0, 0x52, 0x9d, 0xae, 0x6c,
	0xe7, 0xaa, 0xed, 0xd5, 0x9d, 0xee, 0x16, 0x9d, 0xf3, 0xf9, 0x77, 0xbe, 0x7c, 0x1f, 0x59, 0x80,
	0xea, 0xf5, 0x91, 0xc1, 0x01, 0x41, 0xb5, 0x42, 0xab, 0x96, 0x6d, 0x84, 0x44, 0x30, 0x2d, 0x53,
	0x80, 0xbf, 0xb4, 0xd9, 0x31, 0xa9, 0x6b, 0x2e, 0x4b, 0xc3, 0x11, 0xa4, 0xd8, 0x0b, 0x64, 0x7d,
	0x76, 0x31, 0x2a, 0xdd, 0x8c, 0x36, 0x46, 0xa3, 0x8e, 0x32, 0xc7, 0xa2, 0x67, 0x16, 0x1d, 0x58,
	0x74, 0x60, 0xd1, 0x3b, 0x2c, 0xda, 0x67, 0xf1, 0x7b, 0x6f, 0xa0, 0xd6, 0x6a, 0x23, 0xb6, 0xac,
	0xd6, 0x06, 0xec, 0x85, 0x8a, 0xb7, 0xe0, 0xa9, 0x71, 0xe2, 0x05, 0x78, 0x6c, 0xdc, 0x06, 0xf5,
	0x0e, 0x54, 0x59, 0x75, 0xf5, 0x0e, 0x86, 0xbb, 0xf1, 0x87, 0x6e, 0xdd, 0x70, 0xc6, 0x95, 0xd2,
	0xc8, 0xd1, 0xfd, 0x43, 0x0f, 0xc6, 0x1a, 0x10, 0x6a, 0x3b, 0x48, 0xde, 0xf6, 0x5c, 0x8a, 0x35,
	0x47, 0x60, 0xb7, 0x1f, 0x7e, 0xf1, 0xf1, 0x5f, 0x48, 0xc6, 0x4b, 0xeb, 0xaa, 0xe0, 0x08, 0x4b,
	0x6b, 0x2a, 0x4a, 0xc9, 0xa8, 0x45, 0x8e, 0x65, 0x63, 0x60, 0x23, 0x0e, 0x93, 0x20, 0x09, 0xd2,
	0x17, 0xf9, 0xb3, 0x53, 0xfe, 0xd8, 0x84, 0x49, 0x50, 0x10, 0xbb, 0x5b, 0xb9, 0x55, 0xf4, 0x95,
	0xbc, 0xbc, 0xb4, 0x33, 0x09, 0x93, 0x20, 0x1d, 0x4d, 0x63, 0xea, 0x73, 0xb0, 0x8e, 0x69, 0x9f,
	0xd1, 0xef, 0x56, 0x92, 0x3b, 0x45, 0xfe, 0xfc, 0x94, 0x3f, 0xf9, 0x13, 0x84, 0xaf, 0x83, 0x62,
	0x84, 0xe7, 0x71, 0xf4, 0x8d, 0xbc, 0x32, 0x9d, 0x42, 0xb1, 0x87, 0x12, 0x14, 0xaf, 0x24, 0xac,
	0x27, 0x8f, 0x1c, 0x2b, 0x1d, 0x58, 0x3e, 0x1e, 0x6a, 0xe3, 0xb1, 0xc8, 0xc2, 0x8b, 0xe7, 0xc0,
	0xb1, 0x33, 0x30, 0x97, 0x7c, 0x5b, 0x8c, 0x07, 0xc0, 0xcc, 0xbf, 0xff, 0xb4, 0xf8, 0xfb, 0xff,
	0xf7, 0xbb, 0x19, 0xf9, 0x72, 0xf5, 0xde, 0xf7, 0x71, 0x4f, 0x1d, 0x43, 0x8f, 0xfd, 0x94, 0xcb,
	0xe6, 0x27, 0xa7, 0xd7, 0x89, 0xe4, 0x3f, 0xc8, 0x67, 0xa1, 0xbd, 0x93, 0xc6, 0xe8, 0xc3, 0x91,
	0x3e, 0xa0, 0xe8, 0xfc, 0xcd, 0x35, 0x74, 0x65, 0xe3, 0x5f, 0x05, 0xd5, 0x53, 0xd7, 0x43, 0x76,
	0x13, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x76, 0xcb, 0xd2, 0x89, 0x02, 0x00, 0x00,
}
