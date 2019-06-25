// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/ext_authz/v2/ext_authz.proto

package v2

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	matcher "github.com/cilium/proxy/go/envoy/type/matcher"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/lyft/protoc-gen-validate/validate"
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

type ExtAuthz struct {
	// External authorization service configuration.
	//
	// Types that are valid to be assigned to Services:
	//	*ExtAuthz_GrpcService
	//	*ExtAuthz_HttpService
	Services isExtAuthz_Services `protobuf_oneof:"services"`
	//  Changes filter's behaviour on errors:
	//
	//  1. When set to true, the filter will *accept* client request even if the communication with
	//  the authorization service has failed, or if the authorization service has returned a HTTP 5xx
	//  error.
	//
	//  2. When set to false, ext-authz will *reject* client requests and return a *Forbidden*
	//  response if the communication with the authorization service has failed, or if the
	//  authorization service has returned a HTTP 5xx error.
	//
	// Note that errors can be *always* tracked in the :ref:`stats
	// <config_http_filters_ext_authz_stats>`.
	FailureModeAllow bool `protobuf:"varint,2,opt,name=failure_mode_allow,json=failureModeAllow,proto3" json:"failure_mode_allow,omitempty"`
	// Sets the package version the gRPC service should use. This is particularly
	// useful when transitioning from alpha to release versions assuming that both definitions are
	// semantically compatible. Deprecation note: This field is deprecated and should only be used for
	// version upgrade. See release notes for more details.
	UseAlpha bool `protobuf:"varint,4,opt,name=use_alpha,json=useAlpha,proto3" json:"use_alpha,omitempty"` // Deprecated: Do not use.
	// Enables filter to buffer the client request body and send it within the authorization request.
	// A ``x-envoy-auth-partial-body: false|true`` metadata header will be added to the authorization
	// request message indicating if the body data is partial.
	WithRequestBody *BufferSettings `protobuf:"bytes,5,opt,name=with_request_body,json=withRequestBody,proto3" json:"with_request_body,omitempty"`
	// Clears route cache in order to allow the external authorization service to correctly affect
	// routing decisions. Filter clears all cached routes when:
	//
	// 1. The field is set to *true*.
	//
	// 2. The status returned from the authorization service is a HTTP 200 or gRPC 0.
	//
	// 3. At least one *authorization response header* is added to the client request, or is used for
	// altering another client request header.
	//
	ClearRouteCache      bool     `protobuf:"varint,6,opt,name=clear_route_cache,json=clearRouteCache,proto3" json:"clear_route_cache,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtAuthz) Reset()         { *m = ExtAuthz{} }
func (m *ExtAuthz) String() string { return proto.CompactTextString(m) }
func (*ExtAuthz) ProtoMessage()    {}
func (*ExtAuthz) Descriptor() ([]byte, []int) {
	return fileDescriptor_861cd76675296973, []int{0}
}

func (m *ExtAuthz) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtAuthz.Unmarshal(m, b)
}
func (m *ExtAuthz) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtAuthz.Marshal(b, m, deterministic)
}
func (m *ExtAuthz) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtAuthz.Merge(m, src)
}
func (m *ExtAuthz) XXX_Size() int {
	return xxx_messageInfo_ExtAuthz.Size(m)
}
func (m *ExtAuthz) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtAuthz.DiscardUnknown(m)
}

var xxx_messageInfo_ExtAuthz proto.InternalMessageInfo

type isExtAuthz_Services interface {
	isExtAuthz_Services()
}

type ExtAuthz_GrpcService struct {
	GrpcService *core.GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,proto3,oneof"`
}

type ExtAuthz_HttpService struct {
	HttpService *HttpService `protobuf:"bytes,3,opt,name=http_service,json=httpService,proto3,oneof"`
}

func (*ExtAuthz_GrpcService) isExtAuthz_Services() {}

func (*ExtAuthz_HttpService) isExtAuthz_Services() {}

func (m *ExtAuthz) GetServices() isExtAuthz_Services {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *ExtAuthz) GetGrpcService() *core.GrpcService {
	if x, ok := m.GetServices().(*ExtAuthz_GrpcService); ok {
		return x.GrpcService
	}
	return nil
}

