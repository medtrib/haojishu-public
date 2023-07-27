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
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Common_GetCaptcha_FullMethodName             = "/api.common.v1.Common/GetCaptcha"
	Common_VerifyCaptcha_FullMethodName          = "/api.common.v1.Common/VerifyCaptcha"
	Common_FireWallVerify_FullMethodName         = "/api.common.v1.Common/FireWallVerify"
	Common_FireWallBlockIP_FullMethodName        = "/api.common.v1.Common/FireWallBlockIP"
	Common_FireWallUnblockIP_FullMethodName      = "/api.common.v1.Common/FireWallUnblockIP"
	Common_GetBlockedIPs_FullMethodName          = "/api.common.v1.Common/GetBlockedIPs"
	Common_FireWallAddToWhitelist_FullMethodName = "/api.common.v1.Common/FireWallAddToWhitelist"
	Common_RemoveFromWhitelist_FullMethodName    = "/api.common.v1.Common/RemoveFromWhitelist"
	Common_GetWhitelistedIPs_FullMethodName      = "/api.common.v1.Common/GetWhitelistedIPs"
	Common_UpToken_FullMethodName                = "/api.common.v1.Common/UpToken"
	Common_UploadFileBase_FullMethodName         = "/api.common.v1.Common/UploadFileBase"
)

// CommonClient is the client API for Common service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommonClient interface {
	// 获取验证码
	GetCaptcha(ctx context.Context, in *GetCaptchaReq, opts ...grpc.CallOption) (*GetCaptchaRep, error)
	// 验证验证码
	VerifyCaptcha(ctx context.Context, in *VerifyCaptchaReq, opts ...grpc.CallOption) (*VerifyCaptchaRep, error)
	// 防火墙检查
	FireWallVerify(ctx context.Context, in *FireWallVerifyReq, opts ...grpc.CallOption) (*FireWallVerifyRep, error)
	// 防火墙封禁IP
	FireWallBlockIP(ctx context.Context, in *FireWallVerifyReq, opts ...grpc.CallOption) (*FireWallVerifyRep, error)
	// 防火墙解禁IP
	FireWallUnblockIP(ctx context.Context, in *FireWallVerifyReq, opts ...grpc.CallOption) (*FireWallVerifyRep, error)
	// 获取黑单列表
	GetBlockedIPs(ctx context.Context, in *FireWallListReq, opts ...grpc.CallOption) (*FireWallListRep, error)
	// 防火墙添加白名单IP
	FireWallAddToWhitelist(ctx context.Context, in *FireWallVerifyReq, opts ...grpc.CallOption) (*FireWallVerifyRep, error)
	// 防火墙从白名单中删除IP
	RemoveFromWhitelist(ctx context.Context, in *FireWallVerifyReq, opts ...grpc.CallOption) (*FireWallVerifyRep, error)
	// 获取白名单列表
	GetWhitelistedIPs(ctx context.Context, in *FireWallListReq, opts ...grpc.CallOption) (*FireWallListRep, error)
	// 获取Token
	UpToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UpTokenRep, error)
	// 上传文件
	UploadFileBase(ctx context.Context, in *UploadFileBaseReq, opts ...grpc.CallOption) (*UploadFileBaseRep, error)
}

type commonClient struct {
	cc grpc.ClientConnInterface
}

func NewCommonClient(cc grpc.ClientConnInterface) CommonClient {
	return &commonClient{cc}
}

