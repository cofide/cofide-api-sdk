// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"testing"
	"time"

	workloadsuppressionrulesvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/workload_suppression_rule_service/v1alpha1"
	workloadsuppressionrulepb "github.com/cofide/cofide-api-sdk/gen/go/proto/workload_suppression_rule/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func newFake(t *testing.T) *fakeconnect.FakeConnect {
	t.Helper()
	f := fakeconnect.New()
	f.Organizations[test.FakeOrganizationID] = test.FakeOrganization()
	f.TrustZones[test.FakeTrustZoneID] = test.FakeTrustZone()
	f.Clusters[test.FakeClusterID] = test.FakeCluster()
	return f
}

func Test_fakeWorkloadSuppressionRuleClient_CreateWorkloadSuppressionRule(t *testing.T) {
	f := newFake(t)
	client := New(f)

	rule := test.FakeWorkloadSuppressionRule()
	rule.Id = ""

	now := time.Now()
	createdRule, err := client.CreateWorkloadSuppressionRule(t.Context(), rule)
	require.NoError(t, err)
	assert.NotEmpty(t, createdRule.GetId())
	assert.GreaterOrEqual(t, createdRule.GetCreatedAt().AsTime(), now)
	assert.Nil(t, createdRule.GetLastUpdatedAt())
	rule.Id = createdRule.GetId()
	rule.CreatedAt = createdRule.CreatedAt
	assert.EqualExportedValues(t, rule, createdRule)
	assert.EqualExportedValues(t, rule, f.WorkloadSuppressionRules[createdRule.GetId()])
}

func Test_fakeWorkloadSuppressionRuleClient_CreateWorkloadSuppressionRule_InvalidRefs(t *testing.T) {
	f := newFake(t)
	client := New(f)

	invalidOrg := proto.Clone(test.FakeWorkloadSuppressionRule()).(*workloadsuppressionrulepb.WorkloadSuppressionRule)
	invalidOrg.OrgId = "does-not-exist"
	_, err := client.CreateWorkloadSuppressionRule(t.Context(), invalidOrg)
	require.Error(t, err)

	invalidTrustZone := proto.Clone(test.FakeWorkloadSuppressionRule()).(*workloadsuppressionrulepb.WorkloadSuppressionRule)
	invalidTrustZone.GetKubernetesPod().TrustZoneIds = []string{"does-not-exist"}
	_, err = client.CreateWorkloadSuppressionRule(t.Context(), invalidTrustZone)
	require.Error(t, err)

	invalidCluster := proto.Clone(test.FakeWorkloadSuppressionRule()).(*workloadsuppressionrulepb.WorkloadSuppressionRule)
	invalidCluster.GetKubernetesPod().ClusterIds = []string{"does-not-exist"}
	_, err = client.CreateWorkloadSuppressionRule(t.Context(), invalidCluster)
	require.Error(t, err)
}

func Test_fakeWorkloadSuppressionRuleClient_DestroyWorkloadSuppressionRule(t *testing.T) {
	f := newFake(t)
	client := New(f)

	err := client.DestroyWorkloadSuppressionRule(t.Context(), test.FakeWorkloadSuppressionRuleID)
	require.Error(t, err)
	assert.Equal(t, codes.NotFound, status.Code(err))

	f.WorkloadSuppressionRules[test.FakeWorkloadSuppressionRuleID] = test.FakeWorkloadSuppressionRule()

	err = client.DestroyWorkloadSuppressionRule(t.Context(), test.FakeWorkloadSuppressionRuleID)
	require.NoError(t, err)
	assert.Nil(t, f.WorkloadSuppressionRules[test.FakeWorkloadSuppressionRuleID])
}

func Test_fakeWorkloadSuppressionRuleClient_GetWorkloadSuppressionRule(t *testing.T) {
	f := newFake(t)
	client := New(f)

	_, err := client.GetWorkloadSuppressionRule(t.Context(), test.FakeWorkloadSuppressionRuleID)
	require.Error(t, err)
	assert.Equal(t, codes.NotFound, status.Code(err))

	rule := test.FakeWorkloadSuppressionRule()
	f.WorkloadSuppressionRules[test.FakeWorkloadSuppressionRuleID] = rule

	gotRule, err := client.GetWorkloadSuppressionRule(t.Context(), test.FakeWorkloadSuppressionRuleID)
	require.NoError(t, err)
	assert.EqualExportedValues(t, rule, gotRule)
}

