// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.3
// - protoc             v3.19.4
// source: auth/v1/auth.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationAuthAddRole = "/api.auth.v1.Auth/AddRole"
const OperationAuthAddRolesForUser = "/api.auth.v1.Auth/AddRolesForUser"
const OperationAuthCheckAuth = "/api.auth.v1.Auth/CheckAuth"
const OperationAuthCreateApi = "/api.auth.v1.Auth/CreateApi"
const OperationAuthCreateMenu = "/api.auth.v1.Auth/CreateMenu"
const OperationAuthCreateRoleMenu = "/api.auth.v1.Auth/CreateRoleMenu"
const OperationAuthCreateRoleMenuBtn = "/api.auth.v1.Auth/CreateRoleMenuBtn"
const OperationAuthDelRole = "/api.auth.v1.Auth/DelRole"
const OperationAuthDeleteApi = "/api.auth.v1.Auth/DeleteApi"
const OperationAuthDeleteMenu = "/api.auth.v1.Auth/DeleteMenu"
const OperationAuthDeleteRoleForUser = "/api.auth.v1.Auth/DeleteRoleForUser"
const OperationAuthDeleteRolesForUser = "/api.auth.v1.Auth/DeleteRolesForUser"
const OperationAuthEditMenu = "/api.auth.v1.Auth/EditMenu"
const OperationAuthEditRole = "/api.auth.v1.Auth/EditRole"
const OperationAuthFullRoleList = "/api.auth.v1.Auth/FullRoleList"
const OperationAuthGetApiList = "/api.auth.v1.Auth/GetApiList"
const OperationAuthGetApiPageList = "/api.auth.v1.Auth/GetApiPageList"
const OperationAuthGetRoleMenuBtn = "/api.auth.v1.Auth/GetRoleMenuBtn"
const OperationAuthGetRolePolicies = "/api.auth.v1.Auth/GetRolePolicies"
const OperationAuthGetRolesForUser = "/api.auth.v1.Auth/GetRolesForUser"
const OperationAuthGetUsersForRole = "/api.auth.v1.Auth/GetUsersForRole"
const OperationAuthListMenu = "/api.auth.v1.Auth/ListMenu"
const OperationAuthListMenuTree = "/api.auth.v1.Auth/ListMenuTree"
const OperationAuthListRoleMenu = "/api.auth.v1.Auth/ListRoleMenu"
const OperationAuthListRoleMenuTree = "/api.auth.v1.Auth/ListRoleMenuTree"
const OperationAuthPageRoleList = "/api.auth.v1.Auth/PageRoleList"
const OperationAuthSetRolePolicies = "/api.auth.v1.Auth/SetRolePolicies"
const OperationAuthUpdateApi = "/api.auth.v1.Auth/UpdateApi"

