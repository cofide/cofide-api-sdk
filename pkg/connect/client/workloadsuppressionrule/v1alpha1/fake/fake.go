// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"slices"

	workloadsuppressionrulesvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/workload_suppression_rule_service/v1alpha1"
	workloadsuppressionrulepb "github.com/cofide/cofide-api-sdk/gen/go/proto/workload_suppression_rule/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	workloadsuppressionrulev1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/workloadsuppressionrule/v1alpha1"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type fakeWorkloadSuppressionRuleClient struct {
	fake *fakeconnect.FakeConnect
}

// New instantiates a new WorkloadSuppressionRuleClient for communication with a fake Connect API.
func New(fake *fakeconnect.FakeConnect) workloadsuppressionrulev1alpha1.WorkloadSuppressionRuleClient {
	return &fakeWorkloadSuppressionRuleClient{fake: fake}
}

func (c *fakeWorkloadSuppressionRuleClient) CreateWorkloadSuppressionRule(ctx context.Context, rule *workloadsuppressionrulepb.WorkloadSuppressionRule) (*workloadsuppressionrulepb.WorkloadSuppressionRule, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if err := c.fake.ValidateOrganization(rule.GetOrgId()); err != nil {
		return nil, err
	}
	if err := c.validateMatcher(rule); err != nil {
		return nil, err
	}

	rule = clone(rule)
	rule.Id = uuid.New().String()
	rule.CreatedAt = timestamppb.Now()
	c.fake.WorkloadSuppressionRules[rule.GetId()] = rule
	return clone(rule), nil
}

func (c *fakeWorkloadSuppressionRuleClient) DestroyWorkloadSuppressionRule(ctx context.Context, ruleID string) error {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.WorkloadSuppressionRules[ruleID]; !ok {
		return status.Error(codes.NotFound, "workload suppression rule not found")
	}
	delete(c.fake.WorkloadSuppressionRules, ruleID)
	return nil
}

func (c *fakeWorkloadSuppressionRuleClient) GetWorkloadSuppressionRule(ctx context.Context, ruleID string) (*workloadsuppressionrulepb.WorkloadSuppressionRule, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	rule, ok := c.fake.WorkloadSuppressionRules[ruleID]
	if !ok {
		return nil, status.Error(codes.NotFound, "workload suppression rule not found")
	}
	return clone(rule), nil
}

func (c *fakeWorkloadSuppressionRuleClient) ListWorkloadSuppressionRules(ctx context.Context, filter *workloadsuppressionrulesvcpb.ListWorkloadSuppressionRulesRequest_Filter) ([]*workloadsuppressionrulepb.WorkloadSuppressionRule, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	rules := []*workloadsuppressionrulepb.WorkloadSuppressionRule{}
	for _, rule := range c.fake.WorkloadSuppressionRules {
		if ruleMatches(rule, filter) {
			rules = append(rules, clone(rule))
		}
	}
	return rules, nil
}

func ruleMatches(rule *workloadsuppressionrulepb.WorkloadSuppressionRule, filter *workloadsuppressionrulesvcpb.ListWorkloadSuppressionRulesRequest_Filter) bool {
	if filter == nil {
		return true
	}
	if len(filter.GetOrgIds()) > 0 && !slices.Contains(filter.GetOrgIds(), rule.GetOrgId()) {
		return false
	}
	if filter.Enabled != nil && rule.GetEnabled() != filter.GetEnabled() {
		return false
	}
	return true
}

func (c *fakeWorkloadSuppressionRuleClient) UpdateWorkloadSuppressionRule(ctx context.Context, rule *workloadsuppressionrulepb.WorkloadSuppressionRule, updateMask *workloadsuppressionrulesvcpb.UpdateWorkloadSuppressionRuleRequest_UpdateMask) (*workloadsuppressionrulepb.WorkloadSuppressionRule, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	existing, ok := c.fake.WorkloadSuppressionRules[rule.GetId()]
	if !ok {
		return nil, status.Error(codes.NotFound, "workload suppression rule not found")
	}
	if err := c.validateMatcher(rule); err != nil {
		return nil, err
	}

	var new *workloadsuppressionrulepb.WorkloadSuppressionRule
	if updateMask == nil {
		new = clone(rule)
	} else {
		new = clone(existing)
		if updateMask.GetName() {
			new.Name = rule.GetName()
		}
		if updateMask.GetDescription() {
			new.Description = rule.GetDescription()
		}
		if updateMask.GetEnabled() {
			new.Enabled = rule.GetEnabled()
		}
		if updateMask.GetMatcher() {
			new.Matcher = rule.GetMatcher()
		}
	}
	new.Id = existing.GetId()
	new.OrgId = existing.GetOrgId()
	new.CreatedAt = existing.GetCreatedAt()
	new.LastUpdatedAt = timestamppb.Now()

	c.fake.WorkloadSuppressionRules[new.GetId()] = new
	return clone(new), nil
}

func (c *fakeWorkloadSuppressionRuleClient) validateMatcher(rule *workloadsuppressionrulepb.WorkloadSuppressionRule) error {
	pod := rule.GetKubernetesPod()
	if pod == nil {
		return nil
	}
	for _, trustZoneID := range pod.GetTrustZoneIds() {
		if err := c.fake.ValidateTrustZone(trustZoneID); err != nil {
			return err
		}
	}
	for _, clusterID := range pod.GetClusterIds() {
		if err := c.fake.ValidateCluster(clusterID); err != nil {
			return err
		}
	}
	return nil
}

func clone(rule *workloadsuppressionrulepb.WorkloadSuppressionRule) *workloadsuppressionrulepb.WorkloadSuppressionRule {
	return proto.Clone(rule).(*workloadsuppressionrulepb.WorkloadSuppressionRule)
}
