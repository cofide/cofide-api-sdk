// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"
	"regexp"
	"slices"
	"sort"
	"strings"

	workloadsvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/workload_service/v1alpha1"
	exchangepolicypb "github.com/cofide/cofide-api-sdk/gen/go/proto/exchange_policy/v1alpha1"
	identitypb "github.com/cofide/cofide-api-sdk/gen/go/proto/identity/v1alpha1"
	workloadpb "github.com/cofide/cofide-api-sdk/gen/go/proto/workload/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
	workloadv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/workload/v1alpha1"
	"google.golang.org/protobuf/proto"
)

type fakeWorkloadClient struct {
	fake *fakeconnect.FakeConnect
}

// New instantiates a new WorkloadClient for communication with a fake Connect API.
func New(fake *fakeconnect.FakeConnect) workloadv1alpha1.WorkloadClient {
	return &fakeWorkloadClient{
		fake: fake,
	}
}

func (c *fakeWorkloadClient) ListWorkloads(ctx context.Context, filter *workloadsvcpb.ListWorkloadsRequest_Filter) ([]*workloadpb.Workload, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	workloads := []*workloadpb.Workload{}
	for _, workload := range c.fake.Workloads {
		if !workloadMatches(workload, filter) {
			continue
		}
		cloned := proto.Clone(workload).(*workloadpb.Workload)
		cloned.GrantedIdentities = grantedIdentitiesFor(c.fake.Identities, workload.GetId())
		identities := workloadIdentityStrings(cloned)
		cloned.MatchingSubjectExchangePolicyIds = matchingExchangePolicyIDs(c.fake.ExchangePolicies, identities, (*exchangepolicypb.ExchangePolicy).GetSubjectIdentity)
		cloned.MatchingActorExchangePolicyIds = matchingExchangePolicyIDs(c.fake.ExchangePolicies, identities, (*exchangepolicypb.ExchangePolicy).GetActorIdentity)
		workloads = append(workloads, cloned)
	}
	return workloads, nil
}

// grantedIdentitiesFor returns, for workloadID, the {attestation_policy_id, spiffe_id}
// pair for every identity issued to that workload, sorted by attestation policy ID
// (then SPIFFE ID) for deterministic output — the identities map has random
// iteration order.
func grantedIdentitiesFor(identities map[string]*identitypb.Identity, workloadID string) []*workloadpb.GrantedIdentity {
	var granted []*workloadpb.GrantedIdentity
	for _, identity := range identities {
		if identity.GetWorkloadId() != workloadID {
			continue
		}
		granted = append(granted, &workloadpb.GrantedIdentity{
			AttestationPolicyId: identity.GetAttestationPolicyId(),
			SpiffeId:            identity.GetSpiffeId(),
		})
	}
	sort.Slice(granted, func(i, j int) bool {
		if granted[i].GetAttestationPolicyId() != granted[j].GetAttestationPolicyId() {
			return granted[i].GetAttestationPolicyId() < granted[j].GetAttestationPolicyId()
		}
		return granted[i].GetSpiffeId() < granted[j].GetSpiffeId()
	})
	return granted
}

// workloadIdentityStrings returns the identity string(s) that a workload
// would present as either subject or actor in a token exchange, depending on
// its type: the SPIFFE IDs granted to it via attestation policies for
// Kubernetes pods, or the relevant ARN for AWS-native workloads which have no
// SPIRE attestation flow.
func workloadIdentityStrings(workload *workloadpb.Workload) []string {
	switch w := workload.GetWorkload().(type) {
	case *workloadpb.Workload_KubernetesPod:
		var identities []string
		for _, granted := range workload.GetGrantedIdentities() {
			identities = append(identities, granted.GetSpiffeId())
		}
		return identities
	case *workloadpb.Workload_LambdaFunction:
		return []string{w.LambdaFunction.GetIamRoleArn()}
	case *workloadpb.Workload_AgentcoreWorkload:
		return []string{w.AgentcoreWorkload.GetWorkloadIdentityArn()}
	default:
		return nil
	}
}

// matchingExchangePolicyIDs returns the IDs of policies, sorted for
// deterministic output, whose rule selected by ruleFor (e.g. subject_identity
// or actor_identity) is satisfied by any of identities. A policy whose rule is
// unset (no matchers) has no constraint for that role and is never included.
func matchingExchangePolicyIDs(policies map[string]*exchangepolicypb.ExchangePolicy, identities []string, ruleFor func(*exchangepolicypb.ExchangePolicy) *exchangepolicypb.StringSet) []string {
	var matched []string
	for id, policy := range policies {
		rule := ruleFor(policy)
		for _, identity := range identities {
			if stringSetMatches(rule, identity) {
				matched = append(matched, id)
				break
			}
		}
	}
	sort.Strings(matched)
	return matched
}

// stringSetMatches reports whether value satisfies any matcher in set. An
// unset or empty set has no matchers and never matches.
func stringSetMatches(set *exchangepolicypb.StringSet, value string) bool {
	for _, matcher := range set.GetMatchers() {
		switch m := matcher.GetMatch().(type) {
		case *exchangepolicypb.StringMatcher_Exact:
			if m.Exact == value {
				return true
			}
		case *exchangepolicypb.StringMatcher_Glob:
			if globMatch(m.Glob, value) {
				return true
			}
		}
	}
	return false
}

// globMatch reports whether value matches pattern, where '*' matches any
// sequence of characters (including none).
func globMatch(pattern, value string) bool {
	parts := strings.Split(pattern, "*")
	for i, part := range parts {
		parts[i] = regexp.QuoteMeta(part)
	}
	re, err := regexp.Compile("^" + strings.Join(parts, ".*") + "$")
	if err != nil {
		return false
	}
	return re.MatchString(value)
}

