// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake_test

import (
	"context"
	"testing"

	trustbundlepb "github.com/cofide/cofide-api-sdk/gen/go/cofide/common/trust_bundle/v1beta1"
	trustzonepb "github.com/cofide/cofide-api-sdk/gen/go/cofide/trust_zone/v1beta1"
	organizationpb "github.com/cofide/cofide-api-sdk/gen/go/proto/organization/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
	trustzonev1beta1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/trustzone/v1beta1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/trustzone/v1beta1/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	orgID      = "org-1"
	otherOrgID = "org-2"
)

// newFake returns a FakeConnect holding two organizations, and a fake client using it.
func newFake(t *testing.T) (*fakeconnect.FakeConnect, trustzonev1beta1.TrustZoneClient) {
	t.Helper()
	fakeConnect := fakeconnect.New()
	for _, id := range []string{orgID, otherOrgID} {
		fakeConnect.Organizations[id] = &organizationpb.Organization{Id: id, Name: id}
	}
	return fakeConnect, fake.New(fakeConnect)
}

// seed stores a trust zone with the given ID, organization, name and trust domain.
func seed(fakeConnect *fakeconnect.FakeConnect, id, orgID, name, trustDomain string) *trustzonepb.TrustZone {
	trustZone := &trustzonepb.TrustZone{Id: id, OrgId: orgID, Name: name, TrustDomain: trustDomain}
	fakeConnect.TrustZones[id] = proto.CloneOf(trustZone)
	return trustZone
}

func requireCode(t *testing.T, err error, code codes.Code) {
	t.Helper()
	require.Error(t, err)
	assert.Equal(t, code, status.Code(err), err.Error())
}

func Test_CreateTrustZone(t *testing.T) {
	t.Parallel()
	fakeConnect, client := newFake(t)

	input := &trustzonepb.TrustZone{
		Id:                    "ignored-id",
		OrgId:                 orgID,
		Name:                  "tz",
		TrustDomain:           "example.org",
		BundleEndpointUrl:     "https://ignored.example.org",
		JwtIssuerUrl:          "https://ignored.example.org",
		BundleEndpointProfile: trustzonepb.BundleEndpointProfile_BUNDLE_ENDPOINT_PROFILE_HTTPS_WEB,
		IsManagementZone:      true,
		LastUpdatedAt:         timestamppb.Now(),
	}
	created, err := client.CreateTrustZone(t.Context(), input)
	require.NoError(t, err)

	assert.NotEmpty(t, created.GetId())
	assert.NotEqual(t, "ignored-id", created.GetId())
	assert.NotNil(t, created.GetCreatedAt())
	assert.EqualExportedValues(t, &trustzonepb.TrustZone{
		Id:               created.GetId(),
		OrgId:            orgID,
		Name:             "tz",
		TrustDomain:      "example.org",
		IsManagementZone: true,
		CreatedAt:        created.GetCreatedAt(),
	}, created)
	assert.EqualExportedValues(t, created, fakeConnect.TrustZones[created.GetId()])
	assert.Equal(t, "ignored-id", input.GetId(), "input must not be modified")
}

func Test_CreateTrustZone_Errors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		trustZone *trustzonepb.TrustZone
		wantCode  codes.Code
	}{
		{
			name:      "organization does not exist",
			trustZone: &trustzonepb.TrustZone{OrgId: "missing", Name: "tz", TrustDomain: "new.example.org"},
			wantCode:  codes.InvalidArgument,
		},
		{
			name:      "empty name",
			trustZone: &trustzonepb.TrustZone{OrgId: orgID, TrustDomain: "new.example.org"},
			wantCode:  codes.InvalidArgument,
		},
		{
			name:      "empty trust domain",
			trustZone: &trustzonepb.TrustZone{OrgId: orgID, Name: "tz"},
			wantCode:  codes.InvalidArgument,
		},
		{
			name:      "name exists in organization",
			trustZone: &trustzonepb.TrustZone{OrgId: orgID, Name: "existing", TrustDomain: "new.example.org"},
			wantCode:  codes.AlreadyExists,
		},
		{
			name:      "trust domain exists",
			trustZone: &trustzonepb.TrustZone{OrgId: otherOrgID, Name: "tz", TrustDomain: "existing.example.org"},
			wantCode:  codes.AlreadyExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fakeConnect, client := newFake(t)
			seed(fakeConnect, "tz-1", orgID, "existing", "existing.example.org")

			_, err := client.CreateTrustZone(t.Context(), tt.trustZone)
			requireCode(t, err, tt.wantCode)
			assert.Len(t, fakeConnect.TrustZones, 1)
		})
	}
}

