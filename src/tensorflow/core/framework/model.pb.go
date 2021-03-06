// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: tensorflow/core/framework/model.proto

package tensorflow_data

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

// Class of a node in the performance model.
type NodeClass int32

const (
	NodeClass_UNKNOWN               NodeClass = 0
	NodeClass_INTERLEAVE_MANY       NodeClass = 1
	NodeClass_ASYNC_INTERLEAVE_MANY NodeClass = 2
	NodeClass_KNOWN_RATIO           NodeClass = 3
	NodeClass_ASYNC_KNOWN_RATIO     NodeClass = 4
	NodeClass_UNKNOWN_RATIO         NodeClass = 5
)

// Enum value maps for NodeClass.
var (
	NodeClass_name = map[int32]string{
		0: "UNKNOWN",
		1: "INTERLEAVE_MANY",
		2: "ASYNC_INTERLEAVE_MANY",
		3: "KNOWN_RATIO",
		4: "ASYNC_KNOWN_RATIO",
		5: "UNKNOWN_RATIO",
	}
	NodeClass_value = map[string]int32{
		"UNKNOWN":               0,
		"INTERLEAVE_MANY":       1,
		"ASYNC_INTERLEAVE_MANY": 2,
		"KNOWN_RATIO":           3,
		"ASYNC_KNOWN_RATIO":     4,
		"UNKNOWN_RATIO":         5,
	}
)

func (x NodeClass) Enum() *NodeClass {
	p := new(NodeClass)
	*p = x
	return p
}

func (x NodeClass) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeClass) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_core_framework_model_proto_enumTypes[0].Descriptor()
}

func (NodeClass) Type() protoreflect.EnumType {
	return &file_tensorflow_core_framework_model_proto_enumTypes[0]
}

func (x NodeClass) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodeClass.Descriptor instead.
func (NodeClass) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_model_proto_rawDescGZIP(), []int{0}
}

