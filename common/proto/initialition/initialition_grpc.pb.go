// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: initialition.proto

package initialition

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
	Integrate_Integration_FullMethodName = "/template.Integrate/Integration"
)

// IntegrateClient is the client API for Integrate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IntegrateClient interface {
	Integration(ctx context.Context, in *IntegrationRequest, opts ...grpc.CallOption) (*IntegrationRequest, error)
}

type integrateClient struct {
	cc grpc.ClientConnInterface
}

func NewIntegrateClient(cc grpc.ClientConnInterface) IntegrateClient {
	return &integrateClient{cc}
}

func (c *integrateClient) Integration(ctx context.Context, in *IntegrationRequest, opts ...grpc.CallOption) (*IntegrationRequest, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IntegrationRequest)
	err := c.cc.Invoke(ctx, Integrate_Integration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IntegrateServer is the server API for Integrate service.
// All implementations must embed UnimplementedIntegrateServer
// for forward compatibility.
type IntegrateServer interface {
	Integration(context.Context, *IntegrationRequest) (*IntegrationRequest, error)
	mustEmbedUnimplementedIntegrateServer()
}

// UnimplementedIntegrateServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedIntegrateServer struct{}

func (UnimplementedIntegrateServer) Integration(context.Context, *IntegrationRequest) (*IntegrationRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Integration not implemented")
}
func (UnimplementedIntegrateServer) mustEmbedUnimplementedIntegrateServer() {}
func (UnimplementedIntegrateServer) testEmbeddedByValue()                   {}

// UnsafeIntegrateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IntegrateServer will
// result in compilation errors.
type UnsafeIntegrateServer interface {
	mustEmbedUnimplementedIntegrateServer()
}

func RegisterIntegrateServer(s grpc.ServiceRegistrar, srv IntegrateServer) {
	// If the following call pancis, it indicates UnimplementedIntegrateServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Integrate_ServiceDesc, srv)
}

func _Integrate_Integration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntegrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrateServer).Integration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Integrate_Integration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrateServer).Integration(ctx, req.(*IntegrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Integrate_ServiceDesc is the grpc.ServiceDesc for Integrate service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Integrate_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "template.Integrate",
	HandlerType: (*IntegrateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Integration",
			Handler:    _Integrate_Integration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "initialition.proto",
}
