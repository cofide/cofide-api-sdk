// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"

	trustzonesvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/trust_zone_service/v1alpha1"
	trustzonepb "github.com/cofide/cofide-api-sdk/gen/go/proto/trust_zone/v1alpha1"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"google.golang.org/grpc"
)

// TrustZoneClient is an interface for a gRPC client for the v1alpha1 version of the Connect TrustZoneService.
type TrustZoneClient interface {
	CreateTrustZone(ctx context.Context, trustZone *trustzonepb.TrustZone) (*trustzonepb.TrustZone, error)
	DestroyTrustZone(ctx context.Context, trustZoneID string) error
	GetTrustZone(ctx context.Context, trustZoneID string) (*trustzonepb.TrustZone, error)
	ListTrustZones(ctx context.Context, filter *trustzonesvcpb.ListTrustZonesRequest_Filter) ([]*trustzonepb.TrustZone, error)
	UpdateTrustZone(ctx context.Context, trustZone *trustzonepb.TrustZone) (*trustzonepb.TrustZone, error)
	RegisterAgent(ctx context.Context, agent *trustzonesvcpb.Agent, token string, bundle *types.Bundle) (string, error)
	RegisterTrustZoneServer(ctx context.Context, server *trustzonesvcpb.TrustZoneServer, bundle *types.Bundle, trustZoneServerID, token string) error
	UpdateTrustZoneBundle(ctx context.Context, trustZoneID string, bundle *types.Bundle) error
}

type trustZoneClient struct {
	trustZoneClient trustzonesvcpb.TrustZoneServiceClient
}

// New instantiates a new TrustZoneClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) TrustZoneClient {
	return &trustZoneClient{
		trustZoneClient: trustzonesvcpb.NewTrustZoneServiceClient(conn),
	}
}

func (c *trustZoneClient) CreateTrustZone(ctx context.Context, trustZone *trustzonepb.TrustZone) (*trustzonepb.TrustZone, error) {
	resp, err := c.trustZoneClient.CreateTrustZone(ctx, &trustzonesvcpb.CreateTrustZoneRequest{
		TrustZone: trustZone,
	})
	if err != nil {
		return nil, err
	}

	return resp.TrustZone, nil
}

func (c *trustZoneClient) DestroyTrustZone(ctx context.Context, trustZoneID string) error {
	_, err := c.trustZoneClient.DestroyTrustZone(ctx, &trustzonesvcpb.DestroyTrustZoneRequest{
		TrustZoneId: &trustZoneID,
	})
	return err
}

func (c *trustZoneClient) GetTrustZone(ctx context.Context, trustZoneID string) (*trustzonepb.TrustZone, error) {
	resp, err := c.trustZoneClient.GetTrustZone(ctx, &trustzonesvcpb.GetTrustZoneRequest{
		TrustZoneId: &trustZoneID,
	})
	if err != nil {
		return nil, err
	}

	return resp.TrustZone, nil
}

func (c *trustZoneClient) ListTrustZones(ctx context.Context, filter *trustzonesvcpb.ListTrustZonesRequest_Filter) ([]*trustzonepb.TrustZone, error) {
	resp, err := c.trustZoneClient.ListTrustZones(ctx, &trustzonesvcpb.ListTrustZonesRequest{
		Filter: filter,
	})
	if err != nil {
		return nil, err
	}

	return resp.TrustZones, nil
}

func (c *trustZoneClient) UpdateTrustZone(ctx context.Context, trustZone *trustzonepb.TrustZone) (*trustzonepb.TrustZone, error) {
	resp, err := c.trustZoneClient.UpdateTrustZone(ctx, &trustzonesvcpb.UpdateTrustZoneRequest{
		TrustZone: trustZone,
	})
	if err != nil {
		return nil, err
	}

	return resp.TrustZone, nil
}

func (c *trustZoneClient) RegisterAgent(ctx context.Context, agent *trustzonesvcpb.Agent, token string, bundle *types.Bundle) (string, error) {
	resp, err := c.trustZoneClient.RegisterAgent(ctx, &trustzonesvcpb.RegisterAgentRequest{
		Agent:      agent,
		AgentToken: token,
		Bundle:     bundle,
	})
	if err != nil {
		return "", err
	}

	return resp.AgentId, nil
}

func (c *trustZoneClient) RegisterTrustZoneServer(ctx context.Context, server *trustzonesvcpb.TrustZoneServer, bundle *types.Bundle, trustZoneServerID string, token string) error {
	// RegisterTrustZoneServerResponse currently empty
	_, err := c.trustZoneClient.RegisterTrustZoneServer(ctx, &trustzonesvcpb.RegisterTrustZoneServerRequest{
		TrustZoneServer:   server,
		Bundle:            bundle,
		TrustZoneServerId: trustZoneServerID,
		JoinToken:         token,
	})
	return err
}

func (c *trustZoneClient) UpdateTrustZoneBundle(ctx context.Context, trustZoneID string, bundle *types.Bundle) error {
	_, err := c.trustZoneClient.UpdateTrustZoneBundle(ctx, &trustzonesvcpb.UpdateTrustZoneBundleRequest{
		Bundle:      bundle,
		TrustZoneId: trustZoneID,
	})
	if err != nil {
		return err
	}
	return nil
}
