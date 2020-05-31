// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: models/vpp/punt/punt.proto

package vpp_punt

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type L3Protocol int32

const (
	L3Protocol_UNDEFINED_L3 L3Protocol = 0
	L3Protocol_IPv4         L3Protocol = 4
	L3Protocol_IPv6         L3Protocol = 6
	L3Protocol_ALL          L3Protocol = 10
)

var L3Protocol_name = map[int32]string{
	0:  "UNDEFINED_L3",
	4:  "IPv4",
	6:  "IPv6",
	10: "ALL",
}

var L3Protocol_value = map[string]int32{
	"UNDEFINED_L3": 0,
	"IPv4":         4,
	"IPv6":         6,
	"ALL":          10,
}

func (x L3Protocol) String() string {
	return proto.EnumName(L3Protocol_name, int32(x))
}

func (L3Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c6553556423e0b80, []int{0}
}

type L4Protocol int32

const (
	L4Protocol_UNDEFINED_L4 L4Protocol = 0
	L4Protocol_TCP          L4Protocol = 6
	L4Protocol_UDP          L4Protocol = 17
)

var L4Protocol_name = map[int32]string{
	0:  "UNDEFINED_L4",
	6:  "TCP",
	17: "UDP",
}

var L4Protocol_value = map[string]int32{
	"UNDEFINED_L4": 0,
	"TCP":          6,
	"UDP":          17,
}

func (x L4Protocol) String() string {
	return proto.EnumName(L4Protocol_name, int32(x))
}

func (L4Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c6553556423e0b80, []int{1}
}

