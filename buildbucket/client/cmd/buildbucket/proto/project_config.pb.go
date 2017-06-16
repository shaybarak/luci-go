// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/luci/luci-go/buildbucket/client/cmd/buildbucket/proto/project_config.proto

/*
Package buildbucket is a generated protocol buffer package.

It is generated from these files:
	github.com/luci/luci-go/buildbucket/client/cmd/buildbucket/proto/project_config.proto

It has these top-level messages:
	Acl
	AclSet
	Swarming
	Bucket
	BuildbucketCfg
*/
package buildbucket

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Acl_Role int32

const (
	// Can do read-only operations, such as search for builds.
	Acl_READER Acl_Role = 0
	// Same as READER + can schedule and cancel builds.
	Acl_SCHEDULER Acl_Role = 1
	// Can do all write operations.
	Acl_WRITER Acl_Role = 2
)

var Acl_Role_name = map[int32]string{
	0: "READER",
	1: "SCHEDULER",
	2: "WRITER",
}
var Acl_Role_value = map[string]int32{
	"READER":    0,
	"SCHEDULER": 1,
	"WRITER":    2,
}

func (x Acl_Role) Enum() *Acl_Role {
	p := new(Acl_Role)
	*p = x
	return p
}
func (x Acl_Role) String() string {
	return proto.EnumName(Acl_Role_name, int32(x))
}
func (x *Acl_Role) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Acl_Role_value, data, "Acl_Role")
	if err != nil {
		return err
	}
	*x = Acl_Role(value)
	return nil
}
func (Acl_Role) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// A single access control rule.
type Acl struct {
	// Role denotes a list of actions that an identity can perform.
	Role *Acl_Role `protobuf:"varint,1,opt,name=role,enum=buildbucket.Acl_Role" json:"role,omitempty"`
	// Name of the group defined in the auth service.
	Group *string `protobuf:"bytes,2,opt,name=group" json:"group,omitempty"`
	// An email address or a full identity string "kind:name". See auth service
	// on kinds of identities. Anonymous users are "anonymous:anonymous".
	// Either identity or group must be present, not both.
	Identity         *string `protobuf:"bytes,3,opt,name=identity" json:"identity,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Acl) Reset()                    { *m = Acl{} }
func (m *Acl) String() string            { return proto.CompactTextString(m) }
func (*Acl) ProtoMessage()               {}
func (*Acl) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Acl) GetRole() Acl_Role {
	if m != nil && m.Role != nil {
		return *m.Role
	}
	return Acl_READER
}

func (m *Acl) GetGroup() string {
	if m != nil && m.Group != nil {
		return *m.Group
	}
	return ""
}

func (m *Acl) GetIdentity() string {
	if m != nil && m.Identity != nil {
		return *m.Identity
	}
	return ""
}

// A set of Acl messages. Can be referenced in a bucket by name.
type AclSet struct {
	// A name of the ACL set. Required. Must match regex '^[a-z0-9_]+$'.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// List of access control rules.
	// The order does not matter.
	Acls             []*Acl `protobuf:"bytes,2,rep,name=acls" json:"acls,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *AclSet) Reset()                    { *m = AclSet{} }
func (m *AclSet) String() string            { return proto.CompactTextString(m) }
func (*AclSet) ProtoMessage()               {}
func (*AclSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AclSet) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *AclSet) GetAcls() []*Acl {
	if m != nil {
		return m.Acls
	}
	return nil
}

