// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.19.4
// source: coupon.proto

package coupon

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Coupon_AddCoupon_FullMethodName       = "/coupon.Coupon/AddCoupon"
	Coupon_GrantCouponUser_FullMethodName = "/coupon.Coupon/GrantCouponUser"
)

// CouponClient is the client API for Coupon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CouponClient interface {
	AddCoupon(ctx context.Context, in *AddCouponRequest, opts ...grpc.CallOption) (*AddCouponResponse, error)
	GrantCouponUser(ctx context.Context, in *GrantCouponUserRequest, opts ...grpc.CallOption) (*GrantCouponUserResponse, error)
}

type couponClient struct {
	cc grpc.ClientConnInterface
}

func NewCouponClient(cc grpc.ClientConnInterface) CouponClient {
	return &couponClient{cc}
}

func (c *couponClient) AddCoupon(ctx context.Context, in *AddCouponRequest, opts ...grpc.CallOption) (*AddCouponResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddCouponResponse)
	err := c.cc.Invoke(ctx, Coupon_AddCoupon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponClient) GrantCouponUser(ctx context.Context, in *GrantCouponUserRequest, opts ...grpc.CallOption) (*GrantCouponUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GrantCouponUserResponse)
	err := c.cc.Invoke(ctx, Coupon_GrantCouponUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CouponServer is the server API for Coupon service.
// All implementations must embed UnimplementedCouponServer
// for forward compatibility
type CouponServer interface {
	AddCoupon(context.Context, *AddCouponRequest) (*AddCouponResponse, error)
	GrantCouponUser(context.Context, *GrantCouponUserRequest) (*GrantCouponUserResponse, error)
	mustEmbedUnimplementedCouponServer()
}

// UnimplementedCouponServer must be embedded to have forward compatible implementations.
type UnimplementedCouponServer struct {
}

func (UnimplementedCouponServer) AddCoupon(context.Context, *AddCouponRequest) (*AddCouponResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCoupon not implemented")
}
func (UnimplementedCouponServer) GrantCouponUser(context.Context, *GrantCouponUserRequest) (*GrantCouponUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrantCouponUser not implemented")
}
func (UnimplementedCouponServer) mustEmbedUnimplementedCouponServer() {}

// UnsafeCouponServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CouponServer will
// result in compilation errors.
type UnsafeCouponServer interface {
	mustEmbedUnimplementedCouponServer()
}

func RegisterCouponServer(s grpc.ServiceRegistrar, srv CouponServer) {
	s.RegisterService(&Coupon_ServiceDesc, srv)
}

func _Coupon_AddCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCouponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponServer).AddCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coupon_AddCoupon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponServer).AddCoupon(ctx, req.(*AddCouponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coupon_GrantCouponUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrantCouponUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponServer).GrantCouponUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coupon_GrantCouponUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponServer).GrantCouponUser(ctx, req.(*GrantCouponUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Coupon_ServiceDesc is the grpc.ServiceDesc for Coupon service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Coupon_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coupon.Coupon",
	HandlerType: (*CouponServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCoupon",
			Handler:    _Coupon_AddCoupon_Handler,
		},
		{
			MethodName: "GrantCouponUser",
			Handler:    _Coupon_GrantCouponUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coupon.proto",
}
