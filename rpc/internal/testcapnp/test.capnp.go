// Code generated by capnpc-go. DO NOT EDIT.

package testcapnp

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
	server "capnproto.org/go/capnp/v3/server"
	stream "capnproto.org/go/capnp/v3/std/capnp/stream"
	context "context"
)

type PingPong struct{ Client *capnp.Client }

// PingPong_TypeID is the unique identifier for the type PingPong.
const PingPong_TypeID = 0xf004c474c2f8ee7a

func (c PingPong) EchoNum(ctx context.Context, params func(PingPong_echoNum_Params) error) (PingPong_echoNum_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xf004c474c2f8ee7a,
			MethodID:      0,
			InterfaceName: "test.capnp:PingPong",
			MethodName:    "echoNum",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 0}
		s.PlaceArgs = func(s capnp.Struct) error { return params(PingPong_echoNum_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return PingPong_echoNum_Results_Future{Future: ans.Future()}, release
}

func (c PingPong) AddRef() PingPong {
	return PingPong{
		Client: c.Client.AddRef(),
	}
}

func (c PingPong) Release() {
	c.Client.Release()
}

// A PingPong_Server is a PingPong with a local implementation.
type PingPong_Server interface {
	EchoNum(context.Context, PingPong_echoNum) error
}

// PingPong_NewServer creates a new Server from an implementation of PingPong_Server.
func PingPong_NewServer(s PingPong_Server, policy *server.Policy) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(PingPong_Methods(nil, s), s, c, policy)
}

// PingPong_ServerToClient creates a new Client from an implementation of PingPong_Server.
// The caller is responsible for calling Release on the returned Client.
func PingPong_ServerToClient(s PingPong_Server, policy *server.Policy) PingPong {
	return PingPong{Client: capnp.NewClient(PingPong_NewServer(s, policy))}
}

// PingPong_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func PingPong_Methods(methods []server.Method, s PingPong_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xf004c474c2f8ee7a,
			MethodID:      0,
			InterfaceName: "test.capnp:PingPong",
			MethodName:    "echoNum",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.EchoNum(ctx, PingPong_echoNum{call})
		},
	})

	return methods
}

// PingPong_echoNum holds the state for a server call to PingPong.echoNum.
// See server.Call for documentation.
type PingPong_echoNum struct {
	*server.Call
}

// Args returns the call's arguments.
func (c PingPong_echoNum) Args() PingPong_echoNum_Params {
	return PingPong_echoNum_Params{Struct: c.Call.Args()}
}

// AllocResults allocates the results struct.
func (c PingPong_echoNum) AllocResults() (PingPong_echoNum_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return PingPong_echoNum_Results{Struct: r}, err
}

type PingPong_echoNum_Params struct{ capnp.Struct }

// PingPong_echoNum_Params_TypeID is the unique identifier for the type PingPong_echoNum_Params.
const PingPong_echoNum_Params_TypeID = 0xd797e0a99edf0921

func NewPingPong_echoNum_Params(s *capnp.Segment) (PingPong_echoNum_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return PingPong_echoNum_Params{st}, err
}

func NewRootPingPong_echoNum_Params(s *capnp.Segment) (PingPong_echoNum_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return PingPong_echoNum_Params{st}, err
}

func ReadRootPingPong_echoNum_Params(msg *capnp.Message) (PingPong_echoNum_Params, error) {
	root, err := msg.Root()
	return PingPong_echoNum_Params{root.Struct()}, err
}

func (s PingPong_echoNum_Params) String() string {
	str, _ := text.Marshal(0xd797e0a99edf0921, s.Struct)
	return str
}

func (s PingPong_echoNum_Params) N() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s PingPong_echoNum_Params) SetN(v int64) {
	s.Struct.SetUint64(0, uint64(v))
}

// PingPong_echoNum_Params_List is a list of PingPong_echoNum_Params.
type PingPong_echoNum_Params_List = capnp.StructList[PingPong_echoNum_Params]

