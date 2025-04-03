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
	UpdateAttestedNode(ctx context.Context, req *datastorev1alpha1.UpdateAttestedNodeRequest) (*datastorev1alpha1.UpdateAttestedNodeResponse, error)
	GetNodeSelectors(ctx context.Context, req *datastorev1alpha1.GetNodeSelectorsRequest) (*datastorev1alpha1.GetNodeSelectorsResponse, error)
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
	resp, err := c.client.CountAttestedNodes(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *datastoreClient) CreateAttestedNode(ctx context.Context, req *datastorev1alpha1.CreateAttestedNodeRequest) (*datastorev1alpha1.CreateAttestedNodeResponse, error) {
	resp, err := c.client.CreateAttestedNode(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *datastoreClient) DeleteAttestedNode(ctx context.Context, req *datastorev1alpha1.DeleteAttestedNodeRequest) (*datastorev1alpha1.DeleteAttestedNodeResponse, error) {
	resp, err := c.client.DeleteAttestedNode(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *datastoreClient) FetchAttestedNode(ctx context.Context, req *datastorev1alpha1.FetchAttestedNodeRequest) (*datastorev1alpha1.FetchAttestedNodeResponse, error) {
	resp, err := c.client.FetchAttestedNode(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *datastoreClient) UpdateAttestedNode(ctx context.Context, req *datastorev1alpha1.UpdateAttestedNodeRequest) (*datastorev1alpha1.UpdateAttestedNodeResponse, error) {
	resp, err := c.client.UpdateAttestedNode(ctx, &datastorev1alpha1.UpdateAttestedNodeRequest{})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *datastoreClient) GetNodeSelectors(ctx context.Context, req *datastorev1alpha1.GetNodeSelectorsRequest) (*datastorev1alpha1.GetNodeSelectorsResponse, error) {
	resp, err := c.client.GetNodeSelectors(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *datastoreClient) SetNodeSelectors(ctx context.Context, req *datastorev1alpha1.SetNodeSelectorsRequest) (*datastorev1alpha1.SetNodeSelectorsResponse, error) {
	resp, err := c.client.SetNodeSelectors(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
