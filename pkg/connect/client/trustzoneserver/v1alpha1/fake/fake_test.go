// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"testing"
	"time"

	"github.com/cofide/cofide-api-sdk/gen/go/proto/connect/trust_zone_server_service/v1alpha1"
	trustzoneserverpb "github.com/cofide/cofide-api-sdk/gen/go/proto/trust_zone_server/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

func Test_fakeTrustZoneServerClient_CreateTrustZoneServer(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)

	trustZoneServer := test.FakeTrustZoneServer()

	now := time.Now()
	createdTrustZoneServer, err := client.CreateTrustZoneServer(t.Context(), trustZoneServer)
	require.NoError(t, err)
	trustZoneServer.Id = createdTrustZoneServer.Id
	assert.GreaterOrEqual(t, createdTrustZoneServer.GetCreatedAt().AsTime(), now)
	assert.Nil(t, createdTrustZoneServer.GetLastUpdatedAt())
	assert.Nil(t, createdTrustZoneServer.GetDeletedAt())
	trustZoneServer.CreatedAt = createdTrustZoneServer.CreatedAt
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
	now := time.Now()
	updatedTrustZoneServer, err := client.UpdateTrustZoneServer(t.Context(), trustZoneServer, &v1alpha1.UpdateTrustZoneServerRequest_UpdateMask{})
	require.NoError(t, err)
	assert.GreaterOrEqual(t, updatedTrustZoneServer.GetLastUpdatedAt().AsTime(), now)
	assert.Nil(t, updatedTrustZoneServer.GetDeletedAt())
	existingTrustZoneServer.LastUpdatedAt = updatedTrustZoneServer.LastUpdatedAt
	assert.EqualExportedValues(t, existingTrustZoneServer, updatedTrustZoneServer)
	assert.EqualExportedValues(t, updatedTrustZoneServer, fake.TrustZoneServers[updatedTrustZoneServer.GetId()])

	// Changes with non-empty update mask
	now = time.Now()
	updatedTrustZoneServer, err = client.UpdateTrustZoneServer(t.Context(), trustZoneServer, &v1alpha1.UpdateTrustZoneServerRequest_UpdateMask{HelmValues: true})
	require.NoError(t, err)
	assert.GreaterOrEqual(t, updatedTrustZoneServer.GetLastUpdatedAt().AsTime(), now)
	assert.Nil(t, updatedTrustZoneServer.GetDeletedAt())
	trustZoneServer.LastUpdatedAt = updatedTrustZoneServer.LastUpdatedAt
	assert.EqualExportedValues(t, trustZoneServer, updatedTrustZoneServer)
	assert.EqualExportedValues(t, updatedTrustZoneServer, fake.TrustZoneServers[updatedTrustZoneServer.GetId()])
}

func Test_fakeTrustZoneServerClient_DestroyTrustZoneServer(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)

	err := client.DestroyTrustZoneServer(t.Context(), test.FakeTrustZoneServerID)
	require.Error(t, err)

	fake.TrustZoneServers[test.FakeTrustZoneServerID] = test.FakeTrustZoneServer()

	now := time.Now()
	err = client.DestroyTrustZoneServer(t.Context(), test.FakeTrustZoneServerID)
	require.NoError(t, err)
	assert.NotNil(t, fake.TrustZoneServers[test.FakeTrustZoneServerID])
	assert.GreaterOrEqual(t, fake.TrustZoneServers[test.FakeTrustZoneServerID].GetDeletedAt().AsTime(), now)
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

func Test_fakeTrustZoneServerClient_UpdateTrustZoneServerStatus(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)

	trustZoneServer, err := client.CreateTrustZoneServer(t.Context(), test.FakeTrustZoneServer())
	require.NoError(t, err)

	// Simple status change
	now := time.Now()
	err = client.UpdateTrustZoneServerStatus(t.Context(), trustZoneServer.GetId(), trustzoneserverpb.TrustZoneServerStatus_TRUST_ZONE_SERVER_STATUS_PROVISIONED)
	require.NoError(t, err)
	trustZoneServer, err = client.GetTrustZoneServer(t.Context(), trustZoneServer.GetId())
	require.NoError(t, err)
	assert.Equal(t, trustzoneserverpb.TrustZoneServerStatus_TRUST_ZONE_SERVER_STATUS_PROVISIONED, trustZoneServer.GetStatus().GetStatus())
	assert.GreaterOrEqual(t, trustZoneServer.GetStatus().GetLastTransitionTime().AsTime(), now)

	// Status change to deleted removes resource
	err = client.UpdateTrustZoneServerStatus(t.Context(), trustZoneServer.GetId(), trustzoneserverpb.TrustZoneServerStatus_TRUST_ZONE_SERVER_STATUS_DELETED)
	require.NoError(t, err)
	_, err = client.GetTrustZoneServer(t.Context(), trustZoneServer.GetId())
	require.Error(t, err)
	assert.Equal(t, status.Code(err), codes.NotFound)
}
