// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	organizationsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/organization_service/v1alpha1"
	organizationpb "github.com/cofide/cofide-api-sdk/gen/go/proto/organization/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	organizationv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/organization/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type fakeOrganizationClient struct {
	fake *fakeconnect.FakeConnect
}

// New instantiates a new OrganizationClient for communication with a fake Connect API.
func New(fake *fakeconnect.FakeConnect) organizationv1alpha1.OrganizationClient {
	return &fakeOrganizationClient{fake: fake}
}

func (c *fakeOrganizationClient) GetOrganization(ctx context.Context, organizationID string) (*organizationpb.Organization, error) {
	organization, ok := c.fake.Organizations[organizationID]
	if !ok {
		return nil, status.Error(codes.NotFound, "organization not found")
	}
	return clone(organization), nil
}

func (c *fakeOrganizationClient) ListOrganizations(ctx context.Context, filter *organizationsvcpb.ListOrganizationsRequest_Filter) ([]*organizationpb.Organization, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	organizations := []*organizationpb.Organization{}
	for _, organization := range c.fake.Organizations {
		if organizationMatches(organization, filter) {
			organizations = append(organizations, clone(organization))
		}
	}
	return organizations, nil
}

func organizationMatches(organization *organizationpb.Organization, filter *organizationsvcpb.ListOrganizationsRequest_Filter) bool {
	if filter == nil {
		return true
	}
	if filter.Name != nil && organization.GetName() != *filter.Name {
		return false
	}
	return true
}

func clone(organization *organizationpb.Organization) *organizationpb.Organization {
	return proto.Clone(organization).(*organizationpb.Organization)
}
