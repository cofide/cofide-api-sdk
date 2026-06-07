// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"

	cloudorgpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_organization/v1alpha1"
	cloudorgsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/cloud_organization_service/v1alpha1"
	"google.golang.org/grpc"
)

// CloudOrganizationClient is an interface for a gRPC client for the v1alpha1 version of the Connect CloudOrganizationService.
type CloudOrganizationClient interface {
	CreateCloudOrganization(ctx context.Context, cloudOrg *cloudorgpb.CloudOrganization) (*cloudorgpb.CloudOrganization, error)
	GetCloudOrganization(ctx context.Context, cloudOrgID string) (*cloudorgpb.CloudOrganization, error)
	ListCloudOrganizations(ctx context.Context, filter *cloudorgsvcpb.ListCloudOrganizationsRequest_Filter) ([]*cloudorgpb.CloudOrganization, error)
	UpdateCloudOrganization(ctx context.Context, cloudOrg *cloudorgpb.CloudOrganization, updateMask *cloudorgsvcpb.UpdateCloudOrganizationRequest_UpdateMask) (*cloudorgpb.CloudOrganization, error)
	DestroyCloudOrganization(ctx context.Context, cloudOrgID string) error
}

type cloudOrganizationClient struct {
	client cloudorgsvcpb.CloudOrganizationServiceClient
}

// New instantiates a new CloudOrganizationClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) CloudOrganizationClient {
	return &cloudOrganizationClient{
		client: cloudorgsvcpb.NewCloudOrganizationServiceClient(conn),
	}
}

func (c *cloudOrganizationClient) CreateCloudOrganization(ctx context.Context, cloudOrg *cloudorgpb.CloudOrganization) (*cloudorgpb.CloudOrganization, error) {
	resp, err := c.client.CreateCloudOrganization(ctx, &cloudorgsvcpb.CreateCloudOrganizationRequest{
		CloudOrganization: cloudOrg,
	})
	if err != nil {
		return nil, err
	}
	return resp.CloudOrganization, nil
}

func (c *cloudOrganizationClient) GetCloudOrganization(ctx context.Context, cloudOrgID string) (*cloudorgpb.CloudOrganization, error) {
	resp, err := c.client.GetCloudOrganization(ctx, &cloudorgsvcpb.GetCloudOrganizationRequest{
		CloudOrganizationId: cloudOrgID,
	})
	if err != nil {
		return nil, err
	}
	return resp.CloudOrganization, nil
}

func (c *cloudOrganizationClient) ListCloudOrganizations(ctx context.Context, filter *cloudorgsvcpb.ListCloudOrganizationsRequest_Filter) ([]*cloudorgpb.CloudOrganization, error) {
	resp, err := c.client.ListCloudOrganizations(ctx, &cloudorgsvcpb.ListCloudOrganizationsRequest{
		Filter: filter,
	})
	if err != nil {
		return nil, err
	}
	return resp.CloudOrganizations, nil
}

func (c *cloudOrganizationClient) UpdateCloudOrganization(ctx context.Context, cloudOrg *cloudorgpb.CloudOrganization, updateMask *cloudorgsvcpb.UpdateCloudOrganizationRequest_UpdateMask) (*cloudorgpb.CloudOrganization, error) {
	resp, err := c.client.UpdateCloudOrganization(ctx, &cloudorgsvcpb.UpdateCloudOrganizationRequest{
		CloudOrganization: cloudOrg,
		UpdateMask:        updateMask,
	})
	if err != nil {
		return nil, err
	}
	return resp.CloudOrganization, nil
}

func (c *cloudOrganizationClient) DestroyCloudOrganization(ctx context.Context, cloudOrgID string) error {
	_, err := c.client.DestroyCloudOrganization(ctx, &cloudorgsvcpb.DestroyCloudOrganizationRequest{
		CloudOrganizationId: cloudOrgID,
	})
	return err
}
