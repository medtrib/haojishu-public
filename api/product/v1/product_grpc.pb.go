// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: product/v1/product.proto

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
	Product_CreateTerminal_FullMethodName   = "/api.product.v1.Product/CreateTerminal"
	Product_EditTerminal_FullMethodName     = "/api.product.v1.Product/EditTerminal"
	Product_DelTerminal_FullMethodName      = "/api.product.v1.Product/DelTerminal"
	Product_ListTerminal_FullMethodName     = "/api.product.v1.Product/ListTerminal"
	Product_PageListTerminal_FullMethodName = "/api.product.v1.Product/PageListTerminal"
	Product_GetTerminal_FullMethodName      = "/api.product.v1.Product/GetTerminal"
	Product_CreateKind_FullMethodName       = "/api.product.v1.Product/CreateKind"
	Product_EditKind_FullMethodName         = "/api.product.v1.Product/EditKind"
	Product_DelKind_FullMethodName          = "/api.product.v1.Product/DelKind"
	Product_ListKind_FullMethodName         = "/api.product.v1.Product/ListKind"
	Product_PageListKind_FullMethodName     = "/api.product.v1.Product/PageListKind"
	Product_GetKind_FullMethodName          = "/api.product.v1.Product/GetKind"
	Product_CreateProduct_FullMethodName    = "/api.product.v1.Product/CreateProduct"
	Product_EditProduct_FullMethodName      = "/api.product.v1.Product/EditProduct"
	Product_DelProduct_FullMethodName       = "/api.product.v1.Product/DelProduct"
	Product_ListProduct_FullMethodName      = "/api.product.v1.Product/ListProduct"
)

// ProductClient is the client API for Product service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductClient interface {
	// 创建终端
	CreateTerminal(ctx context.Context, in *CreateTerminalReq, opts ...grpc.CallOption) (*ProductStatus, error)
	// 编辑终端
	EditTerminal(ctx context.Context, in *EditTerminalReq, opts ...grpc.CallOption) (*ProductStatus, error)
	// 删除终端
	DelTerminal(ctx context.Context, in *DelIdReq, opts ...grpc.CallOption) (*ProductStatus, error)
	// 获取终端列表
	ListTerminal(ctx context.Context, in *Terminal, opts ...grpc.CallOption) (*ListTerminalRep, error)
	// 分页获取终端列表
	PageListTerminal(ctx context.Context, in *PageListTerminalReq, opts ...grpc.CallOption) (*PageListTerminalRep, error)
	// 查询单条终端信息
	GetTerminal(ctx context.Context, in *Terminal, opts ...grpc.CallOption) (*Terminal, error)
	// 创建产品类型
	CreateKind(ctx context.Context, in *CreateKindReq, opts ...grpc.CallOption) (*ProductStatus, error)
	// 编辑产品类型
	EditKind(ctx context.Context, in *EditKindReq, opts ...grpc.CallOption) (*ProductStatus, error)
	// 删除产品类型
	DelKind(ctx context.Context, in *DelIdReq, opts ...grpc.CallOption) (*ProductStatus, error)
	// 获取产品类型列表
	ListKind(ctx context.Context, in *Kind, opts ...grpc.CallOption) (*ListKindRep, error)
	// 分页获取产品类型列表
	PageListKind(ctx context.Context, in *PageListKindReq, opts ...grpc.CallOption) (*PageListKindRep, error)
	// 查询单条产品类型
	GetKind(ctx context.Context, in *Kind, opts ...grpc.CallOption) (*Kind, error)
	// 创建产品
	CreateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*ProductStatus, error)
	// 编辑产品
	EditProduct(ctx context.Context, in *EditProductReq, opts ...grpc.CallOption) (*ProductStatus, error)
	// 删除产品
	DelProduct(ctx context.Context, in *DelIdReq, opts ...grpc.CallOption) (*ProductStatus, error)
	// 获取产品类型列表
	ListProduct(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*ListProductRep, error)
}

