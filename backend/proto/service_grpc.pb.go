// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.0
// source: service.proto

package proto

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

const (
	StreamingToolsService_ListMediaSource_FullMethodName   = "/proto.StreamingToolsService/ListMediaSource"
	StreamingToolsService_UpsertMediaSource_FullMethodName = "/proto.StreamingToolsService/UpsertMediaSource"
	StreamingToolsService_DeleteMediaSource_FullMethodName = "/proto.StreamingToolsService/DeleteMediaSource"
	StreamingToolsService_ListCamera_FullMethodName        = "/proto.StreamingToolsService/ListCamera"
	StreamingToolsService_UpsertCamera_FullMethodName      = "/proto.StreamingToolsService/UpsertCamera"
	StreamingToolsService_DeleteCamera_FullMethodName      = "/proto.StreamingToolsService/DeleteCamera"
	StreamingToolsService_ListMediaFile_FullMethodName     = "/proto.StreamingToolsService/ListMediaFile"
	StreamingToolsService_UpsertMediaFile_FullMethodName   = "/proto.StreamingToolsService/UpsertMediaFile"
	StreamingToolsService_DeleteMediaFile_FullMethodName   = "/proto.StreamingToolsService/DeleteMediaFile"
)

// StreamingToolsServiceClient is the client API for StreamingToolsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamingToolsServiceClient interface {
	// 列出媒体源
	ListMediaSource(ctx context.Context, in *ListMediaSourceRequest, opts ...grpc.CallOption) (*ListMediaSourceResponse, error)
	// 创建/更新媒体源
	UpsertMediaSource(ctx context.Context, in *UpsertMediaSourceRequest, opts ...grpc.CallOption) (*Response, error)
	// 删除媒体源
	DeleteMediaSource(ctx context.Context, in *DeleteMediaSourceRequest, opts ...grpc.CallOption) (*Response, error)
	// 列出摄像头
	ListCamera(ctx context.Context, in *ListCameraRequest, opts ...grpc.CallOption) (*ListCameraResponse, error)
	// 创建/更新摄像头
	UpsertCamera(ctx context.Context, in *UpsertCameraRequest, opts ...grpc.CallOption) (*Response, error)
	// 删除摄像头
	DeleteCamera(ctx context.Context, in *DeleteCameraRequest, opts ...grpc.CallOption) (*Response, error)
	// 列出媒体文件
	ListMediaFile(ctx context.Context, in *ListMediaFileRequest, opts ...grpc.CallOption) (*ListMediaFileResponse, error)
	// 创建/更新媒体文件
	UpsertMediaFile(ctx context.Context, in *UpsertMediaFileRequest, opts ...grpc.CallOption) (*Response, error)
	// 删除媒体文件
	DeleteMediaFile(ctx context.Context, in *DeleteMediaFileRequest, opts ...grpc.CallOption) (*Response, error)
}

type streamingToolsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamingToolsServiceClient(cc grpc.ClientConnInterface) StreamingToolsServiceClient {
	return &streamingToolsServiceClient{cc}
}

