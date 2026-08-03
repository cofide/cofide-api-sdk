// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package test

import (
	agentpb "github.com/cofide/cofide-api-sdk/gen/go/proto/agent/v1alpha1"
	apbindingpb "github.com/cofide/cofide-api-sdk/gen/go/proto/ap_binding/v1alpha1"
	attestationpolicypb "github.com/cofide/cofide-api-sdk/gen/go/proto/attestation_policy/v1alpha1"
	cloudaccountpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_account/v1alpha1"
	cloudorganizationpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_organization/v1alpha1"
	cloudproviderpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_provider/v1alpha1"
	clusterpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cluster/v1alpha1"
	federatedservicepb "github.com/cofide/cofide-api-sdk/gen/go/proto/federated_service/v1alpha1"
	identitypb "github.com/cofide/cofide-api-sdk/gen/go/proto/identity/v1alpha1"
	organizationpb "github.com/cofide/cofide-api-sdk/gen/go/proto/organization/v1alpha1"
	rolebindingpb "github.com/cofide/cofide-api-sdk/gen/go/proto/role_binding/v1alpha1"
	trustzonepb "github.com/cofide/cofide-api-sdk/gen/go/proto/trust_zone/v1alpha1"
	trustzoneserverpb "github.com/cofide/cofide-api-sdk/gen/go/proto/trust_zone_server/v1alpha1"
	workloadpb "github.com/cofide/cofide-api-sdk/gen/go/proto/workload/v1alpha1"
	workloadsuppressionrulepb "github.com/cofide/cofide-api-sdk/gen/go/proto/workload_suppression_rule/v1alpha1"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
)

const (
	FakeOrganizationID   = "fake-org-id"
	FakeOrganizationName = "fake-org-name"

	FakeTrustZoneID   = "fake-tz-id"
	FakeTrustZoneName = "fake-tz-name"
	FakeTrustDomain   = "fake.trust.domain"

	FakeTrustZoneServerID        = "fake-tzs-id"
	FakeKubernetesNamespace      = "fake-kubernetes-namespace"
	FakeKubernetesServiceAccount = "fake-kubernetes-service-account"

	FakeClusterID   = "fake-cluster-id"
	FakeClusterName = "fake-cluster-name"

	FakeCloudOrganizationID   = "fake-cloud-org-id"
	FakeCloudOrganizationName = "fake-cloud-org-name"
	FakeAWSOrgID              = "o-fakeorgid12"
	FakeIAMRoleARN            = "arn:aws:iam::123456789012:role/fake-role"

	FakeCloudAccountID   = "fake-cloud-account-id"
	FakeCloudAccountName = "fake-cloud-account-name"
	FakeAWSAccountID     = "123456789012"

	FakeAgentToken = "fake-agent-token"
	FakeAgentID    = "fake-agent-id"

	FakeFSID   = "fake-fs-id"
	FakeFSName = "fake-fs-name"

	FakeAttestationPolicyID   = "fake-ap-id"
	FakeAttestationPolicyName = "fake-ap-name"

	FakeAPBindingID = "fake-ap-binding-id"

	FakeWorkloadID      = "fake-workload-id"
	FakeK8sPodUID       = "fake-k8s-pod-uid"
	FakeK8sPodName      = "fake-k8s-pod-name"
	FakeK8sPodNamespace = "fake-k8s-pod-namespace"

	FakeLambdaWorkloadID   = "fake-lambda-workload-id"
	FakeLambdaFunctionName = "fake-lambda-function-name"
	FakeLambdaIAMRoleARN   = "arn:aws:iam::123456789012:role/fake-lambda-role"

	FakeAgentCoreWorkloadID          = "fake-agentcore-workload-id"
	FakeAgentCoreRuntimeName         = "fake-agentcore-runtime-name"
	FakeAgentCoreRoleARN             = "arn:aws:iam::123456789012:role/fake-agentcore-role"
	FakeAgentCoreWorkloadIdentityARN = "arn:aws:bedrock-agentcore:us-east-1:123456789012:workload-identity/fake-agentcore-runtime"

	FakeIdentityID    = "fake-identity-id"
	FakeSPIFFEID      = "spiffe://fake.trust.domain/ns/fake-k8s-pod-namespace/sa/fake-k8s-pod-service-account"
	FakeParentID      = "spiffe://fake.trust.domain/spire/agent/k8s_psat/fake-cluster-name/fake-spire-agent"
	FakeSelectorType  = "fake-selector-type"
	FakeSelectorValue = "fake-selector-value"
	FakeDNSName       = "fake.name.example.org"

	FakeRoleBindingID = "fake-rb-id"
	FakeRoleID        = "fake-role-id"
	FakeUserSubject   = "fake-user-subject"
	FakeResourceID    = FakeAPBindingID
	FakeResourceType  = "AttestationPolicyBinding"

	FakeWorkloadSuppressionRuleID          = "fake-wsr-id"
	FakeWorkloadSuppressionRuleName        = "fake-wsr-name"
	FakeWorkloadSuppressionRuleDescription = "fake-wsr-description"
)

