// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: chat.proto

package chat_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChatV1Client is the client API for ChatV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatV1Client interface {
	WriteMessage(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetMessagesFromChat(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error)
	CreateChat(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetChat(ctx context.Context, in *GetChatRequest, opts ...grpc.CallOption) (*GetChatResponse, error)
	AddUsersToChat(ctx context.Context, in *AddUsersToChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
}

type chatV1Client struct {
	cc grpc.ClientConnInterface
}

func NewChatV1Client(cc grpc.ClientConnInterface) ChatV1Client {
	return &chatV1Client{cc}
}

func (c *chatV1Client) WriteMessage(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/chat_v1.ChatV1/writeMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatV1Client) GetMessagesFromChat(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error) {
	out := new(GetMessagesResponse)
	err := c.cc.Invoke(ctx, "/chat_v1.ChatV1/getMessagesFromChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatV1Client) CreateChat(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/chat_v1.ChatV1/createChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatV1Client) GetChat(ctx context.Context, in *GetChatRequest, opts ...grpc.CallOption) (*GetChatResponse, error) {
	out := new(GetChatResponse)
	err := c.cc.Invoke(ctx, "/chat_v1.ChatV1/getChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatV1Client) AddUsersToChat(ctx context.Context, in *AddUsersToChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/chat_v1.ChatV1/addUsersToChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatV1Client) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/chat_v1.ChatV1/createUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatV1Client) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/chat_v1.ChatV1/getUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatV1Server is the server API for ChatV1 service.
// All implementations must embed UnimplementedChatV1Server
// for forward compatibility
type ChatV1Server interface {
	WriteMessage(context.Context, *WriteRequest) (*emptypb.Empty, error)
	GetMessagesFromChat(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error)
	CreateChat(context.Context, *CreateChatRequest) (*emptypb.Empty, error)
	GetChat(context.Context, *GetChatRequest) (*GetChatResponse, error)
	AddUsersToChat(context.Context, *AddUsersToChatRequest) (*emptypb.Empty, error)
	CreateUser(context.Context, *CreateUserRequest) (*emptypb.Empty, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	mustEmbedUnimplementedChatV1Server()
}

// UnimplementedChatV1Server must be embedded to have forward compatible implementations.
type UnimplementedChatV1Server struct {
}

func (UnimplementedChatV1Server) WriteMessage(context.Context, *WriteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteMessage not implemented")
}
func (UnimplementedChatV1Server) GetMessagesFromChat(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessagesFromChat not implemented")
}
func (UnimplementedChatV1Server) CreateChat(context.Context, *CreateChatRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChat not implemented")
}
func (UnimplementedChatV1Server) GetChat(context.Context, *GetChatRequest) (*GetChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChat not implemented")
}
func (UnimplementedChatV1Server) AddUsersToChat(context.Context, *AddUsersToChatRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUsersToChat not implemented")
}
func (UnimplementedChatV1Server) CreateUser(context.Context, *CreateUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedChatV1Server) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedChatV1Server) mustEmbedUnimplementedChatV1Server() {}

// UnsafeChatV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatV1Server will
// result in compilation errors.
type UnsafeChatV1Server interface {
	mustEmbedUnimplementedChatV1Server()
}

func RegisterChatV1Server(s grpc.ServiceRegistrar, srv ChatV1Server) {
	s.RegisterService(&ChatV1_ServiceDesc, srv)
}

func _ChatV1_WriteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatV1Server).WriteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat_v1.ChatV1/writeMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatV1Server).WriteMessage(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatV1_GetMessagesFromChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatV1Server).GetMessagesFromChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat_v1.ChatV1/getMessagesFromChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatV1Server).GetMessagesFromChat(ctx, req.(*GetMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatV1_CreateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatV1Server).CreateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat_v1.ChatV1/createChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatV1Server).CreateChat(ctx, req.(*CreateChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatV1_GetChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatV1Server).GetChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat_v1.ChatV1/getChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatV1Server).GetChat(ctx, req.(*GetChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatV1_AddUsersToChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUsersToChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatV1Server).AddUsersToChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat_v1.ChatV1/addUsersToChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatV1Server).AddUsersToChat(ctx, req.(*AddUsersToChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatV1_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatV1Server).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat_v1.ChatV1/createUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatV1Server).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatV1_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatV1Server).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat_v1.ChatV1/getUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatV1Server).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatV1_ServiceDesc is the grpc.ServiceDesc for ChatV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat_v1.ChatV1",
	HandlerType: (*ChatV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "writeMessage",
			Handler:    _ChatV1_WriteMessage_Handler,
		},
		{
			MethodName: "getMessagesFromChat",
			Handler:    _ChatV1_GetMessagesFromChat_Handler,
		},
		{
			MethodName: "createChat",
			Handler:    _ChatV1_CreateChat_Handler,
		},
		{
			MethodName: "getChat",
			Handler:    _ChatV1_GetChat_Handler,
		},
		{
			MethodName: "addUsersToChat",
			Handler:    _ChatV1_AddUsersToChat_Handler,
		},
		{
			MethodName: "createUser",
			Handler:    _ChatV1_CreateUser_Handler,
		},
		{
			MethodName: "getUser",
			Handler:    _ChatV1_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
