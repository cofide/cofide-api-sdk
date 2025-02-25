// Copyright 2024 Cofide Limited.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: proto/trust_zone/v1alpha1/trust_zone.proto

package v1alpha1

import (
	v1alpha11 "github.com/cofide/cofide-api-sdk/gen/go/proto/ap_binding/v1alpha1"
	v1alpha1 "github.com/cofide/cofide-api-sdk/gen/go/proto/federation/v1alpha1"
	types "github.com/spiffe/spire-api-sdk/proto/spire/api/types"
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

type BundleEndpointProfile int32

const (
	BundleEndpointProfile_BUNDLE_ENDPOINT_PROFILE_UNSPECIFIED  BundleEndpointProfile = 0
	BundleEndpointProfile_BUNDLE_ENDPOINT_PROFILE_HTTPS_SPIFFE BundleEndpointProfile = 1
	BundleEndpointProfile_BUNDLE_ENDPOINT_PROFILE_HTTPS_WEB    BundleEndpointProfile = 2
)

// Enum value maps for BundleEndpointProfile.
var (
	BundleEndpointProfile_name = map[int32]string{
		0: "BUNDLE_ENDPOINT_PROFILE_UNSPECIFIED",
		1: "BUNDLE_ENDPOINT_PROFILE_HTTPS_SPIFFE",
		2: "BUNDLE_ENDPOINT_PROFILE_HTTPS_WEB",
	}
	BundleEndpointProfile_value = map[string]int32{
		"BUNDLE_ENDPOINT_PROFILE_UNSPECIFIED":  0,
		"BUNDLE_ENDPOINT_PROFILE_HTTPS_SPIFFE": 1,
		"BUNDLE_ENDPOINT_PROFILE_HTTPS_WEB":    2,
	}
)

func (x BundleEndpointProfile) Enum() *BundleEndpointProfile {
	p := new(BundleEndpointProfile)
	*p = x
	return p
}

func (x BundleEndpointProfile) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BundleEndpointProfile) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_trust_zone_v1alpha1_trust_zone_proto_enumTypes[0].Descriptor()
}

func (BundleEndpointProfile) Type() protoreflect.EnumType {
	return &file_proto_trust_zone_v1alpha1_trust_zone_proto_enumTypes[0]
}

func (x BundleEndpointProfile) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BundleEndpointProfile.Descriptor instead.
func (BundleEndpointProfile) EnumDescriptor() ([]byte, []int) {
	return file_proto_trust_zone_v1alpha1_trust_zone_proto_rawDescGZIP(), []int{0}
}

type TrustZone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                  string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TrustDomain           string                 `protobuf:"bytes,2,opt,name=trust_domain,json=trustDomain,proto3" json:"trust_domain,omitempty"`
	BundleEndpointUrl     *string                `protobuf:"bytes,3,opt,name=bundle_endpoint_url,json=bundleEndpointUrl,proto3,oneof" json:"bundle_endpoint_url,omitempty"`
	Bundle                *types.Bundle          `protobuf:"bytes,4,opt,name=bundle,proto3,oneof" json:"bundle,omitempty"`
	Federations           []*v1alpha1.Federation `protobuf:"bytes,5,rep,name=federations,proto3" json:"federations,omitempty"`
	AttestationPolicies   []*v1alpha11.APBinding `protobuf:"bytes,6,rep,name=attestation_policies,json=attestationPolicies,proto3" json:"attestation_policies,omitempty"`
	JwtIssuer             *string                `protobuf:"bytes,7,opt,name=jwt_issuer,json=jwtIssuer,proto3,oneof" json:"jwt_issuer,omitempty"`
	BundleEndpointProfile *BundleEndpointProfile `protobuf:"varint,8,opt,name=bundle_endpoint_profile,json=bundleEndpointProfile,proto3,enum=proto.trust_zone.v1alpha1.BundleEndpointProfile,oneof" json:"bundle_endpoint_profile,omitempty"`
	Id                    *string                `protobuf:"bytes,9,opt,name=id,proto3,oneof" json:"id,omitempty"`
	IsManagementZone      bool                   `protobuf:"varint,10,opt,name=is_management_zone,json=isManagementZone,proto3" json:"is_management_zone,omitempty"`
}

