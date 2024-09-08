// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.1
// source: mls.proto

package mlspb

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

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Sum256  string `protobuf:"bytes,3,opt,name=sum256,proto3" json:"sum256,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mls_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_mls_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_mls_proto_rawDescGZIP(), []int{0}
}

func (x *Chunk) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Chunk) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Chunk) GetSum256() string {
	if x != nil {
		return x.Sum256
	}
	return ""
}

type Closer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *Closer) Reset() {
	*x = Closer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mls_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Closer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Closer) ProtoMessage() {}

func (x *Closer) ProtoReflect() protoreflect.Message {
	mi := &file_mls_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Closer.ProtoReflect.Descriptor instead.
func (*Closer) Descriptor() ([]byte, []int) {
	return file_mls_proto_rawDescGZIP(), []int{1}
}

func (x *Closer) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_mls_proto protoreflect.FileDescriptor

var file_mls_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6c, 0x73,
	0x70, 0x62, 0x22, 0x49, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x6d, 0x32, 0x35, 0x36, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x6d, 0x32, 0x35, 0x36, 0x22, 0x22, 0x0a,
	0x06, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x32, 0x35, 0x0a, 0x0a, 0x4d, 0x4c, 0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x27, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0c, 0x2e, 0x6d, 0x6c, 0x73, 0x70,
	0x62, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x0d, 0x2e, 0x6d, 0x6c, 0x73, 0x70, 0x62, 0x2e,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x72, 0x28, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x6d,
	0x6c, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mls_proto_rawDescOnce sync.Once
	file_mls_proto_rawDescData = file_mls_proto_rawDesc
)

func file_mls_proto_rawDescGZIP() []byte {
	file_mls_proto_rawDescOnce.Do(func() {
		file_mls_proto_rawDescData = protoimpl.X.CompressGZIP(file_mls_proto_rawDescData)
	})
	return file_mls_proto_rawDescData
}

var file_mls_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mls_proto_goTypes = []interface{}{
	(*Chunk)(nil),  // 0: mlspb.Chunk
	(*Closer)(nil), // 1: mlspb.Closer
}
var file_mls_proto_depIdxs = []int32{
	0, // 0: mlspb.MLSService.Upload:input_type -> mlspb.Chunk
	1, // 1: mlspb.MLSService.Upload:output_type -> mlspb.Closer
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mls_proto_init() }
func file_mls_proto_init() {
	if File_mls_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mls_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_mls_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Closer); i {
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
			RawDescriptor: file_mls_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mls_proto_goTypes,
		DependencyIndexes: file_mls_proto_depIdxs,
		MessageInfos:      file_mls_proto_msgTypes,
	}.Build()
	File_mls_proto = out.File
	file_mls_proto_rawDesc = nil
	file_mls_proto_goTypes = nil
	file_mls_proto_depIdxs = nil
}
