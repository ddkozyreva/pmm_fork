// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: actions/v1/actions.proto

package actionsv1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ActionsService_GetAction_FullMethodName            = "/actions.v1.ActionsService/GetAction"
	ActionsService_StartServiceAction_FullMethodName   = "/actions.v1.ActionsService/StartServiceAction"
	ActionsService_StartPTSummaryAction_FullMethodName = "/actions.v1.ActionsService/StartPTSummaryAction"
	ActionsService_CancelAction_FullMethodName         = "/actions.v1.ActionsService/CancelAction"
)

// ActionsServiceClient is the client API for ActionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Actions service provides public Management API methods for Actions.
type ActionsServiceClient interface {
	// GetAction gets result of a given Action.
	GetAction(ctx context.Context, in *GetActionRequest, opts ...grpc.CallOption) (*GetActionResponse, error)
	// StartServiceAction starts a Service Action.
	StartServiceAction(ctx context.Context, in *StartServiceActionRequest, opts ...grpc.CallOption) (*StartServiceActionResponse, error)
	// StartPTSummaryAction starts pt-summary Node Action.
	StartPTSummaryAction(ctx context.Context, in *StartPTSummaryActionRequest, opts ...grpc.CallOption) (*StartPTSummaryActionResponse, error)
	// CancelAction stops an Action.
	CancelAction(ctx context.Context, in *CancelActionRequest, opts ...grpc.CallOption) (*CancelActionResponse, error)
}

type actionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActionsServiceClient(cc grpc.ClientConnInterface) ActionsServiceClient {
	return &actionsServiceClient{cc}
}

func (c *actionsServiceClient) GetAction(ctx context.Context, in *GetActionRequest, opts ...grpc.CallOption) (*GetActionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetActionResponse)
	err := c.cc.Invoke(ctx, ActionsService_GetAction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) StartServiceAction(ctx context.Context, in *StartServiceActionRequest, opts ...grpc.CallOption) (*StartServiceActionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartServiceActionResponse)
	err := c.cc.Invoke(ctx, ActionsService_StartServiceAction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) StartPTSummaryAction(ctx context.Context, in *StartPTSummaryActionRequest, opts ...grpc.CallOption) (*StartPTSummaryActionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartPTSummaryActionResponse)
	err := c.cc.Invoke(ctx, ActionsService_StartPTSummaryAction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsServiceClient) CancelAction(ctx context.Context, in *CancelActionRequest, opts ...grpc.CallOption) (*CancelActionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelActionResponse)
	err := c.cc.Invoke(ctx, ActionsService_CancelAction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActionsServiceServer is the server API for ActionsService service.
// All implementations must embed UnimplementedActionsServiceServer
// for forward compatibility
//
// Actions service provides public Management API methods for Actions.
type ActionsServiceServer interface {
	// GetAction gets result of a given Action.
	GetAction(context.Context, *GetActionRequest) (*GetActionResponse, error)
	// StartServiceAction starts a Service Action.
	StartServiceAction(context.Context, *StartServiceActionRequest) (*StartServiceActionResponse, error)
	// StartPTSummaryAction starts pt-summary Node Action.
	StartPTSummaryAction(context.Context, *StartPTSummaryActionRequest) (*StartPTSummaryActionResponse, error)
	// CancelAction stops an Action.
	CancelAction(context.Context, *CancelActionRequest) (*CancelActionResponse, error)
	mustEmbedUnimplementedActionsServiceServer()
}

// UnimplementedActionsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedActionsServiceServer struct{}

func (UnimplementedActionsServiceServer) GetAction(context.Context, *GetActionRequest) (*GetActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAction not implemented")
}

func (UnimplementedActionsServiceServer) StartServiceAction(context.Context, *StartServiceActionRequest) (*StartServiceActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartServiceAction not implemented")
}

func (UnimplementedActionsServiceServer) StartPTSummaryAction(context.Context, *StartPTSummaryActionRequest) (*StartPTSummaryActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPTSummaryAction not implemented")
}

func (UnimplementedActionsServiceServer) CancelAction(context.Context, *CancelActionRequest) (*CancelActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelAction not implemented")
}
func (UnimplementedActionsServiceServer) mustEmbedUnimplementedActionsServiceServer() {}

// UnsafeActionsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActionsServiceServer will
// result in compilation errors.
type UnsafeActionsServiceServer interface {
	mustEmbedUnimplementedActionsServiceServer()
}

func RegisterActionsServiceServer(s grpc.ServiceRegistrar, srv ActionsServiceServer) {
	s.RegisterService(&ActionsService_ServiceDesc, srv)
}

func _ActionsService_GetAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).GetAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_GetAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).GetAction(ctx, req.(*GetActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_StartServiceAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartServiceActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).StartServiceAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_StartServiceAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).StartServiceAction(ctx, req.(*StartServiceActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_StartPTSummaryAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartPTSummaryActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).StartPTSummaryAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_StartPTSummaryAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).StartPTSummaryAction(ctx, req.(*StartPTSummaryActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsService_CancelAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServiceServer).CancelAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionsService_CancelAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServiceServer).CancelAction(ctx, req.(*CancelActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ActionsService_ServiceDesc is the grpc.ServiceDesc for ActionsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActionsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "actions.v1.ActionsService",
	HandlerType: (*ActionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAction",
			Handler:    _ActionsService_GetAction_Handler,
		},
		{
			MethodName: "StartServiceAction",
			Handler:    _ActionsService_StartServiceAction_Handler,
		},
		{
			MethodName: "StartPTSummaryAction",
			Handler:    _ActionsService_StartPTSummaryAction_Handler,
		},
		{
			MethodName: "CancelAction",
			Handler:    _ActionsService_CancelAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "actions/v1/actions.proto",
}
