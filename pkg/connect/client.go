// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package connect

import (
	"fmt"

	sdkclient "github.com/cofide/cofide-api-sdk/pkg/connect/client"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	// WorkloadAPIAddr is the default Unix socket address for the SPIFFE Workload API
	WorkloadAPIAddr = "unix:///spiffe-workload-api/spire-agent.sock"

	serverAuthoritySubdomain = "connect"
	agentAuthoritySubdomain  = "connect-agent"
	connectSPIFFEIDFormat    = "spiffe://%s/ns/connect/sa/cofide-connect-api"
)

// NewMTLSClient creates a ClientSet and underlying gRPC connection to the Cofide Connect API
// secured with SPIFFE mTLS. The caller is responsible for closing the returned *grpc.ClientConn.
// Additional dial options can be supplied via extraOpts
func NewMTLSClient(
	config *Config,
	x509Source *workloadapi.X509Source,
	bundleSource *workloadapi.BundleSource,
	extraOpts ...grpc.DialOption,
) (sdkclient.ClientSet, *grpc.ClientConn, error) {
	if config == nil {
		return nil, nil, fmt.Errorf("config cannot be nil")
	}

	serverID, err := spiffeid.FromString(fmt.Sprintf(connectSPIFFEIDFormat, config.ConnectTrustDomain))
	if err != nil {
		return nil, nil, fmt.Errorf("invalid connect trust domain: %w", err)
	}

	connectURI := fmt.Sprintf("dns:///%s.%s", serverAuthoritySubdomain, config.ConnectURL)
	authority := fmt.Sprintf("%s.%s", agentAuthoritySubdomain, config.ConnectURL)

	opts := append([]grpc.DialOption{
		grpc.WithAuthority(authority),
		grpc.WithTransportCredentials(
			credentials.NewTLS(
				tlsconfig.MTLSClientConfig(x509Source, bundleSource, tlsconfig.AuthorizeID(serverID)),
			),
		),
	}, extraOpts...)

	conn, err := grpc.NewClient(connectURI, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed creating mTLS gRPC client for Connect API: %w", err)
	}

	return sdkclient.New(conn), conn, nil
}
