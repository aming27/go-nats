// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.4
// source: email.proto

//protoc --go_out=plugins=grpc:. *.proto

package emailService

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Email struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailID   string                 `protobuf:"bytes,1,opt,name=EmailID,proto3" json:"EmailID,omitempty"`
	From      string                 `protobuf:"bytes,2,opt,name=From,proto3" json:"From,omitempty"`
	To        string                 `protobuf:"bytes,3,opt,name=To,proto3" json:"To,omitempty"`
	Subject   string                 `protobuf:"bytes,4,opt,name=Subject,proto3" json:"Subject,omitempty"`
	Message   string                 `protobuf:"bytes,5,opt,name=Message,proto3" json:"Message,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *Email) Reset() {
	*x = Email{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Email) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Email) ProtoMessage() {}

func (x *Email) ProtoReflect() protoreflect.Message {
	mi := &file_email_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Email.ProtoReflect.Descriptor instead.
func (*Email) Descriptor() ([]byte, []int) {
	return file_email_proto_rawDescGZIP(), []int{0}
}

func (x *Email) GetEmailID() string {
	if x != nil {
		return x.EmailID
	}
	return ""
}

func (x *Email) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Email) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Email) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Email) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Email) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_email_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_email_proto_rawDescGZIP(), []int{1}
}

type CreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From    string `protobuf:"bytes,1,opt,name=From,proto3" json:"From,omitempty"`
	To      string `protobuf:"bytes,2,opt,name=To,proto3" json:"To,omitempty"`
	Subject string `protobuf:"bytes,3,opt,name=Subject,proto3" json:"Subject,omitempty"`
	Message string `protobuf:"bytes,4,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *CreateReq) Reset() {
	*x = CreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReq) ProtoMessage() {}

func (x *CreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_email_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReq.ProtoReflect.Descriptor instead.
func (*CreateReq) Descriptor() ([]byte, []int) {
	return file_email_proto_rawDescGZIP(), []int{2}
}

func (x *CreateReq) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *CreateReq) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *CreateReq) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *CreateReq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateRes) Reset() {
	*x = CreateRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRes) ProtoMessage() {}

func (x *CreateRes) ProtoReflect() protoreflect.Message {
	mi := &file_email_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRes.ProtoReflect.Descriptor instead.
func (*CreateRes) Descriptor() ([]byte, []int) {
	return file_email_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRes) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetByIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailID string `protobuf:"bytes,1,opt,name=EmailID,proto3" json:"EmailID,omitempty"`
}

func (x *GetByIDReq) Reset() {
	*x = GetByIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIDReq) ProtoMessage() {}

