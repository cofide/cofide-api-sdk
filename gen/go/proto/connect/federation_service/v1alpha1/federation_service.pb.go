// Copyright 2025 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: proto/connect/federation_service/v1alpha1/federation_service.proto

package v1alpha1

import (
	v1alpha1 "github.com/cofide/cofide-api-sdk/gen/go/proto/federation/v1alpha1"
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

type CreateFederationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Federation    *v1alpha1.Federation   `protobuf:"bytes,1,opt,name=federation,proto3" json:"federation,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateFederationRequest) Reset() {
	*x = CreateFederationRequest{}
	mi := &file_proto_connect_federation_service_v1alpha1_federation_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateFederationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFederationRequest) ProtoMessage() {}

func (x *CreateFederationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_connect_federation_service_v1alpha1_federation_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFederationRequest.ProtoReflect.Descriptor instead.
func (*CreateFederationRequest) Descriptor() ([]byte, []int) {
	return file_proto_connect_federation_service_v1alpha1_federation_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFederationRequest) GetFederation() *v1alpha1.Federation {
	if x != nil {
		return x.Federation
	}
	return nil
}

type CreateFederationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Federation    *v1alpha1.Federation   `protobuf:"bytes,1,opt,name=federation,proto3" json:"federation,omitempty"`
	Success       bool                   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Message       *string                `protobuf:"bytes,3,opt,name=message,proto3,oneof" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateFederationResponse) Reset() {
	*x = CreateFederationResponse{}
	mi := &file_proto_connect_federation_service_v1alpha1_federation_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateFederationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFederationResponse) ProtoMessage() {}

