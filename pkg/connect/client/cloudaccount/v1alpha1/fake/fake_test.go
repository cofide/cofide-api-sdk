// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake_test

import (
	"testing"

	cloudaccountpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_account/v1alpha1"
	crdcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_resource_discovery_config/v1alpha1"
	"github.com/cofide/cofide-api-sdk/gen/go/proto/common/pagination/v1beta1"
	organizationpb "github.com/cofide/cofide-api-sdk/gen/go/proto/organization/v1alpha1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/cloudaccount/v1alpha1/fake"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_CreateCloudAccount(t *testing.T) {
	t.Parallel()

	// Arrange
	fakeConnect := fakeconnect.New()
	org := &organizationpb.Organization{
		Id:   uuid.New().String(),
		Name: "fake",
	}
	fakeConnect.Organizations[org.GetId()] = org
	fakeClient := fake.New(fakeConnect)

	tests := []struct {
		name            string
		input           *cloudaccountpb.CloudAccount
		expectedErr     string
		expectedErrCode codes.Code
	}{
		{
			name: "success",
			input: &cloudaccountpb.CloudAccount{
				OrgId: org.GetId(),
				Name:  "fake-aws-cloud-account-1",
				Provider: &cloudaccountpb.CloudAccount_Aws{
					Aws: &cloudaccountpb.AWSAccount{
						AccountId: "1234566789012",
					},
				},
			},
		},
		{
			name: "error - org does not exist",
			input: &cloudaccountpb.CloudAccount{
				OrgId: "non-existent",
				Name:  "fake-aws-cloud-account-2",
				Provider: &cloudaccountpb.CloudAccount_Aws{
					Aws: &cloudaccountpb.AWSAccount{
						AccountId: "1234566789013",
					},
				},
			},
			expectedErr:     "rpc error: code = InvalidArgument desc = organization not found",
			expectedErrCode: codes.InvalidArgument,
		},
		{
			name: "error - empty name",
			input: &cloudaccountpb.CloudAccount{
				OrgId: org.GetId(),
				Provider: &cloudaccountpb.CloudAccount_Aws{
					Aws: &cloudaccountpb.AWSAccount{
						AccountId: "1234566789014",
					},
				},
			},
			expectedErr:     "rpc error: code = InvalidArgument desc = name must be non-empty",
			expectedErrCode: codes.InvalidArgument,
		},
		// Only AWS accounts are supported now, when other accounts are supported this will become checks that exactly
		// one is set
		{
			name: "error - empty aws account ID",
			input: &cloudaccountpb.CloudAccount{
				OrgId: org.GetId(),
				Name:  "fake-aws-cloud-account-3",
			},
			expectedErr:     "rpc error: code = InvalidArgument desc = aws account ID must be non-empty",
			expectedErrCode: codes.InvalidArgument,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Act
			created, err := fakeClient.CreateCloudAccount(t.Context(), tt.input)

			// Assert - response as expected
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
				assert.Equal(t, tt.expectedErrCode, status.Code(err))
				return
			}
			require.NoError(t, err)
			assert.NotSame(t, tt.input, created)
			assert.NotEmpty(t, created.GetId())
			assert.NotEmpty(t, created.GetCreatedAt())
			assert.Equal(t, tt.input.GetOrgId(), created.GetOrgId())
			assert.Equal(t, tt.input.GetName(), created.GetName())
			assert.EqualExportedValues(t, tt.input.GetAws(), created.GetAws())

			// Assert - stored in fake
			found, err := fakeClient.GetCloudAccount(t.Context(), created.GetId())
			require.NoError(t, err)
			assert.EqualExportedValues(t, created, found)
		})
	}
}