// IPRedirect allows otherwise dropped packet which destination IP address matching some of the VPP addresses
//to redirect to the defined next hop address via the TX interface
type IPRedirect struct {
	L3Protocol           L3Protocol `protobuf:"varint,1,opt,name=l3_protocol,json=l3Protocol,proto3,enum=vpp.punt.L3Protocol" json:"l3_protocol,omitempty"`
	RxInterface          string     `protobuf:"bytes,2,opt,name=rx_interface,json=rxInterface,proto3" json:"rx_interface,omitempty"`
	TxInterface          string     `protobuf:"bytes,3,opt,name=tx_interface,json=txInterface,proto3" json:"tx_interface,omitempty"`
	NextHop              string     `protobuf:"bytes,4,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *IPRedirect) Reset()         { *m = IPRedirect{} }
func (m *IPRedirect) String() string { return proto.CompactTextString(m) }
func (*IPRedirect) ProtoMessage()    {}
func (*IPRedirect) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6553556423e0b80, []int{0}
}
func (m *IPRedirect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPRedirect.Unmarshal(m, b)
}
func (m *IPRedirect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPRedirect.Marshal(b, m, deterministic)
}
func (m *IPRedirect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPRedirect.Merge(m, src)
}
func (m *IPRedirect) XXX_Size() int {
	return xxx_messageInfo_IPRedirect.Size(m)
}
func (m *IPRedirect) XXX_DiscardUnknown() {
	xxx_messageInfo_IPRedirect.DiscardUnknown(m)
}

var xxx_messageInfo_IPRedirect proto.InternalMessageInfo

func (m *IPRedirect) GetL3Protocol() L3Protocol {
	if m != nil {
		return m.L3Protocol
	}
	return L3Protocol_UNDEFINED_L3
}

func (m *IPRedirect) GetRxInterface() string {
	if m != nil {
		return m.RxInterface
	}
	return ""
}

func (m *IPRedirect) GetTxInterface() string {
	if m != nil {
		return m.TxInterface
	}
	return ""
}

func (m *IPRedirect) GetNextHop() string {
	if m != nil {
		return m.NextHop
	}
	return ""
}

func (*IPRedirect) XXX_MessageName() string {
	return "vpp.punt.IPRedirect"
}

// allows otherwise dropped packet which destination IP address matching some of the VPP interface IP addresses to be
//punted to the host. L3 and L4 protocols can be used for filtering
type ToHost struct {
	L3Protocol           L3Protocol `protobuf:"varint,2,opt,name=l3_protocol,json=l3Protocol,proto3,enum=vpp.punt.L3Protocol" json:"l3_protocol,omitempty"`
	L4Protocol           L4Protocol `protobuf:"varint,3,opt,name=l4_protocol,json=l4Protocol,proto3,enum=vpp.punt.L4Protocol" json:"l4_protocol,omitempty"`
	Port                 uint32     `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	SocketPath           string     `protobuf:"bytes,5,opt,name=socket_path,json=socketPath,proto3" json:"socket_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ToHost) Reset()         { *m = ToHost{} }
func (m *ToHost) String() string { return proto.CompactTextString(m) }
func (*ToHost) ProtoMessage()    {}
func (*ToHost) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6553556423e0b80, []int{1}
}
func (m *ToHost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToHost.Unmarshal(m, b)
}
func (m *ToHost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToHost.Marshal(b, m, deterministic)
}
func (m *ToHost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToHost.Merge(m, src)
}
func (m *ToHost) XXX_Size() int {
	return xxx_messageInfo_ToHost.Size(m)
}
func (m *ToHost) XXX_DiscardUnknown() {
	xxx_messageInfo_ToHost.DiscardUnknown(m)
}

var xxx_messageInfo_ToHost proto.InternalMessageInfo

func (m *ToHost) GetL3Protocol() L3Protocol {
	if m != nil {
		return m.L3Protocol
	}
	return L3Protocol_UNDEFINED_L3
}

func (m *ToHost) GetL4Protocol() L4Protocol {
	if m != nil {
		return m.L4Protocol
	}
	return L4Protocol_UNDEFINED_L4
}

func (m *ToHost) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *ToHost) GetSocketPath() string {
	if m != nil {
		return m.SocketPath
	}
	return ""
}

func (*ToHost) XXX_MessageName() string {
	return "vpp.punt.ToHost"
}
func init() {
	proto.RegisterEnum("vpp.punt.L3Protocol", L3Protocol_name, L3Protocol_value)
	proto.RegisterEnum("vpp.punt.L4Protocol", L4Protocol_name, L4Protocol_value)
	proto.RegisterType((*IPRedirect)(nil), "vpp.punt.IPRedirect")
	proto.RegisterType((*ToHost)(nil), "vpp.punt.ToHost")
}

func init() { proto.RegisterFile("models/vpp/punt/punt.proto", fileDescriptor_c6553556423e0b80) }

var fileDescriptor_c6553556423e0b80 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6e, 0xa2, 0x50,
	0x14, 0xc6, 0x45, 0x19, 0x75, 0x8e, 0xce, 0x84, 0xb9, 0x99, 0x05, 0xe3, 0x62, 0x6a, 0x5d, 0x19,
	0x13, 0xa1, 0x29, 0xb4, 0x69, 0x62, 0xd2, 0xa4, 0xad, 0x36, 0x92, 0x10, 0x43, 0x88, 0x6e, 0xba,
	0x21, 0x88, 0x57, 0x20, 0x45, 0xee, 0x0d, 0x5e, 0x89, 0x0f, 0xd4, 0x5d, 0x5f, 0xa4, 0xef, 0xd1,
	0x17, 0x69, 0xb8, 0xf8, 0x87, 0xe8, 0xa6, 0x1b, 0xf2, 0x7d, 0xe7, 0x7c, 0x3f, 0xf2, 0x05, 0x0e,
	0xb4, 0x56, 0x64, 0x81, 0xa3, 0xb5, 0x9a, 0x52, 0xaa, 0xd2, 0x4d, 0xcc, 0xf8, 0x43, 0xa1, 0x09,
	0x61, 0x04, 0xd5, 0x53, 0x4a, 0x95, 0xcc, 0xb7, 0xfa, 0x7e, 0xc8, 0x82, 0xcd, 0x5c, 0xf1, 0xc8,
	0x4a, 0xf5, 0x89, 0x4f, 0x54, 0x1e, 0x98, 0x6f, 0x96, 0xdc, 0x71, 0xc3, 0x55, 0x0e, 0x76, 0xde,
	0x04, 0x00, 0xc3, 0xb2, 0xf1, 0x22, 0x4c, 0xb0, 0xc7, 0xd0, 0x0d, 0x34, 0x22, 0xcd, 0xe1, 0x2b,
	0x8f, 0x44, 0xb2, 0xd0, 0x16, 0xba, 0xbf, 0xaf, 0xff, 0x2a, 0xfb, 0xb7, 0x2b, 0xa6, 0x66, 0xed,
	0x76, 0x36, 0x44, 0x07, 0x8d, 0x2e, 0xa1, 0x99, 0x6c, 0x9d, 0x30, 0x66, 0x38, 0x59, 0xba, 0x1e,
	0x96, 0xcb, 0x6d, 0xa1, 0xfb, 0xd3, 0x6e, 0x24, 0x5b, 0x63, 0x3f, 0xca, 0x22, 0xac, 0x18, 0xa9,
	0xe4, 0x11, 0x56, 0x88, 0xfc, 0x83, 0x7a, 0x8c, 0xb7, 0xcc, 0x09, 0x08, 0x95, 0x45, 0xbe, 0xae,
	0x65, 0x7e, 0x4c, 0x68, 0xe7, 0x5d, 0x80, 0xea, 0x94, 0x8c, 0xc9, 0xfa, 0xac, 0x62, 0xf9, 0x9b,
	0x15, 0x33, 0x4c, 0x3f, 0x62, 0x95, 0x33, 0x4c, 0x2f, 0x60, 0x07, 0x8d, 0x10, 0x88, 0x94, 0x24,
	0x8c, 0xf7, 0xf9, 0x65, 0x73, 0x8d, 0x2e, 0xa0, 0xb1, 0x26, 0xde, 0x2b, 0x66, 0x0e, 0x75, 0x59,
	0x20, 0xff, 0xe0, 0x55, 0x21, 0x1f, 0x59, 0x2e, 0x0b, 0x7a, 0x03, 0x80, 0x63, 0x0b, 0x24, 0x41,
	0x73, 0x36, 0x19, 0x8e, 0x9e, 0x8d, 0xc9, 0x68, 0xe8, 0x98, 0x9a, 0x54, 0x42, 0x75, 0x10, 0x0d,
	0x2b, 0xd5, 0x25, 0x71, 0xa7, 0x6e, 0xa5, 0x2a, 0xaa, 0x41, 0xe5, 0xc1, 0x34, 0x25, 0xe8, 0x5d,
	0x01, 0x1c, 0xbb, 0x9c, 0xc0, 0xba, 0x54, 0xca, 0x82, 0xd3, 0x27, 0x2b, 0x27, 0x66, 0x43, 0x4b,
	0xfa, 0xf3, 0x78, 0xff, 0xf1, 0xf9, 0x5f, 0x78, 0xb9, 0x2b, 0xfc, 0xf8, 0x28, 0xf4, 0x5d, 0x46,
	0xb2, 0x4b, 0xe9, 0xbb, 0x3e, 0x8e, 0x99, 0xea, 0xd2, 0x50, 0x3d, 0x39, 0x9f, 0x41, 0x4a, 0xa9,
	0x93, 0x89, 0x79, 0x95, 0x7f, 0x15, 0xed, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x11, 0xab, 0x16, 0x1e,
	0x61, 0x02, 0x00, 0x00,
}
