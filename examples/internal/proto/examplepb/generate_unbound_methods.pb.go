// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.0
// source: examples/internal/proto/examplepb/generate_unbound_methods.proto

// Generate Unannotated Methods Echo Service
// Similar to echo_service.proto but without annotations and without external configuration.
//
// Generate Unannotated Methods Echo Service API consists of a single service which returns
// a message.

package examplepb

import (
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// GenerateUnboundMethodsSimpleMessage represents a simple message sent to the unannotated GenerateUnboundMethodsEchoService service.
type GenerateUnboundMethodsSimpleMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id represents the message identifier.
	Id       string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Num      int64              `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	Duration *duration.Duration `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *GenerateUnboundMethodsSimpleMessage) Reset() {
	*x = GenerateUnboundMethodsSimpleMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examples_internal_proto_examplepb_generate_unbound_methods_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateUnboundMethodsSimpleMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateUnboundMethodsSimpleMessage) ProtoMessage() {}

func (x *GenerateUnboundMethodsSimpleMessage) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_generate_unbound_methods_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateUnboundMethodsSimpleMessage.ProtoReflect.Descriptor instead.
func (*GenerateUnboundMethodsSimpleMessage) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_generate_unbound_methods_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateUnboundMethodsSimpleMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GenerateUnboundMethodsSimpleMessage) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *GenerateUnboundMethodsSimpleMessage) GetDuration() *duration.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

var File_examples_internal_proto_examplepb_generate_unbound_methods_proto protoreflect.FileDescriptor

var file_examples_internal_proto_examplepb_generate_unbound_methods_proto_rawDesc = []byte{
	0x0a, 0x40, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x70, 0x62, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x6e, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x70, 0x62, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x23, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x6e,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x35, 0x0a, 0x08, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x32, 0xc6, 0x04, 0x0a, 0x21, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55,
	0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x45, 0x63, 0x68,
	0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xb0, 0x01, 0x0a, 0x04, 0x45, 0x63, 0x68,
	0x6f, 0x12, 0x53, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x53, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x55, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0xb4, 0x01, 0x0a, 0x08,
	0x45, 0x63, 0x68, 0x6f, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x53, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x55, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x53, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x73, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0xb6, 0x01, 0x0a, 0x0a, 0x45, 0x63, 0x68, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x53, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x53, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x55, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x53, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x4d, 0x5a, 0x4b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65,
	0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_examples_internal_proto_examplepb_generate_unbound_methods_proto_rawDescOnce sync.Once
	file_examples_internal_proto_examplepb_generate_unbound_methods_proto_rawDescData = file_examples_internal_proto_examplepb_generate_unbound_methods_proto_rawDesc
)

func file_examples_internal_proto_examplepb_generate_unbound_methods_proto_rawDescGZIP() []byte {
	file_examples_internal_proto_examplepb_generate_unbound_methods_proto_rawDescOnce.Do(func() {
		file_examples_internal_proto_examplepb_generate_unbound_methods_proto_rawDescData = protoimpl.X.CompressGZIP(file_examples_internal_proto_examplepb_generate_unbound_methods_proto_rawDescData)
	})
	return file_examples_internal_proto_examplepb_generate_unbound_methods_proto_rawDescData
}

var file_examples_internal_proto_examplepb_generate_unbound_methods_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_examples_internal_proto_examplepb_generate_unbound_methods_proto_goTypes = []interface{}{
	(*GenerateUnboundMethodsSimpleMessage)(nil), // 0: grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsSimpleMessage
	(*duration.Duration)(nil),                   // 1: google.protobuf.Duration
}
var file_examples_internal_proto_examplepb_generate_unbound_methods_proto_depIdxs = []int32{
	1, // 0: grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsSimpleMessage.duration:type_name -> google.protobuf.Duration
	0, // 1: grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsEchoService.Echo:input_type -> grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsSimpleMessage
	0, // 2: grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsEchoService.EchoBody:input_type -> grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsSimpleMessage
	0, // 3: grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsEchoService.EchoDelete:input_type -> grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsSimpleMessage
	0, // 4: grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsEchoService.Echo:output_type -> grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsSimpleMessage
	0, // 5: grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsEchoService.EchoBody:output_type -> grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsSimpleMessage
	0, // 6: grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsEchoService.EchoDelete:output_type -> grpc.gateway.examples.internal.proto.examplepb.GenerateUnboundMethodsSimpleMessage
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_examples_internal_proto_examplepb_generate_unbound_methods_proto_init() }
func file_examples_internal_proto_examplepb_generate_unbound_methods_proto_init() {
	if File_examples_internal_proto_examplepb_generate_unbound_methods_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_examples_internal_proto_examplepb_generate_unbound_methods_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateUnboundMethodsSimpleMessage); i {
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
			RawDescriptor: file_examples_internal_proto_examplepb_generate_unbound_methods_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_examples_internal_proto_examplepb_generate_unbound_methods_proto_goTypes,
		DependencyIndexes: file_examples_internal_proto_examplepb_generate_unbound_methods_proto_depIdxs,
		MessageInfos:      file_examples_internal_proto_examplepb_generate_unbound_methods_proto_msgTypes,
	}.Build()
	File_examples_internal_proto_examplepb_generate_unbound_methods_proto = out.File
	file_examples_internal_proto_examplepb_generate_unbound_methods_proto_rawDesc = nil
	file_examples_internal_proto_examplepb_generate_unbound_methods_proto_goTypes = nil
	file_examples_internal_proto_examplepb_generate_unbound_methods_proto_depIdxs = nil
}
