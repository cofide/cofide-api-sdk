// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	agentpb "github.com/cofide/cofide-api-sdk/gen/go/proto/agent/v1alpha1"
	federatedservicepb "github.com/cofide/cofide-api-sdk/gen/go/proto/federated_service/v1alpha1"
	agentv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/agent/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/google/uuid"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type fakeAgentClient struct {
	fake *fakeconnect.FakeConnect
}

// New instantiates a new AgentClient for communication with a fake Connect API.
func New(fake *fakeconnect.FakeConnect) agentv1alpha1.AgentClient {
	return &fakeAgentClient{
		fake: fake,
	}
}

func (c *fakeAgentClient) CreateAgentJoinToken(ctx context.Context, clusterID, trustZoneID string) (string, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if err := c.fake.ValidateTrustZone(trustZoneID); err != nil {
		return "", err
	}
	if err := c.fake.ValidateCluster(clusterID); err != nil {
		return "", err
	}

	if c.fake.AgentJoinTokens[trustZoneID] == nil {
		c.fake.AgentJoinTokens[trustZoneID] = make(map[string]string)
	}
	token := uuid.New().String()
	c.fake.AgentJoinTokens[trustZoneID][clusterID] = token
	return token, nil
}

func (c *fakeAgentClient) UpdateTrustZoneBundle(ctx context.Context, bundle *types.Bundle) error {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	agentID, err := getAgentIDFromMetadata(ctx)
	if err != nil {
		return err
	}

	if err := c.fake.ValidateAgent(agentID); err != nil {
		return err
	}

	agent := c.fake.Agents[agentID]
	c.fake.TrustZoneBundles[agent.GetTrustZoneId()] = cloneBundle(bundle)
	return nil
}

func (c *fakeAgentClient) UpdateManagedTrustZoneBundle(ctx context.Context, trustZoneID string, bundle *types.Bundle) error {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if err := c.fake.ValidateTrustZone(trustZoneID); err != nil {
		return err
	}

	c.fake.TrustZoneBundles[trustZoneID] = cloneBundle(bundle)
	return nil
}

func (c *fakeAgentClient) UpdateAgentStatus(ctx context.Context, agentStatus *agentpb.AgentStatus) error {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	agentID, err := getAgentIDFromMetadata(ctx)
	if err != nil {
		return err
	}

	if err := c.fake.ValidateAgent(agentID); err != nil {
		return err
	}

	c.fake.AgentStatus[agentID] = cloneStatus(agentStatus)
	return nil
}

func (c *fakeAgentClient) RegisterFederatedService(ctx context.Context, fs *federatedservicepb.FederatedService) (string, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.FederatedServices[fs.Id]; ok {
		return "", status.Error(codes.AlreadyExists, "federated service already exists")
	}

	c.fake.FederatedServices[fs.Id] = cloneFS(fs)
	return fs.Id, nil
}

func (c *fakeAgentClient) DeregisterFederatedService(ctx context.Context, federatedServiceID string) error {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if err := c.fake.ValidateFederatedService(federatedServiceID); err != nil {
		return status.Error(codes.NotFound, "federated service not found")
	}

	delete(c.fake.FederatedServices, federatedServiceID)
	return nil
}

func (c *fakeAgentClient) GetFederatedService(ctx context.Context, federatedServiceID string) (*federatedservicepb.FederatedService, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if err := c.fake.ValidateFederatedService(federatedServiceID); err != nil {
		return nil, status.Error(codes.NotFound, "federated service not found")
	}

	return cloneFS(c.fake.FederatedServices[federatedServiceID]), nil
}

func (c *fakeAgentClient) UpdateFederatedService(ctx context.Context, fs *federatedservicepb.FederatedService) error {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if err := c.fake.ValidateFederatedService(fs.GetId()); err != nil {
		return status.Error(codes.NotFound, "federated service not found")
	}

	c.fake.FederatedServices[fs.GetId()] = cloneFS(fs)
	return nil
}

func (c *fakeAgentClient) ListFederatedServices(ctx context.Context) ([]*federatedservicepb.FederatedService, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	fss := []*federatedservicepb.FederatedService{}
	for _, fs := range c.fake.FederatedServices {
		fss = append(fss, cloneFS(fs))
	}
	return fss, nil
}

func cloneBundle(bundle *types.Bundle) *types.Bundle {
	return proto.Clone(bundle).(*types.Bundle)
}

func cloneStatus(status *agentpb.AgentStatus) *agentpb.AgentStatus {
	return proto.Clone(status).(*agentpb.AgentStatus)
}

func cloneFS(fs *federatedservicepb.FederatedService) *federatedservicepb.FederatedService {
	return proto.Clone(fs).(*federatedservicepb.FederatedService)
}

func getAgentIDFromMetadata(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "missing gRPC metadata")
	}
	agentIDs := md.Get("agent-id")
	if len(agentIDs) != 1 {
		return "", status.Error(codes.Unauthenticated, "missing agent ID gRPC metadata")
	}
	return agentIDs[0], nil
}
