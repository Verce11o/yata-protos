// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: tweets/tweets.proto

package tweets

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

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk       []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
	ContentType string `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweets_tweets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_tweets_tweets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_tweets_tweets_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

func (x *Image) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

type Tweet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TweetId string `protobuf:"bytes,2,opt,name=tweet_id,json=tweetId,proto3" json:"tweet_id,omitempty"`
	Text    string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Image   *Image `protobuf:"bytes,4,opt,name=image,proto3,oneof" json:"image,omitempty"`
}

func (x *Tweet) Reset() {
	*x = Tweet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweets_tweets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tweet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tweet) ProtoMessage() {}

func (x *Tweet) ProtoReflect() protoreflect.Message {
	mi := &file_tweets_tweets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tweet.ProtoReflect.Descriptor instead.
func (*Tweet) Descriptor() ([]byte, []int) {
	return file_tweets_tweets_proto_rawDescGZIP(), []int{1}
}

func (x *Tweet) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Tweet) GetTweetId() string {
	if x != nil {
		return x.TweetId
	}
	return ""
}

func (x *Tweet) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Tweet) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

type CreateTweetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Text   string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Image  *Image `protobuf:"bytes,3,opt,name=image,proto3,oneof" json:"image,omitempty"`
}

func (x *CreateTweetRequest) Reset() {
	*x = CreateTweetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweets_tweets_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTweetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTweetRequest) ProtoMessage() {}

func (x *CreateTweetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tweets_tweets_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTweetRequest.ProtoReflect.Descriptor instead.
func (*CreateTweetRequest) Descriptor() ([]byte, []int) {
	return file_tweets_tweets_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTweetRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateTweetRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *CreateTweetRequest) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

type CreateTweetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TweetId string `protobuf:"bytes,1,opt,name=tweet_id,json=tweetId,proto3" json:"tweet_id,omitempty"`
}

func (x *CreateTweetResponse) Reset() {
	*x = CreateTweetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweets_tweets_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTweetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTweetResponse) ProtoMessage() {}

func (x *CreateTweetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tweets_tweets_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTweetResponse.ProtoReflect.Descriptor instead.
func (*CreateTweetResponse) Descriptor() ([]byte, []int) {
	return file_tweets_tweets_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTweetResponse) GetTweetId() string {
	if x != nil {
		return x.TweetId
	}
	return ""
}

type GetTweetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TweetId string `protobuf:"bytes,1,opt,name=tweet_id,json=tweetId,proto3" json:"tweet_id,omitempty"`
}

func (x *GetTweetRequest) Reset() {
	*x = GetTweetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweets_tweets_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTweetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTweetRequest) ProtoMessage() {}

func (x *GetTweetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tweets_tweets_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTweetRequest.ProtoReflect.Descriptor instead.
func (*GetTweetRequest) Descriptor() ([]byte, []int) {
	return file_tweets_tweets_proto_rawDescGZIP(), []int{4}
}

func (x *GetTweetRequest) GetTweetId() string {
	if x != nil {
		return x.TweetId
	}
	return ""
}

type UpdateTweetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TweetId string `protobuf:"bytes,2,opt,name=tweet_id,json=tweetId,proto3" json:"tweet_id,omitempty"`
	Text    string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Image   *Image `protobuf:"bytes,4,opt,name=image,proto3,oneof" json:"image,omitempty"`
}

func (x *UpdateTweetRequest) Reset() {
	*x = UpdateTweetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweets_tweets_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTweetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTweetRequest) ProtoMessage() {}

func (x *UpdateTweetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tweets_tweets_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTweetRequest.ProtoReflect.Descriptor instead.
func (*UpdateTweetRequest) Descriptor() ([]byte, []int) {
	return file_tweets_tweets_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateTweetRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateTweetRequest) GetTweetId() string {
	if x != nil {
		return x.TweetId
	}
	return ""
}

func (x *UpdateTweetRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *UpdateTweetRequest) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

type DeleteTweetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TweetId string `protobuf:"bytes,2,opt,name=tweet_id,json=tweetId,proto3" json:"tweet_id,omitempty"`
}

func (x *DeleteTweetRequest) Reset() {
	*x = DeleteTweetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweets_tweets_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTweetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTweetRequest) ProtoMessage() {}

func (x *DeleteTweetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tweets_tweets_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTweetRequest.ProtoReflect.Descriptor instead.
func (*DeleteTweetRequest) Descriptor() ([]byte, []int) {
	return file_tweets_tweets_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteTweetRequest) GetTweetId() string {
	if x != nil {
		return x.TweetId
	}
	return ""
}

type DeleteTweetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTweetResponse) Reset() {
	*x = DeleteTweetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tweets_tweets_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTweetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTweetResponse) ProtoMessage() {}

func (x *DeleteTweetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tweets_tweets_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTweetResponse.ProtoReflect.Descriptor instead.
func (*DeleteTweetResponse) Descriptor() ([]byte, []int) {
	return file_tweets_tweets_proto_rawDescGZIP(), []int{7}
}

var File_tweets_tweets_proto protoreflect.FileDescriptor

var file_tweets_tweets_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x77, 0x65, 0x65, 0x74, 0x73, 0x2f, 0x74, 0x77, 0x65, 0x65, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x40, 0x0a, 0x05, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x81, 0x01,
	0x0a, 0x05, 0x54, 0x77, 0x65, 0x65, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x77, 0x65, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x77, 0x65, 0x65, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x26, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x22, 0x73, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x77, 0x65, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x26, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x48, 0x00, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x30, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x74, 0x77, 0x65, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x77, 0x65, 0x65, 0x74, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54,
	0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74,
	0x77, 0x65, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74,
	0x77, 0x65, 0x65, 0x74, 0x49, 0x64, 0x22, 0x8e, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x77, 0x65, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x77, 0x65, 0x65, 0x74, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x26, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x48, 0x00, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x2f, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x74, 0x77, 0x65, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x77, 0x65, 0x65, 0x74, 0x49, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xf6, 0x01, 0x0a, 0x06, 0x54, 0x77, 0x65, 0x65, 0x74, 0x73, 0x12, 0x42, 0x0a, 0x0b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x77, 0x65, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x12, 0x15, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x54, 0x77, 0x65, 0x65, 0x74, 0x12, 0x34,
	0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x77, 0x65, 0x65, 0x74, 0x12, 0x18, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x77, 0x65, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x54,
	0x77, 0x65, 0x65, 0x74, 0x12, 0x42, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x77,
	0x65, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x77, 0x65, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x18, 0x5a, 0x16, 0x76, 0x65, 0x72, 0x63,
	0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x74, 0x77, 0x65, 0x65, 0x74, 0x73, 0x3b, 0x74, 0x77, 0x65, 0x65,
	0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tweets_tweets_proto_rawDescOnce sync.Once
	file_tweets_tweets_proto_rawDescData = file_tweets_tweets_proto_rawDesc
)

func file_tweets_tweets_proto_rawDescGZIP() []byte {
	file_tweets_tweets_proto_rawDescOnce.Do(func() {
		file_tweets_tweets_proto_rawDescData = protoimpl.X.CompressGZIP(file_tweets_tweets_proto_rawDescData)
	})
	return file_tweets_tweets_proto_rawDescData
}

var file_tweets_tweets_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_tweets_tweets_proto_goTypes = []interface{}{
	(*Image)(nil),               // 0: auth.Image
	(*Tweet)(nil),               // 1: auth.Tweet
	(*CreateTweetRequest)(nil),  // 2: auth.CreateTweetRequest
	(*CreateTweetResponse)(nil), // 3: auth.CreateTweetResponse
	(*GetTweetRequest)(nil),     // 4: auth.GetTweetRequest
	(*UpdateTweetRequest)(nil),  // 5: auth.UpdateTweetRequest
	(*DeleteTweetRequest)(nil),  // 6: auth.DeleteTweetRequest
	(*DeleteTweetResponse)(nil), // 7: auth.DeleteTweetResponse
}
var file_tweets_tweets_proto_depIdxs = []int32{
	0, // 0: auth.Tweet.image:type_name -> auth.Image
	0, // 1: auth.CreateTweetRequest.image:type_name -> auth.Image
	0, // 2: auth.UpdateTweetRequest.image:type_name -> auth.Image
	2, // 3: auth.Tweets.CreateTweet:input_type -> auth.CreateTweetRequest
	4, // 4: auth.Tweets.GetTweet:input_type -> auth.GetTweetRequest
	5, // 5: auth.Tweets.UpdateTweet:input_type -> auth.UpdateTweetRequest
	6, // 6: auth.Tweets.DeleteTweet:input_type -> auth.DeleteTweetRequest
	3, // 7: auth.Tweets.CreateTweet:output_type -> auth.CreateTweetResponse
	1, // 8: auth.Tweets.GetTweet:output_type -> auth.Tweet
	1, // 9: auth.Tweets.UpdateTweet:output_type -> auth.Tweet
	7, // 10: auth.Tweets.DeleteTweet:output_type -> auth.DeleteTweetResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tweets_tweets_proto_init() }
func file_tweets_tweets_proto_init() {
	if File_tweets_tweets_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tweets_tweets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_tweets_tweets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tweet); i {
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
		file_tweets_tweets_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTweetRequest); i {
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
		file_tweets_tweets_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTweetResponse); i {
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
		file_tweets_tweets_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTweetRequest); i {
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
		file_tweets_tweets_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTweetRequest); i {
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
		file_tweets_tweets_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTweetRequest); i {
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
		file_tweets_tweets_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTweetResponse); i {
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
	file_tweets_tweets_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_tweets_tweets_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_tweets_tweets_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tweets_tweets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tweets_tweets_proto_goTypes,
		DependencyIndexes: file_tweets_tweets_proto_depIdxs,
		MessageInfos:      file_tweets_tweets_proto_msgTypes,
	}.Build()
	File_tweets_tweets_proto = out.File
	file_tweets_tweets_proto_rawDesc = nil
	file_tweets_tweets_proto_goTypes = nil
	file_tweets_tweets_proto_depIdxs = nil
}
