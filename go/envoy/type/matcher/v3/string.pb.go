// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/v3/string.proto

package envoy_type_matcher_v3

import (
	fmt "fmt"
	_ "github.com/cilium/proxy/go/envoy/annotations"
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

// Specifies the way to match a string.
// [#next-free-field: 7]
type StringMatcher struct {
	// Types that are valid to be assigned to MatchPattern:
	//	*StringMatcher_Exact
	//	*StringMatcher_Prefix
	//	*StringMatcher_Suffix
	//	*StringMatcher_SafeRegex
	MatchPattern isStringMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
	// If true, indicates the exact/prefix/suffix matching should be case insensitive. This has no
	// effect for the safe_regex match.
	// For example, the matcher *data* will match both input string *Data* and *data* if set to true.
	IgnoreCase           bool     `protobuf:"varint,6,opt,name=ignore_case,json=ignoreCase,proto3" json:"ignore_case,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringMatcher) Reset()         { *m = StringMatcher{} }
func (m *StringMatcher) String() string { return proto.CompactTextString(m) }
func (*StringMatcher) ProtoMessage()    {}
func (*StringMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_e33cffa01bf36e0e, []int{0}
}

func (m *StringMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringMatcher.Unmarshal(m, b)
}
func (m *StringMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringMatcher.Marshal(b, m, deterministic)
}
func (m *StringMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringMatcher.Merge(m, src)
}
func (m *StringMatcher) XXX_Size() int {
	return xxx_messageInfo_StringMatcher.Size(m)
}
func (m *StringMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_StringMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_StringMatcher proto.InternalMessageInfo

type isStringMatcher_MatchPattern interface {
	isStringMatcher_MatchPattern()
}

type StringMatcher_Exact struct {
	Exact string `protobuf:"bytes,1,opt,name=exact,proto3,oneof"`
}

type StringMatcher_Prefix struct {
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3,oneof"`
}

type StringMatcher_Suffix struct {
	Suffix string `protobuf:"bytes,3,opt,name=suffix,proto3,oneof"`
}

type StringMatcher_SafeRegex struct {
	SafeRegex *RegexMatcher `protobuf:"bytes,5,opt,name=safe_regex,json=safeRegex,proto3,oneof"`
}

func (*StringMatcher_Exact) isStringMatcher_MatchPattern() {}

func (*StringMatcher_Prefix) isStringMatcher_MatchPattern() {}

func (*StringMatcher_Suffix) isStringMatcher_MatchPattern() {}

func (*StringMatcher_SafeRegex) isStringMatcher_MatchPattern() {}

func (m *StringMatcher) GetMatchPattern() isStringMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (m *StringMatcher) GetExact() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Exact); ok {
		return x.Exact
	}
	return ""
}

func (m *StringMatcher) GetPrefix() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Prefix); ok {
		return x.Prefix
	}
	return ""
}

func (m *StringMatcher) GetSuffix() string {
	if x, ok := m.GetMatchPattern().(*StringMatcher_Suffix); ok {
		return x.Suffix
	}
	return ""
}

func (m *StringMatcher) GetSafeRegex() *RegexMatcher {
	if x, ok := m.GetMatchPattern().(*StringMatcher_SafeRegex); ok {
		return x.SafeRegex
	}
	return nil
}

func (m *StringMatcher) GetIgnoreCase() bool {
	if m != nil {
		return m.IgnoreCase
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StringMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StringMatcher_Exact)(nil),
		(*StringMatcher_Prefix)(nil),
		(*StringMatcher_Suffix)(nil),
		(*StringMatcher_SafeRegex)(nil),
	}
}