func Test_DeleteCloudAccount(t *testing.T) {
	t.Parallel()

	// Arrange
	fakeConnect := fakeconnect.New()
	org := &organizationpb.Organization{
		Id:   uuid.New().String(),
		Name: "fake",
	}
	fakeConnect.Organizations[org.GetId()] = org
	fakeClient := fake.New(fakeConnect)
	existing, err := fakeClient.CreateCloudAccount(
		t.Context(),
		&cloudaccountpb.CloudAccount{
			OrgId: org.GetId(),
			Name:  "fake-aws-cloud-account-1",
			Provider: &cloudaccountpb.CloudAccount_Aws{
				Aws: &cloudaccountpb.AWSAccount{
					AccountId: "1234566789012",
				},
			},
		},
	)
	require.NoError(t, err)
	withConfigs, err := fakeClient.CreateCloudAccount(
		t.Context(),
		&cloudaccountpb.CloudAccount{
			OrgId: org.GetId(),
			Name:  "fake-aws-cloud-account-2",
			Provider: &cloudaccountpb.CloudAccount_Aws{
				Aws: &cloudaccountpb.AWSAccount{
					AccountId: "1234566789013",
				},
			},
		},
	)
	require.NoError(t, err)
	config := &crdcpb.CloudResourceDiscoveryConfig{
		Id:             uuid.New().String(),
		CloudAccountId: withConfigs.GetId(),
		Name:           "fake-config",
	}
	fakeConnect.CloudResourceDiscoveryConfigs[config.GetId()] = config

	tests := []struct {
		name            string
		cloudAccountID  string
		expectedErr     string
		expectedErrCode codes.Code
		expectStored    bool // whether the cloud account should remain in the fake after an error
	}{
		{
			name:           "success",
			cloudAccountID: existing.GetId(),
		},
		{
			name:            "error - does not exist",
			cloudAccountID:  "non-existent",
			expectedErr:     "rpc error: code = NotFound desc = cloud account not found",
			expectedErrCode: codes.NotFound,
		},
		{
			name:            "error - has cloud resource discovery configs",
			cloudAccountID:  withConfigs.GetId(),
			expectedErr:     "rpc error: code = FailedPrecondition desc = cloud account has cloud resource discovery configs",
			expectedErrCode: codes.FailedPrecondition,
			expectStored:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Act
			err := fakeClient.DeleteCloudAccount(t.Context(), tt.cloudAccountID)

			// Assert - response as expected
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
				assert.Equal(t, tt.expectedErrCode, status.Code(err))
				if tt.expectStored {
					_, err = fakeClient.GetCloudAccount(t.Context(), tt.cloudAccountID)
					assert.NoError(t, err)
				}
				return
			}
			require.NoError(t, err)

			// Assert - no longer stored in fake
			_, err = fakeClient.GetCloudAccount(t.Context(), tt.cloudAccountID)
			require.Error(t, err)
			assert.Equal(t, codes.NotFound, status.Code(err))
		})
	}
}

func Test_GetCloudAccount(t *testing.T) {
	t.Parallel()

	// Arrange
	fakeConnect := fakeconnect.New()
	org := &organizationpb.Organization{
		Id:   uuid.New().String(),
		Name: "fake",
	}
	fakeConnect.Organizations[org.GetId()] = org
	fakeClient := fake.New(fakeConnect)
	existing, err := fakeClient.CreateCloudAccount(
		t.Context(),
		&cloudaccountpb.CloudAccount{
			OrgId: org.GetId(),
			Name:  "fake-aws-cloud-account-1",
			Provider: &cloudaccountpb.CloudAccount_Aws{
				Aws: &cloudaccountpb.AWSAccount{
					AccountId: "1234566789012",
				},
			},
		},
	)
	require.NoError(t, err)

	tests := []struct {
		name            string
		cloudAccountID  string
		expected        *cloudaccountpb.CloudAccount
		expectedErr     string
		expectedErrCode codes.Code
	}{
		{
			name:           "success",
			cloudAccountID: existing.GetId(),
			expected:       existing,
		},
		{
			name:            "error - does not exist",
			cloudAccountID:  "non-existent",
			expectedErr:     "rpc error: code = NotFound desc = cloud account not found",
			expectedErrCode: codes.NotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Act
			found, err := fakeClient.GetCloudAccount(t.Context(), tt.cloudAccountID)

			// Assert
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
				assert.Equal(t, status.Code(err), tt.expectedErrCode)
				return
			}
			require.NoError(t, err)
			assert.EqualExportedValues(t, tt.expected, found)
		})
	}
}

