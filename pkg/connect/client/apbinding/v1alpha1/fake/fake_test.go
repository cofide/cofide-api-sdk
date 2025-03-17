// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"testing"

	apbindingpb "github.com/cofide/cofide-api-sdk/gen/go/proto/ap_binding/v1alpha1"
	trustzonepb "github.com/cofide/cofide-api-sdk/gen/go/proto/trust_zone/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func Test_fakeAPBindingClient_CreateAPBinding(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	binding := test.FakeAPBinding()

	_, err := client.CreateAPBinding(ctx, binding)
	require.Error(t, err)

	fake.TrustZones[test.FakeTrustZoneID] = test.FakeTrustZone()

	_, err = client.CreateAPBinding(ctx, binding)
	require.Error(t, err)

	fake.AttestationPolicies[test.FakeAttestationPolicyID] = test.FakeAttestationPolicy()

	createdBinding, err := client.CreateAPBinding(ctx, binding)
	require.NoError(t, err)
	binding.Id = createdBinding.Id
	assert.EqualExportedValues(t, binding, createdBinding)
	assert.EqualExportedValues(t, binding, fake.APBindings[*createdBinding.Id])
}

func Test_fakeAPBindingClient_DestroyAPBinding(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	err := client.DestroyAPBinding(ctx, test.FakeAPBindingID)
	require.Error(t, err)

	fake.APBindings[test.FakeAPBindingID] = test.FakeAPBinding()

	err = client.DestroyAPBinding(ctx, test.FakeAPBindingID)
	require.NoError(t, err)
	require.Nil(t, fake.APBindings[test.FakeAPBindingID])
}

func Test_fakeAPBindingClient_GetAPBinding(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	_, err := client.GetAPBinding(ctx, test.FakeAPBindingID)
	require.Error(t, err)

	binding := test.FakeAPBinding()
	fake.APBindings[test.FakeAPBindingID] = binding

	gotBinding, err := client.GetAPBinding(ctx, test.FakeAPBindingID)
	require.NoError(t, err)
	assert.EqualExportedValues(t, binding, gotBinding)
}

func Test_fakeAPBindingClient_ListAPBindings(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	_, err := client.ListAPBindings(ctx, nil)
	require.NoError(t, err)

	binding := test.FakeAPBinding()
	fake.APBindings[test.FakeAPBindingID] = binding

	bindings, err := client.ListAPBindings(ctx, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*apbindingpb.APBinding{binding}, bindings)
}

func Test_fakeAPBindingClient_UpdateAPBinding(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	binding := test.FakeAPBinding()

	_, err := client.UpdateAPBinding(ctx, binding)
	require.Error(t, err)

	fake.TrustZones[test.FakeTrustZoneID] = test.FakeTrustZone()
	fake.AttestationPolicies[test.FakeAttestationPolicyID] = test.FakeAttestationPolicy()
	fake.APBindings[test.FakeAPBindingID] = binding

	newTrustZoneID := "new-trust-zone-id"
	binding = clone(binding)
	binding.TrustZoneId = &newTrustZoneID
	_, err = client.UpdateAPBinding(ctx, binding)
	require.Error(t, err)

	trustZone := proto.Clone(test.FakeTrustZone()).(*trustzonepb.TrustZone)
	trustZone.Id = &newTrustZoneID
	fake.TrustZones[newTrustZoneID] = trustZone

	updatedBinding, err := client.UpdateAPBinding(ctx, binding)
	require.NoError(t, err)
	assert.EqualExportedValues(t, binding, updatedBinding)
	assert.EqualExportedValues(t, binding, fake.APBindings[test.FakeAPBindingID])
}
