// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"

	workloadsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/workload_service/v1alpha1"
	"google.golang.org/grpc"
)

// WorkloadObservationClient is an interface for a client for the v1alpha1 version of the Connect WorkloadObservationService.
type WorkloadObservationClient interface {
	PublishWorkloadEvents(ctx context.Context, req *workloadsvcpb.PublishWorkloadEventsRequest) (*workloadsvcpb.PublishWorkloadEventsResponse, error)
}

type workloadClient struct {
	workloadClient workloadsvcpb.WorkloadObservationServiceClient
}

// New instantiates a new TrustZoneClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) WorkloadObservationClient {
	return &workloadClient{
		workloadClient: workloadsvcpb.NewWorkloadObservationServiceClient(conn),
	}
}

func (c *workloadClient) PublishWorkloadEvents(ctx context.Context, req *workloadsvcpb.PublishWorkloadEventsRequest) (*workloadsvcpb.PublishWorkloadEventsResponse, error) {
	return nil, nil
}
