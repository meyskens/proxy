// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/buffer/v2/buffer.proto

package v2

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type Buffer struct {
	// The maximum request size that the filter will buffer before the connection
	// manager will stop buffering and return a 413 response.
	MaxRequestBytes      *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=max_request_bytes,json=maxRequestBytes,proto3" json:"max_request_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Buffer) Reset()         { *m = Buffer{} }
func (m *Buffer) String() string { return proto.CompactTextString(m) }
func (*Buffer) ProtoMessage()    {}
func (*Buffer) Descriptor() ([]byte, []int) {
	return fileDescriptor_e782fc75ce4c789f, []int{0}
}

func (m *Buffer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Buffer.Unmarshal(m, b)
}
func (m *Buffer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Buffer.Marshal(b, m, deterministic)
}
func (m *Buffer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Buffer.Merge(m, src)
}
func (m *Buffer) XXX_Size() int {
	return xxx_messageInfo_Buffer.Size(m)
}
func (m *Buffer) XXX_DiscardUnknown() {
	xxx_messageInfo_Buffer.DiscardUnknown(m)
}

var xxx_messageInfo_Buffer proto.InternalMessageInfo

func (m *Buffer) GetMaxRequestBytes() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxRequestBytes
	}
	return nil
}

type BufferPerRoute struct {
	// Types that are valid to be assigned to Override:
	//	*BufferPerRoute_Disabled
	//	*BufferPerRoute_Buffer
	Override             isBufferPerRoute_Override `protobuf_oneof:"override"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *BufferPerRoute) Reset()         { *m = BufferPerRoute{} }
func (m *BufferPerRoute) String() string { return proto.CompactTextString(m) }
func (*BufferPerRoute) ProtoMessage()    {}
func (*BufferPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_e782fc75ce4c789f, []int{1}
}

func (m *BufferPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BufferPerRoute.Unmarshal(m, b)
}
func (m *BufferPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BufferPerRoute.Marshal(b, m, deterministic)
}
func (m *BufferPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BufferPerRoute.Merge(m, src)
}
func (m *BufferPerRoute) XXX_Size() int {
	return xxx_messageInfo_BufferPerRoute.Size(m)
}
func (m *BufferPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_BufferPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_BufferPerRoute proto.InternalMessageInfo

type isBufferPerRoute_Override interface {
	isBufferPerRoute_Override()
}

type BufferPerRoute_Disabled struct {
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3,oneof"`
}

type BufferPerRoute_Buffer struct {
	Buffer *Buffer `protobuf:"bytes,2,opt,name=buffer,proto3,oneof"`
}

func (*BufferPerRoute_Disabled) isBufferPerRoute_Override() {}

func (*BufferPerRoute_Buffer) isBufferPerRoute_Override() {}

func (m *BufferPerRoute) GetOverride() isBufferPerRoute_Override {
	if m != nil {
		return m.Override
	}
	return nil
}

func (m *BufferPerRoute) GetDisabled() bool {
	if x, ok := m.GetOverride().(*BufferPerRoute_Disabled); ok {
		return x.Disabled
	}
	return false
}

func (m *BufferPerRoute) GetBuffer() *Buffer {
	if x, ok := m.GetOverride().(*BufferPerRoute_Buffer); ok {
		return x.Buffer
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BufferPerRoute) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BufferPerRoute_Disabled)(nil),
		(*BufferPerRoute_Buffer)(nil),
	}
}

func init() {
	proto.RegisterType((*Buffer)(nil), "envoy.config.filter.http.buffer.v2.Buffer")
	proto.RegisterType((*BufferPerRoute)(nil), "envoy.config.filter.http.buffer.v2.BufferPerRoute")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/buffer/v2/buffer.proto", fileDescriptor_e782fc75ce4c789f)
}