// Protocol buffer representing the data used by the autotuning modeling
// framework.
type ModelProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output node of this model.
	Output *ModelProto_Node `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	// Counter for node IDs of this model.
	IdCounter int64 `protobuf:"varint,2,opt,name=id_counter,json=idCounter,proto3" json:"id_counter,omitempty"`
	// Indicates whether the modeling framework should collect resource usage,
	// e.g. CPU, memory.
	CollectResourceUsage bool `protobuf:"varint,3,opt,name=collect_resource_usage,json=collectResourceUsage,proto3" json:"collect_resource_usage,omitempty"`
}

func (x *ModelProto) Reset() {
	*x = ModelProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelProto) ProtoMessage() {}

func (x *ModelProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelProto.ProtoReflect.Descriptor instead.
func (*ModelProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_model_proto_rawDescGZIP(), []int{0}
}

func (x *ModelProto) GetOutput() *ModelProto_Node {
	if x != nil {
		return x.Output
	}
	return nil
}

func (x *ModelProto) GetIdCounter() int64 {
	if x != nil {
		return x.IdCounter
	}
	return 0
}

func (x *ModelProto) GetCollectResourceUsage() bool {
	if x != nil {
		return x.CollectResourceUsage
	}
	return false
}

// General representation of a node in the model.
type ModelProto_Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique node ID.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Human-readable name of the node.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// An indication whether autotuning is enabled for this node.
	Autotune bool `protobuf:"varint,3,opt,name=autotune,proto3" json:"autotune,omitempty"`
	// The number of bytes stored in this node's buffer.
	BufferedBytes int64 `protobuf:"varint,4,opt,name=buffered_bytes,json=bufferedBytes,proto3" json:"buffered_bytes,omitempty"`
	// The number of elements stored in this node's buffer.
	BufferedElements int64 `protobuf:"varint,5,opt,name=buffered_elements,json=bufferedElements,proto3" json:"buffered_elements,omitempty"`
	// The number of bytes consumed by the node.
	BytesConsumed int64 `protobuf:"varint,6,opt,name=bytes_consumed,json=bytesConsumed,proto3" json:"bytes_consumed,omitempty"`
	// The number of bytes produced by the node.
	BytesProduced int64 `protobuf:"varint,7,opt,name=bytes_produced,json=bytesProduced,proto3" json:"bytes_produced,omitempty"`
	// The number of elements produced by the node.
	NumElements int64 `protobuf:"varint,8,opt,name=num_elements,json=numElements,proto3" json:"num_elements,omitempty"`
	// The aggregate processing time spent in this node.
	ProcessingTime int64 `protobuf:"varint,9,opt,name=processing_time,json=processingTime,proto3" json:"processing_time,omitempty"`
	// An indication whether this node records metrics about produced and
	// consumed elements.
	RecordMetrics bool `protobuf:"varint,10,opt,name=record_metrics,json=recordMetrics,proto3" json:"record_metrics,omitempty"`
	// Parameters of this node.
	Parameters []*ModelProto_Node_Parameter `protobuf:"bytes,11,rep,name=parameters,proto3" json:"parameters,omitempty"`
	// Statistic of inputs processing time history.
	InputProcessingTimeSum   float64 `protobuf:"fixed64,12,opt,name=input_processing_time_sum,json=inputProcessingTimeSum,proto3" json:"input_processing_time_sum,omitempty"`
	InputProcessingTimeCount int64   `protobuf:"varint,13,opt,name=input_processing_time_count,json=inputProcessingTimeCount,proto3" json:"input_processing_time_count,omitempty"`
	// Inputs of this node.
	Inputs []*ModelProto_Node `protobuf:"bytes,14,rep,name=inputs,proto3" json:"inputs,omitempty"`
	// Class of this node.
	NodeClass NodeClass `protobuf:"varint,15,opt,name=node_class,json=nodeClass,proto3,enum=tensorflow.data.NodeClass" json:"node_class,omitempty"`
	// Ratio of input to output elements. This is only used by KNOWN_RATIO and
	// ASYNC_KNOWN_RATIO nodes.
	Ratio float64 `protobuf:"fixed64,16,opt,name=ratio,proto3" json:"ratio,omitempty"`
	// Ratio identifies how many parallelism calls are introduced by one
	// buffered element. This is only used by ASYNC_KNOWN_RATIO nodes.
	MemoryRatio float64 `protobuf:"fixed64,17,opt,name=memory_ratio,json=memoryRatio,proto3" json:"memory_ratio,omitempty"`
}

func (x *ModelProto_Node) Reset() {
	*x = ModelProto_Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelProto_Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelProto_Node) ProtoMessage() {}

func (x *ModelProto_Node) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelProto_Node.ProtoReflect.Descriptor instead.
func (*ModelProto_Node) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_model_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ModelProto_Node) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ModelProto_Node) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ModelProto_Node) GetAutotune() bool {
	if x != nil {
		return x.Autotune
	}
	return false
}

func (x *ModelProto_Node) GetBufferedBytes() int64 {
	if x != nil {
		return x.BufferedBytes
	}
	return 0
}

func (x *ModelProto_Node) GetBufferedElements() int64 {
	if x != nil {
		return x.BufferedElements
	}
	return 0
}

func (x *ModelProto_Node) GetBytesConsumed() int64 {
	if x != nil {
		return x.BytesConsumed
	}
	return 0
}

func (x *ModelProto_Node) GetBytesProduced() int64 {
	if x != nil {
		return x.BytesProduced
	}
	return 0
}

func (x *ModelProto_Node) GetNumElements() int64 {
	if x != nil {
		return x.NumElements
	}
	return 0
}

func (x *ModelProto_Node) GetProcessingTime() int64 {
	if x != nil {
		return x.ProcessingTime
	}
	return 0
}

func (x *ModelProto_Node) GetRecordMetrics() bool {
	if x != nil {
		return x.RecordMetrics
	}
	return false
}

func (x *ModelProto_Node) GetParameters() []*ModelProto_Node_Parameter {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *ModelProto_Node) GetInputProcessingTimeSum() float64 {
	if x != nil {
		return x.InputProcessingTimeSum
	}
	return 0
}

func (x *ModelProto_Node) GetInputProcessingTimeCount() int64 {
	if x != nil {
		return x.InputProcessingTimeCount
	}
	return 0
}

func (x *ModelProto_Node) GetInputs() []*ModelProto_Node {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *ModelProto_Node) GetNodeClass() NodeClass {
	if x != nil {
		return x.NodeClass
	}
	return NodeClass_UNKNOWN
}

func (x *ModelProto_Node) GetRatio() float64 {
	if x != nil {
		return x.Ratio
	}
	return 0
}

func (x *ModelProto_Node) GetMemoryRatio() float64 {
	if x != nil {
		return x.MemoryRatio
	}
	return 0
}

// Represents a node parameter.
type ModelProto_Node_Parameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Human-readable name of the parameter.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Identifies the model value of the parameter. This can be different from
	// the actual value (e.g. during optimization search).
	Value float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	// The actual value of the parameter.
	StateValue float64 `protobuf:"fixed64,3,opt,name=state_value,json=stateValue,proto3" json:"state_value,omitempty"`
	// Minimum value of the parameter.
	Min float64 `protobuf:"fixed64,4,opt,name=min,proto3" json:"min,omitempty"`
	// Maximum value of the parameter.
	Max float64 `protobuf:"fixed64,5,opt,name=max,proto3" json:"max,omitempty"`
	// Identifies whether the parameter should participate in autotuning.
	Tunable bool `protobuf:"varint,6,opt,name=tunable,proto3" json:"tunable,omitempty"`
}

func (x *ModelProto_Node_Parameter) Reset() {
	*x = ModelProto_Node_Parameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelProto_Node_Parameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelProto_Node_Parameter) ProtoMessage() {}

func (x *ModelProto_Node_Parameter) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelProto_Node_Parameter.ProtoReflect.Descriptor instead.
func (*ModelProto_Node_Parameter) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_model_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *ModelProto_Node_Parameter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ModelProto_Node_Parameter) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *ModelProto_Node_Parameter) GetStateValue() float64 {
	if x != nil {
		return x.StateValue
	}
	return 0
}

func (x *ModelProto_Node_Parameter) GetMin() float64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *ModelProto_Node_Parameter) GetMax() float64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *ModelProto_Node_Parameter) GetTunable() bool {
	if x != nil {
		return x.Tunable
	}
	return false
}

var File_tensorflow_core_framework_model_proto protoreflect.FileDescriptor

var file_tensorflow_core_framework_model_proto_rawDesc = []byte{
	0x0a, 0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x22, 0x84, 0x08, 0x0a, 0x0a, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x38, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x12, 0x34, 0x0a, 0x16, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x14, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x1a, 0xe6, 0x06, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x6f, 0x74, 0x75, 0x6e, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x75, 0x74, 0x6f, 0x74, 0x75, 0x6e, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x65,
	0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72,
	0x65, 0x64, 0x5f, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x10, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x65, 0x64, 0x45, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x62, 0x79, 0x74, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x5f, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x45, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x12, 0x4a, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x12, 0x39, 0x0a, 0x19, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x75, 0x6d, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x16, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x75, 0x6d, 0x12, 0x3d, 0x0a, 0x1b, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x18, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x54, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x06, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x10, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x11, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x1a, 0x94, 0x01, 0x0a, 0x09, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03,
	0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x75, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x74, 0x75, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x2a,
	0x83, 0x01, 0x0a, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x4e,
	0x54, 0x45, 0x52, 0x4c, 0x45, 0x41, 0x56, 0x45, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x10, 0x01, 0x12,
	0x19, 0x0a, 0x15, 0x41, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4c, 0x45,
	0x41, 0x56, 0x45, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x5f, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x41,
	0x53, 0x59, 0x4e, 0x43, 0x5f, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x52, 0x41, 0x54, 0x49, 0x4f,
	0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x52, 0x41,
	0x54, 0x49, 0x4f, 0x10, 0x05, 0x42, 0x03, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tensorflow_core_framework_model_proto_rawDescOnce sync.Once
	file_tensorflow_core_framework_model_proto_rawDescData = file_tensorflow_core_framework_model_proto_rawDesc
)

func file_tensorflow_core_framework_model_proto_rawDescGZIP() []byte {
	file_tensorflow_core_framework_model_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_framework_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_framework_model_proto_rawDescData)
	})
	return file_tensorflow_core_framework_model_proto_rawDescData
}

var file_tensorflow_core_framework_model_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_core_framework_model_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tensorflow_core_framework_model_proto_goTypes = []interface{}{
	(NodeClass)(0),                    // 0: tensorflow.data.NodeClass
	(*ModelProto)(nil),                // 1: tensorflow.data.ModelProto
	(*ModelProto_Node)(nil),           // 2: tensorflow.data.ModelProto.Node
	(*ModelProto_Node_Parameter)(nil), // 3: tensorflow.data.ModelProto.Node.Parameter
}
var file_tensorflow_core_framework_model_proto_depIdxs = []int32{
	2, // 0: tensorflow.data.ModelProto.output:type_name -> tensorflow.data.ModelProto.Node
	3, // 1: tensorflow.data.ModelProto.Node.parameters:type_name -> tensorflow.data.ModelProto.Node.Parameter
	2, // 2: tensorflow.data.ModelProto.Node.inputs:type_name -> tensorflow.data.ModelProto.Node
	0, // 3: tensorflow.data.ModelProto.Node.node_class:type_name -> tensorflow.data.NodeClass
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tensorflow_core_framework_model_proto_init() }
func file_tensorflow_core_framework_model_proto_init() {
	if File_tensorflow_core_framework_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_framework_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelProto); i {
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
		file_tensorflow_core_framework_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelProto_Node); i {
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
		file_tensorflow_core_framework_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelProto_Node_Parameter); i {
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
			RawDescriptor: file_tensorflow_core_framework_model_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_framework_model_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_framework_model_proto_depIdxs,
		EnumInfos:         file_tensorflow_core_framework_model_proto_enumTypes,
		MessageInfos:      file_tensorflow_core_framework_model_proto_msgTypes,
	}.Build()
	File_tensorflow_core_framework_model_proto = out.File
	file_tensorflow_core_framework_model_proto_rawDesc = nil
	file_tensorflow_core_framework_model_proto_goTypes = nil
	file_tensorflow_core_framework_model_proto_depIdxs = nil
}
