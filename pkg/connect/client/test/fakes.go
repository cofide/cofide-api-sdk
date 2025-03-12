// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package test

import (
	agentpb "github.com/cofide/cofide-api-sdk/gen/go/proto/agent/v1alpha1"
	clusterpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cluster/v1alpha1"
	federatedservicepb "github.com/cofide/cofide-api-sdk/gen/go/proto/federated_service/v1alpha1"
	trustzonepb "github.com/cofide/cofide-api-sdk/gen/go/proto/trust_zone/v1alpha1"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
)

const (
	FakeTrustZoneID   = "fake-tz-id"
	FakeTrustZoneName = "fake-tz-name"
	FakeTrustDomain   = "fake.trust.domain"

	FakeClusterID   = "fake-cluster-id"
	FakeClusterName = "fake-cluster-name"

	FakeAgentToken = "fake-agent-token"
	FakeAgentID    = "fake-agent-id"

	FakeFSID   = "fake-fs-id"
	FakeFSName = "fake-fs-name"
)

func FakeTrustZone() *trustzonepb.TrustZone {
	return &trustzonepb.TrustZone{
		Id:          PtrOf(FakeTrustZoneID),
		Name:        FakeTrustZoneName,
		TrustDomain: FakeTrustDomain,
	}
}

func FakeCluster() *clusterpb.Cluster {
	return &clusterpb.Cluster{
		Id:          PtrOf(FakeClusterID),
		Name:        PtrOf(FakeClusterName),
		TrustZoneId: PtrOf(FakeTrustZoneID),
	}
}

func FakeAgent() *agentpb.Agent {
	return &agentpb.Agent{
		Id:          PtrOf(FakeAgentID),
		ClusterId:   PtrOf(FakeClusterID),
		TrustZoneId: PtrOf(FakeTrustZoneID),
	}
}

func FakeFederatedService() *federatedservicepb.FederatedService {
	return &federatedservicepb.FederatedService{
		Id:   FakeFSID,
		Name: FakeFSName,
	}
}

func FakeBundle() *types.Bundle {
	return &types.Bundle{TrustDomain: FakeTrustDomain}
}
