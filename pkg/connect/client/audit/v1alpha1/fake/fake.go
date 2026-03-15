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

func (c *fakeAuditClient) RecordEvent(ctx context.Context, event *auditserverpb.Event) (*auditserverpb.Event, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	event = clone(event)
	event.Id = uuid.New().String()
	event.OccurredAt = timestamppb.Now()
	c.fake.AuditEvents[event.GetId()] = event
	return clone(event), nil
}

func (c *fakeAuditClient) ListEvents(ctx context.Context, filter *auditsvcpb.ListEventsRequest_Filter) ([]*auditserverpb.Event, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	events := []*auditserverpb.Event{}
	for _, event := range c.fake.AuditEvents {
		if c.eventMatches(event, filter) {
			events = append(events, clone(event))
		}
	}
	return events, nil
}

func (c *fakeAuditClient) eventMatches(event *auditserverpb.Event, filter *auditsvcpb.ListEventsRequest_Filter) bool {
	if filter == nil {
		return true
	}
	if filter.EventType != "" && event.GetType() != filter.EventType {
		return false
	}
	if filter.EntityType != "" {
		matched := false
		for _, link := range event.GetLinks() {
			if link.GetType() == filter.EntityType {
				matched = true
				break
			}
		}
		if !matched {
			return false
		}
	}
	if filter.EntityId != "" {
		matched := false
		for _, link := range event.GetLinks() {
			if link.GetId() == filter.EntityId {
				matched = true
				break
			}
		}
		if !matched {
			return false
		}
	}
	if filter.GetOccurredBefore() != nil {
		if event.GetOccurredAt().AsTime().After(filter.GetOccurredBefore().AsTime()) {
			return false
		}
	}
	if filter.GetOccurredAfter() != nil {
		if event.GetOccurredAt().AsTime().Before(filter.GetOccurredBefore().AsTime()) {
			return false
		}
	}
	return true
}

func clone(event *auditserverpb.Event) *auditserverpb.Event {
	return proto.Clone(event).(*auditserverpb.Event)
}
