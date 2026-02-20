// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"
	"testing"

	clusterpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cluster/v1alpha1"
	clustersvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/cluster_service/v1alpha1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	fakeClusterID   = "fake-cluster-id"
	fakeClusterName = "fake-cluster-name"
)

func TestClusterClient_implementsMethods(t *testing.T) {
	test.AssertClientImplementsService(t, &clusterClient{}, clustersvcpb.ClusterService_ServiceDesc)
}

// TestClusterClient_Unimplemented tests ClusterClient against an unimplemented server.
// This ensures that all errors returned are not wrapped can be converted to a gRPC Status using Status.Convert.
func TestClusterClient_Unimplemented(t *testing.T) {
	server := test.NewTestServer(t)
	clustersvcpb.RegisterClusterServiceServer(server.Server, &clustersvcpb.UnimplementedClusterServiceServer{})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	token, err := client.CreateCluster(ctx, nil)
	test.RequireUnimplemented(t, err)
	assert.Empty(t, token)

	err = client.DestroyCluster(ctx, "")
	test.RequireUnimplemented(t, err)

	cluster, err := client.GetCluster(ctx, "")
	test.RequireUnimplemented(t, err)
	assert.Nil(t, cluster)

	clusters, err := client.ListClusters(ctx, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, clusters)

	cluster, err = client.UpdateCluster(ctx, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, cluster)
}

func TestClusterClient(t *testing.T) {
	server := test.NewTestServer(t)
	clustersvcpb.RegisterClusterServiceServer(server.Server, &fakeClusterService{t: t})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	cluster := fakeCluster()

	createdCluster, err := client.CreateCluster(ctx, cluster)
	require.NoError(t, err)
	assert.EqualExportedValues(t, cluster, createdCluster)

	err = client.DestroyCluster(ctx, cluster.GetId())
	require.NoError(t, err)

	gotCluster, err := client.GetCluster(ctx, cluster.GetId())
	require.NoError(t, err)
	assert.Equal(t, cluster.GetId(), gotCluster.GetId())

	filter := &clustersvcpb.ListClustersRequest_Filter{Name: test.PtrOf(fakeClusterName)}
	clusters, err := client.ListClusters(ctx, filter)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*clusterpb.Cluster{fakeCluster()}, clusters)

	updatedCluster, err := client.UpdateCluster(ctx, cluster)
	require.NoError(t, err)
	assert.EqualExportedValues(t, cluster, updatedCluster)
}

// fakeClusterService provides a very simple fake ClusterService implementation with canned responses.
type fakeClusterService struct {
	t *testing.T
}

func (f *fakeClusterService) CreateCluster(ctx context.Context, req *clustersvcpb.CreateClusterRequest) (*clustersvcpb.CreateClusterResponse, error) {
	assert.EqualExportedValues(f.t, fakeCluster(), req.Cluster)
	return &clustersvcpb.CreateClusterResponse{Cluster: req.Cluster}, nil
}

func (f *fakeClusterService) DestroyCluster(ctx context.Context, req *clustersvcpb.DestroyClusterRequest) (*clustersvcpb.DestroyClusterResponse, error) {
	assert.Equal(f.t, fakeClusterID, req.GetClusterId())
	return &clustersvcpb.DestroyClusterResponse{}, nil
}

func (f *fakeClusterService) GetCluster(ctx context.Context, req *clustersvcpb.GetClusterRequest) (*clustersvcpb.GetClusterResponse, error) {
	assert.Equal(f.t, fakeClusterID, req.GetClusterId())
	return &clustersvcpb.GetClusterResponse{Cluster: fakeCluster()}, nil
}

func (f *fakeClusterService) ListClusters(ctx context.Context, req *clustersvcpb.ListClustersRequest) (*clustersvcpb.ListClustersResponse, error) {
	assert.Equal(f.t, fakeClusterName, req.Filter.GetName())
	clusters := []*clusterpb.Cluster{fakeCluster()}
	return &clustersvcpb.ListClustersResponse{Clusters: clusters}, nil
}

func (f *fakeClusterService) UpdateCluster(ctx context.Context, req *clustersvcpb.UpdateClusterRequest) (*clustersvcpb.UpdateClusterResponse, error) {
	assert.EqualExportedValues(f.t, fakeCluster(), req.Cluster)
	return &clustersvcpb.UpdateClusterResponse{Cluster: fakeCluster()}, nil
}

func fakeCluster() *clusterpb.Cluster {
	return &clusterpb.Cluster{
		Id:   test.PtrOf(fakeClusterID),
		Name: test.PtrOf(fakeClusterName),
	}
}
