// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"testing"

	clusterpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cluster/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_fakeClusterClient_CreateCluster(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	cluster := test.FakeCluster()

	_, err := client.CreateCluster(ctx, cluster)
	require.Error(t, err)

	fake.TrustZones[test.FakeTrustZoneID] = test.FakeTrustZone()

	createdCluster, err := client.CreateCluster(ctx, cluster)
	require.NoError(t, err)
	cluster.Id = createdCluster.Id
	assert.EqualExportedValues(t, cluster, createdCluster)
	assert.EqualExportedValues(t, cluster, fake.Clusters[*createdCluster.Id])
}

func Test_fakeClusterClient_DestroyCluster(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	err := client.DestroyCluster(ctx, test.FakeClusterID)
	require.Error(t, err)

	fake.Clusters[test.FakeClusterID] = test.FakeCluster()

	err = client.DestroyCluster(ctx, test.FakeClusterID)
	require.NoError(t, err)
	require.Nil(t, fake.Clusters[test.FakeClusterID])
}

func Test_fakeClusterClient_GetCluster(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	_, err := client.GetCluster(ctx, test.FakeClusterID)
	require.Error(t, err)

	cluster := test.FakeCluster()
	fake.Clusters[test.FakeClusterID] = cluster

	gotCluster, err := client.GetCluster(ctx, test.FakeClusterID)
	require.NoError(t, err)
	assert.EqualExportedValues(t, cluster, gotCluster)
}

func Test_fakeClusterClient_ListClusters(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	_, err := client.ListClusters(ctx, nil)
	require.NoError(t, err)

	cluster := test.FakeCluster()
	fake.Clusters[test.FakeClusterID] = cluster

	clusters, err := client.ListClusters(ctx, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*clusterpb.Cluster{cluster}, clusters)
}

func Test_fakeClusterClient_UpdateCluster(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	cluster := test.FakeCluster()

	_, err := client.UpdateCluster(ctx, cluster)
	require.Error(t, err)

	fake.TrustZones[test.FakeTrustZoneID] = test.FakeTrustZone()
	fake.Clusters[test.FakeClusterID] = cluster

	cluster = clone(cluster)
	cluster.Name = test.PtrOf("new-cluster-name")
	updatedCluster, err := client.UpdateCluster(ctx, cluster)
	require.NoError(t, err)
	assert.EqualExportedValues(t, cluster, updatedCluster)
	assert.EqualExportedValues(t, cluster, fake.Clusters[test.FakeClusterID])
}
