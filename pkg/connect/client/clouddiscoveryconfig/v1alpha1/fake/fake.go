// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	crdcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_discovery_config/v1alpha1"
	"github.com/cofide/cofide-api-sdk/gen/go/proto/common/pagination/v1beta1"
	crdcv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/clouddiscoveryconfig/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type fakeCloudDiscoveryConfigClient struct {
	fake *fakeconnect.FakeConnect
}

// New instantiates a new CloudDiscoveryConfigClient for communication with a fake Connect API.
func New(fake *fakeconnect.FakeConnect) crdcv1alpha1.CloudDiscoveryConfigClient {
	return &fakeCloudDiscoveryConfigClient{fake: fake}
}

func (c *fakeCloudDiscoveryConfigClient) CreateCloudDiscoveryConfig(ctx context.Context, config *crdcpb.CloudDiscoveryConfig) (*crdcpb.CloudDiscoveryConfig, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.CloudAccounts[config.GetCloudAccountId()]; !ok {
		return nil, status.Error(codes.InvalidArgument, "cloud account not found")
	}

	if config.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "name must be non-empty")
	}

	if config.GetDiscoveryConfig() == nil {
		return nil, status.Error(codes.InvalidArgument, "exactly one discovery config must be set")
	}

	config = clone(config)
	config.Id = uuid.New().String()
	config.CreatedAt = timestamppb.Now()
	config.LastUpdatedAt = nil
	// Output only fields are set by Connect
	config.Status = crdcpb.DiscoveryStatus_DISCOVERY_STATUS_UNSPECIFIED
	config.LastSuccessfulDiscovery = nil
	config.StatusLastUpdatedAt = nil
	c.fake.CloudDiscoveryConfigs[config.GetId()] = config
	return clone(config), nil
}

func (c *fakeCloudDiscoveryConfigClient) DeleteCloudDiscoveryConfig(ctx context.Context, configID string) error {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	if _, ok := c.fake.CloudDiscoveryConfigs[configID]; !ok {
		return status.Error(codes.NotFound, "cloud resource discovery config not found")
	}
	delete(c.fake.CloudDiscoveryConfigs, configID)
	return nil
}

func (c *fakeCloudDiscoveryConfigClient) GetCloudDiscoveryConfig(ctx context.Context, configID string) (*crdcpb.CloudDiscoveryConfig, error) {
	c.fake.Mu.RLock()
	defer c.fake.Mu.RUnlock()

	config, ok := c.fake.CloudDiscoveryConfigs[configID]
	if !ok {
		return nil, status.Error(codes.NotFound, "cloud resource discovery config not found")
	}
	return clone(config), nil
}

func (c *fakeCloudDiscoveryConfigClient) ListCloudDiscoveryConfigs(ctx context.Context, filter *crdcpb.ListCloudDiscoveryConfigsRequest_Filter, pagination *v1beta1.PageRequest) ([]*crdcpb.CloudDiscoveryConfig, string, error) {
	c.fake.Mu.RLock()
	defer c.fake.Mu.RUnlock()

	// Filter
	configs := []*crdcpb.CloudDiscoveryConfig{}
	for config := range fakeconnect.SortedValues(c.fake.CloudDiscoveryConfigs) {
		if configMatches(config, filter) {
			configs = append(configs, clone(config))
		}
	}

	// Pagination (token is ID of first item that didn't fit on previous page)
	if pagination.GetPageToken() != "" {
		for len(configs) > 0 && configs[0].GetId() != pagination.GetPageToken() {
			configs = configs[1:]
		}
	}
	var nextPageToken string
	if pagination.GetPageSize() > 0 && len(configs) > int(pagination.GetPageSize()) {
		nextPageToken = configs[pagination.GetPageSize()].GetId()
		configs = configs[:pagination.GetPageSize()]
	}

	return configs, nextPageToken, nil
}

func configMatches(_ *crdcpb.CloudDiscoveryConfig, filter *crdcpb.ListCloudDiscoveryConfigsRequest_Filter) bool {
	if filter == nil {
		return true
	}
	// as filters are added, support them in the fake here
	return true
}

func (c *fakeCloudDiscoveryConfigClient) UpdateAWSLambdaCloudDiscoveryConfig(ctx context.Context, config *crdcpb.CloudDiscoveryConfig, updateMask *crdcpb.UpdateCloudDiscoveryConfigRequest_AWSLambdaDiscoveryConfigUpdateMask) (*crdcpb.CloudDiscoveryConfig, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	existing, ok := c.fake.CloudDiscoveryConfigs[config.GetId()]
	if !ok {
		return nil, status.Error(codes.NotFound, "cloud resource discovery config not found")
	}
	existing = clone(existing)

	if existing.GetAwsLambdaDiscoveryConfig() == nil {
		return nil, status.Error(codes.InvalidArgument, "cloud resource discovery config is not an aws lambda discovery config")
	}

	if updateMask == nil {
		existing = replace(existing, config)
		if existing.GetAwsLambdaDiscoveryConfig() == nil {
			return nil, status.Error(codes.InvalidArgument, "cloud resource discovery config is not an aws lambda discovery config")
		}
	} else {
		if err := applyCommonMask(existing, config, updateMask.GetName(), updateMask.GetDisabled(), updateMask.GetDiscoveryInterval()); err != nil {
			return nil, err
		}
		if config.GetAwsLambdaDiscoveryConfig() == nil {
			return nil, status.Error(codes.InvalidArgument, "cloud resource discovery config is not an aws lambda discovery config")
		}
		existingLambda := existing.GetAwsLambdaDiscoveryConfig()
		if existingLambda.GetRoleChain() == nil {
			existingLambda.RoleChain = &crdcpb.AWSAssumeRoleConfig{}
		}
		if updateMask.GetRegions() {
			existingLambda.Regions = config.GetAwsLambdaDiscoveryConfig().GetRegions()
		}
		applyRoleChainMask(existingLambda.GetRoleChain(), config.GetAwsLambdaDiscoveryConfig().GetRoleChain(), updateMask.GetAssumeThroughOidc(), updateMask.GetAudience(), updateMask.GetRoles())
	}

	return c.store(existing), nil
}

