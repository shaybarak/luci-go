// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/resultdb/internal/proto/internal.proto

package internalpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	v1 "go.chromium.org/luci/resultdb/proto/rpc/v1"
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

// A message for storing all the information attached to a page token.
type PageToken struct {
	// Position marks the cursor's start (exclusive). Its interpretation is
	// implementation-specific. For instance, for a Spanner cursor, this is a
	// string slice representation of the Spanner key corresponding to the entry
	// prior to the one at which to start reading, or empty if the cursor is to
	// start at the beginning.
	Position             []string `protobuf:"bytes,1,rep,name=position,proto3" json:"position,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageToken) Reset()         { *m = PageToken{} }
func (m *PageToken) String() string { return proto.CompactTextString(m) }
func (*PageToken) ProtoMessage()    {}
func (*PageToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c1fcf5b501f5f42, []int{0}
}

func (m *PageToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageToken.Unmarshal(m, b)
}
func (m *PageToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageToken.Marshal(b, m, deterministic)
}
func (m *PageToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageToken.Merge(m, src)
}
func (m *PageToken) XXX_Size() int {
	return xxx_messageInfo_PageToken.Size(m)
}
func (m *PageToken) XXX_DiscardUnknown() {
	xxx_messageInfo_PageToken.DiscardUnknown(m)
}

var xxx_messageInfo_PageToken proto.InternalMessageInfo

func (m *PageToken) GetPosition() []string {
	if m != nil {
		return m.Position
	}
	return nil
}

// A container of artifacts.
// Used to store artifacts in Spanner.
type Artifacts struct {
	ArtifactsV1          []*v1.Artifact `protobuf:"bytes,1,rep,name=artifacts_v1,json=artifactsV1,proto3" json:"artifacts_v1,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Artifacts) Reset()         { *m = Artifacts{} }
func (m *Artifacts) String() string { return proto.CompactTextString(m) }
func (*Artifacts) ProtoMessage()    {}
func (*Artifacts) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c1fcf5b501f5f42, []int{1}
}

func (m *Artifacts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artifacts.Unmarshal(m, b)
}
func (m *Artifacts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artifacts.Marshal(b, m, deterministic)
}
func (m *Artifacts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artifacts.Merge(m, src)
}
func (m *Artifacts) XXX_Size() int {
	return xxx_messageInfo_Artifacts.Size(m)
}
func (m *Artifacts) XXX_DiscardUnknown() {
	xxx_messageInfo_Artifacts.DiscardUnknown(m)
}

var xxx_messageInfo_Artifacts proto.InternalMessageInfo

func (m *Artifacts) GetArtifactsV1() []*v1.Artifact {
	if m != nil {
		return m.ArtifactsV1
	}
	return nil
}

func init() {
	proto.RegisterType((*PageToken)(nil), "luci.resultdb.internal.PageToken")
	proto.RegisterType((*Artifacts)(nil), "luci.resultdb.internal.Artifacts")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/resultdb/internal/proto/internal.proto", fileDescriptor_3c1fcf5b501f5f42)
}

var fileDescriptor_3c1fcf5b501f5f42 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4e, 0xcf, 0xd7, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcd, 0x2c, 0xcd, 0xd5, 0xcb, 0x2f, 0x4a, 0xd7, 0xcf, 0x29, 0x4d, 0xce,
	0xd4, 0x2f, 0x4a, 0x2d, 0x2e, 0xcd, 0x29, 0x49, 0x49, 0xd2, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca,
	0x4b, 0xcc, 0xd1, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x87, 0x73, 0xf5, 0xc0, 0x5c, 0x21, 0x31, 0x90,
	0x62, 0x3d, 0x98, 0x62, 0x3d, 0x98, 0xac, 0x94, 0x0d, 0x7e, 0x43, 0x21, 0x66, 0x15, 0x15, 0x24,
	0xeb, 0x97, 0x19, 0xea, 0x97, 0xa4, 0x16, 0x97, 0xc4, 0x43, 0xa4, 0x20, 0xa6, 0x2a, 0xa9, 0x73,
	0x71, 0x06, 0x24, 0xa6, 0xa7, 0x86, 0xe4, 0x67, 0xa7, 0xe6, 0x09, 0x49, 0x71, 0x71, 0x14, 0xe4,
	0x17, 0x67, 0x96, 0x64, 0xe6, 0xe7, 0x49, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x06, 0xc1, 0xf9, 0x4a,
	0x7e, 0x5c, 0x9c, 0x8e, 0x45, 0x25, 0x99, 0x69, 0x89, 0xc9, 0x25, 0xc5, 0x42, 0x8e, 0x5c, 0x3c,
	0x89, 0x30, 0x4e, 0x7c, 0x99, 0x21, 0x58, 0x31, 0xb7, 0x91, 0x9c, 0x1e, 0xaa, 0x13, 0x8b, 0x0a,
	0x92, 0xf5, 0xca, 0x0c, 0xf5, 0x60, 0xda, 0x82, 0xb8, 0xe1, 0x7a, 0xc2, 0x0c, 0x9d, 0x78, 0xa2,
	0xb8, 0x60, 0x5e, 0x28, 0x48, 0x4a, 0x62, 0x03, 0xbb, 0xc6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0x26, 0x78, 0xdf, 0x3f, 0x22, 0x01, 0x00, 0x00,
}
