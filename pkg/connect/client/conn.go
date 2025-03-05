// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"crypto/tls"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewTLSConnection(connectURI string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(
		&tls.Config{
			InsecureSkipVerify: true, // TODO: migrate away from this flag
		},
	)))
	return grpc.NewClient(
		connectURI,
		opts...,
	)
}

func NewMTLSConnection(connectURI string, connectID spiffeid.ID, source *workloadapi.X509Source, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithTransportCredentials(
		credentials.NewTLS(
			tlsconfig.MTLSClientConfig(
				source,
				source,
				tlsconfig.AuthorizeID(connectID),
			),
		),
	))
	return grpc.NewClient(
		connectURI,
		opts...,
	)
}
