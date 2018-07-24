// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todos/v1/service.proto

package todos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TodoState int32

const (
	TodoState_UNKNOWN  TodoState = 0
	TodoState_ARCHIVED TodoState = 1
	TodoState_ACTIVE   TodoState = 2
)

var TodoState_name = map[int32]string{
	0: "UNKNOWN",
	1: "ARCHIVED",
	2: "ACTIVE",
}
var TodoState_value = map[string]int32{
	"UNKNOWN":  0,
	"ARCHIVED": 1,
	"ACTIVE":   2,
}

func (x TodoState) String() string {
	return proto.EnumName(TodoState_name, int32(x))
}
func (TodoState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_service_33f6f30a333c090b, []int{0}
}

type Todo struct {
	Id                   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string    `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	State                TodoState `protobuf:"varint,3,opt,name=state,proto3,enum=todos.v1.TodoState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_33f6f30a333c090b, []int{0}
}
func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (dst *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(dst, src)
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

func (m *Todo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Todo) GetState() TodoState {
	if m != nil {
		return m.State
	}
	return TodoState_UNKNOWN
}

type TodoResponse struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TodoResponse) Reset()         { *m = TodoResponse{} }
func (m *TodoResponse) String() string { return proto.CompactTextString(m) }
func (*TodoResponse) ProtoMessage()    {}
func (*TodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_33f6f30a333c090b, []int{1}
}
func (m *TodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TodoResponse.Unmarshal(m, b)
}
func (m *TodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TodoResponse.Marshal(b, m, deterministic)
}
func (dst *TodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TodoResponse.Merge(dst, src)
}
func (m *TodoResponse) XXX_Size() int {
	return xxx_messageInfo_TodoResponse.Size(m)
}
func (m *TodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TodoResponse proto.InternalMessageInfo

func (m *TodoResponse) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type ListTodoResponse struct {
	Todos                []*Todo  `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTodoResponse) Reset()         { *m = ListTodoResponse{} }
func (m *ListTodoResponse) String() string { return proto.CompactTextString(m) }
func (*ListTodoResponse) ProtoMessage()    {}
func (*ListTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_33f6f30a333c090b, []int{2}
}
func (m *ListTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodoResponse.Unmarshal(m, b)
}
func (m *ListTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodoResponse.Marshal(b, m, deterministic)
}
func (dst *ListTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodoResponse.Merge(dst, src)
}
func (m *ListTodoResponse) XXX_Size() int {
	return xxx_messageInfo_ListTodoResponse.Size(m)
}
func (m *ListTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodoResponse proto.InternalMessageInfo

func (m *ListTodoResponse) GetTodos() []*Todo {
	if m != nil {
		return m.Todos
	}
	return nil
}

type CreateTodoRequest struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoRequest) Reset()         { *m = CreateTodoRequest{} }
func (m *CreateTodoRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTodoRequest) ProtoMessage()    {}
func (*CreateTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_33f6f30a333c090b, []int{3}
}
func (m *CreateTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoRequest.Unmarshal(m, b)
}
func (m *CreateTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoRequest.Marshal(b, m, deterministic)
}
func (dst *CreateTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoRequest.Merge(dst, src)
}
func (m *CreateTodoRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTodoRequest.Size(m)
}
func (m *CreateTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoRequest proto.InternalMessageInfo

func (m *CreateTodoRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type RemoveTodoRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveTodoRequest) Reset()         { *m = RemoveTodoRequest{} }
func (m *RemoveTodoRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveTodoRequest) ProtoMessage()    {}
func (*RemoveTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_33f6f30a333c090b, []int{4}
}
func (m *RemoveTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveTodoRequest.Unmarshal(m, b)
}
func (m *RemoveTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveTodoRequest.Marshal(b, m, deterministic)
}
func (dst *RemoveTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveTodoRequest.Merge(dst, src)
}
func (m *RemoveTodoRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveTodoRequest.Size(m)
}
func (m *RemoveTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveTodoRequest proto.InternalMessageInfo

func (m *RemoveTodoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetTodoRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTodoRequest) Reset()         { *m = GetTodoRequest{} }
func (m *GetTodoRequest) String() string { return proto.CompactTextString(m) }
func (*GetTodoRequest) ProtoMessage()    {}
func (*GetTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_33f6f30a333c090b, []int{5}
}
func (m *GetTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTodoRequest.Unmarshal(m, b)
}
func (m *GetTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTodoRequest.Marshal(b, m, deterministic)
}
func (dst *GetTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTodoRequest.Merge(dst, src)
}
func (m *GetTodoRequest) XXX_Size() int {
	return xxx_messageInfo_GetTodoRequest.Size(m)
}
func (m *GetTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTodoRequest proto.InternalMessageInfo

func (m *GetTodoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListTodosRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTodosRequest) Reset()         { *m = ListTodosRequest{} }
func (m *ListTodosRequest) String() string { return proto.CompactTextString(m) }
func (*ListTodosRequest) ProtoMessage()    {}
func (*ListTodosRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_33f6f30a333c090b, []int{6}
}
func (m *ListTodosRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodosRequest.Unmarshal(m, b)
}
func (m *ListTodosRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodosRequest.Marshal(b, m, deterministic)
}
func (dst *ListTodosRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodosRequest.Merge(dst, src)
}
func (m *ListTodosRequest) XXX_Size() int {
	return xxx_messageInfo_ListTodosRequest.Size(m)
}
func (m *ListTodosRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodosRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodosRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Todo)(nil), "todos.v1.Todo")
	proto.RegisterType((*TodoResponse)(nil), "todos.v1.TodoResponse")
	proto.RegisterType((*ListTodoResponse)(nil), "todos.v1.ListTodoResponse")
	proto.RegisterType((*CreateTodoRequest)(nil), "todos.v1.CreateTodoRequest")
	proto.RegisterType((*RemoveTodoRequest)(nil), "todos.v1.RemoveTodoRequest")
	proto.RegisterType((*GetTodoRequest)(nil), "todos.v1.GetTodoRequest")
	proto.RegisterType((*ListTodosRequest)(nil), "todos.v1.ListTodosRequest")
	proto.RegisterEnum("todos.v1.TodoState", TodoState_name, TodoState_value)
}

func init() { proto.RegisterFile("todos/v1/service.proto", fileDescriptor_service_33f6f30a333c090b) }

var fileDescriptor_service_33f6f30a333c090b = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x4f, 0xc2, 0x40,
	0x10, 0xc5, 0x6d, 0xf9, 0x3f, 0x90, 0xa6, 0x8c, 0x86, 0x34, 0x78, 0x69, 0x56, 0x0f, 0xe0, 0xa1,
	0x84, 0x7a, 0x31, 0xf1, 0x54, 0x2b, 0x51, 0xa2, 0xc1, 0xa4, 0x22, 0x24, 0xde, 0xd0, 0xee, 0xa1,
	0x89, 0xba, 0xc8, 0xae, 0xfd, 0x26, 0x7e, 0x5f, 0xd3, 0x5d, 0x60, 0x29, 0x7f, 0x3c, 0x76, 0xe6,
	0xbd, 0xd7, 0xfe, 0xde, 0x14, 0x5a, 0x82, 0xc5, 0x8c, 0xf7, 0xd2, 0x7e, 0x8f, 0xd3, 0x45, 0x9a,
	0xbc, 0x53, 0x6f, 0xbe, 0x60, 0x82, 0x61, 0x55, 0xce, 0xbd, 0xb4, 0x4f, 0xa6, 0x50, 0x1c, 0xb3,
	0x98, 0xa1, 0x05, 0x66, 0x12, 0x3b, 0x86, 0x6b, 0x74, 0x6a, 0x91, 0x99, 0xc4, 0x78, 0x02, 0x25,
	0x91, 0x88, 0x0f, 0xea, 0x98, 0x72, 0xa4, 0x1e, 0xb0, 0x0b, 0x25, 0x2e, 0x66, 0x82, 0x3a, 0x05,
	0xd7, 0xe8, 0x58, 0xfe, 0xb1, 0xb7, 0xca, 0xf1, 0xb2, 0x90, 0xe7, 0x6c, 0x15, 0x29, 0x05, 0xf1,
	0xa1, 0x91, 0xcd, 0x22, 0xca, 0xe7, 0xec, 0x8b, 0x53, 0x24, 0x50, 0xcc, 0xc4, 0xf2, 0x15, 0x75,
	0xdf, 0xca, 0x3b, 0x23, 0xb9, 0x23, 0x57, 0x60, 0x3f, 0x26, 0x5c, 0xe4, 0x7c, 0xe7, 0x50, 0x92,
	0x52, 0xc7, 0x70, 0x0b, 0x7b, 0x8c, 0x6a, 0x49, 0xba, 0xd0, 0x0c, 0x17, 0x74, 0x26, 0xa8, 0xf2,
	0x7e, 0xff, 0x50, 0x2e, 0x34, 0x83, 0xb1, 0xc1, 0x40, 0xce, 0xa0, 0x19, 0xd1, 0x4f, 0x96, 0xe6,
	0xa4, 0x5b, 0xf8, 0xc4, 0x05, 0xeb, 0x8e, 0x8a, 0xff, 0x14, 0xa8, 0xbf, 0x95, 0x2f, 0x35, 0x17,
	0x3e, 0xd4, 0xd6, 0x3d, 0x60, 0x1d, 0x2a, 0x2f, 0xa3, 0x87, 0xd1, 0xd3, 0x74, 0x64, 0x1f, 0x61,
	0x03, 0xaa, 0x41, 0x14, 0xde, 0x0f, 0x27, 0x83, 0x5b, 0xdb, 0x40, 0x80, 0x72, 0x10, 0x8e, 0x87,
	0x93, 0x81, 0x6d, 0xfa, 0xbf, 0x26, 0xd4, 0xa5, 0x49, 0x1d, 0x08, 0x03, 0x00, 0x4d, 0x82, 0xa7,
	0x1a, 0x77, 0x87, 0xaf, 0xdd, 0xda, 0xea, 0x62, 0x55, 0x59, 0x00, 0xa0, 0x09, 0x37, 0x23, 0x76,
	0xb8, 0x0f, 0x46, 0x5c, 0x43, 0x65, 0xc9, 0x8f, 0x8e, 0x96, 0xe4, 0x2b, 0x39, 0x68, 0x0e, 0xa1,
	0xb6, 0xae, 0x06, 0xdb, 0x5a, 0xb4, 0xdd, 0x57, 0x7b, 0xcf, 0x6e, 0x15, 0x72, 0x53, 0x79, 0x55,
	0xa7, 0x7d, 0x2b, 0xcb, 0x5f, 0xf6, 0xf2, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x28, 0x98, 0x35, 0x63,
	0xcc, 0x02, 0x00, 0x00,
}
