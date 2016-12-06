// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	VersionRequest
	VersionReply
	StartRequest
	StartReply
	StopRequest
	StopReply
	ImageSignatureRequest
	ImageSignatureReply
	UpdateImageRequest
	UpdateImageReply
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type VersionRequest struct {
}

func (m *VersionRequest) Reset()                    { *m = VersionRequest{} }
func (m *VersionRequest) String() string            { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()               {}
func (*VersionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Ukd server version.
type VersionReply struct {
	Major int32 `protobuf:"varint,1,opt,name=major" json:"major,omitempty"`
	Minor int32 `protobuf:"varint,2,opt,name=minor" json:"minor,omitempty"`
}

func (m *VersionReply) Reset()                    { *m = VersionReply{} }
func (m *VersionReply) String() string            { return proto.CompactTextString(m) }
func (*VersionReply) ProtoMessage()               {}
func (*VersionReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Request message containing image name and image location.
type StartRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Visor    string `protobuf:"bytes,2,opt,name=visor" json:"visor,omitempty"`
	Location string `protobuf:"bytes,3,opt,name=location" json:"location,omitempty"`
}

func (m *StartRequest) Reset()                    { *m = StartRequest{} }
func (m *StartRequest) String() string            { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()               {}
func (*StartRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Response message signalling result of start attempt.
type StartReply struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Ip      string `protobuf:"bytes,2,opt,name=ip" json:"ip,omitempty"`
	Info    string `protobuf:"bytes,3,opt,name=info" json:"info,omitempty"`
}

func (m *StartReply) Reset()                    { *m = StartReply{} }
func (m *StartReply) String() string            { return proto.CompactTextString(m) }
func (*StartReply) ProtoMessage()               {}
func (*StartReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// Request message containing the image name.
type StopRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *StopRequest) Reset()                    { *m = StopRequest{} }
func (m *StopRequest) String() string            { return proto.CompactTextString(m) }
func (*StopRequest) ProtoMessage()               {}
func (*StopRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// Response message signalling result of stop attempt.
type StopReply struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Info    string `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
}

func (m *StopReply) Reset()                    { *m = StopReply{} }
func (m *StopReply) String() string            { return proto.CompactTextString(m) }
func (*StopReply) ProtoMessage()               {}
func (*StopReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ImageSignatureRequest struct {
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
}

func (m *ImageSignatureRequest) Reset()                    { *m = ImageSignatureRequest{} }
func (m *ImageSignatureRequest) String() string            { return proto.CompactTextString(m) }
func (*ImageSignatureRequest) ProtoMessage()               {}
func (*ImageSignatureRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type ImageSignatureReply struct {
	Success   bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Info      string `protobuf:"bytes,3,opt,name=info" json:"info,omitempty"`
}

func (m *ImageSignatureReply) Reset()                    { *m = ImageSignatureReply{} }
func (m *ImageSignatureReply) String() string            { return proto.CompactTextString(m) }
func (*ImageSignatureReply) ProtoMessage()               {}
func (*ImageSignatureReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type UpdateImageRequest struct {
	Base    string `protobuf:"bytes,1,opt,name=base" json:"base,omitempty"`
	Basesig []byte `protobuf:"bytes,2,opt,name=basesig,proto3" json:"basesig,omitempty"`
	Newsig  []byte `protobuf:"bytes,3,opt,name=newsig,proto3" json:"newsig,omitempty"`
	Diff    []byte `protobuf:"bytes,4,opt,name=diff,proto3" json:"diff,omitempty"`
}

func (m *UpdateImageRequest) Reset()                    { *m = UpdateImageRequest{} }
func (m *UpdateImageRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateImageRequest) ProtoMessage()               {}
func (*UpdateImageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type UpdateImageReply struct {
	Success         bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	StagedImagePath string `protobuf:"bytes,2,opt,name=stagedImagePath" json:"stagedImagePath,omitempty"`
	Info            string `protobuf:"bytes,3,opt,name=info" json:"info,omitempty"`
}

func (m *UpdateImageReply) Reset()                    { *m = UpdateImageReply{} }
func (m *UpdateImageReply) String() string            { return proto.CompactTextString(m) }
func (*UpdateImageReply) ProtoMessage()               {}
func (*UpdateImageReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func init() {
	proto.RegisterType((*VersionRequest)(nil), "VersionRequest")
	proto.RegisterType((*VersionReply)(nil), "VersionReply")
	proto.RegisterType((*StartRequest)(nil), "StartRequest")
	proto.RegisterType((*StartReply)(nil), "StartReply")
	proto.RegisterType((*StopRequest)(nil), "StopRequest")
	proto.RegisterType((*StopReply)(nil), "StopReply")
	proto.RegisterType((*ImageSignatureRequest)(nil), "ImageSignatureRequest")
	proto.RegisterType((*ImageSignatureReply)(nil), "ImageSignatureReply")
	proto.RegisterType((*UpdateImageRequest)(nil), "UpdateImageRequest")
	proto.RegisterType((*UpdateImageReply)(nil), "UpdateImageReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Ukd service

type UkdClient interface {
	// Get Server Version.
	GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionReply, error)
	// Start a Unikernel.
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartReply, error)
	// Stop a Unikernel.
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopReply, error)
	// Get image signature.
	GetImageSignature(ctx context.Context, in *ImageSignatureRequest, opts ...grpc.CallOption) (*ImageSignatureReply, error)
	// Update a Unikernel on-disk image.
	UpdateImage(ctx context.Context, in *UpdateImageRequest, opts ...grpc.CallOption) (*UpdateImageReply, error)
}

type ukdClient struct {
	cc *grpc.ClientConn
}

func NewUkdClient(cc *grpc.ClientConn) UkdClient {
	return &ukdClient{cc}
}

func (c *ukdClient) GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionReply, error) {
	out := new(VersionReply)
	err := grpc.Invoke(ctx, "/Ukd/GetVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ukdClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartReply, error) {
	out := new(StartReply)
	err := grpc.Invoke(ctx, "/Ukd/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ukdClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopReply, error) {
	out := new(StopReply)
	err := grpc.Invoke(ctx, "/Ukd/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ukdClient) GetImageSignature(ctx context.Context, in *ImageSignatureRequest, opts ...grpc.CallOption) (*ImageSignatureReply, error) {
	out := new(ImageSignatureReply)
	err := grpc.Invoke(ctx, "/Ukd/GetImageSignature", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ukdClient) UpdateImage(ctx context.Context, in *UpdateImageRequest, opts ...grpc.CallOption) (*UpdateImageReply, error) {
	out := new(UpdateImageReply)
	err := grpc.Invoke(ctx, "/Ukd/UpdateImage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Ukd service

type UkdServer interface {
	// Get Server Version.
	GetVersion(context.Context, *VersionRequest) (*VersionReply, error)
	// Start a Unikernel.
	Start(context.Context, *StartRequest) (*StartReply, error)
	// Stop a Unikernel.
	Stop(context.Context, *StopRequest) (*StopReply, error)
	// Get image signature.
	GetImageSignature(context.Context, *ImageSignatureRequest) (*ImageSignatureReply, error)
	// Update a Unikernel on-disk image.
	UpdateImage(context.Context, *UpdateImageRequest) (*UpdateImageReply, error)
}

func RegisterUkdServer(s *grpc.Server, srv UkdServer) {
	s.RegisterService(&_Ukd_serviceDesc, srv)
}

func _Ukd_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UkdServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ukd/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UkdServer).GetVersion(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ukd_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UkdServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ukd/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UkdServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ukd_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UkdServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ukd/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UkdServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ukd_GetImageSignature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageSignatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UkdServer).GetImageSignature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ukd/GetImageSignature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UkdServer).GetImageSignature(ctx, req.(*ImageSignatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ukd_UpdateImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UkdServer).UpdateImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ukd/UpdateImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UkdServer).UpdateImage(ctx, req.(*UpdateImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ukd_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Ukd",
	HandlerType: (*UkdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _Ukd_GetVersion_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Ukd_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Ukd_Stop_Handler,
		},
		{
			MethodName: "GetImageSignature",
			Handler:    _Ukd_GetImageSignature_Handler,
		},
		{
			MethodName: "UpdateImage",
			Handler:    _Ukd_UpdateImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x53, 0x4d, 0x8f, 0xda, 0x30,
	0x14, 0xe4, 0xbb, 0xe4, 0x11, 0xbe, 0x0c, 0x45, 0x51, 0xd4, 0x43, 0x6b, 0xa9, 0x12, 0x52, 0x25,
	0x1f, 0xda, 0x43, 0xd5, 0x5e, 0x7b, 0xa8, 0xda, 0x53, 0x15, 0x4a, 0xef, 0x86, 0x18, 0x6a, 0x16,
	0x92, 0x6c, 0x6c, 0x76, 0xb5, 0x7f, 0x7e, 0xb5, 0xb6, 0xe3, 0x84, 0xc0, 0x46, 0x39, 0xe5, 0xbd,
	0xf1, 0x78, 0xc6, 0x7e, 0x9e, 0x80, 0x43, 0x13, 0x4e, 0x92, 0x34, 0x96, 0x31, 0x9e, 0xc0, 0xe8,
	0x1f, 0x4b, 0x05, 0x8f, 0xa3, 0x80, 0xdd, 0x9f, 0x99, 0x90, 0xf8, 0x3b, 0xb8, 0x05, 0x92, 0x1c,
	0x9f, 0xd0, 0x1c, 0xba, 0x27, 0x7a, 0x88, 0x53, 0xaf, 0xf9, 0xbe, 0xb9, 0xec, 0x06, 0x59, 0x63,
	0x50, 0x1e, 0x29, 0xb4, 0x65, 0x51, 0xdd, 0xe0, 0xbf, 0xe0, 0xae, 0x24, 0x4d, 0xa5, 0xd5, 0x42,
	0x08, 0x3a, 0x11, 0x3d, 0x31, 0xb3, 0xd5, 0x09, 0x4c, 0xad, 0x77, 0x3e, 0x70, 0x61, 0x77, 0x3a,
	0x41, 0xd6, 0x20, 0x1f, 0xfa, 0xc7, 0x78, 0x4b, 0xa5, 0xb2, 0xf5, 0xda, 0x66, 0xa1, 0xe8, 0xf1,
	0x6f, 0x00, 0xab, 0xaa, 0xcf, 0xe3, 0xc1, 0x1b, 0x71, 0xde, 0x6e, 0x99, 0x10, 0x46, 0xb6, 0x1f,
	0xe4, 0x2d, 0x1a, 0x41, 0x8b, 0x27, 0x56, 0x56, 0x55, 0xda, 0x9d, 0x47, 0xbb, 0xd8, 0xea, 0x99,
	0x1a, 0x7f, 0x80, 0xc1, 0x4a, 0xc6, 0x49, 0xcd, 0x01, 0xf1, 0x37, 0x70, 0x32, 0x4a, 0xbd, 0x5b,
	0xae, 0xde, 0x2a, 0xa9, 0x7f, 0x82, 0xb7, 0xbf, 0x4e, 0x74, 0xcf, 0x56, 0x7c, 0x1f, 0x51, 0x79,
	0x4e, 0x59, 0xc9, 0x27, 0xa1, 0xf2, 0x7f, 0xee, 0xa3, 0x6b, 0x4c, 0x61, 0x76, 0x4b, 0xae, 0x77,
	0x7c, 0x07, 0x8e, 0xc8, 0xb9, 0xc6, 0xd6, 0x0d, 0x2e, 0x40, 0xe5, 0x6d, 0x23, 0x40, 0xeb, 0x24,
	0xa4, 0x92, 0x19, 0xa3, 0xd2, 0x61, 0x36, 0x54, 0x14, 0x97, 0xd6, 0xb5, 0x76, 0xd5, 0x5f, 0x25,
	0x67, 0x95, 0xf3, 0x16, 0x2d, 0xa0, 0x17, 0xb1, 0x47, 0xbd, 0xd0, 0x36, 0x0b, 0xb6, 0xd3, 0x2a,
	0x21, 0xdf, 0xed, 0xbc, 0x8e, 0x41, 0x4d, 0x8d, 0x0f, 0x30, 0xb9, 0xf2, 0xab, 0xbf, 0xcf, 0x12,
	0xc6, 0x42, 0x2a, 0x5e, 0x68, 0xd8, 0x7f, 0xf4, 0x7c, 0xb2, 0x61, 0xde, 0xc2, 0x55, 0x77, 0xfb,
	0xfc, 0xdc, 0x84, 0xf6, 0xfa, 0x2e, 0x44, 0x04, 0xe0, 0x27, 0x93, 0x36, 0xb2, 0x68, 0x4c, 0xae,
	0xe3, 0xec, 0x0f, 0x49, 0x39, 0xcd, 0xb8, 0x81, 0x3e, 0x42, 0xd7, 0xa4, 0x09, 0x0d, 0x49, 0x39,
	0xab, 0xfe, 0x80, 0x5c, 0x42, 0xa6, 0x68, 0x18, 0x3a, 0x3a, 0x05, 0xc8, 0x25, 0xa5, 0xbc, 0xf8,
	0x40, 0x8a, 0x68, 0x28, 0xce, 0x0f, 0x98, 0x2a, 0xeb, 0xeb, 0x47, 0x44, 0x0b, 0x52, 0x19, 0x01,
	0x7f, 0x4e, 0x2a, 0x5e, 0x5b, 0x89, 0x7c, 0x85, 0x41, 0x69, 0x66, 0x68, 0x46, 0x5e, 0xbf, 0x98,
	0x3f, 0x25, 0xb7, 0x63, 0xc5, 0x8d, 0x4d, 0xcf, 0xfc, 0xc1, 0x5f, 0x5e, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x59, 0xf2, 0xf0, 0xd8, 0xce, 0x03, 0x00, 0x00,
}
