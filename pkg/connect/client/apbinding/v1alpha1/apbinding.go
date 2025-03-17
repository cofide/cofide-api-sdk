// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"

	apbindingpb "github.com/cofide/cofide-api-sdk/gen/go/proto/ap_binding/v1alpha1"
	apbindingsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/ap_binding_service/v1alpha1"
	"google.golang.org/grpc"
)

// APBindingClient is an interface for a gRPC client for the v1alpha1 version of the Connect APBindingService.
type APBindingClient interface {
	CreateAPBinding(ctx context.Context, binding *apbindingpb.APBinding) (*apbindingpb.APBinding, error)
	DestroyAPBinding(ctx context.Context, bindingID string) error
	GetAPBinding(ctx context.Context, bindingID string) (*apbindingpb.APBinding, error)
	ListAPBindings(ctx context.Context, filter *apbindingsvcpb.ListAPBindingsRequest_Filter) ([]*apbindingpb.APBinding, error)
	UpdateAPBinding(ctx context.Context, binding *apbindingpb.APBinding) (*apbindingpb.APBinding, error)
}

type apBindingClient struct {
	client apbindingsvcpb.APBindingServiceClient
}

// New instantiates a new APBindingClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) APBindingClient {
	return &apBindingClient{
		client: apbindingsvcpb.NewAPBindingServiceClient(conn),
	}
}

func (c *apBindingClient) CreateAPBinding(ctx context.Context, binding *apbindingpb.APBinding) (*apbindingpb.APBinding, error) {
	resp, err := c.client.CreateAPBinding(ctx, &apbindingsvcpb.CreateAPBindingRequest{
		Binding: binding,
	})
	if err != nil {
		return nil, err
	}
	return resp.Binding, nil
}

func (c *apBindingClient) DestroyAPBinding(ctx context.Context, bindingID string) error {
	_, err := c.client.DestroyAPBinding(ctx, &apbindingsvcpb.DestroyAPBindingRequest{
		BindingId: &bindingID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *apBindingClient) GetAPBinding(ctx context.Context, bindingID string) (*apbindingpb.APBinding, error) {
	resp, err := c.client.GetAPBinding(ctx, &apbindingsvcpb.GetAPBindingRequest{
		BindingId: &bindingID,
	})
	if err != nil {
		return nil, err
	}
	return resp.Binding, nil
}

func (c *apBindingClient) ListAPBindings(ctx context.Context, filter *apbindingsvcpb.ListAPBindingsRequest_Filter) ([]*apbindingpb.APBinding, error) {
	resp, err := c.client.ListAPBindings(ctx, &apbindingsvcpb.ListAPBindingsRequest{
		Filter: filter,
	})
	if err != nil {
		return nil, err
	}
	return resp.Bindings, nil
}

func (c *apBindingClient) UpdateAPBinding(ctx context.Context, binding *apbindingpb.APBinding) (*apbindingpb.APBinding, error) {
	resp, err := c.client.UpdateAPBinding(ctx, &apbindingsvcpb.UpdateAPBindingRequest{
		Binding: binding,
	})
	if err != nil {
		return nil, err
	}
	return resp.Binding, nil
}
