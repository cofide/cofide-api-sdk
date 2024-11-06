// Copyright 2024 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: proto/trust_provider/v1alpha1/trust_provider.proto

package v1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TrustProviderKind int32

const (
	TrustProviderKind_TRUST_PROVIDER_KIND_UNSPECIFIED TrustProviderKind = 0
	TrustProviderKind_TRUST_PROVIDER_KIND_KUBERNETES  TrustProviderKind = 1
)

// Enum value maps for TrustProviderKind.
var (
	TrustProviderKind_name = map[int32]string{
		0: "TRUST_PROVIDER_KIND_UNSPECIFIED",
		1: "TRUST_PROVIDER_KIND_KUBERNETES",
	}
	TrustProviderKind_value = map[string]int32{
		"TRUST_PROVIDER_KIND_UNSPECIFIED": 0,
		"TRUST_PROVIDER_KIND_KUBERNETES":  1,
	}
)

func (x TrustProviderKind) Enum() *TrustProviderKind {
	p := new(TrustProviderKind)
	*p = x
	return p
}

func (x TrustProviderKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TrustProviderKind) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_trust_provider_v1alpha1_trust_provider_proto_enumTypes[0].Descriptor()
}

func (TrustProviderKind) Type() protoreflect.EnumType {
	return &file_proto_trust_provider_v1alpha1_trust_provider_proto_enumTypes[0]
}

func (x TrustProviderKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TrustProviderKind.Descriptor instead.
func (TrustProviderKind) EnumDescriptor() ([]byte, []int) {
	return file_proto_trust_provider_v1alpha1_trust_provider_proto_rawDescGZIP(), []int{0}
}

type TrustProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (x *TrustProvider) Reset() {
	*x = TrustProvider{}
	mi := &file_proto_trust_provider_v1alpha1_trust_provider_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrustProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrustProvider) ProtoMessage() {}

func (x *TrustProvider) ProtoReflect() protoreflect.Message {
	mi := &file_proto_trust_provider_v1alpha1_trust_provider_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrustProvider.ProtoReflect.Descriptor instead.
func (*TrustProvider) Descriptor() ([]byte, []int) {
	return file_proto_trust_provider_v1alpha1_trust_provider_proto_rawDescGZIP(), []int{0}
}

func (x *TrustProvider) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

var File_proto_trust_provider_v1alpha1_trust_provider_proto protoreflect.FileDescriptor

var file_proto_trust_provider_v1alpha1_trust_provider_proto_rawDesc = []byte{
	0x0a, 0x32, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x72, 0x75, 0x73,
	0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x22, 0x23, 0x0a, 0x0d, 0x54, 0x72, 0x75, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x2a, 0x5c, 0x0a, 0x11, 0x54, 0x72, 0x75, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x23, 0x0a,
	0x1f, 0x54, 0x52, 0x55, 0x53, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f,
	0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x22, 0x0a, 0x1e, 0x54, 0x52, 0x55, 0x53, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x56,
	0x49, 0x44, 0x45, 0x52, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x4b, 0x55, 0x42, 0x45, 0x52, 0x4e,
	0x45, 0x54, 0x45, 0x53, 0x10, 0x01, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x66, 0x69, 0x64, 0x65, 0x2f, 0x63, 0x6f, 0x66, 0x69,
	0x64, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_trust_provider_v1alpha1_trust_provider_proto_rawDescOnce sync.Once
	file_proto_trust_provider_v1alpha1_trust_provider_proto_rawDescData = file_proto_trust_provider_v1alpha1_trust_provider_proto_rawDesc
)

func file_proto_trust_provider_v1alpha1_trust_provider_proto_rawDescGZIP() []byte {
	file_proto_trust_provider_v1alpha1_trust_provider_proto_rawDescOnce.Do(func() {
		file_proto_trust_provider_v1alpha1_trust_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_trust_provider_v1alpha1_trust_provider_proto_rawDescData)
	})
	return file_proto_trust_provider_v1alpha1_trust_provider_proto_rawDescData
}

var file_proto_trust_provider_v1alpha1_trust_provider_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_trust_provider_v1alpha1_trust_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_trust_provider_v1alpha1_trust_provider_proto_goTypes = []any{
	(TrustProviderKind)(0), // 0: proto.trust_provider.v1alpha1.TrustProviderKind
	(*TrustProvider)(nil),  // 1: proto.trust_provider.v1alpha1.TrustProvider
}
var file_proto_trust_provider_v1alpha1_trust_provider_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_trust_provider_v1alpha1_trust_provider_proto_init() }
func file_proto_trust_provider_v1alpha1_trust_provider_proto_init() {
	if File_proto_trust_provider_v1alpha1_trust_provider_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_trust_provider_v1alpha1_trust_provider_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_trust_provider_v1alpha1_trust_provider_proto_goTypes,
		DependencyIndexes: file_proto_trust_provider_v1alpha1_trust_provider_proto_depIdxs,
		EnumInfos:         file_proto_trust_provider_v1alpha1_trust_provider_proto_enumTypes,
		MessageInfos:      file_proto_trust_provider_v1alpha1_trust_provider_proto_msgTypes,
	}.Build()
	File_proto_trust_provider_v1alpha1_trust_provider_proto = out.File
	file_proto_trust_provider_v1alpha1_trust_provider_proto_rawDesc = nil
	file_proto_trust_provider_v1alpha1_trust_provider_proto_goTypes = nil
	file_proto_trust_provider_v1alpha1_trust_provider_proto_depIdxs = nil
}
