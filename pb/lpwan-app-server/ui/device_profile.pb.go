// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lpwan-app-server/ui/device_profile.proto

package ui

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type CreateDeviceProfileRequest struct {
	// Device-profile object to create.
	DeviceProfile        *DeviceProfile `protobuf:"bytes,1,opt,name=device_profile,json=deviceProfile,proto3" json:"device_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateDeviceProfileRequest) Reset()         { *m = CreateDeviceProfileRequest{} }
func (m *CreateDeviceProfileRequest) String() string { return proto.CompactTextString(m) }
func (*CreateDeviceProfileRequest) ProtoMessage()    {}
func (*CreateDeviceProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0009d05d6b57291c, []int{0}
}

func (m *CreateDeviceProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDeviceProfileRequest.Unmarshal(m, b)
}
func (m *CreateDeviceProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDeviceProfileRequest.Marshal(b, m, deterministic)
}
func (m *CreateDeviceProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDeviceProfileRequest.Merge(m, src)
}
func (m *CreateDeviceProfileRequest) XXX_Size() int {
	return xxx_messageInfo_CreateDeviceProfileRequest.Size(m)
}
func (m *CreateDeviceProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDeviceProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDeviceProfileRequest proto.InternalMessageInfo

func (m *CreateDeviceProfileRequest) GetDeviceProfile() *DeviceProfile {
	if m != nil {
		return m.DeviceProfile
	}
	return nil
}

type CreateDeviceProfileResponse struct {
	// Device-profile ID (UUID string).
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDeviceProfileResponse) Reset()         { *m = CreateDeviceProfileResponse{} }
func (m *CreateDeviceProfileResponse) String() string { return proto.CompactTextString(m) }
func (*CreateDeviceProfileResponse) ProtoMessage()    {}
func (*CreateDeviceProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0009d05d6b57291c, []int{1}
}

func (m *CreateDeviceProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDeviceProfileResponse.Unmarshal(m, b)
}
func (m *CreateDeviceProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDeviceProfileResponse.Marshal(b, m, deterministic)
}
func (m *CreateDeviceProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDeviceProfileResponse.Merge(m, src)
}
func (m *CreateDeviceProfileResponse) XXX_Size() int {
	return xxx_messageInfo_CreateDeviceProfileResponse.Size(m)
}
func (m *CreateDeviceProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDeviceProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDeviceProfileResponse proto.InternalMessageInfo

func (m *CreateDeviceProfileResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetDeviceProfileRequest struct {
	// Device-profile ID (UUID string).
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDeviceProfileRequest) Reset()         { *m = GetDeviceProfileRequest{} }
func (m *GetDeviceProfileRequest) String() string { return proto.CompactTextString(m) }
func (*GetDeviceProfileRequest) ProtoMessage()    {}
func (*GetDeviceProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0009d05d6b57291c, []int{2}
}

func (m *GetDeviceProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceProfileRequest.Unmarshal(m, b)
}
func (m *GetDeviceProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceProfileRequest.Marshal(b, m, deterministic)
}
func (m *GetDeviceProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceProfileRequest.Merge(m, src)
}
func (m *GetDeviceProfileRequest) XXX_Size() int {
	return xxx_messageInfo_GetDeviceProfileRequest.Size(m)
}
func (m *GetDeviceProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceProfileRequest proto.InternalMessageInfo

func (m *GetDeviceProfileRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetDeviceProfileResponse struct {
	// Device-profile object.
	DeviceProfile *DeviceProfile `protobuf:"bytes,1,opt,name=device_profile,json=deviceProfile,proto3" json:"device_profile,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetDeviceProfileResponse) Reset()         { *m = GetDeviceProfileResponse{} }
func (m *GetDeviceProfileResponse) String() string { return proto.CompactTextString(m) }
func (*GetDeviceProfileResponse) ProtoMessage()    {}
func (*GetDeviceProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0009d05d6b57291c, []int{3}
}

func (m *GetDeviceProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceProfileResponse.Unmarshal(m, b)
}
func (m *GetDeviceProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceProfileResponse.Marshal(b, m, deterministic)
}
func (m *GetDeviceProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceProfileResponse.Merge(m, src)
}
func (m *GetDeviceProfileResponse) XXX_Size() int {
	return xxx_messageInfo_GetDeviceProfileResponse.Size(m)
}
func (m *GetDeviceProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceProfileResponse proto.InternalMessageInfo

func (m *GetDeviceProfileResponse) GetDeviceProfile() *DeviceProfile {
	if m != nil {
		return m.DeviceProfile
	}
	return nil
}

func (m *GetDeviceProfileResponse) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *GetDeviceProfileResponse) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type UpdateDeviceProfileRequest struct {
	// Device-profile object to update.
	DeviceProfile        *DeviceProfile `protobuf:"bytes,1,opt,name=device_profile,json=deviceProfile,proto3" json:"device_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateDeviceProfileRequest) Reset()         { *m = UpdateDeviceProfileRequest{} }
func (m *UpdateDeviceProfileRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateDeviceProfileRequest) ProtoMessage()    {}
func (*UpdateDeviceProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0009d05d6b57291c, []int{4}
}

func (m *UpdateDeviceProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateDeviceProfileRequest.Unmarshal(m, b)
}
func (m *UpdateDeviceProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateDeviceProfileRequest.Marshal(b, m, deterministic)
}
func (m *UpdateDeviceProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateDeviceProfileRequest.Merge(m, src)
}
func (m *UpdateDeviceProfileRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateDeviceProfileRequest.Size(m)
}
func (m *UpdateDeviceProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateDeviceProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateDeviceProfileRequest proto.InternalMessageInfo

func (m *UpdateDeviceProfileRequest) GetDeviceProfile() *DeviceProfile {
	if m != nil {
		return m.DeviceProfile
	}
	return nil
}

type DeleteDeviceProfileRequest struct {
	// Device-profile ID (UUID string).
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteDeviceProfileRequest) Reset()         { *m = DeleteDeviceProfileRequest{} }
func (m *DeleteDeviceProfileRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteDeviceProfileRequest) ProtoMessage()    {}
func (*DeleteDeviceProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0009d05d6b57291c, []int{5}
}

func (m *DeleteDeviceProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDeviceProfileRequest.Unmarshal(m, b)
}
func (m *DeleteDeviceProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDeviceProfileRequest.Marshal(b, m, deterministic)
}
func (m *DeleteDeviceProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDeviceProfileRequest.Merge(m, src)
}
func (m *DeleteDeviceProfileRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteDeviceProfileRequest.Size(m)
}
func (m *DeleteDeviceProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDeviceProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDeviceProfileRequest proto.InternalMessageInfo

func (m *DeleteDeviceProfileRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeviceProfileListItem struct {
	// Device-profile ID (UUID string).
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Device-profile name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Organization ID.
	OrganizationId int64 `protobuf:"varint,3,opt,name=organization_id,json=organizationID,proto3" json:"organization_id,omitempty"`
	// Network-server ID.
	NetworkServerId int64 `protobuf:"varint,4,opt,name=network_server_id,json=networkServerID,proto3" json:"network_server_id,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DeviceProfileListItem) Reset()         { *m = DeviceProfileListItem{} }
func (m *DeviceProfileListItem) String() string { return proto.CompactTextString(m) }
func (*DeviceProfileListItem) ProtoMessage()    {}
func (*DeviceProfileListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_0009d05d6b57291c, []int{6}
}

func (m *DeviceProfileListItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceProfileListItem.Unmarshal(m, b)
}
func (m *DeviceProfileListItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceProfileListItem.Marshal(b, m, deterministic)
}
func (m *DeviceProfileListItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceProfileListItem.Merge(m, src)
}
func (m *DeviceProfileListItem) XXX_Size() int {
	return xxx_messageInfo_DeviceProfileListItem.Size(m)
}
func (m *DeviceProfileListItem) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceProfileListItem.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceProfileListItem proto.InternalMessageInfo

func (m *DeviceProfileListItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeviceProfileListItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeviceProfileListItem) GetOrganizationId() int64 {
	if m != nil {
		return m.OrganizationId
	}
	return 0
}

func (m *DeviceProfileListItem) GetNetworkServerId() int64 {
	if m != nil {
		return m.NetworkServerId
	}
	return 0
}

func (m *DeviceProfileListItem) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *DeviceProfileListItem) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type ListDeviceProfileRequest struct {
	// Max number of items to return.
	Limit int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// Organization id to filter on.
	OrganizationId int64 `protobuf:"varint,3,opt,name=organization_id,json=organizationID,proto3" json:"organization_id,omitempty"`
	// Application id to filter on.
	ApplicationId        int64    `protobuf:"varint,4,opt,name=application_id,json=applicationID,proto3" json:"application_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListDeviceProfileRequest) Reset()         { *m = ListDeviceProfileRequest{} }
func (m *ListDeviceProfileRequest) String() string { return proto.CompactTextString(m) }
func (*ListDeviceProfileRequest) ProtoMessage()    {}
func (*ListDeviceProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0009d05d6b57291c, []int{7}
}

func (m *ListDeviceProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDeviceProfileRequest.Unmarshal(m, b)
}
func (m *ListDeviceProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDeviceProfileRequest.Marshal(b, m, deterministic)
}
func (m *ListDeviceProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDeviceProfileRequest.Merge(m, src)
}
func (m *ListDeviceProfileRequest) XXX_Size() int {
	return xxx_messageInfo_ListDeviceProfileRequest.Size(m)
}
func (m *ListDeviceProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDeviceProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListDeviceProfileRequest proto.InternalMessageInfo

func (m *ListDeviceProfileRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListDeviceProfileRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListDeviceProfileRequest) GetOrganizationId() int64 {
	if m != nil {
		return m.OrganizationId
	}
	return 0
}

func (m *ListDeviceProfileRequest) GetApplicationId() int64 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

type ListDeviceProfileResponse struct {
	// Total number of device-profiles.
	TotalCount           int64                    `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	Result               []*DeviceProfileListItem `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ListDeviceProfileResponse) Reset()         { *m = ListDeviceProfileResponse{} }
func (m *ListDeviceProfileResponse) String() string { return proto.CompactTextString(m) }
func (*ListDeviceProfileResponse) ProtoMessage()    {}
func (*ListDeviceProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0009d05d6b57291c, []int{8}
}

func (m *ListDeviceProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDeviceProfileResponse.Unmarshal(m, b)
}
func (m *ListDeviceProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDeviceProfileResponse.Marshal(b, m, deterministic)
}
func (m *ListDeviceProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDeviceProfileResponse.Merge(m, src)
}
func (m *ListDeviceProfileResponse) XXX_Size() int {
	return xxx_messageInfo_ListDeviceProfileResponse.Size(m)
}
func (m *ListDeviceProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDeviceProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListDeviceProfileResponse proto.InternalMessageInfo

func (m *ListDeviceProfileResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListDeviceProfileResponse) GetResult() []*DeviceProfileListItem {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateDeviceProfileRequest)(nil), "ui.CreateDeviceProfileRequest")
	proto.RegisterType((*CreateDeviceProfileResponse)(nil), "ui.CreateDeviceProfileResponse")
	proto.RegisterType((*GetDeviceProfileRequest)(nil), "ui.GetDeviceProfileRequest")
	proto.RegisterType((*GetDeviceProfileResponse)(nil), "ui.GetDeviceProfileResponse")
	proto.RegisterType((*UpdateDeviceProfileRequest)(nil), "ui.UpdateDeviceProfileRequest")
	proto.RegisterType((*DeleteDeviceProfileRequest)(nil), "ui.DeleteDeviceProfileRequest")
	proto.RegisterType((*DeviceProfileListItem)(nil), "ui.DeviceProfileListItem")
	proto.RegisterType((*ListDeviceProfileRequest)(nil), "ui.ListDeviceProfileRequest")
	proto.RegisterType((*ListDeviceProfileResponse)(nil), "ui.ListDeviceProfileResponse")
}

func init() {
	proto.RegisterFile("lpwan-app-server/ui/device_profile.proto", fileDescriptor_0009d05d6b57291c)
}

var fileDescriptor_0009d05d6b57291c = []byte{
	// 632 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcb, 0x6e, 0xd3, 0x4c,
	0x14, 0x56, 0xec, 0xd4, 0x52, 0x4f, 0xd4, 0x54, 0x19, 0xe5, 0xcf, 0x9f, 0x38, 0x81, 0x80, 0x25,
	0x44, 0x88, 0x48, 0x2c, 0xd2, 0x0d, 0xb0, 0xab, 0x1a, 0x54, 0x45, 0x62, 0x81, 0xcc, 0x65, 0x1b,
	0x4d, 0xed, 0x49, 0x34, 0x60, 0x7b, 0x06, 0x7b, 0xdc, 0x0a, 0x50, 0x37, 0x2c, 0x78, 0x01, 0x36,
	0x3c, 0x01, 0x0f, 0xc3, 0x96, 0x57, 0xe0, 0x41, 0x90, 0x67, 0x26, 0x55, 0x2e, 0x36, 0x97, 0x8a,
	0x8d, 0xe5, 0x39, 0xe7, 0x9b, 0x73, 0xbe, 0xf9, 0xce, 0x05, 0x06, 0x21, 0xbf, 0xc0, 0xf1, 0x08,
	0x73, 0x3e, 0x4a, 0x49, 0x72, 0x4e, 0x12, 0x37, 0xa3, 0x6e, 0x40, 0xce, 0xa9, 0x4f, 0xe6, 0x3c,
	0x61, 0x0b, 0x1a, 0x92, 0x31, 0x4f, 0x98, 0x60, 0xc8, 0xc8, 0xa8, 0xdd, 0x5b, 0x32, 0xb6, 0x0c,
	0x89, 0x8b, 0x39, 0x75, 0x71, 0x1c, 0x33, 0x81, 0x05, 0x65, 0x71, 0xaa, 0x10, 0x76, 0x47, 0xd0,
	0x88, 0xa4, 0x02, 0x47, 0xdc, 0xbd, 0xfa, 0xd3, 0xae, 0x06, 0x89, 0xb8, 0x78, 0xe7, 0xca, 0xaf,
	0x36, 0x39, 0x45, 0x99, 0x75, 0x4a, 0x1d, 0xd1, 0x79, 0x05, 0xf6, 0x49, 0x42, 0xb0, 0x20, 0x53,
	0xc9, 0xe8, 0x99, 0xf2, 0x7a, 0xe4, 0x6d, 0x46, 0x52, 0x81, 0x1e, 0x42, 0x7d, 0x93, 0x69, 0xbb,
	0x72, 0xab, 0x32, 0xa8, 0x4d, 0x1a, 0xe3, 0x8c, 0x8e, 0x37, 0x6f, 0x1c, 0x04, 0xeb, 0x47, 0x67,
	0x04, 0xdd, 0xc2, 0xb8, 0x29, 0x67, 0x71, 0x4a, 0x50, 0x1d, 0x0c, 0x1a, 0xc8, 0x60, 0xfb, 0x9e,
	0x41, 0x03, 0xe7, 0x1e, 0xfc, 0x7f, 0x4a, 0x44, 0x21, 0x87, 0x6d, 0xe8, 0xb7, 0x0a, 0xb4, 0x77,
	0xb1, 0x3a, 0xee, 0xb5, 0x09, 0xa3, 0x47, 0x00, 0xbe, 0x24, 0x1c, 0xcc, 0xb1, 0x68, 0x1b, 0xf2,
	0x96, 0x3d, 0x56, 0xd5, 0x50, 0x5a, 0x9d, 0x65, 0x8b, 0xf1, 0x8b, 0x95, 0xea, 0xde, 0xbe, 0x46,
	0x1f, 0x8b, 0xfc, 0x6a, 0xc6, 0x83, 0xd5, 0x55, 0xf3, 0xf7, 0x57, 0x35, 0xfa, 0x58, 0xe4, 0xf2,
	0xbf, 0x94, 0x87, 0x7f, 0x2c, 0xff, 0x7d, 0xb0, 0xa7, 0x24, 0x24, 0x25, 0x71, 0xb7, 0x25, 0xfd,
	0x64, 0xc0, 0x7f, 0x1b, 0xc0, 0xa7, 0x34, 0x15, 0x33, 0x41, 0xa2, 0x6d, 0x24, 0x42, 0x50, 0x8d,
	0x71, 0x44, 0xa4, 0x3e, 0xfb, 0x9e, 0xfc, 0x47, 0x77, 0xe1, 0x90, 0x25, 0x4b, 0x1c, 0xd3, 0xf7,
	0xb2, 0x57, 0xe7, 0x34, 0x90, 0x1a, 0x98, 0x5e, 0x7d, 0xdd, 0x3c, 0x9b, 0xa2, 0x21, 0x34, 0x62,
	0x22, 0x2e, 0x58, 0xf2, 0x66, 0xae, 0xfa, 0x31, 0x87, 0x56, 0x25, 0xf4, 0x50, 0x3b, 0x9e, 0x4b,
	0xfb, 0x6c, 0xba, 0x55, 0x8e, 0xbd, 0xeb, 0x97, 0xc3, 0xfa, 0x9b, 0x72, 0x7c, 0xa9, 0x40, 0x3b,
	0x7f, 0x7b, 0xa1, 0x6a, 0x4d, 0xd8, 0x0b, 0x69, 0x44, 0x85, 0x94, 0xc3, 0xf4, 0xd4, 0x01, 0xb5,
	0xc0, 0x62, 0x8b, 0x45, 0x4a, 0x54, 0xcf, 0x98, 0x9e, 0x3e, 0xfd, 0xb9, 0x2a, 0x77, 0xa0, 0x8e,
	0x39, 0x0f, 0xa9, 0x7f, 0x85, 0x53, 0x92, 0x1c, 0xac, 0x59, 0x67, 0x53, 0x87, 0x41, 0xa7, 0x80,
	0x99, 0x6e, 0xfb, 0x3e, 0xd4, 0x04, 0x13, 0x38, 0x9c, 0xfb, 0x2c, 0x8b, 0x57, 0x04, 0x41, 0x9a,
	0x4e, 0x72, 0x0b, 0x7a, 0x00, 0x56, 0x42, 0xd2, 0x2c, 0xcc, 0x59, 0x9a, 0x83, 0xda, 0xa4, 0xb3,
	0xd3, 0x41, 0xab, 0x92, 0x7b, 0x1a, 0x38, 0xf9, 0x5a, 0x85, 0xe6, 0x06, 0x22, 0xaf, 0x0d, 0xf5,
	0x09, 0x7a, 0x0d, 0x96, 0x1a, 0x6d, 0x74, 0x33, 0x8f, 0x52, 0xbe, 0x3e, 0xec, 0x7e, 0xa9, 0x5f,
	0xf1, 0x76, 0xfa, 0x1f, 0xbf, 0xff, 0xf8, 0x6c, 0x74, 0x9c, 0xa6, 0xdc, 0x77, 0xaa, 0x85, 0x47,
	0xab, 0x0d, 0xf5, 0xb8, 0x32, 0x44, 0x3e, 0x98, 0xa7, 0x44, 0xa0, 0x6e, 0x1e, 0xa8, 0x64, 0x41,
	0xd8, 0xbd, 0x62, 0xa7, 0x4e, 0x71, 0x5b, 0xa6, 0xe8, 0xa2, 0x4e, 0x51, 0x0a, 0xf7, 0x03, 0x0d,
	0x2e, 0x51, 0x06, 0x96, 0x1a, 0x42, 0xf5, 0xa0, 0xf2, 0x81, 0xb4, 0x5b, 0x3b, 0x6d, 0xf4, 0x24,
	0xdf, 0xb7, 0xce, 0x91, 0x4c, 0x32, 0xb2, 0x07, 0xc5, 0x49, 0xb6, 0xb6, 0x3d, 0x0d, 0x2e, 0xd5,
	0xdb, 0x2c, 0x35, 0xa3, 0x2a, 0x6d, 0xf9, 0xbc, 0x96, 0xa6, 0xd5, 0x6f, 0x1b, 0xfe, 0xe2, 0x6d,
	0x18, 0xaa, 0x79, 0x65, 0x91, 0x14, 0xa9, 0xac, 0xb5, 0xed, 0x1b, 0x25, 0x5e, 0xad, 0x61, 0x4f,
	0xe6, 0x69, 0xa1, 0xc2, 0x32, 0x9d, 0x59, 0x92, 0xd5, 0xd1, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x5d, 0x12, 0x0d, 0x98, 0xe9, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DeviceProfileServiceClient is the client API for DeviceProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeviceProfileServiceClient interface {
	// Create creates the given device-profile.
	Create(ctx context.Context, in *CreateDeviceProfileRequest, opts ...grpc.CallOption) (*CreateDeviceProfileResponse, error)
	// Get returns the device-profile matching the given id.
	Get(ctx context.Context, in *GetDeviceProfileRequest, opts ...grpc.CallOption) (*GetDeviceProfileResponse, error)
	// Update updates the given device-profile.
	Update(ctx context.Context, in *UpdateDeviceProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Delete deletes the device-profile matching the given id.
	Delete(ctx context.Context, in *DeleteDeviceProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// List lists the available device-profiles.
	List(ctx context.Context, in *ListDeviceProfileRequest, opts ...grpc.CallOption) (*ListDeviceProfileResponse, error)
}

type deviceProfileServiceClient struct {
	cc *grpc.ClientConn
}

func NewDeviceProfileServiceClient(cc *grpc.ClientConn) DeviceProfileServiceClient {
	return &deviceProfileServiceClient{cc}
}

func (c *deviceProfileServiceClient) Create(ctx context.Context, in *CreateDeviceProfileRequest, opts ...grpc.CallOption) (*CreateDeviceProfileResponse, error) {
	out := new(CreateDeviceProfileResponse)
	err := c.cc.Invoke(ctx, "/ui.DeviceProfileService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceProfileServiceClient) Get(ctx context.Context, in *GetDeviceProfileRequest, opts ...grpc.CallOption) (*GetDeviceProfileResponse, error) {
	out := new(GetDeviceProfileResponse)
	err := c.cc.Invoke(ctx, "/ui.DeviceProfileService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceProfileServiceClient) Update(ctx context.Context, in *UpdateDeviceProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ui.DeviceProfileService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceProfileServiceClient) Delete(ctx context.Context, in *DeleteDeviceProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ui.DeviceProfileService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceProfileServiceClient) List(ctx context.Context, in *ListDeviceProfileRequest, opts ...grpc.CallOption) (*ListDeviceProfileResponse, error) {
	out := new(ListDeviceProfileResponse)
	err := c.cc.Invoke(ctx, "/ui.DeviceProfileService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceProfileServiceServer is the server API for DeviceProfileService service.
type DeviceProfileServiceServer interface {
	// Create creates the given device-profile.
	Create(context.Context, *CreateDeviceProfileRequest) (*CreateDeviceProfileResponse, error)
	// Get returns the device-profile matching the given id.
	Get(context.Context, *GetDeviceProfileRequest) (*GetDeviceProfileResponse, error)
	// Update updates the given device-profile.
	Update(context.Context, *UpdateDeviceProfileRequest) (*empty.Empty, error)
	// Delete deletes the device-profile matching the given id.
	Delete(context.Context, *DeleteDeviceProfileRequest) (*empty.Empty, error)
	// List lists the available device-profiles.
	List(context.Context, *ListDeviceProfileRequest) (*ListDeviceProfileResponse, error)
}

// UnimplementedDeviceProfileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDeviceProfileServiceServer struct {
}

func (*UnimplementedDeviceProfileServiceServer) Create(ctx context.Context, req *CreateDeviceProfileRequest) (*CreateDeviceProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedDeviceProfileServiceServer) Get(ctx context.Context, req *GetDeviceProfileRequest) (*GetDeviceProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedDeviceProfileServiceServer) Update(ctx context.Context, req *UpdateDeviceProfileRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedDeviceProfileServiceServer) Delete(ctx context.Context, req *DeleteDeviceProfileRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedDeviceProfileServiceServer) List(ctx context.Context, req *ListDeviceProfileRequest) (*ListDeviceProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterDeviceProfileServiceServer(s *grpc.Server, srv DeviceProfileServiceServer) {
	s.RegisterService(&_DeviceProfileService_serviceDesc, srv)
}

func _DeviceProfileService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceProfileServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ui.DeviceProfileService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceProfileServiceServer).Create(ctx, req.(*CreateDeviceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceProfileService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceProfileServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ui.DeviceProfileService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceProfileServiceServer).Get(ctx, req.(*GetDeviceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceProfileService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeviceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceProfileServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ui.DeviceProfileService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceProfileServiceServer).Update(ctx, req.(*UpdateDeviceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceProfileService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceProfileServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ui.DeviceProfileService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceProfileServiceServer).Delete(ctx, req.(*DeleteDeviceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceProfileService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeviceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceProfileServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ui.DeviceProfileService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceProfileServiceServer).List(ctx, req.(*ListDeviceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeviceProfileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ui.DeviceProfileService",
	HandlerType: (*DeviceProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _DeviceProfileService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DeviceProfileService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DeviceProfileService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DeviceProfileService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DeviceProfileService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lpwan-app-server/ui/device_profile.proto",
}
