// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	auditserverpb "github.com/cofide/cofide-api-sdk/gen/go/proto/audit/v1alpha1"
	auditsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/audit_service/v1alpha1"
	auditv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/audit/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type fakeAuditClient struct {
	fake *fakeconnect.FakeConnect
}

func New(fake *fakeconnect.FakeConnect) auditv1alpha1.AuditClient {
	return &fakeAuditClient{
		fake: fake,
	}
}

func (c *fakeAuditClient) AddLog(ctx context.Context, log *auditserverpb.Log) (*auditserverpb.Log, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	log = clone(log)
	log.Id = uuid.New().String()
	log.OccurredAt = timestamppb.Now()
	c.fake.AuditLogs[log.GetId()] = log
	return clone(log), nil
}

func (c *fakeAuditClient) ListLogs(ctx context.Context, filter *auditsvcpb.ListLogsRequest_Filter) ([]*auditserverpb.Log, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	logs := []*auditserverpb.Log{}
	for _, log := range c.fake.AuditLogs {
		if c.logMatches(log, filter) {
			logs = append(logs, clone(log))
		}
	}
	return logs, nil
}

func (c *fakeAuditClient) logMatches(log *auditserverpb.Log, filter *auditsvcpb.ListLogsRequest_Filter) bool {
	if filter == nil {
		return true
	}
	if filter.EventType != "" && log.GetEventType() != filter.EventType {
		return false
	}
	if filter.EntityType != "" && log.GetEntityType() != filter.EntityType {
		return false
	}
	if filter.EntityId != "" && log.GetEntityId() != filter.EntityId {
		return false
	}
	// TODO support time based filters
	return true
}

func clone(log *auditserverpb.Log) *auditserverpb.Log {
	return proto.Clone(log).(*auditserverpb.Log)
}
