// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	crdcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_discovery_config/v1alpha1"
	crdcv1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/clouddiscoveryconfig/v1alpha1"
	fakepagination "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/client/pagination"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
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

func (c *fakeCloudDiscoveryConfigClient) ListCloudDiscoveryConfigs(ctx context.Context, filter *crdcpb.ListCloudDiscoveryConfigsRequest_Filter, requestPagination pagination.Pagination) ([]*crdcpb.CloudDiscoveryConfig, pagination.Pagination, error) {
	c.fake.Mu.RLock()
	defer c.fake.Mu.RUnlock()

	// Filter
	configs := []*crdcpb.CloudDiscoveryConfig{}
	for config := range fakeconnect.SortedValues(c.fake.CloudDiscoveryConfigs) {
		if configMatches(config, filter) {
			configs = append(configs, clone(config))
		}
	}

	configs, pagination := fakepagination.Apply(
		configs,
		func(config *crdcpb.CloudDiscoveryConfig) string { return config.GetId() },
		requestPagination,
	)

	return configs, pagination, nil
}

func configMatches(_ *crdcpb.CloudDiscoveryConfig, filter *crdcpb.ListCloudDiscoveryConfigsRequest_Filter) bool {
	if filter == nil {
		return true
	}
	// as filters are added, support them in the fake here
	return true
}

func (c *fakeCloudDiscoveryConfigClient) UpdateCloudDiscoveryConfig(ctx context.Context, config *crdcpb.CloudDiscoveryConfig, updateMask crdcv1alpha1.UpdateMask) (*crdcpb.CloudDiscoveryConfig, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	// Build the request as the real client does, so that the mask is interpreted as the server would see it.
	req := crdcv1alpha1.NewUpdateRequest(config, updateMask)
	switch mask := req.GetUpdateMask().(type) {
	case *crdcpb.UpdateCloudDiscoveryConfigRequest_AwsLambdaDiscoveryConfigUpdateMask:
		return c.updateAWSLambda(config, mask.AwsLambdaDiscoveryConfigUpdateMask)
	case *crdcpb.UpdateCloudDiscoveryConfigRequest_AwsAgentCoreDiscoveryConfigUpdateMask:
		return c.updateAWSAgentCore(config, mask.AwsAgentCoreDiscoveryConfigUpdateMask)
	case nil:
		// No mask is a full replacement of a config of the same type as the existing config.
		existing, ok := c.fake.CloudDiscoveryConfigs[config.GetId()]
		if !ok {
			return nil, status.Error(codes.NotFound, "cloud resource discovery config not found")
		}
		switch existing.GetDiscoveryConfig().(type) {
		case *crdcpb.CloudDiscoveryConfig_AwsLambdaDiscoveryConfig:
			return c.updateAWSLambda(config, nil)
		case *crdcpb.CloudDiscoveryConfig_AwsAgentCoreDiscoveryConfig:
			return c.updateAWSAgentCore(config, nil)
		default:
			return nil, status.Error(codes.Internal, "cloud resource discovery config has an unsupported type")
		}
	default:
		return nil, status.Error(codes.InvalidArgument, "unsupported update mask")
	}
}

// updateAWSLambda updates an AWS Lambda discovery config. A nil mask is a full replacement.
// The caller must hold the fake's lock.
func (c *fakeCloudDiscoveryConfigClient) updateAWSLambda(config *crdcpb.CloudDiscoveryConfig, updateMask *crdcpb.UpdateCloudDiscoveryConfigRequest_AWSLambdaDiscoveryConfigUpdateMask) (*crdcpb.CloudDiscoveryConfig, error) {
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

// updateAWSAgentCore updates an AWS AgentCore discovery config. A nil mask is a full replacement.
// The caller must hold the fake's lock.
func (c *fakeCloudDiscoveryConfigClient) updateAWSAgentCore(config *crdcpb.CloudDiscoveryConfig, updateMask *crdcpb.UpdateCloudDiscoveryConfigRequest_AWSAgentCoreDiscoveryConfigUpdateMask) (*crdcpb.CloudDiscoveryConfig, error) {
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
