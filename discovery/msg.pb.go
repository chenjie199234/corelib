// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: msg.proto

package discovery

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MsgType int32

const (
	MsgType_Reg   MsgType = 0
	MsgType_UnReg MsgType = 1
	MsgType_Watch MsgType = 2
	MsgType_Push  MsgType = 3
)

// Enum value maps for MsgType.
var (
	MsgType_name = map[int32]string{
		0: "Reg",
		1: "UnReg",
		2: "Watch",
		3: "Push",
	}
	MsgType_value = map[string]int32{
		"Reg":   0,
		"UnReg": 1,
		"Watch": 2,
		"Push":  3,
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
	return file_msg_proto_enumTypes[0].Descriptor()
}

func (MsgType) Type() protoreflect.EnumType {
	return &file_msg_proto_enumTypes[0]
}

func (x MsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgType.Descriptor instead.
func (MsgType) EnumDescriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0}
}

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgType MsgType `protobuf:"varint,1,opt,name=msg_type,json=msgType,proto3,enum=discovery.MsgType" json:"msg_type,omitempty"`
	// Types that are assignable to MsgContent:
	//	*Msg_RegMsg
	//	*Msg_UnregMsg
	//	*Msg_WatchMsg
	//	*Msg_PushMsg
	MsgContent isMsg_MsgContent `protobuf_oneof:"msg_content"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[0]
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
	return file_msg_proto_rawDescGZIP(), []int{0}
}

func (x *Msg) GetMsgType() MsgType {
	if x != nil {
		return x.MsgType
	}
	return MsgType_Reg
}

func (m *Msg) GetMsgContent() isMsg_MsgContent {
	if m != nil {
		return m.MsgContent
	}
	return nil
}

func (x *Msg) GetRegMsg() *RegMsg {
	if x, ok := x.GetMsgContent().(*Msg_RegMsg); ok {
		return x.RegMsg
	}
	return nil
}

func (x *Msg) GetUnregMsg() *UnregMsg {
	if x, ok := x.GetMsgContent().(*Msg_UnregMsg); ok {
		return x.UnregMsg
	}
	return nil
}

func (x *Msg) GetWatchMsg() *WatchMsg {
	if x, ok := x.GetMsgContent().(*Msg_WatchMsg); ok {
		return x.WatchMsg
	}
	return nil
}

func (x *Msg) GetPushMsg() *PushMsg {
	if x, ok := x.GetMsgContent().(*Msg_PushMsg); ok {
		return x.PushMsg
	}
	return nil
}

type isMsg_MsgContent interface {
	isMsg_MsgContent()
}

type Msg_RegMsg struct {
	RegMsg *RegMsg `protobuf:"bytes,2,opt,name=reg_msg,json=regMsg,proto3,oneof"`
}

type Msg_UnregMsg struct {
	UnregMsg *UnregMsg `protobuf:"bytes,3,opt,name=unreg_msg,json=unregMsg,proto3,oneof"`
}

type Msg_WatchMsg struct {
	WatchMsg *WatchMsg `protobuf:"bytes,4,opt,name=watch_msg,json=watchMsg,proto3,oneof"`
}

type Msg_PushMsg struct {
	PushMsg *PushMsg `protobuf:"bytes,5,opt,name=push_msg,json=pushMsg,proto3,oneof"`
}

func (*Msg_RegMsg) isMsg_MsgContent() {}

func (*Msg_UnregMsg) isMsg_MsgContent() {}

func (*Msg_WatchMsg) isMsg_MsgContent() {}

func (*Msg_PushMsg) isMsg_MsgContent() {}

type RegInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WebScheme string `protobuf:"bytes,1,opt,name=web_scheme,json=webScheme,proto3" json:"web_scheme,omitempty"`
	WebIp     string `protobuf:"bytes,2,opt,name=web_ip,json=webIp,proto3" json:"web_ip,omitempty"`
	WebPort   int64  `protobuf:"varint,3,opt,name=web_port,json=webPort,proto3" json:"web_port,omitempty"`
	RpcIp     string `protobuf:"bytes,4,opt,name=rpc_ip,json=rpcIp,proto3" json:"rpc_ip,omitempty"`
	RpcPort   int64  `protobuf:"varint,5,opt,name=rpc_port,json=rpcPort,proto3" json:"rpc_port,omitempty"`
	Addition  []byte `protobuf:"bytes,6,opt,name=addition,proto3" json:"addition,omitempty"`
}

func (x *RegInfo) Reset() {
	*x = RegInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegInfo) ProtoMessage() {}

func (x *RegInfo) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegInfo.ProtoReflect.Descriptor instead.
func (*RegInfo) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{1}
}

func (x *RegInfo) GetWebScheme() string {
	if x != nil {
		return x.WebScheme
	}
	return ""
}

func (x *RegInfo) GetWebIp() string {
	if x != nil {
		return x.WebIp
	}
	return ""
}

func (x *RegInfo) GetWebPort() int64 {
	if x != nil {
		return x.WebPort
	}
	return 0
}

func (x *RegInfo) GetRpcIp() string {
	if x != nil {
		return x.RpcIp
	}
	return ""
}

func (x *RegInfo) GetRpcPort() int64 {
	if x != nil {
		return x.RpcPort
	}
	return 0
}

func (x *RegInfo) GetAddition() []byte {
	if x != nil {
		return x.Addition
	}
	return nil
}

type RegMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppUniqueName string   `protobuf:"bytes,1,opt,name=app_unique_name,json=appUniqueName,proto3" json:"app_unique_name,omitempty"`
	RegInfo       *RegInfo `protobuf:"bytes,2,opt,name=reg_info,json=regInfo,proto3" json:"reg_info,omitempty"`
}

func (x *RegMsg) Reset() {
	*x = RegMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegMsg) ProtoMessage() {}

func (x *RegMsg) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegMsg.ProtoReflect.Descriptor instead.
func (*RegMsg) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{2}
}

func (x *RegMsg) GetAppUniqueName() string {
	if x != nil {
		return x.AppUniqueName
	}
	return ""
}

func (x *RegMsg) GetRegInfo() *RegInfo {
	if x != nil {
		return x.RegInfo
	}
	return nil
}

type UnregMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppUniqueName string `protobuf:"bytes,1,opt,name=app_unique_name,json=appUniqueName,proto3" json:"app_unique_name,omitempty"`
}

func (x *UnregMsg) Reset() {
	*x = UnregMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnregMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnregMsg) ProtoMessage() {}

func (x *UnregMsg) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnregMsg.ProtoReflect.Descriptor instead.
func (*UnregMsg) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{3}
}

func (x *UnregMsg) GetAppUniqueName() string {
	if x != nil {
		return x.AppUniqueName
	}
	return ""
}

type WatchMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppName string `protobuf:"bytes,1,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
}

