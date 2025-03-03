// Code generated by capnpc-go. DO NOT EDIT.

package json

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
	math "math"
	strconv "strconv"
)

const Name = uint64(0xfa5b1fd61c2e7c3d)
const Flatten = uint64(0x82d3e852af0336bf)
const Discriminator = uint64(0xcfa794e8d19a0162)
const Base64 = uint64(0xd7d879450a253e4b)
const Hex = uint64(0xf061e22f0ae5c7b5)
const Notification = uint64(0xa0a054dea32fd98c)

type Value struct{ capnp.Struct }
type Value_Which uint16

const (
	Value_Which_null    Value_Which = 0
	Value_Which_boolean Value_Which = 1
	Value_Which_number  Value_Which = 2
	Value_Which_string_ Value_Which = 3
	Value_Which_array   Value_Which = 4
	Value_Which_object  Value_Which = 5
	Value_Which_call    Value_Which = 6
)

func (w Value_Which) String() string {
	const s = "nullbooleannumberstring_arrayobjectcall"
	switch w {
	case Value_Which_null:
		return s[0:4]
	case Value_Which_boolean:
		return s[4:11]
	case Value_Which_number:
		return s[11:17]
	case Value_Which_string_:
		return s[17:24]
	case Value_Which_array:
		return s[24:29]
	case Value_Which_object:
		return s[29:35]
	case Value_Which_call:
		return s[35:39]

	}
	return "Value_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Value_TypeID is the unique identifier for the type Value.
const Value_TypeID = 0xa3fa7845f919dd83

func NewValue(s *capnp.Segment) (Value, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return Value{st}, err
}

func NewRootValue(s *capnp.Segment) (Value, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return Value{st}, err
}

func ReadRootValue(msg *capnp.Message) (Value, error) {
	root, err := msg.Root()
	return Value{root.Struct()}, err
}

func (s Value) String() string {
	str, _ := text.Marshal(0xa3fa7845f919dd83, s.Struct)
	return str
}

func (s Value) Which() Value_Which {
	return Value_Which(s.Struct.Uint16(0))
}
func (s Value) SetNull() {
	s.Struct.SetUint16(0, 0)

}

func (s Value) Boolean() bool {
	if s.Struct.Uint16(0) != 1 {
		panic("Which() != boolean")
	}
	return s.Struct.Bit(16)
}

func (s Value) SetBoolean(v bool) {
	s.Struct.SetUint16(0, 1)
	s.Struct.SetBit(16, v)
}

func (s Value) Number() float64 {
	if s.Struct.Uint16(0) != 2 {
		panic("Which() != number")
	}
	return math.Float64frombits(s.Struct.Uint64(8))
}

func (s Value) SetNumber(v float64) {
	s.Struct.SetUint16(0, 2)
	s.Struct.SetUint64(8, math.Float64bits(v))
}

func (s Value) String_() (string, error) {
	if s.Struct.Uint16(0) != 3 {
		panic("Which() != string_")
	}
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Value) HasString_() bool {
	if s.Struct.Uint16(0) != 3 {
		return false
	}
	return s.Struct.HasPtr(0)
}

func (s Value) String_Bytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Value) SetString_(v string) error {
	s.Struct.SetUint16(0, 3)
	return s.Struct.SetText(0, v)
}

func (s Value) Array() (Value_List, error) {
	if s.Struct.Uint16(0) != 4 {
		panic("Which() != array")
	}
	p, err := s.Struct.Ptr(0)
	return Value_List{List: p.List()}, err
}

func (s Value) HasArray() bool {
	if s.Struct.Uint16(0) != 4 {
		return false
	}
	return s.Struct.HasPtr(0)
}

