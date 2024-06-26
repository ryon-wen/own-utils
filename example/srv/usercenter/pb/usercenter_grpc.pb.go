// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: pb/usercenter.proto

package pb

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
	UserCenterService_Create_FullMethodName = "/pb.UserCenterService/Create"
	UserCenterService_Update_FullMethodName = "/pb.UserCenterService/Update"
	UserCenterService_Get_FullMethodName    = "/pb.UserCenterService/Get"
	UserCenterService_Delete_FullMethodName = "/pb.UserCenterService/Delete"
)

// UserCenterServiceClient is the client API for UserCenterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserCenterServiceClient interface {
	Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*Empty, error)
	Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*Empty, error)
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*Empty, error)
}

type userCenterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserCenterServiceClient(cc grpc.ClientConnInterface) UserCenterServiceClient {
	return &userCenterServiceClient{cc}
}

func (c *userCenterServiceClient) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, UserCenterService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterServiceClient) Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, UserCenterService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterServiceClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error) {
	out := new(GetResp)
	err := c.cc.Invoke(ctx, UserCenterService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterServiceClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, UserCenterService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserCenterServiceServer is the server API for UserCenterService service.
// All implementations must embed UnimplementedUserCenterServiceServer
// for forward compatibility
type UserCenterServiceServer interface {
	Create(context.Context, *CreateReq) (*Empty, error)
	Update(context.Context, *UpdateReq) (*Empty, error)
	Get(context.Context, *GetReq) (*GetResp, error)
	Delete(context.Context, *DeleteReq) (*Empty, error)
	mustEmbedUnimplementedUserCenterServiceServer()
}

// UnimplementedUserCenterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserCenterServiceServer struct {
}

func (UnimplementedUserCenterServiceServer) Create(context.Context, *CreateReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserCenterServiceServer) Update(context.Context, *UpdateReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserCenterServiceServer) Get(context.Context, *GetReq) (*GetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedUserCenterServiceServer) Delete(context.Context, *DeleteReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedUserCenterServiceServer) mustEmbedUnimplementedUserCenterServiceServer() {}

// UnsafeUserCenterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserCenterServiceServer will
// result in compilation errors.
type UnsafeUserCenterServiceServer interface {
	mustEmbedUnimplementedUserCenterServiceServer()
}

func RegisterUserCenterServiceServer(s grpc.ServiceRegistrar, srv UserCenterServiceServer) {
	s.RegisterService(&UserCenterService_ServiceDesc, srv)
}

func _UserCenterService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCenterService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServiceServer).Create(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenterService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCenterService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServiceServer).Update(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenterService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCenterService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServiceServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenterService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCenterService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServiceServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserCenterService_ServiceDesc is the grpc.ServiceDesc for UserCenterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserCenterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserCenterService",
	HandlerType: (*UserCenterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserCenterService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserCenterService_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserCenterService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserCenterService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/usercenter.proto",
}
