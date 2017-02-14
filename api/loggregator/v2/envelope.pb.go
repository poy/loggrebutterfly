// Code generated by protoc-gen-go.
// source: envelope.proto
// DO NOT EDIT!

/*
Package loggregator is a generated protocol buffer package.

It is generated from these files:
	envelope.proto
	metron.proto

It has these top-level messages:
	Envelope
	Value
	Log
	Counter
	Gauge
	GaugeValue
	Timer
	MetronResponse
*/
package loggregator

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Log_Type int32

const (
	Log_OUT Log_Type = 0
	Log_ERR Log_Type = 1
)

var Log_Type_name = map[int32]string{
	0: "OUT",
	1: "ERR",
}
var Log_Type_value = map[string]int32{
	"OUT": 0,
	"ERR": 1,
}

func (x Log_Type) String() string {
	return proto.EnumName(Log_Type_name, int32(x))
}
func (Log_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type Envelope struct {
	Timestamp  int64             `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	SourceUuid string            `protobuf:"bytes,2,opt,name=source_uuid,json=sourceUuid" json:"source_uuid,omitempty"`
	Tags       map[string]*Value `protobuf:"bytes,3,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Types that are valid to be assigned to Message:
	//	*Envelope_Log
	//	*Envelope_Counter
	//	*Envelope_Gauge
	//	*Envelope_Timer
	Message isEnvelope_Message `protobuf_oneof:"message"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isEnvelope_Message interface {
	isEnvelope_Message()
}

type Envelope_Log struct {
	Log *Log `protobuf:"bytes,4,opt,name=log,oneof"`
}
type Envelope_Counter struct {
	Counter *Counter `protobuf:"bytes,5,opt,name=counter,oneof"`
}
type Envelope_Gauge struct {
	Gauge *Gauge `protobuf:"bytes,6,opt,name=gauge,oneof"`
}
type Envelope_Timer struct {
	Timer *Timer `protobuf:"bytes,7,opt,name=timer,oneof"`
}

func (*Envelope_Log) isEnvelope_Message()     {}
func (*Envelope_Counter) isEnvelope_Message() {}
func (*Envelope_Gauge) isEnvelope_Message()   {}
func (*Envelope_Timer) isEnvelope_Message()   {}

func (m *Envelope) GetMessage() isEnvelope_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Envelope) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Envelope) GetSourceUuid() string {
	if m != nil {
		return m.SourceUuid
	}
	return ""
}

func (m *Envelope) GetTags() map[string]*Value {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Envelope) GetLog() *Log {
	if x, ok := m.GetMessage().(*Envelope_Log); ok {
		return x.Log
	}
	return nil
}

func (m *Envelope) GetCounter() *Counter {
	if x, ok := m.GetMessage().(*Envelope_Counter); ok {
		return x.Counter
	}
	return nil
}

func (m *Envelope) GetGauge() *Gauge {
	if x, ok := m.GetMessage().(*Envelope_Gauge); ok {
		return x.Gauge
	}
	return nil
}

func (m *Envelope) GetTimer() *Timer {
	if x, ok := m.GetMessage().(*Envelope_Timer); ok {
		return x.Timer
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Envelope) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Envelope_OneofMarshaler, _Envelope_OneofUnmarshaler, _Envelope_OneofSizer, []interface{}{
		(*Envelope_Log)(nil),
		(*Envelope_Counter)(nil),
		(*Envelope_Gauge)(nil),
		(*Envelope_Timer)(nil),
	}
}

func _Envelope_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Envelope)
	// message
	switch x := m.Message.(type) {
	case *Envelope_Log:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Log); err != nil {
			return err
		}
	case *Envelope_Counter:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Counter); err != nil {
			return err
		}
	case *Envelope_Gauge:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Gauge); err != nil {
			return err
		}
	case *Envelope_Timer:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Timer); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Envelope.Message has unexpected type %T", x)
	}
	return nil
}

func _Envelope_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Envelope)
	switch tag {
	case 4: // message.log
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Log)
		err := b.DecodeMessage(msg)
		m.Message = &Envelope_Log{msg}
		return true, err
	case 5: // message.counter
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Counter)
		err := b.DecodeMessage(msg)
		m.Message = &Envelope_Counter{msg}
		return true, err
	case 6: // message.gauge
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Gauge)
		err := b.DecodeMessage(msg)
		m.Message = &Envelope_Gauge{msg}
		return true, err
	case 7: // message.timer
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Timer)
		err := b.DecodeMessage(msg)
		m.Message = &Envelope_Timer{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Envelope_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Envelope)
	// message
	switch x := m.Message.(type) {
	case *Envelope_Log:
		s := proto.Size(x.Log)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Envelope_Counter:
		s := proto.Size(x.Counter)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Envelope_Gauge:
		s := proto.Size(x.Gauge)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Envelope_Timer:
		s := proto.Size(x.Timer)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Value struct {
	// Types that are valid to be assigned to Data:
	//	*Value_Text
	//	*Value_Integer
	//	*Value_Decimal
	Data isValue_Data `protobuf_oneof:"data"`
}

func (m *Value) Reset()                    { *m = Value{} }
func (m *Value) String() string            { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()               {}
func (*Value) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isValue_Data interface {
	isValue_Data()
}

type Value_Text struct {
	Text string `protobuf:"bytes,1,opt,name=text,oneof"`
}
type Value_Integer struct {
	Integer int64 `protobuf:"varint,2,opt,name=integer,oneof"`
}
type Value_Decimal struct {
	Decimal float64 `protobuf:"fixed64,3,opt,name=decimal,oneof"`
}

func (*Value_Text) isValue_Data()    {}
func (*Value_Integer) isValue_Data() {}
func (*Value_Decimal) isValue_Data() {}

func (m *Value) GetData() isValue_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Value) GetText() string {
	if x, ok := m.GetData().(*Value_Text); ok {
		return x.Text
	}
	return ""
}

func (m *Value) GetInteger() int64 {
	if x, ok := m.GetData().(*Value_Integer); ok {
		return x.Integer
	}
	return 0
}

func (m *Value) GetDecimal() float64 {
	if x, ok := m.GetData().(*Value_Decimal); ok {
		return x.Decimal
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Value) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Value_OneofMarshaler, _Value_OneofUnmarshaler, _Value_OneofSizer, []interface{}{
		(*Value_Text)(nil),
		(*Value_Integer)(nil),
		(*Value_Decimal)(nil),
	}
}

func _Value_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Value)
	// data
	switch x := m.Data.(type) {
	case *Value_Text:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Text)
	case *Value_Integer:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Integer))
	case *Value_Decimal:
		b.EncodeVarint(3<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.Decimal))
	case nil:
	default:
		return fmt.Errorf("Value.Data has unexpected type %T", x)
	}
	return nil
}

func _Value_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Value)
	switch tag {
	case 1: // data.text
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Data = &Value_Text{x}
		return true, err
	case 2: // data.integer
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Data = &Value_Integer{int64(x)}
		return true, err
	case 3: // data.decimal
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Data = &Value_Decimal{math.Float64frombits(x)}
		return true, err
	default:
		return false, nil
	}
}

func _Value_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Value)
	// data
	switch x := m.Data.(type) {
	case *Value_Text:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Text)))
		n += len(x.Text)
	case *Value_Integer:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Integer))
	case *Value_Decimal:
		n += proto.SizeVarint(3<<3 | proto.WireFixed64)
		n += 8
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Log struct {
	Payload []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Type    Log_Type `protobuf:"varint,2,opt,name=type,enum=loggregator.Log_Type" json:"type,omitempty"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Log) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Log) GetType() Log_Type {
	if m != nil {
		return m.Type
	}
	return Log_OUT
}

