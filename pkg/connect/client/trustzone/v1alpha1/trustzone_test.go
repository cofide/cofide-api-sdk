// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"
	"testing"

	clusterpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cluster/v1alpha1"
	trustzonesvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/trust_zone_service/v1alpha1"
	trustzonepb "github.com/cofide/cofide-api-sdk/gen/go/proto/trust_zone/v1alpha1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	fakeTrustZoneID   = "fake-tz-id"
	fakeTrustZoneName = "fake-tz-name"
	fakeClusterID     = "fake-cluster-id"
	fakeClusterName   = "fake-cluster-name"
	fakeAgentToken    = "fake-agent-token"
	fakeAgentID       = "fake-agent-id"
	fakeTrustDomain   = "fake.trust.domain"
)

// TestTrustZoneClient_Unimplemented tests TrustZoneClient against an unimplemented server.
// This ensures that all errors returned are not wrapped can be converted to a gRPC Status using Status.Convert.
func TestTrustZoneClient_Unimplemented(t *testing.T) {
	server := test.NewTestServer(t)
	trustzonesvcpb.RegisterTrustZoneServiceServer(server.Server, &trustzonesvcpb.UnimplementedTrustZoneServiceServer{})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	trustZone, err := client.CreateTrustZone(ctx, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, trustZone)

	trustZone, err = client.GetTrustZone(ctx, "")
	test.RequireUnimplemented(t, err)
	assert.Nil(t, trustZone)

	trustZones, err := client.ListTrustZones(ctx, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, trustZones)

	token, err := client.RegisterCluster(ctx, "", nil)
	test.RequireUnimplemented(t, err)
	assert.Empty(t, token)

	agentID, err := client.RegisterAgent(ctx, nil, "", nil)
	test.RequireUnimplemented(t, err)
	assert.Empty(t, agentID)
}

func TestTrustZoneClient(t *testing.T) {
	server := test.NewTestServer(t)
	trustzonesvcpb.RegisterTrustZoneServiceServer(server.Server, &fakeTrustZoneService{t: t})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	trustZone := fakeTrustZone()

	createdTrustZone, err := client.CreateTrustZone(ctx, trustZone)
	require.NoError(t, err)
	assert.EqualExportedValues(t, trustZone, createdTrustZone)

	gotTrustZone, err := client.GetTrustZone(ctx, fakeTrustZoneID)
	require.NoError(t, err)
	assert.EqualExportedValues(t, trustZone, gotTrustZone)

	filter := &trustzonesvcpb.ListTrustZonesRequest_Filter{Name: test.PtrOf(fakeTrustZoneName)}
	trustZones, err := client.ListTrustZones(ctx, filter)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*trustzonepb.TrustZone{trustZone}, trustZones)

	cluster := fakeCluster()
	token, err := client.RegisterCluster(ctx, fakeTrustZoneID, cluster)
	require.NoError(t, err)
	assert.Equal(t, fakeAgentToken, token)

	agent := fakeAgent()
	bundle := &types.Bundle{TrustDomain: fakeTrustDomain}
	agentID, err := client.RegisterAgent(ctx, agent, fakeAgentToken, bundle)
	require.NoError(t, err)
	assert.Equal(t, fakeAgentID, agentID)
}

type fakeTrustZoneService struct {
	t *testing.T
}

func (f *fakeTrustZoneService) CreateTrustZone(ctx context.Context, req *trustzonesvcpb.CreateTrustZoneRequest) (*trustzonesvcpb.CreateTrustZoneResponse, error) {
	assert.EqualExportedValues(f.t, fakeTrustZone(), req.TrustZone)
	return &trustzonesvcpb.CreateTrustZoneResponse{TrustZone: req.TrustZone}, nil
}

func (f *fakeTrustZoneService) ListTrustZones(ctx context.Context, req *trustzonesvcpb.ListTrustZonesRequest) (*trustzonesvcpb.ListTrustZonesResponse, error) {
	assert.Equal(f.t, fakeTrustZoneName, req.Filter.GetName())
	trustZones := []*trustzonepb.TrustZone{fakeTrustZone()}
	return &trustzonesvcpb.ListTrustZonesResponse{TrustZones: trustZones}, nil
}

func (f *fakeTrustZoneService) GetTrustZoneDetails(ctx context.Context, req *trustzonesvcpb.GetTrustZoneDetailsRequest) (*trustzonesvcpb.GetTrustZoneDetailsResponse, error) {
	assert.EqualExportedValues(f.t, fakeTrustZoneID, req.TrustZoneId)
	return &trustzonesvcpb.GetTrustZoneDetailsResponse{TrustZone: fakeTrustZone()}, nil
}

func (f *fakeTrustZoneService) RegisterCluster(ctx context.Context, req *trustzonesvcpb.RegisterClusterRequest) (*trustzonesvcpb.RegisterClusterResponse, error) {
	assert.EqualExportedValues(f.t, fakeCluster(), req.Cluster)
	assert.Equal(f.t, fakeTrustZoneID, req.TrustZoneId)
	return &trustzonesvcpb.RegisterClusterResponse{AgentToken: fakeAgentToken}, nil
}

func (f *fakeTrustZoneService) RegisterAgent(ctx context.Context, req *trustzonesvcpb.RegisterAgentRequest) (*trustzonesvcpb.RegisterAgentResponse, error) {
	assert.EqualExportedValues(f.t, fakeAgent(), req.Agent)
	assert.Equal(f.t, fakeAgentToken, req.AgentToken)
	assert.Equal(f.t, fakeTrustDomain, req.Bundle.TrustDomain)
	return &trustzonesvcpb.RegisterAgentResponse{AgentId: fakeAgentID}, nil
}

func fakeTrustZone() *trustzonepb.TrustZone {
	return &trustzonepb.TrustZone{
		Id:   test.PtrOf(fakeTrustZoneID),
		Name: fakeTrustZoneName,
	}
}

func fakeCluster() *clusterpb.Cluster {
	return &clusterpb.Cluster{
		Id:   test.PtrOf(fakeClusterID),
		Name: test.PtrOf(fakeClusterName),
	}
}

func fakeAgent() *trustzonesvcpb.Agent {
	return &trustzonesvcpb.Agent{
		AgentId:     fakeAgentID,
		ClusterId:   fakeClusterID,
		TrustZoneId: fakeTrustZoneID,
	}
}