func (x *WatchMsg) Reset() {
	*x = WatchMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchMsg) ProtoMessage() {}

func (x *WatchMsg) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchMsg.ProtoReflect.Descriptor instead.
func (*WatchMsg) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{4}
}

func (x *WatchMsg) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

type PushMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppName string              `protobuf:"bytes,1,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	Apps    map[string]*RegInfo `protobuf:"bytes,2,rep,name=apps,proto3" json:"apps,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PushMsg) Reset() {
	*x = PushMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsg) ProtoMessage() {}

func (x *PushMsg) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsg.ProtoReflect.Descriptor instead.
func (*PushMsg) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{5}
}

func (x *PushMsg) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *PushMsg) GetApps() map[string]*RegInfo {
	if x != nil {
		return x.Apps
	}
	return nil
}

var File_msg_proto protoreflect.FileDescriptor

var file_msg_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x22, 0x8a, 0x02, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2d,
	0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x12, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x4d, 0x73, 0x67,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a,
	0x07, 0x72, 0x65, 0x67, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x67, 0x4d, 0x73,
	0x67, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x67, 0x4d, 0x73, 0x67, 0x12, 0x32, 0x0a, 0x09, 0x75,
	0x6e, 0x72, 0x65, 0x67, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x55, 0x6e, 0x72, 0x65, 0x67,
	0x4d, 0x73, 0x67, 0x48, 0x00, 0x52, 0x08, 0x75, 0x6e, 0x72, 0x65, 0x67, 0x4d, 0x73, 0x67, 0x12,
	0x32, 0x0a, 0x09, 0x77, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x4d, 0x73, 0x67, 0x48, 0x00, 0x52, 0x08, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x4d, 0x73, 0x67, 0x12, 0x2f, 0x0a, 0x08, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x6d, 0x73, 0x67, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x48, 0x00, 0x52, 0x07, 0x70, 0x75, 0x73,
	0x68, 0x4d, 0x73, 0x67, 0x42, 0x0d, 0x0a, 0x0b, 0x6d, 0x73, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0xa8, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1d, 0x0a, 0x0a, 0x77, 0x65, 0x62, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x65, 0x62, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x15,
	0x0a, 0x06, 0x77, 0x65, 0x62, 0x5f, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x77, 0x65, 0x62, 0x49, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x77, 0x65, 0x62, 0x5f, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x77, 0x65, 0x62, 0x50, 0x6f, 0x72, 0x74,
	0x12, 0x15, 0x0a, 0x06, 0x72, 0x70, 0x63, 0x5f, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x72, 0x70, 0x63, 0x49, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x70, 0x63, 0x5f, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x70, 0x63, 0x50, 0x6f,
	0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5f,
	0x0a, 0x06, 0x52, 0x65, 0x67, 0x4d, 0x73, 0x67, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x5f,
	0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2d, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x52,
	0x65, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x65, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x32, 0x0a, 0x08, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x4d, 0x73, 0x67, 0x12, 0x26, 0x0a, 0x0f, 0x61,
	0x70, 0x70, 0x5f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x25, 0x0a, 0x08, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x73, 0x67, 0x12,
	0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xa3, 0x01, 0x0a, 0x07, 0x50,
	0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x30, 0x0a, 0x04, 0x61, 0x70, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x50, 0x75, 0x73, 0x68,
	0x4d, 0x73, 0x67, 0x2e, 0x41, 0x70, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x61,
	0x70, 0x70, 0x73, 0x1a, 0x4b, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x52, 0x65,
	0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x2a, 0x32, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x52,
	0x65, 0x67, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x6e, 0x52, 0x65, 0x67, 0x10, 0x01, 0x12,
	0x09, 0x0a, 0x05, 0x57, 0x61, 0x74, 0x63, 0x68, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x75,
	0x73, 0x68, 0x10, 0x03, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_proto_rawDescOnce sync.Once
	file_msg_proto_rawDescData = file_msg_proto_rawDesc
)

