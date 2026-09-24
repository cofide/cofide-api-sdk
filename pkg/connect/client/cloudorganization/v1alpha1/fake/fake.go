// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"slices"

	cloudorganizationsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/cloud_organization_service/v1alpha1"
	cloudorganizationpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_organization/v1alpha1"
	cloudorganizationv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/cloudorganization/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type fakeCloudOrganizationClient struct {
	fake *fakeconnect.FakeConnect
}

// New instantiates a new CloudOrganizationClient for communication with a fake Connect API.
func New(fake *fakeconnect.FakeConnect) cloudorganizationv1alpha1.CloudOrganizationClient {
	return &fakeCloudOrganizationClient{fake: fake}
}

func (c *fakeCloudOrganizationClient) CreateCloudOrganization(ctx context.Context, cloudOrganization *cloudorganizationpb.CloudOrganization) (*cloudorganizationpb.CloudOrganization, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if err := c.fake.ValidateOrganization(cloudOrganization.GetOrgId()); err != nil {
		return nil, err
	}
	cloudOrganization = clone(cloudOrganization)
	cloudOrganization.Id = uuid.New().String()
	cloudOrganization.CreatedAt = timestamppb.Now()
	c.fake.CloudOrganizations[cloudOrganization.GetId()] = cloudOrganization
	return clone(cloudOrganization), nil
}

func (c *fakeCloudOrganizationClient) DestroyCloudOrganization(ctx context.Context, cloudOrganizationID string) error {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.CloudOrganizations[cloudOrganizationID]; !ok {
		return status.Error(codes.NotFound, "cloud organization not found")
	}
	delete(c.fake.CloudOrganizations, cloudOrganizationID)
	return nil
}

func (c *fakeCloudOrganizationClient) GetCloudOrganization(ctx context.Context, cloudOrganizationID string) (*cloudorganizationpb.CloudOrganization, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	cloudOrganization, ok := c.fake.CloudOrganizations[cloudOrganizationID]
	if !ok {
		return nil, status.Error(codes.NotFound, "cloud organization not found")
	}
	return clone(cloudOrganization), nil
}

func (c *fakeCloudOrganizationClient) ListCloudOrganizations(ctx context.Context, filter *cloudorganizationsvcpb.ListCloudOrganizationsRequest_Filter) ([]*cloudorganizationpb.CloudOrganization, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	cloudOrganizations := []*cloudorganizationpb.CloudOrganization{}
	for _, cloudOrganization := range c.fake.CloudOrganizations {
		if cloudOrganizationMatches(cloudOrganization, filter) {
			cloudOrganizations = append(cloudOrganizations, clone(cloudOrganization))
		}
	}
	return cloudOrganizations, nil
}

func cloudOrganizationMatches(cloudOrganization *cloudorganizationpb.CloudOrganization, filter *cloudorganizationsvcpb.ListCloudOrganizationsRequest_Filter) bool {
	if filter == nil {
		return true
	}
	if len(filter.GetOrgIds()) > 0 && !slices.Contains(filter.GetOrgIds(), cloudOrganization.GetOrgId()) {
		return false
	}
	return true
}

func (c *fakeCloudOrganizationClient) UpdateCloudOrganization(
	ctx context.Context,
	cloudOrganization *cloudorganizationpb.CloudOrganization,
	updateMask *cloudorganizationsvcpb.UpdateCloudOrganizationRequest_UpdateMask,
) (*cloudorganizationpb.CloudOrganization, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	existing, ok := c.fake.CloudOrganizations[cloudOrganization.GetId()]
	if !ok {
		return nil, status.Error(codes.NotFound, "cloud organization not found")
	}

	var updated *cloudorganizationpb.CloudOrganization
	if updateMask == nil {
		updated = clone(cloudOrganization)
	} else {
		updated = clone(existing)
		if updateMask.GetName() {
			updated.Name = cloudOrganization.GetName()
		}
		if updateMask.GetDiscoveryEnabled() {
			updated.DiscoveryEnabled = cloudOrganization.GetDiscoveryEnabled()
		}
		if updateMask.GetAwsAudience() || updateMask.GetAwsRoleChain() {
			aws := updated.GetAws()
			if aws == nil {
				aws = &cloudorganizationpb.AWSOrganization{}
				updated.Provider = &cloudorganizationpb.CloudOrganization_Aws{Aws: aws}
			}
			if updateMask.GetAwsAudience() {
				aws.Audience = cloudOrganization.GetAws().GetAudience()
			}
			if updateMask.GetAwsRoleChain() {
				aws.RoleChain = cloudOrganization.GetAws().GetRoleChain()
			}
		}
	}

	updated.LastUpdatedAt = timestamppb.Now()

	// existing may have been partially updated or entirely replaced, so explicitly set it to cover both cases
	c.fake.CloudOrganizations[cloudOrganization.GetId()] = updated
	return clone(updated), nil
}

func clone(cloudOrganization *cloudorganizationpb.CloudOrganization) *cloudorganizationpb.CloudOrganization {
	return proto.Clone(cloudOrganization).(*cloudorganizationpb.CloudOrganization)
}
