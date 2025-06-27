// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	identitysvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/identity_service/v1alpha1"
	identitypb "github.com/cofide/cofide-api-sdk/gen/go/proto/identity/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	identityv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/identity/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type fakeIdentityClient struct {
	fake *fakeconnect.FakeConnect
}

// New instantiates a new IdentityClient for communication with a fake Connect API.
func New(fake *fakeconnect.FakeConnect) identityv1alpha1.IdentityClient {
	return &fakeIdentityClient{
		fake: fake,
	}
}

func (c *fakeIdentityClient) GetIdentity(ctx context.Context, identityID string) (*identitypb.Identity, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	identity, ok := c.fake.Identities[identityID]
	if !ok {
		return nil, status.Error(codes.NotFound, "identity not found")
	}
	return clone(identity), nil
}

func (c *fakeIdentityClient) ListIdentities(ctx context.Context, filter *identitysvcpb.ListIdentitiesRequest_Filter) ([]*identitypb.Identity, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	identities := []*identitypb.Identity{}
	for _, identity := range c.fake.Identities {
		if identityMatches(identity, filter) {
			identities = append(identities, clone(identity))
		}
	}
	return identities, nil
}

func identityMatches(identity *identitypb.Identity, filter *identitysvcpb.ListIdentitiesRequest_Filter) bool {
	if filter == nil {
		return true
	}
	if filter.OrgId != nil && identity.GetOrgId() != *filter.OrgId {
		return false
	}
	if filter.TrustZoneId != nil && identity.GetTrustZoneId() != *filter.TrustZoneId {
		return false
	}
	if filter.ClusterId != nil && identity.GetClusterId() != *filter.ClusterId {
		return false
	}
	if filter.AttestationPolicyId != nil && identity.GetAttestationPolicyId() != *filter.AttestationPolicyId {
		return false
	}
	if filter.ApBindingId != nil && identity.GetApBindingId() != *filter.ApBindingId {
		return false
	}
	if filter.WorkloadId != nil && identity.GetWorkloadId() != *filter.WorkloadId {
		return false
	}
	if filter.SpiffeId != nil && identity.GetSpiffeId() != *filter.SpiffeId {
		return false
	}
	return true
}

func clone(identity *identitypb.Identity) *identitypb.Identity {
	return proto.Clone(identity).(*identitypb.Identity)
}
