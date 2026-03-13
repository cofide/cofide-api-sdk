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
	// WorkloadAPIAddr is the default Unix socket address for the SPIFFE Workload API
	WorkloadAPIAddr = "unix:///spiffe-workload-api/spire-agent.sock"

	defaultServerSubdomain = "connect"
	defaultAgentSubdomain  = "connect-agent"
	connectSPIFFEIDFormat  = "spiffe://%s/ns/connect/sa/cofide-connect-api"
)

type MTLSOption func(*mtlsOptions)

type mtlsOptions struct {
	serverSubdomain string
	agentSubdomain  string
}

// WithServerSubdomain overrides the subdomain used for the Connect server URI (default: "connect")
func WithServerSubdomain(subdomain string) MTLSOption {
	return func(o *mtlsOptions) {
		o.serverSubdomain = subdomain
	}
}

// WithAgentSubdomain overrides the subdomain used for the gRPC authority header (default: "connect-agent")
func WithAgentSubdomain(subdomain string) MTLSOption {
	return func(o *mtlsOptions) {
		o.agentSubdomain = subdomain
	}
}

// NewSPIFFEMTLSClient creates a ClientSet and underlying gRPC connection to the Cofide Connect API
// secured with SPIFFE mTLS. The caller is responsible for closing the returned *grpc.ClientConn.
// Additional dial options can be supplied via extraOpts
func NewSPIFFEMTLSClient(
	config *Config,
	x509Source *workloadapi.X509Source,
	bundleSource *workloadapi.BundleSource,
	opts []MTLSOption,
	extraOpts ...grpc.DialOption,
) (ClientSet, *grpc.ClientConn, error) {
	if config == nil {
		return nil, nil, fmt.Errorf("config cannot be nil")
	}

	o := &mtlsOptions{
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
	}, extraOpts...)

	conn, err := grpc.NewClient(connectURI, grpcOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed creating mTLS gRPC client for Connect API: %w", err)
	}

	return New(conn), conn, nil
}
