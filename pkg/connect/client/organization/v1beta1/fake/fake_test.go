// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"testing"

	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
)

func TestFakeOrganizationClient_Unimplemented(t *testing.T) {
	client := New(fakeconnect.New())

	_, err := client.GetOrganization(t.Context(), "")
	test.RequireUnimplemented(t, err)

	_, err = client.GetOrganizationByName(t.Context(), "")
	test.RequireUnimplemented(t, err)

	_, _, err = client.ListOrganizations(t.Context(), nil, pagination.Pagination{})
	test.RequireUnimplemented(t, err)
}
