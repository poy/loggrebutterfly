// Code generated by protoc-gen-go.
// source: master.proto
// DO NOT EDIT!

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

type RoutesInfo struct {
}

func (m *RoutesInfo) Reset()                    { *m = RoutesInfo{} }
func (m *RoutesInfo) String() string            { return proto.CompactTextString(m) }
func (*RoutesInfo) ProtoMessage()               {}
func (*RoutesInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type RoutesResponse struct {
	Routes []*RouteInfo `protobuf:"bytes,1,rep,name=routes" json:"routes,omitempty"`
}

func (m *RoutesResponse) Reset()                    { *m = RoutesResponse{} }
func (m *RoutesResponse) String() string            { return proto.CompactTextString(m) }
func (*RoutesResponse) ProtoMessage()               {}
func (*RoutesResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *RoutesResponse) GetRoutes() []*RouteInfo {
	if m != nil {
		return m.Routes
	}
	return nil
}

type RouteInfo struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Leader string `protobuf:"bytes,2,opt,name=leader" json:"leader,omitempty"`
}

func (m *RouteInfo) Reset()                    { *m = RouteInfo{} }
func (m *RouteInfo) String() string            { return proto.CompactTextString(m) }
func (*RouteInfo) ProtoMessage()               {}
func (*RouteInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *RouteInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteInfo) GetLeader() string {
	if m != nil {
		return m.Leader
	}
	return ""
}

type AnalystsInfo struct {
}

