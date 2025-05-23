// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v5.29.3
// source: grpc/org_sfomuseum_image_embosser.proto

package grpc

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

type EmbossImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=Filename,proto3" json:"Filename,omitempty"`
	Body     []byte `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
	Combined bool   `protobuf:"varint,3,opt,name=Combined,proto3" json:"Combined,omitempty"`
}

func (x *EmbossImageRequest) Reset() {
	*x = EmbossImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_org_sfomuseum_image_embosser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmbossImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmbossImageRequest) ProtoMessage() {}

func (x *EmbossImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_org_sfomuseum_image_embosser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmbossImageRequest.ProtoReflect.Descriptor instead.
func (*EmbossImageRequest) Descriptor() ([]byte, []int) {
	return file_grpc_org_sfomuseum_image_embosser_proto_rawDescGZIP(), []int{0}
}

func (x *EmbossImageRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *EmbossImageRequest) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *EmbossImageRequest) GetCombined() bool {
	if x != nil {
		return x.Combined
	}
	return false
}

type EmbossImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string   `protobuf:"bytes,1,opt,name=Filename,proto3" json:"Filename,omitempty"`
	Body     [][]byte `protobuf:"bytes,2,rep,name=Body,proto3" json:"Body,omitempty"`
	Combined bool     `protobuf:"varint,3,opt,name=Combined,proto3" json:"Combined,omitempty"`
}

func (x *EmbossImageResponse) Reset() {
	*x = EmbossImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_org_sfomuseum_image_embosser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmbossImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmbossImageResponse) ProtoMessage() {}

func (x *EmbossImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_org_sfomuseum_image_embosser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmbossImageResponse.ProtoReflect.Descriptor instead.
func (*EmbossImageResponse) Descriptor() ([]byte, []int) {
	return file_grpc_org_sfomuseum_image_embosser_proto_rawDescGZIP(), []int{1}
}

func (x *EmbossImageResponse) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *EmbossImageResponse) GetBody() [][]byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *EmbossImageResponse) GetCombined() bool {
	if x != nil {
		return x.Combined
	}
	return false
}

var File_grpc_org_sfomuseum_image_embosser_proto protoreflect.FileDescriptor

var file_grpc_org_sfomuseum_image_embosser_proto_rawDesc = []byte{
	0x0a, 0x27, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6f, 0x72, 0x67, 0x5f, 0x73, 0x66, 0x6f, 0x6d, 0x75,
	0x73, 0x65, 0x75, 0x6d, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x65, 0x6d, 0x62, 0x6f, 0x73,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x6f, 0x72, 0x67, 0x5f, 0x73,
	0x66, 0x6f, 0x6d, 0x75, 0x73, 0x65, 0x75, 0x6d, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x65,
	0x6d, 0x62, 0x6f, 0x73, 0x73, 0x65, 0x72, 0x22, 0x60, 0x0a, 0x12, 0x45, 0x6d, 0x62, 0x6f, 0x73,
	0x73, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f, 0x64,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x64, 0x22, 0x61, 0x0a, 0x13, 0x45, 0x6d, 0x62,
	0x6f, 0x73, 0x73, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x42, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x04, 0x42, 0x6f, 0x64, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x64, 0x32, 0x85, 0x01, 0x0a,
	0x0d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x45, 0x6d, 0x62, 0x6f, 0x73, 0x73, 0x65, 0x72, 0x12, 0x74,
	0x0a, 0x0b, 0x45, 0x6d, 0x62, 0x6f, 0x73, 0x73, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x30, 0x2e,
	0x6f, 0x72, 0x67, 0x5f, 0x73, 0x66, 0x6f, 0x6d, 0x75, 0x73, 0x65, 0x75, 0x6d, 0x5f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x65, 0x6d, 0x62, 0x6f, 0x73, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x62,
	0x6f, 0x73, 0x73, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x31, 0x2e, 0x6f, 0x72, 0x67, 0x5f, 0x73, 0x66, 0x6f, 0x6d, 0x75, 0x73, 0x65, 0x75, 0x6d, 0x5f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x65, 0x6d, 0x62, 0x6f, 0x73, 0x73, 0x65, 0x72, 0x2e, 0x45,
	0x6d, 0x62, 0x6f, 0x73, 0x73, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x66, 0x6f, 0x6d, 0x75, 0x73, 0x65, 0x75, 0x6d, 0x2f, 0x67, 0x6f, 0x2d,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x2d, 0x65, 0x6d, 0x62, 0x6f, 0x73, 0x73, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_org_sfomuseum_image_embosser_proto_rawDescOnce sync.Once
	file_grpc_org_sfomuseum_image_embosser_proto_rawDescData = file_grpc_org_sfomuseum_image_embosser_proto_rawDesc
)

func file_grpc_org_sfomuseum_image_embosser_proto_rawDescGZIP() []byte {
	file_grpc_org_sfomuseum_image_embosser_proto_rawDescOnce.Do(func() {
		file_grpc_org_sfomuseum_image_embosser_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_org_sfomuseum_image_embosser_proto_rawDescData)
	})
	return file_grpc_org_sfomuseum_image_embosser_proto_rawDescData
}

var file_grpc_org_sfomuseum_image_embosser_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_org_sfomuseum_image_embosser_proto_goTypes = []interface{}{
	(*EmbossImageRequest)(nil),  // 0: org_sfomuseum_image_embosser.EmbossImageRequest
	(*EmbossImageResponse)(nil), // 1: org_sfomuseum_image_embosser.EmbossImageResponse
}
var file_grpc_org_sfomuseum_image_embosser_proto_depIdxs = []int32{
	0, // 0: org_sfomuseum_image_embosser.ImageEmbosser.EmbossImage:input_type -> org_sfomuseum_image_embosser.EmbossImageRequest
	1, // 1: org_sfomuseum_image_embosser.ImageEmbosser.EmbossImage:output_type -> org_sfomuseum_image_embosser.EmbossImageResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_org_sfomuseum_image_embosser_proto_init() }
func file_grpc_org_sfomuseum_image_embosser_proto_init() {
	if File_grpc_org_sfomuseum_image_embosser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_org_sfomuseum_image_embosser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmbossImageRequest); i {
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
		file_grpc_org_sfomuseum_image_embosser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmbossImageResponse); i {
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
			RawDescriptor: file_grpc_org_sfomuseum_image_embosser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_org_sfomuseum_image_embosser_proto_goTypes,
		DependencyIndexes: file_grpc_org_sfomuseum_image_embosser_proto_depIdxs,
		MessageInfos:      file_grpc_org_sfomuseum_image_embosser_proto_msgTypes,
	}.Build()
	File_grpc_org_sfomuseum_image_embosser_proto = out.File
	file_grpc_org_sfomuseum_image_embosser_proto_rawDesc = nil
	file_grpc_org_sfomuseum_image_embosser_proto_goTypes = nil
	file_grpc_org_sfomuseum_image_embosser_proto_depIdxs = nil
}
