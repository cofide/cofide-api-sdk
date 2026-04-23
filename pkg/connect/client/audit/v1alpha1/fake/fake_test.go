// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"encoding/json"
	"maps"
	"slices"
	"testing"

	auditpb "github.com/cofide/cofide-api-sdk/gen/go/proto/audit/v1alpha1"
	auditsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/audit_service/v1alpha1"
	trustzonepb "github.com/cofide/cofide-api-sdk/gen/go/proto/trust_zone/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_fakeAuditClient_ListEvents(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)

	events, _, err := client.ListEvents(t.Context(), nil, pagination.Pagination{PageSize: 100})
	require.NoError(t, err)
	assert.Empty(t, events)

	trust_zone_id := "trust-zone-1"

	data := map[string]*trustzonepb.TrustZone{
		"trust_zone": {Id: &trust_zone_id},
	}

	dataBytes, err := json.Marshal(data)
	require.NoError(t, err)

	fake.AuditEvents["mock-id"] = &auditpb.Event{
		Id:         "mock-id",
		OccurredAt: timestamppb.Now(),
		Type:       auditpb.EventType_EVENT_TYPE_TRUST_ZONE_CREATION,
		Message:    "trust zone created",
		Links: []*auditpb.EntityLink{
			{Id: "org-1", Type: auditpb.EntityType_ENTITY_TYPE_ORGANIZATION},
			{Id: trust_zone_id, Type: auditpb.EntityType_ENTITY_TYPE_TRUST_ZONE},
		},
		Data:     dataBytes,
		Actor:    "user-1",
		SourceIp: "1.2.3.4",
		Outcome:  auditpb.Outcome_OUTCOME_SUCCESS,
	}

	events, _, err = client.ListEvents(t.Context(), nil, pagination.Pagination{PageSize: 100})
	require.NoError(t, err)
	assert.EqualExportedValues(t, slices.Collect(maps.Values(fake.AuditEvents)), events)
}

func Test_fakeAuditClient_ListEvents_filterMatchLinkless(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)

	fake.AuditEvents["event-1"] = &auditpb.Event{Id: "event-1", Links: []*auditpb.EntityLink{{Id: "org-1", Type: auditpb.EntityType_ENTITY_TYPE_ORGANIZATION}}}
	fake.AuditEvents["event-2"] = &auditpb.Event{Id: "event-2"}

	// match_linkless alone matches only linkless events
	events, _, err := client.ListEvents(t.Context(), &auditsvcpb.ListEventsRequest_Filter{MatchLinkless: true}, pagination.Pagination{PageSize: 100})
	require.NoError(t, err)
	require.Len(t, events, 1)
	assert.Equal(t, "event-2", events[0].GetId())

	// match_linkless combined with entities matches either
	events, _, err = client.ListEvents(t.Context(), &auditsvcpb.ListEventsRequest_Filter{
		Entities:      []*auditsvcpb.ListEventsRequest_Filter_Entity{{Type: auditpb.EntityType_ENTITY_TYPE_ORGANIZATION}},
		MatchLinkless: true,
	}, pagination.Pagination{PageSize: 100})
	require.NoError(t, err)
	assert.Len(t, events, 2)
}

func Test_fakeAuditClient_ListEvents_filterActors(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)

	fake.AuditEvents["event-1"] = &auditpb.Event{Id: "event-1", Actor: "user-1"}
	fake.AuditEvents["event-2"] = &auditpb.Event{Id: "event-2", Actor: "user-2"}

	events, _, err := client.ListEvents(t.Context(), &auditsvcpb.ListEventsRequest_Filter{Actors: []string{"user-1"}}, pagination.Pagination{PageSize: 100})
	require.NoError(t, err)
	require.Len(t, events, 1)
	assert.Equal(t, "event-1", events[0].GetId())

	events, _, err = client.ListEvents(t.Context(), &auditsvcpb.ListEventsRequest_Filter{Actors: []string{"user-1", "user-2"}}, pagination.Pagination{PageSize: 100})
	require.NoError(t, err)
	assert.Len(t, events, 2)
}

