// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	clusterpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cluster/v1alpha1"
	clustersvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/cluster_service/v1alpha1"
	clusterv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/cluster/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type fakeClusterClient struct {
	fake *fakeconnect.FakeConnect
}

// New instantiates a new ClusterClient for communication with a fake Connect API.
func New(fake *fakeconnect.FakeConnect) clusterv1alpha1.ClusterClient {
	return &fakeClusterClient{fake: fake}
}

func (c *fakeClusterClient) CreateCluster(ctx context.Context, cluster *clusterpb.Cluster) (*clusterpb.Cluster, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if err := c.fake.ValidateTrustZone(cluster.GetTrustZoneId()); err != nil {
		return nil, err
	}
	id := uuid.New().String()
	cluster = clone(cluster)
	cluster.Id = &id
	c.fake.Clusters[cluster.GetId()] = cluster
	return clone(cluster), nil
}

func (c *fakeClusterClient) DestroyCluster(ctx context.Context, clusterID string) error {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.Clusters[clusterID]; !ok {
		return status.Error(codes.NotFound, "cluster not found")
	}
	delete(c.fake.Clusters, clusterID)
	return nil
}

func (c *fakeClusterClient) GetCluster(ctx context.Context, clusterID string) (*clusterpb.Cluster, error) {
	cluster, ok := c.fake.Clusters[clusterID]
	if !ok {
		return nil, status.Error(codes.NotFound, "cluster not found")
	}
	return clone(cluster), nil
}

func (c *fakeClusterClient) ListClusters(ctx context.Context, filter *clustersvcpb.ListClustersRequest_Filter) ([]*clusterpb.Cluster, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	clusters := []*clusterpb.Cluster{}
	for _, cluster := range c.fake.Clusters {
		if clusterMatches(cluster, filter) {
			clusters = append(clusters, clone(cluster))
		}
	}
	return clusters, nil
}

func clusterMatches(cluster *clusterpb.Cluster, filter *clustersvcpb.ListClustersRequest_Filter) bool {
	if filter == nil {
		return true
	}
	if filter.Name != nil && cluster.GetName() != *filter.Name {
		return false
	}
	if filter.OrgId != nil && cluster.GetOrgId() != *filter.OrgId {
		return false
	}
	if filter.TrustZoneId != nil && cluster.GetTrustZoneId() != *filter.TrustZoneId {
		return false
	}
	return true
}

func (c *fakeClusterClient) UpdateCluster(ctx context.Context, cluster *clusterpb.Cluster) (*clusterpb.Cluster, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.Clusters[cluster.GetId()]; !ok {
		return nil, status.Error(codes.NotFound, "")
	}
	if _, ok := c.fake.TrustZones[cluster.GetTrustZoneId()]; !ok {
		return nil, status.Error(codes.InvalidArgument, "")
	}
	c.fake.Clusters[cluster.GetId()] = clone(cluster)
	return clone(cluster), nil
}

func clone(cluster *clusterpb.Cluster) *clusterpb.Cluster {
	return proto.Clone(cluster).(*clusterpb.Cluster)
}
