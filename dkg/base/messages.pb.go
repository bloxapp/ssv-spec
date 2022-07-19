// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: base/messages.proto

package base

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

type MessageHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId []byte `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	MsgType   int32  `protobuf:"varint,2,opt,name=msg_type,json=msgType,proto3" json:"msg_type,omitempty"`
	Sender    uint64 `protobuf:"varint,3,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver  uint64 `protobuf:"varint,4,opt,name=receiver,proto3" json:"receiver,omitempty"`
}

func (x *MessageHeader) Reset() {
	*x = MessageHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageHeader) ProtoMessage() {}

func (x *MessageHeader) ProtoReflect() protoreflect.Message {
	mi := &file_base_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageHeader.ProtoReflect.Descriptor instead.
func (*MessageHeader) Descriptor() ([]byte, []int) {
	return file_base_messages_proto_rawDescGZIP(), []int{0}
}

func (x *MessageHeader) GetSessionId() []byte {
	if x != nil {
		return x.SessionId
	}
	return nil
}

func (x *MessageHeader) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *MessageHeader) GetSender() uint64 {
	if x != nil {
		return x.Sender
	}
	return 0
}

func (x *MessageHeader) GetReceiver() uint64 {
	if x != nil {
		return x.Receiver
	}
	return 0
}

//
// A generic message.
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header    *MessageHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data      []byte         `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Signature []byte         `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_base_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_base_messages_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetHeader() *MessageHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Message) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Message) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type LocalKeyShare struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index           uint64   `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	Threshold       uint64   `protobuf:"varint,2,opt,name=Threshold,proto3" json:"Threshold,omitempty"`
	Committee       []uint64 `protobuf:"varint,3,rep,packed,name=Committee,proto3" json:"Committee,omitempty"`
	SharePublicKeys [][]byte `protobuf:"bytes,4,rep,name=SharePublicKeys,proto3" json:"SharePublicKeys,omitempty"`
	PublicKey       []byte   `protobuf:"bytes,5,opt,name=PublicKey,proto3" json:"PublicKey,omitempty"`
	SecretShare     []byte   `protobuf:"bytes,6,opt,name=SecretShare,proto3" json:"SecretShare,omitempty"`
}

func (x *LocalKeyShare) Reset() {
	*x = LocalKeyShare{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalKeyShare) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalKeyShare) ProtoMessage() {}

func (x *LocalKeyShare) ProtoReflect() protoreflect.Message {
	mi := &file_base_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalKeyShare.ProtoReflect.Descriptor instead.
func (*LocalKeyShare) Descriptor() ([]byte, []int) {
	return file_base_messages_proto_rawDescGZIP(), []int{2}
}

func (x *LocalKeyShare) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *LocalKeyShare) GetThreshold() uint64 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

func (x *LocalKeyShare) GetCommittee() []uint64 {
	if x != nil {
		return x.Committee
	}
	return nil
}

func (x *LocalKeyShare) GetSharePublicKeys() [][]byte {
	if x != nil {
		return x.SharePublicKeys
	}
	return nil
}

func (x *LocalKeyShare) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *LocalKeyShare) GetSecretShare() []byte {
	if x != nil {
		return x.SecretShare
	}
	return nil
}

var File_base_messages_proto protoreflect.FileDescriptor

var file_base_messages_proto_rawDesc = []byte{
	0x0a, 0x13, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x73, 0x76, 0x2e, 0x64, 0x6b, 0x67, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x22, 0x7d, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x22, 0x70, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x73, 0x73, 0x76, 0x2e, 0x64, 0x6b, 0x67, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x22, 0xcb, 0x01, 0x0a, 0x0d, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4b, 0x65,
	0x79, 0x53, 0x68, 0x61, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1c, 0x0a, 0x09,
	0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x04, 0x52, 0x09, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0c, 0x52, 0x0f, 0x53, 0x68, 0x61, 0x72, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x6c, 0x6f, 0x78, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x73, 0x76, 0x2d, 0x73, 0x70, 0x65,
	0x63, 0x2f, 0x64, 0x6b, 0x67, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_base_messages_proto_rawDescOnce sync.Once
	file_base_messages_proto_rawDescData = file_base_messages_proto_rawDesc
)

func file_base_messages_proto_rawDescGZIP() []byte {
	file_base_messages_proto_rawDescOnce.Do(func() {
		file_base_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_messages_proto_rawDescData)
	})
	return file_base_messages_proto_rawDescData
}

var file_base_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_base_messages_proto_goTypes = []interface{}{
	(*MessageHeader)(nil), // 0: ssv.dkg.base.MessageHeader
	(*Message)(nil),       // 1: ssv.dkg.base.Message
	(*LocalKeyShare)(nil), // 2: ssv.dkg.base.LocalKeyShare
}
var file_base_messages_proto_depIdxs = []int32{
	0, // 0: ssv.dkg.base.Message.header:type_name -> ssv.dkg.base.MessageHeader
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_base_messages_proto_init() }
func file_base_messages_proto_init() {
	if File_base_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_base_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageHeader); i {
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
		file_base_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_base_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalKeyShare); i {
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
			RawDescriptor: file_base_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_messages_proto_goTypes,
		DependencyIndexes: file_base_messages_proto_depIdxs,
		MessageInfos:      file_base_messages_proto_msgTypes,
	}.Build()
	File_base_messages_proto = out.File
	file_base_messages_proto_rawDesc = nil
	file_base_messages_proto_goTypes = nil
	file_base_messages_proto_depIdxs = nil
}
