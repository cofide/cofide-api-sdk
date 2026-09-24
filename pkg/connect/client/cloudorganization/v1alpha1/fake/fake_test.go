// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"testing"

	cloudorganizationsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/cloud_organization_service/v1alpha1"
	cloudorganizationpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_organization/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_fakeCloudOrganizationClient_CreateCloudOrganization(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	cloudOrganization := test.FakeCloudOrganization()

	_, err := client.CreateCloudOrganization(ctx, cloudOrganization)
	require.Error(t, err)

	fake.Organizations[test.FakeOrganizationID] = test.FakeOrganization()

	createdCloudOrganization, err := client.CreateCloudOrganization(ctx, cloudOrganization)
	require.NoError(t, err)
	cloudOrganization.Id = createdCloudOrganization.Id
	assert.NotNil(t, createdCloudOrganization.GetCreatedAt())
	cloudOrganization.CreatedAt = createdCloudOrganization.CreatedAt
	assert.EqualExportedValues(t, cloudOrganization, createdCloudOrganization)
	assert.EqualExportedValues(t, cloudOrganization, fake.CloudOrganizations[createdCloudOrganization.Id])
}

func Test_fakeCloudOrganizationClient_DestroyCloudOrganization(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	err := client.DestroyCloudOrganization(ctx, test.FakeCloudOrganizationID)
	require.Error(t, err)

	fake.CloudOrganizations[test.FakeCloudOrganizationID] = test.FakeCloudOrganization()

	err = client.DestroyCloudOrganization(ctx, test.FakeCloudOrganizationID)
	require.NoError(t, err)
	require.Nil(t, fake.CloudOrganizations[test.FakeCloudOrganizationID])
}

func Test_fakeCloudOrganizationClient_GetCloudOrganization(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	_, err := client.GetCloudOrganization(ctx, test.FakeCloudOrganizationID)
	require.Error(t, err)

	cloudOrganization := test.FakeCloudOrganization()
	fake.CloudOrganizations[test.FakeCloudOrganizationID] = cloudOrganization

	gotCloudOrganization, err := client.GetCloudOrganization(ctx, test.FakeCloudOrganizationID)
	require.NoError(t, err)
	assert.EqualExportedValues(t, cloudOrganization, gotCloudOrganization)
}

func Test_fakeCloudOrganizationClient_ListCloudOrganizations(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	cloudOrganizations, err := client.ListCloudOrganizations(ctx, nil)
	require.NoError(t, err)
	assert.Empty(t, cloudOrganizations)

	cloudOrganization := test.FakeCloudOrganization()
	fake.CloudOrganizations[test.FakeCloudOrganizationID] = cloudOrganization

	cloudOrganizations, err = client.ListCloudOrganizations(ctx, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*cloudorganizationpb.CloudOrganization{cloudOrganization}, cloudOrganizations)

	cloudOrganizations, err = client.ListCloudOrganizations(ctx, &cloudorganizationsvcpb.ListCloudOrganizationsRequest_Filter{OrgIds: []string{"other-org-id"}})
	require.NoError(t, err)
	assert.Empty(t, cloudOrganizations)
}

func Test_fakeCloudOrganizationClient_UpdateCloudOrganization(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	cloudOrganization := test.FakeCloudOrganization()

	_, err := client.UpdateCloudOrganization(ctx, cloudOrganization, nil)
	require.Error(t, err)

	fake.CloudOrganizations[test.FakeCloudOrganizationID] = test.FakeCloudOrganization()

	cloudOrganization = clone(cloudOrganization)
	cloudOrganization.Name = "new-cloud-org-name"
	updatedCloudOrganization, err := client.UpdateCloudOrganization(ctx, cloudOrganization, nil)
	require.NoError(t, err)
	assert.NotNil(t, updatedCloudOrganization.GetLastUpdatedAt())
	cloudOrganization.LastUpdatedAt = updatedCloudOrganization.LastUpdatedAt
	assert.EqualExportedValues(t, cloudOrganization, updatedCloudOrganization)
	assert.EqualExportedValues(t, cloudOrganization, fake.CloudOrganizations[test.FakeCloudOrganizationID])
}

func Test_fakeCloudOrganizationClient_UpdateCloudOrganization_partial(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	existing := test.FakeCloudOrganization()
	fake.CloudOrganizations[test.FakeCloudOrganizationID] = existing

	update := clone(existing)
	update.Name = "should-not-be-applied"
	update.DiscoveryEnabled = true

	updateMask := &cloudorganizationsvcpb.UpdateCloudOrganizationRequest_UpdateMask{DiscoveryEnabled: true}
	updatedCloudOrganization, err := client.UpdateCloudOrganization(ctx, update, updateMask)
	require.NoError(t, err)
	assert.Equal(t, existing.GetName(), updatedCloudOrganization.GetName())
	assert.True(t, updatedCloudOrganization.GetDiscoveryEnabled())
}
