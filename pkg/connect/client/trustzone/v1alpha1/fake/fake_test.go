// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"testing"

	trustzonesvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/trust_zone_service/v1alpha1"
	trustzonepb "github.com/cofide/cofide-api-sdk/gen/go/proto/trust_zone/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_fakeTrustZoneClient_CreateTrustZone(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	trustZone := test.FakeTrustZone()

	createdTrustZone, err := client.CreateTrustZone(ctx, trustZone)
	require.NoError(t, err)
	trustZone.Id = createdTrustZone.Id
	assert.EqualExportedValues(t, trustZone, createdTrustZone)
	assert.EqualExportedValues(t, trustZone, fake.TrustZones[*createdTrustZone.Id])
}

func Test_fakeTrustZoneClient_DestroyTrustZone(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	err := client.DestroyTrustZone(ctx, test.FakeTrustZoneID)
	require.Error(t, err)

	fake.TrustZones[test.FakeTrustZoneID] = test.FakeTrustZone()

	err = client.DestroyTrustZone(ctx, test.FakeTrustZoneID)
	require.NoError(t, err)
	require.Nil(t, fake.TrustZones[test.FakeTrustZoneID])
}

func Test_fakeTrustZoneClient_GetTrustZone(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	trustZone := test.FakeTrustZone()

	_, err := client.GetTrustZone(ctx, test.FakeTrustZoneID)
	require.Error(t, err)

	fake.TrustZones[test.FakeTrustZoneID] = test.FakeTrustZone()

	gotTrustZone, err := client.GetTrustZone(ctx, test.FakeTrustZoneID)
	require.NoError(t, err)
	assert.EqualExportedValues(t, trustZone, gotTrustZone)
	assert.EqualExportedValues(t, trustZone, fake.TrustZones[test.FakeTrustZoneID])
}

func Test_fakeTrustZoneClient_ListTrustZones(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	trustZone := test.FakeTrustZone()

	trustZones, err := client.ListTrustZones(ctx, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*trustzonepb.TrustZone{}, trustZones)

	fake.TrustZones[test.FakeTrustZoneID] = test.FakeTrustZone()

	trustZones, err = client.ListTrustZones(ctx, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*trustzonepb.TrustZone{trustZone}, trustZones)
}

func Test_fakeTrustZoneClient_RegisterCluster(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	trustZone := test.FakeTrustZone()
	cluster := test.FakeCluster()

	_, err := client.RegisterCluster(ctx, test.FakeTrustZoneID, cluster)
	require.Error(t, err)

	fake.TrustZones[test.FakeTrustZoneID] = trustZone

	agentToken, err := client.RegisterCluster(ctx, test.FakeTrustZoneID, cluster)
	require.NoError(t, err)
	for clusterID := range fake.Clusters {
		cluster.Id = &clusterID
	}
	assert.EqualExportedValues(t, cluster, fake.Clusters[cluster.GetId()])
	assert.Equal(t, fake.AgentJoinTokens[test.FakeTrustZoneID][cluster.GetId()], agentToken)
}

func Test_fakeTrustZoneClient_RegisterAgent(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	agent := fakeAgent()

	_, err := client.RegisterAgent(ctx, agent, test.FakeAgentToken, test.FakeBundle())
	require.Error(t, err)

	fake.TrustZones[test.FakeTrustZoneID] = test.FakeTrustZone()
	fake.Clusters[test.FakeClusterID] = test.FakeCluster()

	agentID, err := client.RegisterAgent(ctx, agent, test.FakeAgentToken, test.FakeBundle())
	require.NoError(t, err)
	expectedAgent := test.FakeAgent()
	expectedAgent.Id = &agentID
	assert.EqualExportedValues(t, expectedAgent, fake.Agents[agentID])
}

func fakeAgent() *trustzonesvcpb.Agent {
	return &trustzonesvcpb.Agent{
		AgentId:     test.FakeAgentID,
		ClusterId:   test.FakeClusterID,
		TrustZoneId: test.FakeTrustZoneID,
	}
}
