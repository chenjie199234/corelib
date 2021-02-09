// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: first/first.proto

package first

import (
	second "./second"
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

type Helloreq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	I32   int32                      `protobuf:"varint,1,opt,name=i32,proto3" json:"i32,omitempty"`            //header:true|gt:6|notin:["8","9"]
	Ri32  []int32                    `protobuf:"varint,2,rep,packed,name=ri32,proto3" json:"ri32,omitempty"`   //empty:false|in:["1","2","3"]
	Ui32  uint32                     `protobuf:"varint,3,opt,name=ui32,proto3" json:"ui32,omitempty"`          //header:true|gt:6|notin:["8","9"]
	Rui32 []uint32                   `protobuf:"varint,4,rep,packed,name=rui32,proto3" json:"rui32,omitempty"` //empty:false|in:["1","2","3"]
	I64   int64                      `protobuf:"varint,5,opt,name=i64,proto3" json:"i64,omitempty"`
	Ri64  []int64                    `protobuf:"varint,6,rep,packed,name=ri64,proto3" json:"ri64,omitempty"`
	Ui64  uint64                     `protobuf:"varint,7,opt,name=ui64,proto3" json:"ui64,omitempty"`
	Rui64 []uint64                   `protobuf:"varint,8,rep,packed,name=rui64,proto3" json:"rui64,omitempty"`
	Bs    []byte                     `protobuf:"bytes,9,opt,name=bs,proto3" json:"bs,omitempty"`    //empty:false
	Rbs   [][]byte                   `protobuf:"bytes,10,rep,name=rbs,proto3" json:"rbs,omitempty"` //empty:false
	Ss    string                     `protobuf:"bytes,11,opt,name=ss,proto3" json:"ss,omitempty"`   //empty:false|header:true
	Rss   []string                   `protobuf:"bytes,12,rep,name=rss,proto3" json:"rss,omitempty"` //empty:false|header:true
	F     float32                    `protobuf:"fixed32,13,opt,name=f,proto3" json:"f,omitempty"`
	Rf    []float32                  `protobuf:"fixed32,14,rep,packed,name=rf,proto3" json:"rf,omitempty"`
	E     second.ABC                 `protobuf:"varint,15,opt,name=e,proto3,enum=second.ABC" json:"e,omitempty"`
	Re    []second.ABC               `protobuf:"varint,16,rep,packed,name=re,proto3,enum=second.ABC" json:"re,omitempty"`
	A     *second.Authinfo           `protobuf:"bytes,17,opt,name=a,proto3" json:"a,omitempty"`
	Ra    []*second.Authinfo         `protobuf:"bytes,18,rep,name=ra,proto3" json:"ra,omitempty"`
	Me    map[int32]second.ABC       `protobuf:"bytes,19,rep,name=me,proto3" json:"me,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=second.ABC"` //empty:false
	Mb    map[int64]*second.Baseinfo `protobuf:"bytes,20,rep,name=mb,proto3" json:"mb,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`                  //empty:false
}

func (x *Helloreq) Reset() {
	*x = Helloreq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_first_first_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Helloreq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Helloreq) ProtoMessage() {}

