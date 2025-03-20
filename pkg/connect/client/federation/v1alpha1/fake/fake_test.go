package fake

import (
	"context"
	"testing"

	federationpb "github.com/cofide/cofide-api-sdk/gen/go/proto/federation/v1alpha1"
	fakeconnect "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/connect"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func TestGetFederation(t *testing.T) {
	fakeConnect := &fakeconnect.FakeConnect{
		Federations: make(map[string]*federationpb.Federation),
	}
	client := New(fakeConnect)

	ctx := context.Background()
	federationID := uuid.New().String()
	federation := &federationpb.Federation{
		Id: &federationID,
	}

	// Add federation to fake storage
	fakeConnect.Federations[federationID] = federation

	t.Run("Federation exists", func(t *testing.T) {
		got, err := client.GetFederation(ctx, federationID)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if !proto.Equal(got, federation) {
			t.Errorf("expected %v, got %v", federation, got)
		}
	})

	t.Run("Federation does not exist", func(t *testing.T) {
		nonExistentID := uuid.New().String()
		_, err := client.GetFederation(ctx, nonExistentID)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
		if status.Code(err) != codes.NotFound {
			t.Errorf("expected NotFound error, got %v", err)
		}
	})
}
