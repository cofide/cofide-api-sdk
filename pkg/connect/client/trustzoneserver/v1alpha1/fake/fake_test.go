// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"testing"

	trustzoneserverpb "github.com/cofide/cofide-api-sdk/gen/go/proto/trust_zone_server/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_fakeTrustZoneServerClient_CreateTrustZoneServer(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	trustZoneServer := test.FakeTrustZoneServer()

	createdTrustZoneServer, err := client.CreateTrustZoneServer(ctx, trustZoneServer)
	require.NoError(t, err)
	trustZoneServer.Id = createdTrustZoneServer.Id
	assert.EqualExportedValues(t, trustZoneServer, createdTrustZoneServer)
	assert.EqualExportedValues(t, trustZoneServer, fake.TrustZoneServers[createdTrustZoneServer.Id])
}

func Test_fakeTrustZoneServerClient_DestroyTrustZoneServer(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	err := client.DestroyTrustZoneServer(ctx, test.FakeTrustZoneServerID)
	require.Error(t, err)

	fake.TrustZoneServers[test.FakeTrustZoneServerID] = test.FakeTrustZoneServer()

	err = client.DestroyTrustZoneServer(ctx, test.FakeTrustZoneServerID)
	require.NoError(t, err)
	require.Nil(t, fake.TrustZoneServers[test.FakeTrustZoneServerID])
}

func Test_fakeTrustZoneServerClient_GetTrustZoneServer(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	trustZoneServer := test.FakeTrustZoneServer()

	_, err := client.GetTrustZoneServer(ctx, test.FakeTrustZoneServerID)
	require.Error(t, err)

	fake.TrustZoneServers[test.FakeTrustZoneServerID] = test.FakeTrustZoneServer()

	gotTrustZoneServer, err := client.GetTrustZoneServer(ctx, test.FakeTrustZoneServerID)
	require.NoError(t, err)
	assert.EqualExportedValues(t, trustZoneServer, gotTrustZoneServer)
	assert.EqualExportedValues(t, trustZoneServer, fake.TrustZoneServers[test.FakeTrustZoneServerID])
}

func Test_fakeTrustZoneServerClient_ListTrustZoneServers(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	trustZoneServer := test.FakeTrustZoneServer()

	trustZoneServers, err := client.ListTrustZoneServers(ctx, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*trustzoneserverpb.TrustZoneServer{}, trustZoneServers)

	fake.TrustZoneServers[test.FakeTrustZoneServerID] = test.FakeTrustZoneServer()

	trustZoneServers, err = client.ListTrustZoneServers(ctx, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*trustzoneserverpb.TrustZoneServer{trustZoneServer}, trustZoneServers)
}
