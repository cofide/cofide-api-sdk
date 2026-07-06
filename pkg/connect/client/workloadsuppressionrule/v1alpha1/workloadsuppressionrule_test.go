// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"
	"testing"

	workloadsuppressionrulesvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/workload_suppression_rule_service/v1alpha1"
	workloadsuppressionrulepb "github.com/cofide/cofide-api-sdk/gen/go/proto/workload_suppression_rule/v1alpha1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	fakeWorkloadSuppressionRuleID   = "fake-wsr-id"
	fakeWorkloadSuppressionRuleName = "fake-wsr-name"
)

func TestWorkloadSuppressionRuleClient_implementsMethods(t *testing.T) {
	test.AssertClientImplementsService(t, &workloadSuppressionRuleClient{}, workloadsuppressionrulesvcpb.WorkloadSuppressionRuleService_ServiceDesc)
}

// TestWorkloadSuppressionRuleClient_Unimplemented tests WorkloadSuppressionRuleClient against an unimplemented server.
// This ensures that all errors returned are not wrapped can be converted to a gRPC Status using Status.Convert.
func TestWorkloadSuppressionRuleClient_Unimplemented(t *testing.T) {
	server := test.NewTestServer(t)
	workloadsuppressionrulesvcpb.RegisterWorkloadSuppressionRuleServiceServer(server.Server, &workloadsuppressionrulesvcpb.UnimplementedWorkloadSuppressionRuleServiceServer{})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	rule, err := client.CreateWorkloadSuppressionRule(ctx, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, rule)

	rule, err = client.GetWorkloadSuppressionRule(ctx, "")
	test.RequireUnimplemented(t, err)
	assert.Nil(t, rule)

	rules, err := client.ListWorkloadSuppressionRules(ctx, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, rules)

	rule, err = client.UpdateWorkloadSuppressionRule(ctx, nil, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, rule)

	err = client.DestroyWorkloadSuppressionRule(ctx, "")
	test.RequireUnimplemented(t, err)
}

func TestWorkloadSuppressionRuleClient(t *testing.T) {
	server := test.NewTestServer(t)
	workloadsuppressionrulesvcpb.RegisterWorkloadSuppressionRuleServiceServer(server.Server, &fakeWorkloadSuppressionRuleService{t: t})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	rule := fakeWorkloadSuppressionRule()

	createdRule, err := client.CreateWorkloadSuppressionRule(ctx, rule)
	require.NoError(t, err)
	assert.EqualExportedValues(t, rule, createdRule)

	gotRule, err := client.GetWorkloadSuppressionRule(ctx, rule.GetId())
	require.NoError(t, err)
	assert.EqualExportedValues(t, rule, gotRule)

	filter := &workloadsuppressionrulesvcpb.ListWorkloadSuppressionRulesRequest_Filter{OrgId: test.PtrOf(test.FakeOrganizationID)}
	rules, err := client.ListWorkloadSuppressionRules(ctx, filter)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*workloadsuppressionrulepb.WorkloadSuppressionRule{rule}, rules)

	updateMask := &workloadsuppressionrulesvcpb.UpdateWorkloadSuppressionRuleRequest_UpdateMask{Enabled: true}
	updatedRule, err := client.UpdateWorkloadSuppressionRule(ctx, rule, updateMask)
	require.NoError(t, err)
	assert.EqualExportedValues(t, rule, updatedRule)

	err = client.DestroyWorkloadSuppressionRule(ctx, rule.GetId())
	require.NoError(t, err)
}

// fakeWorkloadSuppressionRuleService provides a very simple fake WorkloadSuppressionRuleService implementation with canned responses.
type fakeWorkloadSuppressionRuleService struct {
	t *testing.T
}

func (f *fakeWorkloadSuppressionRuleService) CreateWorkloadSuppressionRule(ctx context.Context, req *workloadsuppressionrulesvcpb.CreateWorkloadSuppressionRuleRequest) (*workloadsuppressionrulesvcpb.CreateWorkloadSuppressionRuleResponse, error) {
	assert.EqualExportedValues(f.t, fakeWorkloadSuppressionRule(), req.WorkloadSuppressionRule)
	return &workloadsuppressionrulesvcpb.CreateWorkloadSuppressionRuleResponse{WorkloadSuppressionRule: req.WorkloadSuppressionRule}, nil
}

func (f *fakeWorkloadSuppressionRuleService) GetWorkloadSuppressionRule(ctx context.Context, req *workloadsuppressionrulesvcpb.GetWorkloadSuppressionRuleRequest) (*workloadsuppressionrulesvcpb.GetWorkloadSuppressionRuleResponse, error) {
	assert.Equal(f.t, fakeWorkloadSuppressionRuleID, req.GetWorkloadSuppressionRuleId())
	return &workloadsuppressionrulesvcpb.GetWorkloadSuppressionRuleResponse{WorkloadSuppressionRule: fakeWorkloadSuppressionRule()}, nil
}

func (f *fakeWorkloadSuppressionRuleService) ListWorkloadSuppressionRules(ctx context.Context, req *workloadsuppressionrulesvcpb.ListWorkloadSuppressionRulesRequest) (*workloadsuppressionrulesvcpb.ListWorkloadSuppressionRulesResponse, error) {
	assert.Equal(f.t, test.FakeOrganizationID, req.Filter.GetOrgId())
	rules := []*workloadsuppressionrulepb.WorkloadSuppressionRule{fakeWorkloadSuppressionRule()}
	return &workloadsuppressionrulesvcpb.ListWorkloadSuppressionRulesResponse{WorkloadSuppressionRules: rules}, nil
}

func (f *fakeWorkloadSuppressionRuleService) UpdateWorkloadSuppressionRule(ctx context.Context, req *workloadsuppressionrulesvcpb.UpdateWorkloadSuppressionRuleRequest) (*workloadsuppressionrulesvcpb.UpdateWorkloadSuppressionRuleResponse, error) {
	assert.EqualExportedValues(f.t, fakeWorkloadSuppressionRule(), req.WorkloadSuppressionRule)
	assert.True(f.t, req.UpdateMask.GetEnabled())
	return &workloadsuppressionrulesvcpb.UpdateWorkloadSuppressionRuleResponse{WorkloadSuppressionRule: fakeWorkloadSuppressionRule()}, nil
}

func (f *fakeWorkloadSuppressionRuleService) DestroyWorkloadSuppressionRule(ctx context.Context, req *workloadsuppressionrulesvcpb.DestroyWorkloadSuppressionRuleRequest) (*workloadsuppressionrulesvcpb.DestroyWorkloadSuppressionRuleResponse, error) {
	assert.Equal(f.t, fakeWorkloadSuppressionRuleID, req.GetWorkloadSuppressionRuleId())
	return &workloadsuppressionrulesvcpb.DestroyWorkloadSuppressionRuleResponse{}, nil
}

func fakeWorkloadSuppressionRule() *workloadsuppressionrulepb.WorkloadSuppressionRule {
	return &workloadsuppressionrulepb.WorkloadSuppressionRule{
		Id:      fakeWorkloadSuppressionRuleID,
		OrgId:   test.FakeOrganizationID,
		Name:    fakeWorkloadSuppressionRuleName,
		Enabled: true,
	}
}
