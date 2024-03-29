// Code generated by protoc-gen-go. DO NOT EDIT.
// source: connectioncontext.proto

package connectioncontext

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type IpFamily_Family int32

const (
	IpFamily_IPV4 IpFamily_Family = 0
	IpFamily_IPV6 IpFamily_Family = 1
)

var IpFamily_Family_name = map[int32]string{
	0: "IPV4",
	1: "IPV6",
}
var IpFamily_Family_value = map[string]int32{
	"IPV4": 0,
	"IPV6": 1,
}

func (x IpFamily_Family) String() string {
	return proto.EnumName(IpFamily_Family_name, int32(x))
}
func (IpFamily_Family) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_connectioncontext_05440d8eaa7bf48e, []int{2, 0}
}

type IpNeighbor struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	HardwareAddress      string   `protobuf:"bytes,2,opt,name=hardware_address,json=hardwareAddress,proto3" json:"hardware_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpNeighbor) Reset()         { *m = IpNeighbor{} }
func (m *IpNeighbor) String() string { return proto.CompactTextString(m) }
func (*IpNeighbor) ProtoMessage()    {}
func (*IpNeighbor) Descriptor() ([]byte, []int) {
	return fileDescriptor_connectioncontext_05440d8eaa7bf48e, []int{0}
}
func (m *IpNeighbor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpNeighbor.Unmarshal(m, b)
}
func (m *IpNeighbor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpNeighbor.Marshal(b, m, deterministic)
}
func (dst *IpNeighbor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpNeighbor.Merge(dst, src)
}
func (m *IpNeighbor) XXX_Size() int {
	return xxx_messageInfo_IpNeighbor.Size(m)
}
func (m *IpNeighbor) XXX_DiscardUnknown() {
	xxx_messageInfo_IpNeighbor.DiscardUnknown(m)
}

var xxx_messageInfo_IpNeighbor proto.InternalMessageInfo

func (m *IpNeighbor) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *IpNeighbor) GetHardwareAddress() string {
	if m != nil {
		return m.HardwareAddress
	}
	return ""
}

type Route struct {
	Prefix               string   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_connectioncontext_05440d8eaa7bf48e, []int{1}
}
func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (dst *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(dst, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

type IpFamily struct {
	Family               IpFamily_Family `protobuf:"varint,1,opt,name=family,proto3,enum=connectioncontext.IpFamily_Family" json:"family,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *IpFamily) Reset()         { *m = IpFamily{} }
func (m *IpFamily) String() string { return proto.CompactTextString(m) }
func (*IpFamily) ProtoMessage()    {}
func (*IpFamily) Descriptor() ([]byte, []int) {
	return fileDescriptor_connectioncontext_05440d8eaa7bf48e, []int{2}
}
func (m *IpFamily) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpFamily.Unmarshal(m, b)
}
func (m *IpFamily) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpFamily.Marshal(b, m, deterministic)
}
func (dst *IpFamily) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpFamily.Merge(dst, src)
}
func (m *IpFamily) XXX_Size() int {
	return xxx_messageInfo_IpFamily.Size(m)
}
func (m *IpFamily) XXX_DiscardUnknown() {
	xxx_messageInfo_IpFamily.DiscardUnknown(m)
}

var xxx_messageInfo_IpFamily proto.InternalMessageInfo

func (m *IpFamily) GetFamily() IpFamily_Family {
	if m != nil {
		return m.Family
	}
	return IpFamily_IPV4
}

