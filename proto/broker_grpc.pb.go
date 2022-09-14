// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: broker.proto

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

// BrokerClient is the client API for Broker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BrokerClient interface {
	Register(ctx context.Context, opts ...grpc.CallOption) (Broker_RegisterClient, error)
	Providers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Broker_ProvidersClient, error)
}

type brokerClient struct {
	cc grpc.ClientConnInterface
}

func NewBrokerClient(cc grpc.ClientConnInterface) BrokerClient {
	return &brokerClient{cc}
}

func (c *brokerClient) Register(ctx context.Context, opts ...grpc.CallOption) (Broker_RegisterClient, error) {
	stream, err := c.cc.NewStream(ctx, &Broker_ServiceDesc.Streams[0], "/service.Broker/Register", opts...)
	if err != nil {
		return nil, err
	}
	x := &brokerRegisterClient{stream}
	return x, nil
}

type Broker_RegisterClient interface {
	Send(*Ping) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type brokerRegisterClient struct {
	grpc.ClientStream
}

func (x *brokerRegisterClient) Send(m *Ping) error {
	return x.ClientStream.SendMsg(m)
}

func (x *brokerRegisterClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *brokerClient) Providers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Broker_ProvidersClient, error) {
	stream, err := c.cc.NewStream(ctx, &Broker_ServiceDesc.Streams[1], "/service.Broker/Providers", opts...)
	if err != nil {
		return nil, err
	}
	x := &brokerProvidersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Broker_ProvidersClient interface {
	Recv() (*ProviderList, error)
	grpc.ClientStream
}

type brokerProvidersClient struct {
	grpc.ClientStream
}

func (x *brokerProvidersClient) Recv() (*ProviderList, error) {
	m := new(ProviderList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BrokerServer is the server API for Broker service.
// All implementations must embed UnimplementedBrokerServer
// for forward compatibility
type BrokerServer interface {
	Register(Broker_RegisterServer) error
	Providers(*emptypb.Empty, Broker_ProvidersServer) error
	mustEmbedUnimplementedBrokerServer()
}

// UnimplementedBrokerServer must be embedded to have forward compatible implementations.
type UnimplementedBrokerServer struct {
}

func (UnimplementedBrokerServer) Register(Broker_RegisterServer) error {
	return status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedBrokerServer) Providers(*emptypb.Empty, Broker_ProvidersServer) error {
	return status.Errorf(codes.Unimplemented, "method Providers not implemented")
}
func (UnimplementedBrokerServer) mustEmbedUnimplementedBrokerServer() {}

// UnsafeBrokerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BrokerServer will
// result in compilation errors.
type UnsafeBrokerServer interface {
	mustEmbedUnimplementedBrokerServer()
}

func RegisterBrokerServer(s grpc.ServiceRegistrar, srv BrokerServer) {
	s.RegisterService(&Broker_ServiceDesc, srv)
}

func _Broker_Register_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BrokerServer).Register(&brokerRegisterServer{stream})
}

type Broker_RegisterServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*Ping, error)
	grpc.ServerStream
}

type brokerRegisterServer struct {
	grpc.ServerStream
}

func (x *brokerRegisterServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *brokerRegisterServer) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Broker_Providers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BrokerServer).Providers(m, &brokerProvidersServer{stream})
}

type Broker_ProvidersServer interface {
	Send(*ProviderList) error
	grpc.ServerStream
}

type brokerProvidersServer struct {
	grpc.ServerStream
}

func (x *brokerProvidersServer) Send(m *ProviderList) error {
	return x.ServerStream.SendMsg(m)
}

// Broker_ServiceDesc is the grpc.ServiceDesc for Broker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Broker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Broker",
	HandlerType: (*BrokerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Register",
			Handler:       _Broker_Register_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Providers",
			Handler:       _Broker_Providers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "broker.proto",
}