// Specifies a list of ways to match a string.
type ListStringMatcher struct {
	Patterns             []*StringMatcher `protobuf:"bytes,1,rep,name=patterns,proto3" json:"patterns,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListStringMatcher) Reset()         { *m = ListStringMatcher{} }
func (m *ListStringMatcher) String() string { return proto.CompactTextString(m) }
func (*ListStringMatcher) ProtoMessage()    {}
func (*ListStringMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_e33cffa01bf36e0e, []int{1}
}

func (m *ListStringMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListStringMatcher.Unmarshal(m, b)
}
func (m *ListStringMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListStringMatcher.Marshal(b, m, deterministic)
}
func (m *ListStringMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListStringMatcher.Merge(m, src)
}
func (m *ListStringMatcher) XXX_Size() int {
	return xxx_messageInfo_ListStringMatcher.Size(m)
}
func (m *ListStringMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_ListStringMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_ListStringMatcher proto.InternalMessageInfo

func (m *ListStringMatcher) GetPatterns() []*StringMatcher {
	if m != nil {
		return m.Patterns
	}
	return nil
}

func init() {
	proto.RegisterType((*StringMatcher)(nil), "envoy.type.matcher.v3.StringMatcher")
	proto.RegisterType((*ListStringMatcher)(nil), "envoy.type.matcher.v3.ListStringMatcher")
}

func init() { proto.RegisterFile("envoy/type/matcher/v3/string.proto", fileDescriptor_e33cffa01bf36e0e) }

var fileDescriptor_e33cffa01bf36e0e = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4f, 0x8e, 0xd3, 0x30,
	0x18, 0xc5, 0xc7, 0x69, 0x53, 0x32, 0x8e, 0x46, 0x1a, 0x22, 0xfe, 0x44, 0x5d, 0x40, 0x9a, 0x8e,
	0x20, 0x12, 0x52, 0x22, 0x4d, 0x59, 0xcd, 0xd2, 0x6c, 0xaa, 0x0a, 0x50, 0x15, 0x0e, 0x50, 0x99,
	0xf4, 0x6b, 0xb1, 0x04, 0x76, 0x64, 0xbb, 0x51, 0x7a, 0x03, 0x84, 0x58, 0xb1, 0xe4, 0x08, 0xdc,
	0x83, 0x3b, 0xa1, 0xae, 0x90, 0xed, 0x50, 0x51, 0x35, 0xb3, 0x8b, 0xf3, 0x7e, 0xef, 0x7b, 0xef,
	0xb3, 0x71, 0x0a, 0xbc, 0x11, 0xfb, 0x42, 0xef, 0x6b, 0x28, 0xbe, 0x50, 0x5d, 0x7d, 0x02, 0x59,
	0x34, 0xb3, 0x42, 0x69, 0xc9, 0xf8, 0x36, 0xaf, 0xa5, 0xd0, 0x22, 0x7a, 0x6c, 0x99, 0xdc, 0x30,
	0x79, 0xc7, 0xe4, 0xcd, 0x6c, 0x3c, 0xe9, 0xb7, 0x4a, 0xd8, 0x42, 0xeb, 0x9c, 0xe3, 0xc9, 0x6e,
	0x5d, 0xd3, 0x82, 0x72, 0x2e, 0x34, 0xd5, 0x4c, 0x70, 0x55, 0x34, 0x20, 0x15, 0x13, 0xfc, 0x38,
	0x7c, 0x3c, 0x75, 0x53, 0xfe, 0x67, 0xd6, 0x50, 0x4b, 0xa8, 0xec, 0xa1, 0x83, 0x9e, 0x36, 0xf4,
	0x33, 0x5b, 0x53, 0x0d, 0xc5, 0xbf, 0x0f, 0x27, 0xa4, 0xbf, 0x3c, 0x7c, 0xf5, 0xc1, 0x76, 0x7d,
	0xe7, 0x1a, 0x44, 0x4f, 0xb0, 0x0f, 0x2d, 0xad, 0x74, 0x8c, 0x12, 0x94, 0x5d, 0xce, 0x2f, 0x4a,
	0x77, 0x8c, 0x26, 0x78, 0x54, 0x4b, 0xd8, 0xb0, 0x36, 0xf6, 0x8c, 0x40, 0x1e, 0x1c, 0xc8, 0x50,
	0x7a, 0x09, 0x9a, 0x5f, 0x94, 0x9d, 0x60, 0x10, 0xb5, 0xdb, 0x18, 0x64, 0x70, 0x86, 0x38, 0x21,
	0x7a, 0x8f, 0xb1, 0xa2, 0x1b, 0x58, 0xd9, 0x25, 0x63, 0x3f, 0x41, 0x59, 0x78, 0x3b, 0xcd, 0x7b,
	0xef, 0x27, 0x2f, 0x0d, 0xd3, 0xd5, 0x22, 0xc1, 0x81, 0xf8, 0xdf, 0x90, 0x77, 0x6d, 0x86, 0x5d,
	0x9a, 0x11, 0x56, 0x8d, 0x9e, 0xe3, 0x90, 0x6d, 0xb9, 0x90, 0xb0, 0xaa, 0xa8, 0x82, 0x78, 0x94,
	0xa0, 0x2c, 0x28, 0xb1, 0xfb, 0xf5, 0x86, 0x2a, 0xb8, 0x7b, 0xf9, 0xf3, 0xf7, 0xd7, 0x67, 0x29,
	0x4e, 0x7a, 0x22, 0x4e, 0xf6, 0x26, 0x8f, 0xf0, 0x95, 0x15, 0x56, 0x35, 0xd5, 0x1a, 0x24, 0x8f,
	0x06, 0x7f, 0x08, 0x5a, 0x0c, 0x83, 0xe1, 0xb5, 0x5f, 0xfa, 0xb6, 0x6e, 0xfa, 0x1d, 0xe1, 0x87,
	0x6f, 0x99, 0xd2, 0xa7, 0x17, 0xb6, 0xc0, 0x41, 0x67, 0x51, 0x31, 0x4a, 0x06, 0x59, 0x78, 0x7b,
	0x73, 0xcf, 0x42, 0xa7, 0x81, 0x66, 0xa3, 0x1f, 0xc8, 0x0b, 0x50, 0x79, 0xf4, 0xdf, 0xbd, 0x32,
	0x6d, 0x5f, 0xe0, 0x9b, 0x1e, 0xff, 0x59, 0x30, 0x79, 0x8d, 0xa7, 0x4c, 0xb8, 0xa8, 0x5a, 0x8a,
	0x76, 0xdf, 0x9f, 0x4a, 0x42, 0xe7, 0x5a, 0x9a, 0xf7, 0x5e, 0xa2, 0x8f, 0x23, 0xfb, 0xf0, 0xb3,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xed, 0x99, 0xbe, 0xb9, 0x02, 0x00, 0x00,
}
