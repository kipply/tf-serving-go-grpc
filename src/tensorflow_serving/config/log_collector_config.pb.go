// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: tensorflow_serving/config/log_collector_config.proto

package tensorflow_serving

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

type LogCollectorConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifies the type of the LogCollector we will use to collect these logs.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// The prefix to use for the filenames of the logs.
	FilenamePrefix string `protobuf:"bytes,2,opt,name=filename_prefix,json=filenamePrefix,proto3" json:"filename_prefix,omitempty"`
}

func (x *LogCollectorConfig) Reset() {
	*x = LogCollectorConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_config_log_collector_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogCollectorConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogCollectorConfig) ProtoMessage() {}

func (x *LogCollectorConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_config_log_collector_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogCollectorConfig.ProtoReflect.Descriptor instead.
func (*LogCollectorConfig) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_config_log_collector_config_proto_rawDescGZIP(), []int{0}
}

func (x *LogCollectorConfig) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *LogCollectorConfig) GetFilenamePrefix() string {
	if x != nil {
		return x.FilenamePrefix
	}
	return ""
}

var File_tensorflow_serving_config_log_collector_config_proto protoreflect.FileDescriptor

var file_tensorflow_serving_config_log_collector_config_proto_rawDesc = []byte{
	0x0a, 0x34, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6c, 0x6f, 0x67, 0x5f,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x22, 0x51, 0x0a, 0x12, 0x4c, 0x6f,
	0x67, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x42, 0x03, 0xf8,
	0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_serving_config_log_collector_config_proto_rawDescOnce sync.Once
	file_tensorflow_serving_config_log_collector_config_proto_rawDescData = file_tensorflow_serving_config_log_collector_config_proto_rawDesc
)

func file_tensorflow_serving_config_log_collector_config_proto_rawDescGZIP() []byte {
	file_tensorflow_serving_config_log_collector_config_proto_rawDescOnce.Do(func() {
		file_tensorflow_serving_config_log_collector_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_serving_config_log_collector_config_proto_rawDescData)
	})
	return file_tensorflow_serving_config_log_collector_config_proto_rawDescData
}

var file_tensorflow_serving_config_log_collector_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_serving_config_log_collector_config_proto_goTypes = []interface{}{
	(*LogCollectorConfig)(nil), // 0: tensorflow.serving.LogCollectorConfig
}
var file_tensorflow_serving_config_log_collector_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_serving_config_log_collector_config_proto_init() }
func file_tensorflow_serving_config_log_collector_config_proto_init() {
	if File_tensorflow_serving_config_log_collector_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_serving_config_log_collector_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogCollectorConfig); i {
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
			RawDescriptor: file_tensorflow_serving_config_log_collector_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_serving_config_log_collector_config_proto_goTypes,
		DependencyIndexes: file_tensorflow_serving_config_log_collector_config_proto_depIdxs,
		MessageInfos:      file_tensorflow_serving_config_log_collector_config_proto_msgTypes,
	}.Build()
	File_tensorflow_serving_config_log_collector_config_proto = out.File
	file_tensorflow_serving_config_log_collector_config_proto_rawDesc = nil
	file_tensorflow_serving_config_log_collector_config_proto_goTypes = nil
	file_tensorflow_serving_config_log_collector_config_proto_depIdxs = nil
}
