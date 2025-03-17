// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	agentv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/agent/v1alpha1"
	clusterv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/cluster/v1alpha1"
	trustzonev1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/trustzone/v1alpha1"
	"google.golang.org/grpc"
)

// ClientSet is an interface that provides methods for accessing individual
// versioned clients for specific gRPC services in Cofide Connect API.
type ClientSet interface {
	TrustZoneV1Alpha1() trustzonev1alpha1.TrustZoneClient
	ClusterV1Alpha1() clusterv1alpha1.ClusterClient
	AgentV1Alpha1() agentv1alpha1.AgentClient
}

type clientSet struct {
	trustZoneV1Alpha1 trustzonev1alpha1.TrustZoneClient
	clusterV1Alpha1   clusterv1alpha1.ClusterClient
	agentV1Alpha1     agentv1alpha1.AgentClient
}

// New instantiates a new ClientSet for communication with a Connect API.
func New(conn grpc.ClientConnInterface) ClientSet {
	return &clientSet{
		trustZoneV1Alpha1: trustzonev1alpha1.New(conn),
		clusterV1Alpha1:   clusterv1alpha1.New(conn),
		agentV1Alpha1:     agentv1alpha1.New(conn),
	}
}

func (c *clientSet) TrustZoneV1Alpha1() trustzonev1alpha1.TrustZoneClient {
	return c.trustZoneV1Alpha1
}

func (c *clientSet) ClusterV1Alpha1() clusterv1alpha1.ClusterClient {
	return c.clusterV1Alpha1
}

func (c *clientSet) AgentV1Alpha1() agentv1alpha1.AgentClient {
	return c.agentV1Alpha1
}
