// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

// Package v1beta1 converts between the Cofide v1beta1 trust bundle messages
// and the SPIRE API trust bundle messages they mirror.
//
// SPIRE represents the refresh hint and JWT key expiry as integer seconds,
// with zero meaning unset, where the Cofide messages use Duration and
// Timestamp. Converting to SPIRE drops any fraction of a second, so a key
// never expires later than its Cofide expiry, and a refresh hint under a
// second becomes no hint.
package v1beta1

import (
	"bytes"

	trustbundlepb "github.com/cofide/cofide-api-sdk/gen/go/cofide/common/trust_bundle/v1beta1"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// FromSPIRE converts a SPIRE API trust bundle to a Cofide trust bundle.
// It returns nil if bundle is nil. The result shares no memory with bundle.
func FromSPIRE(bundle *types.Bundle) *trustbundlepb.Bundle {
	if bundle == nil {
		return nil
	}
	return &trustbundlepb.Bundle{
		TrustDomain:     bundle.GetTrustDomain(),
		X509Authorities: convertSlice(bundle.GetX509Authorities(), x509CertificateFromSPIRE),
		JwtAuthorities:  convertSlice(bundle.GetJwtAuthorities(), jwtKeyFromSPIRE),
		RefreshHint:     durationFromSeconds(bundle.GetRefreshHint()),
		SequenceNumber:  bundle.GetSequenceNumber(),
	}
}

// ToSPIRE converts a Cofide trust bundle to a SPIRE API trust bundle.
// It returns nil if bundle is nil. The result shares no memory with bundle.
func ToSPIRE(bundle *trustbundlepb.Bundle) *types.Bundle {
	if bundle == nil {
		return nil
	}
	return &types.Bundle{
		TrustDomain:     bundle.GetTrustDomain(),
		X509Authorities: convertSlice(bundle.GetX509Authorities(), x509CertificateToSPIRE),
		JwtAuthorities:  convertSlice(bundle.GetJwtAuthorities(), jwtKeyToSPIRE),
		RefreshHint:     bundle.GetRefreshHint().GetSeconds(),
		SequenceNumber:  bundle.GetSequenceNumber(),
	}
}

func x509CertificateFromSPIRE(cert *types.X509Certificate) *trustbundlepb.X509Certificate {
	return &trustbundlepb.X509Certificate{
		Asn1:    bytes.Clone(cert.GetAsn1()),
		Tainted: cert.GetTainted(),
	}
}

func x509CertificateToSPIRE(cert *trustbundlepb.X509Certificate) *types.X509Certificate {
	return &types.X509Certificate{
		Asn1:    bytes.Clone(cert.GetAsn1()),
		Tainted: cert.GetTainted(),
	}
}

func jwtKeyFromSPIRE(key *types.JWTKey) *trustbundlepb.JWTKey {
	return &trustbundlepb.JWTKey{
		PublicKey: bytes.Clone(key.GetPublicKey()),
		KeyId:     key.GetKeyId(),
		ExpiresAt: timestampFromSeconds(key.GetExpiresAt()),
		Tainted:   key.GetTainted(),
	}
}

func jwtKeyToSPIRE(key *trustbundlepb.JWTKey) *types.JWTKey {
	return &types.JWTKey{
		PublicKey: bytes.Clone(key.GetPublicKey()),
		KeyId:     key.GetKeyId(),
		ExpiresAt: key.GetExpiresAt().GetSeconds(),
		Tainted:   key.GetTainted(),
	}
}

// durationFromSeconds converts SPIRE's integer seconds to a Duration, with
// zero meaning unset.
func durationFromSeconds(seconds int64) *durationpb.Duration {
	if seconds == 0 {
		return nil
	}
	return &durationpb.Duration{Seconds: seconds}
}

// timestampFromSeconds converts SPIRE's seconds since the Unix epoch to a
// Timestamp, with zero meaning unset.
func timestampFromSeconds(seconds int64) *timestamppb.Timestamp {
	if seconds == 0 {
		return nil
	}
	return &timestamppb.Timestamp{Seconds: seconds}
}

// convertSlice converts each element of in, preserving nil.
func convertSlice[In, Out any](in []In, convert func(In) Out) []Out {
	if in == nil {
		return nil
	}
	out := make([]Out, 0, len(in))
	for _, v := range in {
		out = append(out, convert(v))
	}
	return out
}
