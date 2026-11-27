// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1beta1_test

import (
	"context"
	"testing"

	paginationpb "github.com/cofide/cofide-api-sdk/gen/go/cofide/common/pagination/v1beta1"
	federationpb "github.com/cofide/cofide-api-sdk/gen/go/cofide/federation/v1beta1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/federation/v1beta1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_FederationClient_ImplementsMethods(t *testing.T) {
	test.AssertClientImplementsService(t, v1beta1.New(nil), federationpb.FederationService_ServiceDesc)
}

// Test_FederationClient_Unimplemented tests v1beta1.FederationClient against an unimplemented server.
// This ensures that all errors returned are not wrapped and can be converted to a gRPC Status using Status.Convert.
func Test_FederationClient_Unimplemented(t *testing.T) {
	server := test.NewTestServer(t)
	federationpb.RegisterFederationServiceServer(server.Server, &federationpb.UnimplementedFederationServiceServer{})
	server.Serve()

	conn := server.CreateClientConn()
	client := v1beta1.New(conn)

	_, err := client.CreateFederation(t.Context(), nil)
	test.RequireUnimplemented(t, err)

	err = client.DeleteFederation(t.Context(), "")
	test.RequireUnimplemented(t, err)

	_, err = client.GetFederation(t.Context(), "")
	test.RequireUnimplemented(t, err)

	_, err = client.GetFederationByTrustZones(t.Context(), "", "")
	test.RequireUnimplemented(t, err)

	_, _, err = client.ListFederations(t.Context(), nil, pagination.Pagination{})
	test.RequireUnimplemented(t, err)
}

func Test_CreateFederation(t *testing.T) {
	client, fakeFederationService := clientForTest(t)

	input := &federationpb.Federation{
		TrustZoneId: "test-trust-zone-id",
		Remote:      &federationpb.Federation_RemoteTrustZoneId{RemoteTrustZoneId: "test-remote-trust-zone-id"},
	}
	fakeFederationService.expectedCreateRequest = &federationpb.CreateFederationRequest{
		Federation: input,
	}
	fakeFederationService.createResponse = &federationpb.Federation{
		Id:          "test-id",
		OrgId:       "test-org",
		TrustZoneId: "test-trust-zone-id",
		Remote:      &federationpb.Federation_RemoteTrustZoneId{RemoteTrustZoneId: "test-remote-trust-zone-id"},
	}

	resp, err := client.CreateFederation(t.Context(), input)
	require.NoError(t, err)
	assert.EqualExportedValues(t, fakeFederationService.createResponse, resp)
}

func Test_DeleteFederation(t *testing.T) {
	client, fakeFederationService := clientForTest(t)

	input := "test-federation-id"
	fakeFederationService.expectedDeleteRequest = &federationpb.DeleteFederationRequest{
		FederationId: input,
	}

	err := client.DeleteFederation(t.Context(), input)
	require.NoError(t, err)
}

func Test_GetFederation(t *testing.T) {
	client, fakeFederationService := clientForTest(t)

	input := "test-federation-id"
	fakeFederationService.expectedGetRequest = &federationpb.GetFederationRequest{
		Identifier: &federationpb.GetFederationRequest_FederationId{FederationId: input},
	}
	fakeFederationService.getResponse = &federationpb.Federation{Id: input}

	resp, err := client.GetFederation(t.Context(), input)
	require.NoError(t, err)
	assert.EqualExportedValues(t, fakeFederationService.getResponse, resp)
}

func Test_GetFederationByTrustZones(t *testing.T) {
	client, fakeFederationService := clientForTest(t)

	fakeFederationService.expectedGetRequest = &federationpb.GetFederationRequest{
		Identifier: &federationpb.GetFederationRequest_TrustZones{
			TrustZones: &federationpb.GetFederationRequest_TrustZonePair{
				TrustZoneId:       "test-trust-zone-id",
				RemoteTrustZoneId: "test-remote-trust-zone-id",
			},
		},
	}
	fakeFederationService.getResponse = &federationpb.Federation{
		Id:          "test-id",
		TrustZoneId: "test-trust-zone-id",
		Remote:      &federationpb.Federation_RemoteTrustZoneId{RemoteTrustZoneId: "test-remote-trust-zone-id"},
	}

	resp, err := client.GetFederationByTrustZones(t.Context(), "test-trust-zone-id", "test-remote-trust-zone-id")
	require.NoError(t, err)
	assert.EqualExportedValues(t, fakeFederationService.getResponse, resp)
}

