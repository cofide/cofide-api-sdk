// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: proto/spire/api/types/bundle.proto

package types

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

type Bundle struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the trust domain the bundle belongs to (e.g., "example.org").
	TrustDomain string `protobuf:"bytes,1,opt,name=trust_domain,json=trustDomain,proto3" json:"trust_domain,omitempty"`
	// X.509 authorities for authenticating X509-SVIDs.
	X509Authorities []*X509Certificate `protobuf:"bytes,2,rep,name=x509_authorities,json=x509Authorities,proto3" json:"x509_authorities,omitempty"`
	// JWT authorities for authenticating JWT-SVIDs.
	JwtAuthorities []*JWTKey `protobuf:"bytes,3,rep,name=jwt_authorities,json=jwtAuthorities,proto3" json:"jwt_authorities,omitempty"`
	// A hint on how often the bundle should be refreshed from the bundle
	// provider, in seconds. Can be zero (meaning no hint available).
	RefreshHint int64 `protobuf:"varint,4,opt,name=refresh_hint,json=refreshHint,proto3" json:"refresh_hint,omitempty"`
	// The sequence number of the bundle.
	SequenceNumber uint64 `protobuf:"varint,5,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Bundle) Reset() {
	*x = Bundle{}
	mi := &file_proto_spire_api_types_bundle_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bundle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bundle) ProtoMessage() {}

func (x *Bundle) ProtoReflect() protoreflect.Message {
	mi := &file_proto_spire_api_types_bundle_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bundle.ProtoReflect.Descriptor instead.
func (*Bundle) Descriptor() ([]byte, []int) {
	return file_proto_spire_api_types_bundle_proto_rawDescGZIP(), []int{0}
}

func (x *Bundle) GetTrustDomain() string {
	if x != nil {
		return x.TrustDomain
	}
	return ""
}

func (x *Bundle) GetX509Authorities() []*X509Certificate {
	if x != nil {
		return x.X509Authorities
	}
	return nil
}

func (x *Bundle) GetJwtAuthorities() []*JWTKey {
	if x != nil {
		return x.JwtAuthorities
	}
	return nil
}

func (x *Bundle) GetRefreshHint() int64 {
	if x != nil {
		return x.RefreshHint
	}
	return 0
}

func (x *Bundle) GetSequenceNumber() uint64 {
	if x != nil {
		return x.SequenceNumber
	}
	return 0
}

type X509Certificate struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The ASN.1 DER encoded bytes of the X.509 certificate.
	Asn1 []byte `protobuf:"bytes,1,opt,name=asn1,proto3" json:"asn1,omitempty"`
	// This authority is no longer secure and must not be used.
	Tainted       bool `protobuf:"varint,2,opt,name=tainted,proto3" json:"tainted,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *X509Certificate) Reset() {
	*x = X509Certificate{}
	mi := &file_proto_spire_api_types_bundle_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *X509Certificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*X509Certificate) ProtoMessage() {}

func (x *X509Certificate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_spire_api_types_bundle_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use X509Certificate.ProtoReflect.Descriptor instead.
func (*X509Certificate) Descriptor() ([]byte, []int) {
	return file_proto_spire_api_types_bundle_proto_rawDescGZIP(), []int{1}
}

func (x *X509Certificate) GetAsn1() []byte {
	if x != nil {
		return x.Asn1
	}
	return nil
}

func (x *X509Certificate) GetTainted() bool {
	if x != nil {
		return x.Tainted
	}
	return false
}

type JWTKey struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The PKIX encoded public key.
	PublicKey []byte `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// The key identifier.
	KeyId string `protobuf:"bytes,2,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// When the key expires (seconds since Unix epoch). If zero, the key does
	// not expire.
	ExpiresAt int64 `protobuf:"varint,3,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	// This authority is no longer secure and must not be used
	Tainted       bool `protobuf:"varint,4,opt,name=tainted,proto3" json:"tainted,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JWTKey) Reset() {
	*x = JWTKey{}
	mi := &file_proto_spire_api_types_bundle_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JWTKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWTKey) ProtoMessage() {}

func (x *JWTKey) ProtoReflect() protoreflect.Message {
	mi := &file_proto_spire_api_types_bundle_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWTKey.ProtoReflect.Descriptor instead.
func (*JWTKey) Descriptor() ([]byte, []int) {
	return file_proto_spire_api_types_bundle_proto_rawDescGZIP(), []int{2}
}

func (x *JWTKey) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *JWTKey) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

func (x *JWTKey) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

func (x *JWTKey) GetTainted() bool {
	if x != nil {
		return x.Tainted
	}
	return false
}

