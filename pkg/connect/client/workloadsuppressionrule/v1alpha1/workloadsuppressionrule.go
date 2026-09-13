// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"

	workloadsuppressionrulesvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/workload_suppression_rule_service/v1alpha1"
	workloadsuppressionrulepb "github.com/cofide/cofide-api-sdk/gen/go/proto/workload_suppression_rule/v1alpha1"
	"google.golang.org/grpc"
)

// WorkloadSuppressionRuleClient is an interface for a gRPC client for the v1alpha1 version of the Connect WorkloadSuppressionRuleService.
type WorkloadSuppressionRuleClient interface {
	CreateWorkloadSuppressionRule(ctx context.Context, rule *workloadsuppressionrulepb.WorkloadSuppressionRule) (*workloadsuppressionrulepb.WorkloadSuppressionRule, error)
	DestroyWorkloadSuppressionRule(ctx context.Context, ruleID string) error
	GetWorkloadSuppressionRule(ctx context.Context, ruleID string) (*workloadsuppressionrulepb.WorkloadSuppressionRule, error)
	ListWorkloadSuppressionRules(ctx context.Context, filter *workloadsuppressionrulesvcpb.ListWorkloadSuppressionRulesRequest_Filter) ([]*workloadsuppressionrulepb.WorkloadSuppressionRule, error)
	UpdateWorkloadSuppressionRule(ctx context.Context, rule *workloadsuppressionrulepb.WorkloadSuppressionRule, updateMask *workloadsuppressionrulesvcpb.UpdateWorkloadSuppressionRuleRequest_UpdateMask) (*workloadsuppressionrulepb.WorkloadSuppressionRule, error)
}

type workloadSuppressionRuleClient struct {
	client workloadsuppressionrulesvcpb.WorkloadSuppressionRuleServiceClient
}

// New instantiates a new WorkloadSuppressionRuleClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) WorkloadSuppressionRuleClient {
	return &workloadSuppressionRuleClient{
		client: workloadsuppressionrulesvcpb.NewWorkloadSuppressionRuleServiceClient(conn),
	}
}

func (c *workloadSuppressionRuleClient) CreateWorkloadSuppressionRule(ctx context.Context, rule *workloadsuppressionrulepb.WorkloadSuppressionRule) (*workloadsuppressionrulepb.WorkloadSuppressionRule, error) {
	resp, err := c.client.CreateWorkloadSuppressionRule(ctx, &workloadsuppressionrulesvcpb.CreateWorkloadSuppressionRuleRequest{
		WorkloadSuppressionRule: rule,
	})
	if err != nil {
		return nil, err
	}
	return resp.WorkloadSuppressionRule, nil
}

func (c *workloadSuppressionRuleClient) DestroyWorkloadSuppressionRule(ctx context.Context, ruleID string) error {
	_, err := c.client.DestroyWorkloadSuppressionRule(ctx, &workloadsuppressionrulesvcpb.DestroyWorkloadSuppressionRuleRequest{
		WorkloadSuppressionRuleId: ruleID,
	})
	return err
}

func (c *workloadSuppressionRuleClient) GetWorkloadSuppressionRule(ctx context.Context, ruleID string) (*workloadsuppressionrulepb.WorkloadSuppressionRule, error) {
	resp, err := c.client.GetWorkloadSuppressionRule(ctx, &workloadsuppressionrulesvcpb.GetWorkloadSuppressionRuleRequest{
		WorkloadSuppressionRuleId: ruleID,
	})
	if err != nil {
		return nil, err
	}
	return resp.WorkloadSuppressionRule, nil
}

func (c *workloadSuppressionRuleClient) ListWorkloadSuppressionRules(ctx context.Context, filter *workloadsuppressionrulesvcpb.ListWorkloadSuppressionRulesRequest_Filter) ([]*workloadsuppressionrulepb.WorkloadSuppressionRule, error) {
	resp, err := c.client.ListWorkloadSuppressionRules(ctx, &workloadsuppressionrulesvcpb.ListWorkloadSuppressionRulesRequest{
		Filter: filter,
	})
	if err != nil {
		return nil, err
	}
	return resp.WorkloadSuppressionRules, nil
}

func (c *workloadSuppressionRuleClient) UpdateWorkloadSuppressionRule(ctx context.Context, rule *workloadsuppressionrulepb.WorkloadSuppressionRule, updateMask *workloadsuppressionrulesvcpb.UpdateWorkloadSuppressionRuleRequest_UpdateMask) (*workloadsuppressionrulepb.WorkloadSuppressionRule, error) {
	resp, err := c.client.UpdateWorkloadSuppressionRule(ctx, &workloadsuppressionrulesvcpb.UpdateWorkloadSuppressionRuleRequest{
		WorkloadSuppressionRule: rule,
		UpdateMask:              updateMask,
	})
	if err != nil {
		return nil, err
	}
	return resp.WorkloadSuppressionRule, nil
}
