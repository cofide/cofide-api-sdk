// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"

	organizationsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/organization_service/v1alpha1"
	organizationpb "github.com/cofide/cofide-api-sdk/gen/go/proto/organization/v1alpha1"
	"google.golang.org/grpc"
)

// OrganizationClient is an interface for a gRPC client for the v1alpha1 version of the Connect OrganizationService.
type OrganizationClient interface {
	GetOrganization(ctx context.Context, organizationID string) (*organizationpb.Organization, error)
	ListOrganizations(ctx context.Context, filter *organizationsvcpb.ListOrganizationsRequest_Filter) ([]*organizationpb.Organization, error)
}

type organizationClient struct {
	client organizationsvcpb.OrganizationServiceClient
}

// New instantiates a new OrganizationClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) OrganizationClient {
	return &organizationClient{
		client: organizationsvcpb.NewOrganizationServiceClient(conn),
	}
}

func (c *organizationClient) GetOrganization(ctx context.Context, organizationID string) (*organizationpb.Organization, error) {
	resp, err := c.client.GetOrganization(ctx, &organizationsvcpb.GetOrganizationRequest{
		OrgId: &organizationID,
	})
	if err != nil {
		return nil, err
	}
	return resp.Organization, nil
}

func (c *organizationClient) ListOrganizations(ctx context.Context, filter *organizationsvcpb.ListOrganizationsRequest_Filter) ([]*organizationpb.Organization, error) {
	resp, err := c.client.ListOrganizations(ctx, &organizationsvcpb.ListOrganizationsRequest{
		Filter: filter,
	})
	if err != nil {
		return nil, err
	}
	return resp.Organizations, nil
}
