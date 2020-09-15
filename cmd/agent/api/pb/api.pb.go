// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package pb

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

type EventType int32

const (
	EventType_ADDED    EventType = 0
	EventType_MODIFIED EventType = 1
	EventType_DELETED  EventType = 2
)

var EventType_name = map[int32]string{
	0: "ADDED",
	1: "MODIFIED",
	2: "DELETED",
}

var EventType_value = map[string]int32{
	"ADDED":    0,
	"MODIFIED": 1,
	"DELETED":  2,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

type TagCardinality int32

const (
	TagCardinality_LOW          TagCardinality = 0
	TagCardinality_ORCHESTRATOR TagCardinality = 1
	TagCardinality_HIGH         TagCardinality = 2
)

var TagCardinality_name = map[int32]string{
	0: "LOW",
	1: "ORCHESTRATOR",
	2: "HIGH",
}

var TagCardinality_value = map[string]int32{
	"LOW":          0,
	"ORCHESTRATOR": 1,
	"HIGH":         2,
}

func (x TagCardinality) String() string {
	return proto.EnumName(TagCardinality_name, int32(x))
}

func (TagCardinality) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

type HostnameRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HostnameRequest) Reset()         { *m = HostnameRequest{} }
func (m *HostnameRequest) String() string { return proto.CompactTextString(m) }
func (*HostnameRequest) ProtoMessage()    {}
func (*HostnameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *HostnameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostnameRequest.Unmarshal(m, b)
}
func (m *HostnameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostnameRequest.Marshal(b, m, deterministic)
}
func (m *HostnameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostnameRequest.Merge(m, src)
}
func (m *HostnameRequest) XXX_Size() int {
	return xxx_messageInfo_HostnameRequest.Size(m)
}
func (m *HostnameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HostnameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HostnameRequest proto.InternalMessageInfo

// The response message containing the requested hostname
type HostnameReply struct {
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HostnameReply) Reset()         { *m = HostnameReply{} }
func (m *HostnameReply) String() string { return proto.CompactTextString(m) }
func (*HostnameReply) ProtoMessage()    {}
func (*HostnameReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *HostnameReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostnameReply.Unmarshal(m, b)
}
func (m *HostnameReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostnameReply.Marshal(b, m, deterministic)
}
func (m *HostnameReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostnameReply.Merge(m, src)
}
func (m *HostnameReply) XXX_Size() int {
	return xxx_messageInfo_HostnameReply.Size(m)
}
func (m *HostnameReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HostnameReply.DiscardUnknown(m)
}

var xxx_messageInfo_HostnameReply proto.InternalMessageInfo

func (m *HostnameReply) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

type StreamTagsRequest struct {
	Cardinality          TagCardinality `protobuf:"varint,1,opt,name=cardinality,proto3,enum=pb.TagCardinality" json:"cardinality,omitempty"`
	IncludeFilter        *Filter        `protobuf:"bytes,2,opt,name=includeFilter,proto3" json:"includeFilter,omitempty"`
	ExcludeFilter        *Filter        `protobuf:"bytes,3,opt,name=excludeFilter,proto3" json:"excludeFilter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StreamTagsRequest) Reset()         { *m = StreamTagsRequest{} }
func (m *StreamTagsRequest) String() string { return proto.CompactTextString(m) }
func (*StreamTagsRequest) ProtoMessage()    {}
func (*StreamTagsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *StreamTagsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamTagsRequest.Unmarshal(m, b)
}
func (m *StreamTagsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamTagsRequest.Marshal(b, m, deterministic)
}
func (m *StreamTagsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTagsRequest.Merge(m, src)
}
func (m *StreamTagsRequest) XXX_Size() int {
	return xxx_messageInfo_StreamTagsRequest.Size(m)
}
func (m *StreamTagsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTagsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTagsRequest proto.InternalMessageInfo

func (m *StreamTagsRequest) GetCardinality() TagCardinality {
	if m != nil {
		return m.Cardinality
	}
	return TagCardinality_LOW
}

func (m *StreamTagsRequest) GetIncludeFilter() *Filter {
	if m != nil {
		return m.IncludeFilter
	}
	return nil
}

func (m *StreamTagsRequest) GetExcludeFilter() *Filter {
	if m != nil {
		return m.ExcludeFilter
	}
	return nil
}

type StreamTagsResponse struct {
	Type                 EventType `protobuf:"varint,1,opt,name=type,proto3,enum=pb.EventType" json:"type,omitempty"`
	Entity               *Entity   `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *StreamTagsResponse) Reset()         { *m = StreamTagsResponse{} }
func (m *StreamTagsResponse) String() string { return proto.CompactTextString(m) }
func (*StreamTagsResponse) ProtoMessage()    {}
func (*StreamTagsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *StreamTagsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamTagsResponse.Unmarshal(m, b)
}
func (m *StreamTagsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamTagsResponse.Marshal(b, m, deterministic)
}
func (m *StreamTagsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTagsResponse.Merge(m, src)
}
func (m *StreamTagsResponse) XXX_Size() int {
	return xxx_messageInfo_StreamTagsResponse.Size(m)
}
func (m *StreamTagsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTagsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTagsResponse proto.InternalMessageInfo

func (m *StreamTagsResponse) GetType() EventType {
	if m != nil {
		return m.Type
	}
	return EventType_ADDED
}

func (m *StreamTagsResponse) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

type Filter struct {
	KubeNamespace        string   `protobuf:"bytes,1,opt,name=kubeNamespace,proto3" json:"kubeNamespace,omitempty"`
	Image                string   `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	ContainerName        string   `protobuf:"bytes,3,opt,name=containerName,proto3" json:"containerName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetKubeNamespace() string {
	if m != nil {
		return m.KubeNamespace
	}
	return ""
}

func (m *Filter) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Filter) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

type Entity struct {
	Id                   *EntityId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tags                 []string  `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetId() *EntityId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Entity) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type FetchEntityRequest struct {
	Id                   *EntityId      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Cardinality          TagCardinality `protobuf:"varint,2,opt,name=cardinality,proto3,enum=pb.TagCardinality" json:"cardinality,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FetchEntityRequest) Reset()         { *m = FetchEntityRequest{} }
func (m *FetchEntityRequest) String() string { return proto.CompactTextString(m) }
func (*FetchEntityRequest) ProtoMessage()    {}
func (*FetchEntityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *FetchEntityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchEntityRequest.Unmarshal(m, b)
}
func (m *FetchEntityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchEntityRequest.Marshal(b, m, deterministic)
}
func (m *FetchEntityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchEntityRequest.Merge(m, src)
}
func (m *FetchEntityRequest) XXX_Size() int {
	return xxx_messageInfo_FetchEntityRequest.Size(m)
}
func (m *FetchEntityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchEntityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchEntityRequest proto.InternalMessageInfo

func (m *FetchEntityRequest) GetId() *EntityId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *FetchEntityRequest) GetCardinality() TagCardinality {
	if m != nil {
		return m.Cardinality
	}
	return TagCardinality_LOW
}

type FetchEntityResponse struct {
	Entity               *Entity  `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchEntityResponse) Reset()         { *m = FetchEntityResponse{} }
func (m *FetchEntityResponse) String() string { return proto.CompactTextString(m) }
func (*FetchEntityResponse) ProtoMessage()    {}
func (*FetchEntityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{7}
}

func (m *FetchEntityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchEntityResponse.Unmarshal(m, b)
}
func (m *FetchEntityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchEntityResponse.Marshal(b, m, deterministic)
}
func (m *FetchEntityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchEntityResponse.Merge(m, src)
}
func (m *FetchEntityResponse) XXX_Size() int {
	return xxx_messageInfo_FetchEntityResponse.Size(m)
}
func (m *FetchEntityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchEntityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchEntityResponse proto.InternalMessageInfo

func (m *FetchEntityResponse) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

type EntityId struct {
	Prefix               string   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntityId) Reset()         { *m = EntityId{} }
func (m *EntityId) String() string { return proto.CompactTextString(m) }
func (*EntityId) ProtoMessage()    {}
func (*EntityId) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{8}
}

func (m *EntityId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityId.Unmarshal(m, b)
}
func (m *EntityId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityId.Marshal(b, m, deterministic)
}
func (m *EntityId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityId.Merge(m, src)
}
func (m *EntityId) XXX_Size() int {
	return xxx_messageInfo_EntityId.Size(m)
}
func (m *EntityId) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityId.DiscardUnknown(m)
}

var xxx_messageInfo_EntityId proto.InternalMessageInfo

func (m *EntityId) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *EntityId) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func init() {
	proto.RegisterEnum("pb.EventType", EventType_name, EventType_value)
	proto.RegisterEnum("pb.TagCardinality", TagCardinality_name, TagCardinality_value)
	proto.RegisterType((*HostnameRequest)(nil), "pb.HostnameRequest")
	proto.RegisterType((*HostnameReply)(nil), "pb.HostnameReply")
	proto.RegisterType((*StreamTagsRequest)(nil), "pb.StreamTagsRequest")
	proto.RegisterType((*StreamTagsResponse)(nil), "pb.StreamTagsResponse")
	proto.RegisterType((*Filter)(nil), "pb.Filter")
	proto.RegisterType((*Entity)(nil), "pb.Entity")
	proto.RegisterType((*FetchEntityRequest)(nil), "pb.FetchEntityRequest")
	proto.RegisterType((*FetchEntityResponse)(nil), "pb.FetchEntityResponse")
	proto.RegisterType((*EntityId)(nil), "pb.EntityId")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 625 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcf, 0x4e, 0xdb, 0x4e,
	0x10, 0xc6, 0x0e, 0x84, 0x64, 0x42, 0xc0, 0x19, 0xfe, 0xfc, 0xf8, 0x45, 0x48, 0x4d, 0x57, 0x3d,
	0xa0, 0x54, 0x22, 0x90, 0xd2, 0x43, 0xb9, 0x21, 0x6c, 0x48, 0x24, 0xda, 0x48, 0x8b, 0xab, 0x1e,
	0x7a, 0x40, 0x1b, 0x67, 0x71, 0xb6, 0x4d, 0xec, 0xad, 0xbd, 0x41, 0xe4, 0xda, 0x57, 0xe8, 0x73,
	0xf4, 0x69, 0xfa, 0x00, 0xbd, 0xf4, 0x41, 0x2a, 0xaf, 0x4d, 0x12, 0x13, 0xb5, 0xea, 0xcd, 0xfb,
	0xed, 0xcc, 0x7c, 0xdf, 0x7c, 0xb3, 0x63, 0x28, 0x33, 0x29, 0x8e, 0x64, 0x14, 0xaa, 0x10, 0x4d,
	0xd9, 0xaf, 0x1f, 0xf8, 0x61, 0xe8, 0x8f, 0x78, 0x8b, 0x49, 0xd1, 0x62, 0x41, 0x10, 0x2a, 0xa6,
	0x44, 0x18, 0xc4, 0x69, 0x04, 0xa9, 0xc1, 0x56, 0x27, 0x8c, 0x55, 0xc0, 0xc6, 0x9c, 0xf2, 0x2f,
	0x13, 0x1e, 0x2b, 0xf2, 0x12, 0xaa, 0x73, 0x48, 0x8e, 0xa6, 0x58, 0x87, 0xd2, 0x30, 0x03, 0xf6,
	0x8d, 0x86, 0x71, 0x58, 0xa6, 0xb3, 0x33, 0xf9, 0x6e, 0x40, 0xed, 0x46, 0x45, 0x9c, 0x8d, 0x5d,
	0xe6, 0xc7, 0x59, 0x09, 0x3c, 0x85, 0x8a, 0xc7, 0xa2, 0x81, 0x08, 0xd8, 0x48, 0xa8, 0xa9, 0x4e,
	0xda, 0x6c, 0xe3, 0x91, 0xec, 0x1f, 0xb9, 0xcc, 0xbf, 0x98, 0xdf, 0xd0, 0xc5, 0x30, 0x3c, 0x86,
	0xaa, 0x08, 0xbc, 0xd1, 0x64, 0xc0, 0x2f, 0xc5, 0x48, 0xf1, 0x68, 0xdf, 0x6c, 0x18, 0x87, 0x95,
	0x36, 0x24, 0x79, 0x29, 0x42, 0xf3, 0x01, 0x49, 0x06, 0x7f, 0x58, 0xcc, 0x28, 0x2c, 0x67, 0xe4,
	0x02, 0xc8, 0x47, 0xc0, 0x45, 0xb9, 0xb1, 0x0c, 0x83, 0x98, 0xe3, 0x73, 0x58, 0x55, 0x53, 0xc9,
	0x33, 0xa1, 0xd5, 0x24, 0xdd, 0xb9, 0xe7, 0x81, 0x72, 0xa7, 0x92, 0x53, 0x7d, 0x85, 0x04, 0x8a,
	0x3c, 0x50, 0x49, 0x37, 0x0b, 0xaa, 0x1c, 0x8d, 0xd0, 0xec, 0x86, 0x7c, 0x82, 0x62, 0x26, 0xec,
	0x05, 0x54, 0x3f, 0x4f, 0xfa, 0xfc, 0x1d, 0x1b, 0xf3, 0x58, 0x32, 0xef, 0xd1, 0xb7, 0x3c, 0x88,
	0x3b, 0xb0, 0x26, 0xc6, 0xcc, 0xe7, 0xba, 0x64, 0x99, 0xa6, 0x87, 0x24, 0xd7, 0x0b, 0x03, 0xc5,
	0x44, 0xc0, 0xa3, 0x24, 0x56, 0x37, 0x55, 0xa6, 0x79, 0x90, 0x9c, 0x41, 0x31, 0x65, 0xc7, 0x03,
	0x30, 0xc5, 0x40, 0x13, 0x54, 0xda, 0x1b, 0x73, 0x55, 0xdd, 0x01, 0x35, 0xc5, 0x00, 0x11, 0x56,
	0x15, 0xf3, 0xe3, 0x7d, 0xb3, 0x51, 0x38, 0x2c, 0x53, 0xfd, 0x4d, 0x86, 0x80, 0x97, 0x5c, 0x79,
	0xc3, 0x4c, 0x7e, 0x36, 0xb4, 0xbf, 0xd7, 0x79, 0x32, 0x52, 0xf3, 0x9f, 0x46, 0x4a, 0xde, 0xc0,
	0x76, 0x8e, 0x29, 0xf3, 0x7b, 0x6e, 0xa6, 0xf1, 0x47, 0x33, 0x4f, 0xa1, 0xf4, 0x28, 0x00, 0xf7,
	0xa0, 0x28, 0x23, 0x7e, 0x27, 0x1e, 0x32, 0x1f, 0xb3, 0x13, 0x5a, 0x50, 0x98, 0x88, 0x41, 0x66,
	0x5f, 0xf2, 0xd9, 0x3c, 0x81, 0xf2, 0x6c, 0x72, 0x58, 0x86, 0xb5, 0x73, 0xdb, 0x76, 0x6c, 0x6b,
	0x05, 0x37, 0xa0, 0xf4, 0xb6, 0x67, 0x77, 0x2f, 0xbb, 0x8e, 0x6d, 0x19, 0x58, 0x81, 0x75, 0xdb,
	0xb9, 0x76, 0x5c, 0xc7, 0xb6, 0xcc, 0xe6, 0x6b, 0xd8, 0xcc, 0xb7, 0x80, 0xeb, 0x50, 0xb8, 0xee,
	0x7d, 0xb0, 0x56, 0xd0, 0x82, 0x8d, 0x1e, 0xbd, 0xe8, 0x38, 0x37, 0x2e, 0x3d, 0x77, 0x7b, 0xd4,
	0x32, 0xb0, 0x04, 0xab, 0x9d, 0xee, 0x55, 0xc7, 0x32, 0xdb, 0xef, 0x61, 0xed, 0xdc, 0xe7, 0x81,
	0xc2, 0x6b, 0xa8, 0x5c, 0x71, 0xf5, 0xb8, 0x32, 0xb8, 0x9d, 0xf4, 0xf2, 0x64, 0xa7, 0xea, 0xb5,
	0x3c, 0x28, 0x47, 0x53, 0xb2, 0xfb, 0xf5, 0xc7, 0xaf, 0x6f, 0xe6, 0x16, 0x56, 0x5b, 0xf7, 0x27,
	0x2d, 0x3f, 0x92, 0x5e, 0x2b, 0x59, 0xaa, 0xf6, 0x4f, 0x03, 0x2a, 0xba, 0xee, 0x0d, 0xf7, 0x26,
	0x11, 0xc7, 0x21, 0xec, 0xb8, 0xcc, 0xf7, 0x79, 0x94, 0x3e, 0x5b, 0x6d, 0x89, 0xe0, 0x31, 0xee,
	0x26, 0x15, 0x97, 0x36, 0xaf, 0xbe, 0xf7, 0x14, 0x4e, 0x1d, 0x27, 0xcf, 0x34, 0xdb, 0xff, 0x64,
	0x67, 0xc6, 0x16, 0xeb, 0xa0, 0xdb, 0xe4, 0x41, 0x9c, 0x19, 0xcd, 0x63, 0x03, 0xef, 0xa0, 0x96,
	0x32, 0x2d, 0x4c, 0x0c, 0x75, 0xbd, 0xe5, 0xc7, 0x52, 0xff, 0x6f, 0x09, 0xcf, 0x88, 0x1a, 0x9a,
	0xa8, 0x4e, 0x76, 0x67, 0x44, 0x77, 0x49, 0xd4, 0x6d, 0x3a, 0xd5, 0x33, 0xa3, 0xd9, 0x2f, 0xea,
	0x3f, 0xcf, 0xab, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xe5, 0x00, 0x08, 0xa8, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentClient interface {
	// get the hostname
	GetHostname(ctx context.Context, in *HostnameRequest, opts ...grpc.CallOption) (*HostnameReply, error)
}

type agentClient struct {
	cc *grpc.ClientConn
}

func NewAgentClient(cc *grpc.ClientConn) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) GetHostname(ctx context.Context, in *HostnameRequest, opts ...grpc.CallOption) (*HostnameReply, error) {
	out := new(HostnameReply)
	err := c.cc.Invoke(ctx, "/pb.Agent/GetHostname", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
type AgentServer interface {
	// get the hostname
	GetHostname(context.Context, *HostnameRequest) (*HostnameReply, error)
}

// UnimplementedAgentServer can be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (*UnimplementedAgentServer) GetHostname(ctx context.Context, req *HostnameRequest) (*HostnameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHostname not implemented")
}

func RegisterAgentServer(s *grpc.Server, srv AgentServer) {
	s.RegisterService(&_Agent_serviceDesc, srv)
}

func _Agent_GetHostname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HostnameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetHostname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Agent/GetHostname",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetHostname(ctx, req.(*HostnameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHostname",
			Handler:    _Agent_GetHostname_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// AgentSecureClient is the client API for AgentSecure service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentSecureClient interface {
	TaggerStreamEntities(ctx context.Context, in *StreamTagsRequest, opts ...grpc.CallOption) (AgentSecure_TaggerStreamEntitiesClient, error)
	TaggerFetchEntity(ctx context.Context, in *FetchEntityRequest, opts ...grpc.CallOption) (*FetchEntityResponse, error)
}

type agentSecureClient struct {
	cc *grpc.ClientConn
}

func NewAgentSecureClient(cc *grpc.ClientConn) AgentSecureClient {
	return &agentSecureClient{cc}
}

func (c *agentSecureClient) TaggerStreamEntities(ctx context.Context, in *StreamTagsRequest, opts ...grpc.CallOption) (AgentSecure_TaggerStreamEntitiesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AgentSecure_serviceDesc.Streams[0], "/pb.AgentSecure/TaggerStreamEntities", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentSecureTaggerStreamEntitiesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AgentSecure_TaggerStreamEntitiesClient interface {
	Recv() (*StreamTagsResponse, error)
	grpc.ClientStream
}

type agentSecureTaggerStreamEntitiesClient struct {
	grpc.ClientStream
}

func (x *agentSecureTaggerStreamEntitiesClient) Recv() (*StreamTagsResponse, error) {
	m := new(StreamTagsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentSecureClient) TaggerFetchEntity(ctx context.Context, in *FetchEntityRequest, opts ...grpc.CallOption) (*FetchEntityResponse, error) {
	out := new(FetchEntityResponse)
	err := c.cc.Invoke(ctx, "/pb.AgentSecure/TaggerFetchEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentSecureServer is the server API for AgentSecure service.
type AgentSecureServer interface {
	TaggerStreamEntities(*StreamTagsRequest, AgentSecure_TaggerStreamEntitiesServer) error
	TaggerFetchEntity(context.Context, *FetchEntityRequest) (*FetchEntityResponse, error)
}

// UnimplementedAgentSecureServer can be embedded to have forward compatible implementations.
type UnimplementedAgentSecureServer struct {
}

func (*UnimplementedAgentSecureServer) TaggerStreamEntities(req *StreamTagsRequest, srv AgentSecure_TaggerStreamEntitiesServer) error {
	return status.Errorf(codes.Unimplemented, "method TaggerStreamEntities not implemented")
}
func (*UnimplementedAgentSecureServer) TaggerFetchEntity(ctx context.Context, req *FetchEntityRequest) (*FetchEntityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaggerFetchEntity not implemented")
}

func RegisterAgentSecureServer(s *grpc.Server, srv AgentSecureServer) {
	s.RegisterService(&_AgentSecure_serviceDesc, srv)
}

func _AgentSecure_TaggerStreamEntities_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamTagsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentSecureServer).TaggerStreamEntities(m, &agentSecureTaggerStreamEntitiesServer{stream})
}

type AgentSecure_TaggerStreamEntitiesServer interface {
	Send(*StreamTagsResponse) error
	grpc.ServerStream
}

type agentSecureTaggerStreamEntitiesServer struct {
	grpc.ServerStream
}

func (x *agentSecureTaggerStreamEntitiesServer) Send(m *StreamTagsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _AgentSecure_TaggerFetchEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentSecureServer).TaggerFetchEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AgentSecure/TaggerFetchEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentSecureServer).TaggerFetchEntity(ctx, req.(*FetchEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgentSecure_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AgentSecure",
	HandlerType: (*AgentSecureServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TaggerFetchEntity",
			Handler:    _AgentSecure_TaggerFetchEntity_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TaggerStreamEntities",
			Handler:       _AgentSecure_TaggerStreamEntities_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}
