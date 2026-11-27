package pagination

import (
	paginationpb "github.com/cofide/cofide-api-sdk/gen/go/cofide/common/pagination/v1beta1"
)

type Pagination struct {
	Token    string
	PageSize int32
}

// PageRequestV1Beta1 returns the v1beta1 request for the page p describes.
func (p Pagination) PageRequestV1Beta1() *paginationpb.PageRequest {
	return &paginationpb.PageRequest{PageSize: p.PageSize, PageToken: p.Token}
}

// pageResponse is a page response from a list method, which may be nil if the server sent none. PageResponse
// implements it for each pagination version.
type pageResponse interface {
	GetNextPageToken() string
}

// Next returns the pagination for the page after p, given the response for p. It keeps p's page size, and its token
// is empty if p was the last page.
func (p Pagination) Next(resp pageResponse) Pagination {
	return Pagination{PageSize: p.PageSize, Token: resp.GetNextPageToken()}
}