func (x *Helloreq) ProtoReflect() protoreflect.Message {
	mi := &file_first_first_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Helloreq.ProtoReflect.Descriptor instead.
func (*Helloreq) Descriptor() ([]byte, []int) {
	return file_first_first_proto_rawDescGZIP(), []int{0}
}

func (x *Helloreq) GetI32() int32 {
	if x != nil {
		return x.I32
	}
	return 0
}

func (x *Helloreq) GetRi32() []int32 {
	if x != nil {
		return x.Ri32
	}
	return nil
}

func (x *Helloreq) GetUi32() uint32 {
	if x != nil {
		return x.Ui32
	}
	return 0
}

func (x *Helloreq) GetRui32() []uint32 {
	if x != nil {
		return x.Rui32
	}
	return nil
}

func (x *Helloreq) GetI64() int64 {
	if x != nil {
		return x.I64
	}
	return 0
}

func (x *Helloreq) GetRi64() []int64 {
	if x != nil {
		return x.Ri64
	}
	return nil
}

func (x *Helloreq) GetUi64() uint64 {
	if x != nil {
		return x.Ui64
	}
	return 0
}

func (x *Helloreq) GetRui64() []uint64 {
	if x != nil {
		return x.Rui64
	}
	return nil
}

func (x *Helloreq) GetBs() []byte {
	if x != nil {
		return x.Bs
	}
	return nil
}

func (x *Helloreq) GetRbs() [][]byte {
	if x != nil {
		return x.Rbs
	}
	return nil
}

func (x *Helloreq) GetSs() string {
	if x != nil {
		return x.Ss
	}
	return ""
}

func (x *Helloreq) GetRss() []string {
	if x != nil {
		return x.Rss
	}
	return nil
}

func (x *Helloreq) GetF() float32 {
	if x != nil {
		return x.F
	}
	return 0
}

func (x *Helloreq) GetRf() []float32 {
	if x != nil {
		return x.Rf
	}
	return nil
}

func (x *Helloreq) GetE() second.ABC {
	if x != nil {
		return x.E
	}
	return second.ABC_A
}

func (x *Helloreq) GetRe() []second.ABC {
	if x != nil {
		return x.Re
	}
	return nil
}

func (x *Helloreq) GetA() *second.Authinfo {
	if x != nil {
		return x.A
	}
	return nil
}

func (x *Helloreq) GetRa() []*second.Authinfo {
	if x != nil {
		return x.Ra
	}
	return nil
}

func (x *Helloreq) GetMe() map[int32]second.ABC {
	if x != nil {
		return x.Me
	}
	return nil
}

func (x *Helloreq) GetMb() map[int64]*second.Baseinfo {
	if x != nil {
		return x.Mb
	}
	return nil
}

type Helloresp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Helloresp) Reset() {
	*x = Helloresp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_first_first_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Helloresp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Helloresp) ProtoMessage() {}

func (x *Helloresp) ProtoReflect() protoreflect.Message {
	mi := &file_first_first_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Helloresp.ProtoReflect.Descriptor instead.
func (*Helloresp) Descriptor() ([]byte, []int) {
	return file_first_first_proto_rawDescGZIP(), []int{1}
}

type Worldreq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Worldreq) Reset() {
	*x = Worldreq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_first_first_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Worldreq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Worldreq) ProtoMessage() {}

func (x *Worldreq) ProtoReflect() protoreflect.Message {
	mi := &file_first_first_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Worldreq.ProtoReflect.Descriptor instead.
func (*Worldreq) Descriptor() ([]byte, []int) {
	return file_first_first_proto_rawDescGZIP(), []int{2}
}

type Worldresp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Worldresp) Reset() {
	*x = Worldresp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_first_first_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Worldresp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Worldresp) ProtoMessage() {}

