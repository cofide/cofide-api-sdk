// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"

	attestationpolicypb "github.com/cofide/cofide-api-sdk/gen/go/proto/attestation_policy/v1alpha1"
	attestationpolicysvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/attestation_policy_service/v1alpha1"
	"google.golang.org/grpc"
)

// AttestationPolicyClient is an interface for a gRPC client for the v1alpha1 version of the Connect AttestationPolicyService.
type AttestationPolicyClient interface {
	CreateAttestationPolicy(ctx context.Context, policy *attestationpolicypb.AttestationPolicy) (*attestationpolicypb.AttestationPolicy, error)
	DestroyAttestationPolicy(ctx context.Context, policyID string) error
	GetAttestationPolicy(ctx context.Context, policyID string) (*attestationpolicypb.AttestationPolicy, error)
	ListAttestationPolicies(ctx context.Context, filter *attestationpolicysvcpb.ListAttestationPoliciesRequest_Filter) ([]*attestationpolicypb.AttestationPolicy, error)
	UpdateAttestationPolicy(ctx context.Context, policy *attestationpolicypb.AttestationPolicy) (*attestationpolicypb.AttestationPolicy, error)
}

type attestationPolicyClient struct {
	client attestationpolicysvcpb.AttestationPolicyServiceClient
}

// New instantiates a new AttestationPolicyClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) AttestationPolicyClient {
	return &attestationPolicyClient{
		client: attestationpolicysvcpb.NewAttestationPolicyServiceClient(conn),
	}
}

func (c *attestationPolicyClient) CreateAttestationPolicy(ctx context.Context, policy *attestationpolicypb.AttestationPolicy) (*attestationpolicypb.AttestationPolicy, error) {
	resp, err := c.client.CreateAttestationPolicy(ctx, &attestationpolicysvcpb.CreateAttestationPolicyRequest{
		Policy: policy,
	})
	if err != nil {
		return nil, err
	}
	return resp.Policy, nil
}

func (c *attestationPolicyClient) DestroyAttestationPolicy(ctx context.Context, policyID string) error {
	_, err := c.client.DestroyAttestationPolicy(ctx, &attestationpolicysvcpb.DestroyAttestationPolicyRequest{
		PolicyId: &policyID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *attestationPolicyClient) GetAttestationPolicy(ctx context.Context, policyID string) (*attestationpolicypb.AttestationPolicy, error) {
	resp, err := c.client.GetAttestationPolicy(ctx, &attestationpolicysvcpb.GetAttestationPolicyRequest{
		PolicyId: &policyID,
	})
	if err != nil {
		return nil, err
	}
	return resp.Policy, nil
}

func (c *attestationPolicyClient) ListAttestationPolicies(ctx context.Context, filter *attestationpolicysvcpb.ListAttestationPoliciesRequest_Filter) ([]*attestationpolicypb.AttestationPolicy, error) {
	resp, err := c.client.ListAttestationPolicies(ctx, &attestationpolicysvcpb.ListAttestationPoliciesRequest{
		Filter: filter,
	})
	if err != nil {
		return nil, err
	}
	return resp.Policies, nil
}

func (c *attestationPolicyClient) UpdateAttestationPolicy(ctx context.Context, policy *attestationpolicypb.AttestationPolicy) (*attestationpolicypb.AttestationPolicy, error) {
	resp, err := c.client.UpdateAttestationPolicy(ctx, &attestationpolicysvcpb.UpdateAttestationPolicyRequest{
		Policy: policy,
	})
	if err != nil {
		return nil, err
	}
	return resp.Policy, nil
}
