package fakeconnect

import (
	"sync"

	agentpb "github.com/cofide/cofide-api-sdk/gen/go/proto/agent/v1alpha1"
	attestationpolicypb "github.com/cofide/cofide-api-sdk/gen/go/proto/attestation_policy/v1alpha1"
	clusterpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cluster/v1alpha1"
	federatedservicepb "github.com/cofide/cofide-api-sdk/gen/go/proto/federated_service/v1alpha1"
	trustzonepb "github.com/cofide/cofide-api-sdk/gen/go/proto/trust_zone/v1alpha1"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// FakeConnect implements the state for a fake Connect service.
type FakeConnect struct {
	Mu                  sync.Mutex
	TrustZones          map[string]*trustzonepb.TrustZone
	TrustZoneBundles    map[string]*types.Bundle
	Clusters            map[string]*clusterpb.Cluster
	Agents              map[string]*agentpb.Agent
	AgentJoinTokens     map[string]map[string]string
	AgentStatus         map[string]*agentpb.AgentStatus
	FederatedServices   map[string]*federatedservicepb.FederatedService
	AttestationPolicies map[string]*attestationpolicypb.AttestationPolicy
}

func New() *FakeConnect {
	return &FakeConnect{
		TrustZones:          make(map[string]*trustzonepb.TrustZone),
		TrustZoneBundles:    make(map[string]*types.Bundle),
		Clusters:            make(map[string]*clusterpb.Cluster),
		Agents:              make(map[string]*agentpb.Agent),
		AgentJoinTokens:     make(map[string]map[string]string),
		AgentStatus:         make(map[string]*agentpb.AgentStatus),
		FederatedServices:   make(map[string]*federatedservicepb.FederatedService),
		AttestationPolicies: make(map[string]*attestationpolicypb.AttestationPolicy),
	}
}

func (f *FakeConnect) ValidateTrustZone(trustZoneID string) error {
	if _, ok := f.TrustZones[trustZoneID]; !ok {
		return status.Error(codes.InvalidArgument, "invalid trust zone")
	}
	return nil
}

func (f *FakeConnect) ValidateCluster(clusterID string) error {
	if _, ok := f.Clusters[clusterID]; !ok {
		return status.Error(codes.InvalidArgument, "invalid cluster")
	}
	return nil
}

func (f *FakeConnect) ValidateAgent(agentID string) error {
	if _, ok := f.Agents[agentID]; !ok {
		return status.Error(codes.InvalidArgument, "invalid agent")
	}
	return nil
}

func (f *FakeConnect) ValidateFederatedService(federatedServiceID string) error {
	if _, ok := f.FederatedServices[federatedServiceID]; !ok {
		return status.Error(codes.InvalidArgument, "invalid federated service")
	}
	return nil
}

func (f *FakeConnect) ValidateAttestationPolicy(policyID string) error {
	if _, ok := f.AttestationPolicies[policyID]; !ok {
		return status.Error(codes.InvalidArgument, "invalid attestation policy")
	}
	return nil
}
