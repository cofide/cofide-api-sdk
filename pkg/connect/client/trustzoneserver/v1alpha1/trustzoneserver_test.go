// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"
	"testing"

	trustzoneserversvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/trust_zone_server_service/v1alpha1"
	trustzoneserverpb "github.com/cofide/cofide-api-sdk/gen/go/proto/trust_zone_server/v1alpha1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	fakeTrustZoneServerID        = "fake-tzs-id"
	fakeTrustZoneID              = "fake-tz-id"
	fakeClusterID                = "fake-cluster-id"
	fakeKubernetesNamespace      = "fake-kubernetes-namespace"
	fakeKubernetesServiceAccount = "fake-kubernetes-service-account"
)

// TestTrustZoneServerClient_Unimplemented tests TrustZoneServerClient against an unimplemented server.
// This ensures that all errors returned are not wrapped can be converted to a gRPC Status using Status.Convert.
func TestTrustZoneServerClient_Unimplemented(t *testing.T) {
	server := test.NewTestServer(t)
	trustzoneserversvcpb.RegisterTrustZoneServerServiceServer(server.Server, &trustzoneserversvcpb.UnimplementedTrustZoneServerServiceServer{})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)

	trustZone, err := client.CreateTrustZoneServer(t.Context(), nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, trustZone)

	err = client.DestroyTrustZoneServer(t.Context(), "")
	test.RequireUnimplemented(t, err)

	trustZone, err = client.GetTrustZoneServer(t.Context(), "")
	test.RequireUnimplemented(t, err)
	assert.Nil(t, trustZone)

	trustZones, err := client.ListTrustZoneServers(t.Context(), nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, trustZones)

	trustZone, err = client.UpdateTrustZoneServer(t.Context(), nil, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, trustZone)
}

func TestTrustZoneServerClient(t *testing.T) {
	server := test.NewTestServer(t)
	trustzoneserversvcpb.RegisterTrustZoneServerServiceServer(server.Server, &fakeTrustZoneServerService{t: t})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)

	trustZoneServer := fakeTrustZoneServer()

	createdTrustZone, err := client.CreateTrustZoneServer(t.Context(), trustZoneServer)
	require.NoError(t, err)
	assert.EqualExportedValues(t, trustZoneServer, createdTrustZone)

	err = client.DestroyTrustZoneServer(t.Context(), fakeTrustZoneServerID)
	require.NoError(t, err)

	gotTrustZoneServer, err := client.GetTrustZoneServer(t.Context(), fakeTrustZoneServerID)
	require.NoError(t, err)
	assert.EqualExportedValues(t, trustZoneServer, gotTrustZoneServer)

	filter := &trustzoneserversvcpb.ListTrustZoneServersRequest_Filter{TrustZoneId: fakeTrustZoneID}
	trustZoneServers, err := client.ListTrustZoneServers(t.Context(), filter)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*trustzoneserverpb.TrustZoneServer{trustZoneServer}, trustZoneServers)

	updatedTrustZoneServer, err := client.UpdateTrustZoneServer(t.Context(), trustZoneServer, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, trustZoneServer, updatedTrustZoneServer)
}

type fakeTrustZoneServerService struct {
	t *testing.T
}

func (f *fakeTrustZoneServerService) CreateTrustZoneServer(ctx context.Context, req *trustzoneserversvcpb.CreateTrustZoneServerRequest) (*trustzoneserversvcpb.CreateTrustZoneServerResponse, error) {
	assert.EqualExportedValues(f.t, fakeTrustZoneServer(), req.TrustZoneServer)
	return &trustzoneserversvcpb.CreateTrustZoneServerResponse{TrustZoneServer: req.TrustZoneServer}, nil
}

func (f *fakeTrustZoneServerService) DestroyTrustZoneServer(ctx context.Context, req *trustzoneserversvcpb.DestroyTrustZoneServerRequest) (*trustzoneserversvcpb.DestroyTrustZoneServerResponse, error) {
	assert.Equal(f.t, fakeTrustZoneServerID, req.GetTrustZoneServerId())
	return &trustzoneserversvcpb.DestroyTrustZoneServerResponse{}, nil
}

func (f *fakeTrustZoneServerService) GetTrustZoneServer(ctx context.Context, req *trustzoneserversvcpb.GetTrustZoneServerRequest) (*trustzoneserversvcpb.GetTrustZoneServerResponse, error) {
	assert.Equal(f.t, fakeTrustZoneServerID, req.GetTrustZoneServerId())
	return &trustzoneserversvcpb.GetTrustZoneServerResponse{TrustZoneServer: fakeTrustZoneServer()}, nil
}

func (f *fakeTrustZoneServerService) ListTrustZoneServers(ctx context.Context, req *trustzoneserversvcpb.ListTrustZoneServersRequest) (*trustzoneserversvcpb.ListTrustZoneServersResponse, error) {
	assert.Equal(f.t, fakeTrustZoneID, req.Filter.GetTrustZoneId())
	trustZoneServers := []*trustzoneserverpb.TrustZoneServer{fakeTrustZoneServer()}
	return &trustzoneserversvcpb.ListTrustZoneServersResponse{TrustZoneServers: trustZoneServers}, nil
}

func (f *fakeTrustZoneServerService) UpdateTrustZoneServer(ctx context.Context, req *trustzoneserversvcpb.UpdateTrustZoneServerRequest) (*trustzoneserversvcpb.UpdateTrustZoneServerResponse, error) {
	assert.EqualExportedValues(f.t, fakeTrustZoneServer(), req.TrustZoneServer)
	return &trustzoneserversvcpb.UpdateTrustZoneServerResponse{TrustZoneServer: req.TrustZoneServer}, nil
}

func fakeTrustZoneServer() *trustzoneserverpb.TrustZoneServer {
	return &trustzoneserverpb.TrustZoneServer{
		Id:                       fakeTrustZoneServerID,
		TrustZoneId:              fakeTrustZoneID,
		ClusterId:                fakeClusterID,
		KubernetesNamespace:      fakeKubernetesNamespace,
		KubernetesServiceAccount: fakeKubernetesServiceAccount,
	}
}
