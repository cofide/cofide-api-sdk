// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"testing"

	clustersuppressionrulepb "github.com/cofide/cofide-api-sdk/gen/go/proto/cluster_suppression_rule/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_fakeClusterSuppressionRuleClient_CreateClusterSuppressionRule(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	rule := test.FakeClusterSuppressionRule()

	_, err := client.CreateClusterSuppressionRule(ctx, rule)
	require.Error(t, err)

	fake.Organizations[test.FakeOrganizationID] = test.FakeOrganization()

	createdRule, err := client.CreateClusterSuppressionRule(ctx, rule)
	require.NoError(t, err)
	rule.Id = createdRule.Id
	assert.EqualExportedValues(t, rule, createdRule)
	assert.EqualExportedValues(t, rule, fake.ClusterSuppressionRules[createdRule.Id])
}

func Test_fakeClusterSuppressionRuleClient_CreateClusterSuppressionRule_InvalidScope(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	fake.Organizations[test.FakeOrganizationID] = test.FakeOrganization()

	rule := test.FakeClusterSuppressionRule()
	rule.TrustZoneId = test.PtrOf(test.FakeTrustZoneID)

	_, err := client.CreateClusterSuppressionRule(ctx, rule)
	require.Error(t, err)

	fake.TrustZones[test.FakeTrustZoneID] = test.FakeTrustZone()

	_, err = client.CreateClusterSuppressionRule(ctx, rule)
	require.NoError(t, err)
}

func Test_fakeClusterSuppressionRuleClient_DestroyClusterSuppressionRule(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	err := client.DestroyClusterSuppressionRule(ctx, test.FakeClusterSuppressionRuleID)
	require.Error(t, err)

	fake.ClusterSuppressionRules[test.FakeClusterSuppressionRuleID] = test.FakeClusterSuppressionRule()

	err = client.DestroyClusterSuppressionRule(ctx, test.FakeClusterSuppressionRuleID)
	require.NoError(t, err)
	require.Nil(t, fake.ClusterSuppressionRules[test.FakeClusterSuppressionRuleID])
}

func Test_fakeClusterSuppressionRuleClient_GetClusterSuppressionRule(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	_, err := client.GetClusterSuppressionRule(ctx, test.FakeClusterSuppressionRuleID)
	require.Error(t, err)

	rule := test.FakeClusterSuppressionRule()
	fake.ClusterSuppressionRules[test.FakeClusterSuppressionRuleID] = rule

	gotRule, err := client.GetClusterSuppressionRule(ctx, test.FakeClusterSuppressionRuleID)
	require.NoError(t, err)
	assert.EqualExportedValues(t, rule, gotRule)
}

func Test_fakeClusterSuppressionRuleClient_ListClusterSuppressionRules(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	_, err := client.ListClusterSuppressionRules(ctx, nil)
	require.NoError(t, err)

	rule := test.FakeClusterSuppressionRule()
	fake.ClusterSuppressionRules[test.FakeClusterSuppressionRuleID] = rule

	rules, err := client.ListClusterSuppressionRules(ctx, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*clustersuppressionrulepb.ClusterSuppressionRule{rule}, rules)
}

func Test_fakeClusterSuppressionRuleClient_UpdateClusterSuppressionRule(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	rule := test.FakeClusterSuppressionRule()

	_, err := client.UpdateClusterSuppressionRule(ctx, rule, nil)
	require.Error(t, err)

	fake.Organizations[test.FakeOrganizationID] = test.FakeOrganization()
	fake.ClusterSuppressionRules[test.FakeClusterSuppressionRuleID] = rule

	rule = clone(rule)
	rule.Enabled = false
	updatedRule, err := client.UpdateClusterSuppressionRule(ctx, rule, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, rule, updatedRule)
	assert.EqualExportedValues(t, rule, fake.ClusterSuppressionRules[test.FakeClusterSuppressionRuleID])
}
