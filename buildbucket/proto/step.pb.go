// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/buildbucket/proto/step.proto

package buildbucketpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// A build step.
//
// A step may have children, see name field.
type Step struct {
	// Name of the step, unique within the build.
	// Identifies the step.
	//
	// Pipe character ("|") is reserved to separate parent and child step names.
	// For example, value "a|b" indicates step "b" under step "a".
	// If this is a child step, a parent MUST exist and MUST precede this step in
	// the list of steps.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The timestamp when the step started.
	// Required iff status is STARTED, SUCCESS or FAILURE, or if the step has
	// children.
	// MUST NOT be after start_time/end_time of any of the children.
	StartTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// The timestamp when the step ended.
	// Present iff status is terminal.
	// MUST NOT be before start_time.
	// MUST NOT be before start/end_time of any of the children.
	EndTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Status of the step.
	// Must be specified, i.e. not STATUS_UNSPECIFIED.
	//
	// If the step has children
	//   status MUST NOT be SCHEDULED.
	//   status MUST be STARTED if status of any child is not terminal.
	//
	//   status MUST NOT be "better" than statuses of its children,
	//   where "better" relation is defined by the following order,
	//   from good to bad:
	//     SUCCESS
	//     FAILURE
	//     INFRA_FAILURE
	//     CANCELED
	//   Note that this defines "better" relation only for some statuses.
	//   For those statuses where "better" is not defined, this rule does not
	//   apply.
	Status Status `protobuf:"varint,4,opt,name=status,proto3,enum=buildbucket.v2.Status" json:"status,omitempty"`
	// Logs produced by the step.
	// Log order is up to the step.
	//
	// BigQuery: excluded from rows.
	Logs []*Log `protobuf:"bytes,5,rep,name=logs,proto3" json:"logs,omitempty"`
	// Human-readable summary of the step provided by the step itself,
	// in Markdown format (https://spec.commonmark.org/0.28/).
	//
	// V1 equivalent: combines and supersedes Buildbot's step_text and step links and also supports
	// other formatted text.
	//
	// BigQuery: excluded from rows.
	SummaryMarkdown      string   `protobuf:"bytes,7,opt,name=summary_markdown,json=summaryMarkdown,proto3" json:"summary_markdown,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Step) Reset()         { *m = Step{} }
func (m *Step) String() string { return proto.CompactTextString(m) }
func (*Step) ProtoMessage()    {}
func (*Step) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b6755933d248870, []int{0}
}

func (m *Step) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Step.Unmarshal(m, b)
}
func (m *Step) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Step.Marshal(b, m, deterministic)
}
func (m *Step) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Step.Merge(m, src)
}
func (m *Step) XXX_Size() int {
	return xxx_messageInfo_Step.Size(m)
}
func (m *Step) XXX_DiscardUnknown() {
	xxx_messageInfo_Step.DiscardUnknown(m)
}

var xxx_messageInfo_Step proto.InternalMessageInfo

func (m *Step) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Step) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Step) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *Step) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_STATUS_UNSPECIFIED
}

func (m *Step) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *Step) GetSummaryMarkdown() string {
	if m != nil {
		return m.SummaryMarkdown
	}
	return ""
}

func init() {
	proto.RegisterType((*Step)(nil), "buildbucket.v2.Step")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/buildbucket/proto/step.proto", fileDescriptor_7b6755933d248870)
}

var fileDescriptor_7b6755933d248870 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x31, 0x4f, 0xf3, 0x30,
	0x10, 0x86, 0x95, 0x36, 0x5f, 0xfb, 0xd5, 0x95, 0x0a, 0x32, 0x12, 0x8a, 0xb2, 0x10, 0xb1, 0x10,
	0x16, 0x47, 0xa4, 0x80, 0x84, 0xd8, 0x98, 0x61, 0x49, 0x99, 0x58, 0x2a, 0x27, 0x31, 0x26, 0x6a,
	0xec, 0x8b, 0xec, 0x33, 0x88, 0x7f, 0xc1, 0x4f, 0x46, 0x75, 0x52, 0xa9, 0xb0, 0xd0, 0xed, 0xfc,
	0xde, 0xf3, 0xe8, 0xce, 0x47, 0xae, 0x24, 0xb0, 0xea, 0xcd, 0x80, 0x6a, 0x9c, 0x62, 0x60, 0x64,
	0xd6, 0xba, 0xaa, 0xc9, 0x4a, 0xd7, 0xb4, 0x75, 0xe9, 0xaa, 0x8d, 0xc0, 0xac, 0x33, 0x80, 0x90,
	0x59, 0x14, 0x1d, 0xf3, 0x25, 0x5d, 0xec, 0x75, 0xd9, 0x7b, 0x1e, 0x9f, 0x49, 0x00, 0xd9, 0x8a,
	0x1e, 0x2c, 0xdd, 0x6b, 0x86, 0x8d, 0x12, 0x16, 0xb9, 0x1a, 0x84, 0x78, 0x79, 0xe0, 0x8c, 0x0a,
	0x94, 0x02, 0xdd, 0x4b, 0xe7, 0x5f, 0x23, 0x12, 0xae, 0x50, 0x74, 0x94, 0x92, 0x50, 0x73, 0x25,
	0xa2, 0x20, 0x09, 0xd2, 0x59, 0xe1, 0x6b, 0x7a, 0x47, 0x88, 0x45, 0x6e, 0x70, 0xbd, 0x1d, 0x15,
	0x8d, 0x92, 0x20, 0x9d, 0xe7, 0x31, 0xeb, 0xf7, 0x60, 0xbb, 0x3d, 0xd8, 0xf3, 0x6e, 0x8f, 0x62,
	0xe6, 0xe9, 0xed, 0x9b, 0xde, 0x90, 0xff, 0x42, 0xd7, 0xbd, 0x38, 0xfe, 0x53, 0x9c, 0x0a, 0x5d,
	0x7b, 0x8d, 0x91, 0x89, 0x45, 0x8e, 0xce, 0x46, 0x61, 0x12, 0xa4, 0x8b, 0xfc, 0x94, 0xfd, 0xbc,
	0x02, 0x5b, 0xf9, 0x6e, 0x31, 0x50, 0xf4, 0x82, 0x84, 0x2d, 0x48, 0x1b, 0xfd, 0x4b, 0xc6, 0xe9,
	0x3c, 0x3f, 0xf9, 0x4d, 0x3f, 0x82, 0x2c, 0x3c, 0x40, 0x2f, 0xc9, 0xb1, 0x75, 0x4a, 0x71, 0xf3,
	0xb9, 0x56, 0xdc, 0x6c, 0x6a, 0xf8, 0xd0, 0xd1, 0xd4, 0x7f, 0xf5, 0x68, 0xc8, 0x9f, 0x86, 0xf8,
	0xe1, 0xf6, 0xe5, 0xfa, 0xb0, 0x4b, 0xde, 0xef, 0x25, 0x5d, 0x59, 0x4e, 0x7c, 0xb8, 0xfc, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0x85, 0x01, 0x90, 0x7d, 0xec, 0x01, 0x00, 0x00,
}
