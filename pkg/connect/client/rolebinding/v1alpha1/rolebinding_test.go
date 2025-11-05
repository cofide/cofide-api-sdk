// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"
	"testing"

	rolebindingsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/role_binding_service/v1alpha1"
	rolebindingpb "github.com/cofide/cofide-api-sdk/gen/go/proto/role_binding/v1alpha1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	fakeRoleBindingID = "fake-rb-id"
	fakeRoleID        = "fake-role-id"
	fakeClaimName     = "fake-claim-name"
	fakeClaimValue    = "fake-claim-value"
	fakeResourceType  = rolebindingpb.ResourceType_RESOURCE_TYPE_AP_BINDING
	fakeResourceID    = "fake-apb-id"
)

// TestRoleBindingClient_Unimplemented tests RoleBindingClient against an unimplemented server.
// This ensures that all errors returned are not wrapped can be converted to a gRPC Status using Status.Convert.
func TestRoleBindingClient_Unimplemented(t *testing.T) {
	server := test.NewTestServer(t)
	rolebindingsvcpb.RegisterRoleBindingServiceServer(server.Server, &rolebindingsvcpb.UnimplementedRoleBindingServiceServer{})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	roleBinding, err := client.CreateRoleBinding(ctx, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, roleBinding)

	err = client.DestroyRoleBinding(ctx, "")
	test.RequireUnimplemented(t, err)

	roleBinding, err = client.GetRoleBinding(ctx, "")
	test.RequireUnimplemented(t, err)
	assert.Nil(t, roleBinding)

	roleBindings, err := client.ListRoleBindings(ctx, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, roleBindings)

	roleBinding, err = client.UpdateRoleBinding(ctx, nil)
	test.RequireUnimplemented(t, err)
	assert.Nil(t, roleBinding)
}

func TestRoleBindingClient(t *testing.T) {
	server := test.NewTestServer(t)
	rolebindingsvcpb.RegisterRoleBindingServiceServer(server.Server, &fakeRoleBindingService{t: t})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	roleBinding := fakeRoleBinding()

	createdRoleBinding, err := client.CreateRoleBinding(ctx, roleBinding)
	require.NoError(t, err)
	assert.EqualExportedValues(t, roleBinding, createdRoleBinding)

	err = client.DestroyRoleBinding(ctx, fakeRoleBindingID)
	require.NoError(t, err)

	gotRoleBinding, err := client.GetRoleBinding(ctx, fakeRoleBindingID)
	require.NoError(t, err)
	assert.EqualExportedValues(t, roleBinding, gotRoleBinding)

	filter := &rolebindingsvcpb.ListRoleBindingsRequest_Filter{RoleId: test.PtrOf(fakeRoleID)}
	roleBindings, err := client.ListRoleBindings(ctx, filter)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*rolebindingpb.RoleBinding{roleBinding}, roleBindings)

	updatedRoleBinding, err := client.UpdateRoleBinding(ctx, roleBinding)
	require.NoError(t, err)
	assert.EqualExportedValues(t, roleBinding, updatedRoleBinding)
}

type fakeRoleBindingService struct {
	t *testing.T
}

func (f *fakeRoleBindingService) CreateRoleBinding(ctx context.Context, req *rolebindingsvcpb.CreateRoleBindingRequest) (*rolebindingsvcpb.CreateRoleBindingResponse, error) {
	assert.EqualExportedValues(f.t, fakeRoleBinding(), req.RoleBinding)
	return &rolebindingsvcpb.CreateRoleBindingResponse{RoleBinding: req.RoleBinding}, nil
}

func (f *fakeRoleBindingService) DestroyRoleBinding(ctx context.Context, req *rolebindingsvcpb.DestroyRoleBindingRequest) (*rolebindingsvcpb.DestroyRoleBindingResponse, error) {
	assert.Equal(f.t, fakeRoleBindingID, req.GetRoleBindingId())
	return &rolebindingsvcpb.DestroyRoleBindingResponse{}, nil
}

func (f *fakeRoleBindingService) GetRoleBinding(ctx context.Context, req *rolebindingsvcpb.GetRoleBindingRequest) (*rolebindingsvcpb.GetRoleBindingResponse, error) {
	assert.Equal(f.t, fakeRoleBindingID, req.GetRoleBindingId())
	return &rolebindingsvcpb.GetRoleBindingResponse{RoleBinding: fakeRoleBinding()}, nil
}

func (f *fakeRoleBindingService) ListRoleBindings(ctx context.Context, req *rolebindingsvcpb.ListRoleBindingsRequest) (*rolebindingsvcpb.ListRoleBindingsResponse, error) {
	assert.Equal(f.t, fakeRoleID, req.Filter.GetRoleId())
	roleBindings := []*rolebindingpb.RoleBinding{fakeRoleBinding()}
	return &rolebindingsvcpb.ListRoleBindingsResponse{RoleBindings: roleBindings}, nil
}

func (f *fakeRoleBindingService) UpdateRoleBinding(ctx context.Context, req *rolebindingsvcpb.UpdateRoleBindingRequest) (*rolebindingsvcpb.UpdateRoleBindingResponse, error) {
	assert.EqualExportedValues(f.t, fakeRoleBinding(), req.RoleBinding)
	return &rolebindingsvcpb.UpdateRoleBindingResponse{RoleBinding: req.RoleBinding}, nil
}

func fakeRoleBinding() *rolebindingpb.RoleBinding {
	return &rolebindingpb.RoleBinding{
		Id:     fakeRoleBindingID,
		RoleId: fakeRoleID,
		Principal: &rolebindingpb.RoleBinding_User{
			User: &rolebindingpb.User{
				ClaimReq: &rolebindingpb.ClaimReq{
					Condition: &rolebindingpb.ClaimReq_StringEqualsOp{
						StringEqualsOp: &rolebindingpb.StringEqualsOp{
							Claim: fakeClaimName,
							Value: fakeClaimValue,
						},
					},
				},
			},
		},
		Resource: &rolebindingpb.Resource{
			Type: fakeResourceType,
			Id:   fakeResourceID,
		},
	}
}
