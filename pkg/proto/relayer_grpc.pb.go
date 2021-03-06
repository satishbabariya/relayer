// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: relayer.proto

package proto

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RelayerClient is the client API for Relayer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RelayerClient interface {
}

type relayerClient struct {
	cc grpc.ClientConnInterface
}

func NewRelayerClient(cc grpc.ClientConnInterface) RelayerClient {
	return &relayerClient{cc}
}

// RelayerServer is the server API for Relayer service.
// All implementations must embed UnimplementedRelayerServer
// for forward compatibility
type RelayerServer interface {
	mustEmbedUnimplementedRelayerServer()
}

// UnimplementedRelayerServer must be embedded to have forward compatible implementations.
type UnimplementedRelayerServer struct {
}

func (UnimplementedRelayerServer) mustEmbedUnimplementedRelayerServer() {}

// UnsafeRelayerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelayerServer will
// result in compilation errors.
type UnsafeRelayerServer interface {
	mustEmbedUnimplementedRelayerServer()
}

func RegisterRelayerServer(s grpc.ServiceRegistrar, srv RelayerServer) {
	s.RegisterService(&Relayer_ServiceDesc, srv)
}

// Relayer_ServiceDesc is the grpc.ServiceDesc for Relayer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Relayer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "relayer.Relayer",
	HandlerType: (*RelayerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "relayer.proto",
}
