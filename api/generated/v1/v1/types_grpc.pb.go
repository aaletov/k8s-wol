// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: types.proto

package v1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WolControllerClient is the client API for WolController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WolControllerClient interface {
	WakeUp(ctx context.Context, in *WakeUpRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type wolControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewWolControllerClient(cc grpc.ClientConnInterface) WolControllerClient {
	return &wolControllerClient{cc}
}

func (c *wolControllerClient) WakeUp(ctx context.Context, in *WakeUpRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/v1.WolController/WakeUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WolControllerServer is the server API for WolController service.
// All implementations must embed UnimplementedWolControllerServer
// for forward compatibility
type WolControllerServer interface {
	WakeUp(context.Context, *WakeUpRequest) (*empty.Empty, error)
	mustEmbedUnimplementedWolControllerServer()
}

// UnimplementedWolControllerServer must be embedded to have forward compatible implementations.
type UnimplementedWolControllerServer struct {
}

func (UnimplementedWolControllerServer) WakeUp(context.Context, *WakeUpRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WakeUp not implemented")
}
func (UnimplementedWolControllerServer) mustEmbedUnimplementedWolControllerServer() {}

// UnsafeWolControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WolControllerServer will
// result in compilation errors.
type UnsafeWolControllerServer interface {
	mustEmbedUnimplementedWolControllerServer()
}

func RegisterWolControllerServer(s grpc.ServiceRegistrar, srv WolControllerServer) {
	s.RegisterService(&WolController_ServiceDesc, srv)
}

func _WolController_WakeUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WakeUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WolControllerServer).WakeUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.WolController/WakeUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WolControllerServer).WakeUp(ctx, req.(*WakeUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WolController_ServiceDesc is the grpc.ServiceDesc for WolController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WolController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.WolController",
	HandlerType: (*WolControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WakeUp",
			Handler:    _WolController_WakeUp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "types.proto",
}