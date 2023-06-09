// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: page.proto

package page

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

type GetPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageId uint32 `protobuf:"varint,1,opt,name=PageId,proto3" json:"PageId,omitempty"`
}

func (x *GetPageRequest) Reset() {
	*x = GetPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPageRequest) ProtoMessage() {}

func (x *GetPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_page_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPageRequest.ProtoReflect.Descriptor instead.
func (*GetPageRequest) Descriptor() ([]byte, []int) {
	return file_page_proto_rawDescGZIP(), []int{0}
}

func (x *GetPageRequest) GetPageId() uint32 {
	if x != nil {
		return x.PageId
	}
	return 0
}

type GetPageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageId     uint32   `protobuf:"varint,1,opt,name=PageId,proto3" json:"PageId,omitempty"`
	NextPageId uint32   `protobuf:"varint,2,opt,name=NextPageId,proto3" json:"NextPageId,omitempty"`
	PrevPageId uint32   `protobuf:"varint,3,opt,name=PrevPageId,proto3" json:"PrevPageId,omitempty"`
	ArticleIds []uint32 `protobuf:"varint,4,rep,packed,name=ArticleIds,proto3" json:"ArticleIds,omitempty"`
}

func (x *GetPageResponse) Reset() {
	*x = GetPageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPageResponse) ProtoMessage() {}

func (x *GetPageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_page_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPageResponse.ProtoReflect.Descriptor instead.
func (*GetPageResponse) Descriptor() ([]byte, []int) {
	return file_page_proto_rawDescGZIP(), []int{1}
}

func (x *GetPageResponse) GetPageId() uint32 {
	if x != nil {
		return x.PageId
	}
	return 0
}

func (x *GetPageResponse) GetNextPageId() uint32 {
	if x != nil {
		return x.NextPageId
	}
	return 0
}

func (x *GetPageResponse) GetPrevPageId() uint32 {
	if x != nil {
		return x.PrevPageId
	}
	return 0
}

func (x *GetPageResponse) GetArticleIds() []uint32 {
	if x != nil {
		return x.ArticleIds
	}
	return nil
}

type SetPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageId     uint32   `protobuf:"varint,1,opt,name=PageId,proto3" json:"PageId,omitempty"`
	NextPageId uint32   `protobuf:"varint,2,opt,name=NextPageId,proto3" json:"NextPageId,omitempty"`
	PrevPageId uint32   `protobuf:"varint,3,opt,name=PrevPageId,proto3" json:"PrevPageId,omitempty"`
	ArticleIds []uint32 `protobuf:"varint,4,rep,packed,name=ArticleIds,proto3" json:"ArticleIds,omitempty"`
}

func (x *SetPageRequest) Reset() {
	*x = SetPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPageRequest) ProtoMessage() {}

func (x *SetPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_page_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPageRequest.ProtoReflect.Descriptor instead.
func (*SetPageRequest) Descriptor() ([]byte, []int) {
	return file_page_proto_rawDescGZIP(), []int{2}
}

func (x *SetPageRequest) GetPageId() uint32 {
	if x != nil {
		return x.PageId
	}
	return 0
}

func (x *SetPageRequest) GetNextPageId() uint32 {
	if x != nil {
		return x.NextPageId
	}
	return 0
}

func (x *SetPageRequest) GetPrevPageId() uint32 {
	if x != nil {
		return x.PrevPageId
	}
	return 0
}

func (x *SetPageRequest) GetArticleIds() []uint32 {
	if x != nil {
		return x.ArticleIds
	}
	return nil
}

type SetPageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool `protobuf:"varint,1,opt,name=IsSuccess,proto3" json:"IsSuccess,omitempty"`
}

func (x *SetPageResponse) Reset() {
	*x = SetPageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPageResponse) ProtoMessage() {}

