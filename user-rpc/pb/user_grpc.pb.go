// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: user.proto

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
	Usercenter_UserRegister_FullMethodName = "/pb.usercenter/UserRegister"
	Usercenter_UserLogin_FullMethodName    = "/pb.usercenter/UserLogin"
	Usercenter_UserGetInfo_FullMethodName  = "/pb.usercenter/UserGetInfo"
)

// UsercenterClient is the client API for Usercenter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsercenterClient interface {
	UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error)
	UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error)
	UserGetInfo(ctx context.Context, in *UserGetInfoReq, opts ...grpc.CallOption) (*UserGetInfoResp, error)
}

type usercenterClient struct {
	cc grpc.ClientConnInterface
}

func NewUsercenterClient(cc grpc.ClientConnInterface) UsercenterClient {
	return &usercenterClient{cc}
}

func (c *usercenterClient) UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error) {
	out := new(UserRegisterResp)
	err := c.cc.Invoke(ctx, Usercenter_UserRegister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error) {
	out := new(UserLoginResp)
	err := c.cc.Invoke(ctx, Usercenter_UserLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) UserGetInfo(ctx context.Context, in *UserGetInfoReq, opts ...grpc.CallOption) (*UserGetInfoResp, error) {
	out := new(UserGetInfoResp)
	err := c.cc.Invoke(ctx, Usercenter_UserGetInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsercenterServer is the server API for Usercenter service.
// All implementations must embed UnimplementedUsercenterServer
// for forward compatibility
type UsercenterServer interface {
	UserRegister(context.Context, *UserRegisterReq) (*UserRegisterResp, error)
	UserLogin(context.Context, *UserLoginReq) (*UserLoginResp, error)
	UserGetInfo(context.Context, *UserGetInfoReq) (*UserGetInfoResp, error)
	mustEmbedUnimplementedUsercenterServer()
}

// UnimplementedUsercenterServer must be embedded to have forward compatible implementations.
type UnimplementedUsercenterServer struct {
}

func (UnimplementedUsercenterServer) UserRegister(context.Context, *UserRegisterReq) (*UserRegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (UnimplementedUsercenterServer) UserLogin(context.Context, *UserLoginReq) (*UserLoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedUsercenterServer) UserGetInfo(context.Context, *UserGetInfoReq) (*UserGetInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserGetInfo not implemented")
}
func (UnimplementedUsercenterServer) mustEmbedUnimplementedUsercenterServer() {}

// UnsafeUsercenterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsercenterServer will
// result in compilation errors.
type UnsafeUsercenterServer interface {
	mustEmbedUnimplementedUsercenterServer()
}

func RegisterUsercenterServer(s grpc.ServiceRegistrar, srv UsercenterServer) {
	s.RegisterService(&Usercenter_ServiceDesc, srv)
}

func _Usercenter_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_UserRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).UserRegister(ctx, req.(*UserRegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_UserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).UserLogin(ctx, req.(*UserLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_UserGetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGetInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).UserGetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_UserGetInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).UserGetInfo(ctx, req.(*UserGetInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Usercenter_ServiceDesc is the grpc.ServiceDesc for Usercenter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Usercenter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.usercenter",
	HandlerType: (*UsercenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserRegister",
			Handler:    _Usercenter_UserRegister_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _Usercenter_UserLogin_Handler,
		},
		{
			MethodName: "UserGetInfo",
			Handler:    _Usercenter_UserGetInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
