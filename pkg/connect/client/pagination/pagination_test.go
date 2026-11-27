package pagination_test

import (
	"testing"

	paginationpb "github.com/cofide/cofide-api-sdk/gen/go/cofide/common/pagination/v1beta1"
	paginationv1alpha1pb "github.com/cofide/cofide-api-sdk/gen/go/proto/pagination/v1alpha1"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
	"github.com/stretchr/testify/assert"
)

func TestPagination_PageRequestV1Beta1(t *testing.T) {
	p := pagination.Pagination{PageSize: 100, Token: "test-token"}
	assert.EqualExportedValues(t, &paginationpb.PageRequest{PageSize: 100, PageToken: "test-token"}, p.PageRequestV1Beta1())
}

func TestPagination_Next(t *testing.T) {
	p := pagination.Pagination{PageSize: 100, Token: "test-token"}

	tests := []struct {
		name string
		resp interface{ GetNextPageToken() string }
		want pagination.Pagination
	}{
		{
			name: "v1beta1",
			resp: &paginationpb.PageResponse{NextPageToken: "next-token"},
			want: pagination.Pagination{PageSize: 100, Token: "next-token"},
		},
		{
			name: "v1alpha1",
			resp: &paginationv1alpha1pb.PageResponse{NextPageToken: "next-token"},
			want: pagination.Pagination{PageSize: 100, Token: "next-token"},
		},
		{
			name: "last page",
			resp: &paginationpb.PageResponse{},
			want: pagination.Pagination{PageSize: 100},
		},
		{
			name: "no response",
			resp: (*paginationpb.PageResponse)(nil),
			want: pagination.Pagination{PageSize: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, p.Next(tt.resp))
		})
	}
}
