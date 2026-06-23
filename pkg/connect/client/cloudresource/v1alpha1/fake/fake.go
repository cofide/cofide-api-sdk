// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"context"

	cloudresourcesvcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/connect/cloud_resource_service/v1alpha1"
	cloudresourcepb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_resource/v1alpha1"
	cloudresourcev1alpha1 "github.com/cofide/cofide-api-sdk/pkg/connect/client/cloudresource/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type fakeCloudResourceClient struct {
	fake *fakeconnect.FakeConnect
}

// New instantiates a new CloudResourceClient for communication with a fake Connect API.
func New(fake *fakeconnect.FakeConnect) cloudresourcev1alpha1.CloudResourceClient {
	return &fakeCloudResourceClient{fake: fake}
}

func (c *fakeCloudResourceClient) GetCloudResource(ctx context.Context, cloudResourceID string) (*cloudresourcepb.CloudResource, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	cloudResource, ok := c.fake.CloudResources[cloudResourceID]
	if !ok {
		return nil, status.Error(codes.NotFound, "cloud resource not found")
	}
	return clone(cloudResource), nil
}

func (c *fakeCloudResourceClient) ListCloudResources(ctx context.Context, filter *cloudresourcesvcpb.ListCloudResourcesRequest_Filter) ([]*cloudresourcepb.CloudResource, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	cloudResources := []*cloudresourcepb.CloudResource{}
	for _, cloudResource := range c.fake.CloudResources {
		if cloudResourceMatches(cloudResource, filter) {
			cloudResources = append(cloudResources, clone(cloudResource))
		}
	}
	return cloudResources, nil
}

func (c *fakeCloudResourceClient) UpdateCloudResource(ctx context.Context, cloudResource *cloudresourcepb.CloudResource, updateMask *cloudresourcesvcpb.UpdateCloudResourceRequest_UpdateMask) (*cloudresourcepb.CloudResource, error) {
	c.fake.Mu.Lock()
	defer c.fake.Mu.Unlock()

	existing, ok := c.fake.CloudResources[cloudResource.GetId()]
	if !ok {
		return nil, status.Error(codes.NotFound, "cloud resource not found")
	}
	updated := clone(existing)
	if updateMask == nil || updateMask.Suppressed {
		updated.Suppressed = cloudResource.GetSuppressed()
	}
	c.fake.CloudResources[cloudResource.GetId()] = updated
	return clone(updated), nil
}

func cloudResourceMatches(cloudResource *cloudresourcepb.CloudResource, filter *cloudresourcesvcpb.ListCloudResourcesRequest_Filter) bool {
	if filter == nil {
		return true
	}
	if filter.OrgId != "" && cloudResource.GetOrgId() != filter.OrgId {
		return false
	}
	if filter.CloudAccountId != "" && cloudResource.GetCloudAccountId() != filter.CloudAccountId {
		return false
	}
	if !filter.IncludeSuppressed && cloudResource.GetSuppressed() {
		return false
	}
	return true
}

func clone(cloudResource *cloudresourcepb.CloudResource) *cloudresourcepb.CloudResource {
	return proto.Clone(cloudResource).(*cloudresourcepb.CloudResource)
}
