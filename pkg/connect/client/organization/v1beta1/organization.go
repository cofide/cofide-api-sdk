// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1beta1

import (
	"context"

	organizationpb "github.com/cofide/cofide-api-sdk/gen/go/cofide/organization/v1beta1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
	"google.golang.org/grpc"
)

// OrganizationClient is an interface for a gRPC client for the v1beta1 version of the Connect OrganizationService.
type OrganizationClient interface {
	GetOrganization(ctx context.Context, orgID string) (*organizationpb.Organization, error)
	GetOrganizationByName(ctx context.Context, name string) (*organizationpb.Organization, error)
	ListOrganizations(ctx context.Context, filter *organizationpb.ListOrganizationsRequest_Filter, pagination pagination.Pagination) ([]*organizationpb.Organization, pagination.Pagination, error)
}

type organizationClient struct {
	client organizationpb.OrganizationServiceClient
}

// New instantiates a new OrganizationClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) OrganizationClient {
	return &organizationClient{
		client: organizationpb.NewOrganizationServiceClient(conn),
	}
}

// GetOrganization retrieves an organization by its ID.
func (c *organizationClient) GetOrganization(ctx context.Context, orgID string) (*organizationpb.Organization, error) {
	return c.getOrganization(ctx, &organizationpb.GetOrganizationRequest{
		Identifier: &organizationpb.GetOrganizationRequest_OrgId{OrgId: orgID},
	})
}

// GetOrganizationByName retrieves an organization by its name, which is unique within Connect.
func (c *organizationClient) GetOrganizationByName(ctx context.Context, name string) (*organizationpb.Organization, error) {
	return c.getOrganization(ctx, &organizationpb.GetOrganizationRequest{
		Identifier: &organizationpb.GetOrganizationRequest_Name{Name: name},
	})
}

func (c *organizationClient) getOrganization(ctx context.Context, req *organizationpb.GetOrganizationRequest) (*organizationpb.Organization, error) {
	resp, err := c.client.GetOrganization(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.GetOrganization(), nil
}

func (c *organizationClient) ListOrganizations(ctx context.Context, filter *organizationpb.ListOrganizationsRequest_Filter, requestPagination pagination.Pagination) ([]*organizationpb.Organization, pagination.Pagination, error) {
	resp, err := c.client.ListOrganizations(ctx, &organizationpb.ListOrganizationsRequest{
		Filter:     filter,
		Pagination: requestPagination.PageRequestV1Beta1(),
	})
	if err != nil {
		return nil, pagination.Pagination{}, err
	}
	return resp.GetOrganizations(), requestPagination.Next(resp.GetPagination()), nil
}
