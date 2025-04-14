// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/connect/trust_zone_service/v1alpha1/trust_zone_service.proto

package v1alpha1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha1 "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/trust_zone_service/v1alpha1"
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
	// TrustZoneServiceName is the fully-qualified name of the TrustZoneService service.
	TrustZoneServiceName = "proto.connect.trust_zone_service.v1alpha1.TrustZoneService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TrustZoneServiceCreateTrustZoneProcedure is the fully-qualified name of the TrustZoneService's
	// CreateTrustZone RPC.
	TrustZoneServiceCreateTrustZoneProcedure = "/proto.connect.trust_zone_service.v1alpha1.TrustZoneService/CreateTrustZone"
	// TrustZoneServiceDestroyTrustZoneProcedure is the fully-qualified name of the TrustZoneService's
	// DestroyTrustZone RPC.
	TrustZoneServiceDestroyTrustZoneProcedure = "/proto.connect.trust_zone_service.v1alpha1.TrustZoneService/DestroyTrustZone"
	// TrustZoneServiceGetTrustZoneProcedure is the fully-qualified name of the TrustZoneService's
	// GetTrustZone RPC.
	TrustZoneServiceGetTrustZoneProcedure = "/proto.connect.trust_zone_service.v1alpha1.TrustZoneService/GetTrustZone"
	// TrustZoneServiceListTrustZonesProcedure is the fully-qualified name of the TrustZoneService's
	// ListTrustZones RPC.
	TrustZoneServiceListTrustZonesProcedure = "/proto.connect.trust_zone_service.v1alpha1.TrustZoneService/ListTrustZones"
	// TrustZoneServiceUpdateTrustZoneProcedure is the fully-qualified name of the TrustZoneService's
	// UpdateTrustZone RPC.
	TrustZoneServiceUpdateTrustZoneProcedure = "/proto.connect.trust_zone_service.v1alpha1.TrustZoneService/UpdateTrustZone"
	// TrustZoneServiceGetTrustZoneDetailsProcedure is the fully-qualified name of the
	// TrustZoneService's GetTrustZoneDetails RPC.
	TrustZoneServiceGetTrustZoneDetailsProcedure = "/proto.connect.trust_zone_service.v1alpha1.TrustZoneService/GetTrustZoneDetails"
	// TrustZoneServiceRegisterClusterProcedure is the fully-qualified name of the TrustZoneService's
	// RegisterCluster RPC.
	TrustZoneServiceRegisterClusterProcedure = "/proto.connect.trust_zone_service.v1alpha1.TrustZoneService/RegisterCluster"
	// TrustZoneServiceRegisterAgentProcedure is the fully-qualified name of the TrustZoneService's
	// RegisterAgent RPC.
	TrustZoneServiceRegisterAgentProcedure = "/proto.connect.trust_zone_service.v1alpha1.TrustZoneService/RegisterAgent"
)

// TrustZoneServiceClient is a client for the
// proto.connect.trust_zone_service.v1alpha1.TrustZoneService service.
type TrustZoneServiceClient interface {
	CreateTrustZone(context.Context, *connect.Request[v1alpha1.CreateTrustZoneRequest]) (*connect.Response[v1alpha1.CreateTrustZoneResponse], error)
	DestroyTrustZone(context.Context, *connect.Request[v1alpha1.DestroyTrustZoneRequest]) (*connect.Response[v1alpha1.DestroyTrustZoneResponse], error)
	GetTrustZone(context.Context, *connect.Request[v1alpha1.GetTrustZoneRequest]) (*connect.Response[v1alpha1.GetTrustZoneResponse], error)
	ListTrustZones(context.Context, *connect.Request[v1alpha1.ListTrustZonesRequest]) (*connect.Response[v1alpha1.ListTrustZonesResponse], error)
	UpdateTrustZone(context.Context, *connect.Request[v1alpha1.UpdateTrustZoneRequest]) (*connect.Response[v1alpha1.UpdateTrustZoneResponse], error)
	// DEPRECATED: GetTrustZoneDetails to be replaced with GetTrustZone.
	GetTrustZoneDetails(context.Context, *connect.Request[v1alpha1.GetTrustZoneDetailsRequest]) (*connect.Response[v1alpha1.GetTrustZoneDetailsResponse], error)
	// DEPRECATED: Agent join token creation moved to AgentService.CreateAgentJoinToken.
	// Cluster creation to be moved to ClusterService.CreateCluster.
	RegisterCluster(context.Context, *connect.Request[v1alpha1.RegisterClusterRequest]) (*connect.Response[v1alpha1.RegisterClusterResponse], error)
	RegisterAgent(context.Context, *connect.Request[v1alpha1.RegisterAgentRequest]) (*connect.Response[v1alpha1.RegisterAgentResponse], error)
}

