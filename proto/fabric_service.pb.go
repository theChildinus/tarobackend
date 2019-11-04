// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/fabric_service.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RegisterReq struct {
	Userid               int64    `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterReq) Reset()         { *m = RegisterReq{} }
func (m *RegisterReq) String() string { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()    {}
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b05191437143c17, []int{0}
}

func (m *RegisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReq.Unmarshal(m, b)
}
func (m *RegisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReq.Marshal(b, m, deterministic)
}
func (m *RegisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReq.Merge(m, src)
}
func (m *RegisterReq) XXX_Size() int {
	return xxx_messageInfo_RegisterReq.Size(m)
}
func (m *RegisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReq proto.InternalMessageInfo

func (m *RegisterReq) GetUserid() int64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *RegisterReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type RegisterResp struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResp) Reset()         { *m = RegisterResp{} }
func (m *RegisterResp) String() string { return proto.CompactTextString(m) }
func (*RegisterResp) ProtoMessage()    {}
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b05191437143c17, []int{1}
}

func (m *RegisterResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResp.Unmarshal(m, b)
}
func (m *RegisterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResp.Marshal(b, m, deterministic)
}
func (m *RegisterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResp.Merge(m, src)
}
func (m *RegisterResp) XXX_Size() int {
	return xxx_messageInfo_RegisterResp.Size(m)
}
func (m *RegisterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResp.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResp proto.InternalMessageInfo

func (m *RegisterResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

type DownloadReq struct {
	Userid               int64    `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadReq) Reset()         { *m = DownloadReq{} }
func (m *DownloadReq) String() string { return proto.CompactTextString(m) }
func (*DownloadReq) ProtoMessage()    {}
func (*DownloadReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b05191437143c17, []int{2}
}

func (m *DownloadReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadReq.Unmarshal(m, b)
}
func (m *DownloadReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadReq.Marshal(b, m, deterministic)
}
func (m *DownloadReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadReq.Merge(m, src)
}
func (m *DownloadReq) XXX_Size() int {
	return xxx_messageInfo_DownloadReq.Size(m)
}
func (m *DownloadReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadReq.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadReq proto.InternalMessageInfo

func (m *DownloadReq) GetUserid() int64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *DownloadReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type DownloadResp struct {
	Cert                 string   `protobuf:"bytes,1,opt,name=cert,proto3" json:"cert,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadResp) Reset()         { *m = DownloadResp{} }
func (m *DownloadResp) String() string { return proto.CompactTextString(m) }
func (*DownloadResp) ProtoMessage()    {}
func (*DownloadResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b05191437143c17, []int{3}
}

func (m *DownloadResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadResp.Unmarshal(m, b)
}
func (m *DownloadResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadResp.Marshal(b, m, deterministic)
}
func (m *DownloadResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadResp.Merge(m, src)
}
func (m *DownloadResp) XXX_Size() int {
	return xxx_messageInfo_DownloadResp.Size(m)
}
func (m *DownloadResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadResp.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadResp proto.InternalMessageInfo

func (m *DownloadResp) GetCert() string {
	if m != nil {
		return m.Cert
	}
	return ""
}

type LoginReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Userrand             int64    `protobuf:"varint,2,opt,name=userrand,proto3" json:"userrand,omitempty"`
	Usersign             string   `protobuf:"bytes,3,opt,name=usersign,proto3" json:"usersign,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b05191437143c17, []int{4}
}

func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginReq) GetUserrand() int64 {
	if m != nil {
		return m.Userrand
	}
	return 0
}

func (m *LoginReq) GetUsersign() string {
	if m != nil {
		return m.Usersign
	}
	return ""
}

type LoginResp struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResp) Reset()         { *m = LoginResp{} }
func (m *LoginResp) String() string { return proto.CompactTextString(m) }
func (*LoginResp) ProtoMessage()    {}
func (*LoginResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b05191437143c17, []int{5}
}

func (m *LoginResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResp.Unmarshal(m, b)
}
func (m *LoginResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResp.Marshal(b, m, deterministic)
}
func (m *LoginResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResp.Merge(m, src)
}
func (m *LoginResp) XXX_Size() int {
	return xxx_messageInfo_LoginResp.Size(m)
}
func (m *LoginResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResp.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResp proto.InternalMessageInfo

func (m *LoginResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*RegisterReq)(nil), "proto.RegisterReq")
	proto.RegisterType((*RegisterResp)(nil), "proto.RegisterResp")
	proto.RegisterType((*DownloadReq)(nil), "proto.DownloadReq")
	proto.RegisterType((*DownloadResp)(nil), "proto.DownloadResp")
	proto.RegisterType((*LoginReq)(nil), "proto.LoginReq")
	proto.RegisterType((*LoginResp)(nil), "proto.LoginResp")
}

func init() { proto.RegisterFile("proto/fabric_service.proto", fileDescriptor_4b05191437143c17) }

