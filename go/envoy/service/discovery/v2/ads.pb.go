// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/discovery/v2/ads.proto

package v2

import (
	context "context"
	fmt "fmt"
	v2 "github.com/cilium/proxy/go/envoy/api/v2"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

// [#not-implemented-hide:] Not configuration. Workaround c++ protobuf issue with importing
// services: https://github.com/google/protobuf/issues/4221
type AdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdsDummy) Reset()         { *m = AdsDummy{} }
func (m *AdsDummy) String() string { return proto.CompactTextString(m) }
func (*AdsDummy) ProtoMessage()    {}
func (*AdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_187fd5dcc2dab695, []int{0}
}

func (m *AdsDummy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdsDummy.Unmarshal(m, b)
}
func (m *AdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdsDummy.Marshal(b, m, deterministic)
}
func (m *AdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdsDummy.Merge(m, src)
}
func (m *AdsDummy) XXX_Size() int {
	return xxx_messageInfo_AdsDummy.Size(m)
}
func (m *AdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_AdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_AdsDummy proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdsDummy)(nil), "envoy.service.discovery.v2.AdsDummy")
}

func init() {
	proto.RegisterFile("envoy/service/discovery/v2/ads.proto", fileDescriptor_187fd5dcc2dab695)
}

var fileDescriptor_187fd5dcc2dab695 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x31, 0x03, 0x42, 0x1e, 0x33, 0x41, 0x84, 0x40, 0x2a, 0x1d, 0x3a, 0x5d, 0x90, 0x99,
	0x19, 0x5a, 0xe5, 0x01, 0xaa, 0x76, 0x63, 0x73, 0x93, 0x53, 0x64, 0x41, 0x7a, 0xc6, 0xe7, 0x58,
	0xf8, 0x0d, 0x78, 0x59, 0xde, 0x01, 0x39, 0x86, 0x16, 0x01, 0x61, 0xbe, 0xef, 0xff, 0xff, 0xd3,
	0x27, 0xe7, 0xb8, 0x0f, 0x14, 0x2b, 0x46, 0x17, 0x4c, 0x83, 0x55, 0x6b, 0xb8, 0xa1, 0x80, 0x2e,
	0x56, 0x41, 0x55, 0xba, 0x65, 0xb0, 0x8e, 0x3c, 0x15, 0xe5, 0x48, 0xc1, 0x27, 0x05, 0x07, 0x0a,
	0x82, 0x2a, 0xaf, 0x72, 0x83, 0xb6, 0x26, 0x65, 0x8e, 0xa7, 0x31, 0x39, 0x93, 0xf2, 0x7c, 0xd9,
	0x72, 0x3d, 0xf4, 0x7d, 0x54, 0xef, 0x42, 0x96, 0xcb, 0xae, 0x73, 0xd8, 0x69, 0x8f, 0x6d, 0xfd,
	0x45, 0x6e, 0x73, 0x6b, 0xb1, 0x93, 0x97, 0x5b, 0xef, 0x50, 0xf7, 0x47, 0x66, 0x83, 0x4c, 0x83,
	0x6b, 0x90, 0x8b, 0x6b, 0xc8, 0x2f, 0x68, 0x6b, 0x20, 0x28, 0x38, 0x84, 0x37, 0xf8, 0x32, 0x20,
	0xfb, 0xf2, 0x66, 0xf2, 0xce, 0x96, 0xf6, 0x8c, 0xb3, 0x93, 0x85, 0xb8, 0x13, 0xc5, 0x93, 0xbc,
	0xa8, 0xf1, 0xd9, 0xeb, 0xbf, 0x26, 0x6e, 0x7f, 0x54, 0x24, 0xee, 0xd7, 0xce, 0xfc, 0x7f, 0xe8,
	0xfb, 0xd8, 0xea, 0x41, 0x2e, 0x0c, 0x65, 0xde, 0x3a, 0x7a, 0x8d, 0x30, 0x6d, 0x71, 0x95, 0x2c,
	0xad, 0x93, 0xb1, 0xb5, 0x78, 0x3c, 0x0d, 0xea, 0x4d, 0x88, 0xdd, 0xd9, 0x68, 0xf0, 0xfe, 0x23,
	0x00, 0x00, 0xff, 0xff, 0x7e, 0x66, 0x01, 0x47, 0xa3, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AggregatedDiscoveryServiceClient is the client API for AggregatedDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AggregatedDiscoveryServiceClient interface {
	// This is a gRPC-only API.
	StreamAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (AggregatedDiscoveryService_StreamAggregatedResourcesClient, error)
	DeltaAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (AggregatedDiscoveryService_DeltaAggregatedResourcesClient, error)
}

type aggregatedDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewAggregatedDiscoveryServiceClient(cc *grpc.ClientConn) AggregatedDiscoveryServiceClient {
	return &aggregatedDiscoveryServiceClient{cc}
}

func (c *aggregatedDiscoveryServiceClient) StreamAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (AggregatedDiscoveryService_StreamAggregatedResourcesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AggregatedDiscoveryService_serviceDesc.Streams[0], "/envoy.service.discovery.v2.AggregatedDiscoveryService/StreamAggregatedResources", opts...)
	if err != nil {
		return nil, err
	}
	x := &aggregatedDiscoveryServiceStreamAggregatedResourcesClient{stream}
	return x, nil
}

type AggregatedDiscoveryService_StreamAggregatedResourcesClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type aggregatedDiscoveryServiceStreamAggregatedResourcesClient struct {
	grpc.ClientStream
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aggregatedDiscoveryServiceClient) DeltaAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (AggregatedDiscoveryService_DeltaAggregatedResourcesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AggregatedDiscoveryService_serviceDesc.Streams[1], "/envoy.service.discovery.v2.AggregatedDiscoveryService/DeltaAggregatedResources", opts...)
	if err != nil {
		return nil, err
	}
	x := &aggregatedDiscoveryServiceDeltaAggregatedResourcesClient{stream}
	return x, nil
}

type AggregatedDiscoveryService_DeltaAggregatedResourcesClient interface {
	Send(*v2.DeltaDiscoveryRequest) error
	Recv() (*v2.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type aggregatedDiscoveryServiceDeltaAggregatedResourcesClient struct {
	grpc.ClientStream
}

func (x *aggregatedDiscoveryServiceDeltaAggregatedResourcesClient) Send(m *v2.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceDeltaAggregatedResourcesClient) Recv() (*v2.DeltaDiscoveryResponse, error) {
	m := new(v2.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AggregatedDiscoveryServiceServer is the server API for AggregatedDiscoveryService service.
type AggregatedDiscoveryServiceServer interface {
	// This is a gRPC-only API.
	StreamAggregatedResources(AggregatedDiscoveryService_StreamAggregatedResourcesServer) error
	DeltaAggregatedResources(AggregatedDiscoveryService_DeltaAggregatedResourcesServer) error
}

func RegisterAggregatedDiscoveryServiceServer(s *grpc.Server, srv AggregatedDiscoveryServiceServer) {
	s.RegisterService(&_AggregatedDiscoveryService_serviceDesc, srv)
}

func _AggregatedDiscoveryService_StreamAggregatedResources_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AggregatedDiscoveryServiceServer).StreamAggregatedResources(&aggregatedDiscoveryServiceStreamAggregatedResourcesServer{stream})
}

type AggregatedDiscoveryService_StreamAggregatedResourcesServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type aggregatedDiscoveryServiceStreamAggregatedResourcesServer struct {
	grpc.ServerStream
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AggregatedDiscoveryService_DeltaAggregatedResources_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AggregatedDiscoveryServiceServer).DeltaAggregatedResources(&aggregatedDiscoveryServiceDeltaAggregatedResourcesServer{stream})
}

type AggregatedDiscoveryService_DeltaAggregatedResourcesServer interface {
	Send(*v2.DeltaDiscoveryResponse) error
	Recv() (*v2.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type aggregatedDiscoveryServiceDeltaAggregatedResourcesServer struct {
	grpc.ServerStream
}

func (x *aggregatedDiscoveryServiceDeltaAggregatedResourcesServer) Send(m *v2.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceDeltaAggregatedResourcesServer) Recv() (*v2.DeltaDiscoveryRequest, error) {
	m := new(v2.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AggregatedDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.discovery.v2.AggregatedDiscoveryService",
	HandlerType: (*AggregatedDiscoveryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamAggregatedResources",
			Handler:       _AggregatedDiscoveryService_StreamAggregatedResources_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaAggregatedResources",
			Handler:       _AggregatedDiscoveryService_DeltaAggregatedResources_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/discovery/v2/ads.proto",
}
