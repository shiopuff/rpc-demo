// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.0
// source: bookstore.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BookStoreClient is the client API for BookStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookStoreClient interface {
	// 查询所有书架
	ListShelves(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListShelvesResponse, error)
	// 创建新书架
	CreateShelf(ctx context.Context, in *CreateShelfRequest, opts ...grpc.CallOption) (*Shelf, error)
	// 返回书店内指定的书架
	GetShelf(ctx context.Context, in *GetShelfRequest, opts ...grpc.CallOption) (*Shelf, error)
	// 删除指定书架
	DeleteShelf(ctx context.Context, in *DeleteShelfRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type bookStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewBookStoreClient(cc grpc.ClientConnInterface) BookStoreClient {
	return &bookStoreClient{cc}
}

func (c *bookStoreClient) ListShelves(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListShelvesResponse, error) {
	out := new(ListShelvesResponse)
	err := c.cc.Invoke(ctx, "/bookstore.BookStore/ListShelves", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookStoreClient) CreateShelf(ctx context.Context, in *CreateShelfRequest, opts ...grpc.CallOption) (*Shelf, error) {
	out := new(Shelf)
	err := c.cc.Invoke(ctx, "/bookstore.BookStore/CreateShelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookStoreClient) GetShelf(ctx context.Context, in *GetShelfRequest, opts ...grpc.CallOption) (*Shelf, error) {
	out := new(Shelf)
	err := c.cc.Invoke(ctx, "/bookstore.BookStore/GetShelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookStoreClient) DeleteShelf(ctx context.Context, in *DeleteShelfRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/bookstore.BookStore/DeleteShelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookStoreServer is the server API for BookStore service.
// All implementations must embed UnimplementedBookStoreServer
// for forward compatibility
type BookStoreServer interface {
	// 查询所有书架
	ListShelves(context.Context, *emptypb.Empty) (*ListShelvesResponse, error)
	// 创建新书架
	CreateShelf(context.Context, *CreateShelfRequest) (*Shelf, error)
	// 返回书店内指定的书架
	GetShelf(context.Context, *GetShelfRequest) (*Shelf, error)
	// 删除指定书架
	DeleteShelf(context.Context, *DeleteShelfRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedBookStoreServer()
}

// UnimplementedBookStoreServer must be embedded to have forward compatible implementations.
type UnimplementedBookStoreServer struct {
}

func (UnimplementedBookStoreServer) ListShelves(context.Context, *emptypb.Empty) (*ListShelvesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListShelves not implemented")
}
func (UnimplementedBookStoreServer) CreateShelf(context.Context, *CreateShelfRequest) (*Shelf, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShelf not implemented")
}
func (UnimplementedBookStoreServer) GetShelf(context.Context, *GetShelfRequest) (*Shelf, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShelf not implemented")
}
func (UnimplementedBookStoreServer) DeleteShelf(context.Context, *DeleteShelfRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteShelf not implemented")
}
func (UnimplementedBookStoreServer) mustEmbedUnimplementedBookStoreServer() {}

// UnsafeBookStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookStoreServer will
// result in compilation errors.
type UnsafeBookStoreServer interface {
	mustEmbedUnimplementedBookStoreServer()
}

func RegisterBookStoreServer(s grpc.ServiceRegistrar, srv BookStoreServer) {
	s.RegisterService(&BookStore_ServiceDesc, srv)
}

func _BookStore_ListShelves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookStoreServer).ListShelves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.BookStore/ListShelves",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookStoreServer).ListShelves(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookStore_CreateShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookStoreServer).CreateShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.BookStore/CreateShelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookStoreServer).CreateShelf(ctx, req.(*CreateShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookStore_GetShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookStoreServer).GetShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.BookStore/GetShelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookStoreServer).GetShelf(ctx, req.(*GetShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookStore_DeleteShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookStoreServer).DeleteShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.BookStore/DeleteShelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookStoreServer).DeleteShelf(ctx, req.(*DeleteShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookStore_ServiceDesc is the grpc.ServiceDesc for BookStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bookstore.BookStore",
	HandlerType: (*BookStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListShelves",
			Handler:    _BookStore_ListShelves_Handler,
		},
		{
			MethodName: "CreateShelf",
			Handler:    _BookStore_CreateShelf_Handler,
		},
		{
			MethodName: "GetShelf",
			Handler:    _BookStore_GetShelf_Handler,
		},
		{
			MethodName: "DeleteShelf",
			Handler:    _BookStore_DeleteShelf_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bookstore.proto",
}
