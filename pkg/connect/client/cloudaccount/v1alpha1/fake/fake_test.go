// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"testing"

	cloudaccountsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/cloud_account_service/v1alpha1"
	cloudaccountpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_account/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_fakeCloudAccountClient_CreateCloudAccount(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	cloudAccount := test.FakeCloudAccount()

	_, err := client.CreateCloudAccount(ctx, cloudAccount)
	require.Error(t, err)

	fake.Organizations[test.FakeOrganizationID] = test.FakeOrganization()

	_, err = client.CreateCloudAccount(ctx, cloudAccount)
	require.Error(t, err)

	fake.CloudOrganizations[test.FakeCloudOrganizationID] = test.FakeCloudOrganization()

	createdCloudAccount, err := client.CreateCloudAccount(ctx, cloudAccount)
	require.NoError(t, err)
	cloudAccount.Id = createdCloudAccount.Id
	assert.NotNil(t, createdCloudAccount.GetCreatedAt())
	cloudAccount.CreatedAt = createdCloudAccount.CreatedAt
	assert.EqualExportedValues(t, cloudAccount, createdCloudAccount)
	assert.EqualExportedValues(t, cloudAccount, fake.CloudAccounts[createdCloudAccount.Id])
}

func Test_fakeCloudAccountClient_DestroyCloudAccount(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	err := client.DestroyCloudAccount(ctx, test.FakeCloudAccountID)
	require.Error(t, err)

	fake.CloudAccounts[test.FakeCloudAccountID] = test.FakeCloudAccount()

	err = client.DestroyCloudAccount(ctx, test.FakeCloudAccountID)
	require.NoError(t, err)
	require.Nil(t, fake.CloudAccounts[test.FakeCloudAccountID])
}

func Test_fakeCloudAccountClient_GetCloudAccount(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	_, err := client.GetCloudAccount(ctx, test.FakeCloudAccountID)
	require.Error(t, err)

	cloudAccount := test.FakeCloudAccount()
	fake.CloudAccounts[test.FakeCloudAccountID] = cloudAccount

	gotCloudAccount, err := client.GetCloudAccount(ctx, test.FakeCloudAccountID)
	require.NoError(t, err)
	assert.EqualExportedValues(t, cloudAccount, gotCloudAccount)
}

func Test_fakeCloudAccountClient_ListCloudAccounts(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	cloudAccounts, err := client.ListCloudAccounts(ctx, nil)
	require.NoError(t, err)
	assert.Empty(t, cloudAccounts)

	cloudAccount := test.FakeCloudAccount()
	fake.CloudAccounts[test.FakeCloudAccountID] = cloudAccount

	cloudAccounts, err = client.ListCloudAccounts(ctx, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*cloudaccountpb.CloudAccount{cloudAccount}, cloudAccounts)

	cloudAccounts, err = client.ListCloudAccounts(ctx, &cloudaccountsvcpb.ListCloudAccountsRequest_Filter{CloudOrganizationIds: []string{"other-cloud-org-id"}})
	require.NoError(t, err)
	assert.Empty(t, cloudAccounts)

	suppressed := clone(cloudAccount)
	suppressed.Suppressed = true
	fake.CloudAccounts[test.FakeCloudAccountID] = suppressed

	cloudAccounts, err = client.ListCloudAccounts(ctx, nil)
	require.NoError(t, err)
	assert.Empty(t, cloudAccounts)

	cloudAccounts, err = client.ListCloudAccounts(ctx, &cloudaccountsvcpb.ListCloudAccountsRequest_Filter{IncludeSuppressed: true})
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*cloudaccountpb.CloudAccount{suppressed}, cloudAccounts)
}

func Test_fakeCloudAccountClient_UpdateCloudAccount(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	cloudAccount := test.FakeCloudAccount()

	_, err := client.UpdateCloudAccount(ctx, cloudAccount, nil)
	require.Error(t, err)

	fake.CloudAccounts[test.FakeCloudAccountID] = test.FakeCloudAccount()

	cloudAccount = clone(cloudAccount)
	cloudAccount.Name = "new-cloud-account-name"
	updatedCloudAccount, err := client.UpdateCloudAccount(ctx, cloudAccount, nil)
	require.NoError(t, err)
	assert.NotNil(t, updatedCloudAccount.GetLastUpdatedAt())
	cloudAccount.LastUpdatedAt = updatedCloudAccount.LastUpdatedAt
	assert.EqualExportedValues(t, cloudAccount, updatedCloudAccount)
	assert.EqualExportedValues(t, cloudAccount, fake.CloudAccounts[test.FakeCloudAccountID])
}

func Test_fakeCloudAccountClient_UpdateCloudAccount_partial(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	existing := test.FakeCloudAccount()
	fake.CloudAccounts[test.FakeCloudAccountID] = existing

	update := clone(existing)
	update.Name = "should-not-be-applied"
	update.Suppressed = true

	updateMask := &cloudaccountsvcpb.UpdateCloudAccountRequest_UpdateMask{Suppressed: true}
	updatedCloudAccount, err := client.UpdateCloudAccount(ctx, update, updateMask)
	require.NoError(t, err)
	assert.Equal(t, existing.GetName(), updatedCloudAccount.GetName())
	assert.True(t, updatedCloudAccount.GetSuppressed())
}
