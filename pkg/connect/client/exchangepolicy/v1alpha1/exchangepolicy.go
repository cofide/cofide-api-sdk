// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"

	exchangepolicypb "github.com/cofide/cofide-api-sdk/gen/go/proto/exchange_policy/v1alpha1"
	exchangepolicysvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/exchange_policy_service/v1alpha1"
	"google.golang.org/grpc"
)

// ExchangePolicyClient is an interface for a gRPC client for the v1alpha1 version of the Connect ExchangePolicyService.
type ExchangePolicyClient interface {
	CreateExchangePolicy(ctx context.Context, policy *exchangepolicypb.ExchangePolicy) (*exchangepolicypb.ExchangePolicy, error)
	DestroyExchangePolicy(ctx context.Context, policyID string) error
	GetExchangePolicy(ctx context.Context, policyID string) (*exchangepolicypb.ExchangePolicy, error)
	ListExchangePolicies(ctx context.Context, filter *exchangepolicysvcpb.ListExchangePoliciesRequest_Filter) ([]*exchangepolicypb.ExchangePolicy, error)
	UpdateExchangePolicy(ctx context.Context, policy *exchangepolicypb.ExchangePolicy) (*exchangepolicypb.ExchangePolicy, error)
}

type exchangePolicyClient struct {
	client exchangepolicysvcpb.ExchangePolicyServiceClient
}

// New instantiates a new ExchangePolicyClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) ExchangePolicyClient {
	return &exchangePolicyClient{
		client: exchangepolicysvcpb.NewExchangePolicyServiceClient(conn),
	}
}

func (c *exchangePolicyClient) CreateExchangePolicy(ctx context.Context, policy *exchangepolicypb.ExchangePolicy) (*exchangepolicypb.ExchangePolicy, error) {
	resp, err := c.client.CreateExchangePolicy(ctx, &exchangepolicysvcpb.CreateExchangePolicyRequest{
		ExchangePolicy: policy,
	})
	if err != nil {
		return nil, err
	}
	return resp.ExchangePolicy, nil
}

func (c *exchangePolicyClient) DestroyExchangePolicy(ctx context.Context, policyID string) error {
	_, err := c.client.DestroyExchangePolicy(ctx, &exchangepolicysvcpb.DestroyExchangePolicyRequest{
		ExchangePolicyId: &policyID,
	})
	return err
}

func (c *exchangePolicyClient) GetExchangePolicy(ctx context.Context, policyID string) (*exchangepolicypb.ExchangePolicy, error) {
	resp, err := c.client.GetExchangePolicy(ctx, &exchangepolicysvcpb.GetExchangePolicyRequest{
		ExchangePolicyId: &policyID,
	})
	if err != nil {
		return nil, err
	}
	return resp.ExchangePolicy, nil
}

func (c *exchangePolicyClient) ListExchangePolicies(ctx context.Context, filter *exchangepolicysvcpb.ListExchangePoliciesRequest_Filter) ([]*exchangepolicypb.ExchangePolicy, error) {
	resp, err := c.client.ListExchangePolicies(ctx, &exchangepolicysvcpb.ListExchangePoliciesRequest{
		Filter: filter,
	})
	if err != nil {
		return nil, err
	}
	return resp.ExchangePolicies, nil
}

func (c *exchangePolicyClient) UpdateExchangePolicy(ctx context.Context, policy *exchangepolicypb.ExchangePolicy) (*exchangepolicypb.ExchangePolicy, error) {
	resp, err := c.client.UpdateExchangePolicy(ctx, &exchangepolicysvcpb.UpdateExchangePolicyRequest{
		ExchangePolicy: policy,
	})
	if err != nil {
		return nil, err
	}
	return resp.ExchangePolicy, nil
}
