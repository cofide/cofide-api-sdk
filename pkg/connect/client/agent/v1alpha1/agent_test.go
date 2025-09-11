// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"
	"testing"

	agentpb "github.com/cofide/cofide-api-sdk/gen/go/proto/agent/v1alpha1"
	agentsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/agent_service/v1alpha1"
	federatedservicepb "github.com/cofide/cofide-api-sdk/gen/go/proto/federated_service/v1alpha1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	fakeClusterID   = "fake-cluster-id"
	fakeTrustZoneID = "fake-tz-id"
	fakeAgentToken  = "fake-agent-token"
	fakeTrustDomain = "fake.trust.domain"
	fakeFSID        = "fake-fs-id"
	fakeFSName      = "fake-fs-name"
)

// TestAgentClient_Unimplemented tests AgentClient against an unimplemented server.
// This ensures that all errors returned are not wrapped can be converted to a gRPC Status using Status.Convert.
func TestAgentClient_Unimplemented(t *testing.T) {
	server := test.NewTestServer(t)
	agentsvcpb.RegisterAgentServiceServer(server.Server, &agentsvcpb.UnimplementedAgentServiceServer{})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	token, err := client.CreateAgentJoinToken(ctx, "", "")
	test.RequireUnimplemented(t, err)
	assert.Empty(t, token)

	err = client.UpdateAgentStatus(ctx, nil)
	test.RequireUnimplemented(t, err)

	fsID, err := client.RegisterFederatedService(ctx, nil)
	test.RequireUnimplemented(t, err)
	assert.Empty(t, fsID)

	err = client.DeregisterFederatedService(ctx, "")
	test.RequireUnimplemented(t, err)

	fs, err := client.GetFederatedService(ctx, "")
	test.RequireUnimplemented(t, err)
	assert.Nil(t, fs)

	err = client.UpdateFederatedService(ctx, nil)
	test.RequireUnimplemented(t, err)

	fss, err := client.ListFederatedServices(ctx)
	test.RequireUnimplemented(t, err)
	assert.Empty(t, fss)
}

func TestAgentClient(t *testing.T) {
	server := test.NewTestServer(t)
	agentsvcpb.RegisterAgentServiceServer(server.Server, &fakeAgentService{t: t})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	token, err := client.CreateAgentJoinToken(ctx, fakeClusterID, fakeTrustZoneID)
	require.NoError(t, err)
	assert.Equal(t, fakeAgentToken, token)

	status := &agentpb.AgentStatus{Status: agentpb.AgentStatusCode_AGENT_STATUS_CODE_RUNNING.Enum()}
	err = client.UpdateAgentStatus(ctx, status)
	require.NoError(t, err)

	fakeFS := fakeFederatedService()
	fsID, err := client.RegisterFederatedService(ctx, fakeFS)
	require.NoError(t, err)
	assert.EqualExportedValues(t, fakeFSID, fsID)

	err = client.DeregisterFederatedService(ctx, fakeFSID)
	require.NoError(t, err)

	err = client.UpdateFederatedService(ctx, fakeFS)
	require.NoError(t, err)

	gotFS, err := client.GetFederatedService(ctx, fakeFSID)
	require.NoError(t, err)
	assert.EqualExportedValues(t, fakeFS, gotFS)

	fss, err := client.ListFederatedServices(ctx)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*federatedservicepb.FederatedService{fakeFS}, fss)
}

type fakeAgentService struct {
	t *testing.T
}

func (f *fakeAgentService) CreateAgentJoinToken(ctx context.Context, req *agentsvcpb.CreateAgentJoinTokenRequest) (*agentsvcpb.CreateAgentJoinTokenResponse, error) {
	assert.Equal(f.t, fakeClusterID, req.GetClusterId())
	assert.Equal(f.t, fakeTrustZoneID, req.GetTrustZoneId())
	return &agentsvcpb.CreateAgentJoinTokenResponse{AgentToken: test.PtrOf(fakeAgentToken)}, nil
}

func (f *fakeAgentService) UpdateAgentStatus(ctx context.Context, req *agentsvcpb.UpdateAgentStatusRequest) (*agentsvcpb.UpdateAgentStatusResponse, error) {
	assert.Equal(f.t, agentpb.AgentStatusCode_AGENT_STATUS_CODE_RUNNING, req.Status.GetStatus())
	return &agentsvcpb.UpdateAgentStatusResponse{}, nil
}

func (f *fakeAgentService) RegisterFederatedService(ctx context.Context, req *agentsvcpb.RegisterFederatedServiceRequest) (*agentsvcpb.RegisterFederatedServiceResponse, error) {
	assert.EqualExportedValues(f.t, fakeFederatedService(), req.Service)
	return &agentsvcpb.RegisterFederatedServiceResponse{ServiceId: fakeFSID}, nil
}

func (f *fakeAgentService) DeregisterFederatedService(ctx context.Context, req *agentsvcpb.DeregisterFederatedServiceRequest) (*agentsvcpb.DeregisterFederatedServiceResponse, error) {
	assert.Equal(f.t, fakeFSID, req.ServiceId)
	return &agentsvcpb.DeregisterFederatedServiceResponse{}, nil
}

func (f *fakeAgentService) UpdateFederatedService(ctx context.Context, req *agentsvcpb.UpdateFederatedServiceRequest) (*agentsvcpb.UpdateFederatedServiceResponse, error) {
	assert.EqualExportedValues(f.t, fakeFederatedService(), req.Service)
	return &agentsvcpb.UpdateFederatedServiceResponse{}, nil
}

func (f *fakeAgentService) GetFederatedService(ctx context.Context, req *agentsvcpb.GetFederatedServiceRequest) (*agentsvcpb.GetFederatedServiceResponse, error) {
	assert.Equal(f.t, fakeFSID, req.ServiceId)
	fs := fakeFederatedService()
	return &agentsvcpb.GetFederatedServiceResponse{Service: fs}, nil
}

func (f *fakeAgentService) ListFederatedServices(context.Context, *agentsvcpb.ListFederatedServicesRequest) (*agentsvcpb.ListFederatedServicesResponse, error) {
	fss := []*federatedservicepb.FederatedService{fakeFederatedService()}
	return &agentsvcpb.ListFederatedServicesResponse{Services: fss}, nil
}

func fakeFederatedService() *federatedservicepb.FederatedService {
	return &federatedservicepb.FederatedService{
		Id:   fakeFSID,
		Name: fakeFSName,
	}
}