func Test_CreateTrustZone_SameNameInOtherOrganization(t *testing.T) {
	t.Parallel()
	fakeConnect, client := newFake(t)
	seed(fakeConnect, "tz-1", orgID, "tz", "one.example.org")

	_, err := client.CreateTrustZone(t.Context(), &trustzonepb.TrustZone{OrgId: otherOrgID, Name: "tz", TrustDomain: "two.example.org"})
	require.NoError(t, err)
}

func Test_DeleteTrustZone(t *testing.T) {
	t.Parallel()
	fakeConnect, client := newFake(t)

	err := client.DeleteTrustZone(t.Context(), "tz-1")
	requireCode(t, err, codes.NotFound)

	seed(fakeConnect, "tz-1", orgID, "tz", "example.org")
	fakeConnect.TrustZoneBundles["tz-1"] = &trustbundlepb.Bundle{TrustDomain: "example.org"}
	fakeConnect.AgentJoinTokens["tz-1"] = map[string]string{"cluster-1": "token-1"}

	err = client.DeleteTrustZone(t.Context(), "tz-1")
	require.NoError(t, err)
	assert.NotContains(t, fakeConnect.TrustZones, "tz-1")
	assert.NotContains(t, fakeConnect.TrustZoneBundles, "tz-1")
	assert.NotContains(t, fakeConnect.AgentJoinTokens, "tz-1")
}

func Test_GetTrustZone(t *testing.T) {
	t.Parallel()
	fakeConnect, client := newFake(t)
	want := seed(fakeConnect, "tz-1", orgID, "tz", "example.org")
	seed(fakeConnect, "tz-2", otherOrgID, "tz", "other.example.org")

	got, err := client.GetTrustZone(t.Context(), "tz-1")
	require.NoError(t, err)
	assert.EqualExportedValues(t, want, got)

	_, err = client.GetTrustZone(t.Context(), "missing")
	requireCode(t, err, codes.NotFound)
}

func Test_GetTrustZoneByName(t *testing.T) {
	t.Parallel()
	fakeConnect, client := newFake(t)
	want := seed(fakeConnect, "tz-1", orgID, "tz", "example.org")
	seed(fakeConnect, "tz-2", otherOrgID, "tz", "other.example.org")

	got, err := client.GetTrustZoneByName(t.Context(), orgID, "tz")
	require.NoError(t, err)
	assert.EqualExportedValues(t, want, got)

	_, err = client.GetTrustZoneByName(t.Context(), orgID, "missing")
	requireCode(t, err, codes.NotFound)

	_, err = client.GetTrustZoneByName(t.Context(), "missing", "tz")
	requireCode(t, err, codes.NotFound)
}

func Test_GetTrustZoneByTrustDomain(t *testing.T) {
	t.Parallel()
	fakeConnect, client := newFake(t)
	want := seed(fakeConnect, "tz-1", orgID, "tz", "example.org")
	seed(fakeConnect, "tz-2", otherOrgID, "tz", "other.example.org")

	got, err := client.GetTrustZoneByTrustDomain(t.Context(), "example.org")
	require.NoError(t, err)
	assert.EqualExportedValues(t, want, got)

	_, err = client.GetTrustZoneByTrustDomain(t.Context(), "missing.example.org")
	requireCode(t, err, codes.NotFound)
}

func Test_GetTrustZone_NoAliasing(t *testing.T) {
	t.Parallel()
	fakeConnect, client := newFake(t)
	seed(fakeConnect, "tz-1", orgID, "tz", "example.org")

	got, err := client.GetTrustZone(t.Context(), "tz-1")
	require.NoError(t, err)
	got.Name = "changed"
	assert.Equal(t, "tz", fakeConnect.TrustZones["tz-1"].GetName())
}

