// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1beta1_test

import (
	"context"
	"testing"

	paginationpb "github.com/cofide/cofide-api-sdk/gen/go/cofide/common/pagination/v1beta1"
	organizationpb "github.com/cofide/cofide-api-sdk/gen/go/cofide/organization/v1beta1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/organization/v1beta1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_OrganizationClient_ImplementsMethods(t *testing.T) {
	test.AssertClientImplementsService(t, v1beta1.New(nil), organizationpb.OrganizationService_ServiceDesc)
}

// Test_OrganizationClient_Unimplemented tests v1beta1.OrganizationClient against an unimplemented server.
// This ensures that all errors returned are not wrapped and can be converted to a gRPC Status using Status.Convert.
func Test_OrganizationClient_Unimplemented(t *testing.T) {
	server := test.NewTestServer(t)
	organizationpb.RegisterOrganizationServiceServer(server.Server, &organizationpb.UnimplementedOrganizationServiceServer{})
	server.Serve()

	conn := server.CreateClientConn()
	client := v1beta1.New(conn)

	_, err := client.GetOrganization(t.Context(), "")
	test.RequireUnimplemented(t, err)

	_, err = client.GetOrganizationByName(t.Context(), "")
	test.RequireUnimplemented(t, err)

	_, _, err = client.ListOrganizations(t.Context(), nil, pagination.Pagination{})
	test.RequireUnimplemented(t, err)
}

func Test_GetOrganization(t *testing.T) {
	client, fakeOrganizationService := clientForTest(t)

	input := "test-org-id"
	fakeOrganizationService.expectedGetRequest = &organizationpb.GetOrganizationRequest{
		Identifier: &organizationpb.GetOrganizationRequest_OrgId{OrgId: input},
	}
	fakeOrganizationService.getResponse = &organizationpb.Organization{Id: input}

	resp, err := client.GetOrganization(t.Context(), input)
	require.NoError(t, err)
	assert.EqualExportedValues(t, fakeOrganizationService.getResponse, resp)
}

func Test_GetOrganizationByName(t *testing.T) {
	client, fakeOrganizationService := clientForTest(t)

	input := "test-name"
	fakeOrganizationService.expectedGetRequest = &organizationpb.GetOrganizationRequest{
		Identifier: &organizationpb.GetOrganizationRequest_Name{Name: input},
	}
	fakeOrganizationService.getResponse = &organizationpb.Organization{Id: "test-id", Name: input}

	resp, err := client.GetOrganizationByName(t.Context(), input)
	require.NoError(t, err)
	assert.EqualExportedValues(t, fakeOrganizationService.getResponse, resp)
}

func Test_ListOrganizations(t *testing.T) {
	client, fakeOrganizationService := clientForTest(t)

	filter := &organizationpb.ListOrganizationsRequest_Filter{}
	requestPagination := pagination.Pagination{PageSize: 100, Token: "test-token"}
	fakeOrganizationService.expectedListRequest = &organizationpb.ListOrganizationsRequest{
		Filter:     filter,
		Pagination: &paginationpb.PageRequest{PageSize: requestPagination.PageSize, PageToken: requestPagination.Token},
	}
	fakeOrganizationService.listResponse = []*organizationpb.Organization{
		{Id: "test-id-1", Name: "test-name-1"},
		{Id: "test-id-2", Name: "test-name-2"},
	}

	resp, responsePagination, err := client.ListOrganizations(t.Context(), filter, requestPagination)
	require.NoError(t, err)
	require.Len(t, resp, len(fakeOrganizationService.listResponse))
	for i, organization := range resp {
		assert.EqualExportedValues(t, fakeOrganizationService.listResponse[i], organization)
	}
	assert.Equal(t, pagination.Pagination{PageSize: 100, Token: "next-token"}, responsePagination)
}

func Test_ListOrganizations_NilFilter(t *testing.T) {
	client, fakeOrganizationService := clientForTest(t)

	fakeOrganizationService.expectedListRequest = &organizationpb.ListOrganizationsRequest{
		Pagination: &paginationpb.PageRequest{},
	}

	_, _, err := client.ListOrganizations(t.Context(), nil, pagination.Pagination{})
	require.NoError(t, err)
}

func clientForTest(t *testing.T) (v1beta1.OrganizationClient, *fakeOrganizationService) {
	t.Helper()
	server := test.NewTestServer(t)
	fakeOrganizationService := &fakeOrganizationService{t: t}
	organizationpb.RegisterOrganizationServiceServer(server.Server, fakeOrganizationService)
	server.Serve()

	conn := server.CreateClientConn()
	client := v1beta1.New(conn)

	return client, fakeOrganizationService
}

// fakeOrganizationService provides a very simple fake OrganizationService implementation with canned responses.
// Its purpose is to assert that the inbound request is as expected.
type fakeOrganizationService struct {
	organizationpb.UnimplementedOrganizationServiceServer
	t                   *testing.T
	expectedGetRequest  *organizationpb.GetOrganizationRequest
	expectedListRequest *organizationpb.ListOrganizationsRequest
	getResponse         *organizationpb.Organization
	listResponse        []*organizationpb.Organization
}

func (f *fakeOrganizationService) GetOrganization(ctx context.Context, req *organizationpb.GetOrganizationRequest) (*organizationpb.GetOrganizationResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedGetRequest, req)
	return &organizationpb.GetOrganizationResponse{Organization: f.getResponse}, nil
}

func (f *fakeOrganizationService) ListOrganizations(ctx context.Context, req *organizationpb.ListOrganizationsRequest) (*organizationpb.ListOrganizationsResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedListRequest, req)
	return &organizationpb.ListOrganizationsResponse{
		Organizations: f.listResponse,
		Pagination:    &paginationpb.PageResponse{NextPageToken: "next-token"},
	}, nil
}
