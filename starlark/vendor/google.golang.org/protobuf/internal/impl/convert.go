// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package impl

import (
	"fmt"
	"reflect"

	pref "google.golang.org/protobuf/reflect/protoreflect"
)

// Unwrapper unwraps the value to the underlying value.
// This is implemented by List and Map.
type Unwrapper interface {
	ProtoUnwrap() interface{}
}

// A Converter coverts to/from Go reflect.Value types and protobuf protoreflect.Value types.
type Converter interface {
	PBValueOf(reflect.Value) pref.Value
	GoValueOf(pref.Value) reflect.Value
	New() pref.Value
}

// NewConverter matches a Go type with a protobuf field and returns a Converter
// that converts between the two. Enums must be a named int32 kind that
// implements protoreflect.Enum, and messages must be pointer to a named
// struct type that implements protoreflect.ProtoMessage.
//
// This matcher deliberately supports a wider range of Go types than what
// protoc-gen-go historically generated to be able to automatically wrap some
// v1 messages generated by other forks of protoc-gen-go.
func NewConverter(t reflect.Type, fd pref.FieldDescriptor) Converter {
	switch {
	case fd.IsList():
		return newListConverter(t, fd)
	case fd.IsMap():
		return newMapConverter(t, fd)
	default:
		return newSingularConverter(t, fd)
	}
	panic(fmt.Sprintf("invalid Go type %v for field %v", t, fd.FullName()))
}

var (
	boolType    = reflect.TypeOf(bool(false))
	int32Type   = reflect.TypeOf(int32(0))
	int64Type   = reflect.TypeOf(int64(0))
	uint32Type  = reflect.TypeOf(uint32(0))
	uint64Type  = reflect.TypeOf(uint64(0))
	float32Type = reflect.TypeOf(float32(0))
	float64Type = reflect.TypeOf(float64(0))
	stringType  = reflect.TypeOf(string(""))
	bytesType   = reflect.TypeOf([]byte(nil))
	byteType    = reflect.TypeOf(byte(0))
)

type scalarConverter struct {
	goType, pbType reflect.Type
	def            pref.Value
}

func newSingularConverter(t reflect.Type, fd pref.FieldDescriptor) Converter {
	switch fd.Kind() {
	case pref.BoolKind:
		if t.Kind() == reflect.Bool {
			return &scalarConverter{t, boolType, fd.Default()}
		}
	case pref.Int32Kind, pref.Sint32Kind, pref.Sfixed32Kind:
		if t.Kind() == reflect.Int32 {
			return &scalarConverter{t, int32Type, fd.Default()}
		}
	case pref.Int64Kind, pref.Sint64Kind, pref.Sfixed64Kind:
		if t.Kind() == reflect.Int64 {
			return &scalarConverter{t, int64Type, fd.Default()}
		}
	case pref.Uint32Kind, pref.Fixed32Kind:
		if t.Kind() == reflect.Uint32 {
			return &scalarConverter{t, uint32Type, fd.Default()}
		}
	case pref.Uint64Kind, pref.Fixed64Kind:
		if t.Kind() == reflect.Uint64 {
			return &scalarConverter{t, uint64Type, fd.Default()}
		}
	case pref.FloatKind:
		if t.Kind() == reflect.Float32 {
			return &scalarConverter{t, float32Type, fd.Default()}
		}
	case pref.DoubleKind:
		if t.Kind() == reflect.Float64 {
			return &scalarConverter{t, float64Type, fd.Default()}
		}
	case pref.StringKind:
		if t.Kind() == reflect.String || (t.Kind() == reflect.Slice && t.Elem() == byteType) {
			return &scalarConverter{t, stringType, fd.Default()}
		}
	case pref.BytesKind:
		if t.Kind() == reflect.String || (t.Kind() == reflect.Slice && t.Elem() == byteType) {
			return &scalarConverter{t, bytesType, fd.Default()}
		}
	case pref.EnumKind:
		// Handle enums, which must be a named int32 type.
		if t.Kind() == reflect.Int32 {
			return newEnumConverter(t, fd)
		}
	case pref.MessageKind, pref.GroupKind:
		return newMessageConverter(t)
	}
	panic(fmt.Sprintf("invalid Go type %v for field %v", t, fd.FullName()))
}

func (c *scalarConverter) PBValueOf(v reflect.Value) pref.Value {
	if v.Type() != c.goType {
		panic(fmt.Sprintf("invalid type: got %v, want %v", v.Type(), c.goType))
	}
	if c.goType.Kind() == reflect.String && c.pbType.Kind() == reflect.Slice && v.Len() == 0 {
		return pref.ValueOf([]byte(nil)) // ensure empty string is []byte(nil)
	}
	return pref.ValueOf(v.Convert(c.pbType).Interface())
}

func (c *scalarConverter) GoValueOf(v pref.Value) reflect.Value {
	rv := reflect.ValueOf(v.Interface())
	if rv.Type() != c.pbType {
		panic(fmt.Sprintf("invalid type: got %v, want %v", rv.Type(), c.pbType))
	}
	if c.pbType.Kind() == reflect.String && c.goType.Kind() == reflect.Slice && rv.Len() == 0 {
		return reflect.Zero(c.goType) // ensure empty string is []byte(nil)
	}
	return rv.Convert(c.goType)
}

func (c *scalarConverter) New() pref.Value {
	return c.def
}

type enumConverter struct {
	goType reflect.Type
	def    pref.Value
}

func newEnumConverter(goType reflect.Type, fd pref.FieldDescriptor) Converter {
	return &enumConverter{goType, fd.Default()}
}

func (c *enumConverter) PBValueOf(v reflect.Value) pref.Value {
	if v.Type() != c.goType {
		panic(fmt.Sprintf("invalid type: got %v, want %v", v.Type(), c.goType))
	}
	return pref.ValueOf(pref.EnumNumber(v.Int()))
}

func (c *enumConverter) GoValueOf(v pref.Value) reflect.Value {
	return reflect.ValueOf(v.Enum()).Convert(c.goType)
}

func (c *enumConverter) New() pref.Value {
	return c.def
}

type messageConverter struct {
	goType reflect.Type
}

func newMessageConverter(goType reflect.Type) Converter {
	return &messageConverter{goType}
}

func (c *messageConverter) PBValueOf(v reflect.Value) pref.Value {
	if v.Type() != c.goType {
		panic(fmt.Sprintf("invalid type: got %v, want %v", v.Type(), c.goType))
	}
	if m, ok := v.Interface().(pref.ProtoMessage); ok {
		return pref.ValueOf(m.ProtoReflect())
	}
	return pref.ValueOf(legacyWrapMessage(v).ProtoReflect())
}

func (c *messageConverter) GoValueOf(v pref.Value) reflect.Value {
	m := v.Message()
	var rv reflect.Value
	if u, ok := m.(Unwrapper); ok {
		rv = reflect.ValueOf(u.ProtoUnwrap())
	} else {
		rv = reflect.ValueOf(m.Interface())
	}
	if rv.Type() != c.goType {
		panic(fmt.Sprintf("invalid type: got %v, want %v", rv.Type(), c.goType))
	}
	return rv
}

func (c *messageConverter) New() pref.Value {
	return c.PBValueOf(reflect.New(c.goType.Elem()))
}