func Test_fakeWorkloadSuppressionRuleClient_ListWorkloadSuppressionRules(t *testing.T) {
	f := newFake(t)
	client := New(f)

	rules, err := client.ListWorkloadSuppressionRules(t.Context(), nil)
	require.NoError(t, err)
	assert.Empty(t, rules)

	rule := test.FakeWorkloadSuppressionRule()
	f.WorkloadSuppressionRules[test.FakeWorkloadSuppressionRuleID] = rule

	rules, err = client.ListWorkloadSuppressionRules(t.Context(), nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*workloadsuppressionrulepb.WorkloadSuppressionRule{rule}, rules)

	rules, err = client.ListWorkloadSuppressionRules(t.Context(), &workloadsuppressionrulesvcpb.ListWorkloadSuppressionRulesRequest_Filter{OrgIds: []string{"other-org"}})
	require.NoError(t, err)
	assert.Empty(t, rules)

	disabled := false
	rules, err = client.ListWorkloadSuppressionRules(t.Context(), &workloadsuppressionrulesvcpb.ListWorkloadSuppressionRulesRequest_Filter{Enabled: &disabled})
	require.NoError(t, err)
	assert.Empty(t, rules)
}

func Test_fakeWorkloadSuppressionRuleClient_UpdateWorkloadSuppressionRule(t *testing.T) {
	f := newFake(t)
	client := New(f)

	existingRule, err := client.CreateWorkloadSuppressionRule(t.Context(), test.FakeWorkloadSuppressionRule())
	require.NoError(t, err)

	rule := proto.Clone(existingRule).(*workloadsuppressionrulepb.WorkloadSuppressionRule)
	rule.Name = "new-name"
	rule.Description = "new-description"
	rule.Enabled = false

	// No change with empty update mask
	now := time.Now()
	updatedRule, err := client.UpdateWorkloadSuppressionRule(t.Context(), rule, &workloadsuppressionrulesvcpb.UpdateWorkloadSuppressionRuleRequest_UpdateMask{})
	require.NoError(t, err)
	assert.GreaterOrEqual(t, updatedRule.GetLastUpdatedAt().AsTime(), now)
	existingRule.LastUpdatedAt = updatedRule.LastUpdatedAt
	assert.EqualExportedValues(t, existingRule, updatedRule)
	assert.EqualExportedValues(t, updatedRule, f.WorkloadSuppressionRules[updatedRule.GetId()])

	// Changes with non-empty update mask
	now = time.Now()
	updatedRule, err = client.UpdateWorkloadSuppressionRule(t.Context(), rule, &workloadsuppressionrulesvcpb.UpdateWorkloadSuppressionRuleRequest_UpdateMask{Name: true, Enabled: true})
	require.NoError(t, err)
	assert.GreaterOrEqual(t, updatedRule.GetLastUpdatedAt().AsTime(), now)
	assert.Equal(t, "new-name", updatedRule.GetName())
	assert.Equal(t, existingRule.GetDescription(), updatedRule.GetDescription())
	assert.False(t, updatedRule.GetEnabled())
	assert.EqualExportedValues(t, updatedRule, f.WorkloadSuppressionRules[updatedRule.GetId()])

	// Full replacement with nil update mask
	now = time.Now()
	updatedRule, err = client.UpdateWorkloadSuppressionRule(t.Context(), rule, nil)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, updatedRule.GetLastUpdatedAt().AsTime(), now)
	rule.Id = existingRule.GetId()
	rule.OrgId = existingRule.GetOrgId()
	rule.CreatedAt = existingRule.CreatedAt
	rule.LastUpdatedAt = updatedRule.LastUpdatedAt
	assert.EqualExportedValues(t, rule, updatedRule)
	assert.EqualExportedValues(t, updatedRule, f.WorkloadSuppressionRules[updatedRule.GetId()])
}

func Test_fakeWorkloadSuppressionRuleClient_UpdateWorkloadSuppressionRule_NotFound(t *testing.T) {
	f := newFake(t)
	client := New(f)

	_, err := client.UpdateWorkloadSuppressionRule(t.Context(), test.FakeWorkloadSuppressionRule(), nil)
	require.Error(t, err)
	assert.Equal(t, codes.NotFound, status.Code(err))
}
