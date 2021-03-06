// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/scheduler/appengine/task/gitiles/pb/messages.proto

package pb

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

// Child is the last part and its sha1 tip.
type Child struct {
	Suffix               string   `protobuf:"bytes,1,opt,name=suffix,proto3" json:"suffix,omitempty"`
	Sha1                 []byte   `protobuf:"bytes,2,opt,name=sha1,proto3" json:"sha1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Child) Reset()         { *m = Child{} }
func (m *Child) String() string { return proto.CompactTextString(m) }
func (*Child) ProtoMessage()    {}
func (*Child) Descriptor() ([]byte, []int) {
	return fileDescriptor_404b8bcd5f74310a, []int{0}
}

func (m *Child) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Child.Unmarshal(m, b)
}
func (m *Child) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Child.Marshal(b, m, deterministic)
}
func (m *Child) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Child.Merge(m, src)
}
func (m *Child) XXX_Size() int {
	return xxx_messageInfo_Child.Size(m)
}
func (m *Child) XXX_DiscardUnknown() {
	xxx_messageInfo_Child.DiscardUnknown(m)
}

var xxx_messageInfo_Child proto.InternalMessageInfo

func (m *Child) GetSuffix() string {
	if m != nil {
		return m.Suffix
	}
	return ""
}

func (m *Child) GetSha1() []byte {
	if m != nil {
		return m.Sha1
	}
	return nil
}

// RefSpace is a bunch of children which share the same ref namespace (prefix).
type RefSpace struct {
	Prefix               string   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Children             []*Child `protobuf:"bytes,2,rep,name=children,proto3" json:"children,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefSpace) Reset()         { *m = RefSpace{} }
func (m *RefSpace) String() string { return proto.CompactTextString(m) }
func (*RefSpace) ProtoMessage()    {}
func (*RefSpace) Descriptor() ([]byte, []int) {
	return fileDescriptor_404b8bcd5f74310a, []int{1}
}

func (m *RefSpace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefSpace.Unmarshal(m, b)
}
func (m *RefSpace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefSpace.Marshal(b, m, deterministic)
}
func (m *RefSpace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefSpace.Merge(m, src)
}
func (m *RefSpace) XXX_Size() int {
	return xxx_messageInfo_RefSpace.Size(m)
}
func (m *RefSpace) XXX_DiscardUnknown() {
	xxx_messageInfo_RefSpace.DiscardUnknown(m)
}

var xxx_messageInfo_RefSpace proto.InternalMessageInfo

func (m *RefSpace) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *RefSpace) GetChildren() []*Child {
	if m != nil {
		return m.Children
	}
	return nil
}

// RepositoryState stores tips of all watched refs in a repo.
type RepositoryState struct {
	Spaces               []*RefSpace `protobuf:"bytes,1,rep,name=spaces,proto3" json:"spaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RepositoryState) Reset()         { *m = RepositoryState{} }
func (m *RepositoryState) String() string { return proto.CompactTextString(m) }
func (*RepositoryState) ProtoMessage()    {}
func (*RepositoryState) Descriptor() ([]byte, []int) {
	return fileDescriptor_404b8bcd5f74310a, []int{2}
}

func (m *RepositoryState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepositoryState.Unmarshal(m, b)
}
func (m *RepositoryState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepositoryState.Marshal(b, m, deterministic)
}
func (m *RepositoryState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepositoryState.Merge(m, src)
}
func (m *RepositoryState) XXX_Size() int {
	return xxx_messageInfo_RepositoryState.Size(m)
}
func (m *RepositoryState) XXX_DiscardUnknown() {
	xxx_messageInfo_RepositoryState.DiscardUnknown(m)
}

var xxx_messageInfo_RepositoryState proto.InternalMessageInfo

func (m *RepositoryState) GetSpaces() []*RefSpace {
	if m != nil {
		return m.Spaces
	}
	return nil
}

// DebugState is returned as part of GetDebugJobState RPC response.
type DebugState struct {
	Known                []*DebugState_Ref `protobuf:"bytes,1,rep,name=known,proto3" json:"known,omitempty"`
	Current              []*DebugState_Ref `protobuf:"bytes,2,rep,name=current,proto3" json:"current,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DebugState) Reset()         { *m = DebugState{} }
func (m *DebugState) String() string { return proto.CompactTextString(m) }
func (*DebugState) ProtoMessage()    {}
func (*DebugState) Descriptor() ([]byte, []int) {
	return fileDescriptor_404b8bcd5f74310a, []int{3}
}

func (m *DebugState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugState.Unmarshal(m, b)
}
func (m *DebugState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugState.Marshal(b, m, deterministic)
}
func (m *DebugState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugState.Merge(m, src)
}
func (m *DebugState) XXX_Size() int {
	return xxx_messageInfo_DebugState.Size(m)
}
func (m *DebugState) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugState.DiscardUnknown(m)
}

var xxx_messageInfo_DebugState proto.InternalMessageInfo

func (m *DebugState) GetKnown() []*DebugState_Ref {
	if m != nil {
		return m.Known
	}
	return nil
}

func (m *DebugState) GetCurrent() []*DebugState_Ref {
	if m != nil {
		return m.Current
	}
	return nil
}

