// Code generated by protoc-gen-go. DO NOT EDIT.
// source: time.proto

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:

	time.proto

It has these top-level messages:

	TimeRequest
	Time
*/
package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// TimeRequest requests a time
type TimeRequest struct {
	Interval int32 `protobuf:"varint,1,opt,name=interval" json:"interval,omitempty"`
}

func (m *TimeRequest) Reset()                    { *m = TimeRequest{} }
func (m *TimeRequest) String() string            { return proto.CompactTextString(m) }
func (*TimeRequest) ProtoMessage()               {}
func (*TimeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TimeRequest) GetInterval() int32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

// Time represents a time value
type Time struct {
	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Time) Reset()                    { *m = Time{} }
func (m *Time) String() string            { return proto.CompactTextString(m) }
func (*Time) ProtoMessage()               {}
func (*Time) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Time) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*TimeRequest)(nil), "protobuf.TimeRequest")
	proto.RegisterType((*Time)(nil), "protobuf.Time")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TimeService service

type TimeServiceClient interface {
	// GetTimeStream returns a stream of time messages at specified millis intervals
	GetTimeStream(ctx context.Context, in *TimeRequest, opts ...grpc.CallOption) (TimeService_GetTimeStreamClient, error)
}

type timeServiceClient struct {
	cc *grpc.ClientConn
}

func NewTimeServiceClient(cc *grpc.ClientConn) TimeServiceClient {
	return &timeServiceClient{cc}
}

func (c *timeServiceClient) GetTimeStream(ctx context.Context, in *TimeRequest, opts ...grpc.CallOption) (TimeService_GetTimeStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TimeService_serviceDesc.Streams[0], c.cc, "/protobuf.TimeService/GetTimeStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &timeServiceGetTimeStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TimeService_GetTimeStreamClient interface {
	Recv() (*Time, error)
	grpc.ClientStream
}

type timeServiceGetTimeStreamClient struct {
	grpc.ClientStream
}

func (x *timeServiceGetTimeStreamClient) Recv() (*Time, error) {
	m := new(Time)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TimeService service

type TimeServiceServer interface {
	// GetTimeStream returns a stream of time messages at specified millis intervals
	GetTimeStream(*TimeRequest, TimeService_GetTimeStreamServer) error
}

func RegisterTimeServiceServer(s *grpc.Server, srv TimeServiceServer) {
	s.RegisterService(&_TimeService_serviceDesc, srv)
}

func _TimeService_GetTimeStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TimeServiceServer).GetTimeStream(m, &timeServiceGetTimeStreamServer{stream})
}

type TimeService_GetTimeStreamServer interface {
	Send(*Time) error
	grpc.ServerStream
}

type timeServiceGetTimeStreamServer struct {
	grpc.ServerStream
}

func (x *timeServiceGetTimeStreamServer) Send(m *Time) error {
	return x.ServerStream.SendMsg(m)
}

var _TimeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.TimeService",
	HandlerType: (*TimeServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTimeStream",
			Handler:       _TimeService_GetTimeStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "time.proto",
}

func init() { proto.RegisterFile("time.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xc9, 0xcc, 0x4d,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x49, 0xa5, 0x69, 0x4a, 0x9a, 0x5c,
	0xdc, 0x21, 0x99, 0xb9, 0xa9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x52, 0x5c, 0x1c,
	0x99, 0x79, 0x25, 0xa9, 0x45, 0x65, 0x89, 0x39, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x70,
	0xbe, 0x92, 0x0c, 0x17, 0x0b, 0x48, 0xa9, 0x90, 0x08, 0x17, 0x6b, 0x59, 0x62, 0x4e, 0x69, 0x2a,
	0x58, 0x01, 0x4f, 0x10, 0x84, 0x63, 0xe4, 0x09, 0x31, 0x28, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39,
	0x55, 0xc8, 0x8a, 0x8b, 0xd7, 0x3d, 0xb5, 0x04, 0x2c, 0x52, 0x52, 0x94, 0x9a, 0x98, 0x2b, 0x24,
	0xaa, 0x07, 0xb3, 0x53, 0x0f, 0xc9, 0x42, 0x29, 0x3e, 0x54, 0x61, 0x25, 0x06, 0x03, 0xc6, 0x24,
	0x36, 0xb0, 0x90, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x43, 0xb4, 0x81, 0xb2, 0x00, 0x00,
	0x00,
}
