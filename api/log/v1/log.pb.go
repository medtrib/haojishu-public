// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: api/log/v1/log.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 请求 - 创建
type CreateLogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// trace id
	TraceId string `protobuf:"bytes,2,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// http/rpc
	Component string `protobuf:"bytes,3,opt,name=component,proto3" json:"component,omitempty"`
	// 表
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// 响应时长
	Latency string `protobuf:"bytes,5,opt,name=latency,proto3" json:"latency,omitempty"`
	// 用户id
	UserId string `protobuf:"bytes,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// 请求方式
	Method string `protobuf:"bytes,7,opt,name=method,proto3" json:"method,omitempty"`
	// 请求路径
	Path string `protobuf:"bytes,8,opt,name=path,proto3" json:"path,omitempty"`
	// 请求内容
	Request string `protobuf:"bytes,9,opt,name=request,proto3" json:"request,omitempty"`
	// 响应内容
	Code string `protobuf:"bytes,10,opt,name=code,proto3" json:"code,omitempty"`
	// 提示信息
	Reason string `protobuf:"bytes,11,opt,name=reason,proto3" json:"reason,omitempty"`
	// 请求ip
	Ip string `protobuf:"bytes,12,opt,name=ip,proto3" json:"ip,omitempty"`
	// 创建时间
	CreatedAt string `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// kratos操作路径
	Operation string `protobuf:"bytes,14,opt,name=operation,proto3" json:"operation,omitempty"`
	// 用户名
	Username string `protobuf:"bytes,15,opt,name=username,proto3" json:"username,omitempty"`
	// 角色
	Role string `protobuf:"bytes,16,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *CreateLogReq) Reset() {
	*x = CreateLogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_log_v1_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLogReq) ProtoMessage() {}

func (x *CreateLogReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_log_v1_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLogReq.ProtoReflect.Descriptor instead.
func (*CreateLogReq) Descriptor() ([]byte, []int) {
	return file_api_log_v1_log_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLogReq) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *CreateLogReq) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *CreateLogReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateLogReq) GetLatency() string {
	if x != nil {
		return x.Latency
	}
	return ""
}

func (x *CreateLogReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateLogReq) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *CreateLogReq) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CreateLogReq) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

func (x *CreateLogReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateLogReq) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *CreateLogReq) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *CreateLogReq) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CreateLogReq) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *CreateLogReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateLogReq) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

// 角色状态
type RepStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RepStatus) Reset() {
	*x = RepStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_log_v1_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepStatus) ProtoMessage() {}

func (x *RepStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_log_v1_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepStatus.ProtoReflect.Descriptor instead.
func (*RepStatus) Descriptor() ([]byte, []int) {
	return file_api_log_v1_log_proto_rawDescGZIP(), []int{1}
}

func (x *RepStatus) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// 响应 - 日志信息
type LogInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 日志id
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// trace id
	TraceId string `protobuf:"bytes,2,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// http/rpc
	Component string `protobuf:"bytes,3,opt,name=component,proto3" json:"component,omitempty"`
	// kratos操作路径
	Operation string `protobuf:"bytes,5,opt,name=operation,proto3" json:"operation,omitempty"`
	// 用户id
	UserId string `protobuf:"bytes,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// 请求方式
	Method string `protobuf:"bytes,7,opt,name=method,proto3" json:"method,omitempty"`
	// 请求路径
	Path string `protobuf:"bytes,8,opt,name=path,proto3" json:"path,omitempty"`
	// 请求内容
	Request string `protobuf:"bytes,9,opt,name=request,proto3" json:"request,omitempty"`
	// 响应内容
	Code string `protobuf:"bytes,10,opt,name=code,proto3" json:"code,omitempty"`
	// 提示信息
	Reason string `protobuf:"bytes,11,opt,name=reason,proto3" json:"reason,omitempty"`
	// 请求ip
	Ip string `protobuf:"bytes,12,opt,name=ip,proto3" json:"ip,omitempty"`
	// 创建时间
	CreatedAt string `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// 名称
	Name string `protobuf:"bytes,14,opt,name=name,proto3" json:"name,omitempty"`
	// 响应时长
	Latency string `protobuf:"bytes,15,opt,name=latency,proto3" json:"latency,omitempty"`
	// 用户名
	Username string `protobuf:"bytes,16,opt,name=username,proto3" json:"username,omitempty"`
	// 角色
	Role string `protobuf:"bytes,17,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *LogInfo) Reset() {
	*x = LogInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_log_v1_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogInfo) ProtoMessage() {}

func (x *LogInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_log_v1_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogInfo.ProtoReflect.Descriptor instead.
func (*LogInfo) Descriptor() ([]byte, []int) {
	return file_api_log_v1_log_proto_rawDescGZIP(), []int{2}
}

func (x *LogInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LogInfo) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *LogInfo) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *LogInfo) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *LogInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *LogInfo) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *LogInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *LogInfo) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

func (x *LogInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *LogInfo) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *LogInfo) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *LogInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *LogInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LogInfo) GetLatency() string {
	if x != nil {
		return x.Latency
	}
	return ""
}

