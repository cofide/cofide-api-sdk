// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1_test

import (
	"context"
	"testing"

	crdcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_discovery_config/v1alpha1"
	"github.com/cofide/cofide-api-sdk/gen/go/proto/common/pagination/v1beta1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/clouddiscoveryconfig/v1alpha1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
)

// Test_CloudDiscoveryConfigClient_Unimplemented tests v1alpha1.CloudDiscoveryConfigClient against an unimplemented server.
// This ensures that all errors returned are not wrapped can be converted to a gRPC Status using Status.Convert.
func Test_CloudDiscoveryConfigClient_Unimplemented(t *testing.T) {
	server := test.NewTestServer(t)
	crdcpb.RegisterCloudDiscoveryConfigServiceServer(server.Server, &crdcpb.UnimplementedCloudDiscoveryConfigServiceServer{})
	server.Serve()

	conn := server.CreateClientConn()
	client := v1alpha1.New(conn)

	_, err := client.CreateCloudDiscoveryConfig(t.Context(), nil)
	test.RequireUnimplemented(t, err)

	err = client.DeleteCloudDiscoveryConfig(t.Context(), "")
	test.RequireUnimplemented(t, err)

	_, err = client.GetCloudDiscoveryConfig(t.Context(), "")
	test.RequireUnimplemented(t, err)

	_, _, err = client.ListCloudDiscoveryConfigs(t.Context(), nil, pagination.Pagination{})
	test.RequireUnimplemented(t, err)

	_, err = client.UpdateCloudDiscoveryConfig(t.Context(), nil, nil)
	test.RequireUnimplemented(t, err)
}

func Test_CreateCloudDiscoveryConfig(t *testing.T) {
	client, fakeService := clientForTest(t)

	input := lambdaConfig()
	fakeService.expectedCreateRequest = &crdcpb.CreateCloudDiscoveryConfigRequest{
		CloudDiscoveryConfig: input,
	}

	_, err := client.CreateCloudDiscoveryConfig(t.Context(), input)
	require.NoError(t, err)
}

func Test_DeleteCloudDiscoveryConfig(t *testing.T) {
	client, fakeService := clientForTest(t)

	input := "test-config-id"
	fakeService.expectedDeleteRequest = &crdcpb.DeleteCloudDiscoveryConfigRequest{
		CloudDiscoveryConfigId: input,
	}

	err := client.DeleteCloudDiscoveryConfig(t.Context(), input)
	require.NoError(t, err)
}

func Test_GetCloudDiscoveryConfig(t *testing.T) {
	client, fakeService := clientForTest(t)

	input := "test-config-id"
	fakeService.expectedGetRequest = &crdcpb.GetCloudDiscoveryConfigRequest{
		CloudDiscoveryConfigId: input,
	}

	_, err := client.GetCloudDiscoveryConfig(t.Context(), input)
	require.NoError(t, err)
}

func Test_ListCloudDiscoveryConfigs(t *testing.T) {
	client, fakeService := clientForTest(t)

	filter := &crdcpb.ListCloudDiscoveryConfigsRequest_Filter{}
	pagination := pagination.Pagination{PageSize: 100}
	fakeService.expectedListRequest = &crdcpb.ListCloudDiscoveryConfigsRequest{
		Filter:     filter,
		Pagination: &v1beta1.PageRequest{PageSize: pagination.PageSize},
	}

	_, _, err := client.ListCloudDiscoveryConfigs(t.Context(), filter, pagination)
	require.NoError(t, err)
}

func Test_UpdateCloudDiscoveryConfig_AWSLambda(t *testing.T) {
	client, fakeService := clientForTest(t)

	input := lambdaConfig()
	input.Id = "test-id"
	fakeService.expectedUpdateRequest = &crdcpb.UpdateCloudDiscoveryConfigRequest{
		CloudDiscoveryConfig: input,
		UpdateMask: &crdcpb.UpdateCloudDiscoveryConfigRequest_AwsLambdaDiscoveryConfigUpdateMask{
			AwsLambdaDiscoveryConfigUpdateMask: &crdcpb.UpdateCloudDiscoveryConfigRequest_AWSLambdaDiscoveryConfigUpdateMask{
				Name:     true,
				Regions:  true,
				Audience: true,
			},
		},
	}

	_, err := client.UpdateCloudDiscoveryConfig(t.Context(), input, v1alpha1.AWSLambdaUpdateMask(v1alpha1.AWSLambdaUpdateFields{
		Name:     true,
		Regions:  true,
		Audience: true,
	}))
	require.NoError(t, err)
}

func Test_UpdateCloudDiscoveryConfig_AWSAgentCore(t *testing.T) {
	client, fakeService := clientForTest(t)

	input := agentCoreConfig()
	input.Id = "test-id"
	fakeService.expectedUpdateRequest = &crdcpb.UpdateCloudDiscoveryConfigRequest{
		CloudDiscoveryConfig: input,
		UpdateMask: &crdcpb.UpdateCloudDiscoveryConfigRequest_AwsAgentCoreDiscoveryConfigUpdateMask{
			AwsAgentCoreDiscoveryConfigUpdateMask: &crdcpb.UpdateCloudDiscoveryConfigRequest_AWSAgentCoreDiscoveryConfigUpdateMask{
				Disabled: true,
				Roles:    true,
			},
		},
	}

	_, err := client.UpdateCloudDiscoveryConfig(t.Context(), input, v1alpha1.AWSAgentCoreUpdateMask(v1alpha1.AWSAgentCoreUpdateFields{
		Disabled: true,
		Roles:    true,
	}))
	require.NoError(t, err)
}

