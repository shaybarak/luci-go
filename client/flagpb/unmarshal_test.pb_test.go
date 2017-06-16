// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/luci/luci-go/client/flagpb/unmarshal_test.proto

/*
Package flagpb is a generated protocol buffer package.

It is generated from these files:
	github.com/luci/luci-go/client/flagpb/unmarshal_test.proto

It has these top-level messages:
	M1
	M2
	M3
	MapContainer
*/
package flagpb

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

type E int32

const (
	E_V0 E = 0
	E_V1 E = 1
)

var E_name = map[int32]string{
	0: "V0",
	1: "V1",
}
var E_value = map[string]int32{
	"V0": 0,
	"V1": 1,
}

func (x E) String() string {
	return proto.EnumName(E_name, int32(x))
}
func (E) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type M1 struct {
	S  string  `protobuf:"bytes,1,opt,name=s" json:"s,omitempty"`
	I  int32   `protobuf:"varint,2,opt,name=i" json:"i,omitempty"`
	Ri []int32 `protobuf:"varint,3,rep,packed,name=ri" json:"ri,omitempty"`
	B  bool    `protobuf:"varint,4,opt,name=b" json:"b,omitempty"`
	Rb []bool  `protobuf:"varint,6,rep,packed,name=rb" json:"rb,omitempty"`
	Bb []byte  `protobuf:"bytes,5,opt,name=bb,proto3" json:"bb,omitempty"`
}

func (m *M1) Reset()                    { *m = M1{} }
func (m *M1) String() string            { return proto.CompactTextString(m) }
func (*M1) ProtoMessage()               {}
func (*M1) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *M1) GetS() string {
	if m != nil {
		return m.S
	}
	return ""
}

func (m *M1) GetI() int32 {
	if m != nil {
		return m.I
	}
	return 0
}

func (m *M1) GetRi() []int32 {
	if m != nil {
		return m.Ri
	}
	return nil
}

func (m *M1) GetB() bool {
	if m != nil {
		return m.B
	}
	return false
}

func (m *M1) GetRb() []bool {
	if m != nil {
		return m.Rb
	}
	return nil
}

func (m *M1) GetBb() []byte {
	if m != nil {
		return m.Bb
	}
	return nil
}

type M2 struct {
	M1 *M1 `protobuf:"bytes,1,opt,name=m1" json:"m1,omitempty"`
	E  E   `protobuf:"varint,2,opt,name=e,enum=flagpb.E" json:"e,omitempty"`
}

func (m *M2) Reset()                    { *m = M2{} }
func (m *M2) String() string            { return proto.CompactTextString(m) }
func (*M2) ProtoMessage()               {}
func (*M2) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *M2) GetM1() *M1 {
	if m != nil {
		return m.M1
	}
	return nil
}

func (m *M2) GetE() E {
	if m != nil {
		return m.E
	}
	return E_V0
}

type M3 struct {
	M1 []*M1  `protobuf:"bytes,1,rep,name=m1" json:"m1,omitempty"`
	M2 *M2    `protobuf:"bytes,2,opt,name=m2" json:"m2,omitempty"`
	B  bool   `protobuf:"varint,3,opt,name=b" json:"b,omitempty"`
	S  string `protobuf:"bytes,4,opt,name=s" json:"s,omitempty"`
	Bt []byte `protobuf:"bytes,5,opt,name=bt,proto3" json:"bt,omitempty"`
}

func (m *M3) Reset()                    { *m = M3{} }
func (m *M3) String() string            { return proto.CompactTextString(m) }
func (*M3) ProtoMessage()               {}
func (*M3) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *M3) GetM1() []*M1 {
	if m != nil {
		return m.M1
	}
	return nil
}

func (m *M3) GetM2() *M2 {
	if m != nil {
		return m.M2
	}
	return nil
}

func (m *M3) GetB() bool {
	if m != nil {
		return m.B
	}
	return false
}

func (m *M3) GetS() string {
	if m != nil {
		return m.S
	}
	return ""
}

func (m *M3) GetBt() []byte {
	if m != nil {
		return m.Bt
	}
	return nil
}

type MapContainer struct {
	Ss  map[string]string `protobuf:"bytes,1,rep,name=ss" json:"ss,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Ii  map[int32]int32   `protobuf:"bytes,2,rep,name=ii" json:"ii,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	Sm1 map[string]*M1    `protobuf:"bytes,3,rep,name=sm1" json:"sm1,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *MapContainer) Reset()                    { *m = MapContainer{} }
func (m *MapContainer) String() string            { return proto.CompactTextString(m) }
func (*MapContainer) ProtoMessage()               {}
func (*MapContainer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MapContainer) GetSs() map[string]string {
	if m != nil {
		return m.Ss
	}
	return nil
}

func (m *MapContainer) GetIi() map[int32]int32 {
	if m != nil {
		return m.Ii
	}
	return nil
}

func (m *MapContainer) GetSm1() map[string]*M1 {
	if m != nil {
		return m.Sm1
	}
	return nil
}

func init() {
	proto.RegisterType((*M1)(nil), "flagpb.M1")
	proto.RegisterType((*M2)(nil), "flagpb.M2")
	proto.RegisterType((*M3)(nil), "flagpb.M3")
	proto.RegisterType((*MapContainer)(nil), "flagpb.MapContainer")
	proto.RegisterEnum("flagpb.E", E_name, E_value)
}

