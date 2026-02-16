// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"

	trustzoneserversvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/trust_zone_server_service/v1alpha1"
	trustzoneserverpb "github.com/cofide/cofide-api-sdk/gen/go/proto/trust_zone_server/v1alpha1"
	"google.golang.org/grpc"
)

// TrustZoneServerClient is an interface for a gRPC client for the v1alpha1 version of the Connect TrustZoneServerService.
type TrustZoneServerClient interface {
	CreateTrustZoneServer(ctx context.Context, trustZoneServer *trustzoneserverpb.TrustZoneServer) (*trustzoneserverpb.TrustZoneServer, error)
	DestroyTrustZoneServer(ctx context.Context, trustZoneServerID string) error
	GetTrustZoneServer(ctx context.Context, trustZoneServerID string) (*trustzoneserverpb.TrustZoneServer, error)
	ListTrustZoneServers(ctx context.Context, filter *trustzoneserversvcpb.ListTrustZoneServersRequest_Filter) ([]*trustzoneserverpb.TrustZoneServer, error)
	UpdateTrustZoneServer(ctx context.Context, trustZoneServer *trustzoneserverpb.TrustZoneServer, updateMask *trustzoneserversvcpb.UpdateTrustZoneServerRequest_UpdateMask) (*trustzoneserverpb.TrustZoneServer, error)
	UpdateTrustZoneServerStatus(ctx context.Context, trustZoneServerID string, status trustzoneserverpb.TrustZoneServerStatus) error
}

type trustZoneServerClient struct {
	trustZoneServerClient trustzoneserversvcpb.TrustZoneServerServiceClient
}

// New instantiates a new TrustZoneServerClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) TrustZoneServerClient {
	return &trustZoneServerClient{
		trustZoneServerClient: trustzoneserversvcpb.NewTrustZoneServerServiceClient(conn),
	}
}

func (c *trustZoneServerClient) CreateTrustZoneServer(ctx context.Context, trustZoneServer *trustzoneserverpb.TrustZoneServer) (*trustzoneserverpb.TrustZoneServer, error) {
	resp, err := c.trustZoneServerClient.CreateTrustZoneServer(ctx, &trustzoneserversvcpb.CreateTrustZoneServerRequest{
		TrustZoneServer: trustZoneServer,
	})
	if err != nil {
		return nil, err
	}

	return resp.TrustZoneServer, nil
}

func (c *trustZoneServerClient) DestroyTrustZoneServer(ctx context.Context, trustZoneServerID string) error {
	_, err := c.trustZoneServerClient.DestroyTrustZoneServer(ctx, &trustzoneserversvcpb.DestroyTrustZoneServerRequest{
		TrustZoneServerId: trustZoneServerID,
	})
	return err
}

func (c *trustZoneServerClient) GetTrustZoneServer(ctx context.Context, trustZoneServerID string) (*trustzoneserverpb.TrustZoneServer, error) {
	resp, err := c.trustZoneServerClient.GetTrustZoneServer(ctx, &trustzoneserversvcpb.GetTrustZoneServerRequest{
		TrustZoneServerId: trustZoneServerID,
	})
	if err != nil {
		return nil, err
	}

	return resp.TrustZoneServer, nil
}

func (c *trustZoneServerClient) ListTrustZoneServers(ctx context.Context, filter *trustzoneserversvcpb.ListTrustZoneServersRequest_Filter) ([]*trustzoneserverpb.TrustZoneServer, error) {
	resp, err := c.trustZoneServerClient.ListTrustZoneServers(ctx, &trustzoneserversvcpb.ListTrustZoneServersRequest{
		Filter: filter,
	})
	if err != nil {
		return nil, err
	}

	return resp.TrustZoneServers, nil
}

func (c *trustZoneServerClient) UpdateTrustZoneServer(
	ctx context.Context,
	trustZoneServer *trustzoneserverpb.TrustZoneServer,
	updateMask *trustzoneserversvcpb.UpdateTrustZoneServerRequest_UpdateMask,
) (*trustzoneserverpb.TrustZoneServer, error) {
	resp, err := c.trustZoneServerClient.UpdateTrustZoneServer(ctx, &trustzoneserversvcpb.UpdateTrustZoneServerRequest{
		TrustZoneServer: trustZoneServer,
		UpdateMask:      updateMask,
	})
	if err != nil {
		return nil, err
	}

	return resp.TrustZoneServer, nil
}

func (c *trustZoneServerClient) UpdateTrustZoneServerStatus(
	ctx context.Context,
	trustZoneServerID string,
	status trustzoneserverpb.TrustZoneServerStatus,
) error {
	_, err := c.trustZoneServerClient.UpdateTrustZoneServerStatus(ctx, &trustzoneserversvcpb.UpdateTrustZoneServerStatusRequest{
		TrustZoneServerId: trustZoneServerID,
		Status:            status,
	})
	return err
}
