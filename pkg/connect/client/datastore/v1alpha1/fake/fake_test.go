// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"strconv"
	"testing"

	client "github.com/cofide/cofide-api-sdk/pkg/connect/client/datastore/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"

	datastorev1alpha1 "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/datastore_service/v1alpha1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTest() (*FakeDataStore, client.DataStoreClient) {
	fake := &FakeDataStore{
		AttestedNodes: make(map[string]*datastorev1alpha1.AttestedNode),
		NodeSelectors: make(map[string][]*datastorev1alpha1.Selector),
	}
	connect := fakeconnect.New()
	client := New(connect)
	return fake, client
}

func TestCreateAndFetchAttestedNode(t *testing.T) {
	_, client := setupTest()

	node := &datastorev1alpha1.AttestedNode{
		CertSerialNumber: "1234",
		CertNotAfter:     1000,
		SpiffeId:         "spiffe://example.org/node",
	}

	resp, err := client.CreateAttestedNode(context.Background(), &datastorev1alpha1.CreateAttestedNodeRequest{
		Node: node,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotEmpty(t, resp.Node.SpiffeId)

	fetchResp, err := client.FetchAttestedNode(context.Background(), &datastorev1alpha1.FetchAttestedNodeRequest{
		SpiffeId: resp.Node.SpiffeId,
	})
	require.NoError(t, err)
	assert.Equal(t, resp.Node, fetchResp.Node)
}

func TestUpdateAttestedNode(t *testing.T) {
	_, client := setupTest()

	node := &datastorev1alpha1.AttestedNode{
		CertSerialNumber: "5678",
		CertNotAfter:     2000,
	}

	resp, err := client.CreateAttestedNode(context.Background(), &datastorev1alpha1.CreateAttestedNodeRequest{
		Node: node,
	})
	require.NoError(t, err)

	updatedNode := resp.Node
	updatedNode.CertSerialNumber = "9101"
	updatedNode.CertNotAfter = 3000

	updateResp, err := client.UpdateAttestedNode(context.Background(), &datastorev1alpha1.UpdateAttestedNodeRequest{
		Node: updatedNode,
	})
	require.NoError(t, err)
	assert.Equal(t, updatedNode, updateResp.Node)

	fetchResp, err := client.FetchAttestedNode(context.Background(), &datastorev1alpha1.FetchAttestedNodeRequest{
		SpiffeId: resp.Node.SpiffeId,
	})
	require.NoError(t, err)
	assert.Equal(t, updatedNode, fetchResp.Node)
}

func TestDeleteAttestedNode(t *testing.T) {
	_, client := setupTest()

	node := &datastorev1alpha1.AttestedNode{
		CertSerialNumber: "1122",
		CertNotAfter:     4000,
	}

	resp, err := client.CreateAttestedNode(context.Background(), &datastorev1alpha1.CreateAttestedNodeRequest{
		Node: node,
	})
	require.NoError(t, err)

	deleteResp, err := client.DeleteAttestedNode(context.Background(), &datastorev1alpha1.DeleteAttestedNodeRequest{
		SpiffeId: resp.Node.SpiffeId,
	})
	require.NoError(t, err)
	require.NotNil(t, deleteResp)

	_, err = client.FetchAttestedNode(context.Background(), &datastorev1alpha1.FetchAttestedNodeRequest{
		SpiffeId: resp.Node.SpiffeId,
	})
	require.Error(t, err)
}

func TestCountAttestedNodes(t *testing.T) {
	fake, client := setupTest()

	// Clear existing nodes
	fake.AttestedNodes = make(map[string]*datastorev1alpha1.AttestedNode)

	for i := 0; i < 5; i++ {
		node := &datastorev1alpha1.AttestedNode{
			CertSerialNumber: "count-test",
			CertNotAfter:     int64(i * 1000),
			SpiffeId:         "spiffe://example.org/node" + strconv.Itoa(i),
		}
		_, err := client.CreateAttestedNode(context.Background(), &datastorev1alpha1.CreateAttestedNodeRequest{
			Node: node,
		})
		require.NoError(t, err)
	}

	countResp, err := client.CountAttestedNodes(context.Background(), &datastorev1alpha1.CountAttestedNodesRequest{})
	require.NoError(t, err)
	assert.Equal(t, int32(5), countResp.Count)
}

func TestSetAndGetNodeSelectors(t *testing.T) {
	_, client := setupTest()

	spiffeID := "spiffe://example.org/node"
	selectors := []*datastorev1alpha1.Selector{
		{Type: "type1", Value: "value1"},
		{Type: "type2", Value: "value2"},
	}

	node := &datastorev1alpha1.AttestedNode{
		SpiffeId: spiffeID,
	}
	_, err := client.CreateAttestedNode(context.Background(), &datastorev1alpha1.CreateAttestedNodeRequest{
		Node: node,
	})
	require.NoError(t, err)
	_, err = client.SetNodeSelectors(context.Background(), &datastorev1alpha1.SetNodeSelectorsRequest{
		SpiffeId:  spiffeID,
		Selectors: selectors,
	})
	require.NoError(t, err)

	getResp, err := client.GetNodeSelectors(context.Background(), &datastorev1alpha1.GetNodeSelectorsRequest{
		SpiffeId: spiffeID,
	})
	require.NoError(t, err)
	require.Len(t, getResp.Selectors, 2)
	assert.Equal(t, selectors, getResp.Selectors)
}
