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
	AggregateInfo
	QueryResponse
	AggregateResponse
	AnalystFilter
	TimeRange
	CounterFilter
	LogFilter
	GaugeFilter
	GaugeFilterValue
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
	Filter *AnalystFilter `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
}

func (m *QueryInfo) Reset()                    { *m = QueryInfo{} }
func (m *QueryInfo) String() string            { return proto.CompactTextString(m) }
func (*QueryInfo) ProtoMessage()               {}
func (*QueryInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *QueryInfo) GetFilter() *AnalystFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type AggregateInfo struct {
	Query         *QueryInfo `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
	BucketWidthNs int64      `protobuf:"varint,2,opt,name=bucket_width_ns,json=bucketWidthNs" json:"bucket_width_ns,omitempty"`
}

func (m *AggregateInfo) Reset()                    { *m = AggregateInfo{} }
func (m *AggregateInfo) String() string            { return proto.CompactTextString(m) }
func (*AggregateInfo) ProtoMessage()               {}
func (*AggregateInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AggregateInfo) GetQuery() *QueryInfo {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *AggregateInfo) GetBucketWidthNs() int64 {
	if m != nil {
		return m.BucketWidthNs
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

type AggregateResponse struct {
	Results map[int64]float64 `protobuf:"bytes,1,rep,name=results" json:"results,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
}

func (m *AggregateResponse) Reset()                    { *m = AggregateResponse{} }
func (m *AggregateResponse) String() string            { return proto.CompactTextString(m) }
func (*AggregateResponse) ProtoMessage()               {}
func (*AggregateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AggregateResponse) GetResults() map[int64]float64 {
	if m != nil {
		return m.Results
	}
	return nil
}

type AnalystFilter struct {
	SourceId  string     `protobuf:"bytes,1,opt,name=source_id,json=sourceId" json:"source_id,omitempty"`
	TimeRange *TimeRange `protobuf:"bytes,2,opt,name=time_range,json=timeRange" json:"time_range,omitempty"`
	// Types that are valid to be assigned to Envelopes:
	//	*AnalystFilter_Counter
	//	*AnalystFilter_Log
	//	*AnalystFilter_Gauge
	Envelopes isAnalystFilter_Envelopes `protobuf_oneof:"Envelopes"`
}

func (m *AnalystFilter) Reset()                    { *m = AnalystFilter{} }
func (m *AnalystFilter) String() string            { return proto.CompactTextString(m) }
func (*AnalystFilter) ProtoMessage()               {}
func (*AnalystFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type isAnalystFilter_Envelopes interface {
	isAnalystFilter_Envelopes()
}

type AnalystFilter_Counter struct {
	Counter *CounterFilter `protobuf:"bytes,3,opt,name=counter,oneof"`
}
type AnalystFilter_Log struct {
	Log *LogFilter `protobuf:"bytes,4,opt,name=log,oneof"`
}
type AnalystFilter_Gauge struct {
	Gauge *GaugeFilter `protobuf:"bytes,5,opt,name=gauge,oneof"`
}

func (*AnalystFilter_Counter) isAnalystFilter_Envelopes() {}
func (*AnalystFilter_Log) isAnalystFilter_Envelopes()     {}
func (*AnalystFilter_Gauge) isAnalystFilter_Envelopes()   {}

func (m *AnalystFilter) GetEnvelopes() isAnalystFilter_Envelopes {
	if m != nil {
		return m.Envelopes
	}
	return nil
}

func (m *AnalystFilter) GetSourceId() string {
	if m != nil {
		return m.SourceId
	}
	return ""
}

func (m *AnalystFilter) GetTimeRange() *TimeRange {
	if m != nil {
		return m.TimeRange
	}
	return nil
}

func (m *AnalystFilter) GetCounter() *CounterFilter {
	if x, ok := m.GetEnvelopes().(*AnalystFilter_Counter); ok {
		return x.Counter
	}
	return nil
}

func (m *AnalystFilter) GetLog() *LogFilter {
	if x, ok := m.GetEnvelopes().(*AnalystFilter_Log); ok {
		return x.Log
	}
	return nil
}

func (m *AnalystFilter) GetGauge() *GaugeFilter {
	if x, ok := m.GetEnvelopes().(*AnalystFilter_Gauge); ok {
		return x.Gauge
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AnalystFilter) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AnalystFilter_OneofMarshaler, _AnalystFilter_OneofUnmarshaler, _AnalystFilter_OneofSizer, []interface{}{
		(*AnalystFilter_Counter)(nil),
		(*AnalystFilter_Log)(nil),
		(*AnalystFilter_Gauge)(nil),
	}
}

func _AnalystFilter_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AnalystFilter)
	// Envelopes
	switch x := m.Envelopes.(type) {
	case *AnalystFilter_Counter:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Counter); err != nil {
			return err
		}
	case *AnalystFilter_Log:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Log); err != nil {
			return err
		}
	case *AnalystFilter_Gauge:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Gauge); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AnalystFilter.Envelopes has unexpected type %T", x)
	}
	return nil
}

func _AnalystFilter_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AnalystFilter)
	switch tag {
	case 3: // Envelopes.counter
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CounterFilter)
		err := b.DecodeMessage(msg)
		m.Envelopes = &AnalystFilter_Counter{msg}
		return true, err
	case 4: // Envelopes.log
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LogFilter)
		err := b.DecodeMessage(msg)
		m.Envelopes = &AnalystFilter_Log{msg}
		return true, err
	case 5: // Envelopes.gauge
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GaugeFilter)
		err := b.DecodeMessage(msg)
		m.Envelopes = &AnalystFilter_Gauge{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AnalystFilter_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AnalystFilter)
	// Envelopes
	switch x := m.Envelopes.(type) {
	case *AnalystFilter_Counter:
		s := proto.Size(x.Counter)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AnalystFilter_Log:
		s := proto.Size(x.Log)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AnalystFilter_Gauge:
		s := proto.Size(x.Gauge)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// [start, end)
type TimeRange struct {
	Start int64 `protobuf:"varint,1,opt,name=start" json:"start,omitempty"`
	End   int64 `protobuf:"varint,2,opt,name=end" json:"end,omitempty"`
}

func (m *TimeRange) Reset()                    { *m = TimeRange{} }
func (m *TimeRange) String() string            { return proto.CompactTextString(m) }
func (*TimeRange) ProtoMessage()               {}
func (*TimeRange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

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

type CounterFilter struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CounterFilter) Reset()                    { *m = CounterFilter{} }
func (m *CounterFilter) String() string            { return proto.CompactTextString(m) }
func (*CounterFilter) ProtoMessage()               {}
func (*CounterFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CounterFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type LogFilter struct {
	// Types that are valid to be assigned to Payload:
	//	*LogFilter_Regexp
	//	*LogFilter_Match
	Payload isLogFilter_Payload `protobuf_oneof:"Payload"`
}

func (m *LogFilter) Reset()                    { *m = LogFilter{} }
func (m *LogFilter) String() string            { return proto.CompactTextString(m) }
func (*LogFilter) ProtoMessage()               {}
func (*LogFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type isLogFilter_Payload interface {
	isLogFilter_Payload()
}

type LogFilter_Regexp struct {
	Regexp string `protobuf:"bytes,1,opt,name=regexp,oneof"`
}
type LogFilter_Match struct {
	Match []byte `protobuf:"bytes,2,opt,name=match,proto3,oneof"`
}

func (*LogFilter_Regexp) isLogFilter_Payload() {}
func (*LogFilter_Match) isLogFilter_Payload()  {}

func (m *LogFilter) GetPayload() isLogFilter_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *LogFilter) GetRegexp() string {
	if x, ok := m.GetPayload().(*LogFilter_Regexp); ok {
		return x.Regexp
	}
	return ""
}

func (m *LogFilter) GetMatch() []byte {
	if x, ok := m.GetPayload().(*LogFilter_Match); ok {
		return x.Match
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LogFilter) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LogFilter_OneofMarshaler, _LogFilter_OneofUnmarshaler, _LogFilter_OneofSizer, []interface{}{
		(*LogFilter_Regexp)(nil),
		(*LogFilter_Match)(nil),
	}
}

func _LogFilter_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LogFilter)
	// Payload
	switch x := m.Payload.(type) {
	case *LogFilter_Regexp:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Regexp)
	case *LogFilter_Match:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Match)
	case nil:
	default:
		return fmt.Errorf("LogFilter.Payload has unexpected type %T", x)
	}
	return nil
}

func _LogFilter_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LogFilter)
	switch tag {
	case 1: // Payload.regexp
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Payload = &LogFilter_Regexp{x}
		return true, err
	case 2: // Payload.match
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Payload = &LogFilter_Match{x}
		return true, err
	default:
		return false, nil
	}
}

func _LogFilter_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LogFilter)
	// Payload
	switch x := m.Payload.(type) {
	case *LogFilter_Regexp:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Regexp)))
		n += len(x.Regexp)
	case *LogFilter_Match:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Match)))
		n += len(x.Match)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type GaugeFilter struct {
	Filter map[string]*GaugeFilterValue `protobuf:"bytes,1,rep,name=filter" json:"filter,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *GaugeFilter) Reset()                    { *m = GaugeFilter{} }
func (m *GaugeFilter) String() string            { return proto.CompactTextString(m) }
func (*GaugeFilter) ProtoMessage()               {}
func (*GaugeFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GaugeFilter) GetFilter() map[string]*GaugeFilterValue {
	if m != nil {
		return m.Filter
	}
	return nil
}

type GaugeFilterValue struct {
	Value float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
}

func (m *GaugeFilterValue) Reset()                    { *m = GaugeFilterValue{} }
func (m *GaugeFilterValue) String() string            { return proto.CompactTextString(m) }
func (*GaugeFilterValue) ProtoMessage()               {}
func (*GaugeFilterValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GaugeFilterValue) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryInfo)(nil), "loggrebutterfly.QueryInfo")
	proto.RegisterType((*AggregateInfo)(nil), "loggrebutterfly.AggregateInfo")
	proto.RegisterType((*QueryResponse)(nil), "loggrebutterfly.QueryResponse")
	proto.RegisterType((*AggregateResponse)(nil), "loggrebutterfly.AggregateResponse")
	proto.RegisterType((*AnalystFilter)(nil), "loggrebutterfly.AnalystFilter")
	proto.RegisterType((*TimeRange)(nil), "loggrebutterfly.TimeRange")
	proto.RegisterType((*CounterFilter)(nil), "loggrebutterfly.CounterFilter")
	proto.RegisterType((*LogFilter)(nil), "loggrebutterfly.LogFilter")
	proto.RegisterType((*GaugeFilter)(nil), "loggrebutterfly.GaugeFilter")
	proto.RegisterType((*GaugeFilterValue)(nil), "loggrebutterfly.GaugeFilterValue")
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
	Aggregate(ctx context.Context, in *AggregateInfo, opts ...grpc.CallOption) (*AggregateResponse, error)
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

func (c *analystClient) Aggregate(ctx context.Context, in *AggregateInfo, opts ...grpc.CallOption) (*AggregateResponse, error) {
	out := new(AggregateResponse)
	err := grpc.Invoke(ctx, "/loggrebutterfly.Analyst/Aggregate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Analyst service

type AnalystServer interface {
	Query(context.Context, *QueryInfo) (*QueryResponse, error)
	Aggregate(context.Context, *AggregateInfo) (*AggregateResponse, error)
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

func _Analyst_Aggregate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AggregateInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalystServer).Aggregate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loggrebutterfly.Analyst/Aggregate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalystServer).Aggregate(ctx, req.(*AggregateInfo))
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
		{
			MethodName: "Aggregate",
			Handler:    _Analyst_Aggregate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "analyst.proto",
}

func init() { proto.RegisterFile("analyst.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 614 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x54, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0x8d, 0x9b, 0xa6, 0xf9, 0x79, 0xd2, 0xa8, 0xfd, 0xad, 0x10, 0xb2, 0x02, 0xaa, 0x8a, 0x91,
	0x50, 0x4e, 0x0e, 0x4a, 0x51, 0x81, 0x9e, 0x68, 0xab, 0xf4, 0x8f, 0x84, 0x10, 0x5d, 0x21, 0xb8,
	0x20, 0x45, 0x1b, 0x7b, 0xea, 0x5a, 0x75, 0xbc, 0x66, 0xbd, 0x0e, 0xf8, 0x8b, 0x70, 0xe0, 0xca,
	0x99, 0xef, 0x88, 0x76, 0xd7, 0x76, 0xdd, 0xb4, 0x0d, 0xa7, 0xec, 0xce, 0xbc, 0xf7, 0x76, 0x66,
	0xfc, 0x26, 0xd0, 0x67, 0x09, 0x8b, 0x8b, 0x4c, 0x7a, 0xa9, 0xe0, 0x92, 0x93, 0xad, 0x98, 0x87,
	0xa1, 0xc0, 0x59, 0x2e, 0x25, 0x8a, 0xcb, 0xb8, 0x18, 0x9c, 0x84, 0x91, 0xbc, 0xca, 0x67, 0x9e,
	0xcf, 0xe7, 0x23, 0x96, 0xf2, 0x22, 0xc0, 0xc4, 0xc7, 0xd1, 0x12, 0x6a, 0xc4, 0xd2, 0xa8, 0x8c,
	0x85, 0x4c, 0x72, 0x31, 0x5a, 0x8c, 0x47, 0x98, 0x2c, 0x30, 0xe6, 0x29, 0x1a, 0x61, 0xf7, 0x18,
	0xec, 0x8b, 0x1c, 0x45, 0x71, 0x9e, 0x5c, 0x72, 0xb2, 0x0f, 0x1b, 0x97, 0x51, 0x2c, 0x51, 0x38,
	0xd6, 0xae, 0x35, 0xec, 0x8d, 0x77, 0xbc, 0x25, 0x41, 0xef, 0xd0, 0x54, 0x75, 0xa2, 0x51, 0xb4,
	0x44, 0xbb, 0x11, 0xf4, 0x0f, 0xcb, 0x47, 0x50, 0x0b, 0xbd, 0x84, 0xce, 0x37, 0xa5, 0x5a, 0xea,
	0x0c, 0xee, 0xe8, 0xd4, 0x6f, 0x52, 0x03, 0x24, 0x2f, 0x60, 0x6b, 0x96, 0xfb, 0xd7, 0x28, 0xa7,
	0xdf, 0xa3, 0x40, 0x5e, 0x4d, 0x93, 0xcc, 0x59, 0xdb, 0xb5, 0x86, 0x6d, 0xda, 0x37, 0xe1, 0x2f,
	0x2a, 0xfa, 0x21, 0x73, 0x4f, 0xa1, 0xaf, 0xb9, 0x14, 0xb3, 0x94, 0x27, 0x19, 0x92, 0x7d, 0xb0,
	0xab, 0x96, 0x32, 0xc7, 0xda, 0x6d, 0x0f, 0x7b, 0x63, 0xc7, 0x6b, 0xf4, 0xec, 0x2d, 0xc6, 0xde,
	0xa4, 0x04, 0xd0, 0x1b, 0xa8, 0xfb, 0xcb, 0x82, 0xff, 0xeb, 0xa2, 0x6b, 0xb5, 0x73, 0xe8, 0x0a,
	0xcc, 0xf2, 0x58, 0x56, 0x5a, 0xa3, 0xbb, 0x23, 0x58, 0x26, 0x79, 0xd4, 0x30, 0x26, 0x89, 0x14,
	0x05, 0xad, 0xf8, 0x83, 0x03, 0xd8, 0x6c, 0x26, 0xc8, 0x36, 0xb4, 0xaf, 0xd1, 0x4c, 0xa4, 0x4d,
	0xd5, 0x91, 0x3c, 0x82, 0xce, 0x82, 0xc5, 0x39, 0xea, 0x4e, 0x2d, 0x6a, 0x2e, 0x07, 0x6b, 0x6f,
	0x2c, 0xf7, 0xe7, 0x1a, 0xf4, 0x6f, 0x8d, 0x9a, 0x3c, 0x01, 0x3b, 0xe3, 0xb9, 0xf0, 0x71, 0x1a,
	0x05, 0x5a, 0xc3, 0xa6, 0xff, 0x99, 0xc0, 0x79, 0x40, 0xde, 0x02, 0xc8, 0x68, 0x8e, 0x53, 0xc1,
	0x92, 0xd0, 0xa8, 0xdd, 0x37, 0xf3, 0x4f, 0xd1, 0x1c, 0xa9, 0x42, 0x50, 0x5b, 0x56, 0x47, 0x72,
	0x00, 0x5d, 0x9f, 0xe7, 0x89, 0xfa, 0xe6, 0xed, 0x07, 0xbe, 0xf9, 0xb1, 0xc9, 0x9b, 0x42, 0xce,
	0x5a, 0xb4, 0x22, 0x10, 0x0f, 0xda, 0x31, 0x0f, 0x9d, 0xf5, 0x07, 0xde, 0x7b, 0xcf, 0xc3, 0x9a,
	0xa3, 0x80, 0xe4, 0x15, 0x74, 0x42, 0x96, 0x87, 0xe8, 0x74, 0x34, 0xe3, 0xe9, 0x1d, 0xc6, 0xa9,
	0xca, 0xd6, 0x1c, 0x03, 0x3e, 0xea, 0x81, 0x3d, 0xa9, 0xbf, 0xda, 0x1e, 0xd8, 0x75, 0x1b, 0x6a,
	0x7e, 0x99, 0x64, 0x42, 0x96, 0x33, 0x35, 0x17, 0x35, 0x67, 0x4c, 0x82, 0xd2, 0x3d, 0xea, 0xe8,
	0x3e, 0x87, 0xfe, 0xad, 0x1e, 0x08, 0x81, 0xf5, 0x84, 0xcd, 0xb1, 0x9c, 0xa3, 0x3e, 0xbb, 0x67,
	0x60, 0xd7, 0x05, 0x13, 0x07, 0x36, 0x04, 0x86, 0xf8, 0x23, 0x35, 0x90, 0xb3, 0x16, 0x2d, 0xef,
	0xe4, 0x31, 0x74, 0xe6, 0x4c, 0xfa, 0x57, 0x5a, 0x7f, 0x53, 0x55, 0xa9, 0xaf, 0x47, 0x36, 0x74,
	0x3f, 0xb2, 0x22, 0xe6, 0x2c, 0x70, 0xff, 0x58, 0xd0, 0x6b, 0x74, 0x42, 0xde, 0x35, 0xb6, 0x4a,
	0x59, 0x6a, 0xb8, 0xaa, 0x6f, 0xcf, 0xfc, 0x18, 0x2f, 0x95, 0xbc, 0xc1, 0x57, 0xe8, 0x35, 0xc2,
	0x4d, 0x27, 0xd9, 0xc6, 0x49, 0xaf, 0x9b, 0x4e, 0xea, 0x8d, 0x9f, 0xad, 0x7a, 0xe1, 0xb3, 0x02,
	0x36, 0xcd, 0x36, 0x84, 0xed, 0xe5, 0xf4, 0x8d, 0x35, 0xad, 0x86, 0x35, 0xc7, 0xbf, 0x2d, 0xe8,
	0x96, 0xb6, 0x24, 0x13, 0xe8, 0xe8, 0x45, 0x24, 0x2b, 0x96, 0x7b, 0xb0, 0x73, 0x7f, 0xae, 0xda,
	0x1c, 0xb7, 0x45, 0x2e, 0xc0, 0xae, 0x17, 0x8a, 0xec, 0x3c, 0xbc, 0x6c, 0x5a, 0xce, 0xfd, 0xf7,
	0x32, 0xba, 0xad, 0xd9, 0x86, 0xfe, 0x67, 0xdb, 0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x39, 0x7c,
	0x84, 0x9a, 0x43, 0x05, 0x00, 0x00,
}
