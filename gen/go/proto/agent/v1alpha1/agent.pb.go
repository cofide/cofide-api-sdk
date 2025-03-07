// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: proto/agent/v1alpha1/agent.proto

package v1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type AgentStatusCode int32

const (
	AgentStatusCode_AGENT_STATUS_CODE_UNSPECIFIED AgentStatusCode = 0
	AgentStatusCode_AGENT_STATUS_CODE_RUNNING     AgentStatusCode = 1
	AgentStatusCode_AGENT_STATUS_CODE_STOPPED     AgentStatusCode = 2
	AgentStatusCode_AGENT_STATUS_CODE_ERROR       AgentStatusCode = 3
)

// Enum value maps for AgentStatusCode.
var (
	AgentStatusCode_name = map[int32]string{
		0: "AGENT_STATUS_CODE_UNSPECIFIED",
		1: "AGENT_STATUS_CODE_RUNNING",
		2: "AGENT_STATUS_CODE_STOPPED",
		3: "AGENT_STATUS_CODE_ERROR",
	}
	AgentStatusCode_value = map[string]int32{
		"AGENT_STATUS_CODE_UNSPECIFIED": 0,
		"AGENT_STATUS_CODE_RUNNING":     1,
		"AGENT_STATUS_CODE_STOPPED":     2,
		"AGENT_STATUS_CODE_ERROR":       3,
	}
)

func (x AgentStatusCode) Enum() *AgentStatusCode {
	p := new(AgentStatusCode)
	*p = x
	return p
}

func (x AgentStatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgentStatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_agent_v1alpha1_agent_proto_enumTypes[0].Descriptor()
}

func (AgentStatusCode) Type() protoreflect.EnumType {
	return &file_proto_agent_v1alpha1_agent_proto_enumTypes[0]
}

func (x AgentStatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgentStatusCode.Descriptor instead.
func (AgentStatusCode) EnumDescriptor() ([]byte, []int) {
	return file_proto_agent_v1alpha1_agent_proto_rawDescGZIP(), []int{0}
}

type Agent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *string                `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	ClusterId     *string                `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3,oneof" json:"cluster_id,omitempty"`
	TrustZoneId   *string                `protobuf:"bytes,3,opt,name=trust_zone_id,json=trustZoneId,proto3,oneof" json:"trust_zone_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Agent) Reset() {
	*x = Agent{}
	mi := &file_proto_agent_v1alpha1_agent_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Agent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Agent) ProtoMessage() {}

func (x *Agent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_agent_v1alpha1_agent_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Agent.ProtoReflect.Descriptor instead.
func (*Agent) Descriptor() ([]byte, []int) {
	return file_proto_agent_v1alpha1_agent_proto_rawDescGZIP(), []int{0}
}

func (x *Agent) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *Agent) GetClusterId() string {
	if x != nil && x.ClusterId != nil {
		return *x.ClusterId
	}
	return ""
}

func (x *Agent) GetTrustZoneId() string {
	if x != nil && x.TrustZoneId != nil {
		return *x.TrustZoneId
	}
	return ""
}

type AgentStatus struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        *AgentStatusCode       `protobuf:"varint,1,opt,name=status,proto3,enum=proto.agent.v1alpha1.AgentStatusCode,oneof" json:"status,omitempty"`
	StatusMessage *string                `protobuf:"bytes,2,opt,name=status_message,json=statusMessage,proto3,oneof" json:"status_message,omitempty"`
	LastUpdated   *int64                 `protobuf:"varint,3,opt,name=last_updated,json=lastUpdated,proto3,oneof" json:"last_updated,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AgentStatus) Reset() {
	*x = AgentStatus{}
	mi := &file_proto_agent_v1alpha1_agent_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AgentStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentStatus) ProtoMessage() {}