func (x *TrustZone) Reset() {
	*x = TrustZone{}
	mi := &file_proto_trust_zone_v1alpha1_trust_zone_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrustZone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrustZone) ProtoMessage() {}

func (x *TrustZone) ProtoReflect() protoreflect.Message {
	mi := &file_proto_trust_zone_v1alpha1_trust_zone_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrustZone.ProtoReflect.Descriptor instead.
func (*TrustZone) Descriptor() ([]byte, []int) {
	return file_proto_trust_zone_v1alpha1_trust_zone_proto_rawDescGZIP(), []int{0}
}

func (x *TrustZone) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TrustZone) GetTrustDomain() string {
	if x != nil {
		return x.TrustDomain
	}
	return ""
}

func (x *TrustZone) GetBundleEndpointUrl() string {
	if x != nil && x.BundleEndpointUrl != nil {
		return *x.BundleEndpointUrl
	}
	return ""
}

func (x *TrustZone) GetBundle() *types.Bundle {
	if x != nil {
		return x.Bundle
	}
	return nil
}

func (x *TrustZone) GetFederations() []*v1alpha1.Federation {
	if x != nil {
		return x.Federations
	}
	return nil
}

func (x *TrustZone) GetAttestationPolicies() []*v1alpha11.APBinding {
	if x != nil {
		return x.AttestationPolicies
	}
	return nil
}

func (x *TrustZone) GetJwtIssuer() string {
	if x != nil && x.JwtIssuer != nil {
		return *x.JwtIssuer
	}
	return ""
}

func (x *TrustZone) GetBundleEndpointProfile() BundleEndpointProfile {
	if x != nil && x.BundleEndpointProfile != nil {
		return *x.BundleEndpointProfile
	}
	return BundleEndpointProfile_BUNDLE_ENDPOINT_PROFILE_UNSPECIFIED
}

func (x *TrustZone) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *TrustZone) GetIsManagementZone() bool {
	if x != nil {
		return x.IsManagementZone
	}
	return false
}

var File_proto_trust_zone_v1alpha1_trust_zone_proto protoreflect.FileDescriptor

var file_proto_trust_zone_v1alpha1_trust_zone_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x7a, 0x6f,
	0x6e, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x74, 0x72, 0x75, 0x73,
	0x74, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x2a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61,
	0x70, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x61, 0x70, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x66,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x04, 0x0a, 0x09, 0x54, 0x72, 0x75, 0x73, 0x74, 0x5a, 0x6f, 0x6e,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x75,
	0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x33, 0x0a, 0x13, 0x62, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x11, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a,
	0x06, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x48, 0x01, 0x52, 0x06, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x47, 0x0a, 0x0b, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x57, 0x0a, 0x14,
	0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x61, 0x70, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x50, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x52, 0x13, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0a, 0x6a, 0x77, 0x74, 0x5f, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x09, 0x6a, 0x77, 0x74,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x6d, 0x0a, 0x17, 0x62, 0x75, 0x6e,
	0x64, 0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x48, 0x03, 0x52, 0x15,
	0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a,
	0x12, 0x69, 0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x7a,
	0x6f, 0x6e, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x42, 0x16, 0x0a, 0x14, 0x5f,
	0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f,
	0x75, 0x72, 0x6c, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x6a, 0x77, 0x74, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x42, 0x1a, 0x0a,
	0x18, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64,
	0x2a, 0x91, 0x01, 0x0a, 0x15, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x23, 0x42, 0x55,
	0x4e, 0x44, 0x4c, 0x45, 0x5f, 0x45, 0x4e, 0x44, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x50, 0x52,
	0x4f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x28, 0x0a, 0x24, 0x42, 0x55, 0x4e, 0x44, 0x4c, 0x45, 0x5f, 0x45, 0x4e,
	0x44, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x48,
	0x54, 0x54, 0x50, 0x53, 0x5f, 0x53, 0x50, 0x49, 0x46, 0x46, 0x45, 0x10, 0x01, 0x12, 0x25, 0x0a,
	0x21, 0x42, 0x55, 0x4e, 0x44, 0x4c, 0x45, 0x5f, 0x45, 0x4e, 0x44, 0x50, 0x4f, 0x49, 0x4e, 0x54,
	0x5f, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x48, 0x54, 0x54, 0x50, 0x53, 0x5f, 0x57,
	0x45, 0x42, 0x10, 0x02, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x66, 0x69, 0x64, 0x65, 0x2f, 0x63, 0x6f, 0x66, 0x69, 0x64, 0x65,
	0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x7a, 0x6f, 0x6e, 0x65,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_trust_zone_v1alpha1_trust_zone_proto_rawDescOnce sync.Once
	file_proto_trust_zone_v1alpha1_trust_zone_proto_rawDescData = file_proto_trust_zone_v1alpha1_trust_zone_proto_rawDesc
)