func Test_fakeAuditClient_ListEvents_filterSourceIPs(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)

	fake.AuditEvents["event-1"] = &auditpb.Event{Id: "event-1", SourceIp: "1.2.3.4"}
	fake.AuditEvents["event-2"] = &auditpb.Event{Id: "event-2", SourceIp: "5.6.7.8"}

	events, _, err := client.ListEvents(t.Context(), &auditsvcpb.ListEventsRequest_Filter{SourceIps: []string{"1.2.3.4"}}, pagination.Pagination{PageSize: 100})
	require.NoError(t, err)
	require.Len(t, events, 1)
	assert.Equal(t, "event-1", events[0].GetId())

	events, _, err = client.ListEvents(t.Context(), &auditsvcpb.ListEventsRequest_Filter{SourceIps: []string{"1.2.3.4", "5.6.7.8"}}, pagination.Pagination{PageSize: 100})
	require.NoError(t, err)
	assert.Len(t, events, 2)
}

func Test_fakeAuditClient_ListEvents_filterOutcomes(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)

	fake.AuditEvents["event-1"] = &auditpb.Event{Id: "event-1", Outcome: auditpb.Outcome_OUTCOME_SUCCESS}
	fake.AuditEvents["event-2"] = &auditpb.Event{Id: "event-2", Outcome: auditpb.Outcome_OUTCOME_DENIED}

	events, _, err := client.ListEvents(t.Context(), &auditsvcpb.ListEventsRequest_Filter{Outcomes: []auditpb.Outcome{auditpb.Outcome_OUTCOME_SUCCESS}}, pagination.Pagination{PageSize: 100})
	require.NoError(t, err)
	require.Len(t, events, 1)
	assert.Equal(t, "event-1", events[0].GetId())

	events, _, err = client.ListEvents(t.Context(), &auditsvcpb.ListEventsRequest_Filter{Outcomes: []auditpb.Outcome{auditpb.Outcome_OUTCOME_SUCCESS, auditpb.Outcome_OUTCOME_DENIED}}, pagination.Pagination{PageSize: 100})
	require.NoError(t, err)
	assert.Len(t, events, 2)
}

func Test_fakeAuditClient_ListEvents_filterExcludeEventTypes(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)

	fake.AuditEvents["event-1"] = &auditpb.Event{Id: "event-1", Type: auditpb.EventType_EVENT_TYPE_TRUST_ZONE_CREATION}
	fake.AuditEvents["event-2"] = &auditpb.Event{Id: "event-2", Type: auditpb.EventType_EVENT_TYPE_WORKLOAD_OBSERVATION}

	events, _, err := client.ListEvents(t.Context(), &auditsvcpb.ListEventsRequest_Filter{
		Exclude: &auditsvcpb.ListEventsRequest_Filter_Exclude{
			EventTypes: []auditpb.EventType{auditpb.EventType_EVENT_TYPE_WORKLOAD_OBSERVATION},
		},
	}, pagination.Pagination{PageSize: 100})
	require.NoError(t, err)
	require.Len(t, events, 1)
	assert.Equal(t, "event-1", events[0].GetId())
}

func Test_fakeAuditClient_ListEvents_filterExcludeEntities(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)

	fake.AuditEvents["event-1"] = &auditpb.Event{Id: "event-1", Links: []*auditpb.EntityLink{{Id: "tz-1", Type: auditpb.EntityType_ENTITY_TYPE_TRUST_ZONE}}}
	fake.AuditEvents["event-2"] = &auditpb.Event{Id: "event-2", Links: []*auditpb.EntityLink{{Id: "tz-2", Type: auditpb.EntityType_ENTITY_TYPE_TRUST_ZONE}}}

	// exclude by id
	events, _, err := client.ListEvents(t.Context(), &auditsvcpb.ListEventsRequest_Filter{
		Exclude: &auditsvcpb.ListEventsRequest_Filter_Exclude{
			Entities: []*auditsvcpb.ListEventsRequest_Filter_Entity{{Id: "tz-1"}},
		},
	}, pagination.Pagination{PageSize: 100})
	require.NoError(t, err)
	require.Len(t, events, 1)
	assert.Equal(t, "event-2", events[0].GetId())

	// exclude by type excludes all matching that type
	events, _, err = client.ListEvents(t.Context(), &auditsvcpb.ListEventsRequest_Filter{
		Exclude: &auditsvcpb.ListEventsRequest_Filter_Exclude{
			Entities: []*auditsvcpb.ListEventsRequest_Filter_Entity{{Type: auditpb.EntityType_ENTITY_TYPE_TRUST_ZONE}},
		},
	}, pagination.Pagination{PageSize: 100})
	require.NoError(t, err)
	assert.Empty(t, events)
}

