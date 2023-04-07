// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: list.proto

package list

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

// ListClient is the client API for List service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ListClient interface {
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	PatchList(ctx context.Context, in *PatchListRequest, opts ...grpc.CallOption) (*SetHeadResponse, error)
	InsertList(ctx context.Context, in *InsertListRequest, opts ...grpc.CallOption) (*SetHeadResponse, error)
}

type listClient struct {
	cc grpc.ClientConnInterface
}

func NewListClient(cc grpc.ClientConnInterface) ListClient {
	return &listClient{cc}
}

func (c *listClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/List/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listClient) PatchList(ctx context.Context, in *PatchListRequest, opts ...grpc.CallOption) (*SetHeadResponse, error) {
	out := new(SetHeadResponse)
	err := c.cc.Invoke(ctx, "/List/PatchList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listClient) InsertList(ctx context.Context, in *InsertListRequest, opts ...grpc.CallOption) (*SetHeadResponse, error) {
	out := new(SetHeadResponse)
	err := c.cc.Invoke(ctx, "/List/InsertList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListServer is the server API for List service.
// All implementations must embed UnimplementedListServer
// for forward compatibility
type ListServer interface {
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	PatchList(context.Context, *PatchListRequest) (*SetHeadResponse, error)
	InsertList(context.Context, *InsertListRequest) (*SetHeadResponse, error)
	mustEmbedUnimplementedListServer()
}

// UnimplementedListServer must be embedded to have forward compatible implementations.
type UnimplementedListServer struct {
}

func (UnimplementedListServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedListServer) PatchList(context.Context, *PatchListRequest) (*SetHeadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchList not implemented")
}
func (UnimplementedListServer) InsertList(context.Context, *InsertListRequest) (*SetHeadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertList not implemented")
}
func (UnimplementedListServer) mustEmbedUnimplementedListServer() {}

// UnsafeListServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ListServer will
// result in compilation errors.
type UnsafeListServer interface {
	mustEmbedUnimplementedListServer()
}

func RegisterListServer(s grpc.ServiceRegistrar, srv ListServer) {
	s.RegisterService(&List_ServiceDesc, srv)
}

func _List_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/List/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _List_PatchList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListServer).PatchList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/List/PatchList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListServer).PatchList(ctx, req.(*PatchListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _List_InsertList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListServer).InsertList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/List/InsertList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListServer).InsertList(ctx, req.(*InsertListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// List_ServiceDesc is the grpc.ServiceDesc for List service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var List_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "List",
	HandlerType: (*ListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _List_GetList_Handler,
		},
		{
			MethodName: "PatchList",
			Handler:    _List_PatchList_Handler,
		},
		{
			MethodName: "InsertList",
			Handler:    _List_InsertList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "list.proto",
}