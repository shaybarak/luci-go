// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/luci/luci-go/common/proto/milo/annotations.proto

/*
Package milo is a generated protocol buffer package.

It is generated from these files:
	github.com/luci/luci-go/common/proto/milo/annotations.proto

It has these top-level messages:
	FailureDetails
	Step
	Link
	LogdogStream
	IsolateObject
	DMLink
*/
package milo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Status is the expressed root step of this step or substep.
type Status int32

const (
	// The step is still running.
	Status_RUNNING Status = 0
	// The step has finished successfully.
	Status_SUCCESS Status = 1
	// The step has finished unsuccessfully.
	Status_FAILURE Status = 2
	// The step has been scheduled, but not yet started.
	Status_PENDING Status = 3
)

var Status_name = map[int32]string{
	0: "RUNNING",
	1: "SUCCESS",
	2: "FAILURE",
	3: "PENDING",
}
var Status_value = map[string]int32{
	"RUNNING": 0,
	"SUCCESS": 1,
	"FAILURE": 2,
	"PENDING": 3,
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
	// An unhandled exception occured during execution.
	FailureDetails_EXCEPTION FailureDetails_Type = 1
	// The failure is related to a failed infrastructure component, not an error
	// with the Step itself.
	FailureDetails_INFRA FailureDetails_Type = 2
	// The failure is due to a failed Dungeon Master dependency. This should be
	// used if a Step's external depdendency fails and the Step cannot recover
	// or proceed without it.
	FailureDetails_DM_DEPENDENCY_FAILED FailureDetails_Type = 3
	// The step was cancelled.
	FailureDetails_CANCELLED FailureDetails_Type = 4
	// The failure was due to an resource exhausion. The step was scheduled
	// but never ran, and never will run.
	FailureDetails_EXPIRED FailureDetails_Type = 5
)

var FailureDetails_Type_name = map[int32]string{
	0: "GENERAL",
	1: "EXCEPTION",
	2: "INFRA",
	3: "DM_DEPENDENCY_FAILED",
	4: "CANCELLED",
	5: "EXPIRED",
}
var FailureDetails_Type_value = map[string]int32{
	"GENERAL":              0,
	"EXCEPTION":            1,
	"INFRA":                2,
	"DM_DEPENDENCY_FAILED": 3,
	"CANCELLED":            4,
	"EXPIRED":              5,
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
	FailedDmDependency []*DMLink `protobuf:"bytes,3,rep,name=failed_dm_dependency,json=failedDmDependency" json:"failed_dm_dependency,omitempty"`
}

func (m *FailureDetails) Reset()                    { *m = FailureDetails{} }
func (m *FailureDetails) String() string            { return proto.CompactTextString(m) }
func (*FailureDetails) ProtoMessage()               {}
func (*FailureDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FailureDetails) GetType() FailureDetails_Type {
	if m != nil {
		return m.Type
	}
	return FailureDetails_GENERAL
}

func (m *FailureDetails) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *FailureDetails) GetFailedDmDependency() []*DMLink {
	if m != nil {
		return m.FailedDmDependency
	}
	return nil
}

