// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"testing"

	agentpb "github.com/cofide/cofide-api-sdk/gen/go/proto/agent/v1alpha1"
	federatedservicepb "github.com/cofide/cofide-api-sdk/gen/go/proto/federated_service/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

func Test_fakeAgentClient_CreateAgentJoinToken(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	_, err := client.CreateAgentJoinToken(ctx, test.FakeClusterID, test.FakeTrustZoneID)
	require.Error(t, err)

	fake.TrustZones[test.FakeTrustZoneID] = test.FakeTrustZone()

	_, err = client.CreateAgentJoinToken(ctx, test.FakeClusterID, test.FakeTrustZoneID)
	require.Error(t, err)

	fake.Clusters[test.FakeClusterID] = test.FakeCluster()

	token, err := client.CreateAgentJoinToken(ctx, test.FakeClusterID, test.FakeTrustZoneID)
	require.NoError(t, err)
	assert.Equal(t, fake.AgentJoinTokens[test.FakeTrustZoneID][test.FakeClusterID], token)
}

func Test_fakeAgentClient_UpdateAgentStatus(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	md := metadata.MD{"agent-id": []string{test.FakeAgentID}}
	ctx = metadata.NewIncomingContext(ctx, md)

	agentStatus := &agentpb.AgentStatus{Status: agentpb.AgentStatusCode_AGENT_STATUS_CODE_RUNNING.Enum()}
	err := client.UpdateAgentStatus(ctx, agentStatus)
	require.Error(t, err)

	fake.Agents[test.FakeAgentID] = test.FakeAgent()

	err = client.UpdateAgentStatus(ctx, agentStatus)
	require.NoError(t, err)
	assert.EqualExportedValues(t, agentStatus, fake.AgentStatus[test.FakeAgentID])
}

func Test_fakeAgentClient_RegisterFederatedService(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	fakeFS := test.FakeFederatedService()
	fsID, err := client.RegisterFederatedService(ctx, fakeFS)
	require.NoError(t, err)
	assert.Equal(t, fsID, fakeFS.Id)
	assert.EqualExportedValues(t, fakeFS, fake.FederatedServices[fsID])

	_, err = client.RegisterFederatedService(ctx, fakeFS)
	require.Error(t, err)
}

func Test_fakeAgentClient_DeregisterFederatedService(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	err := client.DeregisterFederatedService(ctx, test.FakeFSID)
	require.Error(t, err)

	fake.FederatedServices[test.FakeFSID] = test.FakeFederatedService()

	err = client.DeregisterFederatedService(ctx, test.FakeFSID)
	require.NoError(t, err)
}

func Test_fakeAgentClient_UpdateFederatedService(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	fakeFS := test.FakeFederatedService()

	err := client.UpdateFederatedService(ctx, fakeFS)
	require.Error(t, err)

	fake.FederatedServices[test.FakeFSID] = fakeFS

	err = client.UpdateFederatedService(ctx, fakeFS)
	require.NoError(t, err)
}

func Test_fakeAgentClient_GetFederatedService(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	_, err := client.GetFederatedService(ctx, test.FakeFSID)
	require.Error(t, err)

	fakeFS := test.FakeFederatedService()
	fake.FederatedServices[test.FakeFSID] = fakeFS

	gotFS, err := client.GetFederatedService(ctx, test.FakeFSID)
	require.NoError(t, err)
	assert.EqualExportedValues(t, fakeFS, gotFS)
}

func Test_fakeAgentClient_ListFederatedServices(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	fss, err := client.ListFederatedServices(ctx)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*federatedservicepb.FederatedService{}, fss)

	fakeFS := test.FakeFederatedService()
	fake.FederatedServices[test.FakeFSID] = fakeFS

	fss, err = client.ListFederatedServices(ctx)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*federatedservicepb.FederatedService{fakeFS}, fss)
}
