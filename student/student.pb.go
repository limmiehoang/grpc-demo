// Code generated by protoc-gen-go. DO NOT EDIT.
// source: student.proto

package student

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Student model
type Model struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Model) Reset()         { *m = Model{} }
func (m *Model) String() string { return proto.CompactTextString(m) }
func (*Model) ProtoMessage()    {}
func (*Model) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{0}
}

func (m *Model) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Model.Unmarshal(m, b)
}
func (m *Model) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Model.Marshal(b, m, deterministic)
}
func (m *Model) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Model.Merge(m, src)
}
func (m *Model) XXX_Size() int {
	return xxx_messageInfo_Model.Size(m)
}
func (m *Model) XXX_DiscardUnknown() {
	xxx_messageInfo_Model.DiscardUnknown(m)
}

var xxx_messageInfo_Model proto.InternalMessageInfo

func (m *Model) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Model) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The find request message containing the student's id.
type FindRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindRequest) Reset()         { *m = FindRequest{} }
func (m *FindRequest) String() string { return proto.CompactTextString(m) }
func (*FindRequest) ProtoMessage()    {}
func (*FindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{1}
}

func (m *FindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindRequest.Unmarshal(m, b)
}
func (m *FindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindRequest.Marshal(b, m, deterministic)
}
func (m *FindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindRequest.Merge(m, src)
}
func (m *FindRequest) XXX_Size() int {
	return xxx_messageInfo_FindRequest.Size(m)
}
func (m *FindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindRequest proto.InternalMessageInfo

func (m *FindRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// The find response message containing the name of the student.
type FindResponse struct {
	Student              *Model   `protobuf:"bytes,1,opt,name=student,proto3" json:"student,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindResponse) Reset()         { *m = FindResponse{} }
func (m *FindResponse) String() string { return proto.CompactTextString(m) }
func (*FindResponse) ProtoMessage()    {}
func (*FindResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{2}
}

func (m *FindResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindResponse.Unmarshal(m, b)
}
func (m *FindResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindResponse.Marshal(b, m, deterministic)
}
func (m *FindResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindResponse.Merge(m, src)
}
func (m *FindResponse) XXX_Size() int {
	return xxx_messageInfo_FindResponse.Size(m)
}
func (m *FindResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindResponse proto.InternalMessageInfo

func (m *FindResponse) GetStudent() *Model {
	if m != nil {
		return m.Student
	}
	return nil
}

// The create request message containing the student's model.
type CreateRequest struct {
	Student              *Model   `protobuf:"bytes,1,opt,name=student,proto3" json:"student,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{3}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetStudent() *Model {
	if m != nil {
		return m.Student
	}
	return nil
}

// The create response message containing the status and message.
type CreateResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{4}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type GetFilter struct {
	Keyword              string   `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFilter) Reset()         { *m = GetFilter{} }
func (m *GetFilter) String() string { return proto.CompactTextString(m) }
func (*GetFilter) ProtoMessage()    {}
func (*GetFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{5}
}

func (m *GetFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFilter.Unmarshal(m, b)
}
func (m *GetFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFilter.Marshal(b, m, deterministic)
}
func (m *GetFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFilter.Merge(m, src)
}
func (m *GetFilter) XXX_Size() int {
	return xxx_messageInfo_GetFilter.Size(m)
}
func (m *GetFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFilter.DiscardUnknown(m)
}

var xxx_messageInfo_GetFilter proto.InternalMessageInfo

func (m *GetFilter) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

func init() {
	proto.RegisterType((*Model)(nil), "student.Model")
	proto.RegisterType((*FindRequest)(nil), "student.FindRequest")
	proto.RegisterType((*FindResponse)(nil), "student.FindResponse")
	proto.RegisterType((*CreateRequest)(nil), "student.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "student.CreateResponse")
	proto.RegisterType((*GetFilter)(nil), "student.GetFilter")
}

func init() { proto.RegisterFile("student.proto", fileDescriptor_94a1c1b032ad0c00) }

var fileDescriptor_94a1c1b032ad0c00 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x6d, 0x82, 0x1a, 0x3b, 0xb1, 0x39, 0x0c, 0x7e, 0x84, 0x80, 0x20, 0x0b, 0x42, 0x41, 0x28,
	0x52, 0x11, 0x54, 0x3c, 0x29, 0xb4, 0x27, 0x2f, 0xf1, 0x17, 0xd4, 0xee, 0x1c, 0x82, 0x35, 0x1b,
	0xb3, 0x13, 0xc4, 0xbf, 0xe7, 0x2f, 0x13, 0x37, 0x99, 0x45, 0xd7, 0x4b, 0x6f, 0x79, 0x93, 0xf7,
	0xe6, 0xbd, 0x37, 0x0b, 0x13, 0xcb, 0x9d, 0xa6, 0x9a, 0x67, 0x4d, 0x6b, 0xd8, 0x60, 0x32, 0x40,
	0x75, 0x01, 0xbb, 0x4f, 0x46, 0xd3, 0x06, 0x33, 0x88, 0x2b, 0x9d, 0x47, 0x67, 0xd1, 0x74, 0x5c,
	0xc6, 0x95, 0x46, 0x84, 0x9d, 0x7a, 0xf5, 0x46, 0x79, 0xec, 0x26, 0xee, 0x5b, 0x9d, 0x42, 0xba,
	0xa8, 0x6a, 0x5d, 0xd2, 0x7b, 0x47, 0x96, 0x43, 0x89, 0xba, 0x81, 0x83, 0xfe, 0xb7, 0x6d, 0x4c,
	0x6d, 0x09, 0xa7, 0x20, 0x36, 0x8e, 0x94, 0xce, 0xb3, 0x99, 0xa4, 0x70, 0x9e, 0xa5, 0x4f, 0x71,
	0x0b, 0x93, 0xc7, 0x96, 0x56, 0x4c, 0xb2, 0x7a, 0x7b, 0xe9, 0x1d, 0x64, 0x22, 0x1d, 0x6c, 0xc3,
	0x26, 0x39, 0x24, 0xb6, 0x5b, 0xaf, 0xc9, 0x5a, 0x57, 0x66, 0xbf, 0x14, 0xa8, 0xce, 0x61, 0xbc,
	0x24, 0x5e, 0x54, 0x1b, 0xa6, 0xf6, 0x87, 0xf6, 0x4a, 0x9f, 0x1f, 0xa6, 0x15, 0xad, 0xc0, 0xf9,
	0x57, 0x04, 0xc9, 0x73, 0x6f, 0x87, 0xf7, 0xfd, 0x09, 0x04, 0x1e, 0xfa, 0x58, 0xbf, 0x0e, 0x53,
	0x1c, 0x05, 0xd3, 0x3e, 0x98, 0x1a, 0xe1, 0x83, 0xf4, 0x14, 0xfd, 0xb1, 0x67, 0xfe, 0xe9, 0x5f,
	0x9c, 0xfc, 0x9b, 0xfb, 0x1d, 0xd7, 0x90, 0x2e, 0x89, 0x87, 0x05, 0x16, 0xd1, 0x33, 0x7d, 0x95,
	0x22, 0x38, 0x96, 0x1a, 0x5d, 0x46, 0x2f, 0x7b, 0xee, 0xe1, 0xaf, 0xbe, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xf1, 0x8a, 0x36, 0xbb, 0x09, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StudentClient is the client API for Student service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StudentClient interface {
	// Finds a student
	FindStudent(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error)
	// Creates a student
	CreateStudent(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Get all students
	GetStudents(ctx context.Context, in *GetFilter, opts ...grpc.CallOption) (Student_GetStudentsClient, error)
}

type studentClient struct {
	cc *grpc.ClientConn
}

func NewStudentClient(cc *grpc.ClientConn) StudentClient {
	return &studentClient{cc}
}

func (c *studentClient) FindStudent(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error) {
	out := new(FindResponse)
	err := c.cc.Invoke(ctx, "/student.Student/FindStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentClient) CreateStudent(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/student.Student/CreateStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentClient) GetStudents(ctx context.Context, in *GetFilter, opts ...grpc.CallOption) (Student_GetStudentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Student_serviceDesc.Streams[0], "/student.Student/GetStudents", opts...)
	if err != nil {
		return nil, err
	}
	x := &studentGetStudentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Student_GetStudentsClient interface {
	Recv() (*Model, error)
	grpc.ClientStream
}

type studentGetStudentsClient struct {
	grpc.ClientStream
}

func (x *studentGetStudentsClient) Recv() (*Model, error) {
	m := new(Model)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StudentServer is the server API for Student service.
type StudentServer interface {
	// Finds a student
	FindStudent(context.Context, *FindRequest) (*FindResponse, error)
	// Creates a student
	CreateStudent(context.Context, *CreateRequest) (*CreateResponse, error)
	// Get all students
	GetStudents(*GetFilter, Student_GetStudentsServer) error
}

// UnimplementedStudentServer can be embedded to have forward compatible implementations.
type UnimplementedStudentServer struct {
}

func (*UnimplementedStudentServer) FindStudent(ctx context.Context, req *FindRequest) (*FindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindStudent not implemented")
}
func (*UnimplementedStudentServer) CreateStudent(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStudent not implemented")
}
func (*UnimplementedStudentServer) GetStudents(req *GetFilter, srv Student_GetStudentsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStudents not implemented")
}

func RegisterStudentServer(s *grpc.Server, srv StudentServer) {
	s.RegisterService(&_Student_serviceDesc, srv)
}

func _Student_FindStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServer).FindStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.Student/FindStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServer).FindStudent(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Student_CreateStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServer).CreateStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.Student/CreateStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServer).CreateStudent(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Student_GetStudents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StudentServer).GetStudents(m, &studentGetStudentsServer{stream})
}

type Student_GetStudentsServer interface {
	Send(*Model) error
	grpc.ServerStream
}

type studentGetStudentsServer struct {
	grpc.ServerStream
}

func (x *studentGetStudentsServer) Send(m *Model) error {
	return x.ServerStream.SendMsg(m)
}

var _Student_serviceDesc = grpc.ServiceDesc{
	ServiceName: "student.Student",
	HandlerType: (*StudentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindStudent",
			Handler:    _Student_FindStudent_Handler,
		},
		{
			MethodName: "CreateStudent",
			Handler:    _Student_CreateStudent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStudents",
			Handler:       _Student_GetStudents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "student.proto",
}