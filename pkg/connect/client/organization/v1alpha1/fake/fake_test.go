// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"testing"

	organizationpb "github.com/cofide/cofide-api-sdk/gen/go/proto/organization/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_fakeOrganizationClient_GetOrganization(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	_, err := client.GetOrganization(ctx, test.FakeOrganizationID)
	require.Error(t, err)

	organization := test.FakeOrganization()
	fake.Organizations[test.FakeOrganizationID] = organization

	gotOrganization, err := client.GetOrganization(ctx, test.FakeOrganizationID)
	require.NoError(t, err)
	assert.EqualExportedValues(t, organization, gotOrganization)
}

func Test_fakeOrganizationClient_ListOrganizations(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	_, err := client.ListOrganizations(ctx, nil)
	require.NoError(t, err)

	organization := test.FakeOrganization()
	fake.Organizations[test.FakeOrganizationID] = organization

	organizations, err := client.ListOrganizations(ctx, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*organizationpb.Organization{organization}, organizations)
}
