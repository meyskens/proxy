// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/data/tap/v2alpha/wrapper.proto

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

// Wrapper for all fully buffered and streamed tap traces that Envoy emits. This is required for
// sending traces over gRPC APIs or more easily persisting binary messages to files.
type TraceWrapper struct {
	// Types that are valid to be assigned to Trace:
	//	*TraceWrapper_HttpBufferedTrace
	//	*TraceWrapper_HttpStreamedTraceSegment
	//	*TraceWrapper_SocketBufferedTrace
	//	*TraceWrapper_SocketStreamedTraceSegment
	Trace                isTraceWrapper_Trace `protobuf_oneof:"trace"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TraceWrapper) Reset()         { *m = TraceWrapper{} }
func (m *TraceWrapper) String() string { return proto.CompactTextString(m) }
func (*TraceWrapper) ProtoMessage()    {}
func (*TraceWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f617738ad092e1c, []int{0}
}

func (m *TraceWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceWrapper.Unmarshal(m, b)
}
func (m *TraceWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceWrapper.Marshal(b, m, deterministic)
}
func (m *TraceWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceWrapper.Merge(m, src)
}
func (m *TraceWrapper) XXX_Size() int {
	return xxx_messageInfo_TraceWrapper.Size(m)
}
func (m *TraceWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_TraceWrapper proto.InternalMessageInfo

type isTraceWrapper_Trace interface {
	isTraceWrapper_Trace()
}

type TraceWrapper_HttpBufferedTrace struct {
	HttpBufferedTrace *HttpBufferedTrace `protobuf:"bytes,1,opt,name=http_buffered_trace,json=httpBufferedTrace,proto3,oneof"`
}

type TraceWrapper_HttpStreamedTraceSegment struct {
	HttpStreamedTraceSegment *HttpStreamedTraceSegment `protobuf:"bytes,2,opt,name=http_streamed_trace_segment,json=httpStreamedTraceSegment,proto3,oneof"`
}

type TraceWrapper_SocketBufferedTrace struct {
	SocketBufferedTrace *SocketBufferedTrace `protobuf:"bytes,3,opt,name=socket_buffered_trace,json=socketBufferedTrace,proto3,oneof"`
}

type TraceWrapper_SocketStreamedTraceSegment struct {
	SocketStreamedTraceSegment *SocketStreamedTraceSegment `protobuf:"bytes,4,opt,name=socket_streamed_trace_segment,json=socketStreamedTraceSegment,proto3,oneof"`
}

func (*TraceWrapper_HttpBufferedTrace) isTraceWrapper_Trace() {}

func (*TraceWrapper_HttpStreamedTraceSegment) isTraceWrapper_Trace() {}

func (*TraceWrapper_SocketBufferedTrace) isTraceWrapper_Trace() {}

func (*TraceWrapper_SocketStreamedTraceSegment) isTraceWrapper_Trace() {}

func (m *TraceWrapper) GetTrace() isTraceWrapper_Trace {
	if m != nil {
		return m.Trace
	}
	return nil
}

func (m *TraceWrapper) GetHttpBufferedTrace() *HttpBufferedTrace {
	if x, ok := m.GetTrace().(*TraceWrapper_HttpBufferedTrace); ok {
		return x.HttpBufferedTrace
	}
	return nil
}

func (m *TraceWrapper) GetHttpStreamedTraceSegment() *HttpStreamedTraceSegment {
	if x, ok := m.GetTrace().(*TraceWrapper_HttpStreamedTraceSegment); ok {
		return x.HttpStreamedTraceSegment
	}
	return nil
}

func (m *TraceWrapper) GetSocketBufferedTrace() *SocketBufferedTrace {
	if x, ok := m.GetTrace().(*TraceWrapper_SocketBufferedTrace); ok {
		return x.SocketBufferedTrace
	}
	return nil
}

func (m *TraceWrapper) GetSocketStreamedTraceSegment() *SocketStreamedTraceSegment {
	if x, ok := m.GetTrace().(*TraceWrapper_SocketStreamedTraceSegment); ok {
		return x.SocketStreamedTraceSegment
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TraceWrapper) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TraceWrapper_HttpBufferedTrace)(nil),
		(*TraceWrapper_HttpStreamedTraceSegment)(nil),
		(*TraceWrapper_SocketBufferedTrace)(nil),
		(*TraceWrapper_SocketStreamedTraceSegment)(nil),
	}
}

func init() {
	proto.RegisterType((*TraceWrapper)(nil), "envoy.data.tap.v2alpha.TraceWrapper")
}

func init() {
	proto.RegisterFile("envoy/data/tap/v2alpha/wrapper.proto", fileDescriptor_9f617738ad092e1c)
}

var fileDescriptor_9f617738ad092e1c = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x4a, 0xc3, 0x30,
	0x1c, 0xc7, 0x57, 0xe7, 0x3c, 0xc4, 0x21, 0x98, 0xa1, 0x96, 0x8a, 0xa0, 0x32, 0x44, 0x11, 0x52,
	0xa9, 0xe0, 0x03, 0xf4, 0xd4, 0xe3, 0x58, 0x05, 0x0f, 0x1e, 0xca, 0x6f, 0x6b, 0x66, 0x87, 0x5b,
	0x13, 0x93, 0x9f, 0x9d, 0x7b, 0x33, 0x8f, 0x9e, 0x7d, 0x13, 0xdf, 0x42, 0x9a, 0x66, 0x97, 0x6a,
	0x76, 0x2b, 0xf9, 0xfe, 0xf9, 0xe4, 0x4b, 0x43, 0x86, 0xbc, 0xac, 0xc4, 0x3a, 0xcc, 0x01, 0x21,
	0x44, 0x90, 0x61, 0x15, 0xc1, 0x42, 0x16, 0x10, 0xae, 0x14, 0x48, 0xc9, 0x15, 0x93, 0x4a, 0xa0,
	0xa0, 0xc7, 0xc6, 0xc5, 0x6a, 0x17, 0x43, 0x90, 0xcc, 0xba, 0x82, 0x0b, 0x47, 0xba, 0x40, 0x94,
	0x4d, 0x34, 0xb8, 0x72, 0x58, 0x50, 0x41, 0xa9, 0xa5, 0x50, 0x68, 0x7d, 0x27, 0x15, 0x2c, 0xe6,
	0x39, 0x20, 0x0f, 0x37, 0x1f, 0x8d, 0x70, 0xf9, 0xdd, 0x25, 0xfd, 0x47, 0x05, 0x53, 0xfe, 0xd4,
	0x5c, 0x89, 0x3e, 0x93, 0x41, 0xdd, 0x9f, 0x4d, 0xde, 0x67, 0x33, 0xae, 0x78, 0x9e, 0x61, 0xad,
	0xfa, 0xde, 0xb9, 0x77, 0xbd, 0x1f, 0xdd, 0xb0, 0xff, 0xaf, 0xca, 0x12, 0x44, 0x19, 0xdb, 0x84,
	0xa9, 0x4b, 0x3a, 0xe3, 0xc3, 0xa2, 0x7d, 0x48, 0xdf, 0xc8, 0xa9, 0x29, 0xd7, 0xa8, 0x38, 0x2c,
	0x37, 0xe5, 0x99, 0xe6, 0x2f, 0x4b, 0x5e, 0xa2, 0xbf, 0x63, 0x20, 0x77, 0xdb, 0x20, 0xa9, 0x4d,
	0x9a, 0xbe, 0xb4, 0xc9, 0x25, 0x9d, 0xb1, 0x5f, 0x38, 0x34, 0x0a, 0xe4, 0x48, 0x8b, 0xe9, 0x2b,
	0xc7, 0xf6, 0xa2, 0xae, 0x81, 0xdd, 0xba, 0x60, 0xa9, 0x09, 0xb5, 0x37, 0x0d, 0xf4, 0xdf, 0x63,
	0xba, 0x22, 0x67, 0x16, 0xe1, 0xd8, 0xb5, 0x6b, 0x50, 0xd1, 0x76, 0x94, 0x63, 0x59, 0xa0, 0x9d,
	0x6a, 0x7c, 0x40, 0x7a, 0x06, 0x44, 0x7b, 0x9f, 0x3f, 0x5f, 0x5d, 0x2f, 0x7e, 0x20, 0xc3, 0xb9,
	0x68, 0x28, 0x52, 0x89, 0x8f, 0xb5, 0x03, 0x18, 0xf7, 0xed, 0xcf, 0x1e, 0xd5, 0x4f, 0x60, 0xe4,
	0x4d, 0xf6, 0xcc, 0x5b, 0xb8, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x05, 0x31, 0xe1, 0x08, 0xaf,
	0x02, 0x00, 0x00,
}
