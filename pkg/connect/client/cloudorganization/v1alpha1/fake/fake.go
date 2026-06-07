// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	cloudorgpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_organization/v1alpha1"
	cloudorgsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/cloud_organization_service/v1alpha1"
	cloudorganizationv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/cloudorganization/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type fakeCloudOrganizationClient struct {
	fake *fakeconnect.FakeConnect
}

// New instantiates a new CloudOrganizationClient for communication with a fake Connect API.
func New(fake *fakeconnect.FakeConnect) cloudorganizationv1alpha1.CloudOrganizationClient {
	return &fakeCloudOrganizationClient{fake: fake}
}

func (c *fakeCloudOrganizationClient) CreateCloudOrganization(ctx context.Context, cloudOrg *cloudorgpb.CloudOrganization) (*cloudorgpb.CloudOrganization, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	id := uuid.New().String()
	cloudOrg = clone(cloudOrg)
	cloudOrg.Id = id
	c.fake.CloudOrganizations[cloudOrg.GetId()] = cloudOrg
	return clone(cloudOrg), nil
}

func (c *fakeCloudOrganizationClient) GetCloudOrganization(ctx context.Context, cloudOrgID string) (*cloudorgpb.CloudOrganization, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	cloudOrg, ok := c.fake.CloudOrganizations[cloudOrgID]
	if !ok {
		return nil, status.Error(codes.NotFound, "cloud organization not found")
	}
	return clone(cloudOrg), nil
}

func (c *fakeCloudOrganizationClient) ListCloudOrganizations(ctx context.Context, filter *cloudorgsvcpb.ListCloudOrganizationsRequest_Filter) ([]*cloudorgpb.CloudOrganization, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	cloudOrgs := []*cloudorgpb.CloudOrganization{}
	for _, cloudOrg := range c.fake.CloudOrganizations {
		if cloudOrgMatches(cloudOrg, filter) {
			cloudOrgs = append(cloudOrgs, clone(cloudOrg))
		}
	}
	return cloudOrgs, nil
}

func cloudOrgMatches(cloudOrg *cloudorgpb.CloudOrganization, filter *cloudorgsvcpb.ListCloudOrganizationsRequest_Filter) bool {
	if filter == nil {
		return true
	}
	if filter.OrgId != "" && cloudOrg.GetOrgId() != filter.OrgId {
		return false
	}
	return true
}

func (c *fakeCloudOrganizationClient) UpdateCloudOrganization(ctx context.Context, cloudOrg *cloudorgpb.CloudOrganization, updateMask *cloudorgsvcpb.UpdateCloudOrganizationRequest_UpdateMask) (*cloudorgpb.CloudOrganization, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.CloudOrganizations[cloudOrg.GetId()]; !ok {
		return nil, status.Error(codes.NotFound, "cloud organization not found")
	}
	c.fake.CloudOrganizations[cloudOrg.GetId()] = clone(cloudOrg)
	return clone(cloudOrg), nil
}

func (c *fakeCloudOrganizationClient) DestroyCloudOrganization(ctx context.Context, cloudOrgID string) error {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.CloudOrganizations[cloudOrgID]; !ok {
		return status.Error(codes.NotFound, "cloud organization not found")
	}
	delete(c.fake.CloudOrganizations, cloudOrgID)
	return nil
}

func clone(cloudOrg *cloudorgpb.CloudOrganization) *cloudorgpb.CloudOrganization {
	return proto.Clone(cloudOrg).(*cloudorgpb.CloudOrganization)
}
