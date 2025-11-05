// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"

	rolebindingsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/role_binding_service/v1alpha1"
	rolebindingpb "github.com/cofide/cofide-api-sdk/gen/go/proto/role_binding/v1alpha1"
	"google.golang.org/grpc"
)

// RoleBindingClient is an interface for a gRPC client for the v1alpha1 version of the Connect RoleBindingService.
type RoleBindingClient interface {
	CreateRoleBinding(ctx context.Context, roleBinding *rolebindingpb.RoleBinding) (*rolebindingpb.RoleBinding, error)
	DestroyRoleBinding(ctx context.Context, roleBindingID string) error
	GetRoleBinding(ctx context.Context, roleBindingID string) (*rolebindingpb.RoleBinding, error)
	ListRoleBindings(ctx context.Context, filter *rolebindingsvcpb.ListRoleBindingsRequest_Filter) ([]*rolebindingpb.RoleBinding, error)
	UpdateRoleBinding(ctx context.Context, roleBinding *rolebindingpb.RoleBinding) (*rolebindingpb.RoleBinding, error)
}

type roleBindingClient struct {
	roleBindingClient rolebindingsvcpb.RoleBindingServiceClient
}

// New instantiates a new RoleBindingClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) RoleBindingClient {
	return &roleBindingClient{
		roleBindingClient: rolebindingsvcpb.NewRoleBindingServiceClient(conn),
	}
}

func (c *roleBindingClient) CreateRoleBinding(ctx context.Context, roleBinding *rolebindingpb.RoleBinding) (*rolebindingpb.RoleBinding, error) {
	resp, err := c.roleBindingClient.CreateRoleBinding(ctx, &rolebindingsvcpb.CreateRoleBindingRequest{
		RoleBinding: roleBinding,
	})
	if err != nil {
		return nil, err
	}

	return resp.RoleBinding, nil
}

func (c *roleBindingClient) DestroyRoleBinding(ctx context.Context, roleBindingID string) error {
	_, err := c.roleBindingClient.DestroyRoleBinding(ctx, &rolebindingsvcpb.DestroyRoleBindingRequest{
		RoleBindingId: roleBindingID,
	})
	return err
}

func (c *roleBindingClient) GetRoleBinding(ctx context.Context, roleBindingID string) (*rolebindingpb.RoleBinding, error) {
	resp, err := c.roleBindingClient.GetRoleBinding(ctx, &rolebindingsvcpb.GetRoleBindingRequest{
		RoleBindingId: roleBindingID,
	})
	if err != nil {
		return nil, err
	}

	return resp.RoleBinding, nil
}

func (c *roleBindingClient) ListRoleBindings(ctx context.Context, filter *rolebindingsvcpb.ListRoleBindingsRequest_Filter) ([]*rolebindingpb.RoleBinding, error) {
	resp, err := c.roleBindingClient.ListRoleBindings(ctx, &rolebindingsvcpb.ListRoleBindingsRequest{
		Filter: filter,
	})
	if err != nil {
		return nil, err
	}

	return resp.RoleBindings, nil
}

func (c *roleBindingClient) UpdateRoleBinding(ctx context.Context, roleBinding *rolebindingpb.RoleBinding) (*rolebindingpb.RoleBinding, error) {
	resp, err := c.roleBindingClient.UpdateRoleBinding(ctx, &rolebindingsvcpb.UpdateRoleBindingRequest{
		RoleBinding: roleBinding,
	})
	if err != nil {
		return nil, err
	}

	return resp.RoleBinding, nil
}
