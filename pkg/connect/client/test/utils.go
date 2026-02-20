package test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
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

// AssertClientImplementsService asserts that an SDK client implements all RPC methods and streams
// provided by a service.
func AssertClientImplementsService(t *testing.T, client any, desc grpc.ServiceDesc) {
	reflectClient := reflect.TypeOf(client)
	for _, method := range desc.Methods {
		_, ok := reflectClient.MethodByName(method.MethodName)
		assert.True(t, ok, "%s client missing method %s", reflectClient.Name(), method.MethodName)
	}
	for _, stream := range desc.Streams {
		_, ok := reflectClient.MethodByName(stream.StreamName)
		assert.True(t, ok, "%s client missing stream %s", reflectClient.Name(), stream.StreamName)
	}
}
