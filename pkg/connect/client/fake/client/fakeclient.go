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
	clusterv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/cluster/v1alpha1"
	fakeclusterv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/cluster/v1alpha1/fake"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	federationV1Alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/federation/v1alpha1"
	fakefederationV1Alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/federation/v1alpha1/fake"
	trustzonev1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/trustzone/v1alpha1"
	faketrustzonev1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/trustzone/v1alpha1/fake"
)

type fakeClientSet struct {
	trustZoneV1Alpha1         trustzonev1alpha1.TrustZoneClient
	clusterV1Alpha1           clusterv1alpha1.ClusterClient
	agentV1Alpha1             agentv1alpha1.AgentClient
	attestationPolicyV1Alpha1 attestationpolicyv1alpha1.AttestationPolicyClient
	apBindingV1Alpha1         apbindingv1alpha1.APBindingClient
	federationV1Alpha1        federationV1Alpha1.FederationClient
}

// New instantiates a new ClientSet that fakes communication with a Connect API.
func New(fake *fakeconnect.FakeConnect) client.ClientSet {
	return &fakeClientSet{
		trustZoneV1Alpha1:         faketrustzonev1alpha1.New(fake),
		clusterV1Alpha1:           fakeclusterv1alpha1.New(fake),
		agentV1Alpha1:             fakeagentv1alpha1.New(fake),
		attestationPolicyV1Alpha1: fakeattestationpolicyv1alpha1.New(fake),
		apBindingV1Alpha1:         fakeapbindingv1alpha1.New(fake),
		federationV1Alpha1:        fakefederationV1Alpha1.New(fake),
	}
}

func (c *fakeClientSet) TrustZoneV1Alpha1() trustzonev1alpha1.TrustZoneClient {
	return c.trustZoneV1Alpha1
}

func (c *fakeClientSet) ClusterV1Alpha1() clusterv1alpha1.ClusterClient {
	return c.clusterV1Alpha1
}

func (c *fakeClientSet) AgentV1Alpha1() agentv1alpha1.AgentClient {
	return c.agentV1Alpha1
}

func (c *fakeClientSet) AttestationPolicyV1Alpha1() attestationpolicyv1alpha1.AttestationPolicyClient {
	return c.attestationPolicyV1Alpha1
}

func (c *fakeClientSet) APBindingV1Alpha1() apbindingv1alpha1.APBindingClient {
	return c.apBindingV1Alpha1
}

func (c *fakeClientSet) FederationV1Alpha1() federationV1Alpha1.FederationClient {
	return c.federationV1Alpha1
}
