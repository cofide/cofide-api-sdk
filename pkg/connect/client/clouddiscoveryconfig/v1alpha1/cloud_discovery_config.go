// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"context"

	"github.com/cofide/cofide-api-sdk/gen/go/cofide/common/pagination/v1beta1"
	crdcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_discovery_config/v1alpha1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
	"google.golang.org/grpc"
)

// CloudDiscoveryConfigClient is an interface for a gRPC client for the v1alpha1 version of the Connect CloudDiscoveryConfigService.
type CloudDiscoveryConfigClient interface {
	CreateCloudDiscoveryConfig(ctx context.Context, config *crdcpb.CloudDiscoveryConfig) (*crdcpb.CloudDiscoveryConfig, error)
	DeleteCloudDiscoveryConfig(ctx context.Context, configID string) error
	GetCloudDiscoveryConfig(ctx context.Context, configID string) (*crdcpb.CloudDiscoveryConfig, error)
	ListCloudDiscoveryConfigs(ctx context.Context, filter *crdcpb.ListCloudDiscoveryConfigsRequest_Filter, pagination pagination.Pagination) ([]*crdcpb.CloudDiscoveryConfig, pagination.Pagination, error)
	UpdateCloudDiscoveryConfig(ctx context.Context, config *crdcpb.CloudDiscoveryConfig, updateMask UpdateMask) (*crdcpb.CloudDiscoveryConfig, error)
}

type cloudDiscoveryConfigClient struct {
	client crdcpb.CloudDiscoveryConfigServiceClient
}

// New instantiates a new CloudDiscoveryConfigClient for communication with a Connect API.
func New(conn grpc.ClientConnInterface) CloudDiscoveryConfigClient {
	return &cloudDiscoveryConfigClient{
		client: crdcpb.NewCloudDiscoveryConfigServiceClient(conn),
	}
}

func (c *cloudDiscoveryConfigClient) CreateCloudDiscoveryConfig(ctx context.Context, config *crdcpb.CloudDiscoveryConfig) (*crdcpb.CloudDiscoveryConfig, error) {
	resp, err := c.client.CreateCloudDiscoveryConfig(ctx, &crdcpb.CreateCloudDiscoveryConfigRequest{
		CloudDiscoveryConfig: config,
	})
	if err != nil {
		return nil, err
	}
	return resp.GetCloudDiscoveryConfig(), nil
}

func (c *cloudDiscoveryConfigClient) DeleteCloudDiscoveryConfig(ctx context.Context, configID string) error {
	_, err := c.client.DeleteCloudDiscoveryConfig(ctx, &crdcpb.DeleteCloudDiscoveryConfigRequest{
		CloudDiscoveryConfigId: configID,
	})
	return err
}

func (c *cloudDiscoveryConfigClient) GetCloudDiscoveryConfig(ctx context.Context, configID string) (*crdcpb.CloudDiscoveryConfig, error) {
	resp, err := c.client.GetCloudDiscoveryConfig(ctx, &crdcpb.GetCloudDiscoveryConfigRequest{
		CloudDiscoveryConfigId: configID,
	})
	if err != nil {
		return nil, err
	}
	return resp.GetCloudDiscoveryConfig(), nil
}

func (c *cloudDiscoveryConfigClient) ListCloudDiscoveryConfigs(ctx context.Context, filter *crdcpb.ListCloudDiscoveryConfigsRequest_Filter, requestPagination pagination.Pagination) ([]*crdcpb.CloudDiscoveryConfig, pagination.Pagination, error) {
	resp, err := c.client.ListCloudDiscoveryConfigs(ctx, &crdcpb.ListCloudDiscoveryConfigsRequest{
		Filter:     filter,
		Pagination: &v1beta1.PageRequest{PageSize: requestPagination.PageSize, PageToken: requestPagination.Token},
	})
	if err != nil {
		return nil, pagination.Pagination{}, err
	}
	return resp.GetCloudDiscoveryConfigs(), pagination.Pagination{PageSize: requestPagination.PageSize, Token: resp.GetPagination().GetNextPageToken()}, nil
}

func (c *cloudDiscoveryConfigClient) UpdateCloudDiscoveryConfig(ctx context.Context, config *crdcpb.CloudDiscoveryConfig, updateMask UpdateMask) (*crdcpb.CloudDiscoveryConfig, error) {
	resp, err := c.client.UpdateCloudDiscoveryConfig(ctx, NewUpdateRequest(config, updateMask))
	if err != nil {
		return nil, err
	}
	return resp.GetCloudDiscoveryConfig(), nil
}
