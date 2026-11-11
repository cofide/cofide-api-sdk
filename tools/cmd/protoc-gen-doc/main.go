// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

// protoc-gen-doc runs github.com/pseudomuto/protoc-gen-doc with AIP-192
// internal comments removed from the proto comments first.
//
// Internal comments, written (-- like this --), are notes for API authors,
// such as api-linter disable comments, and are left out of Google's docs.
// protoc-gen-doc keeps them, so they are removed here, from each comment on
// its own, before protoc-gen-doc renders them.
package main

import (
	"log"
	"strings"

	gendoc "github.com/pseudomuto/protoc-gen-doc"
	"github.com/pseudomuto/protokit"
	"google.golang.org/protobuf/types/pluginpb"
)

// docPlugin is protoc-gen-doc's plugin, with internal comments stripped from
// the request it generates docs for.
type docPlugin struct {
	gendoc.Plugin
}

func (p *docPlugin) Generate(req *pluginpb.CodeGeneratorRequest) (*pluginpb.CodeGeneratorResponse, error) {
	for _, file := range req.GetProtoFile() {
		for _, loc := range file.GetSourceCodeInfo().GetLocation() {
			if loc == nil {
				continue
			}
			if loc.LeadingComments != nil {
				*loc.LeadingComments = stripInternal(*loc.LeadingComments)
			}
			if loc.TrailingComments != nil {
				*loc.TrailingComments = stripInternal(*loc.TrailingComments)
			}
			for i, c := range loc.LeadingDetachedComments {
				loc.LeadingDetachedComments[i] = stripInternal(c)
			}
		}
	}
	return p.Plugin.Generate(req)
}

// stripInternal removes the internal comments from a proto comment.
//
// An internal comment starts with "(--" followed by whitespace, and ends at
// the first "--)" with whitespace before it, as AIP-192 writes them, so that a
// "(--" in ordinary text, such as "(--trust-domain flag)", isn't taken for
// one. An internal comment on lines of its own is removed along with those
// lines, so that it doesn't leave a blank line. Otherwise only the comment and
// the spaces and tabs before it are removed, which keeps the words around it
// apart.
func stripInternal(comment string) string {
	var b strings.Builder
	for {
		start, end := findInternal(comment)
		if start < 0 {
			b.WriteString(comment)
			return b.String()
		}
		before := strings.TrimRight(comment[:start], " \t")
		after := comment[end:]
		if before == "" || strings.HasSuffix(before, "\n") {
			rest := strings.TrimLeft(after, " \t")
			if strings.HasPrefix(rest, "\r\n") {
				after = rest[2:]
			} else if strings.HasPrefix(rest, "\n") || rest == "" {
				after = strings.TrimPrefix(rest, "\n")
			}
		}
		b.WriteString(before)
		comment = after
	}
}

// findInternal returns the start and end of the first internal comment in s,
// or -1, -1 if there is none.
func findInternal(s string) (int, int) {
	for from := 0; ; {
		i := strings.Index(s[from:], "(--")
		if i < 0 {
			return -1, -1
		}
		start := from + i
		if start+3 < len(s) && isSpace(s[start+3]) {
			for j := start + 3; ; {
				k := strings.Index(s[j:], "--)")
				if k < 0 {
					return -1, -1
				}
				if isSpace(s[j+k-1]) {
					return start, j + k + 3
				}
				j += k + 1
			}
		}
		from = start + 1
	}
}

func isSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}

func main() {
	if err := protokit.RunPlugin(new(docPlugin)); err != nil {
		log.Fatal(err)
	}
}
