// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"testing"

	"github.com/cofide/cofide-api-sdk/gen/go/proto/connect/trust_zone_server_service/v1alpha1"
	trustzoneserverpb "github.com/cofide/cofide-api-sdk/gen/go/proto/trust_zone_server/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

func Test_fakeTrustZoneServerClient_CreateTrustZoneServer(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)

	trustZoneServer := test.FakeTrustZoneServer()

	createdTrustZoneServer, err := client.CreateTrustZoneServer(t.Context(), trustZoneServer)
	require.NoError(t, err)
	trustZoneServer.Id = createdTrustZoneServer.Id
	assert.EqualExportedValues(t, trustZoneServer, createdTrustZoneServer)
	assert.EqualExportedValues(t, trustZoneServer, fake.TrustZoneServers[createdTrustZoneServer.Id])
}

func Test_fakeTrustZoneServerClient_UpdateTrustZoneServer(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)

	existingTrustZoneServer, err := client.CreateTrustZoneServer(t.Context(), test.FakeTrustZoneServer())
	require.NoError(t, err)

	mustNewStructPbValue := func(value *structpb.Value, err error) *structpb.Value {
		require.NoError(t, err)
		return value
	}
	trustZoneServer := proto.Clone(existingTrustZoneServer).(*trustzoneserverpb.TrustZoneServer)
	trustZoneServer.HelmValues = &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"field1": mustNewStructPbValue(structpb.NewValue("value1")),
			"field2": mustNewStructPbValue(structpb.NewValue("value2")),
		},
	}

	// No change with empty update mask
	updatedTrustZoneServer, err := client.UpdateTrustZoneServer(t.Context(), trustZoneServer, &v1alpha1.UpdateTrustZoneServerRequest_UpdateMask{})
	require.NoError(t, err)
	assert.EqualExportedValues(t, existingTrustZoneServer, updatedTrustZoneServer)
	assert.EqualExportedValues(t, updatedTrustZoneServer, fake.TrustZoneServers[updatedTrustZoneServer.GetId()])

	// Changes with non-empty update mask
	updatedTrustZoneServer, err = client.UpdateTrustZoneServer(t.Context(), trustZoneServer, &v1alpha1.UpdateTrustZoneServerRequest_UpdateMask{HelmValues: true})
	require.NoError(t, err)
	assert.EqualExportedValues(t, trustZoneServer, updatedTrustZoneServer)
	assert.EqualExportedValues(t, updatedTrustZoneServer, fake.TrustZoneServers[updatedTrustZoneServer.GetId()])
}

func Test_fakeTrustZoneServerClient_DestroyTrustZoneServer(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)

	err := client.DestroyTrustZoneServer(t.Context(), test.FakeTrustZoneServerID)
	require.Error(t, err)

	fake.TrustZoneServers[test.FakeTrustZoneServerID] = test.FakeTrustZoneServer()

	err = client.DestroyTrustZoneServer(t.Context(), test.FakeTrustZoneServerID)
	require.NoError(t, err)
	require.Nil(t, fake.TrustZoneServers[test.FakeTrustZoneServerID])
}

func Test_fakeTrustZoneServerClient_GetTrustZoneServer(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)

	trustZoneServer := test.FakeTrustZoneServer()

	_, err := client.GetTrustZoneServer(t.Context(), test.FakeTrustZoneServerID)
	require.Error(t, err)

	fake.TrustZoneServers[test.FakeTrustZoneServerID] = test.FakeTrustZoneServer()

	gotTrustZoneServer, err := client.GetTrustZoneServer(t.Context(), test.FakeTrustZoneServerID)
	require.NoError(t, err)
	assert.EqualExportedValues(t, trustZoneServer, gotTrustZoneServer)
	assert.EqualExportedValues(t, trustZoneServer, fake.TrustZoneServers[test.FakeTrustZoneServerID])
}

func Test_fakeTrustZoneServerClient_ListTrustZoneServers(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)

	trustZoneServer := test.FakeTrustZoneServer()

	trustZoneServers, err := client.ListTrustZoneServers(t.Context(), nil)
	require.NoError(t, err)
	assert.Empty(t, trustZoneServers)

	fake.TrustZoneServers[test.FakeTrustZoneServerID] = test.FakeTrustZoneServer()

	trustZoneServers, err = client.ListTrustZoneServers(t.Context(), nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*trustzoneserverpb.TrustZoneServer{trustZoneServer}, trustZoneServers)
}

func Test_fakeTrustZoneServerClient_CreateJoinToken(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)

	trustZoneServer := test.FakeTrustZoneServer()

	_, err := client.CreateJoinToken(t.Context(), trustZoneServer.GetId())
	require.Error(t, err)

	fake.TrustZoneServers[trustZoneServer.GetId()] = trustZoneServer

	token, err := client.CreateJoinToken(t.Context(), trustZoneServer.GetId())
	require.NoError(t, err)
	assert.NotEmpty(t, token)
}
