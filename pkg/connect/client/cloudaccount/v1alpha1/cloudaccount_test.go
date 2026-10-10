// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"
	"testing"

	cloudaccountsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/cloud_account_service/v1alpha1"
	cloudaccountpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_account/v1alpha1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCloudAccountClient_implementsMethods(t *testing.T) {
	test.AssertClientImplementsService(t, &cloudAccountClient{}, cloudaccountsvcpb.CloudAccountService_ServiceDesc)
}

// TestCloudAccountClient_Unimplemented tests CloudAccountClient against an unimplemented server.
// This ensures that all errors returned are not wrapped and can be converted to a gRPC Status using Status.Convert.
func TestCloudAccountClient_Unimplemented(t *testing.T) {
	server := test.NewTestServer(t)
	cloudaccountsvcpb.RegisterCloudAccountServiceServer(server.Server, &cloudaccountsvcpb.UnimplementedCloudAccountServiceServer{})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	cloudAccount, err := client.CreateCloudAccount(ctx, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, cloudAccount)

	err = client.DestroyCloudAccount(ctx, "")
	test.RequireUnimplemented(t, err)

	cloudAccount, err = client.GetCloudAccount(ctx, "")
	test.RequireUnimplemented(t, err)
	assert.Nil(t, cloudAccount)

	cloudAccounts, err := client.ListCloudAccounts(ctx, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, cloudAccounts)

	cloudAccount, err = client.UpdateCloudAccount(ctx, nil, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, cloudAccount)
}

func TestCloudAccountClient(t *testing.T) {
	server := test.NewTestServer(t)
	cloudaccountsvcpb.RegisterCloudAccountServiceServer(server.Server, &fakeCloudAccountService{t: t})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	cloudAccount := test.FakeCloudAccount()

	createdCloudAccount, err := client.CreateCloudAccount(ctx, cloudAccount)
	require.NoError(t, err)
	assert.EqualExportedValues(t, cloudAccount, createdCloudAccount)

	err = client.DestroyCloudAccount(ctx, cloudAccount.GetId())
	require.NoError(t, err)

	gotCloudAccount, err := client.GetCloudAccount(ctx, cloudAccount.GetId())
	require.NoError(t, err)
	assert.Equal(t, cloudAccount.GetId(), gotCloudAccount.GetId())

	filter := &cloudaccountsvcpb.ListCloudAccountsRequest_Filter{OrgIds: []string{test.FakeOrganizationID}}
	cloudAccounts, err := client.ListCloudAccounts(ctx, filter)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*cloudaccountpb.CloudAccount{test.FakeCloudAccount()}, cloudAccounts)

	updatedCloudAccount, err := client.UpdateCloudAccount(ctx, cloudAccount, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, cloudAccount, updatedCloudAccount)
}

// fakeCloudAccountService provides a simple fake CloudAccountService implementation with canned responses.
type fakeCloudAccountService struct {
	t *testing.T
}

func (f *fakeCloudAccountService) CreateCloudAccount(ctx context.Context, req *cloudaccountsvcpb.CreateCloudAccountRequest) (*cloudaccountsvcpb.CreateCloudAccountResponse, error) {
	assert.EqualExportedValues(f.t, test.FakeCloudAccount(), req.CloudAccount)
	return &cloudaccountsvcpb.CreateCloudAccountResponse{CloudAccount: req.CloudAccount}, nil
}

func (f *fakeCloudAccountService) DestroyCloudAccount(ctx context.Context, req *cloudaccountsvcpb.DestroyCloudAccountRequest) (*cloudaccountsvcpb.DestroyCloudAccountResponse, error) {
	assert.Equal(f.t, test.FakeCloudAccountID, req.GetCloudAccountId())
	return &cloudaccountsvcpb.DestroyCloudAccountResponse{}, nil
}

func (f *fakeCloudAccountService) GetCloudAccount(ctx context.Context, req *cloudaccountsvcpb.GetCloudAccountRequest) (*cloudaccountsvcpb.GetCloudAccountResponse, error) {
	assert.Equal(f.t, test.FakeCloudAccountID, req.GetCloudAccountId())
	return &cloudaccountsvcpb.GetCloudAccountResponse{CloudAccount: test.FakeCloudAccount()}, nil
}

func (f *fakeCloudAccountService) ListCloudAccounts(ctx context.Context, req *cloudaccountsvcpb.ListCloudAccountsRequest) (*cloudaccountsvcpb.ListCloudAccountsResponse, error) {
	assert.Equal(f.t, []string{test.FakeOrganizationID}, req.Filter.GetOrgIds())
	return &cloudaccountsvcpb.ListCloudAccountsResponse{CloudAccounts: []*cloudaccountpb.CloudAccount{test.FakeCloudAccount()}}, nil
}

func (f *fakeCloudAccountService) UpdateCloudAccount(ctx context.Context, req *cloudaccountsvcpb.UpdateCloudAccountRequest) (*cloudaccountsvcpb.UpdateCloudAccountResponse, error) {
	assert.EqualExportedValues(f.t, test.FakeCloudAccount(), req.CloudAccount)
	return &cloudaccountsvcpb.UpdateCloudAccountResponse{CloudAccount: test.FakeCloudAccount()}, nil
}
