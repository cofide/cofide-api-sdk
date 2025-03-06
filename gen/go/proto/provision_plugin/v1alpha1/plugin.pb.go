// Copyright 2024 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: proto/provision_plugin/v1alpha1/plugin.proto

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

type ValidateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValidateRequest) Reset() {
	*x = ValidateRequest{}
	mi := &file_proto_provision_plugin_v1alpha1_plugin_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateRequest) ProtoMessage() {}

func (x *ValidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provision_plugin_v1alpha1_plugin_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateRequest.ProtoReflect.Descriptor instead.
func (*ValidateRequest) Descriptor() ([]byte, []int) {
	return file_proto_provision_plugin_v1alpha1_plugin_proto_rawDescGZIP(), []int{0}
}

type ValidateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValidateResponse) Reset() {
	*x = ValidateResponse{}
	mi := &file_proto_provision_plugin_v1alpha1_plugin_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateResponse) ProtoMessage() {}

func (x *ValidateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provision_plugin_v1alpha1_plugin_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateResponse.ProtoReflect.Descriptor instead.
func (*ValidateResponse) Descriptor() ([]byte, []int) {
	return file_proto_provision_plugin_v1alpha1_plugin_proto_rawDescGZIP(), []int{1}
}

type DeployRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DataSource    *uint32                `protobuf:"varint,1,opt,name=data_source,json=dataSource,proto3,oneof" json:"data_source,omitempty"`
	KubeCfgFile   *string                `protobuf:"bytes,2,opt,name=kube_cfg_file,json=kubeCfgFile,proto3,oneof" json:"kube_cfg_file,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeployRequest) Reset() {
	*x = DeployRequest{}
	mi := &file_proto_provision_plugin_v1alpha1_plugin_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeployRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployRequest) ProtoMessage() {}

func (x *DeployRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provision_plugin_v1alpha1_plugin_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployRequest.ProtoReflect.Descriptor instead.
func (*DeployRequest) Descriptor() ([]byte, []int) {
	return file_proto_provision_plugin_v1alpha1_plugin_proto_rawDescGZIP(), []int{2}
}

func (x *DeployRequest) GetDataSource() uint32 {
	if x != nil && x.DataSource != nil {
		return *x.DataSource
	}
	return 0
}

func (x *DeployRequest) GetKubeCfgFile() string {
	if x != nil && x.KubeCfgFile != nil {
		return *x.KubeCfgFile
	}
	return ""
}

type DeployResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        *Status                `protobuf:"bytes,1,opt,name=status,proto3,oneof" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeployResponse) Reset() {
	*x = DeployResponse{}
	mi := &file_proto_provision_plugin_v1alpha1_plugin_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeployResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployResponse) ProtoMessage() {}

