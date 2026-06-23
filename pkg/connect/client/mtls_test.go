// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"context"
	"testing"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestNewSPIFFEMTLSClient(t *testing.T) {
	tests := []struct {
		name          string
		connectTarget string
		opts          []Option
		wantErr       string
	}{
		{
			name:          "empty connectTarget",
			connectTarget: "",
			wantErr:       "connectTarget cannot be empty",
		},
		{
			name:          "valid host:port",
			connectTarget: "localhost:8443",
		},
		{
			name:          "valid connect. host",
			connectTarget: "connect.example.org:8443",
		},
		{
			name:          "valid explicit dns scheme",
			connectTarget: "dns:///connect.example.org:8443",
		},
		{
			name:          "valid passthrough scheme",
			connectTarget: "passthrough:///10.0.0.1:8443",
		},
		{
			name:          "valid custom scheme",
			connectTarget: "myresolver:///connect.example.org",
		},
		{
			name:          "valid custom SPIFFE ID path",
			connectTarget: "api.mycompany.com",
		},
		{
			name:          "valid metadata interceptor",
			connectTarget: "api.mycompany.com",
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
			connectSpiffeId := spiffeid.RequireFromString("spiffe://example.org/a/spiffe/id")
			clientSet, conn, err := NewSPIFFEMTLSClient(tt.connectTarget, connectSpiffeId, nil, tt.opts...)
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
