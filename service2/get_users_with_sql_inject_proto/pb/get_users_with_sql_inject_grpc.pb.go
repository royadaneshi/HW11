// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// GetUsersClient is the client API for GetUsers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetUsersClient interface {
	GetData(ctx context.Context, in *GetDataRequest, opts ...grpc.CallOption) (*GetDataResponse, error)
}

type getUsersClient struct {
	cc grpc.ClientConnInterface
}

func NewGetUsersClient(cc grpc.ClientConnInterface) GetUsersClient {
	return &getUsersClient{cc}
}

func (c *getUsersClient) GetData(ctx context.Context, in *GetDataRequest, opts ...grpc.CallOption) (*GetDataResponse, error) {
	out := new(GetDataResponse)
	err := c.cc.Invoke(ctx, "/get_users/GetData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetUsersServer is the server API for GetUsers service.
// All implementations must embed UnimplementedGetUsersServer
// for forward compatibility
type GetUsersServer interface {
	GetData(context.Context, *GetDataRequest) (*GetDataResponse, error)
	mustEmbedUnimplementedGetUsersServer()
}

// UnimplementedGetUsersServer must be embedded to have forward compatible implementations.
type UnimplementedGetUsersServer struct {
}

func (UnimplementedGetUsersServer) GetData(context.Context, *GetDataRequest) (*GetDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}
func (UnimplementedGetUsersServer) mustEmbedUnimplementedGetUsersServer() {}

// UnsafeGetUsersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetUsersServer will
// result in compilation errors.
type UnsafeGetUsersServer interface {
	mustEmbedUnimplementedGetUsersServer()
}

func RegisterGetUsersServer(s grpc.ServiceRegistrar, srv GetUsersServer) {
	s.RegisterService(&GetUsers_ServiceDesc, srv)
}

func _GetUsers_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetUsersServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/get_users/GetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetUsersServer).GetData(ctx, req.(*GetDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetUsers_ServiceDesc is the grpc.ServiceDesc for GetUsers service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetUsers_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "get_users",
	HandlerType: (*GetUsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetData",
			Handler:    _GetUsers_GetData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "get_users_with_sql_inject.proto",
}
