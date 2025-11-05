// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	rolebindingsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/role_binding_service/v1alpha1"
	rolebindingpb "github.com/cofide/cofide-api-sdk/gen/go/proto/role_binding/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	rolebindingv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/rolebinding/v1alpha1"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type fakeRoleBindingClient struct {
	fake *fakeconnect.FakeConnect
}

// New instantiates a new RoleBindingClient for communication with a fake Connect API.
func New(fake *fakeconnect.FakeConnect) rolebindingv1alpha1.RoleBindingClient {
	return &fakeRoleBindingClient{
		fake: fake,
	}
}

func (c *fakeRoleBindingClient) CreateRoleBinding(ctx context.Context, roleBinding *rolebindingpb.RoleBinding) (*rolebindingpb.RoleBinding, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	id := uuid.New().String()
	roleBinding = clone(roleBinding)
	roleBinding.Id = id
	c.fake.RoleBindings[roleBinding.GetId()] = roleBinding
	return clone(roleBinding), nil
}

func (c *fakeRoleBindingClient) DestroyRoleBinding(ctx context.Context, roleBindingID string) error {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.RoleBindings[roleBindingID]; !ok {
		return status.Error(codes.NotFound, "role binding not found")
	}
	delete(c.fake.RoleBindings, roleBindingID)
	return nil
}

func (c *fakeRoleBindingClient) GetRoleBinding(ctx context.Context, roleBindingID string) (*rolebindingpb.RoleBinding, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	roleBinding, ok := c.fake.RoleBindings[roleBindingID]
	if !ok {
		return nil, status.Error(codes.NotFound, "role binding not found")
	}
	return clone(roleBinding), nil
}

func (c *fakeRoleBindingClient) ListRoleBindings(ctx context.Context, filter *rolebindingsvcpb.ListRoleBindingsRequest_Filter) ([]*rolebindingpb.RoleBinding, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	roleBindings := []*rolebindingpb.RoleBinding{}
	for _, roleBinding := range c.fake.RoleBindings {
		if roleBindingMatches(roleBinding, filter) {
			roleBindings = append(roleBindings, clone(roleBinding))
		}
	}
	return roleBindings, nil
}

func roleBindingMatches(roleBinding *rolebindingpb.RoleBinding, filter *rolebindingsvcpb.ListRoleBindingsRequest_Filter) bool {
	if filter == nil {
		return true
	}
	if filter.RoleId != nil && roleBinding.GetRoleId() != *filter.RoleId {
		return false
	}
	if filter.UserId != nil && roleBinding.GetUser().GetId() != *filter.UserId {
		return false
	}
	if filter.GroupId != nil && roleBinding.GetGroup().GetId() != *filter.GroupId {
		return false
	}
	if filter.ResourceId != nil && roleBinding.GetResource().GetId() != *filter.ResourceId {
		return false
	}
	if filter.ResourceType != nil && roleBinding.GetResource().GetType() != *filter.ResourceType {
		return false
	}
	return true
}

func (c *fakeRoleBindingClient) UpdateRoleBinding(ctx context.Context, roleBinding *rolebindingpb.RoleBinding) (*rolebindingpb.RoleBinding, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.RoleBindings[roleBinding.GetId()]; !ok {
		return nil, status.Error(codes.NotFound, "role binding not found")
	}

	c.fake.RoleBindings[roleBinding.GetId()] = clone(roleBinding)
	return clone(roleBinding), nil
}

func clone(roleBinding *rolebindingpb.RoleBinding) *rolebindingpb.RoleBinding {
	return proto.Clone(roleBinding).(*rolebindingpb.RoleBinding)
}
