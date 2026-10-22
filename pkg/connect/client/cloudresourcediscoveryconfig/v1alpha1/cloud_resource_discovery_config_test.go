// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1_test

import (
	"context"
	"testing"

	crdcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_resource_discovery_config/v1alpha1"
	"github.com/cofide/cofide-api-sdk/gen/go/proto/common/pagination/v1beta1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/cloudresourcediscoveryconfig/v1alpha1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
)

// Test_CloudResourceDiscoveryConfigClient_Unimplemented tests v1alpha1.CloudResourceDiscoveryConfigClient against an unimplemented server.
// This ensures that all errors returned are not wrapped can be converted to a gRPC Status using Status.Convert.
func Test_CloudResourceDiscoveryConfigClient_Unimplemented(t *testing.T) {
	server := test.NewTestServer(t)
	crdcpb.RegisterCloudResourceDiscoveryConfigServiceServer(server.Server, &crdcpb.UnimplementedCloudResourceDiscoveryConfigServiceServer{})
	server.Serve()

	conn := server.CreateClientConn()
	client := v1alpha1.New(conn)

	_, err := client.CreateCloudResourceDiscoveryConfig(t.Context(), nil)
	test.RequireUnimplemented(t, err)

	err = client.DeleteCloudResourceDiscoveryConfig(t.Context(), "")
	test.RequireUnimplemented(t, err)

	_, err = client.GetCloudResourceDiscoveryConfig(t.Context(), "")
	test.RequireUnimplemented(t, err)

	_, _, err = client.ListCloudResourceDiscoveryConfigs(t.Context(), nil, nil)
	test.RequireUnimplemented(t, err)

	_, err = client.UpdateAWSLambdaCloudResourceDiscoveryConfig(t.Context(), nil, nil)
	test.RequireUnimplemented(t, err)

	_, err = client.UpdateAWSAgentCoreCloudResourceDiscoveryConfig(t.Context(), nil, nil)
	test.RequireUnimplemented(t, err)
}

func Test_CreateCloudResourceDiscoveryConfig(t *testing.T) {
	client, fakeService := clientForTest(t)

	input := lambdaConfig()
	fakeService.expectedCreateRequest = &crdcpb.CreateCloudResourceDiscoveryConfigRequest{
		CloudResourceDiscoveryConfig: input,
	}

	_, err := client.CreateCloudResourceDiscoveryConfig(t.Context(), input)
	require.NoError(t, err)
}

func Test_DeleteCloudResourceDiscoveryConfig(t *testing.T) {
	client, fakeService := clientForTest(t)

	input := "test-config-id"
	fakeService.expectedDeleteRequest = &crdcpb.DeleteCloudResourceDiscoveryConfigRequest{
		CloudResourceDiscoveryConfigId: input,
	}

	err := client.DeleteCloudResourceDiscoveryConfig(t.Context(), input)
	require.NoError(t, err)
}

func Test_GetCloudResourceDiscoveryConfig(t *testing.T) {
	client, fakeService := clientForTest(t)

	input := "test-config-id"
	fakeService.expectedGetRequest = &crdcpb.GetCloudResourceDiscoveryConfigRequest{
		CloudResourceDiscoveryConfigId: input,
	}

	_, err := client.GetCloudResourceDiscoveryConfig(t.Context(), input)
	require.NoError(t, err)
}

func Test_ListCloudResourceDiscoveryConfigs(t *testing.T) {
	client, fakeService := clientForTest(t)

	filter := &crdcpb.ListCloudResourceDiscoveryConfigsRequest_Filter{}
	pagination := v1beta1.PageRequest{PageSize: 100}
	fakeService.expectedListRequest = &crdcpb.ListCloudResourceDiscoveryConfigsRequest{
		Filter:     filter,
		Pagination: &pagination,
	}

	_, _, err := client.ListCloudResourceDiscoveryConfigs(t.Context(), filter, &pagination)
	require.NoError(t, err)
}

func Test_UpdateAWSLambdaCloudResourceDiscoveryConfig(t *testing.T) {
	client, fakeService := clientForTest(t)

	input := lambdaConfig()
	input.Id = "test-id"
	updateMask := &crdcpb.UpdateCloudResourceDiscoveryConfigRequest_AWSLambdaDiscoveryConfigUpdateMask{
		Name:     true,
		Regions:  true,
		Audience: true,
	}
	fakeService.expectedUpdateRequest = &crdcpb.UpdateCloudResourceDiscoveryConfigRequest{
		CloudResourceDiscoveryConfig: input,
		UpdateMask: &crdcpb.UpdateCloudResourceDiscoveryConfigRequest_AwsLambdaDiscoveryConfigUpdateMask{
			AwsLambdaDiscoveryConfigUpdateMask: updateMask,
		},
	}

	_, err := client.UpdateAWSLambdaCloudResourceDiscoveryConfig(t.Context(), input, updateMask)
	require.NoError(t, err)
}

func Test_UpdateAWSAgentCoreCloudResourceDiscoveryConfig(t *testing.T) {
	client, fakeService := clientForTest(t)

	input := agentCoreConfig()
	input.Id = "test-id"
	updateMask := &crdcpb.UpdateCloudResourceDiscoveryConfigRequest_AWSAgentCoreDiscoveryConfigUpdateMask{
		Disabled: true,
		Roles:    true,
	}
	fakeService.expectedUpdateRequest = &crdcpb.UpdateCloudResourceDiscoveryConfigRequest{
		CloudResourceDiscoveryConfig: input,
		UpdateMask: &crdcpb.UpdateCloudResourceDiscoveryConfigRequest_AwsAgentCoreDiscoveryConfigUpdateMask{
			AwsAgentCoreDiscoveryConfigUpdateMask: updateMask,
		},
	}

	_, err := client.UpdateAWSAgentCoreCloudResourceDiscoveryConfig(t.Context(), input, updateMask)
	require.NoError(t, err)
}

