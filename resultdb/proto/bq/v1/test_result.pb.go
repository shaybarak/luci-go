// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/resultdb/proto/bq/v1/test_result.proto

package bqpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	v1 "go.chromium.org/luci/resultdb/proto/rpc/v1"
	_type "go.chromium.org/luci/resultdb/proto/type"
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

// Represents a test result row in a BigQuery table.
type TestResultRow struct {
	// Invocation level information.
	Invocation *TestResultRow_Invocation `protobuf:"bytes,1,opt,name=invocation,proto3" json:"invocation,omitempty"`
	// A result of a functional test case.
	// Refer to ../../rpc/v1/test_result.proto for definition.
	Result *v1.TestResult `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	// Indicates that the test subject (e.g. a CL) is absolved from blame
	// for this result.
	Exoneration          *TestResultRow_TestExoneration `protobuf:"bytes,3,opt,name=exoneration,proto3" json:"exoneration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *TestResultRow) Reset()         { *m = TestResultRow{} }
func (m *TestResultRow) String() string { return proto.CompactTextString(m) }
func (*TestResultRow) ProtoMessage()    {}
func (*TestResultRow) Descriptor() ([]byte, []int) {
	return fileDescriptor_45e9665d0ffb1b0a, []int{0}
}

func (m *TestResultRow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResultRow.Unmarshal(m, b)
}
func (m *TestResultRow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResultRow.Marshal(b, m, deterministic)
}
func (m *TestResultRow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResultRow.Merge(m, src)
}
func (m *TestResultRow) XXX_Size() int {
	return xxx_messageInfo_TestResultRow.Size(m)
}
func (m *TestResultRow) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResultRow.DiscardUnknown(m)
}

var xxx_messageInfo_TestResultRow proto.InternalMessageInfo

func (m *TestResultRow) GetInvocation() *TestResultRow_Invocation {
	if m != nil {
		return m.Invocation
	}
	return nil
}

func (m *TestResultRow) GetResult() *v1.TestResult {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *TestResultRow) GetExoneration() *TestResultRow_TestExoneration {
	if m != nil {
		return m.Exoneration
	}
	return nil
}

// A subset of luci.resultdb.rpc.v1.Invocation message
// in ../../rpc/v1/invocation.proto.
type TestResultRow_Invocation struct {
	// Invocation-level string key-value pairs.
	// A key can be repeated.
	Tags                 []*_type.StringPair `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TestResultRow_Invocation) Reset()         { *m = TestResultRow_Invocation{} }
func (m *TestResultRow_Invocation) String() string { return proto.CompactTextString(m) }
func (*TestResultRow_Invocation) ProtoMessage()    {}
func (*TestResultRow_Invocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_45e9665d0ffb1b0a, []int{0, 0}
}

func (m *TestResultRow_Invocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResultRow_Invocation.Unmarshal(m, b)
}
func (m *TestResultRow_Invocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResultRow_Invocation.Marshal(b, m, deterministic)
}
func (m *TestResultRow_Invocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResultRow_Invocation.Merge(m, src)
}
func (m *TestResultRow_Invocation) XXX_Size() int {
	return xxx_messageInfo_TestResultRow_Invocation.Size(m)
}
func (m *TestResultRow_Invocation) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResultRow_Invocation.DiscardUnknown(m)
}

var xxx_messageInfo_TestResultRow_Invocation proto.InternalMessageInfo

func (m *TestResultRow_Invocation) GetTags() []*_type.StringPair {
	if m != nil {
		return m.Tags
	}
	return nil
}

type TestResultRow_TestExoneration struct {
	// True if the test subject (e.g. a CL) is absolved from blame for this
	// result, otherwise False.
	// For more details, refer to luci.resultdb.rpc.v1.TestExoneration
	// in ../../rpc/v1/test_result.proto.
	Exonerated           bool     `protobuf:"varint,1,opt,name=exonerated,proto3" json:"exonerated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestResultRow_TestExoneration) Reset()         { *m = TestResultRow_TestExoneration{} }
func (m *TestResultRow_TestExoneration) String() string { return proto.CompactTextString(m) }
func (*TestResultRow_TestExoneration) ProtoMessage()    {}
func (*TestResultRow_TestExoneration) Descriptor() ([]byte, []int) {
	return fileDescriptor_45e9665d0ffb1b0a, []int{0, 1}
}

func (m *TestResultRow_TestExoneration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResultRow_TestExoneration.Unmarshal(m, b)
}
func (m *TestResultRow_TestExoneration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResultRow_TestExoneration.Marshal(b, m, deterministic)
}
func (m *TestResultRow_TestExoneration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResultRow_TestExoneration.Merge(m, src)
}
func (m *TestResultRow_TestExoneration) XXX_Size() int {
	return xxx_messageInfo_TestResultRow_TestExoneration.Size(m)
}
func (m *TestResultRow_TestExoneration) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResultRow_TestExoneration.DiscardUnknown(m)
}