func (x *AgentStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_agent_v1alpha1_agent_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentStatus.ProtoReflect.Descriptor instead.
func (*AgentStatus) Descriptor() ([]byte, []int) {
	return file_proto_agent_v1alpha1_agent_proto_rawDescGZIP(), []int{1}
}

func (x *AgentStatus) GetStatus() AgentStatusCode {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return AgentStatusCode_AGENT_STATUS_CODE_UNSPECIFIED
}

func (x *AgentStatus) GetStatusMessage() string {
	if x != nil && x.StatusMessage != nil {
		return *x.StatusMessage
	}
	return ""
}

func (x *AgentStatus) GetLastUpdated() int64 {
	if x != nil && x.LastUpdated != nil {
		return *x.LastUpdated
	}
	return 0
}

var File_proto_agent_v1alpha1_agent_proto protoreflect.FileDescriptor

var file_proto_agent_v1alpha1_agent_proto_rawDesc = string([]byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x91, 0x01, 0x0a, 0x05, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0d, 0x74,
	0x72, 0x75, 0x73, 0x74, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x02, 0x52, 0x0b, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x74,
	0x72, 0x75, 0x73, 0x74, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x22, 0xd4, 0x01, 0x0a,
	0x0b, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x42, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01,
	0x12, 0x2a, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x02, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x11, 0x0a, 0x0f, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x2a, 0x8f, 0x01, 0x0a, 0x0f, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x1d, 0x41, 0x47, 0x45, 0x4e, 0x54,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x41, 0x47,
	0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x41, 0x47, 0x45,
	0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53,
	0x54, 0x4f, 0x50, 0x50, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x47, 0x45, 0x4e,
	0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x03, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x66, 0x69, 0x64, 0x65, 0x2f, 0x63, 0x6f, 0x66, 0x69, 0x64,
	0x65, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_agent_v1alpha1_agent_proto_rawDescOnce sync.Once
	file_proto_agent_v1alpha1_agent_proto_rawDescData []byte
)

func file_proto_agent_v1alpha1_agent_proto_rawDescGZIP() []byte {
	file_proto_agent_v1alpha1_agent_proto_rawDescOnce.Do(func() {
		file_proto_agent_v1alpha1_agent_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_agent_v1alpha1_agent_proto_rawDesc), len(file_proto_agent_v1alpha1_agent_proto_rawDesc)))
	})
	return file_proto_agent_v1alpha1_agent_proto_rawDescData
}

var file_proto_agent_v1alpha1_agent_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_agent_v1alpha1_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_agent_v1alpha1_agent_proto_goTypes = []any{
	(AgentStatusCode)(0), // 0: proto.agent.v1alpha1.AgentStatusCode
	(*Agent)(nil),        // 1: proto.agent.v1alpha1.Agent
	(*AgentStatus)(nil),  // 2: proto.agent.v1alpha1.AgentStatus
}
var file_proto_agent_v1alpha1_agent_proto_depIdxs = []int32{
	0, // 0: proto.agent.v1alpha1.AgentStatus.status:type_name -> proto.agent.v1alpha1.AgentStatusCode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_agent_v1alpha1_agent_proto_init() }
func file_proto_agent_v1alpha1_agent_proto_init() {
	if File_proto_agent_v1alpha1_agent_proto != nil {
		return
	}
	file_proto_agent_v1alpha1_agent_proto_msgTypes[0].OneofWrappers = []any{}
	file_proto_agent_v1alpha1_agent_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_agent_v1alpha1_agent_proto_rawDesc), len(file_proto_agent_v1alpha1_agent_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_agent_v1alpha1_agent_proto_goTypes,
		DependencyIndexes: file_proto_agent_v1alpha1_agent_proto_depIdxs,
		EnumInfos:         file_proto_agent_v1alpha1_agent_proto_enumTypes,
		MessageInfos:      file_proto_agent_v1alpha1_agent_proto_msgTypes,
	}.Build()
	File_proto_agent_v1alpha1_agent_proto = out.File
	file_proto_agent_v1alpha1_agent_proto_goTypes = nil
	file_proto_agent_v1alpha1_agent_proto_depIdxs = nil
}
