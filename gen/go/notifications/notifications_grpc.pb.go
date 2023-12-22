// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: notifications/notifications.proto

package notifications

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
	Notifications_SubscribeToUser_FullMethodName     = "/notifications.Notifications/SubscribeToUser"
	Notifications_UnSubscribeFromUser_FullMethodName = "/notifications.Notifications/UnSubscribeFromUser"
)

// NotificationsClient is the client API for Notifications service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationsClient interface {
	SubscribeToUser(ctx context.Context, in *SubscribeToUserRequest, opts ...grpc.CallOption) (*SubscribeToUserResponse, error)
	UnSubscribeFromUser(ctx context.Context, in *UnSubscribeFromUserRequest, opts ...grpc.CallOption) (*UnSubscribeFromUserResponse, error)
}

type notificationsClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationsClient(cc grpc.ClientConnInterface) NotificationsClient {
	return &notificationsClient{cc}
}

func (c *notificationsClient) SubscribeToUser(ctx context.Context, in *SubscribeToUserRequest, opts ...grpc.CallOption) (*SubscribeToUserResponse, error) {
	out := new(SubscribeToUserResponse)
	err := c.cc.Invoke(ctx, Notifications_SubscribeToUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) UnSubscribeFromUser(ctx context.Context, in *UnSubscribeFromUserRequest, opts ...grpc.CallOption) (*UnSubscribeFromUserResponse, error) {
	out := new(UnSubscribeFromUserResponse)
	err := c.cc.Invoke(ctx, Notifications_UnSubscribeFromUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationsServer is the server API for Notifications service.
// All implementations must embed UnimplementedNotificationsServer
// for forward compatibility
type NotificationsServer interface {
	SubscribeToUser(context.Context, *SubscribeToUserRequest) (*SubscribeToUserResponse, error)
	UnSubscribeFromUser(context.Context, *UnSubscribeFromUserRequest) (*UnSubscribeFromUserResponse, error)
	mustEmbedUnimplementedNotificationsServer()
}

// UnimplementedNotificationsServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationsServer struct {
}

func (UnimplementedNotificationsServer) SubscribeToUser(context.Context, *SubscribeToUserRequest) (*SubscribeToUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeToUser not implemented")
}
func (UnimplementedNotificationsServer) UnSubscribeFromUser(context.Context, *UnSubscribeFromUserRequest) (*UnSubscribeFromUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnSubscribeFromUser not implemented")
}
func (UnimplementedNotificationsServer) mustEmbedUnimplementedNotificationsServer() {}

// UnsafeNotificationsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationsServer will
// result in compilation errors.
type UnsafeNotificationsServer interface {
	mustEmbedUnimplementedNotificationsServer()
}

func RegisterNotificationsServer(s grpc.ServiceRegistrar, srv NotificationsServer) {
	s.RegisterService(&Notifications_ServiceDesc, srv)
}

func _Notifications_SubscribeToUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeToUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).SubscribeToUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notifications_SubscribeToUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).SubscribeToUser(ctx, req.(*SubscribeToUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_UnSubscribeFromUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnSubscribeFromUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).UnSubscribeFromUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notifications_UnSubscribeFromUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).UnSubscribeFromUser(ctx, req.(*UnSubscribeFromUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Notifications_ServiceDesc is the grpc.ServiceDesc for Notifications service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Notifications_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notifications.Notifications",
	HandlerType: (*NotificationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubscribeToUser",
			Handler:    _Notifications_SubscribeToUser_Handler,
		},
		{
			MethodName: "UnSubscribeFromUser",
			Handler:    _Notifications_UnSubscribeFromUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notifications/notifications.proto",
}