var fileDescriptor_e782fc75ce4c789f = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4f, 0x4a, 0xc3, 0x40,
	0x14, 0xc6, 0x3b, 0xe9, 0x1f, 0xe2, 0x14, 0xb4, 0x0d, 0x82, 0xa5, 0x88, 0x94, 0x6e, 0x94, 0x2e,
	0x66, 0x24, 0xbd, 0xc1, 0x80, 0xa0, 0xae, 0x4a, 0x44, 0x17, 0x6e, 0xca, 0xa4, 0x79, 0x89, 0x91,
	0xb4, 0x13, 0x27, 0x93, 0xd8, 0x5c, 0xc1, 0x43, 0x78, 0x0e, 0x71, 0xd5, 0xeb, 0xf4, 0x16, 0x92,
	0x99, 0xa9, 0x5b, 0xdd, 0x7d, 0xe4, 0xe5, 0xfb, 0xfd, 0xe6, 0x3d, 0x4c, 0x61, 0x53, 0x89, 0x9a,
	0xae, 0xc4, 0x26, 0x4e, 0x13, 0x1a, 0xa7, 0x99, 0x02, 0x49, 0x5f, 0x94, 0xca, 0x69, 0x58, 0xc6,
	0x31, 0x48, 0x5a, 0xf9, 0x36, 0x91, 0x5c, 0x0a, 0x25, 0xbc, 0xa9, 0x2e, 0x10, 0x53, 0x20, 0xa6,
	0x40, 0x9a, 0x02, 0xb1, 0xbf, 0x55, 0xfe, 0xf8, 0x22, 0x11, 0x22, 0xc9, 0x80, 0xea, 0x46, 0x58,
	0xc6, 0xf4, 0x5d, 0xf2, 0x3c, 0x07, 0x59, 0x18, 0xc6, 0xf8, 0xac, 0xe2, 0x59, 0x1a, 0x71, 0x05,
	0xf4, 0x10, 0xec, 0xe0, 0x34, 0x11, 0x89, 0xd0, 0x91, 0x36, 0xc9, 0x7c, 0x9d, 0xae, 0x70, 0x8f,
	0x69, 0xb6, 0xf7, 0x80, 0x87, 0x6b, 0xbe, 0x5d, 0x4a, 0x78, 0x2b, 0xa1, 0x50, 0xcb, 0xb0, 0x56,
	0x50, 0x8c, 0xd0, 0x04, 0x5d, 0xf5, 0xfd, 0x73, 0x62, 0xa4, 0xe4, 0x20, 0x25, 0x8f, 0x77, 0x1b,
	0x35, 0xf7, 0x9f, 0x78, 0x56, 0x02, 0x3b, 0xfa, 0xde, 0xef, 0xda, 0x9d, 0x99, 0x33, 0x69, 0x05,
	0x27, 0x6b, 0xbe, 0x0d, 0x0c, 0x80, 0x35, 0xfd, 0xfb, 0x8e, 0xeb, 0x0c, 0xda, 0xd3, 0x4f, 0x84,
	0x8f, 0x8d, 0x65, 0x01, 0x32, 0x10, 0xa5, 0x02, 0xef, 0x12, 0xbb, 0x51, 0x5a, 0xf0, 0x30, 0x83,
	0x48, 0x4b, 0x5c, 0x8b, 0x79, 0x75, 0x5c, 0x74, 0xdb, 0x0a, 0x7e, 0x87, 0xde, 0x02, 0xf7, 0xcc,
	0xf2, 0x23, 0x47, 0xbf, 0x65, 0x46, 0xfe, 0x3e, 0x12, 0x31, 0x32, 0x86, 0x1b, 0x64, 0xf7, 0x03,
	0x39, 0x83, 0x86, 0x69, 0x39, 0x6c, 0x88, 0x5d, 0x51, 0x81, 0x94, 0x69, 0x04, 0x5e, 0xf7, 0x6b,
	0xbf, 0x6b, 0x23, 0x76, 0x83, 0xaf, 0x53, 0x61, 0xc0, 0xb9, 0x14, 0xdb, 0xfa, 0x1f, 0x0e, 0xd6,
	0xb7, 0x1b, 0x35, 0x27, 0x59, 0xa0, 0x67, 0xa7, 0xf2, 0xc3, 0x9e, 0xbe, 0xcf, 0xfc, 0x27, 0x00,
	0x00, 0xff, 0xff, 0xb5, 0xdc, 0xcb, 0x56, 0xf9, 0x01, 0x00, 0x00,
}
