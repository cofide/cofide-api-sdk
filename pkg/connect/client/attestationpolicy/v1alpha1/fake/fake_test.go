// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"testing"

	attestationpolicypb "github.com/cofide/cofide-api-sdk/gen/go/proto/attestation_policy/v1alpha1"
	"github.com/cofide/cofide-api-sdk/gen/go/proto/connect/attestation_policy_service/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_fakeAttestationPolicyClient_CreateAttestationPolicy(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	policy := test.FakeAttestationPolicy()

	createdPolicy, err := client.CreateAttestationPolicy(ctx, policy)
	require.NoError(t, err)
	policy.Id = createdPolicy.Id
	assert.EqualExportedValues(t, policy, createdPolicy)
	assert.EqualExportedValues(t, policy, fake.AttestationPolicies[*createdPolicy.Id])
}

func Test_fakeAttestationPolicyClient_DestroyAttestationPolicy(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	err := client.DestroyAttestationPolicy(ctx, test.FakeAttestationPolicyID)
	require.Error(t, err)

	fake.AttestationPolicies[test.FakeAttestationPolicyID] = test.FakeAttestationPolicy()

	err = client.DestroyAttestationPolicy(ctx, test.FakeAttestationPolicyID)
	require.NoError(t, err)
	require.Nil(t, fake.AttestationPolicies[test.FakeAttestationPolicyID])
}

func Test_fakeAttestationPolicyClient_GetAttestationPolicy(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	_, err := client.GetAttestationPolicy(ctx, test.FakeAttestationPolicyID)
	require.Error(t, err)

	policy := test.FakeAttestationPolicy()
	fake.AttestationPolicies[test.FakeAttestationPolicyID] = policy

	gotPolicy, err := client.GetAttestationPolicy(ctx, test.FakeAttestationPolicyID)
	require.NoError(t, err)
	assert.EqualExportedValues(t, policy, gotPolicy)
}

func Test_fakeAttestationPolicyClient_ListAttestationPolicies(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	_, err := client.ListAttestationPolicies(ctx, nil)
	require.NoError(t, err)

	policy := test.FakeAttestationPolicy()
	fake.AttestationPolicies[test.FakeAttestationPolicyID] = policy

	policys, err := client.ListAttestationPolicies(ctx, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*attestationpolicypb.AttestationPolicy{policy}, policys)
}

func Test_fakeAttestationPolicyClient_ListAttestationPolicies_TrustZoneIDFilter(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	policy := test.FakeAttestationPolicy()
	fake.AttestationPolicies[test.FakeAttestationPolicyID] = policy
	fake.APBindings[*test.FakeAPBinding().Id] = test.FakeAPBinding()

	policies, err := client.ListAttestationPolicies(ctx, &v1alpha1.ListAttestationPoliciesRequest_Filter{TrustZoneId: test.PtrOf(test.FakeTrustZoneID)})
	require.NoError(t, err)

	assert.EqualExportedValues(t, []*attestationpolicypb.AttestationPolicy{policy}, policies)

	expectEmpty, err := client.ListAttestationPolicies(ctx, &v1alpha1.ListAttestationPoliciesRequest_Filter{TrustZoneId: test.PtrOf("non-existent")})
	require.NoError(t, err)
	require.Empty(t, expectEmpty)
}

func Test_fakeAttestationPolicyClient_UpdateAttestationPolicy(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	policy := test.FakeAttestationPolicy()

	_, err := client.UpdateAttestationPolicy(ctx, policy)
	require.Error(t, err)

	fake.TrustZones[test.FakeTrustZoneID] = test.FakeTrustZone()
	fake.AttestationPolicies[test.FakeAttestationPolicyID] = policy

	policy = clone(policy)
	policy.Name = "new-policy-name"
	updatedPolicy, err := client.UpdateAttestationPolicy(ctx, policy)
	require.NoError(t, err)
	assert.EqualExportedValues(t, policy, updatedPolicy)
	assert.EqualExportedValues(t, policy, fake.AttestationPolicies[test.FakeAttestationPolicyID])
}
