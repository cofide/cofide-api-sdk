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
			workloads = append(workloads, clone(workload))
		}
	}
	return workloads, nil
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

func clone(workload *workloadpb.Workload) *workloadpb.Workload {
	return proto.Clone(workload).(*workloadpb.Workload)
}
