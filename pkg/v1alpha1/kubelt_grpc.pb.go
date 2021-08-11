// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1alpha1

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

// KubeltClient is the client API for Kubelt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KubeltClient interface {
	HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*ApplyResponse, error)
}

type kubeltClient struct {
	cc grpc.ClientConnInterface
}

func NewKubeltClient(cc grpc.ClientConnInterface) KubeltClient {
	return &kubeltClient{cc}
}

func (c *kubeltClient) HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/kubelt.Kubelt/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeltClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/kubelt.Kubelt/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeltClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/kubelt.Kubelt/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubeltClient) Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*ApplyResponse, error) {
	out := new(ApplyResponse)
	err := c.cc.Invoke(ctx, "/kubelt.Kubelt/Apply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KubeltServer is the server API for Kubelt service.
// All implementations must embed UnimplementedKubeltServer
// for forward compatibility
type KubeltServer interface {
	HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	Apply(context.Context, *ApplyRequest) (*ApplyResponse, error)
	mustEmbedUnimplementedKubeltServer()
}

// UnimplementedKubeltServer must be embedded to have forward compatible implementations.
type UnimplementedKubeltServer struct {
}

func (UnimplementedKubeltServer) HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedKubeltServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedKubeltServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedKubeltServer) Apply(context.Context, *ApplyRequest) (*ApplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Apply not implemented")
}
func (UnimplementedKubeltServer) mustEmbedUnimplementedKubeltServer() {}

// UnsafeKubeltServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KubeltServer will
// result in compilation errors.
type UnsafeKubeltServer interface {
	mustEmbedUnimplementedKubeltServer()
}

func RegisterKubeltServer(s grpc.ServiceRegistrar, srv KubeltServer) {
	s.RegisterService(&Kubelt_ServiceDesc, srv)
}

func _Kubelt_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeltServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubelt.Kubelt/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeltServer).HealthCheck(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubelt_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeltServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubelt.Kubelt/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeltServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubelt_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeltServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubelt.Kubelt/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeltServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubelt_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeltServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kubelt.Kubelt/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeltServer).Apply(ctx, req.(*ApplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Kubelt_ServiceDesc is the grpc.ServiceDesc for Kubelt service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Kubelt_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kubelt.Kubelt",
	HandlerType: (*KubeltServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _Kubelt_HealthCheck_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Kubelt_Get_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _Kubelt_Query_Handler,
		},
		{
			MethodName: "Apply",
			Handler:    _Kubelt_Apply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kubelt.proto",
}