// Generic step or substep state.
type Step struct {
	// The display name of the Component.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The command-line invocation of the step, expressed as an argument vector.
	Command *Step_Command `protobuf:"bytes,2,opt,name=command" json:"command,omitempty"`
	// The current running status of the Step.
	Status Status `protobuf:"varint,3,opt,name=status,enum=milo.Status" json:"status,omitempty"`
	// Optional information detailing the failure. This may be populated if the
	// Step's top-level command Status is set to FAILURE.
	FailureDetails *FailureDetails `protobuf:"bytes,4,opt,name=failure_details,json=failureDetails" json:"failure_details,omitempty"`
	// Substeps that this Step is composed of.
	Substep []*Step_Substep `protobuf:"bytes,5,rep,name=substep" json:"substep,omitempty"`
	// A link to this Step's STDOUT stream, if present.
	StdoutStream *LogdogStream `protobuf:"bytes,6,opt,name=stdout_stream,json=stdoutStream" json:"stdout_stream,omitempty"`
	// A link to this Step's STDERR stream, if present.
	StderrStream *LogdogStream `protobuf:"bytes,7,opt,name=stderr_stream,json=stderrStream" json:"stderr_stream,omitempty"`
	// When the step started, expressed as an RFC3339 string using Z (UTC)
	// timezone.
	Started *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=started" json:"started,omitempty"`
	// When the step ended, expressed as an RFC3339 string using Z (UTC) timezone.
	Ended *google_protobuf.Timestamp `protobuf:"bytes,9,opt,name=ended" json:"ended,omitempty"`
	// Arbitrary lines of component text. Each string here is a consecutive line,
	// and should not contain newlines.
	Text []string `protobuf:"bytes,20,rep,name=text" json:"text,omitempty"`
	// The Component's progress.
	Progress *Step_Progress `protobuf:"bytes,21,opt,name=progress" json:"progress,omitempty"`
	// The primary link for this Component. This is the link that interaction
	// with the Component will use.
	Link *Link `protobuf:"bytes,22,opt,name=link" json:"link,omitempty"`
	// Additional links related to the Component. These will be rendered alongside
	// the component.
	OtherLinks []*Link          `protobuf:"bytes,23,rep,name=other_links,json=otherLinks" json:"other_links,omitempty"`
	Property   []*Step_Property `protobuf:"bytes,24,rep,name=property" json:"property,omitempty"`
}

func (m *Step) Reset()                    { *m = Step{} }
func (m *Step) String() string            { return proto.CompactTextString(m) }
func (*Step) ProtoMessage()               {}
func (*Step) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Step) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Step) GetCommand() *Step_Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *Step) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_RUNNING
}

func (m *Step) GetFailureDetails() *FailureDetails {
	if m != nil {
		return m.FailureDetails
	}
	return nil
}

func (m *Step) GetSubstep() []*Step_Substep {
	if m != nil {
		return m.Substep
	}
	return nil
}

func (m *Step) GetStdoutStream() *LogdogStream {
	if m != nil {
		return m.StdoutStream
	}
	return nil
}

func (m *Step) GetStderrStream() *LogdogStream {
	if m != nil {
		return m.StderrStream
	}
	return nil
}

func (m *Step) GetStarted() *google_protobuf.Timestamp {
	if m != nil {
		return m.Started
	}
	return nil
}

func (m *Step) GetEnded() *google_protobuf.Timestamp {
	if m != nil {
		return m.Ended
	}
	return nil
}

func (m *Step) GetText() []string {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *Step) GetProgress() *Step_Progress {
	if m != nil {
		return m.Progress
	}
	return nil
}

func (m *Step) GetLink() *Link {
	if m != nil {
		return m.Link
	}
	return nil
}

func (m *Step) GetOtherLinks() []*Link {
	if m != nil {
		return m.OtherLinks
	}
	return nil
}

func (m *Step) GetProperty() []*Step_Property {
	if m != nil {
		return m.Property
	}
	return nil
}

