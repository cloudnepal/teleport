// Copyright 2023 Gravitational, Inc
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
// source: teleport/kube/v1/kube_service.proto

package kubev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	KubeService_ListKubernetesResources_FullMethodName = "/teleport.kube.v1.KubeService/ListKubernetesResources"
)

// KubeServiceClient is the client API for KubeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// KubeService provides methods to list Kubernetes resources when users are not allowed
// to access the underlying cluster or resources but their `search_as_roles` allow.
type KubeServiceClient interface {
	// ListKubernetesResources lists the Kubernetes resources without leaking details.
	ListKubernetesResources(ctx context.Context, in *ListKubernetesResourcesRequest, opts ...grpc.CallOption) (*ListKubernetesResourcesResponse, error)
}

type kubeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKubeServiceClient(cc grpc.ClientConnInterface) KubeServiceClient {
	return &kubeServiceClient{cc}
}

func (c *kubeServiceClient) ListKubernetesResources(ctx context.Context, in *ListKubernetesResourcesRequest, opts ...grpc.CallOption) (*ListKubernetesResourcesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListKubernetesResourcesResponse)
	err := c.cc.Invoke(ctx, KubeService_ListKubernetesResources_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KubeServiceServer is the server API for KubeService service.
// All implementations must embed UnimplementedKubeServiceServer
// for forward compatibility.
//
// KubeService provides methods to list Kubernetes resources when users are not allowed
// to access the underlying cluster or resources but their `search_as_roles` allow.
type KubeServiceServer interface {
	// ListKubernetesResources lists the Kubernetes resources without leaking details.
	ListKubernetesResources(context.Context, *ListKubernetesResourcesRequest) (*ListKubernetesResourcesResponse, error)
	mustEmbedUnimplementedKubeServiceServer()
}

// UnimplementedKubeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedKubeServiceServer struct{}

func (UnimplementedKubeServiceServer) ListKubernetesResources(context.Context, *ListKubernetesResourcesRequest) (*ListKubernetesResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKubernetesResources not implemented")
}
func (UnimplementedKubeServiceServer) mustEmbedUnimplementedKubeServiceServer() {}
func (UnimplementedKubeServiceServer) testEmbeddedByValue()                     {}

// UnsafeKubeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KubeServiceServer will
// result in compilation errors.
type UnsafeKubeServiceServer interface {
	mustEmbedUnimplementedKubeServiceServer()
}

func RegisterKubeServiceServer(s grpc.ServiceRegistrar, srv KubeServiceServer) {
	// If the following call pancis, it indicates UnimplementedKubeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&KubeService_ServiceDesc, srv)
}

func _KubeService_ListKubernetesResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKubernetesResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubeServiceServer).ListKubernetesResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KubeService_ListKubernetesResources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubeServiceServer).ListKubernetesResources(ctx, req.(*ListKubernetesResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KubeService_ServiceDesc is the grpc.ServiceDesc for KubeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KubeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teleport.kube.v1.KubeService",
	HandlerType: (*KubeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListKubernetesResources",
			Handler:    _KubeService_ListKubernetesResources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teleport/kube/v1/kube_service.proto",
}
