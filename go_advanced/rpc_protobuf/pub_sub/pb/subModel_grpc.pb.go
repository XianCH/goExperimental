// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: subModel.proto

package pb

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

const (
	PubsubService_Publish_FullMethodName   = "/PubsubService/Publish"
	PubsubService_Subscribe_FullMethodName = "/PubsubService/Subscribe"
)

// PubsubServiceClient is the client API for PubsubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PubsubServiceClient interface {
	Publish(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Subscribe(ctx context.Context, in *Request, opts ...grpc.CallOption) (PubsubService_SubscribeClient, error)
}

type pubsubServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPubsubServiceClient(cc grpc.ClientConnInterface) PubsubServiceClient {
	return &pubsubServiceClient{cc}
}

func (c *pubsubServiceClient) Publish(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, PubsubService_Publish_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubsubServiceClient) Subscribe(ctx context.Context, in *Request, opts ...grpc.CallOption) (PubsubService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &PubsubService_ServiceDesc.Streams[0], PubsubService_Subscribe_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &pubsubServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PubsubService_SubscribeClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type pubsubServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *pubsubServiceSubscribeClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PubsubServiceServer is the server API for PubsubService service.
// All implementations must embed UnimplementedPubsubServiceServer
// for forward compatibility
type PubsubServiceServer interface {
	Publish(context.Context, *Request) (*Response, error)
	Subscribe(*Request, PubsubService_SubscribeServer) error
	mustEmbedUnimplementedPubsubServiceServer()
}

// UnimplementedPubsubServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPubsubServiceServer struct {
}

func (UnimplementedPubsubServiceServer) Publish(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedPubsubServiceServer) Subscribe(*Request, PubsubService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedPubsubServiceServer) mustEmbedUnimplementedPubsubServiceServer() {}

// UnsafePubsubServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PubsubServiceServer will
// result in compilation errors.
type UnsafePubsubServiceServer interface {
	mustEmbedUnimplementedPubsubServiceServer()
}

func RegisterPubsubServiceServer(s grpc.ServiceRegistrar, srv PubsubServiceServer) {
	s.RegisterService(&PubsubService_ServiceDesc, srv)
}

func _PubsubService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubsubServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PubsubService_Publish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubsubServiceServer).Publish(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubsubService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PubsubServiceServer).Subscribe(m, &pubsubServiceSubscribeServer{stream})
}

type PubsubService_SubscribeServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type pubsubServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *pubsubServiceSubscribeServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

// PubsubService_ServiceDesc is the grpc.ServiceDesc for PubsubService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PubsubService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PubsubService",
	HandlerType: (*PubsubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _PubsubService_Publish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _PubsubService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "subModel.proto",
}
