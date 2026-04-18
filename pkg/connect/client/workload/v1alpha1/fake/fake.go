// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	workloadsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/workload_service/v1alpha1"
	workloadpb "github.com/cofide/cofide-api-sdk/gen/go/proto/workload/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
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

func (c *fakeWorkloadClient) ListWorkloadEvents(ctx context.Context, filter *workloadpb.ListWorkloadEventsRequest_Filter) ([]*workloadpb.WorkloadEvent, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	events := []*workloadpb.WorkloadEvent{}
	for _, event := range c.fake.WorkloadEvents {
		if workloadEventMatches(event, filter) {
			events = append(events, proto.Clone(event).(*workloadpb.WorkloadEvent))
		}
	}
	return events, nil
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

func workloadEventMatches(event *workloadpb.WorkloadEvent, filter *workloadpb.ListWorkloadEventsRequest_Filter) bool {
	if filter == nil {
		return true
	}
	if filter.OrgId != "" && event.GetOrgId() != filter.OrgId {
		return false
	}
	if filter.TrustZoneId != "" && event.GetTrustZoneId() != filter.TrustZoneId {
		return false
	}
	if filter.ClusterId != "" && event.GetClusterId() != filter.ClusterId {
		return false
	}
	if filter.SpiffeId != nil {
		if delivered := event.GetIdentityDelivered(); delivered == nil || delivered.SpiffeId != *filter.SpiffeId {
			return false
		}
	}
	if filter.EntryId != nil {
		if delivered := event.GetIdentityDelivered(); delivered == nil || delivered.EntryId != *filter.EntryId {
			return false
		}
	}
	return true
}
