// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/worker/proto/worker.proto

package proto

import (
	duration "github.com/golang/protobuf/ptypes/duration"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// Ответ таймаутов вычислителей в текстовом виде
type TimeoutsMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Add      string `protobuf:"bytes,1,opt,name=add,proto3" json:"add,omitempty"`
	Subtract string `protobuf:"bytes,2,opt,name=subtract,proto3" json:"subtract,omitempty"`
	Multiply string `protobuf:"bytes,3,opt,name=multiply,proto3" json:"multiply,omitempty"`
	Divide   string `protobuf:"bytes,4,opt,name=divide,proto3" json:"divide,omitempty"`
}

func (x *TimeoutsMessage) Reset() {
	*x = TimeoutsMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_worker_proto_worker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeoutsMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeoutsMessage) ProtoMessage() {}

func (x *TimeoutsMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_worker_proto_worker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeoutsMessage.ProtoReflect.Descriptor instead.
func (*TimeoutsMessage) Descriptor() ([]byte, []int) {
	return file_pkg_worker_proto_worker_proto_rawDescGZIP(), []int{0}
}

func (x *TimeoutsMessage) GetAdd() string {
	if x != nil {
		return x.Add
	}
	return ""
}

func (x *TimeoutsMessage) GetSubtract() string {
	if x != nil {
		return x.Subtract
	}
	return ""
}

func (x *TimeoutsMessage) GetMultiply() string {
	if x != nil {
		return x.Multiply
	}
	return ""
}

func (x *TimeoutsMessage) GetDivide() string {
	if x != nil {
		return x.Divide
	}
	return ""
}

// Ожидаемое собщение таймаутов
type TimeoutsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Add      *duration.Duration `protobuf:"bytes,1,opt,name=add,proto3" json:"add,omitempty"`
	Subtract *duration.Duration `protobuf:"bytes,2,opt,name=subtract,proto3" json:"subtract,omitempty"`
	Multiply *duration.Duration `protobuf:"bytes,3,opt,name=multiply,proto3" json:"multiply,omitempty"`
	Divide   *duration.Duration `protobuf:"bytes,4,opt,name=divide,proto3" json:"divide,omitempty"`
}

func (x *TimeoutsRequest) Reset() {
	*x = TimeoutsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_worker_proto_worker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeoutsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeoutsRequest) ProtoMessage() {}

func (x *TimeoutsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_worker_proto_worker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeoutsRequest.ProtoReflect.Descriptor instead.
func (*TimeoutsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_worker_proto_worker_proto_rawDescGZIP(), []int{1}
}

func (x *TimeoutsRequest) GetAdd() *duration.Duration {
	if x != nil {
		return x.Add
	}
	return nil
}

func (x *TimeoutsRequest) GetSubtract() *duration.Duration {
	if x != nil {
		return x.Subtract
	}
	return nil
}

func (x *TimeoutsRequest) GetMultiply() *duration.Duration {
	if x != nil {
		return x.Multiply
	}
	return nil
}

func (x *TimeoutsRequest) GetDivide() *duration.Duration {
	if x != nil {
		return x.Divide
	}
	return nil
}

// Определение токена
type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	IsOp  bool   `protobuf:"varint,2,opt,name=isOp,proto3" json:"isOp,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_worker_proto_worker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_worker_proto_worker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_pkg_worker_proto_worker_proto_rawDescGZIP(), []int{2}
}

func (x *Token) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Token) GetIsOp() bool {
	if x != nil {
		return x.IsOp
	}
	return false
}

// Получение выражения для вычисления
type TaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int32    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Expression []*Token `protobuf:"bytes,2,rep,name=expression,proto3" json:"expression,omitempty"`
}

func (x *TaskRequest) Reset() {
	*x = TaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_worker_proto_worker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRequest) ProtoMessage() {}

func (x *TaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_worker_proto_worker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRequest.ProtoReflect.Descriptor instead.
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return file_pkg_worker_proto_worker_proto_rawDescGZIP(), []int{3}
}

func (x *TaskRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *TaskRequest) GetExpression() []*Token {
	if x != nil {
		return x.Expression
	}
	return nil
}

// Ответ на задачу
type TaskRespons struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Value     float32 `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
	FlagError bool    `protobuf:"varint,3,opt,name=flag_error,json=flagError,proto3" json:"flag_error,omitempty"`
}

func (x *TaskRespons) Reset() {
	*x = TaskRespons{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_worker_proto_worker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRespons) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRespons) ProtoMessage() {}

func (x *TaskRespons) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_worker_proto_worker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRespons.ProtoReflect.Descriptor instead.
func (*TaskRespons) Descriptor() ([]byte, []int) {
	return file_pkg_worker_proto_worker_proto_rawDescGZIP(), []int{4}
}

