// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb.proto

It has these top-level messages:
	Message
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

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

type Message struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "pb.Message")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Debug service

type DebugClient interface {
	FetchMessages(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (Debug_FetchMessagesClient, error)
}

type debugClient struct {
	cc *grpc.ClientConn
}

func NewDebugClient(cc *grpc.ClientConn) DebugClient {
	return &debugClient{cc}
}

func (c *debugClient) FetchMessages(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (Debug_FetchMessagesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Debug_serviceDesc.Streams[0], c.cc, "/pb.Debug/FetchMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &debugFetchMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Debug_FetchMessagesClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type debugFetchMessagesClient struct {
	grpc.ClientStream
}

func (x *debugFetchMessagesClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Debug service

type DebugServer interface {
	FetchMessages(*google_protobuf.Empty, Debug_FetchMessagesServer) error
}

func RegisterDebugServer(s *grpc.Server, srv DebugServer) {
	s.RegisterService(&_Debug_serviceDesc, srv)
}

func _Debug_FetchMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(google_protobuf.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DebugServer).FetchMessages(m, &debugFetchMessagesServer{stream})
}

type Debug_FetchMessagesServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type debugFetchMessagesServer struct {
	grpc.ServerStream
}

func (x *debugFetchMessagesServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

var _Debug_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Debug",
	HandlerType: (*DebugServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FetchMessages",
			Handler:       _Debug_FetchMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pb.proto",
}

func init() { proto.RegisterFile("pb.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x28, 0x48, 0xd2, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0xe2, 0x4e, 0xcd, 0x2d, 0x28, 0xa9, 0x84,
	0x08, 0x28, 0x49, 0x72, 0xb1, 0xfb, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x0a, 0xf1, 0x71, 0x31,
	0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x31, 0x65, 0xa6, 0x18, 0x39, 0x72, 0xb1,
	0xba, 0xa4, 0x26, 0x95, 0xa6, 0x0b, 0x59, 0x70, 0xf1, 0xba, 0xa5, 0x96, 0x24, 0x67, 0x40, 0x15,
	0x16, 0x0b, 0x89, 0xe9, 0xa5, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0x42, 0xcc, 0x48, 0x2a, 0x4d, 0xd3,
	0x73, 0x05, 0x19, 0x29, 0xc5, 0xad, 0x57, 0x90, 0xa4, 0x07, 0x55, 0xa5, 0xc4, 0x60, 0xc0, 0x98,
	0xc4, 0x06, 0x56, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xdf, 0xe2, 0xc3, 0x81, 0x00,
	0x00, 0x00,
}
