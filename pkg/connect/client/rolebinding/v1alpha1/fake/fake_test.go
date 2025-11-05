// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"testing"

	rolebindingpb "github.com/cofide/cofide-api-sdk/gen/go/proto/role_binding/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_fakeRoleBindingClient_CreateRoleBinding(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	roleBinding := test.FakeRoleBinding()

	createdRoleBinding, err := client.CreateRoleBinding(ctx, roleBinding)
	require.NoError(t, err)
	roleBinding.Id = createdRoleBinding.Id
	assert.EqualExportedValues(t, roleBinding, createdRoleBinding)
	assert.EqualExportedValues(t, roleBinding, fake.RoleBindings[createdRoleBinding.Id])
}

func Test_fakeRoleBindingClient_DestroyRoleBinding(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	err := client.DestroyRoleBinding(ctx, test.FakeRoleBindingID)
	require.Error(t, err)

	fake.RoleBindings[test.FakeRoleBindingID] = test.FakeRoleBinding()

	err = client.DestroyRoleBinding(ctx, test.FakeRoleBindingID)
	require.NoError(t, err)
	require.Nil(t, fake.RoleBindings[test.FakeRoleBindingID])
}

func Test_fakeRoleBindingClient_GetRoleBinding(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	roleBinding := test.FakeRoleBinding()

	_, err := client.GetRoleBinding(ctx, test.FakeRoleBindingID)
	require.Error(t, err)

	fake.RoleBindings[test.FakeRoleBindingID] = test.FakeRoleBinding()

	gotRoleBinding, err := client.GetRoleBinding(ctx, test.FakeRoleBindingID)
	require.NoError(t, err)
	assert.EqualExportedValues(t, roleBinding, gotRoleBinding)
	assert.EqualExportedValues(t, roleBinding, fake.RoleBindings[test.FakeRoleBindingID])
}

func Test_fakeRoleBindingClient_ListRoleBindings(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	roleBinding := test.FakeRoleBinding()

	roleBindings, err := client.ListRoleBindings(ctx, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*rolebindingpb.RoleBinding{}, roleBindings)

	fake.RoleBindings[test.FakeRoleBindingID] = test.FakeRoleBinding()

	roleBindings, err = client.ListRoleBindings(ctx, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*rolebindingpb.RoleBinding{roleBinding}, roleBindings)
}
