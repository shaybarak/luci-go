// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/lucicfg/external/crostesting/proto/config/source_tree_test_config.proto

package config

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

// Testing restrictions to apply to a source tree.
type TestRestriction struct {
	// Whether to disable hardware tests.
	DisableHwTests bool `protobuf:"varint,1,opt,name=disable_hw_tests,json=disableHwTests,proto3" json:"disable_hw_tests,omitempty"`
	// Whether to disable image tests.
	DisableImageTests bool `protobuf:"varint,2,opt,name=disable_image_tests,json=disableImageTests,proto3" json:"disable_image_tests,omitempty"`
	// Whether to disable virtual machine tests.
	DisableVmTests       bool     `protobuf:"varint,3,opt,name=disable_vm_tests,json=disableVmTests,proto3" json:"disable_vm_tests,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestRestriction) Reset()         { *m = TestRestriction{} }
func (m *TestRestriction) String() string { return proto.CompactTextString(m) }
func (*TestRestriction) ProtoMessage()    {}
func (*TestRestriction) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4d858b6de7c6c4e, []int{0}
}

func (m *TestRestriction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestRestriction.Unmarshal(m, b)
}
func (m *TestRestriction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestRestriction.Marshal(b, m, deterministic)
}
func (m *TestRestriction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRestriction.Merge(m, src)
}
func (m *TestRestriction) XXX_Size() int {
	return xxx_messageInfo_TestRestriction.Size(m)
}
func (m *TestRestriction) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRestriction.DiscardUnknown(m)
}

var xxx_messageInfo_TestRestriction proto.InternalMessageInfo

func (m *TestRestriction) GetDisableHwTests() bool {
	if m != nil {
		return m.DisableHwTests
	}
	return false
}

func (m *TestRestriction) GetDisableImageTests() bool {
	if m != nil {
		return m.DisableImageTests
	}
	return false
}

func (m *TestRestriction) GetDisableVmTests() bool {
	if m != nil {
		return m.DisableVmTests
	}
	return false
}

// A unit of the CrOS codebase. As of 2019-01, this concept maps 1-to-1 with
// "paths" in the the internal and external full.xml files.
type SourceTree struct {
	// The path of a source tree,
	// e.g. "src/platform2" or "src/third_party/kernel/v4.19".
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SourceTree) Reset()         { *m = SourceTree{} }
func (m *SourceTree) String() string { return proto.CompactTextString(m) }
func (*SourceTree) ProtoMessage()    {}
func (*SourceTree) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4d858b6de7c6c4e, []int{1}
}

func (m *SourceTree) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceTree.Unmarshal(m, b)
}
func (m *SourceTree) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceTree.Marshal(b, m, deterministic)
}
func (m *SourceTree) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceTree.Merge(m, src)
}
func (m *SourceTree) XXX_Size() int {
	return xxx_messageInfo_SourceTree.Size(m)
}
func (m *SourceTree) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceTree.DiscardUnknown(m)
}

var xxx_messageInfo_SourceTree proto.InternalMessageInfo

func (m *SourceTree) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// A set of test restrictions for a source tree.
type SourceTreeTestRestriction struct {
	// A CrOS source tree.
	SourceTree *SourceTree `protobuf:"bytes,1,opt,name=source_tree,json=sourceTree,proto3" json:"source_tree,omitempty"`
	// The test restrictions to apply to the source tree.
	TestRestriction      *TestRestriction `protobuf:"bytes,2,opt,name=test_restriction,json=testRestriction,proto3" json:"test_restriction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SourceTreeTestRestriction) Reset()         { *m = SourceTreeTestRestriction{} }
func (m *SourceTreeTestRestriction) String() string { return proto.CompactTextString(m) }
func (*SourceTreeTestRestriction) ProtoMessage()    {}
func (*SourceTreeTestRestriction) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4d858b6de7c6c4e, []int{2}
}

func (m *SourceTreeTestRestriction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceTreeTestRestriction.Unmarshal(m, b)
}
func (m *SourceTreeTestRestriction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceTreeTestRestriction.Marshal(b, m, deterministic)
}
func (m *SourceTreeTestRestriction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceTreeTestRestriction.Merge(m, src)
}
func (m *SourceTreeTestRestriction) XXX_Size() int {
	return xxx_messageInfo_SourceTreeTestRestriction.Size(m)
}
func (m *SourceTreeTestRestriction) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceTreeTestRestriction.DiscardUnknown(m)
}

var xxx_messageInfo_SourceTreeTestRestriction proto.InternalMessageInfo

func (m *SourceTreeTestRestriction) GetSourceTree() *SourceTree {
	if m != nil {
		return m.SourceTree
	}
	return nil
}

func (m *SourceTreeTestRestriction) GetTestRestriction() *TestRestriction {
	if m != nil {
		return m.TestRestriction
	}
	return nil
}

