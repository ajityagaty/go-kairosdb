// Code generated by protoc-gen-go.
// source: types.proto
// DO NOT EDIT!

package kairosdb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Because protobuf doesn't support Map<string, repeated string>
type StringList struct {
	Values []string `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *StringList) Reset()                    { *m = StringList{} }
func (m *StringList) String() string            { return proto.CompactTextString(m) }
func (*StringList) ProtoMessage()               {}
func (*StringList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func init() {
	proto.RegisterType((*StringList)(nil), "kairosdb.StringList")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 90 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0x4e, 0xcc, 0x2c, 0xca, 0x2f, 0x4e,
	0x49, 0x52, 0x52, 0xe1, 0xe2, 0x0a, 0x2e, 0x29, 0xca, 0xcc, 0x4b, 0xf7, 0xc9, 0x2c, 0x2e, 0x11,
	0x12, 0xe3, 0x62, 0x2b, 0x4b, 0xcc, 0x29, 0x4d, 0x2d, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x0c,
	0x82, 0xf2, 0x9c, 0xb8, 0xa2, 0xe0, 0x3a, 0x92, 0xd8, 0xc0, 0x46, 0x18, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xb0, 0xec, 0xcc, 0x5a, 0x51, 0x00, 0x00, 0x00,
}
