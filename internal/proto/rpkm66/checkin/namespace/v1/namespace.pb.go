// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: rpkm66/checkin/namespace/v1/namespace.proto

package v1

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

type Namespace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Namespace) Reset() {
	*x = Namespace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_checkin_namespace_v1_namespace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Namespace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Namespace) ProtoMessage() {}

func (x *Namespace) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_checkin_namespace_v1_namespace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Namespace.ProtoReflect.Descriptor instead.
func (*Namespace) Descriptor() ([]byte, []int) {
	return file_rpkm66_checkin_namespace_v1_namespace_proto_rawDescGZIP(), []int{0}
}

func (x *Namespace) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// GetAllNamespaces
type GetAllNamespacesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllNamespacesRequest) Reset() {
	*x = GetAllNamespacesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_checkin_namespace_v1_namespace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllNamespacesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllNamespacesRequest) ProtoMessage() {}

func (x *GetAllNamespacesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_checkin_namespace_v1_namespace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllNamespacesRequest.ProtoReflect.Descriptor instead.
func (*GetAllNamespacesRequest) Descriptor() ([]byte, []int) {
	return file_rpkm66_checkin_namespace_v1_namespace_proto_rawDescGZIP(), []int{1}
}

type GetAllNamespacesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespaces []*Namespace `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
}

func (x *GetAllNamespacesResponse) Reset() {
	*x = GetAllNamespacesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_checkin_namespace_v1_namespace_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllNamespacesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllNamespacesResponse) ProtoMessage() {}

func (x *GetAllNamespacesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_checkin_namespace_v1_namespace_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllNamespacesResponse.ProtoReflect.Descriptor instead.
func (*GetAllNamespacesResponse) Descriptor() ([]byte, []int) {
	return file_rpkm66_checkin_namespace_v1_namespace_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllNamespacesResponse) GetNamespaces() []*Namespace {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

var File_rpkm66_checkin_namespace_v1_namespace_proto protoreflect.FileDescriptor

var file_rpkm66_checkin_namespace_v1_namespace_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e,
	0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x72,
	0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x1b, 0x0a, 0x09, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x19, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x62, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46,
	0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x69, 0x6e, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x32, 0x96, 0x01, 0x0a, 0x10, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x12, 0x34, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69,
	0x6e, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x1d, 0x5a, 0x1b, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69,
	0x6e, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpkm66_checkin_namespace_v1_namespace_proto_rawDescOnce sync.Once
	file_rpkm66_checkin_namespace_v1_namespace_proto_rawDescData = file_rpkm66_checkin_namespace_v1_namespace_proto_rawDesc
)

func file_rpkm66_checkin_namespace_v1_namespace_proto_rawDescGZIP() []byte {
	file_rpkm66_checkin_namespace_v1_namespace_proto_rawDescOnce.Do(func() {
		file_rpkm66_checkin_namespace_v1_namespace_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpkm66_checkin_namespace_v1_namespace_proto_rawDescData)
	})
	return file_rpkm66_checkin_namespace_v1_namespace_proto_rawDescData
}

var file_rpkm66_checkin_namespace_v1_namespace_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rpkm66_checkin_namespace_v1_namespace_proto_goTypes = []interface{}{
	(*Namespace)(nil),                // 0: rpkm66.checkin.namespace.v1.Namespace
	(*GetAllNamespacesRequest)(nil),  // 1: rpkm66.checkin.namespace.v1.GetAllNamespacesRequest
	(*GetAllNamespacesResponse)(nil), // 2: rpkm66.checkin.namespace.v1.GetAllNamespacesResponse
}
var file_rpkm66_checkin_namespace_v1_namespace_proto_depIdxs = []int32{
	0, // 0: rpkm66.checkin.namespace.v1.GetAllNamespacesResponse.namespaces:type_name -> rpkm66.checkin.namespace.v1.Namespace
	1, // 1: rpkm66.checkin.namespace.v1.NamespaceService.GetAllNamespaces:input_type -> rpkm66.checkin.namespace.v1.GetAllNamespacesRequest
	2, // 2: rpkm66.checkin.namespace.v1.NamespaceService.GetAllNamespaces:output_type -> rpkm66.checkin.namespace.v1.GetAllNamespacesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpkm66_checkin_namespace_v1_namespace_proto_init() }
func file_rpkm66_checkin_namespace_v1_namespace_proto_init() {
	if File_rpkm66_checkin_namespace_v1_namespace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpkm66_checkin_namespace_v1_namespace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Namespace); i {
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
		file_rpkm66_checkin_namespace_v1_namespace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllNamespacesRequest); i {
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
		file_rpkm66_checkin_namespace_v1_namespace_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllNamespacesResponse); i {
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
			RawDescriptor: file_rpkm66_checkin_namespace_v1_namespace_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpkm66_checkin_namespace_v1_namespace_proto_goTypes,
		DependencyIndexes: file_rpkm66_checkin_namespace_v1_namespace_proto_depIdxs,
		MessageInfos:      file_rpkm66_checkin_namespace_v1_namespace_proto_msgTypes,
	}.Build()
	File_rpkm66_checkin_namespace_v1_namespace_proto = out.File
	file_rpkm66_checkin_namespace_v1_namespace_proto_rawDesc = nil
	file_rpkm66_checkin_namespace_v1_namespace_proto_goTypes = nil
	file_rpkm66_checkin_namespace_v1_namespace_proto_depIdxs = nil
}