// Configures test restrictions for all relevant source trees.
// This is the root message.
type SourceTreeTestCfg struct {
	// (Source tree, test restriction) pairs.
	SourceTreeTestRestriction []*SourceTreeTestRestriction `protobuf:"bytes,1,rep,name=source_tree_test_restriction,json=sourceTreeTestRestriction,proto3" json:"source_tree_test_restriction,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                     `json:"-"`
	XXX_unrecognized          []byte                       `json:"-"`
	XXX_sizecache             int32                        `json:"-"`
}

func (m *SourceTreeTestCfg) Reset()         { *m = SourceTreeTestCfg{} }
func (m *SourceTreeTestCfg) String() string { return proto.CompactTextString(m) }
func (*SourceTreeTestCfg) ProtoMessage()    {}
func (*SourceTreeTestCfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4d858b6de7c6c4e, []int{3}
}

func (m *SourceTreeTestCfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceTreeTestCfg.Unmarshal(m, b)
}
func (m *SourceTreeTestCfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceTreeTestCfg.Marshal(b, m, deterministic)
}
func (m *SourceTreeTestCfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceTreeTestCfg.Merge(m, src)
}
func (m *SourceTreeTestCfg) XXX_Size() int {
	return xxx_messageInfo_SourceTreeTestCfg.Size(m)
}
func (m *SourceTreeTestCfg) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceTreeTestCfg.DiscardUnknown(m)
}

var xxx_messageInfo_SourceTreeTestCfg proto.InternalMessageInfo

func (m *SourceTreeTestCfg) GetSourceTreeTestRestriction() []*SourceTreeTestRestriction {
	if m != nil {
		return m.SourceTreeTestRestriction
	}
	return nil
}

func init() {
	proto.RegisterType((*TestRestriction)(nil), "crostesting.TestRestriction")
	proto.RegisterType((*SourceTree)(nil), "crostesting.SourceTree")
	proto.RegisterType((*SourceTreeTestRestriction)(nil), "crostesting.SourceTreeTestRestriction")
	proto.RegisterType((*SourceTreeTestCfg)(nil), "crostesting.SourceTreeTestCfg")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/lucicfg/external/crostesting/proto/config/source_tree_test_config.proto", fileDescriptor_f4d858b6de7c6c4e)
}

var fileDescriptor_f4d858b6de7c6c4e = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x89, 0x15, 0xd1, 0x09, 0xd8, 0x76, 0x3d, 0x68, 0xa1, 0x87, 0x92, 0x83, 0xf4, 0xb4,
	0x0b, 0xf5, 0x22, 0x78, 0x53, 0xa4, 0x7a, 0x8d, 0xc5, 0x83, 0x08, 0x21, 0x5d, 0xa7, 0xdb, 0x85,
	0x26, 0x5b, 0x76, 0xb7, 0xd6, 0x83, 0x5f, 0xc3, 0xab, 0x9f, 0x55, 0x3a, 0x89, 0xe4, 0x8f, 0xf4,
	0xe4, 0x25, 0x09, 0x33, 0xbf, 0x37, 0x6f, 0xde, 0x66, 0xe1, 0x55, 0x19, 0x2e, 0x97, 0xd6, 0x64,
	0x7a, 0x93, 0x71, 0x63, 0x95, 0x58, 0x6d, 0xa4, 0xa6, 0x87, 0x5c, 0x28, 0x81, 0x1f, 0x1e, 0x6d,
	0x9e, 0xae, 0x84, 0xb4, 0xc6, 0x79, 0x74, 0x5e, 0xe7, 0x4a, 0xac, 0xad, 0xf1, 0x46, 0x48, 0x93,
	0x2f, 0xb4, 0x12, 0xce, 0x6c, 0xac, 0xc4, 0xc4, 0x5b, 0xc4, 0x64, 0x07, 0x24, 0x45, 0x9d, 0x13,
	0xc4, 0xc2, 0x9a, 0x2e, 0xfa, 0x0a, 0xa0, 0x3b, 0x43, 0xe7, 0x63, 0x74, 0xde, 0x6a, 0xe9, 0xb5,
	0xc9, 0xd9, 0x18, 0x7a, 0x6f, 0xda, 0xa5, 0xf3, 0x15, 0x26, 0xcb, 0x2d, 0x0d, 0x70, 0x17, 0xc1,
	0x28, 0x18, 0x1f, 0xc7, 0xa7, 0x65, 0xfd, 0x61, 0xbb, 0xd3, 0x38, 0xc6, 0xe1, 0xec, 0x97, 0xd4,
	0x59, 0xaa, 0xb0, 0x84, 0x0f, 0x08, 0xee, 0x97, 0xad, 0xc7, 0x5d, 0xa7, 0xe0, 0x6b, 0x93, 0xdf,
	0xb3, 0x12, 0xee, 0x34, 0x26, 0x3f, 0x67, 0x44, 0x46, 0x23, 0x80, 0x27, 0x4a, 0x31, 0xb3, 0x88,
	0x8c, 0xc1, 0xe1, 0x3a, 0xf5, 0x4b, 0xda, 0xe2, 0x24, 0xa6, 0xef, 0xe8, 0x3b, 0x80, 0x41, 0x85,
	0xb4, 0x33, 0x5c, 0x43, 0x58, 0x3b, 0x05, 0x12, 0x86, 0x93, 0x73, 0x5e, 0x8b, 0xce, 0x2b, 0x71,
	0x0c, 0xae, 0xf2, 0x9a, 0x42, 0x8f, 0xce, 0xcc, 0x56, 0xd3, 0x28, 0x50, 0x38, 0x19, 0x36, 0xe4,
	0x2d, 0xc7, 0xb8, 0xeb, 0x9b, 0x85, 0xe8, 0x13, 0xfa, 0xcd, 0xfd, 0xee, 0x16, 0x8a, 0x29, 0x18,
	0xfe, 0xf9, 0x3b, 0x75, 0xa7, 0x60, 0xd4, 0x19, 0x87, 0x93, 0xcb, 0x3d, 0x8b, 0xb6, 0x3d, 0x07,
	0x6e, 0x5f, 0xeb, 0x76, 0xfa, 0x72, 0xff, 0xbf, 0x5b, 0x74, 0x53, 0xbc, 0xe6, 0x47, 0x54, 0xbc,
	0xfa, 0x09, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x45, 0x82, 0x4c, 0x95, 0x02, 0x00, 0x00,
}