type Counter struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*Counter_Delta
	//	*Counter_Total
	Value isCounter_Value `protobuf_oneof:"value"`
}

func (m *Counter) Reset()                    { *m = Counter{} }
func (m *Counter) String() string            { return proto.CompactTextString(m) }
func (*Counter) ProtoMessage()               {}
func (*Counter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isCounter_Value interface {
	isCounter_Value()
}

type Counter_Delta struct {
	Delta uint64 `protobuf:"varint,2,opt,name=delta,oneof"`
}
type Counter_Total struct {
	Total uint64 `protobuf:"varint,3,opt,name=total,oneof"`
}

func (*Counter_Delta) isCounter_Value() {}
func (*Counter_Total) isCounter_Value() {}

func (m *Counter) GetValue() isCounter_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Counter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Counter) GetDelta() uint64 {
	if x, ok := m.GetValue().(*Counter_Delta); ok {
		return x.Delta
	}
	return 0
}

func (m *Counter) GetTotal() uint64 {
	if x, ok := m.GetValue().(*Counter_Total); ok {
		return x.Total
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Counter) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Counter_OneofMarshaler, _Counter_OneofUnmarshaler, _Counter_OneofSizer, []interface{}{
		(*Counter_Delta)(nil),
		(*Counter_Total)(nil),
	}
}

func _Counter_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Counter)
	// value
	switch x := m.Value.(type) {
	case *Counter_Delta:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Delta))
	case *Counter_Total:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Total))
	case nil:
	default:
		return fmt.Errorf("Counter.Value has unexpected type %T", x)
	}
	return nil
}

func _Counter_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Counter)
	switch tag {
	case 2: // value.delta
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Counter_Delta{x}
		return true, err
	case 3: // value.total
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Counter_Total{x}
		return true, err
	default:
		return false, nil
	}
}

