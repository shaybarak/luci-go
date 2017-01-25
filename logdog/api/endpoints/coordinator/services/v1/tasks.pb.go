// Code generated by protoc-gen-go.
// source: github.com/luci/luci-go/logdog/api/endpoints/coordinator/services/v1/tasks.proto
// DO NOT EDIT!

package logdog

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/duration"
import google_protobuf2 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ArchiveTask is a task queue task description for the archival of a single
// log stream.
type ArchiveTask struct {
	// The name of the project that this stream is bound to.
	Project string `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	// The hash ID of the log stream to archive.
	Id string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	// The archival key of the log stream. If this key doesn't match the key in
	// the log stream state, the request is superfluous and should be deleted.
	Key []byte `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	// Don't waste time archiving the log stream until it is at least this old.
	//
	// This is in place to prevent overly-aggressive archivals from wasting time
	// trying, then failing, becuase the log stream data is still being collected
	// into intermediate storage.
	SettleDelay *google_protobuf.Duration `protobuf:"bytes,4,opt,name=settle_delay,json=settleDelay" json:"settle_delay,omitempty"`
	// The amount of time after the task was created that log stream completeness
	// will be used as a success criteria. If the task's age is older than this
	// value, completeness will not be enforced.
	//
	// The task's age can be calculated by subtracting its lease expiration time
	// (leaseTimestamp) from its enqueued timestamp (enqueueTimestamp).
	CompletePeriod *google_protobuf.Duration `protobuf:"bytes,5,opt,name=complete_period,json=completePeriod" json:"complete_period,omitempty"`
	// The time when this archive task was dispatched.
	//
	// This time is optional, and will be based on the Coordinator's clock. If not
	// zero, it can be used by the Archivist to avoid superfluous archival
	// processing by asserting that IF this time is close to the Archivist's local
	// clock by a specific threshold, it will punt the archival task.
	//
	// Because archival is dispatched by Tumble, the actual encoding of archival
	// parameters is oftentimes delayed such that the request is dispatched to
	// Pub/Sub before the datastore has been updated.
	DispatchedAt *google_protobuf2.Timestamp `protobuf:"bytes,6,opt,name=dispatched_at,json=dispatchedAt" json:"dispatched_at,omitempty"`
}

func (m *ArchiveTask) Reset()                    { *m = ArchiveTask{} }
func (m *ArchiveTask) String() string            { return proto.CompactTextString(m) }
func (*ArchiveTask) ProtoMessage()               {}
func (*ArchiveTask) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ArchiveTask) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *ArchiveTask) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ArchiveTask) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *ArchiveTask) GetSettleDelay() *google_protobuf.Duration {
	if m != nil {
		return m.SettleDelay
	}
	return nil
}

func (m *ArchiveTask) GetCompletePeriod() *google_protobuf.Duration {
	if m != nil {
		return m.CompletePeriod
	}
	return nil
}

func (m *ArchiveTask) GetDispatchedAt() *google_protobuf2.Timestamp {
	if m != nil {
		return m.DispatchedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*ArchiveTask)(nil), "logdog.ArchiveTask")
}

func init() {
	proto.RegisterFile("github.com/luci/luci-go/logdog/api/endpoints/coordinator/services/v1/tasks.proto", fileDescriptor2)
}

var fileDescriptor2 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4b, 0x33, 0x31,
	0x10, 0x86, 0xd9, 0xed, 0xf7, 0x55, 0x4c, 0x6b, 0x95, 0x9c, 0x62, 0x0f, 0x5a, 0x3c, 0xf5, 0xe2,
	0x06, 0xf5, 0x2a, 0x48, 0xa5, 0x3f, 0xa0, 0x2c, 0xbd, 0x97, 0x34, 0x19, 0xb7, 0x63, 0x77, 0x77,
	0x42, 0x32, 0x5b, 0xe8, 0x7f, 0xf0, 0x47, 0x4b, 0x77, 0x5d, 0x04, 0x3d, 0x78, 0x09, 0xc9, 0xe4,
	0x79, 0x5e, 0x5e, 0x46, 0xac, 0x0a, 0xe4, 0x5d, 0xb3, 0xcd, 0x2c, 0x55, 0xba, 0x6c, 0x2c, 0xb6,
	0xc7, 0x7d, 0x41, 0xba, 0xa4, 0xc2, 0x51, 0xa1, 0x8d, 0x47, 0x0d, 0xb5, 0xf3, 0x84, 0x35, 0x47,
	0x6d, 0x89, 0x82, 0xc3, 0xda, 0x30, 0x05, 0x1d, 0x21, 0x1c, 0xd0, 0x42, 0xd4, 0x87, 0x07, 0xcd,
	0x26, 0xee, 0x63, 0xe6, 0x03, 0x31, 0xc9, 0x61, 0x67, 0x4e, 0x6f, 0x0a, 0xa2, 0xa2, 0x04, 0xdd,
	0x4e, 0xb7, 0xcd, 0x9b, 0x76, 0x4d, 0x30, 0x8c, 0x54, 0x77, 0xdc, 0xf4, 0xf6, 0xe7, 0x3f, 0x63,
	0x05, 0x91, 0x4d, 0xe5, 0x3b, 0xe0, 0xee, 0x23, 0x15, 0xa3, 0x45, 0xb0, 0x3b, 0x3c, 0xc0, 0xda,
	0xc4, 0xbd, 0x54, 0xe2, 0xcc, 0x07, 0x7a, 0x07, 0xcb, 0x2a, 0x99, 0x25, 0xf3, 0xf3, 0xbc, 0x7f,
	0xca, 0x89, 0x48, 0xd1, 0xa9, 0xb4, 0x1d, 0xa6, 0xe8, 0xe4, 0x95, 0x18, 0xec, 0xe1, 0xa8, 0x06,
	0xb3, 0x64, 0x3e, 0xce, 0x4f, 0x57, 0xf9, 0x2c, 0xc6, 0x11, 0x98, 0x4b, 0xd8, 0x38, 0x28, 0xcd,
	0x51, 0xfd, 0x9b, 0x25, 0xf3, 0xd1, 0xe3, 0x75, 0xd6, 0x75, 0xc8, 0xfa, 0x0e, 0xd9, 0xf2, 0xab,
	0x63, 0x3e, 0xea, 0xf0, 0xe5, 0x89, 0x96, 0xaf, 0xe2, 0xd2, 0x52, 0xe5, 0x4b, 0x60, 0xd8, 0x78,
	0x08, 0x48, 0x4e, 0xfd, 0xff, 0x2b, 0x60, 0xd2, 0x1b, 0xab, 0x56, 0x90, 0x2f, 0xe2, 0xc2, 0x61,
	0xf4, 0x86, 0xed, 0x0e, 0xdc, 0xc6, 0xb0, 0x1a, 0xb6, 0x09, 0xd3, 0x5f, 0x09, 0xeb, 0x7e, 0x0d,
	0xf9, 0xf8, 0x5b, 0x58, 0xf0, 0x76, 0xd8, 0x12, 0x4f, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0a,
	0x85, 0xe4, 0x2b, 0xb2, 0x01, 0x00, 0x00,
}
