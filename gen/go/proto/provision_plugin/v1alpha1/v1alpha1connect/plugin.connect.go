// Copyright 2024 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/provision_plugin/v1alpha1/plugin.proto

package v1alpha1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha1 "github.com/cofide/cofide-api-sdk/gen/go/proto/provision_plugin/v1alpha1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ProvisionPluginServiceName is the fully-qualified name of the ProvisionPluginService service.
	ProvisionPluginServiceName = "proto.provision_plugin.v1alpha1.ProvisionPluginService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ProvisionPluginServiceValidateProcedure is the fully-qualified name of the
	// ProvisionPluginService's Validate RPC.
	ProvisionPluginServiceValidateProcedure = "/proto.provision_plugin.v1alpha1.ProvisionPluginService/Validate"
	// ProvisionPluginServiceDeployProcedure is the fully-qualified name of the ProvisionPluginService's
	// Deploy RPC.
	ProvisionPluginServiceDeployProcedure = "/proto.provision_plugin.v1alpha1.ProvisionPluginService/Deploy"
	// ProvisionPluginServiceTearDownProcedure is the fully-qualified name of the
	// ProvisionPluginService's TearDown RPC.
	ProvisionPluginServiceTearDownProcedure = "/proto.provision_plugin.v1alpha1.ProvisionPluginService/TearDown"
)

// ProvisionPluginServiceClient is a client for the
// proto.provision_plugin.v1alpha1.ProvisionPluginService service.
type ProvisionPluginServiceClient interface {
	Validate(context.Context, *connect.Request[v1alpha1.ValidateRequest]) (*connect.Response[v1alpha1.ValidateResponse], error)
	Deploy(context.Context, *connect.Request[v1alpha1.DeployRequest]) (*connect.ServerStreamForClient[v1alpha1.DeployResponse], error)
	TearDown(context.Context, *connect.Request[v1alpha1.TearDownRequest]) (*connect.ServerStreamForClient[v1alpha1.TearDownResponse], error)
}

// NewProvisionPluginServiceClient constructs a client for the
// proto.provision_plugin.v1alpha1.ProvisionPluginService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewProvisionPluginServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ProvisionPluginServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	provisionPluginServiceMethods := v1alpha1.File_proto_provision_plugin_v1alpha1_plugin_proto.Services().ByName("ProvisionPluginService").Methods()
	return &provisionPluginServiceClient{
		validate: connect.NewClient[v1alpha1.ValidateRequest, v1alpha1.ValidateResponse](
			httpClient,
			baseURL+ProvisionPluginServiceValidateProcedure,
			connect.WithSchema(provisionPluginServiceMethods.ByName("Validate")),
			connect.WithClientOptions(opts...),
		),
		deploy: connect.NewClient[v1alpha1.DeployRequest, v1alpha1.DeployResponse](
			httpClient,
			baseURL+ProvisionPluginServiceDeployProcedure,
			connect.WithSchema(provisionPluginServiceMethods.ByName("Deploy")),
			connect.WithClientOptions(opts...),
		),
		tearDown: connect.NewClient[v1alpha1.TearDownRequest, v1alpha1.TearDownResponse](
			httpClient,
			baseURL+ProvisionPluginServiceTearDownProcedure,
			connect.WithSchema(provisionPluginServiceMethods.ByName("TearDown")),
			connect.WithClientOptions(opts...),
		),
	}
}

// provisionPluginServiceClient implements ProvisionPluginServiceClient.
type provisionPluginServiceClient struct {
	validate *connect.Client[v1alpha1.ValidateRequest, v1alpha1.ValidateResponse]
	deploy   *connect.Client[v1alpha1.DeployRequest, v1alpha1.DeployResponse]
	tearDown *connect.Client[v1alpha1.TearDownRequest, v1alpha1.TearDownResponse]
}

// Validate calls proto.provision_plugin.v1alpha1.ProvisionPluginService.Validate.
func (c *provisionPluginServiceClient) Validate(ctx context.Context, req *connect.Request[v1alpha1.ValidateRequest]) (*connect.Response[v1alpha1.ValidateResponse], error) {
	return c.validate.CallUnary(ctx, req)
}

