// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package client

// Config holds the configuration needed to establish a connection to a Cofide Connect API
type Config struct {
	// ConnectURL is the base host:port of the Cofide Connect service
	ConnectURL string
	// ConnectTrustDomain is the SPIFFE trust domain of the Cofide Connect service
	ConnectTrustDomain string
}
