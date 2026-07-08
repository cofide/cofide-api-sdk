// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	clustersuppressionrulesvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/cluster_suppression_rule_service/v1alpha1"
	clustersuppressionrulepb "github.com/cofide/cofide-api-sdk/gen/go/proto/cluster_suppression_rule/v1alpha1"
	clustersuppressionrulev1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/clustersuppressionrule/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type fakeClusterSuppressionRuleClient struct {
	fake *fakeconnect.FakeConnect
}

// New instantiates a new ClusterSuppressionRuleClient for communication with a fake Connect API.
func New(fake *fakeconnect.FakeConnect) clustersuppressionrulev1alpha1.ClusterSuppressionRuleClient {
	return &fakeClusterSuppressionRuleClient{fake: fake}
}

func (c *fakeClusterSuppressionRuleClient) CreateClusterSuppressionRule(ctx context.Context, rule *clustersuppressionrulepb.ClusterSuppressionRule) (*clustersuppressionrulepb.ClusterSuppressionRule, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if err := c.validateScope(rule); err != nil {
		return nil, err
	}
	id := uuid.New().String()
	rule = clone(rule)
	rule.Id = id
	c.fake.ClusterSuppressionRules[rule.GetId()] = rule
	return clone(rule), nil
}

func (c *fakeClusterSuppressionRuleClient) GetClusterSuppressionRule(ctx context.Context, ruleID string) (*clustersuppressionrulepb.ClusterSuppressionRule, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	rule, ok := c.fake.ClusterSuppressionRules[ruleID]
	if !ok {
		return nil, status.Error(codes.NotFound, "cluster suppression rule not found")
	}
	return clone(rule), nil
}

func (c *fakeClusterSuppressionRuleClient) ListClusterSuppressionRules(ctx context.Context, filter *clustersuppressionrulesvcpb.ListClusterSuppressionRulesRequest_Filter) ([]*clustersuppressionrulepb.ClusterSuppressionRule, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	rules := []*clustersuppressionrulepb.ClusterSuppressionRule{}
	for _, rule := range c.fake.ClusterSuppressionRules {
		if ruleMatches(rule, filter) {
			rules = append(rules, clone(rule))
		}
	}
	return rules, nil
}

func (c *fakeClusterSuppressionRuleClient) UpdateClusterSuppressionRule(ctx context.Context, rule *clustersuppressionrulepb.ClusterSuppressionRule, updateMask *clustersuppressionrulesvcpb.UpdateClusterSuppressionRuleRequest_UpdateMask) (*clustersuppressionrulepb.ClusterSuppressionRule, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.ClusterSuppressionRules[rule.GetId()]; !ok {
		return nil, status.Error(codes.NotFound, "cluster suppression rule not found")
	}
	if err := c.validateScope(rule); err != nil {
		return nil, err
	}
	c.fake.ClusterSuppressionRules[rule.GetId()] = clone(rule)
	return clone(rule), nil
}

func (c *fakeClusterSuppressionRuleClient) DestroyClusterSuppressionRule(ctx context.Context, ruleID string) error {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.ClusterSuppressionRules[ruleID]; !ok {
		return status.Error(codes.NotFound, "cluster suppression rule not found")
	}
	delete(c.fake.ClusterSuppressionRules, ruleID)
	return nil
}

func (c *fakeClusterSuppressionRuleClient) validateScope(rule *clustersuppressionrulepb.ClusterSuppressionRule) error {
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

func ruleMatches(rule *clustersuppressionrulepb.ClusterSuppressionRule, filter *clustersuppressionrulesvcpb.ListClusterSuppressionRulesRequest_Filter) bool {
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

func clone(rule *clustersuppressionrulepb.ClusterSuppressionRule) *clustersuppressionrulepb.ClusterSuppressionRule {
	return proto.Clone(rule).(*clustersuppressionrulepb.ClusterSuppressionRule)
}
