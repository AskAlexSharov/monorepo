// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

package todo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Todo struct {
	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// @inject_tag: sql:",notnull,default:false"
	Completed bool `protobuf:"varint,4,opt,name=completed,proto3" json:"completed,omitempty"`
	// @inject_tag: sql:"type:timestamptz,default:now()"
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// @inject_tag: sql:"type:timestamptz"
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_4a91dd46d250b516, []int{0}
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

func (m *Todo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Todo) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

func (m *Todo) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Todo) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type CreateRequest struct {
	Item                 *Todo    `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_4a91dd46d250b516, []int{1}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(dst, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetItem() *Todo {
	if m != nil {
		return m.Item
	}
	return nil
}

type CreateResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_4a91dd46d250b516, []int{2}
}
func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (dst *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(dst, src)
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

type CreateBulkRequest struct {
	Items                []*Todo  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBulkRequest) Reset()         { *m = CreateBulkRequest{} }
func (m *CreateBulkRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBulkRequest) ProtoMessage()    {}
func (*CreateBulkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_4a91dd46d250b516, []int{3}
}
func (m *CreateBulkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBulkRequest.Unmarshal(m, b)
}
func (m *CreateBulkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBulkRequest.Marshal(b, m, deterministic)
}
func (dst *CreateBulkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBulkRequest.Merge(dst, src)
}
func (m *CreateBulkRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBulkRequest.Size(m)
}
func (m *CreateBulkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBulkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBulkRequest proto.InternalMessageInfo

func (m *CreateBulkRequest) GetItems() []*Todo {
	if m != nil {
		return m.Items
	}
	return nil
}

type CreateBulkResponse struct {
	Ids                  []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBulkResponse) Reset()         { *m = CreateBulkResponse{} }
func (m *CreateBulkResponse) String() string { return proto.CompactTextString(m) }
func (*CreateBulkResponse) ProtoMessage()    {}
func (*CreateBulkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_4a91dd46d250b516, []int{4}
}
func (m *CreateBulkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBulkResponse.Unmarshal(m, b)
}
func (m *CreateBulkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBulkResponse.Marshal(b, m, deterministic)
}
func (dst *CreateBulkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBulkResponse.Merge(dst, src)
}
func (m *CreateBulkResponse) XXX_Size() int {
	return xxx_messageInfo_CreateBulkResponse.Size(m)
}
func (m *CreateBulkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBulkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBulkResponse proto.InternalMessageInfo

func (m *CreateBulkResponse) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type GetRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_4a91dd46d250b516, []int{5}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetResponse struct {
	Item                 *Todo    `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_4a91dd46d250b516, []int{6}
}
func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (dst *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(dst, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetItem() *Todo {
	if m != nil {
		return m.Item
	}
	return nil
}

type ListRequest struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	NotCompleted         bool     `protobuf:"varint,2,opt,name=not_completed,json=notCompleted,proto3" json:"not_completed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_4a91dd46d250b516, []int{7}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (dst *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(dst, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListRequest) GetNotCompleted() bool {
	if m != nil {
		return m.NotCompleted
	}
	return false
}

type ListResponse struct {
	Items                []*Todo  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_4a91dd46d250b516, []int{8}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (dst *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(dst, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetItems() []*Todo {
	if m != nil {
		return m.Items
	}
	return nil
}

type DeleteRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_4a91dd46d250b516, []int{9}
}
func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(dst, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_4a91dd46d250b516, []int{10}
}
func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(dst, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

type UpdateRequest struct {
	Item                 *Todo    `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_4a91dd46d250b516, []int{11}
}
func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(dst, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetItem() *Todo {
	if m != nil {
		return m.Item
	}
	return nil
}

type UpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_4a91dd46d250b516, []int{12}
}
func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(dst, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

type UpdateBulkRequest struct {
	Items                []*Todo  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBulkRequest) Reset()         { *m = UpdateBulkRequest{} }
func (m *UpdateBulkRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateBulkRequest) ProtoMessage()    {}
func (*UpdateBulkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_4a91dd46d250b516, []int{13}
}
func (m *UpdateBulkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBulkRequest.Unmarshal(m, b)
}
func (m *UpdateBulkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBulkRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateBulkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBulkRequest.Merge(dst, src)
}
func (m *UpdateBulkRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateBulkRequest.Size(m)
}
func (m *UpdateBulkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBulkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBulkRequest proto.InternalMessageInfo

func (m *UpdateBulkRequest) GetItems() []*Todo {
	if m != nil {
		return m.Items
	}
	return nil
}

type UpdateBulkResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBulkResponse) Reset()         { *m = UpdateBulkResponse{} }
func (m *UpdateBulkResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateBulkResponse) ProtoMessage()    {}
func (*UpdateBulkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_4a91dd46d250b516, []int{14}
}
func (m *UpdateBulkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBulkResponse.Unmarshal(m, b)
}
func (m *UpdateBulkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBulkResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateBulkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBulkResponse.Merge(dst, src)
}
func (m *UpdateBulkResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateBulkResponse.Size(m)
}
func (m *UpdateBulkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBulkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBulkResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Todo)(nil), "todo.client.v2.Todo")
	proto.RegisterType((*CreateRequest)(nil), "todo.client.v2.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "todo.client.v2.CreateResponse")
	proto.RegisterType((*CreateBulkRequest)(nil), "todo.client.v2.CreateBulkRequest")
	proto.RegisterType((*CreateBulkResponse)(nil), "todo.client.v2.CreateBulkResponse")
	proto.RegisterType((*GetRequest)(nil), "todo.client.v2.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "todo.client.v2.GetResponse")
	proto.RegisterType((*ListRequest)(nil), "todo.client.v2.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "todo.client.v2.ListResponse")
	proto.RegisterType((*DeleteRequest)(nil), "todo.client.v2.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "todo.client.v2.DeleteResponse")
	proto.RegisterType((*UpdateRequest)(nil), "todo.client.v2.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "todo.client.v2.UpdateResponse")
	proto.RegisterType((*UpdateBulkRequest)(nil), "todo.client.v2.UpdateBulkRequest")
	proto.RegisterType((*UpdateBulkResponse)(nil), "todo.client.v2.UpdateBulkResponse")
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
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Bulk version of CreateTodo
	CreateBulk(ctx context.Context, in *CreateBulkRequest, opts ...grpc.CallOption) (*CreateBulkResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	UpdateBulk(ctx context.Context, in *UpdateBulkRequest, opts ...grpc.CallOption) (*UpdateBulkResponse, error)
}

type todoServiceClient struct {
	cc *grpc.ClientConn
}

func NewTodoServiceClient(cc *grpc.ClientConn) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/todo.client.v2.TodoService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) CreateBulk(ctx context.Context, in *CreateBulkRequest, opts ...grpc.CallOption) (*CreateBulkResponse, error) {
	out := new(CreateBulkResponse)
	err := c.cc.Invoke(ctx, "/todo.client.v2.TodoService/CreateBulk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/todo.client.v2.TodoService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/todo.client.v2.TodoService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/todo.client.v2.TodoService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/todo.client.v2.TodoService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) UpdateBulk(ctx context.Context, in *UpdateBulkRequest, opts ...grpc.CallOption) (*UpdateBulkResponse, error) {
	out := new(UpdateBulkResponse)
	err := c.cc.Invoke(ctx, "/todo.client.v2.TodoService/UpdateBulk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
type TodoServiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Bulk version of CreateTodo
	CreateBulk(context.Context, *CreateBulkRequest) (*CreateBulkResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	UpdateBulk(context.Context, *UpdateBulkRequest) (*UpdateBulkResponse, error)
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.client.v2.TodoService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_CreateBulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBulkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).CreateBulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.client.v2.TodoService/CreateBulk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).CreateBulk(ctx, req.(*CreateBulkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.client.v2.TodoService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.client.v2.TodoService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.client.v2.TodoService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.client.v2.TodoService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_UpdateBulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBulkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).UpdateBulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.client.v2.TodoService/UpdateBulk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).UpdateBulk(ctx, req.(*UpdateBulkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.client.v2.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TodoService_Create_Handler,
		},
		{
			MethodName: "CreateBulk",
			Handler:    _TodoService_CreateBulk_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TodoService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _TodoService_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TodoService_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TodoService_Update_Handler,
		},
		{
			MethodName: "UpdateBulk",
			Handler:    _TodoService_UpdateBulk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor_todo_4a91dd46d250b516) }

var fileDescriptor_todo_4a91dd46d250b516 = []byte{
	// 615 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xc1, 0x6e, 0xd3, 0x4a,
	0x14, 0x95, 0x93, 0xd8, 0x6a, 0xae, 0x9b, 0xbc, 0xf4, 0x2a, 0x0f, 0x59, 0x26, 0x50, 0x63, 0x24,
	0x14, 0x75, 0x61, 0x4b, 0x61, 0x81, 0xda, 0x4d, 0xd5, 0x16, 0xa9, 0x2c, 0x58, 0xb9, 0xed, 0x86,
	0x05, 0x91, 0x13, 0x0f, 0xd5, 0xa8, 0xb6, 0xc7, 0xc4, 0x93, 0x6c, 0x10, 0x1b, 0x7e, 0x81, 0xaf,
	0xe1, 0x3b, 0xf8, 0x03, 0xc4, 0x87, 0x20, 0xcf, 0x8c, 0xeb, 0xc4, 0x4e, 0x80, 0xb2, 0xf3, 0xcc,
	0x3d, 0xe7, 0xdc, 0x3b, 0xf7, 0x1c, 0x19, 0x80, 0xb3, 0x88, 0x79, 0xd9, 0x82, 0x71, 0x86, 0x7d,
	0xf1, 0x3d, 0x8f, 0x29, 0x49, 0xb9, 0xb7, 0x9a, 0xd8, 0xa3, 0x5b, 0xc6, 0x6e, 0x63, 0xe2, 0x87,
	0x19, 0xf5, 0xc3, 0x34, 0x65, 0x3c, 0xe4, 0x94, 0xa5, 0xb9, 0x44, 0xdb, 0x87, 0xaa, 0x2a, 0x4e,
	0xb3, 0xe5, 0x07, 0x9f, 0xd3, 0x84, 0xe4, 0x3c, 0x4c, 0x32, 0x09, 0x70, 0x7f, 0x68, 0xd0, 0xb9,
	0x66, 0x11, 0xc3, 0x3e, 0xb4, 0x68, 0x64, 0x69, 0x8e, 0x36, 0xee, 0x06, 0x2d, 0x1a, 0xe1, 0x10,
	0x74, 0x4e, 0x79, 0x4c, 0xac, 0x96, 0xb8, 0x92, 0x07, 0x74, 0xc0, 0x8c, 0x48, 0x3e, 0x5f, 0xd0,
	0xac, 0xe8, 0x62, 0xb5, 0x45, 0x6d, 0xfd, 0x0a, 0x47, 0xd0, 0x9d, 0xb3, 0x24, 0x8b, 0x09, 0x27,
	0x91, 0xd5, 0x71, 0xb4, 0xf1, 0x5e, 0x50, 0x5d, 0xe0, 0x31, 0xc0, 0x7c, 0x41, 0x42, 0x4e, 0xa2,
	0x69, 0xc8, 0x2d, 0xdd, 0xd1, 0xc6, 0xe6, 0xc4, 0xf6, 0xe4, 0x90, 0x5e, 0x39, 0xa4, 0x77, 0x5d,
	0x0e, 0x19, 0x74, 0x15, 0xfa, 0x8c, 0x17, 0xd4, 0x65, 0x16, 0x95, 0x54, 0xe3, 0xcf, 0x54, 0x85,
	0x3e, 0xe3, 0xee, 0x31, 0xf4, 0x2e, 0x84, 0x4e, 0x40, 0x3e, 0x2e, 0x49, 0xce, 0x71, 0x0c, 0x1d,
	0xca, 0x49, 0x22, 0x9e, 0x6b, 0x4e, 0x86, 0xde, 0xe6, 0x4e, 0xbd, 0x62, 0x21, 0x81, 0x40, 0xb8,
	0x0e, 0xf4, 0x4b, 0x6a, 0x9e, 0xb1, 0x34, 0x27, 0xf5, 0x45, 0xb9, 0xa7, 0x70, 0x20, 0x11, 0xe7,
	0xcb, 0xf8, 0xae, 0x6c, 0x70, 0x04, 0x7a, 0x41, 0xcf, 0x2d, 0xcd, 0x69, 0xef, 0xec, 0x20, 0x21,
	0xee, 0x0b, 0xc0, 0x75, 0x01, 0xd5, 0x66, 0x00, 0x6d, 0x1a, 0x49, 0x7e, 0x37, 0x28, 0x3e, 0xdd,
	0x11, 0xc0, 0x25, 0xe1, 0x65, 0x87, 0xfa, 0x18, 0xaf, 0xc0, 0x14, 0x55, 0x45, 0xff, 0xfb, 0x17,
	0xbe, 0x01, 0xf3, 0x2d, 0xcd, 0xef, 0x75, 0x87, 0xa0, 0xc7, 0x34, 0xa1, 0x5c, 0x30, 0xf5, 0x40,
	0x1e, 0xf0, 0x39, 0xf4, 0x52, 0xc6, 0xa7, 0x95, 0xb3, 0x2d, 0xe1, 0xec, 0x7e, 0xca, 0xf8, 0x45,
	0x79, 0xe7, 0x9e, 0xc0, 0xbe, 0x54, 0x52, 0x33, 0x3c, 0x64, 0x09, 0x87, 0xd0, 0x7b, 0x4d, 0x0a,
	0x99, 0x5d, 0xef, 0x1b, 0x40, 0xbf, 0x04, 0x48, 0xf9, 0xc2, 0xd5, 0x1b, 0x61, 0xf1, 0xc3, 0x5d,
	0x1d, 0x40, 0xbf, 0xa4, 0x2a, 0xb1, 0x53, 0x38, 0x90, 0x37, 0xff, 0xea, 0xe2, 0x10, 0x70, 0x5d,
	0x40, 0xca, 0x4e, 0xbe, 0xe9, 0x60, 0x16, 0xa8, 0x2b, 0xb2, 0x58, 0xd1, 0x39, 0xc1, 0x29, 0x18,
	0xd2, 0x6b, 0x7c, 0x52, 0x17, 0xdb, 0x48, 0xa8, 0xfd, 0x74, 0x57, 0x59, 0xcd, 0xfb, 0xe8, 0xcb,
	0xf7, 0x9f, 0x5f, 0x5b, 0x03, 0x77, 0xcf, 0x5f, 0x4d, 0xfc, 0x02, 0x7a, 0x22, 0x5e, 0x86, 0x09,
	0x40, 0x15, 0x26, 0x7c, 0xb6, 0x5d, 0x65, 0xed, 0x8d, 0xb6, 0xfb, 0x3b, 0x88, 0x6a, 0x66, 0x89,
	0x66, 0xe8, 0xf6, 0xca, 0x66, 0xfe, 0x6c, 0x19, 0xdf, 0x9d, 0x68, 0x47, 0x78, 0x03, 0xed, 0x4b,
	0xc2, 0xd1, 0xae, 0x8b, 0x54, 0x41, 0xb5, 0x1f, 0x6f, 0xad, 0x29, 0xe5, 0xff, 0x85, 0xf2, 0x7f,
	0x58, 0x29, 0x7f, 0xa2, 0xd1, 0x67, 0xbc, 0x82, 0x4e, 0x91, 0x24, 0x6c, 0x70, 0xd7, 0x92, 0x6a,
	0x8f, 0xb6, 0x17, 0x95, 0xf2, 0x40, 0x28, 0x03, 0xde, 0x2f, 0x08, 0xdf, 0x83, 0x21, 0x13, 0xd4,
	0xdc, 0xfd, 0x46, 0xf4, 0x9a, 0xbb, 0xaf, 0x05, 0x4f, 0x0d, 0x7d, 0x54, 0x1b, 0x7a, 0x0a, 0x86,
	0x4c, 0x40, 0x53, 0x7f, 0x23, 0xa7, 0x4d, 0xfd, 0x5a, 0x16, 0x95, 0xb7, 0xf6, 0x16, 0x6f, 0xab,
	0x88, 0x35, 0xbd, 0x6d, 0xe4, 0xb7, 0xe9, 0x6d, 0x33, 0xa1, 0xa5, 0xb7, 0x76, 0xc3, 0xdb, 0x73,
	0xe3, 0x5d, 0xa7, 0x38, 0xcf, 0x0c, 0xf1, 0x73, 0x7d, 0xf9, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x49,
	0xd3, 0x3a, 0x7a, 0x86, 0x06, 0x00, 0x00,
}
