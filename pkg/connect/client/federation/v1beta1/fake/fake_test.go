// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"testing"

	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
)

func TestFakeFederationClient_Unimplemented(t *testing.T) {
	client := New(fakeconnect.New())

	_, err := client.CreateFederation(t.Context(), nil)
	test.RequireUnimplemented(t, err)

	err = client.DeleteFederation(t.Context(), "")
	test.RequireUnimplemented(t, err)

	_, err = client.GetFederation(t.Context(), "")
	test.RequireUnimplemented(t, err)

	_, err = client.GetFederationByTrustZones(t.Context(), "", "")
	test.RequireUnimplemented(t, err)

	_, _, err = client.ListFederations(t.Context(), nil, pagination.Pagination{})
	test.RequireUnimplemented(t, err)
}