type AuthHTTPServer interface {
	// AddRole 添加角色
	AddRole(context.Context, *AddRoleReq) (*Role, error)
	// AddRolesForUser 给用户设置角色
	AddRolesForUser(context.Context, *SetUserForRoleReq) (*RepStatus, error)
	// CheckAuth 检查权限
	CheckAuth(context.Context, *CheckAuthReq) (*RepStatus, error)
	// CreateApi Api创建
	CreateApi(context.Context, *CreateApiReq) (*ApiInfo, error)
	// CreateMenu 创建菜单
	CreateMenu(context.Context, *CreateMenuReq) (*Menu, error)
	// CreateRoleMenu 角色菜单添加
	CreateRoleMenu(context.Context, *CreateRoleMenuReq) (*RepStatus, error)
	// CreateRoleMenuBtn 角色菜单按钮添加
	CreateRoleMenuBtn(context.Context, *CreateRoleMenuBtnReq) (*RepStatus, error)
	// DelRole 删除角色
	DelRole(context.Context, *DelRoleReq) (*RepStatus, error)
	// DeleteApi Api删除
	DeleteApi(context.Context, *DeleteApiReq) (*RepStatus, error)
	// DeleteMenu 删除菜单
	DeleteMenu(context.Context, *IdReq) (*RepStatus, error)
	// DeleteRoleForUser 删除单个用户角色(如果需要删除单个用户的某个角色用这个)
	DeleteRoleForUser(context.Context, *DeleteRoleForUserReq) (*RepStatus, error)
	// DeleteRolesForUser 删除多个用户角色(删除传递用户的所有角色)
	DeleteRolesForUser(context.Context, *DeleteRolesForUserReq) (*RepStatus, error)
	// EditMenu 更新菜单
	EditMenu(context.Context, *EditMenuReq) (*RepStatus, error)
	// EditRole 编辑角色
	EditRole(context.Context, *EditRoleReq) (*RepStatus, error)
	// FullRoleList 获取角色列表(完整)
	FullRoleList(context.Context, *emptypb.Empty) (*FullRoleListRep, error)
	// GetApiList Api列表
	GetApiList(context.Context, *emptypb.Empty) (*GetApiListRep, error)
	// GetApiPageList Api列表分页
	GetApiPageList(context.Context, *GetApiListReq) (*GetApiListPageRep, error)
	// GetRoleMenuBtn 角色菜单按钮 - 列表
	GetRoleMenuBtn(context.Context, *GetRoleMenuBtnReq) (*GetRoleMenuBtnRep, error)
	// GetRolePolicies 获取角色有那些权限
	GetRolePolicies(context.Context, *GetRolePoliciesReq) (*GetRolePoliciesRep, error)
	// GetRolesForUser 获取用户角色
	GetRolesForUser(context.Context, *GetRolesForUserReq) (*GetRolesForUserRep, error)
	// GetUsersForRole 获取角色有那些用户
	GetUsersForRole(context.Context, *GetUsersForRoleReq) (*GetUsersForRoleRep, error)
	// ListMenu 菜单列表(完整)
	ListMenu(context.Context, *emptypb.Empty) (*ListMenuRep, error)
	// ListMenuTree 菜单列表(树)
	ListMenuTree(context.Context, *emptypb.Empty) (*ListMenuRep, error)
	// ListRoleMenu 角色菜单列表(完整)
	ListRoleMenu(context.Context, *ListRoleMenuReq) (*ListMenuRep, error)
	// ListRoleMenuTree 角色菜单 - 树状结构
	ListRoleMenuTree(context.Context, *ListRoleMenuReq) (*ListMenuRep, error)
	// PageRoleList 获取角色列表(分页)
	PageRoleList(context.Context, *PageRoleListReq) (*PageRoleListRep, error)
	// SetRolePolicies 设置角色权限
	SetRolePolicies(context.Context, *SetRolePoliciesReq) (*RepStatus, error)
	// UpdateApi Api更新
	UpdateApi(context.Context, *UpdateApiReq) (*RepStatus, error)
}

func RegisterAuthHTTPServer(s *http.Server, srv AuthHTTPServer) {
	r := s.Route("/")
	r.POST("/auth/v1/AddRole", _Auth_AddRole0_HTTP_Handler(srv))
	r.PUT("/auth/v1/EditRole", _Auth_EditRole0_HTTP_Handler(srv))
	r.DELETE("/auth/v1/DelRole", _Auth_DelRole0_HTTP_Handler(srv))
	r.GET("/auth/v1/FullRoleList", _Auth_FullRoleList0_HTTP_Handler(srv))
	r.GET("/auth/v1/PageRoleList", _Auth_PageRoleList0_HTTP_Handler(srv))
	r.GET("/auth/v1/AddRolesForUser", _Auth_AddRolesForUser0_HTTP_Handler(srv))
	r.GET("/auth/v1/GetRolesForUser", _Auth_GetRolesForUser0_HTTP_Handler(srv))
	r.GET("/auth/v1/GetUsersForRole", _Auth_GetUsersForRole0_HTTP_Handler(srv))
	r.DELETE("/auth/v1/DeleteRoleForUser", _Auth_DeleteRoleForUser0_HTTP_Handler(srv))
	r.DELETE("/auth/v1/DeleteRolesForUser", _Auth_DeleteRolesForUser0_HTTP_Handler(srv))
	r.GET("/auth/v1/GetRolePolicies", _Auth_GetRolePolicies0_HTTP_Handler(srv))
	r.POST("/auth/v1/SetRolePolicies", _Auth_SetRolePolicies0_HTTP_Handler(srv))
	r.GET("/auth/v1/CheckAuth", _Auth_CheckAuth0_HTTP_Handler(srv))
	r.POST("/auth/v1/CreateMenu", _Auth_CreateMenu0_HTTP_Handler(srv))
	r.PUT("/auth/v1/EditMenu", _Auth_EditMenu0_HTTP_Handler(srv))
	r.DELETE("/auth/v1/DeleteMenu", _Auth_DeleteMenu0_HTTP_Handler(srv))
	r.GET("/auth/v1/ListMenu", _Auth_ListMenu0_HTTP_Handler(srv))
	r.GET("/auth/v1/ListMenuTree", _Auth_ListMenuTree0_HTTP_Handler(srv))
	r.POST("/auth/v1/CreateRoleMenu", _Auth_CreateRoleMenu0_HTTP_Handler(srv))
	r.POST("/auth/v1/CreateMenuBtn", _Auth_CreateRoleMenuBtn0_HTTP_Handler(srv))
	r.GET("/auth/v1/ListRoleMenu", _Auth_ListRoleMenu0_HTTP_Handler(srv))
	r.GET("/auth/v1/ListRoleMenuTree", _Auth_ListRoleMenuTree0_HTTP_Handler(srv))
	r.GET("/auth/v1/GetRoleMenuBtn", _Auth_GetRoleMenuBtn0_HTTP_Handler(srv))
	r.GET("/auth/v1/GetApiList", _Auth_GetApiList0_HTTP_Handler(srv))
	r.GET("/auth/v1/GetApiList", _Auth_GetApiPageList0_HTTP_Handler(srv))
	r.POST("/auth/v1/CreateApi", _Auth_CreateApi0_HTTP_Handler(srv))
	r.PUT("/auth/v1/UpdateApi", _Auth_UpdateApi0_HTTP_Handler(srv))
	r.DELETE("/api", _Auth_DeleteApi0_HTTP_Handler(srv))
}

