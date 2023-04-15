// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: signature.proto

package myprotos

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

type SignRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *SignRequest) Reset() {
	*x = SignRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signature_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRequest) ProtoMessage() {}

func (x *SignRequest) ProtoReflect() protoreflect.Message {
	mi := &file_signature_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRequest.ProtoReflect.Descriptor instead.
func (*SignRequest) Descriptor() ([]byte, []int) {
	return file_signature_proto_rawDescGZIP(), []int{0}
}

func (x *SignRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type SignResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SignResponse) Reset() {
	*x = SignResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signature_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignResponse) ProtoMessage() {}

func (x *SignResponse) ProtoReflect() protoreflect.Message {
	mi := &file_signature_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignResponse.ProtoReflect.Descriptor instead.
func (*SignResponse) Descriptor() ([]byte, []int) {
	return file_signature_proto_rawDescGZIP(), []int{1}
}

func (x *SignResponse) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

var File_signature_proto protoreflect.FileDescriptor

var file_signature_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x22, 0x21, 0x0a,
	0x0b, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x22, 0x2c, 0x0a, 0x0c, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x32, 0x49,
	0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x3b, 0x0a, 0x04,
	0x53, 0x69, 0x67, 0x6e, 0x12, 0x17, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x73, 0x69, 0x67, 0x6e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x55, 0x0a, 0x18, 0x63, 0x6f, 0x6d,
	0x2e, 0x64, 0x65, 0x72, 0x73, 0x6f, 0x2e, 0x6e, 0x61, 0x6f, 0x76, 0x6f, 0x75, 0x75, 0x73, 0x61,
	0x72, 0x6a, 0x61, 0x76, 0x61, 0x42, 0x10, 0x53, 0x69, 0x67, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x01, 0x5a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x64,
	0x65, 0x72, 0x73, 0x6f, 0x2f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x5f, 0x63, 0x72, 0x65, 0x75, 0x74,
	0x6f, 0x2f, 0x6d, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0xa2, 0x02, 0x03, 0x53, 0x56, 0x43,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_signature_proto_rawDescOnce sync.Once
	file_signature_proto_rawDescData = file_signature_proto_rawDesc
)

func file_signature_proto_rawDescGZIP() []byte {
	file_signature_proto_rawDescOnce.Do(func() {
		file_signature_proto_rawDescData = protoimpl.X.CompressGZIP(file_signature_proto_rawDescData)
	})
	return file_signature_proto_rawDescData
}

var file_signature_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_signature_proto_goTypes = []interface{}{
	(*SignRequest)(nil),  // 0: signverify.SignRequest
	(*SignResponse)(nil), // 1: signverify.SignResponse
}
var file_signature_proto_depIdxs = []int32{
	0, // 0: signverify.signVerify.Sign:input_type -> signverify.SignRequest
	1, // 1: signverify.signVerify.Sign:output_type -> signverify.SignResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_signature_proto_init() }
func file_signature_proto_init() {
	if File_signature_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_signature_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_signature_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_signature_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_signature_proto_goTypes,
		DependencyIndexes: file_signature_proto_depIdxs,
		MessageInfos:      file_signature_proto_msgTypes,
	}.Build()
	File_signature_proto = out.File
	file_signature_proto_rawDesc = nil
	file_signature_proto_goTypes = nil
	file_signature_proto_depIdxs = nil
}