func (x *LogInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LogInfo) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

// 列表
type GetLogListRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	// 列表
	List []*LogInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *GetLogListRep) Reset() {
	*x = GetLogListRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_log_v1_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogListRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogListRep) ProtoMessage() {}

func (x *GetLogListRep) ProtoReflect() protoreflect.Message {
	mi := &file_api_log_v1_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogListRep.ProtoReflect.Descriptor instead.
func (*GetLogListRep) Descriptor() ([]byte, []int) {
	return file_api_log_v1_log_proto_rawDescGZIP(), []int{3}
}

func (x *GetLogListRep) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GetLogListRep) GetList() []*LogInfo {
	if x != nil {
		return x.List
	}
	return nil
}

// 列表
type GetLogListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 主键id
	Page int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	// 页记录数
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	// 名称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// userId
	UserId string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// 用户名
	Username string `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
	// 角色
	Role string `protobuf:"bytes,6,opt,name=role,proto3" json:"role,omitempty"`
	// 路径
	Operation string `protobuf:"bytes,7,opt,name=operation,proto3" json:"operation,omitempty"`
	// ip
	Ip string `protobuf:"bytes,8,opt,name=ip,proto3" json:"ip,omitempty"`
	// trace_id
	TraceId string `protobuf:"bytes,9,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
}

func (x *GetLogListReq) Reset() {
	*x = GetLogListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_log_v1_log_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogListReq) ProtoMessage() {}

func (x *GetLogListReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_log_v1_log_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogListReq.ProtoReflect.Descriptor instead.
func (*GetLogListReq) Descriptor() ([]byte, []int) {
	return file_api_log_v1_log_proto_rawDescGZIP(), []int{4}
}

func (x *GetLogListReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetLogListReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetLogListReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetLogListReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetLogListReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetLogListReq) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *GetLogListReq) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *GetLogListReq) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *GetLogListReq) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

var File_api_log_v1_log_proto protoreflect.FileDescriptor

var file_api_log_v1_log_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x02, 0x0a, 0x0c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x70, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x25, 0x0a, 0x09, 0x52, 0x65, 0x70,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x88, 0x03, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x70, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x4e, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x27, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f,
	0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0xf7, 0x01, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x22, 0x02, 0x20, 0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x32, 0xba, 0x01, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x5a, 0x0a,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1c, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x57, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x22, 0x13, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6c,
	0x6f, 0x67, 0x42, 0x40, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x65, 0x64, 0x74, 0x72, 0x69, 0x62, 0x2f, 0x68, 0x61, 0x6f, 0x6a, 0x69, 0x73, 0x68, 0x75, 0x2d,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_log_v1_log_proto_rawDescOnce sync.Once
	file_api_log_v1_log_proto_rawDescData = file_api_log_v1_log_proto_rawDesc
)

func file_api_log_v1_log_proto_rawDescGZIP() []byte {
	file_api_log_v1_log_proto_rawDescOnce.Do(func() {
		file_api_log_v1_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_log_v1_log_proto_rawDescData)
	})
	return file_api_log_v1_log_proto_rawDescData
}

var file_api_log_v1_log_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_log_v1_log_proto_goTypes = []interface{}{
	(*CreateLogReq)(nil),  // 0: api.log.v1.CreateLogReq
	(*RepStatus)(nil),     // 1: api.log.v1.RepStatus
	(*LogInfo)(nil),       // 2: api.log.v1.LogInfo
	(*GetLogListRep)(nil), // 3: api.log.v1.GetLogListRep
	(*GetLogListReq)(nil), // 4: api.log.v1.GetLogListReq
}
var file_api_log_v1_log_proto_depIdxs = []int32{
	2, // 0: api.log.v1.GetLogListRep.list:type_name -> api.log.v1.LogInfo
	0, // 1: api.log.v1.Log.CreateLog:input_type -> api.log.v1.CreateLogReq
	4, // 2: api.log.v1.Log.GetLogList:input_type -> api.log.v1.GetLogListReq
	1, // 3: api.log.v1.Log.CreateLog:output_type -> api.log.v1.RepStatus
	3, // 4: api.log.v1.Log.GetLogList:output_type -> api.log.v1.GetLogListRep
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_log_v1_log_proto_init() }
func file_api_log_v1_log_proto_init() {
	if File_api_log_v1_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_log_v1_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLogReq); i {
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
		file_api_log_v1_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepStatus); i {
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
		file_api_log_v1_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogInfo); i {
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
		file_api_log_v1_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogListRep); i {
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
		file_api_log_v1_log_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogListReq); i {
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
			RawDescriptor: file_api_log_v1_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_log_v1_log_proto_goTypes,
		DependencyIndexes: file_api_log_v1_log_proto_depIdxs,
		MessageInfos:      file_api_log_v1_log_proto_msgTypes,
	}.Build()
	File_api_log_v1_log_proto = out.File
	file_api_log_v1_log_proto_rawDesc = nil
	file_api_log_v1_log_proto_goTypes = nil
	file_api_log_v1_log_proto_depIdxs = nil
}
