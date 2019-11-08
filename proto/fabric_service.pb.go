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

type RevokeReq struct {
	Userid               int64    `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokeReq) Reset()         { *m = RevokeReq{} }
func (m *RevokeReq) String() string { return proto.CompactTextString(m) }
func (*RevokeReq) ProtoMessage()    {}
func (*RevokeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b05191437143c17, []int{6}
}

func (m *RevokeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeReq.Unmarshal(m, b)
}
func (m *RevokeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeReq.Marshal(b, m, deterministic)
}
func (m *RevokeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeReq.Merge(m, src)
}
func (m *RevokeReq) XXX_Size() int {
	return xxx_messageInfo_RevokeReq.Size(m)
}
func (m *RevokeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeReq.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeReq proto.InternalMessageInfo

func (m *RevokeReq) GetUserid() int64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *RevokeReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type RevokeResp struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokeResp) Reset()         { *m = RevokeResp{} }
func (m *RevokeResp) String() string { return proto.CompactTextString(m) }
func (*RevokeResp) ProtoMessage()    {}
func (*RevokeResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b05191437143c17, []int{7}
}

func (m *RevokeResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeResp.Unmarshal(m, b)
}
func (m *RevokeResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeResp.Marshal(b, m, deterministic)
}
func (m *RevokeResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeResp.Merge(m, src)
}
func (m *RevokeResp) XXX_Size() int {
	return xxx_messageInfo_RevokeResp.Size(m)
}
func (m *RevokeResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeResp.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeResp proto.InternalMessageInfo

func (m *RevokeResp) GetCode() int64 {
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
	proto.RegisterType((*RevokeReq)(nil), "proto.RevokeReq")
	proto.RegisterType((*RevokeResp)(nil), "proto.RevokeResp")
}

func init() { proto.RegisterFile("proto/fabric_service.proto", fileDescriptor_4b05191437143c17) }

var fileDescriptor_4b05191437143c17 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x51, 0x3d, 0x4f, 0xf3, 0x30,
	0x18, 0xac, 0xdf, 0xbe, 0xad, 0x92, 0x07, 0x0a, 0xc5, 0x48, 0x28, 0xca, 0x42, 0xf0, 0xd4, 0x01,
	0xa5, 0x12, 0x88, 0x19, 0x51, 0x21, 0x26, 0x86, 0xca, 0xec, 0xa0, 0x34, 0x31, 0x91, 0x55, 0xb0,
	0x83, 0x1d, 0xca, 0x9f, 0xe5, 0xc7, 0x20, 0x3b, 0x76, 0x92, 0xf2, 0xb1, 0xc0, 0x94, 0xe7, 0x23,
	0x77, 0xe7, 0xe7, 0x0e, 0xe2, 0x4a, 0xc9, 0x5a, 0xce, 0x1f, 0xb3, 0x95, 0xe2, 0xf9, 0x83, 0x66,
	0x6a, 0xc3, 0x73, 0x96, 0xda, 0x21, 0x1e, 0xd9, 0x0f, 0xb9, 0x82, 0x1d, 0xca, 0x4a, 0xae, 0x6b,
	0xa6, 0x28, 0x7b, 0xc1, 0x47, 0x30, 0x7e, 0xd5, 0x4c, 0xf1, 0x22, 0x42, 0x09, 0x9a, 0x0d, 0xa9,
	0xeb, 0x70, 0x0c, 0x81, 0xa9, 0x44, 0xf6, 0xcc, 0xa2, 0x7f, 0x09, 0x9a, 0x85, 0xb4, 0xed, 0x09,
	0x81, 0xdd, 0x8e, 0x42, 0x57, 0x18, 0xc3, 0xff, 0x5c, 0x16, 0xcc, 0x31, 0xd8, 0xda, 0xc8, 0x5c,
	0xcb, 0x37, 0xf1, 0x24, 0xb3, 0xe2, 0x0f, 0x32, 0x1d, 0x85, 0x93, 0x61, 0xaa, 0xb6, 0x0c, 0x21,
	0xb5, 0x35, 0xb9, 0x87, 0xe0, 0x56, 0x96, 0x5c, 0x18, 0x8d, 0x3e, 0x17, 0xda, 0xe6, 0xf2, 0x3b,
	0x95, 0x89, 0xc2, 0xea, 0x0c, 0x69, 0xdb, 0xfb, 0x9d, 0xe6, 0xa5, 0x88, 0x86, 0x1d, 0xce, 0xf4,
	0xe4, 0x18, 0x42, 0xc7, 0xff, 0xc3, 0x9d, 0x97, 0x10, 0x52, 0xb6, 0x91, 0x6b, 0xf6, 0xdb, 0x2b,
	0x13, 0x00, 0x4f, 0xf0, 0xbd, 0xc4, 0xd9, 0x3b, 0x82, 0xc9, 0x8d, 0x4d, 0xf4, 0xae, 0x09, 0x14,
	0x5f, 0x40, 0xe0, 0x03, 0xc0, 0xb8, 0x89, 0x37, 0xed, 0x85, 0x1a, 0x1f, 0x7e, 0x99, 0xe9, 0x8a,
	0x0c, 0x0c, 0xcc, 0x1b, 0xda, 0xc2, 0x7a, 0x21, 0xb5, 0xb0, 0xbe, 0xeb, 0x64, 0x80, 0x4f, 0x61,
	0x64, 0x3d, 0xc0, 0xfb, 0x6e, 0xef, 0x1d, 0x8f, 0xa7, 0xdb, 0x03, 0xfb, 0xf7, 0x1c, 0xc6, 0xcd,
	0x3d, 0x78, 0xda, 0xbe, 0xc2, 0xf9, 0x13, 0x1f, 0x7c, 0x9a, 0x18, 0xc0, 0xe2, 0x04, 0xf6, 0xa4,
	0x2a, 0xd3, 0xb5, 0x14, 0x65, 0xb3, 0x5e, 0x4c, 0xfc, 0xbb, 0x97, 0xa6, 0x5d, 0xa2, 0xd5, 0xd8,
	0xce, 0xcf, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa2, 0xbc, 0xdc, 0x14, 0xdf, 0x02, 0x00, 0x00,
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
	Revoke(ctx context.Context, in *RevokeReq, opts ...grpc.CallOption) (*RevokeResp, error)
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

func (c *fabricServiceClient) Revoke(ctx context.Context, in *RevokeReq, opts ...grpc.CallOption) (*RevokeResp, error) {
	out := new(RevokeResp)
	err := c.cc.Invoke(ctx, "/proto.FabricService/Revoke", in, out, opts...)
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
	Revoke(context.Context, *RevokeReq) (*RevokeResp, error)
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
func (*UnimplementedFabricServiceServer) Revoke(ctx context.Context, req *RevokeReq) (*RevokeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Revoke not implemented")
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

func _FabricService_Revoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FabricServiceServer).Revoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FabricService/Revoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FabricServiceServer).Revoke(ctx, req.(*RevokeReq))
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
		{
			MethodName: "Revoke",
			Handler:    _FabricService_Revoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/fabric_service.proto",
}
