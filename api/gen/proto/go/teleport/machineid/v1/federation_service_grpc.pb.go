// Copyright 2024 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: teleport/machineid/v1/federation_service.proto

package machineidv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SPIFFEFederationService_GetSPIFFEFederation_FullMethodName    = "/teleport.machineid.v1.SPIFFEFederationService/GetSPIFFEFederation"
	SPIFFEFederationService_ListSPIFFEFederations_FullMethodName  = "/teleport.machineid.v1.SPIFFEFederationService/ListSPIFFEFederations"
	SPIFFEFederationService_DeleteSPIFFEFederation_FullMethodName = "/teleport.machineid.v1.SPIFFEFederationService/DeleteSPIFFEFederation"
	SPIFFEFederationService_CreateSPIFFEFederation_FullMethodName = "/teleport.machineid.v1.SPIFFEFederationService/CreateSPIFFEFederation"
)

// SPIFFEFederationServiceClient is the client API for SPIFFEFederationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// SPIFFEFederationService provides methods to manage SPIFFE Federations
// between trust domains.
type SPIFFEFederationServiceClient interface {
	// GetSPIFFEFederation returns a SPIFFEFederation resource by name.
	GetSPIFFEFederation(ctx context.Context, in *GetSPIFFEFederationRequest, opts ...grpc.CallOption) (*SPIFFEFederation, error)
	// ListSPIFFEFederations returns a list of SPIFFEFederation resources.
	// Follows the pagination semantics of
	// https://cloud.google.com/apis/design/design_patterns#list_pagination
	ListSPIFFEFederations(ctx context.Context, in *ListSPIFFEFederationsRequest, opts ...grpc.CallOption) (*ListSPIFFEFederationsResponse, error)
	// DeleteSPIFFEFederation deletes a SPIFFEFederation resource by name.
	DeleteSPIFFEFederation(ctx context.Context, in *DeleteSPIFFEFederationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// CreateSPIFFEFederation creates a SPIFFEFederation resource.
	CreateSPIFFEFederation(ctx context.Context, in *CreateSPIFFEFederationRequest, opts ...grpc.CallOption) (*SPIFFEFederation, error)
}

type sPIFFEFederationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSPIFFEFederationServiceClient(cc grpc.ClientConnInterface) SPIFFEFederationServiceClient {
	return &sPIFFEFederationServiceClient{cc}
}

func (c *sPIFFEFederationServiceClient) GetSPIFFEFederation(ctx context.Context, in *GetSPIFFEFederationRequest, opts ...grpc.CallOption) (*SPIFFEFederation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SPIFFEFederation)
	err := c.cc.Invoke(ctx, SPIFFEFederationService_GetSPIFFEFederation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sPIFFEFederationServiceClient) ListSPIFFEFederations(ctx context.Context, in *ListSPIFFEFederationsRequest, opts ...grpc.CallOption) (*ListSPIFFEFederationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListSPIFFEFederationsResponse)
	err := c.cc.Invoke(ctx, SPIFFEFederationService_ListSPIFFEFederations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sPIFFEFederationServiceClient) DeleteSPIFFEFederation(ctx context.Context, in *DeleteSPIFFEFederationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SPIFFEFederationService_DeleteSPIFFEFederation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sPIFFEFederationServiceClient) CreateSPIFFEFederation(ctx context.Context, in *CreateSPIFFEFederationRequest, opts ...grpc.CallOption) (*SPIFFEFederation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SPIFFEFederation)
	err := c.cc.Invoke(ctx, SPIFFEFederationService_CreateSPIFFEFederation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SPIFFEFederationServiceServer is the server API for SPIFFEFederationService service.
// All implementations must embed UnimplementedSPIFFEFederationServiceServer
// for forward compatibility.
//
// SPIFFEFederationService provides methods to manage SPIFFE Federations
// between trust domains.
type SPIFFEFederationServiceServer interface {
	// GetSPIFFEFederation returns a SPIFFEFederation resource by name.
	GetSPIFFEFederation(context.Context, *GetSPIFFEFederationRequest) (*SPIFFEFederation, error)
	// ListSPIFFEFederations returns a list of SPIFFEFederation resources.
	// Follows the pagination semantics of
	// https://cloud.google.com/apis/design/design_patterns#list_pagination
	ListSPIFFEFederations(context.Context, *ListSPIFFEFederationsRequest) (*ListSPIFFEFederationsResponse, error)
	// DeleteSPIFFEFederation deletes a SPIFFEFederation resource by name.
	DeleteSPIFFEFederation(context.Context, *DeleteSPIFFEFederationRequest) (*emptypb.Empty, error)
	// CreateSPIFFEFederation creates a SPIFFEFederation resource.
	CreateSPIFFEFederation(context.Context, *CreateSPIFFEFederationRequest) (*SPIFFEFederation, error)
	mustEmbedUnimplementedSPIFFEFederationServiceServer()
}

// UnimplementedSPIFFEFederationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSPIFFEFederationServiceServer struct{}

func (UnimplementedSPIFFEFederationServiceServer) GetSPIFFEFederation(context.Context, *GetSPIFFEFederationRequest) (*SPIFFEFederation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSPIFFEFederation not implemented")
}
func (UnimplementedSPIFFEFederationServiceServer) ListSPIFFEFederations(context.Context, *ListSPIFFEFederationsRequest) (*ListSPIFFEFederationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSPIFFEFederations not implemented")
}
func (UnimplementedSPIFFEFederationServiceServer) DeleteSPIFFEFederation(context.Context, *DeleteSPIFFEFederationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSPIFFEFederation not implemented")
}
func (UnimplementedSPIFFEFederationServiceServer) CreateSPIFFEFederation(context.Context, *CreateSPIFFEFederationRequest) (*SPIFFEFederation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSPIFFEFederation not implemented")
}
func (UnimplementedSPIFFEFederationServiceServer) mustEmbedUnimplementedSPIFFEFederationServiceServer() {
}
func (UnimplementedSPIFFEFederationServiceServer) testEmbeddedByValue() {}

// UnsafeSPIFFEFederationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SPIFFEFederationServiceServer will
// result in compilation errors.
type UnsafeSPIFFEFederationServiceServer interface {
	mustEmbedUnimplementedSPIFFEFederationServiceServer()
}

func RegisterSPIFFEFederationServiceServer(s grpc.ServiceRegistrar, srv SPIFFEFederationServiceServer) {
	// If the following call pancis, it indicates UnimplementedSPIFFEFederationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SPIFFEFederationService_ServiceDesc, srv)
}

func _SPIFFEFederationService_GetSPIFFEFederation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSPIFFEFederationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SPIFFEFederationServiceServer).GetSPIFFEFederation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SPIFFEFederationService_GetSPIFFEFederation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SPIFFEFederationServiceServer).GetSPIFFEFederation(ctx, req.(*GetSPIFFEFederationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SPIFFEFederationService_ListSPIFFEFederations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSPIFFEFederationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SPIFFEFederationServiceServer).ListSPIFFEFederations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SPIFFEFederationService_ListSPIFFEFederations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SPIFFEFederationServiceServer).ListSPIFFEFederations(ctx, req.(*ListSPIFFEFederationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SPIFFEFederationService_DeleteSPIFFEFederation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSPIFFEFederationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SPIFFEFederationServiceServer).DeleteSPIFFEFederation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SPIFFEFederationService_DeleteSPIFFEFederation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SPIFFEFederationServiceServer).DeleteSPIFFEFederation(ctx, req.(*DeleteSPIFFEFederationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SPIFFEFederationService_CreateSPIFFEFederation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSPIFFEFederationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SPIFFEFederationServiceServer).CreateSPIFFEFederation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SPIFFEFederationService_CreateSPIFFEFederation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SPIFFEFederationServiceServer).CreateSPIFFEFederation(ctx, req.(*CreateSPIFFEFederationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SPIFFEFederationService_ServiceDesc is the grpc.ServiceDesc for SPIFFEFederationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SPIFFEFederationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teleport.machineid.v1.SPIFFEFederationService",
	HandlerType: (*SPIFFEFederationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSPIFFEFederation",
			Handler:    _SPIFFEFederationService_GetSPIFFEFederation_Handler,
		},
		{
			MethodName: "ListSPIFFEFederations",
			Handler:    _SPIFFEFederationService_ListSPIFFEFederations_Handler,
		},
		{
			MethodName: "DeleteSPIFFEFederation",
			Handler:    _SPIFFEFederationService_DeleteSPIFFEFederation_Handler,
		},
		{
			MethodName: "CreateSPIFFEFederation",
			Handler:    _SPIFFEFederationService_CreateSPIFFEFederation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teleport/machineid/v1/federation_service.proto",
}