func init() {
	proto.RegisterFile("github.com/luci/luci-go/client/flagpb/unmarshal_test.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0xab, 0x9b, 0x40,
	0x14, 0xc5, 0x3b, 0xd7, 0x68, 0x93, 0x9b, 0x10, 0xc2, 0xb4, 0x50, 0x09, 0x2d, 0x88, 0x2b, 0x29,
	0xad, 0x56, 0x43, 0xa1, 0xcd, 0xb2, 0x25, 0x8b, 0x2e, 0xdc, 0x4c, 0xa1, 0xcb, 0x16, 0x27, 0xd8,
	0x64, 0x78, 0xfe, 0x09, 0xce, 0xf8, 0x20, 0x1f, 0xf2, 0x7d, 0xa7, 0xc7, 0x8c, 0x9a, 0xe7, 0x03,
	0xdf, 0x46, 0xbd, 0x9c, 0x73, 0xee, 0xef, 0xce, 0x1d, 0x71, 0x7f, 0x12, 0xea, 0xdc, 0xf2, 0xf0,
	0x58, 0x97, 0x51, 0xd1, 0x1e, 0x85, 0x79, 0x7c, 0x3e, 0xd5, 0xd1, 0xb1, 0x10, 0x79, 0xa5, 0xa2,
	0xff, 0x45, 0x76, 0xba, 0xf0, 0xa8, 0xad, 0xca, 0xac, 0x91, 0xe7, 0xac, 0xf8, 0xa7, 0x72, 0xa9,
	0xc2, 0x4b, 0x53, 0xab, 0x9a, 0x3a, 0x9d, 0xe8, 0xff, 0x45, 0x48, 0x63, 0xba, 0x42, 0x22, 0x5d,
	0xe2, 0x91, 0x60, 0xc1, 0x88, 0xd4, 0x95, 0x70, 0xc1, 0x23, 0x81, 0xcd, 0x88, 0xa0, 0x6b, 0x84,
	0x46, 0xb8, 0x96, 0x67, 0x05, 0x36, 0x83, 0x46, 0x68, 0x95, 0xbb, 0x33, 0x8f, 0x04, 0x73, 0x46,
	0xb8, 0x51, 0xb9, 0xeb, 0x78, 0x56, 0x30, 0x67, 0xd0, 0x98, 0x9a, 0x73, 0xd7, 0xf6, 0x48, 0xb0,
	0x62, 0xc0, 0xb9, 0xff, 0x1d, 0x21, 0x4d, 0xe8, 0x16, 0xa1, 0x8c, 0x0d, 0x60, 0x99, 0x60, 0xd8,
	0xa1, 0xc3, 0x34, 0x66, 0x50, 0xc6, 0xf4, 0x1d, 0x92, 0xdc, 0xd0, 0xd6, 0xc9, 0x62, 0x90, 0x0e,
	0x8c, 0xe4, 0xfe, 0x19, 0x21, 0xdd, 0xdd, 0xa2, 0xd6, 0x44, 0x54, 0x6b, 0x89, 0xc9, 0x8e, 0xb5,
	0x84, 0x41, 0x99, 0x74, 0x63, 0x5a, 0xc3, 0x98, 0xe6, 0x80, 0xb3, 0xe1, 0x80, 0x7a, 0x48, 0x75,
	0x1b, 0x52, 0xf9, 0x0f, 0x80, 0xab, 0x34, 0xbb, 0xfc, 0xac, 0x2b, 0x95, 0x89, 0x2a, 0x6f, 0xe8,
	0x27, 0x04, 0x29, 0x7b, 0xe8, 0xfb, 0x5b, 0xe3, 0x91, 0x23, 0xfc, 0x2d, 0x0f, 0x95, 0x6a, 0xae,
	0x0c, 0xa4, 0xd4, 0x6e, 0xa1, 0x17, 0xf6, 0xb2, 0xfb, 0x97, 0xe8, 0xdd, 0x42, 0xd0, 0x08, 0x2d,
	0x59, 0xc6, 0x66, 0xa1, 0xcb, 0xe4, 0xc3, 0x74, 0xf3, 0x32, 0xee, 0xfc, 0xda, 0xb9, 0xfd, 0x8a,
	0xaf, 0x7b, 0x1a, 0xdd, 0xa0, 0x75, 0x97, 0x5f, 0xfb, 0x9b, 0xd2, 0x9f, 0xf4, 0x2d, 0xda, 0xf7,
	0x59, 0xd1, 0x76, 0x1b, 0x5c, 0xb0, 0xae, 0xd8, 0xc3, 0x37, 0xa2, 0x63, 0x3d, 0x76, 0x1c, 0xb3,
	0x27, 0x62, 0xf6, 0x38, 0xf6, 0x03, 0xe7, 0x03, 0x7e, 0x02, 0xe7, 0x8d, 0x73, 0xcf, 0x2f, 0xe4,
	0xa9, 0xc7, 0xc7, 0x37, 0x48, 0x0e, 0xd4, 0x41, 0xf8, 0xf3, 0x65, 0xf3, 0xca, 0xbc, 0xe3, 0x0d,
	0xe1, 0x8e, 0xf9, 0xf1, 0x76, 0x8f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xf7, 0xd4, 0x7a, 0xb6,
	0x02, 0x00, 0x00,
}
