// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"

	auditpb "github.com/cofide/cofide-api-sdk/gen/go/proto/audit/v1alpha1"
	auditsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/audit_service/v1alpha1"
	paginationpb "github.com/cofide/cofide-api-sdk/gen/go/proto/pagination/v1alpha1"
	"google.golang.org/grpc"
)

// AuditClient is an interface for a gRPC client for the v1alpha1 version of the Connect AuditService.
type AuditClient interface {
	ListEvents(ctx context.Context, filter *auditsvcpb.ListEventsRequest_Filter, pageSize int32, pageToken string) ([]*auditpb.Event, string, error)
}

type auditClient struct {
	auditClient auditsvcpb.AuditServiceClient
}

// New instantiates a new AuditClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) AuditClient {
	return &auditClient{
		auditClient: auditsvcpb.NewAuditServiceClient(conn),
	}
}

func (c *auditClient) ListEvents(ctx context.Context, filter *auditsvcpb.ListEventsRequest_Filter, pageSize int32, pageToken string) ([]*auditpb.Event, string, error) {
	resp, err := c.auditClient.ListEvents(ctx, &auditsvcpb.ListEventsRequest{
		Filter: filter,
		Pagination: &paginationpb.PageRequest{
			PageSize:  pageSize,
			PageToken: pageToken,
		},
	})
	if err != nil {
		return nil, "", err
	}

	return resp.Events, resp.Pagination.GetNextPageToken(), nil
}