// NewTrustZoneServiceClient constructs a client for the
// proto.connect.trust_zone_service.v1alpha1.TrustZoneService service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTrustZoneServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TrustZoneServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	trustZoneServiceMethods := v1alpha1.File_proto_connect_trust_zone_service_v1alpha1_trust_zone_service_proto.Services().ByName("TrustZoneService").Methods()
	return &trustZoneServiceClient{
		createTrustZone: connect.NewClient[v1alpha1.CreateTrustZoneRequest, v1alpha1.CreateTrustZoneResponse](
			httpClient,
			baseURL+TrustZoneServiceCreateTrustZoneProcedure,
			connect.WithSchema(trustZoneServiceMethods.ByName("CreateTrustZone")),
			connect.WithClientOptions(opts...),
		),
		destroyTrustZone: connect.NewClient[v1alpha1.DestroyTrustZoneRequest, v1alpha1.DestroyTrustZoneResponse](
			httpClient,
			baseURL+TrustZoneServiceDestroyTrustZoneProcedure,
			connect.WithSchema(trustZoneServiceMethods.ByName("DestroyTrustZone")),
			connect.WithClientOptions(opts...),
		),
		getTrustZone: connect.NewClient[v1alpha1.GetTrustZoneRequest, v1alpha1.GetTrustZoneResponse](
			httpClient,
			baseURL+TrustZoneServiceGetTrustZoneProcedure,
			connect.WithSchema(trustZoneServiceMethods.ByName("GetTrustZone")),
			connect.WithClientOptions(opts...),
		),
		listTrustZones: connect.NewClient[v1alpha1.ListTrustZonesRequest, v1alpha1.ListTrustZonesResponse](
			httpClient,
			baseURL+TrustZoneServiceListTrustZonesProcedure,
			connect.WithSchema(trustZoneServiceMethods.ByName("ListTrustZones")),
			connect.WithClientOptions(opts...),
		),
		updateTrustZone: connect.NewClient[v1alpha1.UpdateTrustZoneRequest, v1alpha1.UpdateTrustZoneResponse](
			httpClient,
			baseURL+TrustZoneServiceUpdateTrustZoneProcedure,
			connect.WithSchema(trustZoneServiceMethods.ByName("UpdateTrustZone")),
			connect.WithClientOptions(opts...),
		),
		getTrustZoneDetails: connect.NewClient[v1alpha1.GetTrustZoneDetailsRequest, v1alpha1.GetTrustZoneDetailsResponse](
			httpClient,
			baseURL+TrustZoneServiceGetTrustZoneDetailsProcedure,
			connect.WithSchema(trustZoneServiceMethods.ByName("GetTrustZoneDetails")),
			connect.WithClientOptions(opts...),
		),
		registerCluster: connect.NewClient[v1alpha1.RegisterClusterRequest, v1alpha1.RegisterClusterResponse](
			httpClient,
			baseURL+TrustZoneServiceRegisterClusterProcedure,
			connect.WithSchema(trustZoneServiceMethods.ByName("RegisterCluster")),
			connect.WithClientOptions(opts...),
		),
		registerAgent: connect.NewClient[v1alpha1.RegisterAgentRequest, v1alpha1.RegisterAgentResponse](
			httpClient,
			baseURL+TrustZoneServiceRegisterAgentProcedure,
			connect.WithSchema(trustZoneServiceMethods.ByName("RegisterAgent")),
			connect.WithClientOptions(opts...),
		),
	}
}

