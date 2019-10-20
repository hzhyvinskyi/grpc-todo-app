// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/todo.proto

package todopb

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

type Todo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28b1ccc7160a78f, []int{0}
}

func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (m *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(m, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Todo) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Todo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Todo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type CreateTodoReq struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoReq) Reset()         { *m = CreateTodoReq{} }
func (m *CreateTodoReq) String() string { return proto.CompactTextString(m) }
func (*CreateTodoReq) ProtoMessage()    {}
func (*CreateTodoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28b1ccc7160a78f, []int{1}
}

func (m *CreateTodoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoReq.Unmarshal(m, b)
}
func (m *CreateTodoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoReq.Marshal(b, m, deterministic)
}
func (m *CreateTodoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoReq.Merge(m, src)
}
func (m *CreateTodoReq) XXX_Size() int {
	return xxx_messageInfo_CreateTodoReq.Size(m)
}
func (m *CreateTodoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoReq proto.InternalMessageInfo

func (m *CreateTodoReq) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type CreateTodoRes struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoRes) Reset()         { *m = CreateTodoRes{} }
func (m *CreateTodoRes) String() string { return proto.CompactTextString(m) }
func (*CreateTodoRes) ProtoMessage()    {}
func (*CreateTodoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28b1ccc7160a78f, []int{2}
}

func (m *CreateTodoRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoRes.Unmarshal(m, b)
}
func (m *CreateTodoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoRes.Marshal(b, m, deterministic)
}
func (m *CreateTodoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoRes.Merge(m, src)
}
func (m *CreateTodoRes) XXX_Size() int {
	return xxx_messageInfo_CreateTodoRes.Size(m)
}
func (m *CreateTodoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoRes proto.InternalMessageInfo

func (m *CreateTodoRes) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type ReadTodoReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadTodoReq) Reset()         { *m = ReadTodoReq{} }
func (m *ReadTodoReq) String() string { return proto.CompactTextString(m) }
func (*ReadTodoReq) ProtoMessage()    {}
func (*ReadTodoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28b1ccc7160a78f, []int{3}
}

func (m *ReadTodoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadTodoReq.Unmarshal(m, b)
}
func (m *ReadTodoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadTodoReq.Marshal(b, m, deterministic)
}
func (m *ReadTodoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadTodoReq.Merge(m, src)
}
func (m *ReadTodoReq) XXX_Size() int {
	return xxx_messageInfo_ReadTodoReq.Size(m)
}
func (m *ReadTodoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadTodoReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReadTodoReq proto.InternalMessageInfo

func (m *ReadTodoReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadTodoRes struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadTodoRes) Reset()         { *m = ReadTodoRes{} }
func (m *ReadTodoRes) String() string { return proto.CompactTextString(m) }
func (*ReadTodoRes) ProtoMessage()    {}
func (*ReadTodoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28b1ccc7160a78f, []int{4}
}

func (m *ReadTodoRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadTodoRes.Unmarshal(m, b)
}
func (m *ReadTodoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadTodoRes.Marshal(b, m, deterministic)
}
func (m *ReadTodoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadTodoRes.Merge(m, src)
}
func (m *ReadTodoRes) XXX_Size() int {
	return xxx_messageInfo_ReadTodoRes.Size(m)
}
func (m *ReadTodoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadTodoRes.DiscardUnknown(m)
}

var xxx_messageInfo_ReadTodoRes proto.InternalMessageInfo

func (m *ReadTodoRes) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type UpdateTodoReq struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTodoReq) Reset()         { *m = UpdateTodoReq{} }
func (m *UpdateTodoReq) String() string { return proto.CompactTextString(m) }
func (*UpdateTodoReq) ProtoMessage()    {}
func (*UpdateTodoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28b1ccc7160a78f, []int{5}
}

func (m *UpdateTodoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTodoReq.Unmarshal(m, b)
}
func (m *UpdateTodoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTodoReq.Marshal(b, m, deterministic)
}
func (m *UpdateTodoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTodoReq.Merge(m, src)
}
func (m *UpdateTodoReq) XXX_Size() int {
	return xxx_messageInfo_UpdateTodoReq.Size(m)
}
func (m *UpdateTodoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTodoReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTodoReq proto.InternalMessageInfo

func (m *UpdateTodoReq) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type UpdateTodoRes struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTodoRes) Reset()         { *m = UpdateTodoRes{} }
func (m *UpdateTodoRes) String() string { return proto.CompactTextString(m) }
func (*UpdateTodoRes) ProtoMessage()    {}
func (*UpdateTodoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28b1ccc7160a78f, []int{6}
}

