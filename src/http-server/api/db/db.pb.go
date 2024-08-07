// Code generated by protoc-gen-go. DO NOT EDIT.
// source: http-server/pb/db.proto

package db // import "dmenca/go-learning/http-server/api/db"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type DBGetRequest struct {
	DbId                 string   `protobuf:"bytes,1,opt,name=db_id,json=dbId,proto3" json:"db_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DBGetRequest) Reset()         { *m = DBGetRequest{} }
func (m *DBGetRequest) String() string { return proto.CompactTextString(m) }
func (*DBGetRequest) ProtoMessage()    {}
func (*DBGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_db_b7df01dce7a20228, []int{0}
}
func (m *DBGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DBGetRequest.Unmarshal(m, b)
}
func (m *DBGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DBGetRequest.Marshal(b, m, deterministic)
}
func (dst *DBGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DBGetRequest.Merge(dst, src)
}
func (m *DBGetRequest) XXX_Size() int {
	return xxx_messageInfo_DBGetRequest.Size(m)
}
func (m *DBGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DBGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DBGetRequest proto.InternalMessageInfo

func (m *DBGetRequest) GetDbId() string {
	if m != nil {
		return m.DbId
	}
	return ""
}

type DBGetResponse struct {
	DbInfo               *DBInfo  `protobuf:"bytes,1,opt,name=db_info,json=dbInfo,proto3" json:"db_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DBGetResponse) Reset()         { *m = DBGetResponse{} }
