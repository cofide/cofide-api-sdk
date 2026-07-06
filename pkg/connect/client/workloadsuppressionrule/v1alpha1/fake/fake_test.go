// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"testing"

	workloadsuppressionrulepb "github.com/cofide/cofide-api-sdk/gen/go/proto/workload_suppression_rule/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_fakeWorkloadSuppressionRuleClient_CreateWorkloadSuppressionRule(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	rule := test.FakeWorkloadSuppressionRule()

	_, err := client.CreateWorkloadSuppressionRule(ctx, rule)
	require.Error(t, err)

	fake.Organizations[test.FakeOrganizationID] = test.FakeOrganization()

	createdRule, err := client.CreateWorkloadSuppressionRule(ctx, rule)
	require.NoError(t, err)
	rule.Id = createdRule.Id
	assert.EqualExportedValues(t, rule, createdRule)
	assert.EqualExportedValues(t, rule, fake.WorkloadSuppressionRules[createdRule.Id])
}

func Test_fakeWorkloadSuppressionRuleClient_CreateWorkloadSuppressionRule_InvalidScope(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	fake.Organizations[test.FakeOrganizationID] = test.FakeOrganization()

	rule := test.FakeWorkloadSuppressionRule()
	rule.TrustZoneId = test.PtrOf(test.FakeTrustZoneID)

	_, err := client.CreateWorkloadSuppressionRule(ctx, rule)
	require.Error(t, err)

	fake.TrustZones[test.FakeTrustZoneID] = test.FakeTrustZone()

	_, err = client.CreateWorkloadSuppressionRule(ctx, rule)
	require.NoError(t, err)
}

func Test_fakeWorkloadSuppressionRuleClient_DestroyWorkloadSuppressionRule(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	err := client.DestroyWorkloadSuppressionRule(ctx, test.FakeWorkloadSuppressionRuleID)
	require.Error(t, err)

	fake.WorkloadSuppressionRules[test.FakeWorkloadSuppressionRuleID] = test.FakeWorkloadSuppressionRule()

	err = client.DestroyWorkloadSuppressionRule(ctx, test.FakeWorkloadSuppressionRuleID)
	require.NoError(t, err)
	require.Nil(t, fake.WorkloadSuppressionRules[test.FakeWorkloadSuppressionRuleID])
}

func Test_fakeWorkloadSuppressionRuleClient_GetWorkloadSuppressionRule(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	_, err := client.GetWorkloadSuppressionRule(ctx, test.FakeWorkloadSuppressionRuleID)
	require.Error(t, err)

	rule := test.FakeWorkloadSuppressionRule()
	fake.WorkloadSuppressionRules[test.FakeWorkloadSuppressionRuleID] = rule

	gotRule, err := client.GetWorkloadSuppressionRule(ctx, test.FakeWorkloadSuppressionRuleID)
	require.NoError(t, err)
	assert.EqualExportedValues(t, rule, gotRule)
}

func Test_fakeWorkloadSuppressionRuleClient_ListWorkloadSuppressionRules(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	_, err := client.ListWorkloadSuppressionRules(ctx, nil)
	require.NoError(t, err)

	rule := test.FakeWorkloadSuppressionRule()
	fake.WorkloadSuppressionRules[test.FakeWorkloadSuppressionRuleID] = rule

	rules, err := client.ListWorkloadSuppressionRules(ctx, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*workloadsuppressionrulepb.WorkloadSuppressionRule{rule}, rules)
}

func Test_fakeWorkloadSuppressionRuleClient_UpdateWorkloadSuppressionRule(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	rule := test.FakeWorkloadSuppressionRule()

	_, err := client.UpdateWorkloadSuppressionRule(ctx, rule, nil)
	require.Error(t, err)

	fake.Organizations[test.FakeOrganizationID] = test.FakeOrganization()
	fake.WorkloadSuppressionRules[test.FakeWorkloadSuppressionRuleID] = rule

	rule = clone(rule)
	rule.Enabled = false
	updatedRule, err := client.UpdateWorkloadSuppressionRule(ctx, rule, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, rule, updatedRule)
	assert.EqualExportedValues(t, rule, fake.WorkloadSuppressionRules[test.FakeWorkloadSuppressionRuleID])
}
