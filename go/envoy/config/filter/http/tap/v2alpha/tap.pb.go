// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/tap/v2alpha/tap.proto

package envoy_config_filter_http_tap_v2alpha

import (
	fmt "fmt"
	v2alpha "github.com/cilium/proxy/go/envoy/config/common/tap/v2alpha"
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

// Top level configuration for the tap filter.
type Tap struct {
	// Common configuration for the HTTP tap filter.
	CommonConfig         *v2alpha.CommonExtensionConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *Tap) Reset()         { *m = Tap{} }
func (m *Tap) String() string { return proto.CompactTextString(m) }
func (*Tap) ProtoMessage()    {}
func (*Tap) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee77d938a401b876, []int{0}
}

func (m *Tap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tap.Unmarshal(m, b)
}
func (m *Tap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tap.Marshal(b, m, deterministic)
}
func (m *Tap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tap.Merge(m, src)
}
func (m *Tap) XXX_Size() int {
	return xxx_messageInfo_Tap.Size(m)
}
func (m *Tap) XXX_DiscardUnknown() {
	xxx_messageInfo_Tap.DiscardUnknown(m)
}

var xxx_messageInfo_Tap proto.InternalMessageInfo

func (m *Tap) GetCommonConfig() *v2alpha.CommonExtensionConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*Tap)(nil), "envoy.config.filter.http.tap.v2alpha.Tap")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/tap/v2alpha/tap.proto", fileDescriptor_ee77d938a401b876)
}

var fileDescriptor_ee77d938a401b876 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xb1, 0x4a, 0x03, 0x41,
	0x10, 0x86, 0x59, 0x25, 0x12, 0x4e, 0x05, 0x49, 0xa3, 0xa4, 0x10, 0x11, 0x11, 0x0b, 0x99, 0x85,
	0x0b, 0xf8, 0x00, 0x17, 0xec, 0x83, 0xa4, 0x97, 0x31, 0xd9, 0x24, 0x0b, 0x7b, 0x3b, 0xcb, 0xdd,
	0x78, 0x24, 0xaf, 0x60, 0x63, 0xeb, 0x73, 0x5a, 0x5a, 0x88, 0xdc, 0xce, 0x1e, 0xdc, 0x75, 0x76,
	0xc3, 0xcc, 0x37, 0xff, 0x3f, 0xf3, 0x67, 0x60, 0x7c, 0x43, 0x07, 0xbd, 0x22, 0xbf, 0xb1, 0x5b,
	0xbd, 0xb1, 0x8e, 0x4d, 0xa5, 0x77, 0xcc, 0x41, 0x33, 0x06, 0xdd, 0xe4, 0xe8, 0xc2, 0x0e, 0xdb,
	0x1a, 0x42, 0x45, 0x4c, 0x93, 0xbb, 0xc8, 0x83, 0xf0, 0x20, 0x3c, 0xb4, 0x3c, 0xb4, 0x4c, 0xe2,
	0xa7, 0x8f, 0x03, 0xd5, 0x15, 0x95, 0x25, 0xf9, 0x81, 0xa0, 0xb4, 0x44, 0x73, 0x7a, 0xfd, 0xbe,
	0x0e, 0xa8, 0xd1, 0x7b, 0x62, 0x64, 0x4b, 0xbe, 0xd6, 0xa5, 0xdd, 0x56, 0xc8, 0x26, 0xcd, 0x2f,
	0x1b, 0x74, 0x76, 0x8d, 0x6c, 0x74, 0x57, 0xc8, 0xe0, 0xd6, 0x65, 0xc7, 0x4b, 0x0c, 0x13, 0x93,
	0x9d, 0x8b, 0xde, 0xab, 0x18, 0x5e, 0xa9, 0x1b, 0xf5, 0x70, 0x9a, 0x3f, 0xc1, 0xe0, 0xd6, 0x64,
	0xd9, 0x3b, 0x13, 0xe6, 0xb1, 0xf5, 0xbc, 0x67, 0xe3, 0x6b, 0x4b, 0x7e, 0x1e, 0xc1, 0x62, 0xfc,
	0x53, 0x8c, 0x3e, 0xd4, 0xd1, 0x85, 0x7a, 0x39, 0x93, 0x9d, 0xd4, 0x77, 0xdf, 0x5f, 0xbf, 0x9f,
	0xa3, 0xfb, 0x2e, 0x02, 0xd3, 0xad, 0xd5, 0x29, 0x86, 0xba, 0x97, 0xc3, 0x2c, 0xcb, 0x2d, 0x89,
	0x7f, 0xa8, 0x68, 0x7f, 0x80, 0xff, 0xc4, 0x56, 0x8c, 0x97, 0x18, 0x16, 0xed, 0x67, 0x0b, 0xf5,
	0x76, 0x12, 0x5f, 0x9c, 0xfd, 0x05, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x92, 0x17, 0x8a, 0xa1, 0x01,
	0x00, 0x00,
}