func Test_ListCloudAccounts(t *testing.T) {
	t.Parallel()

	// Arrange
	fakeConnect := fakeconnect.New()
	org := &organizationpb.Organization{
		Id:   uuid.New().String(),
		Name: "fake",
	}
	fakeConnect.Organizations[org.GetId()] = org
	fakeClient := fake.New(fakeConnect)
	existing1, err := fakeClient.CreateCloudAccount(
		t.Context(),
		&cloudaccountpb.CloudAccount{
			OrgId: org.GetId(),
			Name:  "fake-aws-cloud-account-1",
			Provider: &cloudaccountpb.CloudAccount_Aws{
				Aws: &cloudaccountpb.AWSAccount{
					AccountId: "1234566789012",
				},
			},
		},
	)
	require.NoError(t, err)
	existing2, err := fakeClient.CreateCloudAccount(
		t.Context(),
		&cloudaccountpb.CloudAccount{
			OrgId: org.GetId(),
			Name:  "fake-aws-cloud-account-2",
			Provider: &cloudaccountpb.CloudAccount_Aws{
				Aws: &cloudaccountpb.AWSAccount{
					AccountId: "1234566789013",
				},
			},
		},
	)
	require.NoError(t, err)

	// The fake returns results ordered by ID, and IDs are random.
	if existing1.GetId() > existing2.GetId() {
		existing1, existing2 = existing2, existing1
	}

	// When filters are added, add test cases here
	tests := []struct {
		name                  string
		filter                *cloudaccountpb.ListCloudAccountsRequest_Filter
		pagination            *v1beta1.PageRequest
		expected              []*cloudaccountpb.CloudAccount
		expectedNextPageToken string
	}{
		{
			name:     "success",
			expected: []*cloudaccountpb.CloudAccount{existing1, existing2},
		},
		{
			name:                  "success - page with results remaining",
			pagination:            &v1beta1.PageRequest{PageSize: 1},
			expected:              []*cloudaccountpb.CloudAccount{existing1},
			expectedNextPageToken: existing2.GetId(),
		},
		{
			name:       "success - last page",
			pagination: &v1beta1.PageRequest{PageSize: 1, PageToken: existing2.GetId()},
			expected:   []*cloudaccountpb.CloudAccount{existing2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Act
			found, nextPageToken, err := fakeClient.ListCloudAccounts(t.Context(), tt.filter, tt.pagination)

			// Assert
			require.NoError(t, err)
			assert.EqualExportedValues(t, tt.expected, found)
			assert.Equal(t, tt.expectedNextPageToken, nextPageToken)
		})
	}
}

