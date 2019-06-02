// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth/authpb/auth.proto

package authpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Authentication struct {
	Uuid                 []byte   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Authentication) Reset()         { *m = Authentication{} }
func (m *Authentication) String() string { return proto.CompactTextString(m) }
func (*Authentication) ProtoMessage()    {}
func (*Authentication) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_037176464cf267d1, []int{0}
}
func (m *Authentication) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Authentication.Unmarshal(m, b)
}
func (m *Authentication) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Authentication.Marshal(b, m, deterministic)
}
func (dst *Authentication) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Authentication.Merge(dst, src)
}
func (m *Authentication) XXX_Size() int {
	return xxx_messageInfo_Authentication.Size(m)
}
func (m *Authentication) XXX_DiscardUnknown() {
	xxx_messageInfo_Authentication.DiscardUnknown(m)
}

var xxx_messageInfo_Authentication proto.InternalMessageInfo

func (m *Authentication) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

func (m *Authentication) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthenticateRequest struct {
	Auth                 *Authentication `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AuthenticateRequest) Reset()         { *m = AuthenticateRequest{} }
func (m *AuthenticateRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticateRequest) ProtoMessage()    {}
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_037176464cf267d1, []int{1}
}
func (m *AuthenticateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateRequest.Unmarshal(m, b)
}
func (m *AuthenticateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateRequest.Marshal(b, m, deterministic)
}
func (dst *AuthenticateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateRequest.Merge(dst, src)
}
func (m *AuthenticateRequest) XXX_Size() int {
	return xxx_messageInfo_AuthenticateRequest.Size(m)
}
func (m *AuthenticateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateRequest proto.InternalMessageInfo

func (m *AuthenticateRequest) GetAuth() *Authentication {
	if m != nil {
		return m.Auth
	}
	return nil
}

type AuthenticateResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticateResponse) Reset()         { *m = AuthenticateResponse{} }
func (m *AuthenticateResponse) String() string { return proto.CompactTextString(m) }
func (*AuthenticateResponse) ProtoMessage()    {}
func (*AuthenticateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_037176464cf267d1, []int{2}
}
func (m *AuthenticateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateResponse.Unmarshal(m, b)
}
func (m *AuthenticateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateResponse.Marshal(b, m, deterministic)
}
func (dst *AuthenticateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateResponse.Merge(dst, src)
}
func (m *AuthenticateResponse) XXX_Size() int {
	return xxx_messageInfo_AuthenticateResponse.Size(m)
}
func (m *AuthenticateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateResponse proto.InternalMessageInfo

func (m *AuthenticateResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*Authentication)(nil), "auth.Authentication")
	proto.RegisterType((*AuthenticateRequest)(nil), "auth.AuthenticateRequest")
	proto.RegisterType((*AuthenticateResponse)(nil), "auth.AuthenticateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _AuthService_Authenticate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/authpb/auth.proto",
}

func init() { proto.RegisterFile("auth/authpb/auth.proto", fileDescriptor_auth_037176464cf267d1) }

var fileDescriptor_auth_037176464cf267d1 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x2c, 0x2d, 0xc9,
	0xd0, 0x07, 0x11, 0x05, 0x49, 0x60, 0x4a, 0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x88, 0x05, 0xc4,
	0x56, 0x72, 0xe0, 0xe2, 0x73, 0x2c, 0x2d, 0xc9, 0x48, 0xcd, 0x2b, 0xc9, 0x4c, 0x4e, 0x2c, 0xc9,
	0xcc, 0xcf, 0x13, 0x12, 0xe2, 0x62, 0x29, 0x2d, 0xcd, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x09, 0x02, 0xb3, 0x85, 0xa4, 0xb8, 0x38, 0x0a, 0x12, 0x8b, 0x8b, 0xcb, 0xf3, 0x8b, 0x52, 0x24,
	0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x25, 0x7b, 0x2e, 0x61, 0x24, 0x13, 0x52, 0x83,
	0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x34, 0xb8, 0xc0, 0x16, 0x80, 0x8d, 0xe1, 0x36, 0x12,
	0xd1, 0x03, 0xdb, 0x8c, 0x6a, 0x55, 0x10, 0xc4, 0x09, 0x3a, 0x5c, 0x22, 0xa8, 0x06, 0x14, 0x17,
	0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x89, 0x70, 0xb1, 0x96, 0xe4, 0x67, 0xa7, 0xe6, 0x81, 0x8d, 0xe0,
	0x0c, 0x82, 0x70, 0x8c, 0xc2, 0xb8, 0xb8, 0x41, 0xaa, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53,
	0x85, 0xdc, 0xb9, 0x78, 0x90, 0x35, 0x0b, 0x49, 0x62, 0x58, 0x04, 0x73, 0x91, 0x94, 0x14, 0x36,
	0x29, 0x88, 0x5d, 0x4a, 0x0c, 0x4e, 0x1c, 0x51, 0x6c, 0x90, 0x30, 0x4a, 0x62, 0x03, 0x87, 0x8f,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x05, 0xdb, 0xba, 0x36, 0x39, 0x01, 0x00, 0x00,
}
