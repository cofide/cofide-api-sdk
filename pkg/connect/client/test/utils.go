package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RequireUnimplemented requires that an error is a gRPC Status with the Unimplemented code.
func RequireUnimplemented(t *testing.T, err error) {
	t.Helper()
	require.Error(t, err)
	status := status.Convert(err)
	require.NotNil(t, status)
	assert.Equal(t, codes.Unimplemented, status.Code())
}

func PtrOf[T any](val T) *T {
	return &val
}
