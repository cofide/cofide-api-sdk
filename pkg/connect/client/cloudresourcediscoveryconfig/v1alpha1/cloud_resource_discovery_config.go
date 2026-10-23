// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"

	crdcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_resource_discovery_config/v1alpha1"
	"github.com/cofide/cofide-api-sdk/gen/go/proto/common/pagination/v1beta1"
	"google.golang.org/grpc"
)

// CloudResourceDiscoveryConfigClient is an interface for a gRPC client for the v1alpha1 version of the Connect CloudResourceDiscoveryConfigService.
type CloudResourceDiscoveryConfigClient interface {
	CreateCloudResourceDiscoveryConfig(ctx context.Context, config *crdcpb.CloudResourceDiscoveryConfig) (*crdcpb.CloudResourceDiscoveryConfig, error)
	DeleteCloudResourceDiscoveryConfig(ctx context.Context, configID string) error
	GetCloudResourceDiscoveryConfig(ctx context.Context, configID string) (*crdcpb.CloudResourceDiscoveryConfig, error)
	ListCloudResourceDiscoveryConfigs(ctx context.Context, filter *crdcpb.ListCloudResourceDiscoveryConfigsRequest_Filter, pagination *v1beta1.PageRequest) ([]*crdcpb.CloudResourceDiscoveryConfig, string, error)
	UpdateAWSLambdaCloudResourceDiscoveryConfig(ctx context.Context, config *crdcpb.CloudResourceDiscoveryConfig, updateMask *crdcpb.UpdateCloudResourceDiscoveryConfigRequest_AWSLambdaDiscoveryConfigUpdateMask) (*crdcpb.CloudResourceDiscoveryConfig, error)
	UpdateAWSAgentCoreCloudResourceDiscoveryConfig(ctx context.Context, config *crdcpb.CloudResourceDiscoveryConfig, updateMask *crdcpb.UpdateCloudResourceDiscoveryConfigRequest_AWSAgentCoreDiscoveryConfigUpdateMask) (*crdcpb.CloudResourceDiscoveryConfig, error)
}

type cloudResourceDiscoveryConfigClient struct {
	client crdcpb.CloudResourceDiscoveryConfigServiceClient
}

// New instantiates a new CloudResourceDiscoveryConfigClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) CloudResourceDiscoveryConfigClient {
	return &cloudResourceDiscoveryConfigClient{
		client: crdcpb.NewCloudResourceDiscoveryConfigServiceClient(conn),
	}
}

func (c *cloudResourceDiscoveryConfigClient) CreateCloudResourceDiscoveryConfig(ctx context.Context, config *crdcpb.CloudResourceDiscoveryConfig) (*crdcpb.CloudResourceDiscoveryConfig, error) {
	resp, err := c.client.CreateCloudResourceDiscoveryConfig(ctx, &crdcpb.CreateCloudResourceDiscoveryConfigRequest{
		CloudResourceDiscoveryConfig: config,
	})
	if err != nil {
		return nil, err
	}
	return resp.GetCloudResourceDiscoveryConfig(), nil
}

func (c *cloudResourceDiscoveryConfigClient) DeleteCloudResourceDiscoveryConfig(ctx context.Context, configID string) error {
	_, err := c.client.DeleteCloudResourceDiscoveryConfig(ctx, &crdcpb.DeleteCloudResourceDiscoveryConfigRequest{
		CloudResourceDiscoveryConfigId: configID,
	})
	return err
}

func (c *cloudResourceDiscoveryConfigClient) GetCloudResourceDiscoveryConfig(ctx context.Context, configID string) (*crdcpb.CloudResourceDiscoveryConfig, error) {
	resp, err := c.client.GetCloudResourceDiscoveryConfig(ctx, &crdcpb.GetCloudResourceDiscoveryConfigRequest{
		CloudResourceDiscoveryConfigId: configID,
	})
	if err != nil {
		return nil, err
	}
	return resp.GetCloudResourceDiscoveryConfig(), nil
}

func (c *cloudResourceDiscoveryConfigClient) ListCloudResourceDiscoveryConfigs(ctx context.Context, filter *crdcpb.ListCloudResourceDiscoveryConfigsRequest_Filter, pagination *v1beta1.PageRequest) ([]*crdcpb.CloudResourceDiscoveryConfig, string, error) {
	resp, err := c.client.ListCloudResourceDiscoveryConfigs(ctx, &crdcpb.ListCloudResourceDiscoveryConfigsRequest{
		Filter:     filter,
		Pagination: pagination,
	})
	if err != nil {
		return nil, "", err
	}
	return resp.GetCloudResourceDiscoveryConfigs(), resp.GetPagination().GetNextPageToken(), nil
}

func (c *cloudResourceDiscoveryConfigClient) UpdateAWSLambdaCloudResourceDiscoveryConfig(ctx context.Context, config *crdcpb.CloudResourceDiscoveryConfig, updateMask *crdcpb.UpdateCloudResourceDiscoveryConfigRequest_AWSLambdaDiscoveryConfigUpdateMask) (*crdcpb.CloudResourceDiscoveryConfig, error) {
	resp, err := c.client.UpdateCloudResourceDiscoveryConfig(ctx, &crdcpb.UpdateCloudResourceDiscoveryConfigRequest{
		CloudResourceDiscoveryConfig: config,
		UpdateMask: &crdcpb.UpdateCloudResourceDiscoveryConfigRequest_AwsLambdaDiscoveryConfigUpdateMask{
			AwsLambdaDiscoveryConfigUpdateMask: updateMask,
		},
	})
	if err != nil {
		return nil, err
	}
	return resp.GetCloudResourceDiscoveryConfig(), nil
}

func (c *cloudResourceDiscoveryConfigClient) UpdateAWSAgentCoreCloudResourceDiscoveryConfig(ctx context.Context, config *crdcpb.CloudResourceDiscoveryConfig, updateMask *crdcpb.UpdateCloudResourceDiscoveryConfigRequest_AWSAgentCoreDiscoveryConfigUpdateMask) (*crdcpb.CloudResourceDiscoveryConfig, error) {
	resp, err := c.client.UpdateCloudResourceDiscoveryConfig(ctx, &crdcpb.UpdateCloudResourceDiscoveryConfigRequest{
		CloudResourceDiscoveryConfig: config,
		UpdateMask: &crdcpb.UpdateCloudResourceDiscoveryConfigRequest_AwsAgentCoreDiscoveryConfigUpdateMask{
			AwsAgentCoreDiscoveryConfigUpdateMask: updateMask,
		},
	})
	if err != nil {
		return nil, err
	}
	return resp.GetCloudResourceDiscoveryConfig(), nil
}
