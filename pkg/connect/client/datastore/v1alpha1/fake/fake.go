// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"sync"

	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"

	client "github.com/cofide/cofide-api-sdk/pkg/connect/client/datastore/v1alpha1"

	datastorev1alpha1 "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/datastore_service/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type FakeDataStore struct {
	Mu            sync.Mutex
	AttestedNodes map[string]*datastorev1alpha1.AttestedNode
}

type fakeDataStoreClient struct {
	fake *fakeconnect.FakeConnect
}

// New instantiates a new fake DataStoreClient for testing purposes.
func New(fake *fakeconnect.FakeConnect) client.DataStoreClient {
	return &fakeDataStoreClient{
		fake: fake,
	}
}

// ListAttestedNodes implements v1alpha1.DataStoreClient.
func (c *fakeDataStoreClient) ListAttestedNodes(ctx context.Context, req *datastorev1alpha1.ListAttestedNodesRequest) (*datastorev1alpha1.ListAttestedNodesResponse, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	var nodes []*datastorev1alpha1.AttestedNode
	for _, node := range c.fake.AttestedNodes {
		nodes = append(nodes, proto.Clone(node).(*datastorev1alpha1.AttestedNode))
	}

	return &datastorev1alpha1.ListAttestedNodesResponse{Nodes: nodes}, nil
}

// ListNodeSelectors implements v1alpha1.DataStoreClient.
func (c *fakeDataStoreClient) ListNodeSelectors(ctx context.Context, req *datastorev1alpha1.ListNodeSelectorsRequest) (*datastorev1alpha1.ListNodeSelectorsResponse, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()
	nodes := c.fake.AttestedNodes
	var selectors []*datastorev1alpha1.ListNodeSelectorsResponse_NodeSelectors
	for _, node := range nodes {
		newNode := proto.Clone(node).(*datastorev1alpha1.AttestedNode)
		selectors = append(selectors, &datastorev1alpha1.ListNodeSelectorsResponse_NodeSelectors{
			SpiffeId:  newNode.GetSpiffeId(),
			Selectors: newNode.GetSelectors(),
		})
	}

	return &datastorev1alpha1.ListNodeSelectorsResponse{NodeSelectors: selectors}, nil
}

func (c *fakeDataStoreClient) CountAttestedNodes(ctx context.Context, req *datastorev1alpha1.CountAttestedNodesRequest) (*datastorev1alpha1.CountAttestedNodesResponse, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	count := int32(len(c.fake.AttestedNodes))
	return &datastorev1alpha1.CountAttestedNodesResponse{Count: count}, nil
}

func (c *fakeDataStoreClient) CreateAttestedNode(ctx context.Context, req *datastorev1alpha1.CreateAttestedNodeRequest) (*datastorev1alpha1.CreateAttestedNodeResponse, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()
	node := proto.Clone(req.Node).(*datastorev1alpha1.AttestedNode)
	node.Selectors = nil
	c.fake.AttestedNodes[req.GetNode().GetSpiffeId()] = node
	return &datastorev1alpha1.CreateAttestedNodeResponse{Node: node}, nil
}

func (c *fakeDataStoreClient) DeleteAttestedNode(ctx context.Context, req *datastorev1alpha1.DeleteAttestedNodeRequest) (*datastorev1alpha1.DeleteAttestedNodeResponse, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, exists := c.fake.AttestedNodes[req.GetSpiffeId()]; !exists {
		return nil, status.Error(codes.NotFound, "node not found")
	}

	delete(c.fake.AttestedNodes, req.GetSpiffeId())
	return &datastorev1alpha1.DeleteAttestedNodeResponse{}, nil
}

func (c *fakeDataStoreClient) FetchAttestedNode(ctx context.Context, req *datastorev1alpha1.FetchAttestedNodeRequest) (*datastorev1alpha1.FetchAttestedNodeResponse, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	node, exists := c.fake.AttestedNodes[req.GetSpiffeId()]
	if !exists {
		return nil, status.Error(codes.NotFound, "node not found")
	}
	return &datastorev1alpha1.FetchAttestedNodeResponse{Node: proto.Clone(node).(*datastorev1alpha1.AttestedNode)}, nil
}

func (c *fakeDataStoreClient) UpdateAttestedNode(ctx context.Context, req *datastorev1alpha1.UpdateAttestedNodeRequest) (*datastorev1alpha1.UpdateAttestedNodeResponse, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	_, exists := c.fake.AttestedNodes[req.Node.GetSpiffeId()]
	if !exists {
		return nil, status.Error(codes.NotFound, "node not found")
	}
	newNode := proto.Clone(req.Node).(*datastorev1alpha1.AttestedNode)
	newNode.Selectors = nil
	c.fake.AttestedNodes[req.Node.GetSpiffeId()] = newNode
	return &datastorev1alpha1.UpdateAttestedNodeResponse{Node: proto.Clone(newNode).(*datastorev1alpha1.AttestedNode)}, nil
}

func (c *fakeDataStoreClient) GetNodeSelectors(ctx context.Context, req *datastorev1alpha1.GetNodeSelectorsRequest) (*datastorev1alpha1.GetNodeSelectorsResponse, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	node, exists := c.fake.AttestedNodes[req.GetSpiffeId()]
	if !exists {
		return nil, status.Error(codes.NotFound, "node not found")
	}
	newNode := proto.Clone(node).(*datastorev1alpha1.AttestedNode)
	return &datastorev1alpha1.GetNodeSelectorsResponse{Selectors: newNode.GetSelectors()}, nil
}

func (c *fakeDataStoreClient) SetNodeSelectors(ctx context.Context, req *datastorev1alpha1.SetNodeSelectorsRequest) (*datastorev1alpha1.SetNodeSelectorsResponse, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	node, exists := c.fake.AttestedNodes[req.GetSpiffeId()]
	if !exists {
		return nil, status.Error(codes.NotFound, "node not found")
	}

	node.Selectors = req.GetSelectors()
	newNode := proto.Clone(node).(*datastorev1alpha1.AttestedNode)
	c.fake.AttestedNodes[req.GetSpiffeId()] = newNode
	return &datastorev1alpha1.SetNodeSelectorsResponse{}, nil
}