var fileDescriptor_4b05191437143c17 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x90, 0xc1, 0x4e, 0x84, 0x30,
	0x10, 0x86, 0xb7, 0xe2, 0x6e, 0x60, 0x74, 0xd5, 0x8c, 0x89, 0x21, 0x5c, 0x5c, 0x7b, 0xda, 0x83,
	0xc1, 0x44, 0xe3, 0x03, 0xb8, 0x31, 0x9e, 0x3c, 0x6c, 0xea, 0x5d, 0xc3, 0x42, 0x25, 0x8d, 0xda,
	0x62, 0x8b, 0xfa, 0x50, 0xbe, 0xa4, 0x69, 0x69, 0x81, 0x8d, 0xde, 0x3c, 0x31, 0xff, 0x0c, 0xf3,
	0xff, 0x9d, 0x0f, 0xb2, 0x46, 0xab, 0x56, 0x5d, 0x3c, 0x17, 0x1b, 0x2d, 0xca, 0x27, 0xc3, 0xf5,
	0xa7, 0x28, 0x79, 0xee, 0x9a, 0x38, 0x75, 0x1f, 0x7a, 0x03, 0x7b, 0x8c, 0xd7, 0xc2, 0xb4, 0x5c,
	0x33, 0xfe, 0x8e, 0x27, 0x30, 0xfb, 0x30, 0x5c, 0x8b, 0x2a, 0x25, 0x0b, 0xb2, 0x8c, 0x98, 0x57,
	0x98, 0x41, 0x6c, 0x2b, 0x59, 0xbc, 0xf1, 0x74, 0x67, 0x41, 0x96, 0x09, 0xeb, 0x35, 0xa5, 0xb0,
	0x3f, 0x58, 0x98, 0x06, 0x11, 0x76, 0x4b, 0x55, 0x71, 0xef, 0xe0, 0x6a, 0x1b, 0x73, 0xab, 0xbe,
	0xe4, 0xab, 0x2a, 0xaa, 0x7f, 0xc4, 0x0c, 0x16, 0x3e, 0x86, 0xeb, 0xd6, 0x39, 0x24, 0xcc, 0xd5,
	0xf4, 0x11, 0xe2, 0x7b, 0x55, 0x0b, 0x69, 0x33, 0xc6, 0x5e, 0x64, 0xdb, 0x2b, 0xcc, 0x74, 0x21,
	0x2b, 0x97, 0x13, 0xb1, 0x5e, 0x87, 0x99, 0x11, 0xb5, 0x4c, 0xa3, 0x61, 0xcf, 0x6a, 0x7a, 0x0a,
	0x89, 0xf7, 0xff, 0xfb, 0xce, 0xcb, 0x6f, 0x02, 0xf3, 0x3b, 0x87, 0xfb, 0xa1, 0xa3, 0x8d, 0xd7,
	0x10, 0x07, 0x3a, 0x88, 0x1d, 0xfb, 0x7c, 0x44, 0x3c, 0x3b, 0xfe, 0xd5, 0x33, 0x0d, 0x9d, 0xd8,
	0xb5, 0x70, 0x6d, 0xbf, 0x36, 0x22, 0xd8, 0xaf, 0x8d, 0x91, 0xd0, 0x09, 0x9e, 0xc3, 0xd4, 0x3d,
	0x10, 0x0f, 0xfd, 0x3c, 0xe0, 0xc8, 0x8e, 0xb6, 0x1b, 0xf6, 0xef, 0xd5, 0x19, 0x1c, 0x28, 0x5d,
	0xe7, 0x2f, 0x4a, 0xd6, 0xdd, 0x74, 0x35, 0x0f, 0xcf, 0x58, 0x5b, 0xb9, 0x26, 0x9b, 0x99, 0xeb,
	0x5f, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0x72, 0x2c, 0xdb, 0x59, 0x4b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FabricServiceClient is the client API for FabricService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FabricServiceClient interface {
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
	Download(ctx context.Context, in *DownloadReq, opts ...grpc.CallOption) (*DownloadResp, error)
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
}

type fabricServiceClient struct {
	cc *grpc.ClientConn
}

func NewFabricServiceClient(cc *grpc.ClientConn) FabricServiceClient {
	return &fabricServiceClient{cc}
}

func (c *fabricServiceClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	out := new(RegisterResp)
	err := c.cc.Invoke(ctx, "/proto.FabricService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fabricServiceClient) Download(ctx context.Context, in *DownloadReq, opts ...grpc.CallOption) (*DownloadResp, error) {
	out := new(DownloadResp)
	err := c.cc.Invoke(ctx, "/proto.FabricService/Download", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fabricServiceClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/proto.FabricService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FabricServiceServer is the server API for FabricService service.
type FabricServiceServer interface {
	Register(context.Context, *RegisterReq) (*RegisterResp, error)
	Download(context.Context, *DownloadReq) (*DownloadResp, error)
	Login(context.Context, *LoginReq) (*LoginResp, error)
}

// UnimplementedFabricServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFabricServiceServer struct {
}

func (*UnimplementedFabricServiceServer) Register(ctx context.Context, req *RegisterReq) (*RegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedFabricServiceServer) Download(ctx context.Context, req *DownloadReq) (*DownloadResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Download not implemented")
}
func (*UnimplementedFabricServiceServer) Login(ctx context.Context, req *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterFabricServiceServer(s *grpc.Server, srv FabricServiceServer) {
	s.RegisterService(&_FabricService_serviceDesc, srv)
}

func _FabricService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FabricServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FabricService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FabricServiceServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FabricService_Download_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FabricServiceServer).Download(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FabricService/Download",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FabricServiceServer).Download(ctx, req.(*DownloadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FabricService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FabricServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FabricService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FabricServiceServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _FabricService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.FabricService",
	HandlerType: (*FabricServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _FabricService_Register_Handler,
		},
		{
			MethodName: "Download",
			Handler:    _FabricService_Download_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _FabricService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/fabric_service.proto",
}