// Command contains information about a command-line invocation.
type Step_Command struct {
	// The command-line invocation, expressed as an argument vector.
	CommandLine []string `protobuf:"bytes,1,rep,name=command_line,json=commandLine" json:"command_line,omitempty"`
	// The current working directory.
	Cwd string `protobuf:"bytes,2,opt,name=cwd" json:"cwd,omitempty"`
	// Environment represents the state of a process' environment.
	Environ map[string]string `protobuf:"bytes,3,rep,name=environ" json:"environ,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Step_Command) Reset()                    { *m = Step_Command{} }
func (m *Step_Command) String() string            { return proto.CompactTextString(m) }
func (*Step_Command) ProtoMessage()               {}
func (*Step_Command) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Step_Command) GetCommandLine() []string {
	if m != nil {
		return m.CommandLine
	}
	return nil
}

func (m *Step_Command) GetCwd() string {
	if m != nil {
		return m.Cwd
	}
	return ""
}

func (m *Step_Command) GetEnviron() map[string]string {
	if m != nil {
		return m.Environ
	}
	return nil
}

// Sub-steps nested underneath of this step.
type Step_Substep struct {
	// The substep.
	//
	// Types that are valid to be assigned to Substep:
	//	*Step_Substep_Step
	//	*Step_Substep_AnnotationStream
	Substep isStep_Substep_Substep `protobuf_oneof:"substep"`
}

func (m *Step_Substep) Reset()                    { *m = Step_Substep{} }
func (m *Step_Substep) String() string            { return proto.CompactTextString(m) }
func (*Step_Substep) ProtoMessage()               {}
func (*Step_Substep) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

type isStep_Substep_Substep interface {
	isStep_Substep_Substep()
}

type Step_Substep_Step struct {
	Step *Step `protobuf:"bytes,1,opt,name=step,oneof"`
}
type Step_Substep_AnnotationStream struct {
	AnnotationStream *LogdogStream `protobuf:"bytes,2,opt,name=annotation_stream,json=annotationStream,oneof"`
}

func (*Step_Substep_Step) isStep_Substep_Substep()             {}
func (*Step_Substep_AnnotationStream) isStep_Substep_Substep() {}

func (m *Step_Substep) GetSubstep() isStep_Substep_Substep {
	if m != nil {
		return m.Substep
	}
	return nil
}

func (m *Step_Substep) GetStep() *Step {
	if x, ok := m.GetSubstep().(*Step_Substep_Step); ok {
		return x.Step
	}
	return nil
}

func (m *Step_Substep) GetAnnotationStream() *LogdogStream {
	if x, ok := m.GetSubstep().(*Step_Substep_AnnotationStream); ok {
		return x.AnnotationStream
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Step_Substep) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Step_Substep_OneofMarshaler, _Step_Substep_OneofUnmarshaler, _Step_Substep_OneofSizer, []interface{}{
		(*Step_Substep_Step)(nil),
		(*Step_Substep_AnnotationStream)(nil),
	}
}

func _Step_Substep_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Step_Substep)
	// substep
	switch x := m.Substep.(type) {
	case *Step_Substep_Step:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Step); err != nil {
			return err
		}
	case *Step_Substep_AnnotationStream:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AnnotationStream); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Step_Substep.Substep has unexpected type %T", x)
	}
	return nil
}

func _Step_Substep_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Step_Substep)
	switch tag {
	case 1: // substep.step
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Step)
		err := b.DecodeMessage(msg)
		m.Substep = &Step_Substep_Step{msg}
		return true, err
	case 2: // substep.annotation_stream
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LogdogStream)
		err := b.DecodeMessage(msg)
		m.Substep = &Step_Substep_AnnotationStream{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Step_Substep_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Step_Substep)
	// substep
	switch x := m.Substep.(type) {
	case *Step_Substep_Step:
		s := proto.Size(x.Step)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Step_Substep_AnnotationStream:
		s := proto.Size(x.AnnotationStream)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Progress expresses a Component's overall progress. It does this using
// arbitrary "progress units", wich are discrete units of work measured by
// the Component that are either completed or not completed.
//
// A simple construction for "percentage complete" is to set `total` to 100
// and `completed` to the percentage value.
type Step_Progress struct {
	// The total number of progress units. If missing or zero, no progress is
	// expressed.
	Total int32 `protobuf:"varint,1,opt,name=total" json:"total,omitempty"`
	// The number of completed progress units. This must always be less than or
	// equal to `total`. If omitted, it is implied to be zero.
	Completed int32 `protobuf:"varint,2,opt,name=completed" json:"completed,omitempty"`
}

func (m *Step_Progress) Reset()                    { *m = Step_Progress{} }
func (m *Step_Progress) String() string            { return proto.CompactTextString(m) }
func (*Step_Progress) ProtoMessage()               {}
func (*Step_Progress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 2} }

func (m *Step_Progress) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Step_Progress) GetCompleted() int32 {
	if m != nil {
		return m.Completed
	}
	return 0
}

// Property is an arbitrary key/value (build) property.
type Step_Property struct {
	// name is the property name.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// value is the optional property value.
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Step_Property) Reset()                    { *m = Step_Property{} }
func (m *Step_Property) String() string            { return proto.CompactTextString(m) }
func (*Step_Property) ProtoMessage()               {}
func (*Step_Property) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 3} }

func (m *Step_Property) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Step_Property) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// A Link is an optional label followed by a typed link to an external
// resource.
type Link struct {
	// An optional display label for the link.
	Label string `protobuf:"bytes,1,opt,name=label" json:"label,omitempty"`
	// If present, this link is an alias for another link with this name, and
	// should be rendered in relation to that link.
	AliasLabel string `protobuf:"bytes,2,opt,name=alias_label,json=aliasLabel" json:"alias_label,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*Link_Url
	//	*Link_LogdogStream
	//	*Link_IsolateObject
	//	*Link_DmLink
	Value isLink_Value `protobuf_oneof:"value"`
}

