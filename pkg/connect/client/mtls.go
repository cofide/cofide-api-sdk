// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"fmt"

	"github.com/spiffe/go-spiffe/v2/bundle/x509bundle"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/svid/x509svid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type X509Source interface {
	x509svid.Source
	x509bundle.Source
}

type Option func(*options)

type options struct {
	grpcDialOptions []grpc.DialOption
	authority       string
}

// WithGRPCDialOptions appends additional gRPC dial options (e.g. interceptors).
func WithGRPCDialOptions(dialOpts ...grpc.DialOption) Option {
	return func(o *options) {
		o.grpcDialOptions = append(o.grpcDialOptions, dialOpts...)
	}
}

// WithAuthority overrides the authority sent (the TLS SNI) to select the mTLS authentication path.
// If not set no explicit authority will be set.
func WithAuthority(authority string) Option {
	return func(o *options) {
		o.authority = authority
	}
}

// NewSPIFFEMTLSClient creates a ClientSet and underlying gRPC connection to the Cofide Connect API
// secured with SPIFFE mTLS. The caller is responsible for closing the returned *grpc.ClientConn.
//
// connectTarget is the host[:port] or full gRPC URI (e.g. dns:///host:port) to dial.
// connectSpiffeId is the SPIFFE ID of the Connect server.
func NewSPIFFEMTLSClient(
	connectTarget string,
	connectSpiffeId spiffeid.ID,
	x509Source X509Source,
	opts ...Option,
) (ClientSet, *grpc.ClientConn, error) {
	if connectTarget == "" {
		return nil, nil, fmt.Errorf("connectTarget cannot be empty")
	}

	o := &options{}
	for _, opt := range opts {
		opt(o)
	}

	grpcOpts := append([]grpc.DialOption{
		grpc.WithTransportCredentials(
			credentials.NewTLS(
				tlsconfig.MTLSClientConfig(x509Source, x509Source, tlsconfig.AuthorizeID(connectSpiffeId)),
			),
		),
	}, o.grpcDialOptions...)

	if o.authority != "" {
		grpcOpts = append(
			grpcOpts,
			grpc.WithAuthority(o.authority),
		)
	}

	conn, err := grpc.NewClient(connectTarget, grpcOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed creating mTLS gRPC client for Connect API: %w", err)
	}

	return New(conn), conn, nil
}
