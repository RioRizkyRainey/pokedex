// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

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

// AttackHandlerClient is the client API for AttackHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AttackHandlerClient interface {
	GetDamage(ctx context.Context, in *Params, opts ...grpc.CallOption) (*Data, error)
}

type attackHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewAttackHandlerClient(cc grpc.ClientConnInterface) AttackHandlerClient {
	return &attackHandlerClient{cc}
}

func (c *attackHandlerClient) GetDamage(ctx context.Context, in *Params, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/attack.AttackHandler/GetDamage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttackHandlerServer is the server API for AttackHandler service.
// All implementations must embed UnimplementedAttackHandlerServer
// for forward compatibility
type AttackHandlerServer interface {
	GetDamage(context.Context, *Params) (*Data, error)
	mustEmbedUnimplementedAttackHandlerServer()
}

// UnimplementedAttackHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedAttackHandlerServer struct {
}

func (UnimplementedAttackHandlerServer) GetDamage(context.Context, *Params) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDamage not implemented")
}
func (UnimplementedAttackHandlerServer) mustEmbedUnimplementedAttackHandlerServer() {}

// UnsafeAttackHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AttackHandlerServer will
// result in compilation errors.
type UnsafeAttackHandlerServer interface {
	mustEmbedUnimplementedAttackHandlerServer()
}

func RegisterAttackHandlerServer(s grpc.ServiceRegistrar, srv AttackHandlerServer) {
	s.RegisterService(&AttackHandler_ServiceDesc, srv)
}

func _AttackHandler_GetDamage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Params)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttackHandlerServer).GetDamage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/attack.AttackHandler/GetDamage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttackHandlerServer).GetDamage(ctx, req.(*Params))
	}
	return interceptor(ctx, in, info, handler)
}

// AttackHandler_ServiceDesc is the grpc.ServiceDesc for AttackHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AttackHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "attack.AttackHandler",
	HandlerType: (*AttackHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDamage",
			Handler:    _AttackHandler_GetDamage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "attack.proto",
}