func (c *commonClient) GetCaptcha(ctx context.Context, in *GetCaptchaReq, opts ...grpc.CallOption) (*GetCaptchaRep, error) {
	out := new(GetCaptchaRep)
	err := c.cc.Invoke(ctx, Common_GetCaptcha_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) VerifyCaptcha(ctx context.Context, in *VerifyCaptchaReq, opts ...grpc.CallOption) (*VerifyCaptchaRep, error) {
	out := new(VerifyCaptchaRep)
	err := c.cc.Invoke(ctx, Common_VerifyCaptcha_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) FireWallVerify(ctx context.Context, in *FireWallVerifyReq, opts ...grpc.CallOption) (*FireWallVerifyRep, error) {
	out := new(FireWallVerifyRep)
	err := c.cc.Invoke(ctx, Common_FireWallVerify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) FireWallBlockIP(ctx context.Context, in *FireWallVerifyReq, opts ...grpc.CallOption) (*FireWallVerifyRep, error) {
	out := new(FireWallVerifyRep)
	err := c.cc.Invoke(ctx, Common_FireWallBlockIP_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) FireWallUnblockIP(ctx context.Context, in *FireWallVerifyReq, opts ...grpc.CallOption) (*FireWallVerifyRep, error) {
	out := new(FireWallVerifyRep)
	err := c.cc.Invoke(ctx, Common_FireWallUnblockIP_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) GetBlockedIPs(ctx context.Context, in *FireWallListReq, opts ...grpc.CallOption) (*FireWallListRep, error) {
	out := new(FireWallListRep)
	err := c.cc.Invoke(ctx, Common_GetBlockedIPs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) FireWallAddToWhitelist(ctx context.Context, in *FireWallVerifyReq, opts ...grpc.CallOption) (*FireWallVerifyRep, error) {
	out := new(FireWallVerifyRep)
	err := c.cc.Invoke(ctx, Common_FireWallAddToWhitelist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) RemoveFromWhitelist(ctx context.Context, in *FireWallVerifyReq, opts ...grpc.CallOption) (*FireWallVerifyRep, error) {
	out := new(FireWallVerifyRep)
	err := c.cc.Invoke(ctx, Common_RemoveFromWhitelist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) GetWhitelistedIPs(ctx context.Context, in *FireWallListReq, opts ...grpc.CallOption) (*FireWallListRep, error) {
	out := new(FireWallListRep)
	err := c.cc.Invoke(ctx, Common_GetWhitelistedIPs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) UpToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UpTokenRep, error) {
	out := new(UpTokenRep)
	err := c.cc.Invoke(ctx, Common_UpToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) UploadFileBase(ctx context.Context, in *UploadFileBaseReq, opts ...grpc.CallOption) (*UploadFileBaseRep, error) {
	out := new(UploadFileBaseRep)
	err := c.cc.Invoke(ctx, Common_UploadFileBase_FullMethodName, in, out, opts...)
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
	GetCaptcha(context.Context, *GetCaptchaReq) (*GetCaptchaRep, error)
	// 验证验证码
	VerifyCaptcha(context.Context, *VerifyCaptchaReq) (*VerifyCaptchaRep, error)
	// 防火墙检查
	FireWallVerify(context.Context, *FireWallVerifyReq) (*FireWallVerifyRep, error)
	// 防火墙封禁IP
	FireWallBlockIP(context.Context, *FireWallVerifyReq) (*FireWallVerifyRep, error)
	// 防火墙解禁IP
	FireWallUnblockIP(context.Context, *FireWallVerifyReq) (*FireWallVerifyRep, error)
	// 获取黑单列表
	GetBlockedIPs(context.Context, *FireWallListReq) (*FireWallListRep, error)
	// 防火墙添加白名单IP
	FireWallAddToWhitelist(context.Context, *FireWallVerifyReq) (*FireWallVerifyRep, error)
	// 防火墙从白名单中删除IP
	RemoveFromWhitelist(context.Context, *FireWallVerifyReq) (*FireWallVerifyRep, error)
	// 获取白名单列表
	GetWhitelistedIPs(context.Context, *FireWallListReq) (*FireWallListRep, error)
	// 获取Token
	UpToken(context.Context, *emptypb.Empty) (*UpTokenRep, error)
	// 上传文件
	UploadFileBase(context.Context, *UploadFileBaseReq) (*UploadFileBaseRep, error)
	mustEmbedUnimplementedCommonServer()
}

// UnimplementedCommonServer must be embedded to have forward compatible implementations.
type UnimplementedCommonServer struct {
}

func (UnimplementedCommonServer) GetCaptcha(context.Context, *GetCaptchaReq) (*GetCaptchaRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCaptcha not implemented")
}
func (UnimplementedCommonServer) VerifyCaptcha(context.Context, *VerifyCaptchaReq) (*VerifyCaptchaRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyCaptcha not implemented")
}
func (UnimplementedCommonServer) FireWallVerify(context.Context, *FireWallVerifyReq) (*FireWallVerifyRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FireWallVerify not implemented")
}
func (UnimplementedCommonServer) FireWallBlockIP(context.Context, *FireWallVerifyReq) (*FireWallVerifyRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FireWallBlockIP not implemented")
}
func (UnimplementedCommonServer) FireWallUnblockIP(context.Context, *FireWallVerifyReq) (*FireWallVerifyRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FireWallUnblockIP not implemented")
}
func (UnimplementedCommonServer) GetBlockedIPs(context.Context, *FireWallListReq) (*FireWallListRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockedIPs not implemented")
}
func (UnimplementedCommonServer) FireWallAddToWhitelist(context.Context, *FireWallVerifyReq) (*FireWallVerifyRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FireWallAddToWhitelist not implemented")
}
func (UnimplementedCommonServer) RemoveFromWhitelist(context.Context, *FireWallVerifyReq) (*FireWallVerifyRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromWhitelist not implemented")
}
func (UnimplementedCommonServer) GetWhitelistedIPs(context.Context, *FireWallListReq) (*FireWallListRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWhitelistedIPs not implemented")
}
func (UnimplementedCommonServer) UpToken(context.Context, *emptypb.Empty) (*UpTokenRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpToken not implemented")
}
func (UnimplementedCommonServer) UploadFileBase(context.Context, *UploadFileBaseReq) (*UploadFileBaseRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFileBase not implemented")
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
	in := new(GetCaptchaReq)
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
		return srv.(CommonServer).GetCaptcha(ctx, req.(*GetCaptchaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_VerifyCaptcha_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyCaptchaReq)
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
		return srv.(CommonServer).VerifyCaptcha(ctx, req.(*VerifyCaptchaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_FireWallVerify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FireWallVerifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).FireWallVerify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Common_FireWallVerify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).FireWallVerify(ctx, req.(*FireWallVerifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_FireWallBlockIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FireWallVerifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).FireWallBlockIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Common_FireWallBlockIP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).FireWallBlockIP(ctx, req.(*FireWallVerifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_FireWallUnblockIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FireWallVerifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).FireWallUnblockIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Common_FireWallUnblockIP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).FireWallUnblockIP(ctx, req.(*FireWallVerifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_GetBlockedIPs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FireWallListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).GetBlockedIPs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Common_GetBlockedIPs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).GetBlockedIPs(ctx, req.(*FireWallListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_FireWallAddToWhitelist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FireWallVerifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).FireWallAddToWhitelist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Common_FireWallAddToWhitelist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).FireWallAddToWhitelist(ctx, req.(*FireWallVerifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_RemoveFromWhitelist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FireWallVerifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).RemoveFromWhitelist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Common_RemoveFromWhitelist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).RemoveFromWhitelist(ctx, req.(*FireWallVerifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_GetWhitelistedIPs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FireWallListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).GetWhitelistedIPs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Common_GetWhitelistedIPs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).GetWhitelistedIPs(ctx, req.(*FireWallListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_UpToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).UpToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Common_UpToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).UpToken(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_UploadFileBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFileBaseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).UploadFileBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Common_UploadFileBase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).UploadFileBase(ctx, req.(*UploadFileBaseReq))
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
		{
			MethodName: "FireWallVerify",
			Handler:    _Common_FireWallVerify_Handler,
		},
		{
			MethodName: "FireWallBlockIP",
			Handler:    _Common_FireWallBlockIP_Handler,
		},
		{
			MethodName: "FireWallUnblockIP",
			Handler:    _Common_FireWallUnblockIP_Handler,
		},
		{
			MethodName: "GetBlockedIPs",
			Handler:    _Common_GetBlockedIPs_Handler,
		},
		{
			MethodName: "FireWallAddToWhitelist",
			Handler:    _Common_FireWallAddToWhitelist_Handler,
		},
		{
			MethodName: "RemoveFromWhitelist",
			Handler:    _Common_RemoveFromWhitelist_Handler,
		},
		{
			MethodName: "GetWhitelistedIPs",
			Handler:    _Common_GetWhitelistedIPs_Handler,
		},
		{
			MethodName: "UpToken",
			Handler:    _Common_UpToken_Handler,
		},
		{
			MethodName: "UploadFileBase",
			Handler:    _Common_UploadFileBase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/v1/common.proto",
}
