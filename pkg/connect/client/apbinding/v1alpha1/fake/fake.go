// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	apbindingpb "github.com/cofide/cofide-api-sdk/gen/go/proto/ap_binding/v1alpha1"
	apbindingsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/ap_binding_service/v1alpha1"
	apbindingv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/apbinding/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type fakeAPBindingClient struct {
	fake *fakeconnect.FakeConnect
}

// New instantiates a new APBindingClient for communication with a fake Connect API.
func New(fake *fakeconnect.FakeConnect) apbindingv1alpha1.APBindingClient {
	return &fakeAPBindingClient{fake: fake}
}

func (c *fakeAPBindingClient) CreateAPBinding(ctx context.Context, binding *apbindingpb.APBinding) (*apbindingpb.APBinding, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if err := c.fake.ValidateTrustZone(binding.GetTrustZoneId()); err != nil {
		return nil, err
	}
	if err := c.fake.ValidateAttestationPolicy(binding.GetPolicyId()); err != nil {
		return nil, err
	}
	id := uuid.New().String()
	binding = clone(binding)
	binding.Id = &id
	c.fake.APBindings[binding.GetId()] = binding
	return clone(binding), nil
}

func (c *fakeAPBindingClient) DestroyAPBinding(ctx context.Context, bindingID string) error {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.APBindings[bindingID]; !ok {
		return status.Error(codes.NotFound, "attestation policy binding not found")
	}
	delete(c.fake.APBindings, bindingID)
	return nil
}

func (c *fakeAPBindingClient) GetAPBinding(ctx context.Context, bindingID string) (*apbindingpb.APBinding, error) {
	binding, ok := c.fake.APBindings[bindingID]
	if !ok {
		return nil, status.Error(codes.NotFound, "attestation policy binding not found")
	}
	return clone(binding), nil
}

func (c *fakeAPBindingClient) ListAPBindings(ctx context.Context, filter *apbindingsvcpb.ListAPBindingsRequest_Filter) ([]*apbindingpb.APBinding, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	bindings := []*apbindingpb.APBinding{}
	for _, binding := range c.fake.APBindings {
		if bindingMatches(binding, filter) {
			bindings = append(bindings, clone(binding))
		}
	}
	return bindings, nil
}

func bindingMatches(binding *apbindingpb.APBinding, filter *apbindingsvcpb.ListAPBindingsRequest_Filter) bool {
	if filter == nil {
		return true
	}
	if filter.OrgId != nil && binding.GetOrgId() != *filter.OrgId {
		return false
	}
	if filter.TrustZoneId != nil && binding.GetTrustZoneId() != *filter.TrustZoneId {
		return false
	}
	if filter.PolicyId != nil && binding.GetPolicyId() != *filter.PolicyId {
		return false
	}
	return true
}

func (c *fakeAPBindingClient) UpdateAPBinding(ctx context.Context, binding *apbindingpb.APBinding) (*apbindingpb.APBinding, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.APBindings[binding.GetId()]; !ok {
		return nil, status.Error(codes.NotFound, "attestation policy binding not found")
	}
	if err := c.fake.ValidateTrustZone(binding.GetTrustZoneId()); err != nil {
		return nil, err
	}
	if err := c.fake.ValidateAttestationPolicy(binding.GetPolicyId()); err != nil {
		return nil, err
	}
	c.fake.APBindings[binding.GetId()] = clone(binding)
	return clone(binding), nil
}

func clone(binding *apbindingpb.APBinding) *apbindingpb.APBinding {
	return proto.Clone(binding).(*apbindingpb.APBinding)
}
