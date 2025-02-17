// Copyright 2024 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: proto/attestation_policy/v1alpha1/attestation_policy.proto

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

type AttestationPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to Policy:
	//
	//	*AttestationPolicy_Kubernetes
	//	*AttestationPolicy_Static
	Policy isAttestationPolicy_Policy `protobuf_oneof:"policy"`
}

func (x *AttestationPolicy) Reset() {
	*x = AttestationPolicy{}
	mi := &file_proto_attestation_policy_v1alpha1_attestation_policy_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AttestationPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttestationPolicy) ProtoMessage() {}

func (x *AttestationPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_proto_attestation_policy_v1alpha1_attestation_policy_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttestationPolicy.ProtoReflect.Descriptor instead.
func (*AttestationPolicy) Descriptor() ([]byte, []int) {
	return file_proto_attestation_policy_v1alpha1_attestation_policy_proto_rawDescGZIP(), []int{0}
}

func (x *AttestationPolicy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *AttestationPolicy) GetPolicy() isAttestationPolicy_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (x *AttestationPolicy) GetKubernetes() *APKubernetes {
	if x, ok := x.GetPolicy().(*AttestationPolicy_Kubernetes); ok {
		return x.Kubernetes
	}
	return nil
}

func (x *AttestationPolicy) GetStatic() *APStatic {
	if x, ok := x.GetPolicy().(*AttestationPolicy_Static); ok {
		return x.Static
	}
	return nil
}

type isAttestationPolicy_Policy interface {
	isAttestationPolicy_Policy()
}

type AttestationPolicy_Kubernetes struct {
	Kubernetes *APKubernetes `protobuf:"bytes,2,opt,name=kubernetes,proto3,oneof"`
}

type AttestationPolicy_Static struct {
	Static *APStatic `protobuf:"bytes,3,opt,name=static,proto3,oneof"`
}

func (*AttestationPolicy_Kubernetes) isAttestationPolicy_Policy() {}

func (*AttestationPolicy_Static) isAttestationPolicy_Policy() {}

type APKubernetes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NamespaceSelector *APLabelSelector `protobuf:"bytes,1,opt,name=namespace_selector,json=namespaceSelector,proto3,oneof" json:"namespace_selector,omitempty"`
	PodSelector       *APLabelSelector `protobuf:"bytes,2,opt,name=pod_selector,json=podSelector,proto3,oneof" json:"pod_selector,omitempty"`
}

func (x *APKubernetes) Reset() {
	*x = APKubernetes{}
	mi := &file_proto_attestation_policy_v1alpha1_attestation_policy_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *APKubernetes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APKubernetes) ProtoMessage() {}

