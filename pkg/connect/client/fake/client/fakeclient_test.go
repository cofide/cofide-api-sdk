// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFakeClientSet(t *testing.T) {
	client := New(nil)
	require.NotNil(t, client)
	require.NotNil(t, client.TrustZoneV1Alpha1())
	require.NotNil(t, client.ClusterV1Alpha1())
	require.NotNil(t, client.AgentV1Alpha1())
}
