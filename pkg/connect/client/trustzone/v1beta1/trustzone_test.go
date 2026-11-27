// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1beta1_test

import (
	"context"
	"testing"

	paginationpb "github.com/cofide/cofide-api-sdk/gen/go/cofide/common/pagination/v1beta1"
	trustbundlepb "github.com/cofide/cofide-api-sdk/gen/go/cofide/common/trust_bundle/v1beta1"
	trustzonepb "github.com/cofide/cofide-api-sdk/gen/go/cofide/trust_zone/v1beta1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/trustzone/v1beta1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_TrustZoneClient_ImplementsMethods(t *testing.T) {
	test.AssertClientImplementsService(t, v1beta1.New(nil), trustzonepb.TrustZoneService_ServiceDesc)
}

// Test_TrustZoneClient_Unimplemented tests v1beta1.TrustZoneClient against an unimplemented server.
// This ensures that all errors returned are not wrapped and can be converted to a gRPC Status using Status.Convert.
func Test_TrustZoneClient_Unimplemented(t *testing.T) {
	server := test.NewTestServer(t)
	trustzonepb.RegisterTrustZoneServiceServer(server.Server, &trustzonepb.UnimplementedTrustZoneServiceServer{})
	server.Serve()

	conn := server.CreateClientConn()
	client := v1beta1.New(conn)

	_, err := client.CreateTrustZone(t.Context(), nil)
	test.RequireUnimplemented(t, err)

	err = client.DeleteTrustZone(t.Context(), "")
	test.RequireUnimplemented(t, err)

	_, err = client.GetTrustZone(t.Context(), "")
	test.RequireUnimplemented(t, err)

	_, err = client.GetTrustZoneByName(t.Context(), "", "")
	test.RequireUnimplemented(t, err)

	_, err = client.GetTrustZoneByTrustDomain(t.Context(), "")
	test.RequireUnimplemented(t, err)

	_, _, err = client.ListTrustZones(t.Context(), nil, pagination.Pagination{})
	test.RequireUnimplemented(t, err)

	_, err = client.UpdateTrustZone(t.Context(), nil, nil)
	test.RequireUnimplemented(t, err)

	err = client.PublishTrustZoneBundle(t.Context(), nil)
	test.RequireUnimplemented(t, err)

	err = client.RegisterCofideAgent(t.Context(), v1beta1.JoinToken("token"), nil)
	test.RequireUnimplemented(t, err)

	err = client.RegisterTrustZoneServer(t.Context(), v1beta1.JoinToken("token"), nil)
	test.RequireUnimplemented(t, err)
}

func Test_CreateTrustZone(t *testing.T) {
	client, fakeTrustZoneService := clientForTest(t)

	input := &trustzonepb.TrustZone{
		OrgId:       "test-org",
		Name:        "test-name",
		TrustDomain: "example.org",
	}
	fakeTrustZoneService.expectedCreateRequest = &trustzonepb.CreateTrustZoneRequest{
		TrustZone: input,
	}
	fakeTrustZoneService.createResponse = &trustzonepb.TrustZone{
		Id:          "test-id",
		OrgId:       "test-org",
		Name:        "test-name",
		TrustDomain: "example.org",
	}

	resp, err := client.CreateTrustZone(t.Context(), input)
	require.NoError(t, err)
	assert.EqualExportedValues(t, fakeTrustZoneService.createResponse, resp)
}

func Test_DeleteTrustZone(t *testing.T) {
	client, fakeTrustZoneService := clientForTest(t)

	input := "test-trust-zone-id"
	fakeTrustZoneService.expectedDeleteRequest = &trustzonepb.DeleteTrustZoneRequest{
		TrustZoneId: input,
	}

	err := client.DeleteTrustZone(t.Context(), input)
	require.NoError(t, err)
}

func Test_GetTrustZone(t *testing.T) {
	client, fakeTrustZoneService := clientForTest(t)

	input := "test-trust-zone-id"
	fakeTrustZoneService.expectedGetRequest = &trustzonepb.GetTrustZoneRequest{
		Identifier: &trustzonepb.GetTrustZoneRequest_TrustZoneId{TrustZoneId: input},
	}
	fakeTrustZoneService.getResponse = &trustzonepb.TrustZone{Id: input}

	resp, err := client.GetTrustZone(t.Context(), input)
	require.NoError(t, err)
	assert.EqualExportedValues(t, fakeTrustZoneService.getResponse, resp)
}

