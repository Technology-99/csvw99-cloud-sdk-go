// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: mix.proto

package mix

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
	MixRpcService_SendSms_FullMethodName            = "/mix.MixRpcService/SendSms"
	MixRpcService_CheckSms_FullMethodName           = "/mix.MixRpcService/CheckSms"
	MixRpcService_CheckEmsRpc_FullMethodName        = "/mix.MixRpcService/CheckEmsRpc"
	MixRpcService_SendEmsRpc_FullMethodName         = "/mix.MixRpcService/SendEmsRpc"
	MixRpcService_GenerateUploadSign_FullMethodName = "/mix.MixRpcService/GenerateUploadSign"
	MixRpcService_CheckUploadResult_FullMethodName  = "/mix.MixRpcService/CheckUploadResult"
	MixRpcService_Callback_FullMethodName           = "/mix.MixRpcService/Callback"
	MixRpcService_CaptchaGenerate_FullMethodName    = "/mix.MixRpcService/CaptchaGenerate"
)

// MixRpcServiceClient is the client API for MixRpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MixRpcServiceClient interface {
	// note: 短信服务
	SendSms(ctx context.Context, in *SmsParams, opts ...grpc.CallOption) (*SmsResp, error)
	CheckSms(ctx context.Context, in *SmsSearchParams, opts ...grpc.CallOption) (*SmsResp, error)
	// note: 邮件服务
	CheckEmsRpc(ctx context.Context, in *EmsReq, opts ...grpc.CallOption) (*EmsResp, error)
	SendEmsRpc(ctx context.Context, in *EmsReq, opts ...grpc.CallOption) (*EmsResp, error)
	// note: 云存储服务
	GenerateUploadSign(ctx context.Context, in *GenerateUploadSignParams, opts ...grpc.CallOption) (*GenerateUploadSignParamsResp, error)
	CheckUploadResult(ctx context.Context, in *CheckUploadResultParams, opts ...grpc.CallOption) (*CheckUploadResultResp, error)
	Callback(ctx context.Context, in *CallbackParams, opts ...grpc.CallOption) (*CallbackResp, error)
	// note: 验证码服务
	CaptchaGenerate(ctx context.Context, in *CaptchaReq, opts ...grpc.CallOption) (*CaptchaResp, error)
}

type mixRpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMixRpcServiceClient(cc grpc.ClientConnInterface) MixRpcServiceClient {
	return &mixRpcServiceClient{cc}
}

