// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: client_side.proto

package protos

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

type ClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resp string `protobuf:"bytes,1,opt,name=resp,proto3" json:"resp,omitempty"`
}

func (x *ClientResponse) Reset() {
	*x = ClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_side_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientResponse) ProtoMessage() {}

func (x *ClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_side_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientResponse.ProtoReflect.Descriptor instead.
func (*ClientResponse) Descriptor() ([]byte, []int) {
	return file_client_side_proto_rawDescGZIP(), []int{0}
}

func (x *ClientResponse) GetResp() string {
	if x != nil {
		return x.Resp
	}
	return ""
}

type ClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Req string `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
}

func (x *ClientRequest) Reset() {
	*x = ClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_side_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientRequest) ProtoMessage() {}

func (x *ClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_side_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientRequest.ProtoReflect.Descriptor instead.
func (*ClientRequest) Descriptor() ([]byte, []int) {
	return file_client_side_proto_rawDescGZIP(), []int{1}
}

func (x *ClientRequest) GetReq() string {
	if x != nil {
		return x.Req
	}
	return ""
}

var File_client_side_proto protoreflect.FileDescriptor

var file_client_side_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x0e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x73, 0x70, 0x22, 0x21, 0x0a, 0x0d, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65,
	0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x71, 0x32, 0x42, 0x0a, 0x0a,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x64, 0x65, 0x12, 0x34, 0x0a, 0x0f, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x64, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0e, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01,
	0x42, 0x13, 0x5a, 0x11, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_side_proto_rawDescOnce sync.Once
	file_client_side_proto_rawDescData = file_client_side_proto_rawDesc
)

func file_client_side_proto_rawDescGZIP() []byte {
	file_client_side_proto_rawDescOnce.Do(func() {
		file_client_side_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_side_proto_rawDescData)
	})
	return file_client_side_proto_rawDescData
}

var file_client_side_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_client_side_proto_goTypes = []interface{}{
	(*ClientResponse)(nil), // 0: ClientResponse
	(*ClientRequest)(nil),  // 1: ClientRequest
}
var file_client_side_proto_depIdxs = []int32{
	1, // 0: clientSide.ClientSideHello:input_type -> ClientRequest
	0, // 1: clientSide.ClientSideHello:output_type -> ClientResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_client_side_proto_init() }
func file_client_side_proto_init() {
	if File_client_side_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_side_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientResponse); i {
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
		file_client_side_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientRequest); i {
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
			RawDescriptor: file_client_side_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_client_side_proto_goTypes,
		DependencyIndexes: file_client_side_proto_depIdxs,
		MessageInfos:      file_client_side_proto_msgTypes,
	}.Build()
	File_client_side_proto = out.File
	file_client_side_proto_rawDesc = nil
	file_client_side_proto_goTypes = nil
	file_client_side_proto_depIdxs = nil
}
