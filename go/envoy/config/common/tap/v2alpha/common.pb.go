// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/common/tap/v2alpha/common.proto

package envoy_config_common_tap_v2alpha

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	v2alpha "github.com/cilium/proxy/go/envoy/service/tap/v2alpha"
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

// Common configuration for all tap extensions.
type CommonExtensionConfig struct {
	// Types that are valid to be assigned to ConfigType:
	//	*CommonExtensionConfig_AdminConfig
	//	*CommonExtensionConfig_StaticConfig
	//	*CommonExtensionConfig_TapdsConfig
	ConfigType           isCommonExtensionConfig_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *CommonExtensionConfig) Reset()         { *m = CommonExtensionConfig{} }
func (m *CommonExtensionConfig) String() string { return proto.CompactTextString(m) }
func (*CommonExtensionConfig) ProtoMessage()    {}
func (*CommonExtensionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_79cf139d98a2fe3f, []int{0}
}

func (m *CommonExtensionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonExtensionConfig.Unmarshal(m, b)
}
func (m *CommonExtensionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonExtensionConfig.Marshal(b, m, deterministic)
}
func (m *CommonExtensionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonExtensionConfig.Merge(m, src)
}
func (m *CommonExtensionConfig) XXX_Size() int {
	return xxx_messageInfo_CommonExtensionConfig.Size(m)
}
func (m *CommonExtensionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonExtensionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CommonExtensionConfig proto.InternalMessageInfo

type isCommonExtensionConfig_ConfigType interface {
	isCommonExtensionConfig_ConfigType()
}

type CommonExtensionConfig_AdminConfig struct {
	AdminConfig *AdminConfig `protobuf:"bytes,1,opt,name=admin_config,json=adminConfig,proto3,oneof"`
}

type CommonExtensionConfig_StaticConfig struct {
	StaticConfig *v2alpha.TapConfig `protobuf:"bytes,2,opt,name=static_config,json=staticConfig,proto3,oneof"`
}

type CommonExtensionConfig_TapdsConfig struct {
	TapdsConfig *CommonExtensionConfig_TapDSConfig `protobuf:"bytes,3,opt,name=tapds_config,json=tapdsConfig,proto3,oneof"`
}

func (*CommonExtensionConfig_AdminConfig) isCommonExtensionConfig_ConfigType() {}

func (*CommonExtensionConfig_StaticConfig) isCommonExtensionConfig_ConfigType() {}

func (*CommonExtensionConfig_TapdsConfig) isCommonExtensionConfig_ConfigType() {}

func (m *CommonExtensionConfig) GetConfigType() isCommonExtensionConfig_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *CommonExtensionConfig) GetAdminConfig() *AdminConfig {
	if x, ok := m.GetConfigType().(*CommonExtensionConfig_AdminConfig); ok {
		return x.AdminConfig
	}
	return nil
}

func (m *CommonExtensionConfig) GetStaticConfig() *v2alpha.TapConfig {
	if x, ok := m.GetConfigType().(*CommonExtensionConfig_StaticConfig); ok {
		return x.StaticConfig
	}
	return nil
}

func (m *CommonExtensionConfig) GetTapdsConfig() *CommonExtensionConfig_TapDSConfig {
	if x, ok := m.GetConfigType().(*CommonExtensionConfig_TapdsConfig); ok {
		return x.TapdsConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CommonExtensionConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CommonExtensionConfig_AdminConfig)(nil),
		(*CommonExtensionConfig_StaticConfig)(nil),
		(*CommonExtensionConfig_TapdsConfig)(nil),
	}
}

// [#not-implemented-hide:]
type CommonExtensionConfig_TapDSConfig struct {
	// Configuration for the source of TapDS updates for this Cluster.
	ConfigSource *core.ConfigSource `protobuf:"bytes,1,opt,name=config_source,json=configSource,proto3" json:"config_source,omitempty"`
	// Tap config to request from XDS server.
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonExtensionConfig_TapDSConfig) Reset()         { *m = CommonExtensionConfig_TapDSConfig{} }
func (m *CommonExtensionConfig_TapDSConfig) String() string { return proto.CompactTextString(m) }
func (*CommonExtensionConfig_TapDSConfig) ProtoMessage()    {}
func (*CommonExtensionConfig_TapDSConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_79cf139d98a2fe3f, []int{0, 0}
}

