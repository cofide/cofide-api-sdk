// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"
	"testing"

	organizationsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/organization_service/v1alpha1"
	organizationpb "github.com/cofide/cofide-api-sdk/gen/go/proto/organization/v1alpha1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	fakeOrganizationID   = "fake-organization-id"
	fakeOrganizationName = "fake-organization-name"
)

// TestOrganizationClient_Unimplemented tests OrganizationClient against an unimplemented server.
// This ensures that all errors returned are not wrapped can be converted to a gRPC Status using Status.Convert.
func TestOrganizationClient_Unimplemented(t *testing.T) {
	server := test.NewTestServer(t)
	organizationsvcpb.RegisterOrganizationServiceServer(server.Server, &organizationsvcpb.UnimplementedOrganizationServiceServer{})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	cluster, err := client.GetOrganization(ctx, "")
	test.RequireUnimplemented(t, err)
	assert.Nil(t, cluster)

	organizations, err := client.ListOrganizations(ctx, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, organizations)
}

func TestOrganizationClient(t *testing.T) {
	server := test.NewTestServer(t)
	organizationsvcpb.RegisterOrganizationServiceServer(server.Server, &fakeOrganizationService{t: t})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	organization := fakeOrganization()

	gotOrganization, err := client.GetOrganization(ctx, organization.GetId())
	require.NoError(t, err)
	assert.Equal(t, organization.GetId(), gotOrganization.GetId())

	filter := &organizationsvcpb.ListOrganizationsRequest_Filter{Name: test.PtrOf(fakeOrganizationName)}
	organizations, err := client.ListOrganizations(ctx, filter)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*organizationpb.Organization{fakeOrganization()}, organizations)
}

// fakeOrganizationService provides a very simple fake OrganizationService implementation with canned responses.
type fakeOrganizationService struct {
	t *testing.T
}

func (f *fakeOrganizationService) GetOrganization(ctx context.Context, req *organizationsvcpb.GetOrganizationRequest) (*organizationsvcpb.GetOrganizationResponse, error) {
	assert.Equal(f.t, fakeOrganizationID, req.GetOrgId())
	return &organizationsvcpb.GetOrganizationResponse{Organization: fakeOrganization()}, nil
}

func (f *fakeOrganizationService) ListOrganizations(ctx context.Context, req *organizationsvcpb.ListOrganizationsRequest) (*organizationsvcpb.ListOrganizationsResponse, error) {
	assert.Equal(f.t, fakeOrganizationName, req.Filter.GetName())
	organizations := []*organizationpb.Organization{fakeOrganization()}
	return &organizationsvcpb.ListOrganizationsResponse{Organizations: organizations}, nil
}

func fakeOrganization() *organizationpb.Organization {
	return &organizationpb.Organization{
		Id:   fakeOrganizationID,
		Name: fakeOrganizationName,
	}
}