func (m *Link) Reset()                    { *m = Link{} }
func (m *Link) String() string            { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()               {}
func (*Link) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isLink_Value interface {
	isLink_Value()
}

type Link_Url struct {
	Url string `protobuf:"bytes,3,opt,name=url,oneof"`
}
type Link_LogdogStream struct {
	LogdogStream *LogdogStream `protobuf:"bytes,4,opt,name=logdog_stream,json=logdogStream,oneof"`
}
type Link_IsolateObject struct {
	IsolateObject *IsolateObject `protobuf:"bytes,5,opt,name=isolate_object,json=isolateObject,oneof"`
}
type Link_DmLink struct {
	DmLink *DMLink `protobuf:"bytes,6,opt,name=dm_link,json=dmLink,oneof"`
}

func (*Link_Url) isLink_Value()           {}
func (*Link_LogdogStream) isLink_Value()  {}
func (*Link_IsolateObject) isLink_Value() {}
func (*Link_DmLink) isLink_Value()        {}

func (m *Link) GetValue() isLink_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Link) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Link) GetAliasLabel() string {
	if m != nil {
		return m.AliasLabel
	}
	return ""
}

func (m *Link) GetUrl() string {
	if x, ok := m.GetValue().(*Link_Url); ok {
		return x.Url
	}
	return ""
}

func (m *Link) GetLogdogStream() *LogdogStream {
	if x, ok := m.GetValue().(*Link_LogdogStream); ok {
		return x.LogdogStream
	}
	return nil
}

func (m *Link) GetIsolateObject() *IsolateObject {
	if x, ok := m.GetValue().(*Link_IsolateObject); ok {
		return x.IsolateObject
	}
	return nil
}

func (m *Link) GetDmLink() *DMLink {
	if x, ok := m.GetValue().(*Link_DmLink); ok {
		return x.DmLink
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Link) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Link_OneofMarshaler, _Link_OneofUnmarshaler, _Link_OneofSizer, []interface{}{
		(*Link_Url)(nil),
		(*Link_LogdogStream)(nil),
		(*Link_IsolateObject)(nil),
		(*Link_DmLink)(nil),
	}
}

func _Link_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Link)
	// value
	switch x := m.Value.(type) {
	case *Link_Url:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Url)
	case *Link_LogdogStream:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LogdogStream); err != nil {
			return err
		}
	case *Link_IsolateObject:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.IsolateObject); err != nil {
			return err
		}
	case *Link_DmLink:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DmLink); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Link.Value has unexpected type %T", x)
	}
	return nil
}

