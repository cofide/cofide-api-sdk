// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	cloudaccountpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_account/v1alpha1"
	cloudaccountv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/cloudaccount/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
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
		return nil, status.Error(codes.InvalidArgument, "organization not found")
	}

	if cloudAccount.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "name must be non-empty")
	}

	if cloudAccount.GetAws().GetAccountId() == "" {
		return nil, status.Error(codes.InvalidArgument, "aws account ID must be non-empty")
	}

	cloudAccount = clone(cloudAccount)
	cloudAccount.Id = uuid.New().String()
	cloudAccount.CreatedAt = timestamppb.Now()
	cloudAccount.LastUpdatedAt = nil
	c.fake.CloudAccounts[cloudAccount.GetId()] = cloudAccount
	return clone(cloudAccount), nil
}

func (c *fakeCloudAccountClient) DeleteCloudAccount(ctx context.Context, cloudAccountID string) error {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.CloudAccounts[cloudAccountID]; !ok {
		return status.Error(codes.NotFound, "cloud account not found")
	}
	delete(c.fake.CloudAccounts, cloudAccountID)
	return nil
}

func (c *fakeCloudAccountClient) GetCloudAccount(ctx context.Context, cloudAccountID string) (*cloudaccountpb.CloudAccount, error) {
	c.fake.Mu.RLock()
	defer c.fake.Mu.RUnlock()

	cloudAccount, ok := c.fake.CloudAccounts[cloudAccountID]
	if !ok {
		return nil, status.Error(codes.NotFound, "cloud account not found")
	}
	return clone(cloudAccount), nil
}

func (c *fakeCloudAccountClient) ListCloudAccounts(ctx context.Context, filter *cloudaccountpb.ListCloudAccountsRequest_Filter, requestPagination pagination.Pagination) ([]*cloudaccountpb.CloudAccount, pagination.Pagination, error) {
	c.fake.Mu.RLock()
	defer c.fake.Mu.RUnlock()

	// Filter
	cloudAccounts := []*cloudaccountpb.CloudAccount{}
	for cloudAccount := range fakeconnect.SortedValues(c.fake.CloudAccounts) {
		if cloudAccountMatches(cloudAccount, filter) {
			cloudAccounts = append(cloudAccounts, clone(cloudAccount))
		}
	}

	// Pagination (token is ID of first item that didn't fit on previous page)
	if requestPagination.Token != "" {
		for len(cloudAccounts) > 0 && cloudAccounts[0].GetId() != requestPagination.Token {
			if len(cloudAccounts) == 1 {
				cloudAccounts = []*cloudaccountpb.CloudAccount{}
			} else {
				cloudAccounts = cloudAccounts[1:]
			}
		}
	}
	var nextPageToken string
	if requestPagination.PageSize > 0 && len(cloudAccounts) > int(requestPagination.PageSize) {
		nextPageToken = cloudAccounts[requestPagination.PageSize].GetId()
		cloudAccounts = cloudAccounts[:requestPagination.PageSize]
	}

	return cloudAccounts, pagination.Pagination{PageSize: requestPagination.PageSize, Token: nextPageToken}, nil
}

func cloudAccountMatches(_ *cloudaccountpb.CloudAccount, filter *cloudaccountpb.ListCloudAccountsRequest_Filter) bool {
	if filter == nil {
		return true
	}
	// as filters are added, support them in the fake here
	return true
}

func (c *fakeCloudAccountClient) UpdateAWSCloudAccount(ctx context.Context, cloudAccount *cloudaccountpb.CloudAccount, updateMask *cloudaccountpb.UpdateCloudAccountRequest_AWSAccountUpdateMask) (*cloudaccountpb.CloudAccount, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	existing, ok := c.fake.CloudAccounts[cloudAccount.GetId()]
	if !ok {
		return nil, status.Error(codes.NotFound, "cloud account not found")
	}
	existing = clone(existing)

	if existing.GetAws() == nil {
		return nil, status.Error(codes.InvalidArgument, "cloud account is not an aws cloud account")
	}

	if updateMask == nil || updateMask.GetName() {
		if cloudAccount.GetName() == "" {
			return nil, status.Error(codes.InvalidArgument, "cloud account name must not be empty")
		}
		existing.Name = cloudAccount.GetName()
	}
	if updateMask == nil || updateMask.GetAwsAccountId() {
		if cloudAccount.GetAws().GetAccountId() == "" {
			return nil, status.Error(codes.InvalidArgument, "aws account id must not be empty")
		}
		existing.GetAws().AccountId = cloudAccount.GetAws().GetAccountId()
	}

	existing.LastUpdatedAt = timestamppb.Now()

	c.fake.CloudAccounts[cloudAccount.GetId()] = existing
	return clone(existing), nil
}

func clone(cloudAccount *cloudaccountpb.CloudAccount) *cloudaccountpb.CloudAccount {
	return proto.Clone(cloudAccount).(*cloudaccountpb.CloudAccount)
}
