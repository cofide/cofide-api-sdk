// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1beta1_test

import (
	"testing"

	trustzonepb "github.com/cofide/cofide-api-sdk/gen/go/cofide/trust_zone/v1beta1"
	trustzonev1alpha1pb "github.com/cofide/cofide-api-sdk/gen/go/proto/trust_zone/v1alpha1"
	"github.com/cofide/cofide-api-sdk/pkg/trustzone/v1beta1"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func v1alpha1TrustZone() *trustzonev1alpha1pb.TrustZone {
	return &trustzonev1alpha1pb.TrustZone{
		Id:                    new("tz-id"),
		OrgId:                 new("org-id"),
		Name:                  "tz-name",
		TrustDomain:           "example.org",
		BundleEndpointUrl:     new("https://bundle.example.org"),
		JwtIssuer:             new("https://issuer.example.org"),
		BundleEndpointProfile: new(trustzonev1alpha1pb.BundleEndpointProfile_BUNDLE_ENDPOINT_PROFILE_HTTPS_WEB),
		IsManagementZone:      true,
		CreatedAt:             &timestamppb.Timestamp{Seconds: 1700000000},
		LastUpdatedAt:         &timestamppb.Timestamp{Seconds: 1700000001},
	}
}

func v1beta1TrustZone() *trustzonepb.TrustZone {
	return &trustzonepb.TrustZone{
		Id:                    "tz-id",
		OrgId:                 "org-id",
		Name:                  "tz-name",
		TrustDomain:           "example.org",
		BundleEndpointUrl:     "https://bundle.example.org",
		JwtIssuerUrl:          "https://issuer.example.org",
		BundleEndpointProfile: trustzonepb.BundleEndpointProfile_BUNDLE_ENDPOINT_PROFILE_HTTPS_WEB,
		IsManagementZone:      true,
		CreatedAt:             &timestamppb.Timestamp{Seconds: 1700000000},
		LastUpdatedAt:         &timestamppb.Timestamp{Seconds: 1700000001},
	}
}

func TestFromV1Alpha1(t *testing.T) {
	assert.True(t, proto.Equal(v1beta1TrustZone(), v1beta1.FromV1Alpha1(v1alpha1TrustZone())))
}

func TestToV1Alpha1(t *testing.T) {
	assert.True(t, proto.Equal(v1alpha1TrustZone(), v1beta1.ToV1Alpha1(v1beta1TrustZone())))
}

func TestNil(t *testing.T) {
	assert.Nil(t, v1beta1.FromV1Alpha1(nil))
	assert.Nil(t, v1beta1.ToV1Alpha1(nil))
}

// TestEmpty checks that converting to v1alpha1 leaves optional fields unset
// rather than setting them to their zero values.
func TestEmpty(t *testing.T) {
	assert.True(t, proto.Equal(&trustzonepb.TrustZone{}, v1beta1.FromV1Alpha1(&trustzonev1alpha1pb.TrustZone{})))
	assert.True(t, proto.Equal(&trustzonev1alpha1pb.TrustZone{}, v1beta1.ToV1Alpha1(&trustzonepb.TrustZone{})))
}

func TestDeprecatedBundleDropped(t *testing.T) {
	trustZone := v1alpha1TrustZone()
	trustZone.Bundle = &types.Bundle{TrustDomain: "example.org"} //nolint:staticcheck // Checking the deprecated field is dropped.
	assert.True(t, proto.Equal(v1beta1TrustZone(), v1beta1.FromV1Alpha1(trustZone)))
}

func TestNoAliasing(t *testing.T) {
	v1alpha1 := v1alpha1TrustZone()
	v1beta1TZ := v1beta1.FromV1Alpha1(v1alpha1)
	v1alpha1.CreatedAt.Seconds = 0
	v1alpha1.LastUpdatedAt.Seconds = 0
	assert.True(t, proto.Equal(v1beta1TrustZone(), v1beta1TZ))

	v1beta1TZ = v1beta1TrustZone()
	v1alpha1 = v1beta1.ToV1Alpha1(v1beta1TZ)
	v1beta1TZ.CreatedAt.Seconds = 0
	v1beta1TZ.LastUpdatedAt.Seconds = 0
	assert.True(t, proto.Equal(v1alpha1TrustZone(), v1alpha1))
}

// TestFieldsConverted checks that the conversion covers every field of both
// messages, so that a field added to either fails here rather than being
// dropped by conversion. Update the conversion and this list together.
func TestFieldsConverted(t *testing.T) {
	assert.ElementsMatch(t, []protoreflect.Name{
		"id", "org_id", "name", "trust_domain", "bundle_endpoint_url", "jwt_issuer_url",
		"bundle_endpoint_profile", "is_management_zone", "created_at", "last_updated_at",
	}, fieldNames((&trustzonepb.TrustZone{}).ProtoReflect().Descriptor()))
	assert.ElementsMatch(t, []protoreflect.Name{
		"id", "org_id", "name", "trust_domain", "bundle_endpoint_url", "jwt_issuer",
		"bundle_endpoint_profile", "is_management_zone", "created_at", "last_updated_at",
		"bundle",
	}, fieldNames((&trustzonev1alpha1pb.TrustZone{}).ProtoReflect().Descriptor()))
}

// TestBundleEndpointProfilesMatch checks that the two BundleEndpointProfile
// enums have the same values, since the conversion maps them by number.
func TestBundleEndpointProfilesMatch(t *testing.T) {
	v1alpha1Values := trustzonev1alpha1pb.BundleEndpointProfile(0).Descriptor().Values()
	v1beta1Values := trustzonepb.BundleEndpointProfile(0).Descriptor().Values()
	assert.Equal(t, v1alpha1Values.Len(), v1beta1Values.Len())
	for i := range v1alpha1Values.Len() {
		value := v1alpha1Values.Get(i)
		assert.Equal(t, value.Name(), v1beta1Values.ByNumber(value.Number()).Name())
	}
}

func fieldNames(desc protoreflect.MessageDescriptor) []protoreflect.Name {
	var names []protoreflect.Name
	fields := desc.Fields()
	for i := range fields.Len() {
		names = append(names, fields.Get(i).Name())
	}
	return names
}
