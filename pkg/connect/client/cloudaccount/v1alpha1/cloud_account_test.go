// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1_test

import (
	"context"
	"testing"

	cloudaccountpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_account/v1alpha1"
	"github.com/cofide/cofide-api-sdk/gen/go/proto/common/pagination/v1beta1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/cloudaccount/v1alpha1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test_CloudAccountClient_Unimplemented tests v1alpha1.CloudAccountClient against an unimplemented server.
// This ensures that all errors returned are not wrapped can be converted to a gRPC Status using Status.Convert.
func Test_CloudAccountClient_Unimplemented(t *testing.T) {
	server := test.NewTestServer(t)
	cloudaccountpb.RegisterCloudAccountServiceServer(server.Server, &cloudaccountpb.UnimplementedCloudAccountServiceServer{})
	server.Serve()

	conn := server.CreateClientConn()
	client := v1alpha1.New(conn)

	_, err := client.CreateCloudAccount(t.Context(), nil)
	test.RequireUnimplemented(t, err)

	err = client.DeleteCloudAccount(t.Context(), "")
	test.RequireUnimplemented(t, err)

	_, err = client.GetCloudAccount(t.Context(), "")
	test.RequireUnimplemented(t, err)

	_, _, err = client.ListCloudAccounts(t.Context(), nil, pagination.Pagination{})
	test.RequireUnimplemented(t, err)

	_, err = client.UpdateCloudAccount(t.Context(), nil, nil)
	test.RequireUnimplemented(t, err)
}

func Test_CreateCloudAccount(t *testing.T) {
	client, fakeCloudAccountService := clientForTest(t)

	input := &cloudaccountpb.CloudAccount{
		OrgId: "test-org",
		Name:  "test-name",
		Provider: &cloudaccountpb.CloudAccount_Aws{
			Aws: &cloudaccountpb.AWSAccount{
				AccountId: "123456789012",
			},
		},
	}
	fakeCloudAccountService.expectedCreateRequest = &cloudaccountpb.CreateCloudAccountRequest{
		CloudAccount: input,
	}

	_, err := client.CreateCloudAccount(t.Context(), input)
	require.NoError(t, err)
}

func Test_DeleteCloudAccount(t *testing.T) {
	client, fakeCloudAccountService := clientForTest(t)

	input := "test-cloud-account-id"
	fakeCloudAccountService.expectedDeleteRequest = &cloudaccountpb.DeleteCloudAccountRequest{
		CloudAccountId: input,
	}

	err := client.DeleteCloudAccount(t.Context(), input)
	require.NoError(t, err)
}

func Test_GetCloudAccount(t *testing.T) {
	client, fakeCloudAccountService := clientForTest(t)

	input := "test-cloud-account-id"
	fakeCloudAccountService.expectedGetRequest = &cloudaccountpb.GetCloudAccountRequest{
		CloudAccountId: input,
	}

	_, err := client.GetCloudAccount(t.Context(), input)
	require.NoError(t, err)
}

func Test_ListCloudAccounts(t *testing.T) {
	client, fakeCloudAccountService := clientForTest(t)

	filter := &cloudaccountpb.ListCloudAccountsRequest_Filter{}
	requestPagination := pagination.Pagination{PageSize: 100}
	fakeCloudAccountService.expectedListRequest = &cloudaccountpb.ListCloudAccountsRequest{
		Filter:     filter,
		Pagination: &v1beta1.PageRequest{PageSize: requestPagination.PageSize, PageToken: requestPagination.Token},
	}

	_, _, err := client.ListCloudAccounts(t.Context(), filter, requestPagination)
	require.NoError(t, err)
}

func Test_UpdateCloudAccount(t *testing.T) {
	client, fakeCloudAccountService := clientForTest(t)

	input := &cloudaccountpb.CloudAccount{
		Id:   "test-id",
		Name: "test-name",
		Provider: &cloudaccountpb.CloudAccount_Aws{
			Aws: &cloudaccountpb.AWSAccount{
				AccountId: "123456789012",
			},
		},
	}
	fakeCloudAccountService.expectedUpdateRequest = &cloudaccountpb.UpdateCloudAccountRequest{
		CloudAccount: input,
		UpdateMask: &cloudaccountpb.UpdateCloudAccountRequest_AwsAccountUpdateMask{
			AwsAccountUpdateMask: &cloudaccountpb.UpdateCloudAccountRequest_AWSAccountUpdateMask{Name: true},
		},
	}

	_, err := client.UpdateCloudAccount(t.Context(), input, v1alpha1.AWSUpdateMask(true, false))
	require.NoError(t, err)
}

func Test_UpdateCloudAccount_NilUpdateMask(t *testing.T) {
	client, fakeCloudAccountService := clientForTest(t)

	input := &cloudaccountpb.CloudAccount{Id: "test-id", Name: "test-name"}
	// A nil mask must not be sent as an empty mask, which the server would treat as "update nothing".
	fakeCloudAccountService.expectedUpdateRequest = &cloudaccountpb.UpdateCloudAccountRequest{
		CloudAccount: input,
	}

	_, err := client.UpdateCloudAccount(t.Context(), input, nil)
	require.NoError(t, err)
}

func clientForTest(t *testing.T) (v1alpha1.CloudAccountClient, *fakeCloudAccountService) {
	t.Helper()
	server := test.NewTestServer(t)
	fakeCloudAccountService := &fakeCloudAccountService{t: t}
	cloudaccountpb.RegisterCloudAccountServiceServer(server.Server, fakeCloudAccountService)
	server.Serve()

	conn := server.CreateClientConn()
	client := v1alpha1.New(conn)

	return client, fakeCloudAccountService
}

// fakeCloudAccountService provides a very simple fake CloudAccountService implementation with canned responses.
// Its purpose is to assert that the inbound request is as expected.
type fakeCloudAccountService struct {
	t                     *testing.T
	expectedCreateRequest *cloudaccountpb.CreateCloudAccountRequest
	expectedDeleteRequest *cloudaccountpb.DeleteCloudAccountRequest
	expectedGetRequest    *cloudaccountpb.GetCloudAccountRequest
	expectedListRequest   *cloudaccountpb.ListCloudAccountsRequest
	expectedUpdateRequest *cloudaccountpb.UpdateCloudAccountRequest
}

func (f *fakeCloudAccountService) CreateCloudAccount(ctx context.Context, req *cloudaccountpb.CreateCloudAccountRequest) (*cloudaccountpb.CreateCloudAccountResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedCreateRequest, req)
	return &cloudaccountpb.CreateCloudAccountResponse{CloudAccount: req.CloudAccount}, nil
}

