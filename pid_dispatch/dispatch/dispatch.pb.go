// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dispatch.proto

package dispatch

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

type GetPidReq struct {
	Addr                 string   `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	App                  string   `protobuf:"bytes,2,opt,name=app,proto3" json:"app,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPidReq) Reset()         { *m = GetPidReq{} }
func (m *GetPidReq) String() string { return proto.CompactTextString(m) }
func (*GetPidReq) ProtoMessage()    {}
func (*GetPidReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3fbf3dcaa8c6dfa, []int{0}
}

func (m *GetPidReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPidReq.Unmarshal(m, b)
}
func (m *GetPidReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPidReq.Marshal(b, m, deterministic)
}
func (m *GetPidReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPidReq.Merge(m, src)
}
func (m *GetPidReq) XXX_Size() int {
	return xxx_messageInfo_GetPidReq.Size(m)
}
func (m *GetPidReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPidReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetPidReq proto.InternalMessageInfo

func (m *GetPidReq) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *GetPidReq) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

type GetPidRsp struct {
	Pid                  int32    `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPidRsp) Reset()         { *m = GetPidRsp{} }
func (m *GetPidRsp) String() string { return proto.CompactTextString(m) }
func (*GetPidRsp) ProtoMessage()    {}
func (*GetPidRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3fbf3dcaa8c6dfa, []int{1}
}

func (m *GetPidRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPidRsp.Unmarshal(m, b)
}
func (m *GetPidRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPidRsp.Marshal(b, m, deterministic)
}
func (m *GetPidRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPidRsp.Merge(m, src)
}
func (m *GetPidRsp) XXX_Size() int {
	return xxx_messageInfo_GetPidRsp.Size(m)
}
func (m *GetPidRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPidRsp.DiscardUnknown(m)
}

var xxx_messageInfo_GetPidRsp proto.InternalMessageInfo

func (m *GetPidRsp) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func init() {
	proto.RegisterType((*GetPidReq)(nil), "GetPidReq")
	proto.RegisterType((*GetPidRsp)(nil), "GetPidRsp")
}

func init() { proto.RegisterFile("dispatch.proto", fileDescriptor_b3fbf3dcaa8c6dfa) }

var fileDescriptor_b3fbf3dcaa8c6dfa = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xc9, 0x2c, 0x2e,
	0x48, 0x2c, 0x49, 0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x32, 0xe4, 0xe2, 0x74, 0x4f,
	0x2d, 0x09, 0xc8, 0x4c, 0x09, 0x4a, 0x2d, 0x14, 0x12, 0xe2, 0x62, 0x49, 0x4c, 0x49, 0x29, 0x92,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x04, 0xb8, 0x98, 0x13, 0x0b, 0x0a, 0x24,
	0x98, 0xc0, 0x42, 0x20, 0xa6, 0x92, 0x2c, 0x5c, 0x4b, 0x71, 0x01, 0x48, 0xba, 0x20, 0x33, 0x05,
	0xac, 0x83, 0x35, 0x08, 0xc4, 0x34, 0x32, 0xe4, 0xe2, 0x0e, 0xc8, 0x4c, 0x71, 0x81, 0x5a, 0x23,
	0xa4, 0xc4, 0xc5, 0x06, 0x51, 0x2d, 0xc4, 0xa5, 0x07, 0xb7, 0x49, 0x0a, 0xce, 0x2e, 0x2e, 0x50,
	0x62, 0x48, 0x62, 0x03, 0xbb, 0xc5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x9a, 0xf5, 0x1c,
	0x9d, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PidDispatchClient is the client API for PidDispatch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PidDispatchClient interface {
	GetPid(ctx context.Context, in *GetPidReq, opts ...grpc.CallOption) (*GetPidRsp, error)
}

type pidDispatchClient struct {
	cc *grpc.ClientConn
}

func NewPidDispatchClient(cc *grpc.ClientConn) PidDispatchClient {
	return &pidDispatchClient{cc}
}

func (c *pidDispatchClient) GetPid(ctx context.Context, in *GetPidReq, opts ...grpc.CallOption) (*GetPidRsp, error) {
	out := new(GetPidRsp)
	err := c.cc.Invoke(ctx, "/PidDispatch/GetPid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PidDispatchServer is the server API for PidDispatch service.
type PidDispatchServer interface {
	GetPid(context.Context, *GetPidReq) (*GetPidRsp, error)
}

// UnimplementedPidDispatchServer can be embedded to have forward compatible implementations.
type UnimplementedPidDispatchServer struct {
}

func (*UnimplementedPidDispatchServer) GetPid(ctx context.Context, req *GetPidReq) (*GetPidRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPid not implemented")
}

func RegisterPidDispatchServer(s *grpc.Server, srv PidDispatchServer) {
	s.RegisterService(&_PidDispatch_serviceDesc, srv)
}

func _PidDispatch_GetPid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PidDispatchServer).GetPid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PidDispatch/GetPid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PidDispatchServer).GetPid(ctx, req.(*GetPidReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _PidDispatch_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PidDispatch",
	HandlerType: (*PidDispatchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPid",
			Handler:    _PidDispatch_GetPid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dispatch.proto",
}
