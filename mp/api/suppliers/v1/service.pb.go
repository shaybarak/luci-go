// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/mp/api/suppliers/v1/service.proto

package suppliers

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

func init() {
	proto.RegisterFile("go.chromium.org/luci/mp/api/suppliers/v1/service.proto", fileDescriptor_7efc3e636b011554)
}

var fileDescriptor_7efc3e636b011554 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8f, 0x4b, 0xca, 0xc2, 0x30,
	0x14, 0x85, 0xf9, 0xf9, 0xa1, 0xd0, 0x0c, 0x23, 0x38, 0xa8, 0x08, 0xae, 0x20, 0x17, 0x15, 0x44,
	0x47, 0x82, 0x0f, 0x3a, 0xd7, 0x15, 0xb4, 0xe5, 0x1a, 0x03, 0x09, 0x89, 0x79, 0x14, 0x5c, 0xa3,
	0x9b, 0x12, 0x1b, 0xa2, 0xa5, 0xa3, 0x0e, 0x73, 0xce, 0x97, 0xef, 0x72, 0xc8, 0x86, 0x6b, 0xd6,
	0xdc, 0xad, 0x56, 0x22, 0x28, 0xa6, 0x2d, 0x07, 0x19, 0x1a, 0x01, 0xca, 0x40, 0x65, 0x04, 0xb8,
	0x60, 0x8c, 0x14, 0x68, 0x1d, 0xb4, 0x4b, 0x70, 0x68, 0x5b, 0xd1, 0x20, 0x33, 0x56, 0x7b, 0x4d,
	0xf3, 0x6f, 0x57, 0xcc, 0xb8, 0xd6, 0x5c, 0x22, 0x74, 0x45, 0x1d, 0x6e, 0x80, 0xca, 0xf8, 0x67,
	0xe4, 0x8a, 0xed, 0x78, 0x7f, 0x7a, 0xc4, 0x9f, 0xab, 0xd7, 0x1f, 0xc9, 0xaf, 0x29, 0xa3, 0x7b,
	0x92, 0x1d, 0x2d, 0x56, 0x1e, 0xe9, 0x82, 0xfd, 0xc8, 0x18, 0x25, 0xea, 0x82, 0x8f, 0x80, 0xce,
	0x17, 0x93, 0x1e, 0x91, 0x3a, 0x7a, 0x20, 0xd9, 0x09, 0x25, 0x0e, 0x04, 0x31, 0x1a, 0x0a, 0xa6,
	0x2c, 0x4e, 0x62, 0x69, 0x12, 0x3b, 0x7f, 0x26, 0xd1, 0x1d, 0xf9, 0x2f, 0xd1, 0xd3, 0x79, 0x4f,
	0x50, 0xa2, 0x1f, 0x73, 0xbe, 0xce, 0x3a, 0xd5, 0xfa, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x78, 0xff,
	0xf9, 0xec, 0x70, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SuppliersClient is the client API for Suppliers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SuppliersClient interface {
	// Creates a new supplier.
	Create(ctx context.Context, in *CreateSupplierRequest, opts ...grpc.CallOption) (*Supplier, error)
	// Deletes an existing supplier.
	Delete(ctx context.Context, in *DeleteSupplierRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Gets an existing supplier.
	Get(ctx context.Context, in *GetSupplierRequest, opts ...grpc.CallOption) (*Supplier, error)
}
type suppliersPRPCClient struct {
	client *prpc.Client
}

func NewSuppliersPRPCClient(client *prpc.Client) SuppliersClient {
	return &suppliersPRPCClient{client}
}

func (c *suppliersPRPCClient) Create(ctx context.Context, in *CreateSupplierRequest, opts ...grpc.CallOption) (*Supplier, error) {
	out := new(Supplier)
	err := c.client.Call(ctx, "suppliers.Suppliers", "Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *suppliersPRPCClient) Delete(ctx context.Context, in *DeleteSupplierRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.client.Call(ctx, "suppliers.Suppliers", "Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *suppliersPRPCClient) Get(ctx context.Context, in *GetSupplierRequest, opts ...grpc.CallOption) (*Supplier, error) {
	out := new(Supplier)
	err := c.client.Call(ctx, "suppliers.Suppliers", "Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type suppliersClient struct {
	cc *grpc.ClientConn
}

func NewSuppliersClient(cc *grpc.ClientConn) SuppliersClient {
	return &suppliersClient{cc}
}

func (c *suppliersClient) Create(ctx context.Context, in *CreateSupplierRequest, opts ...grpc.CallOption) (*Supplier, error) {
	out := new(Supplier)
	err := c.cc.Invoke(ctx, "/suppliers.Suppliers/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *suppliersClient) Delete(ctx context.Context, in *DeleteSupplierRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/suppliers.Suppliers/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *suppliersClient) Get(ctx context.Context, in *GetSupplierRequest, opts ...grpc.CallOption) (*Supplier, error) {
	out := new(Supplier)
	err := c.cc.Invoke(ctx, "/suppliers.Suppliers/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SuppliersServer is the server API for Suppliers service.
type SuppliersServer interface {
	// Creates a new supplier.
	Create(context.Context, *CreateSupplierRequest) (*Supplier, error)
	// Deletes an existing supplier.
	Delete(context.Context, *DeleteSupplierRequest) (*empty.Empty, error)
	// Gets an existing supplier.
	Get(context.Context, *GetSupplierRequest) (*Supplier, error)
}

func RegisterSuppliersServer(s prpc.Registrar, srv SuppliersServer) {
	s.RegisterService(&_Suppliers_serviceDesc, srv)
}

func _Suppliers_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSupplierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuppliersServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/suppliers.Suppliers/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuppliersServer).Create(ctx, req.(*CreateSupplierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Suppliers_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSupplierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuppliersServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/suppliers.Suppliers/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuppliersServer).Delete(ctx, req.(*DeleteSupplierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Suppliers_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSupplierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuppliersServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/suppliers.Suppliers/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuppliersServer).Get(ctx, req.(*GetSupplierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Suppliers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "suppliers.Suppliers",
	HandlerType: (*SuppliersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Suppliers_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Suppliers_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Suppliers_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/mp/api/suppliers/v1/service.proto",
}