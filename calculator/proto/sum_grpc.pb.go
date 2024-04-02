// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: sum.proto

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

// SumServiceClient is the client API for SumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SumServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	Primes(ctx context.Context, in *PrimesRequest, opts ...grpc.CallOption) (SumService_PrimesClient, error)
	Average(ctx context.Context, opts ...grpc.CallOption) (SumService_AverageClient, error)
}

type sumServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSumServiceClient(cc grpc.ClientConnInterface) SumServiceClient {
	return &sumServiceClient{cc}
}

func (c *sumServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/sum.SumService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) Primes(ctx context.Context, in *PrimesRequest, opts ...grpc.CallOption) (SumService_PrimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &SumService_ServiceDesc.Streams[0], "/sum.SumService/Primes", opts...)
	if err != nil {
		return nil, err
	}
	x := &sumServicePrimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SumService_PrimesClient interface {
	Recv() (*PrimesResponse, error)
	grpc.ClientStream
}

type sumServicePrimesClient struct {
	grpc.ClientStream
}

func (x *sumServicePrimesClient) Recv() (*PrimesResponse, error) {
	m := new(PrimesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sumServiceClient) Average(ctx context.Context, opts ...grpc.CallOption) (SumService_AverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &SumService_ServiceDesc.Streams[1], "/sum.SumService/Average", opts...)
	if err != nil {
		return nil, err
	}
	x := &sumServiceAverageClient{stream}
	return x, nil
}

type SumService_AverageClient interface {
	Send(*AverageRequest) error
	CloseAndRecv() (*AverageResponse, error)
	grpc.ClientStream
}

type sumServiceAverageClient struct {
	grpc.ClientStream
}

func (x *sumServiceAverageClient) Send(m *AverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sumServiceAverageClient) CloseAndRecv() (*AverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SumServiceServer is the server API for SumService service.
// All implementations must embed UnimplementedSumServiceServer
// for forward compatibility
type SumServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	Primes(*PrimesRequest, SumService_PrimesServer) error
	Average(SumService_AverageServer) error
	mustEmbedUnimplementedSumServiceServer()
}

// UnimplementedSumServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSumServiceServer struct {
}

func (UnimplementedSumServiceServer) Sum(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedSumServiceServer) Primes(*PrimesRequest, SumService_PrimesServer) error {
	return status.Errorf(codes.Unimplemented, "method Primes not implemented")
}
func (UnimplementedSumServiceServer) Average(SumService_AverageServer) error {
	return status.Errorf(codes.Unimplemented, "method Average not implemented")
}
func (UnimplementedSumServiceServer) mustEmbedUnimplementedSumServiceServer() {}

// UnsafeSumServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SumServiceServer will
// result in compilation errors.
type UnsafeSumServiceServer interface {
	mustEmbedUnimplementedSumServiceServer()
}

func RegisterSumServiceServer(s grpc.ServiceRegistrar, srv SumServiceServer) {
	s.RegisterService(&SumService_ServiceDesc, srv)
}

func _SumService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_Primes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SumServiceServer).Primes(m, &sumServicePrimesServer{stream})
}

type SumService_PrimesServer interface {
	Send(*PrimesResponse) error
	grpc.ServerStream
}

type sumServicePrimesServer struct {
	grpc.ServerStream
}

func (x *sumServicePrimesServer) Send(m *PrimesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SumService_Average_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SumServiceServer).Average(&sumServiceAverageServer{stream})
}

type SumService_AverageServer interface {
	SendAndClose(*AverageResponse) error
	Recv() (*AverageRequest, error)
	grpc.ServerStream
}

type sumServiceAverageServer struct {
	grpc.ServerStream
}

func (x *sumServiceAverageServer) SendAndClose(m *AverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sumServiceAverageServer) Recv() (*AverageRequest, error) {
	m := new(AverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SumService_ServiceDesc is the grpc.ServiceDesc for SumService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SumService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sum.SumService",
	HandlerType: (*SumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _SumService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Primes",
			Handler:       _SumService_Primes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Average",
			Handler:       _SumService_Average_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "sum.proto",
}