func FakeOrganization() *organizationpb.Organization {
	return &organizationpb.Organization{
		Id:   FakeOrganizationID,
		Name: FakeOrganizationName,
	}
}

func FakeTrustZone() *trustzonepb.TrustZone {
	return &trustzonepb.TrustZone{
		Id:          PtrOf(FakeTrustZoneID),
		Name:        FakeTrustZoneName,
		TrustDomain: FakeTrustDomain,
	}
}

func FakeTrustZoneServer() *trustzoneserverpb.TrustZoneServer {
	return &trustzoneserverpb.TrustZoneServer{
		Id:                       FakeTrustZoneServerID,
		TrustZoneId:              FakeTrustZoneID,
		ClusterId:                FakeClusterID,
		KubernetesNamespace:      FakeKubernetesNamespace,
		KubernetesServiceAccount: FakeKubernetesServiceAccount,
	}
}

func FakeCluster() *clusterpb.Cluster {
	return &clusterpb.Cluster{
		Id:          PtrOf(FakeClusterID),
		Name:        PtrOf(FakeClusterName),
		TrustZoneId: PtrOf(FakeTrustZoneID),
	}
}

func FakeCloudOrganization() *cloudorganizationpb.CloudOrganization {
	return &cloudorganizationpb.CloudOrganization{
		Id:    FakeCloudOrganizationID,
		OrgId: FakeOrganizationID,
		Name:  FakeCloudOrganizationName,
		Provider: &cloudorganizationpb.CloudOrganization_Aws{
			Aws: &cloudorganizationpb.AWSOrganization{
				AwsOrgId: FakeAWSOrgID,
				Audience: "fake-audience",
				RoleChain: []*cloudproviderpb.AWSAssumeRoleConfig{
					{IamRoleArn: FakeIAMRoleARN},
				},
			},
		},
	}
}

func FakeCloudAccount() *cloudaccountpb.CloudAccount {
	return &cloudaccountpb.CloudAccount{
		Id:                  FakeCloudAccountID,
		OrgId:               FakeOrganizationID,
		CloudOrganizationId: PtrOf(FakeCloudOrganizationID),
		Name:                FakeCloudAccountName,
		Provider: &cloudaccountpb.CloudAccount_Aws{
			Aws: &cloudaccountpb.AWSAccount{
				AccountId: FakeAWSAccountID,
			},
		},
	}
}

func FakeAgent() *agentpb.Agent {
	return &agentpb.Agent{
		Id:          PtrOf(FakeAgentID),
		ClusterId:   PtrOf(FakeClusterID),
		TrustZoneId: PtrOf(FakeTrustZoneID),
	}
}

func FakeFederatedService() *federatedservicepb.FederatedService {
	return &federatedservicepb.FederatedService{
		Id:   FakeFSID,
		Name: FakeFSName,
	}
}

func FakeBundle() *types.Bundle {
	return &types.Bundle{TrustDomain: FakeTrustDomain}
}

func FakeAttestationPolicy() *attestationpolicypb.AttestationPolicy {
	return &attestationpolicypb.AttestationPolicy{
		Id:   PtrOf(FakeAttestationPolicyID),
		Name: FakeAttestationPolicyName,
	}
}

func FakeAPBinding() *apbindingpb.APBinding {
	return &apbindingpb.APBinding{
		Id:          PtrOf(FakeAPBindingID),
		TrustZoneId: PtrOf(FakeTrustZoneID),
		PolicyId:    PtrOf(FakeAttestationPolicyID),
	}
}

func FakeK8sPodWorkload() *workloadpb.Workload {
	return &workloadpb.Workload{
		Id:          FakeWorkloadID,
		TrustZoneId: FakeTrustZoneID,
		ClusterId:   FakeClusterID,
		Type:        workloadpb.WorkloadType_WORKLOAD_TYPE_KUBERNETES_POD,
		Workload: &workloadpb.Workload_KubernetesPod{
			KubernetesPod: &workloadpb.KubernetesPod{
				Metadata: &workloadpb.KubernetesMetadata{
					Uid:       FakeK8sPodUID,
					Name:      FakeK8sPodName,
					Namespace: FakeK8sPodNamespace,
				},
			},
		},
	}
}

func FakeLambdaWorkload() *workloadpb.Workload {
	return &workloadpb.Workload{
		Id:          FakeLambdaWorkloadID,
		TrustZoneId: FakeTrustZoneID,
		ClusterId:   FakeClusterID,
		Type:        workloadpb.WorkloadType_WORKLOAD_TYPE_AWS_LAMBDA_FUNCTION,
		Workload: &workloadpb.Workload_LambdaFunction{
			LambdaFunction: &workloadpb.AWSLambdaFunction{
				FunctionName:   FakeLambdaFunctionName,
				IamRoleArn:     FakeLambdaIAMRoleARN,
				CloudAccountId: FakeCloudAccountID,
			},
		},
	}
}

