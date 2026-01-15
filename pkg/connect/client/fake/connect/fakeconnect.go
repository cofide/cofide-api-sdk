package fakeconnect

import (
	"sync"

	agentpb "github.com/cofide/cofide-api-sdk/gen/go/proto/agent/v1alpha1"
	apbindingpb "github.com/cofide/cofide-api-sdk/gen/go/proto/ap_binding/v1alpha1"
	attestationpolicypb "github.com/cofide/cofide-api-sdk/gen/go/proto/attestation_policy/v1alpha1"
	clusterpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cluster/v1alpha1"
	datastoresvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/datastore_service/v1alpha1"
	federatedservicepb "github.com/cofide/cofide-api-sdk/gen/go/proto/federated_service/v1alpha1"
	federationpb "github.com/cofide/cofide-api-sdk/gen/go/proto/federation/v1alpha1"
	identitypb "github.com/cofide/cofide-api-sdk/gen/go/proto/identity/v1alpha1"
	organizationpb "github.com/cofide/cofide-api-sdk/gen/go/proto/organization/v1alpha1"
	rolebindingpb "github.com/cofide/cofide-api-sdk/gen/go/proto/role_binding/v1alpha1"
	trustzonepb "github.com/cofide/cofide-api-sdk/gen/go/proto/trust_zone/v1alpha1"
	trustzoneserverpb "github.com/cofide/cofide-api-sdk/gen/go/proto/trust_zone_server/v1alpha1"
	workloadpb "github.com/cofide/cofide-api-sdk/gen/go/proto/workload/v1alpha1"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// FakeConnect implements the state for a fake Connect service.
type FakeConnect struct {
	Mu                  sync.Mutex
	Organizations       map[string]*organizationpb.Organization
	TrustZones          map[string]*trustzonepb.TrustZone
	TrustZoneBundles    map[string]*types.Bundle
	TrustZoneServers    map[string]*trustzoneserverpb.TrustZoneServer
	Clusters            map[string]*clusterpb.Cluster
	Agents              map[string]*agentpb.Agent
	AgentJoinTokens     map[string]map[string]string
	AgentStatus         map[string]*agentpb.AgentStatus
	FederatedServices   map[string]*federatedservicepb.FederatedService
	AttestationPolicies map[string]*attestationpolicypb.AttestationPolicy
	APBindings          map[string]*apbindingpb.APBinding
	Federations         map[string]*federationpb.Federation
	AttestedNodes       map[string]*datastoresvcpb.AttestedNode
	Workloads           map[string]*workloadpb.Workload
	Identities          map[string]*identitypb.Identity
	RoleBindings        map[string]*rolebindingpb.RoleBinding
}

func New() *FakeConnect {
	return &FakeConnect{
		Organizations:       make(map[string]*organizationpb.Organization),
		TrustZones:          make(map[string]*trustzonepb.TrustZone),
		TrustZoneBundles:    make(map[string]*types.Bundle),
		TrustZoneServers:    make(map[string]*trustzoneserverpb.TrustZoneServer),
		Clusters:            make(map[string]*clusterpb.Cluster),
		Agents:              make(map[string]*agentpb.Agent),
		AgentJoinTokens:     make(map[string]map[string]string),
		AgentStatus:         make(map[string]*agentpb.AgentStatus),
		FederatedServices:   make(map[string]*federatedservicepb.FederatedService),
		AttestationPolicies: make(map[string]*attestationpolicypb.AttestationPolicy),
		APBindings:          make(map[string]*apbindingpb.APBinding),
		Federations:         make(map[string]*federationpb.Federation),
		AttestedNodes:       make(map[string]*datastoresvcpb.AttestedNode),
		Workloads:           make(map[string]*workloadpb.Workload),
		Identities:          make(map[string]*identitypb.Identity),
		RoleBindings:        make(map[string]*rolebindingpb.RoleBinding),
	}
}

func (f *FakeConnect) ValidateOrganization(organizationID string) error {
	if _, ok := f.Organizations[organizationID]; !ok {
		return status.Error(codes.InvalidArgument, "invalid organization")
	}
	return nil
}

func (f *FakeConnect) ValidateTrustZone(trustZoneID string) error {
	if _, ok := f.TrustZones[trustZoneID]; !ok {
		return status.Error(codes.InvalidArgument, "invalid trust zone")
	}
	return nil
}

func (f *FakeConnect) ValidateTrustZoneServer(trustZoneServerID string) error {
	if _, ok := f.TrustZoneServers[trustZoneServerID]; !ok {
		return status.Error(codes.InvalidArgument, "invalid trust zone server")
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

func (f *FakeConnect) ValidateAPBinding(bindingID string) error {
	if _, ok := f.APBindings[bindingID]; !ok {
		return status.Error(codes.InvalidArgument, "invalid attestation policy binding")
	}
	return nil
}

func (f *FakeConnect) ValidateFederation(federationID string) error {
	if _, ok := f.Federations[federationID]; !ok {
		return status.Error(codes.InvalidArgument, "invalid federation")
	}
	return nil
}

func (f *FakeConnect) ValidateWorkload(workloadID string) error {
	if _, ok := f.Workloads[workloadID]; !ok {
		return status.Error(codes.InvalidArgument, "invalid workload")
	}
	return nil
}

func (f *FakeConnect) ValidateIdentity(identityID string) error {
	if _, ok := f.Identities[identityID]; !ok {
		return status.Error(codes.InvalidArgument, "invalid identity")
	}
	return nil
}

func (f *FakeConnect) ValidateRoleBinding(roleBindingID string) error {
	if _, ok := f.RoleBindings[roleBindingID]; !ok {
		return status.Error(codes.InvalidArgument, "invalid role binding")
	}
	return nil
}