func _Auth_AddRole0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddRoleReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthAddRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddRole(ctx, req.(*AddRoleReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Role)
		return ctx.Result(200, reply)
	}
}

func _Auth_EditRole0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in EditRoleReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthEditRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.EditRole(ctx, req.(*EditRoleReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RepStatus)
		return ctx.Result(200, reply)
	}
}

func _Auth_DelRole0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DelRoleReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthDelRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DelRole(ctx, req.(*DelRoleReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RepStatus)
		return ctx.Result(200, reply)
	}
}

func _Auth_FullRoleList0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthFullRoleList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FullRoleList(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FullRoleListRep)
		return ctx.Result(200, reply)
	}
}

func _Auth_PageRoleList0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PageRoleListReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthPageRoleList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PageRoleList(ctx, req.(*PageRoleListReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PageRoleListRep)
		return ctx.Result(200, reply)
	}
}

func _Auth_AddRolesForUser0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetUserForRoleReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthAddRolesForUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddRolesForUser(ctx, req.(*SetUserForRoleReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RepStatus)
		return ctx.Result(200, reply)
	}
}

func _Auth_GetRolesForUser0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetRolesForUserReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthGetRolesForUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetRolesForUser(ctx, req.(*GetRolesForUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetRolesForUserRep)
		return ctx.Result(200, reply)
	}
}

func _Auth_GetUsersForRole0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUsersForRoleReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthGetUsersForRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUsersForRole(ctx, req.(*GetUsersForRoleReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUsersForRoleRep)
		return ctx.Result(200, reply)
	}
}

func _Auth_DeleteRoleForUser0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteRoleForUserReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthDeleteRoleForUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteRoleForUser(ctx, req.(*DeleteRoleForUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RepStatus)
		return ctx.Result(200, reply)
	}
}

func _Auth_DeleteRolesForUser0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteRolesForUserReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthDeleteRolesForUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteRolesForUser(ctx, req.(*DeleteRolesForUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RepStatus)
		return ctx.Result(200, reply)
	}
}

func _Auth_GetRolePolicies0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetRolePoliciesReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthGetRolePolicies)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetRolePolicies(ctx, req.(*GetRolePoliciesReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetRolePoliciesRep)
		return ctx.Result(200, reply)
	}
}

func _Auth_SetRolePolicies0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetRolePoliciesReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthSetRolePolicies)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetRolePolicies(ctx, req.(*SetRolePoliciesReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RepStatus)
		return ctx.Result(200, reply)
	}
}

func _Auth_CheckAuth0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CheckAuthReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthCheckAuth)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CheckAuth(ctx, req.(*CheckAuthReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RepStatus)
		return ctx.Result(200, reply)
	}
}

func _Auth_CreateMenu0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateMenuReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthCreateMenu)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateMenu(ctx, req.(*CreateMenuReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Menu)
		return ctx.Result(200, reply)
	}
}

