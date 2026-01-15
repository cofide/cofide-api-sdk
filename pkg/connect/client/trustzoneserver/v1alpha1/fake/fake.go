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
		if trustZoneServerMatches(trustZoneServer, filter) {
			trustZoneServers = append(trustZoneServers, clone(trustZoneServer))
		}
	}
	return trustZoneServers, nil
}

func trustZoneServerMatches(trustZoneServer *trustzoneserverpb.TrustZoneServer, filter *trustzoneserversvcpb.ListTrustZoneServersRequest_Filter) bool {
	if filter == nil {
		return true
	}
	if filter.TrustZoneId != "" && trustZoneServer.GetTrustZoneId() != filter.TrustZoneId {
		return false
	}
	if filter.ClusterId != "" && trustZoneServer.GetClusterId() != filter.ClusterId {
		return false
	}
	return true
}

func (c *fakeTrustZoneServerClient) UpdateTrustZoneServer(ctx context.Context, trustZoneServer *trustzoneserverpb.TrustZoneServer) (*trustzoneserverpb.TrustZoneServer, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.TrustZoneServers[trustZoneServer.GetId()]; !ok {
		return nil, status.Error(codes.NotFound, "trust zone server not found")
	}

	c.fake.TrustZoneServers[trustZoneServer.GetId()] = clone(trustZoneServer)
	return clone(trustZoneServer), nil
}

func clone(trustZoneServer *trustzoneserverpb.TrustZoneServer) *trustzoneserverpb.TrustZoneServer {
	return proto.Clone(trustZoneServer).(*trustzoneserverpb.TrustZoneServer)
}
