// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: crpc/msg.proto

package crpc

import (
	cerror "github.com/chenjie199234/Corelib/cerror"
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

type MsgType int32

const (
	MsgType_CALL   MsgType = 0
	MsgType_CANCEL MsgType = 1
)

// Enum value maps for MsgType.
var (
	MsgType_name = map[int32]string{
		0: "CALL",
		1: "CANCEL",
	}
	MsgType_value = map[string]int32{
		"CALL":   0,
		"CANCEL": 1,
	}
)

func (x MsgType) Enum() *MsgType {
	p := new(MsgType)
	*p = x
	return p
}

func (x MsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_crpc_msg_proto_enumTypes[0].Descriptor()
}

func (MsgType) Type() protoreflect.EnumType {
	return &file_crpc_msg_proto_enumTypes[0]
}

func (x MsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgType.Descriptor instead.
func (MsgType) EnumDescriptor() ([]byte, []int) {
	return file_crpc_msg_proto_rawDescGZIP(), []int{0}
}

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Callid    uint64            `protobuf:"varint,1,opt,name=callid,proto3" json:"callid,omitempty"`
	Type      MsgType           `protobuf:"varint,2,opt,name=type,proto3,enum=crpc.MsgType" json:"type,omitempty"`
	Path      string            `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Deadline  int64             `protobuf:"varint,4,opt,name=deadline,proto3" json:"deadline,omitempty"`
	Body      []byte            `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	Error     *cerror.Error     `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
	Metadata  map[string]string `protobuf:"bytes,7,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Tracedata map[string]string `protobuf:"bytes,8,rep,name=tracedata,proto3" json:"tracedata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crpc_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_crpc_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_crpc_msg_proto_rawDescGZIP(), []int{0}
}

func (x *Msg) GetCallid() uint64 {
	if x != nil {
		return x.Callid
	}
	return 0
}

func (x *Msg) GetType() MsgType {
	if x != nil {
		return x.Type
	}
	return MsgType_CALL
}

func (x *Msg) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Msg) GetDeadline() int64 {
	if x != nil {
		return x.Deadline
	}
	return 0
}

func (x *Msg) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Msg) GetError() *cerror.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *Msg) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Msg) GetTracedata() map[string]string {
	if x != nil {
		return x.Tracedata
	}
	return nil
}

var File_crpc_msg_proto protoreflect.FileDescriptor

var file_crpc_msg_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x63, 0x72, 0x70, 0x63, 0x1a, 0x13, 0x63, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2f, 0x63,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x03, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x63, 0x72, 0x70, 0x63,
	0x2e, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x72, 0x70, 0x63,
	0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x36, 0x0a, 0x09,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x63, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x3c, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a,
	0x1f, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x41,
	0x4c, 0x4c, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x10, 0x01,
	0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x68, 0x65, 0x6e, 0x6a, 0x69, 0x65, 0x31, 0x39, 0x39, 0x32, 0x33, 0x34, 0x2f, 0x43, 0x6f, 0x72,
	0x65, 0x6c, 0x69, 0x62, 0x2f, 0x63, 0x72, 0x70, 0x63, 0x3b, 0x63, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crpc_msg_proto_rawDescOnce sync.Once
	file_crpc_msg_proto_rawDescData = file_crpc_msg_proto_rawDesc
)

func file_crpc_msg_proto_rawDescGZIP() []byte {
	file_crpc_msg_proto_rawDescOnce.Do(func() {
		file_crpc_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_crpc_msg_proto_rawDescData)
	})
	return file_crpc_msg_proto_rawDescData
}

var file_crpc_msg_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_crpc_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_crpc_msg_proto_goTypes = []interface{}{
	(MsgType)(0),         // 0: crpc.MsgType
	(*Msg)(nil),          // 1: crpc.msg
	nil,                  // 2: crpc.msg.MetadataEntry
	nil,                  // 3: crpc.msg.TracedataEntry
	(*cerror.Error)(nil), // 4: cerror.Error
}
var file_crpc_msg_proto_depIdxs = []int32{
	0, // 0: crpc.msg.type:type_name -> crpc.MsgType
	4, // 1: crpc.msg.error:type_name -> cerror.Error
	2, // 2: crpc.msg.metadata:type_name -> crpc.msg.MetadataEntry
	3, // 3: crpc.msg.tracedata:type_name -> crpc.msg.TracedataEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_crpc_msg_proto_init() }
func file_crpc_msg_proto_init() {
	if File_crpc_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crpc_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
			RawDescriptor: file_crpc_msg_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_crpc_msg_proto_goTypes,
		DependencyIndexes: file_crpc_msg_proto_depIdxs,
		EnumInfos:         file_crpc_msg_proto_enumTypes,
		MessageInfos:      file_crpc_msg_proto_msgTypes,
	}.Build()
	File_crpc_msg_proto = out.File
	file_crpc_msg_proto_rawDesc = nil
	file_crpc_msg_proto_goTypes = nil
	file_crpc_msg_proto_depIdxs = nil
}
