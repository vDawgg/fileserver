// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: unified.proto

package fileserver

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

// RetrieverClient is the client API for Retriever service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RetrieverClient interface {
	//Should also be included in the frontend
	SaveFiles(ctx context.Context, opts ...grpc.CallOption) (Retriever_SaveFilesClient, error)
	GetStructure(ctx context.Context, in *StructureRequest, opts ...grpc.CallOption) (*Structure, error)
	GetFiles(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (Retriever_GetFilesClient, error)
}

type retrieverClient struct {
	cc grpc.ClientConnInterface
}

func NewRetrieverClient(cc grpc.ClientConnInterface) RetrieverClient {
	return &retrieverClient{cc}
}

func (c *retrieverClient) SaveFiles(ctx context.Context, opts ...grpc.CallOption) (Retriever_SaveFilesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Retriever_ServiceDesc.Streams[0], "/fileserver.Retriever/saveFiles", opts...)
	if err != nil {
		return nil, err
	}
	x := &retrieverSaveFilesClient{stream}
	return x, nil
}

type Retriever_SaveFilesClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*UploadStatus, error)
	grpc.ClientStream
}

type retrieverSaveFilesClient struct {
	grpc.ClientStream
}

func (x *retrieverSaveFilesClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *retrieverSaveFilesClient) CloseAndRecv() (*UploadStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *retrieverClient) GetStructure(ctx context.Context, in *StructureRequest, opts ...grpc.CallOption) (*Structure, error) {
	out := new(Structure)
	err := c.cc.Invoke(ctx, "/fileserver.Retriever/getStructure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *retrieverClient) GetFiles(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (Retriever_GetFilesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Retriever_ServiceDesc.Streams[1], "/fileserver.Retriever/getFiles", opts...)
	if err != nil {
		return nil, err
	}
	x := &retrieverGetFilesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Retriever_GetFilesClient interface {
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type retrieverGetFilesClient struct {
	grpc.ClientStream
}

func (x *retrieverGetFilesClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RetrieverServer is the server API for Retriever service.
// All implementations must embed UnimplementedRetrieverServer
// for forward compatibility
type RetrieverServer interface {
	//Should also be included in the frontend
	SaveFiles(Retriever_SaveFilesServer) error
	GetStructure(context.Context, *StructureRequest) (*Structure, error)
	GetFiles(*DownloadRequest, Retriever_GetFilesServer) error
	mustEmbedUnimplementedRetrieverServer()
}

// UnimplementedRetrieverServer must be embedded to have forward compatible implementations.
type UnimplementedRetrieverServer struct {
}

func (UnimplementedRetrieverServer) SaveFiles(Retriever_SaveFilesServer) error {
	return status.Errorf(codes.Unimplemented, "method SaveFiles not implemented")
}
func (UnimplementedRetrieverServer) GetStructure(context.Context, *StructureRequest) (*Structure, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStructure not implemented")
}
func (UnimplementedRetrieverServer) GetFiles(*DownloadRequest, Retriever_GetFilesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFiles not implemented")
}
func (UnimplementedRetrieverServer) mustEmbedUnimplementedRetrieverServer() {}

// UnsafeRetrieverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RetrieverServer will
// result in compilation errors.
type UnsafeRetrieverServer interface {
	mustEmbedUnimplementedRetrieverServer()
}

func RegisterRetrieverServer(s grpc.ServiceRegistrar, srv RetrieverServer) {
	s.RegisterService(&Retriever_ServiceDesc, srv)
}

func _Retriever_SaveFiles_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RetrieverServer).SaveFiles(&retrieverSaveFilesServer{stream})
}

type Retriever_SaveFilesServer interface {
	SendAndClose(*UploadStatus) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type retrieverSaveFilesServer struct {
	grpc.ServerStream
}

func (x *retrieverSaveFilesServer) SendAndClose(m *UploadStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *retrieverSaveFilesServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Retriever_GetStructure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StructureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrieverServer).GetStructure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fileserver.Retriever/getStructure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrieverServer).GetStructure(ctx, req.(*StructureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Retriever_GetFiles_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RetrieverServer).GetFiles(m, &retrieverGetFilesServer{stream})
}

type Retriever_GetFilesServer interface {
	Send(*Chunk) error
	grpc.ServerStream
}

type retrieverGetFilesServer struct {
	grpc.ServerStream
}

func (x *retrieverGetFilesServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

// Retriever_ServiceDesc is the grpc.ServiceDesc for Retriever service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Retriever_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fileserver.Retriever",
	HandlerType: (*RetrieverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getStructure",
			Handler:    _Retriever_GetStructure_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "saveFiles",
			Handler:       _Retriever_SaveFiles_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "getFiles",
			Handler:       _Retriever_GetFiles_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "unified.proto",
}

// AuthenticatorClient is the client API for Authenticator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticatorClient interface {
	//Rpc for getting public keys of the backend
	GetKeys(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Keys, error)
	//Rpc to log in a user
	Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*Token, error)
}

type authenticatorClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticatorClient(cc grpc.ClientConnInterface) AuthenticatorClient {
	return &authenticatorClient{cc}
}

func (c *authenticatorClient) GetKeys(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Keys, error) {
	out := new(Keys)
	err := c.cc.Invoke(ctx, "/fileserver.authenticator/getKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticatorClient) Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/fileserver.authenticator/login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticatorServer is the server API for Authenticator service.
// All implementations must embed UnimplementedAuthenticatorServer
// for forward compatibility
type AuthenticatorServer interface {
	//Rpc for getting public keys of the backend
	GetKeys(context.Context, *Empty) (*Keys, error)
	//Rpc to log in a user
	Login(context.Context, *User) (*Token, error)
	mustEmbedUnimplementedAuthenticatorServer()
}

// UnimplementedAuthenticatorServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticatorServer struct {
}

func (UnimplementedAuthenticatorServer) GetKeys(context.Context, *Empty) (*Keys, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeys not implemented")
}
func (UnimplementedAuthenticatorServer) Login(context.Context, *User) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthenticatorServer) mustEmbedUnimplementedAuthenticatorServer() {}

// UnsafeAuthenticatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticatorServer will
// result in compilation errors.
type UnsafeAuthenticatorServer interface {
	mustEmbedUnimplementedAuthenticatorServer()
}

func RegisterAuthenticatorServer(s grpc.ServiceRegistrar, srv AuthenticatorServer) {
	s.RegisterService(&Authenticator_ServiceDesc, srv)
}

func _Authenticator_GetKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticatorServer).GetKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fileserver.authenticator/getKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticatorServer).GetKeys(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authenticator_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticatorServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fileserver.authenticator/login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticatorServer).Login(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// Authenticator_ServiceDesc is the grpc.ServiceDesc for Authenticator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authenticator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fileserver.authenticator",
	HandlerType: (*AuthenticatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getKeys",
			Handler:    _Authenticator_GetKeys_Handler,
		},
		{
			MethodName: "login",
			Handler:    _Authenticator_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "unified.proto",
}

// AuthorizerClient is the client API for Authorizer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorizerClient interface {
	IsAuthorized(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthReply, error)
	AddAuthorization(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*Added, error)
}

type authorizerClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorizerClient(cc grpc.ClientConnInterface) AuthorizerClient {
	return &authorizerClient{cc}
}

func (c *authorizerClient) IsAuthorized(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthReply, error) {
	out := new(AuthReply)
	err := c.cc.Invoke(ctx, "/fileserver.authorizer/isAuthorized", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizerClient) AddAuthorization(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*Added, error) {
	out := new(Added)
	err := c.cc.Invoke(ctx, "/fileserver.authorizer/addAuthorization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizerServer is the server API for Authorizer service.
// All implementations must embed UnimplementedAuthorizerServer
// for forward compatibility
type AuthorizerServer interface {
	IsAuthorized(context.Context, *AuthRequest) (*AuthReply, error)
	AddAuthorization(context.Context, *AuthRequest) (*Added, error)
	mustEmbedUnimplementedAuthorizerServer()
}

// UnimplementedAuthorizerServer must be embedded to have forward compatible implementations.
type UnimplementedAuthorizerServer struct {
}

func (UnimplementedAuthorizerServer) IsAuthorized(context.Context, *AuthRequest) (*AuthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAuthorized not implemented")
}
func (UnimplementedAuthorizerServer) AddAuthorization(context.Context, *AuthRequest) (*Added, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAuthorization not implemented")
}
func (UnimplementedAuthorizerServer) mustEmbedUnimplementedAuthorizerServer() {}

// UnsafeAuthorizerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorizerServer will
// result in compilation errors.
type UnsafeAuthorizerServer interface {
	mustEmbedUnimplementedAuthorizerServer()
}

func RegisterAuthorizerServer(s grpc.ServiceRegistrar, srv AuthorizerServer) {
	s.RegisterService(&Authorizer_ServiceDesc, srv)
}

func _Authorizer_IsAuthorized_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizerServer).IsAuthorized(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fileserver.authorizer/isAuthorized",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizerServer).IsAuthorized(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorizer_AddAuthorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizerServer).AddAuthorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fileserver.authorizer/addAuthorization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizerServer).AddAuthorization(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Authorizer_ServiceDesc is the grpc.ServiceDesc for Authorizer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authorizer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fileserver.authorizer",
	HandlerType: (*AuthorizerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "isAuthorized",
			Handler:    _Authorizer_IsAuthorized_Handler,
		},
		{
			MethodName: "addAuthorization",
			Handler:    _Authorizer_AddAuthorization_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "unified.proto",
}
