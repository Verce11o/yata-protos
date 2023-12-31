// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: tweets/tweets.proto

package tweets

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
	Tweets_CreateTweet_FullMethodName  = "/tweets.Tweets/CreateTweet"
	Tweets_GetTweet_FullMethodName     = "/tweets.Tweets/GetTweet"
	Tweets_GetAllTweets_FullMethodName = "/tweets.Tweets/GetAllTweets"
	Tweets_UpdateTweet_FullMethodName  = "/tweets.Tweets/UpdateTweet"
	Tweets_DeleteTweet_FullMethodName  = "/tweets.Tweets/DeleteTweet"
)

// TweetsClient is the client API for Tweets service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TweetsClient interface {
	CreateTweet(ctx context.Context, in *CreateTweetRequest, opts ...grpc.CallOption) (*CreateTweetResponse, error)
	GetTweet(ctx context.Context, in *GetTweetRequest, opts ...grpc.CallOption) (*Tweet, error)
	GetAllTweets(ctx context.Context, in *GetAllTweetsRequest, opts ...grpc.CallOption) (*GetAllTweetsResponse, error)
	UpdateTweet(ctx context.Context, in *UpdateTweetRequest, opts ...grpc.CallOption) (*Tweet, error)
	DeleteTweet(ctx context.Context, in *DeleteTweetRequest, opts ...grpc.CallOption) (*DeleteTweetResponse, error)
}

type tweetsClient struct {
	cc grpc.ClientConnInterface
}

func NewTweetsClient(cc grpc.ClientConnInterface) TweetsClient {
	return &tweetsClient{cc}
}

func (c *tweetsClient) CreateTweet(ctx context.Context, in *CreateTweetRequest, opts ...grpc.CallOption) (*CreateTweetResponse, error) {
	out := new(CreateTweetResponse)
	err := c.cc.Invoke(ctx, Tweets_CreateTweet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetsClient) GetTweet(ctx context.Context, in *GetTweetRequest, opts ...grpc.CallOption) (*Tweet, error) {
	out := new(Tweet)
	err := c.cc.Invoke(ctx, Tweets_GetTweet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetsClient) GetAllTweets(ctx context.Context, in *GetAllTweetsRequest, opts ...grpc.CallOption) (*GetAllTweetsResponse, error) {
	out := new(GetAllTweetsResponse)
	err := c.cc.Invoke(ctx, Tweets_GetAllTweets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetsClient) UpdateTweet(ctx context.Context, in *UpdateTweetRequest, opts ...grpc.CallOption) (*Tweet, error) {
	out := new(Tweet)
	err := c.cc.Invoke(ctx, Tweets_UpdateTweet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetsClient) DeleteTweet(ctx context.Context, in *DeleteTweetRequest, opts ...grpc.CallOption) (*DeleteTweetResponse, error) {
	out := new(DeleteTweetResponse)
	err := c.cc.Invoke(ctx, Tweets_DeleteTweet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TweetsServer is the server API for Tweets service.
// All implementations must embed UnimplementedTweetsServer
// for forward compatibility
type TweetsServer interface {
	CreateTweet(context.Context, *CreateTweetRequest) (*CreateTweetResponse, error)
	GetTweet(context.Context, *GetTweetRequest) (*Tweet, error)
	GetAllTweets(context.Context, *GetAllTweetsRequest) (*GetAllTweetsResponse, error)
	UpdateTweet(context.Context, *UpdateTweetRequest) (*Tweet, error)
	DeleteTweet(context.Context, *DeleteTweetRequest) (*DeleteTweetResponse, error)
	mustEmbedUnimplementedTweetsServer()
}

// UnimplementedTweetsServer must be embedded to have forward compatible implementations.
type UnimplementedTweetsServer struct {
}

func (UnimplementedTweetsServer) CreateTweet(context.Context, *CreateTweetRequest) (*CreateTweetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTweet not implemented")
}
func (UnimplementedTweetsServer) GetTweet(context.Context, *GetTweetRequest) (*Tweet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTweet not implemented")
}
func (UnimplementedTweetsServer) GetAllTweets(context.Context, *GetAllTweetsRequest) (*GetAllTweetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTweets not implemented")
}
func (UnimplementedTweetsServer) UpdateTweet(context.Context, *UpdateTweetRequest) (*Tweet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTweet not implemented")
}
func (UnimplementedTweetsServer) DeleteTweet(context.Context, *DeleteTweetRequest) (*DeleteTweetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTweet not implemented")
}
func (UnimplementedTweetsServer) mustEmbedUnimplementedTweetsServer() {}

// UnsafeTweetsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TweetsServer will
// result in compilation errors.
type UnsafeTweetsServer interface {
	mustEmbedUnimplementedTweetsServer()
}

func RegisterTweetsServer(s grpc.ServiceRegistrar, srv TweetsServer) {
	s.RegisterService(&Tweets_ServiceDesc, srv)
}

func _Tweets_CreateTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetsServer).CreateTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tweets_CreateTweet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetsServer).CreateTweet(ctx, req.(*CreateTweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tweets_GetTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetsServer).GetTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tweets_GetTweet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetsServer).GetTweet(ctx, req.(*GetTweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tweets_GetAllTweets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllTweetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetsServer).GetAllTweets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tweets_GetAllTweets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetsServer).GetAllTweets(ctx, req.(*GetAllTweetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tweets_UpdateTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetsServer).UpdateTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tweets_UpdateTweet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetsServer).UpdateTweet(ctx, req.(*UpdateTweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tweets_DeleteTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetsServer).DeleteTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tweets_DeleteTweet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetsServer).DeleteTweet(ctx, req.(*DeleteTweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Tweets_ServiceDesc is the grpc.ServiceDesc for Tweets service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tweets_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tweets.Tweets",
	HandlerType: (*TweetsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTweet",
			Handler:    _Tweets_CreateTweet_Handler,
		},
		{
			MethodName: "GetTweet",
			Handler:    _Tweets_GetTweet_Handler,
		},
		{
			MethodName: "GetAllTweets",
			Handler:    _Tweets_GetAllTweets_Handler,
		},
		{
			MethodName: "UpdateTweet",
			Handler:    _Tweets_UpdateTweet_Handler,
		},
		{
			MethodName: "DeleteTweet",
			Handler:    _Tweets_DeleteTweet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tweets/tweets.proto",
}
