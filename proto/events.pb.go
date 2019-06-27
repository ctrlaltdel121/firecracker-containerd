// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: events.proto

package proto

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type VMStart struct {
	VMID                 string   `protobuf:"bytes,1,opt,name=VMID,json=vMID,proto3" json:"VMID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VMStart) Reset()         { *m = VMStart{} }
func (m *VMStart) String() string { return proto.CompactTextString(m) }
func (*VMStart) ProtoMessage()    {}
func (*VMStart) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{0}
}
func (m *VMStart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VMStart.Unmarshal(m, b)
}
func (m *VMStart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VMStart.Marshal(b, m, deterministic)
}
func (m *VMStart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VMStart.Merge(m, src)
}
func (m *VMStart) XXX_Size() int {
	return xxx_messageInfo_VMStart.Size(m)
}
func (m *VMStart) XXX_DiscardUnknown() {
	xxx_messageInfo_VMStart.DiscardUnknown(m)
}

var xxx_messageInfo_VMStart proto.InternalMessageInfo

func (m *VMStart) GetVMID() string {
	if m != nil {
		return m.VMID
	}
	return ""
}

type VMStop struct {
	VMID                 string   `protobuf:"bytes,1,opt,name=VMID,json=vMID,proto3" json:"VMID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VMStop) Reset()         { *m = VMStop{} }
func (m *VMStop) String() string { return proto.CompactTextString(m) }
func (*VMStop) ProtoMessage()    {}
func (*VMStop) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{1}
}
func (m *VMStop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VMStop.Unmarshal(m, b)
}
func (m *VMStop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VMStop.Marshal(b, m, deterministic)
}
func (m *VMStop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VMStop.Merge(m, src)
}
func (m *VMStop) XXX_Size() int {
	return xxx_messageInfo_VMStop.Size(m)
}
func (m *VMStop) XXX_DiscardUnknown() {
	xxx_messageInfo_VMStop.DiscardUnknown(m)
}

var xxx_messageInfo_VMStop proto.InternalMessageInfo

func (m *VMStop) GetVMID() string {
	if m != nil {
		return m.VMID
	}
	return ""
}

func init() {
	proto.RegisterType((*VMStart)(nil), "VMStart")
	proto.RegisterType((*VMStop)(nil), "VMStop")
}

func init() { proto.RegisterFile("events.proto", fileDescriptor_8f22242cb04491f9) }

var fileDescriptor_8f22242cb04491f9 = []byte{
	// 84 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x2d, 0x4b, 0xcd,
	0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe5, 0x62, 0x0f, 0xf3, 0x0d, 0x2e,
	0x49, 0x2c, 0x2a, 0x11, 0x12, 0xe2, 0x62, 0x09, 0xf3, 0xf5, 0x74, 0x91, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x62, 0x29, 0xf3, 0xf5, 0x74, 0x51, 0x92, 0xe1, 0x62, 0x03, 0x49, 0xe7, 0x17, 0x60,
	0x93, 0x4d, 0x62, 0x03, 0x9b, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x8a, 0x85, 0x02,
	0x53, 0x00, 0x00, 0x00,
}
