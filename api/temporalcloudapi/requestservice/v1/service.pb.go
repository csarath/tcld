// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/requestservice/v1/service.proto

package requestservice

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("api/requestservice/v1/service.proto", fileDescriptor_a6f283241a216f38)
}

var fileDescriptor_a6f283241a216f38 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0x2c, 0xc8, 0xd4,
	0x2f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x29, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2f,
	0x33, 0xd4, 0x87, 0x32, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x44, 0x13, 0x0b, 0x32, 0xf5,
	0x50, 0x15, 0xe9, 0x95, 0x19, 0x4a, 0xe9, 0x60, 0xd7, 0x0b, 0x15, 0x89, 0x2f, 0x4a, 0x2d, 0x2e,
	0xc8, 0xcf, 0x2b, 0x86, 0x1a, 0x62, 0xd4, 0xce, 0xc8, 0xc5, 0x17, 0x04, 0x91, 0x0a, 0x86, 0x28,
	0x16, 0x2a, 0xe5, 0x12, 0x70, 0x4f, 0x2d, 0x81, 0x09, 0x96, 0x24, 0x96, 0x94, 0x16, 0x0b, 0xe9,
	0xe9, 0x61, 0xb5, 0x4c, 0x0f, 0x5d, 0x21, 0x94, 0x23, 0xa5, 0x4f, 0xb4, 0x7a, 0x88, 0x6b, 0x94,
	0x18, 0x9c, 0xd2, 0x2e, 0x3c, 0x94, 0x63, 0xb8, 0xf1, 0x50, 0x8e, 0xe1, 0xc3, 0x43, 0x39, 0xc6,
	0x86, 0x47, 0x72, 0x8c, 0x2b, 0x1e, 0xc9, 0x31, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c,
	0xe3, 0x83, 0x47, 0x72, 0x8c, 0x2f, 0x1e, 0xc9, 0x31, 0x7c, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1,
	0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x19, 0x94, 0xe4, 0x16, 0x14,
	0xe5, 0xe8, 0x25, 0xe7, 0xe4, 0x97, 0xa6, 0xe8, 0x63, 0xf5, 0xb9, 0x35, 0xaa, 0x48, 0x12, 0x1b,
	0xd8, 0xe3, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x9d, 0xf4, 0x8d, 0x64, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RequestServiceClient is the client API for RequestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RequestServiceClient interface {
	// GetRequestStatus get the status of a request.
	GetRequestStatus(ctx context.Context, in *GetRequestStatusRequest, opts ...grpc.CallOption) (*GetRequestStatusResponse, error)
}

type requestServiceClient struct {
	cc *grpc.ClientConn
}

func NewRequestServiceClient(cc *grpc.ClientConn) RequestServiceClient {
	return &requestServiceClient{cc}
}

func (c *requestServiceClient) GetRequestStatus(ctx context.Context, in *GetRequestStatusRequest, opts ...grpc.CallOption) (*GetRequestStatusResponse, error) {
	out := new(GetRequestStatusResponse)
	err := c.cc.Invoke(ctx, "/api.requestservice.v1.RequestService/GetRequestStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RequestServiceServer is the server API for RequestService service.
type RequestServiceServer interface {
	// GetRequestStatus get the status of a request.
	GetRequestStatus(context.Context, *GetRequestStatusRequest) (*GetRequestStatusResponse, error)
}

// UnimplementedRequestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRequestServiceServer struct {
}

func (*UnimplementedRequestServiceServer) GetRequestStatus(ctx context.Context, req *GetRequestStatusRequest) (*GetRequestStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequestStatus not implemented")
}

func RegisterRequestServiceServer(s *grpc.Server, srv RequestServiceServer) {
	s.RegisterService(&_RequestService_serviceDesc, srv)
}

func _RequestService_GetRequestStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequestStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).GetRequestStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.requestservice.v1.RequestService/GetRequestStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).GetRequestStatus(ctx, req.(*GetRequestStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RequestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.requestservice.v1.RequestService",
	HandlerType: (*RequestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRequestStatus",
			Handler:    _RequestService_GetRequestStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/requestservice/v1/service.proto",
}