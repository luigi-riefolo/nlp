// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storer.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The representation of single user entry.
type Entry struct {
	// @inject_tag: csv:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" csv:"id"`
	// @inject_tag: csv:"name"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" csv:"name"`
	// @inject_tag: csv:"email"
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty" csv:"email"`
	// @inject_tag: csv:"mobile_number"
	MobileNumber         string   `protobuf:"bytes,4,opt,name=mobile_number,json=mobileNumber,proto3" json:"mobile_number,omitempty" csv:"mobile_number"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_472706382ed5c451, []int{0}
}

func (m *Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entry.Unmarshal(m, b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return xxx_messageInfo_Entry.Size(m)
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Entry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Entry) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Entry) GetMobileNumber() string {
	if m != nil {
		return m.MobileNumber
	}
	return ""
}

// StoreEntryRequest represents the
// request for the entry to be stored.
type StoreEntryRequest struct {
	Entry                *Entry   `protobuf:"bytes,2,opt,name=entry,proto3" json:"entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreEntryRequest) Reset()         { *m = StoreEntryRequest{} }
func (m *StoreEntryRequest) String() string { return proto.CompactTextString(m) }
func (*StoreEntryRequest) ProtoMessage()    {}
func (*StoreEntryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_472706382ed5c451, []int{1}
}

func (m *StoreEntryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreEntryRequest.Unmarshal(m, b)
}
func (m *StoreEntryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreEntryRequest.Marshal(b, m, deterministic)
}
func (m *StoreEntryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreEntryRequest.Merge(m, src)
}
func (m *StoreEntryRequest) XXX_Size() int {
	return xxx_messageInfo_StoreEntryRequest.Size(m)
}
func (m *StoreEntryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreEntryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StoreEntryRequest proto.InternalMessageInfo

func (m *StoreEntryRequest) GetEntry() *Entry {
	if m != nil {
		return m.Entry
	}
	return nil
}

func init() {
	proto.RegisterType((*Entry)(nil), "storer.Entry")
	proto.RegisterType((*StoreEntryRequest)(nil), "storer.StoreEntryRequest")
}

func init() { proto.RegisterFile("storer.proto", fileDescriptor_472706382ed5c451) }

var fileDescriptor_472706382ed5c451 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xcd, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xed, 0x9a, 0x04, 0x1c, 0x5b, 0xc1, 0x41, 0x64, 0xad, 0x17, 0x49, 0x2f, 0x9e, 0xb6,
	0x50, 0x2f, 0x9e, 0x85, 0x42, 0x4f, 0x1e, 0xe2, 0xcd, 0x8b, 0x64, 0xc9, 0xb4, 0x2c, 0xec, 0x47,
	0xdc, 0x6e, 0x0e, 0xf9, 0xef, 0x25, 0xb3, 0x06, 0x04, 0x6f, 0xf3, 0x7e, 0xf3, 0xf5, 0x1e, 0x2c,
	0xcf, 0x29, 0x44, 0x8a, 0xaa, 0x8f, 0x21, 0x05, 0xac, 0xb2, 0x5a, 0x3f, 0x9e, 0x42, 0x38, 0x59,
	0xda, 0x32, 0xd5, 0xc3, 0x71, 0x4b, 0xae, 0x4f, 0x63, 0x1e, 0xaa, 0x8f, 0x50, 0xee, 0x7d, 0x8a,
	0x23, 0xde, 0x80, 0x30, 0x9d, 0x5c, 0x3c, 0x2d, 0x9e, 0xaf, 0x1a, 0x61, 0x3a, 0x44, 0x28, 0x7c,
	0xeb, 0x48, 0x0a, 0x26, 0x5c, 0xe3, 0x1d, 0x94, 0xe4, 0x5a, 0x63, 0xe5, 0x25, 0xc3, 0x2c, 0x70,
	0x03, 0x2b, 0x17, 0xb4, 0xb1, 0xf4, 0xe5, 0x07, 0xa7, 0x29, 0xca, 0x82, 0xbb, 0xcb, 0x0c, 0xdf,
	0x99, 0xd5, 0xaf, 0x70, 0xfb, 0x31, 0xd9, 0xe1, 0x67, 0x0d, 0x7d, 0x0f, 0x74, 0x4e, 0xb8, 0x81,
	0x92, 0x26, 0xcd, 0x4f, 0xae, 0x77, 0x2b, 0xf5, 0xeb, 0x3f, 0x0f, 0xe5, 0xde, 0xae, 0x81, 0x8a,
	0x37, 0x23, 0x1e, 0xfe, 0xde, 0x38, 0xb4, 0xbe, 0xb3, 0x14, 0xf1, 0x61, 0x5e, 0xfa, 0x77, 0x7e,
	0x7d, 0xaf, 0x72, 0x72, 0x35, 0x27, 0x57, 0xfb, 0x29, 0x79, 0x7d, 0xf1, 0x56, 0x7c, 0x8a, 0x5e,
	0xeb, 0x8a, 0xf9, 0xcb, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x00, 0x18, 0x6f, 0x37, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StorerClient is the client API for Storer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StorerClient interface {
	// StoreEntryHandler stores the received entries.
	StoreEntryHandler(ctx context.Context, in *StoreEntryRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type storerClient struct {
	cc *grpc.ClientConn
}

func NewStorerClient(cc *grpc.ClientConn) StorerClient {
	return &storerClient{cc}
}

func (c *storerClient) StoreEntryHandler(ctx context.Context, in *StoreEntryRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/storer.Storer/StoreEntryHandler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorerServer is the server API for Storer service.
type StorerServer interface {
	// StoreEntryHandler stores the received entries.
	StoreEntryHandler(context.Context, *StoreEntryRequest) (*empty.Empty, error)
}

func RegisterStorerServer(s *grpc.Server, srv StorerServer) {
	s.RegisterService(&_Storer_serviceDesc, srv)
}

func _Storer_StoreEntryHandler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorerServer).StoreEntryHandler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storer.Storer/StoreEntryHandler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorerServer).StoreEntryHandler(ctx, req.(*StoreEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Storer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "storer.Storer",
	HandlerType: (*StorerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreEntryHandler",
			Handler:    _Storer_StoreEntryHandler_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storer.proto",
}
