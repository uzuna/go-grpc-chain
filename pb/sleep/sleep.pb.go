// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sleep/sleep.proto

package sleep

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type SleepRequest struct {
	Duration             *duration.Duration `protobuf:"bytes,1,opt,name=duration,proto3" json:"duration,omitempty"`
	DummyBytes           int32              `protobuf:"varint,2,opt,name=dummy_bytes,json=dummyBytes,proto3" json:"dummy_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SleepRequest) Reset()         { *m = SleepRequest{} }
func (m *SleepRequest) String() string { return proto.CompactTextString(m) }
func (*SleepRequest) ProtoMessage()    {}
func (*SleepRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_78f2de77f7653b60, []int{0}
}

func (m *SleepRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SleepRequest.Unmarshal(m, b)
}
func (m *SleepRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SleepRequest.Marshal(b, m, deterministic)
}
func (m *SleepRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SleepRequest.Merge(m, src)
}
func (m *SleepRequest) XXX_Size() int {
	return xxx_messageInfo_SleepRequest.Size(m)
}
func (m *SleepRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SleepRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SleepRequest proto.InternalMessageInfo

func (m *SleepRequest) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *SleepRequest) GetDummyBytes() int32 {
	if m != nil {
		return m.DummyBytes
	}
	return 0
}

type SleepResponce struct {
	Duration             *duration.Duration `protobuf:"bytes,1,opt,name=duration,proto3" json:"duration,omitempty"`
	DummyBytes           int32              `protobuf:"varint,2,opt,name=dummy_bytes,json=dummyBytes,proto3" json:"dummy_bytes,omitempty"`
	Data                 []byte             `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SleepResponce) Reset()         { *m = SleepResponce{} }
func (m *SleepResponce) String() string { return proto.CompactTextString(m) }
func (*SleepResponce) ProtoMessage()    {}
func (*SleepResponce) Descriptor() ([]byte, []int) {
	return fileDescriptor_78f2de77f7653b60, []int{1}
}

func (m *SleepResponce) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SleepResponce.Unmarshal(m, b)
}
func (m *SleepResponce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SleepResponce.Marshal(b, m, deterministic)
}
func (m *SleepResponce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SleepResponce.Merge(m, src)
}
func (m *SleepResponce) XXX_Size() int {
	return xxx_messageInfo_SleepResponce.Size(m)
}
func (m *SleepResponce) XXX_DiscardUnknown() {
	xxx_messageInfo_SleepResponce.DiscardUnknown(m)
}

var xxx_messageInfo_SleepResponce proto.InternalMessageInfo

func (m *SleepResponce) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *SleepResponce) GetDummyBytes() int32 {
	if m != nil {
		return m.DummyBytes
	}
	return 0
}

func (m *SleepResponce) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*SleepRequest)(nil), "sleep.SleepRequest")
	proto.RegisterType((*SleepResponce)(nil), "sleep.SleepResponce")
}

func init() { proto.RegisterFile("sleep/sleep.proto", fileDescriptor_78f2de77f7653b60) }

var fileDescriptor_78f2de77f7653b60 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0xce, 0x49, 0x4d,
	0x2d, 0xd0, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xac, 0x60, 0x8e, 0x94, 0x5c,
	0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x58, 0x30, 0xa9, 0x34, 0x4d, 0x3f, 0xa5, 0xb4, 0x28,
	0xb1, 0x24, 0x33, 0x3f, 0x0f, 0xa2, 0x4c, 0x29, 0x8d, 0x8b, 0x27, 0x18, 0xa4, 0x30, 0x28, 0xb5,
	0xb0, 0x34, 0xb5, 0xb8, 0x44, 0xc8, 0x94, 0x8b, 0x03, 0xa6, 0x42, 0x82, 0x51, 0x81, 0x51, 0x83,
	0xdb, 0x48, 0x52, 0x0f, 0x62, 0x84, 0x1e, 0xcc, 0x08, 0x3d, 0x17, 0xa8, 0x82, 0x20, 0xb8, 0x52,
	0x21, 0x79, 0x2e, 0xee, 0x94, 0xd2, 0xdc, 0xdc, 0xca, 0xf8, 0xa4, 0xca, 0x92, 0xd4, 0x62, 0x09,
	0x26, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x2e, 0xb0, 0x90, 0x13, 0x48, 0x44, 0xa9, 0x9a, 0x8b, 0x17,
	0x6a, 0x4f, 0x71, 0x41, 0x7e, 0x5e, 0x72, 0x2a, 0xad, 0x2c, 0x12, 0x12, 0xe2, 0x62, 0x49, 0x49,
	0x2c, 0x49, 0x94, 0x60, 0x56, 0x60, 0xd4, 0xe0, 0x09, 0x02, 0xb3, 0x8d, 0x5c, 0xa0, 0x9e, 0x0c,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x32, 0xe1, 0x62, 0x05, 0xf3, 0x85, 0x84, 0xf5, 0x20,
	0x41, 0x86, 0x1c, 0x04, 0x52, 0x22, 0xa8, 0x82, 0x10, 0xf7, 0x2a, 0x31, 0x24, 0xb1, 0x81, 0xdd,
	0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x78, 0xda, 0x52, 0x6d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SleepServiceClient is the client API for SleepService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SleepServiceClient interface {
	Sleep(ctx context.Context, in *SleepRequest, opts ...grpc.CallOption) (*SleepResponce, error)
}

type sleepServiceClient struct {
	cc *grpc.ClientConn
}

func NewSleepServiceClient(cc *grpc.ClientConn) SleepServiceClient {
	return &sleepServiceClient{cc}
}

func (c *sleepServiceClient) Sleep(ctx context.Context, in *SleepRequest, opts ...grpc.CallOption) (*SleepResponce, error) {
	out := new(SleepResponce)
	err := c.cc.Invoke(ctx, "/sleep.SleepService/Sleep", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SleepServiceServer is the server API for SleepService service.
type SleepServiceServer interface {
	Sleep(context.Context, *SleepRequest) (*SleepResponce, error)
}

// UnimplementedSleepServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSleepServiceServer struct {
}

func (*UnimplementedSleepServiceServer) Sleep(ctx context.Context, req *SleepRequest) (*SleepResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sleep not implemented")
}

func RegisterSleepServiceServer(s *grpc.Server, srv SleepServiceServer) {
	s.RegisterService(&_SleepService_serviceDesc, srv)
}

func _SleepService_Sleep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SleepRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SleepServiceServer).Sleep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sleep.SleepService/Sleep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SleepServiceServer).Sleep(ctx, req.(*SleepRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SleepService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sleep.SleepService",
	HandlerType: (*SleepServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sleep",
			Handler:    _SleepService_Sleep_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sleep/sleep.proto",
}