func (x *DeployResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provision_plugin_v1alpha1_plugin_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployResponse.ProtoReflect.Descriptor instead.
func (*DeployResponse) Descriptor() ([]byte, []int) {
	return file_proto_provision_plugin_v1alpha1_plugin_proto_rawDescGZIP(), []int{3}
}

func (x *DeployResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type TearDownRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DataSource    *uint32                `protobuf:"varint,1,opt,name=data_source,json=dataSource,proto3,oneof" json:"data_source,omitempty"`
	KubeCfgFile   *string                `protobuf:"bytes,2,opt,name=kube_cfg_file,json=kubeCfgFile,proto3,oneof" json:"kube_cfg_file,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TearDownRequest) Reset() {
	*x = TearDownRequest{}
	mi := &file_proto_provision_plugin_v1alpha1_plugin_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TearDownRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TearDownRequest) ProtoMessage() {}

func (x *TearDownRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provision_plugin_v1alpha1_plugin_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TearDownRequest.ProtoReflect.Descriptor instead.
func (*TearDownRequest) Descriptor() ([]byte, []int) {
	return file_proto_provision_plugin_v1alpha1_plugin_proto_rawDescGZIP(), []int{4}
}

func (x *TearDownRequest) GetDataSource() uint32 {
	if x != nil && x.DataSource != nil {
		return *x.DataSource
	}
	return 0
}

func (x *TearDownRequest) GetKubeCfgFile() string {
	if x != nil && x.KubeCfgFile != nil {
		return *x.KubeCfgFile
	}
	return ""
}

type TearDownResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        *Status                `protobuf:"bytes,1,opt,name=status,proto3,oneof" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TearDownResponse) Reset() {
	*x = TearDownResponse{}
	mi := &file_proto_provision_plugin_v1alpha1_plugin_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TearDownResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TearDownResponse) ProtoMessage() {}

func (x *TearDownResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provision_plugin_v1alpha1_plugin_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TearDownResponse.ProtoReflect.Descriptor instead.
func (*TearDownResponse) Descriptor() ([]byte, []int) {
	return file_proto_provision_plugin_v1alpha1_plugin_proto_rawDescGZIP(), []int{5}
}

func (x *TearDownResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type Status struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Stage         *string                `protobuf:"bytes,1,opt,name=stage,proto3,oneof" json:"stage,omitempty"`
	Message       *string                `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
	Done          *bool                  `protobuf:"varint,3,opt,name=done,proto3,oneof" json:"done,omitempty"`
	Error         *string                `protobuf:"bytes,4,opt,name=error,proto3,oneof" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Status) Reset() {
	*x = Status{}
	mi := &file_proto_provision_plugin_v1alpha1_plugin_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provision_plugin_v1alpha1_plugin_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_proto_provision_plugin_v1alpha1_plugin_proto_rawDescGZIP(), []int{6}
}

func (x *Status) GetStage() string {
	if x != nil && x.Stage != nil {
		return *x.Stage
	}
	return ""
}

func (x *Status) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (x *Status) GetDone() bool {
	if x != nil && x.Done != nil {
		return *x.Done
	}
	return false
}

func (x *Status) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

var File_proto_provision_plugin_v1alpha1_plugin_proto protoreflect.FileDescriptor

var file_proto_provision_plugin_v1alpha1_plugin_proto_rawDesc = string([]byte{
	0x0a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22,
	0x11, 0x0a, 0x0f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x80, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52,
	0x0a, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x27,
	0x0a, 0x0d, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x63, 0x66, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x6b, 0x75, 0x62, 0x65, 0x43, 0x66, 0x67,
	0x46, 0x69, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6b, 0x75, 0x62, 0x65,
	0x5f, 0x63, 0x66, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x61, 0x0a, 0x0e, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01,
	0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x82, 0x01, 0x0a,
	0x0f, 0x54, 0x65, 0x61, 0x72, 0x44, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0d, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x63,
	0x66, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x0b, 0x6b, 0x75, 0x62, 0x65, 0x43, 0x66, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42,
	0x10, 0x0a, 0x0e, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x63, 0x66, 0x67, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x22, 0x63, 0x0a, 0x10, 0x54, 0x65, 0x61, 0x72, 0x44, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x9f, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x19, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x64,
	0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x04, 0x64, 0x6f, 0x6e,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x6f, 0x6e, 0x65, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xe9, 0x02, 0x0a, 0x16, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x6f, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b, 0x0a, 0x06, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x12, 0x2e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30,
	0x01, 0x12, 0x71, 0x0a, 0x08, 0x54, 0x65, 0x61, 0x72, 0x44, 0x6f, 0x77, 0x6e, 0x12, 0x30, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x54, 0x65, 0x61, 0x72, 0x44, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x54, 0x65, 0x61, 0x72, 0x44, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x30, 0x01, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x66, 0x69, 0x64, 0x65, 0x2f, 0x63, 0x6f, 0x66, 0x69, 0x64, 0x65,
	0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_provision_plugin_v1alpha1_plugin_proto_rawDescOnce sync.Once
	file_proto_provision_plugin_v1alpha1_plugin_proto_rawDescData []byte
)

func file_proto_provision_plugin_v1alpha1_plugin_proto_rawDescGZIP() []byte {
	file_proto_provision_plugin_v1alpha1_plugin_proto_rawDescOnce.Do(func() {
		file_proto_provision_plugin_v1alpha1_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_provision_plugin_v1alpha1_plugin_proto_rawDesc), len(file_proto_provision_plugin_v1alpha1_plugin_proto_rawDesc)))
	})
	return file_proto_provision_plugin_v1alpha1_plugin_proto_rawDescData
}

