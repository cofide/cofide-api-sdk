// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake_test

import (
	"testing"
	"time"

	cloudaccountpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_account/v1alpha1"
	crdcpb "github.com/cofide/cofide-api-sdk/gen/go/proto/cloud_discovery_config/v1alpha1"
	"github.com/cofide/cofide-api-sdk/gen/go/proto/common/pagination/v1beta1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/clouddiscoveryconfig/v1alpha1/fake"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
)

func newFake(t *testing.T) (*fakeconnect.FakeConnect, string) {
	t.Helper()
	fakeConnect := fakeconnect.New()
	cloudAccount := &cloudaccountpb.CloudAccount{
		Id:    uuid.New().String(),
		OrgId: uuid.New().String(),
		Name:  "fake",
	}
	fakeConnect.CloudAccounts[cloudAccount.GetId()] = cloudAccount
	return fakeConnect, cloudAccount.GetId()
}

func roleChain(arn string) *crdcpb.AWSAssumeRoleConfig {
	return &crdcpb.AWSAssumeRoleConfig{
		AssumeThroughOidc: true,
		Audience:          "audience",
		Roles:             []*crdcpb.AWSAssumeRoleConfig_Role{{IamRoleArn: arn, ExternalId: "external-id"}},
	}
}

func lambdaConfig(cloudAccountID, name string) *crdcpb.CloudDiscoveryConfig {
	return &crdcpb.CloudDiscoveryConfig{
		CloudAccountId:    cloudAccountID,
		Name:              name,
		DiscoveryInterval: durationpb.New(time.Minute),
		DiscoveryConfig: &crdcpb.CloudDiscoveryConfig_AwsLambdaDiscoveryConfig{
			AwsLambdaDiscoveryConfig: &crdcpb.AWSLambdaDiscoveryConfig{
				Regions:   []string{"eu-west-1"},
				RoleChain: roleChain("arn:aws:iam::123456789012:role/lambda"),
			},
		},
	}
}

func agentCoreConfig(cloudAccountID, name string) *crdcpb.CloudDiscoveryConfig {
	return &crdcpb.CloudDiscoveryConfig{
		CloudAccountId:    cloudAccountID,
		Name:              name,
		DiscoveryInterval: durationpb.New(time.Minute),
		DiscoveryConfig: &crdcpb.CloudDiscoveryConfig_AwsAgentCoreDiscoveryConfig{
			AwsAgentCoreDiscoveryConfig: &crdcpb.AWSAgentCoreDiscoveryConfig{
				Regions:   []string{"us-east-1"},
				RoleChain: roleChain("arn:aws:iam::123456789012:role/agentcore"),
			},
		},
	}
}

func Test_CreateCloudDiscoveryConfig(t *testing.T) {
	t.Parallel()

	// Arrange
	fakeConnect, cloudAccountID := newFake(t)
	fakeClient := fake.New(fakeConnect)

	noDiscoveryConfig := lambdaConfig(cloudAccountID, "fake-no-config")
	noDiscoveryConfig.DiscoveryConfig = nil
	emptyName := lambdaConfig(cloudAccountID, "")

	tests := []struct {
		name            string
		input           *crdcpb.CloudDiscoveryConfig
		expectedErr     string
		expectedErrCode codes.Code
	}{
		{
			name:  "success - aws lambda",
			input: lambdaConfig(cloudAccountID, "fake-lambda"),
		},
		{
			name:  "success - aws agent core",
			input: agentCoreConfig(cloudAccountID, "fake-agent-core"),
		},
		{
			name:            "error - cloud account does not exist",
			input:           lambdaConfig("non-existent", "fake-lambda-2"),
			expectedErr:     "rpc error: code = InvalidArgument desc = cloud account not found",
			expectedErrCode: codes.InvalidArgument,
		},
		{
			name:            "error - empty name",
			input:           emptyName,
			expectedErr:     "rpc error: code = InvalidArgument desc = name must be non-empty",
			expectedErrCode: codes.InvalidArgument,
		},
		{
			name:            "error - no discovery config",
			input:           noDiscoveryConfig,
			expectedErr:     "rpc error: code = InvalidArgument desc = exactly one discovery config must be set",
			expectedErrCode: codes.InvalidArgument,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Act
			created, err := fakeClient.CreateCloudDiscoveryConfig(t.Context(), tt.input)

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
			assert.Empty(t, created.GetLastUpdatedAt())
			assert.Equal(t, tt.input.GetCloudAccountId(), created.GetCloudAccountId())
			assert.Equal(t, tt.input.GetName(), created.GetName())
			assert.Equal(t, tt.input.GetDisabled(), created.GetDisabled())
			assert.EqualExportedValues(t, tt.input.GetDiscoveryInterval(), created.GetDiscoveryInterval())
			assert.EqualExportedValues(t, tt.input.GetAwsLambdaDiscoveryConfig(), created.GetAwsLambdaDiscoveryConfig())
			assert.EqualExportedValues(t, tt.input.GetAwsAgentCoreDiscoveryConfig(), created.GetAwsAgentCoreDiscoveryConfig())

			// Assert - stored in fake
			found, err := fakeClient.GetCloudDiscoveryConfig(t.Context(), created.GetId())
			require.NoError(t, err)
			assert.EqualExportedValues(t, created, found)
		})
	}
}

