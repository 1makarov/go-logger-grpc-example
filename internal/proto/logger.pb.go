// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logger.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Action int32

const (
	Action_REGISTER Action = 0
	Action_LOGIN    Action = 1
	Action_CREATE   Action = 2
	Action_UPDATE   Action = 3
	Action_GET      Action = 4
	Action_DELETE   Action = 5
)

var Action_name = map[int32]string{
	0: "REGISTER",
	1: "LOGIN",
	2: "CREATE",
	3: "UPDATE",
	4: "GET",
	5: "DELETE",
}

var Action_value = map[string]int32{
	"REGISTER": 0,
	"LOGIN":    1,
	"CREATE":   2,
	"UPDATE":   3,
	"GET":      4,
	"DELETE":   5,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d43b7bfc6b6f7b16, []int{0}
}

type Entity int32

const (
	Entity_BOOK Entity = 0
	Entity_USER Entity = 1
)

var Entity_name = map[int32]string{
	0: "BOOK",
	1: "USER",
}

var Entity_value = map[string]int32{
	"BOOK": 0,
	"USER": 1,
}

func (x Entity) String() string {
	return proto.EnumName(Entity_name, int32(x))
}

func (Entity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d43b7bfc6b6f7b16, []int{1}
}

type ActionRequest struct {
	Action               Action                 `protobuf:"varint,1,opt,name=action,proto3,enum=api.Action" json:"action,omitempty"`
	Entity               Entity                 `protobuf:"varint,2,opt,name=entity,proto3,enum=api.Entity" json:"entity,omitempty"`
	UserId               int32                  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Timestamp            *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ActionRequest) Reset()         { *m = ActionRequest{} }
func (m *ActionRequest) String() string { return proto.CompactTextString(m) }
func (*ActionRequest) ProtoMessage()    {}
func (*ActionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d43b7bfc6b6f7b16, []int{0}
}

func (m *ActionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionRequest.Unmarshal(m, b)
}
func (m *ActionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionRequest.Marshal(b, m, deterministic)
}
func (m *ActionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionRequest.Merge(m, src)
}
func (m *ActionRequest) XXX_Size() int {
	return xxx_messageInfo_ActionRequest.Size(m)
}
func (m *ActionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ActionRequest proto.InternalMessageInfo

func (m *ActionRequest) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return Action_REGISTER
}

func (m *ActionRequest) GetEntity() Entity {
	if m != nil {
		return m.Entity
	}
	return Entity_BOOK
}

func (m *ActionRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ActionRequest) GetTimestamp() *timestamppb.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_d43b7bfc6b6f7b16, []int{1}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("api.Action", Action_name, Action_value)
	proto.RegisterEnum("api.Entity", Entity_name, Entity_value)
	proto.RegisterType((*ActionRequest)(nil), "api.ActionRequest")
	proto.RegisterType((*Empty)(nil), "api.empty")
}

func init() { proto.RegisterFile("logger.proto", fileDescriptor_d43b7bfc6b6f7b16) }

