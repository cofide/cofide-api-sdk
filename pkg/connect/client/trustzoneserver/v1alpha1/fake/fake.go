// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	trustzoneserversvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/trust_zone_server_service/v1alpha1"
	trustzoneserverpb "github.com/cofide/cofide-api-sdk/gen/go/proto/trust_zone_server/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	trustzoneserverv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/trustzoneserver/v1alpha1"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type fakeTrustZoneServerClient struct {
	fake *fakeconnect.FakeConnect
}

// New instantiates a new TrustZoneServerClient for communication with a fake Connect API.
func New(fake *fakeconnect.FakeConnect) trustzoneserverv1alpha1.TrustZoneServerClient {
	return &fakeTrustZoneServerClient{
		fake: fake,
	}
}

func (c *fakeTrustZoneServerClient) CreateTrustZoneServer(ctx context.Context, trustZoneServer *trustzoneserverpb.TrustZoneServer) (*trustzoneserverpb.TrustZoneServer, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	trustZoneServer = clone(trustZoneServer)
	trustZoneServer.Id = uuid.New().String()
	c.fake.TrustZoneServers[trustZoneServer.GetId()] = trustZoneServer
	return clone(trustZoneServer), nil
}

func (c *fakeTrustZoneServerClient) DestroyTrustZoneServer(ctx context.Context, trustZoneServerID string) error {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.TrustZoneServers[trustZoneServerID]; !ok {
		return status.Error(codes.NotFound, "trust zone server not found")
	}
	delete(c.fake.TrustZoneServers, trustZoneServerID)
	return nil
}

func (c *fakeTrustZoneServerClient) GetTrustZoneServer(ctx context.Context, trustZoneServerID string) (*trustzoneserverpb.TrustZoneServer, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	trustZoneServer, ok := c.fake.TrustZoneServers[trustZoneServerID]
	if !ok {
		return nil, status.Error(codes.NotFound, "")
	}
	return clone(trustZoneServer), nil
}

func (c *fakeTrustZoneServerClient) ListTrustZoneServers(ctx context.Context, filter *trustzoneserversvcpb.ListTrustZoneServersRequest_Filter) ([]*trustzoneserverpb.TrustZoneServer, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	trustZoneServers := []*trustzoneserverpb.TrustZoneServer{}
	for _, trustZoneServer := range c.fake.TrustZoneServers {
		if c.trustZoneServerMatches(trustZoneServer, filter) {
			trustZoneServers = append(trustZoneServers, clone(trustZoneServer))
		}
	}
	return trustZoneServers, nil
}

func (c *fakeTrustZoneServerClient) trustZoneServerMatches(trustZoneServer *trustzoneserverpb.TrustZoneServer, filter *trustzoneserversvcpb.ListTrustZoneServersRequest_Filter) bool {
	if filter == nil {
		return true
	}
	if filter.OrgId != "" {
		// Org ID field should be read only for clients, so it may not be set on the stored trust zone server
		// If it is, use that to filter by
		// Otherwise filter by the org ID of a linked trust zone (if one exists)
		serverOrgId := trustZoneServer.GetOrgId()
		if serverOrgId == "" {
			trustZone, ok := c.fake.TrustZones[trustZoneServer.GetTrustZoneId()]
			if ok {
				serverOrgId = trustZone.GetOrgId()
			}
		}
		if serverOrgId != filter.OrgId {
			return false
		}
	}
	if filter.TrustZoneId != "" && trustZoneServer.GetTrustZoneId() != filter.TrustZoneId {
		return false
	}
	if filter.ClusterId != "" && trustZoneServer.GetClusterId() != filter.ClusterId {
		return false
	}
	return true
}

func (c *fakeTrustZoneServerClient) UpdateTrustZoneServer(ctx context.Context, trustZoneServer *trustzoneserverpb.TrustZoneServer, updateMask *trustzoneserversvcpb.UpdateTrustZoneServerRequest_UpdateMask) (*trustzoneserverpb.TrustZoneServer, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	existing, ok := c.fake.TrustZoneServers[trustZoneServer.GetId()]
	if !ok {
		return nil, status.Error(codes.NotFound, "trust zone server not found")
	}

	var new *trustzoneserverpb.TrustZoneServer
	if updateMask == nil {
		new = clone(trustZoneServer)
	} else {
		new = clone(existing)
		// Once update of fields is supported, handle the ones set in the update mask and update here
	}

	// Existing may have been partially updated or entirely replaced, so explicitly set it to cover both cases
	c.fake.TrustZoneServers[trustZoneServer.GetId()] = new
	return clone(new), nil
}

func clone(trustZoneServer *trustzoneserverpb.TrustZoneServer) *trustzoneserverpb.TrustZoneServer {
	return proto.Clone(trustZoneServer).(*trustzoneserverpb.TrustZoneServer)
}