func Test_CreateCloudDiscoveryConfig_IgnoresOutputOnlyFields(t *testing.T) {
	t.Parallel()

	fakeConnect, cloudAccountID := newFake(t)
	fakeClient := fake.New(fakeConnect)

	input := lambdaConfig(cloudAccountID, "fake-lambda")
	input.Id = "client-supplied"
	input.Status = crdcpb.DiscoveryStatus_DISCOVERY_STATUS_ERROR

	created, err := fakeClient.CreateCloudDiscoveryConfig(t.Context(), input)

	require.NoError(t, err)
	assert.NotEqual(t, "client-supplied", created.GetId())
	assert.Equal(t, crdcpb.DiscoveryStatus_DISCOVERY_STATUS_UNSPECIFIED, created.GetStatus())
}

func Test_DeleteCloudDiscoveryConfig(t *testing.T) {
	t.Parallel()

	// Arrange
	fakeConnect, cloudAccountID := newFake(t)
	fakeClient := fake.New(fakeConnect)
	existing, err := fakeClient.CreateCloudDiscoveryConfig(t.Context(), lambdaConfig(cloudAccountID, "fake-lambda"))
	require.NoError(t, err)

	tests := []struct {
		name            string
		configID        string
		expectedErr     string
		expectedErrCode codes.Code
	}{
		{
			name:     "success",
			configID: existing.GetId(),
		},
		{
			name:            "error - does not exist",
			configID:        "non-existent",
			expectedErr:     "rpc error: code = NotFound desc = cloud resource discovery config not found",
			expectedErrCode: codes.NotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Act
			err := fakeClient.DeleteCloudDiscoveryConfig(t.Context(), tt.configID)

			// Assert - response as expected
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
				assert.Equal(t, tt.expectedErrCode, status.Code(err))
				return
			}
			require.NoError(t, err)

			// Assert - no longer stored in fake
			_, err = fakeClient.GetCloudDiscoveryConfig(t.Context(), tt.configID)
			require.Error(t, err)
			assert.Equal(t, codes.NotFound, status.Code(err))
		})
	}
}

func Test_GetCloudDiscoveryConfig(t *testing.T) {
	t.Parallel()

	// Arrange
	fakeConnect, cloudAccountID := newFake(t)
	fakeClient := fake.New(fakeConnect)
	existing, err := fakeClient.CreateCloudDiscoveryConfig(t.Context(), lambdaConfig(cloudAccountID, "fake-lambda"))
	require.NoError(t, err)

	tests := []struct {
		name            string
		configID        string
		expected        *crdcpb.CloudDiscoveryConfig
		expectedErr     string
		expectedErrCode codes.Code
	}{
		{
			name:     "success",
			configID: existing.GetId(),
			expected: existing,
		},
		{
			name:            "error - does not exist",
			configID:        "non-existent",
			expectedErr:     "rpc error: code = NotFound desc = cloud resource discovery config not found",
			expectedErrCode: codes.NotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Act
			found, err := fakeClient.GetCloudDiscoveryConfig(t.Context(), tt.configID)

			// Assert
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
				assert.Equal(t, tt.expectedErrCode, status.Code(err))
				return
			}
			require.NoError(t, err)
			assert.NotSame(t, tt.expected, found)
			assert.EqualExportedValues(t, tt.expected, found)
		})
	}
}

