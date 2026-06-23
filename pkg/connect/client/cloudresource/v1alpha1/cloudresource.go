// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"

	cloudresourcepb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_resource/v1alpha1"
	cloudresourcesvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/cloud_resource_service/v1alpha1"
	"google.golang.org/grpc"
)

// CloudResourceClient is an interface for a gRPC client for the v1alpha1 version of the Connect CloudResourceService.
type CloudResourceClient interface {
	GetCloudResource(ctx context.Context, cloudResourceID string) (*cloudresourcepb.CloudResource, error)
	ListCloudResources(ctx context.Context, filter *cloudresourcesvcpb.ListCloudResourcesRequest_Filter) ([]*cloudresourcepb.CloudResource, error)
	UpdateCloudResource(ctx context.Context, cloudResource *cloudresourcepb.CloudResource, updateMask *cloudresourcesvcpb.UpdateCloudResourceRequest_UpdateMask) (*cloudresourcepb.CloudResource, error)
}

type cloudResourceClient struct {
	client cloudresourcesvcpb.CloudResourceServiceClient
}

// New instantiates a new CloudResourceClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) CloudResourceClient {
	return &cloudResourceClient{
		client: cloudresourcesvcpb.NewCloudResourceServiceClient(conn),
	}
}

func (c *cloudResourceClient) GetCloudResource(ctx context.Context, cloudResourceID string) (*cloudresourcepb.CloudResource, error) {
	resp, err := c.client.GetCloudResource(ctx, &cloudresourcesvcpb.GetCloudResourceRequest{
		CloudResourceId: cloudResourceID,
	})
	if err != nil {
		return nil, err
	}
	return resp.CloudResource, nil
}

func (c *cloudResourceClient) ListCloudResources(ctx context.Context, filter *cloudresourcesvcpb.ListCloudResourcesRequest_Filter) ([]*cloudresourcepb.CloudResource, error) {
	resp, err := c.client.ListCloudResources(ctx, &cloudresourcesvcpb.ListCloudResourcesRequest{
		Filter: filter,
	})
	if err != nil {
		return nil, err
	}
	return resp.CloudResources, nil
}

func (c *cloudResourceClient) UpdateCloudResource(ctx context.Context, cloudResource *cloudresourcepb.CloudResource, updateMask *cloudresourcesvcpb.UpdateCloudResourceRequest_UpdateMask) (*cloudresourcepb.CloudResource, error) {
	resp, err := c.client.UpdateCloudResource(ctx, &cloudresourcesvcpb.UpdateCloudResourceRequest{
		CloudResource: cloudResource,
		UpdateMask:    updateMask,
	})
	if err != nil {
		return nil, err
	}
	return resp.CloudResource, nil
}
