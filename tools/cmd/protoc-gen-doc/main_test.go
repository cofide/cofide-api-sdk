// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package main

import "testing"

func TestStripInternal(t *testing.T) {
	tests := []struct {
		name    string
		comment string
		want    string
	}{
		{
			name:    "no internal comment",
			comment: " A token for the next page.\n",
			want:    " A token for the next page.\n",
		},
		{
			name:    "after the text, over several lines",
			comment: " A token.\n (-- api-linter: core::0158::response-plural-first-field=disabled\n     aip.dev/not-precedent: Not a list response. --)\n",
			want:    " A token.\n",
		},
		{
			name:    "on its own line within a sentence",
			comment: " Foo bar\n (-- api-linter: x=disabled --)\n baz.\n",
			want:    " Foo bar\n baz.\n",
		},
		{
			name:    "between paragraphs",
			comment: " Foo.\n\n (-- note --)\n Bar.\n",
			want:    " Foo.\n\n Bar.\n",
		},
		{
			name:    "inline",
			comment: " A (-- x --) b (-- y --) c.\n",
			want:    " A b c.\n",
		},
		{
			name:    "containing a pipe",
			comment: " A.\n (-- aip.dev/not-precedent: a || b --)\n",
			want:    " A.\n",
		},
		{
			name:    "followed by text on the same line",
			comment: " (-- a --) Public text.\n (-- b --)\n",
			want:    " Public text.\n",
		},
		{
			name:    "with CRLF line endings",
			comment: " A.\r\n (-- x --)\r\n B.\r\n",
			want:    " A.\r\n B.\r\n",
		},
		{
			name:    "at the end with no newline",
			comment: " A.\n (-- x --)",
			want:    " A.\n",
		},
		{
			name:    "unclosed",
			comment: " Set with (-- flag.\n",
			want:    " Set with (-- flag.\n",
		},
		{
			name:    "not in the AIP-192 form",
			comment: " Set with (--trust-domain flag) here (-- note --).\n",
			want:    " Set with (--trust-domain flag) here.\n",
		},
		{
			name:    "not in the AIP-192 form over several lines",
			comment: " (--flag sets it.\n More text.\n (-- x --)\n",
			want:    " (--flag sets it.\n More text.\n",
		},
		{
			name:    "closed without whitespace before it",
			comment: " A (-- x--) b --) c.\n",
			want:    " A c.\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stripInternal(tt.comment); got != tt.want {
				t.Errorf("stripInternal(%q) = %q, want %q", tt.comment, got, tt.want)
			}
		})
	}
}