func Test_ListCloudDiscoveryConfigs(t *testing.T) {
	t.Parallel()

	// Arrange
	fakeConnect, cloudAccountID := newFake(t)
	fakeClient := fake.New(fakeConnect)
	existing1, err := fakeClient.CreateCloudDiscoveryConfig(t.Context(), lambdaConfig(cloudAccountID, "fake-lambda"))
	require.NoError(t, err)
	existing2, err := fakeClient.CreateCloudDiscoveryConfig(t.Context(), agentCoreConfig(cloudAccountID, "fake-agent-core"))
	require.NoError(t, err)

	// The fake returns results ordered by ID, and IDs are random.
	if existing1.GetId() > existing2.GetId() {
		existing1, existing2 = existing2, existing1
	}

	// When filters are added, add test cases here
	tests := []struct {
		name                  string
		filter                *crdcpb.ListCloudDiscoveryConfigsRequest_Filter
		pagination            *v1beta1.PageRequest
		expected              []*crdcpb.CloudDiscoveryConfig
		expectedNextPageToken string
	}{
		{
			name:     "success",
			expected: []*crdcpb.CloudDiscoveryConfig{existing1, existing2},
		},
		{
			name:                  "success - page with results remaining",
			pagination:            &v1beta1.PageRequest{PageSize: 1},
			expected:              []*crdcpb.CloudDiscoveryConfig{existing1},
			expectedNextPageToken: existing2.GetId(),
		},
		{
			name:       "success - last page",
			pagination: &v1beta1.PageRequest{PageSize: 1, PageToken: existing2.GetId()},
			expected:   []*crdcpb.CloudDiscoveryConfig{existing2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Act
			found, nextPageToken, err := fakeClient.ListCloudDiscoveryConfigs(t.Context(), tt.filter, tt.pagination)

			// Assert
			require.NoError(t, err)
			assert.EqualExportedValues(t, tt.expected, found)
			assert.Equal(t, tt.expectedNextPageToken, nextPageToken)
		})
	}
}

func cloneConfig(c *crdcpb.CloudDiscoveryConfig) *crdcpb.CloudDiscoveryConfig {
	return proto.Clone(c).(*crdcpb.CloudDiscoveryConfig)
}

// updatedRoleChain is a role chain with every field different to that returned by roleChain.
func updatedRoleChain() *crdcpb.AWSAssumeRoleConfig {
	return &crdcpb.AWSAssumeRoleConfig{
		AssumeThroughOidc: false,
		Audience:          "updated-audience",
		Roles:             []*crdcpb.AWSAssumeRoleConfig_Role{{IamRoleArn: "arn:aws:iam::999999999999:role/updated"}},
	}
}

// updateInput returns an update input with every mutable field changed, for the discovery config type of existing.
func updateInput(existing *crdcpb.CloudDiscoveryConfig) *crdcpb.CloudDiscoveryConfig {
	input := &crdcpb.CloudDiscoveryConfig{
		Id:                existing.GetId(),
		CloudAccountId:    "ignored",
		Name:              "updated name",
		Disabled:          true,
		DiscoveryInterval: durationpb.New(time.Hour),
	}
	if existing.GetAwsLambdaDiscoveryConfig() != nil {
		input.DiscoveryConfig = &crdcpb.CloudDiscoveryConfig_AwsLambdaDiscoveryConfig{
			AwsLambdaDiscoveryConfig: &crdcpb.AWSLambdaDiscoveryConfig{
				Regions:   []string{"us-west-2", "eu-central-1"},
				RoleChain: updatedRoleChain(),
			},
		}
	} else {
		input.DiscoveryConfig = &crdcpb.CloudDiscoveryConfig_AwsAgentCoreDiscoveryConfig{
			AwsAgentCoreDiscoveryConfig: &crdcpb.AWSAgentCoreDiscoveryConfig{
				Regions:   []string{"us-west-2", "eu-central-1"},
				RoleChain: updatedRoleChain(),
			},
		}
	}
	return input
}

// updateCase describes an update test case that is independent of the discovery config type.
// The mutate function applies the expected changes to a copy of the original config.
type updateCase struct {
	name   string
	mutate func(c *crdcpb.CloudDiscoveryConfig, roleChain *crdcpb.AWSAssumeRoleConfig, setRegions func([]string))
}

