// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"

	workloadsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/workload_service/v1alpha1"
	workloadpb "github.com/cofide/cofide-api-sdk/gen/go/proto/workload/v1alpha1"
	"google.golang.org/grpc"
)

// WorkloadClient is an interface for a client for the v1alpha1 version of the Connect WorkloadService.
type WorkloadClient interface {
	ListWorkloads(ctx context.Context, filter *workloadsvcpb.ListWorkloadsRequest_Filter) ([]*workloadpb.Workload, error)
}

type workloadClient struct {
	workloadClient workloadsvcpb.WorkloadServiceClient
}

// New instantiates a new WorkloadClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) WorkloadClient {
	return &workloadClient{
		workloadClient: workloadsvcpb.NewWorkloadServiceClient(conn),
	}
}

func (c *workloadClient) ListWorkloads(ctx context.Context, filter *workloadsvcpb.ListWorkloadsRequest_Filter) ([]*workloadpb.Workload, error) {
	resp, err := c.workloadClient.ListWorkloads(ctx, &workloadsvcpb.ListWorkloadsRequest{Filter: filter})
	if err != nil {
		return nil, err
	}
	return resp.Workloads, nil
}
