// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/luci/luci-go/common/tsmon/ts_mon_proto/acquisition_task.proto

package ts_mon_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Task_TypeId int32

const (
	Task_MESSAGE_TYPE_ID Task_TypeId = 34049749
)

var Task_TypeId_name = map[int32]string{
	34049749: "MESSAGE_TYPE_ID",
}
var Task_TypeId_value = map[string]int32{
	"MESSAGE_TYPE_ID": 34049749,
}

func (x Task_TypeId) Enum() *Task_TypeId {
	p := new(Task_TypeId)
	*p = x
	return p
}
func (x Task_TypeId) String() string {
	return proto.EnumName(Task_TypeId_name, int32(x))
}
func (x *Task_TypeId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Task_TypeId_value, data, "Task_TypeId")
	if err != nil {
		return err
	}
	*x = Task_TypeId(value)
	return nil
}
func (Task_TypeId) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

type Task struct {
	ProxyEnvironment *string `protobuf:"bytes,5,opt,name=proxy_environment,json=proxyEnvironment" json:"proxy_environment,omitempty"`
	AcquisitionName  *string `protobuf:"bytes,10,opt,name=acquisition_name,json=acquisitionName" json:"acquisition_name,omitempty"`
	ServiceName      *string `protobuf:"bytes,20,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	JobName          *string `protobuf:"bytes,30,opt,name=job_name,json=jobName" json:"job_name,omitempty"`
	DataCenter       *string `protobuf:"bytes,40,opt,name=data_center,json=dataCenter" json:"data_center,omitempty"`
	HostName         *string `protobuf:"bytes,50,opt,name=host_name,json=hostName" json:"host_name,omitempty"`
	TaskNum          *int32  `protobuf:"varint,60,opt,name=task_num,json=taskNum" json:"task_num,omitempty"`
	ProxyZone        *string `protobuf:"bytes,70,opt,name=proxy_zone,json=proxyZone" json:"proxy_zone,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Task) GetProxyEnvironment() string {
	if m != nil && m.ProxyEnvironment != nil {
		return *m.ProxyEnvironment
	}
	return ""
}

func (m *Task) GetAcquisitionName() string {
	if m != nil && m.AcquisitionName != nil {
		return *m.AcquisitionName
	}
	return ""
}

func (m *Task) GetServiceName() string {
	if m != nil && m.ServiceName != nil {
		return *m.ServiceName
	}
	return ""
}

func (m *Task) GetJobName() string {
	if m != nil && m.JobName != nil {
		return *m.JobName
	}
	return ""
}

func (m *Task) GetDataCenter() string {
	if m != nil && m.DataCenter != nil {
		return *m.DataCenter
	}
	return ""
}

func (m *Task) GetHostName() string {
	if m != nil && m.HostName != nil {
		return *m.HostName
	}
	return ""
}

func (m *Task) GetTaskNum() int32 {
	if m != nil && m.TaskNum != nil {
		return *m.TaskNum
	}
	return 0
}

func (m *Task) GetProxyZone() string {
	if m != nil && m.ProxyZone != nil {
		return *m.ProxyZone
	}
	return ""
}

func init() {
	proto.RegisterType((*Task)(nil), "ts_mon.proto.Task")
	proto.RegisterEnum("ts_mon.proto.Task_TypeId", Task_TypeId_name, Task_TypeId_value)
}

func init() {
	proto.RegisterFile("github.com/luci/luci-go/common/tsmon/ts_mon_proto/acquisition_task.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xc1, 0x4a, 0x72, 0x41,
	0x14, 0xc7, 0xf1, 0xe3, 0x33, 0xf5, 0x28, 0x78, 0x1b, 0x22, 0x6e, 0x44, 0x65, 0xae, 0x8c, 0xc8,
	0x0b, 0xad, 0xdb, 0x44, 0xdd, 0xca, 0x45, 0x12, 0xea, 0xa6, 0x36, 0xc3, 0x38, 0x0e, 0x3a, 0xda,
	0xcc, 0xb1, 0x3b, 0x73, 0x25, 0x7b, 0x97, 0x7a, 0x8c, 0xde, 0xa4, 0xf7, 0x89, 0x7b, 0xa6, 0xc8,
	0xcd, 0x81, 0xf3, 0xfb, 0xfd, 0xcf, 0x81, 0x3f, 0xdc, 0x4d, 0xb5, 0x9f, 0xe5, 0xe3, 0xae, 0x44,
	0x93, 0x3c, 0xe7, 0x52, 0xd3, 0x38, 0x9b, 0x62, 0x22, 0xd1, 0x18, 0xb4, 0x89, 0x77, 0x61, 0x72,
	0x83, 0x96, 0x2f, 0x33, 0xf4, 0x98, 0x08, 0xf9, 0x92, 0x6b, 0xa7, 0xbd, 0x46, 0xcb, 0xbd, 0x70,
	0x8b, 0x2e, 0x61, 0xd6, 0x08, 0xa1, 0xb0, 0xb5, 0x3f, 0xff, 0xc1, 0xff, 0x91, 0x70, 0x0b, 0x76,
	0x0a, 0xdb, 0xcb, 0x0c, 0x5f, 0xd7, 0x5c, 0xd9, 0x95, 0xce, 0xd0, 0x1a, 0x65, 0x7d, 0x5c, 0x6e,
	0x95, 0x3a, 0xb5, 0x41, 0x44, 0x22, 0xfd, 0xe3, 0xec, 0x04, 0xa2, 0xcd, 0xef, 0x56, 0x18, 0x15,
	0x03, 0x65, 0x9b, 0x1b, 0xbc, 0x2f, 0x8c, 0x62, 0xc7, 0xd0, 0x70, 0x2a, 0x5b, 0x69, 0xa9, 0x42,
	0x6c, 0x87, 0x62, 0xf5, 0x1f, 0x46, 0x91, 0x3d, 0xa8, 0xce, 0x71, 0x1c, 0xf4, 0x21, 0xe9, 0xca,
	0x1c, 0xc7, 0xa4, 0x8e, 0xa0, 0x3e, 0x11, 0x5e, 0x70, 0xa9, 0xac, 0x57, 0x59, 0xdc, 0x21, 0x0b,
	0x05, 0xba, 0x22, 0xc2, 0xf6, 0xa1, 0x36, 0x43, 0xe7, 0xc3, 0xf1, 0x39, 0xe9, 0x6a, 0x01, 0x7e,
	0x1f, 0x17, 0xc5, 0xb9, 0xcd, 0x4d, 0x7c, 0xd1, 0x2a, 0x75, 0xca, 0x83, 0x4a, 0xb1, 0xf7, 0x73,
	0xc3, 0x0e, 0x00, 0x42, 0xdd, 0x37, 0xb4, 0x2a, 0xbe, 0xa1, 0xc3, 0x1a, 0x91, 0x27, 0xb4, 0xaa,
	0xdd, 0x82, 0xad, 0xd1, 0x7a, 0xa9, 0x7a, 0x13, 0xb6, 0x0b, 0xcd, 0xfb, 0x74, 0x38, 0xbc, 0xbc,
	0x4d, 0xf9, 0xe8, 0xf1, 0x21, 0xe5, 0xbd, 0xeb, 0xe8, 0xeb, 0xfd, 0x23, 0xfa, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0x29, 0x6c, 0x5b, 0xc3, 0x91, 0x01, 0x00, 0x00,
}
