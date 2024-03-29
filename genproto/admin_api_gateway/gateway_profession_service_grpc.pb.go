// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: gateway_profession_service.proto

package admin_api_gateway

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

// ProfessionApiGatewayClient is the client API for ProfessionApiGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfessionApiGatewayClient interface {
	Create(ctx context.Context, in *GatewayCreateProfessionReq, opts ...grpc.CallOption) (*GatewayProfessionId, error)
	Get(ctx context.Context, in *GatewayProfessionId, opts ...grpc.CallOption) (*GatewayProfession, error)
	GetAll(ctx context.Context, in *GatewayGetAllProfessionRequest, opts ...grpc.CallOption) (*GatewayGetAllProfessionResponse, error)
}

type professionApiGatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewProfessionApiGatewayClient(cc grpc.ClientConnInterface) ProfessionApiGatewayClient {
	return &professionApiGatewayClient{cc}
}

func (c *professionApiGatewayClient) Create(ctx context.Context, in *GatewayCreateProfessionReq, opts ...grpc.CallOption) (*GatewayProfessionId, error) {
	out := new(GatewayProfessionId)
	err := c.cc.Invoke(ctx, "/genproto.ProfessionApiGateway/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *professionApiGatewayClient) Get(ctx context.Context, in *GatewayProfessionId, opts ...grpc.CallOption) (*GatewayProfession, error) {
	out := new(GatewayProfession)
	err := c.cc.Invoke(ctx, "/genproto.ProfessionApiGateway/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *professionApiGatewayClient) GetAll(ctx context.Context, in *GatewayGetAllProfessionRequest, opts ...grpc.CallOption) (*GatewayGetAllProfessionResponse, error) {
	out := new(GatewayGetAllProfessionResponse)
	err := c.cc.Invoke(ctx, "/genproto.ProfessionApiGateway/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfessionApiGatewayServer is the server API for ProfessionApiGateway service.
// All implementations must embed UnimplementedProfessionApiGatewayServer
// for forward compatibility
type ProfessionApiGatewayServer interface {
	Create(context.Context, *GatewayCreateProfessionReq) (*GatewayProfessionId, error)
	Get(context.Context, *GatewayProfessionId) (*GatewayProfession, error)
	GetAll(context.Context, *GatewayGetAllProfessionRequest) (*GatewayGetAllProfessionResponse, error)
	mustEmbedUnimplementedProfessionApiGatewayServer()
}

// UnimplementedProfessionApiGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedProfessionApiGatewayServer struct {
}

func (UnimplementedProfessionApiGatewayServer) Create(context.Context, *GatewayCreateProfessionReq) (*GatewayProfessionId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedProfessionApiGatewayServer) Get(context.Context, *GatewayProfessionId) (*GatewayProfession, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedProfessionApiGatewayServer) GetAll(context.Context, *GatewayGetAllProfessionRequest) (*GatewayGetAllProfessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedProfessionApiGatewayServer) mustEmbedUnimplementedProfessionApiGatewayServer() {}

// UnsafeProfessionApiGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfessionApiGatewayServer will
// result in compilation errors.
type UnsafeProfessionApiGatewayServer interface {
	mustEmbedUnimplementedProfessionApiGatewayServer()
}

func RegisterProfessionApiGatewayServer(s grpc.ServiceRegistrar, srv ProfessionApiGatewayServer) {
	s.RegisterService(&ProfessionApiGateway_ServiceDesc, srv)
}

func _ProfessionApiGateway_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayCreateProfessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfessionApiGatewayServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ProfessionApiGateway/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfessionApiGatewayServer).Create(ctx, req.(*GatewayCreateProfessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfessionApiGateway_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayProfessionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfessionApiGatewayServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ProfessionApiGateway/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfessionApiGatewayServer).Get(ctx, req.(*GatewayProfessionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfessionApiGateway_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayGetAllProfessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfessionApiGatewayServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ProfessionApiGateway/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfessionApiGatewayServer).GetAll(ctx, req.(*GatewayGetAllProfessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfessionApiGateway_ServiceDesc is the grpc.ServiceDesc for ProfessionApiGateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfessionApiGateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.ProfessionApiGateway",
	HandlerType: (*ProfessionApiGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ProfessionApiGateway_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ProfessionApiGateway_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ProfessionApiGateway_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway_profession_service.proto",
}
