// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: badgerauth.proto

package pb

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

type Record_State int32

const (
	Record_CREATED Record_State = 0
)

// Enum value maps for Record_State.
var (
	Record_State_name = map[int32]string{
		0: "CREATED",
	}
	Record_State_value = map[string]int32{
		"CREATED": 0,
	}
)

func (x Record_State) Enum() *Record_State {
	p := new(Record_State)
	*p = x
	return p
}

func (x Record_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Record_State) Descriptor() protoreflect.EnumDescriptor {
	return file_badgerauth_proto_enumTypes[0].Descriptor()
}

func (Record_State) Type() protoreflect.EnumType {
	return &file_badgerauth_proto_enumTypes[0]
}

func (x Record_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Record_State.Descriptor instead.
func (Record_State) EnumDescriptor() ([]byte, []int) {
	return file_badgerauth_proto_rawDescGZIP(), []int{0, 0}
}

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// record data
	CreatedAtUnix int64 `protobuf:"varint,1,opt,name=created_at_unix,json=createdAtUnix,proto3" json:"created_at_unix,omitempty"`
	Public        bool  `protobuf:"varint,2,opt,name=public,proto3" json:"public,omitempty"`
	// denormalized information from access grant
	SatelliteAddress string `protobuf:"bytes,3,opt,name=satellite_address,json=satelliteAddress,proto3" json:"satellite_address,omitempty"`
	MacaroonHead     []byte `protobuf:"bytes,4,opt,name=macaroon_head,json=macaroonHead,proto3" json:"macaroon_head,omitempty"`
	ExpiresAtUnix    int64  `protobuf:"varint,5,opt,name=expires_at_unix,json=expiresAtUnix,proto3" json:"expires_at_unix,omitempty"`
	// sensitive data
	EncryptedSecretKey   []byte `protobuf:"bytes,6,opt,name=encrypted_secret_key,json=encryptedSecretKey,proto3" json:"encrypted_secret_key,omitempty"`
	EncryptedAccessGrant []byte `protobuf:"bytes,7,opt,name=encrypted_access_grant,json=encryptedAccessGrant,proto3" json:"encrypted_access_grant,omitempty"`
	// invalidation tracking
	InvalidationReason string `protobuf:"bytes,8,opt,name=invalidation_reason,json=invalidationReason,proto3" json:"invalidation_reason,omitempty"`
	InvalidatedAtUnix  int64  `protobuf:"varint,9,opt,name=invalidated_at_unix,json=invalidatedAtUnix,proto3" json:"invalidated_at_unix,omitempty"`
	// synchronization-related data
	State Record_State `protobuf:"varint,10,opt,name=state,proto3,enum=badgerauth.Record_State" json:"state,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_badgerauth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_badgerauth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_badgerauth_proto_rawDescGZIP(), []int{0}
}

func (x *Record) GetCreatedAtUnix() int64 {
	if x != nil {
		return x.CreatedAtUnix
	}
	return 0
}

func (x *Record) GetPublic() bool {
	if x != nil {
		return x.Public
	}
	return false
}

func (x *Record) GetSatelliteAddress() string {
	if x != nil {
		return x.SatelliteAddress
	}
	return ""
}

func (x *Record) GetMacaroonHead() []byte {
	if x != nil {
		return x.MacaroonHead
	}
	return nil
}

func (x *Record) GetExpiresAtUnix() int64 {
	if x != nil {
		return x.ExpiresAtUnix
	}
	return 0
}

func (x *Record) GetEncryptedSecretKey() []byte {
	if x != nil {
		return x.EncryptedSecretKey
	}
	return nil
}

func (x *Record) GetEncryptedAccessGrant() []byte {
	if x != nil {
		return x.EncryptedAccessGrant
	}
	return nil
}

func (x *Record) GetInvalidationReason() string {
	if x != nil {
		return x.InvalidationReason
	}
	return ""
}

func (x *Record) GetInvalidatedAtUnix() int64 {
	if x != nil {
		return x.InvalidatedAtUnix
	}
	return 0
}

func (x *Record) GetState() Record_State {
	if x != nil {
		return x.State
	}
	return Record_CREATED
}

type ReplicationRequestEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthserviceId string `protobuf:"bytes,1,opt,name=authservice_id,json=authserviceId,proto3" json:"authservice_id,omitempty"`
	Counter       uint64 `protobuf:"varint,2,opt,name=counter,proto3" json:"counter,omitempty"`
}

func (x *ReplicationRequestEntry) Reset() {
	*x = ReplicationRequestEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_badgerauth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplicationRequestEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplicationRequestEntry) ProtoMessage() {}

func (x *ReplicationRequestEntry) ProtoReflect() protoreflect.Message {
	mi := &file_badgerauth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplicationRequestEntry.ProtoReflect.Descriptor instead.
func (*ReplicationRequestEntry) Descriptor() ([]byte, []int) {
	return file_badgerauth_proto_rawDescGZIP(), []int{1}
}

func (x *ReplicationRequestEntry) GetAuthserviceId() string {
	if x != nil {
		return x.AuthserviceId
	}
	return ""
}

func (x *ReplicationRequestEntry) GetCounter() uint64 {
	if x != nil {
		return x.Counter
	}
	return 0
}

type ReplicationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthToken string                     `protobuf:"bytes,1,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
	Entries   []*ReplicationRequestEntry `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *ReplicationRequest) Reset() {
	*x = ReplicationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_badgerauth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplicationRequest) ProtoMessage() {}

func (x *ReplicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_badgerauth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplicationRequest.ProtoReflect.Descriptor instead.
func (*ReplicationRequest) Descriptor() ([]byte, []int) {
	return file_badgerauth_proto_rawDescGZIP(), []int{2}
}

func (x *ReplicationRequest) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

func (x *ReplicationRequest) GetEntries() []*ReplicationRequestEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type ReplicationResponseEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthserviceId     string  `protobuf:"bytes,1,opt,name=authservice_id,json=authserviceId,proto3" json:"authservice_id,omitempty"`
	Counter           uint64  `protobuf:"varint,2,opt,name=counter,proto3" json:"counter,omitempty"`
	EncryptionKeyHash []byte  `protobuf:"bytes,3,opt,name=encryption_key_hash,json=encryptionKeyHash,proto3" json:"encryption_key_hash,omitempty"`
	Record            *Record `protobuf:"bytes,4,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *ReplicationResponseEntry) Reset() {
	*x = ReplicationResponseEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_badgerauth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplicationResponseEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplicationResponseEntry) ProtoMessage() {}

