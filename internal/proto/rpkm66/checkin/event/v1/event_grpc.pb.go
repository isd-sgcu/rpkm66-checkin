// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: rpkm66/checkin/event/v1/event.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	EventService_GetAllEvents_FullMethodName           = "/rpkm66.checkin.event.v1.EventService/GetAllEvents"
	EventService_GetEventById_FullMethodName           = "/rpkm66.checkin.event.v1.EventService/GetEventById"
	EventService_GetEventsByUserId_FullMethodName      = "/rpkm66.checkin.event.v1.EventService/GetEventsByUserId"
	EventService_GetEventsByNamespaceId_FullMethodName = "/rpkm66.checkin.event.v1.EventService/GetEventsByNamespaceId"
)

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventServiceClient interface {
	GetAllEvents(ctx context.Context, in *GetAllEventsRequest, opts ...grpc.CallOption) (*GetAllEventsResponse, error)
	GetEventById(ctx context.Context, in *GetEventByEventIdRequest, opts ...grpc.CallOption) (*GetEventByEventIdResponse, error)
	GetEventsByUserId(ctx context.Context, in *GetEventsByUserIdRequest, opts ...grpc.CallOption) (*GetEventsByUserIdResponse, error)
	GetEventsByNamespaceId(ctx context.Context, in *GetEventsByNamespaceRequestId, opts ...grpc.CallOption) (*GetEventsByNamespaceResponseId, error)
}

type eventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventServiceClient(cc grpc.ClientConnInterface) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) GetAllEvents(ctx context.Context, in *GetAllEventsRequest, opts ...grpc.CallOption) (*GetAllEventsResponse, error) {
	out := new(GetAllEventsResponse)
	err := c.cc.Invoke(ctx, EventService_GetAllEvents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) GetEventById(ctx context.Context, in *GetEventByEventIdRequest, opts ...grpc.CallOption) (*GetEventByEventIdResponse, error) {
	out := new(GetEventByEventIdResponse)
	err := c.cc.Invoke(ctx, EventService_GetEventById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) GetEventsByUserId(ctx context.Context, in *GetEventsByUserIdRequest, opts ...grpc.CallOption) (*GetEventsByUserIdResponse, error) {
	out := new(GetEventsByUserIdResponse)
	err := c.cc.Invoke(ctx, EventService_GetEventsByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) GetEventsByNamespaceId(ctx context.Context, in *GetEventsByNamespaceRequestId, opts ...grpc.CallOption) (*GetEventsByNamespaceResponseId, error) {
	out := new(GetEventsByNamespaceResponseId)
	err := c.cc.Invoke(ctx, EventService_GetEventsByNamespaceId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
// All implementations must embed UnimplementedEventServiceServer
// for forward compatibility
type EventServiceServer interface {
	GetAllEvents(context.Context, *GetAllEventsRequest) (*GetAllEventsResponse, error)
	GetEventById(context.Context, *GetEventByEventIdRequest) (*GetEventByEventIdResponse, error)
	GetEventsByUserId(context.Context, *GetEventsByUserIdRequest) (*GetEventsByUserIdResponse, error)
	GetEventsByNamespaceId(context.Context, *GetEventsByNamespaceRequestId) (*GetEventsByNamespaceResponseId, error)
	mustEmbedUnimplementedEventServiceServer()
}

// UnimplementedEventServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEventServiceServer struct {
}

func (UnimplementedEventServiceServer) GetAllEvents(context.Context, *GetAllEventsRequest) (*GetAllEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllEvents not implemented")
}
func (UnimplementedEventServiceServer) GetEventById(context.Context, *GetEventByEventIdRequest) (*GetEventByEventIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventById not implemented")
}
func (UnimplementedEventServiceServer) GetEventsByUserId(context.Context, *GetEventsByUserIdRequest) (*GetEventsByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventsByUserId not implemented")
}
func (UnimplementedEventServiceServer) GetEventsByNamespaceId(context.Context, *GetEventsByNamespaceRequestId) (*GetEventsByNamespaceResponseId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventsByNamespaceId not implemented")
}
func (UnimplementedEventServiceServer) mustEmbedUnimplementedEventServiceServer() {}

// UnsafeEventServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventServiceServer will
// result in compilation errors.
type UnsafeEventServiceServer interface {
	mustEmbedUnimplementedEventServiceServer()
}

func RegisterEventServiceServer(s grpc.ServiceRegistrar, srv EventServiceServer) {
	s.RegisterService(&EventService_ServiceDesc, srv)
}

func _EventService_GetAllEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetAllEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_GetAllEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetAllEvents(ctx, req.(*GetAllEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_GetEventById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventByEventIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetEventById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_GetEventById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetEventById(ctx, req.(*GetEventByEventIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_GetEventsByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetEventsByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_GetEventsByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetEventsByUserId(ctx, req.(*GetEventsByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_GetEventsByNamespaceId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsByNamespaceRequestId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetEventsByNamespaceId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventService_GetEventsByNamespaceId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetEventsByNamespaceId(ctx, req.(*GetEventsByNamespaceRequestId))
	}
	return interceptor(ctx, in, info, handler)
}

// EventService_ServiceDesc is the grpc.ServiceDesc for EventService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpkm66.checkin.event.v1.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllEvents",
			Handler:    _EventService_GetAllEvents_Handler,
		},
		{
			MethodName: "GetEventById",
			Handler:    _EventService_GetEventById_Handler,
		},
		{
			MethodName: "GetEventsByUserId",
			Handler:    _EventService_GetEventsByUserId_Handler,
		},
		{
			MethodName: "GetEventsByNamespaceId",
			Handler:    _EventService_GetEventsByNamespaceId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpkm66/checkin/event/v1/event.proto",
}