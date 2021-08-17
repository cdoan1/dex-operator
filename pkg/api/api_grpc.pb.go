// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// ApiDexClient is the client API for ApiDex service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiDexClient interface {
	// CreateClient creates a client.
	CreateClient(ctx context.Context, in *CreateClientReq, opts ...grpc.CallOption) (*CreateClientResp, error)
	// UpdateClient updates an existing client
	UpdateClient(ctx context.Context, in *UpdateClientReq, opts ...grpc.CallOption) (*UpdateClientResp, error)
	// DeleteClient deletes the provided client.
	DeleteClient(ctx context.Context, in *DeleteClientReq, opts ...grpc.CallOption) (*DeleteClientResp, error)
	// CreatePassword creates a password.
	CreatePassword(ctx context.Context, in *CreatePasswordReq, opts ...grpc.CallOption) (*CreatePasswordResp, error)
	// UpdatePassword modifies existing password.
	UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*UpdatePasswordResp, error)
	// DeletePassword deletes the password.
	DeletePassword(ctx context.Context, in *DeletePasswordReq, opts ...grpc.CallOption) (*DeletePasswordResp, error)
	// ListPassword lists all password entries.
	ListPasswords(ctx context.Context, in *ListPasswordReq, opts ...grpc.CallOption) (*ListPasswordResp, error)
	// GetVersion returns version information of the server.
	GetVersion(ctx context.Context, in *VersionReq, opts ...grpc.CallOption) (*VersionResp, error)
	// ListRefresh lists all the refresh token entries for a particular user.
	ListRefresh(ctx context.Context, in *ListRefreshReq, opts ...grpc.CallOption) (*ListRefreshResp, error)
	// RevokeRefresh revokes the refresh token for the provided user-client pair.
	//
	// Note that each user-client pair can have only one refresh token at a time.
	RevokeRefresh(ctx context.Context, in *RevokeRefreshReq, opts ...grpc.CallOption) (*RevokeRefreshResp, error)
	// VerifyPassword returns whether a password matches a hash for a specific email or not.
	VerifyPassword(ctx context.Context, in *VerifyPasswordReq, opts ...grpc.CallOption) (*VerifyPasswordResp, error)
}

type apiDexClient struct {
	cc grpc.ClientConnInterface
}

func NewApiDexClient(cc grpc.ClientConnInterface) ApiDexClient {
	return &apiDexClient{cc}
}

