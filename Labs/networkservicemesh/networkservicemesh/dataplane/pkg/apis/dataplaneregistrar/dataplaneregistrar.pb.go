// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dataplaneregistrar.proto

package dataplaneregistrar

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

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

// DataplaneRegistrationRequest is sent by the dataplane to NSM
// to advertise itself and inform NSM about the location of the dataplane socket
// and its initially supported parameters.
type DataplaneRegistrationRequest struct {
	DataplaneName        string   `protobuf:"bytes,1,opt,name=dataplane_name,json=dataplaneName,proto3" json:"dataplane_name,omitempty"`
	DataplaneSocket      string   `protobuf:"bytes,2,opt,name=dataplane_socket,json=dataplaneSocket,proto3" json:"dataplane_socket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataplaneRegistrationRequest) Reset()         { *m = DataplaneRegistrationRequest{} }
func (m *DataplaneRegistrationRequest) String() string { return proto.CompactTextString(m) }
func (*DataplaneRegistrationRequest) ProtoMessage()    {}
func (*DataplaneRegistrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataplaneregistrar_10924475267bfc15, []int{0}
}
func (m *DataplaneRegistrationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataplaneRegistrationRequest.Unmarshal(m, b)
}
func (m *DataplaneRegistrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataplaneRegistrationRequest.Marshal(b, m, deterministic)
}
func (dst *DataplaneRegistrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataplaneRegistrationRequest.Merge(dst, src)
}
func (m *DataplaneRegistrationRequest) XXX_Size() int {
	return xxx_messageInfo_DataplaneRegistrationRequest.Size(m)
}
func (m *DataplaneRegistrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataplaneRegistrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataplaneRegistrationRequest proto.InternalMessageInfo

func (m *DataplaneRegistrationRequest) GetDataplaneName() string {
	if m != nil {
		return m.DataplaneName
	}
	return ""
}

func (m *DataplaneRegistrationRequest) GetDataplaneSocket() string {
	if m != nil {
		return m.DataplaneSocket
	}
	return ""
}

type DataplaneRegistrationReply struct {
	Registered           bool     `protobuf:"varint,1,opt,name=registered,proto3" json:"registered,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataplaneRegistrationReply) Reset()         { *m = DataplaneRegistrationReply{} }
func (m *DataplaneRegistrationReply) String() string { return proto.CompactTextString(m) }
func (*DataplaneRegistrationReply) ProtoMessage()    {}
func (*DataplaneRegistrationReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataplaneregistrar_10924475267bfc15, []int{1}
}
func (m *DataplaneRegistrationReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataplaneRegistrationReply.Unmarshal(m, b)
}
func (m *DataplaneRegistrationReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataplaneRegistrationReply.Marshal(b, m, deterministic)
}
func (dst *DataplaneRegistrationReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataplaneRegistrationReply.Merge(dst, src)
}
func (m *DataplaneRegistrationReply) XXX_Size() int {
	return xxx_messageInfo_DataplaneRegistrationReply.Size(m)
}
func (m *DataplaneRegistrationReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DataplaneRegistrationReply.DiscardUnknown(m)
}

var xxx_messageInfo_DataplaneRegistrationReply proto.InternalMessageInfo

func (m *DataplaneRegistrationReply) GetRegistered() bool {
	if m != nil {
		return m.Registered
	}
	return false
}

// DataplaneUnRegistrationRequest is sent by the dataplane to NSM
// to remove itself from the list of available dataplanes.
type DataplaneUnRegistrationRequest struct {
	DataplaneName        string   `protobuf:"bytes,1,opt,name=dataplane_name,json=dataplaneName,proto3" json:"dataplane_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataplaneUnRegistrationRequest) Reset()         { *m = DataplaneUnRegistrationRequest{} }
func (m *DataplaneUnRegistrationRequest) String() string { return proto.CompactTextString(m) }
func (*DataplaneUnRegistrationRequest) ProtoMessage()    {}
func (*DataplaneUnRegistrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataplaneregistrar_10924475267bfc15, []int{2}
}
func (m *DataplaneUnRegistrationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataplaneUnRegistrationRequest.Unmarshal(m, b)
}
func (m *DataplaneUnRegistrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataplaneUnRegistrationRequest.Marshal(b, m, deterministic)
}
func (dst *DataplaneUnRegistrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataplaneUnRegistrationRequest.Merge(dst, src)
}
func (m *DataplaneUnRegistrationRequest) XXX_Size() int {
	return xxx_messageInfo_DataplaneUnRegistrationRequest.Size(m)
}
func (m *DataplaneUnRegistrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataplaneUnRegistrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataplaneUnRegistrationRequest proto.InternalMessageInfo

func (m *DataplaneUnRegistrationRequest) GetDataplaneName() string {
	if m != nil {
		return m.DataplaneName
	}
	return ""
}

type DataplaneUnRegistrationReply struct {
	UnRegistered         bool     `protobuf:"varint,1,opt,name=un_registered,json=unRegistered,proto3" json:"un_registered,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataplaneUnRegistrationReply) Reset()         { *m = DataplaneUnRegistrationReply{} }
func (m *DataplaneUnRegistrationReply) String() string { return proto.CompactTextString(m) }
func (*DataplaneUnRegistrationReply) ProtoMessage()    {}
func (*DataplaneUnRegistrationReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataplaneregistrar_10924475267bfc15, []int{3}
}
func (m *DataplaneUnRegistrationReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataplaneUnRegistrationReply.Unmarshal(m, b)
}
func (m *DataplaneUnRegistrationReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataplaneUnRegistrationReply.Marshal(b, m, deterministic)
}
func (dst *DataplaneUnRegistrationReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataplaneUnRegistrationReply.Merge(dst, src)
}
func (m *DataplaneUnRegistrationReply) XXX_Size() int {
	return xxx_messageInfo_DataplaneUnRegistrationReply.Size(m)
}
func (m *DataplaneUnRegistrationReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DataplaneUnRegistrationReply.DiscardUnknown(m)
}

var xxx_messageInfo_DataplaneUnRegistrationReply proto.InternalMessageInfo

func (m *DataplaneUnRegistrationReply) GetUnRegistered() bool {
	if m != nil {
		return m.UnRegistered
	}
	return false
}

func init() {
	proto.RegisterType((*DataplaneRegistrationRequest)(nil), "dataplaneregistrar.DataplaneRegistrationRequest")
	proto.RegisterType((*DataplaneRegistrationReply)(nil), "dataplaneregistrar.DataplaneRegistrationReply")
	proto.RegisterType((*DataplaneUnRegistrationRequest)(nil), "dataplaneregistrar.DataplaneUnRegistrationRequest")
	proto.RegisterType((*DataplaneUnRegistrationReply)(nil), "dataplaneregistrar.DataplaneUnRegistrationReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataplaneRegistrationClient is the client API for DataplaneRegistration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataplaneRegistrationClient interface {
	RequestDataplaneRegistration(ctx context.Context, in *DataplaneRegistrationRequest, opts ...grpc.CallOption) (*DataplaneRegistrationReply, error)
	// RequestLiveness is a stream initiated by NSM to inform the dataplane that NSM is still alive and
	// no re-registration is required. Detection a failure on this "channel" will mean
	// that NSM is gone and the dataplane needs to start re-registration logic.
	RequestLiveness(ctx context.Context, opts ...grpc.CallOption) (DataplaneRegistration_RequestLivenessClient, error)
}

type dataplaneRegistrationClient struct {
	cc *grpc.ClientConn
}

func NewDataplaneRegistrationClient(cc *grpc.ClientConn) DataplaneRegistrationClient {
	return &dataplaneRegistrationClient{cc}
}

func (c *dataplaneRegistrationClient) RequestDataplaneRegistration(ctx context.Context, in *DataplaneRegistrationRequest, opts ...grpc.CallOption) (*DataplaneRegistrationReply, error) {
	out := new(DataplaneRegistrationReply)
	err := c.cc.Invoke(ctx, "/dataplaneregistrar.DataplaneRegistration/RequestDataplaneRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataplaneRegistrationClient) RequestLiveness(ctx context.Context, opts ...grpc.CallOption) (DataplaneRegistration_RequestLivenessClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DataplaneRegistration_serviceDesc.Streams[0], "/dataplaneregistrar.DataplaneRegistration/RequestLiveness", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataplaneRegistrationRequestLivenessClient{stream}
	return x, nil
}

type DataplaneRegistration_RequestLivenessClient interface {
	Send(*empty.Empty) error
	Recv() (*empty.Empty, error)
	grpc.ClientStream
}

type dataplaneRegistrationRequestLivenessClient struct {
	grpc.ClientStream
}

func (x *dataplaneRegistrationRequestLivenessClient) Send(m *empty.Empty) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataplaneRegistrationRequestLivenessClient) Recv() (*empty.Empty, error) {
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataplaneRegistrationServer is the server API for DataplaneRegistration service.
type DataplaneRegistrationServer interface {
	RequestDataplaneRegistration(context.Context, *DataplaneRegistrationRequest) (*DataplaneRegistrationReply, error)
	// RequestLiveness is a stream initiated by NSM to inform the dataplane that NSM is still alive and
	// no re-registration is required. Detection a failure on this "channel" will mean
	// that NSM is gone and the dataplane needs to start re-registration logic.
	RequestLiveness(DataplaneRegistration_RequestLivenessServer) error
}

func RegisterDataplaneRegistrationServer(s *grpc.Server, srv DataplaneRegistrationServer) {
	s.RegisterService(&_DataplaneRegistration_serviceDesc, srv)
}

func _DataplaneRegistration_RequestDataplaneRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataplaneRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataplaneRegistrationServer).RequestDataplaneRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dataplaneregistrar.DataplaneRegistration/RequestDataplaneRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataplaneRegistrationServer).RequestDataplaneRegistration(ctx, req.(*DataplaneRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataplaneRegistration_RequestLiveness_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataplaneRegistrationServer).RequestLiveness(&dataplaneRegistrationRequestLivenessServer{stream})
}

type DataplaneRegistration_RequestLivenessServer interface {
	Send(*empty.Empty) error
	Recv() (*empty.Empty, error)
	grpc.ServerStream
}

type dataplaneRegistrationRequestLivenessServer struct {
	grpc.ServerStream
}

func (x *dataplaneRegistrationRequestLivenessServer) Send(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataplaneRegistrationRequestLivenessServer) Recv() (*empty.Empty, error) {
	m := new(empty.Empty)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DataplaneRegistration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dataplaneregistrar.DataplaneRegistration",
	HandlerType: (*DataplaneRegistrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestDataplaneRegistration",
			Handler:    _DataplaneRegistration_RequestDataplaneRegistration_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RequestLiveness",
			Handler:       _DataplaneRegistration_RequestLiveness_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "dataplaneregistrar.proto",
}

// DataplaneUnRegistrationClient is the client API for DataplaneUnRegistration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataplaneUnRegistrationClient interface {
	RequestDataplaneUnRegistration(ctx context.Context, in *DataplaneUnRegistrationRequest, opts ...grpc.CallOption) (*DataplaneUnRegistrationReply, error)
}

type dataplaneUnRegistrationClient struct {
	cc *grpc.ClientConn
}

func NewDataplaneUnRegistrationClient(cc *grpc.ClientConn) DataplaneUnRegistrationClient {
	return &dataplaneUnRegistrationClient{cc}
}

func (c *dataplaneUnRegistrationClient) RequestDataplaneUnRegistration(ctx context.Context, in *DataplaneUnRegistrationRequest, opts ...grpc.CallOption) (*DataplaneUnRegistrationReply, error) {
	out := new(DataplaneUnRegistrationReply)
	err := c.cc.Invoke(ctx, "/dataplaneregistrar.DataplaneUnRegistration/RequestDataplaneUnRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataplaneUnRegistrationServer is the server API for DataplaneUnRegistration service.
type DataplaneUnRegistrationServer interface {
	RequestDataplaneUnRegistration(context.Context, *DataplaneUnRegistrationRequest) (*DataplaneUnRegistrationReply, error)
}

func RegisterDataplaneUnRegistrationServer(s *grpc.Server, srv DataplaneUnRegistrationServer) {
	s.RegisterService(&_DataplaneUnRegistration_serviceDesc, srv)
}

func _DataplaneUnRegistration_RequestDataplaneUnRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataplaneUnRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataplaneUnRegistrationServer).RequestDataplaneUnRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dataplaneregistrar.DataplaneUnRegistration/RequestDataplaneUnRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataplaneUnRegistrationServer).RequestDataplaneUnRegistration(ctx, req.(*DataplaneUnRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataplaneUnRegistration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dataplaneregistrar.DataplaneUnRegistration",
	HandlerType: (*DataplaneUnRegistrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestDataplaneUnRegistration",
			Handler:    _DataplaneUnRegistration_RequestDataplaneUnRegistration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dataplaneregistrar.proto",
}

func init() {
	proto.RegisterFile("dataplaneregistrar.proto", fileDescriptor_dataplaneregistrar_10924475267bfc15)
}

var fileDescriptor_dataplaneregistrar_10924475267bfc15 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x48, 0x49, 0x2c, 0x49,
	0x2c, 0xc8, 0x49, 0xcc, 0x4b, 0x2d, 0x4a, 0x4d, 0xcf, 0x2c, 0x2e, 0x29, 0x4a, 0x2c, 0xd2, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc2, 0x94, 0x91, 0xb2, 0x4a, 0xcf, 0x2c, 0xc9, 0x28, 0x4d,
	0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x07, 0xca, 0xa5, 0xeb, 0x83, 0x15, 0x27, 0x95, 0xa6,
	0x39, 0x94, 0x19, 0xea, 0x19, 0xeb, 0x19, 0xea, 0x17, 0x94, 0x54, 0x16, 0xa4, 0x16, 0xeb, 0xa7,
	0xe6, 0x02, 0x19, 0x10, 0x12, 0x62, 0x9e, 0x52, 0x01, 0x97, 0x8c, 0x0b, 0xcc, 0xc4, 0x20, 0xa8,
	0x89, 0x25, 0x99, 0xf9, 0x79, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0xaa, 0x5c, 0x7c,
	0x70, 0x1b, 0xe3, 0xf3, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x78, 0xe1,
	0xa2, 0x7e, 0x40, 0x41, 0x21, 0x4d, 0x2e, 0x01, 0x84, 0xb2, 0xe2, 0xfc, 0xe4, 0xec, 0xd4, 0x12,
	0x09, 0x26, 0xb0, 0x42, 0x7e, 0xb8, 0x78, 0x30, 0x58, 0x58, 0xc9, 0x86, 0x4b, 0x0a, 0x87, 0x8d,
	0x05, 0x39, 0x95, 0x42, 0x72, 0x5c, 0x5c, 0x10, 0x8f, 0x01, 0xfd, 0x97, 0x02, 0xb6, 0x8b, 0x23,
	0x08, 0x49, 0x44, 0xc9, 0x9d, 0x4b, 0x0e, 0xae, 0x3b, 0x34, 0x8f, 0x7c, 0x17, 0x2b, 0x39, 0x23,
	0x79, 0x1c, 0xdd, 0x20, 0x90, 0x43, 0x94, 0xb9, 0x78, 0x4b, 0xf3, 0xe2, 0x31, 0xdc, 0xc2, 0x53,
	0x0a, 0x55, 0x0b, 0x12, 0x33, 0x7a, 0xc8, 0xc8, 0x25, 0x8a, 0xd5, 0x33, 0x42, 0x0d, 0x8c, 0x5c,
	0x32, 0x50, 0x17, 0x61, 0x57, 0x60, 0xa0, 0x87, 0x25, 0x8e, 0xf1, 0x45, 0x85, 0x94, 0x1e, 0x09,
	0x3a, 0x40, 0x3e, 0x70, 0xe5, 0xe2, 0x87, 0x6a, 0xf5, 0xc9, 0x2c, 0x4b, 0xcd, 0x4b, 0x2d, 0x2e,
	0x16, 0x12, 0xd3, 0x4b, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x83, 0xa5, 0x0f, 0x3d, 0x57, 0x50,
	0x5a, 0x90, 0xc2, 0x21, 0xae, 0xc1, 0x68, 0xc0, 0x68, 0xb4, 0x88, 0x91, 0x4b, 0x1c, 0x47, 0x48,
	0x09, 0xb5, 0x31, 0x72, 0xc9, 0xa1, 0xfb, 0x12, 0x4d, 0x89, 0x11, 0x5e, 0x57, 0x63, 0x8d, 0x42,
	0x29, 0x03, 0x92, 0xf4, 0x00, 0xfd, 0x9a, 0xc4, 0x06, 0x76, 0xb8, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0xfa, 0xfd, 0xea, 0x6d, 0x39, 0x03, 0x00, 0x00,
}
