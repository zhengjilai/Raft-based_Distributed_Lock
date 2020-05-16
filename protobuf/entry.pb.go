// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0-devel
// 	protoc        (unknown)
// source: protobuf/entry.proto

package protobuf

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the entry term
	Term uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	// the entry index
	Index uint64 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	// command name
	CommandName string `protobuf:"bytes,3,opt,name=commandName,proto3" json:"commandName,omitempty"`
	// command content in detail
	CommandContent []byte `protobuf:"bytes,4,opt,name=commandContent,proto3" json:"commandContent,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_entry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_entry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_protobuf_entry_proto_rawDescGZIP(), []int{0}
}

func (x *Entry) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *Entry) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Entry) GetCommandName() string {
	if x != nil {
		return x.CommandName
	}
	return ""
}

func (x *Entry) GetCommandContent() []byte {
	if x != nil {
		return x.CommandContent
	}
	return nil
}

var File_protobuf_entry_proto protoreflect.FileDescriptor

var file_protobuf_entry_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x22, 0x7b, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x20, 0x5a,
	0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_entry_proto_rawDescOnce sync.Once
	file_protobuf_entry_proto_rawDescData = file_protobuf_entry_proto_rawDesc
)

func file_protobuf_entry_proto_rawDescGZIP() []byte {
	file_protobuf_entry_proto_rawDescOnce.Do(func() {
		file_protobuf_entry_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_entry_proto_rawDescData)
	})
	return file_protobuf_entry_proto_rawDescData
}

var file_protobuf_entry_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protobuf_entry_proto_goTypes = []interface{}{
	(*Entry)(nil), // 0: protobuf.Entry
}
var file_protobuf_entry_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_entry_proto_init() }
func file_protobuf_entry_proto_init() {
	if File_protobuf_entry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_entry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
			RawDescriptor: file_protobuf_entry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_entry_proto_goTypes,
		DependencyIndexes: file_protobuf_entry_proto_depIdxs,
		MessageInfos:      file_protobuf_entry_proto_msgTypes,
	}.Build()
	File_protobuf_entry_proto = out.File
	file_protobuf_entry_proto_rawDesc = nil
	file_protobuf_entry_proto_goTypes = nil
	file_protobuf_entry_proto_depIdxs = nil
}
