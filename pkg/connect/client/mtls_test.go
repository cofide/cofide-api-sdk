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
		name          string
		connectTarget string
		trustDomain   string
		opts          []Option
		wantErr       string
	}{
		{
			name:          "empty connectTarget",
			connectTarget: "",
			trustDomain:   "example.org",
			wantErr:       "connectTarget cannot be empty",
		},
		{
			name:          "invalid trust domain",
			connectTarget: "localhost:8443",
			trustDomain:   "not a valid trust domain",
			wantErr:       "invalid Connect trust domain: trust domain characters are limited to lowercase letters, numbers, dots, dashes, and underscores",
		},
		{
			name:          "empty trust domain",
			connectTarget: "localhost:8443",
			trustDomain:   "",
			wantErr:       "invalid Connect trust domain: trust domain is missing",
		},
		{
			name:          "invalid custom SPIFFE ID path",
			connectTarget: "localhost:8443",
			trustDomain:   "example.org",
			opts:          []Option{WithServerSPIFFEIDPath("not/a/valid/path")},
			wantErr:       "invalid server SPIFFE ID: path must have a leading slash",
		},
		{
			name:          "valid host:port",
			connectTarget: "localhost:8443",
			trustDomain:   "example.org",
		},
		{
			name:          "valid connect. host",
			connectTarget: "connect.example.org:8443",
			trustDomain:   "example.org",
		},
		{
			name:          "valid explicit dns scheme",
			connectTarget: "dns:///connect.example.org:8443",
			trustDomain:   "example.org",
		},
		{
			name:          "valid passthrough scheme",
			connectTarget: "passthrough:///10.0.0.1:8443",
			trustDomain:   "example.org",
		},
		{
			name:          "valid custom scheme",
			connectTarget: "myresolver:///connect.example.org",
			trustDomain:   "example.org",
		},
		{
			name:          "valid authority override",
			connectTarget: "api.mycompany.com",
			trustDomain:   "mycompany.com",
			opts:          []Option{WithAuthority("connect-spiffe.mycompany.com")},
		},
		{
			name:          "valid custom SPIFFE ID path",
			connectTarget: "api.mycompany.com",
			trustDomain:   "mycompany.com",
			opts:          []Option{WithServerSPIFFEIDPath("/ns/production/sa/cofide-connect")},
		},
		{
			name:          "valid metadata interceptor",
			connectTarget: "api.mycompany.com",
			trustDomain:   "mycompany.com",
			opts: []Option{
				WithGRPCDialOptions(grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
					ctx = metadata.AppendToOutgoingContext(ctx, "agent-id", "test-agent-id")
					ctx = metadata.AppendToOutgoingContext(ctx, "cluster-id", "test-cluster-id")
					return invoker(ctx, method, req, reply, cc, opts...)
				})),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clientSet, conn, err := NewSPIFFEMTLSClient(tt.connectTarget, tt.trustDomain, nil, nil, tt.opts...)
			if tt.wantErr != "" {
				assert.Nil(t, clientSet)
				assert.Nil(t, conn)
				require.EqualError(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
			assert.NotNil(t, clientSet)
			require.NotNil(t, conn)
			assert.Equal(t, tt.connectTarget, conn.Target())
			require.NoError(t, conn.Close())
		})
	}
}
