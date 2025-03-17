// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"

	clusterpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cluster/v1alpha1"
	clustersvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/cluster_service/v1alpha1"
	"google.golang.org/grpc"
)

// ClusterClient is an interface for a gRPC client for the v1alpha1 version of the Connect ClusterService.
type ClusterClient interface {
	CreateCluster(ctx context.Context, cluster *clusterpb.Cluster) (*clusterpb.Cluster, error)
	DestroyCluster(ctx context.Context, clusterID string) error
	GetCluster(ctx context.Context, clusterID string) (*clusterpb.Cluster, error)
	ListClusters(ctx context.Context, filter *clustersvcpb.ListClustersRequest_Filter) ([]*clusterpb.Cluster, error)
	UpdateCluster(ctx context.Context, cluster *clusterpb.Cluster) (*clusterpb.Cluster, error)
}

type clusterClient struct {
	client clustersvcpb.ClusterServiceClient
}

// New instantiates a new ClusterClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) ClusterClient {
	return &clusterClient{
		client: clustersvcpb.NewClusterServiceClient(conn),
	}
}

func (c *clusterClient) CreateCluster(ctx context.Context, cluster *clusterpb.Cluster) (*clusterpb.Cluster, error) {
	resp, err := c.client.CreateCluster(ctx, &clustersvcpb.CreateClusterRequest{
		Cluster: cluster,
	})
	if err != nil {
		return nil, err
	}
	return resp.Cluster, nil
}

func (c *clusterClient) DestroyCluster(ctx context.Context, clusterID string) error {
	_, err := c.client.DestroyCluster(ctx, &clustersvcpb.DestroyClusterRequest{
		ClusterId: &clusterID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *clusterClient) GetCluster(ctx context.Context, clusterID string) (*clusterpb.Cluster, error) {
	resp, err := c.client.GetCluster(ctx, &clustersvcpb.GetClusterRequest{
		ClusterId: &clusterID,
	})
	if err != nil {
		return nil, err
	}
	return resp.Cluster, nil
}

func (c *clusterClient) ListClusters(ctx context.Context, filter *clustersvcpb.ListClustersRequest_Filter) ([]*clusterpb.Cluster, error) {
	resp, err := c.client.ListClusters(ctx, &clustersvcpb.ListClustersRequest{
		Filter: filter,
	})
	if err != nil {
		return nil, err
	}
	return resp.Clusters, nil
}

func (c *clusterClient) UpdateCluster(ctx context.Context, cluster *clusterpb.Cluster) (*clusterpb.Cluster, error) {
	resp, err := c.client.UpdateCluster(ctx, &clustersvcpb.UpdateClusterRequest{
		Cluster: cluster,
	})
	if err != nil {
		return nil, err
	}
	return resp.Cluster, nil
}
