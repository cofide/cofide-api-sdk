// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"
	"testing"

	attestationpolicypb "github.com/cofide/cofide-api-sdk/gen/go/proto/attestation_policy/v1alpha1"
	attestationpolicysvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/attestation_policy_service/v1alpha1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	fakePolicyID   = "fake-policy-id"
	fakePolicyName = "fake-policy-name"
)

func TestAttestationPolicyClient_implementsMethods(t *testing.T) {
	test.AssertClientImplementsService(t, &attestationPolicyClient{}, attestationpolicysvcpb.AttestationPolicyService_ServiceDesc)
}

// TestAttestationPolicyClient_Unimplemented tests AttestationPolicyClient against an unimplemented server.
// This ensures that all errors returned are not wrapped can be converted to a gRPC Status using Status.Convert.
func TestAttestationPolicyClient_Unimplemented(t *testing.T) {
	server := test.NewTestServer(t)
	attestationpolicysvcpb.RegisterAttestationPolicyServiceServer(server.Server, &attestationpolicysvcpb.UnimplementedAttestationPolicyServiceServer{})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	token, err := client.CreateAttestationPolicy(ctx, nil)
	test.RequireUnimplemented(t, err)
	assert.Empty(t, token)

	err = client.DestroyAttestationPolicy(ctx, "")
	test.RequireUnimplemented(t, err)

	policy, err := client.GetAttestationPolicy(ctx, "")
	test.RequireUnimplemented(t, err)
	assert.Nil(t, policy)

	policies, err := client.ListAttestationPolicies(ctx, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, policies)

	policy, err = client.UpdateAttestationPolicy(ctx, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, policy)
}

func TestAttestationPolicyClient(t *testing.T) {
	server := test.NewTestServer(t)
	attestationpolicysvcpb.RegisterAttestationPolicyServiceServer(server.Server, &fakeAttestationPolicyService{t: t})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	policy := fakePolicy()

	createdPolicy, err := client.CreateAttestationPolicy(ctx, policy)
	require.NoError(t, err)
	assert.EqualExportedValues(t, policy, createdPolicy)

	err = client.DestroyAttestationPolicy(ctx, policy.GetId())
	require.NoError(t, err)

	gotPolicy, err := client.GetAttestationPolicy(ctx, policy.GetId())
	require.NoError(t, err)
	assert.Equal(t, policy.GetId(), gotPolicy.GetId())

	filter := &attestationpolicysvcpb.ListAttestationPoliciesRequest_Filter{Name: test.PtrOf(fakePolicyName)}
	policies, err := client.ListAttestationPolicies(ctx, filter)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*attestationpolicypb.AttestationPolicy{fakePolicy()}, policies)

	updatedPolicy, err := client.UpdateAttestationPolicy(ctx, policy)
	require.NoError(t, err)
	assert.EqualExportedValues(t, policy, updatedPolicy)
}

// fakeAttestationPolicyService provides a very simple fake AttestationPolicyService implementation with canned responses.
type fakeAttestationPolicyService struct {
	t *testing.T
}

func (f *fakeAttestationPolicyService) CreateAttestationPolicy(ctx context.Context, req *attestationpolicysvcpb.CreateAttestationPolicyRequest) (*attestationpolicysvcpb.CreateAttestationPolicyResponse, error) {
	assert.EqualExportedValues(f.t, fakePolicy(), req.Policy)
	return &attestationpolicysvcpb.CreateAttestationPolicyResponse{Policy: req.Policy}, nil
}

func (f *fakeAttestationPolicyService) DestroyAttestationPolicy(ctx context.Context, req *attestationpolicysvcpb.DestroyAttestationPolicyRequest) (*attestationpolicysvcpb.DestroyAttestationPolicyResponse, error) {
	assert.Equal(f.t, fakePolicyID, req.GetPolicyId())
	return &attestationpolicysvcpb.DestroyAttestationPolicyResponse{}, nil
}

func (f *fakeAttestationPolicyService) GetAttestationPolicy(ctx context.Context, req *attestationpolicysvcpb.GetAttestationPolicyRequest) (*attestationpolicysvcpb.GetAttestationPolicyResponse, error) {
	assert.Equal(f.t, fakePolicyID, req.GetPolicyId())
	return &attestationpolicysvcpb.GetAttestationPolicyResponse{Policy: fakePolicy()}, nil
}

func (f *fakeAttestationPolicyService) ListAttestationPolicies(ctx context.Context, req *attestationpolicysvcpb.ListAttestationPoliciesRequest) (*attestationpolicysvcpb.ListAttestationPoliciesResponse, error) {
	assert.Equal(f.t, fakePolicyName, req.Filter.GetName())
	policies := []*attestationpolicypb.AttestationPolicy{fakePolicy()}
	return &attestationpolicysvcpb.ListAttestationPoliciesResponse{Policies: policies}, nil
}

func (f *fakeAttestationPolicyService) UpdateAttestationPolicy(ctx context.Context, req *attestationpolicysvcpb.UpdateAttestationPolicyRequest) (*attestationpolicysvcpb.UpdateAttestationPolicyResponse, error) {
	assert.EqualExportedValues(f.t, fakePolicy(), req.Policy)
	return &attestationpolicysvcpb.UpdateAttestationPolicyResponse{Policy: fakePolicy()}, nil
}

func fakePolicy() *attestationpolicypb.AttestationPolicy {
	return &attestationpolicypb.AttestationPolicy{
		Id:   test.PtrOf(fakePolicyID),
		Name: fakePolicyName,
	}
}