func _Link_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Link)
	switch tag {
	case 3: // value.url
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &Link_Url{x}
		return true, err
	case 4: // value.logdog_stream
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LogdogStream)
		err := b.DecodeMessage(msg)
		m.Value = &Link_LogdogStream{msg}
		return true, err
	case 5: // value.isolate_object
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(IsolateObject)
		err := b.DecodeMessage(msg)
		m.Value = &Link_IsolateObject{msg}
		return true, err
	case 6: // value.dm_link
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DMLink)
		err := b.DecodeMessage(msg)
		m.Value = &Link_DmLink{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Link_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Link)
	// value
	switch x := m.Value.(type) {
	case *Link_Url:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Url)))
		n += len(x.Url)
	case *Link_LogdogStream:
		s := proto.Size(x.LogdogStream)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Link_IsolateObject:
		s := proto.Size(x.IsolateObject)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Link_DmLink:
		s := proto.Size(x.DmLink)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// LogdogStream is a LogDog stream link.
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
func (*LogdogStream) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LogdogStream) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *LogdogStream) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *LogdogStream) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

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
func (*IsolateObject) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *IsolateObject) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *IsolateObject) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

// DMLink is a Dungeon Master execution specification.
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
func (*DMLink) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DMLink) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *DMLink) GetQuest() string {
	if m != nil {
		return m.Quest
	}
	return ""
}

func (m *DMLink) GetAttempt() int64 {
	if m != nil {
		return m.Attempt
	}
	return 0
}

func (m *DMLink) GetExecution() int64 {
	if m != nil {
		return m.Execution
	}
	return 0
}

func init() {
	proto.RegisterType((*FailureDetails)(nil), "milo.FailureDetails")
	proto.RegisterType((*Step)(nil), "milo.Step")
	proto.RegisterType((*Step_Command)(nil), "milo.Step.Command")
	proto.RegisterType((*Step_Substep)(nil), "milo.Step.Substep")
	proto.RegisterType((*Step_Progress)(nil), "milo.Step.Progress")
	proto.RegisterType((*Step_Property)(nil), "milo.Step.Property")
	proto.RegisterType((*Link)(nil), "milo.Link")
	proto.RegisterType((*LogdogStream)(nil), "milo.LogdogStream")
	proto.RegisterType((*IsolateObject)(nil), "milo.IsolateObject")
	proto.RegisterType((*DMLink)(nil), "milo.DMLink")
	proto.RegisterEnum("milo.Status", Status_name, Status_value)
	proto.RegisterEnum("milo.FailureDetails_Type", FailureDetails_Type_name, FailureDetails_Type_value)
}

