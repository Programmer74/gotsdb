// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package proto

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

type KvsStoreRequest struct {
	MsgId                uint32   `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
	Key                  []byte   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KvsStoreRequest) Reset()         { *m = KvsStoreRequest{} }
func (m *KvsStoreRequest) String() string { return proto.CompactTextString(m) }
func (*KvsStoreRequest) ProtoMessage()    {}
func (*KvsStoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}

func (m *KvsStoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KvsStoreRequest.Unmarshal(m, b)
}
func (m *KvsStoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KvsStoreRequest.Marshal(b, m, deterministic)
}
func (m *KvsStoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KvsStoreRequest.Merge(m, src)
}
func (m *KvsStoreRequest) XXX_Size() int {
	return xxx_messageInfo_KvsStoreRequest.Size(m)
}
func (m *KvsStoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KvsStoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KvsStoreRequest proto.InternalMessageInfo

func (m *KvsStoreRequest) GetMsgId() uint32 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *KvsStoreRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KvsStoreRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type KvsStoreResponse struct {
	MsgId                uint32   `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
	Ok                   bool     `protobuf:"varint,2,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KvsStoreResponse) Reset()         { *m = KvsStoreResponse{} }
func (m *KvsStoreResponse) String() string { return proto.CompactTextString(m) }
func (*KvsStoreResponse) ProtoMessage()    {}
func (*KvsStoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}

func (m *KvsStoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KvsStoreResponse.Unmarshal(m, b)
}
func (m *KvsStoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KvsStoreResponse.Marshal(b, m, deterministic)
}
func (m *KvsStoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KvsStoreResponse.Merge(m, src)
}
func (m *KvsStoreResponse) XXX_Size() int {
	return xxx_messageInfo_KvsStoreResponse.Size(m)
}
func (m *KvsStoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_KvsStoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_KvsStoreResponse proto.InternalMessageInfo

func (m *KvsStoreResponse) GetMsgId() uint32 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *KvsStoreResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type KvsKeyExistsRequest struct {
	MsgId                uint32   `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
	Key                  []byte   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KvsKeyExistsRequest) Reset()         { *m = KvsKeyExistsRequest{} }
func (m *KvsKeyExistsRequest) String() string { return proto.CompactTextString(m) }
func (*KvsKeyExistsRequest) ProtoMessage()    {}
func (*KvsKeyExistsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{2}
}