func (x *GetByIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_email_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIDReq.ProtoReflect.Descriptor instead.
func (*GetByIDReq) Descriptor() ([]byte, []int) {
	return file_email_proto_rawDescGZIP(), []int{4}
}

func (x *GetByIDReq) GetEmailID() string {
	if x != nil {
		return x.EmailID
	}
	return ""
}

type GetByIDRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email *Email `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
}

func (x *GetByIDRes) Reset() {
	*x = GetByIDRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIDRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIDRes) ProtoMessage() {}

func (x *GetByIDRes) ProtoReflect() protoreflect.Message {
	mi := &file_email_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIDRes.ProtoReflect.Descriptor instead.
func (*GetByIDRes) Descriptor() ([]byte, []int) {
	return file_email_proto_rawDescGZIP(), []int{5}
}

func (x *GetByIDRes) GetEmail() *Email {
	if x != nil {
		return x.Email
	}
	return nil
}

type SearchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search string `protobuf:"bytes,1,opt,name=Search,proto3" json:"Search,omitempty"`
	Page   int64  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size   int64  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *SearchReq) Reset() {
	*x = SearchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchReq) ProtoMessage() {}

func (x *SearchReq) ProtoReflect() protoreflect.Message {
	mi := &file_email_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchReq.ProtoReflect.Descriptor instead.
func (*SearchReq) Descriptor() ([]byte, []int) {
	return file_email_proto_rawDescGZIP(), []int{6}
}

func (x *SearchReq) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *SearchReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchReq) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type SearchRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64    `protobuf:"varint,1,opt,name=TotalCount,proto3" json:"TotalCount,omitempty"`
	TotalPages int64    `protobuf:"varint,2,opt,name=TotalPages,proto3" json:"TotalPages,omitempty"`
	Page       int64    `protobuf:"varint,3,opt,name=Page,proto3" json:"Page,omitempty"`
	Size       int64    `protobuf:"varint,4,opt,name=Size,proto3" json:"Size,omitempty"`
	HasMore    bool     `protobuf:"varint,5,opt,name=HasMore,proto3" json:"HasMore,omitempty"`
	Emails     []*Email `protobuf:"bytes,6,rep,name=Emails,proto3" json:"Emails,omitempty"`
}

func (x *SearchRes) Reset() {
	*x = SearchRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRes) ProtoMessage() {}

func (x *SearchRes) ProtoReflect() protoreflect.Message {
	mi := &file_email_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRes.ProtoReflect.Descriptor instead.
func (*SearchRes) Descriptor() ([]byte, []int) {
	return file_email_proto_rawDescGZIP(), []int{7}
}

func (x *SearchRes) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *SearchRes) GetTotalPages() int64 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *SearchRes) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchRes) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *SearchRes) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

func (x *SearchRes) GetEmails() []*Email {
	if x != nil {
		return x.Emails
	}
	return nil
}

var File_email_proto protoreflect.FileDescriptor

var file_email_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x46, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x54, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x63, 0x0a, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x54, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x54, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x23, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x26, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x44, 0x22, 0x37, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x05, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x4b, 0x0a, 0x09, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x22, 0xba, 0x01, 0x0a, 0x09, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x48, 0x61, 0x73,
	0x4d, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x48, 0x61, 0x73, 0x4d,
	0x6f, 0x72, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x06, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73,
	0x32, 0xcb, 0x01, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x3f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x18, 0x2e, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x12, 0x3c, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x17, 0x2e, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x10,
	0x5a, 0x0e, 0x2e, 0x3b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_email_proto_rawDescOnce sync.Once
	file_email_proto_rawDescData = file_email_proto_rawDesc
)

func file_email_proto_rawDescGZIP() []byte {
	file_email_proto_rawDescOnce.Do(func() {
		file_email_proto_rawDescData = protoimpl.X.CompressGZIP(file_email_proto_rawDescData)
	})
	return file_email_proto_rawDescData
}

var file_email_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_email_proto_goTypes = []interface{}{
	(*Email)(nil),                 // 0: emailService.Email
	(*Empty)(nil),                 // 1: emailService.Empty
	(*CreateReq)(nil),             // 2: emailService.CreateReq
	(*CreateRes)(nil),             // 3: emailService.CreateRes
	(*GetByIDReq)(nil),            // 4: emailService.GetByIDReq
	(*GetByIDRes)(nil),            // 5: emailService.GetByIDRes
	(*SearchReq)(nil),             // 6: emailService.SearchReq
	(*SearchRes)(nil),             // 7: emailService.SearchRes
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_email_proto_depIdxs = []int32{
	8, // 0: emailService.Email.CreatedAt:type_name -> google.protobuf.Timestamp
	0, // 1: emailService.GetByIDRes.Email:type_name -> emailService.Email
	0, // 2: emailService.SearchRes.Emails:type_name -> emailService.Email
	2, // 3: emailService.EmailService.Create:input_type -> emailService.CreateReq
	4, // 4: emailService.EmailService.GetByID:input_type -> emailService.GetByIDReq
	6, // 5: emailService.EmailService.Search:input_type -> emailService.SearchReq
	3, // 6: emailService.EmailService.Create:output_type -> emailService.CreateRes
	5, // 7: emailService.EmailService.GetByID:output_type -> emailService.GetByIDRes
	7, // 8: emailService.EmailService.Search:output_type -> emailService.SearchRes
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_email_proto_init() }
func file_email_proto_init() {
	if File_email_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_email_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Email); i {
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
		file_email_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_email_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReq); i {
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
		file_email_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRes); i {
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
		file_email_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIDReq); i {
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
		file_email_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIDRes); i {
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
		file_email_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchReq); i {
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
		file_email_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRes); i {
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
			RawDescriptor: file_email_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_email_proto_goTypes,
		DependencyIndexes: file_email_proto_depIdxs,
		MessageInfos:      file_email_proto_msgTypes,
	}.Build()
	File_email_proto = out.File
	file_email_proto_rawDesc = nil
	file_email_proto_goTypes = nil
	file_email_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EmailServiceClient is the client API for EmailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmailServiceClient interface {
	Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateRes, error)
	GetByID(ctx context.Context, in *GetByIDReq, opts ...grpc.CallOption) (*GetByIDRes, error)
	Search(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*SearchRes, error)
}

type emailServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailServiceClient(cc grpc.ClientConnInterface) EmailServiceClient {
	return &emailServiceClient{cc}
}

func (c *emailServiceClient) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateRes, error) {
	out := new(CreateRes)
	err := c.cc.Invoke(ctx, "/emailService.EmailService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) GetByID(ctx context.Context, in *GetByIDReq, opts ...grpc.CallOption) (*GetByIDRes, error) {
	out := new(GetByIDRes)
	err := c.cc.Invoke(ctx, "/emailService.EmailService/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) Search(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*SearchRes, error) {
	out := new(SearchRes)
	err := c.cc.Invoke(ctx, "/emailService.EmailService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServiceServer is the server API for EmailService service.
type EmailServiceServer interface {
	Create(context.Context, *CreateReq) (*CreateRes, error)
	GetByID(context.Context, *GetByIDReq) (*GetByIDRes, error)
	Search(context.Context, *SearchReq) (*SearchRes, error)
}

// UnimplementedEmailServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEmailServiceServer struct {
}

func (*UnimplementedEmailServiceServer) Create(context.Context, *CreateReq) (*CreateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedEmailServiceServer) GetByID(context.Context, *GetByIDReq) (*GetByIDRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (*UnimplementedEmailServiceServer) Search(context.Context, *SearchReq) (*SearchRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}

func RegisterEmailServiceServer(s *grpc.Server, srv EmailServiceServer) {
	s.RegisterService(&_EmailService_serviceDesc, srv)
}

func _EmailService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emailService.EmailService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).Create(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emailService.EmailService/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).GetByID(ctx, req.(*GetByIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emailService.EmailService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).Search(ctx, req.(*SearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _EmailService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "emailService.EmailService",
	HandlerType: (*EmailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _EmailService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _EmailService_GetByID_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _EmailService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "email.proto",
}