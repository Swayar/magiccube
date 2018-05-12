// Code generated by protoc-gen-go. DO NOT EDIT.
// source: basic-transaction.proto

/*
Package transaction is a generated protocol buffer package.

It is generated from these files:
	basic-transaction.proto
	transaction.proto

It has these top-level messages:
	BasicTest
	Test
*/
package transaction

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BasicTest struct {
	Version     uint64 `protobuf:"varint,1,opt,name=version" json:"version"`
	CursorNum   uint64 `protobuf:"varint,2,opt,name=cursor_num,json=cursorNum" json:"cursor_num"`
	CursorLabel uint64 `protobuf:"varint,3,opt,name=cursor_label,json=cursorLabel" json:"cursor_label"`
	Lifetime    uint64 `protobuf:"varint,4,opt,name=lifetime" json:"lifetime"`
	Sender      string `protobuf:"bytes,5,opt,name=sender" json:"sender"`
	Contract    string `protobuf:"bytes,6,opt,name=contract" json:"contract"`
	Method      string `protobuf:"bytes,7,opt,name=method" json:"method"`
	Param       []byte `protobuf:"bytes,8,opt,name=param,proto3" json:"param"`
	SigAlg      uint64 `protobuf:"varint,9,opt,name=sig_alg,json=sigAlg" json:"sig_alg"`
}

func (m *BasicTest) Reset()                    { *m = BasicTest{} }
func (m *BasicTest) String() string            { return proto.CompactTextString(m) }
func (*BasicTest) ProtoMessage()               {}
func (*BasicTest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BasicTest) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *BasicTest) GetCursorNum() uint64 {
	if m != nil {
		return m.CursorNum
	}
	return 0
}

func (m *BasicTest) GetCursorLabel() uint64 {
	if m != nil {
		return m.CursorLabel
	}
	return 0
}

func (m *BasicTest) GetLifetime() uint64 {
	if m != nil {
		return m.Lifetime
	}
	return 0
}

func (m *BasicTest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *BasicTest) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *BasicTest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *BasicTest) GetParam() []byte {
	if m != nil {
		return m.Param
	}
	return nil
}

func (m *BasicTest) GetSigAlg() uint64 {
	if m != nil {
		return m.SigAlg
	}
	return 0
}

func init() {
	proto.RegisterType((*BasicTest)(nil), "transaction.BasicTest")
}

func init() { proto.RegisterFile("basic-transaction.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0x47, 0x89, 0xb6, 0xbb, 0xdd, 0x69, 0x4f, 0x41, 0xec, 0x20, 0x08, 0xab, 0xa7, 0xbd, 0xe8,
	0xc5, 0x4f, 0xa0, 0x67, 0xf1, 0xb0, 0x78, 0x2f, 0xd9, 0x34, 0xae, 0x81, 0xfc, 0x29, 0x99, 0x59,
	0x3f, 0xbe, 0xc8, 0x26, 0x75, 0xe9, 0xf1, 0xbd, 0x37, 0x3f, 0x02, 0x81, 0xfd, 0xa0, 0xc8, 0xea,
	0x27, 0x4e, 0x2a, 0x90, 0xd2, 0x6c, 0x63, 0x78, 0x3e, 0xa5, 0xc8, 0x51, 0x6e, 0x2f, 0xd4, 0xe3,
	0xaf, 0x80, 0xe6, 0x6d, 0x3e, 0xfc, 0x34, 0xc4, 0x12, 0xa1, 0xfe, 0x31, 0x89, 0x6c, 0x0c, 0x28,
	0x5a, 0xd1, 0xad, 0xfa, 0x7f, 0x94, 0xf7, 0x00, 0x7a, 0x4a, 0x14, 0xd3, 0x21, 0x4c, 0x1e, 0xaf,
	0x72, 0x6c, 0x8a, 0xf9, 0x98, 0xbc, 0x7c, 0x80, 0xdd, 0x39, 0x3b, 0x35, 0x18, 0x87, 0xd7, 0xf9,
	0x60, 0x5b, 0xdc, 0xfb, 0xac, 0xe4, 0x1d, 0x6c, 0x9c, 0xfd, 0x32, 0x6c, 0xbd, 0xc1, 0x55, 0xce,
	0x0b, 0xcb, 0x5b, 0xa8, 0xc8, 0x84, 0xa3, 0x49, 0xb8, 0x6e, 0x45, 0xd7, 0xf4, 0x67, 0x9a, 0x37,
	0x3a, 0x06, 0x4e, 0x4a, 0x33, 0x56, 0xb9, 0x2c, 0x3c, 0x6f, 0xbc, 0xe1, 0xef, 0x78, 0xc4, 0xba,
	0x6c, 0x0a, 0xc9, 0x1b, 0x58, 0x9f, 0x54, 0x52, 0x1e, 0x37, 0xad, 0xe8, 0x76, 0x7d, 0x01, 0xb9,
	0x87, 0x9a, 0xec, 0x78, 0x50, 0x6e, 0xc4, 0x26, 0x3f, 0x5e, 0x91, 0x1d, 0x5f, 0xdd, 0x38, 0x54,
	0xf9, 0x53, 0x5e, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x85, 0x3f, 0xb6, 0x68, 0x2f, 0x01, 0x00,
	0x00,
}
