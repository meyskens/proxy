// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/grpc_http1_reverse_bridge/v2alpha1/config.proto

package v2

import (
	fmt "fmt"
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

// gRPC reverse bridge filter configuration
type FilterConfig struct {
	// The content-type to pass to the upstream when the gRPC bridge filter is applied.
	// The filter will also validate that the upstream responds with the same content type.
	ContentType string `protobuf:"bytes,1,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// If true, Envoy will assume that the upstream doesn't understand gRPC frames and
	// strip the gRPC frame from the request, and add it back in to the response. This will
	// hide the gRPC semantics from the upstream, allowing it to receive and respond with a
	// simple binary encoded protobuf.
	WithholdGrpcFrames   bool     `protobuf:"varint,2,opt,name=withhold_grpc_frames,json=withholdGrpcFrames,proto3" json:"withhold_grpc_frames,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterConfig) Reset()         { *m = FilterConfig{} }
func (m *FilterConfig) String() string { return proto.CompactTextString(m) }
func (*FilterConfig) ProtoMessage()    {}
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a86db517160ad0a, []int{0}
}

func (m *FilterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterConfig.Unmarshal(m, b)
}
func (m *FilterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterConfig.Marshal(b, m, deterministic)
}
func (m *FilterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterConfig.Merge(m, src)
}
func (m *FilterConfig) XXX_Size() int {
	return xxx_messageInfo_FilterConfig.Size(m)
}
func (m *FilterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FilterConfig proto.InternalMessageInfo

func (m *FilterConfig) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *FilterConfig) GetWithholdGrpcFrames() bool {
	if m != nil {
		return m.WithholdGrpcFrames
	}
	return false
}

func init() {
	proto.RegisterType((*FilterConfig)(nil), "envoy.config.filter.http.grpc_http1_reverse_bridge.v2alpha1.FilterConfig")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/grpc_http1_reverse_bridge/v2alpha1/config.proto", fileDescriptor_2a86db517160ad0a)
}

var fileDescriptor_2a86db517160ad0a = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xe5, 0x80, 0x10, 0x75, 0x3b, 0x45, 0x48, 0x54, 0x4c, 0x11, 0x53, 0x07, 0x64, 0xd3,
	0x32, 0xb2, 0x05, 0xa9, 0xc0, 0x56, 0x55, 0x4c, 0x0c, 0x58, 0x6e, 0x72, 0x49, 0x2c, 0x05, 0xdb,
	0xba, 0x9e, 0x0c, 0x79, 0x35, 0x26, 0x5e, 0x87, 0xb7, 0x40, 0x71, 0x92, 0x91, 0x89, 0xed, 0x74,
	0xbf, 0xfe, 0x4f, 0xff, 0xc7, 0x9f, 0xc0, 0x06, 0xd7, 0xc9, 0xc2, 0xd9, 0xca, 0xd4, 0xb2, 0x32,
	0x2d, 0x01, 0xca, 0x86, 0xc8, 0xcb, 0x1a, 0x7d, 0xa1, 0xfa, 0x6b, 0xad, 0x10, 0x02, 0xe0, 0x11,
	0xd4, 0x01, 0x4d, 0x59, 0x83, 0x0c, 0x1b, 0xdd, 0xfa, 0x46, 0xaf, 0xc7, 0x96, 0xf0, 0xe8, 0xc8,
	0xa5, 0xf7, 0x91, 0x24, 0xc6, 0xdf, 0x40, 0x12, 0x7d, 0x5f, 0xfc, 0x49, 0x12, 0x13, 0xe9, 0xea,
	0x32, 0xe8, 0xd6, 0x94, 0x9a, 0x40, 0x4e, 0xc7, 0x40, 0xbd, 0xb6, 0x7c, 0xb1, 0x8d, 0xa8, 0x87,
	0xc8, 0x4d, 0x6f, 0xf8, 0xa2, 0x70, 0x96, 0xc0, 0x92, 0xa2, 0xce, 0xc3, 0x92, 0x65, 0x6c, 0x35,
	0xcb, 0x67, 0x5f, 0x3f, 0xdf, 0x27, 0xa7, 0x98, 0x64, 0x6c, 0x3f, 0x1f, 0xe3, 0x97, 0xce, 0x43,
	0x7a, 0xcb, 0x2f, 0x3e, 0x0c, 0x35, 0x8d, 0x6b, 0x4b, 0x15, 0x57, 0x54, 0xa8, 0xdf, 0xe1, 0xb8,
	0x4c, 0x32, 0xb6, 0x3a, 0xdf, 0xa7, 0x53, 0xf6, 0x88, 0xbe, 0xd8, 0xc6, 0x24, 0x7f, 0xe3, 0xcf,
	0xc6, 0x89, 0xa8, 0xe2, 0xd1, 0x7d, 0x76, 0xe2, 0x1f, 0x56, 0xf9, 0x7c, 0x18, 0xbd, 0xeb, 0x4d,
	0x76, 0xec, 0x35, 0x09, 0x9b, 0xc3, 0x59, 0xd4, 0xba, 0xfb, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xff,
	0x54, 0x52, 0xf9, 0x78, 0x01, 0x00, 0x00,
}