func _Auth_EditMenu0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in EditMenuReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthEditMenu)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.EditMenu(ctx, req.(*EditMenuReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RepStatus)
		return ctx.Result(200, reply)
	}
}

func _Auth_DeleteMenu0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IdReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthDeleteMenu)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteMenu(ctx, req.(*IdReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RepStatus)
		return ctx.Result(200, reply)
	}
}

func _Auth_ListMenu0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthListMenu)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListMenu(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMenuRep)
		return ctx.Result(200, reply)
	}
}

func _Auth_ListMenuTree0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthListMenuTree)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListMenuTree(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMenuRep)
		return ctx.Result(200, reply)
	}
}

func _Auth_CreateRoleMenu0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateRoleMenuReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthCreateRoleMenu)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateRoleMenu(ctx, req.(*CreateRoleMenuReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RepStatus)
		return ctx.Result(200, reply)
	}
}

func _Auth_CreateRoleMenuBtn0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateRoleMenuBtnReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthCreateRoleMenuBtn)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateRoleMenuBtn(ctx, req.(*CreateRoleMenuBtnReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RepStatus)
		return ctx.Result(200, reply)
	}
}

func _Auth_ListRoleMenu0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListRoleMenuReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthListRoleMenu)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListRoleMenu(ctx, req.(*ListRoleMenuReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMenuRep)
		return ctx.Result(200, reply)
	}
}

func _Auth_ListRoleMenuTree0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListRoleMenuReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthListRoleMenuTree)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListRoleMenuTree(ctx, req.(*ListRoleMenuReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMenuRep)
		return ctx.Result(200, reply)
	}
}

func _Auth_GetRoleMenuBtn0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetRoleMenuBtnReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthGetRoleMenuBtn)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetRoleMenuBtn(ctx, req.(*GetRoleMenuBtnReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetRoleMenuBtnRep)
		return ctx.Result(200, reply)
	}
}

func _Auth_GetApiList0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthGetApiList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetApiList(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetApiListRep)
		return ctx.Result(200, reply)
	}
}

func _Auth_GetApiPageList0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetApiListReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthGetApiPageList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetApiPageList(ctx, req.(*GetApiListReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetApiListPageRep)
		return ctx.Result(200, reply)
	}
}

func _Auth_CreateApi0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateApiReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthCreateApi)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateApi(ctx, req.(*CreateApiReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ApiInfo)
		return ctx.Result(200, reply)
	}
}

func _Auth_UpdateApi0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateApiReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthUpdateApi)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateApi(ctx, req.(*UpdateApiReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RepStatus)
		return ctx.Result(200, reply)
	}
}

func _Auth_DeleteApi0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteApiReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthDeleteApi)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteApi(ctx, req.(*DeleteApiReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RepStatus)
		return ctx.Result(200, reply)
	}
}

