// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"slices"

	auditpb "github.com/cofide/cofide-api-sdk/gen/go/proto/audit/v1alpha1"
	auditsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/audit_service/v1alpha1"
	auditv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/audit/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"google.golang.org/protobuf/proto"
)

type fakeAuditClient struct {
	fake *fakeconnect.FakeConnect
}

// New instantiates a new TrustZoneServerClient for communication with a fake Connect API.
func New(fake *fakeconnect.FakeConnect) auditv1alpha1.AuditClient {
	return &fakeAuditClient{
		fake: fake,
	}
}

func (c *fakeAuditClient) ListEvents(ctx context.Context, filter *auditsvcpb.ListEventsRequest_Filter, pageSize int32, pageToken string) ([]*auditpb.Event, string, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	events := []*auditpb.Event{}
	for _, event := range c.fake.AuditEvents {
		if c.eventMatches(event, filter) {
			events = append(events, clone(event))
		}
	}
	return events, "", nil
}

func (c *fakeAuditClient) eventMatches(event *auditpb.Event, filter *auditsvcpb.ListEventsRequest_Filter) bool {
	if filter == nil {
		return true
	}
	if len(filter.GetEventTypes()) != 0 {
		if !slices.Contains(filter.GetEventTypes(), event.GetType()) {
			return false
		}
	}
	if len(filter.GetEntities()) != 0 || filter.GetMatchLinkless() {
		var matched bool
		if filter.GetMatchLinkless() && len(event.GetLinks()) == 0 {
			matched = true
		}
		for _, entity := range filter.GetEntities() {
			if entityLinked(event.GetLinks(), entity) {
				matched = true
				break
			}
		}
		if !matched {
			return false
		}
	}
	if filter.GetOccurredBefore().IsValid() {
		if event.GetOccurredAt().AsTime().After(filter.GetOccurredBefore().AsTime()) {
			return false
		}
	}
	if filter.GetOccurredAfter().IsValid() {
		if event.GetOccurredAt().AsTime().Before(filter.GetOccurredAfter().AsTime()) {
			return false
		}
	}
	if len(filter.GetActors()) != 0 {
		if !slices.Contains(filter.GetActors(), event.GetActor()) {
			return false
		}
	}
	if len(filter.GetSourceIps()) != 0 {
		if !slices.Contains(filter.GetSourceIps(), event.GetSourceIp()) {
			return false
		}
	}
	if len(filter.GetOutcomes()) != 0 {
		if !slices.Contains(filter.GetOutcomes(), event.GetOutcome()) {
			return false
		}
	}
	if exclude := filter.GetExclude(); exclude != nil {
		if slices.Contains(exclude.GetEventTypes(), event.GetType()) {
			return false
		}
		for _, entity := range exclude.GetEntities() {
			if entityLinked(event.GetLinks(), entity) {
				return false
			}
		}
		if slices.Contains(exclude.GetActors(), event.GetActor()) {
			return false
		}
		if slices.Contains(exclude.GetSourceIps(), event.GetSourceIp()) {
			return false
		}
		if slices.Contains(exclude.GetOutcomes(), event.GetOutcome()) {
			return false
		}
	}
	return true
}

// entityLinked reports whether any of the event's links match entity.
func entityLinked(links []*auditpb.EntityLink, entity *auditsvcpb.ListEventsRequest_Filter_Entity) bool {
	for _, link := range links {
		idMatch := entity.GetId() == "" || link.GetId() == entity.GetId()
		typeMatch := entity.GetType() == auditpb.EntityType_ENTITY_TYPE_UNSPECIFIED || link.GetType() == entity.GetType()
		if idMatch && typeMatch {
			return true
		}
	}
	return false
}

func clone(event *auditpb.Event) *auditpb.Event {
	return proto.Clone(event).(*auditpb.Event)
}