// Configuration of buildbucket-swarming integration for one bucket.
type Swarming struct {
	// Hostname of the swarming instance, e.g. "chromium-swarm.appspot.com".
	Hostname *string `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	// Used to generate a URL for Build, may contain parameters
	// {swarming_hostname}, {task_id}, {bucket} and {builder}. Defaults to:
	// https://{swarming_hostname}/user/task/{task_id}
	UrlFormat *string `protobuf:"bytes,2,opt,name=url_format,json=urlFormat" json:"url_format,omitempty"`
	// Defines default values for builders.
	BuilderDefaults *Swarming_Builder `protobuf:"bytes,3,opt,name=builder_defaults,json=builderDefaults" json:"builder_defaults,omitempty"`
	// Configuration for each builder.
	// Swarming tasks are created only for builds for builders that are not
	// explicitly specified.
	Builders []*Swarming_Builder `protobuf:"bytes,4,rep,name=builders" json:"builders,omitempty"`
	// Percentage of builds that should use a canary swarming task template.
	// A value from 0 to 100.
	TaskTemplateCanaryPercentage *uint32 `protobuf:"varint,5,opt,name=task_template_canary_percentage,json=taskTemplateCanaryPercentage" json:"task_template_canary_percentage,omitempty"`
	XXX_unrecognized             []byte  `json:"-"`
}

func (m *Swarming) Reset()                    { *m = Swarming{} }
func (m *Swarming) String() string            { return proto.CompactTextString(m) }
func (*Swarming) ProtoMessage()               {}
func (*Swarming) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Swarming) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *Swarming) GetUrlFormat() string {
	if m != nil && m.UrlFormat != nil {
		return *m.UrlFormat
	}
	return ""
}

func (m *Swarming) GetBuilderDefaults() *Swarming_Builder {
	if m != nil {
		return m.BuilderDefaults
	}
	return nil
}

func (m *Swarming) GetBuilders() []*Swarming_Builder {
	if m != nil {
		return m.Builders
	}
	return nil
}

func (m *Swarming) GetTaskTemplateCanaryPercentage() uint32 {
	if m != nil && m.TaskTemplateCanaryPercentage != nil {
		return *m.TaskTemplateCanaryPercentage
	}
	return 0
}

type Swarming_Recipe struct {
	// Repository URL of the recipe package.
	Repository *string `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	// Name of the recipe to run.
	Name *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// colon-separated build properties to set.
	// A property can be overriden by "properties" build parameter.
	//
	// Use this field for string properties and use properties_j for other
	// types.
	Properties []string `protobuf:"bytes,3,rep,name=properties" json:"properties,omitempty"`
	// Same as properties, but the value must valid JSON. For example
	//   properties_j: "a:1"
	// means property a is a number 1, not string "1".
	//
	// If null, it means no property must be defined. In particular, it removes
	// a default value for the property, if any.
	//
	// Fields properties and properties_j can be used together, but cannot both
	// specify values for same property.
	PropertiesJ      []string `protobuf:"bytes,4,rep,name=properties_j,json=propertiesJ" json:"properties_j,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Swarming_Recipe) Reset()                    { *m = Swarming_Recipe{} }
func (m *Swarming_Recipe) String() string            { return proto.CompactTextString(m) }
func (*Swarming_Recipe) ProtoMessage()               {}
func (*Swarming_Recipe) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *Swarming_Recipe) GetRepository() string {
	if m != nil && m.Repository != nil {
		return *m.Repository
	}
	return ""
}

func (m *Swarming_Recipe) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Swarming_Recipe) GetProperties() []string {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *Swarming_Recipe) GetPropertiesJ() []string {
	if m != nil {
		return m.PropertiesJ
	}
	return nil
}

// A builder has a name, a category and specifies what should happen if a
// build is scheduled to that builder.
//
// SECURITY WARNING: if adding more fields to this message, keep in mind that
// a user that has permissions to schedule a build to the bucket, can override
// this config.
type Swarming_Builder struct {
	// Name of the builder. Will be propagated to "builder" build tag and
	// "buildername" recipe property.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Builder category. Will be used for visual grouping, for example in Code Review.
	Category *string `protobuf:"bytes,6,opt,name=category" json:"category,omitempty"`
	// Will be become to swarming task tags.
	// Each tag will end up in "swarming_tag" buildbucket tag, for example
	// "swarming_tag:builder:release"
	SwarmingTags []string `protobuf:"bytes,2,rep,name=swarming_tags,json=swarmingTags" json:"swarming_tags,omitempty"`
	// Colon-delimited key-value pair of task dimensions.
	//
	// If value is not specified ("<key>:"), then it excludes a default value.
	Dimensions []string `protobuf:"bytes,3,rep,name=dimensions" json:"dimensions,omitempty"`
	// CIPD packages to install on the builder.
	CipdPackages []*Swarming_Builder_CipdPackage `protobuf:"bytes,8,rep,name=cipd_packages,json=cipdPackages" json:"cipd_packages,omitempty"`
	// Specifies that a recipe to run.
	Recipe *Swarming_Recipe `protobuf:"bytes,4,opt,name=recipe" json:"recipe,omitempty"`
	// Swarming task priority.
	Priority *uint32 `protobuf:"varint,5,opt,name=priority" json:"priority,omitempty"`
	// Maximum build execution time. Not to be confused with pending time.
	ExecutionTimeoutSecs *uint32 `protobuf:"varint,7,opt,name=execution_timeout_secs,json=executionTimeoutSecs" json:"execution_timeout_secs,omitempty"`
	// Caches that should be present on the bot.
	Caches           []*Swarming_Builder_CacheEntry `protobuf:"bytes,9,rep,name=caches" json:"caches,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *Swarming_Builder) Reset()                    { *m = Swarming_Builder{} }
func (m *Swarming_Builder) String() string            { return proto.CompactTextString(m) }
func (*Swarming_Builder) ProtoMessage()               {}
func (*Swarming_Builder) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 1} }

