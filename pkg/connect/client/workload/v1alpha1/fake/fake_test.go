// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"testing"

	workloadpb "github.com/cofide/cofide-api-sdk/gen/go/proto/workload/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
