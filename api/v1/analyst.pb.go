// Code generated by protoc-gen-go.
// source: analyst.proto
// DO NOT EDIT!

/*
Package loggrebutterfly is a generated protocol buffer package.

It is generated from these files:
	analyst.proto
	data_node.proto
	master.proto

It has these top-level messages:
	QueryInfo
	TimeRange
	QueryResponse
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
import loggregator "github.com/apoydence/loggrebutterfly/api/loggregator/v2"

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

type QueryInfo struct {
	SourceUuid string     `protobuf:"bytes,1,opt,name=source_uuid,json=sourceUuid" json:"source_uuid,omitempty"`
	TimeRange  *TimeRange `protobuf:"bytes,2,opt,name=timeRange" json:"timeRange,omitempty"`
}

func (m *QueryInfo) Reset()                    { *m = QueryInfo{} }
func (m *QueryInfo) String() string            { return proto.CompactTextString(m) }
func (*QueryInfo) ProtoMessage()               {}
func (*QueryInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *QueryInfo) GetSourceUuid() string {
	if m != nil {
		return m.SourceUuid
	}
	return ""
}

func (m *QueryInfo) GetTimeRange() *TimeRange {
	if m != nil {
		return m.TimeRange
	}
	return nil
}

type TimeRange struct {
	Start int64 `protobuf:"varint,1,opt,name=start" json:"start,omitempty"`
	End   int64 `protobuf:"varint,2,opt,name=end" json:"end,omitempty"`
}

func (m *TimeRange) Reset()                    { *m = TimeRange{} }
func (m *TimeRange) String() string            { return proto.CompactTextString(m) }
func (*TimeRange) ProtoMessage()               {}
func (*TimeRange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TimeRange) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *TimeRange) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

type QueryResponse struct {
	Envelopes []*loggregator.Envelope `protobuf:"bytes,1,rep,name=envelopes" json:"envelopes,omitempty"`
}

func (m *QueryResponse) Reset()                    { *m = QueryResponse{} }
func (m *QueryResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()               {}
func (*QueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *QueryResponse) GetEnvelopes() []*loggregator.Envelope {
	if m != nil {
		return m.Envelopes
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryInfo)(nil), "loggrebutterfly.QueryInfo")
	proto.RegisterType((*TimeRange)(nil), "loggrebutterfly.TimeRange")
	proto.RegisterType((*QueryResponse)(nil), "loggrebutterfly.QueryResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Analyst service

type AnalystClient interface {
	Query(ctx context.Context, in *QueryInfo, opts ...grpc.CallOption) (*QueryResponse, error)
}

type analystClient struct {
	cc *grpc.ClientConn
}

func NewAnalystClient(cc *grpc.ClientConn) AnalystClient {
	return &analystClient{cc}
}

func (c *analystClient) Query(ctx context.Context, in *QueryInfo, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := grpc.Invoke(ctx, "/loggrebutterfly.Analyst/Query", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Analyst service

type AnalystServer interface {
	Query(context.Context, *QueryInfo) (*QueryResponse, error)
}

func RegisterAnalystServer(s *grpc.Server, srv AnalystServer) {
	s.RegisterService(&_Analyst_serviceDesc, srv)
}

func _Analyst_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalystServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loggrebutterfly.Analyst/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalystServer).Query(ctx, req.(*QueryInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Analyst_serviceDesc = grpc.ServiceDesc{
	ServiceName: "loggrebutterfly.Analyst",
	HandlerType: (*AnalystServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _Analyst_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "analyst.proto",
}

func init() { proto.RegisterFile("analyst.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0x5f, 0x4b, 0x84, 0x40,
	0x14, 0xc5, 0x33, 0xd9, 0xc2, 0x2b, 0x4b, 0x31, 0x14, 0x88, 0x0f, 0x25, 0x3e, 0xf9, 0x34, 0x82,
	0xbe, 0xf4, 0x1a, 0xb4, 0x41, 0x6f, 0x35, 0xd4, 0x73, 0x8c, 0x7a, 0x35, 0xc1, 0x9d, 0x91, 0xf9,
	0xb3, 0xe0, 0xb7, 0x8f, 0xc6, 0x75, 0x37, 0xb6, 0xde, 0x66, 0xce, 0x3d, 0x9c, 0xfb, 0xbb, 0x07,
	0xd6, 0x5c, 0xf0, 0x61, 0xd2, 0x86, 0x8e, 0x4a, 0x1a, 0x49, 0xae, 0x06, 0xd9, 0x75, 0x0a, 0x2b,
	0x6b, 0x0c, 0xaa, 0x76, 0x98, 0xe2, 0xe7, 0xae, 0x37, 0x5f, 0xb6, 0xa2, 0xb5, 0xdc, 0xe6, 0x7c,
	0x94, 0x53, 0x83, 0xa2, 0xc6, 0xfc, 0xc4, 0x95, 0xf3, 0xb1, 0xdf, 0x6b, 0x1d, 0x37, 0x52, 0xe5,
	0xbb, 0x22, 0x47, 0xb1, 0xc3, 0x41, 0x8e, 0x38, 0x07, 0xa7, 0x2d, 0x04, 0x6f, 0x16, 0xd5, 0xf4,
	0x22, 0x5a, 0x49, 0xee, 0x21, 0xd4, 0xd2, 0xaa, 0x1a, 0x3f, 0xad, 0xed, 0x9b, 0xc8, 0x4b, 0xbc,
	0x2c, 0x60, 0x30, 0x4b, 0x1f, 0xb6, 0x6f, 0xc8, 0x03, 0x04, 0xa6, 0xdf, 0x22, 0xe3, 0xa2, 0xc3,
	0xe8, 0x3c, 0xf1, 0xb2, 0xb0, 0x88, 0xe9, 0xc9, 0x52, 0xfa, 0xbe, 0x38, 0xd8, 0xd1, 0x9c, 0x96,
	0x10, 0x1c, 0x74, 0x72, 0x03, 0x2b, 0x6d, 0xb8, 0x32, 0x6e, 0x83, 0xcf, 0xe6, 0x0f, 0xb9, 0x06,
	0x1f, 0x45, 0xe3, 0x62, 0x7d, 0xf6, 0xf3, 0x4c, 0x9f, 0x60, 0xed, 0xe0, 0x18, 0xea, 0x51, 0x0a,
	0x8d, 0xa4, 0x84, 0x60, 0xe1, 0xd7, 0x91, 0x97, 0xf8, 0x59, 0x58, 0xdc, 0xd2, 0x5f, 0x07, 0xd2,
	0xcd, 0x7e, 0xca, 0x8e, 0xbe, 0xe2, 0x15, 0x2e, 0x1f, 0xe7, 0x32, 0xc9, 0x06, 0x56, 0x2e, 0x90,
	0xfc, 0xa5, 0x3e, 0xb4, 0x10, 0xdf, 0xfd, 0x3f, 0x5b, 0x20, 0xd2, 0xb3, 0xea, 0xc2, 0x75, 0x57,
	0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x16, 0x74, 0xd0, 0xe3, 0xa5, 0x01, 0x00, 0x00,
}