// NewPingPong_echoNum_Params creates a new list of PingPong_echoNum_Params.
func NewPingPong_echoNum_Params_List(s *capnp.Segment, sz int32) (PingPong_echoNum_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return capnp.StructList[PingPong_echoNum_Params]{l}, err
}

// PingPong_echoNum_Params_Future is a wrapper for a PingPong_echoNum_Params promised by a client call.
type PingPong_echoNum_Params_Future struct{ *capnp.Future }

func (p PingPong_echoNum_Params_Future) Struct() (PingPong_echoNum_Params, error) {
	s, err := p.Future.Struct()
	return PingPong_echoNum_Params{s}, err
}

type PingPong_echoNum_Results struct{ capnp.Struct }

// PingPong_echoNum_Results_TypeID is the unique identifier for the type PingPong_echoNum_Results.
const PingPong_echoNum_Results_TypeID = 0x85ddfd96db252600

func NewPingPong_echoNum_Results(s *capnp.Segment) (PingPong_echoNum_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return PingPong_echoNum_Results{st}, err
}

func NewRootPingPong_echoNum_Results(s *capnp.Segment) (PingPong_echoNum_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return PingPong_echoNum_Results{st}, err
}

func ReadRootPingPong_echoNum_Results(msg *capnp.Message) (PingPong_echoNum_Results, error) {
	root, err := msg.Root()
	return PingPong_echoNum_Results{root.Struct()}, err
}

func (s PingPong_echoNum_Results) String() string {
	str, _ := text.Marshal(0x85ddfd96db252600, s.Struct)
	return str
}

func (s PingPong_echoNum_Results) N() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s PingPong_echoNum_Results) SetN(v int64) {
	s.Struct.SetUint64(0, uint64(v))
}

// PingPong_echoNum_Results_List is a list of PingPong_echoNum_Results.
type PingPong_echoNum_Results_List = capnp.StructList[PingPong_echoNum_Results]

// NewPingPong_echoNum_Results creates a new list of PingPong_echoNum_Results.
func NewPingPong_echoNum_Results_List(s *capnp.Segment, sz int32) (PingPong_echoNum_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return capnp.StructList[PingPong_echoNum_Results]{l}, err
}

// PingPong_echoNum_Results_Future is a wrapper for a PingPong_echoNum_Results promised by a client call.
type PingPong_echoNum_Results_Future struct{ *capnp.Future }

func (p PingPong_echoNum_Results_Future) Struct() (PingPong_echoNum_Results, error) {
	s, err := p.Future.Struct()
	return PingPong_echoNum_Results{s}, err
}

type StreamTest struct{ Client *capnp.Client }

// StreamTest_TypeID is the unique identifier for the type StreamTest.
const StreamTest_TypeID = 0xbb3ca85b01eea465

func (c StreamTest) Push(ctx context.Context, params func(StreamTest_push_Params) error) (stream.StreamResult_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xbb3ca85b01eea465,
			MethodID:      0,
			InterfaceName: "test.capnp:StreamTest",
			MethodName:    "push",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		s.PlaceArgs = func(s capnp.Struct) error { return params(StreamTest_push_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return stream.StreamResult_Future{Future: ans.Future()}, release
}

func (c StreamTest) AddRef() StreamTest {
	return StreamTest{
		Client: c.Client.AddRef(),
	}
}

func (c StreamTest) Release() {
	c.Client.Release()
}

// A StreamTest_Server is a StreamTest with a local implementation.
type StreamTest_Server interface {
	Push(context.Context, StreamTest_push) error
}

// StreamTest_NewServer creates a new Server from an implementation of StreamTest_Server.
func StreamTest_NewServer(s StreamTest_Server, policy *server.Policy) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(StreamTest_Methods(nil, s), s, c, policy)
}

// StreamTest_ServerToClient creates a new Client from an implementation of StreamTest_Server.
// The caller is responsible for calling Release on the returned Client.
func StreamTest_ServerToClient(s StreamTest_Server, policy *server.Policy) StreamTest {
	return StreamTest{Client: capnp.NewClient(StreamTest_NewServer(s, policy))}
}

// StreamTest_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func StreamTest_Methods(methods []server.Method, s StreamTest_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbb3ca85b01eea465,
			MethodID:      0,
			InterfaceName: "test.capnp:StreamTest",
			MethodName:    "push",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Push(ctx, StreamTest_push{call})
		},
	})

	return methods
}