// updateCases are the success cases shared by the lambda and agent core update tests, in the same order as the masks
// returned by the mask function passed to runUpdateTests.
var updateCases = []updateCase{
	{
		name: "mask with just name",
		mutate: func(c *crdcpb.CloudDiscoveryConfig, _ *crdcpb.AWSAssumeRoleConfig, _ func([]string)) {
			c.Name = "updated name"
		},
	},
	{
		name: "mask with just disabled",
		mutate: func(c *crdcpb.CloudDiscoveryConfig, _ *crdcpb.AWSAssumeRoleConfig, _ func([]string)) {
			c.Disabled = true
		},
	},
	{
		name: "mask with just discovery interval",
		mutate: func(c *crdcpb.CloudDiscoveryConfig, _ *crdcpb.AWSAssumeRoleConfig, _ func([]string)) {
			c.DiscoveryInterval = durationpb.New(time.Hour)
		},
	},
	{
		name: "mask with just regions",
		mutate: func(_ *crdcpb.CloudDiscoveryConfig, _ *crdcpb.AWSAssumeRoleConfig, setRegions func([]string)) {
			setRegions([]string{"us-west-2", "eu-central-1"})
		},
	},
	{
		name: "mask with just assume through OIDC",
		mutate: func(_ *crdcpb.CloudDiscoveryConfig, rc *crdcpb.AWSAssumeRoleConfig, _ func([]string)) {
			rc.AssumeThroughOidc = false
		},
	},
	{
		name: "mask with just audience",
		mutate: func(_ *crdcpb.CloudDiscoveryConfig, rc *crdcpb.AWSAssumeRoleConfig, _ func([]string)) {
			rc.Audience = "updated-audience"
		},
	},
	{
		name: "mask with just roles",
		mutate: func(_ *crdcpb.CloudDiscoveryConfig, rc *crdcpb.AWSAssumeRoleConfig, _ func([]string)) {
			rc.Roles = updatedRoleChain().GetRoles()
		},
	},
}