func (x *CreateFederationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_connect_federation_service_v1alpha1_federation_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFederationResponse.ProtoReflect.Descriptor instead.
func (*CreateFederationResponse) Descriptor() ([]byte, []int) {
	return file_proto_connect_federation_service_v1alpha1_federation_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFederationResponse) GetFederation() *v1alpha1.Federation {
	if x != nil {
		return x.Federation
	}
	return nil
}

func (x *CreateFederationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateFederationResponse) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

type ListFederationsRequest struct {
	state         protoimpl.MessageState         `protogen:"open.v1"`
	Filter        *ListFederationsRequest_Filter `protobuf:"bytes,1,opt,name=filter,proto3,oneof" json:"filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListFederationsRequest) Reset() {
	*x = ListFederationsRequest{}
	mi := &file_proto_connect_federation_service_v1alpha1_federation_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFederationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFederationsRequest) ProtoMessage() {}

func (x *ListFederationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_connect_federation_service_v1alpha1_federation_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFederationsRequest.ProtoReflect.Descriptor instead.
func (*ListFederationsRequest) Descriptor() ([]byte, []int) {
	return file_proto_connect_federation_service_v1alpha1_federation_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListFederationsRequest) GetFilter() *ListFederationsRequest_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type ListFederationsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Federations   []*v1alpha1.Federation `protobuf:"bytes,1,rep,name=federations,proto3" json:"federations,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListFederationsResponse) Reset() {
	*x = ListFederationsResponse{}
	mi := &file_proto_connect_federation_service_v1alpha1_federation_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFederationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFederationsResponse) ProtoMessage() {}

func (x *ListFederationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_connect_federation_service_v1alpha1_federation_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFederationsResponse.ProtoReflect.Descriptor instead.
func (*ListFederationsResponse) Descriptor() ([]byte, []int) {
	return file_proto_connect_federation_service_v1alpha1_federation_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListFederationsResponse) GetFederations() []*v1alpha1.Federation {
	if x != nil {
		return x.Federations
	}
	return nil
}

type DeleteFederationRequest struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	TrustZoneId       string                 `protobuf:"bytes,1,opt,name=trust_zone_id,json=trustZoneId,proto3" json:"trust_zone_id,omitempty"`
	RemoteTrustZoneId string                 `protobuf:"bytes,2,opt,name=remote_trust_zone_id,json=remoteTrustZoneId,proto3" json:"remote_trust_zone_id,omitempty"`
	OrgId             string                 `protobuf:"bytes,3,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *DeleteFederationRequest) Reset() {
	*x = DeleteFederationRequest{}
	mi := &file_proto_connect_federation_service_v1alpha1_federation_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteFederationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFederationRequest) ProtoMessage() {}

func (x *DeleteFederationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_connect_federation_service_v1alpha1_federation_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFederationRequest.ProtoReflect.Descriptor instead.
func (*DeleteFederationRequest) Descriptor() ([]byte, []int) {
	return file_proto_connect_federation_service_v1alpha1_federation_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteFederationRequest) GetTrustZoneId() string {
	if x != nil {
		return x.TrustZoneId
	}
	return ""
}

func (x *DeleteFederationRequest) GetRemoteTrustZoneId() string {
	if x != nil {
		return x.RemoteTrustZoneId
	}
	return ""
}

func (x *DeleteFederationRequest) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

type DeleteFederationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       *string                `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteFederationResponse) Reset() {
	*x = DeleteFederationResponse{}
	mi := &file_proto_connect_federation_service_v1alpha1_federation_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteFederationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFederationResponse) ProtoMessage() {}

func (x *DeleteFederationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_connect_federation_service_v1alpha1_federation_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFederationResponse.ProtoReflect.Descriptor instead.
func (*DeleteFederationResponse) Descriptor() ([]byte, []int) {
	return file_proto_connect_federation_service_v1alpha1_federation_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteFederationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeleteFederationResponse) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

type ListFederationsRequest_Filter struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	OrgId             *string                `protobuf:"bytes,6,opt,name=org_id,json=orgId,proto3,oneof" json:"org_id,omitempty"`
	From              *string                `protobuf:"bytes,4,opt,name=from,proto3,oneof" json:"from,omitempty"`
	To                *string                `protobuf:"bytes,5,opt,name=to,proto3,oneof" json:"to,omitempty"`
	TrustZoneId       *string                `protobuf:"bytes,2,opt,name=trust_zone_id,json=trustZoneId,proto3,oneof" json:"trust_zone_id,omitempty"`
	RemoteTrustZoneId *string                `protobuf:"bytes,3,opt,name=remote_trust_zone_id,json=remoteTrustZoneId,proto3,oneof" json:"remote_trust_zone_id,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *ListFederationsRequest_Filter) Reset() {
	*x = ListFederationsRequest_Filter{}
	mi := &file_proto_connect_federation_service_v1alpha1_federation_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFederationsRequest_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFederationsRequest_Filter) ProtoMessage() {}

func (x *ListFederationsRequest_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_proto_connect_federation_service_v1alpha1_federation_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFederationsRequest_Filter.ProtoReflect.Descriptor instead.
func (*ListFederationsRequest_Filter) Descriptor() ([]byte, []int) {
	return file_proto_connect_federation_service_v1alpha1_federation_service_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ListFederationsRequest_Filter) GetOrgId() string {
	if x != nil && x.OrgId != nil {
		return *x.OrgId
	}
	return ""
}

func (x *ListFederationsRequest_Filter) GetFrom() string {
	if x != nil && x.From != nil {
		return *x.From
	}
	return ""
}

func (x *ListFederationsRequest_Filter) GetTo() string {
	if x != nil && x.To != nil {
		return *x.To
	}
	return ""
}

func (x *ListFederationsRequest_Filter) GetTrustZoneId() string {
	if x != nil && x.TrustZoneId != nil {
		return *x.TrustZoneId
	}
	return ""
}

func (x *ListFederationsRequest_Filter) GetRemoteTrustZoneId() string {
	if x != nil && x.RemoteTrustZoneId != nil {
		return *x.RemoteTrustZoneId
	}
	return ""
}

var File_proto_connect_federation_service_v1alpha1_federation_service_proto protoreflect.FileDescriptor

var file_proto_connect_federation_service_v1alpha1_federation_service_proto_rawDesc = string([]byte{
	0x0a, 0x42, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f,
	0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x66, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a,
	0x2a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x17, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0a, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0a, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa6, 0x01,
	0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x66, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x84, 0x03, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x46,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x65, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x1a, 0xf7, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x17, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x13, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x02, 0x74, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a,
	0x0d, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0b, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5a, 0x6f, 0x6e,
	0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x14, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x5f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x11, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x54, 0x72,
	0x75, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x66, 0x72, 0x6f, 0x6d,
	0x42, 0x05, 0x0a, 0x03, 0x5f, 0x74, 0x6f, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x74, 0x72, 0x75, 0x73,
	0x74, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x5f,
	0x69, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x62, 0x0a,
	0x17, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0b, 0x66, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x85, 0x01, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x0d, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x49,
	0x64, 0x12, 0x2f, 0x0a, 0x14, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x74, 0x72, 0x75, 0x73,
	0x74, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x54, 0x72, 0x75, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65,
	0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x22, 0x5f, 0x0a, 0x18, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xf0, 0x03, 0x0a, 0x11, 0x46,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x9d, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x9a, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x9d, 0x01,
	0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x53, 0x5a,
	0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x66, 0x69,
	0x64, 0x65, 0x2f, 0x63, 0x6f, 0x66, 0x69, 0x64, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_connect_federation_service_v1alpha1_federation_service_proto_rawDescOnce sync.Once
	file_proto_connect_federation_service_v1alpha1_federation_service_proto_rawDescData []byte
)

func file_proto_connect_federation_service_v1alpha1_federation_service_proto_rawDescGZIP() []byte {
	file_proto_connect_federation_service_v1alpha1_federation_service_proto_rawDescOnce.Do(func() {
		file_proto_connect_federation_service_v1alpha1_federation_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_connect_federation_service_v1alpha1_federation_service_proto_rawDesc), len(file_proto_connect_federation_service_v1alpha1_federation_service_proto_rawDesc)))
	})
	return file_proto_connect_federation_service_v1alpha1_federation_service_proto_rawDescData
}

var file_proto_connect_federation_service_v1alpha1_federation_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_connect_federation_service_v1alpha1_federation_service_proto_goTypes = []any{
	(*CreateFederationRequest)(nil),       // 0: proto.connect.federation_service.v1alpha1.CreateFederationRequest
	(*CreateFederationResponse)(nil),      // 1: proto.connect.federation_service.v1alpha1.CreateFederationResponse
	(*ListFederationsRequest)(nil),        // 2: proto.connect.federation_service.v1alpha1.ListFederationsRequest
	(*ListFederationsResponse)(nil),       // 3: proto.connect.federation_service.v1alpha1.ListFederationsResponse
	(*DeleteFederationRequest)(nil),       // 4: proto.connect.federation_service.v1alpha1.DeleteFederationRequest
	(*DeleteFederationResponse)(nil),      // 5: proto.connect.federation_service.v1alpha1.DeleteFederationResponse
	(*ListFederationsRequest_Filter)(nil), // 6: proto.connect.federation_service.v1alpha1.ListFederationsRequest.Filter
	(*v1alpha1.Federation)(nil),           // 7: proto.federation.v1alpha1.Federation
}
var file_proto_connect_federation_service_v1alpha1_federation_service_proto_depIdxs = []int32{
	7, // 0: proto.connect.federation_service.v1alpha1.CreateFederationRequest.federation:type_name -> proto.federation.v1alpha1.Federation
	7, // 1: proto.connect.federation_service.v1alpha1.CreateFederationResponse.federation:type_name -> proto.federation.v1alpha1.Federation
	6, // 2: proto.connect.federation_service.v1alpha1.ListFederationsRequest.filter:type_name -> proto.connect.federation_service.v1alpha1.ListFederationsRequest.Filter
	7, // 3: proto.connect.federation_service.v1alpha1.ListFederationsResponse.federations:type_name -> proto.federation.v1alpha1.Federation
	0, // 4: proto.connect.federation_service.v1alpha1.FederationService.CreateFederation:input_type -> proto.connect.federation_service.v1alpha1.CreateFederationRequest
	2, // 5: proto.connect.federation_service.v1alpha1.FederationService.ListFederations:input_type -> proto.connect.federation_service.v1alpha1.ListFederationsRequest
	4, // 6: proto.connect.federation_service.v1alpha1.FederationService.DeleteFederation:input_type -> proto.connect.federation_service.v1alpha1.DeleteFederationRequest
	1, // 7: proto.connect.federation_service.v1alpha1.FederationService.CreateFederation:output_type -> proto.connect.federation_service.v1alpha1.CreateFederationResponse
	3, // 8: proto.connect.federation_service.v1alpha1.FederationService.ListFederations:output_type -> proto.connect.federation_service.v1alpha1.ListFederationsResponse
	5, // 9: proto.connect.federation_service.v1alpha1.FederationService.DeleteFederation:output_type -> proto.connect.federation_service.v1alpha1.DeleteFederationResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_connect_federation_service_v1alpha1_federation_service_proto_init() }
func file_proto_connect_federation_service_v1alpha1_federation_service_proto_init() {
	if File_proto_connect_federation_service_v1alpha1_federation_service_proto != nil {
		return
	}
	file_proto_connect_federation_service_v1alpha1_federation_service_proto_msgTypes[1].OneofWrappers = []any{}
	file_proto_connect_federation_service_v1alpha1_federation_service_proto_msgTypes[2].OneofWrappers = []any{}
	file_proto_connect_federation_service_v1alpha1_federation_service_proto_msgTypes[5].OneofWrappers = []any{}
	file_proto_connect_federation_service_v1alpha1_federation_service_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_connect_federation_service_v1alpha1_federation_service_proto_rawDesc), len(file_proto_connect_federation_service_v1alpha1_federation_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_connect_federation_service_v1alpha1_federation_service_proto_goTypes,
		DependencyIndexes: file_proto_connect_federation_service_v1alpha1_federation_service_proto_depIdxs,
		MessageInfos:      file_proto_connect_federation_service_v1alpha1_federation_service_proto_msgTypes,
	}.Build()
	File_proto_connect_federation_service_v1alpha1_federation_service_proto = out.File
	file_proto_connect_federation_service_v1alpha1_federation_service_proto_goTypes = nil
	file_proto_connect_federation_service_v1alpha1_federation_service_proto_depIdxs = nil
}
