// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/ip_tagging/v2/ip_tagging.proto

package v2

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
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

// The type of requests the filter should apply to. The supported types
// are internal, external or both. The
// :ref:`x-forwarded-for<config_http_conn_man_headers_x-forwarded-for_internal_origin>` header is
// used to determine if a request is internal and will result in
// :ref:`x-envoy-internal<config_http_conn_man_headers_x-envoy-internal>`
// being set. The filter defaults to both, and it will apply to all request types.
type IPTagging_RequestType int32

const (
	// Both external and internal requests will be tagged. This is the default value.
	IPTagging_BOTH IPTagging_RequestType = 0
	// Only internal requests will be tagged.
	IPTagging_INTERNAL IPTagging_RequestType = 1
	// Only external requests will be tagged.
	IPTagging_EXTERNAL IPTagging_RequestType = 2
)

var IPTagging_RequestType_name = map[int32]string{
	0: "BOTH",
	1: "INTERNAL",
	2: "EXTERNAL",
}

var IPTagging_RequestType_value = map[string]int32{
	"BOTH":     0,
	"INTERNAL": 1,
	"EXTERNAL": 2,
}

func (x IPTagging_RequestType) String() string {
	return proto.EnumName(IPTagging_RequestType_name, int32(x))
}

func (IPTagging_RequestType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f4275d0b367744d2, []int{0, 0}
}

