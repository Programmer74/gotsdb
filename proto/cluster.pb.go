// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: proto/cluster.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_proto_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_proto_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *PingResponse) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectionString string `protobuf:"bytes,1,opt,name=connectionString,proto3" json:"connectionString,omitempty"`
	Uuid             string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_proto_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *Node) GetConnectionString() string {
	if x != nil {
		return x.ConnectionString
	}
	return ""
}

func (x *Node) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Iam *Node `protobuf:"bytes,1,opt,name=iam,proto3" json:"iam,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_proto_cluster_proto_rawDescGZIP(), []int{3}
}

func (x *HelloRequest) GetIam() *Node {
	if x != nil {
		return x.Iam
	}
	return nil
}

type AliveNodesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AliveNodes []*Node `protobuf:"bytes,1,rep,name=aliveNodes,proto3" json:"aliveNodes,omitempty"`
}

func (x *AliveNodesResponse) Reset() {
	*x = AliveNodesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cluster_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AliveNodesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AliveNodesResponse) ProtoMessage() {}

func (x *AliveNodesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cluster_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AliveNodesResponse.ProtoReflect.Descriptor instead.
func (*AliveNodesResponse) Descriptor() ([]byte, []int) {
	return file_proto_cluster_proto_rawDescGZIP(), []int{4}
}

func (x *AliveNodesResponse) GetAliveNodes() []*Node {
	if x != nil {
		return x.AliveNodes
	}
	return nil
}

var File_proto_cluster_proto protoreflect.FileDescriptor

var file_proto_cluster_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a,
	0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x28, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x22, 0x46, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x03, 0x69, 0x61, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x03, 0x69, 0x61, 0x6d, 0x22, 0x41, 0x0a, 0x12, 0x41, 0x6c, 0x69, 0x76, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a,
	0x0a, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x0a,
	0x61, 0x6c, 0x69, 0x76, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x32, 0x88, 0x06, 0x0a, 0x07, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x6c,
	0x69, 0x76, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x37, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x73, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x6f, 0x69, 0x64, 0x1a, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x4b, 0x76, 0x73,
	0x53, 0x61, 0x76, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x76, 0x73,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x76, 0x73, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x4b, 0x76, 0x73, 0x4b, 0x65, 0x79, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x76,
	0x73, 0x4b, 0x65, 0x79, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x76, 0x73, 0x4b, 0x65, 0x79,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44,
	0x0a, 0x0b, 0x4b, 0x76, 0x73, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x12, 0x19, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x76, 0x73, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4b, 0x76, 0x73, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x4b, 0x76, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x76, 0x73, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4b, 0x76, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x4b, 0x76, 0x73, 0x47, 0x65, 0x74, 0x4b, 0x65,
	0x79, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x76, 0x73, 0x41, 0x6c,
	0x6c, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x76, 0x73, 0x41, 0x6c, 0x6c, 0x4b, 0x65, 0x79, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x54, 0x53, 0x53, 0x61, 0x76,
	0x65, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x53, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x53, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x41, 0x0a, 0x0b, 0x54, 0x53, 0x53, 0x61, 0x76, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x53, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x53, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x54, 0x53, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x53, 0x52, 0x65, 0x74, 0x72,
	0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x53, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x54, 0x53, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x53, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x53, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_cluster_proto_rawDescOnce sync.Once
	file_proto_cluster_proto_rawDescData = file_proto_cluster_proto_rawDesc
)

func file_proto_cluster_proto_rawDescGZIP() []byte {
	file_proto_cluster_proto_rawDescOnce.Do(func() {
		file_proto_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cluster_proto_rawDescData)
	})
	return file_proto_cluster_proto_rawDescData
}

var file_proto_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_cluster_proto_goTypes = []interface{}{
	(*PingRequest)(nil),            // 0: proto.PingRequest
	(*PingResponse)(nil),           // 1: proto.PingResponse
	(*Node)(nil),                   // 2: proto.Node
	(*HelloRequest)(nil),           // 3: proto.HelloRequest
	(*AliveNodesResponse)(nil),     // 4: proto.AliveNodesResponse
	(*Void)(nil),                   // 5: proto.void
	(*KvsStoreRequest)(nil),        // 6: proto.KvsStoreRequest
	(*KvsKeyExistsRequest)(nil),    // 7: proto.KvsKeyExistsRequest
	(*KvsRetrieveRequest)(nil),     // 8: proto.KvsRetrieveRequest
	(*KvsDeleteRequest)(nil),       // 9: proto.KvsDeleteRequest
	(*KvsAllKeysRequest)(nil),      // 10: proto.KvsAllKeysRequest
	(*TSStoreRequest)(nil),         // 11: proto.TSStoreRequest
	(*TSStoreBatchRequest)(nil),    // 12: proto.TSStoreBatchRequest
	(*TSRetrieveRequest)(nil),      // 13: proto.TSRetrieveRequest
	(*TSAvailabilityRequest)(nil),  // 14: proto.TSAvailabilityRequest
	(*KvsStoreResponse)(nil),       // 15: proto.KvsStoreResponse
	(*KvsKeyExistsResponse)(nil),   // 16: proto.KvsKeyExistsResponse
	(*KvsRetrieveResponse)(nil),    // 17: proto.KvsRetrieveResponse
	(*KvsDeleteResponse)(nil),      // 18: proto.KvsDeleteResponse
	(*KvsAllKeysResponse)(nil),     // 19: proto.KvsAllKeysResponse
	(*TSStoreResponse)(nil),        // 20: proto.TSStoreResponse
	(*TSRetrieveResponse)(nil),     // 21: proto.TSRetrieveResponse
	(*TSAvailabilityResponse)(nil), // 22: proto.TSAvailabilityResponse
}
var file_proto_cluster_proto_depIdxs = []int32{
	2,  // 0: proto.HelloRequest.iam:type_name -> proto.Node
	2,  // 1: proto.AliveNodesResponse.aliveNodes:type_name -> proto.Node
	0,  // 2: proto.Cluster.Ping:input_type -> proto.PingRequest
	3,  // 3: proto.Cluster.Hello:input_type -> proto.HelloRequest
	5,  // 4: proto.Cluster.GetAliveNodes:input_type -> proto.void
	6,  // 5: proto.Cluster.KvsSave:input_type -> proto.KvsStoreRequest
	7,  // 6: proto.Cluster.KvsKeyExists:input_type -> proto.KvsKeyExistsRequest
	8,  // 7: proto.Cluster.KvsRetrieve:input_type -> proto.KvsRetrieveRequest
	9,  // 8: proto.Cluster.KvsDelete:input_type -> proto.KvsDeleteRequest
	10, // 9: proto.Cluster.KvsGetKeys:input_type -> proto.KvsAllKeysRequest
	11, // 10: proto.Cluster.TSSave:input_type -> proto.TSStoreRequest
	12, // 11: proto.Cluster.TSSaveBatch:input_type -> proto.TSStoreBatchRequest
	13, // 12: proto.Cluster.TSRetrieve:input_type -> proto.TSRetrieveRequest
	14, // 13: proto.Cluster.TSAvailability:input_type -> proto.TSAvailabilityRequest
	1,  // 14: proto.Cluster.Ping:output_type -> proto.PingResponse
	4,  // 15: proto.Cluster.Hello:output_type -> proto.AliveNodesResponse
	4,  // 16: proto.Cluster.GetAliveNodes:output_type -> proto.AliveNodesResponse
	15, // 17: proto.Cluster.KvsSave:output_type -> proto.KvsStoreResponse
	16, // 18: proto.Cluster.KvsKeyExists:output_type -> proto.KvsKeyExistsResponse
	17, // 19: proto.Cluster.KvsRetrieve:output_type -> proto.KvsRetrieveResponse
	18, // 20: proto.Cluster.KvsDelete:output_type -> proto.KvsDeleteResponse
	19, // 21: proto.Cluster.KvsGetKeys:output_type -> proto.KvsAllKeysResponse
	20, // 22: proto.Cluster.TSSave:output_type -> proto.TSStoreResponse
	20, // 23: proto.Cluster.TSSaveBatch:output_type -> proto.TSStoreResponse
	21, // 24: proto.Cluster.TSRetrieve:output_type -> proto.TSRetrieveResponse
	22, // 25: proto.Cluster.TSAvailability:output_type -> proto.TSAvailabilityResponse
	14, // [14:26] is the sub-list for method output_type
	2,  // [2:14] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_cluster_proto_init() }
func file_proto_cluster_proto_init() {
	if File_proto_cluster_proto != nil {
		return
	}
	file_proto_rpc_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_cluster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_cluster_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AliveNodesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cluster_proto_goTypes,
		DependencyIndexes: file_proto_cluster_proto_depIdxs,
		MessageInfos:      file_proto_cluster_proto_msgTypes,
	}.Build()
	File_proto_cluster_proto = out.File
	file_proto_cluster_proto_rawDesc = nil
	file_proto_cluster_proto_goTypes = nil
	file_proto_cluster_proto_depIdxs = nil
}
