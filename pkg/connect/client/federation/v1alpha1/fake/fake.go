// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	federationsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/federation_service/v1alpha1"
	federationpb "github.com/cofide/cofide-api-sdk/gen/go/proto/federation/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	federationv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/federation/v1alpha1"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type fakeFederationClient struct {
	fake *fakeconnect.FakeConnect
}

// New instantiates a new FederationClient for communication with a fake Connect API.
func New(fake *fakeconnect.FakeConnect) federationv1alpha1.FederationClient {
	return &fakeFederationClient{
		fake: fake,
	}
}

func (c *fakeFederationClient) CreateFederation(ctx context.Context, federation *federationpb.Federation) (*federationpb.Federation, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	id := uuid.New().String()
	federation = proto.Clone(federation).(*federationpb.Federation)
	federation.Id = &id
	c.fake.Federations[federation.GetId()] = federation
	return proto.Clone(federation).(*federationpb.Federation), nil
}

func (c *fakeFederationClient) GetFederation(ctx context.Context, federationID string) (*federationpb.Federation, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	federation, ok := c.fake.Federations[federationID]
	if !ok {
		return nil, status.Error(codes.NotFound, "")
	}
	return proto.Clone(federation).(*federationpb.Federation), nil
}

func (c *fakeFederationClient) ListFederations(ctx context.Context, filter *federationsvcpb.ListFederationsRequest_Filter) ([]*federationpb.Federation, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	federations := []*federationpb.Federation{}
	for _, federation := range c.fake.Federations {
		if federationMatches(federation, filter) {
			federations = append(federations, proto.Clone(federation).(*federationpb.Federation))
		}
	}
	return federations, nil
}

func federationMatches(federation *federationpb.Federation, filter *federationsvcpb.ListFederationsRequest_Filter) bool {
	if filter == nil {
		return true
	}
	if filter.OrgId != nil && federation.GetOrgId() != *filter.OrgId {
		return false
	}
	if filter.RemoteTrustZoneId != nil && federation.GetRemoteTrustZoneId() != *filter.RemoteTrustZoneId {
		return false
	}
	if filter.TrustZoneId != nil && federation.GetTrustZoneId() != *filter.TrustZoneId {
		return false
	}
	return true
}

func (c *fakeFederationClient) DestroyFederation(ctx context.Context, federationID string) error {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()
	_, ok := c.fake.Federations[federationID]
	if !ok {
		return status.Error(codes.NotFound, "")
	}
	delete(c.fake.Federations, federationID)
	return nil
}
