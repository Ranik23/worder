// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: wordcorrector.proto

package gen

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
	WordCorrector_CorrectWord_FullMethodName = "/wordcorrector.WordCorrector/CorrectWord"
)

// WordCorrectorClient is the client API for WordCorrector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WordCorrectorClient interface {
	CorrectWord(ctx context.Context, in *CorrectWordRequest, opts ...grpc.CallOption) (*CorrectWordResponse, error)
}

type wordCorrectorClient struct {
	cc grpc.ClientConnInterface
}

func NewWordCorrectorClient(cc grpc.ClientConnInterface) WordCorrectorClient {
	return &wordCorrectorClient{cc}
}

func (c *wordCorrectorClient) CorrectWord(ctx context.Context, in *CorrectWordRequest, opts ...grpc.CallOption) (*CorrectWordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CorrectWordResponse)
	err := c.cc.Invoke(ctx, WordCorrector_CorrectWord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WordCorrectorServer is the server API for WordCorrector service.
// All implementations must embed UnimplementedWordCorrectorServer
// for forward compatibility.
type WordCorrectorServer interface {
	CorrectWord(context.Context, *CorrectWordRequest) (*CorrectWordResponse, error)
	mustEmbedUnimplementedWordCorrectorServer()
}

// UnimplementedWordCorrectorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWordCorrectorServer struct{}

func (UnimplementedWordCorrectorServer) CorrectWord(context.Context, *CorrectWordRequest) (*CorrectWordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CorrectWord not implemented")
}
func (UnimplementedWordCorrectorServer) mustEmbedUnimplementedWordCorrectorServer() {}
func (UnimplementedWordCorrectorServer) testEmbeddedByValue()                       {}

// UnsafeWordCorrectorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WordCorrectorServer will
// result in compilation errors.
type UnsafeWordCorrectorServer interface {
	mustEmbedUnimplementedWordCorrectorServer()
}

func RegisterWordCorrectorServer(s grpc.ServiceRegistrar, srv WordCorrectorServer) {
	// If the following call pancis, it indicates UnimplementedWordCorrectorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WordCorrector_ServiceDesc, srv)
}

func _WordCorrector_CorrectWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CorrectWordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordCorrectorServer).CorrectWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WordCorrector_CorrectWord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordCorrectorServer).CorrectWord(ctx, req.(*CorrectWordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WordCorrector_ServiceDesc is the grpc.ServiceDesc for WordCorrector service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WordCorrector_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wordcorrector.WordCorrector",
	HandlerType: (*WordCorrectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CorrectWord",
			Handler:    _WordCorrector_CorrectWord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wordcorrector.proto",
}