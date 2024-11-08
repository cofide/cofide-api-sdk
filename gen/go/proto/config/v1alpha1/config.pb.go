// Copyright 2024 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: proto/config/v1alpha1/config.proto

package v1alpha1

import (
	v1alpha11 "github.com/cofide/cofide-api-sdk/gen/go/proto/attestation_policy/v1alpha1"
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

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataSource          *string                        `protobuf:"bytes,1,opt,name=data_source,json=dataSource,proto3,oneof" json:"data_source,omitempty"`
	TrustZones          []*v1alpha1.TrustZone          `protobuf:"bytes,2,rep,name=trust_zones,json=trustZones,proto3" json:"trust_zones,omitempty"`
	AttestationPolicies []*v1alpha11.AttestationPolicy `protobuf:"bytes,3,rep,name=attestation_policies,json=attestationPolicies,proto3" json:"attestation_policies,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_proto_config_v1alpha1_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_proto_config_v1alpha1_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_proto_config_v1alpha1_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetDataSource() string {
	if x != nil && x.DataSource != nil {
		return *x.DataSource
	}
	return ""
}

func (x *Config) GetTrustZones() []*v1alpha1.TrustZone {
	if x != nil {
		return x.TrustZones
	}
	return nil
}

func (x *Config) GetAttestationPolicies() []*v1alpha11.AttestationPolicy {
	if x != nil {
		return x.AttestationPolicies
	}
	return nil
}

var File_proto_config_v1alpha1_config_proto protoreflect.FileDescriptor

var file_proto_config_v1alpha1_config_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x3a, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74,
	0x72, 0x75, 0x73, 0x74, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xee, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x24,
	0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x45, 0x0a, 0x0b, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x7a, 0x6f,
	0x6e, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x72, 0x75, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x52,
	0x0a, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x73, 0x12, 0x67, 0x0a, 0x14, 0x61,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x74,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x13, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x66, 0x69, 0x64, 0x65, 0x2f, 0x63, 0x6f, 0x66, 0x69, 0x64, 0x65,
	0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_config_v1alpha1_config_proto_rawDescOnce sync.Once
	file_proto_config_v1alpha1_config_proto_rawDescData = file_proto_config_v1alpha1_config_proto_rawDesc
)

func file_proto_config_v1alpha1_config_proto_rawDescGZIP() []byte {
	file_proto_config_v1alpha1_config_proto_rawDescOnce.Do(func() {
		file_proto_config_v1alpha1_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_config_v1alpha1_config_proto_rawDescData)
	})
	return file_proto_config_v1alpha1_config_proto_rawDescData
}

var file_proto_config_v1alpha1_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_config_v1alpha1_config_proto_goTypes = []any{
	(*Config)(nil),                      // 0: proto.config.v1alpha1.Config
	(*v1alpha1.TrustZone)(nil),          // 1: proto.trust_zone.v1alpha1.TrustZone
	(*v1alpha11.AttestationPolicy)(nil), // 2: proto.attestation_policy.v1alpha1.AttestationPolicy
}
var file_proto_config_v1alpha1_config_proto_depIdxs = []int32{
	1, // 0: proto.config.v1alpha1.Config.trust_zones:type_name -> proto.trust_zone.v1alpha1.TrustZone
	2, // 1: proto.config.v1alpha1.Config.attestation_policies:type_name -> proto.attestation_policy.v1alpha1.AttestationPolicy
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_config_v1alpha1_config_proto_init() }
func file_proto_config_v1alpha1_config_proto_init() {
	if File_proto_config_v1alpha1_config_proto != nil {
		return
	}
	file_proto_config_v1alpha1_config_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_config_v1alpha1_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_config_v1alpha1_config_proto_goTypes,
		DependencyIndexes: file_proto_config_v1alpha1_config_proto_depIdxs,
		MessageInfos:      file_proto_config_v1alpha1_config_proto_msgTypes,
	}.Build()
	File_proto_config_v1alpha1_config_proto = out.File
	file_proto_config_v1alpha1_config_proto_rawDesc = nil
	file_proto_config_v1alpha1_config_proto_goTypes = nil
	file_proto_config_v1alpha1_config_proto_depIdxs = nil
}
