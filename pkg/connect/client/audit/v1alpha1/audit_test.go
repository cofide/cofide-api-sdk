// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"
	"testing"

	auditpb "github.com/cofide/cofide-api-sdk/gen/go/proto/audit/v1alpha1"
	auditsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/audit_service/v1alpha1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAuditClient_Unimplemented tests AuditClient against an unimplemented server.
// This ensures that all errors returned are not wrapped can be converted to a gRPC Status using Status.Convert.
func TestAuditClient_Unimplemented(t *testing.T) {
	server := test.NewTestServer(t)
	auditsvcpb.RegisterAuditServiceServer(server.Server, &auditsvcpb.UnimplementedAuditServiceServer{})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)

	events, _, err := client.ListEvents(t.Context(), nil, 100, "")
	test.RequireUnimplemented(t, err)
	assert.Nil(t, events)
}

func TestAuditClient(t *testing.T) {
	server := test.NewTestServer(t)
	auditsvcpb.RegisterAuditServiceServer(server.Server, &fakeAuditService{t: t})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)

	filter := &auditsvcpb.ListEventsRequest_Filter{Entities: []*auditsvcpb.ListEventsRequest_Filter_Entity{{Type: auditpb.EntityType_ENTITY_TYPE_ORGANIZATION}}}
	events, _, err := client.ListEvents(t.Context(), filter, 100, "")
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*auditpb.Event{{Id: "mock-id"}}, events)
}

type fakeAuditService struct {
	t *testing.T
}

func (f *fakeAuditService) ListEvents(ctx context.Context, req *auditsvcpb.ListEventsRequest) (*auditsvcpb.ListEventsResponse, error) {
	assert.Equal(f.t, auditpb.EntityType_ENTITY_TYPE_ORGANIZATION, req.Filter.GetEntities()[0].GetType())
	events := []*auditpb.Event{{Id: "mock-id"}}
	return &auditsvcpb.ListEventsResponse{Events: events}, nil
}