func (x *APKubernetes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_attestation_policy_v1alpha1_attestation_policy_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APKubernetes.ProtoReflect.Descriptor instead.
func (*APKubernetes) Descriptor() ([]byte, []int) {
	return file_proto_attestation_policy_v1alpha1_attestation_policy_proto_rawDescGZIP(), []int{1}
}

func (x *APKubernetes) GetNamespaceSelector() *APLabelSelector {
	if x != nil {
		return x.NamespaceSelector
	}
	return nil
}

func (x *APKubernetes) GetPodSelector() *APLabelSelector {
	if x != nil {
		return x.PodSelector
	}
	return nil
}

// This definition has been adapted from the LabelSelector message in Kubernetes.
// https://github.com/kubernetes/apimachinery/blob/master/pkg/apis/meta/v1/generated.proto
type APLabelSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchLabels      map[string]string    `protobuf:"bytes,1,rep,name=match_labels,json=matchLabels,proto3" json:"match_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MatchExpressions []*APMatchExpression `protobuf:"bytes,2,rep,name=match_expressions,json=matchExpressions,proto3" json:"match_expressions,omitempty"`
}

func (x *APLabelSelector) Reset() {
	*x = APLabelSelector{}
	mi := &file_proto_attestation_policy_v1alpha1_attestation_policy_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *APLabelSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APLabelSelector) ProtoMessage() {}

func (x *APLabelSelector) ProtoReflect() protoreflect.Message {
	mi := &file_proto_attestation_policy_v1alpha1_attestation_policy_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APLabelSelector.ProtoReflect.Descriptor instead.
func (*APLabelSelector) Descriptor() ([]byte, []int) {
	return file_proto_attestation_policy_v1alpha1_attestation_policy_proto_rawDescGZIP(), []int{2}
}

func (x *APLabelSelector) GetMatchLabels() map[string]string {
	if x != nil {
		return x.MatchLabels
	}
	return nil
}

func (x *APLabelSelector) GetMatchExpressions() []*APMatchExpression {
	if x != nil {
		return x.MatchExpressions
	}
	return nil
}

type APMatchExpression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Operator string   `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	Values   []string `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *APMatchExpression) Reset() {
	*x = APMatchExpression{}
	mi := &file_proto_attestation_policy_v1alpha1_attestation_policy_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *APMatchExpression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APMatchExpression) ProtoMessage() {}

func (x *APMatchExpression) ProtoReflect() protoreflect.Message {
	mi := &file_proto_attestation_policy_v1alpha1_attestation_policy_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APMatchExpression.ProtoReflect.Descriptor instead.
func (*APMatchExpression) Descriptor() ([]byte, []int) {
	return file_proto_attestation_policy_v1alpha1_attestation_policy_proto_rawDescGZIP(), []int{3}
}

func (x *APMatchExpression) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *APMatchExpression) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *APMatchExpression) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

// APStatic represents a static attestation policy
type APStatic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpiffeId  string            `protobuf:"bytes,1,opt,name=spiffe_id,json=spiffeId,proto3" json:"spiffe_id,omitempty"`
	ParentId  string            `protobuf:"bytes,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Selectors map[string]string `protobuf:"bytes,3,rep,name=selectors,proto3" json:"selectors,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *APStatic) Reset() {
	*x = APStatic{}
	mi := &file_proto_attestation_policy_v1alpha1_attestation_policy_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *APStatic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APStatic) ProtoMessage() {}

func (x *APStatic) ProtoReflect() protoreflect.Message {
	mi := &file_proto_attestation_policy_v1alpha1_attestation_policy_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APStatic.ProtoReflect.Descriptor instead.
func (*APStatic) Descriptor() ([]byte, []int) {
	return file_proto_attestation_policy_v1alpha1_attestation_policy_proto_rawDescGZIP(), []int{4}
}

func (x *APStatic) GetSpiffeId() string {
	if x != nil {
		return x.SpiffeId
	}
	return ""
}

func (x *APStatic) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *APStatic) GetSelectors() map[string]string {
	if x != nil {
		return x.Selectors
	}
	return nil
}

var File_proto_attestation_policy_v1alpha1_attestation_policy_proto protoreflect.FileDescriptor

var file_proto_attestation_policy_v1alpha1_attestation_policy_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22,
	0xcb, 0x01, 0x0a, 0x11, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x51, 0x0a, 0x0a, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x41, 0x50, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x48, 0x00,
	0x52, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x12, 0x45, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x41, 0x50, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x42, 0x08, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0xfa, 0x01,
	0x0a, 0x0c, 0x41, 0x50, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x12, 0x66,
	0x0a, 0x12, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41,
	0x50, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x48, 0x00,
	0x52, 0x11, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x5a, 0x0a, 0x0c, 0x70, 0x6f, 0x64, 0x5f, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x41, 0x50, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x48, 0x01, 0x52, 0x0b, 0x70, 0x6f, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x88,
	0x01, 0x01, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x70, 0x6f,
	0x64, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x9c, 0x02, 0x0a, 0x0f, 0x41,
	0x50, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x66,
	0x0a, 0x0c, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x50, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x61, 0x0a, 0x11, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x50, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x59, 0x0a, 0x11, 0x41, 0x50, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x22, 0xdc, 0x01, 0x0a, 0x08, 0x41, 0x50, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x58, 0x0a, 0x09, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x41, 0x50, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2e, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x1a, 0x3c, 0x0a, 0x0e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x66, 0x69, 0x64, 0x65, 0x2f, 0x63, 0x6f, 0x66, 0x69, 0x64, 0x65, 0x2d,
	0x61, 0x70, 0x69, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_attestation_policy_v1alpha1_attestation_policy_proto_rawDescOnce sync.Once
	file_proto_attestation_policy_v1alpha1_attestation_policy_proto_rawDescData = file_proto_attestation_policy_v1alpha1_attestation_policy_proto_rawDesc
)

