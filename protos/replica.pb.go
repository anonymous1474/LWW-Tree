// Code generated by protoc-gen-go. DO NOT EDIT.
// source: replica.proto

package protos

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

type Request struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e84aa831fb48ea1, []int{0}
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

func (m *Request) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Response struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Clock                int32    `protobuf:"varint,3,opt,name=clock,proto3" json:"clock,omitempty"`
	ID                   int32    `protobuf:"varint,4,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e84aa831fb48ea1, []int{1}
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

func (m *Response) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Response) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Response) GetClock() int32 {
	if m != nil {
		return m.Clock
	}
	return 0
}

func (m *Response) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type Verify struct {
	Vertices             []string `protobuf:"bytes,1,rep,name=vertices,proto3" json:"vertices,omitempty"`
	ID                   int32    `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Verify) Reset()         { *m = Verify{} }
func (m *Verify) String() string { return proto.CompactTextString(m) }
func (*Verify) ProtoMessage()    {}
func (*Verify) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e84aa831fb48ea1, []int{2}
}

func (m *Verify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Verify.Unmarshal(m, b)
}
func (m *Verify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Verify.Marshal(b, m, deterministic)
}
func (m *Verify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Verify.Merge(m, src)
}
func (m *Verify) XXX_Size() int {
	return xxx_messageInfo_Verify.Size(m)
}
func (m *Verify) XXX_DiscardUnknown() {
	xxx_messageInfo_Verify.DiscardUnknown(m)
}

var xxx_messageInfo_Verify proto.InternalMessageInfo

func (m *Verify) GetVertices() []string {
	if m != nil {
		return m.Vertices
	}
	return nil
}

func (m *Verify) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e84aa831fb48ea1, []int{3}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Request)(nil), "protos.Request")
	proto.RegisterType((*Response)(nil), "protos.Response")
	proto.RegisterType((*Verify)(nil), "protos.Verify")
	proto.RegisterType((*Void)(nil), "protos.Void")
}

func init() { proto.RegisterFile("replica.proto", fileDescriptor_1e84aa831fb48ea1) }

var fileDescriptor_1e84aa831fb48ea1 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xc1, 0x4f, 0xc2, 0x30,
	0x14, 0xc6, 0xd9, 0x06, 0x13, 0x9e, 0x0e, 0x49, 0xe3, 0xa1, 0xee, 0x44, 0x7a, 0xe2, 0x22, 0x31,
	0x68, 0x3c, 0x79, 0x31, 0x12, 0x0d, 0xd7, 0x92, 0x10, 0xaf, 0xb3, 0x7b, 0x86, 0x97, 0x4d, 0x3a,
	0xdb, 0x32, 0xc3, 0x7f, 0x6f, 0xb6, 0xb2, 0x79, 0x6a, 0x7f, 0x5f, 0xfb, 0x7d, 0x79, 0xdf, 0x83,
	0xc4, 0x60, 0x55, 0x92, 0xca, 0x96, 0x95, 0xd1, 0x4e, 0xb3, 0xb8, 0x3d, 0xac, 0xb8, 0x85, 0x0b,
	0x89, 0x3f, 0x47, 0xb4, 0x8e, 0x4d, 0x21, 0xa4, 0x9c, 0x07, 0xf3, 0x60, 0x31, 0x92, 0x21, 0xe5,
	0xe2, 0x03, 0xc6, 0x12, 0x6d, 0xa5, 0x0f, 0x16, 0xd9, 0x0c, 0xa2, 0x02, 0x4f, 0xed, 0xe3, 0x44,
	0x36, 0x57, 0x76, 0x03, 0xa3, 0x3a, 0x2b, 0x8f, 0xc8, 0xc3, 0x56, 0xf3, 0xd0, 0xa8, 0xaa, 0xd4,
	0xaa, 0xe0, 0x51, 0x1b, 0xe3, 0xa1, 0x49, 0xde, 0xac, 0xf9, 0xd0, 0x27, 0x6f, 0xd6, 0xe2, 0x11,
	0xe2, 0x1d, 0x1a, 0xfa, 0x3a, 0xb1, 0x14, 0xc6, 0x35, 0x1a, 0x47, 0x0a, 0x2d, 0x0f, 0xe6, 0xd1,
	0x62, 0x22, 0x7b, 0x3e, 0xbb, 0xc2, 0xde, 0x15, 0xc3, 0x70, 0xa7, 0x29, 0x5f, 0xbd, 0x43, 0xb2,
	0x75, 0x06, 0xb3, 0xef, 0x2d, 0x9a, 0x9a, 0x14, 0xb2, 0x27, 0x48, 0xde, 0xd0, 0xa9, 0x7d, 0x3f,
	0xed, 0xb5, 0x2f, 0x69, 0x97, 0xe7, 0x6a, 0xe9, 0xec, 0x5f, 0xf0, 0x5f, 0xc4, 0xe0, 0x3e, 0x58,
	0x3d, 0xc3, 0xe5, 0xeb, 0x3e, 0x73, 0x5d, 0xcc, 0x5d, 0x83, 0xa8, 0x8a, 0x97, 0x83, 0xfd, 0x45,
	0xc3, 0xa6, 0x9d, 0xc7, 0x8f, 0x9a, 0x5e, 0xf5, 0xac, 0x29, 0x17, 0x83, 0x4f, 0xbf, 0xc1, 0x87,
	0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x9e, 0x1f, 0x67, 0x59, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StreamServiceClient is the client API for StreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamServiceClient interface {
	FetchResponse(ctx context.Context, in *Request, opts ...grpc.CallOption) (StreamService_FetchResponseClient, error)
}

