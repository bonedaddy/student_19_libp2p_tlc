// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pubsub.proto

package transport

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type MsgType int32

const (
	MsgType_Raw MsgType = 0
)

var MsgType_name = map[int32]string{
	0: "Raw",
}

var MsgType_value = map[string]int32{
	"Raw": 0,
}

func (x MsgType) Enum() *MsgType {
	p := new(MsgType)
	*p = x
	return p
}

func (x MsgType) String() string {
	return proto.EnumName(MsgType_name, int32(x))
}

func (x *MsgType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MsgType_value, data, "MsgType")
	if err != nil {
		return err
	}
	*x = MsgType(value)
	return nil
}

func (MsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_91df006b05e20cf7, []int{0}
}

type PbMessage struct {
	MsgType              *MsgType     `protobuf:"varint,1,req,name=Msg_type,json=MsgType,enum=transport.MsgType" json:"Msg_type,omitempty"`
	Source               *int64       `protobuf:"varint,2,req,name=source" json:"source,omitempty"`
	Step                 *int64       `protobuf:"varint,3,req,name=step" json:"step,omitempty"`
	History              []*PbMessage `protobuf:"bytes,4,rep,name=history" json:"history,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PbMessage) Reset()         { *m = PbMessage{} }
func (m *PbMessage) String() string { return proto.CompactTextString(m) }
func (*PbMessage) ProtoMessage()    {}
func (*PbMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_91df006b05e20cf7, []int{0}
}

func (m *PbMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbMessage.Unmarshal(m, b)
}
func (m *PbMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbMessage.Marshal(b, m, deterministic)
}
func (m *PbMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbMessage.Merge(m, src)
}
func (m *PbMessage) XXX_Size() int {
	return xxx_messageInfo_PbMessage.Size(m)
}
func (m *PbMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PbMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PbMessage proto.InternalMessageInfo

func (m *PbMessage) GetMsgType() MsgType {
	if m != nil && m.MsgType != nil {
		return *m.MsgType
	}
	return MsgType_Raw
}

func (m *PbMessage) GetSource() int64 {
	if m != nil && m.Source != nil {
		return *m.Source
	}
	return 0
}

func (m *PbMessage) GetStep() int64 {
	if m != nil && m.Step != nil {
		return *m.Step
	}
	return 0
}

func (m *PbMessage) GetHistory() []*PbMessage {
	if m != nil {
		return m.History
	}
	return nil
}

func init() {
	proto.RegisterEnum("transport.MsgType", MsgType_name, MsgType_value)
	proto.RegisterType((*PbMessage)(nil), "transport.PbMessage")
}

func init() { proto.RegisterFile("pubsub.proto", fileDescriptor_91df006b05e20cf7) }

var fileDescriptor_91df006b05e20cf7 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xc1, 0x0a, 0x82, 0x40,
	0x10, 0x86, 0xd3, 0x95, 0xcc, 0x29, 0x22, 0xc6, 0x88, 0x3d, 0x4a, 0x27, 0xe9, 0xb0, 0x07, 0xdf,
	0x43, 0x88, 0xa5, 0x7b, 0x68, 0x2c, 0xd6, 0xa5, 0x1d, 0x76, 0x46, 0xc2, 0x27, 0xe9, 0x75, 0x03,
	0x49, 0xeb, 0x36, 0xf3, 0xff, 0x1f, 0x3f, 0x1f, 0x6c, 0xa8, 0x6f, 0xb9, 0x6f, 0x0d, 0x05, 0x2f,
	0x1e, 0x33, 0x09, 0xcd, 0x93, 0xc9, 0x07, 0x39, 0xbe, 0x23, 0xc8, 0xce, 0x6d, 0xed, 0x98, 0x9b,
	0xce, 0xa1, 0x81, 0x55, 0xcd, 0xdd, 0x55, 0x06, 0x72, 0x3a, 0x2a, 0xe2, 0x72, 0x5b, 0xe5, 0x66,
	0x66, 0xcd, 0x54, 0xd9, 0xb4, 0xe6, 0xee, 0x32, 0x90, 0xc3, 0x03, 0x2c, 0xd9, 0xf7, 0xe1, 0xe6,
	0x74, 0x5c, 0xc4, 0xa5, 0xb2, 0xdf, 0x0f, 0x11, 0x12, 0x16, 0x47, 0x5a, 0x8d, 0xe9, 0x78, 0xa3,
	0x81, 0xf4, 0xfe, 0x60, 0xf1, 0x61, 0xd0, 0x49, 0xa1, 0xca, 0x75, 0xb5, 0xff, 0x9b, 0x9e, 0x15,
	0xec, 0x04, 0x9d, 0xf2, 0x9f, 0x0b, 0xa6, 0xa0, 0x6c, 0xf3, 0xda, 0x2d, 0x3e, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x6c, 0xc9, 0x00, 0x62, 0xc8, 0x00, 0x00, 0x00,
}
