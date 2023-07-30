// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: rpkm66/checkin/event/v1/event.proto

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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId        string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	EventName      string `protobuf:"bytes,2,opt,name=event_name,json=eventName,proto3" json:"event_name,omitempty"`
	NamespaceId    string `protobuf:"bytes,3,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	AdditionalInfo string `protobuf:"bytes,4,opt,name=additional_info,json=additionalInfo,proto3" json:"additional_info,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_checkin_event_v1_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_checkin_event_v1_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_rpkm66_checkin_event_v1_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *Event) GetEventName() string {
	if x != nil {
		return x.EventName
	}
	return ""
}

func (x *Event) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

func (x *Event) GetAdditionalInfo() string {
	if x != nil {
		return x.AdditionalInfo
	}
	return ""
}

type UserEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event   *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	IsTaken bool   `protobuf:"varint,2,opt,name=is_taken,json=isTaken,proto3" json:"is_taken,omitempty"`
	TakenAt int64  `protobuf:"varint,3,opt,name=taken_at,json=takenAt,proto3" json:"taken_at,omitempty"`
}

func (x *UserEvent) Reset() {
	*x = UserEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_checkin_event_v1_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserEvent) ProtoMessage() {}

func (x *UserEvent) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_checkin_event_v1_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserEvent.ProtoReflect.Descriptor instead.
func (*UserEvent) Descriptor() ([]byte, []int) {
	return file_rpkm66_checkin_event_v1_event_proto_rawDescGZIP(), []int{1}
}

func (x *UserEvent) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *UserEvent) GetIsTaken() bool {
	if x != nil {
		return x.IsTaken
	}
	return false
}

func (x *UserEvent) GetTakenAt() int64 {
	if x != nil {
		return x.TakenAt
	}
	return 0
}

// GetAllEvents
type GetAllEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllEventsRequest) Reset() {
	*x = GetAllEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_checkin_event_v1_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllEventsRequest) ProtoMessage() {}

func (x *GetAllEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_checkin_event_v1_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllEventsRequest.ProtoReflect.Descriptor instead.
func (*GetAllEventsRequest) Descriptor() ([]byte, []int) {
	return file_rpkm66_checkin_event_v1_event_proto_rawDescGZIP(), []int{2}
}

type GetAllEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *GetAllEventsResponse) Reset() {
	*x = GetAllEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_checkin_event_v1_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllEventsResponse) ProtoMessage() {}

