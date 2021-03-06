// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lpwan-app-server/mxprotocol-server/gateway_item.proto

package dev

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type GatewayStruct struct {
	Mac                  string               `protobuf:"bytes,1,opt,name=mac,proto3" json:"mac,omitempty"`
	OrgId                int64                `protobuf:"varint,2,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	Description          string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Name                 string               `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GatewayStruct) Reset()         { *m = GatewayStruct{} }
func (m *GatewayStruct) String() string { return proto.CompactTextString(m) }
func (*GatewayStruct) ProtoMessage()    {}
func (*GatewayStruct) Descriptor() ([]byte, []int) {
	return fileDescriptor_d61caa2a00c8122e, []int{0}
}

func (m *GatewayStruct) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayStruct.Unmarshal(m, b)
}
func (m *GatewayStruct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayStruct.Marshal(b, m, deterministic)
}
func (m *GatewayStruct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayStruct.Merge(m, src)
}
func (m *GatewayStruct) XXX_Size() int {
	return xxx_messageInfo_GatewayStruct.Size(m)
}
func (m *GatewayStruct) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayStruct.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayStruct proto.InternalMessageInfo

func (m *GatewayStruct) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *GatewayStruct) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *GatewayStruct) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GatewayStruct) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GatewayStruct) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type GetGatewayByMacRequest struct {
	Mac                  string   `protobuf:"bytes,1,opt,name=mac,proto3" json:"mac,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGatewayByMacRequest) Reset()         { *m = GetGatewayByMacRequest{} }
func (m *GetGatewayByMacRequest) String() string { return proto.CompactTextString(m) }
func (*GetGatewayByMacRequest) ProtoMessage()    {}
func (*GetGatewayByMacRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d61caa2a00c8122e, []int{1}
}

func (m *GetGatewayByMacRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGatewayByMacRequest.Unmarshal(m, b)
}
func (m *GetGatewayByMacRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGatewayByMacRequest.Marshal(b, m, deterministic)
}
func (m *GetGatewayByMacRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGatewayByMacRequest.Merge(m, src)
}
func (m *GetGatewayByMacRequest) XXX_Size() int {
	return xxx_messageInfo_GetGatewayByMacRequest.Size(m)
}
func (m *GetGatewayByMacRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGatewayByMacRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGatewayByMacRequest proto.InternalMessageInfo

func (m *GetGatewayByMacRequest) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

type GetGatewayByMacResponse struct {
	OrgId                int64          `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	GwProfile            *GatewayStruct `protobuf:"bytes,2,opt,name=gw_profile,json=gwProfile,proto3" json:"gw_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetGatewayByMacResponse) Reset()         { *m = GetGatewayByMacResponse{} }
func (m *GetGatewayByMacResponse) String() string { return proto.CompactTextString(m) }
func (*GetGatewayByMacResponse) ProtoMessage()    {}
func (*GetGatewayByMacResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d61caa2a00c8122e, []int{2}
}

func (m *GetGatewayByMacResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGatewayByMacResponse.Unmarshal(m, b)
}
func (m *GetGatewayByMacResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGatewayByMacResponse.Marshal(b, m, deterministic)
}
func (m *GetGatewayByMacResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGatewayByMacResponse.Merge(m, src)
}
func (m *GetGatewayByMacResponse) XXX_Size() int {
	return xxx_messageInfo_GetGatewayByMacResponse.Size(m)
}
func (m *GetGatewayByMacResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGatewayByMacResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetGatewayByMacResponse proto.InternalMessageInfo

func (m *GetGatewayByMacResponse) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *GetGatewayByMacResponse) GetGwProfile() *GatewayStruct {
	if m != nil {
		return m.GwProfile
	}
	return nil
}

type GetGatewayMacListResponse struct {
	GatewayMac           []string `protobuf:"bytes,1,rep,name=gateway_mac,json=gatewayMac,proto3" json:"gateway_mac,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGatewayMacListResponse) Reset()         { *m = GetGatewayMacListResponse{} }
func (m *GetGatewayMacListResponse) String() string { return proto.CompactTextString(m) }
func (*GetGatewayMacListResponse) ProtoMessage()    {}
func (*GetGatewayMacListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d61caa2a00c8122e, []int{3}
}

func (m *GetGatewayMacListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGatewayMacListResponse.Unmarshal(m, b)
}
func (m *GetGatewayMacListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGatewayMacListResponse.Marshal(b, m, deterministic)
}
func (m *GetGatewayMacListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGatewayMacListResponse.Merge(m, src)
}
func (m *GetGatewayMacListResponse) XXX_Size() int {
	return xxx_messageInfo_GetGatewayMacListResponse.Size(m)
}
func (m *GetGatewayMacListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGatewayMacListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetGatewayMacListResponse proto.InternalMessageInfo

func (m *GetGatewayMacListResponse) GetGatewayMac() []string {
	if m != nil {
		return m.GatewayMac
	}
	return nil
}

func init() {
	proto.RegisterType((*GatewayStruct)(nil), "dev.GatewayStruct")
	proto.RegisterType((*GetGatewayByMacRequest)(nil), "dev.GetGatewayByMacRequest")
	proto.RegisterType((*GetGatewayByMacResponse)(nil), "dev.GetGatewayByMacResponse")
	proto.RegisterType((*GetGatewayMacListResponse)(nil), "dev.GetGatewayMacListResponse")
}

func init() {
	proto.RegisterFile("lpwan-app-server/mxprotocol-server/gateway_item.proto", fileDescriptor_d61caa2a00c8122e)
}

var fileDescriptor_d61caa2a00c8122e = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x0b, 0xd3, 0x40,
	0x10, 0xc5, 0x89, 0x69, 0x0b, 0x99, 0xe0, 0xbf, 0x05, 0x6b, 0x4c, 0xc5, 0x86, 0x9c, 0x82, 0xd0,
	0x14, 0x2b, 0x1e, 0x04, 0x2f, 0x0a, 0x52, 0xc4, 0x16, 0x24, 0x7a, 0x0f, 0xdb, 0xcd, 0x74, 0x59,
	0x48, 0xb2, 0xeb, 0x66, 0x9b, 0xd8, 0x8f, 0x24, 0x7e, 0x49, 0x69, 0xd2, 0xd4, 0x9a, 0xf6, 0x36,
	0xcc, 0x7b, 0x99, 0xf9, 0xbd, 0xc9, 0xc2, 0xbb, 0x5c, 0x35, 0xb4, 0x5c, 0x50, 0xa5, 0x16, 0x15,
	0xea, 0x1a, 0xf5, 0xb2, 0xf8, 0xa5, 0xb4, 0x34, 0x92, 0xc9, 0xbc, 0xef, 0x70, 0x6a, 0xb0, 0xa1,
	0xc7, 0x54, 0x18, 0x2c, 0xe2, 0x56, 0x24, 0x76, 0x86, 0xb5, 0x3f, 0xe7, 0x52, 0xf2, 0x1c, 0x97,
	0x6d, 0x6b, 0x77, 0xd8, 0x2f, 0x8d, 0x28, 0xb0, 0x32, 0xb4, 0x50, 0x9d, 0xcb, 0x9f, 0x0d, 0x0d,
	0x58, 0x28, 0x73, 0xec, 0xc4, 0xf0, 0xb7, 0x05, 0x0f, 0xd7, 0xdd, 0xe4, 0xef, 0x46, 0x1f, 0x98,
	0x21, 0x4f, 0xc0, 0x2e, 0x28, 0xf3, 0xac, 0xc0, 0x8a, 0x9c, 0xe4, 0x54, 0x92, 0x67, 0x30, 0x91,
	0x9a, 0xa7, 0x22, 0xf3, 0x1e, 0x04, 0x56, 0x64, 0x27, 0x63, 0xa9, 0xf9, 0x97, 0x8c, 0x04, 0xe0,
	0x66, 0x58, 0x31, 0x2d, 0x94, 0x11, 0xb2, 0xf4, 0xec, 0xf6, 0x83, 0xeb, 0x16, 0x21, 0x30, 0x2a,
	0x69, 0x81, 0xde, 0xa8, 0x95, 0xda, 0x9a, 0xbc, 0x07, 0x60, 0x1a, 0xa9, 0xc1, 0x2c, 0xa5, 0xc6,
	0x1b, 0x07, 0x56, 0xe4, 0xae, 0xfc, 0xb8, 0x43, 0x8c, 0x7b, 0xc4, 0xf8, 0x47, 0x9f, 0x21, 0x71,
	0xce, 0xee, 0x8f, 0x26, 0x7c, 0x0d, 0xd3, 0x35, 0x9a, 0x33, 0xed, 0xa7, 0xe3, 0x96, 0xb2, 0x04,
	0x7f, 0x1e, 0xb0, 0xba, 0xc3, 0x1c, 0x32, 0x78, 0x7e, 0xe3, 0xad, 0x94, 0x2c, 0x2b, 0xbc, 0x8a,
	0x63, 0x5d, 0xc7, 0x79, 0x03, 0xc0, 0x9b, 0x54, 0x69, 0xb9, 0x17, 0x39, 0xb6, 0x49, 0xdd, 0x15,
	0x89, 0x33, 0xac, 0xe3, 0xff, 0xee, 0x93, 0x38, 0xbc, 0xf9, 0xd6, 0x99, 0xc2, 0x0f, 0xf0, 0xe2,
	0xdf, 0x92, 0x2d, 0x65, 0x1b, 0x51, 0x99, 0xcb, 0x9a, 0x39, 0xb8, 0xfd, 0x2f, 0xeb, 0xd8, 0xec,
	0xc8, 0x49, 0x80, 0x5f, 0xcc, 0xab, 0x3f, 0x16, 0x3c, 0xea, 0x47, 0xa3, 0xae, 0x05, 0x43, 0xb2,
	0x81, 0xc7, 0x03, 0x6a, 0x32, 0xeb, 0x10, 0xee, 0xe6, 0xf6, 0x5f, 0xde, 0x17, 0xcf, 0x04, 0x5f,
	0xe1, 0xe9, 0x0d, 0x1e, 0x99, 0xde, 0xdc, 0xfa, 0xf3, 0xe9, 0x39, 0xf8, 0xaf, 0x06, 0xa3, 0x06,
	0x71, 0x76, 0x93, 0xd6, 0xff, 0xf6, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x13, 0x02, 0xd3, 0x72,
	0xab, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GatewayServiceClient is the client API for GatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewayServiceClient interface {
	GetGatewayByMac(ctx context.Context, in *GetGatewayByMacRequest, opts ...grpc.CallOption) (*GetGatewayByMacResponse, error)
	GetGatewayMacList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetGatewayMacListResponse, error)
}

type gatewayServiceClient struct {
	cc *grpc.ClientConn
}

func NewGatewayServiceClient(cc *grpc.ClientConn) GatewayServiceClient {
	return &gatewayServiceClient{cc}
}

func (c *gatewayServiceClient) GetGatewayByMac(ctx context.Context, in *GetGatewayByMacRequest, opts ...grpc.CallOption) (*GetGatewayByMacResponse, error) {
	out := new(GetGatewayByMacResponse)
	err := c.cc.Invoke(ctx, "/dev.GatewayService/GetGatewayByMac", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) GetGatewayMacList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetGatewayMacListResponse, error) {
	out := new(GetGatewayMacListResponse)
	err := c.cc.Invoke(ctx, "/dev.GatewayService/GetGatewayMacList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServiceServer is the server API for GatewayService service.
type GatewayServiceServer interface {
	GetGatewayByMac(context.Context, *GetGatewayByMacRequest) (*GetGatewayByMacResponse, error)
	GetGatewayMacList(context.Context, *empty.Empty) (*GetGatewayMacListResponse, error)
}

// UnimplementedGatewayServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGatewayServiceServer struct {
}

func (*UnimplementedGatewayServiceServer) GetGatewayByMac(ctx context.Context, req *GetGatewayByMacRequest) (*GetGatewayByMacResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGatewayByMac not implemented")
}
func (*UnimplementedGatewayServiceServer) GetGatewayMacList(ctx context.Context, req *empty.Empty) (*GetGatewayMacListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGatewayMacList not implemented")
}

func RegisterGatewayServiceServer(s *grpc.Server, srv GatewayServiceServer) {
	s.RegisterService(&_GatewayService_serviceDesc, srv)
}

func _GatewayService_GetGatewayByMac_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGatewayByMacRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).GetGatewayByMac(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.GatewayService/GetGatewayByMac",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).GetGatewayByMac(ctx, req.(*GetGatewayByMacRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_GetGatewayMacList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).GetGatewayMacList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.GatewayService/GetGatewayMacList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).GetGatewayMacList(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _GatewayService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dev.GatewayService",
	HandlerType: (*GatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGatewayByMac",
			Handler:    _GatewayService_GetGatewayByMac_Handler,
		},
		{
			MethodName: "GetGatewayMacList",
			Handler:    _GatewayService_GetGatewayMacList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lpwan-app-server/mxprotocol-server/gateway_item.proto",
}
