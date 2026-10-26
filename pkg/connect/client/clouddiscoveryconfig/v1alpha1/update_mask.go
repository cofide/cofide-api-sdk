// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	crdcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_discovery_config/v1alpha1"
)

// UpdateMask selects the fields changed by CloudDiscoveryConfigClient.UpdateCloudDiscoveryConfig.
// A nil UpdateMask requests a full replacement of all updatable fields.
// Construct an UpdateMask using a type-specific constructor, e.g. AWSLambdaUpdateMask.
type UpdateMask interface {
	// apply sets the type-specific update mask on the request.
	apply(req *crdcpb.UpdateCloudDiscoveryConfigRequest)
}

// AWSLambdaUpdateFields selects the fields of an AWS Lambda discovery config to update.
type AWSLambdaUpdateFields struct {
	Name              bool
	Disabled          bool
	DiscoveryInterval bool
	Regions           bool
	// AssumeThroughOIDC, Audience and Roles select fields of the assume role config.
	AssumeThroughOIDC bool
	Audience          bool
	Roles             bool
}

type awsLambdaUpdateMask AWSLambdaUpdateFields

// AWSLambdaUpdateMask returns an UpdateMask for an AWS Lambda discovery config.
func AWSLambdaUpdateMask(fields AWSLambdaUpdateFields) UpdateMask {
	return awsLambdaUpdateMask(fields)
}

func (m awsLambdaUpdateMask) apply(req *crdcpb.UpdateCloudDiscoveryConfigRequest) {
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

// AWSAgentCoreUpdateFields selects the fields of an AWS AgentCore discovery config to update.
type AWSAgentCoreUpdateFields struct {
	Name              bool
	Disabled          bool
	DiscoveryInterval bool
	Regions           bool
	// AssumeThroughOIDC, Audience and Roles select fields of the assume role config.
	AssumeThroughOIDC bool
	Audience          bool
	Roles             bool
}

type awsAgentCoreUpdateMask AWSAgentCoreUpdateFields

// AWSAgentCoreUpdateMask returns an UpdateMask for an AWS AgentCore discovery config.
func AWSAgentCoreUpdateMask(fields AWSAgentCoreUpdateFields) UpdateMask {
	return awsAgentCoreUpdateMask(fields)
}

func (m awsAgentCoreUpdateMask) apply(req *crdcpb.UpdateCloudDiscoveryConfigRequest) {
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
