// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package examplepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// NonStandardServiceClient is the client API for NonStandardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NonStandardServiceClient interface {
	// Apply field mask to empty NonStandardMessage and return result.
	Update(ctx context.Context, in *NonStandardUpdateRequest, opts ...grpc.CallOption) (*NonStandardMessage, error)
	// Apply field mask to empty NonStandardMessageWithJSONNames and return result.
	UpdateWithJSONNames(ctx context.Context, in *NonStandardWithJSONNamesUpdateRequest, opts ...grpc.CallOption) (*NonStandardMessageWithJSONNames, error)
}

type nonStandardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNonStandardServiceClient(cc grpc.ClientConnInterface) NonStandardServiceClient {
	return &nonStandardServiceClient{cc}
}

func (c *nonStandardServiceClient) Update(ctx context.Context, in *NonStandardUpdateRequest, opts ...grpc.CallOption) (*NonStandardMessage, error) {
	out := new(NonStandardMessage)
	err := c.cc.Invoke(ctx, "/grpc.gateway.runtime.internal.examplepb.NonStandardService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nonStandardServiceClient) UpdateWithJSONNames(ctx context.Context, in *NonStandardWithJSONNamesUpdateRequest, opts ...grpc.CallOption) (*NonStandardMessageWithJSONNames, error) {
	out := new(NonStandardMessageWithJSONNames)
	err := c.cc.Invoke(ctx, "/grpc.gateway.runtime.internal.examplepb.NonStandardService/UpdateWithJSONNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NonStandardServiceServer is the server API for NonStandardService service.
// All implementations should embed UnimplementedNonStandardServiceServer
// for forward compatibility
type NonStandardServiceServer interface {
	// Apply field mask to empty NonStandardMessage and return result.
	Update(context.Context, *NonStandardUpdateRequest) (*NonStandardMessage, error)
	// Apply field mask to empty NonStandardMessageWithJSONNames and return result.
	UpdateWithJSONNames(context.Context, *NonStandardWithJSONNamesUpdateRequest) (*NonStandardMessageWithJSONNames, error)
}

// UnimplementedNonStandardServiceServer should be embedded to have forward compatible implementations.
type UnimplementedNonStandardServiceServer struct {
}

func (UnimplementedNonStandardServiceServer) Update(context.Context, *NonStandardUpdateRequest) (*NonStandardMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedNonStandardServiceServer) UpdateWithJSONNames(context.Context, *NonStandardWithJSONNamesUpdateRequest) (*NonStandardMessageWithJSONNames, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWithJSONNames not implemented")
}

// UnsafeNonStandardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NonStandardServiceServer will
// result in compilation errors.
type UnsafeNonStandardServiceServer interface {
	mustEmbedUnimplementedNonStandardServiceServer()
}

func RegisterNonStandardServiceServer(s *grpc.Server, srv NonStandardServiceServer) {
	s.RegisterService(&_NonStandardService_serviceDesc, srv)
}

func _NonStandardService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NonStandardUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NonStandardServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.runtime.internal.examplepb.NonStandardService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NonStandardServiceServer).Update(ctx, req.(*NonStandardUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NonStandardService_UpdateWithJSONNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NonStandardWithJSONNamesUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NonStandardServiceServer).UpdateWithJSONNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.runtime.internal.examplepb.NonStandardService/UpdateWithJSONNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NonStandardServiceServer).UpdateWithJSONNames(ctx, req.(*NonStandardWithJSONNamesUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NonStandardService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.runtime.internal.examplepb.NonStandardService",
	HandlerType: (*NonStandardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _NonStandardService_Update_Handler,
		},
		{
			MethodName: "UpdateWithJSONNames",
			Handler:    _NonStandardService_UpdateWithJSONNames_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "runtime/internal/examplepb/non_standard_names.proto",
}