type ExtraPrefixRequest struct {
	AddrFamily           *IpFamily `protobuf:"bytes,1,opt,name=addr_family,json=addrFamily,proto3" json:"addr_family,omitempty"`
	PrefixLen            uint32    `protobuf:"varint,2,opt,name=prefix_len,json=prefixLen,proto3" json:"prefix_len,omitempty"`
	RequiredNumber       uint32    `protobuf:"varint,3,opt,name=required_number,json=requiredNumber,proto3" json:"required_number,omitempty"`
	RequestedNumber      uint32    `protobuf:"varint,4,opt,name=requested_number,json=requestedNumber,proto3" json:"requested_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ExtraPrefixRequest) Reset()         { *m = ExtraPrefixRequest{} }
func (m *ExtraPrefixRequest) String() string { return proto.CompactTextString(m) }
func (*ExtraPrefixRequest) ProtoMessage()    {}
func (*ExtraPrefixRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_connectioncontext_05440d8eaa7bf48e, []int{3}
}
func (m *ExtraPrefixRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraPrefixRequest.Unmarshal(m, b)
}
func (m *ExtraPrefixRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraPrefixRequest.Marshal(b, m, deterministic)
}
func (dst *ExtraPrefixRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraPrefixRequest.Merge(dst, src)
}
func (m *ExtraPrefixRequest) XXX_Size() int {
	return xxx_messageInfo_ExtraPrefixRequest.Size(m)
}
func (m *ExtraPrefixRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraPrefixRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraPrefixRequest proto.InternalMessageInfo

func (m *ExtraPrefixRequest) GetAddrFamily() *IpFamily {
	if m != nil {
		return m.AddrFamily
	}
	return nil
}

func (m *ExtraPrefixRequest) GetPrefixLen() uint32 {
	if m != nil {
		return m.PrefixLen
	}
	return 0
}

func (m *ExtraPrefixRequest) GetRequiredNumber() uint32 {
	if m != nil {
		return m.RequiredNumber
	}
	return 0
}

func (m *ExtraPrefixRequest) GetRequestedNumber() uint32 {
	if m != nil {
		return m.RequestedNumber
	}
	return 0
}

type ConnectionContext struct {
	SrcIpAddr            string                `protobuf:"bytes,1,opt,name=src_ip_addr,json=srcIpAddr,proto3" json:"src_ip_addr,omitempty"`
	DstIpAddr            string                `protobuf:"bytes,2,opt,name=dst_ip_addr,json=dstIpAddr,proto3" json:"dst_ip_addr,omitempty"`
	SrcIpRequired        bool                  `protobuf:"varint,3,opt,name=src_ip_required,json=srcIpRequired,proto3" json:"src_ip_required,omitempty"`
	DstIpRequired        bool                  `protobuf:"varint,4,opt,name=dst_ip_required,json=dstIpRequired,proto3" json:"dst_ip_required,omitempty"`
	Routes               []*Route              `protobuf:"bytes,5,rep,name=routes,proto3" json:"routes,omitempty"`
	ExcludedPrefixes     []string              `protobuf:"bytes,6,rep,name=excluded_prefixes,json=excludedPrefixes,proto3" json:"excluded_prefixes,omitempty"`
	IpNeighbors          []*IpNeighbor         `protobuf:"bytes,7,rep,name=ip_neighbors,json=ipNeighbors,proto3" json:"ip_neighbors,omitempty"`
	ExtraPrefixRequest   []*ExtraPrefixRequest `protobuf:"bytes,8,rep,name=extra_prefix_request,json=extraPrefixRequest,proto3" json:"extra_prefix_request,omitempty"`
	ExtraPrefixes        []string              `protobuf:"bytes,9,rep,name=extra_prefixes,json=extraPrefixes,proto3" json:"extra_prefixes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ConnectionContext) Reset()         { *m = ConnectionContext{} }
func (m *ConnectionContext) String() string { return proto.CompactTextString(m) }
func (*ConnectionContext) ProtoMessage()    {}
func (*ConnectionContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_connectioncontext_05440d8eaa7bf48e, []int{4}
}
func (m *ConnectionContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionContext.Unmarshal(m, b)
}
func (m *ConnectionContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionContext.Marshal(b, m, deterministic)
}
func (dst *ConnectionContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionContext.Merge(dst, src)
}
func (m *ConnectionContext) XXX_Size() int {
	return xxx_messageInfo_ConnectionContext.Size(m)
}
func (m *ConnectionContext) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionContext.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionContext proto.InternalMessageInfo

func (m *ConnectionContext) GetSrcIpAddr() string {
	if m != nil {
		return m.SrcIpAddr
	}
	return ""
}

func (m *ConnectionContext) GetDstIpAddr() string {
	if m != nil {
		return m.DstIpAddr
	}
	return ""
}

func (m *ConnectionContext) GetSrcIpRequired() bool {
	if m != nil {
		return m.SrcIpRequired
	}
	return false
}

func (m *ConnectionContext) GetDstIpRequired() bool {
	if m != nil {
		return m.DstIpRequired
	}
	return false
}

func (m *ConnectionContext) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *ConnectionContext) GetExcludedPrefixes() []string {
	if m != nil {
		return m.ExcludedPrefixes
	}
	return nil
}

