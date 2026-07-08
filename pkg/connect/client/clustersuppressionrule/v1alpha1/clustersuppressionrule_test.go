// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"
	"testing"

	clustersuppressionrulesvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/cluster_suppression_rule_service/v1alpha1"
	clustersuppressionrulepb "github.com/cofide/cofide-api-sdk/gen/go/proto/cluster_suppression_rule/v1alpha1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	fakeClusterSuppressionRuleID   = "fake-csr-id"
	fakeClusterSuppressionRuleName = "fake-csr-name"
)

func TestClusterSuppressionRuleClient_implementsMethods(t *testing.T) {
	test.AssertClientImplementsService(t, &clusterSuppressionRuleClient{}, clustersuppressionrulesvcpb.ClusterSuppressionRuleService_ServiceDesc)
}

// TestClusterSuppressionRuleClient_Unimplemented tests ClusterSuppressionRuleClient against an unimplemented server.
// This ensures that all errors returned are not wrapped can be converted to a gRPC Status using Status.Convert.
func TestClusterSuppressionRuleClient_Unimplemented(t *testing.T) {
	server := test.NewTestServer(t)
	clustersuppressionrulesvcpb.RegisterClusterSuppressionRuleServiceServer(server.Server, &clustersuppressionrulesvcpb.UnimplementedClusterSuppressionRuleServiceServer{})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	rule, err := client.CreateClusterSuppressionRule(ctx, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, rule)

	rule, err = client.GetClusterSuppressionRule(ctx, "")
	test.RequireUnimplemented(t, err)
	assert.Nil(t, rule)

	rules, err := client.ListClusterSuppressionRules(ctx, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, rules)

	rule, err = client.UpdateClusterSuppressionRule(ctx, nil, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, rule)

	err = client.DestroyClusterSuppressionRule(ctx, "")
	test.RequireUnimplemented(t, err)
}

func TestClusterSuppressionRuleClient(t *testing.T) {
	server := test.NewTestServer(t)
	clustersuppressionrulesvcpb.RegisterClusterSuppressionRuleServiceServer(server.Server, &fakeClusterSuppressionRuleService{t: t})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	rule := fakeClusterSuppressionRule()

	createdRule, err := client.CreateClusterSuppressionRule(ctx, rule)
	require.NoError(t, err)
	assert.EqualExportedValues(t, rule, createdRule)

	gotRule, err := client.GetClusterSuppressionRule(ctx, rule.GetId())
	require.NoError(t, err)
	assert.EqualExportedValues(t, rule, gotRule)

	filter := &clustersuppressionrulesvcpb.ListClusterSuppressionRulesRequest_Filter{OrgId: test.PtrOf(test.FakeOrganizationID)}
	rules, err := client.ListClusterSuppressionRules(ctx, filter)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*clustersuppressionrulepb.ClusterSuppressionRule{rule}, rules)

	updateMask := &clustersuppressionrulesvcpb.UpdateClusterSuppressionRuleRequest_UpdateMask{Enabled: true}
	updatedRule, err := client.UpdateClusterSuppressionRule(ctx, rule, updateMask)
	require.NoError(t, err)
	assert.EqualExportedValues(t, rule, updatedRule)

	err = client.DestroyClusterSuppressionRule(ctx, rule.GetId())
	require.NoError(t, err)
}

// fakeClusterSuppressionRuleService provides a very simple fake ClusterSuppressionRuleService implementation with canned responses.
type fakeClusterSuppressionRuleService struct {
	t *testing.T
}

func (f *fakeClusterSuppressionRuleService) CreateClusterSuppressionRule(ctx context.Context, req *clustersuppressionrulesvcpb.CreateClusterSuppressionRuleRequest) (*clustersuppressionrulesvcpb.CreateClusterSuppressionRuleResponse, error) {
	assert.EqualExportedValues(f.t, fakeClusterSuppressionRule(), req.ClusterSuppressionRule)
	return &clustersuppressionrulesvcpb.CreateClusterSuppressionRuleResponse{ClusterSuppressionRule: req.ClusterSuppressionRule}, nil
}

func (f *fakeClusterSuppressionRuleService) GetClusterSuppressionRule(ctx context.Context, req *clustersuppressionrulesvcpb.GetClusterSuppressionRuleRequest) (*clustersuppressionrulesvcpb.GetClusterSuppressionRuleResponse, error) {
	assert.Equal(f.t, fakeClusterSuppressionRuleID, req.GetClusterSuppressionRuleId())
	return &clustersuppressionrulesvcpb.GetClusterSuppressionRuleResponse{ClusterSuppressionRule: fakeClusterSuppressionRule()}, nil
}

func (f *fakeClusterSuppressionRuleService) ListClusterSuppressionRules(ctx context.Context, req *clustersuppressionrulesvcpb.ListClusterSuppressionRulesRequest) (*clustersuppressionrulesvcpb.ListClusterSuppressionRulesResponse, error) {
	assert.Equal(f.t, test.FakeOrganizationID, req.Filter.GetOrgId())
	rules := []*clustersuppressionrulepb.ClusterSuppressionRule{fakeClusterSuppressionRule()}
	return &clustersuppressionrulesvcpb.ListClusterSuppressionRulesResponse{ClusterSuppressionRules: rules}, nil
}

func (f *fakeClusterSuppressionRuleService) UpdateClusterSuppressionRule(ctx context.Context, req *clustersuppressionrulesvcpb.UpdateClusterSuppressionRuleRequest) (*clustersuppressionrulesvcpb.UpdateClusterSuppressionRuleResponse, error) {
	assert.EqualExportedValues(f.t, fakeClusterSuppressionRule(), req.ClusterSuppressionRule)
	assert.True(f.t, req.UpdateMask.GetEnabled())
	return &clustersuppressionrulesvcpb.UpdateClusterSuppressionRuleResponse{ClusterSuppressionRule: fakeClusterSuppressionRule()}, nil
}

func (f *fakeClusterSuppressionRuleService) DestroyClusterSuppressionRule(ctx context.Context, req *clustersuppressionrulesvcpb.DestroyClusterSuppressionRuleRequest) (*clustersuppressionrulesvcpb.DestroyClusterSuppressionRuleResponse, error) {
	assert.Equal(f.t, fakeClusterSuppressionRuleID, req.GetClusterSuppressionRuleId())
	return &clustersuppressionrulesvcpb.DestroyClusterSuppressionRuleResponse{}, nil
}

func fakeClusterSuppressionRule() *clustersuppressionrulepb.ClusterSuppressionRule {
	return &clustersuppressionrulepb.ClusterSuppressionRule{
		Id:      fakeClusterSuppressionRuleID,
		OrgId:   test.FakeOrganizationID,
		Name:    fakeClusterSuppressionRuleName,
		Enabled: true,
	}
}
