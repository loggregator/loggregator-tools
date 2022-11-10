// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: loggregator-api/v2/ingress.proto

package loggregator_v2

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

// IngressClient is the client API for Ingress service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IngressClient interface {
	Sender(ctx context.Context, opts ...grpc.CallOption) (Ingress_SenderClient, error)
	BatchSender(ctx context.Context, opts ...grpc.CallOption) (Ingress_BatchSenderClient, error)
	Send(ctx context.Context, in *EnvelopeBatch, opts ...grpc.CallOption) (*SendResponse, error)
}

type ingressClient struct {
	cc grpc.ClientConnInterface
}

func NewIngressClient(cc grpc.ClientConnInterface) IngressClient {
	return &ingressClient{cc}
}

func (c *ingressClient) Sender(ctx context.Context, opts ...grpc.CallOption) (Ingress_SenderClient, error) {
	stream, err := c.cc.NewStream(ctx, &Ingress_ServiceDesc.Streams[0], "/loggregator.v2.Ingress/Sender", opts...)
	if err != nil {
		return nil, err
	}
	x := &ingressSenderClient{stream}
	return x, nil
}

type Ingress_SenderClient interface {
	Send(*Envelope) error
	CloseAndRecv() (*IngressResponse, error)
	grpc.ClientStream
}

type ingressSenderClient struct {
	grpc.ClientStream
}

func (x *ingressSenderClient) Send(m *Envelope) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ingressSenderClient) CloseAndRecv() (*IngressResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(IngressResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ingressClient) BatchSender(ctx context.Context, opts ...grpc.CallOption) (Ingress_BatchSenderClient, error) {
	stream, err := c.cc.NewStream(ctx, &Ingress_ServiceDesc.Streams[1], "/loggregator.v2.Ingress/BatchSender", opts...)
	if err != nil {
		return nil, err
	}
	x := &ingressBatchSenderClient{stream}
	return x, nil
}

type Ingress_BatchSenderClient interface {
	Send(*EnvelopeBatch) error
	CloseAndRecv() (*BatchSenderResponse, error)
	grpc.ClientStream
}

type ingressBatchSenderClient struct {
	grpc.ClientStream
}

func (x *ingressBatchSenderClient) Send(m *EnvelopeBatch) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ingressBatchSenderClient) CloseAndRecv() (*BatchSenderResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(BatchSenderResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ingressClient) Send(ctx context.Context, in *EnvelopeBatch, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/loggregator.v2.Ingress/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IngressServer is the server API for Ingress service.
// All implementations must embed UnimplementedIngressServer
// for forward compatibility
type IngressServer interface {
	Sender(Ingress_SenderServer) error
	BatchSender(Ingress_BatchSenderServer) error
	Send(context.Context, *EnvelopeBatch) (*SendResponse, error)
	mustEmbedUnimplementedIngressServer()
}

// UnimplementedIngressServer must be embedded to have forward compatible implementations.
type UnimplementedIngressServer struct {
}

func (UnimplementedIngressServer) Sender(Ingress_SenderServer) error {
	return status.Errorf(codes.Unimplemented, "method Sender not implemented")
}
func (UnimplementedIngressServer) BatchSender(Ingress_BatchSenderServer) error {
	return status.Errorf(codes.Unimplemented, "method BatchSender not implemented")
}
func (UnimplementedIngressServer) Send(context.Context, *EnvelopeBatch) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedIngressServer) mustEmbedUnimplementedIngressServer() {}

// UnsafeIngressServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IngressServer will
// result in compilation errors.
type UnsafeIngressServer interface {
	mustEmbedUnimplementedIngressServer()
}

func RegisterIngressServer(s grpc.ServiceRegistrar, srv IngressServer) {
	s.RegisterService(&Ingress_ServiceDesc, srv)
}

func _Ingress_Sender_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IngressServer).Sender(&ingressSenderServer{stream})
}

type Ingress_SenderServer interface {
	SendAndClose(*IngressResponse) error
	Recv() (*Envelope, error)
	grpc.ServerStream
}

type ingressSenderServer struct {
	grpc.ServerStream
}

func (x *ingressSenderServer) SendAndClose(m *IngressResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ingressSenderServer) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Ingress_BatchSender_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IngressServer).BatchSender(&ingressBatchSenderServer{stream})
}

type Ingress_BatchSenderServer interface {
	SendAndClose(*BatchSenderResponse) error
	Recv() (*EnvelopeBatch, error)
	grpc.ServerStream
}

type ingressBatchSenderServer struct {
	grpc.ServerStream
}

func (x *ingressBatchSenderServer) SendAndClose(m *BatchSenderResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ingressBatchSenderServer) Recv() (*EnvelopeBatch, error) {
	m := new(EnvelopeBatch)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Ingress_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnvelopeBatch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IngressServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loggregator.v2.Ingress/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IngressServer).Send(ctx, req.(*EnvelopeBatch))
	}
	return interceptor(ctx, in, info, handler)
}

// Ingress_ServiceDesc is the grpc.ServiceDesc for Ingress service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ingress_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "loggregator.v2.Ingress",
	HandlerType: (*IngressServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Ingress_Send_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Sender",
			Handler:       _Ingress_Sender_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BatchSender",
			Handler:       _Ingress_BatchSender_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "loggregator-api/v2/ingress.proto",
}
