// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	federationpb "github.com/cofide/cofide-api-sdk/gen/go/cofide/federation/v1beta1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	federationv1beta1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/federation/v1beta1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// errUnimplemented is returned by every method, since the fake does not yet implement any of them.
var errUnimplemented = status.Error(codes.Unimplemented, "not implemented by the fake v1beta1 FederationClient")

type fakeFederationClient struct{}

// New instantiates a new FederationClient for communication with a fake Connect API.
// Every method currently returns an Unimplemented error, so it does not yet use the fake Connect.
func New(_ *fakeconnect.FakeConnect) federationv1beta1.FederationClient {
	return &fakeFederationClient{}
}

func (c *fakeFederationClient) CreateFederation(ctx context.Context, federation *federationpb.Federation) (*federationpb.Federation, error) {
	return nil, errUnimplemented
}

func (c *fakeFederationClient) DeleteFederation(ctx context.Context, federationID string) error {
	return errUnimplemented
}

func (c *fakeFederationClient) GetFederation(ctx context.Context, federationID string) (*federationpb.Federation, error) {
	return nil, errUnimplemented
}

func (c *fakeFederationClient) GetFederationByTrustZones(ctx context.Context, trustZoneID string, remoteTrustZoneID string) (*federationpb.Federation, error) {
	return nil, errUnimplemented
}

func (c *fakeFederationClient) ListFederations(ctx context.Context, filter *federationpb.ListFederationsRequest_Filter, requestPagination pagination.Pagination) ([]*federationpb.Federation, pagination.Pagination, error) {
	return nil, pagination.Pagination{}, errUnimplemented
}
