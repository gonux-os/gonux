// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: god.proto

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

// GodClient is the client API for God service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GodClient interface {
	Register(ctx context.Context, in *ActionRegisterRequest, opts ...grpc.CallOption) (*Empty, error)
	Call(ctx context.Context, in *ActionIdRequest, opts ...grpc.CallOption) (*Empty, error)
	WaitFor(ctx context.Context, in *ActionIdRequest, opts ...grpc.CallOption) (*Empty, error)
	Subscribe(ctx context.Context, in *Empty, opts ...grpc.CallOption) (God_SubscribeClient, error)
}

type godClient struct {
	cc grpc.ClientConnInterface
}

func NewGodClient(cc grpc.ClientConnInterface) GodClient {
	return &godClient{cc}
}

func (c *godClient) Register(ctx context.Context, in *ActionRegisterRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.God/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *godClient) Call(ctx context.Context, in *ActionIdRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.God/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *godClient) WaitFor(ctx context.Context, in *ActionIdRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.God/WaitFor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *godClient) Subscribe(ctx context.Context, in *Empty, opts ...grpc.CallOption) (God_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &God_ServiceDesc.Streams[0], "/proto.God/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &godSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type God_SubscribeClient interface {
	Recv() (*ActionCallEvent, error)
	grpc.ClientStream
}

type godSubscribeClient struct {
	grpc.ClientStream
}

func (x *godSubscribeClient) Recv() (*ActionCallEvent, error) {
	m := new(ActionCallEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GodServer is the server API for God service.
// All implementations must embed UnimplementedGodServer
// for forward compatibility
type GodServer interface {
	Register(context.Context, *ActionRegisterRequest) (*Empty, error)
	Call(context.Context, *ActionIdRequest) (*Empty, error)
	WaitFor(context.Context, *ActionIdRequest) (*Empty, error)
	Subscribe(*Empty, God_SubscribeServer) error
	mustEmbedUnimplementedGodServer()
}

// UnimplementedGodServer must be embedded to have forward compatible implementations.
type UnimplementedGodServer struct {
}

func (UnimplementedGodServer) Register(context.Context, *ActionRegisterRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedGodServer) Call(context.Context, *ActionIdRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedGodServer) WaitFor(context.Context, *ActionIdRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitFor not implemented")
}
func (UnimplementedGodServer) Subscribe(*Empty, God_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedGodServer) mustEmbedUnimplementedGodServer() {}

// UnsafeGodServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GodServer will
// result in compilation errors.
type UnsafeGodServer interface {
	mustEmbedUnimplementedGodServer()
}

func RegisterGodServer(s grpc.ServiceRegistrar, srv GodServer) {
	s.RegisterService(&God_ServiceDesc, srv)
}

func _God_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GodServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.God/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GodServer).Register(ctx, req.(*ActionRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _God_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GodServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.God/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GodServer).Call(ctx, req.(*ActionIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _God_WaitFor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GodServer).WaitFor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.God/WaitFor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GodServer).WaitFor(ctx, req.(*ActionIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _God_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GodServer).Subscribe(m, &godSubscribeServer{stream})
}

type God_SubscribeServer interface {
	Send(*ActionCallEvent) error
	grpc.ServerStream
}

type godSubscribeServer struct {
	grpc.ServerStream
}

func (x *godSubscribeServer) Send(m *ActionCallEvent) error {
	return x.ServerStream.SendMsg(m)
}

// God_ServiceDesc is the grpc.ServiceDesc for God service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var God_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.God",
	HandlerType: (*GodServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _God_Register_Handler,
		},
		{
			MethodName: "Call",
			Handler:    _God_Call_Handler,
		},
		{
			MethodName: "WaitFor",
			Handler:    _God_WaitFor_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _God_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "god.proto",
}