func (x *TaskRespons) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *TaskRespons) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *TaskRespons) GetFlagError() bool {
	if x != nil {
		return x.FlagError
	}
	return false
}

var File_pkg_worker_proto_worker_proto protoreflect.FileDescriptor

var file_pkg_worker_proto_worker_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x6b, 0x67, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x0f, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x73,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x64, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x64, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x22, 0xdf, 0x01, 0x0a, 0x0f, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a,
	0x03, 0x61, 0x64, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x61, 0x64, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x73, 0x75,
	0x62, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x73, 0x75, 0x62, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x12, 0x35, 0x0a, 0x08, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x12, 0x31, 0x0a, 0x06, 0x64, 0x69, 0x76, 0x69,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x06, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x22, 0x31, 0x0a, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73,
	0x4f, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x70, 0x22, 0x55,
	0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x5b, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x66, 0x6c, 0x61, 0x67, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x32, 0x81, 0x02, 0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x73, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x73, 0x12, 0x17, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x36, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x13, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x13, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x42, 0x20, 0x5a, 0x1e, 0x47, 0x6f, 0x43, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_worker_proto_worker_proto_rawDescOnce sync.Once
	file_pkg_worker_proto_worker_proto_rawDescData = file_pkg_worker_proto_worker_proto_rawDesc
)

func file_pkg_worker_proto_worker_proto_rawDescGZIP() []byte {
	file_pkg_worker_proto_worker_proto_rawDescOnce.Do(func() {
		file_pkg_worker_proto_worker_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_worker_proto_worker_proto_rawDescData)
	})
	return file_pkg_worker_proto_worker_proto_rawDescData
}

var file_pkg_worker_proto_worker_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_worker_proto_worker_proto_goTypes = []interface{}{
	(*TimeoutsMessage)(nil),   // 0: worker.TimeoutsMessage
	(*TimeoutsRequest)(nil),   // 1: worker.TimeoutsRequest
	(*Token)(nil),             // 2: worker.Token
	(*TaskRequest)(nil),       // 3: worker.TaskRequest
	(*TaskRespons)(nil),       // 4: worker.TaskRespons
	(*duration.Duration)(nil), // 5: google.protobuf.Duration
	(*empty.Empty)(nil),       // 6: google.protobuf.Empty
}
var file_pkg_worker_proto_worker_proto_depIdxs = []int32{
	5, // 0: worker.TimeoutsRequest.add:type_name -> google.protobuf.Duration
	5, // 1: worker.TimeoutsRequest.subtract:type_name -> google.protobuf.Duration
	5, // 2: worker.TimeoutsRequest.multiply:type_name -> google.protobuf.Duration
	5, // 3: worker.TimeoutsRequest.divide:type_name -> google.protobuf.Duration
	2, // 4: worker.TaskRequest.expression:type_name -> worker.Token
	6, // 5: worker.WorkerService.GetTimeouts:input_type -> google.protobuf.Empty
	1, // 6: worker.WorkerService.SetTimeouts:input_type -> worker.TimeoutsRequest
	3, // 7: worker.WorkerService.SetTask:input_type -> worker.TaskRequest
	6, // 8: worker.WorkerService.GetResult:input_type -> google.protobuf.Empty
	0, // 9: worker.WorkerService.GetTimeouts:output_type -> worker.TimeoutsMessage
	6, // 10: worker.WorkerService.SetTimeouts:output_type -> google.protobuf.Empty
	6, // 11: worker.WorkerService.SetTask:output_type -> google.protobuf.Empty
	4, // 12: worker.WorkerService.GetResult:output_type -> worker.TaskRespons
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_worker_proto_worker_proto_init() }
func file_pkg_worker_proto_worker_proto_init() {
	if File_pkg_worker_proto_worker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_worker_proto_worker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeoutsMessage); i {
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
		file_pkg_worker_proto_worker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeoutsRequest); i {
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
		file_pkg_worker_proto_worker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_pkg_worker_proto_worker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskRequest); i {
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
		file_pkg_worker_proto_worker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskRespons); i {
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
			RawDescriptor: file_pkg_worker_proto_worker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_worker_proto_worker_proto_goTypes,
		DependencyIndexes: file_pkg_worker_proto_worker_proto_depIdxs,
		MessageInfos:      file_pkg_worker_proto_worker_proto_msgTypes,
	}.Build()
	File_pkg_worker_proto_worker_proto = out.File
	file_pkg_worker_proto_worker_proto_rawDesc = nil
	file_pkg_worker_proto_worker_proto_goTypes = nil
	file_pkg_worker_proto_worker_proto_depIdxs = nil
}