func (c *streamingToolsServiceClient) ListMediaSource(ctx context.Context, in *ListMediaSourceRequest, opts ...grpc.CallOption) (*ListMediaSourceResponse, error) {
	out := new(ListMediaSourceResponse)
	err := c.cc.Invoke(ctx, StreamingToolsService_ListMediaSource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamingToolsServiceClient) UpsertMediaSource(ctx context.Context, in *UpsertMediaSourceRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, StreamingToolsService_UpsertMediaSource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamingToolsServiceClient) DeleteMediaSource(ctx context.Context, in *DeleteMediaSourceRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, StreamingToolsService_DeleteMediaSource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamingToolsServiceClient) ListCamera(ctx context.Context, in *ListCameraRequest, opts ...grpc.CallOption) (*ListCameraResponse, error) {
	out := new(ListCameraResponse)
	err := c.cc.Invoke(ctx, StreamingToolsService_ListCamera_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamingToolsServiceClient) UpsertCamera(ctx context.Context, in *UpsertCameraRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, StreamingToolsService_UpsertCamera_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamingToolsServiceClient) DeleteCamera(ctx context.Context, in *DeleteCameraRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, StreamingToolsService_DeleteCamera_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamingToolsServiceClient) ListMediaFile(ctx context.Context, in *ListMediaFileRequest, opts ...grpc.CallOption) (*ListMediaFileResponse, error) {
	out := new(ListMediaFileResponse)
	err := c.cc.Invoke(ctx, StreamingToolsService_ListMediaFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamingToolsServiceClient) UpsertMediaFile(ctx context.Context, in *UpsertMediaFileRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, StreamingToolsService_UpsertMediaFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamingToolsServiceClient) DeleteMediaFile(ctx context.Context, in *DeleteMediaFileRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, StreamingToolsService_DeleteMediaFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamingToolsServiceServer is the server API for StreamingToolsService service.
// All implementations must embed UnimplementedStreamingToolsServiceServer
// for forward compatibility
type StreamingToolsServiceServer interface {
	// 列出媒体源
	ListMediaSource(context.Context, *ListMediaSourceRequest) (*ListMediaSourceResponse, error)
	// 创建/更新媒体源
	UpsertMediaSource(context.Context, *UpsertMediaSourceRequest) (*Response, error)
	// 删除媒体源
	DeleteMediaSource(context.Context, *DeleteMediaSourceRequest) (*Response, error)
	// 列出摄像头
	ListCamera(context.Context, *ListCameraRequest) (*ListCameraResponse, error)
	// 创建/更新摄像头
	UpsertCamera(context.Context, *UpsertCameraRequest) (*Response, error)
	// 删除摄像头
	DeleteCamera(context.Context, *DeleteCameraRequest) (*Response, error)
	// 列出媒体文件
	ListMediaFile(context.Context, *ListMediaFileRequest) (*ListMediaFileResponse, error)
	// 创建/更新媒体文件
	UpsertMediaFile(context.Context, *UpsertMediaFileRequest) (*Response, error)
	// 删除媒体文件
	DeleteMediaFile(context.Context, *DeleteMediaFileRequest) (*Response, error)
	mustEmbedUnimplementedStreamingToolsServiceServer()
}

// UnimplementedStreamingToolsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStreamingToolsServiceServer struct {
}

func (UnimplementedStreamingToolsServiceServer) ListMediaSource(context.Context, *ListMediaSourceRequest) (*ListMediaSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMediaSource not implemented")
}
func (UnimplementedStreamingToolsServiceServer) UpsertMediaSource(context.Context, *UpsertMediaSourceRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertMediaSource not implemented")
}
func (UnimplementedStreamingToolsServiceServer) DeleteMediaSource(context.Context, *DeleteMediaSourceRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMediaSource not implemented")
}
func (UnimplementedStreamingToolsServiceServer) ListCamera(context.Context, *ListCameraRequest) (*ListCameraResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCamera not implemented")
}
func (UnimplementedStreamingToolsServiceServer) UpsertCamera(context.Context, *UpsertCameraRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertCamera not implemented")
}
func (UnimplementedStreamingToolsServiceServer) DeleteCamera(context.Context, *DeleteCameraRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCamera not implemented")
}
func (UnimplementedStreamingToolsServiceServer) ListMediaFile(context.Context, *ListMediaFileRequest) (*ListMediaFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMediaFile not implemented")
}
func (UnimplementedStreamingToolsServiceServer) UpsertMediaFile(context.Context, *UpsertMediaFileRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertMediaFile not implemented")
}
func (UnimplementedStreamingToolsServiceServer) DeleteMediaFile(context.Context, *DeleteMediaFileRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMediaFile not implemented")
}
func (UnimplementedStreamingToolsServiceServer) mustEmbedUnimplementedStreamingToolsServiceServer() {}

// UnsafeStreamingToolsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamingToolsServiceServer will
// result in compilation errors.
type UnsafeStreamingToolsServiceServer interface {
	mustEmbedUnimplementedStreamingToolsServiceServer()
}

func RegisterStreamingToolsServiceServer(s grpc.ServiceRegistrar, srv StreamingToolsServiceServer) {
	s.RegisterService(&StreamingToolsService_ServiceDesc, srv)
}

func _StreamingToolsService_ListMediaSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMediaSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamingToolsServiceServer).ListMediaSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamingToolsService_ListMediaSource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamingToolsServiceServer).ListMediaSource(ctx, req.(*ListMediaSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamingToolsService_UpsertMediaSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertMediaSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamingToolsServiceServer).UpsertMediaSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamingToolsService_UpsertMediaSource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamingToolsServiceServer).UpsertMediaSource(ctx, req.(*UpsertMediaSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamingToolsService_DeleteMediaSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMediaSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamingToolsServiceServer).DeleteMediaSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamingToolsService_DeleteMediaSource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamingToolsServiceServer).DeleteMediaSource(ctx, req.(*DeleteMediaSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamingToolsService_ListCamera_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCameraRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamingToolsServiceServer).ListCamera(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamingToolsService_ListCamera_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamingToolsServiceServer).ListCamera(ctx, req.(*ListCameraRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamingToolsService_UpsertCamera_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertCameraRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamingToolsServiceServer).UpsertCamera(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamingToolsService_UpsertCamera_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamingToolsServiceServer).UpsertCamera(ctx, req.(*UpsertCameraRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamingToolsService_DeleteCamera_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCameraRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamingToolsServiceServer).DeleteCamera(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamingToolsService_DeleteCamera_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamingToolsServiceServer).DeleteCamera(ctx, req.(*DeleteCameraRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamingToolsService_ListMediaFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMediaFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamingToolsServiceServer).ListMediaFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamingToolsService_ListMediaFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamingToolsServiceServer).ListMediaFile(ctx, req.(*ListMediaFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamingToolsService_UpsertMediaFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertMediaFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamingToolsServiceServer).UpsertMediaFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamingToolsService_UpsertMediaFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamingToolsServiceServer).UpsertMediaFile(ctx, req.(*UpsertMediaFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamingToolsService_DeleteMediaFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMediaFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamingToolsServiceServer).DeleteMediaFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamingToolsService_DeleteMediaFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamingToolsServiceServer).DeleteMediaFile(ctx, req.(*DeleteMediaFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StreamingToolsService_ServiceDesc is the grpc.ServiceDesc for StreamingToolsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamingToolsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StreamingToolsService",
	HandlerType: (*StreamingToolsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMediaSource",
			Handler:    _StreamingToolsService_ListMediaSource_Handler,
		},
		{
			MethodName: "UpsertMediaSource",
			Handler:    _StreamingToolsService_UpsertMediaSource_Handler,
		},
		{
			MethodName: "DeleteMediaSource",
			Handler:    _StreamingToolsService_DeleteMediaSource_Handler,
		},
		{
			MethodName: "ListCamera",
			Handler:    _StreamingToolsService_ListCamera_Handler,
		},
		{
			MethodName: "UpsertCamera",
			Handler:    _StreamingToolsService_UpsertCamera_Handler,
		},
		{
			MethodName: "DeleteCamera",
			Handler:    _StreamingToolsService_DeleteCamera_Handler,
		},
		{
			MethodName: "ListMediaFile",
			Handler:    _StreamingToolsService_ListMediaFile_Handler,
		},
		{
			MethodName: "UpsertMediaFile",
			Handler:    _StreamingToolsService_UpsertMediaFile_Handler,
		},
		{
			MethodName: "DeleteMediaFile",
			Handler:    _StreamingToolsService_DeleteMediaFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}