// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/luci/luci-go/dm/api/service/v1/finish_attempt.proto

package dm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// FinishAttemptReq sets the final result of an Attempt.
type FinishAttemptReq struct {
	// required
	Auth *Execution_Auth `protobuf:"bytes,1,opt,name=auth" json:"auth,omitempty"`
	// The result data for this Attempt. The `size` field is recalculated after
	// the data field is normalized, and may be omitted.
	Data *JsonResult `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (m *FinishAttemptReq) Reset()                    { *m = FinishAttemptReq{} }
func (m *FinishAttemptReq) String() string            { return proto.CompactTextString(m) }
func (*FinishAttemptReq) ProtoMessage()               {}
func (*FinishAttemptReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *FinishAttemptReq) GetAuth() *Execution_Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *FinishAttemptReq) GetData() *JsonResult {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*FinishAttemptReq)(nil), "dm.FinishAttemptReq")
}

func init() {
	proto.RegisterFile("github.com/luci/luci-go/dm/api/service/v1/finish_attempt.proto", fileDescriptor2)
}

var fileDescriptor2 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xcd, 0xb1, 0xca, 0xc2, 0x30,
	0x14, 0x05, 0x60, 0x5a, 0xca, 0x3f, 0xe4, 0x07, 0x91, 0x4c, 0xc5, 0x49, 0x3a, 0x88, 0x8b, 0x09,
	0xea, 0xe6, 0x20, 0x74, 0xd0, 0xc1, 0x31, 0x0f, 0x60, 0x49, 0x93, 0xd8, 0x04, 0x9a, 0x26, 0x36,
	0x37, 0xc5, 0xc7, 0x97, 0xa6, 0x2f, 0xe0, 0x72, 0x97, 0x73, 0xee, 0x77, 0xd0, 0xb5, 0x33, 0xa0,
	0x63, 0x4b, 0x84, 0xb3, 0xb4, 0x8f, 0xc2, 0xa4, 0x73, 0xe8, 0x1c, 0x95, 0x96, 0x72, 0x6f, 0x68,
	0x50, 0xe3, 0x64, 0x84, 0xa2, 0xd3, 0x91, 0xbe, 0xcc, 0x60, 0x82, 0x6e, 0x38, 0x80, 0xb2, 0x1e,
	0x88, 0x1f, 0x1d, 0x38, 0x9c, 0x4b, 0xbb, 0xb9, 0xfc, 0x6e, 0x74, 0x23, 0xf7, 0xba, 0x91, 0x1c,
	0xf8, 0xf2, 0x5f, 0x3d, 0xd1, 0xfa, 0x9e, 0xdc, 0x7a, 0x61, 0x99, 0x7a, 0xe3, 0x1d, 0x2a, 0x78,
	0x04, 0x5d, 0x66, 0xdb, 0x6c, 0xff, 0x7f, 0xc2, 0x44, 0x5a, 0x72, 0xfb, 0x28, 0x11, 0xc1, 0xb8,
	0x81, 0xd4, 0x11, 0x34, 0x4b, 0x39, 0xae, 0x50, 0x31, 0x4b, 0x65, 0x9e, 0x7a, 0xab, 0xb9, 0xf7,
	0x08, 0x6e, 0x60, 0x2a, 0xc4, 0x1e, 0x58, 0xca, 0xda, 0xbf, 0x34, 0x73, 0xfe, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x33, 0xe6, 0x96, 0x55, 0xe8, 0x00, 0x00, 0x00,
}