func Test_fakeAuditClient_ListEvents_filterExcludeActors(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)

	fake.AuditEvents["event-1"] = &auditpb.Event{Id: "event-1", Actor: "user-1"}
	fake.AuditEvents["event-2"] = &auditpb.Event{Id: "event-2", Actor: "user-2"}

	events, _, err := client.ListEvents(t.Context(), &auditsvcpb.ListEventsRequest_Filter{
		Exclude: &auditsvcpb.ListEventsRequest_Filter_Exclude{Actors: []string{"user-1"}},
	}, pagination.Pagination{PageSize: 100})
	require.NoError(t, err)
	require.Len(t, events, 1)
	assert.Equal(t, "event-2", events[0].GetId())
}

func Test_fakeAuditClient_ListEvents_filterExcludeSourceIPs(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)

	fake.AuditEvents["event-1"] = &auditpb.Event{Id: "event-1", SourceIp: "1.2.3.4"}
	fake.AuditEvents["event-2"] = &auditpb.Event{Id: "event-2", SourceIp: "5.6.7.8"}

	events, _, err := client.ListEvents(t.Context(), &auditsvcpb.ListEventsRequest_Filter{
		Exclude: &auditsvcpb.ListEventsRequest_Filter_Exclude{SourceIps: []string{"1.2.3.4"}},
	}, pagination.Pagination{PageSize: 100})
	require.NoError(t, err)
	require.Len(t, events, 1)
	assert.Equal(t, "event-2", events[0].GetId())
}

func Test_fakeAuditClient_ListEvents_filterExcludeOutcomes(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)

	fake.AuditEvents["event-1"] = &auditpb.Event{Id: "event-1", Outcome: auditpb.Outcome_OUTCOME_SUCCESS}
	fake.AuditEvents["event-2"] = &auditpb.Event{Id: "event-2", Outcome: auditpb.Outcome_OUTCOME_DENIED}

	events, _, err := client.ListEvents(t.Context(), &auditsvcpb.ListEventsRequest_Filter{
		Exclude: &auditsvcpb.ListEventsRequest_Filter_Exclude{Outcomes: []auditpb.Outcome{auditpb.Outcome_OUTCOME_DENIED}},
	}, pagination.Pagination{PageSize: 100})
	require.NoError(t, err)
	require.Len(t, events, 1)
	assert.Equal(t, "event-1", events[0].GetId())
}

func Test_fakeAuditClient_ListEvents_filterIncludeAndExclude(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)

	fake.AuditEvents["event-1"] = &auditpb.Event{
		Id:    "event-1",
		Type:  auditpb.EventType_EVENT_TYPE_TRUST_ZONE_CREATION,
		Links: []*auditpb.EntityLink{{Id: "tz-1", Type: auditpb.EntityType_ENTITY_TYPE_TRUST_ZONE}},
	}
	fake.AuditEvents["event-2"] = &auditpb.Event{
		Id:    "event-2",
		Type:  auditpb.EventType_EVENT_TYPE_WORKLOAD_OBSERVATION,
		Links: []*auditpb.EntityLink{{Id: "tz-1", Type: auditpb.EntityType_ENTITY_TYPE_TRUST_ZONE}},
	}
	fake.AuditEvents["event-3"] = &auditpb.Event{
		Id:    "event-3",
		Type:  auditpb.EventType_EVENT_TYPE_TRUST_ZONE_CREATION,
		Links: []*auditpb.EntityLink{{Id: "tz-2", Type: auditpb.EntityType_ENTITY_TYPE_TRUST_ZONE}},
	}

	// all events for tz-1, excluding workload observations
	events, _, err := client.ListEvents(t.Context(), &auditsvcpb.ListEventsRequest_Filter{
		Entities: []*auditsvcpb.ListEventsRequest_Filter_Entity{{Id: "tz-1"}},
		Exclude: &auditsvcpb.ListEventsRequest_Filter_Exclude{
			EventTypes: []auditpb.EventType{auditpb.EventType_EVENT_TYPE_WORKLOAD_OBSERVATION},
		},
	}, pagination.Pagination{PageSize: 100})
	require.NoError(t, err)
	require.Len(t, events, 1)
	assert.Equal(t, "event-1", events[0].GetId())
}
