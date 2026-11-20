// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

// Package v1beta1 converts between the v1beta1 and v1alpha1 trust zone
// messages.
//
// The v1alpha1 message marks most of its fields optional, where the v1beta1
// message does not, so converting to v1alpha1 leaves an optional field unset
// rather than setting it to its zero value. A v1alpha1 optional field set to
// its zero value therefore comes back unset from a round trip. The deprecated
// v1alpha1 bundle field has no v1beta1 equivalent, so converting to v1beta1
// drops it.
package v1beta1

import (
	trustzonepb "github.com/cofide/cofide-api-sdk/gen/go/cofide/trust_zone/v1beta1"
	trustzonev1alpha1pb "github.com/cofide/cofide-api-sdk/gen/go/proto/trust_zone/v1alpha1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// FromV1Alpha1 converts a v1alpha1 trust zone to a v1beta1 trust zone.
// It returns nil if trustZone is nil. The result shares no memory with trustZone.
func FromV1Alpha1(trustZone *trustzonev1alpha1pb.TrustZone) *trustzonepb.TrustZone {
	if trustZone == nil {
		return nil
	}
	return &trustzonepb.TrustZone{
		Id:                    trustZone.GetId(),
		OrgId:                 trustZone.GetOrgId(),
		Name:                  trustZone.GetName(),
		TrustDomain:           trustZone.GetTrustDomain(),
		BundleEndpointUrl:     trustZone.GetBundleEndpointUrl(),
		JwtIssuerUrl:          trustZone.GetJwtIssuer(),
		BundleEndpointProfile: trustzonepb.BundleEndpointProfile(trustZone.GetBundleEndpointProfile()),
		IsManagementZone:      trustZone.GetIsManagementZone(),
		CreatedAt:             cloneTimestamp(trustZone.GetCreatedAt()),
		LastUpdatedAt:         cloneTimestamp(trustZone.GetLastUpdatedAt()),
	}
}

// ToV1Alpha1 converts a v1beta1 trust zone to a v1alpha1 trust zone.
// It returns nil if trustZone is nil. The result shares no memory with trustZone.
func ToV1Alpha1(trustZone *trustzonepb.TrustZone) *trustzonev1alpha1pb.TrustZone {
	if trustZone == nil {
		return nil
	}
	return &trustzonev1alpha1pb.TrustZone{
		Id:                    optional(trustZone.GetId()),
		OrgId:                 optional(trustZone.GetOrgId()),
		Name:                  trustZone.GetName(),
		TrustDomain:           trustZone.GetTrustDomain(),
		BundleEndpointUrl:     optional(trustZone.GetBundleEndpointUrl()),
		JwtIssuer:             optional(trustZone.GetJwtIssuerUrl()),
		BundleEndpointProfile: optional(trustzonev1alpha1pb.BundleEndpointProfile(trustZone.GetBundleEndpointProfile())),
		IsManagementZone:      trustZone.GetIsManagementZone(),
		CreatedAt:             cloneTimestamp(trustZone.GetCreatedAt()),
		LastUpdatedAt:         cloneTimestamp(trustZone.GetLastUpdatedAt()),
	}
}

// optional returns a pointer to v, or nil if v is its zero value.
func optional[T comparable](v T) *T {
	var zero T
	if v == zero {
		return nil
	}
	return &v
}

func cloneTimestamp(ts *timestamppb.Timestamp) *timestamppb.Timestamp {
	if ts == nil {
		return nil
	}
	return proto.CloneOf(ts)
}