func lambdaConfig() *crdcpb.CloudResourceDiscoveryConfig {
	return &crdcpb.CloudResourceDiscoveryConfig{
		CloudAccountId:    "test-cloud-account-id",
		Name:              "test-name",
		DiscoveryInterval: durationpb.New(60_000_000_000),
		DiscoveryConfig: &crdcpb.CloudResourceDiscoveryConfig_AwsLambdaDiscoveryConfig{
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

func agentCoreConfig() *crdcpb.CloudResourceDiscoveryConfig {
	return &crdcpb.CloudResourceDiscoveryConfig{
		CloudAccountId: "test-cloud-account-id",
		Name:           "test-name",
		DiscoveryConfig: &crdcpb.CloudResourceDiscoveryConfig_AwsAgentCoreDiscoveryConfig{
			AwsAgentCoreDiscoveryConfig: &crdcpb.AWSAgentCoreDiscoveryConfig{
				Regions: []string{"us-east-1"},
				RoleChain: &crdcpb.AWSAssumeRoleConfig{
					Roles: []*crdcpb.AWSAssumeRoleConfig_Role{{IamRoleArn: "arn:aws:iam::123456789012:role/test", ExternalId: "ext"}},
				},
			},
		},
	}
}

func clientForTest(t *testing.T) (v1alpha1.CloudResourceDiscoveryConfigClient, *fakeCloudResourceDiscoveryConfigService) {
	t.Helper()
	server := test.NewTestServer(t)
	fakeService := &fakeCloudResourceDiscoveryConfigService{t: t}
	crdcpb.RegisterCloudResourceDiscoveryConfigServiceServer(server.Server, fakeService)
	server.Serve()

	conn := server.CreateClientConn()
	client := v1alpha1.New(conn)

	return client, fakeService
}

// fakeCloudResourceDiscoveryConfigService provides a very simple fake CloudResourceDiscoveryConfigService implementation with canned responses.
// Its purpose is to assert that the inbound request is as expected.
type fakeCloudResourceDiscoveryConfigService struct {
	t                     *testing.T
	expectedCreateRequest *crdcpb.CreateCloudResourceDiscoveryConfigRequest
	expectedDeleteRequest *crdcpb.DeleteCloudResourceDiscoveryConfigRequest
	expectedGetRequest    *crdcpb.GetCloudResourceDiscoveryConfigRequest
	expectedListRequest   *crdcpb.ListCloudResourceDiscoveryConfigsRequest
	expectedUpdateRequest *crdcpb.UpdateCloudResourceDiscoveryConfigRequest
}

func (f *fakeCloudResourceDiscoveryConfigService) CreateCloudResourceDiscoveryConfig(ctx context.Context, req *crdcpb.CreateCloudResourceDiscoveryConfigRequest) (*crdcpb.CreateCloudResourceDiscoveryConfigResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedCreateRequest, req)
	return &crdcpb.CreateCloudResourceDiscoveryConfigResponse{CloudResourceDiscoveryConfig: req.GetCloudResourceDiscoveryConfig()}, nil
}

func (f *fakeCloudResourceDiscoveryConfigService) DeleteCloudResourceDiscoveryConfig(ctx context.Context, req *crdcpb.DeleteCloudResourceDiscoveryConfigRequest) (*crdcpb.DeleteCloudResourceDiscoveryConfigResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedDeleteRequest, req)
	return &crdcpb.DeleteCloudResourceDiscoveryConfigResponse{}, nil
}

func (f *fakeCloudResourceDiscoveryConfigService) GetCloudResourceDiscoveryConfig(ctx context.Context, req *crdcpb.GetCloudResourceDiscoveryConfigRequest) (*crdcpb.GetCloudResourceDiscoveryConfigResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedGetRequest, req)
	return &crdcpb.GetCloudResourceDiscoveryConfigResponse{CloudResourceDiscoveryConfig: &crdcpb.CloudResourceDiscoveryConfig{Id: req.GetCloudResourceDiscoveryConfigId()}}, nil
}

func (f *fakeCloudResourceDiscoveryConfigService) ListCloudResourceDiscoveryConfigs(ctx context.Context, req *crdcpb.ListCloudResourceDiscoveryConfigsRequest) (*crdcpb.ListCloudResourceDiscoveryConfigsResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedListRequest, req)
	return &crdcpb.ListCloudResourceDiscoveryConfigsResponse{CloudResourceDiscoveryConfigs: []*crdcpb.CloudResourceDiscoveryConfig{}}, nil
}

func (f *fakeCloudResourceDiscoveryConfigService) UpdateCloudResourceDiscoveryConfig(ctx context.Context, req *crdcpb.UpdateCloudResourceDiscoveryConfigRequest) (*crdcpb.UpdateCloudResourceDiscoveryConfigResponse, error) {
	assert.EqualExportedValues(f.t, f.expectedUpdateRequest, req)
	return &crdcpb.UpdateCloudResourceDiscoveryConfigResponse{CloudResourceDiscoveryConfig: &crdcpb.CloudResourceDiscoveryConfig{Id: req.GetCloudResourceDiscoveryConfig().GetId()}}, nil
}