func (m *UpdateTodoRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTodoRes.Unmarshal(m, b)
}
func (m *UpdateTodoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTodoRes.Marshal(b, m, deterministic)
}
func (m *UpdateTodoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTodoRes.Merge(m, src)
}
func (m *UpdateTodoRes) XXX_Size() int {
	return xxx_messageInfo_UpdateTodoRes.Size(m)
}
func (m *UpdateTodoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTodoRes.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTodoRes proto.InternalMessageInfo

func (m *UpdateTodoRes) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type DeleteTodoReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoReq) Reset()         { *m = DeleteTodoReq{} }
func (m *DeleteTodoReq) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoReq) ProtoMessage()    {}
func (*DeleteTodoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28b1ccc7160a78f, []int{7}
}

func (m *DeleteTodoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoReq.Unmarshal(m, b)
}
func (m *DeleteTodoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoReq.Marshal(b, m, deterministic)
}
func (m *DeleteTodoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoReq.Merge(m, src)
}
func (m *DeleteTodoReq) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoReq.Size(m)
}
func (m *DeleteTodoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoReq proto.InternalMessageInfo

func (m *DeleteTodoReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteTodoRes struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoRes) Reset()         { *m = DeleteTodoRes{} }
func (m *DeleteTodoRes) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoRes) ProtoMessage()    {}
func (*DeleteTodoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28b1ccc7160a78f, []int{8}
}

func (m *DeleteTodoRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoRes.Unmarshal(m, b)
}
func (m *DeleteTodoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoRes.Marshal(b, m, deterministic)
}
func (m *DeleteTodoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoRes.Merge(m, src)
}
func (m *DeleteTodoRes) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoRes.Size(m)
}
func (m *DeleteTodoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoRes.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoRes proto.InternalMessageInfo

func (m *DeleteTodoRes) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ListTodoReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTodoReq) Reset()         { *m = ListTodoReq{} }
func (m *ListTodoReq) String() string { return proto.CompactTextString(m) }
func (*ListTodoReq) ProtoMessage()    {}
func (*ListTodoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28b1ccc7160a78f, []int{9}
}

func (m *ListTodoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodoReq.Unmarshal(m, b)
}
func (m *ListTodoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodoReq.Marshal(b, m, deterministic)
}
func (m *ListTodoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodoReq.Merge(m, src)
}
func (m *ListTodoReq) XXX_Size() int {
	return xxx_messageInfo_ListTodoReq.Size(m)
}
func (m *ListTodoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodoReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodoReq proto.InternalMessageInfo

type ListTodoRes struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTodoRes) Reset()         { *m = ListTodoRes{} }
func (m *ListTodoRes) String() string { return proto.CompactTextString(m) }
func (*ListTodoRes) ProtoMessage()    {}
func (*ListTodoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28b1ccc7160a78f, []int{10}
}