// trustZoneServiceClient implements TrustZoneServiceClient.
type trustZoneServiceClient struct {
	createTrustZone     *connect.Client[v1alpha1.CreateTrustZoneRequest, v1alpha1.CreateTrustZoneResponse]
	destroyTrustZone    *connect.Client[v1alpha1.DestroyTrustZoneRequest, v1alpha1.DestroyTrustZoneResponse]
	getTrustZone        *connect.Client[v1alpha1.GetTrustZoneRequest, v1alpha1.GetTrustZoneResponse]
	listTrustZones      *connect.Client[v1alpha1.ListTrustZonesRequest, v1alpha1.ListTrustZonesResponse]
	updateTrustZone     *connect.Client[v1alpha1.UpdateTrustZoneRequest, v1alpha1.UpdateTrustZoneResponse]
	getTrustZoneDetails *connect.Client[v1alpha1.GetTrustZoneDetailsRequest, v1alpha1.GetTrustZoneDetailsResponse]
	registerCluster     *connect.Client[v1alpha1.RegisterClusterRequest, v1alpha1.RegisterClusterResponse]
	registerAgent       *connect.Client[v1alpha1.RegisterAgentRequest, v1alpha1.RegisterAgentResponse]
}

// CreateTrustZone calls proto.connect.trust_zone_service.v1alpha1.TrustZoneService.CreateTrustZone.
func (c *trustZoneServiceClient) CreateTrustZone(ctx context.Context, req *connect.Request[v1alpha1.CreateTrustZoneRequest]) (*connect.Response[v1alpha1.CreateTrustZoneResponse], error) {
	return c.createTrustZone.CallUnary(ctx, req)
}

// DestroyTrustZone calls
// proto.connect.trust_zone_service.v1alpha1.TrustZoneService.DestroyTrustZone.
func (c *trustZoneServiceClient) DestroyTrustZone(ctx context.Context, req *connect.Request[v1alpha1.DestroyTrustZoneRequest]) (*connect.Response[v1alpha1.DestroyTrustZoneResponse], error) {
	return c.destroyTrustZone.CallUnary(ctx, req)
}

// GetTrustZone calls proto.connect.trust_zone_service.v1alpha1.TrustZoneService.GetTrustZone.
func (c *trustZoneServiceClient) GetTrustZone(ctx context.Context, req *connect.Request[v1alpha1.GetTrustZoneRequest]) (*connect.Response[v1alpha1.GetTrustZoneResponse], error) {
	return c.getTrustZone.CallUnary(ctx, req)
}

// ListTrustZones calls proto.connect.trust_zone_service.v1alpha1.TrustZoneService.ListTrustZones.
func (c *trustZoneServiceClient) ListTrustZones(ctx context.Context, req *connect.Request[v1alpha1.ListTrustZonesRequest]) (*connect.Response[v1alpha1.ListTrustZonesResponse], error) {
	return c.listTrustZones.CallUnary(ctx, req)
}

// UpdateTrustZone calls proto.connect.trust_zone_service.v1alpha1.TrustZoneService.UpdateTrustZone.
func (c *trustZoneServiceClient) UpdateTrustZone(ctx context.Context, req *connect.Request[v1alpha1.UpdateTrustZoneRequest]) (*connect.Response[v1alpha1.UpdateTrustZoneResponse], error) {
	return c.updateTrustZone.CallUnary(ctx, req)
}

// GetTrustZoneDetails calls
// proto.connect.trust_zone_service.v1alpha1.TrustZoneService.GetTrustZoneDetails.
func (c *trustZoneServiceClient) GetTrustZoneDetails(ctx context.Context, req *connect.Request[v1alpha1.GetTrustZoneDetailsRequest]) (*connect.Response[v1alpha1.GetTrustZoneDetailsResponse], error) {
	return c.getTrustZoneDetails.CallUnary(ctx, req)
}

// RegisterCluster calls proto.connect.trust_zone_service.v1alpha1.TrustZoneService.RegisterCluster.
func (c *trustZoneServiceClient) RegisterCluster(ctx context.Context, req *connect.Request[v1alpha1.RegisterClusterRequest]) (*connect.Response[v1alpha1.RegisterClusterResponse], error) {
	return c.registerCluster.CallUnary(ctx, req)
}

// RegisterAgent calls proto.connect.trust_zone_service.v1alpha1.TrustZoneService.RegisterAgent.
func (c *trustZoneServiceClient) RegisterAgent(ctx context.Context, req *connect.Request[v1alpha1.RegisterAgentRequest]) (*connect.Response[v1alpha1.RegisterAgentResponse], error) {
	return c.registerAgent.CallUnary(ctx, req)
}

