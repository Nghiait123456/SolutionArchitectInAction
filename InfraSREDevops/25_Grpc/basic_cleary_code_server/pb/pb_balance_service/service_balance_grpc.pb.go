// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.8.0
// source: service_balance.proto

package pb_balance_service

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

// BalanceHandleClient is the client API for BalanceHandle service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BalanceHandleClient interface {
	RequestBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error)
}

type balanceHandleClient struct {
	cc grpc.ClientConnInterface
}

func NewBalanceHandleClient(cc grpc.ClientConnInterface) BalanceHandleClient {
	return &balanceHandleClient{cc}
}

func (c *balanceHandleClient) RequestBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error) {
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, "/balance_service.BalanceHandle/RequestBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BalanceHandleServer is the server API for BalanceHandle service.
// All implementations must embed UnimplementedBalanceHandleServer
// for forward compatibility
type BalanceHandleServer interface {
	RequestBalance(context.Context, *BalanceRequest) (*BalanceResponse, error)
	mustEmbedUnimplementedBalanceHandleServer()
}

// UnimplementedBalanceHandleServer must be embedded to have forward compatible implementations.
type UnimplementedBalanceHandleServer struct {
}

func (UnimplementedBalanceHandleServer) RequestBalance(context.Context, *BalanceRequest) (*BalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestBalance not implemented")
}
func (UnimplementedBalanceHandleServer) mustEmbedUnimplementedBalanceHandleServer() {}

// UnsafeBalanceHandleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BalanceHandleServer will
// result in compilation errors.
type UnsafeBalanceHandleServer interface {
	mustEmbedUnimplementedBalanceHandleServer()
}

func RegisterBalanceHandleServer(s grpc.ServiceRegistrar, srv BalanceHandleServer) {
	s.RegisterService(&BalanceHandle_ServiceDesc, srv)
}

func _BalanceHandle_RequestBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceHandleServer).RequestBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/balance_service.BalanceHandle/RequestBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceHandleServer).RequestBalance(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BalanceHandle_ServiceDesc is the grpc.ServiceDesc for BalanceHandle service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BalanceHandle_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "balance_service.BalanceHandle",
	HandlerType: (*BalanceHandleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestBalance",
			Handler:    _BalanceHandle_RequestBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_balance.proto",
}