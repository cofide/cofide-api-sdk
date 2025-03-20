// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"

	federationsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/federation_service/v1alpha1"
	federationpb "github.com/cofide/cofide-api-sdk/gen/go/proto/federation/v1alpha1"
	"google.golang.org/grpc"
)

// FederationClient is an interface for a gRPC client for the v1alpha1 version of the Connect FederationClient.
type FederationClient interface {
	CreateFederation(ctx context.Context, federation *federationpb.Federation) (*federationpb.Federation, error)
	GetFederation(ctx context.Context, federationID string) (*federationpb.Federation, error)
	ListFederations(ctx context.Context, filter *federationsvcpb.ListFederationsRequest_Filter) ([]*federationpb.Federation, error)
	DestroyFederation(ctx context.Context, federationID string) error
}

type federationClient struct {
	federationClient federationsvcpb.FederationServiceClient
}

// New instantiates a new FederationClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) FederationClient {
	return &federationClient{
		federationClient: federationsvcpb.NewFederationServiceClient(conn),
	}
}

func (c *federationClient) CreateFederation(ctx context.Context, federation *federationpb.Federation) (*federationpb.Federation, error) {
	resp, err := c.federationClient.CreateFederation(ctx, &federationsvcpb.CreateFederationRequest{
		Federation: federation,
	})
	if err != nil {
		return nil, err
	}

	return resp.Federation, nil
}

func (c *federationClient) GetFederation(ctx context.Context, federationID string) (*federationpb.Federation, error) {
	resp, err := c.federationClient.GetFederation(ctx, &federationsvcpb.GetFederationRequest{
		FederationId: federationID,
	})
	if err != nil {
		return nil, err
	}

	return resp.Federation, nil
}

func (c *federationClient) ListFederations(ctx context.Context, filter *federationsvcpb.ListFederationsRequest_Filter) ([]*federationpb.Federation, error) {
	resp, err := c.federationClient.ListFederations(ctx, &federationsvcpb.ListFederationsRequest{
		Filter: filter,
	})
	if err != nil {
		return nil, err
	}

	return resp.Federations, nil
}

func (c *federationClient) DestroyFederation(ctx context.Context, federationID string) error {
	_, err := c.federationClient.DestroyFederation(ctx, &federationsvcpb.DestroyFederationRequest{
		FederationId: &federationID,
	})
	return err
}
