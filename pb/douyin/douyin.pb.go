// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: douyin.proto

package douyin

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OaKeyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *OaKeyReq) Reset() {
	*x = OaKeyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OaKeyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OaKeyReq) ProtoMessage() {}

func (x *OaKeyReq) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OaKeyReq.ProtoReflect.Descriptor instead.
func (*OaKeyReq) Descriptor() ([]byte, []int) {
	return file_douyin_proto_rawDescGZIP(), []int{0}
}

func (x *OaKeyReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type CodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key           string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Code          string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	AnonymousCode string `protobuf:"bytes,3,opt,name=anonymousCode,proto3" json:"anonymousCode,omitempty"`
}

func (x *CodeReq) Reset() {
	*x = CodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeReq) ProtoMessage() {}

func (x *CodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeReq.ProtoReflect.Descriptor instead.
func (*CodeReq) Descriptor() ([]byte, []int) {
	return file_douyin_proto_rawDescGZIP(), []int{1}
}

func (x *CodeReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CodeReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CodeReq) GetAnonymousCode() string {
	if x != nil {
		return x.AnonymousCode
	}
	return ""
}

type CodeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      int32         `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg       string        `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	RequestID string        `protobuf:"bytes,3,opt,name=requestID,proto3" json:"requestID,omitempty"`
	Path      string        `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Data      *CodeRespData `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CodeResp) Reset() {
	*x = CodeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeResp) ProtoMessage() {}

func (x *CodeResp) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeResp.ProtoReflect.Descriptor instead.
func (*CodeResp) Descriptor() ([]byte, []int) {
	return file_douyin_proto_rawDescGZIP(), []int{2}
}

func (x *CodeResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CodeResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CodeResp) GetRequestID() string {
	if x != nil {
		return x.RequestID
	}
	return ""
}

func (x *CodeResp) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CodeResp) GetData() *CodeRespData {
	if x != nil {
		return x.Data
	}
	return nil
}

type CodeRespData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key             string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	SessionKey      string `protobuf:"bytes,2,opt,name=sessionKey,proto3" json:"sessionKey,omitempty"`
	Openid          string `protobuf:"bytes,3,opt,name=openid,proto3" json:"openid,omitempty"`
	AnonymousOpenid string `protobuf:"bytes,4,opt,name=anonymousOpenid,proto3" json:"anonymousOpenid,omitempty"`
	Unionid         string `protobuf:"bytes,5,opt,name=unionid,proto3" json:"unionid,omitempty"`
	ErrNo           int32  `protobuf:"varint,6,opt,name=errNo,proto3" json:"errNo,omitempty"`
	ErrTips         string `protobuf:"bytes,7,opt,name=errTips,proto3" json:"errTips,omitempty"`
}

func (x *CodeRespData) Reset() {
	*x = CodeRespData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeRespData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeRespData) ProtoMessage() {}

func (x *CodeRespData) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeRespData.ProtoReflect.Descriptor instead.
func (*CodeRespData) Descriptor() ([]byte, []int) {
	return file_douyin_proto_rawDescGZIP(), []int{3}
}

func (x *CodeRespData) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CodeRespData) GetSessionKey() string {
	if x != nil {
		return x.SessionKey
	}
	return ""
}

func (x *CodeRespData) GetOpenid() string {
	if x != nil {
		return x.Openid
	}
	return ""
}

func (x *CodeRespData) GetAnonymousOpenid() string {
	if x != nil {
		return x.AnonymousOpenid
	}
	return ""
}

func (x *CodeRespData) GetUnionid() string {
	if x != nil {
		return x.Unionid
	}
	return ""
}

func (x *CodeRespData) GetErrNo() int32 {
	if x != nil {
		return x.ErrNo
	}
	return 0
}

func (x *CodeRespData) GetErrTips() string {
	if x != nil {
		return x.ErrTips
	}
	return ""
}

type OaAccessTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      int32                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg       string                 `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	RequestID string                 `protobuf:"bytes,3,opt,name=requestID,proto3" json:"requestID,omitempty"`
	Path      string                 `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Data      *OaAccessTokenRespData `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OaAccessTokenResp) Reset() {
	*x = OaAccessTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OaAccessTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OaAccessTokenResp) ProtoMessage() {}

func (x *OaAccessTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OaAccessTokenResp.ProtoReflect.Descriptor instead.
func (*OaAccessTokenResp) Descriptor() ([]byte, []int) {
	return file_douyin_proto_rawDescGZIP(), []int{4}
}

func (x *OaAccessTokenResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *OaAccessTokenResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *OaAccessTokenResp) GetRequestID() string {
	if x != nil {
		return x.RequestID
	}
	return ""
}

func (x *OaAccessTokenResp) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *OaAccessTokenResp) GetData() *OaAccessTokenRespData {
	if x != nil {
		return x.Data
	}
	return nil
}

type OaAccessTokenRespData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	ExpiresIn   int64  `protobuf:"varint,2,opt,name=expiresIn,proto3" json:"expiresIn,omitempty"`
	ErrNo       int32  `protobuf:"varint,3,opt,name=errNo,proto3" json:"errNo,omitempty"`
	ErrTips     string `protobuf:"bytes,4,opt,name=errTips,proto3" json:"errTips,omitempty"`
}

func (x *OaAccessTokenRespData) Reset() {
	*x = OaAccessTokenRespData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_douyin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OaAccessTokenRespData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OaAccessTokenRespData) ProtoMessage() {}

func (x *OaAccessTokenRespData) ProtoReflect() protoreflect.Message {
	mi := &file_douyin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OaAccessTokenRespData.ProtoReflect.Descriptor instead.
func (*OaAccessTokenRespData) Descriptor() ([]byte, []int) {
	return file_douyin_proto_rawDescGZIP(), []int{5}
}

func (x *OaAccessTokenRespData) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *OaAccessTokenRespData) GetExpiresIn() int64 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

func (x *OaAccessTokenRespData) GetErrNo() int32 {
	if x != nil {
		return x.ErrNo
	}
	return 0
}

func (x *OaAccessTokenRespData) GetErrTips() string {
	if x != nil {
		return x.ErrTips
	}
	return ""
}

var File_douyin_proto protoreflect.FileDescriptor

var file_douyin_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x22, 0x1c, 0x0a, 0x08, 0x4f, 0x61, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x22, 0x55, 0x0a, 0x07, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x6e,
	0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x8c, 0x01, 0x0a, 0x08,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1c,
	0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xcc, 0x01, 0x0a, 0x0c, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f,
	0x70, 0x65, 0x6e, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f,
	0x75, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x4e, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x72, 0x72, 0x4e, 0x6f, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x54, 0x69, 0x70, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x72, 0x72, 0x54, 0x69, 0x70, 0x73, 0x22, 0x9e, 0x01, 0x0a, 0x11, 0x4f, 0x61,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x4f,
	0x61, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x87, 0x01, 0x0a, 0x15, 0x4f,
	0x61, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x49, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x49, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x4e, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x72, 0x72, 0x4e, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72,
	0x72, 0x54, 0x69, 0x70, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x72, 0x72,
	0x54, 0x69, 0x70, 0x73, 0x32, 0x8e, 0x01, 0x0a, 0x10, 0x44, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x52,
	0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x43, 0x6f, 0x64,
	0x65, 0x32, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0f, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69,
	0x6e, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x49, 0x0a, 0x1a, 0x4f, 0x66,
	0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69,
	0x6e, 0x2e, 0x4f, 0x61, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x64, 0x6f, 0x75,
	0x79, 0x69, 0x6e, 0x2e, 0x4f, 0x61, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x64, 0x6f, 0x75, 0x79, 0x69,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_douyin_proto_rawDescOnce sync.Once
	file_douyin_proto_rawDescData = file_douyin_proto_rawDesc
)

func file_douyin_proto_rawDescGZIP() []byte {
	file_douyin_proto_rawDescOnce.Do(func() {
		file_douyin_proto_rawDescData = protoimpl.X.CompressGZIP(file_douyin_proto_rawDescData)
	})
	return file_douyin_proto_rawDescData
}

var file_douyin_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_douyin_proto_goTypes = []interface{}{
	(*OaKeyReq)(nil),              // 0: douyin.OaKeyReq
	(*CodeReq)(nil),               // 1: douyin.CodeReq
	(*CodeResp)(nil),              // 2: douyin.CodeResp
	(*CodeRespData)(nil),          // 3: douyin.CodeRespData
	(*OaAccessTokenResp)(nil),     // 4: douyin.OaAccessTokenResp
	(*OaAccessTokenRespData)(nil), // 5: douyin.OaAccessTokenRespData
}
var file_douyin_proto_depIdxs = []int32{
	3, // 0: douyin.CodeResp.data:type_name -> douyin.CodeRespData
	5, // 1: douyin.OaAccessTokenResp.data:type_name -> douyin.OaAccessTokenRespData
	1, // 2: douyin.DouyinRpcService.Code2Token:input_type -> douyin.CodeReq
	0, // 3: douyin.DouyinRpcService.OfficialAccountAccessToken:input_type -> douyin.OaKeyReq
	2, // 4: douyin.DouyinRpcService.Code2Token:output_type -> douyin.CodeResp
	4, // 5: douyin.DouyinRpcService.OfficialAccountAccessToken:output_type -> douyin.OaAccessTokenResp
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_douyin_proto_init() }
func file_douyin_proto_init() {
	if File_douyin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_douyin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OaKeyReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_douyin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_douyin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_douyin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeRespData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_douyin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OaAccessTokenResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_douyin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OaAccessTokenRespData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_douyin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_douyin_proto_goTypes,
		DependencyIndexes: file_douyin_proto_depIdxs,
		MessageInfos:      file_douyin_proto_msgTypes,
	}.Build()
	File_douyin_proto = out.File
	file_douyin_proto_rawDesc = nil
	file_douyin_proto_goTypes = nil
	file_douyin_proto_depIdxs = nil
}