func Test_ListTrustZones(t *testing.T) {
	t.Parallel()
	fakeConnect, client := newFake(t)
	tz1 := seed(fakeConnect, "tz-1", orgID, "tz-1", "one.example.org")
	tz2 := seed(fakeConnect, "tz-2", otherOrgID, "tz-2", "two.example.org")
	tz3 := seed(fakeConnect, "tz-3", orgID, "tz-3", "three.example.org")

	tests := []struct {
		name   string
		filter *trustzonepb.ListTrustZonesRequest_Filter
		want   []*trustzonepb.TrustZone
	}{
		{name: "no filter", want: []*trustzonepb.TrustZone{tz1, tz2, tz3}},
		{name: "empty filter", filter: &trustzonepb.ListTrustZonesRequest_Filter{}, want: []*trustzonepb.TrustZone{tz1, tz2, tz3}},
		{name: "one organization", filter: &trustzonepb.ListTrustZonesRequest_Filter{OrgIds: []string{orgID}}, want: []*trustzonepb.TrustZone{tz1, tz3}},
		{name: "two organizations", filter: &trustzonepb.ListTrustZonesRequest_Filter{OrgIds: []string{orgID, otherOrgID}}, want: []*trustzonepb.TrustZone{tz1, tz2, tz3}},
		{name: "unknown organization", filter: &trustzonepb.ListTrustZonesRequest_Filter{OrgIds: []string{"missing"}}, want: []*trustzonepb.TrustZone{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, gotPagination, err := client.ListTrustZones(t.Context(), tt.filter, pagination.Pagination{})
			require.NoError(t, err)
			assert.EqualExportedValues(t, tt.want, got)
			assert.Empty(t, gotPagination.Token)
		})
	}
}

func Test_ListTrustZones_Pagination(t *testing.T) {
	t.Parallel()
	fakeConnect, client := newFake(t)
	tz1 := seed(fakeConnect, "tz-1", orgID, "tz-1", "one.example.org")
	tz2 := seed(fakeConnect, "tz-2", orgID, "tz-2", "two.example.org")
	tz3 := seed(fakeConnect, "tz-3", orgID, "tz-3", "three.example.org")

	got, next, err := client.ListTrustZones(t.Context(), nil, pagination.Pagination{PageSize: 2})
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*trustzonepb.TrustZone{tz1, tz2}, got)
	assert.Equal(t, pagination.Pagination{PageSize: 2, Token: "tz-3"}, next)

	got, next, err = client.ListTrustZones(t.Context(), nil, next)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*trustzonepb.TrustZone{tz3}, got)
	assert.Empty(t, next.Token)
}

func Test_UpdateTrustZone(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		mask *trustzonepb.TrustZoneUpdateMask
	}{
		{name: "no mask", mask: nil},
		{name: "name in mask", mask: &trustzonepb.TrustZoneUpdateMask{Name: true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fakeConnect, client := newFake(t)
			seed(fakeConnect, "tz-1", orgID, "tz", "example.org")

			// Immutable and output only fields in the request are ignored.
			input := &trustzonepb.TrustZone{
				Id:                "tz-1",
				OrgId:             otherOrgID,
				Name:              "new-name",
				TrustDomain:       "changed.example.org",
				BundleEndpointUrl: "https://ignored.example.org",
				IsManagementZone:  true,
			}
			updated, err := client.UpdateTrustZone(t.Context(), input, tt.mask)
			require.NoError(t, err)

			assert.NotNil(t, updated.GetLastUpdatedAt())
			assert.EqualExportedValues(t, &trustzonepb.TrustZone{
				Id:            "tz-1",
				OrgId:         orgID,
				Name:          "new-name",
				TrustDomain:   "example.org",
				LastUpdatedAt: updated.GetLastUpdatedAt(),
			}, updated)
			assert.EqualExportedValues(t, updated, fakeConnect.TrustZones["tz-1"])
		})
	}
}