func (m *CommonExtensionConfig_TapDSConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonExtensionConfig_TapDSConfig.Unmarshal(m, b)
}
func (m *CommonExtensionConfig_TapDSConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonExtensionConfig_TapDSConfig.Marshal(b, m, deterministic)
}
func (m *CommonExtensionConfig_TapDSConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonExtensionConfig_TapDSConfig.Merge(m, src)
}
func (m *CommonExtensionConfig_TapDSConfig) XXX_Size() int {
	return xxx_messageInfo_CommonExtensionConfig_TapDSConfig.Size(m)
}
func (m *CommonExtensionConfig_TapDSConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonExtensionConfig_TapDSConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CommonExtensionConfig_TapDSConfig proto.InternalMessageInfo

func (m *CommonExtensionConfig_TapDSConfig) GetConfigSource() *core.ConfigSource {
	if m != nil {
		return m.ConfigSource
	}
	return nil
}

func (m *CommonExtensionConfig_TapDSConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Configuration for the admin handler. See :ref:`here <config_http_filters_tap_admin_handler>` for
// more information.
type AdminConfig struct {
	// Opaque configuration ID. When requests are made to the admin handler, the passed opaque ID is
	// matched to the configured filter opaque ID to determine which filter to configure.
	ConfigId             string   `protobuf:"bytes,1,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdminConfig) Reset()         { *m = AdminConfig{} }
func (m *AdminConfig) String() string { return proto.CompactTextString(m) }
func (*AdminConfig) ProtoMessage()    {}
func (*AdminConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_79cf139d98a2fe3f, []int{1}
}

func (m *AdminConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdminConfig.Unmarshal(m, b)
}
func (m *AdminConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdminConfig.Marshal(b, m, deterministic)
}
func (m *AdminConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdminConfig.Merge(m, src)
}
func (m *AdminConfig) XXX_Size() int {
	return xxx_messageInfo_AdminConfig.Size(m)
}
func (m *AdminConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AdminConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AdminConfig proto.InternalMessageInfo

func (m *AdminConfig) GetConfigId() string {
	if m != nil {
		return m.ConfigId
	}
	return ""
}

func init() {
	proto.RegisterType((*CommonExtensionConfig)(nil), "envoy.config.common.tap.v2alpha.CommonExtensionConfig")
	proto.RegisterType((*CommonExtensionConfig_TapDSConfig)(nil), "envoy.config.common.tap.v2alpha.CommonExtensionConfig.TapDSConfig")
	proto.RegisterType((*AdminConfig)(nil), "envoy.config.common.tap.v2alpha.AdminConfig")
}

func init() {
	proto.RegisterFile("envoy/config/common/tap/v2alpha/common.proto", fileDescriptor_79cf139d98a2fe3f)
}

var fileDescriptor_79cf139d98a2fe3f = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xd1, 0xea, 0xd3, 0x30,
	0x14, 0xc6, 0xed, 0xf6, 0xff, 0xeb, 0x96, 0x76, 0x20, 0x01, 0x51, 0x2a, 0x6c, 0x63, 0x4c, 0xf1,
	0x62, 0x26, 0xd0, 0x3d, 0x81, 0x99, 0x82, 0x22, 0xc8, 0xdc, 0xbc, 0x1f, 0xc7, 0x36, 0xce, 0xc0,
	0x9a, 0x84, 0x36, 0x2b, 0x9b, 0x8f, 0xe0, 0x8d, 0xb7, 0x3e, 0x86, 0xcf, 0xe6, 0x95, 0xec, 0x42,
	0xa4, 0x49, 0xea, 0x3a, 0x50, 0x76, 0xd7, 0xe6, 0x7c, 0xe7, 0x77, 0xbe, 0xf3, 0x25, 0x68, 0xc6,
	0x65, 0xa5, 0x8e, 0x34, 0x55, 0xf2, 0x93, 0xd8, 0xd2, 0x54, 0xe5, 0xb9, 0x92, 0xd4, 0x80, 0xa6,
	0x55, 0x02, 0x3b, 0xfd, 0x19, 0xfc, 0x11, 0xd1, 0x85, 0x32, 0x0a, 0x8f, 0xac, 0x9a, 0x38, 0x35,
	0xf1, 0x25, 0x03, 0x9a, 0x78, 0x75, 0xfc, 0xc4, 0xe1, 0x40, 0x0b, 0x5a, 0x25, 0x34, 0x55, 0x05,
	0xf7, 0xe8, 0x4d, 0xa9, 0xf6, 0x45, 0xca, 0x1d, 0x27, 0x7e, 0xea, 0x64, 0x25, 0x2f, 0x2a, 0x91,
	0xf2, 0xff, 0xce, 0x8b, 0x87, 0xfb, 0x4c, 0x03, 0x05, 0x29, 0x95, 0x01, 0x23, 0x94, 0x2c, 0x69,
	0x2e, 0xb6, 0x05, 0x98, 0x86, 0xf3, 0xb0, 0x82, 0x9d, 0xc8, 0xc0, 0x70, 0xda, 0x7c, 0xb8, 0xc2,
	0xe4, 0x47, 0x17, 0x3d, 0x58, 0x58, 0xd2, 0xab, 0x83, 0xe1, 0xb2, 0x14, 0x4a, 0x2e, 0xac, 0x0f,
	0xfc, 0x1e, 0x45, 0x90, 0xe5, 0x42, 0x6e, 0x9c, 0xaf, 0x47, 0xc1, 0x38, 0x78, 0x16, 0x26, 0x33,
	0x72, 0x65, 0x33, 0xf2, 0xa2, 0x6e, 0x72, 0x8c, 0xd7, 0x77, 0x56, 0x21, 0x9c, 0x7f, 0xf1, 0x5b,
	0x34, 0x28, 0x6b, 0x7f, 0x69, 0xc3, 0xec, 0x58, 0xe6, 0xd4, 0x33, 0xfd, 0x96, 0x17, 0xb4, 0x0f,
	0xa0, 0xff, 0xb2, 0x22, 0xd7, 0xec, 0x61, 0x5b, 0x14, 0x19, 0xd0, 0x59, 0xd9, 0xb0, 0xba, 0x96,
	0xc5, 0xae, 0xfa, 0xfb, 0xe7, 0xb6, 0xf5, 0x9c, 0x97, 0xeb, 0xb3, 0x6b, 0x4b, 0x76, 0xbf, 0xf1,
	0x17, 0x14, 0xb6, 0xaa, 0xf8, 0x1d, 0x1a, 0x5c, 0xdc, 0x94, 0x0f, 0x66, 0xe4, 0x07, 0x83, 0x16,
	0xa4, 0x4a, 0x48, 0x7d, 0xa3, 0xc4, 0x75, 0xac, 0xad, 0x8c, 0xf5, 0x4e, 0xec, 0xf6, 0x6b, 0xd0,
	0xb9, 0x1f, 0xac, 0xa2, 0xb4, 0x75, 0x8e, 0x1f, 0xa3, 0x1b, 0x09, 0x39, 0xb7, 0x59, 0xf4, 0xd9,
	0xbd, 0x13, 0xbb, 0x29, 0x3a, 0xe3, 0x60, 0x65, 0x0f, 0x19, 0x46, 0xa1, 0x1f, 0x66, 0x8e, 0x9a,
	0xe3, 0xee, 0x2f, 0x16, 0x4c, 0xe6, 0x28, 0x6c, 0x65, 0x8c, 0xa7, 0xa8, 0xef, 0x25, 0x22, 0xb3,
	0x5e, 0x5a, 0x90, 0x9e, 0xab, 0xbc, 0xc9, 0x58, 0xf6, 0xf3, 0xfb, 0xef, 0x6f, 0xb7, 0x63, 0x3c,
	0x74, 0x2e, 0x79, 0xb3, 0x7e, 0x79, 0x11, 0xd1, 0x1c, 0x3d, 0x17, 0xca, 0x2d, 0xa2, 0x0b, 0x75,
	0x38, 0x5e, 0x0b, 0x93, 0x85, 0x2e, 0xcd, 0x65, 0xfd, 0x96, 0x96, 0xc1, 0xc7, 0xbb, 0xf6, 0x51,
	0xcd, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x36, 0xf6, 0x0c, 0x8b, 0x2d, 0x03, 0x00, 0x00,
}