func Test_ListFederations(t *testing.T) {
	client, fakeFederationService := clientForTest(t)

	filter := &federationpb.ListFederationsRequest_Filter{OrgIds: []string{"test-org"}}
	requestPagination := pagination.Pagination{PageSize: 100, Token: "test-token"}
	fakeFederationService.expectedListRequest = &federationpb.ListFederationsRequest{
		Filter:     filter,
		Pagination: &paginationpb.PageRequest{PageSize: requestPagination.PageSize, PageToken: requestPagination.Token},
	}
	fakeFederationService.listResponse = []*federationpb.Federation{
		{Id: "test-id-1", OrgId: "test-org"},
		{Id: "test-id-2", OrgId: "test-org"},
	}

	resp, responsePagination, err := client.ListFederations(t.Context(), filter, requestPagination)
	require.NoError(t, err)
	require.Len(t, resp, len(fakeFederationService.listResponse))
	for i, federation := range resp {
		assert.EqualExportedValues(t, fakeFederationService.listResponse[i], federation)
	}
	assert.Equal(t, pagination.Pagination{PageSize: 100, Token: "next-token"}, responsePagination)
}

func Test_ListFederations_NilFilter(t *testing.T) {
	client, fakeFederationService := clientForTest(t)

	fakeFederationService.expectedListRequest = &federationpb.ListFederationsRequest{
		Pagination: &paginationpb.PageRequest{},
	}

	_, _, err := client.ListFederations(t.Context(), nil, pagination.Pagination{})
	require.NoError(t, err)
}

func clientForTest(t *testing.T) (v1beta1.FederationClient, *fakeFederationService) {
	t.Helper()
	server := test.NewTestServer(t)
	fakeFederationService := &fakeFederationService{t: t}
	federationpb.RegisterFederationServiceServer(server.Server, fakeFederationService)
	server.Serve()

	conn := server.CreateClientConn()
	client := v1beta1.New(conn)

	return client, fakeFederationService
}

// fakeFederationService provides a very simple fake FederationService implementation with canned responses.
// Its purpose is to assert that the inbound request is as expected.
type fakeFederationService struct {
	federationpb.UnimplementedFederationServiceServer
	t                     *testing.T
	expectedCreateRequest *federationpb.CreateFederationRequest
	expectedDeleteRequest *federationpb.DeleteFederationRequest
	expectedGetRequest    *federationpb.GetFederationRequest
	expectedListRequest   *federationpb.ListFederationsRequest
	createResponse        *federationpb.Federation
	getResponse           *federationpb.Federation
	listResponse          []*federationpb.Federation
}

func (f *fakeFederationService) CreateFederation(ctx context.Context, req *federationpb.CreateFederationRequest) (*federationpb.CreateFederationResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedCreateRequest, req)
	return &federationpb.CreateFederationResponse{Federation: f.createResponse}, nil
}

func (f *fakeFederationService) DeleteFederation(ctx context.Context, req *federationpb.DeleteFederationRequest) (*federationpb.DeleteFederationResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedDeleteRequest, req)
	return &federationpb.DeleteFederationResponse{}, nil
}

func (f *fakeFederationService) GetFederation(ctx context.Context, req *federationpb.GetFederationRequest) (*federationpb.GetFederationResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedGetRequest, req)
	return &federationpb.GetFederationResponse{Federation: f.getResponse}, nil
}

func (f *fakeFederationService) ListFederations(ctx context.Context, req *federationpb.ListFederationsRequest) (*federationpb.ListFederationsResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedListRequest, req)
	return &federationpb.ListFederationsResponse{
		Federations: f.listResponse,
		Pagination:  &paginationpb.PageResponse{NextPageToken: "next-token"},
	}, nil
}
