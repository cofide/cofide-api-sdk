// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"

	auditpb "github.com/cofide/cofide-api-sdk/gen/go/proto/audit/v1alpha1"
	auditsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/audit_service/v1alpha1"
	"google.golang.org/grpc"
)

type AuditClient interface {
	RecordEvent(ctx context.Context, log *auditpb.Event) (*auditpb.Event, error)
	ListEvents(ctx context.Context, filter *auditsvcpb.ListEventsRequest_Filter) ([]*auditpb.Event, error)
}

type auditClient struct {
	auditClient auditsvcpb.AuditServiceClient
}

func New(conn grpc.ClientConnInterface) AuditClient {
	return &auditClient{
		auditClient: auditsvcpb.NewAuditServiceClient(conn),
	}
}

func (c *auditClient) RecordEvent(ctx context.Context, event *auditpb.Event) (*auditpb.Event, error) {
	resp, err := c.auditClient.RecordEvent(ctx, &auditsvcpb.RecordEventRequest{
		Event: event,
	})
	if err != nil {
		return nil, err
	}

	return resp.Event, nil
}

func (c *auditClient) ListEvents(ctx context.Context, filter *auditsvcpb.ListEventsRequest_Filter) ([]*auditpb.Event, error) {
	resp, err := c.auditClient.ListEvents(ctx, &auditsvcpb.ListEventsRequest{
		Filter: filter,
	})
	if err != nil {
		return nil, err
	}

	return resp.Events, nil
}