func (m *AnalystsInfo) Reset()                    { *m = AnalystsInfo{} }
func (m *AnalystsInfo) String() string            { return proto.CompactTextString(m) }
func (*AnalystsInfo) ProtoMessage()               {}
func (*AnalystsInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

type AnalystsResponse struct {
	Analysts []*AnalystInfo `protobuf:"bytes,1,rep,name=analysts" json:"analysts,omitempty"`
}

func (m *AnalystsResponse) Reset()                    { *m = AnalystsResponse{} }
func (m *AnalystsResponse) String() string            { return proto.CompactTextString(m) }
func (*AnalystsResponse) ProtoMessage()               {}
func (*AnalystsResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *AnalystsResponse) GetAnalysts() []*AnalystInfo {
	if m != nil {
		return m.Analysts
	}
	return nil
}

type AnalystInfo struct {
	Addr string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
}

func (m *AnalystInfo) Reset()                    { *m = AnalystInfo{} }
func (m *AnalystInfo) String() string            { return proto.CompactTextString(m) }
func (*AnalystInfo) ProtoMessage()               {}
func (*AnalystInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *AnalystInfo) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func init() {
	proto.RegisterType((*RoutesInfo)(nil), "loggrebutterfly.RoutesInfo")
	proto.RegisterType((*RoutesResponse)(nil), "loggrebutterfly.RoutesResponse")
	proto.RegisterType((*RouteInfo)(nil), "loggrebutterfly.RouteInfo")
	proto.RegisterType((*AnalystsInfo)(nil), "loggrebutterfly.AnalystsInfo")
	proto.RegisterType((*AnalystsResponse)(nil), "loggrebutterfly.AnalystsResponse")
	proto.RegisterType((*AnalystInfo)(nil), "loggrebutterfly.AnalystInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Master service

type MasterClient interface {
	Routes(ctx context.Context, in *RoutesInfo, opts ...grpc.CallOption) (*RoutesResponse, error)
	Analysts(ctx context.Context, in *AnalystsInfo, opts ...grpc.CallOption) (*AnalystsResponse, error)
}

type masterClient struct {
	cc *grpc.ClientConn
}

func NewMasterClient(cc *grpc.ClientConn) MasterClient {
	return &masterClient{cc}
}

func (c *masterClient) Routes(ctx context.Context, in *RoutesInfo, opts ...grpc.CallOption) (*RoutesResponse, error) {
	out := new(RoutesResponse)
	err := grpc.Invoke(ctx, "/loggrebutterfly.Master/Routes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) Analysts(ctx context.Context, in *AnalystsInfo, opts ...grpc.CallOption) (*AnalystsResponse, error) {
	out := new(AnalystsResponse)
	err := grpc.Invoke(ctx, "/loggrebutterfly.Master/Analysts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Master service

type MasterServer interface {
	Routes(context.Context, *RoutesInfo) (*RoutesResponse, error)
	Analysts(context.Context, *AnalystsInfo) (*AnalystsResponse, error)
}

func RegisterMasterServer(s *grpc.Server, srv MasterServer) {
	s.RegisterService(&_Master_serviceDesc, srv)
}

func _Master_Routes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoutesInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).Routes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loggrebutterfly.Master/Routes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).Routes(ctx, req.(*RoutesInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_Analysts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnalystsInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).Analysts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loggrebutterfly.Master/Analysts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).Analysts(ctx, req.(*AnalystsInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Master_serviceDesc = grpc.ServiceDesc{
	ServiceName: "loggrebutterfly.Master",
	HandlerType: (*MasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Routes",
			Handler:    _Master_Routes_Handler,
		},
		{
			MethodName: "Analysts",
			Handler:    _Master_Analysts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "master.proto",
}

func init() { proto.RegisterFile("master.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x09, 0x20, 0xab, 0xbd, 0x46, 0x05, 0xdd, 0x80, 0xa2, 0x02, 0xa2, 0xf5, 0xd4, 0x29,
	0x43, 0x18, 0x60, 0x45, 0x62, 0x00, 0x09, 0x18, 0xf2, 0x06, 0xae, 0x72, 0xed, 0x92, 0xda, 0x95,
	0xed, 0x0e, 0x79, 0x1d, 0x9e, 0x14, 0xe5, 0x6c, 0x07, 0x04, 0x64, 0xf3, 0xfd, 0x77, 0xff, 0xaf,
	0xef, 0xce, 0x90, 0xef, 0x95, 0xf3, 0x64, 0xcb, 0x83, 0x35, 0xde, 0xe0, 0x45, 0x6b, 0x76, 0x3b,
	0x4b, 0x9b, 0xa3, 0xf7, 0x64, 0xb7, 0x6d, 0x27, 0x73, 0x80, 0xda, 0x1c, 0x3d, 0xb9, 0x57, 0xbd,
	0x35, 0xf2, 0x19, 0xe6, 0xa1, 0xaa, 0xc9, 0x1d, 0x8c, 0x76, 0x84, 0x15, 0x08, 0xcb, 0x4a, 0x91,
	0x2d, 0xcf, 0xd6, 0xb3, 0x6a, 0x51, 0xfe, 0x4a, 0x28, 0xd9, 0xd0, 0xbb, 0xeb, 0x38, 0x29, 0x1f,
	0x60, 0x3a, 0x88, 0x88, 0x70, 0xae, 0xd5, 0x9e, 0x8a, 0x6c, 0x99, 0xad, 0xa7, 0x35, 0xbf, 0xf1,
	0x0a, 0x44, 0x4b, 0xaa, 0x21, 0x5b, 0x9c, 0xb2, 0x1a, 0x2b, 0x39, 0x87, 0xfc, 0x49, 0xab, 0xb6,
	0x73, 0x3e, 0xe0, 0xbc, 0xc1, 0x65, 0xaa, 0x07, 0xa0, 0x47, 0x98, 0xa8, 0xa8, 0x45, 0xa4, 0x9b,
	0x3f, 0x48, 0xd1, 0xc4, 0x50, 0xc3, 0xb4, 0x5c, 0xc1, 0xec, 0x47, 0xa3, 0x07, 0x53, 0x4d, 0x63,
	0x13, 0x58, 0xff, 0xae, 0x3e, 0x33, 0x10, 0xef, 0x7c, 0x2f, 0x7c, 0x01, 0x11, 0x4e, 0x81, 0xd7,
	0xff, 0xaf, 0xcc, 0x88, 0x8b, 0xbb, 0x91, 0x66, 0xe2, 0x95, 0x27, 0xf8, 0x01, 0x93, 0xb4, 0x05,
	0xde, 0x8e, 0xb1, 0x86, 0xb4, 0xd5, 0x68, 0xfb, 0x3b, 0x6f, 0x23, 0xf8, 0x2b, 0xef, 0xbf, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x1b, 0xcb, 0xbf, 0x48, 0xda, 0x01, 0x00, 0x00,
}
