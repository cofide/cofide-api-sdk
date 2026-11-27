// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1beta1

import (
	"context"

	federationpb "github.com/cofide/cofide-api-sdk/gen/go/cofide/federation/v1beta1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
	"google.golang.org/grpc"
)

// FederationClient is an interface for a gRPC client for the v1beta1 version of the Connect FederationService.
type FederationClient interface {
	CreateFederation(ctx context.Context, federation *federationpb.Federation) (*federationpb.Federation, error)
	DeleteFederation(ctx context.Context, federationID string) error
	GetFederation(ctx context.Context, federationID string) (*federationpb.Federation, error)
	GetFederationByTrustZones(ctx context.Context, trustZoneID string, remoteTrustZoneID string) (*federationpb.Federation, error)
	ListFederations(ctx context.Context, filter *federationpb.ListFederationsRequest_Filter, pagination pagination.Pagination) ([]*federationpb.Federation, pagination.Pagination, error)
}

type federationClient struct {
	client federationpb.FederationServiceClient
}

// New instantiates a new FederationClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) FederationClient {
	return &federationClient{
		client: federationpb.NewFederationServiceClient(conn),
	}
}

func (c *federationClient) CreateFederation(ctx context.Context, federation *federationpb.Federation) (*federationpb.Federation, error) {
	resp, err := c.client.CreateFederation(ctx, &federationpb.CreateFederationRequest{
		Federation: federation,
	})
	if err != nil {
		return nil, err
	}
	return resp.GetFederation(), nil
}

func (c *federationClient) DeleteFederation(ctx context.Context, federationID string) error {
	_, err := c.client.DeleteFederation(ctx, &federationpb.DeleteFederationRequest{
		FederationId: federationID,
	})
	return err
}

// GetFederation retrieves a federation by its ID.
func (c *federationClient) GetFederation(ctx context.Context, federationID string) (*federationpb.Federation, error) {
	return c.getFederation(ctx, &federationpb.GetFederationRequest{
		Identifier: &federationpb.GetFederationRequest_FederationId{FederationId: federationID},
	})
}

// GetFederationByTrustZones retrieves a federation by the local trust zone that trusts the remote one. The order
// matters, since swapping the two IDs identifies a different federation. A federation can only have a remote trust
// zone for now, so this finds any federation, but it cannot find one whose remote is of another kind if one is added.
func (c *federationClient) GetFederationByTrustZones(ctx context.Context, trustZoneID string, remoteTrustZoneID string) (*federationpb.Federation, error) {
	return c.getFederation(ctx, &federationpb.GetFederationRequest{
		Identifier: &federationpb.GetFederationRequest_TrustZones{
			TrustZones: &federationpb.GetFederationRequest_TrustZonePair{
				TrustZoneId:       trustZoneID,
				RemoteTrustZoneId: remoteTrustZoneID,
			},
		},
	})
}

func (c *federationClient) getFederation(ctx context.Context, req *federationpb.GetFederationRequest) (*federationpb.Federation, error) {
	resp, err := c.client.GetFederation(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.GetFederation(), nil
}

func (c *federationClient) ListFederations(ctx context.Context, filter *federationpb.ListFederationsRequest_Filter, requestPagination pagination.Pagination) ([]*federationpb.Federation, pagination.Pagination, error) {
	resp, err := c.client.ListFederations(ctx, &federationpb.ListFederationsRequest{
		Filter:     filter,
		Pagination: requestPagination.PageRequestV1Beta1(),
	})
	if err != nil {
		return nil, pagination.Pagination{}, err
	}
	return resp.GetFederations(), requestPagination.Next(resp.GetPagination()), nil
}