// A nil mask must not be sent as an empty mask, which the server would treat as "update nothing".
func Test_UpdateCloudDiscoveryConfig_NilUpdateMask(t *testing.T) {
	client, fakeService := clientForTest(t)

	input := lambdaConfig()
	input.Id = "test-id"
	fakeService.expectedUpdateRequest = &crdcpb.UpdateCloudDiscoveryConfigRequest{CloudDiscoveryConfig: input}

	_, err := client.UpdateCloudDiscoveryConfig(t.Context(), input, nil)
	require.NoError(t, err)
}

func lambdaConfig() *crdcpb.CloudDiscoveryConfig {
	return &crdcpb.CloudDiscoveryConfig{
		CloudAccountId:    "test-cloud-account-id",
		Name:              "test-name",
		DiscoveryInterval: durationpb.New(60_000_000_000),
		DiscoveryConfig: &crdcpb.CloudDiscoveryConfig_AwsLambdaDiscoveryConfig{
			AwsLambdaDiscoveryConfig: &crdcpb.AWSLambdaDiscoveryConfig{
				Regions: []string{"eu-west-1"},
				RoleChain: &crdcpb.AWSAssumeRoleConfig{
					AssumeThroughOidc: true,
					Audience:          "test-audience",
					Roles:             []*crdcpb.AWSAssumeRoleConfig_Role{{IamRoleArn: "arn:aws:iam::123456789012:role/test"}},
				},
			},
		},
	}
}

func agentCoreConfig() *crdcpb.CloudDiscoveryConfig {
	return &crdcpb.CloudDiscoveryConfig{
		CloudAccountId: "test-cloud-account-id",
		Name:           "test-name",
		DiscoveryConfig: &crdcpb.CloudDiscoveryConfig_AwsAgentCoreDiscoveryConfig{
			AwsAgentCoreDiscoveryConfig: &crdcpb.AWSAgentCoreDiscoveryConfig{
				Regions: []string{"us-east-1"},
				RoleChain: &crdcpb.AWSAssumeRoleConfig{
					Roles: []*crdcpb.AWSAssumeRoleConfig_Role{{IamRoleArn: "arn:aws:iam::123456789012:role/test", ExternalId: "ext"}},
				},
			},
		},
	}
}

func clientForTest(t *testing.T) (v1alpha1.CloudDiscoveryConfigClient, *fakeCloudDiscoveryConfigService) {
	t.Helper()
	server := test.NewTestServer(t)
	fakeService := &fakeCloudDiscoveryConfigService{t: t}
	crdcpb.RegisterCloudDiscoveryConfigServiceServer(server.Server, fakeService)
	server.Serve()

	conn := server.CreateClientConn()
	client := v1alpha1.New(conn)

	return client, fakeService
}

// fakeCloudDiscoveryConfigService provides a very simple fake CloudDiscoveryConfigService implementation with canned responses.
// Its purpose is to assert that the inbound request is as expected.
type fakeCloudDiscoveryConfigService struct {
	t                     *testing.T
	expectedCreateRequest *crdcpb.CreateCloudDiscoveryConfigRequest
	expectedDeleteRequest *crdcpb.DeleteCloudDiscoveryConfigRequest
	expectedGetRequest    *crdcpb.GetCloudDiscoveryConfigRequest
	expectedListRequest   *crdcpb.ListCloudDiscoveryConfigsRequest
	expectedUpdateRequest *crdcpb.UpdateCloudDiscoveryConfigRequest
}

func (f *fakeCloudDiscoveryConfigService) CreateCloudDiscoveryConfig(ctx context.Context, req *crdcpb.CreateCloudDiscoveryConfigRequest) (*crdcpb.CreateCloudDiscoveryConfigResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedCreateRequest, req)
	return &crdcpb.CreateCloudDiscoveryConfigResponse{CloudDiscoveryConfig: req.GetCloudDiscoveryConfig()}, nil
}

func (f *fakeCloudDiscoveryConfigService) DeleteCloudDiscoveryConfig(ctx context.Context, req *crdcpb.DeleteCloudDiscoveryConfigRequest) (*crdcpb.DeleteCloudDiscoveryConfigResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedDeleteRequest, req)
	return &crdcpb.DeleteCloudDiscoveryConfigResponse{}, nil
}

func (f *fakeCloudDiscoveryConfigService) GetCloudDiscoveryConfig(ctx context.Context, req *crdcpb.GetCloudDiscoveryConfigRequest) (*crdcpb.GetCloudDiscoveryConfigResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedGetRequest, req)
	return &crdcpb.GetCloudDiscoveryConfigResponse{CloudDiscoveryConfig: &crdcpb.CloudDiscoveryConfig{Id: req.GetCloudDiscoveryConfigId()}}, nil
}

func (f *fakeCloudDiscoveryConfigService) ListCloudDiscoveryConfigs(ctx context.Context, req *crdcpb.ListCloudDiscoveryConfigsRequest) (*crdcpb.ListCloudDiscoveryConfigsResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedListRequest, req)
	return &crdcpb.ListCloudDiscoveryConfigsResponse{CloudDiscoveryConfigs: []*crdcpb.CloudDiscoveryConfig{}}, nil
}

func (f *fakeCloudDiscoveryConfigService) UpdateCloudDiscoveryConfig(ctx context.Context, req *crdcpb.UpdateCloudDiscoveryConfigRequest) (*crdcpb.UpdateCloudDiscoveryConfigResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedUpdateRequest, req)
	return &crdcpb.UpdateCloudDiscoveryConfigResponse{CloudDiscoveryConfig: &crdcpb.CloudDiscoveryConfig{Id: req.GetCloudDiscoveryConfig().GetId()}}, nil
}
