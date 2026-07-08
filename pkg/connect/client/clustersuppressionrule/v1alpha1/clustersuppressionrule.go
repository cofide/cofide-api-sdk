// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"

	clustersuppressionrulesvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/cluster_suppression_rule_service/v1alpha1"
	clustersuppressionrulepb "github.com/cofide/cofide-api-sdk/gen/go/proto/cluster_suppression_rule/v1alpha1"
	"google.golang.org/grpc"
)

// ClusterSuppressionRuleClient is an interface for a gRPC client for the v1alpha1 version of the Connect ClusterSuppressionRuleService.
type ClusterSuppressionRuleClient interface {
	CreateClusterSuppressionRule(ctx context.Context, rule *clustersuppressionrulepb.ClusterSuppressionRule) (*clustersuppressionrulepb.ClusterSuppressionRule, error)
	GetClusterSuppressionRule(ctx context.Context, ruleID string) (*clustersuppressionrulepb.ClusterSuppressionRule, error)
	ListClusterSuppressionRules(ctx context.Context, filter *clustersuppressionrulesvcpb.ListClusterSuppressionRulesRequest_Filter) ([]*clustersuppressionrulepb.ClusterSuppressionRule, error)
	UpdateClusterSuppressionRule(ctx context.Context, rule *clustersuppressionrulepb.ClusterSuppressionRule, updateMask *clustersuppressionrulesvcpb.UpdateClusterSuppressionRuleRequest_UpdateMask) (*clustersuppressionrulepb.ClusterSuppressionRule, error)
	DestroyClusterSuppressionRule(ctx context.Context, ruleID string) error
}

type clusterSuppressionRuleClient struct {
	client clustersuppressionrulesvcpb.ClusterSuppressionRuleServiceClient
}

// New instantiates a new ClusterSuppressionRuleClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) ClusterSuppressionRuleClient {
	return &clusterSuppressionRuleClient{
		client: clustersuppressionrulesvcpb.NewClusterSuppressionRuleServiceClient(conn),
	}
}

func (c *clusterSuppressionRuleClient) CreateClusterSuppressionRule(ctx context.Context, rule *clustersuppressionrulepb.ClusterSuppressionRule) (*clustersuppressionrulepb.ClusterSuppressionRule, error) {
	resp, err := c.client.CreateClusterSuppressionRule(ctx, &clustersuppressionrulesvcpb.CreateClusterSuppressionRuleRequest{
		ClusterSuppressionRule: rule,
	})
	if err != nil {
		return nil, err
	}
	return resp.ClusterSuppressionRule, nil
}

func (c *clusterSuppressionRuleClient) GetClusterSuppressionRule(ctx context.Context, ruleID string) (*clustersuppressionrulepb.ClusterSuppressionRule, error) {
	resp, err := c.client.GetClusterSuppressionRule(ctx, &clustersuppressionrulesvcpb.GetClusterSuppressionRuleRequest{
		ClusterSuppressionRuleId: ruleID,
	})
	if err != nil {
		return nil, err
	}
	return resp.ClusterSuppressionRule, nil
}

func (c *clusterSuppressionRuleClient) ListClusterSuppressionRules(ctx context.Context, filter *clustersuppressionrulesvcpb.ListClusterSuppressionRulesRequest_Filter) ([]*clustersuppressionrulepb.ClusterSuppressionRule, error) {
	resp, err := c.client.ListClusterSuppressionRules(ctx, &clustersuppressionrulesvcpb.ListClusterSuppressionRulesRequest{
		Filter: filter,
	})
	if err != nil {
		return nil, err
	}
	return resp.ClusterSuppressionRules, nil
}

func (c *clusterSuppressionRuleClient) UpdateClusterSuppressionRule(ctx context.Context, rule *clustersuppressionrulepb.ClusterSuppressionRule, updateMask *clustersuppressionrulesvcpb.UpdateClusterSuppressionRuleRequest_UpdateMask) (*clustersuppressionrulepb.ClusterSuppressionRule, error) {
	resp, err := c.client.UpdateClusterSuppressionRule(ctx, &clustersuppressionrulesvcpb.UpdateClusterSuppressionRuleRequest{
		ClusterSuppressionRule: rule,
		UpdateMask:             updateMask,
	})
	if err != nil {
		return nil, err
	}
	return resp.ClusterSuppressionRule, nil
}

func (c *clusterSuppressionRuleClient) DestroyClusterSuppressionRule(ctx context.Context, ruleID string) error {
	_, err := c.client.DestroyClusterSuppressionRule(ctx, &clustersuppressionrulesvcpb.DestroyClusterSuppressionRuleRequest{
		ClusterSuppressionRuleId: ruleID,
	})
	return err
}
