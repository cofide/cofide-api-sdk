// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"github.com/cofide/cofide-api-sdk/pkg/connect/client"
	agentv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/agent/v1alpha1"
	fakeagentv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/agent/v1alpha1/fake"
	apbindingv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/apbinding/v1alpha1"
	fakeapbindingv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/apbinding/v1alpha1/fake"
	attestationpolicyv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/attestationpolicy/v1alpha1"
	fakeattestationpolicyv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/attestationpolicy/v1alpha1/fake"
	auditv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/audit/v1alpha1"
	fakeauditv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/audit/v1alpha1/fake"
	clusterv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/cluster/v1alpha1"
	fakeclusterv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/cluster/v1alpha1/fake"
	datastorev1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/datastore/v1alpha1"
	fakedatastorev1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/datastore/v1alpha1/fake"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	federationV1Alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/federation/v1alpha1"
	fakefederationV1Alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/federation/v1alpha1/fake"
	identityv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/identity/v1alpha1"
	fakeidentityv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/identity/v1alpha1/fake"
	organizationv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/organization/v1alpha1"
	fakeorganizationv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/organization/v1alpha1/fake"
	rolebindingv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/rolebinding/v1alpha1"
	fakerolebindingv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/rolebinding/v1alpha1/fake"
	trustzonev1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/trustzone/v1alpha1"
	faketrustzonev1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/trustzone/v1alpha1/fake"
	trustzoneserverv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/trustzoneserver/v1alpha1"
	faketrustzoneserverv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/trustzoneserver/v1alpha1/fake"
	workloadv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/workload/v1alpha1"
	fakeworkloadv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/workload/v1alpha1/fake"
)

type fakeClientSet struct {
	agentV1Alpha1             agentv1alpha1.AgentClient
	apBindingV1Alpha1         apbindingv1alpha1.APBindingClient
	attestationPolicyV1Alpha1 attestationpolicyv1alpha1.AttestationPolicyClient
	auditV1Alpha1             auditv1alpha1.AuditClient
	clusterV1Alpha1           clusterv1alpha1.ClusterClient
	datastoreV1Alpha1         datastorev1alpha1.DataStoreClient
	federationV1Alpha1        federationV1Alpha1.FederationClient
	identityV1Alpha1          identityv1alpha1.IdentityClient
	organizationV1Alpha1      organizationv1alpha1.OrganizationClient
	roleBindingV1Alpha1       rolebindingv1alpha1.RoleBindingClient
	trustZoneV1Alpha1         trustzonev1alpha1.TrustZoneClient
	trustZoneServerV1Alpha1   trustzoneserverv1alpha1.TrustZoneServerClient
	workloadV1Alpha1          workloadv1alpha1.WorkloadClient
}

// New instantiates a new ClientSet that fakes communication with a Connect API.
func New(fake *fakeconnect.FakeConnect) client.ClientSet {
	return &fakeClientSet{
		agentV1Alpha1:             fakeagentv1alpha1.New(fake),
		apBindingV1Alpha1:         fakeapbindingv1alpha1.New(fake),
		attestationPolicyV1Alpha1: fakeattestationpolicyv1alpha1.New(fake),
		auditV1Alpha1:             fakeauditv1alpha1.New(fake),
		clusterV1Alpha1:           fakeclusterv1alpha1.New(fake),
		datastoreV1Alpha1:         fakedatastorev1alpha1.New(fake),
		federationV1Alpha1:        fakefederationV1Alpha1.New(fake),
		identityV1Alpha1:          fakeidentityv1alpha1.New(fake),
		organizationV1Alpha1:      fakeorganizationv1alpha1.New(fake),
		roleBindingV1Alpha1:       fakerolebindingv1alpha1.New(fake),
		trustZoneV1Alpha1:         faketrustzonev1alpha1.New(fake),
		trustZoneServerV1Alpha1:   faketrustzoneserverv1alpha1.New(fake),
		workloadV1Alpha1:          fakeworkloadv1alpha1.New(fake),
	}
}

func (c *fakeClientSet) AgentV1Alpha1() agentv1alpha1.AgentClient {
	return c.agentV1Alpha1
}

func (c *fakeClientSet) APBindingV1Alpha1() apbindingv1alpha1.APBindingClient {
	return c.apBindingV1Alpha1
}

func (c *fakeClientSet) AttestationPolicyV1Alpha1() attestationpolicyv1alpha1.AttestationPolicyClient {
	return c.attestationPolicyV1Alpha1
}

func (c *fakeClientSet) AuditV1Alpha1() auditv1alpha1.AuditClient {
	return c.auditV1Alpha1
}

func (c *fakeClientSet) ClusterV1Alpha1() clusterv1alpha1.ClusterClient {
	return c.clusterV1Alpha1
}

func (c *fakeClientSet) DataStoreV1Alpha1() datastorev1alpha1.DataStoreClient {
	return c.datastoreV1Alpha1
}

func (c *fakeClientSet) FederationV1Alpha1() federationV1Alpha1.FederationClient {
	return c.federationV1Alpha1
}

func (c *fakeClientSet) IdentityV1Alpha1() identityv1alpha1.IdentityClient {
	return c.identityV1Alpha1
}

func (c *fakeClientSet) OrganizationV1Alpha1() organizationv1alpha1.OrganizationClient {
	return c.organizationV1Alpha1
}

func (c *fakeClientSet) RoleBindingV1Alpha1() rolebindingv1alpha1.RoleBindingClient {
	return c.roleBindingV1Alpha1
}

func (c *fakeClientSet) TrustZoneV1Alpha1() trustzonev1alpha1.TrustZoneClient {
	return c.trustZoneV1Alpha1
}

func (c *fakeClientSet) TrustZoneServerV1Alpha1() trustzoneserverv1alpha1.TrustZoneServerClient {
	return c.trustZoneServerV1Alpha1
}

func (c *fakeClientSet) WorkloadV1Alpha1() workloadv1alpha1.WorkloadClient {
	return c.workloadV1Alpha1
}
