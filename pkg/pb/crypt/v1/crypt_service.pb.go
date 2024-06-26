// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: crypt/v1/crypt_service.proto

package cryptv1

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

type EncryptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recipients string `protobuf:"bytes,1,opt,name=recipients,proto3" json:"recipients,omitempty"`
	Data       []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EncryptRequest) Reset() {
	*x = EncryptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypt_v1_crypt_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptRequest) ProtoMessage() {}

func (x *EncryptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crypt_v1_crypt_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptRequest.ProtoReflect.Descriptor instead.
func (*EncryptRequest) Descriptor() ([]byte, []int) {
	return file_crypt_v1_crypt_service_proto_rawDescGZIP(), []int{0}
}

func (x *EncryptRequest) GetRecipients() string {
	if x != nil {
		return x.Recipients
	}
	return ""
}

func (x *EncryptRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type EncryptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EncryptResponse) Reset() {
	*x = EncryptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypt_v1_crypt_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptResponse) ProtoMessage() {}

func (x *EncryptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crypt_v1_crypt_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptResponse.ProtoReflect.Descriptor instead.
func (*EncryptResponse) Descriptor() ([]byte, []int) {
	return file_crypt_v1_crypt_service_proto_rawDescGZIP(), []int{1}
}

func (x *EncryptResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_crypt_v1_crypt_service_proto protoreflect.FileDescriptor

var file_crypt_v1_crypt_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x44, 0x0a, 0x0e, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x25,
	0x0a, 0x0f, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x50, 0x0a, 0x0c, 0x43, 0x72, 0x79, 0x70, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x12, 0x18, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x99, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x43, 0x72, 0x79, 0x70, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x64, 0x76, 0x2e, 0x72, 0x75,
	0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x62, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x43, 0x72, 0x79, 0x70, 0x74, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x14, 0x43, 0x72, 0x79, 0x70, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x43, 0x72, 0x79, 0x70, 0x74, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crypt_v1_crypt_service_proto_rawDescOnce sync.Once
	file_crypt_v1_crypt_service_proto_rawDescData = file_crypt_v1_crypt_service_proto_rawDesc
)

func file_crypt_v1_crypt_service_proto_rawDescGZIP() []byte {
	file_crypt_v1_crypt_service_proto_rawDescOnce.Do(func() {
		file_crypt_v1_crypt_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_crypt_v1_crypt_service_proto_rawDescData)
	})
	return file_crypt_v1_crypt_service_proto_rawDescData
}

var file_crypt_v1_crypt_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_crypt_v1_crypt_service_proto_goTypes = []interface{}{
	(*EncryptRequest)(nil),  // 0: crypt.v1.EncryptRequest
	(*EncryptResponse)(nil), // 1: crypt.v1.EncryptResponse
}
var file_crypt_v1_crypt_service_proto_depIdxs = []int32{
	0, // 0: crypt.v1.CryptService.Encrypt:input_type -> crypt.v1.EncryptRequest
	1, // 1: crypt.v1.CryptService.Encrypt:output_type -> crypt.v1.EncryptResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_crypt_v1_crypt_service_proto_init() }
func file_crypt_v1_crypt_service_proto_init() {
	if File_crypt_v1_crypt_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crypt_v1_crypt_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptRequest); i {
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
		file_crypt_v1_crypt_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptResponse); i {
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
			RawDescriptor: file_crypt_v1_crypt_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crypt_v1_crypt_service_proto_goTypes,
		DependencyIndexes: file_crypt_v1_crypt_service_proto_depIdxs,
		MessageInfos:      file_crypt_v1_crypt_service_proto_msgTypes,
	}.Build()
	File_crypt_v1_crypt_service_proto = out.File
	file_crypt_v1_crypt_service_proto_rawDesc = nil
	file_crypt_v1_crypt_service_proto_goTypes = nil
	file_crypt_v1_crypt_service_proto_depIdxs = nil
}