func (c *apiDexClient) CreateClient(ctx context.Context, in *CreateClientReq, opts ...grpc.CallOption) (*CreateClientResp, error) {
	out := new(CreateClientResp)
	err := c.cc.Invoke(ctx, "/api.ApiDex/CreateClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiDexClient) UpdateClient(ctx context.Context, in *UpdateClientReq, opts ...grpc.CallOption) (*UpdateClientResp, error) {
	out := new(UpdateClientResp)
	err := c.cc.Invoke(ctx, "/api.ApiDex/UpdateClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiDexClient) DeleteClient(ctx context.Context, in *DeleteClientReq, opts ...grpc.CallOption) (*DeleteClientResp, error) {
	out := new(DeleteClientResp)
	err := c.cc.Invoke(ctx, "/api.ApiDex/DeleteClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiDexClient) CreatePassword(ctx context.Context, in *CreatePasswordReq, opts ...grpc.CallOption) (*CreatePasswordResp, error) {
	out := new(CreatePasswordResp)
	err := c.cc.Invoke(ctx, "/api.ApiDex/CreatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiDexClient) UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*UpdatePasswordResp, error) {
	out := new(UpdatePasswordResp)
	err := c.cc.Invoke(ctx, "/api.ApiDex/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiDexClient) DeletePassword(ctx context.Context, in *DeletePasswordReq, opts ...grpc.CallOption) (*DeletePasswordResp, error) {
	out := new(DeletePasswordResp)
	err := c.cc.Invoke(ctx, "/api.ApiDex/DeletePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiDexClient) ListPasswords(ctx context.Context, in *ListPasswordReq, opts ...grpc.CallOption) (*ListPasswordResp, error) {
	out := new(ListPasswordResp)
	err := c.cc.Invoke(ctx, "/api.ApiDex/ListPasswords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiDexClient) GetVersion(ctx context.Context, in *VersionReq, opts ...grpc.CallOption) (*VersionResp, error) {
	out := new(VersionResp)
	err := c.cc.Invoke(ctx, "/api.ApiDex/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiDexClient) ListRefresh(ctx context.Context, in *ListRefreshReq, opts ...grpc.CallOption) (*ListRefreshResp, error) {
	out := new(ListRefreshResp)
	err := c.cc.Invoke(ctx, "/api.ApiDex/ListRefresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiDexClient) RevokeRefresh(ctx context.Context, in *RevokeRefreshReq, opts ...grpc.CallOption) (*RevokeRefreshResp, error) {
	out := new(RevokeRefreshResp)
	err := c.cc.Invoke(ctx, "/api.ApiDex/RevokeRefresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiDexClient) VerifyPassword(ctx context.Context, in *VerifyPasswordReq, opts ...grpc.CallOption) (*VerifyPasswordResp, error) {
	out := new(VerifyPasswordResp)
	err := c.cc.Invoke(ctx, "/api.ApiDex/VerifyPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiDexServer is the server API for ApiDex service.
// All implementations must embed UnimplementedApiDexServer
// for forward compatibility
type ApiDexServer interface {
	// CreateClient creates a client.
	CreateClient(context.Context, *CreateClientReq) (*CreateClientResp, error)
	// UpdateClient updates an existing client
	UpdateClient(context.Context, *UpdateClientReq) (*UpdateClientResp, error)
	// DeleteClient deletes the provided client.
	DeleteClient(context.Context, *DeleteClientReq) (*DeleteClientResp, error)
	// CreatePassword creates a password.
	CreatePassword(context.Context, *CreatePasswordReq) (*CreatePasswordResp, error)
	// UpdatePassword modifies existing password.
	UpdatePassword(context.Context, *UpdatePasswordReq) (*UpdatePasswordResp, error)
	// DeletePassword deletes the password.
	DeletePassword(context.Context, *DeletePasswordReq) (*DeletePasswordResp, error)
	// ListPassword lists all password entries.
	ListPasswords(context.Context, *ListPasswordReq) (*ListPasswordResp, error)
	// GetVersion returns version information of the server.
	GetVersion(context.Context, *VersionReq) (*VersionResp, error)
	// ListRefresh lists all the refresh token entries for a particular user.
	ListRefresh(context.Context, *ListRefreshReq) (*ListRefreshResp, error)
	// RevokeRefresh revokes the refresh token for the provided user-client pair.
	//
	// Note that each user-client pair can have only one refresh token at a time.
	RevokeRefresh(context.Context, *RevokeRefreshReq) (*RevokeRefreshResp, error)
	// VerifyPassword returns whether a password matches a hash for a specific email or not.
	VerifyPassword(context.Context, *VerifyPasswordReq) (*VerifyPasswordResp, error)
	mustEmbedUnimplementedApiDexServer()
}

// UnimplementedApiDexServer must be embedded to have forward compatible implementations.
type UnimplementedApiDexServer struct {
}

func (UnimplementedApiDexServer) CreateClient(context.Context, *CreateClientReq) (*CreateClientResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClient not implemented")
}
func (UnimplementedApiDexServer) UpdateClient(context.Context, *UpdateClientReq) (*UpdateClientResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClient not implemented")
}
func (UnimplementedApiDexServer) DeleteClient(context.Context, *DeleteClientReq) (*DeleteClientResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClient not implemented")
}
func (UnimplementedApiDexServer) CreatePassword(context.Context, *CreatePasswordReq) (*CreatePasswordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePassword not implemented")
}
func (UnimplementedApiDexServer) UpdatePassword(context.Context, *UpdatePasswordReq) (*UpdatePasswordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedApiDexServer) DeletePassword(context.Context, *DeletePasswordReq) (*DeletePasswordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePassword not implemented")
}
func (UnimplementedApiDexServer) ListPasswords(context.Context, *ListPasswordReq) (*ListPasswordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPasswords not implemented")
}
func (UnimplementedApiDexServer) GetVersion(context.Context, *VersionReq) (*VersionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedApiDexServer) ListRefresh(context.Context, *ListRefreshReq) (*ListRefreshResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRefresh not implemented")
}
func (UnimplementedApiDexServer) RevokeRefresh(context.Context, *RevokeRefreshReq) (*RevokeRefreshResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeRefresh not implemented")
}
func (UnimplementedApiDexServer) VerifyPassword(context.Context, *VerifyPasswordReq) (*VerifyPasswordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyPassword not implemented")
}
func (UnimplementedApiDexServer) mustEmbedUnimplementedApiDexServer() {}

// UnsafeApiDexServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiDexServer will
// result in compilation errors.
type UnsafeApiDexServer interface {
	mustEmbedUnimplementedApiDexServer()
}

func RegisterApiDexServer(s grpc.ServiceRegistrar, srv ApiDexServer) {
	s.RegisterService(&ApiDex_ServiceDesc, srv)
}

func _ApiDex_CreateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiDexServer).CreateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiDex/CreateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiDexServer).CreateClient(ctx, req.(*CreateClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiDex_UpdateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiDexServer).UpdateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiDex/UpdateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiDexServer).UpdateClient(ctx, req.(*UpdateClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiDex_DeleteClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiDexServer).DeleteClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiDex/DeleteClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiDexServer).DeleteClient(ctx, req.(*DeleteClientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiDex_CreatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiDexServer).CreatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiDex/CreatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiDexServer).CreatePassword(ctx, req.(*CreatePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiDex_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiDexServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiDex/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiDexServer).UpdatePassword(ctx, req.(*UpdatePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiDex_DeletePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiDexServer).DeletePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiDex/DeletePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiDexServer).DeletePassword(ctx, req.(*DeletePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiDex_ListPasswords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiDexServer).ListPasswords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiDex/ListPasswords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiDexServer).ListPasswords(ctx, req.(*ListPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiDex_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiDexServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiDex/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiDexServer).GetVersion(ctx, req.(*VersionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiDex_ListRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRefreshReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiDexServer).ListRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiDex/ListRefresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiDexServer).ListRefresh(ctx, req.(*ListRefreshReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiDex_RevokeRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeRefreshReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiDexServer).RevokeRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiDex/RevokeRefresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiDexServer).RevokeRefresh(ctx, req.(*RevokeRefreshReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiDex_VerifyPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiDexServer).VerifyPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiDex/VerifyPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiDexServer).VerifyPassword(ctx, req.(*VerifyPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiDex_ServiceDesc is the grpc.ServiceDesc for ApiDex service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiDex_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ApiDex",
	HandlerType: (*ApiDexServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClient",
			Handler:    _ApiDex_CreateClient_Handler,
		},
		{
			MethodName: "UpdateClient",
			Handler:    _ApiDex_UpdateClient_Handler,
		},
		{
			MethodName: "DeleteClient",
			Handler:    _ApiDex_DeleteClient_Handler,
		},
		{
			MethodName: "CreatePassword",
			Handler:    _ApiDex_CreatePassword_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _ApiDex_UpdatePassword_Handler,
		},
		{
			MethodName: "DeletePassword",
			Handler:    _ApiDex_DeletePassword_Handler,
		},
		{
			MethodName: "ListPasswords",
			Handler:    _ApiDex_ListPasswords_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _ApiDex_GetVersion_Handler,
		},
		{
			MethodName: "ListRefresh",
			Handler:    _ApiDex_ListRefresh_Handler,
		},
		{
			MethodName: "RevokeRefresh",
			Handler:    _ApiDex_RevokeRefresh_Handler,
		},
		{
			MethodName: "VerifyPassword",
			Handler:    _ApiDex_VerifyPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/api/api.proto",
}