func _Counter_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Counter)
	// value
	switch x := m.Value.(type) {
	case *Counter_Delta:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Delta))
	case *Counter_Total:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Total))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Gauge struct {
	Metrics map[string]*GaugeValue `protobuf:"bytes,1,rep,name=metrics" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Gauge) Reset()                    { *m = Gauge{} }
func (m *Gauge) String() string            { return proto.CompactTextString(m) }
func (*Gauge) ProtoMessage()               {}
func (*Gauge) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Gauge) GetMetrics() map[string]*GaugeValue {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type GaugeValue struct {
	Unit  string  `protobuf:"bytes,1,opt,name=unit" json:"unit,omitempty"`
	Value float64 `protobuf:"fixed64,2,opt,name=value" json:"value,omitempty"`
}

func (m *GaugeValue) Reset()                    { *m = GaugeValue{} }
func (m *GaugeValue) String() string            { return proto.CompactTextString(m) }
func (*GaugeValue) ProtoMessage()               {}
func (*GaugeValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GaugeValue) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *GaugeValue) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Timer struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Start int64  `protobuf:"varint,2,opt,name=start" json:"start,omitempty"`
	Stop  int64  `protobuf:"varint,3,opt,name=stop" json:"stop,omitempty"`
}

func (m *Timer) Reset()                    { *m = Timer{} }
func (m *Timer) String() string            { return proto.CompactTextString(m) }
func (*Timer) ProtoMessage()               {}
func (*Timer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Timer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Timer) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Timer) GetStop() int64 {
	if m != nil {
		return m.Stop
	}
	return 0
}

func init() {
	proto.RegisterType((*Envelope)(nil), "loggregator.Envelope")
	proto.RegisterType((*Value)(nil), "loggregator.Value")
	proto.RegisterType((*Log)(nil), "loggregator.Log")
	proto.RegisterType((*Counter)(nil), "loggregator.Counter")
	proto.RegisterType((*Gauge)(nil), "loggregator.Gauge")
	proto.RegisterType((*GaugeValue)(nil), "loggregator.GaugeValue")
	proto.RegisterType((*Timer)(nil), "loggregator.Timer")
	proto.RegisterEnum("loggregator.Log_Type", Log_Type_name, Log_Type_value)
}

func init() { proto.RegisterFile("envelope.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x93, 0x51, 0x8b, 0xd3, 0x40,
	0x10, 0xc7, 0xbb, 0xb7, 0xc9, 0xc5, 0x4e, 0x8f, 0xa3, 0x2c, 0x55, 0x97, 0x22, 0x5c, 0x09, 0x3e,
	0x44, 0xc1, 0x20, 0x3d, 0x10, 0xf5, 0xb1, 0x52, 0x2c, 0x78, 0x22, 0xac, 0xbd, 0x7b, 0x13, 0x59,
	0x9b, 0x65, 0x09, 0xa6, 0xd9, 0x90, 0x6c, 0x0e, 0xfb, 0x05, 0xfc, 0x0a, 0x7e, 0x5d, 0xd9, 0xd9,
	0xc4, 0x6b, 0x6b, 0xdf, 0x66, 0xe6, 0xff, 0xcb, 0xec, 0xce, 0x7f, 0x36, 0x70, 0xa9, 0xca, 0x7b,
	0x55, 0x98, 0x4a, 0xa5, 0x55, 0x6d, 0xac, 0x61, 0xa3, 0xc2, 0x68, 0x5d, 0x2b, 0x2d, 0xad, 0xa9,
	0xe3, 0xdf, 0x14, 0x1e, 0x2d, 0x3b, 0x9d, 0x3d, 0x83, 0xa1, 0xcd, 0xb7, 0xaa, 0xb1, 0x72, 0x5b,
	0x71, 0x32, 0x23, 0x09, 0x15, 0x0f, 0x05, 0x76, 0x05, 0xa3, 0xc6, 0xb4, 0xf5, 0x46, 0x7d, 0x6f,
	0xdb, 0x3c, 0xe3, 0x67, 0x33, 0x92, 0x0c, 0x05, 0xf8, 0xd2, 0x6d, 0x9b, 0x67, 0xec, 0x1a, 0x02,
	0x2b, 0x75, 0xc3, 0xe9, 0x8c, 0x26, 0xa3, 0xf9, 0x55, 0xba, 0x77, 0x4e, 0xda, 0x9f, 0x91, 0xae,
	0xa5, 0x6e, 0x96, 0xa5, 0xad, 0x77, 0x02, 0x61, 0xf6, 0x1c, 0x68, 0x61, 0x34, 0x0f, 0x66, 0x24,
	0x19, 0xcd, 0xc7, 0x07, 0xdf, 0xdc, 0x18, 0xbd, 0x1a, 0x08, 0x27, 0xb3, 0xd7, 0x10, 0x6d, 0x4c,
	0x5b, 0x5a, 0x55, 0xf3, 0x10, 0xc9, 0xc9, 0x01, 0xf9, 0xc1, 0x6b, 0xab, 0x81, 0xe8, 0x31, 0xf6,
	0x12, 0x42, 0x2d, 0x5b, 0xad, 0xf8, 0x39, 0xf2, 0xec, 0x80, 0xff, 0xe8, 0x94, 0xd5, 0x40, 0x78,
	0xc4, 0xb1, 0x6e, 0xcc, 0x9a, 0x47, 0x27, 0xd8, 0xb5, 0x53, 0x1c, 0x8b, 0xc8, 0xf4, 0x13, 0x0c,
	0xff, 0x8d, 0xc0, 0xc6, 0x40, 0x7f, 0xaa, 0x1d, 0x5a, 0x35, 0x14, 0x2e, 0x64, 0x09, 0x84, 0xf7,
	0xb2, 0x68, 0x15, 0xda, 0x73, 0xdc, 0xea, 0xce, 0x29, 0xc2, 0x03, 0xef, 0xcf, 0xde, 0x92, 0xc5,
	0x10, 0xa2, 0xad, 0x6a, 0x1a, 0xa9, 0x55, 0xfc, 0x0d, 0x42, 0x94, 0xd9, 0x04, 0x02, 0xab, 0x7e,
	0x59, 0xdf, 0x74, 0x35, 0x10, 0x98, 0xb1, 0x29, 0x44, 0x79, 0x69, 0x95, 0x56, 0x35, 0x76, 0xa6,
	0x6e, 0xd4, 0xae, 0xe0, 0xb4, 0x4c, 0x6d, 0xf2, 0xad, 0x2c, 0x38, 0x9d, 0x91, 0x84, 0x38, 0xad,
	0x2b, 0x2c, 0xce, 0x21, 0xc8, 0xa4, 0x95, 0x71, 0x06, 0xf4, 0xc6, 0x68, 0xc6, 0x21, 0xaa, 0xe4,
	0xae, 0x30, 0x32, 0xc3, 0xfe, 0x17, 0xa2, 0x4f, 0xd9, 0x0b, 0x08, 0xec, 0xae, 0xf2, 0xf7, 0xbe,
	0x9c, 0x3f, 0x3e, 0x5e, 0x44, 0xba, 0xde, 0x55, 0x4a, 0x20, 0x12, 0x73, 0x08, 0x5c, 0xc6, 0x22,
	0xa0, 0x5f, 0x6e, 0xd7, 0xe3, 0x81, 0x0b, 0x96, 0x42, 0x8c, 0x49, 0x7c, 0x07, 0x51, 0xb7, 0x0a,
	0xc6, 0x20, 0x28, 0xe5, 0x56, 0x75, 0xde, 0x60, 0xcc, 0x9e, 0x40, 0x98, 0xa9, 0xc2, 0x4a, 0x3c,
	0x24, 0x70, 0x9e, 0x62, 0xea, 0xea, 0xd6, 0xd8, 0xee, 0xfa, 0x58, 0xc7, 0x74, 0x11, 0x75, 0x66,
	0xc6, 0x7f, 0x08, 0x84, 0xb8, 0x33, 0xf6, 0xce, 0x39, 0x66, 0xeb, 0x7c, 0xd3, 0x70, 0x72, 0xe2,
	0x99, 0x21, 0x94, 0x7e, 0xf6, 0x84, 0x7f, 0x66, 0x3d, 0x3f, 0xfd, 0x0a, 0x17, 0xfb, 0xc2, 0x89,
	0xe5, 0xbd, 0x3a, 0x5c, 0xde, 0xd3, 0xff, 0x5b, 0x1f, 0x6f, 0x30, 0x7e, 0x03, 0xf0, 0x20, 0xb8,
	0xa1, 0xdb, 0x32, 0xb7, 0xfd, 0xd0, 0x2e, 0x66, 0x93, 0xfd, 0xa6, 0xa4, 0xfb, 0x36, 0x5e, 0x42,
	0x88, 0x0f, 0xeb, 0xa4, 0x4f, 0x13, 0x08, 0x1b, 0x2b, 0x6b, 0xeb, 0x57, 0x2d, 0x7c, 0xe2, 0xc8,
	0xc6, 0x9a, 0x0a, 0x4d, 0xa2, 0x02, 0xe3, 0x1f, 0xe7, 0xf8, 0x4b, 0x5f, 0xff, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0xdd, 0xa4, 0x15, 0xe6, 0xe4, 0x03, 0x00, 0x00,
}