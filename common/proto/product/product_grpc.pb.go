// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: product.proto

package product

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Product_ProductDetail_FullMethodName      = "/product.Product/ProductDetail"
	Product_EsAddProduct_FullMethodName       = "/product.Product/EsAddProduct"
	Product_GetAllProduct_FullMethodName      = "/product.Product/GetAllProduct"
	Product_EsSearchByKeyWord_FullMethodName  = "/product.Product/EsSearchByKeyWord"
	Product_PriceFind_FullMethodName          = "/product.Product/PriceFind"
	Product_GetCollectProduct_FullMethodName  = "/product.Product/GetCollectProduct"
	Product_ProductCategory_FullMethodName    = "/product.Product/ProductCategory"
	Product_WebsiteProductList_FullMethodName = "/product.Product/WebsiteProductList"
	Product_ProductSort_FullMethodName        = "/product.Product/ProductSort"
	Product_GroupByProductList_FullMethodName = "/product.Product/GroupByProductList"
	Product_ProductRanking_FullMethodName     = "/product.Product/ProductRanking"
	Product_ProductFilter_FullMethodName      = "/product.Product/ProductFilter"
	Product_AddSpikeProduct_FullMethodName    = "/product.Product/AddSpikeProduct"
)

// ProductClient is the client API for Product service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductClient interface {
	ProductDetail(ctx context.Context, in *ProductDetailRequest, opts ...grpc.CallOption) (*ProductDetailResponse, error)
	EsAddProduct(ctx context.Context, in *EsAddProductRequest, opts ...grpc.CallOption) (*EsAddProductResponse, error)
	GetAllProduct(ctx context.Context, in *GetAllProductRequest, opts ...grpc.CallOption) (*GetAllProductResponse, error)
	EsSearchByKeyWord(ctx context.Context, in *EsSearchByKeyWordRequest, opts ...grpc.CallOption) (*EsSearchByKeyWordResponse, error)
	PriceFind(ctx context.Context, in *PriceFindRequest, opts ...grpc.CallOption) (*PriceFindResponse, error)
	GetCollectProduct(ctx context.Context, in *GetCollectProductRequest, opts ...grpc.CallOption) (*GetCollectProductResponse, error)
	ProductCategory(ctx context.Context, in *ProductCategoryRequest, opts ...grpc.CallOption) (*ProductCategoryResponse, error)
	WebsiteProductList(ctx context.Context, in *WebsiteProductListRequest, opts ...grpc.CallOption) (*WebsiteProductListResponse, error)
	ProductSort(ctx context.Context, in *ProductSortRequest, opts ...grpc.CallOption) (*ProductSortResponse, error)
	GroupByProductList(ctx context.Context, in *GroupByProductListRequest, opts ...grpc.CallOption) (*GroupByProductListResponse, error)
	ProductRanking(ctx context.Context, in *ProductRankingRequest, opts ...grpc.CallOption) (*ProductRankingResponse, error)
	ProductFilter(ctx context.Context, in *ProductFilterRequest, opts ...grpc.CallOption) (*ProductFilterResponse, error)
	AddSpikeProduct(ctx context.Context, in *AddSpikeProductReq, opts ...grpc.CallOption) (*AddSpikeProductResp, error)
}

type productClient struct {
	cc grpc.ClientConnInterface
}

func NewProductClient(cc grpc.ClientConnInterface) ProductClient {
	return &productClient{cc}
}

