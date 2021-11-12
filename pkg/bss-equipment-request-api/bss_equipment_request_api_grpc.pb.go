// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package bss_equipment_request_api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BssEquipmentRequestApiServiceClient is the client API for BssEquipmentRequestApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BssEquipmentRequestApiServiceClient interface {
	// DescribeEquipmentRequestV1 - Describe a equipment request
	DescribeEquipmentRequestV1(ctx context.Context, in *DescribeEquipmentRequestV1Request, opts ...grpc.CallOption) (*DescribeEquipmentRequestV1Response, error)
	// CreateEquipmentRequestV1 - Create a new equipment request
	CreateEquipmentRequestV1(ctx context.Context, in *CreateEquipmentRequestV1Request, opts ...grpc.CallOption) (*CreateEquipmentRequestV1Response, error)
	// ListEquipmentRequestV1 - Get list of all equipment requests
	ListEquipmentRequestV1(ctx context.Context, in *ListEquipmentRequestV1Request, opts ...grpc.CallOption) (*ListEquipmentRequestV1Response, error)
	// RemoveEquipmentRequestV1 - Remove one equipment request
	RemoveEquipmentRequestV1(ctx context.Context, in *RemoveEquipmentRequestV1Request, opts ...grpc.CallOption) (*RemoveEquipmentRequestV1Response, error)
	// UpdateEquipmentIDEquipmentRequestV1 - Update equipment id of equipment request (as a example of task4.5 "Реализовать поддержку вариаций типов событий на обновление сущности")
	UpdateEquipmentIDEquipmentRequestV1(ctx context.Context, in *UpdateEquipmentIDEquipmentRequestV1Request, opts ...grpc.CallOption) (*UpdateEquipmentIDEquipmentRequestV1Response, error)
	// UpdateStatusEquipmentRequestV1 - Update status of equipment request (as a example of task4.5 "Реализовать поддержку вариаций типов событий на обновление сущности")
	UpdateStatusEquipmentRequestV1(ctx context.Context, in *UpdateStatusEquipmentRequestV1Request, opts ...grpc.CallOption) (*UpdateStatusEquipmentRequestV1Response, error)
}

type bssEquipmentRequestApiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBssEquipmentRequestApiServiceClient(cc grpc.ClientConnInterface) BssEquipmentRequestApiServiceClient {
	return &bssEquipmentRequestApiServiceClient{cc}
}

