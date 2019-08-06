// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/balance/balance.proto

package balance

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Request struct {
	AccountId            string   `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_be989d5aed6f6a6c, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

type HealthResponse struct {
	Health               string   `protobuf:"bytes,1,opt,name=health,proto3" json:"health,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthResponse) Reset()         { *m = HealthResponse{} }
func (m *HealthResponse) String() string { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()    {}
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_be989d5aed6f6a6c, []int{1}
}

func (m *HealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthResponse.Unmarshal(m, b)
}
func (m *HealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthResponse.Marshal(b, m, deterministic)
}
func (m *HealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthResponse.Merge(m, src)
}
func (m *HealthResponse) XXX_Size() int {
	return xxx_messageInfo_HealthResponse.Size(m)
}
func (m *HealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthResponse proto.InternalMessageInfo

func (m *HealthResponse) GetHealth() string {
	if m != nil {
		return m.Health
	}
	return ""
}

type Response struct {
	Completed            bool     `protobuf:"varint,1,opt,name=completed,proto3" json:"completed,omitempty"`
	Amount               int64    `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency             string   `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_be989d5aed6f6a6c, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

func (m *Response) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Response) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "balance.Request")
	proto.RegisterType((*HealthResponse)(nil), "balance.HealthResponse")
	proto.RegisterType((*Response)(nil), "balance.Response")
}

func init() { proto.RegisterFile("proto/balance/balance.proto", fileDescriptor_be989d5aed6f6a6c) }

var fileDescriptor_be989d5aed6f6a6c = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x4f, 0x4a, 0xc4, 0x30,
	0x14, 0xc6, 0xe9, 0x0c, 0xce, 0x9f, 0xb7, 0x10, 0x27, 0xc8, 0x58, 0x32, 0x15, 0x86, 0xac, 0x8a,
	0x8b, 0x16, 0x75, 0xe7, 0x52, 0x11, 0xd4, 0x65, 0xc1, 0x9d, 0x20, 0x69, 0xe6, 0x39, 0x2d, 0xb6,
	0x49, 0x6d, 0x53, 0x61, 0xb6, 0x5e, 0xc1, 0x6b, 0x78, 0x1b, 0xaf, 0xe0, 0x41, 0xa4, 0x49, 0x5a,
	0xc7, 0x55, 0xf9, 0xbe, 0xf7, 0xfa, 0x23, 0xef, 0x07, 0xab, 0xaa, 0x56, 0x5a, 0xc5, 0x29, 0x2f,
	0xb8, 0x14, 0xd8, 0x7f, 0x23, 0xd3, 0x92, 0xa9, 0x8b, 0x74, 0xb5, 0x55, 0x6a, 0x5b, 0x60, 0x6c,
	0xea, 0xb4, 0x7d, 0x89, 0xb1, 0xac, 0xf4, 0xce, 0x6e, 0xd1, 0xc0, 0x0d, 0x79, 0x95, 0xc7, 0x5c,
	0x4a, 0xa5, 0xb9, 0xce, 0x95, 0x6c, 0xec, 0x94, 0x85, 0x30, 0x4d, 0xf0, 0xad, 0xc5, 0x46, 0x93,
	0x53, 0x00, 0x2e, 0x84, 0x6a, 0xa5, 0x7e, 0xce, 0x37, 0xbe, 0xb7, 0xf6, 0xc2, 0x79, 0x32, 0x77,
	0xcd, 0xfd, 0x86, 0x85, 0x70, 0x78, 0x87, 0xbc, 0xd0, 0x59, 0x82, 0x4d, 0xa5, 0x64, 0x83, 0x64,
	0x09, 0x93, 0xcc, 0x34, 0x6e, 0xd9, 0x25, 0xf6, 0x04, 0xb3, 0x61, 0x27, 0x80, 0xb9, 0x50, 0x65,
	0x55, 0xa0, 0x46, 0xcb, 0x9c, 0x25, 0x7f, 0x45, 0x47, 0xe0, 0x65, 0xc7, 0xf7, 0x47, 0x6b, 0x2f,
	0x1c, 0x27, 0x2e, 0x11, 0x0a, 0x33, 0xd1, 0xd6, 0x35, 0x4a, 0xb1, 0xf3, 0xc7, 0x86, 0x3d, 0xe4,
	0x8b, 0x2f, 0x0f, 0xa6, 0xd7, 0xf6, 0x70, 0xf2, 0x00, 0x07, 0x37, 0x19, 0x8a, 0x57, 0x72, 0x14,
	0xf5, 0x6a, 0xdc, 0x35, 0x74, 0xb1, 0xd7, 0xd8, 0xb7, 0xb0, 0xe0, 0xe3, 0xfb, 0xe7, 0x73, 0xb4,
	0x64, 0x8b, 0xf8, 0xfd, 0x7c, 0x50, 0x2a, 0xba, 0xff, 0xaf, 0xbc, 0x33, 0xf2, 0x08, 0x13, 0x7b,
	0x1f, 0x59, 0x46, 0x56, 0x59, 0xd4, 0xfb, 0x8c, 0x6e, 0x3b, 0x9f, 0xf4, 0x64, 0x40, 0xfe, 0x17,
	0xc1, 0xa8, 0x01, 0x1f, 0x13, 0xb2, 0x0f, 0xb6, 0x32, 0xd2, 0x89, 0x81, 0x5c, 0xfe, 0x06, 0x00,
	0x00, 0xff, 0xff, 0x59, 0x6f, 0x04, 0xb5, 0xca, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BalanceClient is the client API for Balance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BalanceClient interface {
	Check(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthResponse, error)
}

type balanceClient struct {
	cc *grpc.ClientConn
}

func NewBalanceClient(cc *grpc.ClientConn) BalanceClient {
	return &balanceClient{cc}
}

func (c *balanceClient) Check(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/balance.Balance/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balanceClient) Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/balance.Balance/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BalanceServer is the server API for Balance service.
type BalanceServer interface {
	Check(context.Context, *Request) (*Response, error)
	Health(context.Context, *empty.Empty) (*HealthResponse, error)
}

// UnimplementedBalanceServer can be embedded to have forward compatible implementations.
type UnimplementedBalanceServer struct {
}

func (*UnimplementedBalanceServer) Check(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (*UnimplementedBalanceServer) Health(ctx context.Context, req *empty.Empty) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}

func RegisterBalanceServer(s *grpc.Server, srv BalanceServer) {
	s.RegisterService(&_Balance_serviceDesc, srv)
}

func _Balance_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/balance.Balance/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceServer).Check(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Balance_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/balance.Balance/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceServer).Health(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Balance_serviceDesc = grpc.ServiceDesc{
	ServiceName: "balance.Balance",
	HandlerType: (*BalanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Balance_Check_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _Balance_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/balance/balance.proto",
}