func (c *mixRpcServiceClient) SendSms(ctx context.Context, in *SmsParams, opts ...grpc.CallOption) (*SmsResp, error) {
	out := new(SmsResp)
	err := c.cc.Invoke(ctx, MixRpcService_SendSms_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixRpcServiceClient) CheckSms(ctx context.Context, in *SmsSearchParams, opts ...grpc.CallOption) (*SmsResp, error) {
	out := new(SmsResp)
	err := c.cc.Invoke(ctx, MixRpcService_CheckSms_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixRpcServiceClient) CheckEmsRpc(ctx context.Context, in *EmsReq, opts ...grpc.CallOption) (*EmsResp, error) {
	out := new(EmsResp)
	err := c.cc.Invoke(ctx, MixRpcService_CheckEmsRpc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixRpcServiceClient) SendEmsRpc(ctx context.Context, in *EmsReq, opts ...grpc.CallOption) (*EmsResp, error) {
	out := new(EmsResp)
	err := c.cc.Invoke(ctx, MixRpcService_SendEmsRpc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixRpcServiceClient) GenerateUploadSign(ctx context.Context, in *GenerateUploadSignParams, opts ...grpc.CallOption) (*GenerateUploadSignParamsResp, error) {
	out := new(GenerateUploadSignParamsResp)
	err := c.cc.Invoke(ctx, MixRpcService_GenerateUploadSign_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixRpcServiceClient) CheckUploadResult(ctx context.Context, in *CheckUploadResultParams, opts ...grpc.CallOption) (*CheckUploadResultResp, error) {
	out := new(CheckUploadResultResp)
	err := c.cc.Invoke(ctx, MixRpcService_CheckUploadResult_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixRpcServiceClient) Callback(ctx context.Context, in *CallbackParams, opts ...grpc.CallOption) (*CallbackResp, error) {
	out := new(CallbackResp)
	err := c.cc.Invoke(ctx, MixRpcService_Callback_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mixRpcServiceClient) CaptchaGenerate(ctx context.Context, in *CaptchaReq, opts ...grpc.CallOption) (*CaptchaResp, error) {
	out := new(CaptchaResp)
	err := c.cc.Invoke(ctx, MixRpcService_CaptchaGenerate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MixRpcServiceServer is the server API for MixRpcService service.
// All implementations must embed UnimplementedMixRpcServiceServer
// for forward compatibility
type MixRpcServiceServer interface {
	// note: 短信服务
	SendSms(context.Context, *SmsParams) (*SmsResp, error)
	CheckSms(context.Context, *SmsSearchParams) (*SmsResp, error)
	// note: 邮件服务
	CheckEmsRpc(context.Context, *EmsReq) (*EmsResp, error)
	SendEmsRpc(context.Context, *EmsReq) (*EmsResp, error)
	// note: 云存储服务
	GenerateUploadSign(context.Context, *GenerateUploadSignParams) (*GenerateUploadSignParamsResp, error)
	CheckUploadResult(context.Context, *CheckUploadResultParams) (*CheckUploadResultResp, error)
	Callback(context.Context, *CallbackParams) (*CallbackResp, error)
	// note: 验证码服务
	CaptchaGenerate(context.Context, *CaptchaReq) (*CaptchaResp, error)
	mustEmbedUnimplementedMixRpcServiceServer()
}

// UnimplementedMixRpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMixRpcServiceServer struct {
}

func (UnimplementedMixRpcServiceServer) SendSms(context.Context, *SmsParams) (*SmsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSms not implemented")
}
func (UnimplementedMixRpcServiceServer) CheckSms(context.Context, *SmsSearchParams) (*SmsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckSms not implemented")
}
func (UnimplementedMixRpcServiceServer) CheckEmsRpc(context.Context, *EmsReq) (*EmsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckEmsRpc not implemented")
}
func (UnimplementedMixRpcServiceServer) SendEmsRpc(context.Context, *EmsReq) (*EmsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmsRpc not implemented")
}
func (UnimplementedMixRpcServiceServer) GenerateUploadSign(context.Context, *GenerateUploadSignParams) (*GenerateUploadSignParamsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateUploadSign not implemented")
}
func (UnimplementedMixRpcServiceServer) CheckUploadResult(context.Context, *CheckUploadResultParams) (*CheckUploadResultResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUploadResult not implemented")
}
func (UnimplementedMixRpcServiceServer) Callback(context.Context, *CallbackParams) (*CallbackResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Callback not implemented")
}
func (UnimplementedMixRpcServiceServer) CaptchaGenerate(context.Context, *CaptchaReq) (*CaptchaResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CaptchaGenerate not implemented")
}
func (UnimplementedMixRpcServiceServer) mustEmbedUnimplementedMixRpcServiceServer() {}

// UnsafeMixRpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MixRpcServiceServer will
// result in compilation errors.
type UnsafeMixRpcServiceServer interface {
	mustEmbedUnimplementedMixRpcServiceServer()
}

func RegisterMixRpcServiceServer(s grpc.ServiceRegistrar, srv MixRpcServiceServer) {
	s.RegisterService(&MixRpcService_ServiceDesc, srv)
}

func _MixRpcService_SendSms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmsParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixRpcServiceServer).SendSms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MixRpcService_SendSms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixRpcServiceServer).SendSms(ctx, req.(*SmsParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _MixRpcService_CheckSms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmsSearchParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixRpcServiceServer).CheckSms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MixRpcService_CheckSms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixRpcServiceServer).CheckSms(ctx, req.(*SmsSearchParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _MixRpcService_CheckEmsRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixRpcServiceServer).CheckEmsRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MixRpcService_CheckEmsRpc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixRpcServiceServer).CheckEmsRpc(ctx, req.(*EmsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MixRpcService_SendEmsRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixRpcServiceServer).SendEmsRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MixRpcService_SendEmsRpc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixRpcServiceServer).SendEmsRpc(ctx, req.(*EmsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MixRpcService_GenerateUploadSign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateUploadSignParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixRpcServiceServer).GenerateUploadSign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MixRpcService_GenerateUploadSign_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixRpcServiceServer).GenerateUploadSign(ctx, req.(*GenerateUploadSignParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _MixRpcService_CheckUploadResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUploadResultParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixRpcServiceServer).CheckUploadResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MixRpcService_CheckUploadResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixRpcServiceServer).CheckUploadResult(ctx, req.(*CheckUploadResultParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _MixRpcService_Callback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallbackParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixRpcServiceServer).Callback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MixRpcService_Callback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixRpcServiceServer).Callback(ctx, req.(*CallbackParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _MixRpcService_CaptchaGenerate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CaptchaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MixRpcServiceServer).CaptchaGenerate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MixRpcService_CaptchaGenerate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MixRpcServiceServer).CaptchaGenerate(ctx, req.(*CaptchaReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MixRpcService_ServiceDesc is the grpc.ServiceDesc for MixRpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MixRpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mix.MixRpcService",
	HandlerType: (*MixRpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSms",
			Handler:    _MixRpcService_SendSms_Handler,
		},
		{
			MethodName: "CheckSms",
			Handler:    _MixRpcService_CheckSms_Handler,
		},
		{
			MethodName: "CheckEmsRpc",
			Handler:    _MixRpcService_CheckEmsRpc_Handler,
		},
		{
			MethodName: "SendEmsRpc",
			Handler:    _MixRpcService_SendEmsRpc_Handler,
		},
		{
			MethodName: "GenerateUploadSign",
			Handler:    _MixRpcService_GenerateUploadSign_Handler,
		},
		{
			MethodName: "CheckUploadResult",
			Handler:    _MixRpcService_CheckUploadResult_Handler,
		},
		{
			MethodName: "Callback",
			Handler:    _MixRpcService_Callback_Handler,
		},
		{
			MethodName: "CaptchaGenerate",
			Handler:    _MixRpcService_CaptchaGenerate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mix.proto",
}
