// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: proto/connect/federation_service/federation_service.proto

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
	FederationService_CreatedFederation_FullMethodName = "/proto.connect.federation_service.v1alpha1.FederationService/CreatedFederation"
	FederationService_ListFederations_FullMethodName   = "/proto.connect.federation_service.v1alpha1.FederationService/ListFederations"
	FederationService_GetFederation_FullMethodName     = "/proto.connect.federation_service.v1alpha1.FederationService/GetFederation"
	FederationService_DeleteFederation_FullMethodName  = "/proto.connect.federation_service.v1alpha1.FederationService/DeleteFederation"
)

// FederationServiceClient is the client API for FederationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FederationServiceClient interface {
	CreatedFederation(ctx context.Context, in *CreateFederationRequest, opts ...grpc.CallOption) (*CreateFederationResponse, error)
	ListFederations(ctx context.Context, in *ListFederationsRequest, opts ...grpc.CallOption) (*ListFederationsResponse, error)
	GetFederation(ctx context.Context, in *GetFederationRequest, opts ...grpc.CallOption) (*GetFederationResponse, error)
	DeleteFederation(ctx context.Context, in *DeleteFederationRequest, opts ...grpc.CallOption) (*DeleteFederationResponse, error)
}

type federationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFederationServiceClient(cc grpc.ClientConnInterface) FederationServiceClient {
	return &federationServiceClient{cc}
}

func (c *federationServiceClient) CreatedFederation(ctx context.Context, in *CreateFederationRequest, opts ...grpc.CallOption) (*CreateFederationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateFederationResponse)
	err := c.cc.Invoke(ctx, FederationService_CreatedFederation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federationServiceClient) ListFederations(ctx context.Context, in *ListFederationsRequest, opts ...grpc.CallOption) (*ListFederationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListFederationsResponse)
	err := c.cc.Invoke(ctx, FederationService_ListFederations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federationServiceClient) GetFederation(ctx context.Context, in *GetFederationRequest, opts ...grpc.CallOption) (*GetFederationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFederationResponse)
	err := c.cc.Invoke(ctx, FederationService_GetFederation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federationServiceClient) DeleteFederation(ctx context.Context, in *DeleteFederationRequest, opts ...grpc.CallOption) (*DeleteFederationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteFederationResponse)
	err := c.cc.Invoke(ctx, FederationService_DeleteFederation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FederationServiceServer is the server API for FederationService service.
// All implementations should embed UnimplementedFederationServiceServer
// for forward compatibility.
type FederationServiceServer interface {
	CreatedFederation(context.Context, *CreateFederationRequest) (*CreateFederationResponse, error)
	ListFederations(context.Context, *ListFederationsRequest) (*ListFederationsResponse, error)
	GetFederation(context.Context, *GetFederationRequest) (*GetFederationResponse, error)
	DeleteFederation(context.Context, *DeleteFederationRequest) (*DeleteFederationResponse, error)
}

// UnimplementedFederationServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFederationServiceServer struct{}

func (UnimplementedFederationServiceServer) CreatedFederation(context.Context, *CreateFederationRequest) (*CreateFederationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatedFederation not implemented")
}
func (UnimplementedFederationServiceServer) ListFederations(context.Context, *ListFederationsRequest) (*ListFederationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFederations not implemented")
}
func (UnimplementedFederationServiceServer) GetFederation(context.Context, *GetFederationRequest) (*GetFederationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFederation not implemented")
}
func (UnimplementedFederationServiceServer) DeleteFederation(context.Context, *DeleteFederationRequest) (*DeleteFederationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFederation not implemented")
}
func (UnimplementedFederationServiceServer) testEmbeddedByValue() {}

// UnsafeFederationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FederationServiceServer will
// result in compilation errors.
type UnsafeFederationServiceServer interface {
	mustEmbedUnimplementedFederationServiceServer()
}

func RegisterFederationServiceServer(s grpc.ServiceRegistrar, srv FederationServiceServer) {
	// If the following call pancis, it indicates UnimplementedFederationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FederationService_ServiceDesc, srv)
}

func _FederationService_CreatedFederation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFederationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederationServiceServer).CreatedFederation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FederationService_CreatedFederation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederationServiceServer).CreatedFederation(ctx, req.(*CreateFederationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederationService_ListFederations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFederationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederationServiceServer).ListFederations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FederationService_ListFederations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederationServiceServer).ListFederations(ctx, req.(*ListFederationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederationService_GetFederation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFederationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederationServiceServer).GetFederation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FederationService_GetFederation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederationServiceServer).GetFederation(ctx, req.(*GetFederationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederationService_DeleteFederation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFederationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederationServiceServer).DeleteFederation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FederationService_DeleteFederation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederationServiceServer).DeleteFederation(ctx, req.(*DeleteFederationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FederationService_ServiceDesc is the grpc.ServiceDesc for FederationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FederationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.connect.federation_service.v1alpha1.FederationService",
	HandlerType: (*FederationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatedFederation",
			Handler:    _FederationService_CreatedFederation_Handler,
		},
		{
			MethodName: "ListFederations",
			Handler:    _FederationService_ListFederations_Handler,
		},
		{
			MethodName: "GetFederation",
			Handler:    _FederationService_GetFederation_Handler,
		},
		{
			MethodName: "DeleteFederation",
			Handler:    _FederationService_DeleteFederation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/connect/federation_service/federation_service.proto",
}
