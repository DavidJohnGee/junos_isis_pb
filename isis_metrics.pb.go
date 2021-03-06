// Code generated by protoc-gen-go. DO NOT EDIT.
// source: isis_metrics.proto

/*
Package isis_metrics is a generated protocol buffer package.

It is generated from these files:
	isis_metrics.proto

It has these top-level messages:
	IFACE
	UPDATE
*/
package isis_metrics

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

type IFACE struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Metric1          *int32  `protobuf:"varint,2,req,name=metric1" json:"metric1,omitempty"`
	Metric2          *int32  `protobuf:"varint,3,req,name=metric2" json:"metric2,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IFACE) Reset()                    { *m = IFACE{} }
func (m *IFACE) String() string            { return proto.CompactTextString(m) }
func (*IFACE) ProtoMessage()               {}
func (*IFACE) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IFACE) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *IFACE) GetMetric1() int32 {
	if m != nil && m.Metric1 != nil {
		return *m.Metric1
	}
	return 0
}

func (m *IFACE) GetMetric2() int32 {
	if m != nil && m.Metric2 != nil {
		return *m.Metric2
	}
	return 0
}

// Our address book file is just one of these.
type UPDATE struct {
	Name             *string  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Ifaces           []*IFACE `protobuf:"bytes,2,rep,name=ifaces" json:"ifaces,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *UPDATE) Reset()                    { *m = UPDATE{} }
func (m *UPDATE) String() string            { return proto.CompactTextString(m) }
func (*UPDATE) ProtoMessage()               {}
func (*UPDATE) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UPDATE) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *UPDATE) GetIfaces() []*IFACE {
	if m != nil {
		return m.Ifaces
	}
	return nil
}

func init() {
	proto.RegisterType((*IFACE)(nil), "IFACE")
	proto.RegisterType((*UPDATE)(nil), "UPDATE")
}

func init() { proto.RegisterFile("isis_metrics.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0x2c, 0xce, 0x2c,
	0x8e, 0xcf, 0x4d, 0x2d, 0x29, 0xca, 0x4c, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xf2,
	0xe7, 0x62, 0xf5, 0x74, 0x73, 0x74, 0x76, 0x15, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95,
	0x60, 0x54, 0x60, 0xd2, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x24, 0xb8, 0xd8, 0x21, 0xaa, 0x0d, 0x25,
	0x98, 0x14, 0x98, 0x34, 0x58, 0x83, 0x60, 0x5c, 0x84, 0x8c, 0x91, 0x04, 0x33, 0xb2, 0x8c, 0x91,
	0x92, 0x0d, 0x17, 0x5b, 0x68, 0x80, 0x8b, 0x63, 0x08, 0x76, 0x13, 0xe5, 0xb8, 0xd8, 0x32, 0xd3,
	0x12, 0x93, 0x53, 0x8b, 0x25, 0x98, 0x14, 0x98, 0x35, 0xb8, 0x8d, 0xd8, 0xf4, 0xc0, 0xb6, 0x07,
	0x41, 0x45, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x56, 0x86, 0x52, 0x71, 0xa3, 0x00, 0x00, 0x00,
}
