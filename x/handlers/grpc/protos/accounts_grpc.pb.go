// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: protos/accounts.proto

package protos

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
	Accounts_GetPersonDetails_FullMethodName    = "/accounts.v1.accounts.Accounts/GetPersonDetails"
	Accounts_GetFoodPlaceDetails_FullMethodName = "/accounts.v1.accounts.Accounts/GetFoodPlaceDetails"
	Accounts_GetUserId_FullMethodName           = "/accounts.v1.accounts.Accounts/GetUserId"
	Accounts_AddFoodPlaceAccount_FullMethodName = "/accounts.v1.accounts.Accounts/AddFoodPlaceAccount"
	Accounts_AddPersonAccount_FullMethodName    = "/accounts.v1.accounts.Accounts/AddPersonAccount"
)

// AccountsClient is the client API for Accounts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountsClient interface {
	GetPersonDetails(ctx context.Context, in *GetAccountDetailsRequest, opts ...grpc.CallOption) (*GetPersonDetailsResponse, error)
	GetFoodPlaceDetails(ctx context.Context, in *GetAccountDetailsRequest, opts ...grpc.CallOption) (*GetFoodPlaceDetailsResponse, error)
	GetUserId(ctx context.Context, in *GetUserIdRequest, opts ...grpc.CallOption) (*GetUserIdResponse, error)
	AddFoodPlaceAccount(ctx context.Context, in *AddFoodPlaceAccountRequest, opts ...grpc.CallOption) (*AddFoodPlaceAccountResponse, error)
	AddPersonAccount(ctx context.Context, in *AddPersonAccountRequest, opts ...grpc.CallOption) (*AddPersonAccountResponse, error)
}

type accountsClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountsClient(cc grpc.ClientConnInterface) AccountsClient {
	return &accountsClient{cc}
}

func (c *accountsClient) GetPersonDetails(ctx context.Context, in *GetAccountDetailsRequest, opts ...grpc.CallOption) (*GetPersonDetailsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPersonDetailsResponse)
	err := c.cc.Invoke(ctx, Accounts_GetPersonDetails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) GetFoodPlaceDetails(ctx context.Context, in *GetAccountDetailsRequest, opts ...grpc.CallOption) (*GetFoodPlaceDetailsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFoodPlaceDetailsResponse)
	err := c.cc.Invoke(ctx, Accounts_GetFoodPlaceDetails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) GetUserId(ctx context.Context, in *GetUserIdRequest, opts ...grpc.CallOption) (*GetUserIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserIdResponse)
	err := c.cc.Invoke(ctx, Accounts_GetUserId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) AddFoodPlaceAccount(ctx context.Context, in *AddFoodPlaceAccountRequest, opts ...grpc.CallOption) (*AddFoodPlaceAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddFoodPlaceAccountResponse)
	err := c.cc.Invoke(ctx, Accounts_AddFoodPlaceAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) AddPersonAccount(ctx context.Context, in *AddPersonAccountRequest, opts ...grpc.CallOption) (*AddPersonAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddPersonAccountResponse)
	err := c.cc.Invoke(ctx, Accounts_AddPersonAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountsServer is the server API for Accounts service.
// All implementations must embed UnimplementedAccountsServer
// for forward compatibility
type AccountsServer interface {
	GetPersonDetails(context.Context, *GetAccountDetailsRequest) (*GetPersonDetailsResponse, error)
	GetFoodPlaceDetails(context.Context, *GetAccountDetailsRequest) (*GetFoodPlaceDetailsResponse, error)
	GetUserId(context.Context, *GetUserIdRequest) (*GetUserIdResponse, error)
	AddFoodPlaceAccount(context.Context, *AddFoodPlaceAccountRequest) (*AddFoodPlaceAccountResponse, error)
	AddPersonAccount(context.Context, *AddPersonAccountRequest) (*AddPersonAccountResponse, error)
	mustEmbedUnimplementedAccountsServer()
}

// UnimplementedAccountsServer must be embedded to have forward compatible implementations.
type UnimplementedAccountsServer struct {
}

func (UnimplementedAccountsServer) GetPersonDetails(context.Context, *GetAccountDetailsRequest) (*GetPersonDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersonDetails not implemented")
}
func (UnimplementedAccountsServer) GetFoodPlaceDetails(context.Context, *GetAccountDetailsRequest) (*GetFoodPlaceDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFoodPlaceDetails not implemented")
}
func (UnimplementedAccountsServer) GetUserId(context.Context, *GetUserIdRequest) (*GetUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserId not implemented")
}
func (UnimplementedAccountsServer) AddFoodPlaceAccount(context.Context, *AddFoodPlaceAccountRequest) (*AddFoodPlaceAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFoodPlaceAccount not implemented")
}
func (UnimplementedAccountsServer) AddPersonAccount(context.Context, *AddPersonAccountRequest) (*AddPersonAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPersonAccount not implemented")
}
func (UnimplementedAccountsServer) mustEmbedUnimplementedAccountsServer() {}

// UnsafeAccountsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountsServer will
// result in compilation errors.
type UnsafeAccountsServer interface {
	mustEmbedUnimplementedAccountsServer()
}

func RegisterAccountsServer(s grpc.ServiceRegistrar, srv AccountsServer) {
	s.RegisterService(&Accounts_ServiceDesc, srv)
}

func _Accounts_GetPersonDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).GetPersonDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Accounts_GetPersonDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).GetPersonDetails(ctx, req.(*GetAccountDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_GetFoodPlaceDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).GetFoodPlaceDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Accounts_GetFoodPlaceDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).GetFoodPlaceDetails(ctx, req.(*GetAccountDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_GetUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).GetUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Accounts_GetUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).GetUserId(ctx, req.(*GetUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_AddFoodPlaceAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFoodPlaceAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).AddFoodPlaceAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Accounts_AddFoodPlaceAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).AddFoodPlaceAccount(ctx, req.(*AddFoodPlaceAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_AddPersonAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPersonAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).AddPersonAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Accounts_AddPersonAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).AddPersonAccount(ctx, req.(*AddPersonAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Accounts_ServiceDesc is the grpc.ServiceDesc for Accounts service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Accounts_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accounts.v1.accounts.Accounts",
	HandlerType: (*AccountsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPersonDetails",
			Handler:    _Accounts_GetPersonDetails_Handler,
		},
		{
			MethodName: "GetFoodPlaceDetails",
			Handler:    _Accounts_GetFoodPlaceDetails_Handler,
		},
		{
			MethodName: "GetUserId",
			Handler:    _Accounts_GetUserId_Handler,
		},
		{
			MethodName: "AddFoodPlaceAccount",
			Handler:    _Accounts_AddFoodPlaceAccount_Handler,
		},
		{
			MethodName: "AddPersonAccount",
			Handler:    _Accounts_AddPersonAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/accounts.proto",
}
