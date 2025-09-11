// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"

	agentpb "github.com/cofide/cofide-api-sdk/gen/go/proto/agent/v1alpha1"
	agentsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/agent_service/v1alpha1"
	federatedservicepb "github.com/cofide/cofide-api-sdk/gen/go/proto/federated_service/v1alpha1"
	"google.golang.org/grpc"
)

// AgentClient is an interface for a gRPC client for the v1alpha1 version of the Connect AgentService.
type AgentClient interface {
	CreateAgentJoinToken(ctx context.Context, clusterID, trustZoneID string) (string, error)
	UpdateAgentStatus(ctx context.Context, status *agentpb.AgentStatus) error
	// DEPRECATED: Federated services will move to a separate client.
	RegisterFederatedService(ctx context.Context, fs *federatedservicepb.FederatedService) (string, error)
	DeregisterFederatedService(ctx context.Context, federatedServiceID string) error
	GetFederatedService(ctx context.Context, federatedServiceID string) (*federatedservicepb.FederatedService, error)
	UpdateFederatedService(ctx context.Context, fs *federatedservicepb.FederatedService) error
	ListFederatedServices(ctx context.Context) ([]*federatedservicepb.FederatedService, error)
}

type agentClient struct {
	client agentsvcpb.AgentServiceClient
}

// New instantiates a new AgentClient for communication with a Connect API.
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
		return "", err
	}
	return resp.GetAgentToken(), nil
}

func (c *agentClient) UpdateAgentStatus(ctx context.Context, status *agentpb.AgentStatus) error {
	_, err := c.client.UpdateAgentStatus(ctx, &agentsvcpb.UpdateAgentStatusRequest{Status: status})
	if err != nil {
		return err
	}
	return nil
}

func (c *agentClient) RegisterFederatedService(ctx context.Context, fs *federatedservicepb.FederatedService) (string, error) {
	resp, err := c.client.RegisterFederatedService(ctx, &agentsvcpb.RegisterFederatedServiceRequest{
		Service: fs,
	})
	if err != nil {
		return "", err
	}
	return resp.ServiceId, nil
}

func (c *agentClient) DeregisterFederatedService(ctx context.Context, federatedServiceID string) error {
	_, err := c.client.DeregisterFederatedService(ctx, &agentsvcpb.DeregisterFederatedServiceRequest{
		ServiceId: federatedServiceID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *agentClient) GetFederatedService(ctx context.Context, federatedServiceID string) (*federatedservicepb.FederatedService, error) {
	resp, err := c.client.GetFederatedService(ctx, &agentsvcpb.GetFederatedServiceRequest{
		ServiceId: federatedServiceID,
	})
	if err != nil {
		return nil, err
	}
	return resp.Service, nil
}

func (c *agentClient) UpdateFederatedService(ctx context.Context, fs *federatedservicepb.FederatedService) error {
	_, err := c.client.UpdateFederatedService(ctx, &agentsvcpb.UpdateFederatedServiceRequest{
		Service: fs,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *agentClient) ListFederatedServices(ctx context.Context) ([]*federatedservicepb.FederatedService, error) {
	resp, err := c.client.ListFederatedServices(ctx, &agentsvcpb.ListFederatedServicesRequest{})
	if err != nil {
		return nil, err
	}
	return resp.Services, nil
}
