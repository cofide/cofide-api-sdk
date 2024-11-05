// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: proto/ap_binding/v1alpha1/ap_binding.proto

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

type APBinding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrustZone     string   `protobuf:"bytes,1,opt,name=trust_zone,json=trustZone,proto3" json:"trust_zone,omitempty"`
	Policy        string   `protobuf:"bytes,2,opt,name=policy,proto3" json:"policy,omitempty"`
	FederatesWith []string `protobuf:"bytes,3,rep,name=federates_with,json=federatesWith,proto3" json:"federates_with,omitempty"`
}

func (x *APBinding) Reset() {
	*x = APBinding{}
	mi := &file_proto_ap_binding_v1alpha1_ap_binding_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *APBinding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APBinding) ProtoMessage() {}

func (x *APBinding) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ap_binding_v1alpha1_ap_binding_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APBinding.ProtoReflect.Descriptor instead.
func (*APBinding) Descriptor() ([]byte, []int) {
	return file_proto_ap_binding_v1alpha1_ap_binding_proto_rawDescGZIP(), []int{0}
}

func (x *APBinding) GetTrustZone() string {
	if x != nil {
		return x.TrustZone
	}
	return ""
}

func (x *APBinding) GetPolicy() string {
	if x != nil {
		return x.Policy
	}
	return ""
}

func (x *APBinding) GetFederatesWith() []string {
	if x != nil {
		return x.FederatesWith
	}
	return nil
}

var File_proto_ap_binding_v1alpha1_ap_binding_proto protoreflect.FileDescriptor

var file_proto_ap_binding_v1alpha1_ap_binding_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x70, 0x5f, 0x62,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x69, 0x0a, 0x09, 0x41, 0x50, 0x42, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x7a, 0x6f,
	0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5a,
	0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x66,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x73, 0x57, 0x69,
	0x74, 0x68, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x66, 0x69, 0x64, 0x65, 0x2f, 0x63, 0x6f, 0x66, 0x69, 0x64, 0x65, 0x2d, 0x61,
	0x70, 0x69, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x70, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ap_binding_v1alpha1_ap_binding_proto_rawDescOnce sync.Once
	file_proto_ap_binding_v1alpha1_ap_binding_proto_rawDescData = file_proto_ap_binding_v1alpha1_ap_binding_proto_rawDesc
)

func file_proto_ap_binding_v1alpha1_ap_binding_proto_rawDescGZIP() []byte {
	file_proto_ap_binding_v1alpha1_ap_binding_proto_rawDescOnce.Do(func() {
		file_proto_ap_binding_v1alpha1_ap_binding_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ap_binding_v1alpha1_ap_binding_proto_rawDescData)
	})
	return file_proto_ap_binding_v1alpha1_ap_binding_proto_rawDescData
}

var file_proto_ap_binding_v1alpha1_ap_binding_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_ap_binding_v1alpha1_ap_binding_proto_goTypes = []any{
	(*APBinding)(nil), // 0: proto.ap_binding.v1alpha1.APBinding
}
var file_proto_ap_binding_v1alpha1_ap_binding_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_ap_binding_v1alpha1_ap_binding_proto_init() }
func file_proto_ap_binding_v1alpha1_ap_binding_proto_init() {
	if File_proto_ap_binding_v1alpha1_ap_binding_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_ap_binding_v1alpha1_ap_binding_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_ap_binding_v1alpha1_ap_binding_proto_goTypes,
		DependencyIndexes: file_proto_ap_binding_v1alpha1_ap_binding_proto_depIdxs,
		MessageInfos:      file_proto_ap_binding_v1alpha1_ap_binding_proto_msgTypes,
	}.Build()
	File_proto_ap_binding_v1alpha1_ap_binding_proto = out.File
	file_proto_ap_binding_v1alpha1_ap_binding_proto_rawDesc = nil
	file_proto_ap_binding_v1alpha1_ap_binding_proto_goTypes = nil
	file_proto_ap_binding_v1alpha1_ap_binding_proto_depIdxs = nil
}
