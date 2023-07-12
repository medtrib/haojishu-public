// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: api/auth/v1/auth.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 添加角色
type AddRoleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AddRoleReq) Reset() {
	*x = AddRoleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_v1_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRoleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRoleReq) ProtoMessage() {}

func (x *AddRoleReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRoleReq.ProtoReflect.Descriptor instead.
func (*AddRoleReq) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AddRoleReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 添加角色回复
type AddRoleRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 角色Id
	RoleId uint64 `protobuf:"varint,1,opt,name=roleId,proto3" json:"roleId,omitempty"`
	// 角色名称
	RoleName string `protobuf:"bytes,2,opt,name=roleName,proto3" json:"roleName,omitempty"`
}

func (x *AddRoleRep) Reset() {
	*x = AddRoleRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_v1_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRoleRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRoleRep) ProtoMessage() {}

func (x *AddRoleRep) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRoleRep.ProtoReflect.Descriptor instead.
func (*AddRoleRep) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AddRoleRep) GetRoleId() uint64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *AddRoleRep) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

// 编辑角色
type EditRoleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 角色Id
	RoleId uint64 `protobuf:"varint,1,opt,name=roleId,proto3" json:"roleId,omitempty"`
	// 角色名称
	RoleName string `protobuf:"bytes,2,opt,name=roleName,proto3" json:"roleName,omitempty"`
}

func (x *EditRoleReq) Reset() {
	*x = EditRoleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_v1_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditRoleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditRoleReq) ProtoMessage() {}

func (x *EditRoleReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditRoleReq.ProtoReflect.Descriptor instead.
func (*EditRoleReq) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_proto_rawDescGZIP(), []int{2}
}

func (x *EditRoleReq) GetRoleId() uint64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *EditRoleReq) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

// 角色状态回复
type RoleStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RoleStatus) Reset() {
	*x = RoleStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_v1_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleStatus) ProtoMessage() {}

func (x *RoleStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleStatus.ProtoReflect.Descriptor instead.
func (*RoleStatus) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_proto_rawDescGZIP(), []int{3}
}

func (x *RoleStatus) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// 删除角色
type DelRoleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId uint64 `protobuf:"varint,1,opt,name=roleId,proto3" json:"roleId,omitempty"`
}

func (x *DelRoleReq) Reset() {
	*x = DelRoleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_v1_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelRoleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelRoleReq) ProtoMessage() {}

func (x *DelRoleReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelRoleReq.ProtoReflect.Descriptor instead.
func (*DelRoleReq) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_proto_rawDescGZIP(), []int{4}
}

func (x *DelRoleReq) GetRoleId() uint64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

// 完整角色列表回复
type FullRoleListRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Role `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *FullRoleListRep) Reset() {
	*x = FullRoleListRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_v1_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FullRoleListRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FullRoleListRep) ProtoMessage() {}

func (x *FullRoleListRep) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FullRoleListRep.ProtoReflect.Descriptor instead.
func (*FullRoleListRep) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_proto_rawDescGZIP(), []int{5}
}

func (x *FullRoleListRep) GetList() []*Role {
	if x != nil {
		return x.List
	}
	return nil
}

// 请求 分页角色列表
type PageRoleListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     uint64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize uint64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	RoleName string `protobuf:"bytes,3,opt,name=roleName,proto3" json:"roleName,omitempty"`
}

func (x *PageRoleListReq) Reset() {
	*x = PageRoleListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_v1_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageRoleListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageRoleListReq) ProtoMessage() {}

func (x *PageRoleListReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageRoleListReq.ProtoReflect.Descriptor instead.
func (*PageRoleListReq) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_proto_rawDescGZIP(), []int{6}
}

func (x *PageRoleListReq) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PageRoleListReq) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PageRoleListReq) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

// 分页角色列表
type PageRoleListRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint64  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List  []*Role `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *PageRoleListRep) Reset() {
	*x = PageRoleListRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_v1_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageRoleListRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageRoleListRep) ProtoMessage() {}

func (x *PageRoleListRep) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageRoleListRep.ProtoReflect.Descriptor instead.
func (*PageRoleListRep) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_proto_rawDescGZIP(), []int{7}
}

func (x *PageRoleListRep) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PageRoleListRep) GetList() []*Role {
	if x != nil {
		return x.List
	}
	return nil
}

// 角色
type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 角色Id
	RoleId uint64 `protobuf:"varint,1,opt,name=roleId,proto3" json:"roleId,omitempty"`
	// 角色名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 创建时间
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// 更新时间
	UpdatedAt string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_auth_v1_auth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_api_auth_v1_auth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_api_auth_v1_auth_proto_rawDescGZIP(), []int{8}
}

func (x *Role) GetRoleId() uint64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Role) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Role) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_api_auth_v1_auth_proto protoreflect.FileDescriptor

