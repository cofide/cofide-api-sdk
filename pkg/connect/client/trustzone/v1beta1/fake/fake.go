// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"slices"

	trustbundlepb "github.com/cofide/cofide-api-sdk/gen/go/cofide/common/trust_bundle/v1beta1"
	trustzonepb "github.com/cofide/cofide-api-sdk/gen/go/cofide/trust_zone/v1beta1"
	fakepagination "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/client/pagination"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
	trustzonev1beta1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/trustzone/v1beta1"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type fakeTrustZoneClient struct {
	fake *fakeconnect.FakeConnect
}

// New instantiates a new TrustZoneClient for communication with a fake Connect API.
func New(fake *fakeconnect.FakeConnect) trustzonev1beta1.TrustZoneClient {
	return &fakeTrustZoneClient{
		fake: fake,
	}
}

func (c *fakeTrustZoneClient) CreateTrustZone(ctx context.Context, trustZone *trustzonepb.TrustZone) (*trustzonepb.TrustZone, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.Organizations[trustZone.GetOrgId()]; !ok {
		return nil, status.Error(codes.InvalidArgument, "organization not found")
	}
	if trustZone.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "name must be non-empty")
	}
	if trustZone.GetTrustDomain() == "" {
		return nil, status.Error(codes.InvalidArgument, "trust domain must be non-empty")
	}
	if c.findByName(trustZone.GetOrgId(), trustZone.GetName()) != nil {
		return nil, status.Error(codes.AlreadyExists, "trust zone name already exists in organization")
	}
	if c.findByTrustDomain(trustZone.GetTrustDomain()) != nil {
		return nil, status.Error(codes.AlreadyExists, "trust domain already exists")
	}

	trustZone = proto.CloneOf(trustZone)
	trustZone.Id = uuid.New().String()
	// Connect sets these from its configured trust bundle store, and ignores any values in the request.
	trustZone.BundleEndpointUrl = ""
	trustZone.JwtIssuerUrl = ""
	trustZone.BundleEndpointProfile = trustzonepb.BundleEndpointProfile_BUNDLE_ENDPOINT_PROFILE_UNSPECIFIED
	trustZone.CreatedAt = timestamppb.Now()
	trustZone.LastUpdatedAt = nil
	c.fake.TrustZones[trustZone.GetId()] = trustZone
	return proto.CloneOf(trustZone), nil
}

func (c *fakeTrustZoneClient) DeleteTrustZone(ctx context.Context, trustZoneID string) error {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.TrustZones[trustZoneID]; !ok {
		return status.Error(codes.NotFound, "trust zone not found")
	}
	delete(c.fake.TrustZones, trustZoneID)
	delete(c.fake.TrustZoneBundles, trustZoneID)
	delete(c.fake.AgentJoinTokens, trustZoneID)
	return nil
}

func (c *fakeTrustZoneClient) GetTrustZone(ctx context.Context, trustZoneID string) (*trustzonepb.TrustZone, error) {
	c.fake.Mu.RLock()
	defer c.fake.Mu.RUnlock()

	return found(c.fake.TrustZones[trustZoneID])
}

func (c *fakeTrustZoneClient) GetTrustZoneByName(ctx context.Context, orgID string, name string) (*trustzonepb.TrustZone, error) {
	c.fake.Mu.RLock()
	defer c.fake.Mu.RUnlock()

	return found(c.findByName(orgID, name))
}

func (c *fakeTrustZoneClient) GetTrustZoneByTrustDomain(ctx context.Context, trustDomain string) (*trustzonepb.TrustZone, error) {
	c.fake.Mu.RLock()
	defer c.fake.Mu.RUnlock()

	return found(c.findByTrustDomain(trustDomain))
}

func (c *fakeTrustZoneClient) ListTrustZones(ctx context.Context, filter *trustzonepb.ListTrustZonesRequest_Filter, requestPagination pagination.Pagination) ([]*trustzonepb.TrustZone, pagination.Pagination, error) {
	c.fake.Mu.RLock()
	defer c.fake.Mu.RUnlock()

	trustZones := []*trustzonepb.TrustZone{}
	for trustZone := range fakeconnect.SortedValues(c.fake.TrustZones) {
		if trustZoneMatches(trustZone, filter) {
			trustZones = append(trustZones, proto.CloneOf(trustZone))
		}
	}

	trustZones, pagination := fakepagination.Apply(
		trustZones,
		func(trustZone *trustzonepb.TrustZone) string { return trustZone.GetId() },
		requestPagination,
	)
	return trustZones, pagination, nil
}

func trustZoneMatches(trustZone *trustzonepb.TrustZone, filter *trustzonepb.ListTrustZonesRequest_Filter) bool {
	if len(filter.GetOrgIds()) > 0 && !slices.Contains(filter.GetOrgIds(), trustZone.GetOrgId()) {
		return false
	}
	return true
}

