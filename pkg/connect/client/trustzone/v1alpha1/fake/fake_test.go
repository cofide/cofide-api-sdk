// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"testing"

	clusterpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cluster/v1alpha1"
	trustzonesvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/trust_zone_service/v1alpha1"
	trustzonepb "github.com/cofide/cofide-api-sdk/gen/go/proto/trust_zone/v1alpha1"
	trustzoneserverpb "github.com/cofide/cofide-api-sdk/gen/go/proto/trust_zone_server/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func Test_fakeTrustZoneClient_RegisterTrustZoneServer(t *testing.T) {
	t.Parallel()
	fake := fakeconnect.New()
	client := New(fake)

	trustZone1 := &trustzonepb.TrustZone{Id: test.PtrOf("trust-zone-1"), TrustDomain: "trust-domain-1"}
	fake.TrustZones[trustZone1.GetId()] = trustZone1
	trustZone2 := &trustzonepb.TrustZone{Id: test.PtrOf("trust-zone-2"), TrustDomain: "trust-domain-2"}
	fake.TrustZones[trustZone2.GetId()] = trustZone2
	trustZone3 := &trustzonepb.TrustZone{Id: test.PtrOf("trust-zone-3"), TrustDomain: "trust-domain-3"}
	fake.TrustZones[trustZone3.GetId()] = trustZone3
	trustZone4 := &trustzonepb.TrustZone{Id: test.PtrOf("trust-zone-4"), TrustDomain: "trust-domain-4"}
	fake.TrustZones[trustZone4.GetId()] = trustZone4
	cluster1 := &clusterpb.Cluster{Id: test.PtrOf("cluster-1"), TrustZoneId: test.PtrOf(trustZone1.GetId())}
	fake.Clusters[cluster1.GetId()] = cluster1
	trustZoneSever1 := &trustzoneserverpb.TrustZoneServer{Id: "trust-zone-server-1", TrustZoneId: trustZone2.GetId(), ClusterId: cluster1.GetId()}
	fake.TrustZoneServers[trustZoneSever1.GetId()] = trustZoneSever1
	trustZoneSever2 := &trustzoneserverpb.TrustZoneServer{Id: "trust-zone-server-2", TrustZoneId: trustZone3.GetId(), ClusterId: cluster1.GetId()}
	fake.TrustZoneServers[trustZoneSever2.GetId()] = trustZoneSever2
	trustZoneSever3 := &trustzoneserverpb.TrustZoneServer{Id: "trust-zone-server-3", TrustZoneId: trustZone3.GetId(), ClusterId: cluster1.GetId()}
	fake.TrustZoneServers[trustZoneSever3.GetId()] = trustZoneSever3
	token := "join-token"
	fake.TrustZoneServerJoinTokens[trustZoneSever2.GetId()] = map[string]struct{}{token: struct{}{}}

	type registrationInput struct {
		trustZoneServer   *trustzonesvcpb.TrustZoneServer
		bundle            *types.Bundle
		trustZoneServerID string
		joinToken         string
	}
	tests := []struct {
		name                      string
		input                     registrationInput
		expectedBundleTrustZoneID string
		wantErrMsg                string
		wantErrCode               codes.Code
	}{
		{
			name: "Registration using cluster ID",
			input: registrationInput{
				trustZoneServer: &trustzonesvcpb.TrustZoneServer{ClusterId: cluster1.GetId()},
				bundle:          &types.Bundle{TrustDomain: trustZone1.GetTrustDomain()},
			},
			expectedBundleTrustZoneID: trustZone1.GetId(),
		},
		{
			name: "Registration using trust zone server ID",
			input: registrationInput{
				bundle:            &types.Bundle{TrustDomain: trustZone2.GetTrustDomain()},
				trustZoneServerID: trustZoneSever1.GetId(),
			},
			expectedBundleTrustZoneID: trustZone2.GetId(),
		},
		{
			name: "Registration using trust zone server ID and valid join token",
			input: registrationInput{
				bundle:            &types.Bundle{TrustDomain: trustZone3.GetTrustDomain()},
				trustZoneServerID: trustZoneSever2.GetId(),
				joinToken:         token,
			},
			expectedBundleTrustZoneID: trustZone3.GetId(),
		},
		{
			name: "Registration using trust zone server ID and invalid join token",
			input: registrationInput{
				bundle:            &types.Bundle{TrustDomain: trustZone3.GetTrustDomain()},
				trustZoneServerID: trustZoneSever3.GetId(),
				joinToken:         "invalid",
			},
			wantErrMsg:  "rpc error: code = Unauthenticated desc = invalid token",
			wantErrCode: codes.Unauthenticated,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := client.RegisterTrustZoneServer(t.Context(), tt.input.trustZoneServer, tt.input.bundle, tt.input.trustZoneServerID, tt.input.joinToken)
			if tt.wantErrMsg != "" {
				require.EqualError(t, err, tt.wantErrMsg)
				statusFromErr, ok := status.FromError(err)
				require.True(t, ok)
				assert.Equal(t, tt.wantErrCode, statusFromErr.Code())
				return
			}
			require.NoError(t, err)
			assert.EqualExportedValues(t, tt.input.bundle, fake.TrustZoneBundles[tt.expectedBundleTrustZoneID])
		})
	}
}

func Test_fakeTrustZoneClient_UpdateTrustZoneBundle(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	fakeBundle := test.FakeBundle()
	err := client.UpdateTrustZoneBundle(ctx, test.FakeTrustZoneID, fakeBundle)
	require.Error(t, err)

	fake.TrustZones[test.FakeTrustZoneID] = test.FakeTrustZone()

	err = client.UpdateTrustZoneBundle(ctx, test.FakeTrustZoneID, fakeBundle)
	require.NoError(t, err)
	assert.Equal(t, fake.TrustZoneBundles[test.FakeTrustZoneID], fakeBundle)
	assert.NotSame(t, fake.TrustZoneBundles[test.FakeTrustZoneID], fakeBundle)

	fakeBundle.SequenceNumber = 42
	err = client.UpdateTrustZoneBundle(ctx, test.FakeTrustZoneID, fakeBundle)
	require.NoError(t, err)
	assert.Equal(t, fake.TrustZoneBundles[test.FakeTrustZoneID], fakeBundle)
}

func fakeAgent() *trustzonesvcpb.Agent {
	return &trustzonesvcpb.Agent{
		AgentId:     test.FakeAgentID,
		ClusterId:   test.FakeClusterID,
		TrustZoneId: test.FakeTrustZoneID,
	}
}
