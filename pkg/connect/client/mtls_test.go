// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestNewSPIFFEMTLSClient(t *testing.T) {
	tests := []struct {
		name      string
		config    *Config
		opts      []MTLSOption
		extraOpts []grpc.DialOption
		wantErr   string
	}{
		{
			name:    "nil config",
			config:  nil,
			wantErr: "config cannot be nil",
		},
		{
			name: "invalid trust domain",
			config: &Config{
				ConnectURL:         "localhost:8080",
				ConnectTrustDomain: "not a valid trust domain",
			},
			wantErr: "invalid connect trust domain",
		},
		{
			name: "empty trust domain",
			config: &Config{
				ConnectURL:         "localhost:8080",
				ConnectTrustDomain: "",
			},
			wantErr: "invalid connect trust domain",
		},
		{
			name: "valid config",
			config: &Config{
				ConnectURL:         "localhost:8080",
				ConnectTrustDomain: "example.org",
			},
		},
		{
			name: "valid config with custom subdomains",
			config: &Config{
				ConnectURL:         "localhost:8080",
				ConnectTrustDomain: "example.org",
			},
			opts: []MTLSOption{
				WithServerSubdomain("custom-connect"),
				WithAgentSubdomain("custom-connect-agent"),
			},
		},
		{
			name: "valid config with metadata interceptor",
			config: &Config{
				ConnectURL:         "localhost:8080",
				ConnectTrustDomain: "example.org",
			},
			extraOpts: []grpc.DialOption{
				grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
					ctx = metadata.AppendToOutgoingContext(ctx, "agent-id", "test-agent-id")
					ctx = metadata.AppendToOutgoingContext(ctx, "cluster-id", "test-cluster-id")
					return invoker(ctx, method, req, reply, cc, opts...)
				}),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clientSet, conn, err := NewSPIFFEMTLSClient(tt.config, nil, nil, tt.opts, tt.extraOpts...)
			if tt.wantErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tt.wantErr)
				assert.Nil(t, clientSet)
				assert.Nil(t, conn)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, clientSet)
				require.NotNil(t, conn)
				require.NoError(t, conn.Close())
			}
		})
	}
}