func Test_GetTrustZoneByName(t *testing.T) {
	client, fakeTrustZoneService := clientForTest(t)

	fakeTrustZoneService.expectedGetRequest = &trustzonepb.GetTrustZoneRequest{
		Identifier: &trustzonepb.GetTrustZoneRequest_NameWithOrg{
			NameWithOrg: &trustzonepb.GetTrustZoneRequest_NameWithOrgId{OrgId: "test-org", Name: "test-name"},
		},
	}
	fakeTrustZoneService.getResponse = &trustzonepb.TrustZone{Id: "test-id", OrgId: "test-org", Name: "test-name"}

	resp, err := client.GetTrustZoneByName(t.Context(), "test-org", "test-name")
	require.NoError(t, err)
	assert.EqualExportedValues(t, fakeTrustZoneService.getResponse, resp)
}

func Test_GetTrustZoneByTrustDomain(t *testing.T) {
	client, fakeTrustZoneService := clientForTest(t)

	input := "example.org"
	fakeTrustZoneService.expectedGetRequest = &trustzonepb.GetTrustZoneRequest{
		Identifier: &trustzonepb.GetTrustZoneRequest_TrustDomain{TrustDomain: input},
	}
	fakeTrustZoneService.getResponse = &trustzonepb.TrustZone{Id: "test-id", TrustDomain: input}

	resp, err := client.GetTrustZoneByTrustDomain(t.Context(), input)
	require.NoError(t, err)
	assert.EqualExportedValues(t, fakeTrustZoneService.getResponse, resp)
}

func Test_ListTrustZones(t *testing.T) {
	client, fakeTrustZoneService := clientForTest(t)

	filter := &trustzonepb.ListTrustZonesRequest_Filter{OrgIds: []string{"test-org"}}
	requestPagination := pagination.Pagination{PageSize: 100, Token: "test-token"}
	fakeTrustZoneService.expectedListRequest = &trustzonepb.ListTrustZonesRequest{
		Filter:     filter,
		Pagination: &paginationpb.PageRequest{PageSize: requestPagination.PageSize, PageToken: requestPagination.Token},
	}
	fakeTrustZoneService.listResponse = []*trustzonepb.TrustZone{
		{Id: "test-id-1", OrgId: "test-org"},
		{Id: "test-id-2", OrgId: "test-org"},
	}

	resp, responsePagination, err := client.ListTrustZones(t.Context(), filter, requestPagination)
	require.NoError(t, err)
	require.Len(t, resp, len(fakeTrustZoneService.listResponse))
	for i, trustZone := range resp {
		assert.EqualExportedValues(t, fakeTrustZoneService.listResponse[i], trustZone)
	}
	assert.Equal(t, pagination.Pagination{PageSize: 100, Token: "next-token"}, responsePagination)
}

func Test_ListTrustZones_NilFilter(t *testing.T) {
	client, fakeTrustZoneService := clientForTest(t)

	fakeTrustZoneService.expectedListRequest = &trustzonepb.ListTrustZonesRequest{
		Pagination: &paginationpb.PageRequest{},
	}

	_, _, err := client.ListTrustZones(t.Context(), nil, pagination.Pagination{})
	require.NoError(t, err)
}

func Test_UpdateTrustZone(t *testing.T) {
	client, fakeTrustZoneService := clientForTest(t)

	input := &trustzonepb.TrustZone{Id: "test-id", Name: "test-name"}
	updateMask := &trustzonepb.TrustZoneUpdateMask{Name: true}
	fakeTrustZoneService.expectedUpdateRequest = &trustzonepb.UpdateTrustZoneRequest{
		TrustZone:  input,
		UpdateMask: updateMask,
	}
	fakeTrustZoneService.updateResponse = &trustzonepb.TrustZone{Id: "test-id", OrgId: "test-org", Name: "test-name"}

	resp, err := client.UpdateTrustZone(t.Context(), input, updateMask)
	require.NoError(t, err)
	assert.EqualExportedValues(t, fakeTrustZoneService.updateResponse, resp)
}

