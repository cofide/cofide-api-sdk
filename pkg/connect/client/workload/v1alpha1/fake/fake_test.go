// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"testing"
	"time"

	workloadsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/workload_service/v1alpha1"
	workloadpb "github.com/cofide/cofide-api-sdk/gen/go/proto/workload/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_fakeWorkloadClient_ListWorkloads(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	workload := test.FakeK8sPodWorkload()

	workloads, err := client.ListWorkloads(ctx, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*workloadpb.Workload{}, workloads)

	fake.Workloads[test.FakeWorkloadID] = test.FakeK8sPodWorkload()

	workloads, err = client.ListWorkloads(ctx, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*workloadpb.Workload{workload}, workloads)
}

func Test_fakeWorkloadClient_ListWorkloadEvents_filterObservedTimeRange(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()
	now := time.Now()

	recent := &workloadpb.WorkloadEvent{
		OrgId:             "org-1",
		ObservedTimestamp: timestamppb.New(now.Add(-5 * time.Minute)),
	}
	stale := &workloadpb.WorkloadEvent{
		OrgId:             "org-1",
		ObservedTimestamp: timestamppb.New(now.Add(-2 * time.Hour)),
	}
	fake.WorkloadEvents = []*workloadpb.WorkloadEvent{recent, stale}

	events, _, err := client.ListWorkloadEvents(ctx, &workloadsvcpb.ListWorkloadEventsRequest_Filter{
		OrgId:          "org-1",
		ObservedAfter:  timestamppb.New(now.Add(-time.Hour)),
		ObservedBefore: timestamppb.New(now),
	}, pagination.Pagination{PageSize: 100})
	require.NoError(t, err)

	assert.EqualExportedValues(t, []*workloadpb.WorkloadEvent{recent}, events)
}

func Test_fakeWorkloadClient_ListWorkloadEvents_filterAgentSpiffeID(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	matching := &workloadpb.WorkloadEvent{
		AgentSpiffeId: "spiffe://example.org/agent/matching",
	}
	other := &workloadpb.WorkloadEvent{
		AgentSpiffeId: "spiffe://example.org/agent/other",
	}
	fake.WorkloadEvents = []*workloadpb.WorkloadEvent{matching, other}

	events, _, err := client.ListWorkloadEvents(ctx, &workloadsvcpb.ListWorkloadEventsRequest_Filter{
		AgentSpiffeId: "spiffe://example.org/agent/matching",
	}, pagination.Pagination{PageSize: 100})
	require.NoError(t, err)

	assert.EqualExportedValues(t, []*workloadpb.WorkloadEvent{matching}, events)
}

func Test_fakeWorkloadClient_ListWorkloadEvents_filterEventTypes(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	workloadAttested := &workloadpb.WorkloadEvent{
		Event: &workloadpb.WorkloadEvent_WorkloadAttested{
			WorkloadAttested: &workloadpb.WorkloadAttestedEvent{},
		},
	}
	workloadAttestationFailed := &workloadpb.WorkloadEvent{
		Event: &workloadpb.WorkloadEvent_WorkloadAttestationFailed{
			WorkloadAttestationFailed: &workloadpb.WorkloadAttestationFailedEvent{
				Error: "context cancelled",
			},
		},
	}
	identityDelivered := &workloadpb.WorkloadEvent{
		Event: &workloadpb.WorkloadEvent_IdentityDelivered{
			IdentityDelivered: &workloadpb.IdentityDeliveredEvent{
				EntryId:  "entry-1",
				SpiffeId: "spiffe://example.org/workload",
			},
		},
	}
	noIdentity := &workloadpb.WorkloadEvent{
		Event: &workloadpb.WorkloadEvent_NoIdentity{
			NoIdentity: &workloadpb.NoIdentityEvent{Error: "no matching entries"},
		},
	}
	fake.WorkloadEvents = []*workloadpb.WorkloadEvent{workloadAttested, workloadAttestationFailed, identityDelivered, noIdentity}

	events, _, err := client.ListWorkloadEvents(ctx, &workloadsvcpb.ListWorkloadEventsRequest_Filter{
		EventTypes: []workloadpb.WorkloadEventType{
			workloadpb.WorkloadEventType_WORKLOAD_EVENT_TYPE_WORKLOAD_ATTESTATION_FAILED,
			workloadpb.WorkloadEventType_WORKLOAD_EVENT_TYPE_IDENTITY_DELIVERED,
			workloadpb.WorkloadEventType_WORKLOAD_EVENT_TYPE_NO_IDENTITY,
		},
	}, pagination.Pagination{PageSize: 100})
	require.NoError(t, err)

	assert.EqualExportedValues(t, []*workloadpb.WorkloadEvent{workloadAttestationFailed, identityDelivered, noIdentity}, events)
}

func Test_fakeWorkloadEventsStream_Send_skipsNilEvents(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	stream, err := client.PublishWorkloadEvents(ctx)
	require.NoError(t, err)

	event := &workloadpb.WorkloadEvent{OrgId: "org-1"}
	require.NoError(t, stream.Send([]*workloadpb.WorkloadEvent{event, nil}))

	assert.EqualExportedValues(t, []*workloadpb.WorkloadEvent{event}, fake.WorkloadEvents)
}
