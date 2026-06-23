// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	cloudaccountpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_account/v1alpha1"
	cloudaccountsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/cloud_account_service/v1alpha1"
	cloudaccountv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/cloudaccount/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
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

	id := uuid.New().String()
	cloudAccount = clone(cloudAccount)
	cloudAccount.Id = id
	c.fake.CloudAccounts[cloudAccount.GetId()] = cloudAccount
	return clone(cloudAccount), nil
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
	if filter == nil {
		return true
	}
	if filter.OrgId != "" && cloudAccount.GetOrgId() != filter.OrgId {
		return false
	}
	if filter.CloudOrganizationId != "" && cloudAccount.GetCloudOrganizationId() != filter.CloudOrganizationId {
		return false
	}
	if !filter.IncludeSuppressed && cloudAccount.GetSuppressed() {
		return false
	}
	return true
}

func (c *fakeCloudAccountClient) UpdateCloudAccount(ctx context.Context, cloudAccount *cloudaccountpb.CloudAccount, updateMask *cloudaccountsvcpb.UpdateCloudAccountRequest_UpdateMask) (*cloudaccountpb.CloudAccount, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.CloudAccounts[cloudAccount.GetId()]; !ok {
		return nil, status.Error(codes.NotFound, "cloud account not found")
	}
	c.fake.CloudAccounts[cloudAccount.GetId()] = clone(cloudAccount)
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

func clone(cloudAccount *cloudaccountpb.CloudAccount) *cloudaccountpb.CloudAccount {
	return proto.Clone(cloudAccount).(*cloudaccountpb.CloudAccount)
}
