// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"
	"testing"

	cloudorganizationsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/cloud_organization_service/v1alpha1"
	cloudorganizationpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_organization/v1alpha1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCloudOrganizationClient_implementsMethods(t *testing.T) {
	test.AssertClientImplementsService(t, &cloudOrganizationClient{}, cloudorganizationsvcpb.CloudOrganizationService_ServiceDesc)
}

// TestCloudOrganizationClient_Unimplemented tests CloudOrganizationClient against an unimplemented server.
// This ensures that all errors returned are not wrapped and can be converted to a gRPC Status using Status.Convert.
func TestCloudOrganizationClient_Unimplemented(t *testing.T) {
	server := test.NewTestServer(t)
	cloudorganizationsvcpb.RegisterCloudOrganizationServiceServer(server.Server, &cloudorganizationsvcpb.UnimplementedCloudOrganizationServiceServer{})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	cloudOrganization, err := client.CreateCloudOrganization(ctx, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, cloudOrganization)

	err = client.DestroyCloudOrganization(ctx, "")
	test.RequireUnimplemented(t, err)

	cloudOrganization, err = client.GetCloudOrganization(ctx, "")
	test.RequireUnimplemented(t, err)
	assert.Nil(t, cloudOrganization)

	cloudOrganizations, err := client.ListCloudOrganizations(ctx, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, cloudOrganizations)

	cloudOrganization, err = client.UpdateCloudOrganization(ctx, nil, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, cloudOrganization)
}

func TestCloudOrganizationClient(t *testing.T) {
	server := test.NewTestServer(t)
	cloudorganizationsvcpb.RegisterCloudOrganizationServiceServer(server.Server, &fakeCloudOrganizationService{t: t})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	cloudOrganization := test.FakeCloudOrganization()

	createdCloudOrganization, err := client.CreateCloudOrganization(ctx, cloudOrganization)
	require.NoError(t, err)
	assert.EqualExportedValues(t, cloudOrganization, createdCloudOrganization)

	err = client.DestroyCloudOrganization(ctx, cloudOrganization.GetId())
	require.NoError(t, err)

	gotCloudOrganization, err := client.GetCloudOrganization(ctx, cloudOrganization.GetId())
	require.NoError(t, err)
	assert.Equal(t, cloudOrganization.GetId(), gotCloudOrganization.GetId())

	filter := &cloudorganizationsvcpb.ListCloudOrganizationsRequest_Filter{OrgIds: []string{test.FakeOrganizationID}}
	cloudOrganizations, err := client.ListCloudOrganizations(ctx, filter)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*cloudorganizationpb.CloudOrganization{test.FakeCloudOrganization()}, cloudOrganizations)

	updatedCloudOrganization, err := client.UpdateCloudOrganization(ctx, cloudOrganization, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, cloudOrganization, updatedCloudOrganization)
}

// fakeCloudOrganizationService provides a simple fake CloudOrganizationService implementation with canned responses.
type fakeCloudOrganizationService struct {
	t *testing.T
}

func (f *fakeCloudOrganizationService) CreateCloudOrganization(ctx context.Context, req *cloudorganizationsvcpb.CreateCloudOrganizationRequest) (*cloudorganizationsvcpb.CreateCloudOrganizationResponse, error) {
	assert.EqualExportedValues(f.t, test.FakeCloudOrganization(), req.CloudOrganization)
	return &cloudorganizationsvcpb.CreateCloudOrganizationResponse{CloudOrganization: req.CloudOrganization}, nil
}

func (f *fakeCloudOrganizationService) DestroyCloudOrganization(ctx context.Context, req *cloudorganizationsvcpb.DestroyCloudOrganizationRequest) (*cloudorganizationsvcpb.DestroyCloudOrganizationResponse, error) {
	assert.Equal(f.t, test.FakeCloudOrganizationID, req.GetCloudOrganizationId())
	return &cloudorganizationsvcpb.DestroyCloudOrganizationResponse{}, nil
}

func (f *fakeCloudOrganizationService) GetCloudOrganization(ctx context.Context, req *cloudorganizationsvcpb.GetCloudOrganizationRequest) (*cloudorganizationsvcpb.GetCloudOrganizationResponse, error) {
	assert.Equal(f.t, test.FakeCloudOrganizationID, req.GetCloudOrganizationId())
	return &cloudorganizationsvcpb.GetCloudOrganizationResponse{CloudOrganization: test.FakeCloudOrganization()}, nil
}

func (f *fakeCloudOrganizationService) ListCloudOrganizations(ctx context.Context, req *cloudorganizationsvcpb.ListCloudOrganizationsRequest) (*cloudorganizationsvcpb.ListCloudOrganizationsResponse, error) {
	assert.Equal(f.t, []string{test.FakeOrganizationID}, req.Filter.GetOrgIds())
	return &cloudorganizationsvcpb.ListCloudOrganizationsResponse{CloudOrganizations: []*cloudorganizationpb.CloudOrganization{test.FakeCloudOrganization()}}, nil
}

func (f *fakeCloudOrganizationService) UpdateCloudOrganization(ctx context.Context, req *cloudorganizationsvcpb.UpdateCloudOrganizationRequest) (*cloudorganizationsvcpb.UpdateCloudOrganizationResponse, error) {
	assert.EqualExportedValues(f.t, test.FakeCloudOrganization(), req.CloudOrganization)
	return &cloudorganizationsvcpb.UpdateCloudOrganizationResponse{CloudOrganization: test.FakeCloudOrganization()}, nil
}