type AuthHTTPClient interface {
	AddRole(ctx context.Context, req *AddRoleReq, opts ...http.CallOption) (rsp *Role, err error)
	AddRolesForUser(ctx context.Context, req *SetUserForRoleReq, opts ...http.CallOption) (rsp *RepStatus, err error)
	CheckAuth(ctx context.Context, req *CheckAuthReq, opts ...http.CallOption) (rsp *RepStatus, err error)
	CreateApi(ctx context.Context, req *CreateApiReq, opts ...http.CallOption) (rsp *ApiInfo, err error)
	CreateMenu(ctx context.Context, req *CreateMenuReq, opts ...http.CallOption) (rsp *Menu, err error)
	CreateRoleMenu(ctx context.Context, req *CreateRoleMenuReq, opts ...http.CallOption) (rsp *RepStatus, err error)
	CreateRoleMenuBtn(ctx context.Context, req *CreateRoleMenuBtnReq, opts ...http.CallOption) (rsp *RepStatus, err error)
	DelRole(ctx context.Context, req *DelRoleReq, opts ...http.CallOption) (rsp *RepStatus, err error)
	DeleteApi(ctx context.Context, req *DeleteApiReq, opts ...http.CallOption) (rsp *RepStatus, err error)
	DeleteMenu(ctx context.Context, req *IdReq, opts ...http.CallOption) (rsp *RepStatus, err error)
	DeleteRoleForUser(ctx context.Context, req *DeleteRoleForUserReq, opts ...http.CallOption) (rsp *RepStatus, err error)
	DeleteRolesForUser(ctx context.Context, req *DeleteRolesForUserReq, opts ...http.CallOption) (rsp *RepStatus, err error)
	EditMenu(ctx context.Context, req *EditMenuReq, opts ...http.CallOption) (rsp *RepStatus, err error)
	EditRole(ctx context.Context, req *EditRoleReq, opts ...http.CallOption) (rsp *RepStatus, err error)
	FullRoleList(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *FullRoleListRep, err error)
	GetApiList(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *GetApiListRep, err error)
	GetApiPageList(ctx context.Context, req *GetApiListReq, opts ...http.CallOption) (rsp *GetApiListPageRep, err error)
	GetRoleMenuBtn(ctx context.Context, req *GetRoleMenuBtnReq, opts ...http.CallOption) (rsp *GetRoleMenuBtnRep, err error)
	GetRolePolicies(ctx context.Context, req *GetRolePoliciesReq, opts ...http.CallOption) (rsp *GetRolePoliciesRep, err error)
	GetRolesForUser(ctx context.Context, req *GetRolesForUserReq, opts ...http.CallOption) (rsp *GetRolesForUserRep, err error)
	GetUsersForRole(ctx context.Context, req *GetUsersForRoleReq, opts ...http.CallOption) (rsp *GetUsersForRoleRep, err error)
	ListMenu(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *ListMenuRep, err error)
	ListMenuTree(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *ListMenuRep, err error)
	ListRoleMenu(ctx context.Context, req *ListRoleMenuReq, opts ...http.CallOption) (rsp *ListMenuRep, err error)
	ListRoleMenuTree(ctx context.Context, req *ListRoleMenuReq, opts ...http.CallOption) (rsp *ListMenuRep, err error)
	PageRoleList(ctx context.Context, req *PageRoleListReq, opts ...http.CallOption) (rsp *PageRoleListRep, err error)
	SetRolePolicies(ctx context.Context, req *SetRolePoliciesReq, opts ...http.CallOption) (rsp *RepStatus, err error)
	UpdateApi(ctx context.Context, req *UpdateApiReq, opts ...http.CallOption) (rsp *RepStatus, err error)
}

type AuthHTTPClientImpl struct {
	cc *http.Client
}

func NewAuthHTTPClient(client *http.Client) AuthHTTPClient {
	return &AuthHTTPClientImpl{client}
}

