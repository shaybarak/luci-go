// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/luci/luci-go/logdog/api/endpoints/coordinator/logs/v1/state.proto

package logdog

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// LogStreamState is a bidirectional state value used in UpdateStream calls.
//
// LogStreamState is embeddable in Endpoints request/response structs.
type LogStreamState struct {
	// ProtoVersion is the protobuf version for this stream.
	ProtoVersion string `protobuf:"bytes,1,opt,name=proto_version,json=protoVersion" json:"proto_version,omitempty"`
	// The time when the log stream was registered with the Coordinator.
	Created *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created" json:"created,omitempty"`
	// The stream index of the log stream's terminal message. If the value is -1,
	// the log is still streaming.
	TerminalIndex int64 `protobuf:"varint,3,opt,name=terminal_index,json=terminalIndex" json:"terminal_index,omitempty"`
	// If non-nil, the log stream is archived, and this field contains archival
	// details.
	Archive *LogStreamState_ArchiveInfo `protobuf:"bytes,4,opt,name=archive" json:"archive,omitempty"`
	// Indicates the purged state of a log. A log that has been purged is only
	// acknowledged to administrative clients.
	Purged bool `protobuf:"varint,5,opt,name=purged" json:"purged,omitempty"`
}

func (m *LogStreamState) Reset()                    { *m = LogStreamState{} }
func (m *LogStreamState) String() string            { return proto.CompactTextString(m) }
func (*LogStreamState) ProtoMessage()               {}
func (*LogStreamState) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *LogStreamState) GetProtoVersion() string {
	if m != nil {
		return m.ProtoVersion
	}
	return ""
}

func (m *LogStreamState) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *LogStreamState) GetTerminalIndex() int64 {
	if m != nil {
		return m.TerminalIndex
	}
	return 0
}

func (m *LogStreamState) GetArchive() *LogStreamState_ArchiveInfo {
	if m != nil {
		return m.Archive
	}
	return nil
}

func (m *LogStreamState) GetPurged() bool {
	if m != nil {
		return m.Purged
	}
	return false
}

// ArchiveInfo contains archive details for the log stream.
type LogStreamState_ArchiveInfo struct {
	// The Google Storage URL where the log stream's index is archived.
	IndexUrl string `protobuf:"bytes,1,opt,name=index_url,json=indexUrl" json:"index_url,omitempty"`
	// The Google Storage URL where the log stream's raw stream data is archived.
	StreamUrl string `protobuf:"bytes,2,opt,name=stream_url,json=streamUrl" json:"stream_url,omitempty"`
	// The Google Storage URL where the log stream's assembled data is archived.
	DataUrl string `protobuf:"bytes,3,opt,name=data_url,json=dataUrl" json:"data_url,omitempty"`
	// If true, all log entries between 0 and terminal_index were archived. If
	// false, this indicates that the log stream was not completely loaded into
	// intermediate storage when the archival interval expired.
	Complete bool `protobuf:"varint,4,opt,name=complete" json:"complete,omitempty"`
	// The number of log
	LogEntryCount int64 `protobuf:"varint,5,opt,name=log_entry_count,json=logEntryCount" json:"log_entry_count,omitempty"`
}

