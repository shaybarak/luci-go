// Code generated by protoc-gen-go.
// source: annotations.proto
// DO NOT EDIT!

/*
Package milo is a generated protocol buffer package.

It is generated from these files:
	annotations.proto

It has these top-level messages:
	FailureDetails
	Step
	Component
	Command
	Progress
	LogdogStream
	IsolateObject
	DMLink
*/
package milo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/luci/luci-go/common/proto/google"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Status is the expressed root step of this step or substep.
type Status int32

const (
	// The step is still running.
	Status_RUNNING Status = 0
	// The step has finished successfully.
	Status_SUCCESS Status = 1
	// The step has finished unsuccessfully.
	Status_FAILURE Status = 2
	// The step has finished unexpectedly.
	Status_EXCEPTION Status = 3
)

var Status_name = map[int32]string{
	0: "RUNNING",
	1: "SUCCESS",
	2: "FAILURE",
	3: "EXCEPTION",
}
var Status_value = map[string]int32{
	"RUNNING":   0,
	"SUCCESS":   1,
	"FAILURE":   2,
	"EXCEPTION": 3,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Type is the type of failure.
type FailureDetails_Type int32

const (
	// The failure is a general failure.
	FailureDetails_GENERAL FailureDetails_Type = 0
	// The failure is related to a failed infrastructure component, not an error
	// with the Step itself.
	FailureDetails_INFRA FailureDetails_Type = 1
	// The failure is due to a failed Dungeon Master dependency. This should be
	// used if a Step's external depdendency fails and the Step cannot recover
	// or proceed without it.
	FailureDetails_DM_DEPENDENCY_FAILED FailureDetails_Type = 2
)

var FailureDetails_Type_name = map[int32]string{
	0: "GENERAL",
	1: "INFRA",
	2: "DM_DEPENDENCY_FAILED",
}
var FailureDetails_Type_value = map[string]int32{
	"GENERAL":              0,
	"INFRA":                1,
	"DM_DEPENDENCY_FAILED": 2,
}

func (x FailureDetails_Type) String() string {
	return proto.EnumName(FailureDetails_Type_name, int32(x))
}
func (FailureDetails_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// FailureType provides more details on the nature of the Status.
type FailureDetails struct {
	Type FailureDetails_Type `protobuf:"varint,1,opt,name=type,enum=milo.FailureDetails_Type" json:"type,omitempty"`
	// An optional string describing the failure.
	Text string `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	// If the failure type is DEPENDENCY_FAILED, the failed dependencies should be
	// listed here.
	FailedDmDependency []*DMLink `protobuf:"bytes,3,rep,name=failed_dm_dependency" json:"failed_dm_dependency,omitempty"`
}

func (m *FailureDetails) Reset()                    { *m = FailureDetails{} }
func (m *FailureDetails) String() string            { return proto.CompactTextString(m) }
func (*FailureDetails) ProtoMessage()               {}
func (*FailureDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FailureDetails) GetFailedDmDependency() []*DMLink {
	if m != nil {
		return m.FailedDmDependency
	}
	return nil
}

// Generic step or substep state.
type Step struct {
	// The command-line invocation of the step, expressed as an argument vector.
	Command *Command `protobuf:"bytes,1,opt,name=command" json:"command,omitempty"`
	// Optional information detailing the failure. This may be populated if the
	// Step's top-level command Status is set to FAILURE.
	FailureDetails *FailureDetails `protobuf:"bytes,2,opt,name=failure_details" json:"failure_details,omitempty"`
	// Base Component information describing the high-level Step.
	StepComponent *Component `protobuf:"bytes,3,opt,name=step_component" json:"step_component,omitempty"`
	// Sub-components that are part of the Step.
	Components []*Component `protobuf:"bytes,4,rep,name=components" json:"components,omitempty"`
	// Substeps will be constructed as extensions on the parent LogDog stream.
	//
	// For example, if the parent's logging base path is:
	// luci/dm/QUEST/ATTEMPT/EXECUTION/+/
	//
	// Its substep #0 will have logging base path:
	// luci/dm/QUEST/ATTEMPT/EXECUTION/+/steps/0
	//
	// ... which can have known log stream names appended to it for the full
	// log stream path. The following appendages are part of the standard
	// Milo protocol expectations:
	// - .../stdout: A text stream containing the Step's STDOUT.
	// - .../stderr: A text stream containing the Step's STDERR.
	// - .../annotation: A text stream containing the Step's annotation stream
	//                   protobuf (Step message protobuf).
	//
	// For example:
	// - luci/dm/QUEST/ATTEMPT/EXECUTION/+/steps/0/stdout
	// - luci/dm/QUEST/ATTEMPT/EXECUTION/+/steps/0/annotations
	SubstepLogdogNameBase []string `protobuf:"bytes,5,rep,name=substep_logdog_name_base" json:"substep_logdog_name_base,omitempty"`
}

func (m *Step) Reset()                    { *m = Step{} }
func (m *Step) String() string            { return proto.CompactTextString(m) }
func (*Step) ProtoMessage()               {}
func (*Step) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Step) GetCommand() *Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *Step) GetFailureDetails() *FailureDetails {
	if m != nil {
		return m.FailureDetails
	}
	return nil
}

func (m *Step) GetStepComponent() *Component {
	if m != nil {
		return m.StepComponent
	}
	return nil
}

func (m *Step) GetComponents() []*Component {
	if m != nil {
		return m.Components
	}
	return nil
}

// A Component represents a renderable state.
type Component struct {
	// The display name of the Component.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Type classifies the information associated with the Note.
	Status Status `protobuf:"varint,2,opt,name=status,enum=milo.Status" json:"status,omitempty"`
	// When the step started, expressed as an RFC3339 string using Z (UTC)
	// timezone.
	Started *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=started" json:"started,omitempty"`
	// When the step ended, expressed as an RFC3339 string using Z (UTC) timezone.
	Ended *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=ended" json:"ended,omitempty"`
	// Arbitrary lines of component text. Each string here is a consecutive line,
	// and should not contain newlines.
	Text []string `protobuf:"bytes,5,rep,name=text" json:"text,omitempty"`
	// The Component's progress.
	Progress *Progress `protobuf:"bytes,6,opt,name=progress" json:"progress,omitempty"`
	// The primary link for this Component. This is the link that interaction
	// with the Component will use.
	Link *Component_Link `protobuf:"bytes,7,opt,name=link" json:"link,omitempty"`
	// Additional links related to the Component. These will be rendered alongside
	// the component.
	OtherLinks []*Component_Link     `protobuf:"bytes,8,rep,name=other_links" json:"other_links,omitempty"`
	Property   []*Component_Property `protobuf:"bytes,9,rep,name=property" json:"property,omitempty"`
}

func (m *Component) Reset()                    { *m = Component{} }
func (m *Component) String() string            { return proto.CompactTextString(m) }
func (*Component) ProtoMessage()               {}
func (*Component) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Component) GetStarted() *google_protobuf.Timestamp {
	if m != nil {
		return m.Started
	}
	return nil
}

func (m *Component) GetEnded() *google_protobuf.Timestamp {
	if m != nil {
		return m.Ended
	}
	return nil
}

func (m *Component) GetProgress() *Progress {
	if m != nil {
		return m.Progress
	}
	return nil
}

func (m *Component) GetLink() *Component_Link {
	if m != nil {
		return m.Link
	}
	return nil
}

func (m *Component) GetOtherLinks() []*Component_Link {
	if m != nil {
		return m.OtherLinks
	}
	return nil
}

func (m *Component) GetProperty() []*Component_Property {
	if m != nil {
		return m.Property
	}
	return nil
}

// A Link is an optional label followed by a typed link to an external
// resource.
type Component_Link struct {
	// An optional display label for the link.
	Label string `protobuf:"bytes,1,opt,name=label" json:"label,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*Component_Link_Url
	//	*Component_Link_LogdogStream
	//	*Component_Link_IsolateObject
	//	*Component_Link_DmLink
	Value isComponent_Link_Value `protobuf_oneof:"value"`
}

func (m *Component_Link) Reset()                    { *m = Component_Link{} }
func (m *Component_Link) String() string            { return proto.CompactTextString(m) }
func (*Component_Link) ProtoMessage()               {}
func (*Component_Link) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type isComponent_Link_Value interface {
	isComponent_Link_Value()
}

type Component_Link_Url struct {
	Url string `protobuf:"bytes,2,opt,name=url,oneof"`
}
type Component_Link_LogdogStream struct {
	LogdogStream *LogdogStream `protobuf:"bytes,3,opt,name=logdog_stream,oneof"`
}
type Component_Link_IsolateObject struct {
	IsolateObject *IsolateObject `protobuf:"bytes,4,opt,name=isolate_object,oneof"`
}
type Component_Link_DmLink struct {
	DmLink *DMLink `protobuf:"bytes,5,opt,name=dm_link,oneof"`
}

func (*Component_Link_Url) isComponent_Link_Value()           {}
func (*Component_Link_LogdogStream) isComponent_Link_Value()  {}
func (*Component_Link_IsolateObject) isComponent_Link_Value() {}
func (*Component_Link_DmLink) isComponent_Link_Value()        {}

func (m *Component_Link) GetValue() isComponent_Link_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Component_Link) GetUrl() string {
	if x, ok := m.GetValue().(*Component_Link_Url); ok {
		return x.Url
	}
	return ""
}

func (m *Component_Link) GetLogdogStream() *LogdogStream {
	if x, ok := m.GetValue().(*Component_Link_LogdogStream); ok {
		return x.LogdogStream
	}
	return nil
}

func (m *Component_Link) GetIsolateObject() *IsolateObject {
	if x, ok := m.GetValue().(*Component_Link_IsolateObject); ok {
		return x.IsolateObject
	}
	return nil
}

func (m *Component_Link) GetDmLink() *DMLink {
	if x, ok := m.GetValue().(*Component_Link_DmLink); ok {
		return x.DmLink
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Component_Link) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Component_Link_OneofMarshaler, _Component_Link_OneofUnmarshaler, _Component_Link_OneofSizer, []interface{}{
		(*Component_Link_Url)(nil),
		(*Component_Link_LogdogStream)(nil),
		(*Component_Link_IsolateObject)(nil),
		(*Component_Link_DmLink)(nil),
	}
}

func _Component_Link_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Component_Link)
	// value
	switch x := m.Value.(type) {
	case *Component_Link_Url:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Url)
	case *Component_Link_LogdogStream:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LogdogStream); err != nil {
			return err
		}
	case *Component_Link_IsolateObject:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.IsolateObject); err != nil {
			return err
		}
	case *Component_Link_DmLink:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DmLink); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Component_Link.Value has unexpected type %T", x)
	}
	return nil
}

func _Component_Link_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Component_Link)
	switch tag {
	case 2: // value.url
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &Component_Link_Url{x}
		return true, err
	case 3: // value.logdog_stream
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LogdogStream)
		err := b.DecodeMessage(msg)
		m.Value = &Component_Link_LogdogStream{msg}
		return true, err
	case 4: // value.isolate_object
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(IsolateObject)
		err := b.DecodeMessage(msg)
		m.Value = &Component_Link_IsolateObject{msg}
		return true, err
	case 5: // value.dm_link
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DMLink)
		err := b.DecodeMessage(msg)
		m.Value = &Component_Link_DmLink{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Component_Link_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Component_Link)
	// value
	switch x := m.Value.(type) {
	case *Component_Link_Url:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Url)))
		n += len(x.Url)
	case *Component_Link_LogdogStream:
		s := proto.Size(x.LogdogStream)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Component_Link_IsolateObject:
		s := proto.Size(x.IsolateObject)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Component_Link_DmLink:
		s := proto.Size(x.DmLink)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Property is an arbitrary key/value (build) property.
type Component_Property struct {
	// name is the property name.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// value is the optional property value.
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Component_Property) Reset()                    { *m = Component_Property{} }
func (m *Component_Property) String() string            { return proto.CompactTextString(m) }
func (*Component_Property) ProtoMessage()               {}
func (*Component_Property) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 1} }

// Command contains information about a command-line invocation.
type Command struct {
	// The command-line invocation, expressed as an argument vector.
	CommandLine []string `protobuf:"bytes,1,rep,name=command_line" json:"command_line,omitempty"`
	// The current working directory.
	Cwd     string               `protobuf:"bytes,2,opt,name=cwd" json:"cwd,omitempty"`
	Environ *Command_Environment `protobuf:"bytes,3,opt,name=environ" json:"environ,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Command) GetEnviron() *Command_Environment {
	if m != nil {
		return m.Environ
	}
	return nil
}

// Environment represents the state of a process' environment.
type Command_Environment struct {
	// The entries that compose the environment.
	Entries []*Command_Environment_Entry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
}

func (m *Command_Environment) Reset()                    { *m = Command_Environment{} }
func (m *Command_Environment) String() string            { return proto.CompactTextString(m) }
func (*Command_Environment) ProtoMessage()               {}
func (*Command_Environment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

func (m *Command_Environment) GetEntries() []*Command_Environment_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

// Entry is a single key/value environment entry.
type Command_Environment_Entry struct {
	// Name is the name of the environment variable.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Value is the value of the environment variable. This may be empty.
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Command_Environment_Entry) Reset()                    { *m = Command_Environment_Entry{} }
func (m *Command_Environment_Entry) String() string            { return proto.CompactTextString(m) }
func (*Command_Environment_Entry) ProtoMessage()               {}
func (*Command_Environment_Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0, 0} }

// Progress expresses a Component's overall progress. It does this using
// arbitrary "progress units", wich are discrete units of work measured by the
// Component that are either completed or not completed.
//
// A simple construction for "percentage complete" is to set `total` to 100 and
// `completed` to the percentage value.
type Progress struct {
	// The total number of progress units. If missing or zero, no progress is
	// expressed.
	Total int32 `protobuf:"varint,1,opt,name=total" json:"total,omitempty"`
	// The number of completed progress units. This must always be less than or
	// equal to `total`. If omitted, it is implied to be zero.
	Completed int32 `protobuf:"varint,2,opt,name=completed" json:"completed,omitempty"`
}

func (m *Progress) Reset()                    { *m = Progress{} }
func (m *Progress) String() string            { return proto.CompactTextString(m) }
func (*Progress) ProtoMessage()               {}
func (*Progress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// LogdogLink is a LogDog stream link.
type LogdogStream struct {
	// The stream's server. If omitted, the server is the same server that this
	// annotation stream is homed on.
	Server string `protobuf:"bytes,1,opt,name=server" json:"server,omitempty"`
	// The log Prefix. If empty, the prefix is the same prefix as this annotation
	// stream.
	Prefix string `protobuf:"bytes,2,opt,name=prefix" json:"prefix,omitempty"`
	// The log name.
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *LogdogStream) Reset()                    { *m = LogdogStream{} }
func (m *LogdogStream) String() string            { return proto.CompactTextString(m) }
func (*LogdogStream) ProtoMessage()               {}
func (*LogdogStream) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// IsolateObject is an Isolate service object specification.
type IsolateObject struct {
	// The Isolate server. If empty, this is the default Isolate server specified
	// by the project's LUCI config.
	Server string `protobuf:"bytes,1,opt,name=server" json:"server,omitempty"`
	// The isolate object hash.
	Hash string `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
}

func (m *IsolateObject) Reset()                    { *m = IsolateObject{} }
func (m *IsolateObject) String() string            { return proto.CompactTextString(m) }
func (*IsolateObject) ProtoMessage()               {}
func (*IsolateObject) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// Dependency is a Dungeon Master execution specification.
type DMLink struct {
	// The Dungeon Master server. If empty, this is the default Isolate server
	// specified by the project's LUCI config.
	Server string `protobuf:"bytes,1,opt,name=server" json:"server,omitempty"`
	// The quest name.
	Quest string `protobuf:"bytes,2,opt,name=quest" json:"quest,omitempty"`
	// The attempt number.
	Attempt int64 `protobuf:"varint,3,opt,name=attempt" json:"attempt,omitempty"`
	// The execution number.
	Execution int64 `protobuf:"varint,4,opt,name=execution" json:"execution,omitempty"`
}

func (m *DMLink) Reset()                    { *m = DMLink{} }
func (m *DMLink) String() string            { return proto.CompactTextString(m) }
func (*DMLink) ProtoMessage()               {}
func (*DMLink) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*FailureDetails)(nil), "milo.FailureDetails")
	proto.RegisterType((*Step)(nil), "milo.Step")
	proto.RegisterType((*Component)(nil), "milo.Component")
	proto.RegisterType((*Component_Link)(nil), "milo.Component.Link")
	proto.RegisterType((*Component_Property)(nil), "milo.Component.Property")
	proto.RegisterType((*Command)(nil), "milo.Command")
	proto.RegisterType((*Command_Environment)(nil), "milo.Command.Environment")
	proto.RegisterType((*Command_Environment_Entry)(nil), "milo.Command.Environment.Entry")
	proto.RegisterType((*Progress)(nil), "milo.Progress")
	proto.RegisterType((*LogdogStream)(nil), "milo.LogdogStream")
	proto.RegisterType((*IsolateObject)(nil), "milo.IsolateObject")
	proto.RegisterType((*DMLink)(nil), "milo.DMLink")
	proto.RegisterEnum("milo.Status", Status_name, Status_value)
	proto.RegisterEnum("milo.FailureDetails_Type", FailureDetails_Type_name, FailureDetails_Type_value)
}

var fileDescriptor0 = []byte{
	// 791 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x54, 0x5f, 0x6f, 0xe3, 0x44,
	0x10, 0x6f, 0xea, 0x38, 0xae, 0xc7, 0x4d, 0x9a, 0x5b, 0xfa, 0xe0, 0x8b, 0x10, 0xad, 0x02, 0xd2,
	0x1d, 0x85, 0xfa, 0x50, 0x78, 0xe1, 0x01, 0x21, 0x95, 0xc4, 0xbd, 0x8b, 0xe8, 0xf9, 0xa2, 0xa4,
	0x95, 0xe0, 0xc9, 0x72, 0xe2, 0x49, 0x6a, 0xf0, 0x3f, 0xec, 0x75, 0x69, 0xbe, 0x05, 0x9f, 0x03,
	0x89, 0xcf, 0xc0, 0x13, 0xdf, 0x8b, 0xd9, 0x5d, 0x27, 0x34, 0xd7, 0xc2, 0xbd, 0xed, 0xee, 0xfc,
	0x66, 0xe6, 0x37, 0xf3, 0x9b, 0x1d, 0x78, 0x16, 0xa4, 0x69, 0xc6, 0x03, 0x1e, 0x65, 0x69, 0xe9,
	0xe4, 0x45, 0xc6, 0x33, 0xd6, 0x4c, 0xa2, 0x38, 0xeb, 0x9d, 0xac, 0xb2, 0x6c, 0x15, 0xe3, 0x2b,
	0xf9, 0x36, 0xaf, 0x96, 0xaf, 0x78, 0x94, 0x60, 0xc9, 0x83, 0x24, 0x57, 0xb0, 0xfe, 0x9f, 0x0d,
	0xe8, 0x5c, 0x06, 0x51, 0x5c, 0x15, 0x38, 0x42, 0x4e, 0x87, 0x92, 0xbd, 0x80, 0x26, 0x5f, 0xe7,
	0x68, 0x37, 0x4e, 0x1b, 0x2f, 0x3b, 0x83, 0xe7, 0x8e, 0x08, 0xe4, 0xec, 0x62, 0x9c, 0x6b, 0x02,
	0xb0, 0x43, 0x02, 0xe2, 0x3d, 0xb7, 0xf7, 0x09, 0x68, 0xb2, 0x33, 0x38, 0x5e, 0x92, 0x0d, 0x43,
	0x3f, 0x4c, 0xfc, 0x10, 0x73, 0x4c, 0x43, 0x4c, 0x17, 0x6b, 0x5b, 0x3b, 0xd5, 0x5e, 0x5a, 0x83,
	0x43, 0x15, 0x66, 0xf4, 0xf6, 0x2a, 0x4a, 0x7f, 0xe9, 0x7f, 0x03, 0x4d, 0x19, 0xc1, 0x02, 0xe3,
	0xb5, 0xeb, 0xb9, 0xd3, 0x8b, 0xab, 0xee, 0x1e, 0x33, 0x41, 0x1f, 0x7b, 0x97, 0xd3, 0x8b, 0x6e,
	0x83, 0xd9, 0x70, 0x3c, 0x7a, 0xeb, 0x8f, 0xdc, 0x89, 0xeb, 0x8d, 0x5c, 0x6f, 0xf8, 0x93, 0x7f,
	0x79, 0x31, 0xbe, 0x72, 0x47, 0xdd, 0xfd, 0xfe, 0xdf, 0x0d, 0x68, 0xce, 0x38, 0xe6, 0xec, 0x13,
	0x30, 0x16, 0x59, 0x92, 0x04, 0x69, 0x28, 0x89, 0x5a, 0x83, 0xb6, 0xca, 0x30, 0x54, 0x8f, 0xec,
	0x1c, 0x8e, 0x96, 0x8a, 0x33, 0x91, 0x91, 0xa4, 0x25, 0x4f, 0x6b, 0x70, 0xfc, 0x54, 0x41, 0x54,
	0x74, 0xa7, 0xa4, 0xb0, 0x3e, 0xc5, 0xcc, 0xb3, 0x14, 0x53, 0x4e, 0xbc, 0x05, 0xfa, 0x68, 0x1b,
	0x55, 0x3d, 0xb3, 0x4f, 0x01, 0xb6, 0x98, 0xd2, 0x6e, 0xca, 0xe2, 0x1e, 0x81, 0x4e, 0xc1, 0x2e,
	0xab, 0xb9, 0x0c, 0x18, 0x67, 0xab, 0x30, 0x5b, 0xf9, 0x69, 0x90, 0xa0, 0x3f, 0x0f, 0x4a, 0xb4,
	0x75, 0x72, 0x31, 0xfb, 0xbf, 0x37, 0xc1, 0xfc, 0x17, 0x4f, 0x9d, 0x14, 0x00, 0x59, 0x89, 0xc9,
	0x3e, 0x86, 0x16, 0x49, 0xc4, 0x2b, 0xc5, 0xb8, 0xb3, 0xe9, 0xdd, 0x4c, 0xbe, 0xb1, 0x2f, 0xc0,
	0x20, 0x6b, 0xc1, 0x31, 0xac, 0x29, 0xf6, 0x1c, 0x25, 0xb2, 0xb3, 0x11, 0xd9, 0xb9, 0xde, 0x88,
	0xcc, 0x3e, 0x07, 0x5d, 0xe8, 0x10, 0x12, 0xd1, 0x0f, 0x41, 0x37, 0x6a, 0x4a, 0x7e, 0x54, 0xc1,
	0x01, 0x61, 0x56, 0x05, 0x96, 0xa5, 0xdd, 0x92, 0xbe, 0x1d, 0xc5, 0x62, 0x52, 0xbf, 0xb2, 0x3e,
	0x34, 0x63, 0xd2, 0xd2, 0x36, 0x1e, 0x76, 0x75, 0x5b, 0x92, 0x23, 0x74, 0xa6, 0xf4, 0x56, 0xc6,
	0x6f, 0xb1, 0xf0, 0x05, 0xb2, 0xb4, 0x0f, 0x64, 0xb7, 0x9e, 0x86, 0x9e, 0xc9, 0x84, 0x39, 0x16,
	0x7c, 0x6d, 0x9b, 0x12, 0x67, 0xbf, 0x8f, 0x9b, 0xd4, 0xf6, 0xde, 0x1f, 0x34, 0x04, 0xd2, 0xa9,
	0x0d, 0x7a, 0x1c, 0xcc, 0x31, 0xae, 0x1b, 0xd7, 0x06, 0xad, 0x2a, 0x62, 0x35, 0x8f, 0x6f, 0xf6,
	0xa8, 0x53, 0xed, 0xba, 0xfb, 0x25, 0x2f, 0x30, 0x48, 0xea, 0x7e, 0x31, 0x15, 0xf7, 0x4a, 0x9a,
	0x66, 0xd2, 0x42, 0xe0, 0x73, 0xe8, 0x44, 0x65, 0x16, 0x07, 0x1c, 0xfd, 0x6c, 0xfe, 0x33, 0x2e,
	0x78, 0xdd, 0xb2, 0x8f, 0x14, 0x7a, 0xac, 0x6c, 0xef, 0xa4, 0x89, 0xe0, 0x27, 0x60, 0xd0, 0x98,
	0xcb, 0x06, 0xe8, 0x12, 0xb7, 0x33, 0xe0, 0x6f, 0xf6, 0xbe, 0x37, 0x40, 0xbf, 0x0b, 0xe2, 0x0a,
	0x7b, 0x2f, 0xe0, 0x60, 0x43, 0xfc, 0x3d, 0x9d, 0xdb, 0x35, 0x44, 0x11, 0xee, 0xff, 0xd5, 0x00,
	0x63, 0x33, 0xbd, 0xc7, 0x70, 0x58, 0x4f, 0xb7, 0xc8, 0x21, 0x1c, 0x84, 0x28, 0x16, 0x68, 0x8b,
	0xdf, 0xc2, 0xed, 0x7f, 0x33, 0x30, 0xbd, 0x8b, 0x8a, 0x2c, 0xad, 0xeb, 0x7a, 0xbe, 0xf3, 0x01,
	0x1c, 0x57, 0x19, 0x13, 0xea, 0x5c, 0x0f, 0xc1, 0x7a, 0x70, 0x65, 0x5f, 0x09, 0x57, 0x5e, 0x44,
	0x58, 0xca, 0xc0, 0xd6, 0xe0, 0xe4, 0x3f, 0x5d, 0xe9, 0xcc, 0x8b, 0x75, 0xef, 0x33, 0xd0, 0xe5,
	0xe1, 0xff, 0x2b, 0xf8, 0x52, 0x96, 0xaa, 0xc6, 0x83, 0x4c, 0x9c, 0x56, 0x92, 0x92, 0x46, 0x67,
	0xcf, 0xc0, 0x14, 0xdf, 0x26, 0x46, 0x31, 0xb7, 0x02, 0xad, 0xf7, 0xbf, 0x85, 0xc3, 0x87, 0x1a,
	0xb0, 0x0e, 0x8d, 0x3d, 0x16, 0x77, 0x58, 0xd4, 0xc1, 0xe9, 0x9e, 0x17, 0xb8, 0x8c, 0xee, 0xeb,
	0x82, 0x37, 0xa9, 0x35, 0x99, 0xeb, 0x1c, 0xda, 0x3b, 0x9a, 0x3c, 0x72, 0x27, 0xf8, 0x6d, 0x50,
	0xde, 0xd6, 0xd4, 0x7e, 0x80, 0x96, 0x92, 0xe6, 0x11, 0x8e, 0x88, 0xfe, 0x5a, 0xd1, 0x1f, 0xa8,
	0xb3, 0x1c, 0x81, 0x11, 0x70, 0x8e, 0x49, 0xae, 0x36, 0x80, 0x26, 0x98, 0xe3, 0x3d, 0x2e, 0x2a,
	0xb1, 0x5c, 0xe5, 0x4c, 0x68, 0x67, 0xdf, 0x41, 0xab, 0xfe, 0x8c, 0xb4, 0xc0, 0xa6, 0x37, 0x9e,
	0x37, 0xf6, 0x5e, 0xd3, 0x02, 0xa3, 0xcb, 0xec, 0x66, 0x38, 0x74, 0x67, 0x33, 0x5a, 0x61, 0x74,
	0x11, 0x4b, 0xeb, 0x66, 0xea, 0x76, 0xf7, 0x29, 0x87, 0xe9, 0xfe, 0x38, 0x74, 0x27, 0xd7, 0xe3,
	0x77, 0x5e, 0x57, 0x9b, 0xb7, 0xe4, 0xf7, 0xfb, 0xfa, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x55,
	0x8a, 0x21, 0x70, 0xb7, 0x05, 0x00, 0x00,
}