func Test_UpdateTrustZone_EmptyMask(t *testing.T) {
	t.Parallel()
	fakeConnect, client := newFake(t)
	want := seed(fakeConnect, "tz-1", orgID, "tz", "example.org")

	// A mask selecting no fields updates nothing, not even the last updated time, and so cannot fail validation.
	updated, err := client.UpdateTrustZone(t.Context(), &trustzonepb.TrustZone{Id: "tz-1"}, &trustzonepb.TrustZoneUpdateMask{})
	require.NoError(t, err)
	assert.EqualExportedValues(t, want, updated)
	assert.EqualExportedValues(t, want, fakeConnect.TrustZones["tz-1"])
}

func Test_UpdateTrustZone_Errors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		trustZone *trustzonepb.TrustZone
		wantCode  codes.Code
	}{
		{
			name:      "trust zone does not exist",
			trustZone: &trustzonepb.TrustZone{Id: "missing", Name: "tz"},
			wantCode:  codes.NotFound,
		},
		{
			name:      "empty name",
			trustZone: &trustzonepb.TrustZone{Id: "tz-1"},
			wantCode:  codes.InvalidArgument,
		},
		{
			name:      "name exists in organization",
			trustZone: &trustzonepb.TrustZone{Id: "tz-1", Name: "other"},
			wantCode:  codes.AlreadyExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fakeConnect, client := newFake(t)
			want := seed(fakeConnect, "tz-1", orgID, "tz", "one.example.org")
			seed(fakeConnect, "tz-2", orgID, "other", "two.example.org")

			_, err := client.UpdateTrustZone(t.Context(), tt.trustZone, nil)
			requireCode(t, err, tt.wantCode)
			assert.EqualExportedValues(t, want, fakeConnect.TrustZones["tz-1"])
		})
	}
}

func Test_UpdateTrustZone_KeepName(t *testing.T) {
	t.Parallel()
	fakeConnect, client := newFake(t)
	seed(fakeConnect, "tz-1", orgID, "tz", "example.org")

	// A trust zone's own name does not conflict with itself.
	_, err := client.UpdateTrustZone(t.Context(), &trustzonepb.TrustZone{Id: "tz-1", Name: "tz"}, nil)
	require.NoError(t, err)
}

func Test_PublishTrustZoneBundle(t *testing.T) {
	t.Parallel()
	fakeConnect, client := newFake(t)
	seed(fakeConnect, "tz-1", orgID, "tz", "example.org")

	err := client.PublishTrustZoneBundle(t.Context(), nil)
	requireCode(t, err, codes.InvalidArgument)

	err = client.PublishTrustZoneBundle(t.Context(), &trustbundlepb.Bundle{TrustDomain: "missing.example.org"})
	requireCode(t, err, codes.InvalidArgument)

	bundle := &trustbundlepb.Bundle{TrustDomain: "example.org", SequenceNumber: 1}
	err = client.PublishTrustZoneBundle(t.Context(), bundle)
	require.NoError(t, err)
	assert.EqualExportedValues(t, bundle, fakeConnect.TrustZoneBundles["tz-1"])
	assert.NotSame(t, bundle, fakeConnect.TrustZoneBundles["tz-1"])

	// Publishing replaces the stored bundle.
	bundle = &trustbundlepb.Bundle{TrustDomain: "example.org", SequenceNumber: 2}
	err = client.PublishTrustZoneBundle(t.Context(), bundle)
	require.NoError(t, err)
	assert.EqualExportedValues(t, bundle, fakeConnect.TrustZoneBundles["tz-1"])
}

// credential is both a CofideAgentCredential and a TrustZoneServerCredential, so that one set of cases can test
// both registration methods.
type credential interface {
	trustzonev1beta1.CofideAgentCredential
	trustzonev1beta1.TrustZoneServerCredential
}