// StreamTest_push holds the state for a server call to StreamTest.push.
// See server.Call for documentation.
type StreamTest_push struct {
	*server.Call
}

// Args returns the call's arguments.
func (c StreamTest_push) Args() StreamTest_push_Params {
	return StreamTest_push_Params{Struct: c.Call.Args()}
}

// AllocResults allocates the results struct.
func (c StreamTest_push) AllocResults() (stream.StreamResult, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return stream.StreamResult{Struct: r}, err
}

type StreamTest_push_Params struct{ capnp.Struct }

// StreamTest_push_Params_TypeID is the unique identifier for the type StreamTest_push_Params.
const StreamTest_push_Params_TypeID = 0xf838dca6c8721bdb

func NewStreamTest_push_Params(s *capnp.Segment) (StreamTest_push_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return StreamTest_push_Params{st}, err
}

func NewRootStreamTest_push_Params(s *capnp.Segment) (StreamTest_push_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return StreamTest_push_Params{st}, err
}

func ReadRootStreamTest_push_Params(msg *capnp.Message) (StreamTest_push_Params, error) {
	root, err := msg.Root()
	return StreamTest_push_Params{root.Struct()}, err
}

func (s StreamTest_push_Params) String() string {
	str, _ := text.Marshal(0xf838dca6c8721bdb, s.Struct)
	return str
}

func (s StreamTest_push_Params) Data() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s StreamTest_push_Params) HasData() bool {
	return s.Struct.HasPtr(0)
}

func (s StreamTest_push_Params) SetData(v []byte) error {
	return s.Struct.SetData(0, v)
}

// StreamTest_push_Params_List is a list of StreamTest_push_Params.
type StreamTest_push_Params_List = capnp.StructList[StreamTest_push_Params]

// NewStreamTest_push_Params creates a new list of StreamTest_push_Params.
func NewStreamTest_push_Params_List(s *capnp.Segment, sz int32) (StreamTest_push_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[StreamTest_push_Params]{l}, err
}

// StreamTest_push_Params_Future is a wrapper for a StreamTest_push_Params promised by a client call.
type StreamTest_push_Params_Future struct{ *capnp.Future }

func (p StreamTest_push_Params_Future) Struct() (StreamTest_push_Params, error) {
	s, err := p.Future.Struct()
	return StreamTest_push_Params{s}, err
}

type CapArgsTest struct{ Client *capnp.Client }

// CapArgsTest_TypeID is the unique identifier for the type CapArgsTest.
const CapArgsTest_TypeID = 0xb86bce7f916a10cc

