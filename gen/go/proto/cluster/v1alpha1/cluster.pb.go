// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: proto/cluster/v1alpha1/cluster.proto

package v1alpha1

import (
	v1alpha1 "github.com/cofide/cofide-api-sdk/gen/go/proto/trust_provider/v1alpha1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Cluster struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Id    *string                `protobuf:"bytes,8,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Name  *string                `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	OrgId *string                `protobuf:"bytes,9,opt,name=org_id,json=orgId,proto3,oneof" json:"org_id,omitempty"`
	// DEPRECATED: replaced by trust_zone_id.
	TrustZone         *string                 `protobuf:"bytes,2,opt,name=trust_zone,json=trustZone,proto3,oneof" json:"trust_zone,omitempty"`
	TrustZoneId       *string                 `protobuf:"bytes,10,opt,name=trust_zone_id,json=trustZoneId,proto3,oneof" json:"trust_zone_id,omitempty"`
	KubernetesContext *string                 `protobuf:"bytes,3,opt,name=kubernetes_context,json=kubernetesContext,proto3,oneof" json:"kubernetes_context,omitempty"`
	TrustProvider     *v1alpha1.TrustProvider `protobuf:"bytes,4,opt,name=trust_provider,json=trustProvider,proto3,oneof" json:"trust_provider,omitempty"`
	ExtraHelmValues   *structpb.Struct        `protobuf:"bytes,5,opt,name=extra_helm_values,json=extraHelmValues,proto3,oneof" json:"extra_helm_values,omitempty"`
	Profile           *string                 `protobuf:"bytes,6,opt,name=profile,proto3,oneof" json:"profile,omitempty"`
	ExternalServer    *bool                   `protobuf:"varint,7,opt,name=external_server,json=externalServer,proto3,oneof" json:"external_server,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	mi := &file_proto_cluster_v1alpha1_cluster_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cluster_v1alpha1_cluster_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_proto_cluster_v1alpha1_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *Cluster) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *Cluster) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Cluster) GetOrgId() string {
	if x != nil && x.OrgId != nil {
		return *x.OrgId
	}
	return ""
}

func (x *Cluster) GetTrustZone() string {
	if x != nil && x.TrustZone != nil {
		return *x.TrustZone
	}
	return ""
}

func (x *Cluster) GetTrustZoneId() string {
	if x != nil && x.TrustZoneId != nil {
		return *x.TrustZoneId
	}
	return ""
}

func (x *Cluster) GetKubernetesContext() string {
	if x != nil && x.KubernetesContext != nil {
		return *x.KubernetesContext
	}
	return ""
}

func (x *Cluster) GetTrustProvider() *v1alpha1.TrustProvider {
	if x != nil {
		return x.TrustProvider
	}
	return nil
}

func (x *Cluster) GetExtraHelmValues() *structpb.Struct {
	if x != nil {
		return x.ExtraHelmValues
	}
	return nil
}

func (x *Cluster) GetProfile() string {
	if x != nil && x.Profile != nil {
		return *x.Profile
	}
	return ""
}

func (x *Cluster) GetExternalServer() bool {
	if x != nil && x.ExternalServer != nil {
		return *x.ExternalServer
	}
	return false
}

var File_proto_cluster_v1alpha1_cluster_proto protoreflect.FileDescriptor

var file_proto_cluster_v1alpha1_cluster_proto_rawDesc = string([]byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x74, 0x72, 0x75, 0x73,
	0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe1, 0x04, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x13, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x06, 0x6f, 0x72,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x6f, 0x72,
	0x67, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f,
	0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x09, 0x74, 0x72,
	0x75, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0d, 0x74, 0x72,
	0x75, 0x73, 0x74, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x04, 0x52, 0x0b, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x12, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x05, 0x52, 0x11, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x88, 0x01, 0x01, 0x12, 0x58, 0x0a, 0x0e, 0x74, 0x72, 0x75, 0x73, 0x74,
	0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x54, 0x72, 0x75, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x48, 0x06, 0x52,
	0x0d, 0x74, 0x72, 0x75, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x88, 0x01,
	0x01, 0x12, 0x48, 0x0a, 0x11, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x68, 0x65, 0x6c, 0x6d, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x48, 0x07, 0x52, 0x0f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x48, 0x65,
	0x6c, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x09, 0x52, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x72, 0x67,
	0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x7a, 0x6f,
	0x6e, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x7a, 0x6f, 0x6e,
	0x65, 0x5f, 0x69, 0x64, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x42, 0x11, 0x0a, 0x0f, 0x5f,
	0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x42, 0x14,
	0x0a, 0x12, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x68, 0x65, 0x6c, 0x6d, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x42, 0x12, 0x0a, 0x10, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x66, 0x69, 0x64, 0x65, 0x2f, 0x63, 0x6f, 0x66, 0x69, 0x64, 0x65,
	0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_cluster_v1alpha1_cluster_proto_rawDescOnce sync.Once
	file_proto_cluster_v1alpha1_cluster_proto_rawDescData []byte
)

func file_proto_cluster_v1alpha1_cluster_proto_rawDescGZIP() []byte {
	file_proto_cluster_v1alpha1_cluster_proto_rawDescOnce.Do(func() {
		file_proto_cluster_v1alpha1_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_cluster_v1alpha1_cluster_proto_rawDesc), len(file_proto_cluster_v1alpha1_cluster_proto_rawDesc)))
	})
	return file_proto_cluster_v1alpha1_cluster_proto_rawDescData
}

var file_proto_cluster_v1alpha1_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_cluster_v1alpha1_cluster_proto_goTypes = []any{
	(*Cluster)(nil),                // 0: proto.cluster.v1alpha1.Cluster
	(*v1alpha1.TrustProvider)(nil), // 1: proto.trust_provider.v1alpha1.TrustProvider
	(*structpb.Struct)(nil),        // 2: google.protobuf.Struct
}
var file_proto_cluster_v1alpha1_cluster_proto_depIdxs = []int32{
	1, // 0: proto.cluster.v1alpha1.Cluster.trust_provider:type_name -> proto.trust_provider.v1alpha1.TrustProvider
	2, // 1: proto.cluster.v1alpha1.Cluster.extra_helm_values:type_name -> google.protobuf.Struct
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_cluster_v1alpha1_cluster_proto_init() }
func file_proto_cluster_v1alpha1_cluster_proto_init() {
	if File_proto_cluster_v1alpha1_cluster_proto != nil {
		return
	}
	file_proto_cluster_v1alpha1_cluster_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_cluster_v1alpha1_cluster_proto_rawDesc), len(file_proto_cluster_v1alpha1_cluster_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_cluster_v1alpha1_cluster_proto_goTypes,
		DependencyIndexes: file_proto_cluster_v1alpha1_cluster_proto_depIdxs,
		MessageInfos:      file_proto_cluster_v1alpha1_cluster_proto_msgTypes,
	}.Build()
	File_proto_cluster_v1alpha1_cluster_proto = out.File
	file_proto_cluster_v1alpha1_cluster_proto_goTypes = nil
	file_proto_cluster_v1alpha1_cluster_proto_depIdxs = nil
}