func (m *Swarming_Builder) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Swarming_Builder) GetCategory() string {
	if m != nil && m.Category != nil {
		return *m.Category
	}
	return ""
}

func (m *Swarming_Builder) GetSwarmingTags() []string {
	if m != nil {
		return m.SwarmingTags
	}
	return nil
}

func (m *Swarming_Builder) GetDimensions() []string {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *Swarming_Builder) GetCipdPackages() []*Swarming_Builder_CipdPackage {
	if m != nil {
		return m.CipdPackages
	}
	return nil
}

func (m *Swarming_Builder) GetRecipe() *Swarming_Recipe {
	if m != nil {
		return m.Recipe
	}
	return nil
}

func (m *Swarming_Builder) GetPriority() uint32 {
	if m != nil && m.Priority != nil {
		return *m.Priority
	}
	return 0
}

func (m *Swarming_Builder) GetExecutionTimeoutSecs() uint32 {
	if m != nil && m.ExecutionTimeoutSecs != nil {
		return *m.ExecutionTimeoutSecs
	}
	return 0
}

func (m *Swarming_Builder) GetCaches() []*Swarming_Builder_CacheEntry {
	if m != nil {
		return m.Caches
	}
	return nil
}

type Swarming_Builder_CipdPackage struct {
	// A template of a full CIPD package name, e.g
	// "infra/tools/authutil/${platform}". This can be parametrized using
	// ${platform} and ${os_ver} parameters, where ${platform} will be
	// expanded into "<os>-<architecture>" and ${os_ver} will be expanded to
	// OS version name.
	PackageName *string `protobuf:"bytes,1,opt,name=package_name,json=packageName" json:"package_name,omitempty"`
	// Path to dir, relative to the task working dir, where to install the
	// package. The path cannot be empty or start with a slash.
	Path *string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	// Valid package version for all packages matched by package name.
	Version          *string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Swarming_Builder_CipdPackage) Reset()         { *m = Swarming_Builder_CipdPackage{} }
func (m *Swarming_Builder_CipdPackage) String() string { return proto.CompactTextString(m) }
func (*Swarming_Builder_CipdPackage) ProtoMessage()    {}
func (*Swarming_Builder_CipdPackage) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 1, 0}
}

func (m *Swarming_Builder_CipdPackage) GetPackageName() string {
	if m != nil && m.PackageName != nil {
		return *m.PackageName
	}
	return ""
}

func (m *Swarming_Builder_CipdPackage) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *Swarming_Builder_CipdPackage) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

// Describes a named cache that should be present on the bot.
// See also https://github.com/luci/luci-py/blob/3a2941345cf011a96bcd83d76328395323245bb5/appengine/swarming/swarming_rpcs.py#L166
type Swarming_Builder_CacheEntry struct {
	// Unique name of the cache. Required. Length is limited to 4096.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Relative path to the directory that will be linked to the named cache.
	// Required.
	// A path cannot be shared among multiple caches or CIPD installations.
	// A task will fail if a file/dir with the same name already exists.
	Path             *string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Swarming_Builder_CacheEntry) Reset()         { *m = Swarming_Builder_CacheEntry{} }
func (m *Swarming_Builder_CacheEntry) String() string { return proto.CompactTextString(m) }
func (*Swarming_Builder_CacheEntry) ProtoMessage()    {}
func (*Swarming_Builder_CacheEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 1, 1}
}

func (m *Swarming_Builder_CacheEntry) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Swarming_Builder_CacheEntry) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

