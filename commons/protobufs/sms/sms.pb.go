// Code generated by protoc-gen-go.
// source: sms/sms.proto
// DO NOT EDIT!

/*
Package sms is a generated protocol buffer package.

It is generated from these files:
	sms/sms.proto

It has these top-level messages:
	SendInfo
	Status
*/
package sms

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

type SendInfo_Type int32

const (
	SendInfo_DEFAUTL         SendInfo_Type = 0
	SendInfo_REGISTERCODE    SendInfo_Type = 1
	SendInfo_RECOVERPASSWORD SendInfo_Type = 2
)

var SendInfo_Type_name = map[int32]string{
	0: "DEFAUTL",
	1: "REGISTERCODE",
	2: "RECOVERPASSWORD",
}
var SendInfo_Type_value = map[string]int32{
	"DEFAUTL":         0,
	"REGISTERCODE":    1,
	"RECOVERPASSWORD": 2,
}

func (x SendInfo_Type) String() string {
	return proto.EnumName(SendInfo_Type_name, int32(x))
}
func (SendInfo_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type SendInfo struct {
	Phone   string        `protobuf:"bytes,1,opt,name=phone" json:"phone,omitempty"`
	Type    SendInfo_Type `protobuf:"varint,2,opt,name=type,enum=sms.SendInfo_Type" json:"type,omitempty"`
	Content string        `protobuf:"bytes,3,opt,name=content" json:"content,omitempty"`
}

func (m *SendInfo) Reset()                    { *m = SendInfo{} }
func (m *SendInfo) String() string            { return proto.CompactTextString(m) }
func (*SendInfo) ProtoMessage()               {}
func (*SendInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Status struct {
	Status bool `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*SendInfo)(nil), "sms.SendInfo")
	proto.RegisterType((*Status)(nil), "sms.Status")
	proto.RegisterEnum("sms.SendInfo_Type", SendInfo_Type_name, SendInfo_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for SMS service

type SMSClient interface {
	Send(ctx context.Context, in *SendInfo, opts ...grpc.CallOption) (*Status, error)
}

type sMSClient struct {
	cc *grpc.ClientConn
}

func NewSMSClient(cc *grpc.ClientConn) SMSClient {
	return &sMSClient{cc}
}

func (c *sMSClient) Send(ctx context.Context, in *SendInfo, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/sms.SMS/Send", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SMS service

type SMSServer interface {
	Send(context.Context, *SendInfo) (*Status, error)
}

func RegisterSMSServer(s *grpc.Server, srv SMSServer) {
	s.RegisterService(&_SMS_serviceDesc, srv)
}

func _SMS_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SMSServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sms.SMS/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SMSServer).Send(ctx, req.(*SendInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _SMS_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sms.SMS",
	HandlerType: (*SMSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _SMS_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("sms/sms.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x8f, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0xbb, 0x4d, 0x4c, 0xeb, 0xd4, 0x6a, 0x18, 0x45, 0x82, 0xa7, 0xb0, 0x88, 0x14, 0x84,
	0x08, 0xf5, 0xe6, 0xad, 0x34, 0xab, 0x14, 0x94, 0xc8, 0x6c, 0xd4, 0xb3, 0x1f, 0x2b, 0x5e, 0xb2,
	0xbb, 0x38, 0xeb, 0xa1, 0x7f, 0xc6, 0xdf, 0x2a, 0xd9, 0xda, 0x43, 0x6f, 0xf3, 0x30, 0xef, 0x3b,
	0x3c, 0x03, 0x53, 0xee, 0xf8, 0x8a, 0x3b, 0xae, 0xfc, 0xb7, 0x0b, 0x0e, 0x13, 0xee, 0x58, 0xfe,
	0x0a, 0x18, 0x6b, 0x63, 0x3f, 0x56, 0xf6, 0xd3, 0xe1, 0x09, 0xec, 0xf9, 0x2f, 0x67, 0x4d, 0x21,
	0x4a, 0x31, 0xdb, 0xa7, 0x0d, 0xe0, 0x05, 0xa4, 0x61, 0xed, 0x4d, 0x31, 0x2c, 0xc5, 0xec, 0x70,
	0x8e, 0x55, 0x7f, 0x61, 0x5b, 0xa9, 0xda, 0xb5, 0x37, 0x14, 0xf7, 0x58, 0xc0, 0xe8, 0xdd, 0xd9,
	0x60, 0x6c, 0x28, 0x92, 0xd8, 0xdf, 0xa2, 0xbc, 0x81, 0xb4, 0xcf, 0xe1, 0x04, 0x46, 0xb5, 0xba,
	0x5d, 0x3c, 0xb5, 0xf7, 0xf9, 0x00, 0x73, 0x38, 0x20, 0x75, 0xb7, 0xd2, 0xad, 0xa2, 0x65, 0x53,
	0xab, 0x5c, 0xe0, 0x31, 0x1c, 0x91, 0x5a, 0x36, 0xcf, 0x8a, 0x1e, 0x17, 0x5a, 0xbf, 0x34, 0x54,
	0xe7, 0x43, 0x59, 0x42, 0xa6, 0xc3, 0x6b, 0xf8, 0x61, 0x3c, 0x85, 0x8c, 0xe3, 0x14, 0xf5, 0xc6,
	0xf4, 0x4f, 0xf3, 0x4b, 0x48, 0xf4, 0x83, 0xc6, 0x73, 0x48, 0x7b, 0x2b, 0x9c, 0xee, 0x08, 0x9e,
	0x4d, 0x36, 0x18, 0xa3, 0x72, 0xf0, 0x96, 0xc5, 0xdf, 0xaf, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xe4, 0xfc, 0x9d, 0xc6, 0x0c, 0x01, 0x00, 0x00,
}