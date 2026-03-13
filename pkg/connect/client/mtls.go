// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"fmt"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	defaultServerSubdomain = "connect"
	defaultAgentSubdomain  = "connect-agent"
	connectSPIFFEIDFormat  = "spiffe://%s/ns/connect/sa/cofide-connect-api"
)

type Option func(*options)

type options struct {
	serverSubdomain string
	agentSubdomain  string
	grpcDialOptions []grpc.DialOption
}

// WithServerSubdomain overrides the subdomain used for the Connect server URI (default: "connect")
func WithServerSubdomain(subdomain string) Option {
	return func(o *options) {
		o.serverSubdomain = subdomain
	}
}

// WithAgentSubdomain overrides the subdomain used for the gRPC authority header (default: "connect-agent")
func WithAgentSubdomain(subdomain string) Option {
	return func(o *options) {
		o.agentSubdomain = subdomain
	}
}

// WithGRPCDialOptions appends additional gRPC dial options (e.g. interceptors).
func WithGRPCDialOptions(dialOpts ...grpc.DialOption) Option {
	return func(o *options) {
		o.grpcDialOptions = append(o.grpcDialOptions, dialOpts...)
	}
}

// NewSPIFFEMTLSClient creates a ClientSet and underlying gRPC connection to the Cofide Connect API
// secured with SPIFFE mTLS. The caller is responsible for closing the returned *grpc.ClientConn.
func NewSPIFFEMTLSClient(
	config *Config,
	x509Source *workloadapi.X509Source,
	bundleSource *workloadapi.BundleSource,
	opts ...Option,
) (ClientSet, *grpc.ClientConn, error) {
	if config == nil {
		return nil, nil, fmt.Errorf("config cannot be nil")
	}

	o := &options{
		serverSubdomain: defaultServerSubdomain,
		agentSubdomain:  defaultAgentSubdomain,
	}
	for _, opt := range opts {
		opt(o)
	}

	serverID, err := spiffeid.FromString(fmt.Sprintf(connectSPIFFEIDFormat, config.ConnectTrustDomain))
	if err != nil {
		return nil, nil, fmt.Errorf("invalid connect trust domain: %w", err)
	}

	connectURI := fmt.Sprintf("dns:///%s.%s", o.serverSubdomain, config.ConnectURL)
	authority := fmt.Sprintf("%s.%s", o.agentSubdomain, config.ConnectURL)

	grpcOpts := append([]grpc.DialOption{
		grpc.WithAuthority(authority),
		grpc.WithTransportCredentials(
			credentials.NewTLS(
				tlsconfig.MTLSClientConfig(x509Source, bundleSource, tlsconfig.AuthorizeID(serverID)),
			),
		),
	}, o.grpcDialOptions...)

	conn, err := grpc.NewClient(connectURI, grpcOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed creating mTLS gRPC client for Connect API: %w", err)
	}

	return New(conn), conn, nil
}