var file_api_auth_v1_auth_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x0a, 0x41, 0x64, 0x64,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xff,
	0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x56, 0x0a, 0x0b, 0x45, 0x64, 0x69,
	0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20,
	0x00, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x08, 0x72, 0x6f, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07,
	0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x01, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x26, 0x0a, 0x0a, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x2d, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00,
	0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x0f, 0x46, 0x75, 0x6c, 0x6c,
	0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x12, 0x25, 0x0a, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x22, 0x7d, 0x0a, 0x0f, 0x50, 0x61, 0x67, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x32, 0x04, 0x10, 0x32, 0x20, 0x00, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x26, 0x0a, 0x08, 0x72, 0x6f, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07,
	0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x01, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x4e, 0x0a, 0x0f, 0x50, 0x61, 0x67, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x25, 0x0a, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x22, 0x70, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c,
	0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x32, 0xe4, 0x03, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x58, 0x0a, 0x07,
	0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x41,
	0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x5b, 0x0a, 0x08, 0x45, 0x64, 0x69, 0x74, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x64, 0x69, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a,
	0x1a, 0x11, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x45, 0x64, 0x69, 0x74, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x55, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x17,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x2a, 0x10, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x76, 0x31, 0x2f, 0x44, 0x65, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x63, 0x0a, 0x0c, 0x46, 0x75,
	0x6c, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x75, 0x6c, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x76, 0x31, 0x2f, 0x46, 0x75, 0x6c, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x69, 0x0a, 0x0c, 0x50, 0x61, 0x67, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x22, 0x1d, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x2f, 0x0a, 0x0b, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x1e, 0x68, 0x61, 0x6f,
	0x6a, 0x69, 0x73, 0x68, 0x75, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_auth_v1_auth_proto_rawDescOnce sync.Once
	file_api_auth_v1_auth_proto_rawDescData = file_api_auth_v1_auth_proto_rawDesc
)

func file_api_auth_v1_auth_proto_rawDescGZIP() []byte {
	file_api_auth_v1_auth_proto_rawDescOnce.Do(func() {
		file_api_auth_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_auth_v1_auth_proto_rawDescData)
	})
	return file_api_auth_v1_auth_proto_rawDescData
}

var file_api_auth_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_auth_v1_auth_proto_goTypes = []interface{}{
	(*AddRoleReq)(nil),      // 0: api.auth.v1.AddRoleReq
	(*AddRoleRep)(nil),      // 1: api.auth.v1.AddRoleRep
	(*EditRoleReq)(nil),     // 2: api.auth.v1.EditRoleReq
	(*RoleStatus)(nil),      // 3: api.auth.v1.RoleStatus
	(*DelRoleReq)(nil),      // 4: api.auth.v1.DelRoleReq
	(*FullRoleListRep)(nil), // 5: api.auth.v1.FullRoleListRep
	(*PageRoleListReq)(nil), // 6: api.auth.v1.PageRoleListReq
	(*PageRoleListRep)(nil), // 7: api.auth.v1.PageRoleListRep
	(*Role)(nil),            // 8: api.auth.v1.Role
	(*emptypb.Empty)(nil),   // 9: google.protobuf.Empty
}
var file_api_auth_v1_auth_proto_depIdxs = []int32{
	8, // 0: api.auth.v1.FullRoleListRep.list:type_name -> api.auth.v1.Role
	8, // 1: api.auth.v1.PageRoleListRep.list:type_name -> api.auth.v1.Role
	0, // 2: api.auth.v1.Auth.AddRole:input_type -> api.auth.v1.AddRoleReq
	2, // 3: api.auth.v1.Auth.EditRole:input_type -> api.auth.v1.EditRoleReq
	4, // 4: api.auth.v1.Auth.DelRole:input_type -> api.auth.v1.DelRoleReq
	9, // 5: api.auth.v1.Auth.FullRoleList:input_type -> google.protobuf.Empty
	6, // 6: api.auth.v1.Auth.PageRoleList:input_type -> api.auth.v1.PageRoleListReq
	1, // 7: api.auth.v1.Auth.AddRole:output_type -> api.auth.v1.AddRoleRep
	3, // 8: api.auth.v1.Auth.EditRole:output_type -> api.auth.v1.RoleStatus
	3, // 9: api.auth.v1.Auth.DelRole:output_type -> api.auth.v1.RoleStatus
	5, // 10: api.auth.v1.Auth.FullRoleList:output_type -> api.auth.v1.FullRoleListRep
	7, // 11: api.auth.v1.Auth.PageRoleList:output_type -> api.auth.v1.PageRoleListRep
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_auth_v1_auth_proto_init() }
func file_api_auth_v1_auth_proto_init() {
	if File_api_auth_v1_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_auth_v1_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRoleReq); i {
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
		file_api_auth_v1_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRoleRep); i {
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
		file_api_auth_v1_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditRoleReq); i {
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
		file_api_auth_v1_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleStatus); i {
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
		file_api_auth_v1_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelRoleReq); i {
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
		file_api_auth_v1_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FullRoleListRep); i {
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
		file_api_auth_v1_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageRoleListReq); i {
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
		file_api_auth_v1_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageRoleListRep); i {
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
		file_api_auth_v1_auth_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
			RawDescriptor: file_api_auth_v1_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_auth_v1_auth_proto_goTypes,
		DependencyIndexes: file_api_auth_v1_auth_proto_depIdxs,
		MessageInfos:      file_api_auth_v1_auth_proto_msgTypes,
	}.Build()
	File_api_auth_v1_auth_proto = out.File
	file_api_auth_v1_auth_proto_rawDesc = nil
	file_api_auth_v1_auth_proto_goTypes = nil
	file_api_auth_v1_auth_proto_depIdxs = nil
}