func Test_UpdateAWSCloudAccount(t *testing.T) {
	t.Parallel()

	// Arrange
	fakeConnect := fakeconnect.New()
	org := &organizationpb.Organization{
		Id:   uuid.New().String(),
		Name: "fake",
	}
	fakeConnect.Organizations[org.GetId()] = org
	fakeClient := fake.New(fakeConnect)
	existing1, err := fakeClient.CreateCloudAccount(
		t.Context(),
		&cloudaccountpb.CloudAccount{
			OrgId: org.GetId(),
			Name:  "fake-aws-cloud-account-1",
			Provider: &cloudaccountpb.CloudAccount_Aws{
				Aws: &cloudaccountpb.AWSAccount{
					AccountId: "1234566789012",
				},
			},
		},
	)
	require.NoError(t, err)
	existing2, err := fakeClient.CreateCloudAccount(
		t.Context(),
		&cloudaccountpb.CloudAccount{
			OrgId: org.GetId(),
			Name:  "fake-aws-cloud-account-2",
			Provider: &cloudaccountpb.CloudAccount_Aws{
				Aws: &cloudaccountpb.AWSAccount{
					AccountId: "1234566789013",
				},
			},
		},
	)
	require.NoError(t, err)
	existing3, err := fakeClient.CreateCloudAccount(
		t.Context(),
		&cloudaccountpb.CloudAccount{
			OrgId: org.GetId(),
			Name:  "fake-aws-cloud-account-3",
			Provider: &cloudaccountpb.CloudAccount_Aws{
				Aws: &cloudaccountpb.AWSAccount{
					AccountId: "1234566789014",
				},
			},
		},
	)
	require.NoError(t, err)

	tests := []struct {
		name            string
		input           *cloudaccountpb.CloudAccount
		updateMask      *cloudaccountpb.UpdateCloudAccountRequest_AWSAccountUpdateMask
		expected        *cloudaccountpb.CloudAccount
		expectedErr     string
		expectedErrCode codes.Code
	}{
		{
			name: "success - empty mask updates all fields",
			input: &cloudaccountpb.CloudAccount{
				Id:   existing1.GetId(),
				Name: "updated name",
				Provider: &cloudaccountpb.CloudAccount_Aws{
					Aws: &cloudaccountpb.AWSAccount{
						AccountId: "999999999999",
					},
				},
			},
			expected: &cloudaccountpb.CloudAccount{
				Name: "updated name",
				Provider: &cloudaccountpb.CloudAccount_Aws{
					Aws: &cloudaccountpb.AWSAccount{
						AccountId: "999999999999",
					},
				},
			},
		},
		{
			name: "success - mask with just name",
			input: &cloudaccountpb.CloudAccount{
				Id:   existing2.GetId(),
				Name: "updated name",
				Provider: &cloudaccountpb.CloudAccount_Aws{
					Aws: &cloudaccountpb.AWSAccount{
						AccountId: "999999999999",
					},
				},
			},
			updateMask: &cloudaccountpb.UpdateCloudAccountRequest_AWSAccountUpdateMask{
				Name: true,
			},
			expected: &cloudaccountpb.CloudAccount{
				Name:     "updated name",
				Provider: existing2.GetProvider(),
			},
		},
		{
			name: "success - mask with just aws account ID",
			input: &cloudaccountpb.CloudAccount{
				Id:   existing3.GetId(),
				Name: "updated name",
				Provider: &cloudaccountpb.CloudAccount_Aws{
					Aws: &cloudaccountpb.AWSAccount{
						AccountId: "999999999999",
					},
				},
			},
			updateMask: &cloudaccountpb.UpdateCloudAccountRequest_AWSAccountUpdateMask{
				AwsAccountId: true,
			},
			expected: &cloudaccountpb.CloudAccount{
				Name: existing3.GetName(),
				Provider: &cloudaccountpb.CloudAccount_Aws{
					Aws: &cloudaccountpb.AWSAccount{
						AccountId: "999999999999",
					},
				},
			},
		},
		{
			name: "error - not found",
			input: &cloudaccountpb.CloudAccount{
				Id: "non-existent",
			},
			expectedErr:     "rpc error: code = NotFound desc = cloud account not found",
			expectedErrCode: codes.NotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var original *cloudaccountpb.CloudAccount
			if tt.expectedErr == "" {
				var err error
				original, err = fakeClient.GetCloudAccount(t.Context(), tt.input.GetId())
				require.NoError(t, err)
			}

			// Act
			updated, err := fakeClient.UpdateAWSCloudAccount(t.Context(), tt.input, tt.updateMask)

			// Assert - response as expected
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
				assert.Equal(t, tt.expectedErrCode, status.Code(err))
				return
			}
			require.NoError(t, err)
			assert.NotSame(t, tt.input, updated)
			assert.Equal(t, tt.input.GetId(), updated.GetId())       // ID unchanged
			assert.Equal(t, original.GetOrgId(), updated.GetOrgId()) // org ID unchanged
			assert.Equal(t, tt.expected.GetName(), updated.GetName())
			assert.EqualExportedValues(t, tt.expected.GetAws(), updated.GetAws())
			assert.Equal(t, original.GetCreatedAt(), updated.GetCreatedAt()) // created at unchanged
			assert.NotEmpty(t, updated.GetLastUpdatedAt())

			// Assert - stored in fake
			found, err := fakeClient.GetCloudAccount(t.Context(), updated.GetId())
			require.NoError(t, err)
			assert.EqualExportedValues(t, updated, found)
		})
	}
}