func (c *fakeTrustZoneClient) UpdateTrustZone(ctx context.Context, trustZone *trustzonepb.TrustZone, updateMask *trustzonepb.TrustZoneUpdateMask) (*trustzonepb.TrustZone, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	existing, ok := c.fake.TrustZones[trustZone.GetId()]
	if !ok {
		return nil, status.Error(codes.NotFound, "trust zone not found")
	}
	existing = proto.CloneOf(existing)

	// A mask selecting no fields updates nothing. No mask is a full replacement of the updatable fields, of which name
	// is the only one.
	if updateMask != nil && !updateMask.GetName() {
		return existing, nil
	}
	if trustZone.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "name must be non-empty")
	}
	if other := c.findByName(existing.GetOrgId(), trustZone.GetName()); other != nil && other.GetId() != existing.GetId() {
		return nil, status.Error(codes.AlreadyExists, "trust zone name already exists in organization")
	}
	existing.Name = trustZone.GetName()
	existing.LastUpdatedAt = timestamppb.Now()
	c.fake.TrustZones[existing.GetId()] = existing
	return proto.CloneOf(existing), nil
}

func (c *fakeTrustZoneClient) PublishTrustZoneBundle(ctx context.Context, bundle *trustbundlepb.Bundle) error {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	trustZone, err := c.trustZoneForBundle(bundle)
	if err != nil {
		return err
	}
	c.fake.TrustZoneBundles[trustZone.GetId()] = proto.CloneOf(bundle)
	return nil
}

func (c *fakeTrustZoneClient) RegisterCofideAgent(ctx context.Context, credential trustzonev1beta1.CofideAgentCredential, bundle *trustbundlepb.Bundle) error {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	return c.register(credential, bundle)
}

func (c *fakeTrustZoneClient) RegisterTrustZoneServer(ctx context.Context, credential trustzonev1beta1.TrustZoneServerCredential, bundle *trustbundlepb.Bundle) error {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	return c.register(credential, bundle)
}

// register checks a registration credential for the trust zone identified by the bundle's trust domain, then stores
// the bundle. A join token must be one issued for the trust zone, and is single use, so a successful registration
// removes it. The fake cannot verify a Kubernetes PSAT, so it accepts any non-empty one.
func (c *fakeTrustZoneClient) register(credential any, bundle *trustbundlepb.Bundle) error {
	trustZone, err := c.trustZoneForBundle(bundle)
	if err != nil {
		return err
	}

	switch credential := credential.(type) {
	case trustzonev1beta1.JoinToken:
		if credential == "" {
			return status.Error(codes.InvalidArgument, "join token must be non-empty")
		}
		tokens := c.fake.AgentJoinTokens[trustZone.GetId()]
		clusterID, ok := tokenKey(tokens, string(credential))
		if !ok {
			return status.Error(codes.Unauthenticated, "invalid join token")
		}
		delete(tokens, clusterID)
	case trustzonev1beta1.K8sPSAT:
		if credential == "" {
			return status.Error(codes.InvalidArgument, "k8s psat must be non-empty")
		}
	default:
		return status.Error(codes.InvalidArgument, "a join token or k8s psat must be provided")
	}

	c.fake.TrustZoneBundles[trustZone.GetId()] = proto.CloneOf(bundle)
	return nil
}

// tokenKey returns the cluster ID a join token was issued for, from a trust zone's join tokens keyed by cluster ID.
func tokenKey(tokens map[string]string, token string) (string, bool) {
	for clusterID, t := range tokens {
		if t == token {
			return clusterID, true
		}
	}
	return "", false
}

// trustZoneForBundle returns the trust zone identified by the bundle's trust domain.
func (c *fakeTrustZoneClient) trustZoneForBundle(bundle *trustbundlepb.Bundle) (*trustzonepb.TrustZone, error) {
	if bundle == nil {
		return nil, status.Error(codes.InvalidArgument, "bundle must be provided")
	}
	trustZone := c.findByTrustDomain(bundle.GetTrustDomain())
	if trustZone == nil {
		return nil, status.Error(codes.InvalidArgument, "no trust zone for bundle trust domain")
	}
	return trustZone, nil
}

func (c *fakeTrustZoneClient) findByName(orgID, name string) *trustzonepb.TrustZone {
	for trustZone := range fakeconnect.SortedValues(c.fake.TrustZones) {
		if trustZone.GetOrgId() == orgID && trustZone.GetName() == name {
			return trustZone
		}
	}
	return nil
}

func (c *fakeTrustZoneClient) findByTrustDomain(trustDomain string) *trustzonepb.TrustZone {
	for trustZone := range fakeconnect.SortedValues(c.fake.TrustZones) {
		if trustZone.GetTrustDomain() == trustDomain {
			return trustZone
		}
	}
	return nil
}

// found returns a copy of trustZone, or a NotFound error if it is nil.
func found(trustZone *trustzonepb.TrustZone) (*trustzonepb.TrustZone, error) {
	if trustZone == nil {
		return nil, status.Error(codes.NotFound, "trust zone not found")
	}
	return proto.CloneOf(trustZone), nil
}
