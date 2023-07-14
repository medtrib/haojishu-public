// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: auth/v1/auth.proto

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
	Auth_AddRole_FullMethodName            = "/api.auth.v1.Auth/AddRole"
	Auth_EditRole_FullMethodName           = "/api.auth.v1.Auth/EditRole"
	Auth_DelRole_FullMethodName            = "/api.auth.v1.Auth/DelRole"
	Auth_FullRoleList_FullMethodName       = "/api.auth.v1.Auth/FullRoleList"
	Auth_PageRoleList_FullMethodName       = "/api.auth.v1.Auth/PageRoleList"
	Auth_AddRolesForUser_FullMethodName    = "/api.auth.v1.Auth/AddRolesForUser"
	Auth_GetRolesForUser_FullMethodName    = "/api.auth.v1.Auth/GetRolesForUser"
	Auth_GetUsersForRole_FullMethodName    = "/api.auth.v1.Auth/GetUsersForRole"
	Auth_DeleteRoleForUser_FullMethodName  = "/api.auth.v1.Auth/DeleteRoleForUser"
	Auth_DeleteRolesForUser_FullMethodName = "/api.auth.v1.Auth/DeleteRolesForUser"
	Auth_GetRolePolicies_FullMethodName    = "/api.auth.v1.Auth/GetRolePolicies"
	Auth_SetRolePolicies_FullMethodName    = "/api.auth.v1.Auth/SetRolePolicies"
	Auth_CheckAuth_FullMethodName          = "/api.auth.v1.Auth/CheckAuth"
	Auth_CreateMenu_FullMethodName         = "/api.auth.v1.Auth/CreateMenu"
	Auth_EditMenu_FullMethodName           = "/api.auth.v1.Auth/EditMenu"
	Auth_DeleteMenu_FullMethodName         = "/api.auth.v1.Auth/DeleteMenu"
	Auth_ListMenu_FullMethodName           = "/api.auth.v1.Auth/ListMenu"
	Auth_ListMenuTree_FullMethodName       = "/api.auth.v1.Auth/ListMenuTree"
)

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	// 添加角色
	AddRole(ctx context.Context, in *AddRoleReq, opts ...grpc.CallOption) (*AddRoleRep, error)
	// 编辑角色
	EditRole(ctx context.Context, in *EditRoleReq, opts ...grpc.CallOption) (*RepStatus, error)
	// 删除角色
	DelRole(ctx context.Context, in *DelRoleReq, opts ...grpc.CallOption) (*RepStatus, error)
	// 获取角色列表(完整)
	FullRoleList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FullRoleListRep, error)
	// 获取角色列表(分页)
	PageRoleList(ctx context.Context, in *PageRoleListReq, opts ...grpc.CallOption) (*PageRoleListRep, error)
	// 给用户设置角色
	AddRolesForUser(ctx context.Context, in *SetUserForRoleReq, opts ...grpc.CallOption) (*RepStatus, error)
	// 获取用户角色
	GetRolesForUser(ctx context.Context, in *GetRolesForUserReq, opts ...grpc.CallOption) (*GetRolesForUserRep, error)
	// 获取角色有那些用户
	GetUsersForRole(ctx context.Context, in *GetUsersForRoleReq, opts ...grpc.CallOption) (*GetUsersForRoleRep, error)
	// 删除单个用户角色(如果需要删除单个用户的某个角色用这个)
	DeleteRoleForUser(ctx context.Context, in *DeleteRoleForUserReq, opts ...grpc.CallOption) (*RepStatus, error)
	// 删除多个用户角色(删除传递用户的所有角色)
	DeleteRolesForUser(ctx context.Context, in *DeleteRolesForUserReq, opts ...grpc.CallOption) (*RepStatus, error)
	// 获取角色有那些权限
	GetRolePolicies(ctx context.Context, in *GetRolePoliciesReq, opts ...grpc.CallOption) (*GetRolePoliciesRep, error)
	// 设置角色权限
	SetRolePolicies(ctx context.Context, in *SetRolePoliciesReq, opts ...grpc.CallOption) (*RepStatus, error)
	// 检查权限
	CheckAuth(ctx context.Context, in *CheckAuthReq, opts ...grpc.CallOption) (*RepStatus, error)
	// 创建菜单
	CreateMenu(ctx context.Context, in *CreateMenuReq, opts ...grpc.CallOption) (*Menu, error)
	// 更新菜单
	EditMenu(ctx context.Context, in *EditMenuReq, opts ...grpc.CallOption) (*RepStatus, error)
	// 删除菜单
	DeleteMenu(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*RepStatus, error)
	// 菜单列表(完整)
	ListMenu(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListMenuRep, error)
	// 菜单列表(树)
	ListMenuTree(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListMenuRep, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) AddRole(ctx context.Context, in *AddRoleReq, opts ...grpc.CallOption) (*AddRoleRep, error) {
	out := new(AddRoleRep)
	err := c.cc.Invoke(ctx, Auth_AddRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) EditRole(ctx context.Context, in *EditRoleReq, opts ...grpc.CallOption) (*RepStatus, error) {
	out := new(RepStatus)
	err := c.cc.Invoke(ctx, Auth_EditRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DelRole(ctx context.Context, in *DelRoleReq, opts ...grpc.CallOption) (*RepStatus, error) {
	out := new(RepStatus)
	err := c.cc.Invoke(ctx, Auth_DelRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) FullRoleList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FullRoleListRep, error) {
	out := new(FullRoleListRep)
	err := c.cc.Invoke(ctx, Auth_FullRoleList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) PageRoleList(ctx context.Context, in *PageRoleListReq, opts ...grpc.CallOption) (*PageRoleListRep, error) {
	out := new(PageRoleListRep)
	err := c.cc.Invoke(ctx, Auth_PageRoleList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) AddRolesForUser(ctx context.Context, in *SetUserForRoleReq, opts ...grpc.CallOption) (*RepStatus, error) {
	out := new(RepStatus)
	err := c.cc.Invoke(ctx, Auth_AddRolesForUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetRolesForUser(ctx context.Context, in *GetRolesForUserReq, opts ...grpc.CallOption) (*GetRolesForUserRep, error) {
	out := new(GetRolesForUserRep)
	err := c.cc.Invoke(ctx, Auth_GetRolesForUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetUsersForRole(ctx context.Context, in *GetUsersForRoleReq, opts ...grpc.CallOption) (*GetUsersForRoleRep, error) {
	out := new(GetUsersForRoleRep)
	err := c.cc.Invoke(ctx, Auth_GetUsersForRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DeleteRoleForUser(ctx context.Context, in *DeleteRoleForUserReq, opts ...grpc.CallOption) (*RepStatus, error) {
	out := new(RepStatus)
	err := c.cc.Invoke(ctx, Auth_DeleteRoleForUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DeleteRolesForUser(ctx context.Context, in *DeleteRolesForUserReq, opts ...grpc.CallOption) (*RepStatus, error) {
	out := new(RepStatus)
	err := c.cc.Invoke(ctx, Auth_DeleteRolesForUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetRolePolicies(ctx context.Context, in *GetRolePoliciesReq, opts ...grpc.CallOption) (*GetRolePoliciesRep, error) {
	out := new(GetRolePoliciesRep)
	err := c.cc.Invoke(ctx, Auth_GetRolePolicies_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SetRolePolicies(ctx context.Context, in *SetRolePoliciesReq, opts ...grpc.CallOption) (*RepStatus, error) {
	out := new(RepStatus)
	err := c.cc.Invoke(ctx, Auth_SetRolePolicies_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CheckAuth(ctx context.Context, in *CheckAuthReq, opts ...grpc.CallOption) (*RepStatus, error) {
	out := new(RepStatus)
	err := c.cc.Invoke(ctx, Auth_CheckAuth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CreateMenu(ctx context.Context, in *CreateMenuReq, opts ...grpc.CallOption) (*Menu, error) {
	out := new(Menu)
	err := c.cc.Invoke(ctx, Auth_CreateMenu_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) EditMenu(ctx context.Context, in *EditMenuReq, opts ...grpc.CallOption) (*RepStatus, error) {
	out := new(RepStatus)
	err := c.cc.Invoke(ctx, Auth_EditMenu_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DeleteMenu(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*RepStatus, error) {
	out := new(RepStatus)
	err := c.cc.Invoke(ctx, Auth_DeleteMenu_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ListMenu(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListMenuRep, error) {
	out := new(ListMenuRep)
	err := c.cc.Invoke(ctx, Auth_ListMenu_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ListMenuTree(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListMenuRep, error) {
	out := new(ListMenuRep)
	err := c.cc.Invoke(ctx, Auth_ListMenuTree_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	// 添加角色
	AddRole(context.Context, *AddRoleReq) (*AddRoleRep, error)
	// 编辑角色
	EditRole(context.Context, *EditRoleReq) (*RepStatus, error)
	// 删除角色
	DelRole(context.Context, *DelRoleReq) (*RepStatus, error)
	// 获取角色列表(完整)
	FullRoleList(context.Context, *emptypb.Empty) (*FullRoleListRep, error)
	// 获取角色列表(分页)
	PageRoleList(context.Context, *PageRoleListReq) (*PageRoleListRep, error)
	// 给用户设置角色
	AddRolesForUser(context.Context, *SetUserForRoleReq) (*RepStatus, error)
	// 获取用户角色
	GetRolesForUser(context.Context, *GetRolesForUserReq) (*GetRolesForUserRep, error)
	// 获取角色有那些用户
	GetUsersForRole(context.Context, *GetUsersForRoleReq) (*GetUsersForRoleRep, error)
	// 删除单个用户角色(如果需要删除单个用户的某个角色用这个)
	DeleteRoleForUser(context.Context, *DeleteRoleForUserReq) (*RepStatus, error)
	// 删除多个用户角色(删除传递用户的所有角色)
	DeleteRolesForUser(context.Context, *DeleteRolesForUserReq) (*RepStatus, error)
	// 获取角色有那些权限
	GetRolePolicies(context.Context, *GetRolePoliciesReq) (*GetRolePoliciesRep, error)
	// 设置角色权限
	SetRolePolicies(context.Context, *SetRolePoliciesReq) (*RepStatus, error)
	// 检查权限
	CheckAuth(context.Context, *CheckAuthReq) (*RepStatus, error)
	// 创建菜单
	CreateMenu(context.Context, *CreateMenuReq) (*Menu, error)
	// 更新菜单
	EditMenu(context.Context, *EditMenuReq) (*RepStatus, error)
	// 删除菜单
	DeleteMenu(context.Context, *IdReq) (*RepStatus, error)
	// 菜单列表(完整)
	ListMenu(context.Context, *emptypb.Empty) (*ListMenuRep, error)
	// 菜单列表(树)
	ListMenuTree(context.Context, *emptypb.Empty) (*ListMenuRep, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) AddRole(context.Context, *AddRoleReq) (*AddRoleRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRole not implemented")
}
func (UnimplementedAuthServer) EditRole(context.Context, *EditRoleReq) (*RepStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditRole not implemented")
}
func (UnimplementedAuthServer) DelRole(context.Context, *DelRoleReq) (*RepStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelRole not implemented")
}
func (UnimplementedAuthServer) FullRoleList(context.Context, *emptypb.Empty) (*FullRoleListRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FullRoleList not implemented")
}
func (UnimplementedAuthServer) PageRoleList(context.Context, *PageRoleListReq) (*PageRoleListRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PageRoleList not implemented")
}
func (UnimplementedAuthServer) AddRolesForUser(context.Context, *SetUserForRoleReq) (*RepStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRolesForUser not implemented")
}
func (UnimplementedAuthServer) GetRolesForUser(context.Context, *GetRolesForUserReq) (*GetRolesForUserRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRolesForUser not implemented")
}
func (UnimplementedAuthServer) GetUsersForRole(context.Context, *GetUsersForRoleReq) (*GetUsersForRoleRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersForRole not implemented")
}
func (UnimplementedAuthServer) DeleteRoleForUser(context.Context, *DeleteRoleForUserReq) (*RepStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoleForUser not implemented")
}
func (UnimplementedAuthServer) DeleteRolesForUser(context.Context, *DeleteRolesForUserReq) (*RepStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRolesForUser not implemented")
}
func (UnimplementedAuthServer) GetRolePolicies(context.Context, *GetRolePoliciesReq) (*GetRolePoliciesRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRolePolicies not implemented")
}
func (UnimplementedAuthServer) SetRolePolicies(context.Context, *SetRolePoliciesReq) (*RepStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRolePolicies not implemented")
}
func (UnimplementedAuthServer) CheckAuth(context.Context, *CheckAuthReq) (*RepStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAuth not implemented")
}
func (UnimplementedAuthServer) CreateMenu(context.Context, *CreateMenuReq) (*Menu, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMenu not implemented")
}
func (UnimplementedAuthServer) EditMenu(context.Context, *EditMenuReq) (*RepStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditMenu not implemented")
}
func (UnimplementedAuthServer) DeleteMenu(context.Context, *IdReq) (*RepStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMenu not implemented")
}
func (UnimplementedAuthServer) ListMenu(context.Context, *emptypb.Empty) (*ListMenuRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMenu not implemented")
}
func (UnimplementedAuthServer) ListMenuTree(context.Context, *emptypb.Empty) (*ListMenuRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMenuTree not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_AddRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AddRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_AddRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AddRole(ctx, req.(*AddRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_EditRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).EditRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_EditRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).EditRole(ctx, req.(*EditRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_DelRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).DelRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_DelRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).DelRole(ctx, req.(*DelRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_FullRoleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).FullRoleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_FullRoleList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).FullRoleList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_PageRoleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageRoleListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).PageRoleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_PageRoleList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).PageRoleList(ctx, req.(*PageRoleListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_AddRolesForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserForRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AddRolesForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_AddRolesForUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AddRolesForUser(ctx, req.(*SetUserForRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetRolesForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRolesForUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetRolesForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetRolesForUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetRolesForUser(ctx, req.(*GetRolesForUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetUsersForRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersForRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetUsersForRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetUsersForRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetUsersForRole(ctx, req.(*GetUsersForRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_DeleteRoleForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleForUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).DeleteRoleForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_DeleteRoleForUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).DeleteRoleForUser(ctx, req.(*DeleteRoleForUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_DeleteRolesForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRolesForUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).DeleteRolesForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_DeleteRolesForUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).DeleteRolesForUser(ctx, req.(*DeleteRolesForUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetRolePolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRolePoliciesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetRolePolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetRolePolicies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetRolePolicies(ctx, req.(*GetRolePoliciesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SetRolePolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRolePoliciesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SetRolePolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_SetRolePolicies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SetRolePolicies(ctx, req.(*SetRolePoliciesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CheckAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CheckAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_CheckAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CheckAuth(ctx, req.(*CheckAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CreateMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMenuReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CreateMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_CreateMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CreateMenu(ctx, req.(*CreateMenuReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_EditMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditMenuReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).EditMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_EditMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).EditMenu(ctx, req.(*EditMenuReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_DeleteMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).DeleteMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_DeleteMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).DeleteMenu(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ListMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ListMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_ListMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ListMenu(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ListMenuTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ListMenuTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_ListMenuTree_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ListMenuTree(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.auth.v1.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddRole",
			Handler:    _Auth_AddRole_Handler,
		},
		{
			MethodName: "EditRole",
			Handler:    _Auth_EditRole_Handler,
		},
		{
			MethodName: "DelRole",
			Handler:    _Auth_DelRole_Handler,
		},
		{
			MethodName: "FullRoleList",
			Handler:    _Auth_FullRoleList_Handler,
		},
		{
			MethodName: "PageRoleList",
			Handler:    _Auth_PageRoleList_Handler,
		},
		{
			MethodName: "AddRolesForUser",
			Handler:    _Auth_AddRolesForUser_Handler,
		},
		{
			MethodName: "GetRolesForUser",
			Handler:    _Auth_GetRolesForUser_Handler,
		},
		{
			MethodName: "GetUsersForRole",
			Handler:    _Auth_GetUsersForRole_Handler,
		},
		{
			MethodName: "DeleteRoleForUser",
			Handler:    _Auth_DeleteRoleForUser_Handler,
		},
		{
			MethodName: "DeleteRolesForUser",
			Handler:    _Auth_DeleteRolesForUser_Handler,
		},
		{
			MethodName: "GetRolePolicies",
			Handler:    _Auth_GetRolePolicies_Handler,
		},
		{
			MethodName: "SetRolePolicies",
			Handler:    _Auth_SetRolePolicies_Handler,
		},
		{
			MethodName: "CheckAuth",
			Handler:    _Auth_CheckAuth_Handler,
		},
		{
			MethodName: "CreateMenu",
			Handler:    _Auth_CreateMenu_Handler,
		},
		{
			MethodName: "EditMenu",
			Handler:    _Auth_EditMenu_Handler,
		},
		{
			MethodName: "DeleteMenu",
			Handler:    _Auth_DeleteMenu_Handler,
		},
		{
			MethodName: "ListMenu",
			Handler:    _Auth_ListMenu_Handler,
		},
		{
			MethodName: "ListMenuTree",
			Handler:    _Auth_ListMenuTree_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/v1/auth.proto",
}
