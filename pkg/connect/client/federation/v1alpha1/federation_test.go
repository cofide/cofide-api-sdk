// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"
	"testing"

	federationsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/federation_service/v1alpha1"
	federationpb "github.com/cofide/cofide-api-sdk/gen/go/proto/federation/v1alpha1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
)

const (
	fakeFederationID      = "fake-federation-id"
	fakeTrustZoneId       = "fake-trust-zone-id"
	fakeRemoteTrustZoneId = "fake-remote-trust-zone-id"
)

func TestFederationClient_Unimplemented(t *testing.T) {
	server := test.NewTestServer(t)
	federationsvcpb.RegisterFederationServiceServer(server.Server, &federationsvcpb.UnimplementedFederationServiceServer{})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	federation, err := client.CreateFederation(ctx, nil)
	test.RequireUnimplemented(t, err)
	if federation != nil {
		t.Errorf("expected nil federation, got %v", federation)
	}

	federation, err = client.GetFederation(ctx, "")
	test.RequireUnimplemented(t, err)
	if federation != nil {
		t.Errorf("expected nil federation, got %v", federation)
	}

	federations, err := client.ListFederations(ctx, &federationsvcpb.ListFederationsRequest_Filter{
		TrustZoneId: test.PtrOf(fakeTrustZoneId),
	})
	test.RequireUnimplemented(t, err)
	if federations != nil {
		t.Errorf("expected nil federations, got %v", federations)
	}

	err = client.DestroyFederation(ctx, "")
	test.RequireUnimplemented(t, err)
}

func TestFederationClient(t *testing.T) {
	server := test.NewTestServer(t)
	federationsvcpb.RegisterFederationServiceServer(server.Server, &fakeFederationService{t: t})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	federation := fakeFederation()

	createdFederation, err := client.CreateFederation(ctx, federation)
	if err != nil {
		t.Fatalf("CreateFederation failed: %v", err)
	}
	if !equalFederations(federation, createdFederation) {
		t.Errorf("expected %v, got %v", federation, createdFederation)
	}

	gotFederation, err := client.GetFederation(ctx, fakeFederationID)
	if err != nil {
		t.Fatalf("GetFederation failed: %v", err)
	}
	if !equalFederations(federation, gotFederation) {
		t.Errorf("expected %v, got %v", federation, gotFederation)
	}

	filter := &federationsvcpb.ListFederationsRequest_Filter{
		TrustZoneId: test.PtrOf(fakeTrustZoneId),
	}
	federations, err := client.ListFederations(ctx, filter)
	if err != nil {
		t.Fatalf("ListFederations failed: %v", err)
	}
	if len(federations) != 1 || !equalFederations(federation, federations[0]) {
		t.Errorf("expected %v, got %v", []*federationpb.Federation{federation}, federations)
	}

	err = client.DestroyFederation(ctx, fakeFederationID)
	if err != nil {
		t.Fatalf("DestroyFederation failed: %v", err)
	}
}

type fakeFederationService struct {
	t *testing.T
}

func (f *fakeFederationService) CreateFederation(ctx context.Context, req *federationsvcpb.CreateFederationRequest) (*federationsvcpb.CreateFederationResponse, error) {
	if !equalFederations(fakeFederation(), req.Federation) {
		f.t.Errorf("expected %v, got %v", fakeFederation(), req.Federation)
	}
	return &federationsvcpb.CreateFederationResponse{Federation: req.Federation}, nil
}

func (f *fakeFederationService) GetFederation(ctx context.Context, req *federationsvcpb.GetFederationRequest) (*federationsvcpb.GetFederationResponse, error) {
	if req.FederationId != fakeFederationID {
		f.t.Errorf("expected %v, got %v", fakeFederationID, req.FederationId)
	}
	return &federationsvcpb.GetFederationResponse{Federation: fakeFederation()}, nil
}

func (f *fakeFederationService) ListFederations(ctx context.Context, req *federationsvcpb.ListFederationsRequest) (*federationsvcpb.ListFederationsResponse, error) {
	if req.Filter.GetTrustZoneId() != fakeTrustZoneId {
		f.t.Errorf("expected %v, got %v", fakeTrustZoneId, req.Filter.GetTrustZoneId())
	}
	return &federationsvcpb.ListFederationsResponse{Federations: []*federationpb.Federation{fakeFederation()}}, nil
}

func (f *fakeFederationService) DestroyFederation(ctx context.Context, req *federationsvcpb.DestroyFederationRequest) (*federationsvcpb.DestroyFederationResponse, error) {
	if req.GetFederationId() != fakeFederationID {
		f.t.Errorf("expected %v, got %v", fakeFederationID, req.GetFederationId())
	}
	return &federationsvcpb.DestroyFederationResponse{}, nil
}

func fakeFederation() *federationpb.Federation {
	return &federationpb.Federation{
		Id:                test.PtrOf(fakeFederationID),
		TrustZoneId:       test.PtrOf(fakeTrustZoneId),
		RemoteTrustZoneId: test.PtrOf(fakeRemoteTrustZoneId),
	}
}

func equalFederations(a, b *federationpb.Federation) bool {
	return *a.Id == *b.Id && *a.TrustZoneId == *b.TrustZoneId && *a.RemoteTrustZoneId == *b.RemoteTrustZoneId
}
