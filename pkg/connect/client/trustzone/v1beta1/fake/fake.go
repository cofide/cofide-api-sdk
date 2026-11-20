// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	trustbundlepb "github.com/cofide/cofide-api-sdk/gen/go/cofide/common/trust_bundle/v1beta1"
	trustzonepb "github.com/cofide/cofide-api-sdk/gen/go/cofide/trust_zone/v1beta1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
	trustzonev1beta1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/trustzone/v1beta1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// errUnimplemented is returned by every method, since the fake does not yet implement any of them.
var errUnimplemented = status.Error(codes.Unimplemented, "not implemented by the fake v1beta1 TrustZoneClient")

type fakeTrustZoneClient struct{}

// New instantiates a new TrustZoneClient for communication with a fake Connect API.
// Every method currently returns an Unimplemented error, so it does not yet use the fake Connect.
func New(_ *fakeconnect.FakeConnect) trustzonev1beta1.TrustZoneClient {
	return &fakeTrustZoneClient{}
}

func (c *fakeTrustZoneClient) CreateTrustZone(ctx context.Context, trustZone *trustzonepb.TrustZone) (*trustzonepb.TrustZone, error) {
	return nil, errUnimplemented
}

func (c *fakeTrustZoneClient) DeleteTrustZone(ctx context.Context, trustZoneID string) error {
	return errUnimplemented
}

func (c *fakeTrustZoneClient) GetTrustZone(ctx context.Context, trustZoneID string) (*trustzonepb.TrustZone, error) {
	return nil, errUnimplemented
}

func (c *fakeTrustZoneClient) GetTrustZoneByName(ctx context.Context, orgID string, name string) (*trustzonepb.TrustZone, error) {
	return nil, errUnimplemented
}

func (c *fakeTrustZoneClient) GetTrustZoneByTrustDomain(ctx context.Context, trustDomain string) (*trustzonepb.TrustZone, error) {
	return nil, errUnimplemented
}

func (c *fakeTrustZoneClient) ListTrustZones(ctx context.Context, filter *trustzonepb.ListTrustZonesRequest_Filter, requestPagination pagination.Pagination) ([]*trustzonepb.TrustZone, pagination.Pagination, error) {
	return nil, pagination.Pagination{}, errUnimplemented
}

func (c *fakeTrustZoneClient) UpdateTrustZone(ctx context.Context, trustZone *trustzonepb.TrustZone, updateMask *trustzonepb.TrustZoneUpdateMask) (*trustzonepb.TrustZone, error) {
	return nil, errUnimplemented
}

func (c *fakeTrustZoneClient) PublishTrustZoneBundle(ctx context.Context, bundle *trustbundlepb.Bundle) error {
	return errUnimplemented
}

func (c *fakeTrustZoneClient) RegisterCofideAgent(ctx context.Context, credential trustzonev1beta1.CofideAgentCredential, bundle *trustbundlepb.Bundle) error {
	return errUnimplemented
}

func (c *fakeTrustZoneClient) RegisterTrustZoneServer(ctx context.Context, credential trustzonev1beta1.TrustZoneServerCredential, bundle *trustbundlepb.Bundle) error {
	return errUnimplemented
}
