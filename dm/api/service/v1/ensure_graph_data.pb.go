// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/luci/luci-go/dm/api/service/v1/ensure_graph_data.proto

package dm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import templateproto "github.com/luci/luci-go/common/data/text/templateproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TemplateInstantiation struct {
	// project is the luci-config project which defines the template.
	Project string `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	// ref is the git ref of the project that defined this template. If omitted,
	// this will use the template definition from the project-wide configuration
	// and not the configuration located on a particular ref (like
	// 'refs/heads/master').
	Ref string `protobuf:"bytes,2,opt,name=ref" json:"ref,omitempty"`
	// specifier specifies the actual template name, as well as any substitution
	// parameters which that template might require.
	Specifier *templateproto.Specifier `protobuf:"bytes,4,opt,name=specifier" json:"specifier,omitempty"`
}

func (m *TemplateInstantiation) Reset()                    { *m = TemplateInstantiation{} }
func (m *TemplateInstantiation) String() string            { return proto.CompactTextString(m) }
func (*TemplateInstantiation) ProtoMessage()               {}
func (*TemplateInstantiation) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *TemplateInstantiation) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *TemplateInstantiation) GetRef() string {
	if m != nil {
		return m.Ref
	}
	return ""
}

func (m *TemplateInstantiation) GetSpecifier() *templateproto.Specifier {
	if m != nil {
		return m.Specifier
	}
	return nil
}

// EnsureGraphDataReq allows you to assert the existence of Attempts in DM's
// graph, and allows you to declare dependencies from one Attempt to another.
//
// You can declare Attempts by any combination of:
//   * Providing a quest description plus a list of Attempt numbers for that
//     quest.
//   * Providing a template instantiation (for a project-declared quest
//     template) plus a list of Attempt numbers for that quest.
//   * Providing a raw set of quest_id -> attempt numbers for quests that you
//     already know that DM has a definition for.
//
// In response, DM will tell you what the IDs of all supplied Quests/Attempts
// are.
//
// To create a dependencies, call this method while running as part of an
// execution by filling the for_execution field. All attempts named as described
// above will become dependencies for the indicated execution. It is only
// possible for a currently-running execution to create dependencies for its own
// Attempt. In particular, it is not possible to create dependencies as
// a non-execution user (e.g. a human), nor is it possible for an execution to
// create attempts on behalf of some other execution.
//
// If the attempts were being created as dependencies, and were already in the
// Finished state, this request can also opt to include the AttemptResults
// directly.
type EnsureGraphDataReq struct {
	// Quest is a list of quest descriptors. DM will ensure that the
	// corresponding Quests exist. If they don't, they'll be created.
	Quest []*Quest_Desc `protobuf:"bytes,1,rep,name=quest" json:"quest,omitempty"`
	// QuestAttempt allows the addition of attempts which are derived from
	// the quest bodies provided above.
	// Each entry here maps 1:1 with the equivalent quest.
	QuestAttempt []*AttemptList_Nums `protobuf:"bytes,2,rep,name=quest_attempt,json=questAttempt" json:"quest_attempt,omitempty"`
	// TemplateQuest allows the addition of quests which are derived from
	// Templates, as defined on a per-project basis.
	TemplateQuest []*TemplateInstantiation `protobuf:"bytes,3,rep,name=template_quest,json=templateQuest" json:"template_quest,omitempty"`
	// TemplateAttempt allows the addition of attempts which are derived from
	// Templates. This must be equal in length to template_quest.
	// Each entry here maps 1:1 with the equivalent quest in template_quest.
	TemplateAttempt []*AttemptList_Nums `protobuf:"bytes,4,rep,name=template_attempt,json=templateAttempt" json:"template_attempt,omitempty"`
	// RawAttempts is a list that asserts that the following attempts should
	// exist. The quest ids in this list must be already-known to DM, NOT
	// included in the quest field above. This is useful when you know the ID of
	// the Quest, but not the actual definition of the quest.
	RawAttempts *AttemptList `protobuf:"bytes,5,opt,name=raw_attempts,json=rawAttempts" json:"raw_attempts,omitempty"`
	// ForExecution is an authentication pair (Execution_ID, Token).
	//
	// If this is provided then it will serve as authorization for the creation of
	// any `quests` included, and any `attempts` indicated will be set as
	// dependencies for the execution.
	//
	// If this omitted, then the request requires some user/bot authentication,
	// and any quests/attempts provided will be made standalone (e.g. nothing will
	// depend on them).
	ForExecution *Execution_Auth             `protobuf:"bytes,6,opt,name=for_execution,json=forExecution" json:"for_execution,omitempty"`
	Limit        *EnsureGraphDataReq_Limit   `protobuf:"bytes,7,opt,name=limit" json:"limit,omitempty"`
	Include      *EnsureGraphDataReq_Include `protobuf:"bytes,8,opt,name=include" json:"include,omitempty"`
}

func (m *EnsureGraphDataReq) Reset()                    { *m = EnsureGraphDataReq{} }
func (m *EnsureGraphDataReq) String() string            { return proto.CompactTextString(m) }
func (*EnsureGraphDataReq) ProtoMessage()               {}
func (*EnsureGraphDataReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *EnsureGraphDataReq) GetQuest() []*Quest_Desc {
	if m != nil {
		return m.Quest
	}
	return nil
}

func (m *EnsureGraphDataReq) GetQuestAttempt() []*AttemptList_Nums {
	if m != nil {
		return m.QuestAttempt
	}
	return nil
}

func (m *EnsureGraphDataReq) GetTemplateQuest() []*TemplateInstantiation {
	if m != nil {
		return m.TemplateQuest
	}
	return nil
}

func (m *EnsureGraphDataReq) GetTemplateAttempt() []*AttemptList_Nums {
	if m != nil {
		return m.TemplateAttempt
	}
	return nil
}

func (m *EnsureGraphDataReq) GetRawAttempts() *AttemptList {
	if m != nil {
		return m.RawAttempts
	}
	return nil
}

func (m *EnsureGraphDataReq) GetForExecution() *Execution_Auth {
	if m != nil {
		return m.ForExecution
	}
	return nil
}

func (m *EnsureGraphDataReq) GetLimit() *EnsureGraphDataReq_Limit {
	if m != nil {
		return m.Limit
	}
	return nil
}

func (m *EnsureGraphDataReq) GetInclude() *EnsureGraphDataReq_Include {
	if m != nil {
		return m.Include
	}
	return nil
}

type EnsureGraphDataReq_Limit struct {
	// MaxDataSize sets the maximum amount of 'Data' (in bytes) that can be
	// returned, if include.attempt_result is set. If this limit is hit, then
	// the appropriate 'partial' value will be set for that object, but the base
	// object would still be included in the result.
	//
	// If this limit is 0, a default limit of 16MB will be used. If this limit
	// exceeds 30MB, it will be reduced to 30MB.
	MaxDataSize uint32 `protobuf:"varint,3,opt,name=max_data_size,json=maxDataSize" json:"max_data_size,omitempty"`
}

func (m *EnsureGraphDataReq_Limit) Reset()                    { *m = EnsureGraphDataReq_Limit{} }
func (m *EnsureGraphDataReq_Limit) String() string            { return proto.CompactTextString(m) }
func (*EnsureGraphDataReq_Limit) ProtoMessage()               {}
func (*EnsureGraphDataReq_Limit) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 0} }

func (m *EnsureGraphDataReq_Limit) GetMaxDataSize() uint32 {
	if m != nil {
		return m.MaxDataSize
	}
	return 0
}

type EnsureGraphDataReq_Include struct {
	Attempt *EnsureGraphDataReq_Include_Options `protobuf:"bytes,4,opt,name=attempt" json:"attempt,omitempty"`
}

func (m *EnsureGraphDataReq_Include) Reset()                    { *m = EnsureGraphDataReq_Include{} }
func (m *EnsureGraphDataReq_Include) String() string            { return proto.CompactTextString(m) }
func (*EnsureGraphDataReq_Include) ProtoMessage()               {}
func (*EnsureGraphDataReq_Include) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 1} }

func (m *EnsureGraphDataReq_Include) GetAttempt() *EnsureGraphDataReq_Include_Options {
	if m != nil {
		return m.Attempt
	}
	return nil
}

type EnsureGraphDataReq_Include_Options struct {
	// Instructs finished objects to include the Result field.
	Result bool `protobuf:"varint,3,opt,name=result" json:"result,omitempty"`
}

func (m *EnsureGraphDataReq_Include_Options) Reset()         { *m = EnsureGraphDataReq_Include_Options{} }
func (m *EnsureGraphDataReq_Include_Options) String() string { return proto.CompactTextString(m) }
func (*EnsureGraphDataReq_Include_Options) ProtoMessage()    {}
func (*EnsureGraphDataReq_Include_Options) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{1, 1, 0}
}

func (m *EnsureGraphDataReq_Include_Options) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type EnsureGraphDataRsp struct {
	// accepted is true when all new graph data was journaled successfully. This
	// means that `quests`, `attempts`, `template_quest`, `template_attempt` were
	// all well-formed and are scheduled to be added. They will 'eventually' be
	// readable via other APIs (like WalkGraph), but when they are, they'll have
	// the IDs reflected in this response.
	//
	// If `attempts` referrs to quests that don't exist and weren't provided in
	// `quests`, those quests will be listed in `result` with the DNE flag set.
	//
	// If `template_quest` had errors (missing template, bad params, etc.), the
	// errors will be located in `template_error`. If all of the templates parsed
	// successfully, the quest ids for those rendered `template_quest` will be in
	// `template_ids`.
	Accepted bool `protobuf:"varint,1,opt,name=accepted" json:"accepted,omitempty"`
	// quest_ids will be populated with the Quest.IDs of any quests defined
	// by quest in the initial request. Its length is guaranteed to match
	// the length of quest, if there were no errors.
	QuestIds []*Quest_ID `protobuf:"bytes,2,rep,name=quest_ids,json=questIds" json:"quest_ids,omitempty"`
	// template_ids will be populated with the Quest.IDs of any templates defined
	// by template_quest in the initial request. Its length is guaranteed to match
	// the length of template_quest, if there were no errors.
	TemplateIds []*Quest_ID `protobuf:"bytes,3,rep,name=template_ids,json=templateIds" json:"template_ids,omitempty"`
	// template_error is either empty if there were no template errors, or the
	// length of template_quest. Non-empty strings are errors.
	TemplateError []string `protobuf:"bytes,4,rep,name=template_error,json=templateError" json:"template_error,omitempty"`
	// result holds the graph data pertaining to the request, containing any
	// graph state that already existed at the time of the call. Any new data
	// that was added to the graph state (accepted==true) will appear with
	// `DNE==true`.
	//
	// Quest data will always be returned for any Quests which exist.
	//
	// If accepted==false, you can inspect this to determine why:
	//   * Quests (without data) mentioned by the `attempts` field that do not
	//     exist will have `DNE==true`.
	//
	// This also can be used to make adding dependencies a stateless
	// single-request action:
	//   * Attempts requested (assuming the corresponding Quest exists) will
	//     contain their current state. If Include.AttemptResult was true, the
	//     results will be populated (with the size limit mentioned in the request
	//     documentation).
	Result *GraphData `protobuf:"bytes,5,opt,name=result" json:"result,omitempty"`
	// (if `for_execution` was specified) ShouldHalt indicates that the request
	// was accepted by DM, and the execution should halt (DM will re-execute the
	// Attempt when it becomes unblocked). If this is true, then the execution's
	// auth Token is also revoked and will no longer work for futher API calls.
	//
	// If `for_execution` was provided in the request and this is false, it means
	// that the execution may continue executing.
	ShouldHalt bool `protobuf:"varint,6,opt,name=should_halt,json=shouldHalt" json:"should_halt,omitempty"`
}

func (m *EnsureGraphDataRsp) Reset()                    { *m = EnsureGraphDataRsp{} }
func (m *EnsureGraphDataRsp) String() string            { return proto.CompactTextString(m) }
func (*EnsureGraphDataRsp) ProtoMessage()               {}
func (*EnsureGraphDataRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *EnsureGraphDataRsp) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

func (m *EnsureGraphDataRsp) GetQuestIds() []*Quest_ID {
	if m != nil {
		return m.QuestIds
	}
	return nil
}

func (m *EnsureGraphDataRsp) GetTemplateIds() []*Quest_ID {
	if m != nil {
		return m.TemplateIds
	}
	return nil
}

func (m *EnsureGraphDataRsp) GetTemplateError() []string {
	if m != nil {
		return m.TemplateError
	}
	return nil
}

func (m *EnsureGraphDataRsp) GetResult() *GraphData {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *EnsureGraphDataRsp) GetShouldHalt() bool {
	if m != nil {
		return m.ShouldHalt
	}
	return false
}

func init() {
	proto.RegisterType((*TemplateInstantiation)(nil), "dm.TemplateInstantiation")
	proto.RegisterType((*EnsureGraphDataReq)(nil), "dm.EnsureGraphDataReq")
	proto.RegisterType((*EnsureGraphDataReq_Limit)(nil), "dm.EnsureGraphDataReq.Limit")
	proto.RegisterType((*EnsureGraphDataReq_Include)(nil), "dm.EnsureGraphDataReq.Include")
	proto.RegisterType((*EnsureGraphDataReq_Include_Options)(nil), "dm.EnsureGraphDataReq.Include.Options")
	proto.RegisterType((*EnsureGraphDataRsp)(nil), "dm.EnsureGraphDataRsp")
}

func init() {
	proto.RegisterFile("github.com/luci/luci-go/dm/api/service/v1/ensure_graph_data.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 648 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x5b, 0x6f, 0xd3, 0x30,
	0x18, 0x55, 0x9a, 0xa4, 0x49, 0xdd, 0x76, 0x8b, 0x2c, 0x40, 0x21, 0x42, 0x30, 0x55, 0x0c, 0x95,
	0x07, 0x12, 0x51, 0x04, 0x63, 0xbc, 0xb0, 0x49, 0xab, 0xa0, 0xd5, 0x04, 0xc2, 0xe3, 0x3d, 0xf2,
	0x12, 0x77, 0x35, 0xca, 0x6d, 0xb6, 0xb3, 0x95, 0xbd, 0xf1, 0x47, 0xf8, 0x07, 0xfc, 0x42, 0x5e,
	0x90, 0x9d, 0x4b, 0xc7, 0x6e, 0xda, 0x4b, 0x94, 0xef, 0x7c, 0xe7, 0x7c, 0x37, 0x1d, 0x83, 0xfd,
	0x13, 0x2a, 0x96, 0xe5, 0xb1, 0x1f, 0xe5, 0x69, 0x90, 0x94, 0x11, 0x55, 0x9f, 0x57, 0x27, 0x79,
	0x10, 0xa7, 0x01, 0x2e, 0x68, 0xc0, 0x09, 0x3b, 0xa3, 0x11, 0x09, 0xce, 0x5e, 0x07, 0x24, 0xe3,
	0x25, 0x23, 0xe1, 0x09, 0xc3, 0xc5, 0x32, 0x8c, 0xb1, 0xc0, 0x7e, 0xc1, 0x72, 0x91, 0xc3, 0x4e,
	0x9c, 0x7a, 0xd3, 0xdb, 0xca, 0x44, 0x79, 0x9a, 0xe6, 0x59, 0x20, 0xe9, 0x81, 0x20, 0x2b, 0x11,
	0x08, 0x92, 0x16, 0x09, 0x16, 0x44, 0x69, 0xdb, 0xa8, 0x2a, 0xe5, 0x7d, 0xb8, 0xff, 0x34, 0x57,
	0xc7, 0xf0, 0xde, 0xde, 0x5f, 0x2b, 0x7e, 0x16, 0x84, 0x57, 0xb2, 0xd1, 0x2f, 0x0d, 0x3c, 0xfc,
	0x5e, 0x4f, 0x31, 0xcb, 0xb8, 0xc0, 0x99, 0xa0, 0x58, 0xd0, 0x3c, 0x83, 0x2e, 0xb0, 0x0a, 0x96,
	0xff, 0x20, 0x91, 0x70, 0xb5, 0x2d, 0x6d, 0xdc, 0x43, 0x4d, 0x08, 0x1d, 0xa0, 0x33, 0xb2, 0x70,
	0x3b, 0x0a, 0x95, 0xbf, 0xf0, 0x1d, 0xe8, 0xf1, 0x82, 0x44, 0x74, 0x41, 0x09, 0x73, 0x8d, 0x2d,
	0x6d, 0xdc, 0x9f, 0xb8, 0xfe, 0x7f, 0xab, 0xfa, 0x47, 0x4d, 0x1e, 0xad, 0xa9, 0x73, 0xc3, 0xd6,
	0x1d, 0x63, 0xf4, 0xc7, 0x04, 0x70, 0xaa, 0xae, 0xfb, 0x49, 0x6e, 0x75, 0x80, 0x05, 0x46, 0xe4,
	0x14, 0x3e, 0x07, 0xe6, 0x69, 0x49, 0xb8, 0x6c, 0xaf, 0x8f, 0xfb, 0x93, 0x0d, 0x3f, 0x4e, 0xfd,
	0x6f, 0x12, 0xf0, 0x0f, 0x08, 0x8f, 0x50, 0x95, 0x84, 0xbb, 0x60, 0xa8, 0x7e, 0x42, 0x2c, 0x64,
	0x43, 0xe1, 0x76, 0x14, 0xfb, 0x81, 0x64, 0xef, 0x57, 0xd0, 0x21, 0xe5, 0xc2, 0xff, 0x52, 0xa6,
	0x1c, 0x0d, 0x14, 0xb5, 0x86, 0xe1, 0x1e, 0xd8, 0x68, 0x66, 0x0c, 0xab, 0x4e, 0xba, 0xd2, 0x3e,
	0x96, 0xda, 0x1b, 0x8f, 0x82, 0x86, 0x8d, 0x40, 0x0d, 0x02, 0x3f, 0x02, 0xa7, 0xad, 0xd0, 0xf4,
	0x37, 0xee, 0xe8, 0xbf, 0xd9, 0xb0, 0x9b, 0x11, 0x26, 0x60, 0xc0, 0xf0, 0x79, 0xa3, 0xe5, 0xae,
	0xa9, 0x6e, 0xb7, 0x79, 0x45, 0x8c, 0xfa, 0x0c, 0x9f, 0xd7, 0x31, 0x87, 0x3b, 0x60, 0xb8, 0xc8,
	0x59, 0x48, 0x56, 0x24, 0x2a, 0xe5, 0x50, 0x6e, 0x57, 0x89, 0xa0, 0x14, 0x4d, 0x1b, 0xd0, 0xdf,
	0x2f, 0xc5, 0x12, 0x0d, 0x16, 0x39, 0x6b, 0x21, 0x38, 0x01, 0x66, 0x42, 0x53, 0x2a, 0x5c, 0x4b,
	0x09, 0x9e, 0x28, 0xc1, 0xb5, 0xbb, 0xfb, 0x87, 0x92, 0x83, 0x2a, 0x2a, 0x7c, 0x0f, 0x2c, 0x9a,
	0x45, 0x49, 0x19, 0x13, 0xd7, 0x56, 0xaa, 0xa7, 0xb7, 0xa8, 0x66, 0x15, 0x0b, 0x35, 0x74, 0x6f,
	0x07, 0x98, 0xaa, 0x12, 0x1c, 0x81, 0x61, 0x8a, 0x57, 0xca, 0xab, 0x21, 0xa7, 0x17, 0xc4, 0xd5,
	0xb7, 0xb4, 0xf1, 0x10, 0xf5, 0x53, 0xbc, 0x92, 0xe2, 0x23, 0x7a, 0x41, 0xe6, 0x86, 0xad, 0x39,
	0x9d, 0xb9, 0x61, 0x77, 0x1c, 0xdd, 0xfb, 0xad, 0x01, 0xab, 0xae, 0x06, 0xf7, 0x80, 0xb5, 0xbe,
	0xab, 0x6c, 0xff, 0xe2, 0xee, 0xf6, 0xfe, 0xd7, 0x42, 0xae, 0xca, 0x51, 0x23, 0xf3, 0x76, 0x81,
	0x55, 0x63, 0xf0, 0x11, 0xe8, 0x32, 0xc2, 0xcb, 0x44, 0xa8, 0x09, 0x6c, 0x54, 0x47, 0x97, 0x9b,
	0xcf, 0x0d, 0xdb, 0x70, 0xcc, 0xb9, 0x61, 0x9b, 0x4e, 0xb7, 0xc5, 0x75, 0xc7, 0x68, 0x91, 0xae,
	0x63, 0x8d, 0xfe, 0x6a, 0xd7, 0xfd, 0xca, 0x0b, 0xe8, 0x01, 0x1b, 0x47, 0x11, 0x29, 0x04, 0x89,
	0xd5, 0x8b, 0xb1, 0x51, 0x1b, 0xc3, 0x97, 0xa0, 0x57, 0xb9, 0x94, 0xc6, 0xbc, 0x76, 0xe8, 0x60,
	0xed, 0xe7, 0xd9, 0x01, 0xb2, 0x55, 0x7a, 0x16, 0x73, 0x18, 0x80, 0x41, 0xeb, 0x29, 0xc9, 0xd6,
	0x6f, 0x60, 0xf7, 0x1b, 0x86, 0x14, 0x6c, 0x5f, 0xb2, 0x31, 0x61, 0x2c, 0x67, 0xca, 0x82, 0xbd,
	0xb5, 0x57, 0xa7, 0x12, 0x84, 0xdb, 0xed, 0xf6, 0x95, 0xc9, 0x86, 0xb2, 0xe2, 0x7a, 0x81, 0x3a,
	0x09, 0x9f, 0x81, 0x3e, 0x5f, 0xe6, 0x65, 0x12, 0x87, 0x4b, 0x9c, 0x08, 0xe5, 0x2d, 0x1b, 0x81,
	0x0a, 0xfa, 0x8c, 0x13, 0x71, 0xdc, 0x55, 0xef, 0xf9, 0xcd, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xbd, 0x90, 0x48, 0x61, 0x3b, 0x05, 0x00, 0x00,
}
