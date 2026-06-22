// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"testing"
	"time"

	workloadpb "github.com/cofide/cofide-api-sdk/gen/go/proto/workload/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
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

func Test_fakeWorkloadClient_ListWorkloadEvents_filterMaxAge(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	recent := &workloadpb.WorkloadEvent{
		OrgId:             "org-1",
		ObservedTimestamp: timestamppb.New(time.Now().Add(-5 * time.Minute)),
	}
	stale := &workloadpb.WorkloadEvent{
		OrgId:             "org-1",
		ObservedTimestamp: timestamppb.New(time.Now().Add(-2 * time.Hour)),
	}
	fake.WorkloadEvents = []*workloadpb.WorkloadEvent{recent, stale}

	events, _, err := client.ListWorkloadEvents(ctx, &workloadpb.ListWorkloadEventsRequest_Filter{
		OrgId:  "org-1",
		MaxAge: durationpb.New(time.Hour),
	}, pagination.Pagination{PageSize: 100})
	require.NoError(t, err)

	assert.EqualExportedValues(t, []*workloadpb.WorkloadEvent{recent}, events)
}