func Test_UpdateAWSLambdaCloudDiscoveryConfig(t *testing.T) {
	t.Parallel()

	// Arrange
	fakeConnect, cloudAccountID := newFake(t)
	fakeClient := fake.New(fakeConnect)
	create := func(c *crdcpb.CloudDiscoveryConfig) *crdcpb.CloudDiscoveryConfig {
		created, err := fakeClient.CreateCloudDiscoveryConfig(t.Context(), c)
		require.NoError(t, err)
		return created
	}

	masks := []*crdcpb.UpdateCloudDiscoveryConfigRequest_AWSLambdaDiscoveryConfigUpdateMask{
		{Name: true},
		{Disabled: true},
		{DiscoveryInterval: true},
		{Regions: true},
		{AssumeThroughOidc: true},
		{Audience: true},
		{Roles: true},
	}
	require.Len(t, masks, len(updateCases))

	type testCase struct {
		name            string
		input           *crdcpb.CloudDiscoveryConfig
		updateMask      *crdcpb.UpdateCloudDiscoveryConfigRequest_AWSLambdaDiscoveryConfigUpdateMask
		original        *crdcpb.CloudDiscoveryConfig
		expected        *crdcpb.CloudDiscoveryConfig
		expectedErr     string
		expectedErrCode codes.Code
	}

	original := create(lambdaConfig(cloudAccountID, "fake-lambda-full"))
	expectedFull := updateInput(original)
	expectedFull.CloudAccountId = original.GetCloudAccountId()
	tests := []testCase{
		{
			name:     "success - empty mask updates all fields",
			input:    updateInput(original),
			original: original,
			expected: expectedFull,
		},
	}
	for i, uc := range updateCases {
		existing := create(lambdaConfig(cloudAccountID, "fake-lambda-"+uc.name))
		expected := cloneConfig(existing)
		uc.mutate(expected, expected.GetAwsLambdaDiscoveryConfig().GetRoleChain(), func(r []string) {
			expected.GetAwsLambdaDiscoveryConfig().Regions = r
		})
		tests = append(tests, testCase{
			name:       "success - " + uc.name,
			input:      updateInput(existing),
			updateMask: masks[i],
			original:   existing,
			expected:   expected,
		})
	}
	existingAgentCore := create(agentCoreConfig(cloudAccountID, "fake-agent-core"))
	emptyNameInput := updateInput(original)
	emptyNameInput.Name = ""
	tests = append(tests,
		testCase{
			name:            "error - not found",
			input:           &crdcpb.CloudDiscoveryConfig{Id: "non-existent"},
			expectedErr:     "rpc error: code = NotFound desc = cloud resource discovery config not found",
			expectedErrCode: codes.NotFound,
		},
		testCase{
			name:            "error - existing config is not an aws lambda config",
			input:           updateInput(existingAgentCore),
			expectedErr:     "rpc error: code = InvalidArgument desc = cloud resource discovery config is not an aws lambda discovery config",
			expectedErrCode: codes.InvalidArgument,
		},
		testCase{
			name:            "error - input is not an aws lambda config",
			input:           updateInput(existingAgentCore),
			updateMask:      &crdcpb.UpdateCloudDiscoveryConfigRequest_AWSLambdaDiscoveryConfigUpdateMask{Regions: true},
			original:        original,
			expectedErr:     "rpc error: code = InvalidArgument desc = cloud resource discovery config is not an aws lambda discovery config",
			expectedErrCode: codes.InvalidArgument,
		},
		testCase{
			name:            "error - empty name in mask",
			input:           emptyNameInput,
			updateMask:      &crdcpb.UpdateCloudDiscoveryConfigRequest_AWSLambdaDiscoveryConfigUpdateMask{Name: true},
			original:        original,
			expectedErr:     "rpc error: code = InvalidArgument desc = cloud resource discovery config name must not be empty",
			expectedErrCode: codes.InvalidArgument,
		},
	)
	// The "input is not an aws lambda config" case needs the input to target an existing lambda config.
	for i := range tests {
		if tests[i].name == "error - input is not an aws lambda config" {
			tests[i].input.Id = original.GetId()
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Act
			updated, err := fakeClient.UpdateAWSLambdaCloudDiscoveryConfig(t.Context(), tt.input, tt.updateMask)

			// Assert - response as expected
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
				assert.Equal(t, tt.expectedErrCode, status.Code(err))
				return
			}
			require.NoError(t, err)
			assertUpdated(t, tt.input, tt.original, tt.expected, updated)
			assert.EqualExportedValues(t, tt.expected.GetAwsLambdaDiscoveryConfig(), updated.GetAwsLambdaDiscoveryConfig())

			// Assert - stored in fake
			found, err := fakeClient.GetCloudDiscoveryConfig(t.Context(), updated.GetId())
			require.NoError(t, err)
			assert.EqualExportedValues(t, updated, found)
		})
	}
}

