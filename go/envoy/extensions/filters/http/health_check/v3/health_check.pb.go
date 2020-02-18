// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/health_check/v3/health_check.proto

package envoy_extensions_filters_http_health_check_v3

import (
	fmt "fmt"
	v31 "github.com/cilium/proxy/go/envoy/config/route/v3"
	v3 "github.com/cilium/proxy/go/envoy/type/v3"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// [#next-free-field: 6]
type HealthCheck struct {
	// Specifies whether the filter operates in pass through mode or not.
	PassThroughMode *wrappers.BoolValue `protobuf:"bytes,1,opt,name=pass_through_mode,json=passThroughMode,proto3" json:"pass_through_mode,omitempty"`
	// If operating in pass through mode, the amount of time in milliseconds
	// that the filter should cache the upstream response.
	CacheTime *duration.Duration `protobuf:"bytes,3,opt,name=cache_time,json=cacheTime,proto3" json:"cache_time,omitempty"`
	// If operating in non-pass-through mode, specifies a set of upstream cluster
	// names and the minimum percentage of servers in each of those clusters that
	// must be healthy or degraded in order for the filter to return a 200.
	ClusterMinHealthyPercentages map[string]*v3.Percent `protobuf:"bytes,4,rep,name=cluster_min_healthy_percentages,json=clusterMinHealthyPercentages,proto3" json:"cluster_min_healthy_percentages,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Specifies a set of health check request headers to match on. The health check filter will
	// check a request’s headers against all the specified headers. To specify the health check
	// endpoint, set the ``:path`` header to match on.
	Headers              []*v31.HeaderMatcher `protobuf:"bytes,5,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *HealthCheck) Reset()         { *m = HealthCheck{} }
func (m *HealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()    {}
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_554d4249f82380e7, []int{0}
}

func (m *HealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck.Merge(m, src)
}
func (m *HealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck.Size(m)
}
func (m *HealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck proto.InternalMessageInfo

func (m *HealthCheck) GetPassThroughMode() *wrappers.BoolValue {
	if m != nil {
		return m.PassThroughMode
	}
	return nil
}

func (m *HealthCheck) GetCacheTime() *duration.Duration {
	if m != nil {
		return m.CacheTime
	}
	return nil
}

func (m *HealthCheck) GetClusterMinHealthyPercentages() map[string]*v3.Percent {
	if m != nil {
		return m.ClusterMinHealthyPercentages
	}
	return nil
}

func (m *HealthCheck) GetHeaders() []*v31.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

func init() {
	proto.RegisterType((*HealthCheck)(nil), "envoy.extensions.filters.http.health_check.v3.HealthCheck")
	proto.RegisterMapType((map[string]*v3.Percent)(nil), "envoy.extensions.filters.http.health_check.v3.HealthCheck.ClusterMinHealthyPercentagesEntry")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/health_check/v3/health_check.proto", fileDescriptor_554d4249f82380e7)
}

var fileDescriptor_554d4249f82380e7 = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x5d, 0x6b, 0x13, 0x41,
	0x14, 0x65, 0x93, 0x46, 0xdb, 0xc9, 0x83, 0x71, 0x1f, 0x34, 0x46, 0xa9, 0xad, 0xf8, 0xd0, 0x87,
	0x3a, 0x0b, 0x89, 0x48, 0x69, 0x41, 0x24, 0x55, 0x10, 0x21, 0x10, 0x97, 0xe2, 0x93, 0xb0, 0x4c,
	0x67, 0x6f, 0x76, 0x87, 0xee, 0xce, 0x0c, 0x33, 0xb3, 0x6b, 0xf7, 0x1f, 0x88, 0x3f, 0xc1, 0x57,
	0x1f, 0xfd, 0x1d, 0xfe, 0x29, 0x9f, 0x64, 0x3e, 0x4a, 0x13, 0xfc, 0xa2, 0x6f, 0xb3, 0x7b, 0xcf,
	0x3d, 0xe7, 0xdc, 0x73, 0x2f, 0x7a, 0x05, 0xbc, 0x15, 0x5d, 0x02, 0x97, 0x06, 0xb8, 0x66, 0x82,
	0xeb, 0x64, 0xc5, 0x2a, 0x03, 0x4a, 0x27, 0xa5, 0x31, 0x32, 0x29, 0x81, 0x54, 0xa6, 0xcc, 0x68,
	0x09, 0xf4, 0x22, 0x69, 0x67, 0x1b, 0xdf, 0x58, 0x2a, 0x61, 0x44, 0xfc, 0xcc, 0x31, 0xe0, 0x6b,
	0x06, 0x1c, 0x18, 0xb0, 0x65, 0xc0, 0x1b, 0x1d, 0xed, 0x6c, 0x72, 0xe8, 0x05, 0xa9, 0xe0, 0x2b,
	0x56, 0x24, 0x4a, 0x34, 0x06, 0x2c, 0xb1, 0x7b, 0x64, 0x54, 0xd4, 0x52, 0x70, 0xe0, 0x46, 0x7b,
	0xf2, 0xc9, 0x43, 0x8f, 0x36, 0x9d, 0x74, 0x28, 0x09, 0x8a, 0x02, 0x37, 0xa1, 0xb8, 0x5b, 0x08,
	0x51, 0x54, 0x90, 0xb8, 0xaf, 0xf3, 0x66, 0x95, 0xe4, 0x8d, 0x22, 0x86, 0x09, 0xfe, 0xb7, 0xfa,
	0x27, 0x45, 0xa4, 0xb4, 0xce, 0x7c, 0x7d, 0xbf, 0xc9, 0x25, 0x49, 0x08, 0xe7, 0xc2, 0xb8, 0x36,
	0x9d, 0xb4, 0xa0, 0xec, 0x08, 0x8c, 0x17, 0x01, 0x72, 0xbf, 0x25, 0x15, 0xcb, 0x89, 0xf5, 0x18,
	0x1e, 0xbe, 0xf0, 0xe4, 0xfb, 0x16, 0x1a, 0xbe, 0x75, 0xa3, 0x9d, 0xda, 0xc9, 0xe2, 0x25, 0xba,
	0x2b, 0x89, 0xd6, 0x99, 0x29, 0x95, 0x68, 0x8a, 0x32, 0xab, 0x45, 0x0e, 0xe3, 0x68, 0x2f, 0x3a,
	0x18, 0x4e, 0x27, 0xd8, 0xfb, 0xc0, 0x57, 0x3e, 0xf0, 0x5c, 0x88, 0xea, 0x03, 0xa9, 0x1a, 0x98,
	0x6f, 0xff, 0x9c, 0x0f, 0xbe, 0x44, 0xbd, 0x51, 0x94, 0xde, 0xb1, 0xed, 0x67, 0xbe, 0x7b, 0x21,
	0x72, 0x88, 0x8f, 0x10, 0xa2, 0x84, 0x96, 0x90, 0x19, 0x56, 0xc3, 0xb8, 0xef, 0xa8, 0x1e, 0xfc,
	0x46, 0xf5, 0x3a, 0x8c, 0x9c, 0xee, 0x38, 0xf0, 0x19, 0xab, 0x21, 0xfe, 0x16, 0xa1, 0xc7, 0xb4,
	0x6a, 0xb4, 0x01, 0x95, 0xd5, 0x8c, 0x67, 0x7e, 0x05, 0x5d, 0x16, 0xd2, 0x23, 0x05, 0xe8, 0xf1,
	0xd6, 0x5e, 0xff, 0x60, 0x38, 0xfd, 0x88, 0x6f, 0xb4, 0x3c, 0xbc, 0x36, 0x31, 0x3e, 0xf5, 0x0a,
	0x0b, 0xc6, 0xfd, 0xdf, 0x6e, 0x79, 0x4d, 0xff, 0x86, 0x1b, 0xd5, 0xa5, 0x8f, 0xe8, 0x3f, 0x20,
	0xf1, 0x4b, 0x74, 0xbb, 0x04, 0x92, 0x83, 0xd2, 0xe3, 0x81, 0x33, 0xf3, 0x34, 0x98, 0xf1, 0xa7,
	0x81, 0xdd, 0x45, 0x04, 0xd1, 0x1c, 0xd4, 0x82, 0x18, 0x5a, 0x82, 0x4a, 0xaf, 0x9a, 0x26, 0x05,
	0xda, 0xff, 0xaf, 0x85, 0x78, 0x84, 0xfa, 0x17, 0xd0, 0xb9, 0x45, 0xec, 0xa4, 0xf6, 0x19, 0x1f,
	0xa2, 0x41, 0x6b, 0xa3, 0x1f, 0xf7, 0x5c, 0xa2, 0xf7, 0x82, 0xa8, 0xbd, 0x30, 0x2b, 0x16, 0x18,
	0x52, 0x0f, 0x3a, 0xee, 0x1d, 0x45, 0xc7, 0x27, 0x5f, 0x7f, 0x7c, 0xde, 0x7d, 0x81, 0x9e, 0x6f,
	0xb8, 0xf3, 0x31, 0xfd, 0x29, 0xa5, 0xe9, 0x7a, 0x4a, 0xef, 0xb6, 0xb6, 0x7b, 0xa3, 0xfe, 0xfc,
	0x3d, 0x3a, 0x61, 0xc2, 0x2b, 0x49, 0x25, 0x2e, 0xbb, 0x9b, 0xc5, 0x3e, 0x1f, 0xad, 0x31, 0x2e,
	0xed, 0xe6, 0x97, 0xd1, 0xf9, 0x2d, 0x77, 0x02, 0xb3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb6,
	0xaf, 0x1e, 0x44, 0xc1, 0x03, 0x00, 0x00,
}
