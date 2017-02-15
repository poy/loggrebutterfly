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
	AnalystsInfo
	AnalystsResponse
	AnalystInfo
*/
package loggrebutterfly

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import loggregator_v2 "github.com/apoydence/loggrebutterfly/api/loggregator/v2"

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
	SourceId string     `protobuf:"bytes,1,opt,name=source_uuid,json=sourceUuid" json:"source_uuid,omitempty"`
	TimeRange  *TimeRange `protobuf:"bytes,2,opt,name=timeRange" json:"timeRange,omitempty"`
}

func (m *QueryInfo) Reset()                    { *m = QueryInfo{} }
func (m *QueryInfo) String() string            { return proto.CompactTextString(m) }
func (*QueryInfo) ProtoMessage()               {}
func (*QueryInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *QueryInfo) GetSourceId() string {
	if m != nil {
		return m.SourceId
	}
	return ""
}

func (m *QueryInfo) GetTimeRange() *TimeRange {
	if m != nil {
		return m.TimeRange
	}
	return nil
}

// [start, end)
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
	Envelopes []*loggregator_v2.Envelope `protobuf:"bytes,1,rep,name=envelopes" json:"envelopes,omitempty"`
}

func (m *QueryResponse) Reset()                    { *m = QueryResponse{} }
func (m *QueryResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()               {}
func (*QueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *QueryResponse) GetEnvelopes() []*loggregator_v2.Envelope {
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
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4b, 0x84, 0x40,
	0x1c, 0xc5, 0x33, 0xd9, 0xc2, 0xaf, 0x2c, 0xc5, 0xd0, 0x41, 0x3c, 0x94, 0x78, 0xf2, 0x34, 0x82,
	0x0b, 0xd1, 0xb5, 0xc3, 0x16, 0xdd, 0x6a, 0xa8, 0x73, 0x8c, 0xfa, 0xd5, 0x04, 0x77, 0x46, 0xe6,
	0x87, 0xe0, 0x7f, 0x1f, 0x8d, 0xeb, 0x6e, 0x6c, 0xdd, 0x66, 0xde, 0x7c, 0x78, 0xef, 0xcd, 0x83,
	0x35, 0x17, 0xbc, 0x9f, 0xb4, 0xa1, 0x83, 0x92, 0x46, 0x92, 0xab, 0x5e, 0xb6, 0xad, 0xc2, 0xd2,
	0x1a, 0x83, 0xaa, 0xe9, 0xa7, 0xf8, 0xa9, 0xed, 0xcc, 0x97, 0x2d, 0x69, 0x25, 0x77, 0x39, 0x1f,
	0xe4, 0x54, 0xa3, 0xa8, 0x30, 0x3f, 0xa1, 0x72, 0x3e, 0x74, 0x7b, 0xad, 0xe5, 0x46, 0xaa, 0x7c,
	0x2c, 0x72, 0x14, 0x23, 0xf6, 0x72, 0xc0, 0xd9, 0x38, 0x6d, 0x20, 0x78, 0xb3, 0xa8, 0xa6, 0x17,
	0xd1, 0x48, 0x72, 0x07, 0xa1, 0x96, 0x56, 0x55, 0xf8, 0x69, 0x6d, 0x57, 0x47, 0x5e, 0xe2, 0x65,
	0x01, 0x83, 0x59, 0xfa, 0xb0, 0x5d, 0x4d, 0x1e, 0x20, 0x30, 0xdd, 0x0e, 0x19, 0x17, 0x2d, 0x46,
	0xe7, 0x89, 0x97, 0x85, 0x45, 0x4c, 0x4f, 0x42, 0xe9, 0xfb, 0x42, 0xb0, 0x23, 0x9c, 0x6e, 0x20,
	0x38, 0xe8, 0xe4, 0x06, 0x56, 0xda, 0x70, 0x65, 0x5c, 0x82, 0xcf, 0xe6, 0x0b, 0xb9, 0x06, 0x1f,
	0x45, 0xed, 0x6c, 0x7d, 0xf6, 0x73, 0x4c, 0x9f, 0x61, 0xed, 0xca, 0x31, 0xd4, 0x83, 0x14, 0x1a,
	0xc9, 0x3d, 0x04, 0x4b, 0x7f, 0x1d, 0x79, 0x89, 0x9f, 0x85, 0x45, 0x44, 0x7f, 0x7d, 0x90, 0x8e,
	0x05, 0xdd, 0xee, 0x01, 0x76, 0x44, 0x8b, 0x57, 0xb8, 0x7c, 0x9c, 0xf7, 0x24, 0x5b, 0x58, 0x39,
	0x4f, 0xf2, 0xb7, 0xf8, 0x61, 0x88, 0xf8, 0xf6, 0xff, 0xb7, 0xa5, 0x47, 0x7a, 0x56, 0x5e, 0xb8,
	0xf9, 0x36, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x90, 0xad, 0xfa, 0xa8, 0x01, 0x00, 0x00,
}
