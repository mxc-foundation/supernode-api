// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lpwan-app-server/gateway/heartbeat.proto

package heatbeat

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

// MiningRequest sends gateway list to m2m
type HeartbeatRequest struct {
	GatewayMac           string   `protobuf:"bytes,1,opt,name=gateway_mac,json=gatewayMac,proto3" json:"gateway_mac,omitempty"`
	Model                string   `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	ConfigHash           string   `protobuf:"bytes,3,opt,name=config_hash,json=configHash,proto3" json:"config_hash,omitempty"`
	OsVersion            string   `protobuf:"bytes,4,opt,name=os_version,json=osVersion,proto3" json:"os_version,omitempty"`
	Statistics           string   `protobuf:"bytes,5,opt,name=statistics,proto3" json:"statistics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatRequest) Reset()         { *m = HeartbeatRequest{} }
func (m *HeartbeatRequest) String() string { return proto.CompactTextString(m) }
func (*HeartbeatRequest) ProtoMessage()    {}
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c90f16355bbb9e3, []int{0}
}

func (m *HeartbeatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatRequest.Unmarshal(m, b)
}
func (m *HeartbeatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatRequest.Marshal(b, m, deterministic)
}
func (m *HeartbeatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatRequest.Merge(m, src)
}
func (m *HeartbeatRequest) XXX_Size() int {
	return xxx_messageInfo_HeartbeatRequest.Size(m)
}
func (m *HeartbeatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatRequest proto.InternalMessageInfo

func (m *HeartbeatRequest) GetGatewayMac() string {
	if m != nil {
		return m.GatewayMac
	}
	return ""
}

func (m *HeartbeatRequest) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *HeartbeatRequest) GetConfigHash() string {
	if m != nil {
		return m.ConfigHash
	}
	return ""
}

func (m *HeartbeatRequest) GetOsVersion() string {
	if m != nil {
		return m.OsVersion
	}
	return ""
}

func (m *HeartbeatRequest) GetStatistics() string {
	if m != nil {
		return m.Statistics
	}
	return ""
}

