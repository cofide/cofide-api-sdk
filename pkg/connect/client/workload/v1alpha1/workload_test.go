// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"
	"io"
	"testing"

	workloadsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/workload_service/v1alpha1"
	workloadpb "github.com/cofide/cofide-api-sdk/gen/go/proto/workload/v1alpha1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

func TestWorkloadClient_implementsMethods(t *testing.T) {
	test.AssertClientImplementsService(t, &workloadClient{}, workloadsvcpb.WorkloadService_ServiceDesc)
}

// TestWorkloadClient_Unimplemented tests WorkloadClient against an unimplemented server.
// This ensures that all errors returned are not wrapped can be converted to a gRPC Status using Status.Convert.
func TestWorkloadClient_Unimplemented(t *testing.T) {
	server := test.NewTestServer(t)
	workloadsvcpb.RegisterWorkloadServiceServer(server.Server, &workloadsvcpb.UnimplementedWorkloadServiceServer{})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	workloads, err := client.ListWorkloads(ctx, nil)
	test.RequireUnimplemented(t, err)
	if workloads != nil {
		t.Errorf("expected nil workloads, got %v", workloads)
	}

	stream, err := client.PublishWorkloads(ctx)
	require.NoError(t, err)
	_ = stream.Send([]*workloadpb.Workload{{Id: "workload-1"}})
	test.RequireUnimplemented(t, stream.Close())

	eventsStream, err := client.PublishWorkloadEvents(ctx)
	require.NoError(t, err)
	_ = eventsStream.Send([]*workloadpb.WorkloadEvent{{OrgId: "org-1"}})
	test.RequireUnimplemented(t, eventsStream.Close())
}

func TestWorkloadClient_ListWorkloads(t *testing.T) {
	fakeService := &fakeWorkloadService{
		workloads: []*workloadpb.Workload{
			{Id: "workload-1", OrgId: "org-1"},
			{Id: "workload-2", OrgId: "org-1"},
		},
	}
	server := test.NewTestServer(t)
	workloadsvcpb.RegisterWorkloadServiceServer(server.Server, fakeService)
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)

	filter := &workloadsvcpb.ListWorkloadsRequest_Filter{OrgId: proto.String("org-1")}
	workloads, err := client.ListWorkloads(t.Context(), filter)
	require.NoError(t, err)

	require.Len(t, workloads, 2)
	assert.Equal(t, "workload-1", workloads[0].GetId())
	assert.Equal(t, "workload-2", workloads[1].GetId())
	assert.Equal(t, "org-1", fakeService.receivedFilter.GetOrgId())
}

func TestWorkloadClient_PublishWorkloads(t *testing.T) {
	fakeService := &fakeWorkloadService{}
	server := test.NewTestServer(t)
	workloadsvcpb.RegisterWorkloadServiceServer(server.Server, fakeService)
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	stream, err := client.PublishWorkloads(ctx)
	require.NoError(t, err)

	workloads := []*workloadpb.Workload{
		{Id: "workload-1", OrgId: "org-1"},
		{Id: "workload-2", OrgId: "org-1"},
	}
	require.NoError(t, stream.Send(workloads))
	require.NoError(t, stream.Send([]*workloadpb.Workload{{Id: "workload-3", OrgId: "org-1"}}))
	require.NoError(t, stream.Close())

	assert.Len(t, fakeService.received, 3)
	assert.Equal(t, "workload-1", fakeService.received[0].GetId())
	assert.Equal(t, "workload-2", fakeService.received[1].GetId())
	assert.Equal(t, "workload-3", fakeService.received[2].GetId())
}

func TestWorkloadClient_PublishWorkloadEvents(t *testing.T) {
	fakeService := &fakeWorkloadService{}
	server := test.NewTestServer(t)
	workloadsvcpb.RegisterWorkloadServiceServer(server.Server, fakeService)
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	stream, err := client.PublishWorkloadEvents(ctx)
	require.NoError(t, err)

	events := []*workloadpb.WorkloadEvent{
		{Id: "event-1", OrgId: "org-1"},
		{Id: "event-2", OrgId: "org-1"},
	}
	require.NoError(t, stream.Send(events))
	require.NoError(t, stream.Send([]*workloadpb.WorkloadEvent{{Id: "event-3", OrgId: "org-1"}}))
	require.NoError(t, stream.Close())

	assert.Len(t, fakeService.receivedEvents, 3)
	assert.Equal(t, "event-1", fakeService.receivedEvents[0].GetId())
	assert.Equal(t, "event-2", fakeService.receivedEvents[1].GetId())
	assert.Equal(t, "event-3", fakeService.receivedEvents[2].GetId())
}

type fakeWorkloadService struct {
	workloadsvcpb.UnimplementedWorkloadServiceServer

	workloads      []*workloadpb.Workload
	receivedFilter *workloadsvcpb.ListWorkloadsRequest_Filter

	received       []*workloadpb.Workload
	receivedEvents []*workloadpb.WorkloadEvent
}

func (f *fakeWorkloadService) ListWorkloads(ctx context.Context, req *workloadsvcpb.ListWorkloadsRequest) (*workloadsvcpb.ListWorkloadsResponse, error) {
	f.receivedFilter = req.GetFilter()
	return &workloadsvcpb.ListWorkloadsResponse{Workloads: f.workloads}, nil
}

func (f *fakeWorkloadService) PublishWorkloads(stream grpc.ClientStreamingServer[workloadsvcpb.PublishWorkloadsRequest, workloadsvcpb.PublishWorkloadsResponse]) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&workloadsvcpb.PublishWorkloadsResponse{})
		}
		if err != nil {
			return err
		}
		f.received = append(f.received, req.GetWorkloads()...)
	}
}

func (f *fakeWorkloadService) PublishWorkloadEvents(stream grpc.ClientStreamingServer[workloadsvcpb.PublishWorkloadEventsRequest, workloadsvcpb.PublishWorkloadEventsResponse]) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&workloadsvcpb.PublishWorkloadEventsResponse{})
		}
		if err != nil {
			return err
		}
		f.receivedEvents = append(f.receivedEvents, req.GetEvents()...)
	}
}