type BundleMask struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// x509_authorities field mask.
	X509Authorities bool `protobuf:"varint,2,opt,name=x509_authorities,json=x509Authorities,proto3" json:"x509_authorities,omitempty"`
	// jwt_authorities field mask.
	JwtAuthorities bool `protobuf:"varint,3,opt,name=jwt_authorities,json=jwtAuthorities,proto3" json:"jwt_authorities,omitempty"`
	// refresh_hint field mask.
	RefreshHint bool `protobuf:"varint,4,opt,name=refresh_hint,json=refreshHint,proto3" json:"refresh_hint,omitempty"`
	// sequence_number field mask.
	SequenceNumber bool `protobuf:"varint,5,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *BundleMask) Reset() {
	*x = BundleMask{}
	mi := &file_proto_spire_api_types_bundle_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BundleMask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BundleMask) ProtoMessage() {}

func (x *BundleMask) ProtoReflect() protoreflect.Message {
	mi := &file_proto_spire_api_types_bundle_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BundleMask.ProtoReflect.Descriptor instead.
func (*BundleMask) Descriptor() ([]byte, []int) {
	return file_proto_spire_api_types_bundle_proto_rawDescGZIP(), []int{3}
}

func (x *BundleMask) GetX509Authorities() bool {
	if x != nil {
		return x.X509Authorities
	}
	return false
}

func (x *BundleMask) GetJwtAuthorities() bool {
	if x != nil {
		return x.JwtAuthorities
	}
	return false
}

func (x *BundleMask) GetRefreshHint() bool {
	if x != nil {
		return x.RefreshHint
	}
	return false
}

func (x *BundleMask) GetSequenceNumber() bool {
	if x != nil {
		return x.SequenceNumber
	}
	return false
}

var File_proto_spire_api_types_bundle_proto protoreflect.FileDescriptor

var file_proto_spire_api_types_bundle_proto_rawDesc = string([]byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x86, 0x02, 0x0a, 0x06, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x75, 0x73, 0x74, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x12, 0x4b, 0x0a, 0x10, 0x78, 0x35, 0x30, 0x39, 0x5f, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x58, 0x35, 0x30, 0x39, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x0f, 0x78, 0x35, 0x30, 0x39, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x40, 0x0a, 0x0f, 0x6a, 0x77, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x69, 0x72,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4a, 0x57, 0x54, 0x4b,
	0x65, 0x79, 0x52, 0x0e, 0x6a, 0x77, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x68, 0x69,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x48, 0x69, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x3f,
	0x0a, 0x0f, 0x58, 0x35, 0x30, 0x39, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x73, 0x6e, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x61, 0x73, 0x6e, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x74, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x64, 0x22,
	0x77, 0x0a, 0x06, 0x4a, 0x57, 0x54, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x15, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x74, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x74, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x64, 0x22, 0xac, 0x01, 0x0a, 0x0a, 0x42, 0x75, 0x6e,
	0x64, 0x6c, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x29, 0x0a, 0x10, 0x78, 0x35, 0x30, 0x39, 0x5f,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0f, 0x78, 0x35, 0x30, 0x39, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6a, 0x77, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6a, 0x77, 0x74,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x68, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x48, 0x69, 0x6e, 0x74, 0x12, 0x27,
	0x0a, 0x0f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x2f, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_spire_api_types_bundle_proto_rawDescOnce sync.Once
	file_proto_spire_api_types_bundle_proto_rawDescData []byte
)

func file_proto_spire_api_types_bundle_proto_rawDescGZIP() []byte {
	file_proto_spire_api_types_bundle_proto_rawDescOnce.Do(func() {
		file_proto_spire_api_types_bundle_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_spire_api_types_bundle_proto_rawDesc), len(file_proto_spire_api_types_bundle_proto_rawDesc)))
	})
	return file_proto_spire_api_types_bundle_proto_rawDescData
}

var file_proto_spire_api_types_bundle_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_spire_api_types_bundle_proto_goTypes = []any{
	(*Bundle)(nil),          // 0: spire.api.types.Bundle
	(*X509Certificate)(nil), // 1: spire.api.types.X509Certificate
	(*JWTKey)(nil),          // 2: spire.api.types.JWTKey
	(*BundleMask)(nil),      // 3: spire.api.types.BundleMask
}
var file_proto_spire_api_types_bundle_proto_depIdxs = []int32{
	1, // 0: spire.api.types.Bundle.x509_authorities:type_name -> spire.api.types.X509Certificate
	2, // 1: spire.api.types.Bundle.jwt_authorities:type_name -> spire.api.types.JWTKey
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_spire_api_types_bundle_proto_init() }
func file_proto_spire_api_types_bundle_proto_init() {
	if File_proto_spire_api_types_bundle_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_spire_api_types_bundle_proto_rawDesc), len(file_proto_spire_api_types_bundle_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_spire_api_types_bundle_proto_goTypes,
		DependencyIndexes: file_proto_spire_api_types_bundle_proto_depIdxs,
		MessageInfos:      file_proto_spire_api_types_bundle_proto_msgTypes,
	}.Build()
	File_proto_spire_api_types_bundle_proto = out.File
	file_proto_spire_api_types_bundle_proto_goTypes = nil
	file_proto_spire_api_types_bundle_proto_depIdxs = nil
}
