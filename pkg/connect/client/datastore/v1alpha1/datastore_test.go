// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"
	"testing"

	datastorev1alpha1 "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/datastore_service/v1alpha1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	fakeSpiffeID     = "spiffe://example.org/node/fake-node"
	fakeSelectorType = "k8s"
	fakeSelectorVal  = "pod:default/example"
)

func TestDataStoreClient_implementsMethods(t *testing.T) {
	test.AssertClientImplementsService(t, &datastoreClient{}, datastorev1alpha1.DataStoreService_ServiceDesc)
}

func TestDataStoreClient_Unimplemented(t *testing.T) {
	server := test.NewTestServer(t)
	datastorev1alpha1.RegisterDataStoreServiceServer(server.Server, &datastorev1alpha1.UnimplementedDataStoreServiceServer{})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	_, err := client.CountAttestedNodes(ctx, &datastorev1alpha1.CountAttestedNodesRequest{})
	test.RequireUnimplemented(t, err)

	_, err = client.CreateAttestedNode(ctx, &datastorev1alpha1.CreateAttestedNodeRequest{})
	test.RequireUnimplemented(t, err)

	_, err = client.DeleteAttestedNode(ctx, &datastorev1alpha1.DeleteAttestedNodeRequest{})
	test.RequireUnimplemented(t, err)

	_, err = client.FetchAttestedNode(ctx, &datastorev1alpha1.FetchAttestedNodeRequest{})
	test.RequireUnimplemented(t, err)

	_, err = client.ListAttestedNodes(ctx, &datastorev1alpha1.ListAttestedNodesRequest{})
	test.RequireUnimplemented(t, err)

	_, err = client.UpdateAttestedNode(ctx, &datastorev1alpha1.UpdateAttestedNodeRequest{})
	test.RequireUnimplemented(t, err)

	_, err = client.GetNodeSelectors(ctx, &datastorev1alpha1.GetNodeSelectorsRequest{})
	test.RequireUnimplemented(t, err)

	_, err = client.SetNodeSelectors(ctx, &datastorev1alpha1.SetNodeSelectorsRequest{})
	test.RequireUnimplemented(t, err)

	_, err = client.ListNodeSelectors(ctx, &datastorev1alpha1.ListNodeSelectorsRequest{})
	test.RequireUnimplemented(t, err)
}

func TestDataStoreClient(t *testing.T) {
	server := test.NewTestServer(t)
	datastorev1alpha1.RegisterDataStoreServiceServer(server.Server, &fakeDataStoreService{t: t})
	server.Serve()

	conn := server.CreateClientConn()
	client := New(conn)
	ctx := context.Background()

	node := fakeAttestedNode()

	// Test CountAttestedNodes
	countResp, err := client.CountAttestedNodes(ctx, &datastorev1alpha1.CountAttestedNodesRequest{})
	require.NoError(t, err)
	assert.Equal(t, int32(1), countResp.Count)

	// Test CreateAttestedNode
	createResp, err := client.CreateAttestedNode(ctx, &datastorev1alpha1.CreateAttestedNodeRequest{Node: node})
	require.NoError(t, err)
	assert.EqualExportedValues(t, node, createResp.Node)

	// Test DeleteAttestedNode
	deleteResp, err := client.DeleteAttestedNode(ctx, &datastorev1alpha1.DeleteAttestedNodeRequest{SpiffeId: fakeSpiffeID})
	require.NoError(t, err)
	assert.NotNil(t, deleteResp)

	// Test FetchAttestedNode
	fetchResp, err := client.FetchAttestedNode(ctx, &datastorev1alpha1.FetchAttestedNodeRequest{SpiffeId: fakeSpiffeID})
	require.NoError(t, err)
	assert.EqualExportedValues(t, node, fetchResp.Node)

	// Test ListAttestedNodes
	listResp, err := client.ListAttestedNodes(ctx, &datastorev1alpha1.ListAttestedNodesRequest{})
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*datastorev1alpha1.AttestedNode{node}, listResp.Nodes)

	// Test UpdateAttestedNode
	updateResp, err := client.UpdateAttestedNode(ctx, &datastorev1alpha1.UpdateAttestedNodeRequest{Node: node})
	require.NoError(t, err)
	assert.EqualExportedValues(t, node, updateResp.Node)

	// Test GetNodeSelectors
	selectors := fakeNodeSelectors()
	getSelectorsResp, err := client.GetNodeSelectors(ctx, &datastorev1alpha1.GetNodeSelectorsRequest{SpiffeId: fakeSpiffeID})
	require.NoError(t, err)
	assert.EqualExportedValues(t, selectors, getSelectorsResp.Selectors)

	// Test ListNodeSelectors
	listSelectorsResp, err := client.ListNodeSelectors(ctx, &datastorev1alpha1.ListNodeSelectorsRequest{})
	require.NoError(t, err)
	assert.EqualExportedValues(t, map[string]*datastorev1alpha1.ListNodeSelectorsResponse_NodeSelectors{
		fakeSpiffeID: {
			Selectors: selectors,
		},
	}, listSelectorsResp.Selectors)

	// Test SetNodeSelectors
	setSelectorsResp, err := client.SetNodeSelectors(ctx, &datastorev1alpha1.SetNodeSelectorsRequest{SpiffeId: fakeSpiffeID, Selectors: selectors})
	require.NoError(t, err)
	assert.NotNil(t, setSelectorsResp)
}

