// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: kb_index_service.proto

package proto

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

// KBIndexServiceClient is the client API for KBIndexService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KBIndexServiceClient interface {
	Create(ctx context.Context, in *KBIndexRequest, opts ...grpc.CallOption) (*KBIndexReply, error)
	Modify(ctx context.Context, in *KBIndexRequest, opts ...grpc.CallOption) (*KBIndexReply, error)
	Delete(ctx context.Context, in *KBIndexRequest, opts ...grpc.CallOption) (*KBIndexReply, error)
}

type kBIndexServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKBIndexServiceClient(cc grpc.ClientConnInterface) KBIndexServiceClient {
	return &kBIndexServiceClient{cc}
}

func (c *kBIndexServiceClient) Create(ctx context.Context, in *KBIndexRequest, opts ...grpc.CallOption) (*KBIndexReply, error) {
	out := new(KBIndexReply)
	err := c.cc.Invoke(ctx, "/KBIndexService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kBIndexServiceClient) Modify(ctx context.Context, in *KBIndexRequest, opts ...grpc.CallOption) (*KBIndexReply, error) {
	out := new(KBIndexReply)
	err := c.cc.Invoke(ctx, "/KBIndexService/Modify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kBIndexServiceClient) Delete(ctx context.Context, in *KBIndexRequest, opts ...grpc.CallOption) (*KBIndexReply, error) {
	out := new(KBIndexReply)
	err := c.cc.Invoke(ctx, "/KBIndexService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KBIndexServiceServer is the server API for KBIndexService service.
// All implementations must embed UnimplementedKBIndexServiceServer
// for forward compatibility
type KBIndexServiceServer interface {
	Create(context.Context, *KBIndexRequest) (*KBIndexReply, error)
	Modify(context.Context, *KBIndexRequest) (*KBIndexReply, error)
	Delete(context.Context, *KBIndexRequest) (*KBIndexReply, error)
	mustEmbedUnimplementedKBIndexServiceServer()
}

// UnimplementedKBIndexServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKBIndexServiceServer struct {
}

func (UnimplementedKBIndexServiceServer) Create(context.Context, *KBIndexRequest) (*KBIndexReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedKBIndexServiceServer) Modify(context.Context, *KBIndexRequest) (*KBIndexReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Modify not implemented")
}
func (UnimplementedKBIndexServiceServer) Delete(context.Context, *KBIndexRequest) (*KBIndexReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedKBIndexServiceServer) mustEmbedUnimplementedKBIndexServiceServer() {}

// UnsafeKBIndexServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KBIndexServiceServer will
// result in compilation errors.
type UnsafeKBIndexServiceServer interface {
	mustEmbedUnimplementedKBIndexServiceServer()
}

func RegisterKBIndexServiceServer(s grpc.ServiceRegistrar, srv KBIndexServiceServer) {
	s.RegisterService(&KBIndexService_ServiceDesc, srv)
}

func _KBIndexService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KBIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KBIndexServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KBIndexService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KBIndexServiceServer).Create(ctx, req.(*KBIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KBIndexService_Modify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KBIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KBIndexServiceServer).Modify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KBIndexService/Modify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KBIndexServiceServer).Modify(ctx, req.(*KBIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KBIndexService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KBIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KBIndexServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KBIndexService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KBIndexServiceServer).Delete(ctx, req.(*KBIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KBIndexService_ServiceDesc is the grpc.ServiceDesc for KBIndexService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KBIndexService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "KBIndexService",
	HandlerType: (*KBIndexServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _KBIndexService_Create_Handler,
		},
		{
			MethodName: "Modify",
			Handler:    _KBIndexService_Modify_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _KBIndexService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kb_index_service.proto",
}