func (m *ConnectionContext) GetIpNeighbors() []*IpNeighbor {
	if m != nil {
		return m.IpNeighbors
	}
	return nil
}

func (m *ConnectionContext) GetExtraPrefixRequest() []*ExtraPrefixRequest {
	if m != nil {
		return m.ExtraPrefixRequest
	}
	return nil
}

func (m *ConnectionContext) GetExtraPrefixes() []string {
	if m != nil {
		return m.ExtraPrefixes
	}
	return nil
}

func init() {
	proto.RegisterType((*IpNeighbor)(nil), "connectioncontext.IpNeighbor")
	proto.RegisterType((*Route)(nil), "connectioncontext.Route")
	proto.RegisterType((*IpFamily)(nil), "connectioncontext.IpFamily")
	proto.RegisterType((*ExtraPrefixRequest)(nil), "connectioncontext.ExtraPrefixRequest")
	proto.RegisterType((*ConnectionContext)(nil), "connectioncontext.ConnectionContext")
	proto.RegisterEnum("connectioncontext.IpFamily_Family", IpFamily_Family_name, IpFamily_Family_value)
}

func init() {
	proto.RegisterFile("connectioncontext.proto", fileDescriptor_connectioncontext_05440d8eaa7bf48e)
}

var fileDescriptor_connectioncontext_05440d8eaa7bf48e = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x53, 0xdf, 0x8b, 0xd3, 0x40,
	0x10, 0xb6, 0xd7, 0x5e, 0x6c, 0xa6, 0xf6, 0xd7, 0x2a, 0x1a, 0xd0, 0xd3, 0x23, 0x70, 0xfe, 0x40,
	0x28, 0x72, 0x8a, 0x0f, 0xe2, 0x83, 0x7a, 0xa8, 0x14, 0xe4, 0x28, 0xfb, 0xa0, 0xe0, 0x4b, 0x48,
	0x9b, 0x39, 0x2f, 0xd0, 0x4b, 0xe2, 0xee, 0x16, 0xeb, 0xff, 0xa7, 0xff, 0x97, 0xbb, 0xb3, 0x93,
	0x58, 0x48, 0xb9, 0xa7, 0x4e, 0xbf, 0xf9, 0x66, 0xe6, 0xdb, 0x6f, 0x26, 0x70, 0x6f, 0x55, 0x16,
	0x05, 0xae, 0x4c, 0x5e, 0x16, 0x36, 0x32, 0xb8, 0x35, 0xb3, 0x4a, 0x95, 0xa6, 0x14, 0xd3, 0x56,
	0x22, 0xfe, 0x0c, 0x30, 0xaf, 0xce, 0x31, 0xff, 0x71, 0xb9, 0x2c, 0x95, 0x18, 0xc1, 0x41, 0x5e,
	0x45, 0x9d, 0xe3, 0xce, 0xd3, 0x50, 0xda, 0x48, 0x3c, 0x83, 0xc9, 0x65, 0xaa, 0xb2, 0x5f, 0xa9,
	0xc2, 0x24, 0xcd, 0x32, 0x85, 0x5a, 0x47, 0x07, 0x94, 0x1d, 0xd7, 0xf8, 0x7b, 0x0f, 0xc7, 0x8f,
	0xe0, 0x50, 0x96, 0x1b, 0x83, 0xe2, 0x2e, 0x04, 0x95, 0xc2, 0x8b, 0x7c, 0xcb, 0x7d, 0xf8, 0x5f,
	0x9c, 0x41, 0x7f, 0x5e, 0x7d, 0x4a, 0xaf, 0xf2, 0xf5, 0x6f, 0xf1, 0x06, 0x82, 0x0b, 0x8a, 0x88,
	0x33, 0x3a, 0x8d, 0x67, 0x6d, 0xc9, 0x35, 0x79, 0xe6, 0x7f, 0x24, 0x57, 0xc4, 0x0f, 0x20, 0xe0,
	0x2e, 0x7d, 0xe8, 0xcd, 0x17, 0x5f, 0x5f, 0x4d, 0x6e, 0x70, 0xf4, 0x7a, 0xd2, 0x89, 0xff, 0x76,
	0x40, 0x7c, 0xdc, 0x1a, 0x95, 0x2e, 0x68, 0xaa, 0xc4, 0x9f, 0x1b, 0xd4, 0x46, 0xbc, 0x85, 0x81,
	0xd3, 0x9f, 0xec, 0x4c, 0x1d, 0x9c, 0xde, 0xbf, 0x66, 0xaa, 0x04, 0xc7, 0xe7, 0x41, 0x47, 0x00,
	0xfe, 0x11, 0xc9, 0x1a, 0x0b, 0x32, 0x60, 0x28, 0x43, 0x8f, 0x7c, 0xc1, 0x42, 0x3c, 0x81, 0xb1,
	0xb2, 0x73, 0x72, 0x85, 0x59, 0x52, 0x6c, 0xae, 0x96, 0xa8, 0xa2, 0x2e, 0x71, 0x46, 0x35, 0x7c,
	0x4e, 0xa8, 0xb3, 0x53, 0x79, 0x41, 0xff, 0x99, 0x3d, 0x62, 0x8e, 0x1b, 0xdc, 0x53, 0xe3, 0x3f,
	0x5d, 0x98, 0x9e, 0x35, 0xea, 0xce, 0xbc, 0x3a, 0xf1, 0x10, 0x06, 0x5a, 0xad, 0x92, 0xbc, 0xa2,
	0x6d, 0xb0, 0xc1, 0xa1, 0x85, 0xe6, 0x95, 0xdb, 0x83, 0xcb, 0x67, 0xda, 0x34, 0x79, 0xbf, 0xaa,
	0xd0, 0x42, 0x9c, 0x7f, 0x0c, 0x63, 0xae, 0xaf, 0x95, 0x91, 0xd2, 0xbe, 0x1c, 0x52, 0x0f, 0xc9,
	0xa0, 0xe3, 0x71, 0x9f, 0x86, 0xd7, 0xf3, 0x3c, 0xea, 0xd5, 0xf0, 0x5e, 0x40, 0xa0, 0xdc, 0xd2,
	0x75, 0x74, 0x78, 0xdc, 0xb5, 0x8e, 0x46, 0x7b, 0x1c, 0xa5, 0xab, 0x90, 0xcc, 0x13, 0xcf, 0x61,
	0x8a, 0xdb, 0xd5, 0x7a, 0x93, 0x59, 0x07, 0xbc, 0x83, 0xb6, 0x38, 0xb0, 0xc5, 0xa1, 0x9c, 0xd4,
	0x89, 0x05, 0xe3, 0xe2, 0x1d, 0xdc, 0xb2, 0x12, 0x0a, 0xbe, 0x4e, 0x1d, 0xdd, 0xa4, 0x21, 0x47,
	0x7b, 0xd7, 0x56, 0xdf, 0xb0, 0x1c, 0xe4, 0x4d, 0xac, 0xc5, 0x37, 0xb8, 0x83, 0xee, 0x1a, 0x78,
	0x56, 0xc2, 0x36, 0x47, 0x7d, 0xea, 0x74, 0xb2, 0xa7, 0x53, 0xfb, 0x78, 0xa4, 0xc0, 0xf6, 0x41,
	0x9d, 0xc0, 0x68, 0xb7, 0xb1, 0x7d, 0x44, 0x48, 0x8f, 0x18, 0xee, 0x70, 0x51, 0x7f, 0xb8, 0xfd,
	0xbd, 0xfd, 0xcd, 0x2d, 0x03, 0xfa, 0x1a, 0x5f, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x01, 0xa6,
	0xaf, 0xfc, 0xa8, 0x03, 0x00, 0x00,
}