func (c *productClient) ProductDetail(ctx context.Context, in *ProductDetailRequest, opts ...grpc.CallOption) (*ProductDetailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductDetailResponse)
	err := c.cc.Invoke(ctx, Product_ProductDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) EsAddProduct(ctx context.Context, in *EsAddProductRequest, opts ...grpc.CallOption) (*EsAddProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EsAddProductResponse)
	err := c.cc.Invoke(ctx, Product_EsAddProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) GetAllProduct(ctx context.Context, in *GetAllProductRequest, opts ...grpc.CallOption) (*GetAllProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllProductResponse)
	err := c.cc.Invoke(ctx, Product_GetAllProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) EsSearchByKeyWord(ctx context.Context, in *EsSearchByKeyWordRequest, opts ...grpc.CallOption) (*EsSearchByKeyWordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EsSearchByKeyWordResponse)
	err := c.cc.Invoke(ctx, Product_EsSearchByKeyWord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) PriceFind(ctx context.Context, in *PriceFindRequest, opts ...grpc.CallOption) (*PriceFindResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PriceFindResponse)
	err := c.cc.Invoke(ctx, Product_PriceFind_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) GetCollectProduct(ctx context.Context, in *GetCollectProductRequest, opts ...grpc.CallOption) (*GetCollectProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCollectProductResponse)
	err := c.cc.Invoke(ctx, Product_GetCollectProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) ProductCategory(ctx context.Context, in *ProductCategoryRequest, opts ...grpc.CallOption) (*ProductCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductCategoryResponse)
	err := c.cc.Invoke(ctx, Product_ProductCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) WebsiteProductList(ctx context.Context, in *WebsiteProductListRequest, opts ...grpc.CallOption) (*WebsiteProductListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WebsiteProductListResponse)
	err := c.cc.Invoke(ctx, Product_WebsiteProductList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) ProductSort(ctx context.Context, in *ProductSortRequest, opts ...grpc.CallOption) (*ProductSortResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductSortResponse)
	err := c.cc.Invoke(ctx, Product_ProductSort_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) GroupByProductList(ctx context.Context, in *GroupByProductListRequest, opts ...grpc.CallOption) (*GroupByProductListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GroupByProductListResponse)
	err := c.cc.Invoke(ctx, Product_GroupByProductList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) ProductRanking(ctx context.Context, in *ProductRankingRequest, opts ...grpc.CallOption) (*ProductRankingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductRankingResponse)
	err := c.cc.Invoke(ctx, Product_ProductRanking_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) ProductFilter(ctx context.Context, in *ProductFilterRequest, opts ...grpc.CallOption) (*ProductFilterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductFilterResponse)
	err := c.cc.Invoke(ctx, Product_ProductFilter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) AddSpikeProduct(ctx context.Context, in *AddSpikeProductReq, opts ...grpc.CallOption) (*AddSpikeProductResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddSpikeProductResp)
	err := c.cc.Invoke(ctx, Product_AddSpikeProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServer is the server API for Product service.
// All implementations must embed UnimplementedProductServer
// for forward compatibility.
type ProductServer interface {
	ProductDetail(context.Context, *ProductDetailRequest) (*ProductDetailResponse, error)
	EsAddProduct(context.Context, *EsAddProductRequest) (*EsAddProductResponse, error)
	GetAllProduct(context.Context, *GetAllProductRequest) (*GetAllProductResponse, error)
	EsSearchByKeyWord(context.Context, *EsSearchByKeyWordRequest) (*EsSearchByKeyWordResponse, error)
	PriceFind(context.Context, *PriceFindRequest) (*PriceFindResponse, error)
	GetCollectProduct(context.Context, *GetCollectProductRequest) (*GetCollectProductResponse, error)
	ProductCategory(context.Context, *ProductCategoryRequest) (*ProductCategoryResponse, error)
	WebsiteProductList(context.Context, *WebsiteProductListRequest) (*WebsiteProductListResponse, error)
	ProductSort(context.Context, *ProductSortRequest) (*ProductSortResponse, error)
	GroupByProductList(context.Context, *GroupByProductListRequest) (*GroupByProductListResponse, error)
	ProductRanking(context.Context, *ProductRankingRequest) (*ProductRankingResponse, error)
	ProductFilter(context.Context, *ProductFilterRequest) (*ProductFilterResponse, error)
	AddSpikeProduct(context.Context, *AddSpikeProductReq) (*AddSpikeProductResp, error)
	mustEmbedUnimplementedProductServer()
}

// UnimplementedProductServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProductServer struct{}

func (UnimplementedProductServer) ProductDetail(context.Context, *ProductDetailRequest) (*ProductDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductDetail not implemented")
}
func (UnimplementedProductServer) EsAddProduct(context.Context, *EsAddProductRequest) (*EsAddProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EsAddProduct not implemented")
}
func (UnimplementedProductServer) GetAllProduct(context.Context, *GetAllProductRequest) (*GetAllProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProduct not implemented")
}
func (UnimplementedProductServer) EsSearchByKeyWord(context.Context, *EsSearchByKeyWordRequest) (*EsSearchByKeyWordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EsSearchByKeyWord not implemented")
}
func (UnimplementedProductServer) PriceFind(context.Context, *PriceFindRequest) (*PriceFindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PriceFind not implemented")
}
func (UnimplementedProductServer) GetCollectProduct(context.Context, *GetCollectProductRequest) (*GetCollectProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollectProduct not implemented")
}
func (UnimplementedProductServer) ProductCategory(context.Context, *ProductCategoryRequest) (*ProductCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductCategory not implemented")
}
func (UnimplementedProductServer) WebsiteProductList(context.Context, *WebsiteProductListRequest) (*WebsiteProductListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WebsiteProductList not implemented")
}
func (UnimplementedProductServer) ProductSort(context.Context, *ProductSortRequest) (*ProductSortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductSort not implemented")
}
func (UnimplementedProductServer) GroupByProductList(context.Context, *GroupByProductListRequest) (*GroupByProductListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupByProductList not implemented")
}
func (UnimplementedProductServer) ProductRanking(context.Context, *ProductRankingRequest) (*ProductRankingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductRanking not implemented")
}
func (UnimplementedProductServer) ProductFilter(context.Context, *ProductFilterRequest) (*ProductFilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductFilter not implemented")
}
func (UnimplementedProductServer) AddSpikeProduct(context.Context, *AddSpikeProductReq) (*AddSpikeProductResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSpikeProduct not implemented")
}
func (UnimplementedProductServer) mustEmbedUnimplementedProductServer() {}
func (UnimplementedProductServer) testEmbeddedByValue()                 {}

// UnsafeProductServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServer will
// result in compilation errors.
type UnsafeProductServer interface {
	mustEmbedUnimplementedProductServer()
}

func RegisterProductServer(s grpc.ServiceRegistrar, srv ProductServer) {
	// If the following call pancis, it indicates UnimplementedProductServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Product_ServiceDesc, srv)
}

func _Product_ProductDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).ProductDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_ProductDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).ProductDetail(ctx, req.(*ProductDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_EsAddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EsAddProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).EsAddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_EsAddProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).EsAddProduct(ctx, req.(*EsAddProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_GetAllProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetAllProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_GetAllProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetAllProduct(ctx, req.(*GetAllProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_EsSearchByKeyWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EsSearchByKeyWordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).EsSearchByKeyWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_EsSearchByKeyWord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).EsSearchByKeyWord(ctx, req.(*EsSearchByKeyWordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_PriceFind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PriceFindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).PriceFind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_PriceFind_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).PriceFind(ctx, req.(*PriceFindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_GetCollectProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCollectProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetCollectProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_GetCollectProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetCollectProduct(ctx, req.(*GetCollectProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_ProductCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).ProductCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_ProductCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).ProductCategory(ctx, req.(*ProductCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_WebsiteProductList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebsiteProductListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).WebsiteProductList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_WebsiteProductList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).WebsiteProductList(ctx, req.(*WebsiteProductListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_ProductSort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductSortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).ProductSort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_ProductSort_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).ProductSort(ctx, req.(*ProductSortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_GroupByProductList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupByProductListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GroupByProductList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_GroupByProductList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GroupByProductList(ctx, req.(*GroupByProductListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_ProductRanking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRankingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).ProductRanking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_ProductRanking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).ProductRanking(ctx, req.(*ProductRankingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_ProductFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).ProductFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_ProductFilter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).ProductFilter(ctx, req.(*ProductFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_AddSpikeProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSpikeProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).AddSpikeProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_AddSpikeProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).AddSpikeProduct(ctx, req.(*AddSpikeProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Product_ServiceDesc is the grpc.ServiceDesc for Product service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Product_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.Product",
	HandlerType: (*ProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProductDetail",
			Handler:    _Product_ProductDetail_Handler,
		},
		{
			MethodName: "EsAddProduct",
			Handler:    _Product_EsAddProduct_Handler,
		},
		{
			MethodName: "GetAllProduct",
			Handler:    _Product_GetAllProduct_Handler,
		},
		{
			MethodName: "EsSearchByKeyWord",
			Handler:    _Product_EsSearchByKeyWord_Handler,
		},
		{
			MethodName: "PriceFind",
			Handler:    _Product_PriceFind_Handler,
		},
		{
			MethodName: "GetCollectProduct",
			Handler:    _Product_GetCollectProduct_Handler,
		},
		{
			MethodName: "ProductCategory",
			Handler:    _Product_ProductCategory_Handler,
		},
		{
			MethodName: "WebsiteProductList",
			Handler:    _Product_WebsiteProductList_Handler,
		},
		{
			MethodName: "ProductSort",
			Handler:    _Product_ProductSort_Handler,
		},
		{
			MethodName: "GroupByProductList",
			Handler:    _Product_GroupByProductList_Handler,
		},
		{
			MethodName: "ProductRanking",
			Handler:    _Product_ProductRanking_Handler,
		},
		{
			MethodName: "ProductFilter",
			Handler:    _Product_ProductFilter_Handler,
		},
		{
			MethodName: "AddSpikeProduct",
			Handler:    _Product_AddSpikeProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