var file_proto_provision_plugin_v1alpha1_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_provision_plugin_v1alpha1_plugin_proto_goTypes = []any{
	(*ValidateRequest)(nil),  // 0: proto.provision_plugin.v1alpha1.ValidateRequest
	(*ValidateResponse)(nil), // 1: proto.provision_plugin.v1alpha1.ValidateResponse
	(*DeployRequest)(nil),    // 2: proto.provision_plugin.v1alpha1.DeployRequest
	(*DeployResponse)(nil),   // 3: proto.provision_plugin.v1alpha1.DeployResponse
	(*TearDownRequest)(nil),  // 4: proto.provision_plugin.v1alpha1.TearDownRequest
	(*TearDownResponse)(nil), // 5: proto.provision_plugin.v1alpha1.TearDownResponse
	(*Status)(nil),           // 6: proto.provision_plugin.v1alpha1.Status
}
var file_proto_provision_plugin_v1alpha1_plugin_proto_depIdxs = []int32{
	6, // 0: proto.provision_plugin.v1alpha1.DeployResponse.status:type_name -> proto.provision_plugin.v1alpha1.Status
	6, // 1: proto.provision_plugin.v1alpha1.TearDownResponse.status:type_name -> proto.provision_plugin.v1alpha1.Status
	0, // 2: proto.provision_plugin.v1alpha1.ProvisionPluginService.Validate:input_type -> proto.provision_plugin.v1alpha1.ValidateRequest
	2, // 3: proto.provision_plugin.v1alpha1.ProvisionPluginService.Deploy:input_type -> proto.provision_plugin.v1alpha1.DeployRequest
	4, // 4: proto.provision_plugin.v1alpha1.ProvisionPluginService.TearDown:input_type -> proto.provision_plugin.v1alpha1.TearDownRequest
	1, // 5: proto.provision_plugin.v1alpha1.ProvisionPluginService.Validate:output_type -> proto.provision_plugin.v1alpha1.ValidateResponse
	3, // 6: proto.provision_plugin.v1alpha1.ProvisionPluginService.Deploy:output_type -> proto.provision_plugin.v1alpha1.DeployResponse
	5, // 7: proto.provision_plugin.v1alpha1.ProvisionPluginService.TearDown:output_type -> proto.provision_plugin.v1alpha1.TearDownResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_provision_plugin_v1alpha1_plugin_proto_init() }
func file_proto_provision_plugin_v1alpha1_plugin_proto_init() {
	if File_proto_provision_plugin_v1alpha1_plugin_proto != nil {
		return
	}
	file_proto_provision_plugin_v1alpha1_plugin_proto_msgTypes[2].OneofWrappers = []any{}
	file_proto_provision_plugin_v1alpha1_plugin_proto_msgTypes[3].OneofWrappers = []any{}
	file_proto_provision_plugin_v1alpha1_plugin_proto_msgTypes[4].OneofWrappers = []any{}
	file_proto_provision_plugin_v1alpha1_plugin_proto_msgTypes[5].OneofWrappers = []any{}
	file_proto_provision_plugin_v1alpha1_plugin_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_provision_plugin_v1alpha1_plugin_proto_rawDesc), len(file_proto_provision_plugin_v1alpha1_plugin_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_provision_plugin_v1alpha1_plugin_proto_goTypes,
		DependencyIndexes: file_proto_provision_plugin_v1alpha1_plugin_proto_depIdxs,
		MessageInfos:      file_proto_provision_plugin_v1alpha1_plugin_proto_msgTypes,
	}.Build()
	File_proto_provision_plugin_v1alpha1_plugin_proto = out.File
	file_proto_provision_plugin_v1alpha1_plugin_proto_goTypes = nil
	file_proto_provision_plugin_v1alpha1_plugin_proto_depIdxs = nil
}
