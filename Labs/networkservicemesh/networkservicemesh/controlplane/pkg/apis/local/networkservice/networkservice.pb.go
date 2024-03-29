// Code generated by protoc-gen-go. DO NOT EDIT.
// source: networkservice.proto

package networkservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import connection "github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/local/connection"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NetworkServiceRequest struct {
	Connection           *connection.Connection  `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	MechanismPreferences []*connection.Mechanism `protobuf:"bytes,2,rep,name=mechanism_preferences,json=mechanismPreferences,proto3" json:"mechanism_preferences,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *NetworkServiceRequest) Reset()         { *m = NetworkServiceRequest{} }
func (m *NetworkServiceRequest) String() string { return proto.CompactTextString(m) }
func (*NetworkServiceRequest) ProtoMessage()    {}
func (*NetworkServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_networkservice_2e42ac5a48e51550, []int{0}
}
func (m *NetworkServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkServiceRequest.Unmarshal(m, b)
}
func (m *NetworkServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkServiceRequest.Marshal(b, m, deterministic)
}
func (dst *NetworkServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkServiceRequest.Merge(dst, src)
}
func (m *NetworkServiceRequest) XXX_Size() int {
	return xxx_messageInfo_NetworkServiceRequest.Size(m)
}
func (m *NetworkServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkServiceRequest proto.InternalMessageInfo

func (m *NetworkServiceRequest) GetConnection() *connection.Connection {
	if m != nil {
		return m.Connection
	}
	return nil
}

func (m *NetworkServiceRequest) GetMechanismPreferences() []*connection.Mechanism {
	if m != nil {
		return m.MechanismPreferences
	}
	return nil
}

func init() {
	proto.RegisterType((*NetworkServiceRequest)(nil), "local.networkservice.NetworkServiceRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkServiceClient is the client API for NetworkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkServiceClient interface {
	Request(ctx context.Context, in *NetworkServiceRequest, opts ...grpc.CallOption) (*connection.Connection, error)
	Close(ctx context.Context, in *connection.Connection, opts ...grpc.CallOption) (*empty.Empty, error)
}

type networkServiceClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServiceClient(cc *grpc.ClientConn) NetworkServiceClient {
	return &networkServiceClient{cc}
}

func (c *networkServiceClient) Request(ctx context.Context, in *NetworkServiceRequest, opts ...grpc.CallOption) (*connection.Connection, error) {
	out := new(connection.Connection)
	err := c.cc.Invoke(ctx, "/local.networkservice.NetworkService/Request", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) Close(ctx context.Context, in *connection.Connection, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/local.networkservice.NetworkService/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServiceServer is the server API for NetworkService service.
type NetworkServiceServer interface {
	Request(context.Context, *NetworkServiceRequest) (*connection.Connection, error)
	Close(context.Context, *connection.Connection) (*empty.Empty, error)
}

func RegisterNetworkServiceServer(s *grpc.Server, srv NetworkServiceServer) {
	s.RegisterService(&_NetworkService_serviceDesc, srv)
}

func _NetworkService_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/local.networkservice.NetworkService/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Request(ctx, req.(*NetworkServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(connection.Connection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/local.networkservice.NetworkService/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Close(ctx, req.(*connection.Connection))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "local.networkservice.NetworkService",
	HandlerType: (*NetworkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _NetworkService_Request_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _NetworkService_Close_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "networkservice.proto",
}

func init() {
	proto.RegisterFile("networkservice.proto", fileDescriptor_networkservice_2e42ac5a48e51550)
}

var fileDescriptor_networkservice_2e42ac5a48e51550 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x52, 0x41, 0x4b, 0x33, 0x31,
	0x10, 0x65, 0xbf, 0x0f, 0x15, 0x22, 0x14, 0x09, 0xad, 0x94, 0xd5, 0x43, 0xf1, 0x24, 0x08, 0x09,
	0x6d, 0x6f, 0xa2, 0x20, 0x16, 0x8f, 0x4a, 0xa9, 0x9e, 0xbc, 0xc8, 0x36, 0x4c, 0xb3, 0x4b, 0xb3,
	0x99, 0x98, 0x64, 0x2b, 0xfe, 0x20, 0xf1, 0x6f, 0xba, 0xdd, 0xd8, 0x76, 0x17, 0x97, 0x5e, 0xc2,
	0x90, 0xf7, 0xde, 0xcc, 0x7b, 0x93, 0x90, 0xae, 0x06, 0xff, 0x81, 0x76, 0xe9, 0xc0, 0xae, 0x32,
	0x01, 0xcc, 0x58, 0xf4, 0x48, 0xbb, 0x0a, 0x45, 0xa2, 0x58, 0x13, 0x8b, 0x53, 0x99, 0xf9, 0xb4,
	0x98, 0x33, 0x81, 0x39, 0x6f, 0x42, 0x39, 0xb8, 0xb4, 0xed, 0x4a, 0xa0, 0xf6, 0x16, 0x95, 0x51,
	0x89, 0x06, 0x6e, 0x96, 0x92, 0x27, 0x26, 0x73, 0xbc, 0x6a, 0xbe, 0xc6, 0x34, 0x08, 0x9f, 0xa1,
	0xae, 0x95, 0x61, 0x7e, 0x7c, 0x5d, 0x9b, 0x24, 0xb1, 0x54, 0x4b, 0x5e, 0x01, 0xf3, 0x62, 0x71,
	0xb7, 0x1a, 0xb2, 0x31, 0x1b, 0x72, 0xe3, 0x3f, 0x0d, 0x38, 0x0e, 0x79, 0x59, 0x84, 0x33, 0x68,
	0x2f, 0xbe, 0x23, 0xd2, 0x7b, 0x0a, 0x56, 0x9e, 0x83, 0x95, 0x19, 0xbc, 0x17, 0xe0, 0x3c, 0xbd,
	0x21, 0x64, 0x37, 0xa9, 0x1f, 0x0d, 0xa2, 0xcb, 0xe3, 0xd1, 0x39, 0x0b, 0x51, 0x6b, 0x16, 0x26,
	0xdb, 0x72, 0x56, 0xe3, 0xd3, 0x29, 0xe9, 0xe5, 0x20, 0xd2, 0x44, 0x67, 0x2e, 0x7f, 0x33, 0x16,
	0x16, 0x60, 0x41, 0x0b, 0x70, 0xfd, 0x7f, 0x83, 0xff, 0x65, 0xa3, 0xb3, 0xbf, 0x8d, 0x1e, 0x37,
	0xf4, 0x59, 0x77, 0xab, 0x9c, 0xee, 0x84, 0xa3, 0xaf, 0x88, 0x74, 0x9a, 0x4e, 0xe9, 0x0b, 0x39,
	0xda, 0xb8, 0xbd, 0x62, 0x6d, 0x8f, 0xc0, 0x5a, 0xa3, 0xc5, 0x7b, 0x63, 0xd0, 0x5b, 0x72, 0x30,
	0x51, 0xe8, 0x80, 0xee, 0xa5, 0xc5, 0xa7, 0x4c, 0x22, 0x4a, 0xf5, 0xfb, 0x09, 0xca, 0x5d, 0xb3,
	0x87, 0xf5, 0x5e, 0xef, 0x4f, 0x5e, 0x3b, 0x4d, 0x13, 0xf3, 0xc3, 0x8a, 0x31, 0xfe, 0x09, 0x00,
	0x00, 0xff, 0xff, 0xab, 0x3a, 0xc7, 0xc7, 0x3e, 0x02, 0x00, 0x00,
}