func (m *ExtAuthz) GetHttpService() *HttpService {
	if x, ok := m.GetServices().(*ExtAuthz_HttpService); ok {
		return x.HttpService
	}
	return nil
}

func (m *ExtAuthz) GetFailureModeAllow() bool {
	if m != nil {
		return m.FailureModeAllow
	}
	return false
}

// Deprecated: Do not use.
func (m *ExtAuthz) GetUseAlpha() bool {
	if m != nil {
		return m.UseAlpha
	}
	return false
}

func (m *ExtAuthz) GetWithRequestBody() *BufferSettings {
	if m != nil {
		return m.WithRequestBody
	}
	return nil
}

func (m *ExtAuthz) GetClearRouteCache() bool {
	if m != nil {
		return m.ClearRouteCache
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExtAuthz) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExtAuthz_GrpcService)(nil),
		(*ExtAuthz_HttpService)(nil),
	}
}

// Configuration for buffering the request data.
type BufferSettings struct {
	// Sets the maximum size of a message body that the filter will hold in memory. Envoy will return
	// *HTTP 413* and will *not* initiate the authorization process when buffer reaches the number
	// set in this field. Note that this setting will have precedence over :ref:`failure_mode_allow
	// <envoy_api_field_config.filter.http.ext_authz.v2.ExtAuthz.failure_mode_allow>`.
	MaxRequestBytes uint32 `protobuf:"varint,1,opt,name=max_request_bytes,json=maxRequestBytes,proto3" json:"max_request_bytes,omitempty"`
	// When this field is true, Envoy will buffer the message until *max_request_bytes* is reached.
	// The authorization request will be dispatched and no 413 HTTP error will be returned by the
	// filter.
	AllowPartialMessage  bool     `protobuf:"varint,2,opt,name=allow_partial_message,json=allowPartialMessage,proto3" json:"allow_partial_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BufferSettings) Reset()         { *m = BufferSettings{} }
func (m *BufferSettings) String() string { return proto.CompactTextString(m) }
func (*BufferSettings) ProtoMessage()    {}
func (*BufferSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_861cd76675296973, []int{1}
}

func (m *BufferSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BufferSettings.Unmarshal(m, b)
}
func (m *BufferSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BufferSettings.Marshal(b, m, deterministic)
}
func (m *BufferSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BufferSettings.Merge(m, src)
}
func (m *BufferSettings) XXX_Size() int {
	return xxx_messageInfo_BufferSettings.Size(m)
}
func (m *BufferSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_BufferSettings.DiscardUnknown(m)
}

var xxx_messageInfo_BufferSettings proto.InternalMessageInfo

func (m *BufferSettings) GetMaxRequestBytes() uint32 {
	if m != nil {
		return m.MaxRequestBytes
	}
	return 0
}

func (m *BufferSettings) GetAllowPartialMessage() bool {
	if m != nil {
		return m.AllowPartialMessage
	}
	return false
}

// HttpService is used for raw HTTP communication between the filter and the authorization service.
// When configured, the filter will parse the client request and use these attributes to call the
// authorization server. Depending on the response, the filter may reject or accept the client
// request. Note that in any of these events, metadata can be added, removed or overridden by the
// filter:
//
// *On authorization request*, a list of allowed request headers may be supplied. See
// :ref:`allowed_headers
// <envoy_api_field_config.filter.http.ext_authz.v2.AuthorizationRequest.allowed_headers>`
// for details. Additional headers metadata may be added to the authorization request. See
// :ref:`headers_to_add
// <envoy_api_field_config.filter.http.ext_authz.v2.AuthorizationRequest.headers_to_add>` for
// details.
//
// On authorization response status HTTP 200 OK, the filter will allow traffic to the upstream and
// additional headers metadata may be added to the original client request. See
// :ref:`allowed_upstream_headers
// <envoy_api_field_config.filter.http.ext_authz.v2.AuthorizationResponse.allowed_upstream_headers>`
// for details.
//
// On other authorization response statuses, the filter will not allow traffic. Additional headers
// metadata as well as body may be added to the client's response. See :ref:`allowed_client_headers
// <envoy_api_field_config.filter.http.ext_authz.v2.AuthorizationResponse.allowed_client_headers>`
// for details.
type HttpService struct {
	// Sets the HTTP server URI which the authorization requests must be sent to.
	ServerUri *core.HttpUri `protobuf:"bytes,1,opt,name=server_uri,json=serverUri,proto3" json:"server_uri,omitempty"`
	// Sets a prefix to the value of authorization request header *Path*.
	PathPrefix string `protobuf:"bytes,2,opt,name=path_prefix,json=pathPrefix,proto3" json:"path_prefix,omitempty"`
	// Settings used for controlling authorization request metadata.
	AuthorizationRequest *AuthorizationRequest `protobuf:"bytes,7,opt,name=authorization_request,json=authorizationRequest,proto3" json:"authorization_request,omitempty"`
	// Settings used for controlling authorization response metadata.
	AuthorizationResponse *AuthorizationResponse `protobuf:"bytes,8,opt,name=authorization_response,json=authorizationResponse,proto3" json:"authorization_response,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *HttpService) Reset()         { *m = HttpService{} }
func (m *HttpService) String() string { return proto.CompactTextString(m) }
func (*HttpService) ProtoMessage()    {}
func (*HttpService) Descriptor() ([]byte, []int) {
	return fileDescriptor_861cd76675296973, []int{2}
}

func (m *HttpService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpService.Unmarshal(m, b)
}
func (m *HttpService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpService.Marshal(b, m, deterministic)
}
func (m *HttpService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpService.Merge(m, src)
}
func (m *HttpService) XXX_Size() int {
	return xxx_messageInfo_HttpService.Size(m)
}
func (m *HttpService) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpService.DiscardUnknown(m)
}

var xxx_messageInfo_HttpService proto.InternalMessageInfo

func (m *HttpService) GetServerUri() *core.HttpUri {
	if m != nil {
		return m.ServerUri
	}
	return nil
}

func (m *HttpService) GetPathPrefix() string {
	if m != nil {
		return m.PathPrefix
	}
	return ""
}

func (m *HttpService) GetAuthorizationRequest() *AuthorizationRequest {
	if m != nil {
		return m.AuthorizationRequest
	}
	return nil
}

func (m *HttpService) GetAuthorizationResponse() *AuthorizationResponse {
	if m != nil {
		return m.AuthorizationResponse
	}
	return nil
}

type AuthorizationRequest struct {
	// Authorization request will include the client request headers that have a correspondent match
	// in the :ref:`list <envoy_api_msg_type.matcher.ListStringMatcher>`. Note that in addition to the
	// user's supplied matchers:
	//
	// 1. *Host*, *Method*, *Path* and *Content-Length* are automatically included to the list.
	//
	// 2. *Content-Length* will be set to 0 and the request to the authorization service will not have
	// a message body.
	//
	AllowedHeaders *matcher.ListStringMatcher `protobuf:"bytes,1,opt,name=allowed_headers,json=allowedHeaders,proto3" json:"allowed_headers,omitempty"`
	// Sets a list of headers that will be included to the request to authorization service. Note that
	// client request of the same key will be overridden.
	HeadersToAdd         []*core.HeaderValue `protobuf:"bytes,2,rep,name=headers_to_add,json=headersToAdd,proto3" json:"headers_to_add,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *AuthorizationRequest) Reset()         { *m = AuthorizationRequest{} }
func (m *AuthorizationRequest) String() string { return proto.CompactTextString(m) }
func (*AuthorizationRequest) ProtoMessage()    {}
func (*AuthorizationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_861cd76675296973, []int{3}
}

func (m *AuthorizationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizationRequest.Unmarshal(m, b)
}
func (m *AuthorizationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizationRequest.Marshal(b, m, deterministic)
}
func (m *AuthorizationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizationRequest.Merge(m, src)
}
func (m *AuthorizationRequest) XXX_Size() int {
	return xxx_messageInfo_AuthorizationRequest.Size(m)
}
func (m *AuthorizationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizationRequest proto.InternalMessageInfo

func (m *AuthorizationRequest) GetAllowedHeaders() *matcher.ListStringMatcher {
	if m != nil {
		return m.AllowedHeaders
	}
	return nil
}

func (m *AuthorizationRequest) GetHeadersToAdd() []*core.HeaderValue {
	if m != nil {
		return m.HeadersToAdd
	}
	return nil
}

type AuthorizationResponse struct {
	// When this :ref:`list <envoy_api_msg_type.matcher.ListStringMatcher>` is set, authorization
	// response headers that have a correspondent match will be added to the original client request.
	// Note that coexistent headers will be overridden.
	AllowedUpstreamHeaders *matcher.ListStringMatcher `protobuf:"bytes,1,opt,name=allowed_upstream_headers,json=allowedUpstreamHeaders,proto3" json:"allowed_upstream_headers,omitempty"`
	// When this :ref:`list <envoy_api_msg_type.matcher.ListStringMatcher>`. is set, authorization
	// response headers that have a correspondent match will be added to the client's response. Note
	// that when this list is *not* set, all the authorization response headers, except *Authority
	// (Host)* will be in the response to the client. When a header is included in this list, *Path*,
	// *Status*, *Content-Length*, *WWWAuthenticate* and *Location* are automatically added.
	AllowedClientHeaders *matcher.ListStringMatcher `protobuf:"bytes,2,opt,name=allowed_client_headers,json=allowedClientHeaders,proto3" json:"allowed_client_headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *AuthorizationResponse) Reset()         { *m = AuthorizationResponse{} }
func (m *AuthorizationResponse) String() string { return proto.CompactTextString(m) }
func (*AuthorizationResponse) ProtoMessage()    {}
func (*AuthorizationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_861cd76675296973, []int{4}
}

func (m *AuthorizationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizationResponse.Unmarshal(m, b)
}
func (m *AuthorizationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizationResponse.Marshal(b, m, deterministic)
}
func (m *AuthorizationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizationResponse.Merge(m, src)
}
func (m *AuthorizationResponse) XXX_Size() int {
	return xxx_messageInfo_AuthorizationResponse.Size(m)
}
func (m *AuthorizationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizationResponse proto.InternalMessageInfo

func (m *AuthorizationResponse) GetAllowedUpstreamHeaders() *matcher.ListStringMatcher {
	if m != nil {
		return m.AllowedUpstreamHeaders
	}
	return nil
}

func (m *AuthorizationResponse) GetAllowedClientHeaders() *matcher.ListStringMatcher {
	if m != nil {
		return m.AllowedClientHeaders
	}
	return nil
}

// Extra settings on a per virtualhost/route/weighted-cluster level.
type ExtAuthzPerRoute struct {
	// Types that are valid to be assigned to Override:
	//	*ExtAuthzPerRoute_Disabled
	//	*ExtAuthzPerRoute_CheckSettings
	Override             isExtAuthzPerRoute_Override `protobuf_oneof:"override"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ExtAuthzPerRoute) Reset()         { *m = ExtAuthzPerRoute{} }
func (m *ExtAuthzPerRoute) String() string { return proto.CompactTextString(m) }
func (*ExtAuthzPerRoute) ProtoMessage()    {}
func (*ExtAuthzPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_861cd76675296973, []int{5}
}

func (m *ExtAuthzPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtAuthzPerRoute.Unmarshal(m, b)
}
func (m *ExtAuthzPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtAuthzPerRoute.Marshal(b, m, deterministic)
}
func (m *ExtAuthzPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtAuthzPerRoute.Merge(m, src)
}
func (m *ExtAuthzPerRoute) XXX_Size() int {
	return xxx_messageInfo_ExtAuthzPerRoute.Size(m)
}
func (m *ExtAuthzPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtAuthzPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_ExtAuthzPerRoute proto.InternalMessageInfo

type isExtAuthzPerRoute_Override interface {
	isExtAuthzPerRoute_Override()
}

type ExtAuthzPerRoute_Disabled struct {
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3,oneof"`
}

type ExtAuthzPerRoute_CheckSettings struct {
	CheckSettings *CheckSettings `protobuf:"bytes,2,opt,name=check_settings,json=checkSettings,proto3,oneof"`
}

func (*ExtAuthzPerRoute_Disabled) isExtAuthzPerRoute_Override() {}

func (*ExtAuthzPerRoute_CheckSettings) isExtAuthzPerRoute_Override() {}

func (m *ExtAuthzPerRoute) GetOverride() isExtAuthzPerRoute_Override {
	if m != nil {
		return m.Override
	}
	return nil
}

func (m *ExtAuthzPerRoute) GetDisabled() bool {
	if x, ok := m.GetOverride().(*ExtAuthzPerRoute_Disabled); ok {
		return x.Disabled
	}
	return false
}

func (m *ExtAuthzPerRoute) GetCheckSettings() *CheckSettings {
	if x, ok := m.GetOverride().(*ExtAuthzPerRoute_CheckSettings); ok {
		return x.CheckSettings
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExtAuthzPerRoute) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExtAuthzPerRoute_Disabled)(nil),
		(*ExtAuthzPerRoute_CheckSettings)(nil),
	}
}

// Extra settings for the check request. You can use this to provide extra context for the
// external authorization server on specific virtual hosts \ routes. For example, adding a context
// extension on the virtual host level can give the ext-authz server information on what virtual
// host is used without needing to parse the host header. If CheckSettings is specified in multiple
// per-filter-configs, they will be merged in order, and the result will be used.
type CheckSettings struct {
	// Context extensions to set on the CheckRequest's
	// :ref:`AttributeContext.context_extensions<envoy_api_field_service.auth.v2.AttributeContext.context_extensions>`
	//
	// Merge semantics for this field are such that keys from more specific configs override.
	//
	// .. note::
	//
	//   These settings are only applied to a filter configured with a
	//   :ref:`grpc_service<envoy_api_field_config.filter.http.ext_authz.v2.ExtAuthz.grpc_service>`.
	ContextExtensions    map[string]string `protobuf:"bytes,1,rep,name=context_extensions,json=contextExtensions,proto3" json:"context_extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CheckSettings) Reset()         { *m = CheckSettings{} }
func (m *CheckSettings) String() string { return proto.CompactTextString(m) }
func (*CheckSettings) ProtoMessage()    {}
func (*CheckSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_861cd76675296973, []int{6}
}

func (m *CheckSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckSettings.Unmarshal(m, b)
}
func (m *CheckSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckSettings.Marshal(b, m, deterministic)
}
func (m *CheckSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckSettings.Merge(m, src)
}
func (m *CheckSettings) XXX_Size() int {
	return xxx_messageInfo_CheckSettings.Size(m)
}
func (m *CheckSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckSettings.DiscardUnknown(m)
}

var xxx_messageInfo_CheckSettings proto.InternalMessageInfo

func (m *CheckSettings) GetContextExtensions() map[string]string {
	if m != nil {
		return m.ContextExtensions
	}
	return nil
}

func init() {
	proto.RegisterType((*ExtAuthz)(nil), "envoy.config.filter.http.ext_authz.v2.ExtAuthz")
	proto.RegisterType((*BufferSettings)(nil), "envoy.config.filter.http.ext_authz.v2.BufferSettings")
	proto.RegisterType((*HttpService)(nil), "envoy.config.filter.http.ext_authz.v2.HttpService")
	proto.RegisterType((*AuthorizationRequest)(nil), "envoy.config.filter.http.ext_authz.v2.AuthorizationRequest")
	proto.RegisterType((*AuthorizationResponse)(nil), "envoy.config.filter.http.ext_authz.v2.AuthorizationResponse")
	proto.RegisterType((*ExtAuthzPerRoute)(nil), "envoy.config.filter.http.ext_authz.v2.ExtAuthzPerRoute")
	proto.RegisterType((*CheckSettings)(nil), "envoy.config.filter.http.ext_authz.v2.CheckSettings")
	proto.RegisterMapType((map[string]string)(nil), "envoy.config.filter.http.ext_authz.v2.CheckSettings.ContextExtensionsEntry")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/ext_authz/v2/ext_authz.proto", fileDescriptor_861cd76675296973)
}

var fileDescriptor_861cd76675296973 = []byte{
	// 922 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x6e, 0xe3, 0xc4,
	0x17, 0xaf, 0x9d, 0xa4, 0xeb, 0x4e, 0xbf, 0xd2, 0xf9, 0xb7, 0xfd, 0x5b, 0x15, 0xda, 0x56, 0x11,
	0x2b, 0x56, 0x2b, 0x64, 0x4b, 0x59, 0x2a, 0xf1, 0x75, 0xd3, 0x64, 0x2b, 0xaa, 0x42, 0x57, 0x95,
	0x4b, 0x41, 0x82, 0x0b, 0x6b, 0x6a, 0x9f, 0xc4, 0xc3, 0x3a, 0x1e, 0x33, 0x33, 0xce, 0x26, 0xe5,
	0x0d, 0xb8, 0xe4, 0x39, 0x78, 0x00, 0xe0, 0x6a, 0x5f, 0x81, 0x0b, 0x24, 0xae, 0xb9, 0x62, 0xdf,
	0x02, 0xcd, 0x87, 0x9b, 0x76, 0x37, 0x48, 0xa1, 0x77, 0x33, 0xe7, 0xfc, 0xce, 0xef, 0xfc, 0xe6,
	0x7c, 0xd8, 0xe8, 0x10, 0x8a, 0x31, 0x9b, 0x86, 0x09, 0x2b, 0x06, 0x74, 0x18, 0x0e, 0x68, 0x2e,
	0x81, 0x87, 0x99, 0x94, 0x65, 0x08, 0x13, 0x19, 0x93, 0x4a, 0x66, 0xd7, 0xe1, 0xb8, 0x3b, 0xbb,
	0x04, 0x25, 0x67, 0x92, 0xe1, 0x47, 0x3a, 0x2c, 0x30, 0x61, 0x81, 0x09, 0x0b, 0x54, 0x58, 0x30,
	0x43, 0x8e, 0xbb, 0x7b, 0xef, 0x18, 0x76, 0x52, 0x52, 0x45, 0x92, 0x30, 0x0e, 0xe1, 0x15, 0x11,
	0x60, 0x48, 0xf6, 0xde, 0x7d, 0xdb, 0x3b, 0xe4, 0x65, 0x12, 0x0b, 0xe0, 0x63, 0x9a, 0xd4, 0xa8,
	0x83, 0xb7, 0x51, 0x2a, 0x51, 0x5c, 0x71, 0x6a, 0x11, 0xfb, 0x06, 0x21, 0xa7, 0x25, 0x84, 0x23,
	0x22, 0x93, 0x0c, 0x78, 0x28, 0x24, 0xa7, 0xc5, 0xd0, 0x02, 0xfe, 0x3f, 0x26, 0x39, 0x4d, 0x89,
	0x84, 0xb0, 0x3e, 0x58, 0xc7, 0xf6, 0x90, 0x0d, 0x99, 0x3e, 0x86, 0xea, 0x64, 0xac, 0x9d, 0x9f,
	0x1a, 0xc8, 0x3b, 0x9e, 0xc8, 0x23, 0xf5, 0x0a, 0xdc, 0x47, 0x6b, 0xb7, 0x45, 0xf9, 0xce, 0x81,
	0xf3, 0x78, 0xb5, 0xfb, 0x30, 0x30, 0x05, 0x20, 0x25, 0x0d, 0xc6, 0xdd, 0x40, 0xa9, 0x0a, 0x3e,
	0xe3, 0x65, 0x72, 0x61, 0x50, 0x27, 0x4b, 0xd1, 0xea, 0x70, 0x76, 0xc5, 0x5f, 0xa3, 0x35, 0xad,
	0xb9, 0x26, 0x69, 0x68, 0x92, 0x6e, 0xb0, 0x50, 0x15, 0x83, 0x13, 0x29, 0xcb, 0x5b, 0xc4, 0xd9,
	0xec, 0x8a, 0xdf, 0x47, 0x78, 0x40, 0x68, 0x5e, 0x71, 0x88, 0x47, 0x2c, 0x85, 0x98, 0xe4, 0x39,
	0x7b, 0xe9, 0xbb, 0x07, 0xce, 0x63, 0x2f, 0x6a, 0x5b, 0xcf, 0x19, 0x4b, 0xe1, 0x48, 0xd9, 0xf1,
	0x3e, 0x5a, 0xa9, 0x84, 0x02, 0x95, 0x19, 0xf1, 0x9b, 0x0a, 0xd4, 0x73, 0x7d, 0x27, 0xf2, 0x2a,
	0x01, 0x47, 0xca, 0x86, 0x09, 0xda, 0x7a, 0x49, 0x65, 0x16, 0x73, 0xf8, 0xbe, 0x02, 0x21, 0xe3,
	0x2b, 0x96, 0x4e, 0xfd, 0x96, 0x16, 0x7b, 0xb8, 0xa0, 0xd8, 0x5e, 0x35, 0x18, 0x00, 0xbf, 0x00,
	0x29, 0x69, 0x31, 0x14, 0xd1, 0xa6, 0xe2, 0x8b, 0x0c, 0x5d, 0x8f, 0xa5, 0x53, 0xfc, 0x04, 0x6d,
	0x25, 0x39, 0x10, 0x1e, 0x73, 0x56, 0x49, 0x88, 0x13, 0x92, 0x64, 0xe0, 0x2f, 0x6b, 0xc1, 0x9b,
	0xda, 0x11, 0x29, 0x7b, 0x5f, 0x99, 0x7b, 0x08, 0x79, 0xb6, 0x62, 0xa2, 0xf3, 0x03, 0xda, 0xb8,
	0x4b, 0x8d, 0x0f, 0xd1, 0xd6, 0x88, 0x4c, 0x66, 0x5a, 0xa7, 0x12, 0x84, 0x6e, 0xcf, 0x7a, 0x6f,
	0xe5, 0xb7, 0xd7, 0xaf, 0x1a, 0xcd, 0x27, 0xee, 0xc1, 0x52, 0xb4, 0x39, 0x22, 0x93, 0x3a, 0xbf,
	0x42, 0xe0, 0x2e, 0xda, 0xd1, 0x55, 0x8a, 0x4b, 0xc2, 0x25, 0x25, 0x79, 0x3c, 0x02, 0x21, 0xc8,
	0x10, 0x6c, 0xd5, 0xfe, 0xa7, 0x9d, 0xe7, 0xc6, 0x77, 0x66, 0x5c, 0x9d, 0xbf, 0x5d, 0xb4, 0x7a,
	0xab, 0x0b, 0xf8, 0x23, 0x84, 0x94, 0x30, 0xe0, 0x6a, 0x0a, 0xed, 0x48, 0xec, 0xcd, 0x19, 0x09,
	0x15, 0x73, 0xc9, 0x69, 0xb4, 0x62, 0xd0, 0x97, 0x9c, 0xe2, 0x7d, 0xb4, 0x5a, 0x12, 0x99, 0xc5,
	0x25, 0x87, 0x01, 0x9d, 0xe8, 0xa4, 0x2b, 0x11, 0x52, 0xa6, 0x73, 0x6d, 0xc1, 0x25, 0xda, 0x51,
	0xc5, 0x64, 0x9c, 0x5e, 0x13, 0x49, 0x59, 0x51, 0x3f, 0xd0, 0x7f, 0xa0, 0xd3, 0x7c, 0xb2, 0x60,
	0x1f, 0x8e, 0x6e, 0x73, 0xd8, 0x02, 0x44, 0xdb, 0x64, 0x8e, 0x15, 0x0b, 0xb4, 0xfb, 0x66, 0x46,
	0x51, 0xb2, 0x42, 0x80, 0xef, 0xe9, 0x94, 0x9f, 0xde, 0x2f, 0xa5, 0xe1, 0x88, 0x76, 0xc8, 0x3c,
	0xf3, 0x69, 0xd3, 0x6b, 0xb4, 0x9b, 0xa7, 0x4d, 0xaf, 0xd9, 0x6e, 0x9d, 0x36, 0xbd, 0x56, 0x7b,
	0xf9, 0xb4, 0xe9, 0x2d, 0xb7, 0x1f, 0x74, 0x7e, 0x76, 0xd0, 0xf6, 0x3c, 0xed, 0xf8, 0x39, 0xda,
	0xd4, 0xad, 0x81, 0x34, 0xce, 0x80, 0xa4, 0xc0, 0x85, 0x2d, 0xfc, 0x23, 0x2b, 0x4f, 0xed, 0x7f,
	0x60, 0xf7, 0x3f, 0xf8, 0x82, 0x0a, 0x79, 0xa1, 0xbf, 0x01, 0x67, 0xc6, 0x12, 0x6d, 0xd8, 0xe8,
	0x13, 0x13, 0x8c, 0x9f, 0xa1, 0x0d, 0xcb, 0x13, 0x4b, 0x16, 0x93, 0x34, 0xf5, 0xdd, 0x83, 0xc6,
	0xbf, 0xac, 0xb6, 0x89, 0xf9, 0x8a, 0xe4, 0x15, 0x44, 0x6b, 0x36, 0xea, 0x4b, 0x76, 0x94, 0xa6,
	0x9d, 0x3f, 0x1c, 0xb4, 0x33, 0xf7, 0xdd, 0x38, 0x46, 0x7e, 0xad, 0xb7, 0x2a, 0x85, 0xe4, 0x40,
	0x46, 0xf7, 0x13, 0xbe, 0x6b, 0x69, 0x2e, 0x2d, 0x4b, 0xfd, 0x80, 0x6f, 0x51, 0xed, 0x89, 0x93,
	0x9c, 0x42, 0x21, 0x6f, 0xe8, 0xdd, 0xff, 0x42, 0xbf, 0x6d, 0x49, 0xfa, 0x9a, 0xc3, 0x92, 0x77,
	0x7e, 0x75, 0x50, 0xbb, 0xfe, 0x06, 0x9e, 0x83, 0x59, 0x4a, 0xfc, 0x1e, 0xf2, 0x52, 0x2a, 0xc8,
	0x55, 0x0e, 0xa9, 0x7e, 0x82, 0x67, 0x17, 0xed, 0x3b, 0xd7, 0x73, 0x4e, 0x96, 0xa2, 0x1b, 0x27,
	0xa6, 0x68, 0x23, 0xc9, 0x20, 0x79, 0x11, 0x0b, 0xbb, 0xac, 0x56, 0xd2, 0x07, 0x0b, 0x4e, 0x52,
	0x5f, 0x05, 0xd7, 0x8b, 0xde, 0x43, 0x2a, 0x49, 0xeb, 0x47, 0xc7, 0x6d, 0xab, 0x2c, 0xeb, 0xc9,
	0x1d, 0xe7, 0x16, 0xf2, 0xd8, 0x18, 0x38, 0xa7, 0x29, 0xe0, 0xd6, 0x2f, 0xaf, 0x5f, 0x35, 0x9c,
	0xce, 0xef, 0x0e, 0x5a, 0xbf, 0xc3, 0x80, 0xaf, 0x11, 0x4e, 0x58, 0x21, 0x55, 0x0e, 0x98, 0x48,
	0x28, 0x04, 0x65, 0x85, 0xea, 0x82, 0xea, 0xf7, 0xe7, 0xf7, 0xd1, 0x14, 0xf4, 0x0d, 0xdd, 0xf1,
	0x0d, 0xdb, 0x71, 0x21, 0xf9, 0x34, 0xda, 0x4a, 0xde, 0xb4, 0xef, 0x3d, 0x43, 0xbb, 0xf3, 0xc1,
	0xb8, 0x8d, 0x1a, 0x2f, 0x60, 0xaa, 0x2b, 0xb9, 0x12, 0xa9, 0x23, 0xde, 0x46, 0xad, 0xb1, 0x1a,
	0x32, 0xfb, 0x59, 0x30, 0x97, 0x8f, 0xdd, 0x0f, 0x9d, 0xde, 0xf3, 0x3f, 0xff, 0x7a, 0xe8, 0xa0,
	0xa7, 0x94, 0x19, 0xb5, 0x25, 0x67, 0x93, 0xe9, 0x62, 0xc2, 0x7b, 0xeb, 0x37, 0x7d, 0x54, 0x7f,
	0xb7, 0x73, 0xe7, 0x1b, 0x77, 0xdc, 0xbd, 0x5a, 0xd6, 0xbf, 0xba, 0xa7, 0xff, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x72, 0x68, 0x16, 0x9a, 0x00, 0x08, 0x00, 0x00,
}
