// Code generated by protoc-gen-go. DO NOT EDIT.
// source: guestbook.proto

package guestbook

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

type AddLeaveRequest struct {
	Leave                *Leave   `protobuf:"bytes,1,opt,name=leave,proto3" json:"leave,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddLeaveRequest) Reset()         { *m = AddLeaveRequest{} }
func (m *AddLeaveRequest) String() string { return proto.CompactTextString(m) }
func (*AddLeaveRequest) ProtoMessage()    {}
func (*AddLeaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbda7dd58e0f267b, []int{0}
}

func (m *AddLeaveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddLeaveRequest.Unmarshal(m, b)
}
func (m *AddLeaveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddLeaveRequest.Marshal(b, m, deterministic)
}
func (m *AddLeaveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddLeaveRequest.Merge(m, src)
}
func (m *AddLeaveRequest) XXX_Size() int {
	return xxx_messageInfo_AddLeaveRequest.Size(m)
}
func (m *AddLeaveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddLeaveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddLeaveRequest proto.InternalMessageInfo

func (m *AddLeaveRequest) GetLeave() *Leave {
	if m != nil {
		return m.Leave
	}
	return nil
}

type AddLeaveResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddLeaveResponse) Reset()         { *m = AddLeaveResponse{} }
func (m *AddLeaveResponse) String() string { return proto.CompactTextString(m) }
func (*AddLeaveResponse) ProtoMessage()    {}
func (*AddLeaveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbda7dd58e0f267b, []int{1}
}

func (m *AddLeaveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddLeaveResponse.Unmarshal(m, b)
}
func (m *AddLeaveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddLeaveResponse.Marshal(b, m, deterministic)
}
func (m *AddLeaveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddLeaveResponse.Merge(m, src)
}
func (m *AddLeaveResponse) XXX_Size() int {
	return xxx_messageInfo_AddLeaveResponse.Size(m)
}
func (m *AddLeaveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddLeaveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddLeaveResponse proto.InternalMessageInfo

type Leave struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Leave) Reset()         { *m = Leave{} }
func (m *Leave) String() string { return proto.CompactTextString(m) }
func (*Leave) ProtoMessage()    {}
func (*Leave) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbda7dd58e0f267b, []int{2}
}

func (m *Leave) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Leave.Unmarshal(m, b)
}
func (m *Leave) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Leave.Marshal(b, m, deterministic)
}
func (m *Leave) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Leave.Merge(m, src)
}
func (m *Leave) XXX_Size() int {
	return xxx_messageInfo_Leave.Size(m)
}
func (m *Leave) XXX_DiscardUnknown() {
	xxx_messageInfo_Leave.DiscardUnknown(m)
}

var xxx_messageInfo_Leave proto.InternalMessageInfo

func (m *Leave) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Leave) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type GetLeaveRequest struct {
	Offset               uint32   `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                uint32   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLeaveRequest) Reset()         { *m = GetLeaveRequest{} }
func (m *GetLeaveRequest) String() string { return proto.CompactTextString(m) }
func (*GetLeaveRequest) ProtoMessage()    {}
func (*GetLeaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbda7dd58e0f267b, []int{3}
}