type DebugState_Ref struct {
	Ref                  string   `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	Commit               string   `protobuf:"bytes,2,opt,name=commit,proto3" json:"commit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DebugState_Ref) Reset()         { *m = DebugState_Ref{} }
func (m *DebugState_Ref) String() string { return proto.CompactTextString(m) }
func (*DebugState_Ref) ProtoMessage()    {}
func (*DebugState_Ref) Descriptor() ([]byte, []int) {
	return fileDescriptor_404b8bcd5f74310a, []int{3, 0}
}

func (m *DebugState_Ref) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugState_Ref.Unmarshal(m, b)
}
func (m *DebugState_Ref) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugState_Ref.Marshal(b, m, deterministic)
}
func (m *DebugState_Ref) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugState_Ref.Merge(m, src)
}
func (m *DebugState_Ref) XXX_Size() int {
	return xxx_messageInfo_DebugState_Ref.Size(m)
}
func (m *DebugState_Ref) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugState_Ref.DiscardUnknown(m)
}

var xxx_messageInfo_DebugState_Ref proto.InternalMessageInfo

func (m *DebugState_Ref) GetRef() string {
	if m != nil {
		return m.Ref
	}
	return ""
}

func (m *DebugState_Ref) GetCommit() string {
	if m != nil {
		return m.Commit
	}
	return ""
}

func init() {
	proto.RegisterType((*Child)(nil), "gitiles.messages.Child")
	proto.RegisterType((*RefSpace)(nil), "gitiles.messages.RefSpace")
	proto.RegisterType((*RepositoryState)(nil), "gitiles.messages.RepositoryState")
	proto.RegisterType((*DebugState)(nil), "gitiles.messages.DebugState")
	proto.RegisterType((*DebugState_Ref)(nil), "gitiles.messages.DebugState.Ref")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/scheduler/appengine/task/gitiles/pb/messages.proto", fileDescriptor_404b8bcd5f74310a)
}

var fileDescriptor_404b8bcd5f74310a = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x3b, 0x4f, 0xc3, 0x30,
	0x14, 0x85, 0x95, 0xbe, 0x68, 0x2f, 0x48, 0x54, 0x1e, 0x20, 0xea, 0x14, 0x65, 0xea, 0x14, 0x8b,
	0x56, 0x62, 0x60, 0xe4, 0x21, 0x76, 0x77, 0x40, 0x62, 0x4b, 0xdc, 0x9b, 0xc4, 0x6a, 0x12, 0x5b,
	0x7e, 0x08, 0xf8, 0x59, 0xfc, 0x43, 0xe4, 0xd4, 0x01, 0x44, 0x07, 0xb6, 0x7b, 0x93, 0x73, 0xbe,
	0x63, 0x1f, 0xc3, 0x73, 0x25, 0x33, 0x5e, 0x6b, 0xd9, 0x0a, 0xd7, 0x66, 0x52, 0x57, 0xb4, 0x71,
	0x5c, 0x50, 0xc3, 0x6b, 0xdc, 0xbb, 0x06, 0x35, 0xcd, 0x95, 0xc2, 0xae, 0x12, 0x1d, 0x52, 0x9b,
	0x9b, 0x03, 0xad, 0x84, 0x15, 0x0d, 0x1a, 0xaa, 0x0a, 0xda, 0xa2, 0x31, 0x79, 0x85, 0x26, 0x53,
	0x5a, 0x5a, 0x49, 0x96, 0xe1, 0x57, 0x36, 0x7c, 0x4f, 0xb7, 0x30, 0x7d, 0xa8, 0x45, 0xb3, 0x27,
	0x57, 0x30, 0x33, 0xae, 0x2c, 0xc5, 0x7b, 0x1c, 0x25, 0xd1, 0x7a, 0xc1, 0xc2, 0x46, 0x08, 0x4c,
	0x4c, 0x9d, 0xdf, 0xc4, 0xa3, 0x24, 0x5a, 0x5f, 0xb0, 0x7e, 0x4e, 0x5f, 0x60, 0xce, 0xb0, 0xdc,
	0xa9, 0x9c, 0xa3, 0xf7, 0x29, 0x8d, 0xbf, 0x7c, 0xc7, 0x8d, 0x6c, 0x61, 0xce, 0x3d, 0x58, 0x63,
	0x17, 0x8f, 0x92, 0xf1, 0xfa, 0x7c, 0x73, 0x9d, 0xfd, 0x4d, 0xcf, 0xfa, 0x68, 0xf6, 0x2d, 0x4c,
	0x9f, 0xe0, 0x92, 0xa1, 0x92, 0x46, 0x58, 0xa9, 0x3f, 0x76, 0x36, 0xb7, 0x48, 0x36, 0x30, 0x33,
	0x3e, 0xc8, 0xc4, 0x51, 0x4f, 0x59, 0x9d, 0x52, 0x86, 0xb3, 0xb0, 0xa0, 0x4c, 0x3f, 0x23, 0x80,
	0x47, 0x2c, 0x5c, 0x75, 0x44, 0xdc, 0xc2, 0xf4, 0xd0, 0xc9, 0xb7, 0x2e, 0x10, 0x92, 0x53, 0xc2,
	0x8f, 0xd8, 0xc3, 0xd8, 0x51, 0x4e, 0xee, 0xe0, 0x8c, 0x3b, 0xad, 0xb1, 0xb3, 0xe1, 0x06, 0xff,
	0x3b, 0x07, 0xc3, 0x8a, 0xc2, 0x98, 0x61, 0x49, 0x96, 0x30, 0xd6, 0x58, 0x86, 0x6a, 0xfc, 0xe8,
	0xfb, 0xe2, 0xb2, 0x6d, 0x85, 0xed, 0x1b, 0x5d, 0xb0, 0xb0, 0xdd, 0x4f, 0x5e, 0x47, 0xaa, 0x28,
	0x66, 0xfd, 0x3b, 0x6d, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x50, 0xc2, 0xcc, 0xf2, 0x01,
	0x00, 0x00,
}
