// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lpwan-app-server/ui/server_info.proto

package ui

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ServerRegion int32

const (
	ServerRegion_NOT_DEFINED ServerRegion = 0
	ServerRegion_AVERAGE     ServerRegion = 1
	ServerRegion_RESTRICTED  ServerRegion = 2
)

var ServerRegion_name = map[int32]string{
	0: "NOT_DEFINED",
	1: "AVERAGE",
	2: "RESTRICTED",
}

var ServerRegion_value = map[string]int32{
	"NOT_DEFINED": 0,
	"AVERAGE":     1,
	"RESTRICTED":  2,
}

func (x ServerRegion) String() string {
	return proto.EnumName(ServerRegion_name, int32(x))
}

func (ServerRegion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9135f600200f7bdc, []int{0}
}

type GetAppserverVersionResponse struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAppserverVersionResponse) Reset()         { *m = GetAppserverVersionResponse{} }
func (m *GetAppserverVersionResponse) String() string { return proto.CompactTextString(m) }
func (*GetAppserverVersionResponse) ProtoMessage()    {}
func (*GetAppserverVersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9135f600200f7bdc, []int{0}
}

func (m *GetAppserverVersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAppserverVersionResponse.Unmarshal(m, b)
}
func (m *GetAppserverVersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAppserverVersionResponse.Marshal(b, m, deterministic)
}
func (m *GetAppserverVersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAppserverVersionResponse.Merge(m, src)
}
func (m *GetAppserverVersionResponse) XXX_Size() int {
	return xxx_messageInfo_GetAppserverVersionResponse.Size(m)
}
func (m *GetAppserverVersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAppserverVersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAppserverVersionResponse proto.InternalMessageInfo

func (m *GetAppserverVersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type GetM2MServerVersionResponse struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetM2MServerVersionResponse) Reset()         { *m = GetM2MServerVersionResponse{} }
func (m *GetM2MServerVersionResponse) String() string { return proto.CompactTextString(m) }
func (*GetM2MServerVersionResponse) ProtoMessage()    {}
func (*GetM2MServerVersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9135f600200f7bdc, []int{1}
}

func (m *GetM2MServerVersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetM2MServerVersionResponse.Unmarshal(m, b)
}
func (m *GetM2MServerVersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetM2MServerVersionResponse.Marshal(b, m, deterministic)
}
func (m *GetM2MServerVersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetM2MServerVersionResponse.Merge(m, src)
}
func (m *GetM2MServerVersionResponse) XXX_Size() int {
	return xxx_messageInfo_GetM2MServerVersionResponse.Size(m)
}
func (m *GetM2MServerVersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetM2MServerVersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetM2MServerVersionResponse proto.InternalMessageInfo

func (m *GetM2MServerVersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type GetServerRegionResponse struct {
	ServerRegion         string   `protobuf:"bytes,1,opt,name=server_region,json=serverRegion,proto3" json:"server_region,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServerRegionResponse) Reset()         { *m = GetServerRegionResponse{} }
func (m *GetServerRegionResponse) String() string { return proto.CompactTextString(m) }
func (*GetServerRegionResponse) ProtoMessage()    {}
func (*GetServerRegionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9135f600200f7bdc, []int{2}
}

func (m *GetServerRegionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServerRegionResponse.Unmarshal(m, b)
}
func (m *GetServerRegionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServerRegionResponse.Marshal(b, m, deterministic)
}
func (m *GetServerRegionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServerRegionResponse.Merge(m, src)
}
func (m *GetServerRegionResponse) XXX_Size() int {
	return xxx_messageInfo_GetServerRegionResponse.Size(m)
}
func (m *GetServerRegionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServerRegionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetServerRegionResponse proto.InternalMessageInfo

func (m *GetServerRegionResponse) GetServerRegion() string {
	if m != nil {
		return m.ServerRegion
	}
	return ""
}

func init() {
	proto.RegisterEnum("ui.ServerRegion", ServerRegion_name, ServerRegion_value)
	proto.RegisterType((*GetAppserverVersionResponse)(nil), "ui.GetAppserverVersionResponse")
	proto.RegisterType((*GetM2MServerVersionResponse)(nil), "ui.GetM2MServerVersionResponse")
	proto.RegisterType((*GetServerRegionResponse)(nil), "ui.GetServerRegionResponse")
}

func init() {
	proto.RegisterFile("lpwan-app-server/ui/server_info.proto", fileDescriptor_9135f600200f7bdc)
}

var fileDescriptor_9135f600200f7bdc = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x4b, 0x02, 0x41,
	0x18, 0xc6, 0x73, 0x0f, 0x49, 0xaf, 0x96, 0x3a, 0x41, 0x89, 0x46, 0xc9, 0x9a, 0x11, 0xc6, 0xee,
	0x82, 0x1d, 0xba, 0x44, 0x20, 0x39, 0x89, 0x07, 0x0d, 0x56, 0xf1, 0x2a, 0x6b, 0x8c, 0x32, 0xa0,
	0x33, 0xc3, 0xee, 0xac, 0x51, 0xdd, 0xfa, 0x0a, 0x7d, 0xb4, 0xbe, 0x42, 0x1f, 0xa3, 0x43, 0x38,
	0x33, 0x82, 0xf8, 0xa7, 0xe8, 0xb2, 0xbc, 0x3b, 0xfb, 0x3e, 0xf3, 0x63, 0x7f, 0x0f, 0x54, 0x26,
	0xe2, 0x39, 0x60, 0x4e, 0x20, 0x84, 0x13, 0x91, 0x70, 0x46, 0x42, 0x2f, 0xa6, 0x9e, 0x9e, 0x06,
	0x94, 0x8d, 0xb8, 0x2b, 0x42, 0x2e, 0x39, 0xb2, 0x62, 0x5a, 0x38, 0x19, 0x73, 0x3e, 0x9e, 0x10,
	0x2f, 0x10, 0xd4, 0x0b, 0x18, 0xe3, 0x32, 0x90, 0x94, 0xb3, 0x48, 0x6f, 0x14, 0x72, 0x64, 0x2a,
	0xe4, 0x8b, 0xa7, 0x9e, 0xfa, 0xc8, 0xbe, 0x81, 0x62, 0x93, 0xc8, 0xba, 0x10, 0xfa, 0xbe, 0x3e,
	0x09, 0x23, 0xca, 0x99, 0x4f, 0x22, 0xc1, 0x59, 0x44, 0x50, 0x1e, 0x92, 0x33, 0x7d, 0x94, 0x4f,
	0x94, 0x12, 0x97, 0x7b, 0xfe, 0xe2, 0xd5, 0x04, 0xdb, 0xb5, 0x76, 0xf7, 0x9f, 0xc1, 0x3b, 0x38,
	0x6e, 0x12, 0xa9, 0x53, 0x3e, 0x19, 0x2f, 0x87, 0xca, 0xb0, 0x6f, 0x7e, 0x2b, 0x54, 0x1f, 0x4c,
	0x34, 0x1d, 0x2d, 0x2d, 0x57, 0x6f, 0x21, 0xbd, 0x1c, 0x46, 0x19, 0x48, 0x75, 0x1e, 0x7b, 0x83,
	0x06, 0x7e, 0x68, 0x75, 0x70, 0x23, 0xbb, 0x83, 0x52, 0x90, 0xac, 0xf7, 0xb1, 0x5f, 0x6f, 0xe2,
	0x6c, 0x02, 0x1d, 0x00, 0xf8, 0xb8, 0xdb, 0xf3, 0x5b, 0xf7, 0x3d, 0xdc, 0xc8, 0x5a, 0xb5, 0x6f,
	0x0b, 0x72, 0x3a, 0xde, 0x62, 0x23, 0x3e, 0x9f, 0xe8, 0x13, 0x41, 0xaf, 0x70, 0xb8, 0xc1, 0x02,
	0x3a, 0x72, 0xb5, 0x4e, 0xed, 0x6a, 0x18, 0x8f, 0x5c, 0x3c, 0x57, 0x57, 0x38, 0x73, 0x63, 0xea,
	0xfe, 0xa2, 0xcd, 0xae, 0xbe, 0x7f, 0x7e, 0x7d, 0x58, 0xe7, 0xc8, 0x56, 0x45, 0xe8, 0x1d, 0x67,
	0x5e, 0x95, 0x17, 0x2c, 0x22, 0x8e, 0xf1, 0x81, 0xde, 0x14, 0x7b, 0x55, 0xe4, 0x9f, 0xec, 0x6d,
	0xe6, 0xed, 0x2b, 0xc5, 0xae, 0xa0, 0xf2, 0x1a, 0x7b, 0x5a, 0x9b, 0x3a, 0x2b, 0x70, 0x06, 0x99,
	0x95, 0x32, 0xb6, 0x82, 0x8b, 0x06, 0xbc, 0xa9, 0x39, 0xfb, 0x42, 0x41, 0x4b, 0xe8, 0x74, 0x0d,
	0x6a, 0x66, 0x5d, 0xe8, 0x70, 0x57, 0x5d, 0x7a, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0x08, 0x83,
	0x53, 0xce, 0xd3, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServerInfoServiceClient is the client API for ServerInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerInfoServiceClient interface {
	// get version
	GetAppserverVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetAppserverVersionResponse, error)
	GetM2MServerVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetM2MServerVersionResponse, error)
	GetServerRegion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetServerRegionResponse, error)
}

type serverInfoServiceClient struct {
	cc *grpc.ClientConn
}

func NewServerInfoServiceClient(cc *grpc.ClientConn) ServerInfoServiceClient {
	return &serverInfoServiceClient{cc}
}

func (c *serverInfoServiceClient) GetAppserverVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetAppserverVersionResponse, error) {
	out := new(GetAppserverVersionResponse)
	err := c.cc.Invoke(ctx, "/ui.ServerInfoService/GetAppserverVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverInfoServiceClient) GetM2MServerVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetM2MServerVersionResponse, error) {
	out := new(GetM2MServerVersionResponse)
	err := c.cc.Invoke(ctx, "/ui.ServerInfoService/GetM2MServerVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverInfoServiceClient) GetServerRegion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetServerRegionResponse, error) {
	out := new(GetServerRegionResponse)
	err := c.cc.Invoke(ctx, "/ui.ServerInfoService/GetServerRegion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerInfoServiceServer is the server API for ServerInfoService service.
type ServerInfoServiceServer interface {
	// get version
	GetAppserverVersion(context.Context, *empty.Empty) (*GetAppserverVersionResponse, error)
	GetM2MServerVersion(context.Context, *empty.Empty) (*GetM2MServerVersionResponse, error)
	GetServerRegion(context.Context, *empty.Empty) (*GetServerRegionResponse, error)
}

// UnimplementedServerInfoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServerInfoServiceServer struct {
}

func (*UnimplementedServerInfoServiceServer) GetAppserverVersion(ctx context.Context, req *empty.Empty) (*GetAppserverVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppserverVersion not implemented")
}
func (*UnimplementedServerInfoServiceServer) GetM2MServerVersion(ctx context.Context, req *empty.Empty) (*GetM2MServerVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetM2MServerVersion not implemented")
}
func (*UnimplementedServerInfoServiceServer) GetServerRegion(ctx context.Context, req *empty.Empty) (*GetServerRegionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerRegion not implemented")
}

func RegisterServerInfoServiceServer(s *grpc.Server, srv ServerInfoServiceServer) {
	s.RegisterService(&_ServerInfoService_serviceDesc, srv)
}

func _ServerInfoService_GetAppserverVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerInfoServiceServer).GetAppserverVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ui.ServerInfoService/GetAppserverVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerInfoServiceServer).GetAppserverVersion(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerInfoService_GetM2MServerVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerInfoServiceServer).GetM2MServerVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ui.ServerInfoService/GetM2MServerVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerInfoServiceServer).GetM2MServerVersion(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerInfoService_GetServerRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerInfoServiceServer).GetServerRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ui.ServerInfoService/GetServerRegion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerInfoServiceServer).GetServerRegion(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServerInfoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ui.ServerInfoService",
	HandlerType: (*ServerInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAppserverVersion",
			Handler:    _ServerInfoService_GetAppserverVersion_Handler,
		},
		{
			MethodName: "GetM2MServerVersion",
			Handler:    _ServerInfoService_GetM2MServerVersion_Handler,
		},
		{
			MethodName: "GetServerRegion",
			Handler:    _ServerInfoService_GetServerRegion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lpwan-app-server/ui/server_info.proto",
}
