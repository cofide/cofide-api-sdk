// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"slices"

	cloudaccountsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/cloud_account_service/v1alpha1"
	cloudaccountpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_account/v1alpha1"
	cloudaccountv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/cloudaccount/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type fakeCloudAccountClient struct {
	fake *fakeconnect.FakeConnect
}

// New instantiates a new CloudAccountClient for communication with a fake Connect API.
func New(fake *fakeconnect.FakeConnect) cloudaccountv1alpha1.CloudAccountClient {
	return &fakeCloudAccountClient{fake: fake}
}

func (c *fakeCloudAccountClient) CreateCloudAccount(ctx context.Context, cloudAccount *cloudaccountpb.CloudAccount) (*cloudaccountpb.CloudAccount, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if err := c.fake.ValidateOrganization(cloudAccount.GetOrgId()); err != nil {
		return nil, err
	}
	if cloudAccount.CloudOrganizationId != nil {
		if err := c.fake.ValidateCloudOrganization(cloudAccount.GetCloudOrganizationId()); err != nil {
			return nil, err
		}
	}
	cloudAccount = clone(cloudAccount)
	cloudAccount.Id = uuid.New().String()
	cloudAccount.CreatedAt = timestamppb.Now()
	c.fake.CloudAccounts[cloudAccount.GetId()] = cloudAccount
	return clone(cloudAccount), nil
}

func (c *fakeCloudAccountClient) DestroyCloudAccount(ctx context.Context, cloudAccountID string) error {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.CloudAccounts[cloudAccountID]; !ok {
		return status.Error(codes.NotFound, "cloud account not found")
	}
	delete(c.fake.CloudAccounts, cloudAccountID)
	return nil
}

func (c *fakeCloudAccountClient) GetCloudAccount(ctx context.Context, cloudAccountID string) (*cloudaccountpb.CloudAccount, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	cloudAccount, ok := c.fake.CloudAccounts[cloudAccountID]
	if !ok {
		return nil, status.Error(codes.NotFound, "cloud account not found")
	}
	return clone(cloudAccount), nil
}

func (c *fakeCloudAccountClient) ListCloudAccounts(ctx context.Context, filter *cloudaccountsvcpb.ListCloudAccountsRequest_Filter) ([]*cloudaccountpb.CloudAccount, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	cloudAccounts := []*cloudaccountpb.CloudAccount{}
	for _, cloudAccount := range c.fake.CloudAccounts {
		if cloudAccountMatches(cloudAccount, filter) {
			cloudAccounts = append(cloudAccounts, clone(cloudAccount))
		}
	}
	return cloudAccounts, nil
}

func cloudAccountMatches(cloudAccount *cloudaccountpb.CloudAccount, filter *cloudaccountsvcpb.ListCloudAccountsRequest_Filter) bool {
	if cloudAccount.GetSuppressed() && !filter.GetIncludeSuppressed() {
		return false
	}
	if filter == nil {
		return true
	}
	if len(filter.GetOrgIds()) > 0 && !slices.Contains(filter.GetOrgIds(), cloudAccount.GetOrgId()) {
		return false
	}
	if len(filter.GetCloudOrganizationIds()) > 0 && !slices.Contains(filter.GetCloudOrganizationIds(), cloudAccount.GetCloudOrganizationId()) {
		return false
	}
	return true
}

func (c *fakeCloudAccountClient) UpdateCloudAccount(
	ctx context.Context,
	cloudAccount *cloudaccountpb.CloudAccount,
	updateMask *cloudaccountsvcpb.UpdateCloudAccountRequest_UpdateMask,
) (*cloudaccountpb.CloudAccount, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	existing, ok := c.fake.CloudAccounts[cloudAccount.GetId()]
	if !ok {
		return nil, status.Error(codes.NotFound, "cloud account not found")
	}

	var updated *cloudaccountpb.CloudAccount
	if updateMask == nil {
		updated = clone(cloudAccount)
	} else {
		updated = clone(existing)
		if updateMask.GetName() {
			updated.Name = cloudAccount.GetName()
		}
		if updateMask.GetSuppressed() {
			updated.Suppressed = cloudAccount.GetSuppressed()
		}
		if updateMask.GetManagedByDiscovery() {
			updated.ManagedByDiscovery = cloudAccount.GetManagedByDiscovery()
		}
		if updateMask.GetAwsLambdaDiscoveryConfig() || updateMask.GetAwsAgentCoreDiscoveryConfig() {
			aws := updated.GetAws()
			if aws == nil {
				aws = &cloudaccountpb.AWSAccount{}
				updated.Provider = &cloudaccountpb.CloudAccount_Aws{Aws: aws}
			}
			if updateMask.GetAwsLambdaDiscoveryConfig() {
				aws.LambdaDiscoveryConfig = cloudAccount.GetAws().GetLambdaDiscoveryConfig()
			}
			if updateMask.GetAwsAgentCoreDiscoveryConfig() {
				aws.AgentCoreDiscoveryConfig = cloudAccount.GetAws().GetAgentCoreDiscoveryConfig()
			}
		}
	}

	updated.LastUpdatedAt = timestamppb.Now()

	// existing may have been partially updated or entirely replaced, so explicitly set it to cover both cases
	c.fake.CloudAccounts[cloudAccount.GetId()] = updated
	return clone(updated), nil
}

func clone(cloudAccount *cloudaccountpb.CloudAccount) *cloudaccountpb.CloudAccount {
	return proto.Clone(cloudAccount).(*cloudaccountpb.CloudAccount)
}
