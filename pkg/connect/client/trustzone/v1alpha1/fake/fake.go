// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	agentpb "github.com/cofide/cofide-api-sdk/gen/go/proto/agent/v1alpha1"
	trustzonesvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/trust_zone_service/v1alpha1"
	trustzonepb "github.com/cofide/cofide-api-sdk/gen/go/proto/trust_zone/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	trustzonev1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/trustzone/v1alpha1"
	"github.com/google/uuid"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type fakeTrustZoneClient struct {
	fake *fakeconnect.FakeConnect
}

// New instantiates a new TrustZoneClient for communication with a fake Connect API.
func New(fake *fakeconnect.FakeConnect) trustzonev1alpha1.TrustZoneClient {
	return &fakeTrustZoneClient{
		fake: fake,
	}
}

func (c *fakeTrustZoneClient) CreateTrustZone(ctx context.Context, trustZone *trustzonepb.TrustZone) (*trustzonepb.TrustZone, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	id := uuid.New().String()
	trustZone = clone(trustZone)
	trustZone.Id = &id
	c.fake.TrustZones[trustZone.GetId()] = trustZone
	return clone(trustZone), nil
}

func (c *fakeTrustZoneClient) DestroyTrustZone(ctx context.Context, trustZoneID string) error {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.TrustZones[trustZoneID]; !ok {
		return status.Error(codes.NotFound, "trust zone not found")
	}
	delete(c.fake.TrustZones, trustZoneID)
	return nil
}

func (c *fakeTrustZoneClient) GetTrustZone(ctx context.Context, trustZoneID string) (*trustzonepb.TrustZone, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	trustZone, ok := c.fake.TrustZones[trustZoneID]
	if !ok {
		return nil, status.Error(codes.NotFound, "")
	}
	return clone(trustZone), nil
}

func (c *fakeTrustZoneClient) ListTrustZones(ctx context.Context, filter *trustzonesvcpb.ListTrustZonesRequest_Filter) ([]*trustzonepb.TrustZone, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	trustZones := []*trustzonepb.TrustZone{}
	for _, trustZone := range c.fake.TrustZones {
		if trustZoneMatches(trustZone, filter) {
			trustZones = append(trustZones, clone(trustZone))
		}
	}
	return trustZones, nil
}

func trustZoneMatches(trustZone *trustzonepb.TrustZone, filter *trustzonesvcpb.ListTrustZonesRequest_Filter) bool {
	if filter == nil {
		return true
	}
	if filter.Name != nil && trustZone.GetName() != *filter.Name {
		return false
	}
	if filter.OrgId != nil && trustZone.GetOrgId() != *filter.OrgId {
		return false
	}
	if filter.TrustDomain != nil && trustZone.GetTrustDomain() != *filter.TrustDomain {
		return false
	}
	return true
}

func (c *fakeTrustZoneClient) UpdateTrustZone(ctx context.Context, trustZone *trustzonepb.TrustZone) (*trustzonepb.TrustZone, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.TrustZones[trustZone.GetId()]; !ok {
		return nil, status.Error(codes.NotFound, "trust zone not found")
	}

	c.fake.TrustZones[trustZone.GetId()] = clone(trustZone)
	return clone(trustZone), nil
}

func (c *fakeTrustZoneClient) RegisterAgent(ctx context.Context, agent *trustzonesvcpb.Agent, token string, bundle *types.Bundle) (string, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.TrustZones[agent.TrustZoneId]; !ok {
		return "", status.Error(codes.InvalidArgument, "invalid trust zone")
	}
	if _, ok := c.fake.Clusters[agent.ClusterId]; !ok {
		return "", status.Error(codes.InvalidArgument, "invalid cluster")
	}
	id := uuid.New().String()
	newAgent := agentpb.Agent{
		Id:          &id,
		ClusterId:   &agent.ClusterId,
		TrustZoneId: &agent.TrustZoneId,
	}
	c.fake.Agents[id] = &newAgent
	c.fake.TrustZoneBundles[agent.GetTrustZoneId()] = cloneBundle(bundle)
	return id, nil
}

func (c *fakeTrustZoneClient) RegisterTrustZoneServer(ctx context.Context, server *trustzonesvcpb.TrustZoneServer, bundle *types.Bundle) error {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	cluster, ok := c.fake.Clusters[server.ClusterId]
	if !ok {
		return status.Error(codes.InvalidArgument, "invalid cluster")
	}

	c.fake.TrustZoneBundles[cluster.GetTrustZoneId()] = cloneBundle(bundle)
	return nil
}

func (c *fakeTrustZoneClient) UpdateTrustZoneBundle(ctx context.Context, bundle *types.Bundle) error {
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

func (c *fakeTrustZoneClient) UpdateManagedTrustZoneBundle(ctx context.Context, trustZoneID string, bundle *types.Bundle) error {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if err := c.fake.ValidateTrustZone(trustZoneID); err != nil {
		return err
	}

	c.fake.TrustZoneBundles[trustZoneID] = cloneBundle(bundle)
	return nil
}

func clone(trustZone *trustzonepb.TrustZone) *trustzonepb.TrustZone {
	return proto.Clone(trustZone).(*trustzonepb.TrustZone)
}

func cloneBundle(bundle *types.Bundle) *types.Bundle {
	return proto.Clone(bundle).(*types.Bundle)
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