// Deploy calls proto.provision_plugin.v1alpha1.ProvisionPluginService.Deploy.
func (c *provisionPluginServiceClient) Deploy(ctx context.Context, req *connect.Request[v1alpha1.DeployRequest]) (*connect.ServerStreamForClient[v1alpha1.DeployResponse], error) {
	return c.deploy.CallServerStream(ctx, req)
}

// TearDown calls proto.provision_plugin.v1alpha1.ProvisionPluginService.TearDown.
func (c *provisionPluginServiceClient) TearDown(ctx context.Context, req *connect.Request[v1alpha1.TearDownRequest]) (*connect.ServerStreamForClient[v1alpha1.TearDownResponse], error) {
	return c.tearDown.CallServerStream(ctx, req)
}

// ProvisionPluginServiceHandler is an implementation of the
// proto.provision_plugin.v1alpha1.ProvisionPluginService service.
type ProvisionPluginServiceHandler interface {
	Validate(context.Context, *connect.Request[v1alpha1.ValidateRequest]) (*connect.Response[v1alpha1.ValidateResponse], error)
	Deploy(context.Context, *connect.Request[v1alpha1.DeployRequest], *connect.ServerStream[v1alpha1.DeployResponse]) error
	TearDown(context.Context, *connect.Request[v1alpha1.TearDownRequest], *connect.ServerStream[v1alpha1.TearDownResponse]) error
}

// NewProvisionPluginServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewProvisionPluginServiceHandler(svc ProvisionPluginServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	provisionPluginServiceMethods := v1alpha1.File_proto_provision_plugin_v1alpha1_plugin_proto.Services().ByName("ProvisionPluginService").Methods()
	provisionPluginServiceValidateHandler := connect.NewUnaryHandler(
		ProvisionPluginServiceValidateProcedure,
		svc.Validate,
		connect.WithSchema(provisionPluginServiceMethods.ByName("Validate")),
		connect.WithHandlerOptions(opts...),
	)
	provisionPluginServiceDeployHandler := connect.NewServerStreamHandler(
		ProvisionPluginServiceDeployProcedure,
		svc.Deploy,
		connect.WithSchema(provisionPluginServiceMethods.ByName("Deploy")),
		connect.WithHandlerOptions(opts...),
	)
	provisionPluginServiceTearDownHandler := connect.NewServerStreamHandler(
		ProvisionPluginServiceTearDownProcedure,
		svc.TearDown,
		connect.WithSchema(provisionPluginServiceMethods.ByName("TearDown")),
		connect.WithHandlerOptions(opts...),
	)
	return "/proto.provision_plugin.v1alpha1.ProvisionPluginService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ProvisionPluginServiceValidateProcedure:
			provisionPluginServiceValidateHandler.ServeHTTP(w, r)
		case ProvisionPluginServiceDeployProcedure:
			provisionPluginServiceDeployHandler.ServeHTTP(w, r)
		case ProvisionPluginServiceTearDownProcedure:
			provisionPluginServiceTearDownHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedProvisionPluginServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedProvisionPluginServiceHandler struct{}

func (UnimplementedProvisionPluginServiceHandler) Validate(context.Context, *connect.Request[v1alpha1.ValidateRequest]) (*connect.Response[v1alpha1.ValidateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.provision_plugin.v1alpha1.ProvisionPluginService.Validate is not implemented"))
}

func (UnimplementedProvisionPluginServiceHandler) Deploy(context.Context, *connect.Request[v1alpha1.DeployRequest], *connect.ServerStream[v1alpha1.DeployResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("proto.provision_plugin.v1alpha1.ProvisionPluginService.Deploy is not implemented"))
}

func (UnimplementedProvisionPluginServiceHandler) TearDown(context.Context, *connect.Request[v1alpha1.TearDownRequest], *connect.ServerStream[v1alpha1.TearDownResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("proto.provision_plugin.v1alpha1.ProvisionPluginService.TearDown is not implemented"))
}