// Defines one bucket in buildbucket.cfg
type Bucket struct {
	// Name of the bucket. Names are unique within one instance of buildbucket.
	// If another project already uses this name, a config will be rejected.
	// Name reservation is first-come first-serve.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// List of access control rules for the bucket.
	// The order does not matter.
	Acls []*Acl `protobuf:"bytes,2,rep,name=acls" json:"acls,omitempty"`
	// A list of ACL set names. Each ACL in each referenced ACL set will be
	// included in this bucket.
	// The order does not matter.
	AclSets []string `protobuf:"bytes,4,rep,name=acl_sets,json=aclSets" json:"acl_sets,omitempty"`
	// Buildbucket-swarming integration.
	Swarming         *Swarming `protobuf:"bytes,3,opt,name=swarming" json:"swarming,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Bucket) Reset()                    { *m = Bucket{} }
func (m *Bucket) String() string            { return proto.CompactTextString(m) }
func (*Bucket) ProtoMessage()               {}
func (*Bucket) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Bucket) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Bucket) GetAcls() []*Acl {
	if m != nil {
		return m.Acls
	}
	return nil
}

func (m *Bucket) GetAclSets() []string {
	if m != nil {
		return m.AclSets
	}
	return nil
}

func (m *Bucket) GetSwarming() *Swarming {
	if m != nil {
		return m.Swarming
	}
	return nil
}

// Schema of buildbucket.cfg file, a project config.
type BuildbucketCfg struct {
	// All buckets defined for this project.
	Buckets []*Bucket `protobuf:"bytes,1,rep,name=buckets" json:"buckets,omitempty"`
	// A list of ACL sets. Names must be unique.
	AclSets          []*AclSet `protobuf:"bytes,2,rep,name=acl_sets,json=aclSets" json:"acl_sets,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *BuildbucketCfg) Reset()                    { *m = BuildbucketCfg{} }
func (m *BuildbucketCfg) String() string            { return proto.CompactTextString(m) }
func (*BuildbucketCfg) ProtoMessage()               {}
func (*BuildbucketCfg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BuildbucketCfg) GetBuckets() []*Bucket {
	if m != nil {
		return m.Buckets
	}
	return nil
}

func (m *BuildbucketCfg) GetAclSets() []*AclSet {
	if m != nil {
		return m.AclSets
	}
	return nil
}

func init() {
	proto.RegisterType((*Acl)(nil), "buildbucket.Acl")
	proto.RegisterType((*AclSet)(nil), "buildbucket.AclSet")
	proto.RegisterType((*Swarming)(nil), "buildbucket.Swarming")
	proto.RegisterType((*Swarming_Recipe)(nil), "buildbucket.Swarming.Recipe")
	proto.RegisterType((*Swarming_Builder)(nil), "buildbucket.Swarming.Builder")
	proto.RegisterType((*Swarming_Builder_CipdPackage)(nil), "buildbucket.Swarming.Builder.CipdPackage")
	proto.RegisterType((*Swarming_Builder_CacheEntry)(nil), "buildbucket.Swarming.Builder.CacheEntry")
	proto.RegisterType((*Bucket)(nil), "buildbucket.Bucket")
	proto.RegisterType((*BuildbucketCfg)(nil), "buildbucket.BuildbucketCfg")
	proto.RegisterEnum("buildbucket.Acl_Role", Acl_Role_name, Acl_Role_value)
}

func init() {
	proto.RegisterFile("github.com/luci/luci-go/buildbucket/client/cmd/buildbucket/proto/project_config.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 736 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0xdb, 0x38,
	0x10, 0x5e, 0xd9, 0x8a, 0x7f, 0xc6, 0x71, 0xd6, 0xe0, 0x26, 0x0b, 0xad, 0x91, 0x64, 0xbd, 0xde,
	0x3d, 0x38, 0x87, 0xc8, 0xd8, 0x20, 0x97, 0xbd, 0xad, 0xed, 0xb8, 0x48, 0x8b, 0x22, 0x08, 0x68,
	0x07, 0xbd, 0x55, 0x60, 0x68, 0x5a, 0x66, 0x22, 0x89, 0x02, 0x49, 0xb5, 0xf5, 0xa9, 0xaf, 0xd0,
	0xbe, 0x48, 0x5f, 0xa2, 0x2f, 0x56, 0x88, 0x92, 0x65, 0xa5, 0x30, 0x1a, 0xa0, 0x17, 0x81, 0xf3,
	0xf1, 0x9b, 0xe1, 0x37, 0xf3, 0x51, 0x84, 0x3b, 0x9f, 0xeb, 0x55, 0x72, 0xef, 0x52, 0x11, 0x0e,
	0x83, 0x84, 0x72, 0xf3, 0x39, 0xf7, 0xc5, 0xf0, 0x3e, 0xe1, 0xc1, 0xe2, 0x3e, 0xa1, 0x8f, 0x4c,
	0x0f, 0x69, 0xc0, 0x59, 0xa4, 0x87, 0x34, 0x5c, 0x3c, 0x81, 0x63, 0x29, 0xb4, 0x48, 0xbf, 0x0f,
	0x8c, 0x6a, 0x8f, 0x8a, 0x68, 0xc9, 0x7d, 0xd7, 0x80, 0xa8, 0x55, 0xe2, 0xf5, 0x3f, 0x5b, 0x50,
	0x1d, 0xd1, 0x00, 0x9d, 0x81, 0x2d, 0x45, 0xc0, 0x1c, 0xab, 0x67, 0x0d, 0x0e, 0x2e, 0x8e, 0xdc,
	0x12, 0xc7, 0x1d, 0xd1, 0xc0, 0xc5, 0x22, 0x60, 0xd8, 0x50, 0xd0, 0x21, 0xec, 0xf9, 0x52, 0x24,
	0xb1, 0x53, 0xe9, 0x59, 0x83, 0x26, 0xce, 0x02, 0xd4, 0x85, 0x06, 0x5f, 0xb0, 0x48, 0x73, 0xbd,
	0x76, 0xaa, 0x66, 0xa3, 0x88, 0xfb, 0xe7, 0x60, 0xa7, 0xf9, 0x08, 0xa0, 0x86, 0xa7, 0xa3, 0xab,
	0x29, 0xee, 0xfc, 0x82, 0xda, 0xd0, 0x9c, 0x4d, 0xae, 0xa7, 0x57, 0x77, 0xaf, 0xa7, 0xb8, 0x63,
	0xa5, 0x5b, 0x6f, 0xf0, 0xcb, 0xf9, 0x14, 0x77, 0x2a, 0xfd, 0x31, 0xd4, 0x46, 0x34, 0x98, 0x31,
	0x8d, 0x10, 0xd8, 0x11, 0x09, 0x33, 0x55, 0x4d, 0x6c, 0xd6, 0xe8, 0x1f, 0xb0, 0x09, 0x0d, 0x94,
	0x53, 0xe9, 0x55, 0x07, 0xad, 0x8b, 0xce, 0xf7, 0x4a, 0xb1, 0xd9, 0xed, 0x7f, 0xad, 0x43, 0x63,
	0xf6, 0x9e, 0xc8, 0x90, 0x47, 0x7e, 0xaa, 0x6d, 0x25, 0x94, 0x2e, 0x95, 0x2a, 0x62, 0x74, 0x02,
	0x90, 0xc8, 0xc0, 0x5b, 0x0a, 0x19, 0x12, 0x9d, 0xb7, 0xd4, 0x4c, 0x64, 0xf0, 0xc2, 0x00, 0xe8,
	0x1a, 0x3a, 0xe6, 0x00, 0x26, 0xbd, 0x05, 0x5b, 0x92, 0x24, 0xd0, 0xca, 0xb4, 0xd7, 0xba, 0x38,
	0x79, 0x72, 0xf2, 0xe6, 0x2c, 0x77, 0x9c, 0xb1, 0xf1, 0xaf, 0x79, 0xda, 0x55, 0x9e, 0x85, 0xfe,
	0x83, 0x46, 0x0e, 0x29, 0xc7, 0x36, 0xda, 0x9f, 0xa9, 0x50, 0xd0, 0xd1, 0x14, 0xfe, 0xd4, 0x44,
	0x3d, 0x7a, 0x9a, 0x85, 0x71, 0x40, 0x34, 0xf3, 0x28, 0x89, 0x88, 0x5c, 0x7b, 0x31, 0x93, 0x94,
	0x45, 0x9a, 0xf8, 0xcc, 0xd9, 0xeb, 0x59, 0x83, 0x36, 0x3e, 0x4e, 0x69, 0xf3, 0x9c, 0x35, 0x31,
	0xa4, 0xdb, 0x82, 0xd3, 0xfd, 0x08, 0x35, 0xcc, 0x28, 0x8f, 0x19, 0x3a, 0x05, 0x90, 0x2c, 0x16,
	0x8a, 0x6b, 0x21, 0xd7, 0xf9, 0x48, 0x4a, 0x48, 0x31, 0xf7, 0x4a, 0x69, 0xee, 0xa7, 0x00, 0xb1,
	0x14, 0x31, 0x93, 0x9a, 0xb3, 0x74, 0x06, 0xd5, 0x34, 0x67, 0x8b, 0xa0, 0xbf, 0x60, 0x7f, 0x1b,
	0x79, 0x0f, 0xa6, 0xc7, 0x26, 0x6e, 0x6d, 0xb1, 0x57, 0xdd, 0x2f, 0x36, 0xd4, 0xf3, 0xee, 0x76,
	0x5a, 0xdb, 0x85, 0x06, 0x25, 0x9a, 0xf9, 0xa9, 0xa8, 0x5a, 0xe6, 0xd3, 0x26, 0x46, 0x7f, 0x43,
	0x5b, 0xe5, 0x13, 0xf2, 0x34, 0xf1, 0x33, 0xff, 0x9b, 0x78, 0x7f, 0x03, 0xce, 0x89, 0xaf, 0x52,
	0x8d, 0x0b, 0x1e, 0xb2, 0x48, 0x71, 0x11, 0x15, 0x1a, 0xb7, 0x08, 0xba, 0x81, 0x36, 0xe5, 0xf1,
	0xc2, 0x8b, 0x09, 0x7d, 0x24, 0x3e, 0x53, 0x4e, 0xc3, 0x18, 0x71, 0xf6, 0x43, 0x23, 0xdc, 0x09,
	0x8f, 0x17, 0xb7, 0x59, 0x06, 0xde, 0xa7, 0xdb, 0x40, 0xa1, 0x4b, 0xa8, 0x49, 0x33, 0x51, 0xc7,
	0x36, 0x77, 0xe2, 0x78, 0x77, 0xa1, 0x6c, 0xea, 0x38, 0xe7, 0xa6, 0x6d, 0xc6, 0x92, 0x0b, 0x99,
	0xfe, 0x2a, 0x99, 0x6f, 0x45, 0x8c, 0x2e, 0xe1, 0x77, 0xf6, 0x81, 0xd1, 0x44, 0x73, 0x11, 0x79,
	0x9a, 0x87, 0x4c, 0x24, 0xda, 0x53, 0x8c, 0x2a, 0xa7, 0x6e, 0x98, 0x87, 0xc5, 0xee, 0x3c, 0xdb,
	0x9c, 0x31, 0xaa, 0xd0, 0xff, 0x50, 0xa3, 0x84, 0xae, 0x98, 0x72, 0x9a, 0xa6, 0xa1, 0xc1, 0x33,
	0x0d, 0xa5, 0xdc, 0x69, 0xa4, 0xe5, 0x1a, 0xe7, 0x79, 0xdd, 0xb7, 0xd0, 0x2a, 0xb5, 0x69, 0xcc,
	0xcc, 0x96, 0x5e, 0xc9, 0xa5, 0x56, 0x8e, 0xdd, 0xa4, 0x66, 0x21, 0xb0, 0x63, 0xa2, 0x57, 0x9b,
	0x3b, 0x92, 0xae, 0x91, 0x03, 0xf5, 0x77, 0x4c, 0xa6, 0xb3, 0xce, 0xdf, 0x80, 0x4d, 0xd8, 0xbd,
	0x04, 0xd8, 0x9e, 0xba, 0xd3, 0xfc, 0x1d, 0xf5, 0xfa, 0x9f, 0x2c, 0xa8, 0x8d, 0x4d, 0x13, 0x3f,
	0xff, 0x14, 0xa0, 0x3f, 0xa0, 0x41, 0x68, 0xe0, 0x29, 0xa6, 0x55, 0x7e, 0x29, 0xeb, 0xc4, 0x3c,
	0x2f, 0x0a, 0xfd, 0x0b, 0x8d, 0xcd, 0xfd, 0xc9, 0xff, 0xea, 0xa3, 0x9d, 0x93, 0xc3, 0x05, 0xad,
	0x2f, 0xe0, 0x60, 0xbc, 0x65, 0x4c, 0x96, 0x3e, 0x3a, 0x87, 0x7a, 0x16, 0x28, 0xc7, 0x32, 0x42,
	0x7e, 0x7b, 0x52, 0x23, 0xd3, 0x8f, 0x37, 0x1c, 0xe4, 0x96, 0xe4, 0x54, 0x76, 0xf0, 0xb3, 0xa7,
	0xaf, 0xd0, 0xf8, 0x2d, 0x00, 0x00, 0xff, 0xff, 0xb2, 0xed, 0xa5, 0x27, 0x06, 0x06, 0x00, 0x00,
}
