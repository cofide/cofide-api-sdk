// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"

	identitysvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/identity_service/v1alpha1"
	identitypb "github.com/cofide/cofide-api-sdk/gen/go/proto/identity/v1alpha1"
	"google.golang.org/grpc"
)

// IdentityClient is an interface for a client for the v1alpha1 version of the Connect IdentityService.
type IdentityClient interface {
	GetIdentity(ctx context.Context, identityID string) (*identitypb.Identity, error)
	ListIdentities(ctx context.Context, filter *identitysvcpb.ListIdentitiesRequest_Filter) ([]*identitypb.Identity, error)
}

type identityClient struct {
	identityClient identitysvcpb.IdentityServiceClient
}

// New instantiates a new IdentityClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) IdentityClient {
	return &identityClient{
		identityClient: identitysvcpb.NewIdentityServiceClient(conn),
	}
}

func (c *identityClient) GetIdentity(ctx context.Context, identityID string) (*identitypb.Identity, error) {
	resp, err := c.identityClient.GetIdentity(ctx, &identitysvcpb.GetIdentityRequest{IdentityId: identityID})
	if err != nil {
		return nil, err
	}
	return resp.Identity, nil
}

func (c *identityClient) ListIdentities(ctx context.Context, filter *identitysvcpb.ListIdentitiesRequest_Filter) ([]*identitypb.Identity, error) {
	resp, err := c.identityClient.ListIdentities(ctx, &identitysvcpb.ListIdentitiesRequest{Filter: filter})
	if err != nil {
		return nil, err
	}
	return resp.Identities, nil
}