type IPTagging struct {
	// The type of request the filter should apply to.
	RequestType IPTagging_RequestType `protobuf:"varint,1,opt,name=request_type,json=requestType,proto3,enum=envoy.config.filter.http.ip_tagging.v2.IPTagging_RequestType" json:"request_type,omitempty"`
	// [#comment:TODO(ccaraman): Extend functionality to load IP tags from file system.
	// Tracked by issue https://github.com/envoyproxy/envoy/issues/2695]
	// The set of IP tags for the filter.
	IpTags               []*IPTagging_IPTag `protobuf:"bytes,4,rep,name=ip_tags,json=ipTags,proto3" json:"ip_tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *IPTagging) Reset()         { *m = IPTagging{} }
func (m *IPTagging) String() string { return proto.CompactTextString(m) }
func (*IPTagging) ProtoMessage()    {}
func (*IPTagging) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4275d0b367744d2, []int{0}
}

func (m *IPTagging) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPTagging.Unmarshal(m, b)
}
func (m *IPTagging) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPTagging.Marshal(b, m, deterministic)
}
func (m *IPTagging) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPTagging.Merge(m, src)
}
func (m *IPTagging) XXX_Size() int {
	return xxx_messageInfo_IPTagging.Size(m)
}
func (m *IPTagging) XXX_DiscardUnknown() {
	xxx_messageInfo_IPTagging.DiscardUnknown(m)
}

var xxx_messageInfo_IPTagging proto.InternalMessageInfo

func (m *IPTagging) GetRequestType() IPTagging_RequestType {
	if m != nil {
		return m.RequestType
	}
	return IPTagging_BOTH
}

func (m *IPTagging) GetIpTags() []*IPTagging_IPTag {
	if m != nil {
		return m.IpTags
	}
	return nil
}

// Supplies the IP tag name and the IP address subnets.
type IPTagging_IPTag struct {
	// Specifies the IP tag name to apply.
	IpTagName string `protobuf:"bytes,1,opt,name=ip_tag_name,json=ipTagName,proto3" json:"ip_tag_name,omitempty"`
	// A list of IP address subnets that will be tagged with
	// ip_tag_name. Both IPv4 and IPv6 are supported.
	IpList               []*core.CidrRange `protobuf:"bytes,2,rep,name=ip_list,json=ipList,proto3" json:"ip_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *IPTagging_IPTag) Reset()         { *m = IPTagging_IPTag{} }
func (m *IPTagging_IPTag) String() string { return proto.CompactTextString(m) }
func (*IPTagging_IPTag) ProtoMessage()    {}
func (*IPTagging_IPTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4275d0b367744d2, []int{0, 0}
}

func (m *IPTagging_IPTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPTagging_IPTag.Unmarshal(m, b)
}
func (m *IPTagging_IPTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPTagging_IPTag.Marshal(b, m, deterministic)
}
func (m *IPTagging_IPTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPTagging_IPTag.Merge(m, src)
}
func (m *IPTagging_IPTag) XXX_Size() int {
	return xxx_messageInfo_IPTagging_IPTag.Size(m)
}
func (m *IPTagging_IPTag) XXX_DiscardUnknown() {
	xxx_messageInfo_IPTagging_IPTag.DiscardUnknown(m)
}

var xxx_messageInfo_IPTagging_IPTag proto.InternalMessageInfo

func (m *IPTagging_IPTag) GetIpTagName() string {
	if m != nil {
		return m.IpTagName
	}
	return ""
}

func (m *IPTagging_IPTag) GetIpList() []*core.CidrRange {
	if m != nil {
		return m.IpList
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.filter.http.ip_tagging.v2.IPTagging_RequestType", IPTagging_RequestType_name, IPTagging_RequestType_value)
	proto.RegisterType((*IPTagging)(nil), "envoy.config.filter.http.ip_tagging.v2.IPTagging")
	proto.RegisterType((*IPTagging_IPTag)(nil), "envoy.config.filter.http.ip_tagging.v2.IPTagging.IPTag")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/ip_tagging/v2/ip_tagging.proto", fileDescriptor_f4275d0b367744d2)
}

var fileDescriptor_f4275d0b367744d2 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x4d, 0xf7, 0xc7, 0x2d, 0x1d, 0x63, 0xf4, 0xe2, 0x18, 0xa2, 0x63, 0x07, 0xd9, 0x29,
	0x85, 0x4e, 0xd9, 0xc9, 0x83, 0x95, 0x81, 0x83, 0x31, 0x47, 0xe9, 0x41, 0x44, 0x1c, 0x71, 0xcd,
	0x6a, 0xa0, 0x6b, 0x63, 0x12, 0x8b, 0xbd, 0xfa, 0x11, 0xfc, 0x38, 0x9e, 0xfc, 0x3a, 0x82, 0x1f,
	0x42, 0x92, 0x54, 0xdd, 0x71, 0xde, 0xf2, 0xf2, 0xe6, 0xf9, 0x3d, 0xcf, 0x93, 0xc0, 0x31, 0x49,
	0xf3, 0xac, 0x70, 0x57, 0x59, 0xba, 0xa6, 0xb1, 0xbb, 0xa6, 0x89, 0x24, 0xdc, 0x7d, 0x94, 0x92,
	0xb9, 0x94, 0x2d, 0x25, 0x8e, 0x63, 0x9a, 0xc6, 0x6e, 0xee, 0x6d, 0x4d, 0x88, 0xf1, 0x4c, 0x66,
	0xce, 0x89, 0x16, 0x22, 0x23, 0x44, 0x46, 0x88, 0x94, 0x10, 0x6d, 0x5d, 0xcd, 0xbd, 0xde, 0xb1,
	0x31, 0xc0, 0x8c, 0x2a, 0xcc, 0x2a, 0xe3, 0xc4, 0xc5, 0x51, 0xc4, 0x89, 0x10, 0x06, 0xd4, 0x3b,
	0xc8, 0x71, 0x42, 0x23, 0x2c, 0x89, 0xfb, 0x73, 0x30, 0x8b, 0xc1, 0x97, 0x05, 0x9b, 0xd3, 0x45,
	0x68, 0x50, 0x4e, 0x02, 0x5b, 0x9c, 0x3c, 0x3d, 0x13, 0x21, 0x97, 0xb2, 0x60, 0xa4, 0x0b, 0xfa,
	0x60, 0xd8, 0xf6, 0xce, 0xd1, 0x6e, 0x31, 0xd0, 0x2f, 0x08, 0x05, 0x86, 0x12, 0x16, 0x8c, 0xf8,
	0xf0, 0xfd, 0xf3, 0xa3, 0x52, 0x7b, 0x05, 0x56, 0x07, 0x04, 0x36, 0xff, 0x5b, 0x38, 0x77, 0x70,
	0xdf, 0xe8, 0x45, 0xb7, 0xda, 0xaf, 0x0c, 0x6d, 0x6f, 0xfc, 0x7f, 0x23, 0x7d, 0x2a, 0x2d, 0xde,
	0x80, 0xd5, 0x00, 0x41, 0x9d, 0xb2, 0x10, 0xc7, 0xa2, 0x77, 0x0f, 0x6b, 0x7a, 0xe9, 0x1c, 0x41,
	0xdb, 0xa8, 0x97, 0x29, 0xde, 0x98, 0x4e, 0xcd, 0xa0, 0xa9, 0x6f, 0xcd, 0xf1, 0x86, 0x38, 0x67,
	0x3a, 0x46, 0x42, 0x85, 0xec, 0x5a, 0x3a, 0xc6, 0x61, 0x19, 0x03, 0x33, 0xaa, 0xcc, 0xd4, 0x73,
	0xa2, 0x4b, 0x1a, 0xf1, 0x00, 0xa7, 0x31, 0x51, 0xfc, 0x19, 0x15, 0x72, 0x30, 0x82, 0xf6, 0x56,
	0x4b, 0xa7, 0x01, 0xab, 0xfe, 0x75, 0x78, 0xd5, 0xd9, 0x73, 0x5a, 0xb0, 0x31, 0x9d, 0x87, 0x93,
	0x60, 0x7e, 0x31, 0xeb, 0x00, 0x35, 0x4d, 0x6e, 0xca, 0xc9, 0xf2, 0x67, 0xf0, 0x94, 0x66, 0x06,
	0xcf, 0x78, 0xf6, 0x52, 0xec, 0x58, 0xd8, 0x6f, 0x4f, 0x59, 0xd9, 0x78, 0xa1, 0xbe, 0x6d, 0x01,
	0x6e, 0xad, 0xdc, 0x7b, 0xa8, 0xeb, 0x3f, 0x1c, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x10,
	0x24, 0xee, 0x60, 0x02, 0x00, 0x00,
}
