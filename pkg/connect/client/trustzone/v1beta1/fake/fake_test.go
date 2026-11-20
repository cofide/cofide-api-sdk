// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package fake

import (
	"testing"

	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/test"
)

func TestFakeTrustZoneClient_Unimplemented(t *testing.T) {
	client := New(fakeconnect.New())

	_, err := client.CreateTrustZone(t.Context(), nil)
	test.RequireUnimplemented(t, err)

	err = client.DeleteTrustZone(t.Context(), "")
	test.RequireUnimplemented(t, err)

	_, err = client.GetTrustZone(t.Context(), "")
	test.RequireUnimplemented(t, err)

	_, err = client.GetTrustZoneByName(t.Context(), "", "")
	test.RequireUnimplemented(t, err)

	_, err = client.GetTrustZoneByTrustDomain(t.Context(), "")
	test.RequireUnimplemented(t, err)

	_, _, err = client.ListTrustZones(t.Context(), nil, pagination.Pagination{})
	test.RequireUnimplemented(t, err)

	_, err = client.UpdateTrustZone(t.Context(), nil, nil)
	test.RequireUnimplemented(t, err)

	err = client.PublishTrustZoneBundle(t.Context(), nil)
	test.RequireUnimplemented(t, err)

	err = client.RegisterCofideAgent(t.Context(), nil, nil)
	test.RequireUnimplemented(t, err)

	err = client.RegisterTrustZoneServer(t.Context(), nil, nil)
	test.RequireUnimplemented(t, err)
}
