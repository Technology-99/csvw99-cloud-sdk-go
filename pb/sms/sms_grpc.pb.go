// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: sms.proto

package sms

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
	SmsRpcService_SendSms_FullMethodName  = "/sms.SmsRpcService/SendSms"
	SmsRpcService_CheckSms_FullMethodName = "/sms.SmsRpcService/CheckSms"
)

// SmsRpcServiceClient is the client API for SmsRpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmsRpcServiceClient interface {
	SendSms(ctx context.Context, in *SmsParams, opts ...grpc.CallOption) (*SmsResp, error)
	CheckSms(ctx context.Context, in *SmsParams, opts ...grpc.CallOption) (*SmsResp, error)
}

type smsRpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSmsRpcServiceClient(cc grpc.ClientConnInterface) SmsRpcServiceClient {
	return &smsRpcServiceClient{cc}
}

func (c *smsRpcServiceClient) SendSms(ctx context.Context, in *SmsParams, opts ...grpc.CallOption) (*SmsResp, error) {
	out := new(SmsResp)
	err := c.cc.Invoke(ctx, SmsRpcService_SendSms_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsRpcServiceClient) CheckSms(ctx context.Context, in *SmsParams, opts ...grpc.CallOption) (*SmsResp, error) {
	out := new(SmsResp)
	err := c.cc.Invoke(ctx, SmsRpcService_CheckSms_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmsRpcServiceServer is the server API for SmsRpcService service.
// All implementations must embed UnimplementedSmsRpcServiceServer
// for forward compatibility
type SmsRpcServiceServer interface {
	SendSms(context.Context, *SmsParams) (*SmsResp, error)
	CheckSms(context.Context, *SmsParams) (*SmsResp, error)
	mustEmbedUnimplementedSmsRpcServiceServer()
}

// UnimplementedSmsRpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSmsRpcServiceServer struct {
}

func (UnimplementedSmsRpcServiceServer) SendSms(context.Context, *SmsParams) (*SmsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSms not implemented")
}
func (UnimplementedSmsRpcServiceServer) CheckSms(context.Context, *SmsParams) (*SmsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckSms not implemented")
}
func (UnimplementedSmsRpcServiceServer) mustEmbedUnimplementedSmsRpcServiceServer() {}

// UnsafeSmsRpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmsRpcServiceServer will
// types in compilation errors.
type UnsafeSmsRpcServiceServer interface {
	mustEmbedUnimplementedSmsRpcServiceServer()
}

func RegisterSmsRpcServiceServer(s grpc.ServiceRegistrar, srv SmsRpcServiceServer) {
	s.RegisterService(&SmsRpcService_ServiceDesc, srv)
}

func _SmsRpcService_SendSms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmsParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsRpcServiceServer).SendSms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmsRpcService_SendSms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsRpcServiceServer).SendSms(ctx, req.(*SmsParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmsRpcService_CheckSms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmsParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsRpcServiceServer).CheckSms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmsRpcService_CheckSms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsRpcServiceServer).CheckSms(ctx, req.(*SmsParams))
	}
	return interceptor(ctx, in, info, handler)
}

// SmsRpcService_ServiceDesc is the grpc.ServiceDesc for SmsRpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SmsRpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sms.SmsRpcService",
	HandlerType: (*SmsRpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSms",
			Handler:    _SmsRpcService_SendSms_Handler,
		},
		{
			MethodName: "CheckSms",
			Handler:    _SmsRpcService_CheckSms_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sms.proto",
}
