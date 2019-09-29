// Code generated by protoc-gen-go. DO NOT EDIT.
// source: provider.proto

package provider

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

type ServerInfo struct {
	App                  string   `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	Set                  string   `protobuf:"bytes,2,opt,name=set,proto3" json:"set,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Addr                 string   `protobuf:"bytes,4,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerInfo) Reset()         { *m = ServerInfo{} }
func (m *ServerInfo) String() string { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()    {}
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6a9f3c02af3d1c8, []int{0}
}

func (m *ServerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerInfo.Unmarshal(m, b)
}
func (m *ServerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerInfo.Marshal(b, m, deterministic)
}
func (m *ServerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerInfo.Merge(m, src)
}
func (m *ServerInfo) XXX_Size() int {
	return xxx_messageInfo_ServerInfo.Size(m)
}
func (m *ServerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ServerInfo proto.InternalMessageInfo

func (m *ServerInfo) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *ServerInfo) GetSet() string {
	if m != nil {
		return m.Set
	}
	return ""
}

func (m *ServerInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServerInfo) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type RegistryReq struct {
	NewServer            *ServerInfo `protobuf:"bytes,1,opt,name=new_server,json=newServer,proto3" json:"new_server,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RegistryReq) Reset()         { *m = RegistryReq{} }
func (m *RegistryReq) String() string { return proto.CompactTextString(m) }
func (*RegistryReq) ProtoMessage()    {}
func (*RegistryReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6a9f3c02af3d1c8, []int{1}
}

func (m *RegistryReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistryReq.Unmarshal(m, b)
}
func (m *RegistryReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistryReq.Marshal(b, m, deterministic)
}
func (m *RegistryReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistryReq.Merge(m, src)
}
func (m *RegistryReq) XXX_Size() int {
	return xxx_messageInfo_RegistryReq.Size(m)
}
func (m *RegistryReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistryReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegistryReq proto.InternalMessageInfo

func (m *RegistryReq) GetNewServer() *ServerInfo {
	if m != nil {
		return m.NewServer
	}
	return nil
}

type RegistryRsp struct {
	ServerList           []*ServerInfo `protobuf:"bytes,1,rep,name=server_list,json=serverList,proto3" json:"server_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RegistryRsp) Reset()         { *m = RegistryRsp{} }
func (m *RegistryRsp) String() string { return proto.CompactTextString(m) }
func (*RegistryRsp) ProtoMessage()    {}
func (*RegistryRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6a9f3c02af3d1c8, []int{2}
}

func (m *RegistryRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistryRsp.Unmarshal(m, b)
}
func (m *RegistryRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistryRsp.Marshal(b, m, deterministic)
}
func (m *RegistryRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistryRsp.Merge(m, src)
}
func (m *RegistryRsp) XXX_Size() int {
	return xxx_messageInfo_RegistryRsp.Size(m)
}
func (m *RegistryRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistryRsp.DiscardUnknown(m)
}

var xxx_messageInfo_RegistryRsp proto.InternalMessageInfo

func (m *RegistryRsp) GetServerList() []*ServerInfo {
	if m != nil {
		return m.ServerList
	}
	return nil
}

type HeartBeatReq struct {
	Addr                 string   `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartBeatReq) Reset()         { *m = HeartBeatReq{} }
func (m *HeartBeatReq) String() string { return proto.CompactTextString(m) }
func (*HeartBeatReq) ProtoMessage()    {}
func (*HeartBeatReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6a9f3c02af3d1c8, []int{3}
}

func (m *HeartBeatReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartBeatReq.Unmarshal(m, b)
}
func (m *HeartBeatReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartBeatReq.Marshal(b, m, deterministic)
}
func (m *HeartBeatReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartBeatReq.Merge(m, src)
}
func (m *HeartBeatReq) XXX_Size() int {
	return xxx_messageInfo_HeartBeatReq.Size(m)
}
func (m *HeartBeatReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartBeatReq.DiscardUnknown(m)
}

var xxx_messageInfo_HeartBeatReq proto.InternalMessageInfo

func (m *HeartBeatReq) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type HeartBeatRsp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartBeatRsp) Reset()         { *m = HeartBeatRsp{} }
func (m *HeartBeatRsp) String() string { return proto.CompactTextString(m) }
func (*HeartBeatRsp) ProtoMessage()    {}
func (*HeartBeatRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6a9f3c02af3d1c8, []int{4}
}

func (m *HeartBeatRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartBeatRsp.Unmarshal(m, b)
}
func (m *HeartBeatRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartBeatRsp.Marshal(b, m, deterministic)
}
func (m *HeartBeatRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartBeatRsp.Merge(m, src)
}
func (m *HeartBeatRsp) XXX_Size() int {
	return xxx_messageInfo_HeartBeatRsp.Size(m)
}
func (m *HeartBeatRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartBeatRsp.DiscardUnknown(m)
}

var xxx_messageInfo_HeartBeatRsp proto.InternalMessageInfo

type SubscribeReq struct {
	Addr                 string   `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeReq) Reset()         { *m = SubscribeReq{} }
func (m *SubscribeReq) String() string { return proto.CompactTextString(m) }
func (*SubscribeReq) ProtoMessage()    {}
func (*SubscribeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6a9f3c02af3d1c8, []int{5}
}

func (m *SubscribeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeReq.Unmarshal(m, b)
}
func (m *SubscribeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeReq.Marshal(b, m, deterministic)
}
func (m *SubscribeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeReq.Merge(m, src)
}
func (m *SubscribeReq) XXX_Size() int {
	return xxx_messageInfo_SubscribeReq.Size(m)
}
func (m *SubscribeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeReq.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeReq proto.InternalMessageInfo

func (m *SubscribeReq) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type PublishRsp struct {
	UpdateServerList     []*ServerInfo `protobuf:"bytes,1,rep,name=update_server_list,json=updateServerList,proto3" json:"update_server_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PublishRsp) Reset()         { *m = PublishRsp{} }
func (m *PublishRsp) String() string { return proto.CompactTextString(m) }
func (*PublishRsp) ProtoMessage()    {}
func (*PublishRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6a9f3c02af3d1c8, []int{6}
}

func (m *PublishRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRsp.Unmarshal(m, b)
}
func (m *PublishRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRsp.Marshal(b, m, deterministic)
}
func (m *PublishRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRsp.Merge(m, src)
}
func (m *PublishRsp) XXX_Size() int {
	return xxx_messageInfo_PublishRsp.Size(m)
}
func (m *PublishRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRsp.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRsp proto.InternalMessageInfo

func (m *PublishRsp) GetUpdateServerList() []*ServerInfo {
	if m != nil {
		return m.UpdateServerList
	}
	return nil
}

func init() {
	proto.RegisterType((*ServerInfo)(nil), "ServerInfo")
	proto.RegisterType((*RegistryReq)(nil), "RegistryReq")
	proto.RegisterType((*RegistryRsp)(nil), "RegistryRsp")
	proto.RegisterType((*HeartBeatReq)(nil), "HeartBeatReq")
	proto.RegisterType((*HeartBeatRsp)(nil), "HeartBeatRsp")
	proto.RegisterType((*SubscribeReq)(nil), "SubscribeReq")
	proto.RegisterType((*PublishRsp)(nil), "PublishRsp")
}

func init() { proto.RegisterFile("provider.proto", fileDescriptor_c6a9f3c02af3d1c8) }

var fileDescriptor_c6a9f3c02af3d1c8 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x4a, 0x03, 0x31,
	0x14, 0x6c, 0x6c, 0x29, 0xf6, 0x6d, 0x5b, 0x4a, 0x4e, 0x4b, 0x4f, 0x25, 0xa7, 0xa2, 0x25, 0x48,
	0x3d, 0x15, 0x6f, 0x5e, 0x54, 0x10, 0x91, 0x5d, 0x04, 0x6f, 0x4b, 0xd6, 0x7d, 0xd6, 0x40, 0xbb,
	0x1b, 0x93, 0x6c, 0x8b, 0x1f, 0xe1, 0x3f, 0x4b, 0x12, 0xda, 0x06, 0x41, 0xbc, 0xbd, 0x37, 0x99,
	0x61, 0x66, 0x92, 0xc0, 0x58, 0xe9, 0x66, 0x27, 0x2b, 0xd4, 0x5c, 0xe9, 0xc6, 0x36, 0xec, 0x15,
	0x20, 0x47, 0xbd, 0x43, 0xfd, 0x50, 0xbf, 0x37, 0x74, 0x02, 0x5d, 0xa1, 0x54, 0x4a, 0x66, 0x64,
	0x3e, 0xc8, 0xdc, 0xe8, 0x10, 0x83, 0x36, 0x3d, 0x0b, 0x88, 0x41, 0x4b, 0x29, 0xf4, 0x6a, 0xb1,
	0xc5, 0xb4, 0xeb, 0x21, 0x3f, 0x3b, 0x4c, 0x54, 0x95, 0x4e, 0x7b, 0x01, 0x73, 0x33, 0x5b, 0x41,
	0x92, 0xe1, 0x5a, 0x1a, 0xab, 0xbf, 0x32, 0xfc, 0xa4, 0x17, 0x00, 0x35, 0xee, 0x0b, 0xe3, 0xcd,
	0xbc, 0x43, 0xb2, 0x4c, 0xf8, 0xc9, 0x3b, 0x1b, 0xd4, 0xb8, 0x0f, 0x2b, 0xbb, 0x89, 0xa4, 0x46,
	0xd1, 0x05, 0x24, 0x41, 0x56, 0x6c, 0xa4, 0xb1, 0x29, 0x99, 0x75, 0x7f, 0x6b, 0x21, 0x9c, 0x3f,
	0x4a, 0x63, 0x19, 0x83, 0xe1, 0x3d, 0x0a, 0x6d, 0x6f, 0x51, 0x58, 0x67, 0x7c, 0xc8, 0x46, 0xa2,
	0x6c, 0xe3, 0x98, 0x63, 0x94, 0xd3, 0xe4, 0x6d, 0x69, 0xde, 0xb4, 0x2c, 0xf1, 0x2f, 0xcd, 0x1d,
	0xc0, 0x73, 0x5b, 0x6e, 0xa4, 0xf9, 0x70, 0x99, 0x56, 0x40, 0x5b, 0x55, 0x09, 0x8b, 0xc5, 0x3f,
	0xd1, 0x26, 0x81, 0x96, 0x1f, 0x03, 0x2e, 0xbf, 0x09, 0xf4, 0x9f, 0xc4, 0x56, 0xd6, 0x6b, 0x3a,
	0x87, 0xf3, 0x43, 0x51, 0x3a, 0xe4, 0xd1, 0x75, 0x4d, 0xa3, 0xcd, 0x28, 0xd6, 0xa1, 0x97, 0x30,
	0x38, 0x26, 0xa6, 0x23, 0x1e, 0x37, 0x9c, 0xc6, 0xab, 0x27, 0x2f, 0x00, 0x5e, 0xbc, 0xab, 0x7f,
	0xd4, 0x11, 0x8f, 0xbb, 0x4d, 0x13, 0x7e, 0xaa, 0xc1, 0x3a, 0x57, 0xa4, 0xec, 0xfb, 0x9f, 0x70,
	0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xf9, 0xad, 0xc1, 0x1b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NamingClient is the client API for Naming service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NamingClient interface {
	Registry(ctx context.Context, in *RegistryReq, opts ...grpc.CallOption) (*RegistryRsp, error)
	HeartBeat(ctx context.Context, in *HeartBeatReq, opts ...grpc.CallOption) (*HeartBeatRsp, error)
	UpdateInfo(ctx context.Context, in *SubscribeReq, opts ...grpc.CallOption) (Naming_UpdateInfoClient, error)
}

type namingClient struct {
	cc *grpc.ClientConn
}

func NewNamingClient(cc *grpc.ClientConn) NamingClient {
	return &namingClient{cc}
}

func (c *namingClient) Registry(ctx context.Context, in *RegistryReq, opts ...grpc.CallOption) (*RegistryRsp, error) {
	out := new(RegistryRsp)
	err := c.cc.Invoke(ctx, "/Naming/Registry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namingClient) HeartBeat(ctx context.Context, in *HeartBeatReq, opts ...grpc.CallOption) (*HeartBeatRsp, error) {
	out := new(HeartBeatRsp)
	err := c.cc.Invoke(ctx, "/Naming/HeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namingClient) UpdateInfo(ctx context.Context, in *SubscribeReq, opts ...grpc.CallOption) (Naming_UpdateInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Naming_serviceDesc.Streams[0], "/Naming/UpdateInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &namingUpdateInfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Naming_UpdateInfoClient interface {
	Recv() (*PublishRsp, error)
	grpc.ClientStream
}

type namingUpdateInfoClient struct {
	grpc.ClientStream
}

func (x *namingUpdateInfoClient) Recv() (*PublishRsp, error) {
	m := new(PublishRsp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NamingServer is the server API for Naming service.
type NamingServer interface {
	Registry(context.Context, *RegistryReq) (*RegistryRsp, error)
	HeartBeat(context.Context, *HeartBeatReq) (*HeartBeatRsp, error)
	UpdateInfo(*SubscribeReq, Naming_UpdateInfoServer) error
}

// UnimplementedNamingServer can be embedded to have forward compatible implementations.
type UnimplementedNamingServer struct {
}

func (*UnimplementedNamingServer) Registry(ctx context.Context, req *RegistryReq) (*RegistryRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Registry not implemented")
}
func (*UnimplementedNamingServer) HeartBeat(ctx context.Context, req *HeartBeatReq) (*HeartBeatRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartBeat not implemented")
}
func (*UnimplementedNamingServer) UpdateInfo(req *SubscribeReq, srv Naming_UpdateInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateInfo not implemented")
}

func RegisterNamingServer(s *grpc.Server, srv NamingServer) {
	s.RegisterService(&_Naming_serviceDesc, srv)
}

func _Naming_Registry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamingServer).Registry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Naming/Registry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamingServer).Registry(ctx, req.(*RegistryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Naming_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartBeatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamingServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Naming/HeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamingServer).HeartBeat(ctx, req.(*HeartBeatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Naming_UpdateInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NamingServer).UpdateInfo(m, &namingUpdateInfoServer{stream})
}

type Naming_UpdateInfoServer interface {
	Send(*PublishRsp) error
	grpc.ServerStream
}

type namingUpdateInfoServer struct {
	grpc.ServerStream
}

func (x *namingUpdateInfoServer) Send(m *PublishRsp) error {
	return x.ServerStream.SendMsg(m)
}

var _Naming_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Naming",
	HandlerType: (*NamingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Registry",
			Handler:    _Naming_Registry_Handler,
		},
		{
			MethodName: "HeartBeat",
			Handler:    _Naming_HeartBeat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UpdateInfo",
			Handler:       _Naming_UpdateInfo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "provider.proto",
}