func (s Value) SetArray(v Value_List) error {
	s.Struct.SetUint16(0, 4)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewArray sets the array field to a newly
// allocated Value_List, preferring placement in s's segment.
func (s Value) NewArray(n int32) (Value_List, error) {
	s.Struct.SetUint16(0, 4)
	l, err := NewValue_List(s.Struct.Segment(), n)
	if err != nil {
		return Value_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Value) Object() (Value_Field_List, error) {
	if s.Struct.Uint16(0) != 5 {
		panic("Which() != object")
	}
	p, err := s.Struct.Ptr(0)
	return Value_Field_List{List: p.List()}, err
}

func (s Value) HasObject() bool {
	if s.Struct.Uint16(0) != 5 {
		return false
	}
	return s.Struct.HasPtr(0)
}

func (s Value) SetObject(v Value_Field_List) error {
	s.Struct.SetUint16(0, 5)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewObject sets the object field to a newly
// allocated Value_Field_List, preferring placement in s's segment.
func (s Value) NewObject(n int32) (Value_Field_List, error) {
	s.Struct.SetUint16(0, 5)
	l, err := NewValue_Field_List(s.Struct.Segment(), n)
	if err != nil {
		return Value_Field_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Value) Call() (Value_Call, error) {
	if s.Struct.Uint16(0) != 6 {
		panic("Which() != call")
	}
	p, err := s.Struct.Ptr(0)
	return Value_Call{Struct: p.Struct()}, err
}

func (s Value) HasCall() bool {
	if s.Struct.Uint16(0) != 6 {
		return false
	}
	return s.Struct.HasPtr(0)
}

func (s Value) SetCall(v Value_Call) error {
	s.Struct.SetUint16(0, 6)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewCall sets the call field to a newly
// allocated Value_Call struct, preferring placement in s's segment.
func (s Value) NewCall() (Value_Call, error) {
	s.Struct.SetUint16(0, 6)
	ss, err := NewValue_Call(s.Struct.Segment())
	if err != nil {
		return Value_Call{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Value_List is a list of Value.
type Value_List = capnp.StructList[Value]

// NewValue creates a new list of Value.
func NewValue_List(s *capnp.Segment, sz int32) (Value_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1}, sz)
	return capnp.StructList[Value]{l}, err
}

// Value_Future is a wrapper for a Value promised by a client call.
type Value_Future struct{ *capnp.Future }

func (p Value_Future) Struct() (Value, error) {
	s, err := p.Future.Struct()
	return Value{s}, err
}

func (p Value_Future) Call() Value_Call_Future {
	return Value_Call_Future{Future: p.Future.Field(0, nil)}
}

type Value_Field struct{ capnp.Struct }

// Value_Field_TypeID is the unique identifier for the type Value_Field.
const Value_Field_TypeID = 0xe31026e735d69ddf

func NewValue_Field(s *capnp.Segment) (Value_Field, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Value_Field{st}, err
}

func NewRootValue_Field(s *capnp.Segment) (Value_Field, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Value_Field{st}, err
}

func ReadRootValue_Field(msg *capnp.Message) (Value_Field, error) {
	root, err := msg.Root()
	return Value_Field{root.Struct()}, err
}

func (s Value_Field) String() string {
	str, _ := text.Marshal(0xe31026e735d69ddf, s.Struct)
	return str
}

func (s Value_Field) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Value_Field) HasName() bool {
	return s.Struct.HasPtr(0)
}

func (s Value_Field) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Value_Field) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Value_Field) Value() (Value, error) {
	p, err := s.Struct.Ptr(1)
	return Value{Struct: p.Struct()}, err
}

func (s Value_Field) HasValue() bool {
	return s.Struct.HasPtr(1)
}

func (s Value_Field) SetValue(v Value) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewValue sets the value field to a newly
// allocated Value struct, preferring placement in s's segment.
func (s Value_Field) NewValue() (Value, error) {
	ss, err := NewValue(s.Struct.Segment())
	if err != nil {
		return Value{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// Value_Field_List is a list of Value_Field.
type Value_Field_List = capnp.StructList[Value_Field]

// NewValue_Field creates a new list of Value_Field.
func NewValue_Field_List(s *capnp.Segment, sz int32) (Value_Field_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return capnp.StructList[Value_Field]{l}, err
}

// Value_Field_Future is a wrapper for a Value_Field promised by a client call.
type Value_Field_Future struct{ *capnp.Future }

func (p Value_Field_Future) Struct() (Value_Field, error) {
	s, err := p.Future.Struct()
	return Value_Field{s}, err
}

func (p Value_Field_Future) Value() Value_Future {
	return Value_Future{Future: p.Future.Field(1, nil)}
}

type Value_Call struct{ capnp.Struct }

// Value_Call_TypeID is the unique identifier for the type Value_Call.
const Value_Call_TypeID = 0xa0d9f6eca1c93d48

func NewValue_Call(s *capnp.Segment) (Value_Call, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Value_Call{st}, err
}

func NewRootValue_Call(s *capnp.Segment) (Value_Call, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Value_Call{st}, err
}

func ReadRootValue_Call(msg *capnp.Message) (Value_Call, error) {
	root, err := msg.Root()
	return Value_Call{root.Struct()}, err
}

func (s Value_Call) String() string {
	str, _ := text.Marshal(0xa0d9f6eca1c93d48, s.Struct)
	return str
}

func (s Value_Call) Function() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Value_Call) HasFunction() bool {
	return s.Struct.HasPtr(0)
}

func (s Value_Call) FunctionBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Value_Call) SetFunction(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Value_Call) Params() (Value_List, error) {
	p, err := s.Struct.Ptr(1)
	return Value_List{List: p.List()}, err
}

func (s Value_Call) HasParams() bool {
	return s.Struct.HasPtr(1)
}

func (s Value_Call) SetParams(v Value_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewParams sets the params field to a newly
// allocated Value_List, preferring placement in s's segment.
func (s Value_Call) NewParams(n int32) (Value_List, error) {
	l, err := NewValue_List(s.Struct.Segment(), n)
	if err != nil {
		return Value_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// Value_Call_List is a list of Value_Call.
type Value_Call_List = capnp.StructList[Value_Call]

// NewValue_Call creates a new list of Value_Call.
func NewValue_Call_List(s *capnp.Segment, sz int32) (Value_Call_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return capnp.StructList[Value_Call]{l}, err
}

// Value_Call_Future is a wrapper for a Value_Call promised by a client call.
type Value_Call_Future struct{ *capnp.Future }

func (p Value_Call_Future) Struct() (Value_Call, error) {
	s, err := p.Future.Struct()
	return Value_Call{s}, err
}

type FlattenOptions struct{ capnp.Struct }

// FlattenOptions_TypeID is the unique identifier for the type FlattenOptions.
const FlattenOptions_TypeID = 0xc4df13257bc2ea61

func NewFlattenOptions(s *capnp.Segment) (FlattenOptions, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return FlattenOptions{st}, err
}

func NewRootFlattenOptions(s *capnp.Segment) (FlattenOptions, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return FlattenOptions{st}, err
}

func ReadRootFlattenOptions(msg *capnp.Message) (FlattenOptions, error) {
	root, err := msg.Root()
	return FlattenOptions{root.Struct()}, err
}

func (s FlattenOptions) String() string {
	str, _ := text.Marshal(0xc4df13257bc2ea61, s.Struct)
	return str
}

func (s FlattenOptions) Prefix() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s FlattenOptions) HasPrefix() bool {
	return s.Struct.HasPtr(0)
}

func (s FlattenOptions) PrefixBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s FlattenOptions) SetPrefix(v string) error {
	return s.Struct.SetText(0, v)
}

// FlattenOptions_List is a list of FlattenOptions.
type FlattenOptions_List = capnp.StructList[FlattenOptions]

// NewFlattenOptions creates a new list of FlattenOptions.
func NewFlattenOptions_List(s *capnp.Segment, sz int32) (FlattenOptions_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[FlattenOptions]{l}, err
}

// FlattenOptions_Future is a wrapper for a FlattenOptions promised by a client call.
type FlattenOptions_Future struct{ *capnp.Future }

func (p FlattenOptions_Future) Struct() (FlattenOptions, error) {
	s, err := p.Future.Struct()
	return FlattenOptions{s}, err
}

type DiscriminatorOptions struct{ capnp.Struct }

// DiscriminatorOptions_TypeID is the unique identifier for the type DiscriminatorOptions.
const DiscriminatorOptions_TypeID = 0xc2f8c20c293e5319

func NewDiscriminatorOptions(s *capnp.Segment) (DiscriminatorOptions, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return DiscriminatorOptions{st}, err
}

func NewRootDiscriminatorOptions(s *capnp.Segment) (DiscriminatorOptions, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return DiscriminatorOptions{st}, err
}

func ReadRootDiscriminatorOptions(msg *capnp.Message) (DiscriminatorOptions, error) {
	root, err := msg.Root()
	return DiscriminatorOptions{root.Struct()}, err
}

func (s DiscriminatorOptions) String() string {
	str, _ := text.Marshal(0xc2f8c20c293e5319, s.Struct)
	return str
}

func (s DiscriminatorOptions) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s DiscriminatorOptions) HasName() bool {
	return s.Struct.HasPtr(0)
}

func (s DiscriminatorOptions) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s DiscriminatorOptions) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s DiscriminatorOptions) ValueName() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s DiscriminatorOptions) HasValueName() bool {
	return s.Struct.HasPtr(1)
}

func (s DiscriminatorOptions) ValueNameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s DiscriminatorOptions) SetValueName(v string) error {
	return s.Struct.SetText(1, v)
}