type fakeDataStoreService struct {
	t *testing.T
	datastorev1alpha1.UnimplementedDataStoreServiceServer
}

func (f *fakeDataStoreService) CountAttestedNodes(ctx context.Context, req *datastorev1alpha1.CountAttestedNodesRequest) (*datastorev1alpha1.CountAttestedNodesResponse, error) {
	return &datastorev1alpha1.CountAttestedNodesResponse{Count: 1}, nil
}

func (f *fakeDataStoreService) CreateAttestedNode(ctx context.Context, req *datastorev1alpha1.CreateAttestedNodeRequest) (*datastorev1alpha1.CreateAttestedNodeResponse, error) {
	assert.Equal(f.t, fakeAttestedNode(), req.Node)
	return &datastorev1alpha1.CreateAttestedNodeResponse{Node: req.Node}, nil
}

func (f *fakeDataStoreService) DeleteAttestedNode(ctx context.Context, req *datastorev1alpha1.DeleteAttestedNodeRequest) (*datastorev1alpha1.DeleteAttestedNodeResponse, error) {
	assert.Equal(f.t, fakeSpiffeID, req.GetSpiffeId())
	return &datastorev1alpha1.DeleteAttestedNodeResponse{}, nil
}

func (f *fakeDataStoreService) FetchAttestedNode(ctx context.Context, req *datastorev1alpha1.FetchAttestedNodeRequest) (*datastorev1alpha1.FetchAttestedNodeResponse, error) {
	assert.Equal(f.t, fakeSpiffeID, req.GetSpiffeId())
	return &datastorev1alpha1.FetchAttestedNodeResponse{Node: fakeAttestedNode()}, nil
}

func (f *fakeDataStoreService) ListAttestedNodes(ctx context.Context, req *datastorev1alpha1.ListAttestedNodesRequest) (*datastorev1alpha1.ListAttestedNodesResponse, error) {
	nodes := []*datastorev1alpha1.AttestedNode{fakeAttestedNode()}
	return &datastorev1alpha1.ListAttestedNodesResponse{Nodes: nodes}, nil
}

func (f *fakeDataStoreService) UpdateAttestedNode(ctx context.Context, req *datastorev1alpha1.UpdateAttestedNodeRequest) (*datastorev1alpha1.UpdateAttestedNodeResponse, error) {
	assert.Equal(f.t, fakeAttestedNode(), req.Node)
	return &datastorev1alpha1.UpdateAttestedNodeResponse{Node: req.Node}, nil
}

func (f *fakeDataStoreService) GetNodeSelectors(ctx context.Context, req *datastorev1alpha1.GetNodeSelectorsRequest) (*datastorev1alpha1.GetNodeSelectorsResponse, error) {
	assert.Equal(f.t, fakeSpiffeID, req.GetSpiffeId())
	return &datastorev1alpha1.GetNodeSelectorsResponse{Selectors: fakeNodeSelectors()}, nil
}

func (f *fakeDataStoreService) SetNodeSelectors(ctx context.Context, req *datastorev1alpha1.SetNodeSelectorsRequest) (*datastorev1alpha1.SetNodeSelectorsResponse, error) {
	assert.Equal(f.t, fakeSpiffeID, req.GetSpiffeId())
	assert.Equal(f.t, fakeNodeSelectors(), req.Selectors)
	return &datastorev1alpha1.SetNodeSelectorsResponse{}, nil
}

func (f *fakeDataStoreService) ListNodeSelectors(ctx context.Context, req *datastorev1alpha1.ListNodeSelectorsRequest) (*datastorev1alpha1.ListNodeSelectorsResponse, error) {
	nodeSelectors := &datastorev1alpha1.ListNodeSelectorsResponse_NodeSelectors{
		Selectors: fakeNodeSelectors(),
	}
	selectors := make(map[string]*datastorev1alpha1.ListNodeSelectorsResponse_NodeSelectors)
	selectors[fakeSpiffeID] = nodeSelectors
	return &datastorev1alpha1.ListNodeSelectorsResponse{Selectors: selectors}, nil
}

func fakeAttestedNode() *datastorev1alpha1.AttestedNode {
	return &datastorev1alpha1.AttestedNode{
		SpiffeId: fakeSpiffeID,
	}
}

func fakeNodeSelectors() []*datastorev1alpha1.Selector {
	return []*datastorev1alpha1.Selector{
		{
			Type:  fakeSelectorType,
			Value: fakeSelectorVal,
		},
	}
}