func (c *AuthHTTPClientImpl) AddRole(ctx context.Context, in *AddRoleReq, opts ...http.CallOption) (*Role, error) {
	var out Role
	pattern := "/auth/v1/AddRole"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthAddRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) AddRolesForUser(ctx context.Context, in *SetUserForRoleReq, opts ...http.CallOption) (*RepStatus, error) {
	var out RepStatus
	pattern := "/auth/v1/AddRolesForUser"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAuthAddRolesForUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) CheckAuth(ctx context.Context, in *CheckAuthReq, opts ...http.CallOption) (*RepStatus, error) {
	var out RepStatus
	pattern := "/auth/v1/CheckAuth"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAuthCheckAuth))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) CreateApi(ctx context.Context, in *CreateApiReq, opts ...http.CallOption) (*ApiInfo, error) {
	var out ApiInfo
	pattern := "/auth/v1/CreateApi"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthCreateApi))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) CreateMenu(ctx context.Context, in *CreateMenuReq, opts ...http.CallOption) (*Menu, error) {
	var out Menu
	pattern := "/auth/v1/CreateMenu"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthCreateMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) CreateRoleMenu(ctx context.Context, in *CreateRoleMenuReq, opts ...http.CallOption) (*RepStatus, error) {
	var out RepStatus
	pattern := "/auth/v1/CreateRoleMenu"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthCreateRoleMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) CreateRoleMenuBtn(ctx context.Context, in *CreateRoleMenuBtnReq, opts ...http.CallOption) (*RepStatus, error) {
	var out RepStatus
	pattern := "/auth/v1/CreateMenuBtn"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthCreateRoleMenuBtn))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) DelRole(ctx context.Context, in *DelRoleReq, opts ...http.CallOption) (*RepStatus, error) {
	var out RepStatus
	pattern := "/auth/v1/DelRole"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAuthDelRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) DeleteApi(ctx context.Context, in *DeleteApiReq, opts ...http.CallOption) (*RepStatus, error) {
	var out RepStatus
	pattern := "/api"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAuthDeleteApi))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) DeleteMenu(ctx context.Context, in *IdReq, opts ...http.CallOption) (*RepStatus, error) {
	var out RepStatus
	pattern := "/auth/v1/DeleteMenu"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAuthDeleteMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) DeleteRoleForUser(ctx context.Context, in *DeleteRoleForUserReq, opts ...http.CallOption) (*RepStatus, error) {
	var out RepStatus
	pattern := "/auth/v1/DeleteRoleForUser"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAuthDeleteRoleForUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) DeleteRolesForUser(ctx context.Context, in *DeleteRolesForUserReq, opts ...http.CallOption) (*RepStatus, error) {
	var out RepStatus
	pattern := "/auth/v1/DeleteRolesForUser"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAuthDeleteRolesForUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) EditMenu(ctx context.Context, in *EditMenuReq, opts ...http.CallOption) (*RepStatus, error) {
	var out RepStatus
	pattern := "/auth/v1/EditMenu"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthEditMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) EditRole(ctx context.Context, in *EditRoleReq, opts ...http.CallOption) (*RepStatus, error) {
	var out RepStatus
	pattern := "/auth/v1/EditRole"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthEditRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) FullRoleList(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*FullRoleListRep, error) {
	var out FullRoleListRep
	pattern := "/auth/v1/FullRoleList"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAuthFullRoleList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) GetApiList(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*GetApiListRep, error) {
	var out GetApiListRep
	pattern := "/auth/v1/GetApiList"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAuthGetApiList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) GetApiPageList(ctx context.Context, in *GetApiListReq, opts ...http.CallOption) (*GetApiListPageRep, error) {
	var out GetApiListPageRep
	pattern := "/auth/v1/GetApiList"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAuthGetApiPageList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) GetRoleMenuBtn(ctx context.Context, in *GetRoleMenuBtnReq, opts ...http.CallOption) (*GetRoleMenuBtnRep, error) {
	var out GetRoleMenuBtnRep
	pattern := "/auth/v1/GetRoleMenuBtn"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAuthGetRoleMenuBtn))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) GetRolePolicies(ctx context.Context, in *GetRolePoliciesReq, opts ...http.CallOption) (*GetRolePoliciesRep, error) {
	var out GetRolePoliciesRep
	pattern := "/auth/v1/GetRolePolicies"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAuthGetRolePolicies))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) GetRolesForUser(ctx context.Context, in *GetRolesForUserReq, opts ...http.CallOption) (*GetRolesForUserRep, error) {
	var out GetRolesForUserRep
	pattern := "/auth/v1/GetRolesForUser"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAuthGetRolesForUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) GetUsersForRole(ctx context.Context, in *GetUsersForRoleReq, opts ...http.CallOption) (*GetUsersForRoleRep, error) {
	var out GetUsersForRoleRep
	pattern := "/auth/v1/GetUsersForRole"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAuthGetUsersForRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) ListMenu(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*ListMenuRep, error) {
	var out ListMenuRep
	pattern := "/auth/v1/ListMenu"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAuthListMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) ListMenuTree(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*ListMenuRep, error) {
	var out ListMenuRep
	pattern := "/auth/v1/ListMenuTree"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAuthListMenuTree))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) ListRoleMenu(ctx context.Context, in *ListRoleMenuReq, opts ...http.CallOption) (*ListMenuRep, error) {
	var out ListMenuRep
	pattern := "/auth/v1/ListRoleMenu"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAuthListRoleMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) ListRoleMenuTree(ctx context.Context, in *ListRoleMenuReq, opts ...http.CallOption) (*ListMenuRep, error) {
	var out ListMenuRep
	pattern := "/auth/v1/ListRoleMenuTree"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAuthListRoleMenuTree))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) PageRoleList(ctx context.Context, in *PageRoleListReq, opts ...http.CallOption) (*PageRoleListRep, error) {
	var out PageRoleListRep
	pattern := "/auth/v1/PageRoleList"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAuthPageRoleList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) SetRolePolicies(ctx context.Context, in *SetRolePoliciesReq, opts ...http.CallOption) (*RepStatus, error) {
	var out RepStatus
	pattern := "/auth/v1/SetRolePolicies"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthSetRolePolicies))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) UpdateApi(ctx context.Context, in *UpdateApiReq, opts ...http.CallOption) (*RepStatus, error) {
	var out RepStatus
	pattern := "/auth/v1/UpdateApi"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthUpdateApi))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