func init() {
	proto.RegisterFile("github.com/luci/luci-go/common/proto/milo/annotations.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 981 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0x5d, 0x6f, 0xe3, 0x44,
	0x14, 0x8d, 0x63, 0x27, 0x6e, 0x6e, 0xd2, 0x62, 0x86, 0xb0, 0x98, 0x08, 0xd1, 0x10, 0x21, 0x51,
	0xc1, 0xae, 0x83, 0xca, 0x4a, 0xb0, 0x1f, 0xac, 0x94, 0x4d, 0xdc, 0x4d, 0xa4, 0x6c, 0xb6, 0x9a,
	0xb4, 0xd2, 0xf2, 0x14, 0x39, 0xf1, 0x24, 0x35, 0xb5, 0x3d, 0xc6, 0x1e, 0x97, 0xe6, 0xe7, 0xf0,
	0x23, 0xf8, 0x6f, 0x3c, 0xf0, 0x80, 0xe6, 0xc3, 0x89, 0x0b, 0xed, 0xf2, 0x62, 0xcd, 0xb9, 0xf7,
	0xdc, 0x99, 0xb9, 0x67, 0x8e, 0x67, 0xe0, 0xc5, 0x26, 0x60, 0x57, 0xf9, 0xd2, 0x59, 0xd1, 0xa8,
	0x1f, 0xe6, 0xab, 0x40, 0x7c, 0x9e, 0x6c, 0x68, 0x7f, 0x45, 0xa3, 0x88, 0xc6, 0xfd, 0x24, 0xa5,
	0x8c, 0xf6, 0xa3, 0x20, 0xa4, 0x7d, 0x2f, 0x8e, 0x29, 0xf3, 0x58, 0x40, 0xe3, 0xcc, 0x11, 0x61,
	0x64, 0xf0, 0x78, 0xe7, 0x78, 0x43, 0xe9, 0x26, 0x24, 0x92, 0xba, 0xcc, 0xd7, 0x7d, 0x16, 0x44,
	0x24, 0x63, 0x5e, 0x94, 0x48, 0x5a, 0xef, 0x2f, 0x0d, 0x8e, 0xce, 0xbc, 0x20, 0xcc, 0x53, 0x32,
	0x22, 0xcc, 0x0b, 0xc2, 0x0c, 0x3d, 0x01, 0x83, 0x6d, 0x13, 0x62, 0x6b, 0x5d, 0xed, 0xe4, 0xe8,
	0xf4, 0x73, 0x87, 0x4f, 0xe4, 0xdc, 0xe5, 0x38, 0x17, 0xdb, 0x84, 0x60, 0x41, 0x43, 0x08, 0x0c,
	0x46, 0x6e, 0x99, 0x5d, 0xed, 0x6a, 0x27, 0x0d, 0x2c, 0xc6, 0xe8, 0x15, 0xb4, 0xd7, 0x5e, 0x10,
	0x12, 0x7f, 0xe1, 0x47, 0x0b, 0x9f, 0x24, 0x24, 0xf6, 0x49, 0xbc, 0xda, 0xda, 0x7a, 0x57, 0x3f,
	0x69, 0x9e, 0xb6, 0xe4, 0x94, 0xa3, 0xb7, 0xd3, 0x20, 0xbe, 0xc6, 0x48, 0x32, 0x47, 0xd1, 0x68,
	0xc7, 0xeb, 0xad, 0xc0, 0xe0, 0x2b, 0xa0, 0x26, 0x98, 0x6f, 0xdc, 0x99, 0x8b, 0x07, 0x53, 0xab,
	0x82, 0x0e, 0xa1, 0xe1, 0xbe, 0x1f, 0xba, 0xe7, 0x17, 0x93, 0x77, 0x33, 0x4b, 0x43, 0x0d, 0xa8,
	0x4d, 0x66, 0x67, 0x78, 0x60, 0x55, 0x91, 0x0d, 0xed, 0xd1, 0xdb, 0xc5, 0xc8, 0x3d, 0x77, 0x67,
	0x23, 0x77, 0x36, 0xfc, 0x65, 0x71, 0x36, 0x98, 0x4c, 0xdd, 0x91, 0xa5, 0xf3, 0x9a, 0xe1, 0x60,
	0x36, 0x74, 0xa7, 0x1c, 0x1a, 0x7c, 0x3e, 0xf7, 0xfd, 0xf9, 0x04, 0xbb, 0x23, 0xab, 0xd6, 0xfb,
	0xe3, 0x00, 0x8c, 0x39, 0x23, 0x09, 0xef, 0x20, 0xf6, 0x22, 0xd9, 0x70, 0x03, 0x8b, 0x31, 0x7a,
	0x0c, 0x26, 0x57, 0xd9, 0x8b, 0x7d, 0xd1, 0x58, 0xf3, 0x14, 0xc9, 0x4d, 0xf3, 0x02, 0x67, 0x28,
	0x33, 0xb8, 0xa0, 0xa0, 0xaf, 0xa1, 0x9e, 0x31, 0x8f, 0xe5, 0x99, 0xad, 0x0b, 0xd1, 0x5a, 0x05,
	0x99, 0xc7, 0xb0, 0xca, 0xa1, 0x9f, 0xe1, 0xa3, 0xb5, 0x94, 0x71, 0xe1, 0x4b, 0x1d, 0x6d, 0x43,
	0xcc, 0xdd, 0xbe, 0x4f, 0x63, 0x7c, 0xb4, 0xbe, 0x7b, 0x2e, 0x8f, 0xc1, 0xcc, 0xf2, 0x65, 0xc6,
	0x48, 0x62, 0xd7, 0x84, 0x8e, 0xe5, 0x2d, 0xcd, 0x65, 0x06, 0x17, 0x14, 0xf4, 0x23, 0x1c, 0x66,
	0xcc, 0xa7, 0x39, 0x5b, 0x64, 0x2c, 0x25, 0x5e, 0x64, 0xd7, 0xcb, 0x6d, 0x4c, 0xe9, 0xc6, 0xa7,
	0x9b, 0xb9, 0xc8, 0xe0, 0x96, 0x24, 0x4a, 0xa4, 0x0a, 0x49, 0x9a, 0x16, 0x85, 0xe6, 0x07, 0x0b,
	0x49, 0x9a, 0xaa, 0xc2, 0xa7, 0x60, 0x66, 0xcc, 0x4b, 0x19, 0xf1, 0xed, 0x03, 0x51, 0xd2, 0x71,
	0xa4, 0xfb, 0x9c, 0xc2, 0x7d, 0xce, 0x45, 0xe1, 0x3e, 0x5c, 0x50, 0xd1, 0xf7, 0x50, 0xe3, 0xa7,
	0xee, 0xdb, 0x8d, 0xff, 0xad, 0x91, 0xc4, 0x9d, 0xe1, 0xda, 0x5d, 0x7d, 0x67, 0xb8, 0x3e, 0x1c,
	0x24, 0x29, 0xdd, 0xa4, 0x24, 0xcb, 0xec, 0x4f, 0xc5, 0x44, 0x9f, 0x94, 0xc4, 0x39, 0x57, 0x29,
	0xbc, 0x23, 0xa1, 0x2f, 0xc1, 0x08, 0x83, 0xf8, 0xda, 0x7e, 0x24, 0xc8, 0xa0, 0x9a, 0xe3, 0x7e,
	0x14, 0x71, 0xf4, 0x1d, 0x34, 0x29, 0xbb, 0x22, 0xe9, 0x82, 0xa3, 0xcc, 0xfe, 0x4c, 0x08, 0x5e,
	0xa6, 0x81, 0x48, 0xf3, 0x61, 0xa6, 0x56, 0x4f, 0x48, 0xca, 0xb6, 0xb6, 0x2d, 0x98, 0xff, 0x5a,
	0x5d, 0xa4, 0xf0, 0x8e, 0xd4, 0xf9, 0x53, 0x03, 0x53, 0x99, 0x08, 0x7d, 0x05, 0x2d, 0x65, 0x23,
	0xbe, 0x16, 0x77, 0x21, 0x6f, 0xab, 0xa9, 0x62, 0xd3, 0x20, 0x26, 0xc8, 0x02, 0x7d, 0xf5, 0xbb,
	0xaf, 0xfe, 0x30, 0x3e, 0x44, 0xcf, 0xc0, 0x24, 0xf1, 0x4d, 0x90, 0xd2, 0x58, 0xfd, 0x53, 0xc7,
	0xff, 0xb5, 0xa7, 0xe3, 0x4a, 0x86, 0x1b, 0xb3, 0x74, 0x8b, 0x0b, 0x7e, 0xe7, 0x39, 0xb4, 0xca,
	0x09, 0x3e, 0xf9, 0x35, 0xd9, 0x2a, 0xf3, 0xf3, 0x21, 0x6a, 0x43, 0xed, 0xc6, 0x0b, 0x73, 0xa2,
	0x16, 0x94, 0xe0, 0x79, 0xf5, 0x27, 0xad, 0xb3, 0x05, 0x53, 0x19, 0x0d, 0x75, 0xc1, 0x10, 0x56,
	0xd4, 0xca, 0x02, 0xf2, 0xe5, 0xc7, 0x15, 0x2c, 0x32, 0x68, 0x00, 0x1f, 0xef, 0xaf, 0xa5, 0xc2,
	0x4c, 0xd5, 0x87, 0xcc, 0x34, 0xae, 0x60, 0x6b, 0x4f, 0x97, 0xb1, 0xd7, 0x8d, 0x9d, 0xe5, 0x3b,
	0xaf, 0xe0, 0xa0, 0x38, 0x46, 0xbe, 0x41, 0x46, 0x99, 0x17, 0x8a, 0xc5, 0x6b, 0x58, 0x02, 0xf4,
	0x05, 0x34, 0x56, 0x34, 0x4a, 0x42, 0xc2, 0x1d, 0x58, 0x15, 0x99, 0x7d, 0xa0, 0xf3, 0x54, 0xd4,
	0x0b, 0xf9, 0xef, 0xfd, 0xe1, 0xef, 0x6d, 0xba, 0xf7, 0xb7, 0x06, 0x06, 0x3f, 0x63, 0x9e, 0x0e,
	0xbd, 0x25, 0x09, 0x55, 0x8d, 0x04, 0xe8, 0x18, 0x9a, 0x5e, 0x18, 0x78, 0xd9, 0x42, 0xe6, 0x64,
	0x29, 0x88, 0xd0, 0x54, 0x10, 0x10, 0xe8, 0x79, 0x1a, 0x8a, 0x5b, 0xa1, 0x31, 0xae, 0x60, 0x0e,
	0xd0, 0x33, 0x38, 0x0c, 0x45, 0xe3, 0x85, 0x26, 0xc6, 0x07, 0x34, 0x69, 0x85, 0x25, 0x8c, 0x5e,
	0xc2, 0x51, 0x90, 0xd1, 0xd0, 0x63, 0x64, 0x41, 0x97, 0xbf, 0x92, 0x15, 0xb3, 0x6b, 0x65, 0xb3,
	0x4f, 0x64, 0xee, 0x9d, 0x48, 0x8d, 0x2b, 0xf8, 0x30, 0x28, 0x07, 0xd0, 0x37, 0x60, 0xfa, 0x91,
	0x30, 0xb4, 0xba, 0x0c, 0xee, 0x5c, 0xc4, 0xe3, 0x0a, 0xae, 0xfb, 0x11, 0x1f, 0xbd, 0x36, 0x95,
	0x16, 0x3d, 0x0c, 0xad, 0xf2, 0x7e, 0xd0, 0x23, 0xa8, 0x67, 0x24, 0xbd, 0x21, 0xa9, 0x92, 0x41,
	0x21, 0x1e, 0x4f, 0x52, 0xb2, 0x0e, 0x6e, 0x95, 0x04, 0x0a, 0xed, 0x84, 0xd6, 0xf7, 0x42, 0xf7,
	0x5e, 0xc0, 0xe1, 0x9d, 0x7d, 0x3e, 0x38, 0x29, 0x02, 0xe3, 0xca, 0xcb, 0xae, 0x8a, 0x87, 0x85,
	0x8f, 0x7b, 0x31, 0xd4, 0xe5, 0x6e, 0x1f, 0xac, 0x6a, 0x43, 0xed, 0xb7, 0x9c, 0x64, 0xc5, 0x7b,
	0x24, 0x01, 0xb2, 0xc1, 0xf4, 0x18, 0x23, 0x51, 0xc2, 0xc4, 0x5e, 0x74, 0x5c, 0x40, 0xee, 0x1a,
	0x72, 0x4b, 0x56, 0x39, 0x77, 0x9d, 0x38, 0x09, 0x1d, 0xef, 0x03, 0xdf, 0xbe, 0x84, 0xba, 0xbc,
	0xc4, 0xf9, 0xd3, 0x81, 0x2f, 0x67, 0xb3, 0xc9, 0xec, 0x8d, 0x55, 0xe1, 0x60, 0x7e, 0x39, 0x1c,
	0xba, 0xf3, 0xb9, 0xa5, 0x71, 0xc0, 0xdf, 0x9b, 0x4b, 0xec, 0x5a, 0x55, 0x0e, 0xf8, 0x2b, 0xc4,
	0x69, 0xfa, 0xb2, 0x2e, 0x2e, 0xb1, 0x1f, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x81, 0x81,
	0x89, 0xc9, 0x07, 0x00, 0x00,
}
