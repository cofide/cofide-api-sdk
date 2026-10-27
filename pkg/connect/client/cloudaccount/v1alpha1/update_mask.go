// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	cloudaccountpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_account/v1alpha1"
)

// UpdateMask selects the fields changed by CloudAccountClient.UpdateCloudAccount.
// A nil UpdateMask requests a full replacement of all updatable fields.
// Construct an UpdateMask using a provider-specific constructor, e.g. AWSUpdateMask.
type UpdateMask interface {
	// apply sets the provider-specific update mask on the request.
	apply(req *cloudaccountpb.UpdateCloudAccountRequest)
}

// AWSUpdateMask is an UpdateMask for an AWS cloud account.
// Set fields to true to update them: name to update the name, and awsAccountID to update the AWS account ID.
type AWSUpdateMask struct {
	Name         bool
	AWSAccountID bool
}

func (m AWSUpdateMask) apply(req *cloudaccountpb.UpdateCloudAccountRequest) {
	req.UpdateMask = &cloudaccountpb.UpdateCloudAccountRequest_AwsAccountUpdateMask{
		AwsAccountUpdateMask: &cloudaccountpb.UpdateCloudAccountRequest_AWSAccountUpdateMask{
			Name:         m.Name,
			AwsAccountId: m.AWSAccountID,
		},
	}
}

// NewUpdateRequest builds an UpdateCloudAccountRequest from a cloud account and an update mask.
// The update mask is left unset if mask is nil, which requests a full replacement.
func NewUpdateRequest(cloudAccount *cloudaccountpb.CloudAccount, mask UpdateMask) *cloudaccountpb.UpdateCloudAccountRequest {
	req := &cloudaccountpb.UpdateCloudAccountRequest{CloudAccount: cloudAccount}
	if mask != nil {
		mask.apply(req)
	}
	return req
}
