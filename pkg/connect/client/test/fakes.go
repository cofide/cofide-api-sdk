// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package test

import (
	agentpb "github.com/cofide/cofide-api-sdk/gen/go/proto/agent/v1alpha1"
	apbindingpb "github.com/cofide/cofide-api-sdk/gen/go/proto/ap_binding/v1alpha1"
	attestationpolicypb "github.com/cofide/cofide-api-sdk/gen/go/proto/attestation_policy/v1alpha1"
	clusterpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cluster/v1alpha1"
	federatedservicepb "github.com/cofide/cofide-api-sdk/gen/go/proto/federated_service/v1alpha1"
	identitypb "github.com/cofide/cofide-api-sdk/gen/go/proto/identity/v1alpha1"
	trustzonepb "github.com/cofide/cofide-api-sdk/gen/go/proto/trust_zone/v1alpha1"
	workloadpb "github.com/cofide/cofide-api-sdk/gen/go/proto/workload/v1alpha1"
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

	FakeAttestationPolicyID   = "fake-ap-id"
	FakeAttestationPolicyName = "fake-ap-name"

	FakeAPBindingID = "fake-ap-binding-id"

	FakeWorkloadID      = "fake-workload-id"
	FakeK8sPodUID       = "fake-k8s-pod-uid"
	FakeK8sPodName      = "fake-k8s-pod-name"
	FakeK8sPodNamespace = "fake-k8s-pod-namespace"

	FakeIdentityID    = "fake-identity-id"
	FakeSPIFFEID      = "spiffe://fake.trust.domain/ns/fake-k8s-pod-namespace/sa/fake-k8s-pod-service-account"
	FakeParentID      = "spiffe://fake.trust.domain/spire/agent/k8s_psat/fake-cluster-name/fake-spire-agent"
	FakeSelectorType  = "fake-selector-type"
	FakeSelectorValue = "fake-selector-value"
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

func FakeAttestationPolicy() *attestationpolicypb.AttestationPolicy {
	return &attestationpolicypb.AttestationPolicy{
		Id:   PtrOf(FakeAttestationPolicyID),
		Name: FakeAttestationPolicyName,
	}
}

func FakeAPBinding() *apbindingpb.APBinding {
	return &apbindingpb.APBinding{
		Id:          PtrOf(FakeAPBindingID),
		TrustZoneId: PtrOf(FakeTrustZoneID),
		PolicyId:    PtrOf(FakeAttestationPolicyID),
	}
}

func FakeK8sPodWorkload() *workloadpb.Workload {
	return &workloadpb.Workload{
		Id:          FakeWorkloadID,
		TrustZoneId: FakeTrustZoneID,
		ClusterId:   FakeClusterID,
		Type:        workloadpb.WorkloadType_WORKLOAD_TYPE_KUBERNETES_POD,
		Workload: &workloadpb.Workload_KubernetesPod{
			KubernetesPod: &workloadpb.KubernetesPod{
				Metadata: &workloadpb.KubernetesMetadata{
					Uid:       FakeK8sPodUID,
					Name:      FakeK8sPodName,
					Namespace: FakeK8sPodNamespace,
				},
			},
		},
	}
}

func FakeIdentity() *identitypb.Identity {
	return &identitypb.Identity{
		Id:          FakeIdentityID,
		TrustZoneId: FakeTrustZoneID,
		ClusterId:   FakeClusterID,
		WorkloadId:  FakeWorkloadID,
		SpiffeId:    FakeSPIFFEID,
		ParentId:    FakeParentID,
		Selectors: []*identitypb.Selector{
			{
				Type:  FakeSelectorType,
				Value: FakeSelectorValue,
			},
		},
		Federations: []*identitypb.IdentityFederation{
			{
				TrustZoneId: PtrOf(FakeTrustZoneID),
			},
		},
	}
}
