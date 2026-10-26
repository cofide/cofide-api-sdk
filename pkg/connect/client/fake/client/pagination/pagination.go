package pagination

import (
	"slices"

	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
)

func Apply[T any](input []T, token func(T) string, requestPagination pagination.Pagination) ([]T, pagination.Pagination) {
	// If input fits on one page, it will be the output
	output := input

	// First item in output should be the first one the token matches (if present)
	// That is the first item that did not fit on the previous page
	if requestPagination.Token != "" {
		output = slices.DeleteFunc(output, func(item T) bool {
			return token(item) < requestPagination.Token
		})
	}

	// Token should be first item that doesn't fit on the current page
	var nextPageToken string
	if requestPagination.PageSize > 0 && len(output) > int(requestPagination.PageSize) {
		nextPageToken = token(output[requestPagination.PageSize])
		output = output[:requestPagination.PageSize]
	}

	return output, pagination.Pagination{PageSize: requestPagination.PageSize, Token: nextPageToken}
}