func FakeAgentCoreWorkload() *workloadpb.Workload {
	return &workloadpb.Workload{
		Id:          FakeAgentCoreWorkloadID,
		TrustZoneId: FakeTrustZoneID,
		ClusterId:   FakeClusterID,
		Type:        workloadpb.WorkloadType_WORKLOAD_TYPE_AWS_AGENTCORE_RUNTIME,
		Workload: &workloadpb.Workload_AgentcoreWorkload{
			AgentcoreWorkload: &workloadpb.AWSAgentCoreRuntime{
				AgentRuntimeName:    FakeAgentCoreRuntimeName,
				RoleArn:             FakeAgentCoreRoleARN,
				WorkloadIdentityArn: FakeAgentCoreWorkloadIdentityARN,
				CloudAccountId:      FakeCloudAccountID,
			},
		},
	}
}

func FakeIdentity() *identitypb.Identity {
	return &identitypb.Identity{
		Id:                  FakeIdentityID,
		TrustZoneId:         FakeTrustZoneID,
		ClusterId:           FakeClusterID,
		WorkloadId:          FakeWorkloadID,
		AttestationPolicyId: FakeAttestationPolicyID,
		SpiffeId:            FakeSPIFFEID,
		ParentId:            FakeParentID,
		Selectors: []*identitypb.Selector{
			{
				Type:  FakeSelectorType,
				Value: FakeSelectorValue,
			},
		},
		DnsNames: []string{
			FakeDNSName,
		},
		Federations: []*identitypb.IdentityFederation{
			{
				Federation: &identitypb.IdentityFederation_TrustZoneId{
					TrustZoneId: FakeTrustZoneID,
				},
			},
		},
	}
}

func FakeWorkloadSuppressionRule() *workloadsuppressionrulepb.WorkloadSuppressionRule {
	return &workloadsuppressionrulepb.WorkloadSuppressionRule{
		Id:          FakeWorkloadSuppressionRuleID,
		OrgId:       FakeOrganizationID,
		Name:        FakeWorkloadSuppressionRuleName,
		Description: FakeWorkloadSuppressionRuleDescription,
		Enabled:     true,
		Matcher: &workloadsuppressionrulepb.WorkloadSuppressionRule_KubernetesPod{
			KubernetesPod: &workloadsuppressionrulepb.KubernetesPodMatcher{
				TrustZoneIds: []string{FakeTrustZoneID},
				ClusterIds:   []string{FakeClusterID},
				Namespaces:   []string{FakeK8sPodNamespace},
			},
		},
	}
}

func FakeWorkloadSuppressionRuleLambda() *workloadsuppressionrulepb.WorkloadSuppressionRule {
	return &workloadsuppressionrulepb.WorkloadSuppressionRule{
		Id:          FakeWorkloadSuppressionRuleID,
		OrgId:       FakeOrganizationID,
		Name:        FakeWorkloadSuppressionRuleName,
		Description: FakeWorkloadSuppressionRuleDescription,
		Enabled:     true,
		Matcher: &workloadsuppressionrulepb.WorkloadSuppressionRule_AwsLambdaFunction{
			AwsLambdaFunction: &workloadsuppressionrulepb.AWSLambdaFunctionMatcher{
				CloudAccountIds: []string{FakeCloudAccountID},
			},
		},
	}
}

func FakeWorkloadSuppressionRuleAgentCore() *workloadsuppressionrulepb.WorkloadSuppressionRule {
	return &workloadsuppressionrulepb.WorkloadSuppressionRule{
		Id:          FakeWorkloadSuppressionRuleID,
		OrgId:       FakeOrganizationID,
		Name:        FakeWorkloadSuppressionRuleName,
		Description: FakeWorkloadSuppressionRuleDescription,
		Enabled:     true,
		Matcher: &workloadsuppressionrulepb.WorkloadSuppressionRule_AwsAgentcoreRuntime{
			AwsAgentcoreRuntime: &workloadsuppressionrulepb.AWSAgentCoreRuntimeMatcher{
				CloudAccountIds: []string{FakeCloudAccountID},
			},
		},
	}
}

func FakeRoleBinding() *rolebindingpb.RoleBinding {
	return &rolebindingpb.RoleBinding{
		Id:     FakeRoleBindingID,
		RoleId: FakeRoleID,
		Principal: &rolebindingpb.RoleBinding_User{
			User: &rolebindingpb.User{
				Subject: FakeUserSubject,
			},
		},
		Resource: &rolebindingpb.Resource{
			Id:   FakeResourceID,
			Type: FakeResourceType,
		},
	}
}
