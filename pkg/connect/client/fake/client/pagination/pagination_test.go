// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package pagination_test

import (
	"testing"

	fakepagination "github.com/cofide/cofide-api-sdk/pkg/connect/client/fake/client/pagination"
	"github.com/cofide/cofide-api-sdk/pkg/connect/client/pagination"
	"github.com/stretchr/testify/assert"
)

func Test_Apply(t *testing.T) {
	t.Parallel()

	// Items are sorted, as in the fakes, and are their own tokens.
	items := []string{"a", "b", "c", "d"}

	tests := []struct {
		name               string
		input              []string
		request            pagination.Pagination
		expected           []string
		expectedPagination pagination.Pagination
	}{
		{
			name:     "no page size returns everything",
			input:    items,
			expected: items,
		},
		{
			name:               "empty input",
			input:              []string{},
			request:            pagination.Pagination{PageSize: 2},
			expected:           []string{},
			expectedPagination: pagination.Pagination{PageSize: 2},
		},
		{
			name:               "page size larger than input has no next token",
			input:              items,
			request:            pagination.Pagination{PageSize: 10},
			expected:           items,
			expectedPagination: pagination.Pagination{PageSize: 10},
		},
		{
			name:               "page size equal to input has no next token",
			input:              items,
			request:            pagination.Pagination{PageSize: 4},
			expected:           items,
			expectedPagination: pagination.Pagination{PageSize: 4},
		},
		{
			name:               "first page has next token for first item on the next page",
			input:              items,
			request:            pagination.Pagination{PageSize: 2},
			expected:           []string{"a", "b"},
			expectedPagination: pagination.Pagination{PageSize: 2, Token: "c"},
		},
		{
			name:               "middle page starts at token",
			input:              items,
			request:            pagination.Pagination{PageSize: 1, Token: "b"},
			expected:           []string{"b"},
			expectedPagination: pagination.Pagination{PageSize: 1, Token: "c"},
		},
		{
			name:               "last page has no next token",
			input:              items,
			request:            pagination.Pagination{PageSize: 2, Token: "c"},
			expected:           []string{"c", "d"},
			expectedPagination: pagination.Pagination{PageSize: 2},
		},
		{
			name:     "token without page size returns the remainder",
			input:    items,
			request:  pagination.Pagination{Token: "c"},
			expected: []string{"c", "d"},
		},
		{
			name:               "token for deleted item resumes at next item",
			input:              []string{"a", "c", "d"},
			request:            pagination.Pagination{PageSize: 1, Token: "b"},
			expected:           []string{"c"},
			expectedPagination: pagination.Pagination{PageSize: 1, Token: "d"},
		},
		{
			name:               "token for deleted final item returns nothing",
			input:              []string{"a", "b"},
			request:            pagination.Pagination{PageSize: 1, Token: "c"},
			expected:           []string{},
			expectedPagination: pagination.Pagination{PageSize: 1},
		},
		{
			name:               "token past the end returns nothing",
			input:              items,
			request:            pagination.Pagination{PageSize: 2, Token: "z"},
			expected:           []string{},
			expectedPagination: pagination.Pagination{PageSize: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Act - Apply may modify its input, so pass a copy
			output, next := fakepagination.Apply(
				append([]string{}, tt.input...),
				func(s string) string { return s },
				tt.request,
			)

			// Assert
			assert.Equal(t, tt.expected, output)
			assert.Equal(t, tt.expectedPagination, next)
		})
	}
}