func file_proto_attestation_policy_v1alpha1_attestation_policy_proto_rawDescGZIP() []byte {
	file_proto_attestation_policy_v1alpha1_attestation_policy_proto_rawDescOnce.Do(func() {
		file_proto_attestation_policy_v1alpha1_attestation_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_attestation_policy_v1alpha1_attestation_policy_proto_rawDescData)
	})
	return file_proto_attestation_policy_v1alpha1_attestation_policy_proto_rawDescData
}

var file_proto_attestation_policy_v1alpha1_attestation_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_attestation_policy_v1alpha1_attestation_policy_proto_goTypes = []any{
	(*AttestationPolicy)(nil), // 0: proto.attestation_policy.v1alpha1.AttestationPolicy
	(*APKubernetes)(nil),      // 1: proto.attestation_policy.v1alpha1.APKubernetes
	(*APLabelSelector)(nil),   // 2: proto.attestation_policy.v1alpha1.APLabelSelector
	(*APMatchExpression)(nil), // 3: proto.attestation_policy.v1alpha1.APMatchExpression
	(*APStatic)(nil),          // 4: proto.attestation_policy.v1alpha1.APStatic
	nil,                       // 5: proto.attestation_policy.v1alpha1.APLabelSelector.MatchLabelsEntry
	nil,                       // 6: proto.attestation_policy.v1alpha1.APStatic.SelectorsEntry
}
var file_proto_attestation_policy_v1alpha1_attestation_policy_proto_depIdxs = []int32{
	1, // 0: proto.attestation_policy.v1alpha1.AttestationPolicy.kubernetes:type_name -> proto.attestation_policy.v1alpha1.APKubernetes
	4, // 1: proto.attestation_policy.v1alpha1.AttestationPolicy.static:type_name -> proto.attestation_policy.v1alpha1.APStatic
	2, // 2: proto.attestation_policy.v1alpha1.APKubernetes.namespace_selector:type_name -> proto.attestation_policy.v1alpha1.APLabelSelector
	2, // 3: proto.attestation_policy.v1alpha1.APKubernetes.pod_selector:type_name -> proto.attestation_policy.v1alpha1.APLabelSelector
	5, // 4: proto.attestation_policy.v1alpha1.APLabelSelector.match_labels:type_name -> proto.attestation_policy.v1alpha1.APLabelSelector.MatchLabelsEntry
	3, // 5: proto.attestation_policy.v1alpha1.APLabelSelector.match_expressions:type_name -> proto.attestation_policy.v1alpha1.APMatchExpression
	6, // 6: proto.attestation_policy.v1alpha1.APStatic.selectors:type_name -> proto.attestation_policy.v1alpha1.APStatic.SelectorsEntry
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_attestation_policy_v1alpha1_attestation_policy_proto_init() }
func file_proto_attestation_policy_v1alpha1_attestation_policy_proto_init() {
	if File_proto_attestation_policy_v1alpha1_attestation_policy_proto != nil {
		return
	}
	file_proto_attestation_policy_v1alpha1_attestation_policy_proto_msgTypes[0].OneofWrappers = []any{
		(*AttestationPolicy_Kubernetes)(nil),
		(*AttestationPolicy_Static)(nil),
	}
	file_proto_attestation_policy_v1alpha1_attestation_policy_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_attestation_policy_v1alpha1_attestation_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_attestation_policy_v1alpha1_attestation_policy_proto_goTypes,
		DependencyIndexes: file_proto_attestation_policy_v1alpha1_attestation_policy_proto_depIdxs,
		MessageInfos:      file_proto_attestation_policy_v1alpha1_attestation_policy_proto_msgTypes,
	}.Build()
	File_proto_attestation_policy_v1alpha1_attestation_policy_proto = out.File
	file_proto_attestation_policy_v1alpha1_attestation_policy_proto_rawDesc = nil
	file_proto_attestation_policy_v1alpha1_attestation_policy_proto_goTypes = nil
	file_proto_attestation_policy_v1alpha1_attestation_policy_proto_depIdxs = nil
}
