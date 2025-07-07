// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	agentpb "github.com/cofide/cofide-api-sdk/gen/go/proto/agent/v1alpha1"
	clusterpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cluster/v1alpha1"
	trustzonesvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/trust_zone_service/v1alpha1"
	trustzonepb "github.com/cofide/cofide-api-sdk/gen/go/proto/trust_zone/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	trustzonev1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/trustzone/v1alpha1"
	"github.com/google/uuid"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"google.golang.org/grpc/codes"
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

// DEPRECATED: Agent join token creation moved to AgentService.CreateAgentJoinToken.
// Cluster creation to be moved to ClusterService.CreateCluster.
func (c *fakeTrustZoneClient) RegisterCluster(ctx context.Context, trustZoneID string, cluster *clusterpb.Cluster) (string, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.TrustZones[trustZoneID]; !ok {
		return "", status.Error(codes.InvalidArgument, "")
	}
	cluster = cloneCluster(cluster)
	id := uuid.New().String()
	cluster.Id = &id
	cluster.TrustZoneId = &trustZoneID
	c.fake.Clusters[id] = cluster
	if c.fake.AgentJoinTokens[trustZoneID] == nil {
		c.fake.AgentJoinTokens[trustZoneID] = make(map[string]string)
	}
	token := uuid.New().String()
	c.fake.AgentJoinTokens[trustZoneID][cluster.GetId()] = token
	return token, nil
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

	if _, ok := c.fake.Clusters[server.ClusterId]; !ok {
		return status.Error(codes.InvalidArgument, "invalid cluster")
	}

	return nil
}


func clone(trustZone *trustzonepb.TrustZone) *trustzonepb.TrustZone {
	return proto.Clone(trustZone).(*trustzonepb.TrustZone)
}

func cloneCluster(cluster *clusterpb.Cluster) *clusterpb.Cluster {
	return proto.Clone(cluster).(*clusterpb.Cluster)
}

func cloneBundle(bundle *types.Bundle) *types.Bundle {
	return proto.Clone(bundle).(*types.Bundle)
}
