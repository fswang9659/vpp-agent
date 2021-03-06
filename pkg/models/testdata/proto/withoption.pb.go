// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/withoption.proto

package model

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "go.ligato.io/vpp-agent/v3/proto/ligato/generic"
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

type WithOption struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Caption              string   `protobuf:"bytes,2,opt,name=caption,proto3" json:"caption,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WithOption) Reset()         { *m = WithOption{} }
func (m *WithOption) String() string { return proto.CompactTextString(m) }
func (*WithOption) ProtoMessage()    {}
func (*WithOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f7848ce778760ed, []int{0}
}

func (m *WithOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WithOption.Unmarshal(m, b)
}
func (m *WithOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WithOption.Marshal(b, m, deterministic)
}
func (m *WithOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithOption.Merge(m, src)
}
func (m *WithOption) XXX_Size() int {
	return xxx_messageInfo_WithOption.Size(m)
}
func (m *WithOption) XXX_DiscardUnknown() {
	xxx_messageInfo_WithOption.DiscardUnknown(m)
}

var xxx_messageInfo_WithOption proto.InternalMessageInfo

func (m *WithOption) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *WithOption) GetCaption() string {
	if m != nil {
		return m.Caption
	}
	return ""
}

func init() {
	proto.RegisterType((*WithOption)(nil), "model.WithOption")
}

func init() { proto.RegisterFile("proto/withoption.proto", fileDescriptor_2f7848ce778760ed) }

var fileDescriptor_2f7848ce778760ed = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xcf, 0x2c, 0xc9, 0xc8, 0x2f, 0x28, 0xc9, 0xcc, 0xcf, 0xd3, 0x03, 0x0b, 0x08,
	0xb1, 0xe6, 0xe6, 0xa7, 0xa4, 0xe6, 0x48, 0xc9, 0xe4, 0x64, 0xa6, 0x27, 0x96, 0xe4, 0xeb, 0xa7,
	0xa7, 0xe6, 0xa5, 0x16, 0x65, 0x26, 0xeb, 0x43, 0xd4, 0x14, 0x43, 0x14, 0x29, 0x05, 0x72, 0x71,
	0x85, 0x67, 0x96, 0x64, 0xf8, 0x83, 0x05, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x58, 0x83, 0x98, 0x32, 0x53, 0x84, 0x24, 0xb8, 0xd8, 0x93, 0x13, 0xc1, 0x52, 0x12,
	0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x95, 0xe4, 0xa7, 0x43, 0x12, 0xa2, 0x5c, 0x5c,
	0x60, 0x2b, 0x4a, 0x52, 0x8b, 0x4b, 0x8a, 0xa5, 0xd8, 0xcb, 0x21, 0x26, 0x27, 0xb1, 0x81, 0x4d,
	0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x59, 0x7c, 0x47, 0xa0, 0x98, 0x00, 0x00, 0x00,
}
