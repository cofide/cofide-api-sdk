// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	clusterpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cluster/v1alpha1"
	clustersvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/cluster_service/v1alpha1"
	clusterv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/cluster/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type fakeClusterClient struct {
	clusters map[string]*clusterpb.Cluster
}

// New instantiates a new ClusterClient for communication with a Connect API.
func New() clusterv1alpha1.ClusterClient {
	return &fakeClusterClient{}
}

func (c *fakeClusterClient) CreateCluster(ctx context.Context, cluster *clusterpb.Cluster) (*clusterpb.Cluster, error) {
	// TODO: set ID, check exists?
	c.clusters[cluster.GetId()] = cluster
	return cluster, nil
}

func (c *fakeClusterClient) DestroyCluster(ctx context.Context, clusterID string) error {
	delete(c.clusters, clusterID)
	return nil
}

func (c *fakeClusterClient) GetCluster(ctx context.Context, clusterID string) (*clusterpb.Cluster, error) {
	cluster, ok := c.clusters[clusterID]
	if !ok {
		return nil, status.Error(codes.NotFound, "")
	}
	return cluster, nil
}

func (c *fakeClusterClient) ListClusters(ctx context.Context, filter *clustersvcpb.ListClustersRequest_Filter) ([]*clusterpb.Cluster, error) {
	clusters := []*clusterpb.Cluster{}
	for _, cluster := range c.clusters {
		clusters = append(clusters, cluster)
	}
	return clusters, nil
}

func (c *fakeClusterClient) UpdateCluster(ctx context.Context, cluster *clusterpb.Cluster) (*clusterpb.Cluster, error) {
	c.clusters[cluster.GetId()] = cluster
	return cluster, nil
}