// DiscriminatorOptions_List is a list of DiscriminatorOptions.
type DiscriminatorOptions_List = capnp.StructList[DiscriminatorOptions]

// NewDiscriminatorOptions creates a new list of DiscriminatorOptions.
func NewDiscriminatorOptions_List(s *capnp.Segment, sz int32) (DiscriminatorOptions_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return capnp.StructList[DiscriminatorOptions]{l}, err
}

// DiscriminatorOptions_Future is a wrapper for a DiscriminatorOptions promised by a client call.
type DiscriminatorOptions_Future struct{ *capnp.Future }

func (p DiscriminatorOptions_Future) Struct() (DiscriminatorOptions, error) {
	s, err := p.Future.Struct()
	return DiscriminatorOptions{s}, err
}

const schema_8ef99297a43a5e34 = "x\xda\x8c\x94_h[e\x18\xc6\x9f\xe7\xfb\xce9\xeb" +
	"\x9a\xd4\xe6x\"Z\xd8\x88\x17\xab\xda\xa2\xed\xd6\xd5)" +
	"\x81-:\xed(\x8a\xb3_#\xde\x88\x7fN\xb2S\xcd" +
	"89\x09I\xaa\xad\x0a\x83y'\xfe\xc1!\x08\x03\xa1" +
	"2\xc1+Q\xc4\x82\x15E\x17&2A\xa4*\xd3^" +
	"\xe8\x8a:\xecd\x82\xbbP\xdaj=\xf2\x9d\xcc&i" +
	"\x1dx\x17\xbe\xef\x97\xe7}\xde\xe7}\xbf\xb3\xd3\x12\xb7" +
	"\x19\xbb\xban\xb1 \xd4\xa8i\x85\x1f\xed\x91o\x8d/" +
	"}u\x14*f.\x86\xc3\x0f\xa5_\x7f\xe5\xd8\xca\x0b" +
	"\x00\x1d\xca\xe3\xceV\xb9\x05\xc8\x1aR\x12\x0c\x9f[\x18" +
	"<\xf1\xfd}33x)f\x8a6\xf4wQw\xd6" +
	"\x84F\x97E\x84\x8e\xee=\xfd\xda\x85?\x16f`\xdb" +
	"\x0c\x9f\xf9\xaegedj\xf5\x04L\x8d8\xe7\xc5;" +
	"\xce\xc5\xe8\xd7\xaf\xe2\x09\xb4\\\xab\x18[dG\xb8E" +
	"3\x0f\xcag\x1dO^\x0f\xec\x9e\x96/j\xe9\x9e\xec" +
	"\xbe\xbex}\xb9\x0e;\xc6&\x1dI\xef\x1e1\x05\x1d" +
	"e\xea\xff\xddcjm\xf7\x97\xfaS\xbd\xce\xd9S\x1b" +
	"`j\xe2M\xf3Kg.bg\xcd\x0c\x18\xe6x|" +
	"~\xe9\xe57\xbe\xd0I\x8c\xb5\xb5\xb7`~\xe6\x9c\xd3" +
	"dv\xd1\x8c\xda\xbb{_o\xe7\xc8\xf4\xb7\xdfh\xf4" +
	"\xda6t\xde<\xe6,D\xe8\xd7\x0d\xf4\xec\xabgn" +
	"\xfe\xf9\xba\xc4\x8f\xff\x95\xc4'\xe6\xfb\xce\xe7\x91\x83\xd3" +
	"\x91\xdb\xd9O\xcfu\x0e\xfe\xe0\xfe\xb6Y\xb6\xd7z\xd2" +
	"\xe9\xb3\xb4\xec\x0e+\x92\xdd\xfb\xf4\xc0\xb63\xa9\x07V" +
	"1\x1f3\x97\xda\x87q\x95u\xd4\xe9\x89\xd8\xa4f\x0f" +
	"\x86\x87\xab\xa5` \xef\x96\x19\x94\xd3\x13\xbe[\xabI" +
	"/\x18#\x99h\xc6\x03\xaeS\"(\xa7\x83R\xad0" +
	"Q\xc8\xbb\xb5B)\xc0\x18\x09\xd9&r\xbf\xebOz" +
	"\xa9\x81;\\\xdf\x1f#U\x874\x00\x83\x80\xddw\x17" +
	"\xa0n\x90T\xc3\x826\x99\xa4>\xdc\x95\x06\xd4\x8d\x92" +
	"jT0\x9c\x98\x0c\xf2\x91*\xc08\x04\xe3`\xa6\xec" +
	"V\xdcb\x95W\x80cR\xbbZO\x09\xd4\x87\x9bK" +
	"\x03\xaa\x83\xad\xd1n\x1dj\xd98\xb3?u\xa0\xe0\xf9" +
	"\x87\xba\xb5?\xb5M\x1a\xf10\x8c\xdc\xcd\xf6\x03\xeam" +
	"I\xf5\x81\xe0v\xfe\x1d&\x1a\xfe\xe6\xf6\x03\xea]I" +
	"\xf5\xb1\xe0v\xb1\x162I\x01\xd8\x1fj\xdb\xefI\xaa" +
	"S\x82]\xf2\xaf0I\x09\xd8'\xd3\xf6\xc9\x94Z\x94" +
	"T\x17\x04\xbb\x8c?\xc3$\x0d\xc0>?\x04\xa8\x9f$" +
	"\xc7)\xd8e\xae\x86I\x9a\x80\xbd\xa6%\x96%\xb3I" +
	"}l\xad\x84IZ\x80c\xb3\x1f\xc8\xc6)\x99\xbd\x86" +
	"\x82\xdd\xc1\xa4\xef\xc3:\x92+\x95|\xcf\x0dH\x08\x12" +
	"\xcc\x04\x93\xc5\x9cWa\x0c\x8210S\xadU\x0a\xc1" +
	"\xa3\xca\xa0\x08/>?x\xf5\x95\x8f\xcc\xd5\xa1\x0c\xc1" +
	"\xdb\x13d\x1c\xb0\xb9\xffH\x03y\xb8\x19m\xca\xadT" +
	"\xdc\xe9\xcb&\x9b)\xe5\x0e{\xf9Z\xf3~=\xd1\xc6" +
	"}w\xde\xf5}&\x9a\xd9\x82L\xb4\x0cD\x06\xe5\xf4" +
	"\x9d\x85j\xbeR(\x16\x02\xb7V\xaa\xdc[\xd6\xb3\xad" +
	"b\xc3V\xe8\xdcwH\xaa\x9d-[q\xd3\xf8\xa5\xad" +
	"\xb8U'\xe0\x16\xbd\x7f=\x87\x8f\xeb\x11\x1ft\x8b`" +
	"\xf3\xacu9\x0f\xe8\x15\xf6\x82F-Vu-\xa3Q" +
	"\x8b\xb4\xbbt\xe4\x1d\x92*)\x98)W\xbc\x89\xc2\xd4" +
	"%\x11\xd8\xec\x04\xda\x94\x0e\xb5z\x07\x1aob\xfd\xfb" +
	"\xb2a\xf3rn\xd5\xdb\xc3\xe1\xcb\xbc\x86\xcc@\xb4s" +
	"\xff\xa7\xf1\xa1\xe6\x1bik<\x155\xbeaJ\xadq" +
	"#\x13\x94\xd3\x8fyS\x9b-h\x99\xc8~\x1c\xe2\x9f" +
	"\x00\x00\x00\xff\xff\x06\xfd\x9d\x93"

func init() {
	schemas.Register(schema_8ef99297a43a5e34,
		0x82d3e852af0336bf,
		0xa0a054dea32fd98c,
		0xa0d9f6eca1c93d48,
		0xa3fa7845f919dd83,
		0xc2f8c20c293e5319,
		0xc4df13257bc2ea61,
		0xcfa794e8d19a0162,
		0xd7d879450a253e4b,
		0xe31026e735d69ddf,
		0xf061e22f0ae5c7b5,
		0xfa5b1fd61c2e7c3d)
}