func (x *Worldresp) ProtoReflect() protoreflect.Message {
	mi := &file_first_first_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Worldresp.ProtoReflect.Descriptor instead.
func (*Worldresp) Descriptor() ([]byte, []int) {
	return file_first_first_proto_rawDescGZIP(), []int{3}
}

var File_first_first_proto protoreflect.FileDescriptor

var file_first_first_proto_rawDesc = []byte{
	0x0a, 0x11, 0x66, 0x69, 0x72, 0x73, 0x74, 0x2f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x1a, 0x13, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x2f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe5, 0x04, 0x0a, 0x08, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x72, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x33, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x69, 0x33, 0x32, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x69, 0x33, 0x32, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x04, 0x72, 0x69,
	0x33, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x69, 0x33, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x75, 0x69, 0x33, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x75, 0x69, 0x33, 0x32, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x05, 0x72, 0x75, 0x69, 0x33, 0x32, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x36, 0x34, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x69, 0x36, 0x34, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x69, 0x36, 0x34, 0x18, 0x06, 0x20, 0x03, 0x28, 0x03, 0x52, 0x04, 0x72, 0x69,
	0x36, 0x34, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x69, 0x36, 0x34, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x04, 0x75, 0x69, 0x36, 0x34, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x75, 0x69, 0x36, 0x34, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x04, 0x52, 0x05, 0x72, 0x75, 0x69, 0x36, 0x34, 0x12, 0x0e, 0x0a, 0x02,
	0x62, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x62, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x72, 0x62, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x03, 0x72, 0x62, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x73, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x72, 0x73, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x72, 0x73, 0x73,
	0x12, 0x0c, 0x0a, 0x01, 0x66, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x66, 0x12, 0x0e,
	0x0a, 0x02, 0x72, 0x66, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x02, 0x52, 0x02, 0x72, 0x66, 0x12, 0x19,
	0x0a, 0x01, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x2e, 0x41, 0x42, 0x43, 0x52, 0x01, 0x65, 0x12, 0x1b, 0x0a, 0x02, 0x72, 0x65, 0x18,
	0x10, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x2e, 0x41,
	0x42, 0x43, 0x52, 0x02, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x01, 0x61, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x69,
	0x6e, 0x66, 0x6f, 0x52, 0x01, 0x61, 0x12, 0x20, 0x0a, 0x02, 0x72, 0x61, 0x18, 0x12, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x69, 0x6e, 0x66, 0x6f, 0x52, 0x02, 0x72, 0x61, 0x12, 0x27, 0x0a, 0x02, 0x6d, 0x65, 0x18, 0x13,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x72, 0x65, 0x71, 0x2e, 0x4d, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x02, 0x6d,
	0x65, 0x12, 0x27, 0x0a, 0x02, 0x6d, 0x62, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x72, 0x65, 0x71, 0x2e, 0x4d,
	0x62, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x02, 0x6d, 0x62, 0x1a, 0x42, 0x0a, 0x07, 0x4d, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x2e,
	0x41, 0x42, 0x43, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x47,
	0x0a, 0x07, 0x4d, 0x62, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x0b, 0x0a, 0x09, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x72, 0x65, 0x73, 0x70, 0x22, 0x0a, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x72, 0x65, 0x71,
	0x22, 0x0b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x72, 0x65, 0x73, 0x70, 0x32, 0x5f, 0x0a,
	0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x12,
	0x0f, 0x2e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x72, 0x65, 0x71,
	0x1a, 0x10, 0x2e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x72, 0x65,
	0x73, 0x70, 0x12, 0x2a, 0x0a, 0x05, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x0f, 0x2e, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x72, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x72, 0x65, 0x73, 0x70, 0x42, 0x0f,
	0x5a, 0x0d, 0x2e, 0x2f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x3b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_first_first_proto_rawDescOnce sync.Once
	file_first_first_proto_rawDescData = file_first_first_proto_rawDesc
)

func file_first_first_proto_rawDescGZIP() []byte {
	file_first_first_proto_rawDescOnce.Do(func() {
		file_first_first_proto_rawDescData = protoimpl.X.CompressGZIP(file_first_first_proto_rawDescData)
	})
	return file_first_first_proto_rawDescData
}

var file_first_first_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_first_first_proto_goTypes = []interface{}{
	(*Helloreq)(nil),        // 0: first.helloreq
	(*Helloresp)(nil),       // 1: first.helloresp
	(*Worldreq)(nil),        // 2: first.worldreq
	(*Worldresp)(nil),       // 3: first.worldresp
	nil,                     // 4: first.helloreq.MeEntry
	nil,                     // 5: first.helloreq.MbEntry
	(second.ABC)(0),         // 6: second.ABC
	(*second.Authinfo)(nil), // 7: second.authinfo
	(*second.Baseinfo)(nil), // 8: second.baseinfo
}
var file_first_first_proto_depIdxs = []int32{
	6,  // 0: first.helloreq.e:type_name -> second.ABC
	6,  // 1: first.helloreq.re:type_name -> second.ABC
	7,  // 2: first.helloreq.a:type_name -> second.authinfo
	7,  // 3: first.helloreq.ra:type_name -> second.authinfo
	4,  // 4: first.helloreq.me:type_name -> first.helloreq.MeEntry
	5,  // 5: first.helloreq.mb:type_name -> first.helloreq.MbEntry
	6,  // 6: first.helloreq.MeEntry.value:type_name -> second.ABC
	8,  // 7: first.helloreq.MbEntry.value:type_name -> second.baseinfo
	0,  // 8: first.first.hello:input_type -> first.helloreq
	2,  // 9: first.first.world:input_type -> first.worldreq
	1,  // 10: first.first.hello:output_type -> first.helloresp
	3,  // 11: first.first.world:output_type -> first.worldresp
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_first_first_proto_init() }
func file_first_first_proto_init() {
	if File_first_first_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_first_first_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Helloreq); i {
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
		file_first_first_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Helloresp); i {
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
		file_first_first_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Worldreq); i {
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
		file_first_first_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Worldresp); i {
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
			RawDescriptor: file_first_first_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_first_first_proto_goTypes,
		DependencyIndexes: file_first_first_proto_depIdxs,
		MessageInfos:      file_first_first_proto_msgTypes,
	}.Build()
	File_first_first_proto = out.File
	file_first_first_proto_rawDesc = nil
	file_first_first_proto_goTypes = nil
	file_first_first_proto_depIdxs = nil
}
