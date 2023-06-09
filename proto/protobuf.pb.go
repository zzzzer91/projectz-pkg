// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: protobuf.proto

package proto

import (
	context "context"
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

type Protobuf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Protobuf) Reset() {
	*x = Protobuf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Protobuf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Protobuf) ProtoMessage() {}

func (x *Protobuf) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Protobuf.ProtoReflect.Descriptor instead.
func (*Protobuf) Descriptor() ([]byte, []int) {
	return file_protobuf_proto_rawDescGZIP(), []int{0}
}

func (x *Protobuf) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Protobuf) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Protobuf) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_protobuf_proto protoreflect.FileDescriptor

var file_protobuf_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x44, 0x0a, 0x08, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a,
	0x7a, 0x7a, 0x7a, 0x65, 0x72, 0x39, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7a,
	0x2d, 0x70, 0x6b, 0x67, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_proto_rawDescOnce sync.Once
	file_protobuf_proto_rawDescData = file_protobuf_proto_rawDesc
)

func file_protobuf_proto_rawDescGZIP() []byte {
	file_protobuf_proto_rawDescOnce.Do(func() {
		file_protobuf_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_proto_rawDescData)
	})
	return file_protobuf_proto_rawDescData
}

var file_protobuf_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protobuf_proto_goTypes = []interface{}{
	(*Protobuf)(nil), // 0: protobuf.Protobuf
}
var file_protobuf_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_proto_init() }
func file_protobuf_proto_init() {
	if File_protobuf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Protobuf); i {
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
			RawDescriptor: file_protobuf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_proto_goTypes,
		DependencyIndexes: file_protobuf_proto_depIdxs,
		MessageInfos:      file_protobuf_proto_msgTypes,
	}.Build()
	File_protobuf_proto = out.File
	file_protobuf_proto_rawDesc = nil
	file_protobuf_proto_goTypes = nil
	file_protobuf_proto_depIdxs = nil
}

var _ context.Context
