// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.5.0
// source: Zhaoh0326.proto

package stubs

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

// Zhaoh0326Client is the client API for Zhaoh0326 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Zhaoh0326Client interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type zhaoh0326Client struct {
	cc grpc.ClientConnInterface
}

func NewZhaoh0326Client(cc grpc.ClientConnInterface) Zhaoh0326Client {
	return &zhaoh0326Client{cc}
}

func (c *zhaoh0326Client) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/sample_module_0326.sample_package_0326.Zhaoh0326/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Zhaoh0326Server is the server API for Zhaoh0326 service.
// All implementations must embed UnimplementedZhaoh0326Server
// for forward compatibility
type Zhaoh0326Server interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedZhaoh0326Server()
}

// UnimplementedZhaoh0326Server must be embedded to have forward compatible implementations.
type UnimplementedZhaoh0326Server struct {
}

func (UnimplementedZhaoh0326Server) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedZhaoh0326Server) mustEmbedUnimplementedZhaoh0326Server() {}

// UnsafeZhaoh0326Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Zhaoh0326Server will
// result in compilation errors.
type UnsafeZhaoh0326Server interface {
	mustEmbedUnimplementedZhaoh0326Server()
}

func RegisterZhaoh0326Server(s grpc.ServiceRegistrar, srv Zhaoh0326Server) {
	s.RegisterService(&Zhaoh0326_ServiceDesc, srv)
}

func _Zhaoh0326_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Zhaoh0326Server).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sample_module_0326.sample_package_0326.Zhaoh0326/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Zhaoh0326Server).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Zhaoh0326_ServiceDesc is the grpc.ServiceDesc for Zhaoh0326 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Zhaoh0326_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sample_module_0326.sample_package_0326.Zhaoh0326",
	HandlerType: (*Zhaoh0326Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Zhaoh0326_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Zhaoh0326.proto",
}
