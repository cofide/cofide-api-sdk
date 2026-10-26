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
	DeleteCloudAccount(ctx context.Context, cloudAccountID string) error
	GetCloudAccount(ctx context.Context, cloudAccountID string) (*cloudaccountpb.CloudAccount, error)
	ListCloudAccounts(ctx context.Context, filter *cloudaccountpb.ListCloudAccountsRequest_Filter, pagination pagination.Pagination) ([]*cloudaccountpb.CloudAccount, pagination.Pagination, error)
	UpdateAWSCloudAccount(ctx context.Context, cloudAccount *cloudaccountpb.CloudAccount, updateMask *cloudaccountpb.UpdateCloudAccountRequest_AWSAccountUpdateMask) (*cloudaccountpb.CloudAccount, error)
}

type cloudAccountClient struct {
	client cloudaccountpb.CloudAccountServiceClient
}

// New instantiates a new CloudAccountClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) CloudAccountClient {
	return &cloudAccountClient{
		client: cloudaccountpb.NewCloudAccountServiceClient(conn),
	}
}

func (c *cloudAccountClient) CreateCloudAccount(ctx context.Context, cloudAccount *cloudaccountpb.CloudAccount) (*cloudaccountpb.CloudAccount, error) {
	resp, err := c.client.CreateCloudAccount(ctx, &cloudaccountpb.CreateCloudAccountRequest{
		CloudAccount: cloudAccount,
	})
	if err != nil {
		return nil, err
	}
	return resp.GetCloudAccount(), nil
}

func (c *cloudAccountClient) DeleteCloudAccount(ctx context.Context, cloudAccountID string) error {
	_, err := c.client.DeleteCloudAccount(ctx, &cloudaccountpb.DeleteCloudAccountRequest{
		CloudAccountId: cloudAccountID,
	})
	return err
}

func (c *cloudAccountClient) GetCloudAccount(ctx context.Context, cloudAccountID string) (*cloudaccountpb.CloudAccount, error) {
	resp, err := c.client.GetCloudAccount(ctx, &cloudaccountpb.GetCloudAccountRequest{
		CloudAccountId: cloudAccountID,
	})
	if err != nil {
		return nil, err
	}
	return resp.GetCloudAccount(), nil
}

func (c *cloudAccountClient) ListCloudAccounts(ctx context.Context, filter *cloudaccountpb.ListCloudAccountsRequest_Filter, requestPagination pagination.Pagination) ([]*cloudaccountpb.CloudAccount, pagination.Pagination, error) {
	resp, err := c.client.ListCloudAccounts(ctx, &cloudaccountpb.ListCloudAccountsRequest{
		Filter:     filter,
		Pagination: &v1beta1.PageRequest{PageSize: requestPagination.PageSize, PageToken: requestPagination.Token},
	})
	if err != nil {
		return nil, pagination.Pagination{}, err
	}
	return resp.GetCloudAccounts(), pagination.Pagination{PageSize: requestPagination.PageSize, Token: resp.GetPagination().GetNextPageToken()}, nil
}

func (c *cloudAccountClient) UpdateAWSCloudAccount(ctx context.Context, cloudAccount *cloudaccountpb.CloudAccount, updateMask *cloudaccountpb.UpdateCloudAccountRequest_AWSAccountUpdateMask) (*cloudaccountpb.CloudAccount, error) {
	req := &cloudaccountpb.UpdateCloudAccountRequest{CloudAccount: cloudAccount}
	// Only set the oneof when a mask is provided. A wrapper around a nil mask is encoded as an empty (present) mask,
	// which the server would treat as "update no fields" rather than a full replacement.
	if updateMask != nil {
		req.UpdateMask = &cloudaccountpb.UpdateCloudAccountRequest_AwsAccountUpdateMask{
			AwsAccountUpdateMask: updateMask,
		}
	}
	resp, err := c.client.UpdateCloudAccount(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.GetCloudAccount(), nil
}
