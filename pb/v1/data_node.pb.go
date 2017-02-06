// Code generated by protoc-gen-go.
// source: data_node.proto
// DO NOT EDIT!

/*
Package loggrebutterfly is a generated protocol buffer package.

It is generated from these files:
	data_node.proto
	master.proto

It has these top-level messages:
	WriteInfo
	WriteResponse
	ReadInfo
	ReadData
	RoutesInfo
	RoutesResponse
	RouteInfo
*/
package loggrebutterfly

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

type WriteInfo struct {
	Payload []byte `protobuf:"bytes,1,opt,name=Payload,json=payload,proto3" json:"Payload,omitempty"`
}

func (m *WriteInfo) Reset()                    { *m = WriteInfo{} }
func (m *WriteInfo) String() string            { return proto.CompactTextString(m) }
func (*WriteInfo) ProtoMessage()               {}
func (*WriteInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *WriteInfo) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type WriteResponse struct {
}

func (m *WriteResponse) Reset()                    { *m = WriteResponse{} }
func (m *WriteResponse) String() string            { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()               {}
func (*WriteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ReadInfo struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ReadInfo) Reset()                    { *m = ReadInfo{} }
func (m *ReadInfo) String() string            { return proto.CompactTextString(m) }
func (*ReadInfo) ProtoMessage()               {}
func (*ReadInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReadInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ReadData struct {
	Payload []byte `protobuf:"bytes,1,opt,name=Payload,json=payload,proto3" json:"Payload,omitempty"`
}

func (m *ReadData) Reset()                    { *m = ReadData{} }
func (m *ReadData) String() string            { return proto.CompactTextString(m) }
func (*ReadData) ProtoMessage()               {}
func (*ReadData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadData) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*WriteInfo)(nil), "loggrebutterfly.WriteInfo")
	proto.RegisterType((*WriteResponse)(nil), "loggrebutterfly.WriteResponse")
	proto.RegisterType((*ReadInfo)(nil), "loggrebutterfly.ReadInfo")
	proto.RegisterType((*ReadData)(nil), "loggrebutterfly.ReadData")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DataNode service

type DataNodeClient interface {
	Write(ctx context.Context, opts ...grpc.CallOption) (DataNode_WriteClient, error)
	Read(ctx context.Context, in *ReadInfo, opts ...grpc.CallOption) (DataNode_ReadClient, error)
}

type dataNodeClient struct {
	cc *grpc.ClientConn
}

func NewDataNodeClient(cc *grpc.ClientConn) DataNodeClient {
	return &dataNodeClient{cc}
}

func (c *dataNodeClient) Write(ctx context.Context, opts ...grpc.CallOption) (DataNode_WriteClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DataNode_serviceDesc.Streams[0], c.cc, "/loggrebutterfly.DataNode/Write", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataNodeWriteClient{stream}
	return x, nil
}

type DataNode_WriteClient interface {
	Send(*WriteInfo) error
	CloseAndRecv() (*WriteResponse, error)
	grpc.ClientStream
}

type dataNodeWriteClient struct {
	grpc.ClientStream
}

func (x *dataNodeWriteClient) Send(m *WriteInfo) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataNodeWriteClient) CloseAndRecv() (*WriteResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(WriteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataNodeClient) Read(ctx context.Context, in *ReadInfo, opts ...grpc.CallOption) (DataNode_ReadClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DataNode_serviceDesc.Streams[1], c.cc, "/loggrebutterfly.DataNode/Read", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataNodeReadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataNode_ReadClient interface {
	Recv() (*ReadData, error)
	grpc.ClientStream
}

type dataNodeReadClient struct {
	grpc.ClientStream
}

func (x *dataNodeReadClient) Recv() (*ReadData, error) {
	m := new(ReadData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for DataNode service

type DataNodeServer interface {
	Write(DataNode_WriteServer) error
	Read(*ReadInfo, DataNode_ReadServer) error
}

func RegisterDataNodeServer(s *grpc.Server, srv DataNodeServer) {
	s.RegisterService(&_DataNode_serviceDesc, srv)
}

func _DataNode_Write_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataNodeServer).Write(&dataNodeWriteServer{stream})
}

type DataNode_WriteServer interface {
	SendAndClose(*WriteResponse) error
	Recv() (*WriteInfo, error)
	grpc.ServerStream
}

type dataNodeWriteServer struct {
	grpc.ServerStream
}

func (x *dataNodeWriteServer) SendAndClose(m *WriteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataNodeWriteServer) Recv() (*WriteInfo, error) {
	m := new(WriteInfo)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DataNode_Read_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataNodeServer).Read(m, &dataNodeReadServer{stream})
}

type DataNode_ReadServer interface {
	Send(*ReadData) error
	grpc.ServerStream
}

type dataNodeReadServer struct {
	grpc.ServerStream
}

func (x *dataNodeReadServer) Send(m *ReadData) error {
	return x.ServerStream.SendMsg(m)
}

var _DataNode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "loggrebutterfly.DataNode",
	HandlerType: (*DataNodeServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Write",
			Handler:       _DataNode_Write_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Read",
			Handler:       _DataNode_Read_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "data_node.proto",
}

func init() { proto.RegisterFile("data_node.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x49, 0x2c, 0x49,
	0x8c, 0xcf, 0xcb, 0x4f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcf, 0xc9, 0x4f,
	0x4f, 0x2f, 0x4a, 0x4d, 0x2a, 0x2d, 0x29, 0x49, 0x2d, 0x4a, 0xcb, 0xa9, 0x54, 0x52, 0xe5, 0xe2,
	0x0c, 0x2f, 0xca, 0x2c, 0x49, 0xf5, 0xcc, 0x4b, 0xcb, 0x17, 0x92, 0xe0, 0x62, 0x0f, 0x48, 0xac,
	0xcc, 0xc9, 0x4f, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x62, 0x2f, 0x80, 0x70, 0x95,
	0xf8, 0xb9, 0x78, 0xc1, 0xca, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x95, 0xe4, 0xb8,
	0x38, 0x82, 0x52, 0x13, 0x53, 0xc0, 0xda, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0xc1, 0x7a,
	0x38, 0x83, 0xc0, 0x6c, 0x25, 0x15, 0x88, 0xbc, 0x4b, 0x62, 0x49, 0x22, 0x6e, 0x63, 0x8d, 0xa6,
	0x32, 0x72, 0x71, 0x80, 0x94, 0xf8, 0xe5, 0xa7, 0xa4, 0x0a, 0xb9, 0x73, 0xb1, 0x82, 0xed, 0x10,
	0x92, 0xd2, 0x43, 0x73, 0xa5, 0x1e, 0xdc, 0x89, 0x52, 0x72, 0xd8, 0xe5, 0xe0, 0xee, 0x62, 0xd0,
	0x60, 0x14, 0x72, 0xe0, 0x62, 0x01, 0xd9, 0x2d, 0x24, 0x89, 0xa1, 0x16, 0xe6, 0x64, 0x29, 0xec,
	0x52, 0x20, 0xa7, 0x28, 0x31, 0x18, 0x30, 0x26, 0xb1, 0x81, 0x43, 0xcb, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x12, 0xf0, 0xe2, 0xe4, 0x40, 0x01, 0x00, 0x00,
}