func file_msg_proto_rawDescGZIP() []byte {
	file_msg_proto_rawDescOnce.Do(func() {
		file_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_proto_rawDescData)
	})
	return file_msg_proto_rawDescData
}

var file_msg_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_msg_proto_goTypes = []interface{}{
	(MsgType)(0),     // 0: discovery.MsgType
	(*Msg)(nil),      // 1: discovery.msg
	(*RegInfo)(nil),  // 2: discovery.RegInfo
	(*RegMsg)(nil),   // 3: discovery.RegMsg
	(*UnregMsg)(nil), // 4: discovery.UnregMsg
	(*WatchMsg)(nil), // 5: discovery.WatchMsg
	(*PushMsg)(nil),  // 6: discovery.PushMsg
	nil,              // 7: discovery.PushMsg.AppsEntry
}
var file_msg_proto_depIdxs = []int32{
	0, // 0: discovery.msg.msg_type:type_name -> discovery.MsgType
	3, // 1: discovery.msg.reg_msg:type_name -> discovery.RegMsg
	4, // 2: discovery.msg.unreg_msg:type_name -> discovery.UnregMsg
	5, // 3: discovery.msg.watch_msg:type_name -> discovery.WatchMsg
	6, // 4: discovery.msg.push_msg:type_name -> discovery.PushMsg
	2, // 5: discovery.RegMsg.reg_info:type_name -> discovery.RegInfo
	7, // 6: discovery.PushMsg.apps:type_name -> discovery.PushMsg.AppsEntry
	2, // 7: discovery.PushMsg.AppsEntry.value:type_name -> discovery.RegInfo
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_msg_proto_init() }
func file_msg_proto_init() {
	if File_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegInfo); i {
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
		file_msg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegMsg); i {
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
		file_msg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnregMsg); i {
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
		file_msg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchMsg); i {
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
		file_msg_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMsg); i {
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
	file_msg_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Msg_RegMsg)(nil),
		(*Msg_UnregMsg)(nil),
		(*Msg_WatchMsg)(nil),
		(*Msg_PushMsg)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msg_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msg_proto_goTypes,
		DependencyIndexes: file_msg_proto_depIdxs,
		EnumInfos:         file_msg_proto_enumTypes,
		MessageInfos:      file_msg_proto_msgTypes,
	}.Build()
	File_msg_proto = out.File
	file_msg_proto_rawDesc = nil
	file_msg_proto_goTypes = nil
	file_msg_proto_depIdxs = nil
}