var fileDescriptor_d43b7bfc6b6f7b16 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x6b, 0xc2, 0x30,
	0x18, 0x86, 0x8d, 0xb5, 0x51, 0x3f, 0xdd, 0x08, 0xdf, 0x65, 0x45, 0x06, 0x13, 0xc7, 0x40, 0x3c,
	0x54, 0x70, 0x97, 0x5d, 0x75, 0x06, 0x91, 0x89, 0x8e, 0x58, 0x2f, 0xbb, 0x8c, 0x6e, 0xcd, 0x4a,
	0x40, 0x4d, 0xa7, 0xe9, 0xc1, 0xff, 0xb4, 0x1f, 0x39, 0x92, 0xd6, 0xcd, 0x53, 0xde, 0xbc, 0xdf,
	0x73, 0x78, 0x78, 0xa1, 0xbd, 0xd5, 0x69, 0x2a, 0x0f, 0x61, 0x76, 0xd0, 0x46, 0xa3, 0x17, 0x67,
	0xaa, 0x73, 0x97, 0x6a, 0x9d, 0x6e, 0xe5, 0xd0, 0x55, 0x1f, 0xf9, 0xd7, 0xd0, 0xa8, 0x9d, 0x3c,
	0x9a, 0x78, 0x97, 0x15, 0x54, 0xef, 0x87, 0xc0, 0xd5, 0xf8, 0xd3, 0x28, 0xbd, 0x17, 0xf2, 0x3b,
	0x97, 0x47, 0x83, 0xf7, 0x40, 0x63, 0x57, 0x04, 0xa4, 0x4b, 0xfa, 0xd7, 0xa3, 0x56, 0x18, 0x67,
	0x2a, 0x2c, 0x99, 0xf2, 0x64, 0x21, 0xb9, 0x37, 0xca, 0x9c, 0x82, 0xea, 0x05, 0xc4, 0x5d, 0x25,
	0xca, 0x13, 0xde, 0x40, 0x3d, 0x3f, 0xca, 0xc3, 0xbb, 0x4a, 0x02, 0xaf, 0x4b, 0xfa, 0xbe, 0xa0,
	0xf6, 0x3b, 0x4f, 0xf0, 0x09, 0x9a, 0x7f, 0x1e, 0x41, 0xad, 0x4b, 0xfa, 0xad, 0x51, 0x27, 0x2c,
	0x4c, 0xc3, 0xb3, 0x69, 0x18, 0x9d, 0x09, 0xf1, 0x0f, 0xf7, 0xea, 0xe0, 0xcb, 0x5d, 0x66, 0x4e,
	0x83, 0x25, 0xd0, 0x42, 0x09, 0xdb, 0xd0, 0x10, 0x7c, 0x36, 0x5f, 0x47, 0x5c, 0xb0, 0x0a, 0x36,
	0xc1, 0x5f, 0xac, 0x66, 0xf3, 0x25, 0x23, 0x08, 0x40, 0x9f, 0x05, 0x1f, 0x47, 0x9c, 0x55, 0x6d,
	0xde, 0xbc, 0x4e, 0x6d, 0xf6, 0xb0, 0x0e, 0xde, 0x8c, 0x47, 0xac, 0x66, 0xcb, 0x29, 0x5f, 0xf0,
	0x88, 0x33, 0x7f, 0x70, 0x0b, 0xb4, 0xb0, 0xc7, 0x06, 0xd4, 0x26, 0xab, 0xd5, 0x0b, 0xab, 0xd8,
	0xb4, 0x59, 0x73, 0xc1, 0xc8, 0x68, 0x08, 0x74, 0xe1, 0xb6, 0xc5, 0x07, 0xf0, 0xc6, 0x49, 0x82,
	0x78, 0x39, 0x4a, 0x31, 0x5c, 0x07, 0x5c, 0xe7, 0xf4, 0x26, 0x8d, 0x37, 0x5a, 0x6e, 0x4e, 0xdd,
	0xf3, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x1a, 0xd5, 0x13, 0x9d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoggerClient is the client API for Logger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoggerClient interface {
	Add(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Empty, error)
}

type loggerClient struct {
	cc *grpc.ClientConn
}

func NewLoggerClient(cc *grpc.ClientConn) LoggerClient {
	return &loggerClient{cc}
}

func (c *loggerClient) Add(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/api.Logger/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoggerServer is the server API for Logger service.
type LoggerServer interface {
	Add(context.Context, *ActionRequest) (*Empty, error)
}

// UnimplementedLoggerServer can be embedded to have forward compatible implementations.
type UnimplementedLoggerServer struct {
}

func (*UnimplementedLoggerServer) Add(ctx context.Context, req *ActionRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}

func RegisterLoggerServer(s *grpc.Server, srv LoggerServer) {
	s.RegisterService(&_Logger_serviceDesc, srv)
}

func _Logger_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggerServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Logger/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggerServer).Add(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Logger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Logger",
	HandlerType: (*LoggerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Logger_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logger.proto",
}