func (m *GetLeaveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLeaveRequest.Unmarshal(m, b)
}
func (m *GetLeaveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLeaveRequest.Marshal(b, m, deterministic)
}
func (m *GetLeaveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLeaveRequest.Merge(m, src)
}
func (m *GetLeaveRequest) XXX_Size() int {
	return xxx_messageInfo_GetLeaveRequest.Size(m)
}
func (m *GetLeaveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLeaveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLeaveRequest proto.InternalMessageInfo

func (m *GetLeaveRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetLeaveRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetLeaveResponse struct {
	Leaves               []*Leave `protobuf:"bytes,1,rep,name=leaves,proto3" json:"leaves,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLeaveResponse) Reset()         { *m = GetLeaveResponse{} }
func (m *GetLeaveResponse) String() string { return proto.CompactTextString(m) }
func (*GetLeaveResponse) ProtoMessage()    {}
func (*GetLeaveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbda7dd58e0f267b, []int{4}
}

func (m *GetLeaveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLeaveResponse.Unmarshal(m, b)
}
func (m *GetLeaveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLeaveResponse.Marshal(b, m, deterministic)
}
func (m *GetLeaveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLeaveResponse.Merge(m, src)
}
func (m *GetLeaveResponse) XXX_Size() int {
	return xxx_messageInfo_GetLeaveResponse.Size(m)
}
func (m *GetLeaveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLeaveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLeaveResponse proto.InternalMessageInfo

func (m *GetLeaveResponse) GetLeaves() []*Leave {
	if m != nil {
		return m.Leaves
	}
	return nil
}

func init() {
	proto.RegisterType((*AddLeaveRequest)(nil), "guestbook.AddLeaveRequest")
	proto.RegisterType((*AddLeaveResponse)(nil), "guestbook.AddLeaveResponse")
	proto.RegisterType((*Leave)(nil), "guestbook.Leave")
	proto.RegisterType((*GetLeaveRequest)(nil), "guestbook.GetLeaveRequest")
	proto.RegisterType((*GetLeaveResponse)(nil), "guestbook.GetLeaveResponse")
}

func init() { proto.RegisterFile("guestbook.proto", fileDescriptor_fbda7dd58e0f267b) }

var fileDescriptor_fbda7dd58e0f267b = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x1b, 0x25, 0xd1, 0x8e, 0x94, 0x2c, 0x83, 0xc8, 0x52, 0x2f, 0x65, 0x0f, 0xd2, 0x53,
	0x0f, 0xf5, 0x20, 0x82, 0x20, 0x0a, 0xd2, 0x8b, 0xa7, 0xf5, 0x13, 0xf4, 0xcf, 0x54, 0x96, 0xa6,
	0x99, 0xda, 0x5d, 0xfb, 0x79, 0xfc, 0xa8, 0xb2, 0xbb, 0x89, 0x31, 0x21, 0xc7, 0xf7, 0x5e, 0xe6,
	0xc7, 0x7b, 0x1b, 0xc8, 0x3f, 0xbf, 0xc9, 0xba, 0x15, 0xf3, 0x6e, 0x76, 0x38, 0xb2, 0x63, 0x1c,
	0xfe, 0x19, 0xea, 0x11, 0xf2, 0x97, 0xcd, 0xe6, 0x9d, 0x96, 0x27, 0xd2, 0xf4, 0xe5, 0x6d, 0xbc,
	0x83, 0xb4, 0xf0, 0x5a, 0x26, 0x93, 0x64, 0x7a, 0x35, 0x17, 0xb3, 0xe6, 0x3c, 0x7e, 0x17, 0x63,
	0x85, 0x20, 0x9a, 0x53, 0x7b, 0xe0, 0xd2, 0x92, 0x7a, 0x80, 0x34, 0x18, 0x78, 0x0d, 0x29, 0xed,
	0x97, 0xa6, 0x08, 0x90, 0xa1, 0x8e, 0x02, 0x25, 0x5c, 0xac, 0xb9, 0x74, 0x54, 0x3a, 0x79, 0x16,
	0xfc, 0x5a, 0xaa, 0x67, 0xc8, 0x17, 0xe4, 0x5a, 0x3d, 0x6e, 0x20, 0xe3, 0xed, 0xd6, 0x92, 0x0b,
	0x8c, 0x91, 0xae, 0x94, 0x47, 0x17, 0x66, 0x6f, 0x22, 0x62, 0xa4, 0xa3, 0x50, 0x4f, 0x20, 0x1a,
	0x40, 0x6c, 0x83, 0x53, 0xc8, 0x42, 0x55, 0x2b, 0x93, 0xc9, 0x79, 0xef, 0x94, 0x2a, 0x9f, 0xff,
	0x24, 0x20, 0x16, 0x3e, 0x7b, 0x65, 0xde, 0x7d, 0xd0, 0xf1, 0x64, 0xd6, 0x84, 0x6f, 0x70, 0x59,
	0x0f, 0xc4, 0xf1, 0xbf, 0xd3, 0xce, 0x83, 0x8d, 0x6f, 0x7b, 0xb3, 0xea, 0x45, 0x06, 0x1e, 0x53,
	0x37, 0x6b, 0x61, 0x3a, 0x7b, 0x5b, 0x98, 0xee, 0x14, 0x35, 0x58, 0x65, 0xe1, 0xdf, 0xdd, 0xff,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x62, 0xad, 0x64, 0xce, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GuestBookServiceClient is the client API for GuestBookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GuestBookServiceClient interface {
	//添加留言
	AddLeave(ctx context.Context, in *AddLeaveRequest, opts ...grpc.CallOption) (*AddLeaveResponse, error)
	//查看留言
	GetLeave(ctx context.Context, in *GetLeaveRequest, opts ...grpc.CallOption) (*GetLeaveResponse, error)
}

type guestBookServiceClient struct {
	cc *grpc.ClientConn
}

func NewGuestBookServiceClient(cc *grpc.ClientConn) GuestBookServiceClient {
	return &guestBookServiceClient{cc}
}

func (c *guestBookServiceClient) AddLeave(ctx context.Context, in *AddLeaveRequest, opts ...grpc.CallOption) (*AddLeaveResponse, error) {
	out := new(AddLeaveResponse)
	err := c.cc.Invoke(ctx, "/guestbook.GuestBookService/AddLeave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guestBookServiceClient) GetLeave(ctx context.Context, in *GetLeaveRequest, opts ...grpc.CallOption) (*GetLeaveResponse, error) {
	out := new(GetLeaveResponse)
	err := c.cc.Invoke(ctx, "/guestbook.GuestBookService/GetLeave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GuestBookServiceServer is the server API for GuestBookService service.
type GuestBookServiceServer interface {
	//添加留言
	AddLeave(context.Context, *AddLeaveRequest) (*AddLeaveResponse, error)
	//查看留言
	GetLeave(context.Context, *GetLeaveRequest) (*GetLeaveResponse, error)
}

// UnimplementedGuestBookServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGuestBookServiceServer struct {
}

func (*UnimplementedGuestBookServiceServer) AddLeave(ctx context.Context, req *AddLeaveRequest) (*AddLeaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLeave not implemented")
}
func (*UnimplementedGuestBookServiceServer) GetLeave(ctx context.Context, req *GetLeaveRequest) (*GetLeaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeave not implemented")
}

func RegisterGuestBookServiceServer(s *grpc.Server, srv GuestBookServiceServer) {
	s.RegisterService(&_GuestBookService_serviceDesc, srv)
}

func _GuestBookService_AddLeave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLeaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestBookServiceServer).AddLeave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/guestbook.GuestBookService/AddLeave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestBookServiceServer).AddLeave(ctx, req.(*AddLeaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuestBookService_GetLeave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLeaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestBookServiceServer).GetLeave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/guestbook.GuestBookService/GetLeave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestBookServiceServer).GetLeave(ctx, req.(*GetLeaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GuestBookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "guestbook.GuestBookService",
	HandlerType: (*GuestBookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddLeave",
			Handler:    _GuestBookService_AddLeave_Handler,
		},
		{
			MethodName: "GetLeave",
			Handler:    _GuestBookService_GetLeave_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "guestbook.proto",
}