func (c CapArgsTest) Call(ctx context.Context, params func(CapArgsTest_call_Params) error) (CapArgsTest_call_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xb86bce7f916a10cc,
			MethodID:      0,
			InterfaceName: "test.capnp:CapArgsTest",
			MethodName:    "call",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		s.PlaceArgs = func(s capnp.Struct) error { return params(CapArgsTest_call_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return CapArgsTest_call_Results_Future{Future: ans.Future()}, release
}
func (c CapArgsTest) Self(ctx context.Context, params func(CapArgsTest_self_Params) error) (CapArgsTest_self_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xb86bce7f916a10cc,
			MethodID:      1,
			InterfaceName: "test.capnp:CapArgsTest",
			MethodName:    "self",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		s.PlaceArgs = func(s capnp.Struct) error { return params(CapArgsTest_self_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return CapArgsTest_self_Results_Future{Future: ans.Future()}, release
}

func (c CapArgsTest) AddRef() CapArgsTest {
	return CapArgsTest{
		Client: c.Client.AddRef(),
	}
}

func (c CapArgsTest) Release() {
	c.Client.Release()
}

// A CapArgsTest_Server is a CapArgsTest with a local implementation.
type CapArgsTest_Server interface {
	Call(context.Context, CapArgsTest_call) error

	Self(context.Context, CapArgsTest_self) error
}

// CapArgsTest_NewServer creates a new Server from an implementation of CapArgsTest_Server.
func CapArgsTest_NewServer(s CapArgsTest_Server, policy *server.Policy) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(CapArgsTest_Methods(nil, s), s, c, policy)
}

// CapArgsTest_ServerToClient creates a new Client from an implementation of CapArgsTest_Server.
// The caller is responsible for calling Release on the returned Client.
func CapArgsTest_ServerToClient(s CapArgsTest_Server, policy *server.Policy) CapArgsTest {
	return CapArgsTest{Client: capnp.NewClient(CapArgsTest_NewServer(s, policy))}
}

// CapArgsTest_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func CapArgsTest_Methods(methods []server.Method, s CapArgsTest_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xb86bce7f916a10cc,
			MethodID:      0,
			InterfaceName: "test.capnp:CapArgsTest",
			MethodName:    "call",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Call(ctx, CapArgsTest_call{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xb86bce7f916a10cc,
			MethodID:      1,
			InterfaceName: "test.capnp:CapArgsTest",
			MethodName:    "self",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Self(ctx, CapArgsTest_self{call})
		},
	})

	return methods
}

// CapArgsTest_call holds the state for a server call to CapArgsTest.call.
// See server.Call for documentation.
type CapArgsTest_call struct {
	*server.Call
}

// Args returns the call's arguments.
func (c CapArgsTest_call) Args() CapArgsTest_call_Params {
	return CapArgsTest_call_Params{Struct: c.Call.Args()}
}

// AllocResults allocates the results struct.
func (c CapArgsTest_call) AllocResults() (CapArgsTest_call_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CapArgsTest_call_Results{Struct: r}, err
}

// CapArgsTest_self holds the state for a server call to CapArgsTest.self.
// See server.Call for documentation.
type CapArgsTest_self struct {
	*server.Call
}

// Args returns the call's arguments.
func (c CapArgsTest_self) Args() CapArgsTest_self_Params {
	return CapArgsTest_self_Params{Struct: c.Call.Args()}
}

// AllocResults allocates the results struct.
func (c CapArgsTest_self) AllocResults() (CapArgsTest_self_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return CapArgsTest_self_Results{Struct: r}, err
}

type CapArgsTest_call_Params struct{ capnp.Struct }

// CapArgsTest_call_Params_TypeID is the unique identifier for the type CapArgsTest_call_Params.
const CapArgsTest_call_Params_TypeID = 0x80087e4e698768a2

func NewCapArgsTest_call_Params(s *capnp.Segment) (CapArgsTest_call_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return CapArgsTest_call_Params{st}, err
}

func NewRootCapArgsTest_call_Params(s *capnp.Segment) (CapArgsTest_call_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return CapArgsTest_call_Params{st}, err
}

func ReadRootCapArgsTest_call_Params(msg *capnp.Message) (CapArgsTest_call_Params, error) {
	root, err := msg.Root()
	return CapArgsTest_call_Params{root.Struct()}, err
}

func (s CapArgsTest_call_Params) String() string {
	str, _ := text.Marshal(0x80087e4e698768a2, s.Struct)
	return str
}

func (s CapArgsTest_call_Params) Cap() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s CapArgsTest_call_Params) HasCap() bool {
	return s.Struct.HasPtr(0)
}

