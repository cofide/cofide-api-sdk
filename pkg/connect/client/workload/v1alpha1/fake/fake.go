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
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	for workload := range fakeconnect.SortedValues(c.fake.Workloads) {
		if workloadMatches(workload, filter) {
			workloads = append(workloads, proto.Clone(workload).(*workloadpb.Workload))
		}
	}
	return workloads, nil
}

func (c *fakeWorkloadClient) PublishWorkloads(ctx context.Context) (workloadv1alpha1.WorkloadsStream, error) {
	return &fakeWorkloadsStream{fake: c.fake}, nil
}

type fakeWorkloadsStream struct {
	fake *fakeconnect.FakeConnect
}

// Send stores the workloads, as the real Connect API does for workloads published by an observer.
//
// The identity of a workload is the combination of its cluster ID, type and platform ID (see platformID). The fake
// generates the workload ID: if a workload with the same identity is already stored it is replaced and keeps its ID,
// otherwise the workload is added with a new ID. A workload marked as deleted removes the stored workload with the
// same identity, if there is one, so it is no longer returned by ListWorkloads.
//
// A workload that Connect would never accept is dropped without failing the stream, and the publisher is not told.
// Such a workload is recorded in the fake's RejectedWorkloads. These are workloads that are not marked as deleted but
// have an ID, and workloads with an unknown type or without a platform ID.
func (s *fakeWorkloadsStream) Send(workloads []*workloadpb.Workload) error {
	s.fake.Mu.Lock()
	defer s.fake.Mu.Unlock()

	for _, workload := range workloads {
		if workload == nil {
			continue
		}
		if err := s.publishWorkload(workload); err != nil {
			s.fake.RejectedWorkloads = append(s.fake.RejectedWorkloads, proto.Clone(workload).(*workloadpb.Workload))
		}
	}
	return nil
}

// publishWorkload stores or deletes a workload, or returns an error if Connect would reject it. The caller must hold
// the lock.
func (s *fakeWorkloadsStream) publishWorkload(workload *workloadpb.Workload) error {
	// Connect generates the ID. The ID is ignored when deleting, which finds the workload by its identity.
	if !workload.GetDeleted() && workload.GetId() != "" {
		return status.Error(codes.InvalidArgument, "workload ID must be empty")
	}
	id, err := platformID(workload)
	if err != nil {
		return err
	}
	existingID, found := s.findObserved(workload, id)

	if workload.GetDeleted() {
		if found {
			delete(s.fake.Workloads, existingID)
		}
		return nil
	}

	stored := proto.Clone(workload).(*workloadpb.Workload)
	if found {
		stored.Id = existingID
	} else {
		stored.Id = uuid.New().String()
	}
	s.fake.Workloads[stored.Id] = stored
	return nil
}

// findObserved returns the ID of the stored workload with the same cluster ID, type and platform ID as workload. The
// caller must hold the lock.
func (s *fakeWorkloadsStream) findObserved(workload *workloadpb.Workload, wantPlatformID string) (string, bool) {
	for id, existing := range s.fake.Workloads {
		if existing.GetClusterId() != workload.GetClusterId() || existing.GetType() != workload.GetType() {
			continue
		}
		// Workloads stored with an unsupported platform identity cannot match.
		if existingPlatformID, err := platformID(existing); err == nil && existingPlatformID == wantPlatformID {
			return id, true
		}
	}
	return "", false
}

// platformID returns the identifier of a workload within its platform, which together with its cluster ID and type
// identifies it to Connect. For Kubernetes pods this is <namespace>/<name>, which allows referencing a pod after it
// has been deleted. For AWS resources it is the ARN.
func platformID(workload *workloadpb.Workload) (string, error) {
	var id string
	switch wl := workload.GetWorkload().(type) {
	case *workloadpb.Workload_KubernetesPod:
		namespace := wl.KubernetesPod.GetMetadata().GetNamespace()
		name := wl.KubernetesPod.GetMetadata().GetName()
		if namespace == "" || name == "" {
			return "", status.Error(codes.InvalidArgument, "kubernetes pod namespace and name must be set")
		}
		return namespace + "/" + name, nil
	case *workloadpb.Workload_AwsLambdaFunction:
		id = wl.AwsLambdaFunction.GetFunctionArn()
	case *workloadpb.Workload_AwsAgentcoreRuntime:
		id = wl.AwsAgentcoreRuntime.GetAgentRuntimeArn()
	default:
		return "", status.Errorf(codes.InvalidArgument, "unknown workload field: %T", wl)
	}
	if id == "" {
		return "", status.Error(codes.InvalidArgument, "workload ARN must be set")
	}
	return id, nil
}

func (s *fakeWorkloadsStream) Close() error {
	return nil
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
