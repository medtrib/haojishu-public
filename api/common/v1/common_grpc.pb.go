// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: common/v1/common.proto

package v1

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
	Common_GetCaptcha_FullMethodName    = "/api.common.v1.Common/GetCaptcha"
	Common_VerifyCaptcha_FullMethodName = "/api.common.v1.Common/VerifyCaptcha"
)

// CommonClient is the client API for Common service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommonClient interface {
	// 获取验证码
	GetCaptcha(ctx context.Context, in *GetCaptchaRequest, opts ...grpc.CallOption) (*GetCaptchaReply, error)
	// 验证验证码
	VerifyCaptcha(ctx context.Context, in *VerifyCaptchaRequest, opts ...grpc.CallOption) (*VerifyCaptchaReply, error)
}

type commonClient struct {
	cc grpc.ClientConnInterface
}

func NewCommonClient(cc grpc.ClientConnInterface) CommonClient {
	return &commonClient{cc}
}

func (c *commonClient) GetCaptcha(ctx context.Context, in *GetCaptchaRequest, opts ...grpc.CallOption) (*GetCaptchaReply, error) {
	out := new(GetCaptchaReply)
	err := c.cc.Invoke(ctx, Common_GetCaptcha_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) VerifyCaptcha(ctx context.Context, in *VerifyCaptchaRequest, opts ...grpc.CallOption) (*VerifyCaptchaReply, error) {
	out := new(VerifyCaptchaReply)
	err := c.cc.Invoke(ctx, Common_VerifyCaptcha_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommonServer is the server API for Common service.
// All implementations must embed UnimplementedCommonServer
// for forward compatibility
type CommonServer interface {
	// 获取验证码
	GetCaptcha(context.Context, *GetCaptchaRequest) (*GetCaptchaReply, error)
	// 验证验证码
	VerifyCaptcha(context.Context, *VerifyCaptchaRequest) (*VerifyCaptchaReply, error)
	mustEmbedUnimplementedCommonServer()
}

// UnimplementedCommonServer must be embedded to have forward compatible implementations.
type UnimplementedCommonServer struct {
}

func (UnimplementedCommonServer) GetCaptcha(context.Context, *GetCaptchaRequest) (*GetCaptchaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCaptcha not implemented")
}
func (UnimplementedCommonServer) VerifyCaptcha(context.Context, *VerifyCaptchaRequest) (*VerifyCaptchaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyCaptcha not implemented")
}
func (UnimplementedCommonServer) mustEmbedUnimplementedCommonServer() {}

// UnsafeCommonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommonServer will
// result in compilation errors.
type UnsafeCommonServer interface {
	mustEmbedUnimplementedCommonServer()
}

func RegisterCommonServer(s grpc.ServiceRegistrar, srv CommonServer) {
	s.RegisterService(&Common_ServiceDesc, srv)
}

func _Common_GetCaptcha_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCaptchaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).GetCaptcha(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Common_GetCaptcha_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).GetCaptcha(ctx, req.(*GetCaptchaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_VerifyCaptcha_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyCaptchaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).VerifyCaptcha(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Common_VerifyCaptcha_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).VerifyCaptcha(ctx, req.(*VerifyCaptchaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Common_ServiceDesc is the grpc.ServiceDesc for Common service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Common_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.common.v1.Common",
	HandlerType: (*CommonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCaptcha",
			Handler:    _Common_GetCaptcha_Handler,
		},
		{
			MethodName: "VerifyCaptcha",
			Handler:    _Common_VerifyCaptcha_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/v1/common.proto",
}
