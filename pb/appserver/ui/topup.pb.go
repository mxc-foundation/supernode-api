// Code generated by protoc-gen-go. DO NOT EDIT.
// source: topup.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Money int32

const (
	Money_ETH     Money = 0
	Money_ETH_MXC Money = 1
	Money_TETH    Money = 2
)

var Money_name = map[int32]string{
	0: "ETH",
	1: "ETH_MXC",
	2: "TETH",
}

var Money_value = map[string]int32{
	"ETH":     0,
	"ETH_MXC": 1,
	"TETH":    2,
}

func (x Money) String() string {
	return proto.EnumName(Money_name, int32(x))
}

func (Money) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8eec749941d0cb6c, []int{0}
}

type GetTopUpHistoryRequest struct {
	OrgId                int64    `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int64    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTopUpHistoryRequest) Reset()         { *m = GetTopUpHistoryRequest{} }
func (m *GetTopUpHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetTopUpHistoryRequest) ProtoMessage()    {}
func (*GetTopUpHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eec749941d0cb6c, []int{0}
}

func (m *GetTopUpHistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopUpHistoryRequest.Unmarshal(m, b)
}
func (m *GetTopUpHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopUpHistoryRequest.Marshal(b, m, deterministic)
}
func (m *GetTopUpHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopUpHistoryRequest.Merge(m, src)
}
func (m *GetTopUpHistoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetTopUpHistoryRequest.Size(m)
}
func (m *GetTopUpHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopUpHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopUpHistoryRequest proto.InternalMessageInfo

func (m *GetTopUpHistoryRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *GetTopUpHistoryRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetTopUpHistoryRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type TopUpHistory struct {
	Amount               float64  `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
	CreatedAt            string   `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	TxHash               string   `protobuf:"bytes,3,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopUpHistory) Reset()         { *m = TopUpHistory{} }
func (m *TopUpHistory) String() string { return proto.CompactTextString(m) }
func (*TopUpHistory) ProtoMessage()    {}
func (*TopUpHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eec749941d0cb6c, []int{1}
}

func (m *TopUpHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopUpHistory.Unmarshal(m, b)
}
func (m *TopUpHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopUpHistory.Marshal(b, m, deterministic)
}
func (m *TopUpHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopUpHistory.Merge(m, src)
}
func (m *TopUpHistory) XXX_Size() int {
	return xxx_messageInfo_TopUpHistory.Size(m)
}
func (m *TopUpHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_TopUpHistory.DiscardUnknown(m)
}

var xxx_messageInfo_TopUpHistory proto.InternalMessageInfo

func (m *TopUpHistory) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TopUpHistory) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *TopUpHistory) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

