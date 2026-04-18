// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"

	workloadsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/workload_service/v1alpha1"
	workloadpb "github.com/cofide/cofide-api-sdk/gen/go/proto/workload/v1alpha1"
	"google.golang.org/grpc"
)

// WorkloadEventsStream is the write end of a PublishWorkloadEvents client stream.
// The caller is responsible for batching, reconnection, and calling Close when done.
type WorkloadEventsStream interface {
	Send(events []*workloadpb.WorkloadEvent) error
	Close() error
}

// WorkloadClient is an interface for a client for the v1alpha1 version of the Connect WorkloadService.
type WorkloadClient interface {
	ListWorkloads(ctx context.Context, filter *workloadsvcpb.ListWorkloadsRequest_Filter) ([]*workloadpb.Workload, error)
	ListWorkloadEvents(ctx context.Context, filter *workloadpb.ListWorkloadEventsRequest_Filter) ([]*workloadpb.WorkloadEvent, error)
	PublishWorkloadEvents(ctx context.Context) (WorkloadEventsStream, error)
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

func (c *workloadClient) ListWorkloadEvents(ctx context.Context, filter *workloadpb.ListWorkloadEventsRequest_Filter) ([]*workloadpb.WorkloadEvent, error) {
	resp, err := c.workloadClient.ListWorkloadEvents(ctx, &workloadpb.ListWorkloadEventsRequest{Filter: filter})
	if err != nil {
		return nil, err
	}
	return resp.Events, nil
}

func (c *workloadClient) PublishWorkloadEvents(ctx context.Context) (WorkloadEventsStream, error) {
	stream, err := c.workloadClient.PublishWorkloadEvents(ctx)
	if err != nil {
		return nil, err
	}
	return &workloadEventsStream{stream: stream}, nil
}

type workloadEventsStream struct {
	stream grpc.ClientStreamingClient[workloadpb.PublishWorkloadEventsRequest, workloadpb.PublishWorkloadEventsResponse]
}

func (s *workloadEventsStream) Send(events []*workloadpb.WorkloadEvent) error {
	return s.stream.Send(&workloadpb.PublishWorkloadEventsRequest{Events: events})
}

func (s *workloadEventsStream) Close() error {
	_, err := s.stream.CloseAndRecv()
	return err
}