func (x *GetAllEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_checkin_event_v1_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllEventsResponse.ProtoReflect.Descriptor instead.
func (*GetAllEventsResponse) Descriptor() ([]byte, []int) {
	return file_rpkm66_checkin_event_v1_event_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllEventsResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

// GetEventByEventId
type GetEventByEventIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
}

func (x *GetEventByEventIdRequest) Reset() {
	*x = GetEventByEventIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_checkin_event_v1_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventByEventIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventByEventIdRequest) ProtoMessage() {}

func (x *GetEventByEventIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_checkin_event_v1_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventByEventIdRequest.ProtoReflect.Descriptor instead.
func (*GetEventByEventIdRequest) Descriptor() ([]byte, []int) {
	return file_rpkm66_checkin_event_v1_event_proto_rawDescGZIP(), []int{4}
}

func (x *GetEventByEventIdRequest) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

type GetEventByEventIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *GetEventByEventIdResponse) Reset() {
	*x = GetEventByEventIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_checkin_event_v1_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventByEventIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventByEventIdResponse) ProtoMessage() {}

func (x *GetEventByEventIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_checkin_event_v1_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventByEventIdResponse.ProtoReflect.Descriptor instead.
func (*GetEventByEventIdResponse) Descriptor() ([]byte, []int) {
	return file_rpkm66_checkin_event_v1_event_proto_rawDescGZIP(), []int{5}
}

func (x *GetEventByEventIdResponse) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

// GetEventsByUserId
type GetEventsByUserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetEventsByUserIdRequest) Reset() {
	*x = GetEventsByUserIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_checkin_event_v1_event_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsByUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsByUserIdRequest) ProtoMessage() {}

func (x *GetEventsByUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_checkin_event_v1_event_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsByUserIdRequest.ProtoReflect.Descriptor instead.
func (*GetEventsByUserIdRequest) Descriptor() ([]byte, []int) {
	return file_rpkm66_checkin_event_v1_event_proto_rawDescGZIP(), []int{6}
}

func (x *GetEventsByUserIdRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetEventsByUserIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*UserEvent `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *GetEventsByUserIdResponse) Reset() {
	*x = GetEventsByUserIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_checkin_event_v1_event_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsByUserIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsByUserIdResponse) ProtoMessage() {}

func (x *GetEventsByUserIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_checkin_event_v1_event_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsByUserIdResponse.ProtoReflect.Descriptor instead.
func (*GetEventsByUserIdResponse) Descriptor() ([]byte, []int) {
	return file_rpkm66_checkin_event_v1_event_proto_rawDescGZIP(), []int{7}
}

func (x *GetEventsByUserIdResponse) GetEvents() []*UserEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

// GetEventsByNamespaceId
type GetEventsByNamespaceIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NamespaceId string `protobuf:"bytes,1,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
}

func (x *GetEventsByNamespaceIdRequest) Reset() {
	*x = GetEventsByNamespaceIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_checkin_event_v1_event_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsByNamespaceIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsByNamespaceIdRequest) ProtoMessage() {}

func (x *GetEventsByNamespaceIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_checkin_event_v1_event_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsByNamespaceIdRequest.ProtoReflect.Descriptor instead.
func (*GetEventsByNamespaceIdRequest) Descriptor() ([]byte, []int) {
	return file_rpkm66_checkin_event_v1_event_proto_rawDescGZIP(), []int{8}
}

func (x *GetEventsByNamespaceIdRequest) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

type GetEventsByNamespaceIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *GetEventsByNamespaceIdResponse) Reset() {
	*x = GetEventsByNamespaceIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpkm66_checkin_event_v1_event_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsByNamespaceIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsByNamespaceIdResponse) ProtoMessage() {}

func (x *GetEventsByNamespaceIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpkm66_checkin_event_v1_event_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsByNamespaceIdResponse.ProtoReflect.Descriptor instead.
func (*GetEventsByNamespaceIdResponse) Descriptor() ([]byte, []int) {
	return file_rpkm66_checkin_event_v1_event_proto_rawDescGZIP(), []int{9}
}

func (x *GetEventsByNamespaceIdResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_rpkm66_checkin_event_v1_event_proto protoreflect.FileDescriptor

var file_rpkm66_checkin_event_v1_event_proto_rawDesc = []byte{
	0x0a, 0x23, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x8d,
	0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x77,
	0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x70, 0x6b,
	0x6d, 0x36, 0x36, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08,
	0x74, 0x61, 0x6b, 0x65, 0x6e, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x74, 0x61, 0x6b, 0x65, 0x6e, 0x41, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4e,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x35,
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x42, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x69, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x33, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x57, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x70, 0x6b,
	0x6d, 0x36, 0x36, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x42, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0x58, 0x0a, 0x1e, 0x47, 0x65,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x06,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72,
	0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x32, 0x8a, 0x04, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2c, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x7c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x42, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x31, 0x2e, 0x72, 0x70, 0x6b, 0x6d,
	0x36, 0x36, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x72,
	0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42,
	0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x7f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x31, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36,
	0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x72, 0x70, 0x6b,
	0x6d, 0x36, 0x36, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03,
	0x88, 0x02, 0x01, 0x12, 0x8b, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x36,
	0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2e,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x19, 0x5a, 0x17, 0x72, 0x70, 0x6b, 0x6d, 0x36, 0x36, 0x2f, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x69, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpkm66_checkin_event_v1_event_proto_rawDescOnce sync.Once
	file_rpkm66_checkin_event_v1_event_proto_rawDescData = file_rpkm66_checkin_event_v1_event_proto_rawDesc
)

func file_rpkm66_checkin_event_v1_event_proto_rawDescGZIP() []byte {
	file_rpkm66_checkin_event_v1_event_proto_rawDescOnce.Do(func() {
		file_rpkm66_checkin_event_v1_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpkm66_checkin_event_v1_event_proto_rawDescData)
	})
	return file_rpkm66_checkin_event_v1_event_proto_rawDescData
}

var file_rpkm66_checkin_event_v1_event_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_rpkm66_checkin_event_v1_event_proto_goTypes = []interface{}{
	(*Event)(nil),                          // 0: rpkm66.checkin.event.v1.Event
	(*UserEvent)(nil),                      // 1: rpkm66.checkin.event.v1.UserEvent
	(*GetAllEventsRequest)(nil),            // 2: rpkm66.checkin.event.v1.GetAllEventsRequest
	(*GetAllEventsResponse)(nil),           // 3: rpkm66.checkin.event.v1.GetAllEventsResponse
	(*GetEventByEventIdRequest)(nil),       // 4: rpkm66.checkin.event.v1.GetEventByEventIdRequest
	(*GetEventByEventIdResponse)(nil),      // 5: rpkm66.checkin.event.v1.GetEventByEventIdResponse
	(*GetEventsByUserIdRequest)(nil),       // 6: rpkm66.checkin.event.v1.GetEventsByUserIdRequest
	(*GetEventsByUserIdResponse)(nil),      // 7: rpkm66.checkin.event.v1.GetEventsByUserIdResponse
	(*GetEventsByNamespaceIdRequest)(nil),  // 8: rpkm66.checkin.event.v1.GetEventsByNamespaceIdRequest
	(*GetEventsByNamespaceIdResponse)(nil), // 9: rpkm66.checkin.event.v1.GetEventsByNamespaceIdResponse
}
var file_rpkm66_checkin_event_v1_event_proto_depIdxs = []int32{
	0, // 0: rpkm66.checkin.event.v1.UserEvent.event:type_name -> rpkm66.checkin.event.v1.Event
	0, // 1: rpkm66.checkin.event.v1.GetAllEventsResponse.events:type_name -> rpkm66.checkin.event.v1.Event
	0, // 2: rpkm66.checkin.event.v1.GetEventByEventIdResponse.event:type_name -> rpkm66.checkin.event.v1.Event
	1, // 3: rpkm66.checkin.event.v1.GetEventsByUserIdResponse.events:type_name -> rpkm66.checkin.event.v1.UserEvent
	0, // 4: rpkm66.checkin.event.v1.GetEventsByNamespaceIdResponse.events:type_name -> rpkm66.checkin.event.v1.Event
	2, // 5: rpkm66.checkin.event.v1.EventService.GetAllEvents:input_type -> rpkm66.checkin.event.v1.GetAllEventsRequest
	4, // 6: rpkm66.checkin.event.v1.EventService.GetEventByEventId:input_type -> rpkm66.checkin.event.v1.GetEventByEventIdRequest
	6, // 7: rpkm66.checkin.event.v1.EventService.GetEventsByUserId:input_type -> rpkm66.checkin.event.v1.GetEventsByUserIdRequest
	8, // 8: rpkm66.checkin.event.v1.EventService.GetEventsByNamespaceId:input_type -> rpkm66.checkin.event.v1.GetEventsByNamespaceIdRequest
	3, // 9: rpkm66.checkin.event.v1.EventService.GetAllEvents:output_type -> rpkm66.checkin.event.v1.GetAllEventsResponse
	5, // 10: rpkm66.checkin.event.v1.EventService.GetEventByEventId:output_type -> rpkm66.checkin.event.v1.GetEventByEventIdResponse
	7, // 11: rpkm66.checkin.event.v1.EventService.GetEventsByUserId:output_type -> rpkm66.checkin.event.v1.GetEventsByUserIdResponse
	9, // 12: rpkm66.checkin.event.v1.EventService.GetEventsByNamespaceId:output_type -> rpkm66.checkin.event.v1.GetEventsByNamespaceIdResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_rpkm66_checkin_event_v1_event_proto_init() }
func file_rpkm66_checkin_event_v1_event_proto_init() {
	if File_rpkm66_checkin_event_v1_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpkm66_checkin_event_v1_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_rpkm66_checkin_event_v1_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserEvent); i {
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
		file_rpkm66_checkin_event_v1_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllEventsRequest); i {
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
		file_rpkm66_checkin_event_v1_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllEventsResponse); i {
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
		file_rpkm66_checkin_event_v1_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventByEventIdRequest); i {
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
		file_rpkm66_checkin_event_v1_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventByEventIdResponse); i {
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
		file_rpkm66_checkin_event_v1_event_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsByUserIdRequest); i {
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
		file_rpkm66_checkin_event_v1_event_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsByUserIdResponse); i {
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
		file_rpkm66_checkin_event_v1_event_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsByNamespaceIdRequest); i {
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
		file_rpkm66_checkin_event_v1_event_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsByNamespaceIdResponse); i {
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
			RawDescriptor: file_rpkm66_checkin_event_v1_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpkm66_checkin_event_v1_event_proto_goTypes,
		DependencyIndexes: file_rpkm66_checkin_event_v1_event_proto_depIdxs,
		MessageInfos:      file_rpkm66_checkin_event_v1_event_proto_msgTypes,
	}.Build()
	File_rpkm66_checkin_event_v1_event_proto = out.File
	file_rpkm66_checkin_event_v1_event_proto_rawDesc = nil
	file_rpkm66_checkin_event_v1_event_proto_goTypes = nil
	file_rpkm66_checkin_event_v1_event_proto_depIdxs = nil
}
