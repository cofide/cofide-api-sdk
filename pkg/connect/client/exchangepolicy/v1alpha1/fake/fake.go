// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	exchangepolicypb "github.com/cofide/cofide-api-sdk/gen/go/proto/exchange_policy/v1alpha1"
	exchangepolicysvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/exchange_policy_service/v1alpha1"
	exchangepolicyv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/exchangepolicy/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type fakeExchangePolicyClient struct {
	fake *fakeconnect.FakeConnect
}

// New instantiates a new ExchangePolicyClient for communication with a fake Connect API.
func New(fake *fakeconnect.FakeConnect) exchangepolicyv1alpha1.ExchangePolicyClient {
	return &fakeExchangePolicyClient{fake: fake}
}

func (c *fakeExchangePolicyClient) CreateExchangePolicy(ctx context.Context, policy *exchangepolicypb.ExchangePolicy) (*exchangepolicypb.ExchangePolicy, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if err := c.fake.ValidateTrustZone(policy.GetTrustZoneId()); err != nil {
		return nil, err
	}
	id := uuid.New().String()
	policy = clone(policy)
	policy.Id = &id
	c.fake.ExchangePolicies[policy.GetId()] = policy
	return clone(policy), nil
}

func (c *fakeExchangePolicyClient) DestroyExchangePolicy(ctx context.Context, policyID string) error {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.ExchangePolicies[policyID]; !ok {
		return status.Error(codes.NotFound, "exchange policy not found")
	}
	delete(c.fake.ExchangePolicies, policyID)
	return nil
}

func (c *fakeExchangePolicyClient) GetExchangePolicy(ctx context.Context, policyID string) (*exchangepolicypb.ExchangePolicy, error) {
	policy, ok := c.fake.ExchangePolicies[policyID]
	if !ok {
		return nil, status.Error(codes.NotFound, "exchange policy not found")
	}
	return clone(policy), nil
}

func (c *fakeExchangePolicyClient) ListExchangePolicies(ctx context.Context, filter *exchangepolicysvcpb.ListExchangePoliciesRequest_Filter) ([]*exchangepolicypb.ExchangePolicy, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	policies := []*exchangepolicypb.ExchangePolicy{}
	for _, policy := range c.fake.ExchangePolicies {
		if policyMatches(policy, filter) {
			policies = append(policies, clone(policy))
		}
	}
	return policies, nil
}

func policyMatches(policy *exchangepolicypb.ExchangePolicy, filter *exchangepolicysvcpb.ListExchangePoliciesRequest_Filter) bool {
	if filter == nil {
		return true
	}
	if filter.OrgId != nil && policy.GetOrgId() != *filter.OrgId {
		return false
	}
	if filter.TrustZoneId != nil && policy.GetTrustZoneId() != *filter.TrustZoneId {
		return false
	}
	if filter.Name != nil && policy.GetName() != *filter.Name {
		return false
	}
	return true
}

func (c *fakeExchangePolicyClient) UpdateExchangePolicy(ctx context.Context, policy *exchangepolicypb.ExchangePolicy) (*exchangepolicypb.ExchangePolicy, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.ExchangePolicies[policy.GetId()]; !ok {
		return nil, status.Error(codes.NotFound, "exchange policy not found")
	}
	if err := c.fake.ValidateTrustZone(policy.GetTrustZoneId()); err != nil {
		return nil, err
	}
	c.fake.ExchangePolicies[policy.GetId()] = clone(policy)
	return clone(policy), nil
}

func clone(policy *exchangepolicypb.ExchangePolicy) *exchangepolicypb.ExchangePolicy {
	return proto.Clone(policy).(*exchangepolicypb.ExchangePolicy)
}