func (x *ReplicationResponseEntry) ProtoReflect() protoreflect.Message {
	mi := &file_badgerauth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplicationResponseEntry.ProtoReflect.Descriptor instead.
func (*ReplicationResponseEntry) Descriptor() ([]byte, []int) {
	return file_badgerauth_proto_rawDescGZIP(), []int{3}
}

func (x *ReplicationResponseEntry) GetAuthserviceId() string {
	if x != nil {
		return x.AuthserviceId
	}
	return ""
}

func (x *ReplicationResponseEntry) GetCounter() uint64 {
	if x != nil {
		return x.Counter
	}
	return 0
}

func (x *ReplicationResponseEntry) GetEncryptionKeyHash() []byte {
	if x != nil {
		return x.EncryptionKeyHash
	}
	return nil
}

func (x *ReplicationResponseEntry) GetRecord() *Record {
	if x != nil {
		return x.Record
	}
	return nil
}

type ReplicationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*ReplicationResponseEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *ReplicationResponse) Reset() {
	*x = ReplicationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_badgerauth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplicationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplicationResponse) ProtoMessage() {}

func (x *ReplicationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_badgerauth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplicationResponse.ProtoReflect.Descriptor instead.
func (*ReplicationResponse) Descriptor() ([]byte, []int) {
	return file_badgerauth_proto_rawDescGZIP(), []int{4}
}

func (x *ReplicationResponse) GetEntries() []*ReplicationResponseEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

var File_badgerauth_proto protoreflect.FileDescriptor

var file_badgerauth_proto_rawDesc = []byte{
	0x0a, 0x10, 0x62, 0x61, 0x64, 0x67, 0x65, 0x72, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x62, 0x61, 0x64, 0x67, 0x65, 0x72, 0x61, 0x75, 0x74, 0x68, 0x22, 0xd1,
	0x03, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x5f, 0x75, 0x6e, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x55, 0x6e, 0x69,
	0x78, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x61, 0x74,
	0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x63, 0x61, 0x72, 0x6f,
	0x6f, 0x6e, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6d,
	0x61, 0x63, 0x61, 0x72, 0x6f, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x5f, 0x75, 0x6e, 0x69, 0x78, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x55,
	0x6e, 0x69, 0x78, 0x12, 0x30, 0x0a, 0x14, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64,
	0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x12, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x16, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x14, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x69,
	0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x13,
	0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x5f, 0x75,
	0x6e, 0x69, 0x78, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x69, 0x6e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x55, 0x6e, 0x69, 0x78, 0x12, 0x2e, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x62, 0x61,
	0x64, 0x67, 0x65, 0x72, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x14, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44,
	0x10, 0x00, 0x22, 0x5a, 0x0a, 0x17, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x25, 0x0a,
	0x0e, 0x61, 0x75, 0x74, 0x68, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x72,
	0x0a, 0x12, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x3d, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x62, 0x61, 0x64, 0x67, 0x65, 0x72, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x22, 0xb7, 0x01, 0x0a, 0x18, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x25, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x12, 0x2e, 0x0a, 0x13, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b,
	0x65, 0x79, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x65,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x2a, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x62, 0x61, 0x64, 0x67, 0x65, 0x72, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x55, 0x0a, 0x13,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x62, 0x61, 0x64, 0x67, 0x65, 0x72, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x32, 0x62, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x09, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x62, 0x61, 0x64, 0x67, 0x65, 0x72, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x62, 0x61, 0x64, 0x67, 0x65, 0x72, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2c, 0x5a, 0x2a, 0x73, 0x74, 0x6f, 0x72, 0x6a,
	0x2e, 0x69, 0x6f, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2d, 0x6d, 0x74, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x62, 0x61, 0x64, 0x67, 0x65, 0x72, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_badgerauth_proto_rawDescOnce sync.Once
	file_badgerauth_proto_rawDescData = file_badgerauth_proto_rawDesc
)

func file_badgerauth_proto_rawDescGZIP() []byte {
	file_badgerauth_proto_rawDescOnce.Do(func() {
		file_badgerauth_proto_rawDescData = protoimpl.X.CompressGZIP(file_badgerauth_proto_rawDescData)
	})
	return file_badgerauth_proto_rawDescData
}

var file_badgerauth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_badgerauth_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_badgerauth_proto_goTypes = []interface{}{
	(Record_State)(0),                // 0: badgerauth.Record.State
	(*Record)(nil),                   // 1: badgerauth.Record
	(*ReplicationRequestEntry)(nil),  // 2: badgerauth.ReplicationRequestEntry
	(*ReplicationRequest)(nil),       // 3: badgerauth.ReplicationRequest
	(*ReplicationResponseEntry)(nil), // 4: badgerauth.ReplicationResponseEntry
	(*ReplicationResponse)(nil),      // 5: badgerauth.ReplicationResponse
}
var file_badgerauth_proto_depIdxs = []int32{
	0, // 0: badgerauth.Record.state:type_name -> badgerauth.Record.State
	2, // 1: badgerauth.ReplicationRequest.entries:type_name -> badgerauth.ReplicationRequestEntry
	1, // 2: badgerauth.ReplicationResponseEntry.record:type_name -> badgerauth.Record
	4, // 3: badgerauth.ReplicationResponse.entries:type_name -> badgerauth.ReplicationResponseEntry
	3, // 4: badgerauth.ReplicationService.Replicate:input_type -> badgerauth.ReplicationRequest
	5, // 5: badgerauth.ReplicationService.Replicate:output_type -> badgerauth.ReplicationResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_badgerauth_proto_init() }
func file_badgerauth_proto_init() {
	if File_badgerauth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_badgerauth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
		file_badgerauth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplicationRequestEntry); i {
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
		file_badgerauth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplicationRequest); i {
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
		file_badgerauth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplicationResponseEntry); i {
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
		file_badgerauth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplicationResponse); i {
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
			RawDescriptor: file_badgerauth_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_badgerauth_proto_goTypes,
		DependencyIndexes: file_badgerauth_proto_depIdxs,
		EnumInfos:         file_badgerauth_proto_enumTypes,
		MessageInfos:      file_badgerauth_proto_msgTypes,
	}.Build()
	File_badgerauth_proto = out.File
	file_badgerauth_proto_rawDesc = nil
	file_badgerauth_proto_goTypes = nil
	file_badgerauth_proto_depIdxs = nil
}
