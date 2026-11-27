// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	organizationpb "github.com/cofide/cofide-api-sdk/gen/go/cofide/organization/v1beta1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	organizationv1beta1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/organization/v1beta1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// errUnimplemented is returned by every method, since the fake does not yet implement any of them.
var errUnimplemented = status.Error(codes.Unimplemented, "not implemented by the fake v1beta1 OrganizationClient")

type fakeOrganizationClient struct{}

// New instantiates a new OrganizationClient for communication with a fake Connect API.
// Every method currently returns an Unimplemented error, so it does not yet use the fake Connect.
func New(_ *fakeconnect.FakeConnect) organizationv1beta1.OrganizationClient {
	return &fakeOrganizationClient{}
}

func (c *fakeOrganizationClient) GetOrganization(ctx context.Context, orgID string) (*organizationpb.Organization, error) {
	return nil, errUnimplemented
}

func (c *fakeOrganizationClient) GetOrganizationByName(ctx context.Context, name string) (*organizationpb.Organization, error) {
	return nil, errUnimplemented
}

func (c *fakeOrganizationClient) ListOrganizations(ctx context.Context, filter *organizationpb.ListOrganizationsRequest_Filter, requestPagination pagination.Pagination) ([]*organizationpb.Organization, pagination.Pagination, error) {
	return nil, pagination.Pagination{}, errUnimplemented
}
