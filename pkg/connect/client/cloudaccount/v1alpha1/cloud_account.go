// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"

	cloudaccountpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_account/v1alpha1"
	"github.com/cofide/cofide-api-sdk/gen/go/proto/common/pagination/v1beta1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
	"google.golang.org/grpc"
)

// CloudAccountClient is an interface for a gRPC client for the v1alpha1 version of the Connect CloudAccountService.
type CloudAccountClient interface {
	CreateCloudAccount(ctx context.Context, cloudAccount *cloudaccountpb.CloudAccount) (*cloudaccountpb.CloudAccount, error)
	GetCloudAccount(ctx context.Context, cloudAccountID string) (*cloudaccountpb.CloudAccount, error)
	ListCloudAccounts(ctx context.Context, filter *cloudaccountpb.ListCloudAccountsRequest_Filter, pagination pagination.Pagination) ([]*cloudaccountpb.CloudAccount, error)
	UpdateCloudAccount(ctx context.Context, cloudAccount *cloudaccountpb.CloudAccount, updateMask *cloudaccountpb.UpdateCloudAccountRequest_UpdateMask) (*cloudaccountpb.CloudAccount, error)
	DeleteCloudAccount(ctx context.Context, cloudAccountID string) error
}

type exchangePolicyClient struct {
	client cloudaccountpb.CloudAccountServiceClient
}

// New instantiates a new CloudAccountClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) CloudAccountClient {
	return &exchangePolicyClient{
		client: cloudaccountpb.NewCloudAccountServiceClient(conn),
	}
}

func (c *exchangePolicyClient) CreateCloudAccount(ctx context.Context, cloudAccount *cloudaccountpb.CloudAccount) (*cloudaccountpb.CloudAccount, error) {
	resp, err := c.client.CreateCloudAccount(ctx, &cloudaccountpb.CreateCloudAccountRequest{
		CloudAccount: cloudAccount,
	})
	if err != nil {
		return nil, err
	}
	return resp.CloudAccount, nil
}

func (c *exchangePolicyClient) GetCloudAccount(ctx context.Context, cloudAccountID string) (*cloudaccountpb.CloudAccount, error) {
	resp, err := c.client.GetCloudAccount(ctx, &cloudaccountpb.GetCloudAccountRequest{
		CloudAccountId: cloudAccountID,
	})
	if err != nil {
		return nil, err
	}
	return resp.CloudAccount, nil
}

func (c *exchangePolicyClient) ListCloudAccounts(
	ctx context.Context,
	filter *cloudaccountpb.ListCloudAccountsRequest_Filter,
	pagination pagination.Pagination,
) ([]*cloudaccountpb.CloudAccount, error) {
	resp, err := c.client.ListCloudAccounts(ctx, &cloudaccountpb.ListCloudAccountsRequest{
		Filter: filter,
		Pagination: &v1beta1.PageRequest{
			PageSize:  pagination.PageSize,
			PageToken: pagination.Token,
		},
	})
	if err != nil {
		return nil, err
	}
	return resp.CloudAccounts, nil
}

func (c *exchangePolicyClient) UpdateCloudAccount(ctx context.Context, cloudAccount *cloudaccountpb.CloudAccount, updateMask *cloudaccountpb.UpdateCloudAccountRequest_UpdateMask) (*cloudaccountpb.CloudAccount, error) {
	resp, err := c.client.UpdateCloudAccount(ctx, &cloudaccountpb.UpdateCloudAccountRequest{
		CloudAccount: cloudAccount,
		UpdateMask:   updateMask,
	})
	if err != nil {
		return nil, err
	}
	return resp.CloudAccount, nil
}

func (c *exchangePolicyClient) DeleteCloudAccount(ctx context.Context, cloudAccountID string) error {
	_, err := c.client.DeleteCloudAccount(ctx, &cloudaccountpb.DeleteCloudAccountRequest{
		CloudAccountId: cloudAccountID,
	})
	return err
}