type GetTopUpHistoryResponse struct {
	Count                int64           `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	TopupHistory         []*TopUpHistory `protobuf:"bytes,2,rep,name=topup_history,json=topupHistory,proto3" json:"topup_history,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetTopUpHistoryResponse) Reset()         { *m = GetTopUpHistoryResponse{} }
func (m *GetTopUpHistoryResponse) String() string { return proto.CompactTextString(m) }
func (*GetTopUpHistoryResponse) ProtoMessage()    {}
func (*GetTopUpHistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eec749941d0cb6c, []int{2}
}

func (m *GetTopUpHistoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopUpHistoryResponse.Unmarshal(m, b)
}
func (m *GetTopUpHistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopUpHistoryResponse.Marshal(b, m, deterministic)
}
func (m *GetTopUpHistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopUpHistoryResponse.Merge(m, src)
}
func (m *GetTopUpHistoryResponse) XXX_Size() int {
	return xxx_messageInfo_GetTopUpHistoryResponse.Size(m)
}
func (m *GetTopUpHistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopUpHistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopUpHistoryResponse proto.InternalMessageInfo

func (m *GetTopUpHistoryResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *GetTopUpHistoryResponse) GetTopupHistory() []*TopUpHistory {
	if m != nil {
		return m.TopupHistory
	}
	return nil
}

type GetTopUpDestinationRequest struct {
	OrgId                int64    `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTopUpDestinationRequest) Reset()         { *m = GetTopUpDestinationRequest{} }
func (m *GetTopUpDestinationRequest) String() string { return proto.CompactTextString(m) }
func (*GetTopUpDestinationRequest) ProtoMessage()    {}
func (*GetTopUpDestinationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eec749941d0cb6c, []int{3}
}

func (m *GetTopUpDestinationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopUpDestinationRequest.Unmarshal(m, b)
}
func (m *GetTopUpDestinationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopUpDestinationRequest.Marshal(b, m, deterministic)
}
func (m *GetTopUpDestinationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopUpDestinationRequest.Merge(m, src)
}
func (m *GetTopUpDestinationRequest) XXX_Size() int {
	return xxx_messageInfo_GetTopUpDestinationRequest.Size(m)
}
func (m *GetTopUpDestinationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopUpDestinationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopUpDestinationRequest proto.InternalMessageInfo

func (m *GetTopUpDestinationRequest) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

type GetTopUpDestinationResponse struct {
	ActiveAccount        string   `protobuf:"bytes,1,opt,name=active_account,json=activeAccount,proto3" json:"active_account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTopUpDestinationResponse) Reset()         { *m = GetTopUpDestinationResponse{} }
func (m *GetTopUpDestinationResponse) String() string { return proto.CompactTextString(m) }
func (*GetTopUpDestinationResponse) ProtoMessage()    {}
func (*GetTopUpDestinationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eec749941d0cb6c, []int{4}
}

func (m *GetTopUpDestinationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopUpDestinationResponse.Unmarshal(m, b)
}
func (m *GetTopUpDestinationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopUpDestinationResponse.Marshal(b, m, deterministic)
}
func (m *GetTopUpDestinationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopUpDestinationResponse.Merge(m, src)
}
func (m *GetTopUpDestinationResponse) XXX_Size() int {
	return xxx_messageInfo_GetTopUpDestinationResponse.Size(m)
}
func (m *GetTopUpDestinationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopUpDestinationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopUpDestinationResponse proto.InternalMessageInfo

func (m *GetTopUpDestinationResponse) GetActiveAccount() string {
	if m != nil {
		return m.ActiveAccount
	}
	return ""
}

func init() {
	proto.RegisterEnum("api.Money", Money_name, Money_value)
	proto.RegisterType((*GetTopUpHistoryRequest)(nil), "api.GetTopUpHistoryRequest")
	proto.RegisterType((*TopUpHistory)(nil), "api.TopUpHistory")
	proto.RegisterType((*GetTopUpHistoryResponse)(nil), "api.GetTopUpHistoryResponse")
	proto.RegisterType((*GetTopUpDestinationRequest)(nil), "api.GetTopUpDestinationRequest")
	proto.RegisterType((*GetTopUpDestinationResponse)(nil), "api.GetTopUpDestinationResponse")
}

func init() { proto.RegisterFile("topup.proto", fileDescriptor_8eec749941d0cb6c) }

var fileDescriptor_8eec749941d0cb6c = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x5d, 0x4f, 0xd4, 0x40,
	0x14, 0xb5, 0xad, 0xdb, 0xb5, 0x77, 0x41, 0x71, 0x10, 0x68, 0xba, 0x18, 0x37, 0x4d, 0x8c, 0xc4,
	0xc4, 0xdd, 0x04, 0x12, 0xdf, 0x89, 0x10, 0xeb, 0x03, 0x2f, 0xb5, 0x26, 0xbe, 0x68, 0x33, 0x76,
	0x87, 0x76, 0x12, 0xe8, 0x1d, 0x3b, 0xb7, 0x04, 0x5e, 0xfd, 0x0b, 0xfe, 0x34, 0xff, 0x82, 0x3f,
	0xc0, 0x9f, 0x60, 0x76, 0x66, 0xfc, 0x40, 0x2a, 0x8f, 0xe7, 0xf4, 0xcc, 0xb9, 0xe7, 0x9e, 0x5e,
	0x98, 0x10, 0xaa, 0x5e, 0xcd, 0x55, 0x87, 0x84, 0x2c, 0xe0, 0x4a, 0x26, 0xbb, 0x35, 0x62, 0x7d,
	0x26, 0x16, 0x5c, 0xc9, 0x05, 0x6f, 0x5b, 0x24, 0x4e, 0x12, 0x5b, 0x6d, 0x25, 0xe9, 0x07, 0xd8,
	0x7e, 0x2d, 0xa8, 0x40, 0xf5, 0x4e, 0x65, 0x52, 0x13, 0x76, 0x57, 0xb9, 0xf8, 0xdc, 0x0b, 0x4d,
	0x6c, 0x0b, 0x42, 0xec, 0xea, 0x52, 0x2e, 0x63, 0x6f, 0xe6, 0xed, 0x05, 0xf9, 0x08, 0xbb, 0xfa,
	0xcd, 0x92, 0x6d, 0x43, 0x88, 0xa7, 0xa7, 0x5a, 0x50, 0xec, 0x1b, 0xda, 0x21, 0xf6, 0x08, 0x46,
	0x67, 0xf2, 0x5c, 0x52, 0x1c, 0x58, 0xb5, 0x01, 0xe9, 0x47, 0x58, 0xfb, 0xdb, 0x7b, 0xf5, 0x9a,
	0x9f, 0x63, 0xdf, 0x92, 0x31, 0xf5, 0x72, 0x87, 0xd8, 0x63, 0x80, 0xaa, 0x13, 0x9c, 0xc4, 0xb2,
	0xe4, 0xd6, 0x39, 0xca, 0x23, 0xc7, 0x1c, 0x12, 0xdb, 0x81, 0x31, 0x5d, 0x96, 0x0d, 0xd7, 0x8d,
	0xb1, 0x8f, 0xf2, 0x90, 0x2e, 0x33, 0xae, 0x9b, 0xb4, 0x86, 0x9d, 0x1b, 0xf1, 0xb5, 0xc2, 0x56,
	0x8b, 0x55, 0xa0, 0xea, 0xf7, 0xa4, 0x20, 0xb7, 0x80, 0xbd, 0x84, 0x75, 0xd3, 0x50, 0xd9, 0x58,
	0x79, 0xec, 0xcf, 0x82, 0xbd, 0xc9, 0xfe, 0xc3, 0x39, 0x57, 0x72, 0x7e, 0xcd, 0x67, 0xcd, 0xe8,
	0x1c, 0x4a, 0x0f, 0x20, 0xf9, 0x35, 0xe8, 0x48, 0x68, 0x92, 0xad, 0x69, 0xf1, 0xf6, 0xae, 0xd2,
	0x23, 0x98, 0x0e, 0x3e, 0x72, 0x09, 0x9f, 0xc2, 0x7d, 0x5e, 0x91, 0xbc, 0x10, 0x25, 0xaf, 0xfe,
	0x44, 0x8d, 0xf2, 0x75, 0xcb, 0x1e, 0x5a, 0xf2, 0xf9, 0x33, 0x18, 0x9d, 0x60, 0x2b, 0xae, 0xd8,
	0x18, 0x82, 0xe3, 0x22, 0xdb, 0xb8, 0xc3, 0x26, 0x30, 0x3e, 0x2e, 0xb2, 0xf2, 0xe4, 0xfd, 0xab,
	0x0d, 0x8f, 0xdd, 0x83, 0xbb, 0xc5, 0x8a, 0xf6, 0xf7, 0x7f, 0x78, 0xae, 0xed, 0xb7, 0xa2, 0xbb,
	0x90, 0x95, 0x60, 0x12, 0x1e, 0xfc, 0xd3, 0x0e, 0x9b, 0x9a, 0x45, 0x87, 0x7f, 0x79, 0xb2, 0x3b,
	0xfc, 0xd1, 0xc6, 0x4d, 0xa7, 0x5f, 0xbe, 0x7d, 0xff, 0xea, 0x6f, 0xb1, 0x4d, 0x73, 0x4a, 0x84,
	0xea, 0x45, 0xaf, 0x16, 0xae, 0x46, 0xd6, 0xc3, 0xe6, 0xc0, 0xaa, 0xec, 0xc9, 0x35, 0xc7, 0x9b,
	0xcd, 0x25, 0xb3, 0xff, 0x0b, 0x6e, 0x1b, 0xeb, 0x0a, 0xfb, 0x14, 0x9a, 0x2b, 0x3e, 0xf8, 0x19,
	0x00, 0x00, 0xff, 0xff, 0x0f, 0x68, 0xfa, 0x56, 0xf7, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TopUpServiceClient is the client API for TopUpService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TopUpServiceClient interface {
	GetTopUpHistory(ctx context.Context, in *GetTopUpHistoryRequest, opts ...grpc.CallOption) (*GetTopUpHistoryResponse, error)
	GetTopUpDestination(ctx context.Context, in *GetTopUpDestinationRequest, opts ...grpc.CallOption) (*GetTopUpDestinationResponse, error)
}

type topUpServiceClient struct {
	cc *grpc.ClientConn
}

func NewTopUpServiceClient(cc *grpc.ClientConn) TopUpServiceClient {
	return &topUpServiceClient{cc}
}

func (c *topUpServiceClient) GetTopUpHistory(ctx context.Context, in *GetTopUpHistoryRequest, opts ...grpc.CallOption) (*GetTopUpHistoryResponse, error) {
	out := new(GetTopUpHistoryResponse)
	err := c.cc.Invoke(ctx, "/api.TopUpService/GetTopUpHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topUpServiceClient) GetTopUpDestination(ctx context.Context, in *GetTopUpDestinationRequest, opts ...grpc.CallOption) (*GetTopUpDestinationResponse, error) {
	out := new(GetTopUpDestinationResponse)
	err := c.cc.Invoke(ctx, "/api.TopUpService/GetTopUpDestination", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopUpServiceServer is the server API for TopUpService service.
type TopUpServiceServer interface {
	GetTopUpHistory(context.Context, *GetTopUpHistoryRequest) (*GetTopUpHistoryResponse, error)
	GetTopUpDestination(context.Context, *GetTopUpDestinationRequest) (*GetTopUpDestinationResponse, error)
}

// UnimplementedTopUpServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTopUpServiceServer struct {
}

func (*UnimplementedTopUpServiceServer) GetTopUpHistory(ctx context.Context, req *GetTopUpHistoryRequest) (*GetTopUpHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopUpHistory not implemented")
}
func (*UnimplementedTopUpServiceServer) GetTopUpDestination(ctx context.Context, req *GetTopUpDestinationRequest) (*GetTopUpDestinationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopUpDestination not implemented")
}

func RegisterTopUpServiceServer(s *grpc.Server, srv TopUpServiceServer) {
	s.RegisterService(&_TopUpService_serviceDesc, srv)
}

func _TopUpService_GetTopUpHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopUpHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopUpServiceServer).GetTopUpHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TopUpService/GetTopUpHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopUpServiceServer).GetTopUpHistory(ctx, req.(*GetTopUpHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopUpService_GetTopUpDestination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopUpDestinationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopUpServiceServer).GetTopUpDestination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TopUpService/GetTopUpDestination",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopUpServiceServer).GetTopUpDestination(ctx, req.(*GetTopUpDestinationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TopUpService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.TopUpService",
	HandlerType: (*TopUpServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTopUpHistory",
			Handler:    _TopUpService_GetTopUpHistory_Handler,
		},
		{
			MethodName: "GetTopUpDestination",
			Handler:    _TopUpService_GetTopUpDestination_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "topup.proto",
}
