// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"testing"
	"time"

	workloadsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/workload_service/v1alpha1"
	exchangepolicypb "github.com/cofide/cofide-api-sdk/gen/go/proto/exchange_policy/v1alpha1"
	identitypb "github.com/cofide/cofide-api-sdk/gen/go/proto/identity/v1alpha1"
	workloadpb "github.com/cofide/cofide-api-sdk/gen/go/proto/workload/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_fakeWorkloadClient_ListWorkloads(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	workload := test.FakeK8sPodWorkload()

	workloads, err := client.ListWorkloads(ctx, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*workloadpb.Workload{}, workloads)

	fake.Workloads[test.FakeWorkloadID] = test.FakeK8sPodWorkload()

	workloads, err = client.ListWorkloads(ctx, nil)
	require.NoError(t, err)
	assert.EqualExportedValues(t, []*workloadpb.Workload{workload}, workloads)
	assert.Empty(t, workloads[0].GetGrantedIdentities())
}

func Test_fakeWorkloadClient_ListWorkloads_grantedIdentities(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	fake.Workloads[test.FakeWorkloadID] = test.FakeK8sPodWorkload()
	fake.Identities[test.FakeIdentityID] = test.FakeIdentity()
	fake.Identities["other-identity-id"] = &identitypb.Identity{
		WorkloadId:          "other-workload-id",
		AttestationPolicyId: "other-ap-id",
		SpiffeId:            "spiffe://fake.trust.domain/other",
	}

	workloads, err := client.ListWorkloads(ctx, nil)
	require.NoError(t, err)
	require.Len(t, workloads, 1)

	assert.EqualExportedValues(t, []*workloadpb.GrantedIdentity{
		{
			AttestationPolicyId: test.FakeAttestationPolicyID,
			SpiffeId:            test.FakeSPIFFEID,
		},
	}, workloads[0].GetGrantedIdentities())
}

func Test_fakeWorkloadClient_ListWorkloads_matchingExchangePolicies(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	fake.Workloads[test.FakeWorkloadID] = test.FakeK8sPodWorkload()
	fake.Workloads[test.FakeLambdaWorkloadID] = test.FakeLambdaWorkload()
	fake.Workloads[test.FakeAgentCoreWorkloadID] = test.FakeAgentCoreWorkload()
	fake.Identities[test.FakeIdentityID] = test.FakeIdentity()

	fake.ExchangePolicies["subject-exact"] = &exchangepolicypb.ExchangePolicy{
		Id:              "subject-exact",
		SubjectIdentity: exactStringSet(test.FakeSPIFFEID),
	}
	fake.ExchangePolicies["subject-glob"] = &exchangepolicypb.ExchangePolicy{
		Id:              "subject-glob",
		SubjectIdentity: globStringSet("spiffe://fake.trust.domain/ns/*/sa/*"),
	}
	fake.ExchangePolicies["actor-exact"] = &exchangepolicypb.ExchangePolicy{
		Id:            "actor-exact",
		ActorIdentity: exactStringSet(test.FakeSPIFFEID),
	}
	fake.ExchangePolicies["subject-and-actor"] = &exchangepolicypb.ExchangePolicy{
		Id:              "subject-and-actor",
		SubjectIdentity: exactStringSet(test.FakeSPIFFEID),
		ActorIdentity:   exactStringSet(test.FakeSPIFFEID),
	}
	fake.ExchangePolicies["no-match"] = &exchangepolicypb.ExchangePolicy{
		Id:              "no-match",
		SubjectIdentity: exactStringSet("spiffe://fake.trust.domain/other"),
	}
	fake.ExchangePolicies["no-rule"] = &exchangepolicypb.ExchangePolicy{
		Id: "no-rule",
	}
	fake.ExchangePolicies["lambda-subject"] = &exchangepolicypb.ExchangePolicy{
		Id:              "lambda-subject",
		SubjectIdentity: exactStringSet(test.FakeLambdaIAMRoleARN),
	}
	fake.ExchangePolicies["agentcore-subject"] = &exchangepolicypb.ExchangePolicy{
		Id:              "agentcore-subject",
		SubjectIdentity: exactStringSet(test.FakeAgentCoreWorkloadIdentityARN),
	}
	fake.ExchangePolicies["agentcore-role-arn-not-matched"] = &exchangepolicypb.ExchangePolicy{
		Id:              "agentcore-role-arn-not-matched",
		SubjectIdentity: exactStringSet(test.FakeAgentCoreRoleARN),
	}

	workloads, err := client.ListWorkloads(ctx, nil)
	require.NoError(t, err)
	require.Len(t, workloads, 3)

	byID := make(map[string]*workloadpb.Workload, len(workloads))
	for _, w := range workloads {
		byID[w.GetId()] = w
	}

	k8sPod := byID[test.FakeWorkloadID]
	assert.ElementsMatch(t, []string{"subject-exact", "subject-glob", "subject-and-actor"}, k8sPod.GetMatchingSubjectExchangePolicyIds())
	assert.ElementsMatch(t, []string{"actor-exact", "subject-and-actor"}, k8sPod.GetMatchingActorExchangePolicyIds())

	lambda := byID[test.FakeLambdaWorkloadID]
	assert.ElementsMatch(t, []string{"lambda-subject"}, lambda.GetMatchingSubjectExchangePolicyIds())
	assert.Empty(t, lambda.GetMatchingActorExchangePolicyIds())

	agentCore := byID[test.FakeAgentCoreWorkloadID]
	assert.ElementsMatch(t, []string{"agentcore-subject"}, agentCore.GetMatchingSubjectExchangePolicyIds())
	assert.Empty(t, agentCore.GetMatchingActorExchangePolicyIds())
}

