// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: proto/connect/agent_service/v1alpha1/agent_service.proto

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
	AgentService_UpdateTrustZoneBundle_FullMethodName      = "/proto.connect.agent_service.v1alpha1.AgentService/UpdateTrustZoneBundle"
	AgentService_UpdateTrustZoneAuthority_FullMethodName   = "/proto.connect.agent_service.v1alpha1.AgentService/UpdateTrustZoneAuthority"
	AgentService_UpdateAgentStatus_FullMethodName          = "/proto.connect.agent_service.v1alpha1.AgentService/UpdateAgentStatus"
	AgentService_RegisterFederatedService_FullMethodName   = "/proto.connect.agent_service.v1alpha1.AgentService/RegisterFederatedService"
	AgentService_DeregisterFederatedService_FullMethodName = "/proto.connect.agent_service.v1alpha1.AgentService/DeregisterFederatedService"
)

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentServiceClient interface {
	UpdateTrustZoneBundle(ctx context.Context, in *UpdateTrustZoneBundleRequest, opts ...grpc.CallOption) (*UpdateTrustZoneBundleResponse, error)
	UpdateTrustZoneAuthority(ctx context.Context, in *UpdateTrustZoneAuthorityRequest, opts ...grpc.CallOption) (*UpdateTrustZoneAuthorityResponse, error)
	UpdateAgentStatus(ctx context.Context, in *UpdateAgentStatusRequest, opts ...grpc.CallOption) (*UpdateAgentStatusResponse, error)
	RegisterFederatedService(ctx context.Context, in *RegisterFederatedServiceRequest, opts ...grpc.CallOption) (*RegisterFederatedServiceResponse, error)
	DeregisterFederatedService(ctx context.Context, in *DeregisterFederatedServiceRequest, opts ...grpc.CallOption) (*DeregisterFederatedServiceResponse, error)
}

type agentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentServiceClient(cc grpc.ClientConnInterface) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) UpdateTrustZoneBundle(ctx context.Context, in *UpdateTrustZoneBundleRequest, opts ...grpc.CallOption) (*UpdateTrustZoneBundleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTrustZoneBundleResponse)
	err := c.cc.Invoke(ctx, AgentService_UpdateTrustZoneBundle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) UpdateTrustZoneAuthority(ctx context.Context, in *UpdateTrustZoneAuthorityRequest, opts ...grpc.CallOption) (*UpdateTrustZoneAuthorityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTrustZoneAuthorityResponse)
	err := c.cc.Invoke(ctx, AgentService_UpdateTrustZoneAuthority_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) UpdateAgentStatus(ctx context.Context, in *UpdateAgentStatusRequest, opts ...grpc.CallOption) (*UpdateAgentStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateAgentStatusResponse)
	err := c.cc.Invoke(ctx, AgentService_UpdateAgentStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) RegisterFederatedService(ctx context.Context, in *RegisterFederatedServiceRequest, opts ...grpc.CallOption) (*RegisterFederatedServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterFederatedServiceResponse)
	err := c.cc.Invoke(ctx, AgentService_RegisterFederatedService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) DeregisterFederatedService(ctx context.Context, in *DeregisterFederatedServiceRequest, opts ...grpc.CallOption) (*DeregisterFederatedServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeregisterFederatedServiceResponse)
	err := c.cc.Invoke(ctx, AgentService_DeregisterFederatedService_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServiceServer is the server API for AgentService service.
// All implementations should embed UnimplementedAgentServiceServer
// for forward compatibility.
type AgentServiceServer interface {
	UpdateTrustZoneBundle(context.Context, *UpdateTrustZoneBundleRequest) (*UpdateTrustZoneBundleResponse, error)
	UpdateTrustZoneAuthority(context.Context, *UpdateTrustZoneAuthorityRequest) (*UpdateTrustZoneAuthorityResponse, error)
	UpdateAgentStatus(context.Context, *UpdateAgentStatusRequest) (*UpdateAgentStatusResponse, error)
	RegisterFederatedService(context.Context, *RegisterFederatedServiceRequest) (*RegisterFederatedServiceResponse, error)
	DeregisterFederatedService(context.Context, *DeregisterFederatedServiceRequest) (*DeregisterFederatedServiceResponse, error)
}

// UnimplementedAgentServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAgentServiceServer struct{}

func (UnimplementedAgentServiceServer) UpdateTrustZoneBundle(context.Context, *UpdateTrustZoneBundleRequest) (*UpdateTrustZoneBundleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTrustZoneBundle not implemented")
}
func (UnimplementedAgentServiceServer) UpdateTrustZoneAuthority(context.Context, *UpdateTrustZoneAuthorityRequest) (*UpdateTrustZoneAuthorityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTrustZoneAuthority not implemented")
}
func (UnimplementedAgentServiceServer) UpdateAgentStatus(context.Context, *UpdateAgentStatusRequest) (*UpdateAgentStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAgentStatus not implemented")
}
func (UnimplementedAgentServiceServer) RegisterFederatedService(context.Context, *RegisterFederatedServiceRequest) (*RegisterFederatedServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterFederatedService not implemented")
}
func (UnimplementedAgentServiceServer) DeregisterFederatedService(context.Context, *DeregisterFederatedServiceRequest) (*DeregisterFederatedServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeregisterFederatedService not implemented")
}
func (UnimplementedAgentServiceServer) testEmbeddedByValue() {}

// UnsafeAgentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServiceServer will
// result in compilation errors.
type UnsafeAgentServiceServer interface {
	mustEmbedUnimplementedAgentServiceServer()
}

func RegisterAgentServiceServer(s grpc.ServiceRegistrar, srv AgentServiceServer) {
	// If the following call pancis, it indicates UnimplementedAgentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AgentService_ServiceDesc, srv)
}

func _AgentService_UpdateTrustZoneBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTrustZoneBundleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).UpdateTrustZoneBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_UpdateTrustZoneBundle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).UpdateTrustZoneBundle(ctx, req.(*UpdateTrustZoneBundleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_UpdateTrustZoneAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTrustZoneAuthorityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).UpdateTrustZoneAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_UpdateTrustZoneAuthority_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).UpdateTrustZoneAuthority(ctx, req.(*UpdateTrustZoneAuthorityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_UpdateAgentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAgentStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).UpdateAgentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_UpdateAgentStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).UpdateAgentStatus(ctx, req.(*UpdateAgentStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_RegisterFederatedService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterFederatedServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).RegisterFederatedService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_RegisterFederatedService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).RegisterFederatedService(ctx, req.(*RegisterFederatedServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_DeregisterFederatedService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeregisterFederatedServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).DeregisterFederatedService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_DeregisterFederatedService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).DeregisterFederatedService(ctx, req.(*DeregisterFederatedServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentService_ServiceDesc is the grpc.ServiceDesc for AgentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.connect.agent_service.v1alpha1.AgentService",
	HandlerType: (*AgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateTrustZoneBundle",
			Handler:    _AgentService_UpdateTrustZoneBundle_Handler,
		},
		{
			MethodName: "UpdateTrustZoneAuthority",
			Handler:    _AgentService_UpdateTrustZoneAuthority_Handler,
		},
		{
			MethodName: "UpdateAgentStatus",
			Handler:    _AgentService_UpdateAgentStatus_Handler,
		},
		{
			MethodName: "RegisterFederatedService",
			Handler:    _AgentService_RegisterFederatedService_Handler,
		},
		{
			MethodName: "DeregisterFederatedService",
			Handler:    _AgentService_DeregisterFederatedService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/connect/agent_service/v1alpha1/agent_service.proto",
}