// TrustZoneServiceHandler is an implementation of the
// proto.connect.trust_zone_service.v1alpha1.TrustZoneService service.
type TrustZoneServiceHandler interface {
	CreateTrustZone(context.Context, *connect.Request[v1alpha1.CreateTrustZoneRequest]) (*connect.Response[v1alpha1.CreateTrustZoneResponse], error)
	DestroyTrustZone(context.Context, *connect.Request[v1alpha1.DestroyTrustZoneRequest]) (*connect.Response[v1alpha1.DestroyTrustZoneResponse], error)
	GetTrustZone(context.Context, *connect.Request[v1alpha1.GetTrustZoneRequest]) (*connect.Response[v1alpha1.GetTrustZoneResponse], error)
	ListTrustZones(context.Context, *connect.Request[v1alpha1.ListTrustZonesRequest]) (*connect.Response[v1alpha1.ListTrustZonesResponse], error)
	UpdateTrustZone(context.Context, *connect.Request[v1alpha1.UpdateTrustZoneRequest]) (*connect.Response[v1alpha1.UpdateTrustZoneResponse], error)
	// DEPRECATED: GetTrustZoneDetails to be replaced with GetTrustZone.
	GetTrustZoneDetails(context.Context, *connect.Request[v1alpha1.GetTrustZoneDetailsRequest]) (*connect.Response[v1alpha1.GetTrustZoneDetailsResponse], error)
	// DEPRECATED: Agent join token creation moved to AgentService.CreateAgentJoinToken.
	// Cluster creation to be moved to ClusterService.CreateCluster.
	RegisterCluster(context.Context, *connect.Request[v1alpha1.RegisterClusterRequest]) (*connect.Response[v1alpha1.RegisterClusterResponse], error)
	RegisterAgent(context.Context, *connect.Request[v1alpha1.RegisterAgentRequest]) (*connect.Response[v1alpha1.RegisterAgentResponse], error)
}

// NewTrustZoneServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTrustZoneServiceHandler(svc TrustZoneServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	trustZoneServiceMethods := v1alpha1.File_proto_connect_trust_zone_service_v1alpha1_trust_zone_service_proto.Services().ByName("TrustZoneService").Methods()
	trustZoneServiceCreateTrustZoneHandler := connect.NewUnaryHandler(
		TrustZoneServiceCreateTrustZoneProcedure,
		svc.CreateTrustZone,
		connect.WithSchema(trustZoneServiceMethods.ByName("CreateTrustZone")),
		connect.WithHandlerOptions(opts...),
	)
	trustZoneServiceDestroyTrustZoneHandler := connect.NewUnaryHandler(
		TrustZoneServiceDestroyTrustZoneProcedure,
		svc.DestroyTrustZone,
		connect.WithSchema(trustZoneServiceMethods.ByName("DestroyTrustZone")),
		connect.WithHandlerOptions(opts...),
	)
	trustZoneServiceGetTrustZoneHandler := connect.NewUnaryHandler(
		TrustZoneServiceGetTrustZoneProcedure,
		svc.GetTrustZone,
		connect.WithSchema(trustZoneServiceMethods.ByName("GetTrustZone")),
		connect.WithHandlerOptions(opts...),
	)
	trustZoneServiceListTrustZonesHandler := connect.NewUnaryHandler(
		TrustZoneServiceListTrustZonesProcedure,
		svc.ListTrustZones,
		connect.WithSchema(trustZoneServiceMethods.ByName("ListTrustZones")),
		connect.WithHandlerOptions(opts...),
	)
	trustZoneServiceUpdateTrustZoneHandler := connect.NewUnaryHandler(
		TrustZoneServiceUpdateTrustZoneProcedure,
		svc.UpdateTrustZone,
		connect.WithSchema(trustZoneServiceMethods.ByName("UpdateTrustZone")),
		connect.WithHandlerOptions(opts...),
	)
	trustZoneServiceGetTrustZoneDetailsHandler := connect.NewUnaryHandler(
		TrustZoneServiceGetTrustZoneDetailsProcedure,
		svc.GetTrustZoneDetails,
		connect.WithSchema(trustZoneServiceMethods.ByName("GetTrustZoneDetails")),
		connect.WithHandlerOptions(opts...),
	)
	trustZoneServiceRegisterClusterHandler := connect.NewUnaryHandler(
		TrustZoneServiceRegisterClusterProcedure,
		svc.RegisterCluster,
		connect.WithSchema(trustZoneServiceMethods.ByName("RegisterCluster")),
		connect.WithHandlerOptions(opts...),
	)
	trustZoneServiceRegisterAgentHandler := connect.NewUnaryHandler(
		TrustZoneServiceRegisterAgentProcedure,
		svc.RegisterAgent,
		connect.WithSchema(trustZoneServiceMethods.ByName("RegisterAgent")),
		connect.WithHandlerOptions(opts...),
	)
	return "/proto.connect.trust_zone_service.v1alpha1.TrustZoneService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TrustZoneServiceCreateTrustZoneProcedure:
			trustZoneServiceCreateTrustZoneHandler.ServeHTTP(w, r)
		case TrustZoneServiceDestroyTrustZoneProcedure:
			trustZoneServiceDestroyTrustZoneHandler.ServeHTTP(w, r)
		case TrustZoneServiceGetTrustZoneProcedure:
			trustZoneServiceGetTrustZoneHandler.ServeHTTP(w, r)
		case TrustZoneServiceListTrustZonesProcedure:
			trustZoneServiceListTrustZonesHandler.ServeHTTP(w, r)
		case TrustZoneServiceUpdateTrustZoneProcedure:
			trustZoneServiceUpdateTrustZoneHandler.ServeHTTP(w, r)
		case TrustZoneServiceGetTrustZoneDetailsProcedure:
			trustZoneServiceGetTrustZoneDetailsHandler.ServeHTTP(w, r)
		case TrustZoneServiceRegisterClusterProcedure:
			trustZoneServiceRegisterClusterHandler.ServeHTTP(w, r)
		case TrustZoneServiceRegisterAgentProcedure:
			trustZoneServiceRegisterAgentHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTrustZoneServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTrustZoneServiceHandler struct{}