func (m *ListTodoRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodoRes.Unmarshal(m, b)
}
func (m *ListTodoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodoRes.Marshal(b, m, deterministic)
}
func (m *ListTodoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodoRes.Merge(m, src)
}
func (m *ListTodoRes) XXX_Size() int {
	return xxx_messageInfo_ListTodoRes.Size(m)
}
func (m *ListTodoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodoRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodoRes proto.InternalMessageInfo

func (m *ListTodoRes) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

func init() {
	proto.RegisterType((*Todo)(nil), "todo.Todo")
	proto.RegisterType((*CreateTodoReq)(nil), "todo.CreateTodoReq")
	proto.RegisterType((*CreateTodoRes)(nil), "todo.CreateTodoRes")
	proto.RegisterType((*ReadTodoReq)(nil), "todo.ReadTodoReq")
	proto.RegisterType((*ReadTodoRes)(nil), "todo.ReadTodoRes")
	proto.RegisterType((*UpdateTodoReq)(nil), "todo.UpdateTodoReq")
	proto.RegisterType((*UpdateTodoRes)(nil), "todo.UpdateTodoRes")
	proto.RegisterType((*DeleteTodoReq)(nil), "todo.DeleteTodoReq")
	proto.RegisterType((*DeleteTodoRes)(nil), "todo.DeleteTodoRes")
	proto.RegisterType((*ListTodoReq)(nil), "todo.ListTodoReq")
	proto.RegisterType((*ListTodoRes)(nil), "todo.ListTodoRes")
}

func init() { proto.RegisterFile("proto/todo.proto", fileDescriptor_d28b1ccc7160a78f) }

var fileDescriptor_d28b1ccc7160a78f = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0xa5, 0xb1, 0xb6, 0x71, 0x42, 0x45, 0x47, 0xc1, 0x10, 0x50, 0x4b, 0x4e, 0x7a, 0xb0, 0x2d,
	0x11, 0xfc, 0x01, 0xea, 0x45, 0xf0, 0x14, 0xf5, 0xe2, 0x45, 0xda, 0xec, 0x20, 0x0b, 0xc5, 0x8d,
	0x99, 0xad, 0x7f, 0xc6, 0x3f, 0x2b, 0xbb, 0x71, 0xc9, 0x47, 0x03, 0xd1, 0x5b, 0xde, 0x9b, 0x79,
	0xfb, 0xde, 0xbe, 0x25, 0x70, 0x90, 0x17, 0x4a, 0xab, 0xb9, 0x56, 0x42, 0xcd, 0xec, 0x27, 0x0e,
	0xcd, 0x77, 0xfc, 0x0e, 0xc3, 0x67, 0x25, 0x14, 0xee, 0x83, 0x27, 0x45, 0x38, 0x98, 0x0e, 0x2e,
	0xf6, 0x52, 0x4f, 0x0a, 0x3c, 0x81, 0xf1, 0x86, 0xa9, 0x78, 0x93, 0x22, 0xf4, 0x2c, 0x39, 0x32,
	0xf0, 0x41, 0xe0, 0x31, 0xec, 0x6a, 0xa9, 0xd7, 0x14, 0xee, 0x58, 0xba, 0x04, 0x38, 0x85, 0x40,
	0x10, 0x67, 0x85, 0xcc, 0xb5, 0x54, 0x1f, 0xe1, 0xd0, 0xce, 0xea, 0x54, 0x3c, 0x87, 0xc9, 0x5d,
	0x41, 0x4b, 0x4d, 0xc6, 0x2e, 0xa5, 0x4f, 0x3c, 0x03, 0x9b, 0xc0, 0x7a, 0x06, 0x09, 0xcc, 0x6c,
	0x34, 0x3b, 0x2c, 0x93, 0xb5, 0x04, 0xdc, 0x2b, 0x38, 0x85, 0x20, 0xa5, 0xa5, 0x70, 0xe7, 0xb7,
	0x6e, 0x14, 0x5f, 0xd5, 0xc7, 0xfc, 0x17, 0xfb, 0x97, 0x5c, 0xfc, 0x2f, 0x6f, 0x5d, 0xd0, 0xef,
	0x70, 0x0e, 0x93, 0x7b, 0x5a, 0x53, 0xe5, 0xd0, 0x4e, 0x7c, 0xd9, 0x5c, 0x60, 0x0c, 0x61, 0xcc,
	0x9b, 0x2c, 0x23, 0x66, 0xbb, 0xe5, 0xa7, 0x0e, 0xc6, 0x13, 0x08, 0x1e, 0x25, 0xeb, 0xdf, 0x93,
	0xcc, 0x5d, 0x2b, 0xd8, 0x9b, 0x24, 0xf9, 0xf6, 0x20, 0x30, 0xf0, 0x89, 0x8a, 0x2f, 0x99, 0x11,
	0xde, 0x00, 0x54, 0xd5, 0xe3, 0x51, 0xb9, 0xdf, 0x78, 0xbd, 0xa8, 0x83, 0x64, 0x5c, 0x80, 0xef,
	0x2a, 0xc6, 0xc3, 0x72, 0xa1, 0xf6, 0x22, 0xd1, 0x16, 0xc5, 0xc6, 0xa9, 0x2a, 0xcd, 0x39, 0x35,
	0x7a, 0x8f, 0x3a, 0x48, 0xab, 0xab, 0xaa, 0x71, 0xba, 0x46, 0x9b, 0x51, 0x07, 0xc9, 0x98, 0x80,
	0xef, 0x8a, 0x71, 0x09, 0x6b, 0xbd, 0x45, 0x5b, 0x14, 0x2f, 0x06, 0xb7, 0xfe, 0xeb, 0xc8, 0xb0,
	0xf9, 0x6a, 0x35, 0xb2, 0x7f, 0xce, 0xf5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x63, 0x31, 0xac,
	0x47, 0x4d, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoServiceClient interface {
	CreateTodo(ctx context.Context, in *CreateTodoReq, opts ...grpc.CallOption) (*CreateTodoRes, error)
	ReadTodo(ctx context.Context, in *ReadTodoReq, opts ...grpc.CallOption) (*ReadTodoRes, error)
	UpdateTodo(ctx context.Context, in *UpdateTodoReq, opts ...grpc.CallOption) (*UpdateTodoRes, error)
	DeleteTodo(ctx context.Context, in *DeleteTodoReq, opts ...grpc.CallOption) (*DeleteTodoRes, error)
	ListTodo(ctx context.Context, in *ListTodoReq, opts ...grpc.CallOption) (TodoService_ListTodoClient, error)
}

type todoServiceClient struct {
	cc *grpc.ClientConn
}

func NewTodoServiceClient(cc *grpc.ClientConn) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) CreateTodo(ctx context.Context, in *CreateTodoReq, opts ...grpc.CallOption) (*CreateTodoRes, error) {
	out := new(CreateTodoRes)
	err := c.cc.Invoke(ctx, "/todo.TodoService/CreateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) ReadTodo(ctx context.Context, in *ReadTodoReq, opts ...grpc.CallOption) (*ReadTodoRes, error) {
	out := new(ReadTodoRes)
	err := c.cc.Invoke(ctx, "/todo.TodoService/ReadTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) UpdateTodo(ctx context.Context, in *UpdateTodoReq, opts ...grpc.CallOption) (*UpdateTodoRes, error) {
	out := new(UpdateTodoRes)
	err := c.cc.Invoke(ctx, "/todo.TodoService/UpdateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) DeleteTodo(ctx context.Context, in *DeleteTodoReq, opts ...grpc.CallOption) (*DeleteTodoRes, error) {
	out := new(DeleteTodoRes)
	err := c.cc.Invoke(ctx, "/todo.TodoService/DeleteTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) ListTodo(ctx context.Context, in *ListTodoReq, opts ...grpc.CallOption) (TodoService_ListTodoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TodoService_serviceDesc.Streams[0], "/todo.TodoService/ListTodo", opts...)
	if err != nil {
		return nil, err
	}
	x := &todoServiceListTodoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TodoService_ListTodoClient interface {
	Recv() (*ListTodoRes, error)
	grpc.ClientStream
}

type todoServiceListTodoClient struct {
	grpc.ClientStream
}

func (x *todoServiceListTodoClient) Recv() (*ListTodoRes, error) {
	m := new(ListTodoRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TodoServiceServer is the server API for TodoService service.
type TodoServiceServer interface {
	CreateTodo(context.Context, *CreateTodoReq) (*CreateTodoRes, error)
	ReadTodo(context.Context, *ReadTodoReq) (*ReadTodoRes, error)
	UpdateTodo(context.Context, *UpdateTodoReq) (*UpdateTodoRes, error)
	DeleteTodo(context.Context, *DeleteTodoReq) (*DeleteTodoRes, error)
	ListTodo(*ListTodoReq, TodoService_ListTodoServer) error
}

// UnimplementedTodoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (*UnimplementedTodoServiceServer) CreateTodo(ctx context.Context, req *CreateTodoReq) (*CreateTodoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (*UnimplementedTodoServiceServer) ReadTodo(ctx context.Context, req *ReadTodoReq) (*ReadTodoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadTodo not implemented")
}
func (*UnimplementedTodoServiceServer) UpdateTodo(ctx context.Context, req *UpdateTodoReq) (*UpdateTodoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodo not implemented")
}
func (*UnimplementedTodoServiceServer) DeleteTodo(ctx context.Context, req *DeleteTodoReq) (*DeleteTodoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodo not implemented")
}
func (*UnimplementedTodoServiceServer) ListTodo(req *ListTodoReq, srv TodoService_ListTodoServer) error {
	return status.Errorf(codes.Unimplemented, "method ListTodo not implemented")
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).CreateTodo(ctx, req.(*CreateTodoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_ReadTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadTodoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).ReadTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/ReadTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).ReadTodo(ctx, req.(*ReadTodoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_UpdateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTodoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).UpdateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/UpdateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).UpdateTodo(ctx, req.(*UpdateTodoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_DeleteTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTodoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).DeleteTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/DeleteTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).DeleteTodo(ctx, req.(*DeleteTodoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_ListTodo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListTodoReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TodoServiceServer).ListTodo(m, &todoServiceListTodoServer{stream})
}

type TodoService_ListTodoServer interface {
	Send(*ListTodoRes) error
	grpc.ServerStream
}

type todoServiceListTodoServer struct {
	grpc.ServerStream
}

func (x *todoServiceListTodoServer) Send(m *ListTodoRes) error {
	return x.ServerStream.SendMsg(m)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodo",
			Handler:    _TodoService_CreateTodo_Handler,
		},
		{
			MethodName: "ReadTodo",
			Handler:    _TodoService_ReadTodo_Handler,
		},
		{
			MethodName: "UpdateTodo",
			Handler:    _TodoService_UpdateTodo_Handler,
		},
		{
			MethodName: "DeleteTodo",
			Handler:    _TodoService_DeleteTodo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListTodo",
			Handler:       _TodoService_ListTodo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/todo.proto",
}