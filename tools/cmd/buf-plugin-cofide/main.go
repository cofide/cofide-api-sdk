// Copyright 2026 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

// buf-plugin-cofide is a buf check plugin with Cofide's own lint rules.
//
// GO_IDENT_CLASH reports proto names that protoc-gen-go renames by appending
// underscores. protoc-gen-go names a oneof field's wrapper type
// <Message>_<Field>, the same way it names a nested message or enum, and a
// struct field the same as a generated method such as Reset or GetFoo. When
// two names clash, it appends underscores to one of them until they don't,
// giving names such as GetTrustZoneRequest_NameWithOrgId_. It has no option to
// name them differently, so the proto names have to change instead.
//
// Rather than copy protoc-gen-go's naming rules, the rule asks protoc-gen-go's
// own library, protogen, which names protoc-gen-go would generate.
package main

import (
	"context"
	"fmt"
	"path"
	"strings"

	"buf.build/go/bufplugin/check"
	"buf.build/go/bufplugin/descriptor"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/pluginpb"
)

const goIdentClashRuleID = "GO_IDENT_CLASH"

var spec = &check.Spec{
	Rules: []*check.RuleSpec{
		{
			ID:      goIdentClashRuleID,
			Purpose: "Checks that protoc-gen-go does not append underscores to Go names to avoid clashes.",
			Type:    check.RuleTypeLint,
			Handler: check.RuleHandlerFunc(checkGoIdentClash),
		},
	},
}

func main() {
	check.Main(spec)
}

func checkGoIdentClash(ctx context.Context, w check.ResponseWriter, req check.Request) error {
	plugin, err := newProtogenPlugin(req.FileDescriptors())
	if err != nil {
		return err
	}
	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}
		for _, message := range file.Messages {
			checkMessage(w, message)
		}
	}
	return nil
}

func checkMessage(w check.ResponseWriter, message *protogen.Message) {
	for _, field := range message.Fields {
		if strings.HasSuffix(field.GoName, "_") {
			w.AddAnnotation(
				check.WithDescriptor(field.Desc),
				check.WithMessagef(
					"Field %q has Go name %s, as its name clashes with a method protoc-gen-go generates for %s. Rename the field.",
					field.Desc.Name(), field.GoName, message.GoIdent.GoName,
				),
			)
		}
		// protogen names a wrapper for a proto3 optional field too, as it's in
		// a synthetic oneof, but protoc-gen-go doesn't generate one.
		if field.Oneof != nil && !field.Oneof.Desc.IsSynthetic() && strings.HasSuffix(field.GoIdent.GoName, "_") {
			w.AddAnnotation(
				check.WithDescriptor(field.Desc),
				check.WithMessagef(
					"Oneof field %q has Go wrapper type %s, as its name clashes with a nested type of %s. Rename the field or the nested type.",
					field.Desc.Name(), field.GoIdent.GoName, message.Desc.Name(),
				),
			)
		}
	}
	for _, nested := range message.Messages {
		checkMessage(w, nested)
	}
}

// newProtogenPlugin builds a protogen.Plugin for the files, as protoc-gen-go
// would for a request to generate code for the files being linted.
func newProtogenPlugin(files []descriptor.FileDescriptor) (*protogen.Plugin, error) {
	req := &pluginpb.CodeGeneratorRequest{}
	// protogen needs a Go import path for every file. Names within a file don't
	// depend on it, so each directory is given its own placeholder path, which
	// also overrides any go_package option.
	var params []string
	seen := map[string]bool{}
	var add func(protoreflect.FileDescriptor)
	add = func(file protoreflect.FileDescriptor) {
		if seen[file.Path()] {
			return
		}
		seen[file.Path()] = true
		// protogen needs a file's imports before the file itself.
		imports := file.Imports()
		for i := range imports.Len() {
			add(imports.Get(i).FileDescriptor)
		}
		fileProto := protodesc.ToFileDescriptorProto(file)
		// Drop the go_package option, so that the placeholder import path is
		// used for the package name too.
		if fileProto.Options != nil {
			fileProto.Options = proto.CloneOf(fileProto.Options)
			fileProto.Options.GoPackage = nil
		}
		req.ProtoFile = append(req.ProtoFile, fileProto)
		params = append(params, fmt.Sprintf("M%s=cofide.invalid/lint/%s", file.Path(), path.Dir(file.Path())))
	}
	for _, file := range files {
		add(file.ProtoreflectFileDescriptor())
		if !file.IsImport() {
			req.FileToGenerate = append(req.FileToGenerate, file.ProtoreflectFileDescriptor().Path())
		}
	}
	req.Parameter = proto.String(strings.Join(params, ","))
	return protogen.Options{}.New(req)
}

