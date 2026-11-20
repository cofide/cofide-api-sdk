// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"testing"

	"buf.build/go/bufplugin/check/checktest"
)

func TestSpec(t *testing.T) {
	checktest.SpecTest(t, spec)
}

func TestGoIdentClash(t *testing.T) {
	checktest.CheckTest{
		Request: &checktest.RequestSpec{
			Files: &checktest.ProtoFileSpec{
				DirPaths:  []string{"testdata"},
				FilePaths: []string{"clash/clash.proto"},
			},
			RuleIDs: []string{goIdentClashRuleID},
		},
		Spec: spec,
		ExpectedAnnotations: []checktest.ExpectedAnnotation{
			clash(16, 4, 39, `Oneof field "name_with_org_id" has Go wrapper type GetThingRequest_NameWithOrgId_, as its name clashes with a nested type of GetThingRequest. Rename the field or the nested type.`),
			clash(17, 4, 18, `Oneof field "kind" has Go wrapper type GetThingRequest_Kind_, as its name clashes with a nested type of GetThingRequest. Rename the field or the nested type.`),
			clash(26, 6, 20, `Oneof field "pair" has Go wrapper type Outer_Inner_Pair_, as its name clashes with a nested type of Inner. Rename the field or the nested type.`),
			clash(32, 2, 19, `Field "reset" has Go name Reset_, as its name clashes with a method protoc-gen-go generates for Methods. Rename the field.`),
			clash(33, 2, 24, `Field "descriptor" has Go name Descriptor_, as its name clashes with a method protoc-gen-go generates for Methods. Rename the field.`),
			clash(35, 2, 21, `Field "get_foo" has Go name GetFoo_, as its name clashes with a method protoc-gen-go generates for Methods. Rename the field.`),
		},
	}.Run(t)
}

// TestImports checks that clashes are only reported in the files being
// linted, not in the files they import.
func TestImports(t *testing.T) {
	checktest.CheckTest{
		Request: &checktest.RequestSpec{
			Files: &checktest.ProtoFileSpec{
				DirPaths:  []string{"testdata"},
				FilePaths: []string{"ok/ok.proto"},
			},
			RuleIDs: []string{goIdentClashRuleID},
		},
		Spec: spec,
	}.Run(t)
}

// clash returns the annotation expected for a clash on one line of
// testdata/clash/clash.proto. Lines and columns are zero-indexed.
func clash(line, startColumn, endColumn int, message string) checktest.ExpectedAnnotation {
	return checktest.ExpectedAnnotation{
		RuleID:  goIdentClashRuleID,
		Message: message,
		FileLocation: &checktest.ExpectedFileLocation{
			FileName:    "clash/clash.proto",
			StartLine:   line,
			StartColumn: startColumn,
			EndLine:     line,
			EndColumn:   endColumn,
		},
	}
}
