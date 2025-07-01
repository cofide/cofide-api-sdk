// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"testing"

	identitypb "github.com/cofide/cofide-api-sdk/gen/go/proto/identity/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_fakeIdentityClient_GetIdentity(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	_, err := client.GetIdentity(ctx, test.FakeIdentityID)
	require.Error(t, err)

	identity := test.FakeIdentity()
	fake.Identities[test.FakeIdentityID] = identity

	gotIdentity, err := client.GetIdentity(ctx, test.FakeIdentityID)
	require.NoError(t, err)
	assert.EqualExportedValues(t, identity, gotIdentity)
}

func Test_fakeIdentityClient_ListIdentities(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	identity := test.FakeIdentity()

	identities, err := client.ListIdentities(ctx, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*identitypb.Identity{}, identities)

	fake.Identities[test.FakeIdentityID] = test.FakeIdentity()

	identities, err = client.ListIdentities(ctx, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*identitypb.Identity{identity}, identities)
}
