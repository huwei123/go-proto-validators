// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.3
// source: github.com/mwitkow/go-proto-validators/examples/nested.proto

package validator_examples

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
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

type InnerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SomeInteger int32 `protobuf:"varint,1,opt,name=some_integer,json=someInteger,proto3" json:"some_integer,omitempty"`
}

func (x *InnerMessage) Reset() {
	*x = InnerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_mwitkow_go_proto_validators_examples_nested_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerMessage) ProtoMessage() {}

func (x *InnerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_mwitkow_go_proto_validators_examples_nested_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerMessage.ProtoReflect.Descriptor instead.
func (*InnerMessage) Descriptor() ([]byte, []int) {
	return file_github_com_mwitkow_go_proto_validators_examples_nested_proto_rawDescGZIP(), []int{0}
}

func (x *InnerMessage) GetSomeInteger() int32 {
	if x != nil {
		return x.SomeInteger
	}
	return 0
}

type OuterMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImportantString string        `protobuf:"bytes,1,opt,name=important_string,json=importantString,proto3" json:"important_string,omitempty"`
	Inner           *InnerMessage `protobuf:"bytes,2,opt,name=inner,proto3" json:"inner,omitempty"`
}

func (x *OuterMessage) Reset() {
	*x = OuterMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_mwitkow_go_proto_validators_examples_nested_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OuterMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OuterMessage) ProtoMessage() {}

func (x *OuterMessage) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_mwitkow_go_proto_validators_examples_nested_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OuterMessage.ProtoReflect.Descriptor instead.
func (*OuterMessage) Descriptor() ([]byte, []int) {
	return file_github_com_mwitkow_go_proto_validators_examples_nested_proto_rawDescGZIP(), []int{1}
}

func (x *OuterMessage) GetImportantString() string {
	if x != nil {
		return x.ImportantString
	}
	return ""
}

func (x *OuterMessage) GetInner() *InnerMessage {
	if x != nil {
		return x.Inner
	}
	return nil
}

var File_github_com_mwitkow_go_proto_validators_examples_nested_proto protoreflect.FileDescriptor

var file_github_com_mwitkow_go_proto_validators_examples_nested_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69,
	0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x2f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x1a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x0c, 0x49,
	0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x0c, 0x73,
	0x6f, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x08, 0xe2, 0xdf, 0x1f, 0x04, 0x10, 0x00, 0x18, 0x64, 0x52, 0x0b, 0x73, 0x6f, 0x6d,
	0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x22, 0x8d, 0x01, 0x0a, 0x0c, 0x4f, 0x75, 0x74,
	0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3d, 0x0a, 0x10, 0x69, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x12, 0xe2, 0xdf, 0x1f, 0x0e, 0x0a, 0x0c, 0x5e, 0x5b, 0x61, 0x2d, 0x7a,
	0x5d, 0x7b, 0x32, 0x2c, 0x35, 0x7d, 0x24, 0x52, 0x0f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x61,
	0x6e, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x3e, 0x0a, 0x05, 0x69, 0x6e, 0x6e, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x6e,
	0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20,
	0x01, 0x52, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_mwitkow_go_proto_validators_examples_nested_proto_rawDescOnce sync.Once
	file_github_com_mwitkow_go_proto_validators_examples_nested_proto_rawDescData = file_github_com_mwitkow_go_proto_validators_examples_nested_proto_rawDesc
)

func file_github_com_mwitkow_go_proto_validators_examples_nested_proto_rawDescGZIP() []byte {
	file_github_com_mwitkow_go_proto_validators_examples_nested_proto_rawDescOnce.Do(func() {
		file_github_com_mwitkow_go_proto_validators_examples_nested_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_mwitkow_go_proto_validators_examples_nested_proto_rawDescData)
	})
	return file_github_com_mwitkow_go_proto_validators_examples_nested_proto_rawDescData
}

var file_github_com_mwitkow_go_proto_validators_examples_nested_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_mwitkow_go_proto_validators_examples_nested_proto_goTypes = []interface{}{
	(*InnerMessage)(nil), // 0: validator.examples.InnerMessage
	(*OuterMessage)(nil), // 1: validator.examples.OuterMessage
}
var file_github_com_mwitkow_go_proto_validators_examples_nested_proto_depIdxs = []int32{
	0, // 0: validator.examples.OuterMessage.inner:type_name -> validator.examples.InnerMessage
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_github_com_mwitkow_go_proto_validators_examples_nested_proto_init() }
func file_github_com_mwitkow_go_proto_validators_examples_nested_proto_init() {
	if File_github_com_mwitkow_go_proto_validators_examples_nested_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_mwitkow_go_proto_validators_examples_nested_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InnerMessage); i {
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
		file_github_com_mwitkow_go_proto_validators_examples_nested_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OuterMessage); i {
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
			RawDescriptor: file_github_com_mwitkow_go_proto_validators_examples_nested_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_mwitkow_go_proto_validators_examples_nested_proto_goTypes,
		DependencyIndexes: file_github_com_mwitkow_go_proto_validators_examples_nested_proto_depIdxs,
		MessageInfos:      file_github_com_mwitkow_go_proto_validators_examples_nested_proto_msgTypes,
	}.Build()
	File_github_com_mwitkow_go_proto_validators_examples_nested_proto = out.File
	file_github_com_mwitkow_go_proto_validators_examples_nested_proto_rawDesc = nil
	file_github_com_mwitkow_go_proto_validators_examples_nested_proto_goTypes = nil
	file_github_com_mwitkow_go_proto_validators_examples_nested_proto_depIdxs = nil
}
