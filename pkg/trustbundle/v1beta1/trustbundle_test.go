// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1beta1_test

import (
	"testing"

	trustbundlepb "github.com/cofide/cofide-api-sdk/gen/go/cofide/common/trust_bundle/v1beta1"
	"github.com/cofide/cofide-api-sdk/pkg/trustbundle/v1beta1"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func spireBundle() *types.Bundle {
	return &types.Bundle{
		TrustDomain: "example.org",
		X509Authorities: []*types.X509Certificate{
			{Asn1: []byte("cert-1")},
			{Asn1: []byte("cert-2"), Tainted: true},
		},
		JwtAuthorities: []*types.JWTKey{
			{PublicKey: []byte("key-1"), KeyId: "kid-1", ExpiresAt: 1700000000},
			{PublicKey: []byte("key-2"), KeyId: "kid-2", Tainted: true},
		},
		RefreshHint:    300,
		SequenceNumber: 42,
	}
}

func cofideBundle() *trustbundlepb.Bundle {
	return &trustbundlepb.Bundle{
		TrustDomain: "example.org",
		X509Authorities: []*trustbundlepb.X509Certificate{
			{Asn1: []byte("cert-1")},
			{Asn1: []byte("cert-2"), Tainted: true},
		},
		JwtAuthorities: []*trustbundlepb.JWTKey{
			{PublicKey: []byte("key-1"), KeyId: "kid-1", ExpiresAt: &timestamppb.Timestamp{Seconds: 1700000000}},
			{PublicKey: []byte("key-2"), KeyId: "kid-2", Tainted: true},
		},
		RefreshHint:    &durationpb.Duration{Seconds: 300},
		SequenceNumber: 42,
	}
}

func TestFromSPIRE(t *testing.T) {
	assert.True(t, proto.Equal(cofideBundle(), v1beta1.FromSPIRE(spireBundle())))
}

func TestToSPIRE(t *testing.T) {
	assert.True(t, proto.Equal(spireBundle(), v1beta1.ToSPIRE(cofideBundle())))
}

func TestNil(t *testing.T) {
	assert.Nil(t, v1beta1.FromSPIRE(nil))
	assert.Nil(t, v1beta1.ToSPIRE(nil))
}

func TestEmpty(t *testing.T) {
	assert.True(t, proto.Equal(&trustbundlepb.Bundle{}, v1beta1.FromSPIRE(&types.Bundle{})))
	assert.True(t, proto.Equal(&types.Bundle{}, v1beta1.ToSPIRE(&trustbundlepb.Bundle{})))
}

func TestNoAliasing(t *testing.T) {
	spire := spireBundle()
	cofide := v1beta1.FromSPIRE(spire)
	spire.X509Authorities[0].Asn1[0] = 'X'
	spire.JwtAuthorities[0].PublicKey[0] = 'X'
	assert.True(t, proto.Equal(cofideBundle(), cofide))

	cofide = cofideBundle()
	spire = v1beta1.ToSPIRE(cofide)
	cofide.X509Authorities[0].Asn1[0] = 'X'
	cofide.JwtAuthorities[0].PublicKey[0] = 'X'
	assert.True(t, proto.Equal(spireBundle(), spire))
}

// TestFractionalSeconds checks that converting to SPIRE drops fractions of a
// second, so that a key never expires later than its Cofide expiry.
func TestFractionalSeconds(t *testing.T) {
	bundle := &trustbundlepb.Bundle{
		JwtAuthorities: []*trustbundlepb.JWTKey{
			{ExpiresAt: &timestamppb.Timestamp{Seconds: 1700000000, Nanos: 999999999}},
		},
		RefreshHint: &durationpb.Duration{Seconds: 300, Nanos: 999999999},
	}
	spire := v1beta1.ToSPIRE(bundle)
	assert.Equal(t, int64(1700000000), spire.GetJwtAuthorities()[0].GetExpiresAt())
	assert.Equal(t, int64(300), spire.GetRefreshHint())
}

// TestFieldsMatch checks that the Cofide messages have the same field names and
// numbers as the SPIRE messages they mirror, so that a field added to SPIRE's
// messages on an upgrade of spire-api-sdk fails here rather than being dropped
// by conversion. Field types are not compared, since the Cofide messages use
// well-known types where SPIRE uses integer seconds.
func TestFieldsMatch(t *testing.T) {
	pairs := []struct {
		spire, cofide protoreflect.MessageDescriptor
	}{
		{(&types.Bundle{}).ProtoReflect().Descriptor(), (&trustbundlepb.Bundle{}).ProtoReflect().Descriptor()},
		{(&types.X509Certificate{}).ProtoReflect().Descriptor(), (&trustbundlepb.X509Certificate{}).ProtoReflect().Descriptor()},
		{(&types.JWTKey{}).ProtoReflect().Descriptor(), (&trustbundlepb.JWTKey{}).ProtoReflect().Descriptor()},
	}
	for _, pair := range pairs {
		t.Run(string(pair.spire.Name()), func(t *testing.T) {
			spireFields := fieldSummaries(pair.spire)
			require.NotEmpty(t, spireFields)
			assert.Equal(t, spireFields, fieldSummaries(pair.cofide))
		})
	}
}

type fieldSummary struct {
	Number      protoreflect.FieldNumber
	Name        protoreflect.Name
	Cardinality protoreflect.Cardinality
}

func fieldSummaries(desc protoreflect.MessageDescriptor) []fieldSummary {
	var summaries []fieldSummary
	fields := desc.Fields()
	for i := range fields.Len() {
		field := fields.Get(i)
		summaries = append(summaries, fieldSummary{
			Number:      field.Number(),
			Name:        field.Name(),
			Cardinality: field.Cardinality(),
		})
	}
	return summaries
}
