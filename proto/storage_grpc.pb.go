// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageClient interface {
	Pull(ctx context.Context, in *StorageRef, opts ...grpc.CallOption) (Storage_PullClient, error)
	Push(ctx context.Context, opts ...grpc.CallOption) (Storage_PushClient, error)
}

type storageClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageClient(cc grpc.ClientConnInterface) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) Pull(ctx context.Context, in *StorageRef, opts ...grpc.CallOption) (Storage_PullClient, error) {
	stream, err := c.cc.NewStream(ctx, &Storage_ServiceDesc.Streams[0], "/service.Storage/Pull", opts...)
	if err != nil {
		return nil, err
	}
	x := &storagePullClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Storage_PullClient interface {
	Recv() (*StorageParcel, error)
	grpc.ClientStream
}

type storagePullClient struct {
	grpc.ClientStream
}

func (x *storagePullClient) Recv() (*StorageParcel, error) {
	m := new(StorageParcel)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storageClient) Push(ctx context.Context, opts ...grpc.CallOption) (Storage_PushClient, error) {
	stream, err := c.cc.NewStream(ctx, &Storage_ServiceDesc.Streams[1], "/service.Storage/Push", opts...)
	if err != nil {
		return nil, err
	}
	x := &storagePushClient{stream}
	return x, nil
}

type Storage_PushClient interface {
	Send(*StorageParcel) error
	CloseAndRecv() (*StorageRef, error)
	grpc.ClientStream
}

type storagePushClient struct {
	grpc.ClientStream
}

func (x *storagePushClient) Send(m *StorageParcel) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storagePushClient) CloseAndRecv() (*StorageRef, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StorageRef)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StorageServer is the server API for Storage service.
// All implementations must embed UnimplementedStorageServer
// for forward compatibility
type StorageServer interface {
	Pull(*StorageRef, Storage_PullServer) error
	Push(Storage_PushServer) error
	mustEmbedUnimplementedStorageServer()
}

// UnimplementedStorageServer must be embedded to have forward compatible implementations.
type UnimplementedStorageServer struct {
}

func (UnimplementedStorageServer) Pull(*StorageRef, Storage_PullServer) error {
	return status.Errorf(codes.Unimplemented, "method Pull not implemented")
}
func (UnimplementedStorageServer) Push(Storage_PushServer) error {
	return status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (UnimplementedStorageServer) mustEmbedUnimplementedStorageServer() {}

// UnsafeStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServer will
// result in compilation errors.
type UnsafeStorageServer interface {
	mustEmbedUnimplementedStorageServer()
}

func RegisterStorageServer(s grpc.ServiceRegistrar, srv StorageServer) {
	s.RegisterService(&Storage_ServiceDesc, srv)
}

func _Storage_Pull_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StorageRef)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StorageServer).Pull(m, &storagePullServer{stream})
}

type Storage_PullServer interface {
	Send(*StorageParcel) error
	grpc.ServerStream
}

type storagePullServer struct {
	grpc.ServerStream
}

func (x *storagePullServer) Send(m *StorageParcel) error {
	return x.ServerStream.SendMsg(m)
}

func _Storage_Push_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StorageServer).Push(&storagePushServer{stream})
}

type Storage_PushServer interface {
	SendAndClose(*StorageRef) error
	Recv() (*StorageParcel, error)
	grpc.ServerStream
}

type storagePushServer struct {
	grpc.ServerStream
}

func (x *storagePushServer) SendAndClose(m *StorageRef) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storagePushServer) Recv() (*StorageParcel, error) {
	m := new(StorageParcel)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Storage_ServiceDesc is the grpc.ServiceDesc for Storage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Storage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Pull",
			Handler:       _Storage_Pull_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Push",
			Handler:       _Storage_Push_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "storage.proto",
}