var xxx_messageInfo_TestResultRow_TestExoneration proto.InternalMessageInfo

func (m *TestResultRow_TestExoneration) GetExonerated() bool {
	if m != nil {
		return m.Exonerated
	}
	return false
}

func init() {
	proto.RegisterType((*TestResultRow)(nil), "luci.resultdb.bq.v1.TestResultRow")
	proto.RegisterType((*TestResultRow_Invocation)(nil), "luci.resultdb.bq.v1.TestResultRow.Invocation")
	proto.RegisterType((*TestResultRow_TestExoneration)(nil), "luci.resultdb.bq.v1.TestResultRow.TestExoneration")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/resultdb/proto/bq/v1/test_result.proto", fileDescriptor_45e9665d0ffb1b0a)
}

var fileDescriptor_45e9665d0ffb1b0a = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xd9, 0x26, 0x43, 0x5e, 0x11, 0x21, 0x5e, 0x4a, 0x0f, 0xa3, 0x78, 0xda, 0xc5, 0xc4,
	0x56, 0x04, 0x61, 0x1e, 0x44, 0xf0, 0xe0, 0x41, 0x90, 0xb8, 0x93, 0x17, 0x69, 0xb2, 0x50, 0x03,
	0x6b, 0x5f, 0x9b, 0xa6, 0x55, 0x3f, 0xaf, 0x5f, 0x44, 0x9a, 0x3a, 0x57, 0xcb, 0x0e, 0x3d, 0xf6,
	0xf5, 0xff, 0xfb, 0xbd, 0x3f, 0x8f, 0xc0, 0x2a, 0x45, 0x2a, 0xdf, 0x0d, 0x66, 0xba, 0xce, 0x28,
	0x9a, 0x94, 0x6d, 0x6b, 0xa9, 0x99, 0x51, 0x55, 0xbd, 0xb5, 0x1b, 0xc1, 0x0a, 0x83, 0x16, 0x99,
	0x28, 0x59, 0x13, 0x31, 0xab, 0x2a, 0xfb, 0xd6, 0xfd, 0xa1, 0x6e, 0x4e, 0xce, 0xda, 0x30, 0xdd,
	0x85, 0xa9, 0x28, 0x69, 0x13, 0x05, 0xb7, 0x63, 0x8c, 0xa6, 0x90, 0x07, 0x95, 0xc1, 0xf5, 0x18,
	0xda, 0x7e, 0x15, 0x8a, 0x49, 0xcc, 0x32, 0xcc, 0x3b, 0xec, 0xfc, 0x7b, 0x0a, 0x27, 0x6b, 0x55,
	0x59, 0xee, 0x82, 0x1c, 0x3f, 0xc8, 0x13, 0x80, 0xce, 0x1b, 0x94, 0x89, 0xd5, 0x98, 0xfb, 0x93,
	0x70, 0xb2, 0xf4, 0xe2, 0x0b, 0x7a, 0xa0, 0x30, 0xfd, 0xc7, 0xd1, 0xc7, 0x3f, 0x88, 0xf7, 0x04,
	0xe4, 0x06, 0xe6, 0x1d, 0xe6, 0x4f, 0x9d, 0x2a, 0x1c, 0xa8, 0x4c, 0x21, 0x07, 0xae, 0xdf, 0x3c,
	0x59, 0x83, 0xa7, 0x3e, 0x31, 0x57, 0xa6, 0x6b, 0x32, 0x73, 0x78, 0x3c, 0xa2, 0x49, 0xfb, 0xf5,
	0xb0, 0x27, 0x79, 0x5f, 0x13, 0xdc, 0x01, 0xec, 0x9b, 0x92, 0x18, 0x8e, 0x6c, 0x92, 0x56, 0xfe,
	0x24, 0x9c, 0x2d, 0xbd, 0x78, 0x31, 0x90, 0xb7, 0xe7, 0xa2, 0x2f, 0xd6, 0xe8, 0x3c, 0x7d, 0x4e,
	0xb4, 0xe1, 0x2e, 0x1b, 0x44, 0x70, 0x3a, 0xd8, 0x40, 0x16, 0x00, 0xbb, 0x1d, 0x6a, 0xe3, 0x6e,
	0x76, 0xcc, 0x7b, 0x93, 0xfb, 0xcb, 0x57, 0x3a, 0xfa, 0xb9, 0xac, 0x44, 0x59, 0x08, 0x31, 0x77,
	0x93, 0xab, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x12, 0x55, 0xdc, 0xad, 0x67, 0x02, 0x00, 0x00,
}