func (f *fakeCloudAccountService) DeleteCloudAccount(ctx context.Context, req *cloudaccountpb.DeleteCloudAccountRequest) (*cloudaccountpb.DeleteCloudAccountResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedDeleteRequest, req)
	return &cloudaccountpb.DeleteCloudAccountResponse{}, nil
}

func (f *fakeCloudAccountService) GetCloudAccount(ctx context.Context, req *cloudaccountpb.GetCloudAccountRequest) (*cloudaccountpb.GetCloudAccountResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedGetRequest, req)
	return &cloudaccountpb.GetCloudAccountResponse{CloudAccount: &cloudaccountpb.CloudAccount{Id: req.GetCloudAccountId()}}, nil
}

func (f *fakeCloudAccountService) ListCloudAccounts(ctx context.Context, req *cloudaccountpb.ListCloudAccountsRequest) (*cloudaccountpb.ListCloudAccountsResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedListRequest, req)
	return &cloudaccountpb.ListCloudAccountsResponse{CloudAccounts: []*cloudaccountpb.CloudAccount{}}, nil
}

func (f *fakeCloudAccountService) UpdateCloudAccount(ctx context.Context, req *cloudaccountpb.UpdateCloudAccountRequest) (*cloudaccountpb.UpdateCloudAccountResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedUpdateRequest, req)
	return &cloudaccountpb.UpdateCloudAccountResponse{CloudAccount: &cloudaccountpb.CloudAccount{Id: req.GetCloudAccount().GetId()}}, nil
}