func (x *SetPageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_page_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPageResponse.ProtoReflect.Descriptor instead.
func (*SetPageResponse) Descriptor() ([]byte, []int) {
	return file_page_proto_rawDescGZIP(), []int{3}
}

func (x *SetPageResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

type InsertPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageId     uint32   `protobuf:"varint,1,opt,name=PageId,proto3" json:"PageId,omitempty"`
	NextPageId uint32   `protobuf:"varint,2,opt,name=NextPageId,proto3" json:"NextPageId,omitempty"`
	PrevPageId uint32   `protobuf:"varint,3,opt,name=PrevPageId,proto3" json:"PrevPageId,omitempty"`
	ArticleIds []uint32 `protobuf:"varint,4,rep,packed,name=ArticleIds,proto3" json:"ArticleIds,omitempty"`
}

func (x *InsertPageRequest) Reset() {
	*x = InsertPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertPageRequest) ProtoMessage() {}

func (x *InsertPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_page_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertPageRequest.ProtoReflect.Descriptor instead.
func (*InsertPageRequest) Descriptor() ([]byte, []int) {
	return file_page_proto_rawDescGZIP(), []int{4}
}

func (x *InsertPageRequest) GetPageId() uint32 {
	if x != nil {
		return x.PageId
	}
	return 0
}

func (x *InsertPageRequest) GetNextPageId() uint32 {
	if x != nil {
		return x.NextPageId
	}
	return 0
}

func (x *InsertPageRequest) GetPrevPageId() uint32 {
	if x != nil {
		return x.PrevPageId
	}
	return 0
}

func (x *InsertPageRequest) GetArticleIds() []uint32 {
	if x != nil {
		return x.ArticleIds
	}
	return nil
}

type InsertPageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool   `protobuf:"varint,1,opt,name=IsSuccess,proto3" json:"IsSuccess,omitempty"`
	NewPageId uint32 `protobuf:"varint,2,opt,name=NewPageId,proto3" json:"NewPageId,omitempty"`
}

func (x *InsertPageResponse) Reset() {
	*x = InsertPageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertPageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertPageResponse) ProtoMessage() {}

func (x *InsertPageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_page_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertPageResponse.ProtoReflect.Descriptor instead.
func (*InsertPageResponse) Descriptor() ([]byte, []int) {
	return file_page_proto_rawDescGZIP(), []int{5}
}

func (x *InsertPageResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *InsertPageResponse) GetNewPageId() uint32 {
	if x != nil {
		return x.NewPageId
	}
	return 0
}

type DeletePageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageId uint32 `protobuf:"varint,1,opt,name=PageId,proto3" json:"PageId,omitempty"`
	Hour   int32  `protobuf:"varint,2,opt,name=Hour,proto3" json:"Hour,omitempty"`
}

func (x *DeletePageRequest) Reset() {
	*x = DeletePageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePageRequest) ProtoMessage() {}

func (x *DeletePageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_page_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePageRequest.ProtoReflect.Descriptor instead.
func (*DeletePageRequest) Descriptor() ([]byte, []int) {
	return file_page_proto_rawDescGZIP(), []int{6}
}

func (x *DeletePageRequest) GetPageId() uint32 {
	if x != nil {
		return x.PageId
	}
	return 0
}

func (x *DeletePageRequest) GetHour() int32 {
	if x != nil {
		return x.Hour
	}
	return 0
}

type DeletePageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RowAffected uint32 `protobuf:"varint,1,opt,name=RowAffected,proto3" json:"RowAffected,omitempty"`
}

func (x *DeletePageResponse) Reset() {
	*x = DeletePageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePageResponse) ProtoMessage() {}

func (x *DeletePageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_page_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePageResponse.ProtoReflect.Descriptor instead.
func (*DeletePageResponse) Descriptor() ([]byte, []int) {
	return file_page_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePageResponse) GetRowAffected() uint32 {
	if x != nil {
		return x.RowAffected
	}
	return 0
}

var File_page_proto protoreflect.FileDescriptor

var file_page_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x61,
	0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x50, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x72, 0x65, 0x76, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x50, 0x72, 0x65, 0x76, 0x50, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49,
	0x64, 0x73, 0x22, 0x88, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x50, 0x72, 0x65, 0x76, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x50, 0x72, 0x65, 0x76, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x0a, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x22, 0x2f, 0x0a,
	0x0f, 0x53, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x49, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x8b,
	0x01, 0x0a, 0x11, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x4e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x50, 0x72, 0x65, 0x76, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x50, 0x72, 0x65, 0x76, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x0a, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x22, 0x50, 0x0a, 0x12,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x49, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x4e, 0x65, 0x77, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x3f,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x50, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x48,
	0x6f, 0x75, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x48, 0x6f, 0x75, 0x72, 0x22,
	0x36, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x6f, 0x77, 0x41, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x52, 0x6f, 0x77, 0x41,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x32, 0x98, 0x02, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65,
	0x12, 0x2c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x0f, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c,
	0x0a, 0x07, 0x53, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x0f, 0x2e, 0x53, 0x65, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x53, 0x65, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x2e, 0x49, 0x6e, 0x73,
	0x65, 0x72, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x67,
	0x65, 0x43, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x48, 0x6f, 0x75, 0x72, 0x42, 0x65, 0x66, 0x6f,
	0x72, 0x65, 0x12, 0x12, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x67,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_page_proto_rawDescOnce sync.Once
	file_page_proto_rawDescData = file_page_proto_rawDesc
)

func file_page_proto_rawDescGZIP() []byte {
	file_page_proto_rawDescOnce.Do(func() {
		file_page_proto_rawDescData = protoimpl.X.CompressGZIP(file_page_proto_rawDescData)
	})
	return file_page_proto_rawDescData
}

var file_page_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_page_proto_goTypes = []interface{}{
	(*GetPageRequest)(nil),     // 0: GetPageRequest
	(*GetPageResponse)(nil),    // 1: GetPageResponse
	(*SetPageRequest)(nil),     // 2: SetPageRequest
	(*SetPageResponse)(nil),    // 3: SetPageResponse
	(*InsertPageRequest)(nil),  // 4: InsertPageRequest
	(*InsertPageResponse)(nil), // 5: InsertPageResponse
	(*DeletePageRequest)(nil),  // 6: DeletePageRequest
	(*DeletePageResponse)(nil), // 7: DeletePageResponse
}
var file_page_proto_depIdxs = []int32{
	0, // 0: Page.GetPage:input_type -> GetPageRequest
	2, // 1: Page.SetPage:input_type -> SetPageRequest
	4, // 2: Page.InsertPage:input_type -> InsertPageRequest
	6, // 3: Page.DeletePageCertainHourBefore:input_type -> DeletePageRequest
	6, // 4: Page.DeletePage:input_type -> DeletePageRequest
	1, // 5: Page.GetPage:output_type -> GetPageResponse
	3, // 6: Page.SetPage:output_type -> SetPageResponse
	5, // 7: Page.InsertPage:output_type -> InsertPageResponse
	7, // 8: Page.DeletePageCertainHourBefore:output_type -> DeletePageResponse
	7, // 9: Page.DeletePage:output_type -> DeletePageResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_page_proto_init() }
func file_page_proto_init() {
	if File_page_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_page_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPageRequest); i {
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
		file_page_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPageResponse); i {
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
		file_page_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPageRequest); i {
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
		file_page_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPageResponse); i {
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
		file_page_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertPageRequest); i {
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
		file_page_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertPageResponse); i {
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
		file_page_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePageRequest); i {
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
		file_page_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePageResponse); i {
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
			RawDescriptor: file_page_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_page_proto_goTypes,
		DependencyIndexes: file_page_proto_depIdxs,
		MessageInfos:      file_page_proto_msgTypes,
	}.Build()
	File_page_proto = out.File
	file_page_proto_rawDesc = nil
	file_page_proto_goTypes = nil
	file_page_proto_depIdxs = nil
}
