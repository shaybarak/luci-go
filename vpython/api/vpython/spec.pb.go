// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/luci/luci-go/vpython/api/vpython/spec.proto

package vpython

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Spec is a "vpython" environment specification.
type Spec struct {
	// The Python version to use. This should be of the form:
	// "Major[.Minor[.Patch]]"
	//
	// If specified,
	// - The Major version will be enforced absolutely. Python 3 will not be
	//   preferred over Python 2 because '3' is greater than '2'.
	// - The remaining versions, if specified, will be regarded as *minimum*
	//   versions. In other words, if "2.7.4" is specified and the system has
	//   "2.7.12", that will suffice. Similarly, "2.6" would accept a "2.7"
	//   interpreter.
	//
	// If empty, the default Python interpreter ("python") will be used.
	PythonVersion string          `protobuf:"bytes,1,opt,name=python_version,json=pythonVersion" json:"python_version,omitempty"`
	Wheel         []*Spec_Package `protobuf:"bytes,2,rep,name=wheel" json:"wheel,omitempty"`
	// The VirtualEnv package.
	//
	// This should be left empty to use the `vpython` default package
	// (recommended).
	Virtualenv *Spec_Package `protobuf:"bytes,3,opt,name=virtualenv" json:"virtualenv,omitempty"`
	// Specification-provided PEP425 verification tags.
	//
	// By default, verification will be performed against a default set of
	// environment parameters. However, a given specification may offer its own
	// set of PEP425 tags representing the systems that it wants to be verified
	// against.
	VerifyPep425Tag []*PEP425Tag `protobuf:"bytes,4,rep,name=verify_pep425_tag,json=verifyPep425Tag" json:"verify_pep425_tag,omitempty"`
}

func (m *Spec) Reset()                    { *m = Spec{} }
func (m *Spec) String() string            { return proto.CompactTextString(m) }
func (*Spec) ProtoMessage()               {}
func (*Spec) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Spec) GetPythonVersion() string {
	if m != nil {
		return m.PythonVersion
	}
	return ""
}

func (m *Spec) GetWheel() []*Spec_Package {
	if m != nil {
		return m.Wheel
	}
	return nil
}

func (m *Spec) GetVirtualenv() *Spec_Package {
	if m != nil {
		return m.Virtualenv
	}
	return nil
}

func (m *Spec) GetVerifyPep425Tag() []*PEP425Tag {
	if m != nil {
		return m.VerifyPep425Tag
	}
	return nil
}

// A definition for a remote package. The type of package depends on the
// configured package resolver.
type Spec_Package struct {
	// The name of the package.
	//
	// - For CIPD, this is the package name.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The package version.
	//
	// - For CIPD, this will be any recognized CIPD version (i.e., ID, tag, or
	//   ref).
	Version string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	// Optional PEP425 tags to determine whether this package is included on the
	// target system. If no match tags are specified, this package will always
	// be included. If match tags are specified, the package will be included if
	// any system PEP425 tags match at least one of the match tags.
	//
	// A match will succeed if any system PEP425 tag field matches the
	// corresponding field in the PEP425 tag. If the match tag omits a field
	// (partial), that field will not be considered. For example, if a match
	// tag specifies just an ABI field, any system PEP425 tag with that ABI will
	// be considered a successful match, regardless of other field values.
	MatchTag []*PEP425Tag `protobuf:"bytes,3,rep,name=match_tag,json=matchTag" json:"match_tag,omitempty"`
}

func (m *Spec_Package) Reset()                    { *m = Spec_Package{} }
func (m *Spec_Package) String() string            { return proto.CompactTextString(m) }
func (*Spec_Package) ProtoMessage()               {}
func (*Spec_Package) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

func (m *Spec_Package) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Spec_Package) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Spec_Package) GetMatchTag() []*PEP425Tag {
	if m != nil {
		return m.MatchTag
	}
	return nil
}

func init() {
	proto.RegisterType((*Spec)(nil), "vpython.Spec")
	proto.RegisterType((*Spec_Package)(nil), "vpython.Spec.Package")
}

func init() {
	proto.RegisterFile("github.com/luci/luci-go/vpython/api/vpython/spec.proto", fileDescriptor2)
}

var fileDescriptor2 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xe9, 0x1f, 0xad, 0x3b, 0xa2, 0x62, 0x40, 0x28, 0x7b, 0x2a, 0x82, 0xb0, 0x20, 0xb6,
	0x50, 0xb7, 0xe2, 0xc9, 0x9b, 0xf7, 0x52, 0x8b, 0xd7, 0x92, 0x0d, 0x63, 0x1a, 0x6c, 0x9b, 0xd0,
	0x4d, 0x23, 0xfb, 0xd9, 0xfc, 0x72, 0x62, 0xd2, 0x5d, 0xbc, 0xec, 0xc1, 0x4b, 0x98, 0xbc, 0xbc,
	0xf9, 0xbd, 0xcc, 0xc0, 0x13, 0x17, 0xba, 0x9d, 0x36, 0x29, 0x93, 0x7d, 0xd6, 0x4d, 0x4c, 0xd8,
	0xe3, 0x81, 0xcb, 0xcc, 0xa8, 0x9d, 0x6e, 0xe5, 0x90, 0x51, 0x25, 0x0e, 0xf5, 0x56, 0x21, 0x4b,
	0xd5, 0x28, 0xb5, 0x24, 0xd1, 0xac, 0x2d, 0x9f, 0xff, 0x03, 0x50, 0xa8, 0xd6, 0x79, 0xe1, 0x10,
	0xb7, 0xdf, 0x3e, 0x84, 0x6f, 0x0a, 0x19, 0xb9, 0x83, 0x4b, 0xf7, 0xde, 0x18, 0x1c, 0xb7, 0x42,
	0x0e, 0xb1, 0x97, 0x78, 0xab, 0x45, 0x75, 0xe1, 0xd4, 0x77, 0x27, 0x92, 0x7b, 0x38, 0xf9, 0x6a,
	0x11, 0xbb, 0xd8, 0x4f, 0x82, 0xd5, 0x79, 0x7e, 0x93, 0xce, 0xd4, 0xf4, 0x17, 0x92, 0x96, 0x94,
	0x7d, 0x52, 0x8e, 0x95, 0xf3, 0x90, 0x02, 0xc0, 0x88, 0x51, 0x4f, 0xb4, 0xc3, 0xc1, 0xc4, 0x41,
	0xe2, 0x1d, 0xef, 0xf8, 0x63, 0x24, 0x2f, 0x70, 0x6d, 0x70, 0x14, 0x1f, 0xbb, 0xc6, 0x7d, 0xb5,
	0xd1, 0x94, 0xc7, 0xa1, 0xcd, 0x23, 0x87, 0xee, 0xf2, 0xb5, 0x5c, 0xe7, 0x45, 0x4d, 0x79, 0x75,
	0xe5, 0xcc, 0xa5, 0xf5, 0xd6, 0x94, 0x2f, 0x5b, 0x88, 0x66, 0x2c, 0x21, 0x10, 0x0e, 0xb4, 0xc7,
	0x79, 0x16, 0x5b, 0x93, 0x18, 0xa2, 0xfd, 0x88, 0xbe, 0x95, 0xf7, 0x57, 0x92, 0xc1, 0xa2, 0xa7,
	0x9a, 0xb5, 0x36, 0x30, 0x38, 0x1a, 0x78, 0x66, 0x4d, 0x35, 0xe5, 0x9b, 0x53, 0xbb, 0xc4, 0xc7,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xe6, 0x2e, 0x37, 0xc1, 0x01, 0x00, 0x00,
}
