// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"

	cloudorganizationpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_organization/v1alpha1"
	cloudorganizationsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/cloud_organization_service/v1alpha1"
	"google.golang.org/grpc"
)

// CloudOrganizationClient is an interface for a gRPC client for the v1alpha1 version of the Connect CloudOrganizationService.
type CloudOrganizationClient interface {
	CreateCloudOrganization(ctx context.Context, cloudOrganization *cloudorganizationpb.CloudOrganization) (*cloudorganizationpb.CloudOrganization, error)
	DestroyCloudOrganization(ctx context.Context, cloudOrganizationID string) error
	GetCloudOrganization(ctx context.Context, cloudOrganizationID string) (*cloudorganizationpb.CloudOrganization, error)
	ListCloudOrganizations(ctx context.Context, filter *cloudorganizationsvcpb.ListCloudOrganizationsRequest_Filter) ([]*cloudorganizationpb.CloudOrganization, error)
	UpdateCloudOrganization(ctx context.Context, cloudOrganization *cloudorganizationpb.CloudOrganization, updateMask *cloudorganizationsvcpb.UpdateCloudOrganizationRequest_UpdateMask) (*cloudorganizationpb.CloudOrganization, error)
}

type cloudOrganizationClient struct {
	client cloudorganizationsvcpb.CloudOrganizationServiceClient
}

// New instantiates a new CloudOrganizationClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) CloudOrganizationClient {
	return &cloudOrganizationClient{
		client: cloudorganizationsvcpb.NewCloudOrganizationServiceClient(conn),
	}
}

func (c *cloudOrganizationClient) CreateCloudOrganization(ctx context.Context, cloudOrganization *cloudorganizationpb.CloudOrganization) (*cloudorganizationpb.CloudOrganization, error) {
	resp, err := c.client.CreateCloudOrganization(ctx, &cloudorganizationsvcpb.CreateCloudOrganizationRequest{
		CloudOrganization: cloudOrganization,
	})
	if err != nil {
		return nil, err
	}
	return resp.CloudOrganization, nil
}

func (c *cloudOrganizationClient) DestroyCloudOrganization(ctx context.Context, cloudOrganizationID string) error {
	_, err := c.client.DestroyCloudOrganization(ctx, &cloudorganizationsvcpb.DestroyCloudOrganizationRequest{
		CloudOrganizationId: cloudOrganizationID,
	})
	return err
}

func (c *cloudOrganizationClient) GetCloudOrganization(ctx context.Context, cloudOrganizationID string) (*cloudorganizationpb.CloudOrganization, error) {
	resp, err := c.client.GetCloudOrganization(ctx, &cloudorganizationsvcpb.GetCloudOrganizationRequest{
		CloudOrganizationId: cloudOrganizationID,
	})
	if err != nil {
		return nil, err
	}
	return resp.CloudOrganization, nil
}

func (c *cloudOrganizationClient) ListCloudOrganizations(ctx context.Context, filter *cloudorganizationsvcpb.ListCloudOrganizationsRequest_Filter) ([]*cloudorganizationpb.CloudOrganization, error) {
	resp, err := c.client.ListCloudOrganizations(ctx, &cloudorganizationsvcpb.ListCloudOrganizationsRequest{
		Filter: filter,
	})
	if err != nil {
		return nil, err
	}
	return resp.CloudOrganizations, nil
}

func (c *cloudOrganizationClient) UpdateCloudOrganization(
	ctx context.Context,
	cloudOrganization *cloudorganizationpb.CloudOrganization,
	updateMask *cloudorganizationsvcpb.UpdateCloudOrganizationRequest_UpdateMask,
) (*cloudorganizationpb.CloudOrganization, error) {
	resp, err := c.client.UpdateCloudOrganization(ctx, &cloudorganizationsvcpb.UpdateCloudOrganizationRequest{
		CloudOrganization: cloudOrganization,
		UpdateMask:        updateMask,
	})
	if err != nil {
		return nil, err
	}
	return resp.CloudOrganization, nil
}
