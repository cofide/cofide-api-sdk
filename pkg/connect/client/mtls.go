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

const defaultServerSPIFFEIDPath = "/ns/connect/sa/cofide-connect-api"

type Option func(*options)

type options struct {
	grpcDialOptions    []grpc.DialOption
	serverSPIFFEIDPath string
	connectAuthority   string
}

// WithGRPCDialOptions appends additional gRPC dial options (e.g. interceptors).
func WithGRPCDialOptions(dialOpts ...grpc.DialOption) Option {
	return func(o *options) {
		o.grpcDialOptions = append(o.grpcDialOptions, dialOpts...)
	}
}

// WithServerSPIFFEIDPath overrides the path component of the Connect server's SPIFFE ID.
// The full ID is spiffe://<trustDomain><path>. Defaults to /ns/connect/sa/cofide-connect-api.
// Set this when Connect is deployed in a different namespace or with a different service account.
func WithServerSPIFFEIDPath(path string) Option {
	return func(o *options) {
		o.serverSPIFFEIDPath = path
	}
}

// WithAuthority overrides the gRPC :authority header and TLS SNI sent to the Connect server.
// By default the authority is derived from the connectTarget: if the host begins with "connect."
// it is replaced with "connect-mtls.", otherwise the host is used as-is.
func WithAuthority(authority string) Option {
	return func(o *options) {
		o.connectAuthority = authority
	}
}

// NewSPIFFEMTLSClient creates a ClientSet and underlying gRPC connection to the Cofide Connect API
// secured with SPIFFE mTLS. The caller is responsible for closing the returned *grpc.ClientConn.
//
// connectTarget is the host[:port] or full gRPC URI (e.g. dns:///host:port) to dial.
// trustDomain is the SPIFFE trust domain of the Connect server.
// The gRPC :authority header and TLS SNI default to connectTarget with "mtls." prefix; use WithAuthority to override.
// The expected server SPIFFE ID path defaults to /ns/connect/sa/cofide-connect-api; use WithServerSPIFFEIDPath to override.
func NewSPIFFEMTLSClient(
	connectTarget string,
	trustDomain string,
	x509Source x509svid.Source,
	bundleSource x509bundle.Source,
	opts ...Option,
) (ClientSet, *grpc.ClientConn, error) {
	if connectTarget == "" {
		return nil, nil, fmt.Errorf("connectTarget cannot be empty")
	}
	connectTrustDomain, err := spiffeid.TrustDomainFromString(trustDomain)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid Connect trust domain: %w", err)
	}

	o := &options{
		connectAuthority:   "mtls." + connectTarget,
		serverSPIFFEIDPath: defaultServerSPIFFEIDPath,
	}
	for _, opt := range opts {
		opt(o)
	}

	serverID, err := spiffeid.FromPath(connectTrustDomain, o.serverSPIFFEIDPath)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid server SPIFFE ID: %w", err)
	}

	grpcOpts := append([]grpc.DialOption{
		grpc.WithAuthority(o.connectAuthority),
		grpc.WithTransportCredentials(
			credentials.NewTLS(
				tlsconfig.MTLSClientConfig(x509Source, bundleSource, tlsconfig.AuthorizeID(serverID)),
			),
		),
	}, o.grpcDialOptions...)

	conn, err := grpc.NewClient(connectTarget, grpcOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed creating mTLS gRPC client for Connect API: %w", err)
	}

	return New(conn), conn, nil
}
