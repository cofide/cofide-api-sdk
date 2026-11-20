// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1beta1

import (
	trustzonepb "github.com/cofide/cofide-api-sdk/gen/go/cofide/trust_zone/v1beta1"
)

// CofideAgentCredential authenticates the registration of a Cofide Agent with its trust zone.
// Use JoinToken or K8sPSAT.
type CofideAgentCredential interface {
	applyToCofideAgent(req *trustzonepb.RegisterCofideAgentRequest)
}

// TrustZoneServerCredential authenticates the registration of a trust zone server with its trust zone.
// Use JoinToken or K8sPSAT.
type TrustZoneServerCredential interface {
	applyToTrustZoneServer(req *trustzonepb.RegisterTrustZoneServerRequest)
}

// JoinToken is a CofideAgentCredential and TrustZoneServerCredential holding a join token issued for the trust zone.
type JoinToken string

func (t JoinToken) applyToCofideAgent(req *trustzonepb.RegisterCofideAgentRequest) {
	req.Auth = &trustzonepb.RegisterCofideAgentRequest_JoinToken{JoinToken: string(t)}
}

func (t JoinToken) applyToTrustZoneServer(req *trustzonepb.RegisterTrustZoneServerRequest) {
	req.Auth = &trustzonepb.RegisterTrustZoneServerRequest_JoinToken{JoinToken: string(t)}
}

// K8sPSAT is a CofideAgentCredential and TrustZoneServerCredential holding a Kubernetes projected service account token issued by a cluster in the trust zone.
type K8sPSAT string

func (t K8sPSAT) applyToCofideAgent(req *trustzonepb.RegisterCofideAgentRequest) {
	req.Auth = &trustzonepb.RegisterCofideAgentRequest_K8SPsat{K8SPsat: string(t)}
}

func (t K8sPSAT) applyToTrustZoneServer(req *trustzonepb.RegisterTrustZoneServerRequest) {
	req.Auth = &trustzonepb.RegisterTrustZoneServerRequest_K8SPsat{K8SPsat: string(t)}
}
