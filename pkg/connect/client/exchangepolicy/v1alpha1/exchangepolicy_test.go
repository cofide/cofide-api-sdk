// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"
	"testing"

	exchangepolicysvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/exchange_policy_service/v1alpha1"
	exchangepolicypb "github.com/cofide/cofide-api-sdk/gen/go/proto/exchange_policy/v1alpha1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	fakeExchangePolicyID   = "fake-exchange-policy-id"
	fakeExchangePolicyName = "fake-exchange-policy-name"
	fakeTrustZoneID        = "fake-tz-id"
)

func TestExchangePolicyClient_implementsMethods(t *testing.T) {
	test.AssertClientImplementsService(t, &exchangePolicyClient{}, exchangepolicysvcpb.ExchangePolicyService_ServiceDesc)
}

// TestExchangePolicyClient_Unimplemented tests ExchangePolicyClient against an unimplemented server.
// This ensures that all errors returned are not wrapped and can be converted to a gRPC Status using Status.Convert.
func TestExchangePolicyClient_Unimplemented(t *testing.T) {
	server := test.NewTestServer(t)
	exchangepolicysvcpb.RegisterExchangePolicyServiceServer(server.Server, &exchangepolicysvcpb.UnimplementedExchangePolicyServiceServer{})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	policy, err := client.CreateExchangePolicy(ctx, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, policy)

	err = client.DestroyExchangePolicy(ctx, "")
	test.RequireUnimplemented(t, err)

	policy, err = client.GetExchangePolicy(ctx, "")
	test.RequireUnimplemented(t, err)
	assert.Nil(t, policy)

	policies, err := client.ListExchangePolicies(ctx, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, policies)

	policy, err = client.UpdateExchangePolicy(ctx, nil, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, policy)
}

func TestExchangePolicyClient(t *testing.T) {
	server := test.NewTestServer(t)
	exchangepolicysvcpb.RegisterExchangePolicyServiceServer(server.Server, &fakeExchangePolicyService{t: t})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	policy := fakeExchangePolicy()

	createdPolicy, err := client.CreateExchangePolicy(ctx, policy)
	require.NoError(t, err)
	assert.EqualExportedValues(t, policy, createdPolicy)

	err = client.DestroyExchangePolicy(ctx, policy.GetId())
	require.NoError(t, err)

	gotPolicy, err := client.GetExchangePolicy(ctx, policy.GetId())
	require.NoError(t, err)
	assert.Equal(t, policy.GetId(), gotPolicy.GetId())

	filter := &exchangepolicysvcpb.ListExchangePoliciesRequest_Filter{TrustZoneId: fakeTrustZoneID}
	policies, err := client.ListExchangePolicies(ctx, filter)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*exchangepolicypb.ExchangePolicy{fakeExchangePolicy()}, policies)

	updatedPolicy, err := client.UpdateExchangePolicy(ctx, policy, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, policy, updatedPolicy)
}

// fakeExchangePolicyService provides a simple fake ExchangePolicyService implementation with canned responses.
type fakeExchangePolicyService struct {
	t *testing.T
}

func (f *fakeExchangePolicyService) CreateExchangePolicy(ctx context.Context, req *exchangepolicysvcpb.CreateExchangePolicyRequest) (*exchangepolicysvcpb.CreateExchangePolicyResponse, error) {
	assert.EqualExportedValues(f.t, fakeExchangePolicy(), req.ExchangePolicy)
	return &exchangepolicysvcpb.CreateExchangePolicyResponse{ExchangePolicy: req.ExchangePolicy}, nil
}

func (f *fakeExchangePolicyService) DestroyExchangePolicy(ctx context.Context, req *exchangepolicysvcpb.DestroyExchangePolicyRequest) (*exchangepolicysvcpb.DestroyExchangePolicyResponse, error) {
	assert.Equal(f.t, fakeExchangePolicyID, req.GetExchangePolicyId())
	return &exchangepolicysvcpb.DestroyExchangePolicyResponse{}, nil
}

func (f *fakeExchangePolicyService) GetExchangePolicy(ctx context.Context, req *exchangepolicysvcpb.GetExchangePolicyRequest) (*exchangepolicysvcpb.GetExchangePolicyResponse, error) {
	assert.Equal(f.t, fakeExchangePolicyID, req.GetExchangePolicyId())
	return &exchangepolicysvcpb.GetExchangePolicyResponse{ExchangePolicy: fakeExchangePolicy()}, nil
}

func (f *fakeExchangePolicyService) ListExchangePolicies(ctx context.Context, req *exchangepolicysvcpb.ListExchangePoliciesRequest) (*exchangepolicysvcpb.ListExchangePoliciesResponse, error) {
	assert.Equal(f.t, fakeTrustZoneID, req.Filter.GetTrustZoneId())
	return &exchangepolicysvcpb.ListExchangePoliciesResponse{ExchangePolicies: []*exchangepolicypb.ExchangePolicy{fakeExchangePolicy()}}, nil
}

func (f *fakeExchangePolicyService) UpdateExchangePolicy(ctx context.Context, req *exchangepolicysvcpb.UpdateExchangePolicyRequest) (*exchangepolicysvcpb.UpdateExchangePolicyResponse, error) {
	assert.EqualExportedValues(f.t, fakeExchangePolicy(), req.ExchangePolicy)
	return &exchangepolicysvcpb.UpdateExchangePolicyResponse{ExchangePolicy: fakeExchangePolicy()}, nil
}

func fakeExchangePolicy() *exchangepolicypb.ExchangePolicy {
	return &exchangepolicypb.ExchangePolicy{
		Id:          fakeExchangePolicyID,
		Name:        fakeExchangePolicyName,
		TrustZoneId: fakeTrustZoneID,
	}
}