func Test_UpdateTrustZone_NilUpdateMask(t *testing.T) {
	client, fakeTrustZoneService := clientForTest(t)

	input := &trustzonepb.TrustZone{Id: "test-id", Name: "test-name"}
	// A nil mask must not be sent as an empty mask, which the server would treat as "update nothing".
	fakeTrustZoneService.expectedUpdateRequest = &trustzonepb.UpdateTrustZoneRequest{
		TrustZone: input,
	}
	fakeTrustZoneService.updateResponse = &trustzonepb.TrustZone{Id: "test-id", OrgId: "test-org", Name: "test-name"}

	resp, err := client.UpdateTrustZone(t.Context(), input, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, fakeTrustZoneService.updateResponse, resp)
}

func Test_PublishTrustZoneBundle(t *testing.T) {
	client, fakeTrustZoneService := clientForTest(t)

	bundle := &trustbundlepb.Bundle{TrustDomain: "example.org"}
	fakeTrustZoneService.expectedPublishBundleRequest = &trustzonepb.PublishTrustZoneBundleRequest{
		Bundle: bundle,
	}

	err := client.PublishTrustZoneBundle(t.Context(), bundle)
	require.NoError(t, err)
}

func Test_RegisterCofideAgent(t *testing.T) {
	bundle := &trustbundlepb.Bundle{TrustDomain: "example.org"}
	tests := []struct {
		name       string
		credential v1beta1.CofideAgentCredential
		want       *trustzonepb.RegisterCofideAgentRequest
	}{
		{
			name:       "join token",
			credential: v1beta1.JoinToken("test-token"),
			want: &trustzonepb.RegisterCofideAgentRequest{
				Auth:   &trustzonepb.RegisterCofideAgentRequest_JoinToken{JoinToken: "test-token"},
				Bundle: bundle,
			},
		},
		{
			name:       "k8s psat",
			credential: v1beta1.K8sPSAT("test-psat"),
			want: &trustzonepb.RegisterCofideAgentRequest{
				Auth:   &trustzonepb.RegisterCofideAgentRequest_K8SPsat{K8SPsat: "test-psat"},
				Bundle: bundle,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, fakeTrustZoneService := clientForTest(t)
			fakeTrustZoneService.expectedRegisterCofideAgentRequest = tt.want

			err := client.RegisterCofideAgent(t.Context(), tt.credential, bundle)
			require.NoError(t, err)
		})
	}
}

func Test_RegisterTrustZoneServer(t *testing.T) {
	bundle := &trustbundlepb.Bundle{TrustDomain: "example.org"}
	tests := []struct {
		name       string
		credential v1beta1.TrustZoneServerCredential
		want       *trustzonepb.RegisterTrustZoneServerRequest
	}{
		{
			name:       "join token",
			credential: v1beta1.JoinToken("test-token"),
			want: &trustzonepb.RegisterTrustZoneServerRequest{
				Auth:   &trustzonepb.RegisterTrustZoneServerRequest_JoinToken{JoinToken: "test-token"},
				Bundle: bundle,
			},
		},
		{
			name:       "k8s psat",
			credential: v1beta1.K8sPSAT("test-psat"),
			want: &trustzonepb.RegisterTrustZoneServerRequest{
				Auth:   &trustzonepb.RegisterTrustZoneServerRequest_K8SPsat{K8SPsat: "test-psat"},
				Bundle: bundle,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, fakeTrustZoneService := clientForTest(t)
			fakeTrustZoneService.expectedRegisterTrustZoneServerRequest = tt.want

			err := client.RegisterTrustZoneServer(t.Context(), tt.credential, bundle)
			require.NoError(t, err)
		})
	}
}

func Test_Register_NilCredential(t *testing.T) {
	// The fake service fails the test if it receives a request, since it expects none.
	client, _ := clientForTest(t)
	bundle := &trustbundlepb.Bundle{TrustDomain: "example.org"}

	err := client.RegisterCofideAgent(t.Context(), nil, bundle)
	require.Error(t, err)
	assert.Equal(t, codes.InvalidArgument, status.Code(err))

	err = client.RegisterTrustZoneServer(t.Context(), nil, bundle)
	require.Error(t, err)
	assert.Equal(t, codes.InvalidArgument, status.Code(err))
}

