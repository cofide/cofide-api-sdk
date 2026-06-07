// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"

	cloudaccountpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_account/v1alpha1"
	cloudaccountsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/cloud_account_service/v1alpha1"
	"google.golang.org/grpc"
)

// CloudAccountClient is an interface for a gRPC client for the v1alpha1 version of the Connect CloudAccountService.
type CloudAccountClient interface {
	CreateCloudAccount(ctx context.Context, cloudAccount *cloudaccountpb.CloudAccount) (*cloudaccountpb.CloudAccount, error)
	GetCloudAccount(ctx context.Context, cloudAccountID string) (*cloudaccountpb.CloudAccount, error)
	ListCloudAccounts(ctx context.Context, filter *cloudaccountsvcpb.ListCloudAccountsRequest_Filter) ([]*cloudaccountpb.CloudAccount, error)
	UpdateCloudAccount(ctx context.Context, cloudAccount *cloudaccountpb.CloudAccount, updateMask *cloudaccountsvcpb.UpdateCloudAccountRequest_UpdateMask) (*cloudaccountpb.CloudAccount, error)
	DestroyCloudAccount(ctx context.Context, cloudAccountID string) error
}

type cloudAccountClient struct {
	client cloudaccountsvcpb.CloudAccountServiceClient
}

// New instantiates a new CloudAccountClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) CloudAccountClient {
	return &cloudAccountClient{
		client: cloudaccountsvcpb.NewCloudAccountServiceClient(conn),
	}
}

func (c *cloudAccountClient) CreateCloudAccount(ctx context.Context, cloudAccount *cloudaccountpb.CloudAccount) (*cloudaccountpb.CloudAccount, error) {
	resp, err := c.client.CreateCloudAccount(ctx, &cloudaccountsvcpb.CreateCloudAccountRequest{
		CloudAccount: cloudAccount,
	})
	if err != nil {
		return nil, err
	}
	return resp.CloudAccount, nil
}

func (c *cloudAccountClient) GetCloudAccount(ctx context.Context, cloudAccountID string) (*cloudaccountpb.CloudAccount, error) {
	resp, err := c.client.GetCloudAccount(ctx, &cloudaccountsvcpb.GetCloudAccountRequest{
		CloudAccountId: cloudAccountID,
	})
	if err != nil {
		return nil, err
	}
	return resp.CloudAccount, nil
}

func (c *cloudAccountClient) ListCloudAccounts(ctx context.Context, filter *cloudaccountsvcpb.ListCloudAccountsRequest_Filter) ([]*cloudaccountpb.CloudAccount, error) {
	resp, err := c.client.ListCloudAccounts(ctx, &cloudaccountsvcpb.ListCloudAccountsRequest{
		Filter: filter,
	})
	if err != nil {
		return nil, err
	}
	return resp.CloudAccounts, nil
}

func (c *cloudAccountClient) UpdateCloudAccount(ctx context.Context, cloudAccount *cloudaccountpb.CloudAccount, updateMask *cloudaccountsvcpb.UpdateCloudAccountRequest_UpdateMask) (*cloudaccountpb.CloudAccount, error) {
	resp, err := c.client.UpdateCloudAccount(ctx, &cloudaccountsvcpb.UpdateCloudAccountRequest{
		CloudAccount: cloudAccount,
		UpdateMask:   updateMask,
	})
	if err != nil {
		return nil, err
	}
	return resp.CloudAccount, nil
}

func (c *cloudAccountClient) DestroyCloudAccount(ctx context.Context, cloudAccountID string) error {
	_, err := c.client.DestroyCloudAccount(ctx, &cloudaccountsvcpb.DestroyCloudAccountRequest{
		CloudAccountId: cloudAccountID,
	})
	return err
}
