// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

// TestServer provides a gRPC server for testing with an in-memory bufconn transport.
type TestServer struct {
	t        *testing.T
	Listener *bufconn.Listener
	Server   *grpc.Server
}

func NewTestServer(t *testing.T) *TestServer {
	return &TestServer{
		t:        t,
		Listener: bufconn.Listen(bufSize),
		Server:   grpc.NewServer(),
	}
}

func (ts *TestServer) Serve() {
	go func() {
		err := ts.Server.Serve(ts.Listener)
		require.NoError(ts.t, err)
	}()
}

// BuildTestConn creates a gRPC client connection for the provided TestServer.
func (ts *TestServer) CreateClientConn() *grpc.ClientConn {
	bufDialer := func(context.Context, string) (net.Conn, error) {
		return ts.Listener.Dial()
	}

	conn, err := grpc.NewClient(
		"passthrough://bufnet",
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(ts.t, err)
	return conn
}
