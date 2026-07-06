// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	workloadsuppressionrulesvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/workload_suppression_rule_service/v1alpha1"
	workloadsuppressionrulepb "github.com/cofide/cofide-api-sdk/gen/go/proto/workload_suppression_rule/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	workloadsuppressionrulev1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/workloadsuppressionrule/v1alpha1"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
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

	if err := c.validateScope(rule); err != nil {
		return nil, err
	}
	id := uuid.New().String()
	rule = clone(rule)
	rule.Id = id
	c.fake.WorkloadSuppressionRules[rule.GetId()] = rule
	return clone(rule), nil
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

func (c *fakeWorkloadSuppressionRuleClient) UpdateWorkloadSuppressionRule(ctx context.Context, rule *workloadsuppressionrulepb.WorkloadSuppressionRule, updateMask *workloadsuppressionrulesvcpb.UpdateWorkloadSuppressionRuleRequest_UpdateMask) (*workloadsuppressionrulepb.WorkloadSuppressionRule, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.WorkloadSuppressionRules[rule.GetId()]; !ok {
		return nil, status.Error(codes.NotFound, "workload suppression rule not found")
	}
	if err := c.validateScope(rule); err != nil {
		return nil, err
	}
	c.fake.WorkloadSuppressionRules[rule.GetId()] = clone(rule)
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

func (c *fakeWorkloadSuppressionRuleClient) validateScope(rule *workloadsuppressionrulepb.WorkloadSuppressionRule) error {
	if err := c.fake.ValidateOrganization(rule.GetOrgId()); err != nil {
		return err
	}
	if rule.TrustZoneId != nil {
		if err := c.fake.ValidateTrustZone(rule.GetTrustZoneId()); err != nil {
			return err
		}
	}
	if rule.ClusterId != nil {
		if err := c.fake.ValidateCluster(rule.GetClusterId()); err != nil {
			return err
		}
	}
	return nil
}

func ruleMatches(rule *workloadsuppressionrulepb.WorkloadSuppressionRule, filter *workloadsuppressionrulesvcpb.ListWorkloadSuppressionRulesRequest_Filter) bool {
	if filter == nil {
		return true
	}
	if filter.OrgId != nil && rule.GetOrgId() != *filter.OrgId {
		return false
	}
	if filter.TrustZoneId != nil && rule.GetTrustZoneId() != *filter.TrustZoneId {
		return false
	}
	if filter.ClusterId != nil && rule.GetClusterId() != *filter.ClusterId {
		return false
	}
	if filter.Enabled != nil && rule.GetEnabled() != *filter.Enabled {
		return false
	}
	return true
}

func clone(rule *workloadsuppressionrulepb.WorkloadSuppressionRule) *workloadsuppressionrulepb.WorkloadSuppressionRule {
	return proto.Clone(rule).(*workloadsuppressionrulepb.WorkloadSuppressionRule)
}