func (c *fakeCloudDiscoveryConfigClient) UpdateAWSAgentCoreCloudDiscoveryConfig(ctx context.Context, config *crdcpb.CloudDiscoveryConfig, updateMask *crdcpb.UpdateCloudDiscoveryConfigRequest_AWSAgentCoreDiscoveryConfigUpdateMask) (*crdcpb.CloudDiscoveryConfig, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	existing, ok := c.fake.CloudDiscoveryConfigs[config.GetId()]
	if !ok {
		return nil, status.Error(codes.NotFound, "cloud resource discovery config not found")
	}
	existing = clone(existing)

	if existing.GetAwsAgentCoreDiscoveryConfig() == nil {
		return nil, status.Error(codes.InvalidArgument, "cloud resource discovery config is not an aws agent core discovery config")
	}

	if updateMask == nil {
		existing = replace(existing, config)
		if existing.GetAwsAgentCoreDiscoveryConfig() == nil {
			return nil, status.Error(codes.InvalidArgument, "cloud resource discovery config is not an aws agent core discovery config")
		}
	} else {
		if err := applyCommonMask(existing, config, updateMask.GetName(), updateMask.GetDisabled(), updateMask.GetDiscoveryInterval()); err != nil {
			return nil, err
		}
		if config.GetAwsAgentCoreDiscoveryConfig() == nil {
			return nil, status.Error(codes.InvalidArgument, "cloud resource discovery config is not an aws agent core discovery config")
		}
		existingAgentCore := existing.GetAwsAgentCoreDiscoveryConfig()
		if existingAgentCore.GetRoleChain() == nil {
			existingAgentCore.RoleChain = &crdcpb.AWSAssumeRoleConfig{}
		}
		if updateMask.GetRegions() {
			existingAgentCore.Regions = config.GetAwsAgentCoreDiscoveryConfig().GetRegions()
		}
		applyRoleChainMask(existingAgentCore.GetRoleChain(), config.GetAwsAgentCoreDiscoveryConfig().GetRoleChain(), updateMask.GetAssumeThroughOidc(), updateMask.GetAudience(), updateMask.GetRoles())
	}

	return c.store(existing), nil
}

// store sets the last updated time, saves the config in the fake and returns a copy.
// The caller must hold the fake's lock.
func (c *fakeCloudDiscoveryConfigClient) store(config *crdcpb.CloudDiscoveryConfig) *crdcpb.CloudDiscoveryConfig {
	config.LastUpdatedAt = timestamppb.Now()
	c.fake.CloudDiscoveryConfigs[config.GetId()] = config
	return clone(config)
}

// replace returns a full replacement of existing with config, preserving immutable and output only properties.
func replace(existing, config *crdcpb.CloudDiscoveryConfig) *crdcpb.CloudDiscoveryConfig {
	replacement := clone(config)
	replacement.CloudAccountId = existing.GetCloudAccountId()
	replacement.Status = existing.GetStatus()
	replacement.LastSuccessfulDiscovery = existing.GetLastSuccessfulDiscovery()
	replacement.StatusLastUpdatedAt = existing.GetStatusLastUpdatedAt()
	replacement.CreatedAt = existing.GetCreatedAt()
	return replacement
}

// applyCommonMask applies the update mask fields common to all discovery config types.
func applyCommonMask(existing, config *crdcpb.CloudDiscoveryConfig, name, disabled, discoveryInterval bool) error {
	if name {
		if config.GetName() == "" {
			return status.Error(codes.InvalidArgument, "cloud resource discovery config name must not be empty")
		}
		existing.Name = config.GetName()
	}
	if disabled {
		existing.Disabled = config.GetDisabled()
	}
	if discoveryInterval {
		existing.DiscoveryInterval = config.GetDiscoveryInterval()
	}
	return nil
}

// applyRoleChainMask applies the update mask fields for an assume role config.
func applyRoleChainMask(existing, config *crdcpb.AWSAssumeRoleConfig, assumeThroughOIDC, audience, roles bool) {
	if assumeThroughOIDC {
		existing.AssumeThroughOidc = config.GetAssumeThroughOidc()
	}
	if audience {
		existing.Audience = config.GetAudience()
	}
	if roles {
		existing.Roles = config.GetRoles()
	}
}

func clone(config *crdcpb.CloudDiscoveryConfig) *crdcpb.CloudDiscoveryConfig {
	return proto.Clone(config).(*crdcpb.CloudDiscoveryConfig)
}