func (m *DBGetResponse) String() string { return proto.CompactTextString(m) }
func (*DBGetResponse) ProtoMessage()    {}
func (*DBGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_db_b7df01dce7a20228, []int{1}
}
func (m *DBGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DBGetResponse.Unmarshal(m, b)
}
func (m *DBGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DBGetResponse.Marshal(b, m, deterministic)
}
func (dst *DBGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DBGetResponse.Merge(dst, src)
}
func (m *DBGetResponse) XXX_Size() int {
	return xxx_messageInfo_DBGetResponse.Size(m)
}
func (m *DBGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DBGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DBGetResponse proto.InternalMessageInfo

func (m *DBGetResponse) GetDbInfo() *DBInfo {
	if m != nil {
		return m.DbInfo
	}
	return nil
}

type DBInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DbId                 string   `protobuf:"bytes,2,opt,name=db_id,json=dbId,proto3" json:"db_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DBInfo) Reset()         { *m = DBInfo{} }
func (m *DBInfo) String() string { return proto.CompactTextString(m) }
func (*DBInfo) ProtoMessage()    {}
func (*DBInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_db_b7df01dce7a20228, []int{2}
}
func (m *DBInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DBInfo.Unmarshal(m, b)
}
func (m *DBInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DBInfo.Marshal(b, m, deterministic)
}
func (dst *DBInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DBInfo.Merge(dst, src)
}
func (m *DBInfo) XXX_Size() int {
	return xxx_messageInfo_DBInfo.Size(m)
}
func (m *DBInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DBInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DBInfo proto.InternalMessageInfo

func (m *DBInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DBInfo) GetDbId() string {
	if m != nil {
		return m.DbId
	}
	return ""
}

func init() {
	proto.RegisterType((*DBGetRequest)(nil), "dmenca.go_learning.http_server.DBGetRequest")
	proto.RegisterType((*DBGetResponse)(nil), "dmenca.go_learning.http_server.DBGetResponse")
	proto.RegisterType((*DBInfo)(nil), "dmenca.go_learning.http_server.DBInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DBApiClient is the client API for DBApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DBApiClient interface {
	DBGet(ctx context.Context, in *DBGetRequest, opts ...grpc.CallOption) (*DBGetResponse, error)
}

type dBApiClient struct {
	cc *grpc.ClientConn
}

func NewDBApiClient(cc *grpc.ClientConn) DBApiClient {
	return &dBApiClient{cc}
}

func (c *dBApiClient) DBGet(ctx context.Context, in *DBGetRequest, opts ...grpc.CallOption) (*DBGetResponse, error) {
	out := new(DBGetResponse)
	err := c.cc.Invoke(ctx, "/dmenca.go_learning.http_server.DBApi/DBGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DBApiServer is the server API for DBApi service.
type DBApiServer interface {
	DBGet(context.Context, *DBGetRequest) (*DBGetResponse, error)
}

func RegisterDBApiServer(s *grpc.Server, srv DBApiServer) {
	s.RegisterService(&_DBApi_serviceDesc, srv)
}

func _DBApi_DBGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DBGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBApiServer).DBGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dmenca.go_learning.http_server.DBApi/DBGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBApiServer).DBGet(ctx, req.(*DBGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DBApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dmenca.go_learning.http_server.DBApi",
	HandlerType: (*DBApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DBGet",
			Handler:    _DBApi_DBGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "http-server/pb/db.proto",
}

func init() { proto.RegisterFile("http-server/pb/db.proto", fileDescriptor_db_b7df01dce7a20228) }

var fileDescriptor_db_b7df01dce7a20228 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcd, 0x4a, 0xfb, 0x40,
	0x14, 0xc5, 0x69, 0x69, 0xf3, 0xe7, 0x3f, 0xea, 0x66, 0x44, 0x2a, 0x41, 0xab, 0x44, 0xfc, 0x58,
	0x98, 0x0c, 0xad, 0x0f, 0x20, 0x86, 0x40, 0xe9, 0x4e, 0xb2, 0x74, 0x13, 0x66, 0x9c, 0x9b, 0x38,
	0xd0, 0xcc, 0x8c, 0x33, 0xd3, 0xba, 0x10, 0x37, 0xba, 0x76, 0xe5, 0xa3, 0xf9, 0x0a, 0x3e, 0x88,
	0xe4, 0xa3, 0x44, 0x5c, 0xa8, 0xab, 0x5c, 0x72, 0x7e, 0x9c, 0x7b, 0xce, 0x1d, 0x34, 0xba, 0x73,
	0x4e, 0x87, 0x16, 0xcc, 0x0a, 0x0c, 0xd1, 0x8c, 0x70, 0x16, 0x69, 0xa3, 0x9c, 0xc2, 0x63, 0x5e,
	0x82, 0xbc, 0xa5, 0x51, 0xa1, 0xb2, 0x05, 0x50, 0x23, 0x85, 0x2c, 0xa2, 0x8a, 0xcd, 0x1a, 0xd6,
	0x3f, 0x28, 0x94, 0x2a, 0x16, 0x40, 0x6a, 0x9a, 0x2d, 0x73, 0xe2, 0x44, 0x09, 0xd6, 0xd1, 0x52,
	0x37, 0x06, 0xfe, 0xf8, 0x3b, 0xf0, 0x60, 0xa8, 0xd6, 0x60, 0x6c, 0xab, 0xef, 0xb5, 0x3a, 0xd5,
	0x82, 0x50, 0x29, 0x95, 0xa3, 0x4e, 0x28, 0xd9, 0xaa, 0xc1, 0x11, 0xda, 0x4c, 0xe2, 0x19, 0xb8,
	0x14, 0xee, 0x97, 0x60, 0x1d, 0xde, 0x46, 0x43, 0xce, 0x32, 0xc1, 0x77, 0x7b, 0x87, 0xbd, 0xb3,
	0xff, 0xe9, 0x80, 0xb3, 0x39, 0x0f, 0xae, 0xd1, 0x56, 0x0b, 0x59, 0xad, 0xa4, 0x05, 0x7c, 0x89,
	0xfe, 0x55, 0x94, 0xcc, 0x55, 0xcd, 0x6d, 0x4c, 0x4f, 0xa2, 0x9f, 0x6b, 0x44, 0x49, 0x3c, 0x97,
	0xb9, 0x4a, 0x3d, 0xce, 0xaa, 0x6f, 0x30, 0x41, 0x5e, 0xf3, 0x07, 0x63, 0x34, 0x90, 0xb4, 0x84,
	0xf5, 0xbe, 0x6a, 0xee, 0x42, 0xf4, 0xbb, 0x10, 0xd3, 0xd7, 0x1e, 0x1a, 0x26, 0xf1, 0x95, 0x16,
	0xf8, 0xa5, 0x9e, 0x66, 0xe0, 0xf0, 0xf9, 0xef, 0x6b, 0xbb, 0x6e, 0x7e, 0xf8, 0x47, 0xba, 0x29,
	0x19, 0xec, 0x3f, 0xbf, 0x7f, 0xbc, 0xf5, 0x47, 0x78, 0x87, 0xac, 0x26, 0x84, 0x53, 0x47, 0x19,
	0xb5, 0x60, 0xc9, 0x63, 0x1d, 0xed, 0x29, 0x3e, 0xbd, 0x39, 0x6e, 0xec, 0x48, 0xa1, 0xc2, 0xb5,
	0x1d, 0xf9, 0xfa, 0xcc, 0xd5, 0xc5, 0x39, 0x63, 0x5e, 0x7d, 0xe9, 0x8b, 0xcf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xf2, 0xcd, 0x85, 0x0b, 0x03, 0x02, 0x00, 0x00,
}
