// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"slices"

	workloadsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/workload_service/v1alpha1"
	workloadpb "github.com/cofide/cofide-api-sdk/gen/go/proto/workload/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
	workloadv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/workload/v1alpha1"
	"google.golang.org/protobuf/proto"
)

type fakeWorkloadClient struct {
	fake *fakeconnect.FakeConnect
}

// New instantiates a new WorkloadClient for communication with a fake Connect API.
func New(fake *fakeconnect.FakeConnect) workloadv1alpha1.WorkloadClient {
	return &fakeWorkloadClient{
		fake: fake,
	}
}

func (c *fakeWorkloadClient) ListWorkloads(ctx context.Context, filter *workloadsvcpb.ListWorkloadsRequest_Filter) ([]*workloadpb.Workload, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	workloads := []*workloadpb.Workload{}
	for _, workload := range c.fake.Workloads {
		if workloadMatches(workload, filter) {
			workloads = append(workloads, proto.Clone(workload).(*workloadpb.Workload))
		}
	}
	return workloads, nil
}

func (c *fakeWorkloadClient) ListWorkloadEvents(ctx context.Context, filter *workloadsvcpb.ListWorkloadEventsRequest_Filter, requestPagination pagination.Pagination) ([]*workloadpb.WorkloadEvent, pagination.Pagination, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	events := []*workloadpb.WorkloadEvent{}
	for _, event := range c.fake.WorkloadEvents {
		if workloadEventMatches(event, filter) {
			events = append(events, proto.Clone(event).(*workloadpb.WorkloadEvent))
		}
	}
	return events, pagination.Pagination{PageSize: requestPagination.PageSize}, nil
}

func (c *fakeWorkloadClient) PublishWorkloadEvents(ctx context.Context) (workloadv1alpha1.WorkloadEventsStream, error) {
	return &fakeWorkloadEventsStream{fake: c.fake}, nil
}

type fakeWorkloadEventsStream struct {
	fake *fakeconnect.FakeConnect
}

func (s *fakeWorkloadEventsStream) Send(events []*workloadpb.WorkloadEvent) error {
	s.fake.Mu.Lock()
	defer s.fake.Mu.Unlock()

	for _, event := range events {
		if event == nil {
			continue
		}
		s.fake.WorkloadEvents = append(s.fake.WorkloadEvents, proto.Clone(event).(*workloadpb.WorkloadEvent))
	}
	return nil
}

func (s *fakeWorkloadEventsStream) Close() error {
	return nil
}

func workloadMatches(workload *workloadpb.Workload, filter *workloadsvcpb.ListWorkloadsRequest_Filter) bool {
	if filter == nil {
		return true
	}
	if filter.OrgId != nil && workload.GetOrgId() != *filter.OrgId {
		return false
	}
	if filter.TrustZoneId != nil && workload.GetTrustZoneId() != *filter.TrustZoneId {
		return false
	}
	if filter.ClusterId != nil && workload.GetClusterId() != *filter.ClusterId {
		return false
	}
	return true
}

func workloadEventMatches(event *workloadpb.WorkloadEvent, filter *workloadsvcpb.ListWorkloadEventsRequest_Filter) bool {
	if filter == nil {
		return true
	}
	if len(filter.GetOrgIds()) > 0 && !slices.Contains(filter.GetOrgIds(), event.GetOrgId()) {
		return false
	}
	if len(filter.GetTrustZoneIds()) > 0 && !slices.Contains(filter.GetTrustZoneIds(), event.GetTrustZoneId()) {
		return false
	}
	if len(filter.GetClusterIds()) > 0 && !slices.Contains(filter.GetClusterIds(), event.GetClusterId()) {
		return false
	}
	if len(filter.GetAgentSpiffeIds()) > 0 && !slices.Contains(filter.GetAgentSpiffeIds(), event.GetAgentSpiffeId()) {
		return false
	}
	if len(filter.GetWorkloadIds()) > 0 && !slices.Contains(filter.GetWorkloadIds(), event.GetWorkloadId()) {
		return false
	}
	if len(filter.GetEventTypes()) > 0 && !workloadEventTypeMatches(event, filter.GetEventTypes()) {
		return false
	}
	observedBefore := filter.GetObservedBefore()
	observedAfter := filter.GetObservedAfter()
	hasObservedBefore := observedBefore.IsValid()
	hasObservedAfter := observedAfter.IsValid()
	if hasObservedBefore || hasObservedAfter {
		observed := event.GetObservedTimestamp()
		if !observed.IsValid() {
			return false
		}
		if hasObservedBefore && observed.AsTime().After(observedBefore.AsTime()) {
			return false
		}
		if hasObservedAfter && observed.AsTime().Before(observedAfter.AsTime()) {
			return false
		}
	}
	if len(filter.GetSpiffeIds()) > 0 {
		if delivered := event.GetIdentityDelivered(); delivered == nil || !slices.Contains(filter.GetSpiffeIds(), delivered.GetSpiffeId()) {
			return false
		}
	}
	if len(filter.GetEntryIds()) > 0 {
		if delivered := event.GetIdentityDelivered(); delivered == nil || !slices.Contains(filter.GetEntryIds(), delivered.GetEntryId()) {
			return false
		}
	}
	return true
}

func workloadEventTypeMatches(event *workloadpb.WorkloadEvent, eventTypes []workloadpb.WorkloadEventType) bool {
	actual := workloadEventType(event)
	for _, eventType := range eventTypes {
		if eventType == actual {
			return true
		}
	}
	return false
}

func workloadEventType(event *workloadpb.WorkloadEvent) workloadpb.WorkloadEventType {
	switch event.GetEvent().(type) {
	case *workloadpb.WorkloadEvent_WorkloadAttested:
		return workloadpb.WorkloadEventType_WORKLOAD_EVENT_TYPE_WORKLOAD_ATTESTED
	case *workloadpb.WorkloadEvent_WorkloadAttestationFailed:
		return workloadpb.WorkloadEventType_WORKLOAD_EVENT_TYPE_WORKLOAD_ATTESTATION_FAILED
	case *workloadpb.WorkloadEvent_IdentityDelivered:
		return workloadpb.WorkloadEventType_WORKLOAD_EVENT_TYPE_IDENTITY_DELIVERED
	case *workloadpb.WorkloadEvent_NoIdentity:
		return workloadpb.WorkloadEventType_WORKLOAD_EVENT_TYPE_NO_IDENTITY
	default:
		return workloadpb.WorkloadEventType_WORKLOAD_EVENT_TYPE_UNSPECIFIED
	}
}