func Test_Register(t *testing.T) {
	t.Parallel()

	methods := []struct {
		name     string
		register func(context.Context, trustzonev1beta1.TrustZoneClient, credential, *trustbundlepb.Bundle) error
	}{
		{
			name: "RegisterCofideAgent",
			register: func(ctx context.Context, client trustzonev1beta1.TrustZoneClient, credential credential, bundle *trustbundlepb.Bundle) error {
				return client.RegisterCofideAgent(ctx, credential, bundle)
			},
		},
		{
			name: "RegisterTrustZoneServer",
			register: func(ctx context.Context, client trustzonev1beta1.TrustZoneClient, credential credential, bundle *trustbundlepb.Bundle) error {
				return client.RegisterTrustZoneServer(ctx, credential, bundle)
			},
		},
	}
	tests := []struct {
		name        string
		credential  credential
		trustDomain string
		wantCode    codes.Code
	}{
		{name: "join token", credential: trustzonev1beta1.JoinToken("token-1"), trustDomain: "one.example.org"},
		{name: "k8s psat", credential: trustzonev1beta1.K8sPSAT("psat"), trustDomain: "one.example.org"},
		{name: "join token for other trust zone", credential: trustzonev1beta1.JoinToken("token-2"), trustDomain: "one.example.org", wantCode: codes.Unauthenticated},
		{name: "unknown join token", credential: trustzonev1beta1.JoinToken("unknown"), trustDomain: "one.example.org", wantCode: codes.Unauthenticated},
		{name: "empty join token", credential: trustzonev1beta1.JoinToken(""), trustDomain: "one.example.org", wantCode: codes.InvalidArgument},
		{name: "empty k8s psat", credential: trustzonev1beta1.K8sPSAT(""), trustDomain: "one.example.org", wantCode: codes.InvalidArgument},
		{name: "no credential", trustDomain: "one.example.org", wantCode: codes.InvalidArgument},
		{name: "unknown trust domain", credential: trustzonev1beta1.K8sPSAT("psat"), trustDomain: "missing.example.org", wantCode: codes.InvalidArgument},
	}
	for _, method := range methods {
		t.Run(method.name, func(t *testing.T) {
			t.Parallel()
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					t.Parallel()
					fakeConnect, client := newFake(t)
					seed(fakeConnect, "tz-1", orgID, "one", "one.example.org")
					seed(fakeConnect, "tz-2", orgID, "two", "two.example.org")
					fakeConnect.AgentJoinTokens["tz-1"] = map[string]string{"cluster-1": "token-1"}
					fakeConnect.AgentJoinTokens["tz-2"] = map[string]string{"cluster-2": "token-2"}

					bundle := &trustbundlepb.Bundle{TrustDomain: tt.trustDomain, SequenceNumber: 1}
					err := method.register(t.Context(), client, tt.credential, bundle)
					if tt.wantCode != codes.OK {
						requireCode(t, err, tt.wantCode)
						assert.Empty(t, fakeConnect.TrustZoneBundles)
						// A failed registration leaves every join token usable.
						assert.Equal(t, map[string]string{"cluster-1": "token-1"}, fakeConnect.AgentJoinTokens["tz-1"])
						assert.Equal(t, map[string]string{"cluster-2": "token-2"}, fakeConnect.AgentJoinTokens["tz-2"])
						return
					}
					require.NoError(t, err)
					assert.EqualExportedValues(t, bundle, fakeConnect.TrustZoneBundles["tz-1"])

					if _, ok := tt.credential.(trustzonev1beta1.JoinToken); ok {
						// Join tokens are single use.
						assert.Empty(t, fakeConnect.AgentJoinTokens["tz-1"])
						err = method.register(t.Context(), client, tt.credential, bundle)
						requireCode(t, err, codes.Unauthenticated)
					}
				})
			}
		})
	}

	t.Run("no bundle", func(t *testing.T) {
		t.Parallel()
		_, client := newFake(t)
		err := client.RegisterCofideAgent(t.Context(), trustzonev1beta1.K8sPSAT("psat"), nil)
		requireCode(t, err, codes.InvalidArgument)
		err = client.RegisterTrustZoneServer(t.Context(), trustzonev1beta1.K8sPSAT("psat"), nil)
		requireCode(t, err, codes.InvalidArgument)
	})
}
