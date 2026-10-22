package fakeconnect

import (
	"cmp"
	"iter"
	"maps"
	"slices"
	"sync"

	agentpb "github.com/cofide/cofide-api-sdk/gen/go/proto/agent/v1alpha1"
	attestationpolicypb "github.com/cofide/cofide-api-sdk/gen/go/proto/attestation_policy/v1alpha1"
	auditpb "github.com/cofide/cofide-api-sdk/gen/go/proto/audit/v1alpha1"
	cloudaccountpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_account/v1alpha1"
	clusterpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cluster/v1alpha1"
	datastoresvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/datastore_service/v1alpha1"
	exchangepolicypb "github.com/cofide/cofide-api-sdk/gen/go/proto/exchange_policy/v1alpha1"
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
//
// Mu guards all of the maps below. Fake client methods must hold it for the
// whole of any operation that reads or writes them: Lock for operations that
// modify state, RLock for read-only operations such as Get and List. This
// includes the Validate* helpers, which read the maps without locking and so
// must only be called with Mu held.
type FakeConnect struct {
	Mu                  sync.RWMutex
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
	ExchangePolicies    map[string]*exchangepolicypb.ExchangePolicy
	Federations         map[string]*federationpb.Federation
	AttestedNodes       map[string]*datastoresvcpb.AttestedNode
	Workloads           map[string]*workloadpb.Workload
	WorkloadEvents      []*workloadpb.WorkloadEvent
	Identities          map[string]*identitypb.Identity
	RoleBindings        map[string]*rolebindingpb.RoleBinding
	AuditEvents         map[string]*auditpb.Event
	CloudAccounts       map[string]*cloudaccountpb.CloudAccount
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
		ExchangePolicies:    make(map[string]*exchangepolicypb.ExchangePolicy),
		Federations:         make(map[string]*federationpb.Federation),
		AttestedNodes:       make(map[string]*datastoresvcpb.AttestedNode),
		Workloads:           make(map[string]*workloadpb.Workload),
		WorkloadEvents:      []*workloadpb.WorkloadEvent{},
		Identities:          make(map[string]*identitypb.Identity),
		RoleBindings:        make(map[string]*rolebindingpb.RoleBinding),
		AuditEvents:         make(map[string]*auditpb.Event),
		CloudAccounts:       make(map[string]*cloudaccountpb.CloudAccount),
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

// SortedValues returns an iterator over the values of m in ascending key order.
// Fake List methods use it so that their results are deterministic, which Go map
// iteration is not.
func SortedValues[M ~map[K]V, K cmp.Ordered, V any](m M) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, k := range slices.Sorted(maps.Keys(m)) {
			if !yield(m[k]) {
				return
			}
		}
	}
}