func (UnimplementedTrustZoneServiceHandler) CreateTrustZone(context.Context, *connect.Request[v1alpha1.CreateTrustZoneRequest]) (*connect.Response[v1alpha1.CreateTrustZoneResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.connect.trust_zone_service.v1alpha1.TrustZoneService.CreateTrustZone is not implemented"))
}

func (UnimplementedTrustZoneServiceHandler) DestroyTrustZone(context.Context, *connect.Request[v1alpha1.DestroyTrustZoneRequest]) (*connect.Response[v1alpha1.DestroyTrustZoneResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.connect.trust_zone_service.v1alpha1.TrustZoneService.DestroyTrustZone is not implemented"))
}

func (UnimplementedTrustZoneServiceHandler) GetTrustZone(context.Context, *connect.Request[v1alpha1.GetTrustZoneRequest]) (*connect.Response[v1alpha1.GetTrustZoneResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.connect.trust_zone_service.v1alpha1.TrustZoneService.GetTrustZone is not implemented"))
}

func (UnimplementedTrustZoneServiceHandler) ListTrustZones(context.Context, *connect.Request[v1alpha1.ListTrustZonesRequest]) (*connect.Response[v1alpha1.ListTrustZonesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.connect.trust_zone_service.v1alpha1.TrustZoneService.ListTrustZones is not implemented"))
}

func (UnimplementedTrustZoneServiceHandler) UpdateTrustZone(context.Context, *connect.Request[v1alpha1.UpdateTrustZoneRequest]) (*connect.Response[v1alpha1.UpdateTrustZoneResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.connect.trust_zone_service.v1alpha1.TrustZoneService.UpdateTrustZone is not implemented"))
}

func (UnimplementedTrustZoneServiceHandler) GetTrustZoneDetails(context.Context, *connect.Request[v1alpha1.GetTrustZoneDetailsRequest]) (*connect.Response[v1alpha1.GetTrustZoneDetailsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.connect.trust_zone_service.v1alpha1.TrustZoneService.GetTrustZoneDetails is not implemented"))
}

func (UnimplementedTrustZoneServiceHandler) RegisterCluster(context.Context, *connect.Request[v1alpha1.RegisterClusterRequest]) (*connect.Response[v1alpha1.RegisterClusterResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.connect.trust_zone_service.v1alpha1.TrustZoneService.RegisterCluster is not implemented"))
}

func (UnimplementedTrustZoneServiceHandler) RegisterAgent(context.Context, *connect.Request[v1alpha1.RegisterAgentRequest]) (*connect.Response[v1alpha1.RegisterAgentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.connect.trust_zone_service.v1alpha1.TrustZoneService.RegisterAgent is not implemented"))
}