func exactStringSet(values ...string) *exchangepolicypb.StringSet {
	set := &exchangepolicypb.StringSet{}
	for _, value := range values {
		set.Matchers = append(set.Matchers, &exchangepolicypb.StringMatcher{
			Match: &exchangepolicypb.StringMatcher_Exact{Exact: value},
		})
	}
	return set
}

func globStringSet(patterns ...string) *exchangepolicypb.StringSet {
	set := &exchangepolicypb.StringSet{}
	for _, pattern := range patterns {
		set.Matchers = append(set.Matchers, &exchangepolicypb.StringMatcher{
			Match: &exchangepolicypb.StringMatcher_Glob{Glob: pattern},
		})
	}
	return set
}

func Test_fakeWorkloadClient_ListWorkloadEvents_filterObservedTimeRange(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()
	now := time.Now()

	recent := &workloadpb.WorkloadEvent{
		OrgId:             "org-1",
		ObservedTimestamp: timestamppb.New(now.Add(-5 * time.Minute)),
	}
	stale := &workloadpb.WorkloadEvent{
		OrgId:             "org-1",
		ObservedTimestamp: timestamppb.New(now.Add(-2 * time.Hour)),
	}
	fake.WorkloadEvents = []*workloadpb.WorkloadEvent{recent, stale}

	events, _, err := client.ListWorkloadEvents(ctx, &workloadsvcpb.ListWorkloadEventsRequest_Filter{
		OrgIds:         []string{"org-1"},
		ObservedAfter:  timestamppb.New(now.Add(-time.Hour)),
		ObservedBefore: timestamppb.New(now),
	}, pagination.Pagination{PageSize: 100})
	require.NoError(t, err)

	assert.EqualExportedValues(t, []*workloadpb.WorkloadEvent{recent}, events)
}