type productClient struct {
	cc grpc.ClientConnInterface
}

func NewProductClient(cc grpc.ClientConnInterface) ProductClient {
	return &productClient{cc}
}

func (c *productClient) CreateTerminal(ctx context.Context, in *CreateTerminalReq, opts ...grpc.CallOption) (*ProductStatus, error) {
	out := new(ProductStatus)
	err := c.cc.Invoke(ctx, Product_CreateTerminal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) EditTerminal(ctx context.Context, in *EditTerminalReq, opts ...grpc.CallOption) (*ProductStatus, error) {
	out := new(ProductStatus)
	err := c.cc.Invoke(ctx, Product_EditTerminal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) DelTerminal(ctx context.Context, in *DelIdReq, opts ...grpc.CallOption) (*ProductStatus, error) {
	out := new(ProductStatus)
	err := c.cc.Invoke(ctx, Product_DelTerminal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) ListTerminal(ctx context.Context, in *Terminal, opts ...grpc.CallOption) (*ListTerminalRep, error) {
	out := new(ListTerminalRep)
	err := c.cc.Invoke(ctx, Product_ListTerminal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) PageListTerminal(ctx context.Context, in *PageListTerminalReq, opts ...grpc.CallOption) (*PageListTerminalRep, error) {
	out := new(PageListTerminalRep)
	err := c.cc.Invoke(ctx, Product_PageListTerminal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) GetTerminal(ctx context.Context, in *Terminal, opts ...grpc.CallOption) (*Terminal, error) {
	out := new(Terminal)
	err := c.cc.Invoke(ctx, Product_GetTerminal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) CreateKind(ctx context.Context, in *CreateKindReq, opts ...grpc.CallOption) (*ProductStatus, error) {
	out := new(ProductStatus)
	err := c.cc.Invoke(ctx, Product_CreateKind_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) EditKind(ctx context.Context, in *EditKindReq, opts ...grpc.CallOption) (*ProductStatus, error) {
	out := new(ProductStatus)
	err := c.cc.Invoke(ctx, Product_EditKind_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) DelKind(ctx context.Context, in *DelIdReq, opts ...grpc.CallOption) (*ProductStatus, error) {
	out := new(ProductStatus)
	err := c.cc.Invoke(ctx, Product_DelKind_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) ListKind(ctx context.Context, in *Kind, opts ...grpc.CallOption) (*ListKindRep, error) {
	out := new(ListKindRep)
	err := c.cc.Invoke(ctx, Product_ListKind_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) PageListKind(ctx context.Context, in *PageListKindReq, opts ...grpc.CallOption) (*PageListKindRep, error) {
	out := new(PageListKindRep)
	err := c.cc.Invoke(ctx, Product_PageListKind_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) GetKind(ctx context.Context, in *Kind, opts ...grpc.CallOption) (*Kind, error) {
	out := new(Kind)
	err := c.cc.Invoke(ctx, Product_GetKind_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) CreateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*ProductStatus, error) {
	out := new(ProductStatus)
	err := c.cc.Invoke(ctx, Product_CreateProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) EditProduct(ctx context.Context, in *EditProductReq, opts ...grpc.CallOption) (*ProductStatus, error) {
	out := new(ProductStatus)
	err := c.cc.Invoke(ctx, Product_EditProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) DelProduct(ctx context.Context, in *DelIdReq, opts ...grpc.CallOption) (*ProductStatus, error) {
	out := new(ProductStatus)
	err := c.cc.Invoke(ctx, Product_DelProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) ListProduct(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*ListProductRep, error) {
	out := new(ListProductRep)
	err := c.cc.Invoke(ctx, Product_ListProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServer is the server API for Product service.
// All implementations must embed UnimplementedProductServer
// for forward compatibility
type ProductServer interface {
	// 创建终端
	CreateTerminal(context.Context, *CreateTerminalReq) (*ProductStatus, error)
	// 编辑终端
	EditTerminal(context.Context, *EditTerminalReq) (*ProductStatus, error)
	// 删除终端
	DelTerminal(context.Context, *DelIdReq) (*ProductStatus, error)
	// 获取终端列表
	ListTerminal(context.Context, *Terminal) (*ListTerminalRep, error)
	// 分页获取终端列表
	PageListTerminal(context.Context, *PageListTerminalReq) (*PageListTerminalRep, error)
	// 查询单条终端信息
	GetTerminal(context.Context, *Terminal) (*Terminal, error)
	// 创建产品类型
	CreateKind(context.Context, *CreateKindReq) (*ProductStatus, error)
	// 编辑产品类型
	EditKind(context.Context, *EditKindReq) (*ProductStatus, error)
	// 删除产品类型
	DelKind(context.Context, *DelIdReq) (*ProductStatus, error)
	// 获取产品类型列表
	ListKind(context.Context, *Kind) (*ListKindRep, error)
	// 分页获取产品类型列表
	PageListKind(context.Context, *PageListKindReq) (*PageListKindRep, error)
	// 查询单条产品类型
	GetKind(context.Context, *Kind) (*Kind, error)
	// 创建产品
	CreateProduct(context.Context, *CreateProductReq) (*ProductStatus, error)
	// 编辑产品
	EditProduct(context.Context, *EditProductReq) (*ProductStatus, error)
	// 删除产品
	DelProduct(context.Context, *DelIdReq) (*ProductStatus, error)
	// 获取产品类型列表
	ListProduct(context.Context, *ProductInfo) (*ListProductRep, error)
	mustEmbedUnimplementedProductServer()
}

// UnimplementedProductServer must be embedded to have forward compatible implementations.
type UnimplementedProductServer struct {
}

func (UnimplementedProductServer) CreateTerminal(context.Context, *CreateTerminalReq) (*ProductStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTerminal not implemented")
}
func (UnimplementedProductServer) EditTerminal(context.Context, *EditTerminalReq) (*ProductStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditTerminal not implemented")
}
func (UnimplementedProductServer) DelTerminal(context.Context, *DelIdReq) (*ProductStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelTerminal not implemented")
}
func (UnimplementedProductServer) ListTerminal(context.Context, *Terminal) (*ListTerminalRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTerminal not implemented")
}
func (UnimplementedProductServer) PageListTerminal(context.Context, *PageListTerminalReq) (*PageListTerminalRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PageListTerminal not implemented")
}
func (UnimplementedProductServer) GetTerminal(context.Context, *Terminal) (*Terminal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTerminal not implemented")
}
func (UnimplementedProductServer) CreateKind(context.Context, *CreateKindReq) (*ProductStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKind not implemented")
}
func (UnimplementedProductServer) EditKind(context.Context, *EditKindReq) (*ProductStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditKind not implemented")
}
func (UnimplementedProductServer) DelKind(context.Context, *DelIdReq) (*ProductStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelKind not implemented")
}
func (UnimplementedProductServer) ListKind(context.Context, *Kind) (*ListKindRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKind not implemented")
}
func (UnimplementedProductServer) PageListKind(context.Context, *PageListKindReq) (*PageListKindRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PageListKind not implemented")
}
func (UnimplementedProductServer) GetKind(context.Context, *Kind) (*Kind, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKind not implemented")
}
func (UnimplementedProductServer) CreateProduct(context.Context, *CreateProductReq) (*ProductStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedProductServer) EditProduct(context.Context, *EditProductReq) (*ProductStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditProduct not implemented")
}
func (UnimplementedProductServer) DelProduct(context.Context, *DelIdReq) (*ProductStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelProduct not implemented")
}
func (UnimplementedProductServer) ListProduct(context.Context, *ProductInfo) (*ListProductRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProduct not implemented")
}
func (UnimplementedProductServer) mustEmbedUnimplementedProductServer() {}

// UnsafeProductServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServer will
// result in compilation errors.
type UnsafeProductServer interface {
	mustEmbedUnimplementedProductServer()
}

func RegisterProductServer(s grpc.ServiceRegistrar, srv ProductServer) {
	s.RegisterService(&Product_ServiceDesc, srv)
}

func _Product_CreateTerminal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTerminalReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).CreateTerminal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_CreateTerminal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).CreateTerminal(ctx, req.(*CreateTerminalReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_EditTerminal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditTerminalReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).EditTerminal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_EditTerminal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).EditTerminal(ctx, req.(*EditTerminalReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_DelTerminal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).DelTerminal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_DelTerminal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).DelTerminal(ctx, req.(*DelIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_ListTerminal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Terminal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).ListTerminal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_ListTerminal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).ListTerminal(ctx, req.(*Terminal))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_PageListTerminal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageListTerminalReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).PageListTerminal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_PageListTerminal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).PageListTerminal(ctx, req.(*PageListTerminalReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_GetTerminal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Terminal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetTerminal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_GetTerminal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetTerminal(ctx, req.(*Terminal))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_CreateKind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateKindReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).CreateKind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_CreateKind_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).CreateKind(ctx, req.(*CreateKindReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_EditKind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditKindReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).EditKind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_EditKind_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).EditKind(ctx, req.(*EditKindReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_DelKind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).DelKind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_DelKind_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).DelKind(ctx, req.(*DelIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_ListKind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Kind)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).ListKind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_ListKind_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).ListKind(ctx, req.(*Kind))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_PageListKind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageListKindReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).PageListKind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_PageListKind_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).PageListKind(ctx, req.(*PageListKindReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_GetKind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Kind)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetKind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_GetKind_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetKind(ctx, req.(*Kind))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_CreateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).CreateProduct(ctx, req.(*CreateProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_EditProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).EditProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_EditProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).EditProduct(ctx, req.(*EditProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_DelProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).DelProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_DelProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).DelProduct(ctx, req.(*DelIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_ListProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).ListProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_ListProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).ListProduct(ctx, req.(*ProductInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// Product_ServiceDesc is the grpc.ServiceDesc for Product service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Product_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.product.v1.Product",
	HandlerType: (*ProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTerminal",
			Handler:    _Product_CreateTerminal_Handler,
		},
		{
			MethodName: "EditTerminal",
			Handler:    _Product_EditTerminal_Handler,
		},
		{
			MethodName: "DelTerminal",
			Handler:    _Product_DelTerminal_Handler,
		},
		{
			MethodName: "ListTerminal",
			Handler:    _Product_ListTerminal_Handler,
		},
		{
			MethodName: "PageListTerminal",
			Handler:    _Product_PageListTerminal_Handler,
		},
		{
			MethodName: "GetTerminal",
			Handler:    _Product_GetTerminal_Handler,
		},
		{
			MethodName: "CreateKind",
			Handler:    _Product_CreateKind_Handler,
		},
		{
			MethodName: "EditKind",
			Handler:    _Product_EditKind_Handler,
		},
		{
			MethodName: "DelKind",
			Handler:    _Product_DelKind_Handler,
		},
		{
			MethodName: "ListKind",
			Handler:    _Product_ListKind_Handler,
		},
		{
			MethodName: "PageListKind",
			Handler:    _Product_PageListKind_Handler,
		},
		{
			MethodName: "GetKind",
			Handler:    _Product_GetKind_Handler,
		},
		{
			MethodName: "CreateProduct",
			Handler:    _Product_CreateProduct_Handler,
		},
		{
			MethodName: "EditProduct",
			Handler:    _Product_EditProduct_Handler,
		},
		{
			MethodName: "DelProduct",
			Handler:    _Product_DelProduct_Handler,
		},
		{
			MethodName: "ListProduct",
			Handler:    _Product_ListProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product/v1/product.proto",
}