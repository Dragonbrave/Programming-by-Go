package gym

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

// Person 谁在健身
type Person struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Actions              []string `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c3788ee4222e6c, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetActions() []string {
	if m != nil {
		return m.Actions
	}
	return nil
}

//结果
type Reply struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c3788ee4222e6c, []int{1}
}

func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Reply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Person)(nil), "lightweight.Person")
	proto.RegisterType((*Reply)(nil), "lightweight.Reply")
}

func init() {
	proto.RegisterFile("gym.proto", fileDescriptor_54c3788ee4222e6c)
}

var fileDescriptor_54c3788ee4222e6c = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0xaf, 0xcc, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xce, 0xc9, 0x4c, 0xcf, 0x28, 0x29, 0x4f, 0x05, 0x91,
	0x4a, 0x66, 0x5c, 0x6c, 0x01, 0xa9, 0x45, 0xc5, 0xf9, 0x79, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89,
	0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x04, 0x17, 0x7b, 0x62,
	0x72, 0x49, 0x66, 0x7e, 0x5e, 0xb1, 0x04, 0x93, 0x02, 0xb3, 0x06, 0x67, 0x10, 0x8c, 0xab, 0xa4,
	0xcb, 0xc5, 0x1a, 0x94, 0x5a, 0x90, 0x53, 0x09, 0xd2, 0x96, 0x9c, 0x9f, 0x02, 0xd1, 0xc6, 0x1a,
	0x04, 0x66, 0x0b, 0x09, 0x70, 0x31, 0xe7, 0x16, 0xa7, 0x4b, 0x30, 0x81, 0x4d, 0x02, 0x31, 0x8d,
	0x1c, 0xb8, 0x98, 0xdd, 0x2b, 0x73, 0x85, 0x2c, 0xb9, 0x78, 0x9c, 0xf2, 0x53, 0x2a, 0x9d, 0x4a,
	0x33, 0x73, 0x52, 0x32, 0xf3, 0xd2, 0x85, 0x84, 0xf5, 0x90, 0xdc, 0xa2, 0x07, 0x71, 0x88, 0x94,
	0x10, 0x8a, 0x20, 0xd8, 0x16, 0x25, 0x86, 0x24, 0x36, 0xb0, 0xe3, 0x8d, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xda, 0xd4, 0x71, 0xad, 0xc9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GymClient is the client API for Gym service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GymClient interface {
	BodyBuilding(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Reply, error)
}

type gymClient struct {
	cc grpc.ClientConnInterface
}

func NewGymClient(cc grpc.ClientConnInterface) GymClient {
	return &gymClient{cc}
}

func (c *gymClient) BodyBuilding(ctx context.Context, in *Person, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/lightweight.Gym/BodyBuilding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GymServer is the server API for Gym service.
type GymServer interface {
	BodyBuilding(context.Context, *Person) (*Reply, error)
}

// UnimplementedGymServer can be embedded to have forward compatible implementations.
type UnimplementedGymServer struct {
}

func (*UnimplementedGymServer) BodyBuilding(ctx context.Context, req *Person) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BodyBuilding not implemented")
}

func RegisterGymServer(s *grpc.Server, srv GymServer) {
	s.RegisterService(&_Gym_serviceDesc, srv)
}

func _Gym_BodyBuilding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Person)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GymServer).BodyBuilding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lightweight.Gym/BodyBuilding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GymServer).BodyBuilding(ctx, req.(*Person))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gym_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lightweight.Gym",
	HandlerType: (*GymServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BodyBuilding",
			Handler:    _Gym_BodyBuilding_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gym.proto",
}