func Test_UpdateAWSAgentCoreCloudDiscoveryConfig(t *testing.T) {
	t.Parallel()

	// Arrange
	fakeConnect, cloudAccountID := newFake(t)
	fakeClient := fake.New(fakeConnect)
	create := func(c *crdcpb.CloudDiscoveryConfig) *crdcpb.CloudDiscoveryConfig {
		created, err := fakeClient.CreateCloudDiscoveryConfig(t.Context(), c)
		require.NoError(t, err)
		return created
	}

	masks := []*crdcpb.UpdateCloudDiscoveryConfigRequest_AWSAgentCoreDiscoveryConfigUpdateMask{
		{Name: true},
		{Disabled: true},
		{DiscoveryInterval: true},
		{Regions: true},
		{AssumeThroughOidc: true},
		{Audience: true},
		{Roles: true},
	}
	require.Len(t, masks, len(updateCases))

	type testCase struct {
		name            string
		input           *crdcpb.CloudDiscoveryConfig
		updateMask      *crdcpb.UpdateCloudDiscoveryConfigRequest_AWSAgentCoreDiscoveryConfigUpdateMask
		original        *crdcpb.CloudDiscoveryConfig
		expected        *crdcpb.CloudDiscoveryConfig
		expectedErr     string
		expectedErrCode codes.Code
	}

	original := create(agentCoreConfig(cloudAccountID, "fake-agent-core-full"))
	expectedFull := updateInput(original)
	expectedFull.CloudAccountId = original.GetCloudAccountId()
	tests := []testCase{
		{
			name:     "success - empty mask updates all fields",
			input:    updateInput(original),
			original: original,
			expected: expectedFull,
		},
	}
	for i, uc := range updateCases {
		existing := create(agentCoreConfig(cloudAccountID, "fake-agent-core-"+uc.name))
		expected := cloneConfig(existing)
		uc.mutate(expected, expected.GetAwsAgentCoreDiscoveryConfig().GetRoleChain(), func(r []string) {
			expected.GetAwsAgentCoreDiscoveryConfig().Regions = r
		})
		tests = append(tests, testCase{
			name:       "success - " + uc.name,
			input:      updateInput(existing),
			updateMask: masks[i],
			original:   existing,
			expected:   expected,
		})
	}
	existingLambda := create(lambdaConfig(cloudAccountID, "fake-lambda"))
	emptyNameInput := updateInput(original)
	emptyNameInput.Name = ""
	wrongTypeInput := updateInput(existingLambda)
	wrongTypeInput.Id = original.GetId()
	tests = append(tests,
		testCase{
			name:            "error - not found",
			input:           &crdcpb.CloudDiscoveryConfig{Id: "non-existent"},
			expectedErr:     "rpc error: code = NotFound desc = cloud resource discovery config not found",
			expectedErrCode: codes.NotFound,
		},
		testCase{
			name:            "error - existing config is not an aws agent core config",
			input:           updateInput(existingLambda),
			expectedErr:     "rpc error: code = InvalidArgument desc = cloud resource discovery config is not an aws agent core discovery config",
			expectedErrCode: codes.InvalidArgument,
		},
		testCase{
			name:            "error - input is not an aws agent core config",
			input:           wrongTypeInput,
			updateMask:      &crdcpb.UpdateCloudDiscoveryConfigRequest_AWSAgentCoreDiscoveryConfigUpdateMask{Regions: true},
			expectedErr:     "rpc error: code = InvalidArgument desc = cloud resource discovery config is not an aws agent core discovery config",
			expectedErrCode: codes.InvalidArgument,
		},
		testCase{
			name:            "error - empty name in mask",
			input:           emptyNameInput,
			updateMask:      &crdcpb.UpdateCloudDiscoveryConfigRequest_AWSAgentCoreDiscoveryConfigUpdateMask{Name: true},
			expectedErr:     "rpc error: code = InvalidArgument desc = cloud resource discovery config name must not be empty",
			expectedErrCode: codes.InvalidArgument,
		},
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Act
			updated, err := fakeClient.UpdateAWSAgentCoreCloudDiscoveryConfig(t.Context(), tt.input, tt.updateMask)

			// Assert - response as expected
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
				assert.Equal(t, tt.expectedErrCode, status.Code(err))
				return
			}
			require.NoError(t, err)
			assertUpdated(t, tt.input, tt.original, tt.expected, updated)
			assert.EqualExportedValues(t, tt.expected.GetAwsAgentCoreDiscoveryConfig(), updated.GetAwsAgentCoreDiscoveryConfig())

			// Assert - stored in fake
			found, err := fakeClient.GetCloudDiscoveryConfig(t.Context(), updated.GetId())
			require.NoError(t, err)
			assert.EqualExportedValues(t, updated, found)
		})
	}
}

// assertUpdated asserts the type independent properties of an updated config.
func assertUpdated(t *testing.T, input, original, expected, updated *crdcpb.CloudDiscoveryConfig) {
	t.Helper()
	assert.NotSame(t, input, updated)
	assert.Equal(t, input.GetId(), updated.GetId())                            // ID unchanged
	assert.Equal(t, original.GetCloudAccountId(), updated.GetCloudAccountId()) // cloud account ID unchanged
	assert.Equal(t, expected.GetName(), updated.GetName())
	assert.Equal(t, expected.GetDisabled(), updated.GetDisabled())
	assert.EqualExportedValues(t, expected.GetDiscoveryInterval(), updated.GetDiscoveryInterval())
	assert.Equal(t, original.GetStatus(), updated.GetStatus()) // output only fields unchanged
	assert.Equal(t, original.GetLastSuccessfulDiscovery(), updated.GetLastSuccessfulDiscovery())
	assert.Equal(t, original.GetStatusLastUpdatedAt(), updated.GetStatusLastUpdatedAt())
	assert.Equal(t, original.GetCreatedAt(), updated.GetCreatedAt()) // created at unchanged
	assert.NotEmpty(t, updated.GetLastUpdatedAt())
}
