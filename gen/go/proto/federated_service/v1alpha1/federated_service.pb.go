// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: proto/federated_service/v1alpha1/federated_service.proto

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

type FederatedService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            string            `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ClusterName          string            `protobuf:"bytes,4,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	TrustDomain          string            `protobuf:"bytes,5,opt,name=trust_domain,json=trustDomain,proto3" json:"trust_domain,omitempty"`
	GatewayEntries       []*GatewayEntry   `protobuf:"bytes,6,rep,name=gateway_entries,json=gatewayEntries,proto3" json:"gateway_entries,omitempty"`
	WorkloadLabels       map[string]string `protobuf:"bytes,7,rep,name=workload_labels,json=workloadLabels,proto3" json:"workload_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ExportedTrustDomains []string          `protobuf:"bytes,8,rep,name=exported_trust_domains,json=exportedTrustDomains,proto3" json:"exported_trust_domains,omitempty"`
	Port                 uint32            `protobuf:"varint,9,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *FederatedService) Reset() {
	*x = FederatedService{}
	mi := &file_proto_federated_service_v1alpha1_federated_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FederatedService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FederatedService) ProtoMessage() {}

func (x *FederatedService) ProtoReflect() protoreflect.Message {
	mi := &file_proto_federated_service_v1alpha1_federated_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FederatedService.ProtoReflect.Descriptor instead.
func (*FederatedService) Descriptor() ([]byte, []int) {
	return file_proto_federated_service_v1alpha1_federated_service_proto_rawDescGZIP(), []int{0}
}

func (x *FederatedService) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FederatedService) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FederatedService) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *FederatedService) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *FederatedService) GetTrustDomain() string {
	if x != nil {
		return x.TrustDomain
	}
	return ""
}

func (x *FederatedService) GetGatewayEntries() []*GatewayEntry {
	if x != nil {
		return x.GatewayEntries
	}
	return nil
}

func (x *FederatedService) GetWorkloadLabels() map[string]string {
	if x != nil {
		return x.WorkloadLabels
	}
	return nil
}

func (x *FederatedService) GetExportedTrustDomains() []string {
	if x != nil {
		return x.ExportedTrustDomains
	}
	return nil
}

func (x *FederatedService) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type GatewayEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Ip       string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *GatewayEntry) Reset() {
	*x = GatewayEntry{}
	mi := &file_proto_federated_service_v1alpha1_federated_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GatewayEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayEntry) ProtoMessage() {}

func (x *GatewayEntry) ProtoReflect() protoreflect.Message {
	mi := &file_proto_federated_service_v1alpha1_federated_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayEntry.ProtoReflect.Descriptor instead.
func (*GatewayEntry) Descriptor() ([]byte, []int) {
	return file_proto_federated_service_v1alpha1_federated_service_proto_rawDescGZIP(), []int{1}
}

func (x *GatewayEntry) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *GatewayEntry) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GatewayEntry) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

var File_proto_federated_service_v1alpha1_federated_service_proto protoreflect.FileDescriptor

var file_proto_federated_service_v1alpha1_federated_service_proto_rawDesc = []byte{
	0x0a, 0x38, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0xf1, 0x03, 0x0a,
	0x10, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72,
	0x75, 0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x57, 0x0a, 0x0f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x6f, 0x0a, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f,
	0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x14, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x54, 0x72, 0x75,
	0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x41, 0x0a,
	0x13, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x4e, 0x0a, 0x0c, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70,
	0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x66, 0x69, 0x64, 0x65, 0x2f, 0x63, 0x6f, 0x66, 0x69, 0x64, 0x65, 0x2d, 0x61, 0x70, 0x69,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_federated_service_v1alpha1_federated_service_proto_rawDescOnce sync.Once
	file_proto_federated_service_v1alpha1_federated_service_proto_rawDescData = file_proto_federated_service_v1alpha1_federated_service_proto_rawDesc
)

func file_proto_federated_service_v1alpha1_federated_service_proto_rawDescGZIP() []byte {
	file_proto_federated_service_v1alpha1_federated_service_proto_rawDescOnce.Do(func() {
		file_proto_federated_service_v1alpha1_federated_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_federated_service_v1alpha1_federated_service_proto_rawDescData)
	})
	return file_proto_federated_service_v1alpha1_federated_service_proto_rawDescData
}

var file_proto_federated_service_v1alpha1_federated_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_federated_service_v1alpha1_federated_service_proto_goTypes = []any{
	(*FederatedService)(nil), // 0: proto.federated_service.v1alpha1.FederatedService
	(*GatewayEntry)(nil),     // 1: proto.federated_service.v1alpha1.GatewayEntry
	nil,                      // 2: proto.federated_service.v1alpha1.FederatedService.WorkloadLabelsEntry
}
var file_proto_federated_service_v1alpha1_federated_service_proto_depIdxs = []int32{
	1, // 0: proto.federated_service.v1alpha1.FederatedService.gateway_entries:type_name -> proto.federated_service.v1alpha1.GatewayEntry
	2, // 1: proto.federated_service.v1alpha1.FederatedService.workload_labels:type_name -> proto.federated_service.v1alpha1.FederatedService.WorkloadLabelsEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_federated_service_v1alpha1_federated_service_proto_init() }
func file_proto_federated_service_v1alpha1_federated_service_proto_init() {
	if File_proto_federated_service_v1alpha1_federated_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_federated_service_v1alpha1_federated_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_federated_service_v1alpha1_federated_service_proto_goTypes,
		DependencyIndexes: file_proto_federated_service_v1alpha1_federated_service_proto_depIdxs,
		MessageInfos:      file_proto_federated_service_v1alpha1_federated_service_proto_msgTypes,
	}.Build()
	File_proto_federated_service_v1alpha1_federated_service_proto = out.File
	file_proto_federated_service_v1alpha1_federated_service_proto_rawDesc = nil
	file_proto_federated_service_v1alpha1_federated_service_proto_goTypes = nil
	file_proto_federated_service_v1alpha1_federated_service_proto_depIdxs = nil
}