func (s CapArgsTest_call_Params) SetCap(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

// CapArgsTest_call_Params_List is a list of CapArgsTest_call_Params.
type CapArgsTest_call_Params_List = capnp.StructList[CapArgsTest_call_Params]

// NewCapArgsTest_call_Params creates a new list of CapArgsTest_call_Params.
func NewCapArgsTest_call_Params_List(s *capnp.Segment, sz int32) (CapArgsTest_call_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[CapArgsTest_call_Params]{l}, err
}

// CapArgsTest_call_Params_Future is a wrapper for a CapArgsTest_call_Params promised by a client call.
type CapArgsTest_call_Params_Future struct{ *capnp.Future }

func (p CapArgsTest_call_Params_Future) Struct() (CapArgsTest_call_Params, error) {
	s, err := p.Future.Struct()
	return CapArgsTest_call_Params{s}, err
}

func (p CapArgsTest_call_Params_Future) Cap() *capnp.Future {
	return p.Future.Field(0, nil)
}

type CapArgsTest_call_Results struct{ capnp.Struct }

// CapArgsTest_call_Results_TypeID is the unique identifier for the type CapArgsTest_call_Results.
const CapArgsTest_call_Results_TypeID = 0x96fbc50dc2f0200d

func NewCapArgsTest_call_Results(s *capnp.Segment) (CapArgsTest_call_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CapArgsTest_call_Results{st}, err
}

func NewRootCapArgsTest_call_Results(s *capnp.Segment) (CapArgsTest_call_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CapArgsTest_call_Results{st}, err
}

func ReadRootCapArgsTest_call_Results(msg *capnp.Message) (CapArgsTest_call_Results, error) {
	root, err := msg.Root()
	return CapArgsTest_call_Results{root.Struct()}, err
}

func (s CapArgsTest_call_Results) String() string {
	str, _ := text.Marshal(0x96fbc50dc2f0200d, s.Struct)
	return str
}

// CapArgsTest_call_Results_List is a list of CapArgsTest_call_Results.
type CapArgsTest_call_Results_List = capnp.StructList[CapArgsTest_call_Results]

// NewCapArgsTest_call_Results creates a new list of CapArgsTest_call_Results.
func NewCapArgsTest_call_Results_List(s *capnp.Segment, sz int32) (CapArgsTest_call_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[CapArgsTest_call_Results]{l}, err
}

// CapArgsTest_call_Results_Future is a wrapper for a CapArgsTest_call_Results promised by a client call.
type CapArgsTest_call_Results_Future struct{ *capnp.Future }

func (p CapArgsTest_call_Results_Future) Struct() (CapArgsTest_call_Results, error) {
	s, err := p.Future.Struct()
	return CapArgsTest_call_Results{s}, err
}

type CapArgsTest_self_Params struct{ capnp.Struct }

// CapArgsTest_self_Params_TypeID is the unique identifier for the type CapArgsTest_self_Params.
const CapArgsTest_self_Params_TypeID = 0xe2553e5a663abb7d

func NewCapArgsTest_self_Params(s *capnp.Segment) (CapArgsTest_self_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CapArgsTest_self_Params{st}, err
}

func NewRootCapArgsTest_self_Params(s *capnp.Segment) (CapArgsTest_self_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CapArgsTest_self_Params{st}, err
}

func ReadRootCapArgsTest_self_Params(msg *capnp.Message) (CapArgsTest_self_Params, error) {
	root, err := msg.Root()
	return CapArgsTest_self_Params{root.Struct()}, err
}

func (s CapArgsTest_self_Params) String() string {
	str, _ := text.Marshal(0xe2553e5a663abb7d, s.Struct)
	return str
}

// CapArgsTest_self_Params_List is a list of CapArgsTest_self_Params.
type CapArgsTest_self_Params_List = capnp.StructList[CapArgsTest_self_Params]

// NewCapArgsTest_self_Params creates a new list of CapArgsTest_self_Params.
func NewCapArgsTest_self_Params_List(s *capnp.Segment, sz int32) (CapArgsTest_self_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[CapArgsTest_self_Params]{l}, err
}

// CapArgsTest_self_Params_Future is a wrapper for a CapArgsTest_self_Params promised by a client call.
type CapArgsTest_self_Params_Future struct{ *capnp.Future }

func (p CapArgsTest_self_Params_Future) Struct() (CapArgsTest_self_Params, error) {
	s, err := p.Future.Struct()
	return CapArgsTest_self_Params{s}, err
}

type CapArgsTest_self_Results struct{ capnp.Struct }

// CapArgsTest_self_Results_TypeID is the unique identifier for the type CapArgsTest_self_Results.
const CapArgsTest_self_Results_TypeID = 0x9746cc05cbff1132

func NewCapArgsTest_self_Results(s *capnp.Segment) (CapArgsTest_self_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return CapArgsTest_self_Results{st}, err
}

func NewRootCapArgsTest_self_Results(s *capnp.Segment) (CapArgsTest_self_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return CapArgsTest_self_Results{st}, err
}

func ReadRootCapArgsTest_self_Results(msg *capnp.Message) (CapArgsTest_self_Results, error) {
	root, err := msg.Root()
	return CapArgsTest_self_Results{root.Struct()}, err
}

func (s CapArgsTest_self_Results) String() string {
	str, _ := text.Marshal(0x9746cc05cbff1132, s.Struct)
	return str
}

func (s CapArgsTest_self_Results) Self() CapArgsTest {
	p, _ := s.Struct.Ptr(0)
	return CapArgsTest{Client: p.Interface().Client()}
}

func (s CapArgsTest_self_Results) HasSelf() bool {
	return s.Struct.HasPtr(0)
}

func (s CapArgsTest_self_Results) SetSelf(v CapArgsTest) error {
	if !v.Client.IsValid() {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// CapArgsTest_self_Results_List is a list of CapArgsTest_self_Results.
type CapArgsTest_self_Results_List = capnp.StructList[CapArgsTest_self_Results]

// NewCapArgsTest_self_Results creates a new list of CapArgsTest_self_Results.
func NewCapArgsTest_self_Results_List(s *capnp.Segment, sz int32) (CapArgsTest_self_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[CapArgsTest_self_Results]{l}, err
}

// CapArgsTest_self_Results_Future is a wrapper for a CapArgsTest_self_Results promised by a client call.
type CapArgsTest_self_Results_Future struct{ *capnp.Future }

func (p CapArgsTest_self_Results_Future) Struct() (CapArgsTest_self_Results, error) {
	s, err := p.Future.Struct()
	return CapArgsTest_self_Results{s}, err
}

func (p CapArgsTest_self_Results_Future) Self() CapArgsTest {
	return CapArgsTest{Client: p.Future.Field(0, nil).Client()}
}

const schema_ef12a34b9807e19c = "x\xda\x94\x93Mh\xd3`\x18\xc7\x9f\xa7\xef\x9b\xa5C" +
	"B}\x93\xa1\xa8\x07ul\x82;\x14\xeb\x0e\xea\x10\xdb" +
	")8A\x18\x89\x1f\x07\xf5 \xa1fm\xb5MC\x93" +
	"\"\x0a:/\xb2\xb3\x97ME\x05\xe7\x07z\x9c \xa8" +
	"\xbb\xe8p\xa08\xbc\x89\xc8\x1c~\xdc\x9d;\x0eD#" +
	"\xef\x9bfilw\xf0V\xfa\x7f\xde\xff\xef\xff|d" +
	"\xc7\x05\xcc\xd1\x8c\xb2\xab\x03\x12\xc6!\xa9\xc3\x9f,\x8e" +
	"\x95\x86/'\xaf\x00[\x8f\x00\x12\xca\x00\xfdH6\"" +
	"\xa0\xd6I\xb2\x80\x7f\xb6\xf5\xce\x8f\xff^\xb8jt!" +
	"\x02P.o'\xdd\\\xcep\xd9W\xb6,\xcd(\xb3" +
	"\xbf\xc6\x83\xf7B7\xb8N\xfd\x9d\xcc\x7f'\xcd\x1d\x9c" +
	"hv\xde\x13<\x1d\x14O\xe7\xd6\x9e\xbd6\xfa\xfe\xdc" +
	"3`k\x88\x7f\xeb\x9b|\xfd\xf0=\xf5'\x00j&" +
	"y\xa1\x95\x88\x0c\xa0YdL{\xca\x7f\xf9\xd6\xfdE" +
	"<\xf5h\xeftK\xf1m\xf2D{ \x8a\xef\x92!" +
	"\xed\xad(\xde\xda\xf9\xe5\xce\xe3\xaf\x13\x1f\xa1)\xf5T" +
	"\xd0\xd4s\x81\xbe4=0rr\xdf\xf1\xefM\xa9?" +
	"q\x9d\xfa\x17\x17\x97g\xbc\xd7t\xa9\x85\xf3\x8aL\x06" +
	"\xee\xda,\x19\xd2~\x08\xce\xfc\xa6\xda\x9b\x87\x9fw/" +
	"\x03[\xb7\xd2\xe2\x07\xa2r\xce\x02\xc9\xc2\x09\xdf\xb3\\" +
	"/\x9d7\x1db;\x03\x07Lg\xb0Vp\x8f\x05\x7f" +
	"\x95\xcb=\xbaY3I\xc55(\xa1\x00\x14\x01\x98\xd2" +
	"\x0d`$\x09\x1a]\x09\x94\xf3\xa6\x83*%\x80\xa8\x02" +
	"\xc6\x9c\xf4\x92]\xd0\xabv!m\xe5\x8b\xd5\xe1z\xa5" +
	"\xe7\x88\xe5\xd6\xe5\xb2\x17\xb3R#+\xb4Q\x82\x04J" +
	"\xff\xd8\xb4\x04\x0amV\xadr\xad\xf2H[X_\x04" +
	"K\xf1\"d\xd1v\x01\x915\x911\xf4\xcc\x06\xa6:" +
	"\xa2\x91$\x12\xc0\xca)bxS,\xd3\x07\x09\xd6+" +
	"c\xb40\x0c\xaf\x8am\xe0\x9a\"\xa7x\xf2\\@\xcd" +
	"\xa1\x8eq\xd0Q\xaff\x99\x9b+!\x87\x0aN\xb85" +
	"\xb4\xa7^\x9e\xef\xbfy\xfa\x06c\xdcK\x92SN\xdd" +
	"-\xc6M\xda\x8e\xbb\xcd\xe2\xfeo\xdab\x8e\x0d\x97X" +
	"^\x81\x92\xabv!J\x1b\xde2B\xe3Kdl\xbf" +
	"H;\xda\x88\xd3\x1aXt-\x9aN\xf3\x8e\x04\xa9\x82" +
	"\xab.\xec\x8c\xe9\x99\xa8@\x02\x15\xc0\xbf\x01\x00\x00\xff" +
	"\xff\x19\x83 \xe4"

func init() {
	schemas.Register(schema_ef12a34b9807e19c,
		0x80087e4e698768a2,
		0x85ddfd96db252600,
		0x96fbc50dc2f0200d,
		0x9746cc05cbff1132,
		0xb86bce7f916a10cc,
		0xbb3ca85b01eea465,
		0xd797e0a99edf0921,
		0xe2553e5a663abb7d,
		0xf004c474c2f8ee7a,
		0xf838dca6c8721bdb)
}
