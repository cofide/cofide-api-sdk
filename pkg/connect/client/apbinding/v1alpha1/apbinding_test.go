// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"
	"testing"

	apbindingpb "github.com/cofide/cofide-api-sdk/gen/go/proto/ap_binding/v1alpha1"
	apbindingsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/ap_binding_service/v1alpha1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	fakeAPBindingID         = "fake-ap-binding-id"
	fakeAPBindingName       = "fake-ap-binding-name"
	fakeTrustZoneID         = "fake-tz-id"
	fakeAttestationPolicyID = "fake-ap-id"
)

func TestAPBindingClient_implementsMethods(t *testing.T) {
	test.AssertClientImplementsService(t, &apBindingClient{}, apbindingsvcpb.APBindingService_ServiceDesc)
}

// TestAPBindingClient_Unimplemented tests APBindingClient against an unimplemented server.
// This ensures that all errors returned are not wrapped can be converted to a gRPC Status using Status.Convert.
func TestAPBindingClient_Unimplemented(t *testing.T) {
	server := test.NewTestServer(t)
	apbindingsvcpb.RegisterAPBindingServiceServer(server.Server, &apbindingsvcpb.UnimplementedAPBindingServiceServer{})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	token, err := client.CreateAPBinding(ctx, nil)
	test.RequireUnimplemented(t, err)
	assert.Empty(t, token)

	err = client.DestroyAPBinding(ctx, "")
	test.RequireUnimplemented(t, err)

	binding, err := client.GetAPBinding(ctx, "")
	test.RequireUnimplemented(t, err)
	assert.Nil(t, binding)

	bindings, err := client.ListAPBindings(ctx, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, bindings)

	binding, err = client.UpdateAPBinding(ctx, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, binding)
}

func TestAPBindingClient(t *testing.T) {
	server := test.NewTestServer(t)
	apbindingsvcpb.RegisterAPBindingServiceServer(server.Server, &fakeAPBindingService{t: t})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	binding := fakeAPBinding()

	createdAPBinding, err := client.CreateAPBinding(ctx, binding)
	require.NoError(t, err)
	assert.EqualExportedValues(t, binding, createdAPBinding)

	err = client.DestroyAPBinding(ctx, binding.GetId())
	require.NoError(t, err)

	gotAPBinding, err := client.GetAPBinding(ctx, binding.GetId())
	require.NoError(t, err)
	assert.Equal(t, binding.GetId(), gotAPBinding.GetId())

	filter := &apbindingsvcpb.ListAPBindingsRequest_Filter{TrustZoneId: test.PtrOf(fakeTrustZoneID)}
	bindings, err := client.ListAPBindings(ctx, filter)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*apbindingpb.APBinding{fakeAPBinding()}, bindings)

	updatedAPBinding, err := client.UpdateAPBinding(ctx, binding)
	require.NoError(t, err)
	assert.EqualExportedValues(t, binding, updatedAPBinding)
}

// fakeAPBindingService provides a very simple fake APBindingService implementation with canned responses.
type fakeAPBindingService struct {
	t *testing.T
}

func (f *fakeAPBindingService) CreateAPBinding(ctx context.Context, req *apbindingsvcpb.CreateAPBindingRequest) (*apbindingsvcpb.CreateAPBindingResponse, error) {
	assert.EqualExportedValues(f.t, fakeAPBinding(), req.Binding)
	return &apbindingsvcpb.CreateAPBindingResponse{Binding: req.Binding}, nil
}

func (f *fakeAPBindingService) DestroyAPBinding(ctx context.Context, req *apbindingsvcpb.DestroyAPBindingRequest) (*apbindingsvcpb.DestroyAPBindingResponse, error) {
	assert.Equal(f.t, fakeAPBindingID, req.GetBindingId())
	return &apbindingsvcpb.DestroyAPBindingResponse{}, nil
}

func (f *fakeAPBindingService) GetAPBinding(ctx context.Context, req *apbindingsvcpb.GetAPBindingRequest) (*apbindingsvcpb.GetAPBindingResponse, error) {
	assert.Equal(f.t, fakeAPBindingID, req.GetBindingId())
	return &apbindingsvcpb.GetAPBindingResponse{Binding: fakeAPBinding()}, nil
}

func (f *fakeAPBindingService) ListAPBindings(ctx context.Context, req *apbindingsvcpb.ListAPBindingsRequest) (*apbindingsvcpb.ListAPBindingsResponse, error) {
	assert.Equal(f.t, fakeTrustZoneID, req.Filter.GetTrustZoneId())
	bindings := []*apbindingpb.APBinding{fakeAPBinding()}
	return &apbindingsvcpb.ListAPBindingsResponse{Bindings: bindings}, nil
}

func (f *fakeAPBindingService) UpdateAPBinding(ctx context.Context, req *apbindingsvcpb.UpdateAPBindingRequest) (*apbindingsvcpb.UpdateAPBindingResponse, error) {
	assert.EqualExportedValues(f.t, fakeAPBinding(), req.Binding)
	return &apbindingsvcpb.UpdateAPBindingResponse{Binding: fakeAPBinding()}, nil
}

func fakeAPBinding() *apbindingpb.APBinding {
	return &apbindingpb.APBinding{
		Id:          test.PtrOf(fakeAPBindingID),
		TrustZoneId: test.PtrOf(fakeTrustZoneID),
		PolicyId:    test.PtrOf(fakeAttestationPolicyID),
	}
}