func (m *LogStreamState_ArchiveInfo) Reset()                    { *m = LogStreamState_ArchiveInfo{} }
func (m *LogStreamState_ArchiveInfo) String() string            { return proto.CompactTextString(m) }
func (*LogStreamState_ArchiveInfo) ProtoMessage()               {}
func (*LogStreamState_ArchiveInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

func (m *LogStreamState_ArchiveInfo) GetIndexUrl() string {
	if m != nil {
		return m.IndexUrl
	}
	return ""
}

func (m *LogStreamState_ArchiveInfo) GetStreamUrl() string {
	if m != nil {
		return m.StreamUrl
	}
	return ""
}

func (m *LogStreamState_ArchiveInfo) GetDataUrl() string {
	if m != nil {
		return m.DataUrl
	}
	return ""
}

func (m *LogStreamState_ArchiveInfo) GetComplete() bool {
	if m != nil {
		return m.Complete
	}
	return false
}

func (m *LogStreamState_ArchiveInfo) GetLogEntryCount() int64 {
	if m != nil {
		return m.LogEntryCount
	}
	return 0
}

func init() {
	proto.RegisterType((*LogStreamState)(nil), "logdog.LogStreamState")
	proto.RegisterType((*LogStreamState_ArchiveInfo)(nil), "logdog.LogStreamState.ArchiveInfo")
}

func init() {
	proto.RegisterFile("github.com/luci/luci-go/logdog/api/endpoints/coordinator/logs/v1/state.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x6b, 0xdb, 0x30,
	0x14, 0xc6, 0x71, 0xbc, 0x25, 0x8e, 0xb2, 0x64, 0xa0, 0xc3, 0xf0, 0x3c, 0xc6, 0x4c, 0xc6, 0x86,
	0x2f, 0x93, 0xd8, 0xd6, 0x63, 0x2f, 0xa5, 0xf4, 0x10, 0xc8, 0xc9, 0x69, 0x7b, 0x35, 0x8a, 0xad,
	0x28, 0x02, 0x59, 0xcf, 0xc8, 0x72, 0x68, 0xff, 0x8e, 0xfe, 0x13, 0xfd, 0x33, 0x8b, 0xa4, 0xb8,
	0xb4, 0x17, 0xc3, 0xfb, 0x7d, 0xdf, 0xd3, 0xf7, 0x3d, 0xa3, 0xad, 0x90, 0xf6, 0x38, 0xec, 0x49,
	0x0d, 0x2d, 0x55, 0x43, 0x2d, 0xfd, 0xe7, 0x8f, 0x00, 0xaa, 0x40, 0x34, 0x20, 0x28, 0xeb, 0x24,
	0xe5, 0xba, 0xe9, 0x40, 0x6a, 0xdb, 0xd3, 0x1a, 0xc0, 0x34, 0x52, 0x33, 0x0b, 0xc6, 0x19, 0x7a,
	0x7a, 0xfa, 0x4b, 0x7b, 0xcb, 0x2c, 0x27, 0x9d, 0x01, 0x0b, 0x78, 0x1a, 0xb6, 0xb2, 0x1f, 0x02,
	0x40, 0x28, 0x4e, 0x3d, 0xdd, 0x0f, 0x07, 0x6a, 0x65, 0xcb, 0x7b, 0xcb, 0xda, 0x2e, 0x18, 0xd7,
	0x4f, 0x31, 0x5a, 0x6d, 0x41, 0xec, 0xac, 0xe1, 0xac, 0xdd, 0xb9, 0x17, 0xf0, 0x4f, 0xb4, 0xf4,
	0x5a, 0x75, 0xe2, 0xa6, 0x97, 0xa0, 0xd3, 0x28, 0x8f, 0x8a, 0x79, 0xf9, 0xc9, 0xc3, 0xfb, 0xc0,
	0xf0, 0x05, 0x9a, 0xd5, 0x86, 0x33, 0xcb, 0x9b, 0x74, 0x92, 0x47, 0xc5, 0xe2, 0x5f, 0x46, 0x42,
	0x14, 0x19, 0xa3, 0xc8, 0xed, 0x18, 0x55, 0x8e, 0x56, 0xfc, 0x0b, 0xad, 0x2c, 0x37, 0xad, 0xd4,
	0x4c, 0x55, 0x52, 0x37, 0xfc, 0x21, 0x8d, 0xf3, 0xa8, 0x88, 0xcb, 0xe5, 0x48, 0x37, 0x0e, 0xe2,
	0x4b, 0x34, 0x63, 0xa6, 0x3e, 0xca, 0x13, 0x4f, 0x3f, 0xf8, 0xc7, 0xd7, 0x24, 0xdc, 0x43, 0xde,
	0x57, 0x25, 0x57, 0xc1, 0xb5, 0xd1, 0x07, 0x28, 0xc7, 0x15, 0xfc, 0x05, 0x4d, 0xbb, 0xc1, 0x08,
	0xde, 0xa4, 0x1f, 0xf3, 0xa8, 0x48, 0xca, 0xf3, 0x94, 0x3d, 0x47, 0x68, 0xf1, 0x66, 0x01, 0x7f,
	0x43, 0x73, 0xdf, 0xa1, 0x1a, 0x8c, 0x3a, 0xdf, 0x98, 0x78, 0x70, 0x67, 0x14, 0xfe, 0x8e, 0x50,
	0xef, 0x83, 0xbc, 0x3a, 0xf1, 0xea, 0x3c, 0x10, 0x27, 0x7f, 0x45, 0x49, 0xc3, 0x2c, 0xf3, 0x62,
	0xec, 0xc5, 0x99, 0x9b, 0x9d, 0x94, 0xa1, 0xa4, 0x86, 0xb6, 0x53, 0xdc, 0x86, 0xf6, 0x49, 0xf9,
	0x3a, 0xe3, 0xdf, 0xe8, 0xb3, 0x02, 0x51, 0x71, 0x6d, 0xcd, 0x63, 0x55, 0xc3, 0xa0, 0xad, 0xef,
	0x18, 0x97, 0x4b, 0x05, 0xe2, 0xc6, 0xd1, 0x6b, 0x07, 0xf7, 0x53, 0xff, 0x13, 0xff, 0xbf, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xee, 0xce, 0xc7, 0x46, 0x15, 0x02, 0x00, 0x00,
}
