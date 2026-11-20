// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1beta1

import (
	"context"

	paginationpb "github.com/cofide/cofide-api-sdk/gen/go/cofide/common/pagination/v1beta1"
	trustbundlepb "github.com/cofide/cofide-api-sdk/gen/go/cofide/common/trust_bundle/v1beta1"
	trustzonepb "github.com/cofide/cofide-api-sdk/gen/go/cofide/trust_zone/v1beta1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
	"google.golang.org/grpc"
)

// TrustZoneClient is an interface for a gRPC client for the v1beta1 version of the Connect TrustZoneService.
type TrustZoneClient interface {
	CreateTrustZone(ctx context.Context, trustZone *trustzonepb.TrustZone) (*trustzonepb.TrustZone, error)
	DeleteTrustZone(ctx context.Context, trustZoneID string) error
	GetTrustZone(ctx context.Context, trustZoneID string) (*trustzonepb.TrustZone, error)
	GetTrustZoneByName(ctx context.Context, orgID string, name string) (*trustzonepb.TrustZone, error)
	GetTrustZoneByTrustDomain(ctx context.Context, trustDomain string) (*trustzonepb.TrustZone, error)
	ListTrustZones(ctx context.Context, filter *trustzonepb.ListTrustZonesRequest_Filter, pagination pagination.Pagination) ([]*trustzonepb.TrustZone, pagination.Pagination, error)
	UpdateTrustZone(ctx context.Context, trustZone *trustzonepb.TrustZone, updateMask *trustzonepb.TrustZoneUpdateMask) (*trustzonepb.TrustZone, error)
	PublishTrustZoneBundle(ctx context.Context, bundle *trustbundlepb.Bundle) error
	RegisterCofideAgent(ctx context.Context, credential CofideAgentCredential, bundle *trustbundlepb.Bundle) error
	RegisterTrustZoneServer(ctx context.Context, credential TrustZoneServerCredential, bundle *trustbundlepb.Bundle) error
}

type trustZoneClient struct {
	client trustzonepb.TrustZoneServiceClient
}

// New instantiates a new TrustZoneClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) TrustZoneClient {
	return &trustZoneClient{
		client: trustzonepb.NewTrustZoneServiceClient(conn),
	}
}

func (c *trustZoneClient) CreateTrustZone(ctx context.Context, trustZone *trustzonepb.TrustZone) (*trustzonepb.TrustZone, error) {
	resp, err := c.client.CreateTrustZone(ctx, &trustzonepb.CreateTrustZoneRequest{
		TrustZone: trustZone,
	})
	if err != nil {
		return nil, err
	}
	return resp.GetTrustZone(), nil
}

func (c *trustZoneClient) DeleteTrustZone(ctx context.Context, trustZoneID string) error {
	_, err := c.client.DeleteTrustZone(ctx, &trustzonepb.DeleteTrustZoneRequest{
		TrustZoneId: trustZoneID,
	})
	return err
}

func (c *trustZoneClient) GetTrustZone(ctx context.Context, trustZoneID string) (*trustzonepb.TrustZone, error) {
	return c.getTrustZone(ctx, &trustzonepb.GetTrustZoneRequest{
		Identifier: &trustzonepb.GetTrustZoneRequest_TrustZoneId{TrustZoneId: trustZoneID},
	})
}

// GetTrustZoneByName retrieves a trust zone by its name, which is unique within an organization.
func (c *trustZoneClient) GetTrustZoneByName(ctx context.Context, orgID string, name string) (*trustzonepb.TrustZone, error) {
	return c.getTrustZone(ctx, &trustzonepb.GetTrustZoneRequest{
		Identifier: &trustzonepb.GetTrustZoneRequest_NameWithOrg{
			NameWithOrg: &trustzonepb.GetTrustZoneRequest_NameWithOrgId{OrgId: orgID, Name: name},
		},
	})
}

// GetTrustZoneByTrustDomain retrieves a trust zone by its SPIFFE trust domain.
func (c *trustZoneClient) GetTrustZoneByTrustDomain(ctx context.Context, trustDomain string) (*trustzonepb.TrustZone, error) {
	return c.getTrustZone(ctx, &trustzonepb.GetTrustZoneRequest{
		Identifier: &trustzonepb.GetTrustZoneRequest_TrustDomain{TrustDomain: trustDomain},
	})
}

func (c *trustZoneClient) getTrustZone(ctx context.Context, req *trustzonepb.GetTrustZoneRequest) (*trustzonepb.TrustZone, error) {
	resp, err := c.client.GetTrustZone(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.GetTrustZone(), nil
}

func (c *trustZoneClient) ListTrustZones(ctx context.Context, filter *trustzonepb.ListTrustZonesRequest_Filter, requestPagination pagination.Pagination) ([]*trustzonepb.TrustZone, pagination.Pagination, error) {
	resp, err := c.client.ListTrustZones(ctx, &trustzonepb.ListTrustZonesRequest{
		Filter:     filter,
		Pagination: &paginationpb.PageRequest{PageSize: requestPagination.PageSize, PageToken: requestPagination.Token},
	})
	if err != nil {
		return nil, pagination.Pagination{}, err
	}
	return resp.GetTrustZones(), pagination.Pagination{PageSize: requestPagination.PageSize, Token: resp.GetPagination().GetNextPageToken()}, nil
}

// UpdateTrustZone updates a trust zone, identified by its ID. A nil updateMask requests a full replacement of all
// updatable fields.
func (c *trustZoneClient) UpdateTrustZone(ctx context.Context, trustZone *trustzonepb.TrustZone, updateMask *trustzonepb.TrustZoneUpdateMask) (*trustzonepb.TrustZone, error) {
	resp, err := c.client.UpdateTrustZone(ctx, &trustzonepb.UpdateTrustZoneRequest{
		TrustZone:  trustZone,
		UpdateMask: updateMask,
	})
	if err != nil {
		return nil, err
	}
	return resp.GetTrustZone(), nil
}

func (c *trustZoneClient) PublishTrustZoneBundle(ctx context.Context, bundle *trustbundlepb.Bundle) error {
	_, err := c.client.PublishTrustZoneBundle(ctx, &trustzonepb.PublishTrustZoneBundleRequest{
		Bundle: bundle,
	})
	return err
}

func (c *trustZoneClient) RegisterCofideAgent(ctx context.Context, credential CofideAgentCredential, bundle *trustbundlepb.Bundle) error {
	req := &trustzonepb.RegisterCofideAgentRequest{Bundle: bundle}
	if credential != nil {
		credential.applyToCofideAgent(req)
	}
	_, err := c.client.RegisterCofideAgent(ctx, req)
	return err
}

func (c *trustZoneClient) RegisterTrustZoneServer(ctx context.Context, credential TrustZoneServerCredential, bundle *trustbundlepb.Bundle) error {
	req := &trustzonepb.RegisterTrustZoneServerRequest{Bundle: bundle}
	if credential != nil {
		credential.applyToTrustZoneServer(req)
	}
	_, err := c.client.RegisterTrustZoneServer(ctx, req)
	return err
}