func file_proto_trust_zone_v1alpha1_trust_zone_proto_rawDescGZIP() []byte {
	file_proto_trust_zone_v1alpha1_trust_zone_proto_rawDescOnce.Do(func() {
		file_proto_trust_zone_v1alpha1_trust_zone_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_trust_zone_v1alpha1_trust_zone_proto_rawDescData)
	})
	return file_proto_trust_zone_v1alpha1_trust_zone_proto_rawDescData
}

var file_proto_trust_zone_v1alpha1_trust_zone_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_trust_zone_v1alpha1_trust_zone_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_trust_zone_v1alpha1_trust_zone_proto_goTypes = []any{
	(BundleEndpointProfile)(0),  // 0: proto.trust_zone.v1alpha1.BundleEndpointProfile
	(*TrustZone)(nil),           // 1: proto.trust_zone.v1alpha1.TrustZone
	(*types.Bundle)(nil),        // 2: spire.api.types.Bundle
	(*v1alpha1.Federation)(nil), // 3: proto.federation.v1alpha1.Federation
	(*v1alpha11.APBinding)(nil), // 4: proto.ap_binding.v1alpha1.APBinding
}
var file_proto_trust_zone_v1alpha1_trust_zone_proto_depIdxs = []int32{
	2, // 0: proto.trust_zone.v1alpha1.TrustZone.bundle:type_name -> spire.api.types.Bundle
	3, // 1: proto.trust_zone.v1alpha1.TrustZone.federations:type_name -> proto.federation.v1alpha1.Federation
	4, // 2: proto.trust_zone.v1alpha1.TrustZone.attestation_policies:type_name -> proto.ap_binding.v1alpha1.APBinding
	0, // 3: proto.trust_zone.v1alpha1.TrustZone.bundle_endpoint_profile:type_name -> proto.trust_zone.v1alpha1.BundleEndpointProfile
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_trust_zone_v1alpha1_trust_zone_proto_init() }
func file_proto_trust_zone_v1alpha1_trust_zone_proto_init() {
	if File_proto_trust_zone_v1alpha1_trust_zone_proto != nil {
		return
	}
	file_proto_trust_zone_v1alpha1_trust_zone_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_trust_zone_v1alpha1_trust_zone_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_trust_zone_v1alpha1_trust_zone_proto_goTypes,
		DependencyIndexes: file_proto_trust_zone_v1alpha1_trust_zone_proto_depIdxs,
		EnumInfos:         file_proto_trust_zone_v1alpha1_trust_zone_proto_enumTypes,
		MessageInfos:      file_proto_trust_zone_v1alpha1_trust_zone_proto_msgTypes,
	}.Build()
	File_proto_trust_zone_v1alpha1_trust_zone_proto = out.File
	file_proto_trust_zone_v1alpha1_trust_zone_proto_rawDesc = nil
	file_proto_trust_zone_v1alpha1_trust_zone_proto_goTypes = nil
	file_proto_trust_zone_v1alpha1_trust_zone_proto_depIdxs = nil
}