type streamServiceClient struct {
	cc *grpc.ClientConn
}

func NewStreamServiceClient(cc *grpc.ClientConn) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) FetchResponse(ctx context.Context, in *Request, opts ...grpc.CallOption) (StreamService_FetchResponseClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[0], "/protos.StreamService/FetchResponse", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceFetchResponseClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamService_FetchResponseClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type streamServiceFetchResponseClient struct {
	grpc.ClientStream
}

func (x *streamServiceFetchResponseClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServiceServer is the server API for StreamService service.
type StreamServiceServer interface {
	FetchResponse(*Request, StreamService_FetchResponseServer) error
}

// UnimplementedStreamServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStreamServiceServer struct {
}

func (*UnimplementedStreamServiceServer) FetchResponse(req *Request, srv StreamService_FetchResponseServer) error {
	return status.Errorf(codes.Unimplemented, "method FetchResponse not implemented")
}

func RegisterStreamServiceServer(s *grpc.Server, srv StreamServiceServer) {
	s.RegisterService(&_StreamService_serviceDesc, srv)
}

func _StreamService_FetchResponse_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServiceServer).FetchResponse(m, &streamServiceFetchResponseServer{stream})
}

type StreamService_FetchResponseServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type streamServiceFetchResponseServer struct {
	grpc.ServerStream
}

func (x *streamServiceFetchResponseServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

var _StreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FetchResponse",
			Handler:       _StreamService_FetchResponse_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "replica.proto",
}

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	CheckAnswer(ctx context.Context, in *Verify, opts ...grpc.CallOption) (*Void, error)
}

type chatServiceClient struct {
	cc *grpc.ClientConn
}

func NewChatServiceClient(cc *grpc.ClientConn) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) CheckAnswer(ctx context.Context, in *Verify, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.ChatService/CheckAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	CheckAnswer(context.Context, *Verify) (*Void, error)
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) CheckAnswer(ctx context.Context, req *Verify) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAnswer not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_CheckAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Verify)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CheckAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ChatService/CheckAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CheckAnswer(ctx, req.(*Verify))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckAnswer",
			Handler:    _ChatService_CheckAnswer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "replica.proto",
}