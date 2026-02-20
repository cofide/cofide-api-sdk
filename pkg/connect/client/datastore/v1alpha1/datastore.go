// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"

	datastorev1alpha1 "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/datastore_service/v1alpha1"
	"google.golang.org/grpc"
)

// DataStoreClient is an interface for a client for the v1alpha1 version of the Connect DataStoreService.
type DataStoreClient interface {
	CountAttestedNodes(ctx context.Context, req *datastorev1alpha1.CountAttestedNodesRequest) (*datastorev1alpha1.CountAttestedNodesResponse, error)
	CreateAttestedNode(ctx context.Context, req *datastorev1alpha1.CreateAttestedNodeRequest) (*datastorev1alpha1.CreateAttestedNodeResponse, error)
	DeleteAttestedNode(ctx context.Context, req *datastorev1alpha1.DeleteAttestedNodeRequest) (*datastorev1alpha1.DeleteAttestedNodeResponse, error)
	FetchAttestedNode(ctx context.Context, req *datastorev1alpha1.FetchAttestedNodeRequest) (*datastorev1alpha1.FetchAttestedNodeResponse, error)
	ListAttestedNodes(ctx context.Context, req *datastorev1alpha1.ListAttestedNodesRequest) (*datastorev1alpha1.ListAttestedNodesResponse, error)
	UpdateAttestedNode(ctx context.Context, req *datastorev1alpha1.UpdateAttestedNodeRequest) (*datastorev1alpha1.UpdateAttestedNodeResponse, error)
	PruneAttestedExpiredNodes(ctx context.Context, req *datastorev1alpha1.PruneAttestedExpiredNodesRequest) (*datastorev1alpha1.PruneAttestedExpiredNodesResponse, error)
	GetNodeSelectors(ctx context.Context, req *datastorev1alpha1.GetNodeSelectorsRequest) (*datastorev1alpha1.GetNodeSelectorsResponse, error)
	ListNodeSelectors(ctx context.Context, req *datastorev1alpha1.ListNodeSelectorsRequest) (*datastorev1alpha1.ListNodeSelectorsResponse, error)
	SetNodeSelectors(ctx context.Context, req *datastorev1alpha1.SetNodeSelectorsRequest) (*datastorev1alpha1.SetNodeSelectorsResponse, error)
}

type datastoreClient struct {
	client datastorev1alpha1.DataStoreServiceClient
}

// New instantiates a new DataStoreClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) DataStoreClient {
	return &datastoreClient{
		client: datastorev1alpha1.NewDataStoreServiceClient(conn),
	}
}

func (c *datastoreClient) CountAttestedNodes(ctx context.Context, req *datastorev1alpha1.CountAttestedNodesRequest) (*datastorev1alpha1.CountAttestedNodesResponse, error) {
	return c.client.CountAttestedNodes(ctx, req)
}

func (c *datastoreClient) CreateAttestedNode(ctx context.Context, req *datastorev1alpha1.CreateAttestedNodeRequest) (*datastorev1alpha1.CreateAttestedNodeResponse, error) {
	return c.client.CreateAttestedNode(ctx, req)
}

func (c *datastoreClient) DeleteAttestedNode(ctx context.Context, req *datastorev1alpha1.DeleteAttestedNodeRequest) (*datastorev1alpha1.DeleteAttestedNodeResponse, error) {
	return c.client.DeleteAttestedNode(ctx, req)
}

func (c *datastoreClient) FetchAttestedNode(ctx context.Context, req *datastorev1alpha1.FetchAttestedNodeRequest) (*datastorev1alpha1.FetchAttestedNodeResponse, error) {
	return c.client.FetchAttestedNode(ctx, req)
}

func (c *datastoreClient) ListAttestedNodes(ctx context.Context, req *datastorev1alpha1.ListAttestedNodesRequest) (*datastorev1alpha1.ListAttestedNodesResponse, error) {
	return c.client.ListAttestedNodes(ctx, req)
}

func (c *datastoreClient) UpdateAttestedNode(ctx context.Context, req *datastorev1alpha1.UpdateAttestedNodeRequest) (*datastorev1alpha1.UpdateAttestedNodeResponse, error) {
	return c.client.UpdateAttestedNode(ctx, req)
}

func (c *datastoreClient) PruneAttestedExpiredNodes(ctx context.Context, req *datastorev1alpha1.PruneAttestedExpiredNodesRequest) (*datastorev1alpha1.PruneAttestedExpiredNodesResponse, error) {
	return c.client.PruneAttestedExpiredNodes(ctx, req)
}

func (c *datastoreClient) GetNodeSelectors(ctx context.Context, req *datastorev1alpha1.GetNodeSelectorsRequest) (*datastorev1alpha1.GetNodeSelectorsResponse, error) {
	return c.client.GetNodeSelectors(ctx, req)
}

func (c *datastoreClient) ListNodeSelectors(ctx context.Context, req *datastorev1alpha1.ListNodeSelectorsRequest) (*datastorev1alpha1.ListNodeSelectorsResponse, error) {
	return c.client.ListNodeSelectors(ctx, req)
}

func (c *datastoreClient) SetNodeSelectors(ctx context.Context, req *datastorev1alpha1.SetNodeSelectorsRequest) (*datastorev1alpha1.SetNodeSelectorsResponse, error) {
	return c.client.SetNodeSelectors(ctx, req)
}