func clientForTest(t *testing.T) (v1beta1.TrustZoneClient, *fakeTrustZoneService) {
	t.Helper()
	server := test.NewTestServer(t)
	fakeTrustZoneService := &fakeTrustZoneService{t: t}
	trustzonepb.RegisterTrustZoneServiceServer(server.Server, fakeTrustZoneService)
	server.Serve()

	conn := server.CreateClientConn()
	client := v1beta1.New(conn)

	return client, fakeTrustZoneService
}

// fakeTrustZoneService provides a very simple fake TrustZoneService implementation with canned responses.
// Its purpose is to assert that the inbound request is as expected.
type fakeTrustZoneService struct {
	trustzonepb.UnimplementedTrustZoneServiceServer
	t                                      *testing.T
	expectedCreateRequest                  *trustzonepb.CreateTrustZoneRequest
	expectedDeleteRequest                  *trustzonepb.DeleteTrustZoneRequest
	expectedGetRequest                     *trustzonepb.GetTrustZoneRequest
	expectedListRequest                    *trustzonepb.ListTrustZonesRequest
	expectedUpdateRequest                  *trustzonepb.UpdateTrustZoneRequest
	expectedPublishBundleRequest           *trustzonepb.PublishTrustZoneBundleRequest
	expectedRegisterCofideAgentRequest     *trustzonepb.RegisterCofideAgentRequest
	expectedRegisterTrustZoneServerRequest *trustzonepb.RegisterTrustZoneServerRequest
	createResponse                         *trustzonepb.TrustZone
	getResponse                            *trustzonepb.TrustZone
	listResponse                           []*trustzonepb.TrustZone
	updateResponse                         *trustzonepb.TrustZone
}

func (f *fakeTrustZoneService) CreateTrustZone(ctx context.Context, req *trustzonepb.CreateTrustZoneRequest) (*trustzonepb.CreateTrustZoneResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedCreateRequest, req)
	return &trustzonepb.CreateTrustZoneResponse{TrustZone: f.createResponse}, nil
}

func (f *fakeTrustZoneService) DeleteTrustZone(ctx context.Context, req *trustzonepb.DeleteTrustZoneRequest) (*trustzonepb.DeleteTrustZoneResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedDeleteRequest, req)
	return &trustzonepb.DeleteTrustZoneResponse{}, nil
}

func (f *fakeTrustZoneService) GetTrustZone(ctx context.Context, req *trustzonepb.GetTrustZoneRequest) (*trustzonepb.GetTrustZoneResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedGetRequest, req)
	return &trustzonepb.GetTrustZoneResponse{TrustZone: f.getResponse}, nil
}

func (f *fakeTrustZoneService) ListTrustZones(ctx context.Context, req *trustzonepb.ListTrustZonesRequest) (*trustzonepb.ListTrustZonesResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedListRequest, req)
	return &trustzonepb.ListTrustZonesResponse{
		TrustZones: f.listResponse,
		Pagination: &paginationpb.PageResponse{NextPageToken: "next-token"},
	}, nil
}

func (f *fakeTrustZoneService) UpdateTrustZone(ctx context.Context, req *trustzonepb.UpdateTrustZoneRequest) (*trustzonepb.UpdateTrustZoneResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedUpdateRequest, req)
	return &trustzonepb.UpdateTrustZoneResponse{TrustZone: f.updateResponse}, nil
}

func (f *fakeTrustZoneService) PublishTrustZoneBundle(ctx context.Context, req *trustzonepb.PublishTrustZoneBundleRequest) (*trustzonepb.PublishTrustZoneBundleResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedPublishBundleRequest, req)
	return &trustzonepb.PublishTrustZoneBundleResponse{}, nil
}

func (f *fakeTrustZoneService) RegisterCofideAgent(ctx context.Context, req *trustzonepb.RegisterCofideAgentRequest) (*trustzonepb.RegisterCofideAgentResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedRegisterCofideAgentRequest, req)
	return &trustzonepb.RegisterCofideAgentResponse{}, nil
}

func (f *fakeTrustZoneService) RegisterTrustZoneServer(ctx context.Context, req *trustzonepb.RegisterTrustZoneServerRequest) (*trustzonepb.RegisterTrustZoneServerResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedRegisterTrustZoneServerRequest, req)
	return &trustzonepb.RegisterTrustZoneServerResponse{}, nil
}
