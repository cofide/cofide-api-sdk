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
	AddLog(ctx context.Context, log *auditpb.Log) (*auditpb.Log, error)
	ListLogs(ctx context.Context, filter *auditsvcpb.ListLogsRequest_Filter) ([]*auditpb.Log, error)
}

type auditClient struct {
	auditClient auditsvcpb.AuditServiceClient
}

func New(conn grpc.ClientConnInterface) AuditClient {
	return &auditClient{
		auditClient: auditsvcpb.NewAuditServiceClient(conn),
	}
}

func (c *auditClient) AddLog(ctx context.Context, log *auditpb.Log) (*auditpb.Log, error) {
	resp, err := c.auditClient.AddLog(ctx, &auditsvcpb.AddLogRequest{
		Log: log,
	})
	if err != nil {
		return nil, err
	}

	return resp.Log, nil
}

func (c *auditClient) ListLogs(ctx context.Context, filter *auditsvcpb.ListLogsRequest_Filter) ([]*auditpb.Log, error) {
	resp, err := c.auditClient.ListLogs(ctx, &auditsvcpb.ListLogsRequest{
		Filter: filter,
	})
	if err != nil {
		return nil, err
	}

	return resp.Logs, nil
}
