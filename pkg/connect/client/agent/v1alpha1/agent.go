// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"
	"fmt"
	"log/slog"

	agentpb "github.com/cofide/cofide-api-sdk/gen/go/proto/agent/v1alpha1"
	agentsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/agent_service/v1alpha1"
	federatedservicepb "github.com/cofide/cofide-api-sdk/gen/go/proto/federated_service/v1alpha1"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"google.golang.org/grpc"
)

type AgentClient interface {
	CreateAgentJoinToken(ctx context.Context, clusterID, trustZoneID string) (string, error)
	UpdateTrustZoneBundle(ctx context.Context, bundle *types.Bundle) error
	UpdateManagedTrustZoneBundle(ctx context.Context, trustZoneID string, bundle *types.Bundle) error
	UpdateAgentStatus(ctx context.Context, status *agentpb.AgentStatus) error
	// DEPRECATED: Federated services will move to a separate client.
	RegisterFederatedService(ctx context.Context, fs *federatedservicepb.FederatedService) (string, error)
	DeregisterFederatedService(ctx context.Context, federatedServiceID string) error
	GetFederatedService(ctx context.Context, federatedServiceID string) (*federatedservicepb.FederatedService, error)
	UpdateFederatedService(ctx context.Context, fs *federatedservicepb.FederatedService) error
}

type agentClient struct {
	client agentsvcpb.AgentServiceClient
}

// New instantiates a new ClusterClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) AgentClient {
	return &agentClient{
		client: agentsvcpb.NewAgentServiceClient(conn),
	}
}

func (c *agentClient) CreateAgentJoinToken(ctx context.Context, clusterID, trustZoneID string) (string, error) {
	resp, err := c.client.CreateAgentJoinToken(ctx, &agentsvcpb.CreateAgentJoinTokenRequest{
		ClusterId:   &clusterID,
		TrustZoneId: &trustZoneID,
	})
	if err != nil {
		// FIXME: Does this lose the status & code?
		return "", fmt.Errorf("failed to create agent join token: %w", err)
	}

	return resp.GetAgentToken(), nil
}

func (c *agentClient) UpdateTrustZoneBundle(ctx context.Context, bundle *types.Bundle) error {
	trustzoneBundleUpdateRequest := &agentsvcpb.UpdateTrustZoneBundleRequest{
		Bundle: bundle,
	}
	return c.doUpdateTrustZoneBundle(ctx, trustzoneBundleUpdateRequest)
}

func (c *agentClient) UpdateManagedTrustZoneBundle(ctx context.Context, trustZoneID string, bundle *types.Bundle) error {
	trustzoneBundleUpdateRequest := &agentsvcpb.UpdateTrustZoneBundleRequest{
		Bundle:      bundle,
		TrustZoneId: trustZoneID,
	}
	return c.doUpdateTrustZoneBundle(ctx, trustzoneBundleUpdateRequest)
}

func (c *agentClient) doUpdateTrustZoneBundle(ctx context.Context, trustzoneBundleUpdateRequest *agentsvcpb.UpdateTrustZoneBundleRequest) error {
	resp, err := c.client.UpdateTrustZoneBundle(ctx, trustzoneBundleUpdateRequest)

	if err != nil {
		return fmt.Errorf("failed to update bundle: %w", err)
	}
	slog.Info("successfully updated trust bundle", "response", resp.Message)
	return nil
}

func (c *agentClient) UpdateAgentStatus(ctx context.Context, status *agentpb.AgentStatus) error {
	slog.Debug("updating agent status", "status", status.Status)
	_, err := c.client.UpdateAgentStatus(ctx, &agentsvcpb.UpdateAgentStatusRequest{Status: status})
	if err != nil {
		return fmt.Errorf("failed to update agent status: %w", err)
	}
	slog.Debug("updated agent status", "agent", status.Status)
	return nil
}

func (c *agentClient) RegisterFederatedService(ctx context.Context, fs *federatedservicepb.FederatedService) (string, error) {
	resp, err := c.client.RegisterFederatedService(ctx, &agentsvcpb.RegisterFederatedServiceRequest{
		Service: fs,
	})
	if err != nil {
		return "", fmt.Errorf("failed to register federated service: %w", err)
	}

	slog.Info("successfully registered FederatedService", "service_id", resp.ServiceId)

	return resp.ServiceId, nil
}

func (c *agentClient) DeregisterFederatedService(ctx context.Context, federatedServiceID string) error {
	_, err := c.client.DeregisterFederatedService(ctx, &agentsvcpb.DeregisterFederatedServiceRequest{
		ServiceId: federatedServiceID,
	})

	if err != nil {
		return fmt.Errorf("failed to de-register federated service: %w", err)
	}

	slog.Info("successfully de-registered FederatedService", "service_id", federatedServiceID)

	return nil
}

func (c *agentClient) GetFederatedService(ctx context.Context, federatedServiceID string) (*federatedservicepb.FederatedService, error) {
	resp, err := c.client.GetFederatedService(ctx, &agentsvcpb.GetFederatedServiceRequest{
		ServiceId: federatedServiceID,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get federated service: %w", err)
	}

	slog.Info("successfully fetched FederatedService", "service_id", federatedServiceID)

	return resp.Service, nil
}

func (c *agentClient) UpdateFederatedService(ctx context.Context, fs *federatedservicepb.FederatedService) error {
	_, err := c.client.UpdateFederatedService(ctx, &agentsvcpb.UpdateFederatedServiceRequest{
		Service: fs,
	})

	if err != nil {
		return fmt.Errorf("failed to update federated service: %w", err)
	}

	slog.Info("successfully updated FederatedService", "service_id", fs.Id)

	return nil
}
