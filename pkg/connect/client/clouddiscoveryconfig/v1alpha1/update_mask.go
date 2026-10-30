// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	crdcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_discovery_config/v1alpha1"
)

// UpdateMask selects the fields changed by CloudDiscoveryConfigClient.UpdateCloudDiscoveryConfig.
// A nil UpdateMask requests a full replacement of all updatable fields.
// Use the UpdateMask type matching the discovery config type, e.g. AWSLambdaUpdateMask.
type UpdateMask interface {
	// apply sets the type-specific update mask on the request.
	apply(req *crdcpb.UpdateCloudDiscoveryConfigRequest)
}

// AWSLambdaUpdateMask is an UpdateMask for an AWS Lambda discovery config.
// It selects the fields of the config to update.
type AWSLambdaUpdateMask struct {
	Name              bool
	Disabled          bool
	DiscoveryInterval bool
	Regions           bool
	// AssumeThroughOIDC, Audience and Roles select fields of the assume role config.
	AssumeThroughOIDC bool
	Audience          bool
	Roles             bool
}

func (m AWSLambdaUpdateMask) apply(req *crdcpb.UpdateCloudDiscoveryConfigRequest) {
	req.UpdateMask = &crdcpb.UpdateCloudDiscoveryConfigRequest_AwsLambdaDiscoveryConfigUpdateMask{
		AwsLambdaDiscoveryConfigUpdateMask: &crdcpb.UpdateCloudDiscoveryConfigRequest_AWSLambdaDiscoveryConfigUpdateMask{
			Name:              m.Name,
			Disabled:          m.Disabled,
			DiscoveryInterval: m.DiscoveryInterval,
			Regions:           m.Regions,
			AssumeThroughOidc: m.AssumeThroughOIDC,
			Audience:          m.Audience,
			Roles:             m.Roles,
		},
	}
}

// AWSAgentCoreUpdateMask is an UpdateMask for an AWS AgentCore discovery config.
// It selects the fields of the config to update.
type AWSAgentCoreUpdateMask struct {
	Name              bool
	Disabled          bool
	DiscoveryInterval bool
	Regions           bool
	// AssumeThroughOIDC, Audience and Roles select fields of the assume role config.
	AssumeThroughOIDC bool
	Audience          bool
	Roles             bool
}

func (m AWSAgentCoreUpdateMask) apply(req *crdcpb.UpdateCloudDiscoveryConfigRequest) {
	req.UpdateMask = &crdcpb.UpdateCloudDiscoveryConfigRequest_AwsAgentCoreDiscoveryConfigUpdateMask{
		AwsAgentCoreDiscoveryConfigUpdateMask: &crdcpb.UpdateCloudDiscoveryConfigRequest_AWSAgentCoreDiscoveryConfigUpdateMask{
			Name:              m.Name,
			Disabled:          m.Disabled,
			DiscoveryInterval: m.DiscoveryInterval,
			Regions:           m.Regions,
			AssumeThroughOidc: m.AssumeThroughOIDC,
			Audience:          m.Audience,
			Roles:             m.Roles,
		},
	}
}

// NewUpdateRequest builds an UpdateCloudDiscoveryConfigRequest from a config and an update mask.
// The update mask is left unset if mask is nil, which requests a full replacement.
func NewUpdateRequest(config *crdcpb.CloudDiscoveryConfig, mask UpdateMask) *crdcpb.UpdateCloudDiscoveryConfigRequest {
	req := &crdcpb.UpdateCloudDiscoveryConfigRequest{CloudDiscoveryConfig: config}
	if mask != nil {
		mask.apply(req)
	}
	return req
}
