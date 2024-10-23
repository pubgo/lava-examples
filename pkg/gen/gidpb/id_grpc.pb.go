// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: gid/id.proto

package gidpb

import (
	context "context"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Id_Generate_FullMethodName       = "/example.gid.Id/Generate"
	Id_TypeStream_FullMethodName     = "/example.gid.Id/TypeStream"
	Id_Types_FullMethodName          = "/example.gid.Id/Types"
	Id_PutTypes_FullMethodName       = "/example.gid.Id/PutTypes"
	Id_Chat_FullMethodName           = "/example.gid.Id/Chat"
	Id_Chat1_FullMethodName          = "/example.gid.Id/Chat1"
	Id_UploadDownload_FullMethodName = "/example.gid.Id/UploadDownload"
	Id_DoProxy_FullMethodName        = "/example.gid.Id/DoProxy"
	Id_ProxyExecEvent_FullMethodName = "/example.gid.Id/ProxyExecEvent"
	Id_EventChanged_FullMethodName   = "/example.gid.Id/EventChanged"
)

// IdClient is the client API for Id service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Id 生成随机ID服务
type IdClient interface {
	// Generate 生成ID
	Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error)
	// 返回流
	TypeStream(ctx context.Context, in *TypesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[TypesResponse], error)
	// Types id类型
	Types(ctx context.Context, in *TypesRequest, opts ...grpc.CallOption) (*TypesResponse, error)
	PutTypes(ctx context.Context, in *TypesRequest, opts ...grpc.CallOption) (*TypesResponse, error)
	// 聊天
	Chat(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ChatMessage, ChatMessage], error)
	// ws: chat1
	Chat1(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ChatMessage, ChatMessage], error)
	UploadDownload(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	DoProxy(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ProxyExecEvent(ctx context.Context, in *DoProxyEventReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	EventChanged(ctx context.Context, in *DoProxyEventReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type idClient struct {
	cc grpc.ClientConnInterface
}

func NewIdClient(cc grpc.ClientConnInterface) IdClient {
	return &idClient{cc}
}

func (c *idClient) Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateResponse)
	err := c.cc.Invoke(ctx, Id_Generate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *idClient) TypeStream(ctx context.Context, in *TypesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[TypesResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Id_ServiceDesc.Streams[0], Id_TypeStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[TypesRequest, TypesResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Id_TypeStreamClient = grpc.ServerStreamingClient[TypesResponse]

func (c *idClient) Types(ctx context.Context, in *TypesRequest, opts ...grpc.CallOption) (*TypesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TypesResponse)
	err := c.cc.Invoke(ctx, Id_Types_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *idClient) PutTypes(ctx context.Context, in *TypesRequest, opts ...grpc.CallOption) (*TypesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TypesResponse)
	err := c.cc.Invoke(ctx, Id_PutTypes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *idClient) Chat(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ChatMessage, ChatMessage], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Id_ServiceDesc.Streams[1], Id_Chat_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ChatMessage, ChatMessage]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Id_ChatClient = grpc.BidiStreamingClient[ChatMessage, ChatMessage]

func (c *idClient) Chat1(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ChatMessage, ChatMessage], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Id_ServiceDesc.Streams[2], Id_Chat1_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ChatMessage, ChatMessage]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Id_Chat1Client = grpc.BidiStreamingClient[ChatMessage, ChatMessage]

func (c *idClient) UploadDownload(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, Id_UploadDownload_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *idClient) DoProxy(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Id_DoProxy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *idClient) ProxyExecEvent(ctx context.Context, in *DoProxyEventReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Id_ProxyExecEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *idClient) EventChanged(ctx context.Context, in *DoProxyEventReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Id_EventChanged_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdServer is the server API for Id service.
// All implementations should embed UnimplementedIdServer
// for forward compatibility.
//
// Id 生成随机ID服务
type IdServer interface {
	// Generate 生成ID
	Generate(context.Context, *GenerateRequest) (*GenerateResponse, error)
	// 返回流
	TypeStream(*TypesRequest, grpc.ServerStreamingServer[TypesResponse]) error
	// Types id类型
	Types(context.Context, *TypesRequest) (*TypesResponse, error)
	PutTypes(context.Context, *TypesRequest) (*TypesResponse, error)
	// 聊天
	Chat(grpc.BidiStreamingServer[ChatMessage, ChatMessage]) error
	// ws: chat1
	Chat1(grpc.BidiStreamingServer[ChatMessage, ChatMessage]) error
	UploadDownload(context.Context, *UploadFileRequest) (*httpbody.HttpBody, error)
	DoProxy(context.Context, *Empty) (*emptypb.Empty, error)
	ProxyExecEvent(context.Context, *DoProxyEventReq) (*emptypb.Empty, error)
	EventChanged(context.Context, *DoProxyEventReq) (*emptypb.Empty, error)
}

// UnimplementedIdServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedIdServer struct{}

func (UnimplementedIdServer) Generate(context.Context, *GenerateRequest) (*GenerateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (UnimplementedIdServer) TypeStream(*TypesRequest, grpc.ServerStreamingServer[TypesResponse]) error {
	return status.Errorf(codes.Unimplemented, "method TypeStream not implemented")
}
func (UnimplementedIdServer) Types(context.Context, *TypesRequest) (*TypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Types not implemented")
}
func (UnimplementedIdServer) PutTypes(context.Context, *TypesRequest) (*TypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutTypes not implemented")
}
func (UnimplementedIdServer) Chat(grpc.BidiStreamingServer[ChatMessage, ChatMessage]) error {
	return status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (UnimplementedIdServer) Chat1(grpc.BidiStreamingServer[ChatMessage, ChatMessage]) error {
	return status.Errorf(codes.Unimplemented, "method Chat1 not implemented")
}
func (UnimplementedIdServer) UploadDownload(context.Context, *UploadFileRequest) (*httpbody.HttpBody, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadDownload not implemented")
}
func (UnimplementedIdServer) DoProxy(context.Context, *Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoProxy not implemented")
}
func (UnimplementedIdServer) ProxyExecEvent(context.Context, *DoProxyEventReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProxyExecEvent not implemented")
}
func (UnimplementedIdServer) EventChanged(context.Context, *DoProxyEventReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EventChanged not implemented")
}
func (UnimplementedIdServer) testEmbeddedByValue() {}

// UnsafeIdServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdServer will
// result in compilation errors.
type UnsafeIdServer interface {
	mustEmbedUnimplementedIdServer()
}

func RegisterIdServer(s grpc.ServiceRegistrar, srv IdServer) {
	// If the following call pancis, it indicates UnimplementedIdServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Id_ServiceDesc, srv)
}

func _Id_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Id_Generate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdServer).Generate(ctx, req.(*GenerateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Id_TypeStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TypesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IdServer).TypeStream(m, &grpc.GenericServerStream[TypesRequest, TypesResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Id_TypeStreamServer = grpc.ServerStreamingServer[TypesResponse]

func _Id_Types_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdServer).Types(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Id_Types_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdServer).Types(ctx, req.(*TypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Id_PutTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdServer).PutTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Id_PutTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdServer).PutTypes(ctx, req.(*TypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Id_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IdServer).Chat(&grpc.GenericServerStream[ChatMessage, ChatMessage]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Id_ChatServer = grpc.BidiStreamingServer[ChatMessage, ChatMessage]

func _Id_Chat1_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IdServer).Chat1(&grpc.GenericServerStream[ChatMessage, ChatMessage]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Id_Chat1Server = grpc.BidiStreamingServer[ChatMessage, ChatMessage]

func _Id_UploadDownload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdServer).UploadDownload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Id_UploadDownload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdServer).UploadDownload(ctx, req.(*UploadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Id_DoProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdServer).DoProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Id_DoProxy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdServer).DoProxy(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Id_ProxyExecEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoProxyEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdServer).ProxyExecEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Id_ProxyExecEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdServer).ProxyExecEvent(ctx, req.(*DoProxyEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Id_EventChanged_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoProxyEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdServer).EventChanged(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Id_EventChanged_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdServer).EventChanged(ctx, req.(*DoProxyEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Id_ServiceDesc is the grpc.ServiceDesc for Id service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Id_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.gid.Id",
	HandlerType: (*IdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _Id_Generate_Handler,
		},
		{
			MethodName: "Types",
			Handler:    _Id_Types_Handler,
		},
		{
			MethodName: "PutTypes",
			Handler:    _Id_PutTypes_Handler,
		},
		{
			MethodName: "UploadDownload",
			Handler:    _Id_UploadDownload_Handler,
		},
		{
			MethodName: "DoProxy",
			Handler:    _Id_DoProxy_Handler,
		},
		{
			MethodName: "ProxyExecEvent",
			Handler:    _Id_ProxyExecEvent_Handler,
		},
		{
			MethodName: "EventChanged",
			Handler:    _Id_EventChanged_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TypeStream",
			Handler:       _Id_TypeStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Chat",
			Handler:       _Id_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Chat1",
			Handler:       _Id_Chat1_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "gid/id.proto",
}
