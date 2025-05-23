// Copyright 2024 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: proto/provision_plugin/v1alpha1/plugin.proto

package v1alpha1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ProvisionPluginService_Validate_FullMethodName      = "/proto.provision_plugin.v1alpha1.ProvisionPluginService/Validate"
	ProvisionPluginService_Deploy_FullMethodName        = "/proto.provision_plugin.v1alpha1.ProvisionPluginService/Deploy"
	ProvisionPluginService_TearDown_FullMethodName      = "/proto.provision_plugin.v1alpha1.ProvisionPluginService/TearDown"
	ProvisionPluginService_GetHelmValues_FullMethodName = "/proto.provision_plugin.v1alpha1.ProvisionPluginService/GetHelmValues"
)

// ProvisionPluginServiceClient is the client API for ProvisionPluginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProvisionPluginServiceClient interface {
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
	Deploy(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DeployResponse], error)
	TearDown(ctx context.Context, in *TearDownRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[TearDownResponse], error)
	GetHelmValues(ctx context.Context, in *GetHelmValuesRequest, opts ...grpc.CallOption) (*GetHelmValuesResponse, error)
}

type provisionPluginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProvisionPluginServiceClient(cc grpc.ClientConnInterface) ProvisionPluginServiceClient {
	return &provisionPluginServiceClient{cc}
}

func (c *provisionPluginServiceClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidateResponse)
	err := c.cc.Invoke(ctx, ProvisionPluginService_Validate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *provisionPluginServiceClient) Deploy(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DeployResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ProvisionPluginService_ServiceDesc.Streams[0], ProvisionPluginService_Deploy_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[DeployRequest, DeployResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ProvisionPluginService_DeployClient = grpc.ServerStreamingClient[DeployResponse]

func (c *provisionPluginServiceClient) TearDown(ctx context.Context, in *TearDownRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[TearDownResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ProvisionPluginService_ServiceDesc.Streams[1], ProvisionPluginService_TearDown_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[TearDownRequest, TearDownResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ProvisionPluginService_TearDownClient = grpc.ServerStreamingClient[TearDownResponse]

func (c *provisionPluginServiceClient) GetHelmValues(ctx context.Context, in *GetHelmValuesRequest, opts ...grpc.CallOption) (*GetHelmValuesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetHelmValuesResponse)
	err := c.cc.Invoke(ctx, ProvisionPluginService_GetHelmValues_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProvisionPluginServiceServer is the server API for ProvisionPluginService service.
// All implementations should embed UnimplementedProvisionPluginServiceServer
// for forward compatibility.
type ProvisionPluginServiceServer interface {
	Validate(context.Context, *ValidateRequest) (*ValidateResponse, error)
	Deploy(*DeployRequest, grpc.ServerStreamingServer[DeployResponse]) error
	TearDown(*TearDownRequest, grpc.ServerStreamingServer[TearDownResponse]) error
	GetHelmValues(context.Context, *GetHelmValuesRequest) (*GetHelmValuesResponse, error)
}

// UnimplementedProvisionPluginServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProvisionPluginServiceServer struct{}

func (UnimplementedProvisionPluginServiceServer) Validate(context.Context, *ValidateRequest) (*ValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedProvisionPluginServiceServer) Deploy(*DeployRequest, grpc.ServerStreamingServer[DeployResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Deploy not implemented")
}
func (UnimplementedProvisionPluginServiceServer) TearDown(*TearDownRequest, grpc.ServerStreamingServer[TearDownResponse]) error {
	return status.Errorf(codes.Unimplemented, "method TearDown not implemented")
}
func (UnimplementedProvisionPluginServiceServer) GetHelmValues(context.Context, *GetHelmValuesRequest) (*GetHelmValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHelmValues not implemented")
}
func (UnimplementedProvisionPluginServiceServer) testEmbeddedByValue() {}

// UnsafeProvisionPluginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProvisionPluginServiceServer will
// result in compilation errors.
type UnsafeProvisionPluginServiceServer interface {
	mustEmbedUnimplementedProvisionPluginServiceServer()
}

func RegisterProvisionPluginServiceServer(s grpc.ServiceRegistrar, srv ProvisionPluginServiceServer) {
	// If the following call pancis, it indicates UnimplementedProvisionPluginServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProvisionPluginService_ServiceDesc, srv)
}

func _ProvisionPluginService_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisionPluginServiceServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProvisionPluginService_Validate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisionPluginServiceServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProvisionPluginService_Deploy_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DeployRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProvisionPluginServiceServer).Deploy(m, &grpc.GenericServerStream[DeployRequest, DeployResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ProvisionPluginService_DeployServer = grpc.ServerStreamingServer[DeployResponse]

func _ProvisionPluginService_TearDown_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TearDownRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProvisionPluginServiceServer).TearDown(m, &grpc.GenericServerStream[TearDownRequest, TearDownResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ProvisionPluginService_TearDownServer = grpc.ServerStreamingServer[TearDownResponse]

func _ProvisionPluginService_GetHelmValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHelmValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisionPluginServiceServer).GetHelmValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProvisionPluginService_GetHelmValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisionPluginServiceServer).GetHelmValues(ctx, req.(*GetHelmValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProvisionPluginService_ServiceDesc is the grpc.ServiceDesc for ProvisionPluginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProvisionPluginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.provision_plugin.v1alpha1.ProvisionPluginService",
	HandlerType: (*ProvisionPluginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Validate",
			Handler:    _ProvisionPluginService_Validate_Handler,
		},
		{
			MethodName: "GetHelmValues",
			Handler:    _ProvisionPluginService_GetHelmValues_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Deploy",
			Handler:       _ProvisionPluginService_Deploy_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "TearDown",
			Handler:       _ProvisionPluginService_TearDown_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/provision_plugin/v1alpha1/plugin.proto",
}