type HeartbeatResponse struct {
	NewFirmwareLink      string   `protobuf:"bytes,1,opt,name=new_firmware_link,json=newFirmwareLink,proto3" json:"new_firmware_link,omitempty"`
	Config               string   `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatResponse) Reset()         { *m = HeartbeatResponse{} }
func (m *HeartbeatResponse) String() string { return proto.CompactTextString(m) }
func (*HeartbeatResponse) ProtoMessage()    {}
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c90f16355bbb9e3, []int{1}
}

func (m *HeartbeatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatResponse.Unmarshal(m, b)
}
func (m *HeartbeatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatResponse.Marshal(b, m, deterministic)
}
func (m *HeartbeatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatResponse.Merge(m, src)
}
func (m *HeartbeatResponse) XXX_Size() int {
	return xxx_messageInfo_HeartbeatResponse.Size(m)
}
func (m *HeartbeatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatResponse proto.InternalMessageInfo

func (m *HeartbeatResponse) GetNewFirmwareLink() string {
	if m != nil {
		return m.NewFirmwareLink
	}
	return ""
}

func (m *HeartbeatResponse) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

func init() {
	proto.RegisterType((*HeartbeatRequest)(nil), "heatbeat.HeartbeatRequest")
	proto.RegisterType((*HeartbeatResponse)(nil), "heatbeat.HeartbeatResponse")
}

func init() {
	proto.RegisterFile("lpwan-app-server/gateway/heartbeat.proto", fileDescriptor_4c90f16355bbb9e3)
}

var fileDescriptor_4c90f16355bbb9e3 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x89, 0xda, 0x62, 0xae, 0x0b, 0xed, 0x20, 0x12, 0x2a, 0xfe, 0xd0, 0x55, 0x11, 0x9a,
	0x82, 0xbe, 0x82, 0x48, 0x17, 0xba, 0xa9, 0xa0, 0xee, 0xc2, 0x6d, 0xbc, 0x6d, 0x86, 0x26, 0x33,
	0xe3, 0xdc, 0x31, 0x83, 0x8f, 0xe4, 0x5b, 0x8a, 0x99, 0xa9, 0x16, 0x71, 0x79, 0xbe, 0xc3, 0x81,
	0x8f, 0x03, 0xe3, 0xda, 0x78, 0x54, 0x13, 0x34, 0x66, 0xc2, 0x64, 0x5b, 0xb2, 0xd3, 0x15, 0x3a,
	0xf2, 0xf8, 0x31, 0xad, 0x08, 0xad, 0x5b, 0x10, 0xba, 0xdc, 0x58, 0xed, 0xb4, 0xd8, 0xaf, 0x08,
	0xbb, 0x3c, 0xfa, 0x4c, 0xe0, 0x68, 0xb6, 0x69, 0xe7, 0xf4, 0xf6, 0x4e, 0xec, 0xc4, 0x05, 0x1c,
	0xc4, 0x65, 0xd1, 0x60, 0x99, 0x25, 0x97, 0xc9, 0x38, 0x9d, 0x43, 0x44, 0x0f, 0x58, 0x8a, 0x63,
	0xe8, 0x35, 0xfa, 0x95, 0xea, 0x6c, 0xa7, 0xab, 0x42, 0xf8, 0x9e, 0x95, 0x5a, 0x2d, 0xe5, 0xaa,
	0xa8, 0x90, 0xab, 0x6c, 0x37, 0xcc, 0x02, 0x9a, 0x21, 0x57, 0xe2, 0x0c, 0x40, 0x73, 0xd1, 0x92,
	0x65, 0xa9, 0x55, 0xb6, 0xd7, 0xf5, 0xa9, 0xe6, 0xa7, 0x00, 0xc4, 0x39, 0x00, 0x3b, 0x74, 0x92,
	0x9d, 0x2c, 0x39, 0xeb, 0x85, 0xf9, 0x2f, 0x19, 0x3d, 0xc3, 0x60, 0x4b, 0x95, 0x8d, 0x56, 0x4c,
	0xe2, 0x0a, 0x06, 0x8a, 0x7c, 0xb1, 0x94, 0xb6, 0xf1, 0x68, 0xa9, 0xa8, 0xa5, 0x5a, 0x47, 0xe3,
	0x43, 0x45, 0xfe, 0x2e, 0xf2, 0x7b, 0xa9, 0xd6, 0xe2, 0x04, 0xfa, 0xc1, 0x26, 0x7a, 0xc7, 0x74,
	0xfd, 0xb2, 0xf5, 0xc1, 0x23, 0xd9, 0x56, 0x96, 0x24, 0x6e, 0x21, 0xfd, 0x61, 0x62, 0x98, 0x6f,
	0x0e, 0xcb, 0xff, 0x9e, 0x35, 0x3c, 0xfd, 0xb7, 0x0b, 0x76, 0x8b, 0x7e, 0xf7, 0xf7, 0xcd, 0x57,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x46, 0x13, 0x45, 0xa9, 0x9b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HeartbeatServiceClient is the client API for HeartbeatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HeartbeatServiceClient interface {
	Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error)
}

type heartbeatServiceClient struct {
	cc *grpc.ClientConn
}

func NewHeartbeatServiceClient(cc *grpc.ClientConn) HeartbeatServiceClient {
	return &heartbeatServiceClient{cc}
}

func (c *heartbeatServiceClient) Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error) {
	out := new(HeartbeatResponse)
	err := c.cc.Invoke(ctx, "/heatbeat.HeartbeatService/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeartbeatServiceServer is the server API for HeartbeatService service.
type HeartbeatServiceServer interface {
	Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error)
}

// UnimplementedHeartbeatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHeartbeatServiceServer struct {
}

func (*UnimplementedHeartbeatServiceServer) Heartbeat(ctx context.Context, req *HeartbeatRequest) (*HeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}

func RegisterHeartbeatServiceServer(s *grpc.Server, srv HeartbeatServiceServer) {
	s.RegisterService(&_HeartbeatService_serviceDesc, srv)
}

func _HeartbeatService_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartbeatServiceServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heatbeat.HeartbeatService/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartbeatServiceServer).Heartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HeartbeatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "heatbeat.HeartbeatService",
	HandlerType: (*HeartbeatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Heartbeat",
			Handler:    _HeartbeatService_Heartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lpwan-app-server/gateway/heartbeat.proto",
}