func Test_fakeWorkloadClient_ListWorkloadEvents_filterAgentSpiffeID(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	matching := &workloadpb.WorkloadEvent{
		AgentSpiffeId: "spiffe://example.org/agent/matching",
	}
	other := &workloadpb.WorkloadEvent{
		AgentSpiffeId: "spiffe://example.org/agent/other",
	}
	fake.WorkloadEvents = []*workloadpb.WorkloadEvent{matching, other}

	events, _, err := client.ListWorkloadEvents(ctx, &workloadsvcpb.ListWorkloadEventsRequest_Filter{
		AgentSpiffeIds: []string{"spiffe://example.org/agent/matching"},
	}, pagination.Pagination{PageSize: 100})
	require.NoError(t, err)

	assert.EqualExportedValues(t, []*workloadpb.WorkloadEvent{matching}, events)
}

func Test_fakeWorkloadClient_ListWorkloadEvents_filterWorkloadID(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	matching := &workloadpb.WorkloadEvent{
		WorkloadId: "workload-matching",
	}
	other := &workloadpb.WorkloadEvent{
		WorkloadId: "workload-other",
	}
	fake.WorkloadEvents = []*workloadpb.WorkloadEvent{matching, other}

	events, _, err := client.ListWorkloadEvents(ctx, &workloadsvcpb.ListWorkloadEventsRequest_Filter{
		WorkloadIds: []string{"workload-matching"},
	}, pagination.Pagination{PageSize: 100})
	require.NoError(t, err)

	assert.EqualExportedValues(t, []*workloadpb.WorkloadEvent{matching}, events)
}

func Test_fakeWorkloadClient_ListWorkloadEvents_filterEventTypes(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	workloadAttested := &workloadpb.WorkloadEvent{
		Event: &workloadpb.WorkloadEvent_WorkloadAttested{
			WorkloadAttested: &workloadpb.WorkloadAttestedEvent{},
		},
	}
	workloadAttestationFailed := &workloadpb.WorkloadEvent{
		Event: &workloadpb.WorkloadEvent_WorkloadAttestationFailed{
			WorkloadAttestationFailed: &workloadpb.WorkloadAttestationFailedEvent{
				Error: "context cancelled",
			},
		},
	}
	identityDelivered := &workloadpb.WorkloadEvent{
		Event: &workloadpb.WorkloadEvent_IdentityDelivered{
			IdentityDelivered: &workloadpb.IdentityDeliveredEvent{
				EntryId:  "entry-1",
				SpiffeId: "spiffe://example.org/workload",
			},
		},
	}
	noIdentity := &workloadpb.WorkloadEvent{
		Event: &workloadpb.WorkloadEvent_NoIdentity{
			NoIdentity: &workloadpb.NoIdentityEvent{Error: "no matching entries"},
		},
	}
	fake.WorkloadEvents = []*workloadpb.WorkloadEvent{workloadAttested, workloadAttestationFailed, identityDelivered, noIdentity}

	events, _, err := client.ListWorkloadEvents(ctx, &workloadsvcpb.ListWorkloadEventsRequest_Filter{
		EventTypes: []workloadpb.WorkloadEventType{
			workloadpb.WorkloadEventType_WORKLOAD_EVENT_TYPE_WORKLOAD_ATTESTATION_FAILED,
			workloadpb.WorkloadEventType_WORKLOAD_EVENT_TYPE_IDENTITY_DELIVERED,
			workloadpb.WorkloadEventType_WORKLOAD_EVENT_TYPE_NO_IDENTITY,
		},
	}, pagination.Pagination{PageSize: 100})
	require.NoError(t, err)

	assert.EqualExportedValues(t, []*workloadpb.WorkloadEvent{workloadAttestationFailed, identityDelivered, noIdentity}, events)
}

func Test_fakeWorkloadEventsStream_Send_skipsNilEvents(t *testing.T) {
	fake := fakeconnect.New()
	client := New(fake)
	ctx := context.Background()

	stream, err := client.PublishWorkloadEvents(ctx)
	require.NoError(t, err)

	event := &workloadpb.WorkloadEvent{OrgId: "org-1"}
	require.NoError(t, stream.Send([]*workloadpb.WorkloadEvent{event, nil}))

	assert.EqualExportedValues(t, []*workloadpb.WorkloadEvent{event}, fake.WorkloadEvents)
}
