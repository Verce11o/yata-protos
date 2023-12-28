// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: notifications/notifications.proto

package notifications

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotificationId   string                 `protobuf:"bytes,1,opt,name=notification_id,json=notificationId,proto3" json:"notification_id,omitempty"`
	UserId           string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SenderId         string                 `protobuf:"bytes,3,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	Read             bool                   `protobuf:"varint,4,opt,name=read,proto3" json:"read,omitempty"`
	LicenseExpiredAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=license_expired_at,json=licenseExpiredAt,proto3" json:"license_expired_at,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_notifications_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_notifications_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_notifications_notifications_proto_rawDescGZIP(), []int{0}
}

func (x *Notification) GetNotificationId() string {
	if x != nil {
		return x.NotificationId
	}
	return ""
}

func (x *Notification) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Notification) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *Notification) GetRead() bool {
	if x != nil {
		return x.Read
	}
	return false
}

func (x *Notification) GetLicenseExpiredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LicenseExpiredAt
	}
	return nil
}

type GetNotificationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetNotificationsRequest) Reset() {
	*x = GetNotificationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_notifications_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotificationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotificationsRequest) ProtoMessage() {}

func (x *GetNotificationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_notifications_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotificationsRequest.ProtoReflect.Descriptor instead.
func (*GetNotificationsRequest) Descriptor() ([]byte, []int) {
	return file_notifications_notifications_proto_rawDescGZIP(), []int{1}
}

func (x *GetNotificationsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetNotificationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notifications []*Notification `protobuf:"bytes,1,rep,name=notifications,proto3" json:"notifications,omitempty"`
}

func (x *GetNotificationsResponse) Reset() {
	*x = GetNotificationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_notifications_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotificationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotificationsResponse) ProtoMessage() {}

func (x *GetNotificationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_notifications_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotificationsResponse.ProtoReflect.Descriptor instead.
func (*GetNotificationsResponse) Descriptor() ([]byte, []int) {
	return file_notifications_notifications_proto_rawDescGZIP(), []int{2}
}

func (x *GetNotificationsResponse) GetNotifications() []*Notification {
	if x != nil {
		return x.Notifications
	}
	return nil
}

type MarkNotificationAsReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	NotificationId string `protobuf:"bytes,2,opt,name=notification_id,json=notificationId,proto3" json:"notification_id,omitempty"`
}

func (x *MarkNotificationAsReadRequest) Reset() {
	*x = MarkNotificationAsReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_notifications_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkNotificationAsReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkNotificationAsReadRequest) ProtoMessage() {}

func (x *MarkNotificationAsReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_notifications_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkNotificationAsReadRequest.ProtoReflect.Descriptor instead.
func (*MarkNotificationAsReadRequest) Descriptor() ([]byte, []int) {
	return file_notifications_notifications_proto_rawDescGZIP(), []int{3}
}

func (x *MarkNotificationAsReadRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *MarkNotificationAsReadRequest) GetNotificationId() string {
	if x != nil {
		return x.NotificationId
	}
	return ""
}

type MarkNotificationAsReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MarkNotificationAsReadResponse) Reset() {
	*x = MarkNotificationAsReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_notifications_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkNotificationAsReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkNotificationAsReadResponse) ProtoMessage() {}

func (x *MarkNotificationAsReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_notifications_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkNotificationAsReadResponse.ProtoReflect.Descriptor instead.
func (*MarkNotificationAsReadResponse) Descriptor() ([]byte, []int) {
	return file_notifications_notifications_proto_rawDescGZIP(), []int{4}
}

type ReadAllNotificationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ReadAllNotificationsRequest) Reset() {
	*x = ReadAllNotificationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_notifications_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAllNotificationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAllNotificationsRequest) ProtoMessage() {}

func (x *ReadAllNotificationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_notifications_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAllNotificationsRequest.ProtoReflect.Descriptor instead.
func (*ReadAllNotificationsRequest) Descriptor() ([]byte, []int) {
	return file_notifications_notifications_proto_rawDescGZIP(), []int{5}
}

func (x *ReadAllNotificationsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ReadAllNotificationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReadAllNotificationsResponse) Reset() {
	*x = ReadAllNotificationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_notifications_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAllNotificationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAllNotificationsResponse) ProtoMessage() {}

func (x *ReadAllNotificationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_notifications_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAllNotificationsResponse.ProtoReflect.Descriptor instead.
func (*ReadAllNotificationsResponse) Descriptor() ([]byte, []int) {
	return file_notifications_notifications_proto_rawDescGZIP(), []int{6}
}

type SubscribeToUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ToUserId string `protobuf:"bytes,2,opt,name=to_user_id,json=toUserId,proto3" json:"to_user_id,omitempty"`
}

func (x *SubscribeToUserRequest) Reset() {
	*x = SubscribeToUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_notifications_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeToUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeToUserRequest) ProtoMessage() {}

func (x *SubscribeToUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_notifications_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeToUserRequest.ProtoReflect.Descriptor instead.
func (*SubscribeToUserRequest) Descriptor() ([]byte, []int) {
	return file_notifications_notifications_proto_rawDescGZIP(), []int{7}
}

func (x *SubscribeToUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SubscribeToUserRequest) GetToUserId() string {
	if x != nil {
		return x.ToUserId
	}
	return ""
}

type SubscribeToUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscribeToUserResponse) Reset() {
	*x = SubscribeToUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_notifications_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeToUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeToUserResponse) ProtoMessage() {}

func (x *SubscribeToUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_notifications_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeToUserResponse.ProtoReflect.Descriptor instead.
func (*SubscribeToUserResponse) Descriptor() ([]byte, []int) {
	return file_notifications_notifications_proto_rawDescGZIP(), []int{8}
}

type UnSubscribeFromUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ToUserId string `protobuf:"bytes,2,opt,name=to_user_id,json=toUserId,proto3" json:"to_user_id,omitempty"`
}

func (x *UnSubscribeFromUserRequest) Reset() {
	*x = UnSubscribeFromUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_notifications_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnSubscribeFromUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnSubscribeFromUserRequest) ProtoMessage() {}

func (x *UnSubscribeFromUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_notifications_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnSubscribeFromUserRequest.ProtoReflect.Descriptor instead.
func (*UnSubscribeFromUserRequest) Descriptor() ([]byte, []int) {
	return file_notifications_notifications_proto_rawDescGZIP(), []int{9}
}

func (x *UnSubscribeFromUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UnSubscribeFromUserRequest) GetToUserId() string {
	if x != nil {
		return x.ToUserId
	}
	return ""
}

type UnSubscribeFromUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnSubscribeFromUserResponse) Reset() {
	*x = UnSubscribeFromUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_notifications_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnSubscribeFromUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnSubscribeFromUserResponse) ProtoMessage() {}

func (x *UnSubscribeFromUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_notifications_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnSubscribeFromUserResponse.ProtoReflect.Descriptor instead.
func (*UnSubscribeFromUserResponse) Descriptor() ([]byte, []int) {
	return file_notifications_notifications_proto_rawDescGZIP(), []int{10}
}

var File_notifications_notifications_proto protoreflect.FileDescriptor

var file_notifications_notifications_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x01, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x72, 0x65, 0x61, 0x64, 0x12, 0x48, 0x0a, 0x12, 0x6c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x10, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x32, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x41, 0x0a, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x61, 0x0a, 0x1d, 0x4d, 0x61, 0x72, 0x6b, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27,
	0x0a, 0x0f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x20, 0x0a, 0x1e, 0x4d, 0x61, 0x72, 0x6b, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x52, 0x65, 0x61,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x0a, 0x1b, 0x52, 0x65, 0x61,
	0x64, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x1e, 0x0a, 0x1c, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x4f, 0x0a, 0x16, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x19, 0x0a, 0x17, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53, 0x0a,
	0x1a, 0x55, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x46, 0x72, 0x6f, 0x6d,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x1d, 0x0a, 0x1b, 0x55, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x46, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0xac, 0x04, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x60, 0x0a, 0x0f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6c, 0x0a, 0x13, 0x55, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x55, 0x6e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x55, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x75, 0x0a, 0x16, 0x4d, 0x61, 0x72, 0x6b,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x52, 0x65,
	0x61, 0x64, 0x12, 0x2c, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2d, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x6f, 0x0a, 0x14, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2a, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x26, 0x5a, 0x24, 0x76, 0x65, 0x72, 0x63, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3b, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notifications_notifications_proto_rawDescOnce sync.Once
	file_notifications_notifications_proto_rawDescData = file_notifications_notifications_proto_rawDesc
)

func file_notifications_notifications_proto_rawDescGZIP() []byte {
	file_notifications_notifications_proto_rawDescOnce.Do(func() {
		file_notifications_notifications_proto_rawDescData = protoimpl.X.CompressGZIP(file_notifications_notifications_proto_rawDescData)
	})
	return file_notifications_notifications_proto_rawDescData
}

var file_notifications_notifications_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_notifications_notifications_proto_goTypes = []interface{}{
	(*Notification)(nil),                   // 0: notifications.Notification
	(*GetNotificationsRequest)(nil),        // 1: notifications.GetNotificationsRequest
	(*GetNotificationsResponse)(nil),       // 2: notifications.GetNotificationsResponse
	(*MarkNotificationAsReadRequest)(nil),  // 3: notifications.MarkNotificationAsReadRequest
	(*MarkNotificationAsReadResponse)(nil), // 4: notifications.MarkNotificationAsReadResponse
	(*ReadAllNotificationsRequest)(nil),    // 5: notifications.ReadAllNotificationsRequest
	(*ReadAllNotificationsResponse)(nil),   // 6: notifications.ReadAllNotificationsResponse
	(*SubscribeToUserRequest)(nil),         // 7: notifications.SubscribeToUserRequest
	(*SubscribeToUserResponse)(nil),        // 8: notifications.SubscribeToUserResponse
	(*UnSubscribeFromUserRequest)(nil),     // 9: notifications.UnSubscribeFromUserRequest
	(*UnSubscribeFromUserResponse)(nil),    // 10: notifications.UnSubscribeFromUserResponse
	(*timestamppb.Timestamp)(nil),          // 11: google.protobuf.Timestamp
}
var file_notifications_notifications_proto_depIdxs = []int32{
	11, // 0: notifications.Notification.license_expired_at:type_name -> google.protobuf.Timestamp
	0,  // 1: notifications.GetNotificationsResponse.notifications:type_name -> notifications.Notification
	7,  // 2: notifications.Notifications.SubscribeToUser:input_type -> notifications.SubscribeToUserRequest
	9,  // 3: notifications.Notifications.UnSubscribeFromUser:input_type -> notifications.UnSubscribeFromUserRequest
	1,  // 4: notifications.Notifications.GetNotifications:input_type -> notifications.GetNotificationsRequest
	3,  // 5: notifications.Notifications.MarkNotificationAsRead:input_type -> notifications.MarkNotificationAsReadRequest
	5,  // 6: notifications.Notifications.ReadAllNotifications:input_type -> notifications.ReadAllNotificationsRequest
	8,  // 7: notifications.Notifications.SubscribeToUser:output_type -> notifications.SubscribeToUserResponse
	10, // 8: notifications.Notifications.UnSubscribeFromUser:output_type -> notifications.UnSubscribeFromUserResponse
	2,  // 9: notifications.Notifications.GetNotifications:output_type -> notifications.GetNotificationsResponse
	4,  // 10: notifications.Notifications.MarkNotificationAsRead:output_type -> notifications.MarkNotificationAsReadResponse
	6,  // 11: notifications.Notifications.ReadAllNotifications:output_type -> notifications.ReadAllNotificationsResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_notifications_notifications_proto_init() }
func file_notifications_notifications_proto_init() {
	if File_notifications_notifications_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notifications_notifications_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
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
		file_notifications_notifications_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotificationsRequest); i {
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
		file_notifications_notifications_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotificationsResponse); i {
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
		file_notifications_notifications_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkNotificationAsReadRequest); i {
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
		file_notifications_notifications_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkNotificationAsReadResponse); i {
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
		file_notifications_notifications_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAllNotificationsRequest); i {
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
		file_notifications_notifications_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAllNotificationsResponse); i {
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
		file_notifications_notifications_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeToUserRequest); i {
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
		file_notifications_notifications_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeToUserResponse); i {
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
		file_notifications_notifications_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnSubscribeFromUserRequest); i {
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
		file_notifications_notifications_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnSubscribeFromUserResponse); i {
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
			RawDescriptor: file_notifications_notifications_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notifications_notifications_proto_goTypes,
		DependencyIndexes: file_notifications_notifications_proto_depIdxs,
		MessageInfos:      file_notifications_notifications_proto_msgTypes,
	}.Build()
	File_notifications_notifications_proto = out.File
	file_notifications_notifications_proto_rawDesc = nil
	file_notifications_notifications_proto_goTypes = nil
	file_notifications_notifications_proto_depIdxs = nil
}
