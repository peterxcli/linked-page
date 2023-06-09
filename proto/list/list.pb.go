// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: list.proto

package list

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

type GetListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListId uint32 `protobuf:"varint,1,opt,name=ListId,proto3" json:"ListId,omitempty"`
	UserId uint32 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *GetListRequest) Reset() {
	*x = GetListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListRequest) ProtoMessage() {}

func (x *GetListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListRequest.ProtoReflect.Descriptor instead.
func (*GetListRequest) Descriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{0}
}

func (x *GetListRequest) GetListId() uint32 {
	if x != nil {
		return x.ListId
	}
	return 0
}

func (x *GetListRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListId uint32 `protobuf:"varint,1,opt,name=ListId,proto3" json:"ListId,omitempty"`
	UserId uint32 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	HeadId uint32 `protobuf:"varint,3,opt,name=HeadId,proto3" json:"HeadId,omitempty"`
}

func (x *GetListResponse) Reset() {
	*x = GetListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListResponse) ProtoMessage() {}

func (x *GetListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListResponse.ProtoReflect.Descriptor instead.
func (*GetListResponse) Descriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{1}
}

func (x *GetListResponse) GetListId() uint32 {
	if x != nil {
		return x.ListId
	}
	return 0
}

func (x *GetListResponse) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetListResponse) GetHeadId() uint32 {
	if x != nil {
		return x.HeadId
	}
	return 0
}

type PatchListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListId uint32 `protobuf:"varint,1,opt,name=ListId,proto3" json:"ListId,omitempty"`
	UserId uint32 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	HeadId uint32 `protobuf:"varint,3,opt,name=HeadId,proto3" json:"HeadId,omitempty"`
}

func (x *PatchListRequest) Reset() {
	*x = PatchListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchListRequest) ProtoMessage() {}

func (x *PatchListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchListRequest.ProtoReflect.Descriptor instead.
func (*PatchListRequest) Descriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{2}
}

func (x *PatchListRequest) GetListId() uint32 {
	if x != nil {
		return x.ListId
	}
	return 0
}

func (x *PatchListRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PatchListRequest) GetHeadId() uint32 {
	if x != nil {
		return x.HeadId
	}
	return 0
}

type InsertListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListId uint32 `protobuf:"varint,1,opt,name=ListId,proto3" json:"ListId,omitempty"`
	UserId uint32 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	HeadId uint32 `protobuf:"varint,3,opt,name=HeadId,proto3" json:"HeadId,omitempty"`
}

func (x *InsertListRequest) Reset() {
	*x = InsertListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertListRequest) ProtoMessage() {}

func (x *InsertListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_list_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertListRequest.ProtoReflect.Descriptor instead.
func (*InsertListRequest) Descriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{3}
}

func (x *InsertListRequest) GetListId() uint32 {
	if x != nil {
		return x.ListId
	}
	return 0
}

func (x *InsertListRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *InsertListRequest) GetHeadId() uint32 {
	if x != nil {
		return x.HeadId
	}
	return 0
}

type SetHeadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool `protobuf:"varint,1,opt,name=IsSuccess,proto3" json:"IsSuccess,omitempty"`
}

func (x *SetHeadResponse) Reset() {
	*x = SetHeadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_list_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetHeadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetHeadResponse) ProtoMessage() {}

func (x *SetHeadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_list_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetHeadResponse.ProtoReflect.Descriptor instead.
func (*SetHeadResponse) Descriptor() ([]byte, []int) {
	return file_list_proto_rawDescGZIP(), []int{4}
}

func (x *SetHeadResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

var File_list_proto protoreflect.FileDescriptor

var file_list_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x59,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x48, 0x65, 0x61, 0x64, 0x49, 0x64, 0x22, 0x5a, 0x0a, 0x10, 0x50, 0x61, 0x74,
	0x63, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x48, 0x65, 0x61, 0x64, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x48,
	0x65, 0x61, 0x64, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x11, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65,
	0x61, 0x64, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x48, 0x65, 0x61, 0x64,
	0x49, 0x64, 0x22, 0x2f, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x49, 0x73, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x32, 0x9a, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x50, 0x61,
	0x74, 0x63, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x53, 0x65, 0x74,
	0x48, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0a,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x49, 0x6e, 0x73,
	0x65, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x53, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x0e, 0x5a, 0x0c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_list_proto_rawDescOnce sync.Once
	file_list_proto_rawDescData = file_list_proto_rawDesc
)

func file_list_proto_rawDescGZIP() []byte {
	file_list_proto_rawDescOnce.Do(func() {
		file_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_list_proto_rawDescData)
	})
	return file_list_proto_rawDescData
}

var file_list_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_list_proto_goTypes = []interface{}{
	(*GetListRequest)(nil),    // 0: GetListRequest
	(*GetListResponse)(nil),   // 1: GetListResponse
	(*PatchListRequest)(nil),  // 2: PatchListRequest
	(*InsertListRequest)(nil), // 3: InsertListRequest
	(*SetHeadResponse)(nil),   // 4: SetHeadResponse
}
var file_list_proto_depIdxs = []int32{
	0, // 0: List.GetList:input_type -> GetListRequest
	2, // 1: List.PatchList:input_type -> PatchListRequest
	3, // 2: List.InsertList:input_type -> InsertListRequest
	1, // 3: List.GetList:output_type -> GetListResponse
	4, // 4: List.PatchList:output_type -> SetHeadResponse
	4, // 5: List.InsertList:output_type -> SetHeadResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_list_proto_init() }
func file_list_proto_init() {
	if File_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListRequest); i {
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
		file_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListResponse); i {
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
		file_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchListRequest); i {
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
		file_list_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertListRequest); i {
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
		file_list_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetHeadResponse); i {
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
			RawDescriptor: file_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_list_proto_goTypes,
		DependencyIndexes: file_list_proto_depIdxs,
		MessageInfos:      file_list_proto_msgTypes,
	}.Build()
	File_list_proto = out.File
	file_list_proto_rawDesc = nil
	file_list_proto_goTypes = nil
	file_list_proto_depIdxs = nil
}
