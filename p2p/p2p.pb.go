// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative p2p/p2p.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.0
// source: p2p/p2p.proto

package p2p

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

type Send struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port int64 `protobuf:"varint,1,opt,name=Port,proto3" json:"Port,omitempty"`
}

func (x *Send) Reset() {
	*x = Send{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_p2p_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Send) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Send) ProtoMessage() {}

func (x *Send) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_p2p_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Send.ProtoReflect.Descriptor instead.
func (*Send) Descriptor() ([]byte, []int) {
	return file_p2p_p2p_proto_rawDescGZIP(), []int{0}
}

func (x *Send) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_p2p_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_p2p_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_p2p_p2p_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_p2p_p2p_proto protoreflect.FileDescriptor

var file_p2p_p2p_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x32, 0x70, 0x2f, 0x70, 0x32, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x50, 0x32, 0x50, 0x22, 0x1a, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x50, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74,
	0x22, 0x24, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x5e, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x12, 0x09, 0x2e, 0x50, 0x32, 0x50, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x1a, 0x0d, 0x2e, 0x50, 0x32,
	0x50, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x0a,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x09, 0x2e, 0x50, 0x32, 0x50,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x1a, 0x0d, 0x2e, 0x50, 0x32, 0x50, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2e, 0x2f, 0x70, 0x32, 0x70,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_p2p_p2p_proto_rawDescOnce sync.Once
	file_p2p_p2p_proto_rawDescData = file_p2p_p2p_proto_rawDesc
)

func file_p2p_p2p_proto_rawDescGZIP() []byte {
	file_p2p_p2p_proto_rawDescOnce.Do(func() {
		file_p2p_p2p_proto_rawDescData = protoimpl.X.CompressGZIP(file_p2p_p2p_proto_rawDescData)
	})
	return file_p2p_p2p_proto_rawDescData
}

var file_p2p_p2p_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_p2p_p2p_proto_goTypes = []interface{}{
	(*Send)(nil),     // 0: P2P.Send
	(*Response)(nil), // 1: P2P.Response
}
var file_p2p_p2p_proto_depIdxs = []int32{
	0, // 0: P2P.ChatService.Connect:input_type -> P2P.Send
	0, // 1: P2P.ChatService.Disconnect:input_type -> P2P.Send
	1, // 2: P2P.ChatService.Connect:output_type -> P2P.Response
	1, // 3: P2P.ChatService.Disconnect:output_type -> P2P.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_p2p_p2p_proto_init() }
func file_p2p_p2p_proto_init() {
	if File_p2p_p2p_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_p2p_p2p_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Send); i {
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
		file_p2p_p2p_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_p2p_p2p_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_p2p_p2p_proto_goTypes,
		DependencyIndexes: file_p2p_p2p_proto_depIdxs,
		MessageInfos:      file_p2p_p2p_proto_msgTypes,
	}.Build()
	File_p2p_p2p_proto = out.File
	file_p2p_p2p_proto_rawDesc = nil
	file_p2p_p2p_proto_goTypes = nil
	file_p2p_p2p_proto_depIdxs = nil
}