func (c *fakeWorkloadClient) ListWorkloadEvents(ctx context.Context, filter *workloadsvcpb.ListWorkloadEventsRequest_Filter, requestPagination pagination.Pagination) ([]*workloadpb.WorkloadEvent, pagination.Pagination, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	events := []*workloadpb.WorkloadEvent{}
	for _, event := range c.fake.WorkloadEvents {
		if workloadEventMatches(event, filter) {
			events = append(events, proto.Clone(event).(*workloadpb.WorkloadEvent))
		}
	}
	return events, pagination.Pagination{PageSize: requestPagination.PageSize}, nil
}

func (c *fakeWorkloadClient) PublishWorkloadEvents(ctx context.Context) (workloadv1alpha1.WorkloadEventsStream, error) {
	return &fakeWorkloadEventsStream{fake: c.fake}, nil
}

type fakeWorkloadEventsStream struct {
	fake *fakeconnect.FakeConnect
}

func (s *fakeWorkloadEventsStream) Send(events []*workloadpb.WorkloadEvent) error {
	s.fake.Mu.Lock()
	defer s.fake.Mu.Unlock()

	for _, event := range events {
		if event == nil {
			continue
		}
		s.fake.WorkloadEvents = append(s.fake.WorkloadEvents, proto.Clone(event).(*workloadpb.WorkloadEvent))
	}
	return nil
}

func (s *fakeWorkloadEventsStream) Close() error {
	return nil
}

func workloadMatches(workload *workloadpb.Workload, filter *workloadsvcpb.ListWorkloadsRequest_Filter) bool {
	if filter == nil {
		return true
	}
	if filter.OrgId != nil && workload.GetOrgId() != *filter.OrgId {
		return false
	}
	if filter.TrustZoneId != nil && workload.GetTrustZoneId() != *filter.TrustZoneId {
		return false
	}
	if filter.ClusterId != nil && workload.GetClusterId() != *filter.ClusterId {
		return false
	}
	return true
}

func workloadEventMatches(event *workloadpb.WorkloadEvent, filter *workloadsvcpb.ListWorkloadEventsRequest_Filter) bool {
	if filter == nil {
		return true
	}
	if len(filter.GetOrgIds()) > 0 && !slices.Contains(filter.GetOrgIds(), event.GetOrgId()) {
		return false
	}
	if len(filter.GetTrustZoneIds()) > 0 && !slices.Contains(filter.GetTrustZoneIds(), event.GetTrustZoneId()) {
		return false
	}
	if len(filter.GetClusterIds()) > 0 && !slices.Contains(filter.GetClusterIds(), event.GetClusterId()) {
		return false
	}
	if len(filter.GetAgentSpiffeIds()) > 0 && !slices.Contains(filter.GetAgentSpiffeIds(), event.GetAgentSpiffeId()) {
		return false
	}
	if len(filter.GetWorkloadIds()) > 0 && !slices.Contains(filter.GetWorkloadIds(), event.GetWorkloadId()) {
		return false
	}
	if len(filter.GetEventTypes()) > 0 && !workloadEventTypeMatches(event, filter.GetEventTypes()) {
		return false
	}
	observedBefore := filter.GetObservedBefore()
	observedAfter := filter.GetObservedAfter()
	hasObservedBefore := observedBefore.IsValid()
	hasObservedAfter := observedAfter.IsValid()
	if hasObservedBefore || hasObservedAfter {
		observed := event.GetObservedTimestamp()
		if !observed.IsValid() {
			return false
		}
		if hasObservedBefore && observed.AsTime().After(observedBefore.AsTime()) {
			return false
		}
		if hasObservedAfter && observed.AsTime().Before(observedAfter.AsTime()) {
			return false
		}
	}
	if len(filter.GetSpiffeIds()) > 0 {
		if delivered := event.GetIdentityDelivered(); delivered == nil || !slices.Contains(filter.GetSpiffeIds(), delivered.GetSpiffeId()) {
			return false
		}
	}
	if len(filter.GetEntryIds()) > 0 {
		if delivered := event.GetIdentityDelivered(); delivered == nil || !slices.Contains(filter.GetEntryIds(), delivered.GetEntryId()) {
			return false
		}
	}
	return true
}

func workloadEventTypeMatches(event *workloadpb.WorkloadEvent, eventTypes []workloadpb.WorkloadEventType) bool {
	actual := workloadEventType(event)
	for _, eventType := range eventTypes {
		if eventType == actual {
			return true
		}
	}
	return false
}

func workloadEventType(event *workloadpb.WorkloadEvent) workloadpb.WorkloadEventType {
	switch event.GetEvent().(type) {
	case *workloadpb.WorkloadEvent_WorkloadAttested:
		return workloadpb.WorkloadEventType_WORKLOAD_EVENT_TYPE_WORKLOAD_ATTESTED
	case *workloadpb.WorkloadEvent_WorkloadAttestationFailed:
		return workloadpb.WorkloadEventType_WORKLOAD_EVENT_TYPE_WORKLOAD_ATTESTATION_FAILED
	case *workloadpb.WorkloadEvent_IdentityDelivered:
		return workloadpb.WorkloadEventType_WORKLOAD_EVENT_TYPE_IDENTITY_DELIVERED
	case *workloadpb.WorkloadEvent_NoIdentity:
		return workloadpb.WorkloadEventType_WORKLOAD_EVENT_TYPE_NO_IDENTITY
	default:
		return workloadpb.WorkloadEventType_WORKLOAD_EVENT_TYPE_UNSPECIFIED
	}
}