func (m *KvsKeyExistsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KvsKeyExistsRequest.Unmarshal(m, b)
}
func (m *KvsKeyExistsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KvsKeyExistsRequest.Marshal(b, m, deterministic)
}
func (m *KvsKeyExistsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KvsKeyExistsRequest.Merge(m, src)
}
func (m *KvsKeyExistsRequest) XXX_Size() int {
	return xxx_messageInfo_KvsKeyExistsRequest.Size(m)
}
func (m *KvsKeyExistsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KvsKeyExistsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KvsKeyExistsRequest proto.InternalMessageInfo

func (m *KvsKeyExistsRequest) GetMsgId() uint32 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *KvsKeyExistsRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type KvsKeyExistsResponse struct {
	MsgId                uint32   `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
	Exists               bool     `protobuf:"varint,2,opt,name=exists,proto3" json:"exists,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KvsKeyExistsResponse) Reset()         { *m = KvsKeyExistsResponse{} }
func (m *KvsKeyExistsResponse) String() string { return proto.CompactTextString(m) }
func (*KvsKeyExistsResponse) ProtoMessage()    {}
func (*KvsKeyExistsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{3}
}

func (m *KvsKeyExistsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KvsKeyExistsResponse.Unmarshal(m, b)
}
func (m *KvsKeyExistsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KvsKeyExistsResponse.Marshal(b, m, deterministic)
}
func (m *KvsKeyExistsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KvsKeyExistsResponse.Merge(m, src)
}
func (m *KvsKeyExistsResponse) XXX_Size() int {
	return xxx_messageInfo_KvsKeyExistsResponse.Size(m)
}
func (m *KvsKeyExistsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_KvsKeyExistsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_KvsKeyExistsResponse proto.InternalMessageInfo

func (m *KvsKeyExistsResponse) GetMsgId() uint32 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *KvsKeyExistsResponse) GetExists() bool {
	if m != nil {
		return m.Exists
	}
	return false
}

type KvsRetrieveRequest struct {
	MsgId                uint32   `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
	Key                  []byte   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KvsRetrieveRequest) Reset()         { *m = KvsRetrieveRequest{} }
func (m *KvsRetrieveRequest) String() string { return proto.CompactTextString(m) }
func (*KvsRetrieveRequest) ProtoMessage()    {}
func (*KvsRetrieveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{4}
}

func (m *KvsRetrieveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KvsRetrieveRequest.Unmarshal(m, b)
}
func (m *KvsRetrieveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KvsRetrieveRequest.Marshal(b, m, deterministic)
}
func (m *KvsRetrieveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KvsRetrieveRequest.Merge(m, src)
}
func (m *KvsRetrieveRequest) XXX_Size() int {
	return xxx_messageInfo_KvsRetrieveRequest.Size(m)
}
func (m *KvsRetrieveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KvsRetrieveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KvsRetrieveRequest proto.InternalMessageInfo

func (m *KvsRetrieveRequest) GetMsgId() uint32 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *KvsRetrieveRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type KvsRetrieveResponse struct {
	MsgId                uint32   `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KvsRetrieveResponse) Reset()         { *m = KvsRetrieveResponse{} }
func (m *KvsRetrieveResponse) String() string { return proto.CompactTextString(m) }
func (*KvsRetrieveResponse) ProtoMessage()    {}
func (*KvsRetrieveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{5}
}

func (m *KvsRetrieveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KvsRetrieveResponse.Unmarshal(m, b)
}
func (m *KvsRetrieveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KvsRetrieveResponse.Marshal(b, m, deterministic)
}
func (m *KvsRetrieveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KvsRetrieveResponse.Merge(m, src)
}
func (m *KvsRetrieveResponse) XXX_Size() int {
	return xxx_messageInfo_KvsRetrieveResponse.Size(m)
}
func (m *KvsRetrieveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_KvsRetrieveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_KvsRetrieveResponse proto.InternalMessageInfo

func (m *KvsRetrieveResponse) GetMsgId() uint32 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *KvsRetrieveResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type KvsDeleteRequest struct {
	MsgId                uint32   `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
	Key                  []byte   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KvsDeleteRequest) Reset()         { *m = KvsDeleteRequest{} }
func (m *KvsDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*KvsDeleteRequest) ProtoMessage()    {}
func (*KvsDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{6}
}

func (m *KvsDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KvsDeleteRequest.Unmarshal(m, b)
}
func (m *KvsDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KvsDeleteRequest.Marshal(b, m, deterministic)
}
func (m *KvsDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KvsDeleteRequest.Merge(m, src)
}
func (m *KvsDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_KvsDeleteRequest.Size(m)
}
func (m *KvsDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KvsDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KvsDeleteRequest proto.InternalMessageInfo

func (m *KvsDeleteRequest) GetMsgId() uint32 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *KvsDeleteRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type KvsDeleteResponse struct {
	MsgId                uint32   `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
	Ok                   bool     `protobuf:"varint,2,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KvsDeleteResponse) Reset()         { *m = KvsDeleteResponse{} }
func (m *KvsDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*KvsDeleteResponse) ProtoMessage()    {}
func (*KvsDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{7}
}

func (m *KvsDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KvsDeleteResponse.Unmarshal(m, b)
}
func (m *KvsDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KvsDeleteResponse.Marshal(b, m, deterministic)
}
func (m *KvsDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KvsDeleteResponse.Merge(m, src)
}
func (m *KvsDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_KvsDeleteResponse.Size(m)
}
func (m *KvsDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_KvsDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_KvsDeleteResponse proto.InternalMessageInfo

func (m *KvsDeleteResponse) GetMsgId() uint32 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *KvsDeleteResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type TSPoints struct {
	Points               map[uint64]float64 `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TSPoints) Reset()         { *m = TSPoints{} }
func (m *TSPoints) String() string { return proto.CompactTextString(m) }
func (*TSPoints) ProtoMessage()    {}
func (*TSPoints) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{8}
}

func (m *TSPoints) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TSPoints.Unmarshal(m, b)
}
func (m *TSPoints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TSPoints.Marshal(b, m, deterministic)
}
func (m *TSPoints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TSPoints.Merge(m, src)
}
func (m *TSPoints) XXX_Size() int {
	return xxx_messageInfo_TSPoints.Size(m)
}
func (m *TSPoints) XXX_DiscardUnknown() {
	xxx_messageInfo_TSPoints.DiscardUnknown(m)
}

var xxx_messageInfo_TSPoints proto.InternalMessageInfo

func (m *TSPoints) GetPoints() map[uint64]float64 {
	if m != nil {
		return m.Points
	}
	return nil
}

type TSStoreRequest struct {
	MsgId                uint32               `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
	DataSource           string               `protobuf:"bytes,2,opt,name=dataSource,proto3" json:"dataSource,omitempty"`
	Values               map[string]*TSPoints `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ExpirationMillis     uint64               `protobuf:"varint,4,opt,name=expirationMillis,proto3" json:"expirationMillis,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TSStoreRequest) Reset()         { *m = TSStoreRequest{} }
func (m *TSStoreRequest) String() string { return proto.CompactTextString(m) }
func (*TSStoreRequest) ProtoMessage()    {}
func (*TSStoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{9}
}

func (m *TSStoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TSStoreRequest.Unmarshal(m, b)
}
func (m *TSStoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TSStoreRequest.Marshal(b, m, deterministic)
}
func (m *TSStoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TSStoreRequest.Merge(m, src)
}
func (m *TSStoreRequest) XXX_Size() int {
	return xxx_messageInfo_TSStoreRequest.Size(m)
}
func (m *TSStoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TSStoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TSStoreRequest proto.InternalMessageInfo

func (m *TSStoreRequest) GetMsgId() uint32 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *TSStoreRequest) GetDataSource() string {
	if m != nil {
		return m.DataSource
	}
	return ""
}

func (m *TSStoreRequest) GetValues() map[string]*TSPoints {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *TSStoreRequest) GetExpirationMillis() uint64 {
	if m != nil {
		return m.ExpirationMillis
	}
	return 0
}

type TSStoreResponse struct {
	MsgId                uint32   `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
	Ok                   bool     `protobuf:"varint,2,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TSStoreResponse) Reset()         { *m = TSStoreResponse{} }
func (m *TSStoreResponse) String() string { return proto.CompactTextString(m) }
func (*TSStoreResponse) ProtoMessage()    {}
func (*TSStoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{10}
}

func (m *TSStoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TSStoreResponse.Unmarshal(m, b)
}
func (m *TSStoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TSStoreResponse.Marshal(b, m, deterministic)
}
func (m *TSStoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TSStoreResponse.Merge(m, src)
}
func (m *TSStoreResponse) XXX_Size() int {
	return xxx_messageInfo_TSStoreResponse.Size(m)
}
func (m *TSStoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TSStoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TSStoreResponse proto.InternalMessageInfo

func (m *TSStoreResponse) GetMsgId() uint32 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *TSStoreResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type TSRetrieveRequest struct {
	MsgId                uint32   `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
	DataSource           string   `protobuf:"bytes,2,opt,name=dataSource,proto3" json:"dataSource,omitempty"`
	FromTimestamp        uint64   `protobuf:"varint,3,opt,name=fromTimestamp,proto3" json:"fromTimestamp,omitempty"`
	ToTimestamp          uint64   `protobuf:"varint,4,opt,name=toTimestamp,proto3" json:"toTimestamp,omitempty"`
	Tags                 []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TSRetrieveRequest) Reset()         { *m = TSRetrieveRequest{} }
func (m *TSRetrieveRequest) String() string { return proto.CompactTextString(m) }
func (*TSRetrieveRequest) ProtoMessage()    {}
func (*TSRetrieveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{11}
}

func (m *TSRetrieveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TSRetrieveRequest.Unmarshal(m, b)
}
func (m *TSRetrieveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TSRetrieveRequest.Marshal(b, m, deterministic)
}
func (m *TSRetrieveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TSRetrieveRequest.Merge(m, src)
}
func (m *TSRetrieveRequest) XXX_Size() int {
	return xxx_messageInfo_TSRetrieveRequest.Size(m)
}
func (m *TSRetrieveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TSRetrieveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TSRetrieveRequest proto.InternalMessageInfo

func (m *TSRetrieveRequest) GetMsgId() uint32 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *TSRetrieveRequest) GetDataSource() string {
	if m != nil {
		return m.DataSource
	}
	return ""
}

func (m *TSRetrieveRequest) GetFromTimestamp() uint64 {
	if m != nil {
		return m.FromTimestamp
	}
	return 0
}

func (m *TSRetrieveRequest) GetToTimestamp() uint64 {
	if m != nil {
		return m.ToTimestamp
	}
	return 0
}

func (m *TSRetrieveRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type TSRetrieveResponse struct {
	MsgId                uint32               `protobuf:"varint,1,opt,name=msgId,proto3" json:"msgId,omitempty"`
	DataSource           string               `protobuf:"bytes,2,opt,name=dataSource,proto3" json:"dataSource,omitempty"`
	FromTimestamp        uint64               `protobuf:"varint,3,opt,name=fromTimestamp,proto3" json:"fromTimestamp,omitempty"`
	ToTimestamp          uint64               `protobuf:"varint,4,opt,name=toTimestamp,proto3" json:"toTimestamp,omitempty"`
	Values               map[string]*TSPoints `protobuf:"bytes,5,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TSRetrieveResponse) Reset()         { *m = TSRetrieveResponse{} }
func (m *TSRetrieveResponse) String() string { return proto.CompactTextString(m) }
func (*TSRetrieveResponse) ProtoMessage()    {}
func (*TSRetrieveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{12}
}

func (m *TSRetrieveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TSRetrieveResponse.Unmarshal(m, b)
}
func (m *TSRetrieveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TSRetrieveResponse.Marshal(b, m, deterministic)
}
func (m *TSRetrieveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TSRetrieveResponse.Merge(m, src)
}
func (m *TSRetrieveResponse) XXX_Size() int {
	return xxx_messageInfo_TSRetrieveResponse.Size(m)
}
func (m *TSRetrieveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TSRetrieveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TSRetrieveResponse proto.InternalMessageInfo

func (m *TSRetrieveResponse) GetMsgId() uint32 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *TSRetrieveResponse) GetDataSource() string {
	if m != nil {
		return m.DataSource
	}
	return ""
}

func (m *TSRetrieveResponse) GetFromTimestamp() uint64 {
	if m != nil {
		return m.FromTimestamp
	}
	return 0
}

func (m *TSRetrieveResponse) GetToTimestamp() uint64 {
	if m != nil {
		return m.ToTimestamp
	}
	return 0
}

func (m *TSRetrieveResponse) GetValues() map[string]*TSPoints {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*KvsStoreRequest)(nil), "proto.KvsStoreRequest")
	proto.RegisterType((*KvsStoreResponse)(nil), "proto.KvsStoreResponse")
	proto.RegisterType((*KvsKeyExistsRequest)(nil), "proto.KvsKeyExistsRequest")
	proto.RegisterType((*KvsKeyExistsResponse)(nil), "proto.KvsKeyExistsResponse")
	proto.RegisterType((*KvsRetrieveRequest)(nil), "proto.KvsRetrieveRequest")
	proto.RegisterType((*KvsRetrieveResponse)(nil), "proto.KvsRetrieveResponse")
	proto.RegisterType((*KvsDeleteRequest)(nil), "proto.KvsDeleteRequest")
	proto.RegisterType((*KvsDeleteResponse)(nil), "proto.KvsDeleteResponse")
	proto.RegisterType((*TSPoints)(nil), "proto.TSPoints")
	proto.RegisterMapType((map[uint64]float64)(nil), "proto.TSPoints.PointsEntry")
	proto.RegisterType((*TSStoreRequest)(nil), "proto.TSStoreRequest")
	proto.RegisterMapType((map[string]*TSPoints)(nil), "proto.TSStoreRequest.ValuesEntry")
	proto.RegisterType((*TSStoreResponse)(nil), "proto.TSStoreResponse")
	proto.RegisterType((*TSRetrieveRequest)(nil), "proto.TSRetrieveRequest")
	proto.RegisterType((*TSRetrieveResponse)(nil), "proto.TSRetrieveResponse")
	proto.RegisterMapType((map[string]*TSPoints)(nil), "proto.TSRetrieveResponse.ValuesEntry")
}

func init() {
	proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1)
}

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 581 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x55, 0x92, 0x36, 0xac, 0x37, 0xdb, 0xda, 0x99, 0x51, 0xb2, 0x4c, 0x42, 0x25, 0x62, 0x52,
	0xc5, 0x43, 0x1f, 0xba, 0x87, 0xad, 0x15, 0x43, 0x1a, 0xea, 0x34, 0x41, 0x85, 0x40, 0x49, 0xc4,
	0x7b, 0xd8, 0xcc, 0x14, 0xf5, 0xc3, 0x21, 0x76, 0xa3, 0x95, 0x3f, 0xc1, 0x2f, 0xe0, 0x91, 0x1f,
	0xc9, 0x1b, 0x8a, 0xe3, 0x25, 0x6e, 0x9a, 0x95, 0x45, 0x42, 0x7b, 0x8a, 0x7d, 0x7d, 0xef, 0xf1,
	0xb9, 0xc7, 0x27, 0x17, 0x1a, 0x51, 0x78, 0xd5, 0x0b, 0x23, 0xc2, 0x08, 0xaa, 0xf3, 0x8f, 0xfd,
	0x09, 0x9a, 0xe3, 0x98, 0xba, 0x8c, 0x44, 0xd8, 0xc1, 0xdf, 0x17, 0x98, 0x32, 0xb4, 0x0f, 0xf5,
	0x19, 0xbd, 0x79, 0x7f, 0x6d, 0x2a, 0x1d, 0xa5, 0xbb, 0xe3, 0xa4, 0x1b, 0xd4, 0x02, 0x6d, 0x82,
	0x97, 0xa6, 0xda, 0x51, 0xba, 0xdb, 0x4e, 0xb2, 0x4c, 0xf2, 0x62, 0x7f, 0xba, 0xc0, 0xa6, 0xc6,
	0x63, 0xe9, 0xc6, 0x3e, 0x85, 0x56, 0x0e, 0x48, 0x43, 0x32, 0xa7, 0xf8, 0x1e, 0xc4, 0x5d, 0x50,
	0xc9, 0x84, 0x03, 0x6e, 0x39, 0x2a, 0x99, 0xd8, 0x67, 0xf0, 0x74, 0x1c, 0xd3, 0x31, 0x5e, 0x5e,
	0xdc, 0x06, 0x94, 0xd1, 0x8a, 0x74, 0xec, 0x11, 0xec, 0xaf, 0x96, 0x6f, 0xbc, 0xbc, 0x0d, 0x3a,
	0xe6, 0x79, 0x82, 0x80, 0xd8, 0xd9, 0x6f, 0x00, 0x8d, 0x63, 0xea, 0x60, 0x16, 0x05, 0x38, 0xae,
	0x2a, 0x89, 0x7d, 0xce, 0x5b, 0xc8, 0xab, 0x37, 0x52, 0xc8, 0xf4, 0x53, 0x65, 0xfd, 0x86, 0x5c,
	0xbf, 0x11, 0x9e, 0x62, 0x56, 0xf9, 0xfa, 0x01, 0xec, 0x49, 0xb5, 0x95, 0xc4, 0xff, 0x01, 0x5b,
	0x9e, 0xfb, 0x99, 0x04, 0x73, 0x46, 0xd1, 0x31, 0xe8, 0x21, 0x5f, 0x99, 0x4a, 0x47, 0xeb, 0x1a,
	0xfd, 0xc3, 0xd4, 0x32, 0xbd, 0xbb, 0x84, 0x5e, 0xfa, 0xb9, 0x98, 0xb3, 0x68, 0xe9, 0x88, 0x54,
	0x6b, 0x00, 0x86, 0x14, 0xbe, 0x23, 0x97, 0xdc, 0x59, 0x2b, 0xd8, 0x25, 0xb9, 0x54, 0x11, 0xed,
	0x0e, 0xd5, 0x53, 0xc5, 0xfe, 0xa3, 0xc0, 0xae, 0xe7, 0x3e, 0xc0, 0x83, 0x2f, 0x00, 0xae, 0x7d,
	0xe6, 0xbb, 0x64, 0x11, 0x5d, 0xa5, 0x38, 0x0d, 0x47, 0x8a, 0xa0, 0x01, 0xe8, 0x1c, 0x95, 0x9a,
	0x1a, 0x27, 0xfe, 0x32, 0x23, 0x2e, 0x83, 0xf7, 0xbe, 0xf0, 0x1c, 0x41, 0x3f, 0x2d, 0x40, 0xaf,
	0xa1, 0x85, 0x6f, 0xc3, 0x20, 0xf2, 0x59, 0x40, 0xe6, 0x1f, 0x83, 0xe9, 0x34, 0xa0, 0x66, 0x8d,
	0x93, 0x5f, 0x8b, 0x5b, 0x1f, 0xc0, 0x90, 0x20, 0xe4, 0x56, 0x1b, 0x69, 0xab, 0x47, 0x72, 0xab,
	0x46, 0xbf, 0x59, 0xd0, 0x4f, 0xee, 0xfd, 0x04, 0x9a, 0x19, 0xbb, 0x4a, 0x0f, 0xf6, 0x5b, 0x81,
	0x3d, 0xcf, 0x7d, 0x98, 0x51, 0xff, 0xa5, 0xdb, 0x2b, 0xd8, 0xf9, 0x16, 0x91, 0x99, 0x17, 0xcc,
	0x30, 0x65, 0xfe, 0x2c, 0xe4, 0x7f, 0x74, 0xcd, 0x59, 0x0d, 0xa2, 0x0e, 0x18, 0x8c, 0xe4, 0x39,
	0xa9, 0x3a, 0x72, 0x08, 0x21, 0xa8, 0x31, 0xff, 0x86, 0x9a, 0xf5, 0x8e, 0xd6, 0x6d, 0x38, 0x7c,
	0x6d, 0xff, 0x52, 0x01, 0xc9, 0x3c, 0x37, 0x36, 0xf9, 0x58, 0x44, 0xcf, 0x32, 0xa3, 0xd4, 0xb9,
	0x51, 0x8e, 0xb2, 0x17, 0x2a, 0x12, 0x2d, 0x33, 0xcb, 0xff, 0x34, 0x40, 0xff, 0xa7, 0x06, 0xfa,
	0x25, 0xf1, 0xdc, 0xd1, 0x3b, 0x34, 0x84, 0x27, 0xc9, 0xe8, 0xf4, 0x63, 0x8c, 0xda, 0xa2, 0xa2,
	0x30, 0x9b, 0xad, 0xe7, 0x6b, 0x71, 0xa1, 0xe7, 0x25, 0x6c, 0xcb, 0xd3, 0x0f, 0x59, 0x79, 0x62,
	0x71, 0xa2, 0x5a, 0x87, 0xa5, 0x67, 0x02, 0x68, 0x04, 0x86, 0x34, 0xc2, 0xd0, 0x41, 0x9e, 0x5b,
	0xf0, 0x9a, 0x65, 0x95, 0x1d, 0x09, 0x94, 0xb7, 0xd0, 0xc8, 0x26, 0x11, 0x92, 0x48, 0xaf, 0xcc,
	0x35, 0xcb, 0x5c, 0x3f, 0x10, 0xf5, 0x27, 0xa0, 0x7b, 0x2e, 0x57, 0xe2, 0x59, 0xe9, 0x3f, 0x6c,
	0xb5, 0x8b, 0x61, 0x51, 0x78, 0x0e, 0x90, 0x3f, 0x22, 0x32, 0x4b, 0xde, 0x35, 0xad, 0x3f, 0xb8,
	0xf7, 0xc5, 0xbf, 0xea, 0xfc, 0xe4, 0xf8, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x11, 0x99,
	0x1c, 0x2d, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GoTSDBClient is the client API for GoTSDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GoTSDBClient interface {
	KvsSave(ctx context.Context, in *KvsStoreRequest, opts ...grpc.CallOption) (*KvsStoreResponse, error)
	KvsKeyExists(ctx context.Context, in *KvsKeyExistsRequest, opts ...grpc.CallOption) (*KvsKeyExistsResponse, error)
	KvsRetrieve(ctx context.Context, in *KvsRetrieveRequest, opts ...grpc.CallOption) (*KvsRetrieveResponse, error)
	KvsDelete(ctx context.Context, in *KvsDeleteRequest, opts ...grpc.CallOption) (*KvsDeleteResponse, error)
	TSSave(ctx context.Context, in *TSStoreRequest, opts ...grpc.CallOption) (*TSStoreResponse, error)
	TSRetrieve(ctx context.Context, in *TSRetrieveRequest, opts ...grpc.CallOption) (*TSRetrieveResponse, error)
}

type goTSDBClient struct {
	cc grpc.ClientConnInterface
}

func NewGoTSDBClient(cc grpc.ClientConnInterface) GoTSDBClient {
	return &goTSDBClient{cc}
}

func (c *goTSDBClient) KvsSave(ctx context.Context, in *KvsStoreRequest, opts ...grpc.CallOption) (*KvsStoreResponse, error) {
	out := new(KvsStoreResponse)
	err := c.cc.Invoke(ctx, "/proto.GoTSDB/KvsSave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goTSDBClient) KvsKeyExists(ctx context.Context, in *KvsKeyExistsRequest, opts ...grpc.CallOption) (*KvsKeyExistsResponse, error) {
	out := new(KvsKeyExistsResponse)
	err := c.cc.Invoke(ctx, "/proto.GoTSDB/KvsKeyExists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goTSDBClient) KvsRetrieve(ctx context.Context, in *KvsRetrieveRequest, opts ...grpc.CallOption) (*KvsRetrieveResponse, error) {
	out := new(KvsRetrieveResponse)
	err := c.cc.Invoke(ctx, "/proto.GoTSDB/KvsRetrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goTSDBClient) KvsDelete(ctx context.Context, in *KvsDeleteRequest, opts ...grpc.CallOption) (*KvsDeleteResponse, error) {
	out := new(KvsDeleteResponse)
	err := c.cc.Invoke(ctx, "/proto.GoTSDB/KvsDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goTSDBClient) TSSave(ctx context.Context, in *TSStoreRequest, opts ...grpc.CallOption) (*TSStoreResponse, error) {
	out := new(TSStoreResponse)
	err := c.cc.Invoke(ctx, "/proto.GoTSDB/TSSave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goTSDBClient) TSRetrieve(ctx context.Context, in *TSRetrieveRequest, opts ...grpc.CallOption) (*TSRetrieveResponse, error) {
	out := new(TSRetrieveResponse)
	err := c.cc.Invoke(ctx, "/proto.GoTSDB/TSRetrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoTSDBServer is the server API for GoTSDB service.
type GoTSDBServer interface {
	KvsSave(context.Context, *KvsStoreRequest) (*KvsStoreResponse, error)
	KvsKeyExists(context.Context, *KvsKeyExistsRequest) (*KvsKeyExistsResponse, error)
	KvsRetrieve(context.Context, *KvsRetrieveRequest) (*KvsRetrieveResponse, error)
	KvsDelete(context.Context, *KvsDeleteRequest) (*KvsDeleteResponse, error)
	TSSave(context.Context, *TSStoreRequest) (*TSStoreResponse, error)
	TSRetrieve(context.Context, *TSRetrieveRequest) (*TSRetrieveResponse, error)
}

// UnimplementedGoTSDBServer can be embedded to have forward compatible implementations.
type UnimplementedGoTSDBServer struct {
}

func (*UnimplementedGoTSDBServer) KvsSave(ctx context.Context, req *KvsStoreRequest) (*KvsStoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KvsSave not implemented")
}
func (*UnimplementedGoTSDBServer) KvsKeyExists(ctx context.Context, req *KvsKeyExistsRequest) (*KvsKeyExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KvsKeyExists not implemented")
}
func (*UnimplementedGoTSDBServer) KvsRetrieve(ctx context.Context, req *KvsRetrieveRequest) (*KvsRetrieveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KvsRetrieve not implemented")
}
func (*UnimplementedGoTSDBServer) KvsDelete(ctx context.Context, req *KvsDeleteRequest) (*KvsDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KvsDelete not implemented")
}
func (*UnimplementedGoTSDBServer) TSSave(ctx context.Context, req *TSStoreRequest) (*TSStoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TSSave not implemented")
}
func (*UnimplementedGoTSDBServer) TSRetrieve(ctx context.Context, req *TSRetrieveRequest) (*TSRetrieveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TSRetrieve not implemented")
}

func RegisterGoTSDBServer(s *grpc.Server, srv GoTSDBServer) {
	s.RegisterService(&_GoTSDB_serviceDesc, srv)
}

func _GoTSDB_KvsSave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KvsStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoTSDBServer).KvsSave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GoTSDB/KvsSave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoTSDBServer).KvsSave(ctx, req.(*KvsStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoTSDB_KvsKeyExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KvsKeyExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoTSDBServer).KvsKeyExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GoTSDB/KvsKeyExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoTSDBServer).KvsKeyExists(ctx, req.(*KvsKeyExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoTSDB_KvsRetrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KvsRetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoTSDBServer).KvsRetrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GoTSDB/KvsRetrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoTSDBServer).KvsRetrieve(ctx, req.(*KvsRetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoTSDB_KvsDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KvsDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoTSDBServer).KvsDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GoTSDB/KvsDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoTSDBServer).KvsDelete(ctx, req.(*KvsDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoTSDB_TSSave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TSStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoTSDBServer).TSSave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GoTSDB/TSSave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoTSDBServer).TSSave(ctx, req.(*TSStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoTSDB_TSRetrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TSRetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoTSDBServer).TSRetrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GoTSDB/TSRetrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoTSDBServer).TSRetrieve(ctx, req.(*TSRetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GoTSDB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GoTSDB",
	HandlerType: (*GoTSDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "KvsSave",
			Handler:    _GoTSDB_KvsSave_Handler,
		},
		{
			MethodName: "KvsKeyExists",
			Handler:    _GoTSDB_KvsKeyExists_Handler,
		},
		{
			MethodName: "KvsRetrieve",
			Handler:    _GoTSDB_KvsRetrieve_Handler,
		},
		{
			MethodName: "KvsDelete",
			Handler:    _GoTSDB_KvsDelete_Handler,
		},
		{
			MethodName: "TSSave",
			Handler:    _GoTSDB_TSSave_Handler,
		},
		{
			MethodName: "TSRetrieve",
			Handler:    _GoTSDB_TSRetrieve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}
