// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	attestationpolicypb "github.com/cofide/cofide-api-sdk/gen/go/proto/attestation_policy/v1alpha1"
	attestationpolicysvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/attestation_policy_service/v1alpha1"
	policyv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/attestationpolicy/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type fakeAttestationPolicyClient struct {
	fake *fakeconnect.FakeConnect
}

// New instantiates a new AttestationPolicyClient for communication with a fake Connect API.
func New(fake *fakeconnect.FakeConnect) policyv1alpha1.AttestationPolicyClient {
	return &fakeAttestationPolicyClient{fake: fake}
}

func (c *fakeAttestationPolicyClient) CreateAttestationPolicy(ctx context.Context, policy *attestationpolicypb.AttestationPolicy) (*attestationpolicypb.AttestationPolicy, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	id := uuid.New().String()
	policy = clone(policy)
	policy.Id = &id
	c.fake.AttestationPolicies[policy.GetId()] = policy

	return clone(policy), nil
}

func (c *fakeAttestationPolicyClient) DestroyAttestationPolicy(ctx context.Context, policyID string) error {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.AttestationPolicies[policyID]; !ok {
		return status.Error(codes.NotFound, "policy not found")
	}
	delete(c.fake.AttestationPolicies, policyID)
	return nil
}

func (c *fakeAttestationPolicyClient) GetAttestationPolicy(ctx context.Context, policyID string) (*attestationpolicypb.AttestationPolicy, error) {
	policy, ok := c.fake.AttestationPolicies[policyID]
	if !ok {
		return nil, status.Error(codes.NotFound, "policy not found")
	}
	return clone(policy), nil
}

func (c *fakeAttestationPolicyClient) ListAttestationPolicies(ctx context.Context, filter *attestationpolicysvcpb.ListAttestationPoliciesRequest_Filter) ([]*attestationpolicypb.AttestationPolicy, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	policies := []*attestationpolicypb.AttestationPolicy{}
	for _, policy := range c.fake.AttestationPolicies {
		if c.policyMatches(policy, filter) {
			policies = append(policies, clone(policy))
		}
	}
	return policies, nil
}

func (c *fakeAttestationPolicyClient) policyMatches(policy *attestationpolicypb.AttestationPolicy, filter *attestationpolicysvcpb.ListAttestationPoliciesRequest_Filter) bool {
	if filter == nil {
		return true
	}
	if filter.Name != nil && policy.GetName() != *filter.Name {
		return false
	}
	if filter.OrgId != nil && policy.GetOrgId() != *filter.OrgId {
		return false
	}
	if filter.TrustZoneId != nil {
		for _, apb := range c.fake.APBindings {
			if apb.GetTrustZoneId() == *filter.TrustZoneId {
				// If we find an APB with this trust_zone_id
				// then exit the search through the APBs
				break
			}
			// No trust_zone_id matches found => return false
			return false
		}
	}
	return true
}

func (c *fakeAttestationPolicyClient) UpdateAttestationPolicy(ctx context.Context, policy *attestationpolicypb.AttestationPolicy) (*attestationpolicypb.AttestationPolicy, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.AttestationPolicies[policy.GetId()]; !ok {
		return nil, status.Error(codes.NotFound, "")
	}
	c.fake.AttestationPolicies[policy.GetId()] = clone(policy)
	return clone(policy), nil
}

func clone(policy *attestationpolicypb.AttestationPolicy) *attestationpolicypb.AttestationPolicy {
	return proto.Clone(policy).(*attestationpolicypb.AttestationPolicy)
}
