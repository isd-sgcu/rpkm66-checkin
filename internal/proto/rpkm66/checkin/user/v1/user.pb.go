// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: rpkm66/checkin/user/v1/user.proto

package v1

import (
	v1 "github.com/isd-sgcu/rpkm66-checkin/internal/proto/rpkm66/checkin/event/v1"
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

// AddEvent
type AddEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AddEventRequest) Reset() {
	*x = AddEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_checkin_user_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEventRequest) ProtoMessage() {}

func (x *AddEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_checkin_user_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEventRequest.ProtoReflect.Descriptor instead.
func (*AddEventRequest) Descriptor() ([]byte, []int) {
	return file_rpkm66_checkin_user_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *AddEventRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddEventRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AddEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *v1.Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *AddEventResponse) Reset() {
	*x = AddEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_checkin_user_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEventResponse) ProtoMessage() {}

func (x *AddEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_checkin_user_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEventResponse.ProtoReflect.Descriptor instead.
func (*AddEventResponse) Descriptor() ([]byte, []int) {
	return file_rpkm66_checkin_user_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *AddEventResponse) GetEvent() *v1.Event {
	if x != nil {
		return x.Event
	}
	return nil
}

var File_rpkm66_checkin_user_v1_user_proto protoreflect.FileDescriptor

var file_rpkm66_checkin_user_v1_user_proto_rawDesc = []byte{
	0x0a, 0x21, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x69, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x23, 0x72, 0x70, 0x6b,
	0x6d, 0x36, 0x36, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x40, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x48, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x32, 0x6e, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x08, 0x41,
	0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36,
	0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69,
	0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16,
	0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpkm66_checkin_user_v1_user_proto_rawDescOnce sync.Once
	file_rpkm66_checkin_user_v1_user_proto_rawDescData = file_rpkm66_checkin_user_v1_user_proto_rawDesc
)

func file_rpkm66_checkin_user_v1_user_proto_rawDescGZIP() []byte {
	file_rpkm66_checkin_user_v1_user_proto_rawDescOnce.Do(func() {
		file_rpkm66_checkin_user_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpkm66_checkin_user_v1_user_proto_rawDescData)
	})
	return file_rpkm66_checkin_user_v1_user_proto_rawDescData
}

var file_rpkm66_checkin_user_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpkm66_checkin_user_v1_user_proto_goTypes = []interface{}{
	(*AddEventRequest)(nil),  // 0: rpkm66.checkin.user.v1.AddEventRequest
	(*AddEventResponse)(nil), // 1: rpkm66.checkin.user.v1.AddEventResponse
	(*v1.Event)(nil),         // 2: rpkm66.checkin.event.v1.Event
}
var file_rpkm66_checkin_user_v1_user_proto_depIdxs = []int32{
	2, // 0: rpkm66.checkin.user.v1.AddEventResponse.event:type_name -> rpkm66.checkin.event.v1.Event
	0, // 1: rpkm66.checkin.user.v1.UserService.AddEvent:input_type -> rpkm66.checkin.user.v1.AddEventRequest
	1, // 2: rpkm66.checkin.user.v1.UserService.AddEvent:output_type -> rpkm66.checkin.user.v1.AddEventResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpkm66_checkin_user_v1_user_proto_init() }
func file_rpkm66_checkin_user_v1_user_proto_init() {
	if File_rpkm66_checkin_user_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpkm66_checkin_user_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEventRequest); i {
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
		file_rpkm66_checkin_user_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEventResponse); i {
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
			RawDescriptor: file_rpkm66_checkin_user_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpkm66_checkin_user_v1_user_proto_goTypes,
		DependencyIndexes: file_rpkm66_checkin_user_v1_user_proto_depIdxs,
		MessageInfos:      file_rpkm66_checkin_user_v1_user_proto_msgTypes,
	}.Build()
	File_rpkm66_checkin_user_v1_user_proto = out.File
	file_rpkm66_checkin_user_v1_user_proto_rawDesc = nil
	file_rpkm66_checkin_user_v1_user_proto_goTypes = nil
	file_rpkm66_checkin_user_v1_user_proto_depIdxs = nil
}
