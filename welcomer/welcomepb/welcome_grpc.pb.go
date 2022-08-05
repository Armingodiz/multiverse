// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: welcomer/welcomepb/welcome.proto

package welcomepb

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WelcomeServiceClient is the client API for WelcomeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WelcomeServiceClient interface {
}

type welcomeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWelcomeServiceClient(cc grpc.ClientConnInterface) WelcomeServiceClient {
	return &welcomeServiceClient{cc}
}

// WelcomeServiceServer is the server API for WelcomeService service.
// All implementations must embed UnimplementedWelcomeServiceServer
// for forward compatibility
type WelcomeServiceServer interface {
	mustEmbedUnimplementedWelcomeServiceServer()
}

// UnimplementedWelcomeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWelcomeServiceServer struct {
}

func (UnimplementedWelcomeServiceServer) mustEmbedUnimplementedWelcomeServiceServer() {}

// UnsafeWelcomeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WelcomeServiceServer will
// result in compilation errors.
type UnsafeWelcomeServiceServer interface {
	mustEmbedUnimplementedWelcomeServiceServer()
}

func RegisterWelcomeServiceServer(s grpc.ServiceRegistrar, srv WelcomeServiceServer) {
	s.RegisterService(&WelcomeService_ServiceDesc, srv)
}

// WelcomeService_ServiceDesc is the grpc.ServiceDesc for WelcomeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WelcomeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "welcomer.WelcomeService",
	HandlerType: (*WelcomeServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "welcomer/welcomepb/welcome.proto",
}
