// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: gid/proxy.proto

package gidpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type EchoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hello string `protobuf:"bytes,1,opt,name=hello,proto3" json:"hello,omitempty"`
}

func (x *EchoReq) Reset() {
	*x = EchoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gid_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoReq) ProtoMessage() {}

func (x *EchoReq) ProtoReflect() protoreflect.Message {
	mi := &file_gid_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoReq.ProtoReflect.Descriptor instead.
func (*EchoReq) Descriptor() ([]byte, []int) {
	return file_gid_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *EchoReq) GetHello() string {
	if x != nil {
		return x.Hello
	}
	return ""
}

type EchoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hello string `protobuf:"bytes,1,opt,name=hello,proto3" json:"hello,omitempty"`
}

func (x *EchoRsp) Reset() {
	*x = EchoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gid_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoRsp) ProtoMessage() {}

func (x *EchoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_gid_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoRsp.ProtoReflect.Descriptor instead.
func (*EchoRsp) Descriptor() ([]byte, []int) {
	return file_gid_proxy_proto_rawDescGZIP(), []int{1}
}

func (x *EchoRsp) GetHello() string {
	if x != nil {
		return x.Hello
	}
	return ""
}

var File_gid_proxy_proto protoreflect.FileDescriptor

var file_gid_proxy_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x69, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x67, 0x69, 0x64, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x07,
	0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x22, 0x1f, 0x0a,
	0x07, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x32, 0x52,
	0x0a, 0x07, 0x49, 0x64, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x47, 0x0a, 0x04, 0x45, 0x63, 0x68,
	0x6f, 0x12, 0x14, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x67, 0x69, 0x64, 0x2e,
	0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x67, 0x69, 0x64, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x73, 0x70, 0x22, 0x13, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08, 0x2f, 0x69, 0x64, 0x2f, 0x65, 0x63,
	0x68, 0x6f, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x75, 0x62, 0x67, 0x6f, 0x2f, 0x6c, 0x61, 0x76, 0x61, 0x2d, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x69, 0x64,
	0x70, 0x62, 0x3b, 0x67, 0x69, 0x64, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gid_proxy_proto_rawDescOnce sync.Once
	file_gid_proxy_proto_rawDescData = file_gid_proxy_proto_rawDesc
)

func file_gid_proxy_proto_rawDescGZIP() []byte {
	file_gid_proxy_proto_rawDescOnce.Do(func() {
		file_gid_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_gid_proxy_proto_rawDescData)
	})
	return file_gid_proxy_proto_rawDescData
}

var file_gid_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gid_proxy_proto_goTypes = []any{
	(*EchoReq)(nil), // 0: example.gid.EchoReq
	(*EchoRsp)(nil), // 1: example.gid.EchoRsp
}
var file_gid_proxy_proto_depIdxs = []int32{
	0, // 0: example.gid.IdProxy.Echo:input_type -> example.gid.EchoReq
	1, // 1: example.gid.IdProxy.Echo:output_type -> example.gid.EchoRsp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gid_proxy_proto_init() }
func file_gid_proxy_proto_init() {
	if File_gid_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gid_proxy_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*EchoReq); i {
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
		file_gid_proxy_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*EchoRsp); i {
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
			RawDescriptor: file_gid_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gid_proxy_proto_goTypes,
		DependencyIndexes: file_gid_proxy_proto_depIdxs,
		MessageInfos:      file_gid_proxy_proto_msgTypes,
	}.Build()
	File_gid_proxy_proto = out.File
	file_gid_proxy_proto_rawDesc = nil
	file_gid_proxy_proto_goTypes = nil
	file_gid_proxy_proto_depIdxs = nil
}
