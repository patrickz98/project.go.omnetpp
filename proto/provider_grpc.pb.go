// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: provider.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProviderClient is the client API for Provider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProviderClient interface {
	Info(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ProviderInfo, error)
	Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Utilization, error)
	Allocate(ctx context.Context, opts ...grpc.CallOption) (Provider_AllocateClient, error)
	GetSession(ctx context.Context, in *Simulation, opts ...grpc.CallOption) (*Session, error)
	SetSession(ctx context.Context, in *Session, opts ...grpc.CallOption) (*Session, error)
	// Todo: close session properly
	// rpc CloseSession (Session)      returns (google.protobuf.Empty);
	Extract(ctx context.Context, in *Bundle, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Compile(ctx context.Context, in *Simulation, opts ...grpc.CallOption) (*Binary, error)
	ListRunNums(ctx context.Context, in *Simulation, opts ...grpc.CallOption) (*SimulationRunList, error)
	Run(ctx context.Context, in *SimulationRun, opts ...grpc.CallOption) (*StorageRef, error)
}

type providerClient struct {
	cc grpc.ClientConnInterface
}

func NewProviderClient(cc grpc.ClientConnInterface) ProviderClient {
	return &providerClient{cc}
}

func (c *providerClient) Info(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ProviderInfo, error) {
	out := new(ProviderInfo)
	err := c.cc.Invoke(ctx, "/service.Provider/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Utilization, error) {
	out := new(Utilization)
	err := c.cc.Invoke(ctx, "/service.Provider/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) Allocate(ctx context.Context, opts ...grpc.CallOption) (Provider_AllocateClient, error) {
	stream, err := c.cc.NewStream(ctx, &Provider_ServiceDesc.Streams[0], "/service.Provider/Allocate", opts...)
	if err != nil {
		return nil, err
	}
	x := &providerAllocateClient{stream}
	return x, nil
}

type Provider_AllocateClient interface {
	Send(*FreeSlot) error
	Recv() (*AllocateSlot, error)
	grpc.ClientStream
}

type providerAllocateClient struct {
	grpc.ClientStream
}

func (x *providerAllocateClient) Send(m *FreeSlot) error {
	return x.ClientStream.SendMsg(m)
}

func (x *providerAllocateClient) Recv() (*AllocateSlot, error) {
	m := new(AllocateSlot)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *providerClient) GetSession(ctx context.Context, in *Simulation, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/service.Provider/GetSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) SetSession(ctx context.Context, in *Session, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/service.Provider/SetSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) Extract(ctx context.Context, in *Bundle, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/service.Provider/Extract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) Compile(ctx context.Context, in *Simulation, opts ...grpc.CallOption) (*Binary, error) {
	out := new(Binary)
	err := c.cc.Invoke(ctx, "/service.Provider/Compile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) ListRunNums(ctx context.Context, in *Simulation, opts ...grpc.CallOption) (*SimulationRunList, error) {
	out := new(SimulationRunList)
	err := c.cc.Invoke(ctx, "/service.Provider/ListRunNums", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) Run(ctx context.Context, in *SimulationRun, opts ...grpc.CallOption) (*StorageRef, error) {
	out := new(StorageRef)
	err := c.cc.Invoke(ctx, "/service.Provider/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProviderServer is the server API for Provider service.
// All implementations must embed UnimplementedProviderServer
// for forward compatibility
type ProviderServer interface {
	Info(context.Context, *emptypb.Empty) (*ProviderInfo, error)
	Status(context.Context, *emptypb.Empty) (*Utilization, error)
	Allocate(Provider_AllocateServer) error
	GetSession(context.Context, *Simulation) (*Session, error)
	SetSession(context.Context, *Session) (*Session, error)
	// Todo: close session properly
	// rpc CloseSession (Session)      returns (google.protobuf.Empty);
	Extract(context.Context, *Bundle) (*emptypb.Empty, error)
	Compile(context.Context, *Simulation) (*Binary, error)
	ListRunNums(context.Context, *Simulation) (*SimulationRunList, error)
	Run(context.Context, *SimulationRun) (*StorageRef, error)
	mustEmbedUnimplementedProviderServer()
}

// UnimplementedProviderServer must be embedded to have forward compatible implementations.
type UnimplementedProviderServer struct {
}

func (UnimplementedProviderServer) Info(context.Context, *emptypb.Empty) (*ProviderInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedProviderServer) Status(context.Context, *emptypb.Empty) (*Utilization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedProviderServer) Allocate(Provider_AllocateServer) error {
	return status.Errorf(codes.Unimplemented, "method Allocate not implemented")
}
func (UnimplementedProviderServer) GetSession(context.Context, *Simulation) (*Session, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSession not implemented")
}
func (UnimplementedProviderServer) SetSession(context.Context, *Session) (*Session, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSession not implemented")
}
func (UnimplementedProviderServer) Extract(context.Context, *Bundle) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Extract not implemented")
}
func (UnimplementedProviderServer) Compile(context.Context, *Simulation) (*Binary, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Compile not implemented")
}
func (UnimplementedProviderServer) ListRunNums(context.Context, *Simulation) (*SimulationRunList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRunNums not implemented")
}
func (UnimplementedProviderServer) Run(context.Context, *SimulationRun) (*StorageRef, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (UnimplementedProviderServer) mustEmbedUnimplementedProviderServer() {}

// UnsafeProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProviderServer will
// result in compilation errors.
type UnsafeProviderServer interface {
	mustEmbedUnimplementedProviderServer()
}

func RegisterProviderServer(s grpc.ServiceRegistrar, srv ProviderServer) {
	s.RegisterService(&Provider_ServiceDesc, srv)
}

func _Provider_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Provider/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).Info(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Provider/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).Status(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_Allocate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProviderServer).Allocate(&providerAllocateServer{stream})
}

type Provider_AllocateServer interface {
	Send(*AllocateSlot) error
	Recv() (*FreeSlot, error)
	grpc.ServerStream
}

type providerAllocateServer struct {
	grpc.ServerStream
}

func (x *providerAllocateServer) Send(m *AllocateSlot) error {
	return x.ServerStream.SendMsg(m)
}

func (x *providerAllocateServer) Recv() (*FreeSlot, error) {
	m := new(FreeSlot)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Provider_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Simulation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Provider/GetSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).GetSession(ctx, req.(*Simulation))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_SetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).SetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Provider/SetSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).SetSession(ctx, req.(*Session))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_Extract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bundle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).Extract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Provider/Extract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).Extract(ctx, req.(*Bundle))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_Compile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Simulation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).Compile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Provider/Compile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).Compile(ctx, req.(*Simulation))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_ListRunNums_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Simulation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).ListRunNums(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Provider/ListRunNums",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).ListRunNums(ctx, req.(*Simulation))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimulationRun)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Provider/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).Run(ctx, req.(*SimulationRun))
	}
	return interceptor(ctx, in, info, handler)
}

// Provider_ServiceDesc is the grpc.ServiceDesc for Provider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Provider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Provider",
	HandlerType: (*ProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _Provider_Info_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Provider_Status_Handler,
		},
		{
			MethodName: "GetSession",
			Handler:    _Provider_GetSession_Handler,
		},
		{
			MethodName: "SetSession",
			Handler:    _Provider_SetSession_Handler,
		},
		{
			MethodName: "Extract",
			Handler:    _Provider_Extract_Handler,
		},
		{
			MethodName: "Compile",
			Handler:    _Provider_Compile_Handler,
		},
		{
			MethodName: "ListRunNums",
			Handler:    _Provider_ListRunNums_Handler,
		},
		{
			MethodName: "Run",
			Handler:    _Provider_Run_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Allocate",
			Handler:       _Provider_Allocate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "provider.proto",
}
