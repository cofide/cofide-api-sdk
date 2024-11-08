// Copyright 2024 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: proto/cofidectl_plugin/v1alpha1/plugin.proto

package v1alpha1

import (
	v1alpha1 "github.com/cofide/cofide-api-sdk/gen/go/proto/trust_zone/v1alpha1"
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

type ListTrustZonesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListTrustZonesRequest) Reset() {
	*x = ListTrustZonesRequest{}
	mi := &file_proto_cofidectl_plugin_v1alpha1_plugin_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTrustZonesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTrustZonesRequest) ProtoMessage() {}

func (x *ListTrustZonesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cofidectl_plugin_v1alpha1_plugin_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTrustZonesRequest.ProtoReflect.Descriptor instead.
func (*ListTrustZonesRequest) Descriptor() ([]byte, []int) {
	return file_proto_cofidectl_plugin_v1alpha1_plugin_proto_rawDescGZIP(), []int{0}
}

type ListTrustZonesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrustZones []*v1alpha1.TrustZone `protobuf:"bytes,1,rep,name=trust_zones,json=trustZones,proto3" json:"trust_zones,omitempty"`
}

func (x *ListTrustZonesResponse) Reset() {
	*x = ListTrustZonesResponse{}
	mi := &file_proto_cofidectl_plugin_v1alpha1_plugin_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTrustZonesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTrustZonesResponse) ProtoMessage() {}

func (x *ListTrustZonesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cofidectl_plugin_v1alpha1_plugin_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTrustZonesResponse.ProtoReflect.Descriptor instead.
func (*ListTrustZonesResponse) Descriptor() ([]byte, []int) {
	return file_proto_cofidectl_plugin_v1alpha1_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *ListTrustZonesResponse) GetTrustZones() []*v1alpha1.TrustZone {
	if x != nil {
		return x.TrustZones
	}
	return nil
}

var File_proto_cofidectl_plugin_v1alpha1_plugin_proto protoreflect.FileDescriptor

var file_proto_cofidectl_plugin_v1alpha1_plugin_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x66, 0x69, 0x64, 0x65, 0x63, 0x74,
	0x6c, 0x5f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x66, 0x69, 0x64, 0x65, 0x63, 0x74, 0x6c, 0x5f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a,
	0x2a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x7a, 0x6f, 0x6e,
	0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74,
	0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a, 0x15, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x72, 0x75, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x5f, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x75, 0x73,
	0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45,
	0x0a, 0x0b, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x72, 0x75, 0x73,
	0x74, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x54, 0x72, 0x75, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x0a, 0x74, 0x72, 0x75, 0x73, 0x74,
	0x5a, 0x6f, 0x6e, 0x65, 0x73, 0x32, 0x9d, 0x01, 0x0a, 0x17, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x81, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x75, 0x73, 0x74, 0x5a,
	0x6f, 0x6e, 0x65, 0x73, 0x12, 0x36, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x66,
	0x69, 0x64, 0x65, 0x63, 0x74, 0x6c, 0x5f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x75, 0x73, 0x74,
	0x5a, 0x6f, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x66, 0x69, 0x64, 0x65, 0x63, 0x74, 0x6c, 0x5f, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x72, 0x75, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x66, 0x69, 0x64, 0x65, 0x2f, 0x63, 0x6f, 0x66, 0x69, 0x64,
	0x65, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x66, 0x69, 0x64, 0x65, 0x63, 0x74, 0x6c,
	0x5f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_cofidectl_plugin_v1alpha1_plugin_proto_rawDescOnce sync.Once
	file_proto_cofidectl_plugin_v1alpha1_plugin_proto_rawDescData = file_proto_cofidectl_plugin_v1alpha1_plugin_proto_rawDesc
)

func file_proto_cofidectl_plugin_v1alpha1_plugin_proto_rawDescGZIP() []byte {
	file_proto_cofidectl_plugin_v1alpha1_plugin_proto_rawDescOnce.Do(func() {
		file_proto_cofidectl_plugin_v1alpha1_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cofidectl_plugin_v1alpha1_plugin_proto_rawDescData)
	})
	return file_proto_cofidectl_plugin_v1alpha1_plugin_proto_rawDescData
}

var file_proto_cofidectl_plugin_v1alpha1_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_cofidectl_plugin_v1alpha1_plugin_proto_goTypes = []any{
	(*ListTrustZonesRequest)(nil),  // 0: proto.cofidectl_plugin.v1alpha1.ListTrustZonesRequest
	(*ListTrustZonesResponse)(nil), // 1: proto.cofidectl_plugin.v1alpha1.ListTrustZonesResponse
	(*v1alpha1.TrustZone)(nil),     // 2: proto.trust_zone.v1alpha1.TrustZone
}
var file_proto_cofidectl_plugin_v1alpha1_plugin_proto_depIdxs = []int32{
	2, // 0: proto.cofidectl_plugin.v1alpha1.ListTrustZonesResponse.trust_zones:type_name -> proto.trust_zone.v1alpha1.TrustZone
	0, // 1: proto.cofidectl_plugin.v1alpha1.DataSourcePluginService.ListTrustZones:input_type -> proto.cofidectl_plugin.v1alpha1.ListTrustZonesRequest
	1, // 2: proto.cofidectl_plugin.v1alpha1.DataSourcePluginService.ListTrustZones:output_type -> proto.cofidectl_plugin.v1alpha1.ListTrustZonesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_cofidectl_plugin_v1alpha1_plugin_proto_init() }
func file_proto_cofidectl_plugin_v1alpha1_plugin_proto_init() {
	if File_proto_cofidectl_plugin_v1alpha1_plugin_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_cofidectl_plugin_v1alpha1_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cofidectl_plugin_v1alpha1_plugin_proto_goTypes,
		DependencyIndexes: file_proto_cofidectl_plugin_v1alpha1_plugin_proto_depIdxs,
		MessageInfos:      file_proto_cofidectl_plugin_v1alpha1_plugin_proto_msgTypes,
	}.Build()
	File_proto_cofidectl_plugin_v1alpha1_plugin_proto = out.File
	file_proto_cofidectl_plugin_v1alpha1_plugin_proto_rawDesc = nil
	file_proto_cofidectl_plugin_v1alpha1_plugin_proto_goTypes = nil
	file_proto_cofidectl_plugin_v1alpha1_plugin_proto_depIdxs = nil
}