func (c *bssEquipmentRequestApiServiceClient) DescribeEquipmentRequestV1(ctx context.Context, in *DescribeEquipmentRequestV1Request, opts ...grpc.CallOption) (*DescribeEquipmentRequestV1Response, error) {
	out := new(DescribeEquipmentRequestV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/DescribeEquipmentRequestV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bssEquipmentRequestApiServiceClient) CreateEquipmentRequestV1(ctx context.Context, in *CreateEquipmentRequestV1Request, opts ...grpc.CallOption) (*CreateEquipmentRequestV1Response, error) {
	out := new(CreateEquipmentRequestV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/CreateEquipmentRequestV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bssEquipmentRequestApiServiceClient) ListEquipmentRequestV1(ctx context.Context, in *ListEquipmentRequestV1Request, opts ...grpc.CallOption) (*ListEquipmentRequestV1Response, error) {
	out := new(ListEquipmentRequestV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/ListEquipmentRequestV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bssEquipmentRequestApiServiceClient) RemoveEquipmentRequestV1(ctx context.Context, in *RemoveEquipmentRequestV1Request, opts ...grpc.CallOption) (*RemoveEquipmentRequestV1Response, error) {
	out := new(RemoveEquipmentRequestV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/RemoveEquipmentRequestV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bssEquipmentRequestApiServiceClient) UpdateEquipmentIDEquipmentRequestV1(ctx context.Context, in *UpdateEquipmentIDEquipmentRequestV1Request, opts ...grpc.CallOption) (*UpdateEquipmentIDEquipmentRequestV1Response, error) {
	out := new(UpdateEquipmentIDEquipmentRequestV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/UpdateEquipmentIDEquipmentRequestV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bssEquipmentRequestApiServiceClient) UpdateStatusEquipmentRequestV1(ctx context.Context, in *UpdateStatusEquipmentRequestV1Request, opts ...grpc.CallOption) (*UpdateStatusEquipmentRequestV1Response, error) {
	out := new(UpdateStatusEquipmentRequestV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/UpdateStatusEquipmentRequestV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BssEquipmentRequestApiServiceServer is the server API for BssEquipmentRequestApiService service.
// All implementations must embed UnimplementedBssEquipmentRequestApiServiceServer
// for forward compatibility
type BssEquipmentRequestApiServiceServer interface {
	// DescribeEquipmentRequestV1 - Describe a equipment request
	DescribeEquipmentRequestV1(context.Context, *DescribeEquipmentRequestV1Request) (*DescribeEquipmentRequestV1Response, error)
	// CreateEquipmentRequestV1 - Create a new equipment request
	CreateEquipmentRequestV1(context.Context, *CreateEquipmentRequestV1Request) (*CreateEquipmentRequestV1Response, error)
	// ListEquipmentRequestV1 - Get list of all equipment requests
	ListEquipmentRequestV1(context.Context, *ListEquipmentRequestV1Request) (*ListEquipmentRequestV1Response, error)
	// RemoveEquipmentRequestV1 - Remove one equipment request
	RemoveEquipmentRequestV1(context.Context, *RemoveEquipmentRequestV1Request) (*RemoveEquipmentRequestV1Response, error)
	// UpdateEquipmentIDEquipmentRequestV1 - Update equipment id of equipment request (as a example of task4.5 "Реализовать поддержку вариаций типов событий на обновление сущности")
	UpdateEquipmentIDEquipmentRequestV1(context.Context, *UpdateEquipmentIDEquipmentRequestV1Request) (*UpdateEquipmentIDEquipmentRequestV1Response, error)
	// UpdateStatusEquipmentRequestV1 - Update status of equipment request (as a example of task4.5 "Реализовать поддержку вариаций типов событий на обновление сущности")
	UpdateStatusEquipmentRequestV1(context.Context, *UpdateStatusEquipmentRequestV1Request) (*UpdateStatusEquipmentRequestV1Response, error)
	mustEmbedUnimplementedBssEquipmentRequestApiServiceServer()
}

// UnimplementedBssEquipmentRequestApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBssEquipmentRequestApiServiceServer struct {
}

func (UnimplementedBssEquipmentRequestApiServiceServer) DescribeEquipmentRequestV1(context.Context, *DescribeEquipmentRequestV1Request) (*DescribeEquipmentRequestV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeEquipmentRequestV1 not implemented")
}
func (UnimplementedBssEquipmentRequestApiServiceServer) CreateEquipmentRequestV1(context.Context, *CreateEquipmentRequestV1Request) (*CreateEquipmentRequestV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEquipmentRequestV1 not implemented")
}
func (UnimplementedBssEquipmentRequestApiServiceServer) ListEquipmentRequestV1(context.Context, *ListEquipmentRequestV1Request) (*ListEquipmentRequestV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEquipmentRequestV1 not implemented")
}
func (UnimplementedBssEquipmentRequestApiServiceServer) RemoveEquipmentRequestV1(context.Context, *RemoveEquipmentRequestV1Request) (*RemoveEquipmentRequestV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveEquipmentRequestV1 not implemented")
}
func (UnimplementedBssEquipmentRequestApiServiceServer) UpdateEquipmentIDEquipmentRequestV1(context.Context, *UpdateEquipmentIDEquipmentRequestV1Request) (*UpdateEquipmentIDEquipmentRequestV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEquipmentIDEquipmentRequestV1 not implemented")
}
func (UnimplementedBssEquipmentRequestApiServiceServer) UpdateStatusEquipmentRequestV1(context.Context, *UpdateStatusEquipmentRequestV1Request) (*UpdateStatusEquipmentRequestV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatusEquipmentRequestV1 not implemented")
}
func (UnimplementedBssEquipmentRequestApiServiceServer) mustEmbedUnimplementedBssEquipmentRequestApiServiceServer() {
}

// UnsafeBssEquipmentRequestApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BssEquipmentRequestApiServiceServer will
// result in compilation errors.
type UnsafeBssEquipmentRequestApiServiceServer interface {
	mustEmbedUnimplementedBssEquipmentRequestApiServiceServer()
}

func RegisterBssEquipmentRequestApiServiceServer(s grpc.ServiceRegistrar, srv BssEquipmentRequestApiServiceServer) {
	s.RegisterService(&BssEquipmentRequestApiService_ServiceDesc, srv)
}

func _BssEquipmentRequestApiService_DescribeEquipmentRequestV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeEquipmentRequestV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BssEquipmentRequestApiServiceServer).DescribeEquipmentRequestV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/DescribeEquipmentRequestV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BssEquipmentRequestApiServiceServer).DescribeEquipmentRequestV1(ctx, req.(*DescribeEquipmentRequestV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _BssEquipmentRequestApiService_CreateEquipmentRequestV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEquipmentRequestV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BssEquipmentRequestApiServiceServer).CreateEquipmentRequestV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/CreateEquipmentRequestV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BssEquipmentRequestApiServiceServer).CreateEquipmentRequestV1(ctx, req.(*CreateEquipmentRequestV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _BssEquipmentRequestApiService_ListEquipmentRequestV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEquipmentRequestV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BssEquipmentRequestApiServiceServer).ListEquipmentRequestV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/ListEquipmentRequestV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BssEquipmentRequestApiServiceServer).ListEquipmentRequestV1(ctx, req.(*ListEquipmentRequestV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _BssEquipmentRequestApiService_RemoveEquipmentRequestV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveEquipmentRequestV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BssEquipmentRequestApiServiceServer).RemoveEquipmentRequestV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/RemoveEquipmentRequestV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BssEquipmentRequestApiServiceServer).RemoveEquipmentRequestV1(ctx, req.(*RemoveEquipmentRequestV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _BssEquipmentRequestApiService_UpdateEquipmentIDEquipmentRequestV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEquipmentIDEquipmentRequestV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BssEquipmentRequestApiServiceServer).UpdateEquipmentIDEquipmentRequestV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/UpdateEquipmentIDEquipmentRequestV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BssEquipmentRequestApiServiceServer).UpdateEquipmentIDEquipmentRequestV1(ctx, req.(*UpdateEquipmentIDEquipmentRequestV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _BssEquipmentRequestApiService_UpdateStatusEquipmentRequestV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatusEquipmentRequestV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BssEquipmentRequestApiServiceServer).UpdateStatusEquipmentRequestV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService/UpdateStatusEquipmentRequestV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BssEquipmentRequestApiServiceServer).UpdateStatusEquipmentRequestV1(ctx, req.(*UpdateStatusEquipmentRequestV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// BssEquipmentRequestApiService_ServiceDesc is the grpc.ServiceDesc for BssEquipmentRequestApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BssEquipmentRequestApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ozonmp.bss_equipment_request_api.v1.BssEquipmentRequestApiService",
	HandlerType: (*BssEquipmentRequestApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribeEquipmentRequestV1",
			Handler:    _BssEquipmentRequestApiService_DescribeEquipmentRequestV1_Handler,
		},
		{
			MethodName: "CreateEquipmentRequestV1",
			Handler:    _BssEquipmentRequestApiService_CreateEquipmentRequestV1_Handler,
		},
		{
			MethodName: "ListEquipmentRequestV1",
			Handler:    _BssEquipmentRequestApiService_ListEquipmentRequestV1_Handler,
		},
		{
			MethodName: "RemoveEquipmentRequestV1",
			Handler:    _BssEquipmentRequestApiService_RemoveEquipmentRequestV1_Handler,
		},
		{
			MethodName: "UpdateEquipmentIDEquipmentRequestV1",
			Handler:    _BssEquipmentRequestApiService_UpdateEquipmentIDEquipmentRequestV1_Handler,
		},
		{
			MethodName: "UpdateStatusEquipmentRequestV1",
			Handler:    _BssEquipmentRequestApiService_UpdateStatusEquipmentRequestV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ozonmp/bss_equipment_request_api/v1/bss_equipment_request_api.proto",
}
