// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Task struct {
	Id                   string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Parameters           map[string]*Parameter `protobuf:"bytes,2,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Blueprint            string                `protobuf:"bytes,3,opt,name=blueprint,proto3" json:"blueprint,omitempty"`
	IgnoreErrors         bool                  `protobuf:"varint,4,opt,name=ignore_errors,json=ignoreErrors,proto3" json:"ignore_errors,omitempty"`
	Steps                []*Step               `protobuf:"bytes,5,rep,name=steps,proto3" json:"steps,omitempty"`
	Status               string                `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	StartedAt            *timestamp.Timestamp  `protobuf:"bytes,7,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp  `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{0}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Task) GetParameters() map[string]*Parameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *Task) GetBlueprint() string {
	if m != nil {
		return m.Blueprint
	}
	return ""
}

func (m *Task) GetIgnoreErrors() bool {
	if m != nil {
		return m.IgnoreErrors
	}
	return false
}

func (m *Task) GetSteps() []*Step {
	if m != nil {
		return m.Steps
	}
	return nil
}

func (m *Task) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Task) GetStartedAt() *timestamp.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func (m *Task) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type Parameter struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	IsSecret             bool     `protobuf:"varint,3,opt,name=is_secret,json=isSecret,proto3" json:"is_secret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Parameter) Reset()         { *m = Parameter{} }
func (m *Parameter) String() string { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()    {}
func (*Parameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{1}
}

func (m *Parameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Parameter.Unmarshal(m, b)
}
func (m *Parameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Parameter.Marshal(b, m, deterministic)
}
func (m *Parameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parameter.Merge(m, src)
}
func (m *Parameter) XXX_Size() int {
	return xxx_messageInfo_Parameter.Size(m)
}
func (m *Parameter) XXX_DiscardUnknown() {
	xxx_messageInfo_Parameter.DiscardUnknown(m)
}

var xxx_messageInfo_Parameter proto.InternalMessageInfo

func (m *Parameter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Parameter) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Parameter) GetIsSecret() bool {
	if m != nil {
		return m.IsSecret
	}
	return false
}

type TaskDocker struct {
	Registries           []*DockerRegistry `protobuf:"bytes,1,rep,name=registries,proto3" json:"registries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TaskDocker) Reset()         { *m = TaskDocker{} }
func (m *TaskDocker) String() string { return proto.CompactTextString(m) }
func (*TaskDocker) ProtoMessage()    {}
func (*TaskDocker) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{2}
}

func (m *TaskDocker) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskDocker.Unmarshal(m, b)
}
func (m *TaskDocker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskDocker.Marshal(b, m, deterministic)
}
func (m *TaskDocker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskDocker.Merge(m, src)
}
func (m *TaskDocker) XXX_Size() int {
	return xxx_messageInfo_TaskDocker.Size(m)
}
func (m *TaskDocker) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskDocker.DiscardUnknown(m)
}

var xxx_messageInfo_TaskDocker proto.InternalMessageInfo

func (m *TaskDocker) GetRegistries() []*DockerRegistry {
	if m != nil {
		return m.Registries
	}
	return nil
}

type DockerRegistry struct {
	Address              string            `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Use                  string            `protobuf:"bytes,2,opt,name=use,proto3" json:"use,omitempty"`
	Arguments            map[string]string `protobuf:"bytes,3,rep,name=arguments,proto3" json:"arguments,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AuthorizationToken   string            `protobuf:"bytes,4,opt,name=authorization_token,json=authorizationToken,proto3" json:"authorization_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DockerRegistry) Reset()         { *m = DockerRegistry{} }
func (m *DockerRegistry) String() string { return proto.CompactTextString(m) }
func (*DockerRegistry) ProtoMessage()    {}
func (*DockerRegistry) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{3}
}

func (m *DockerRegistry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DockerRegistry.Unmarshal(m, b)
}
func (m *DockerRegistry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DockerRegistry.Marshal(b, m, deterministic)
}
func (m *DockerRegistry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DockerRegistry.Merge(m, src)
}
func (m *DockerRegistry) XXX_Size() int {
	return xxx_messageInfo_DockerRegistry.Size(m)
}
func (m *DockerRegistry) XXX_DiscardUnknown() {
	xxx_messageInfo_DockerRegistry.DiscardUnknown(m)
}

var xxx_messageInfo_DockerRegistry proto.InternalMessageInfo

func (m *DockerRegistry) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *DockerRegistry) GetUse() string {
	if m != nil {
		return m.Use
	}
	return ""
}

func (m *DockerRegistry) GetArguments() map[string]string {
	if m != nil {
		return m.Arguments
	}
	return nil
}

func (m *DockerRegistry) GetAuthorizationToken() string {
	if m != nil {
		return m.AuthorizationToken
	}
	return ""
}

func init() {
	proto.RegisterType((*Task)(nil), "velocity.v1.Task")
	proto.RegisterMapType((map[string]*Parameter)(nil), "velocity.v1.Task.ParametersEntry")
	proto.RegisterType((*Parameter)(nil), "velocity.v1.Parameter")
	proto.RegisterType((*TaskDocker)(nil), "velocity.v1.TaskDocker")
	proto.RegisterType((*DockerRegistry)(nil), "velocity.v1.DockerRegistry")
	proto.RegisterMapType((map[string]string)(nil), "velocity.v1.DockerRegistry.ArgumentsEntry")
}

func init() { proto.RegisterFile("task.proto", fileDescriptor_ce5d8dd45b4a91ff) }

var fileDescriptor_ce5d8dd45b4a91ff = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x55, 0xd2, 0x6e, 0xb7, 0x99, 0x42, 0x01, 0x83, 0x56, 0x56, 0x16, 0x89, 0x52, 0x0e, 0x54,
	0x08, 0xa5, 0xda, 0xe5, 0xc2, 0xd7, 0xa5, 0x88, 0x95, 0xe0, 0x86, 0xbc, 0xe5, 0xc2, 0xa5, 0x72,
	0x9b, 0x21, 0x58, 0x69, 0xe2, 0xc8, 0x9e, 0x54, 0x2a, 0x3f, 0x89, 0xff, 0xc7, 0x1d, 0xc5, 0x49,
	0xba, 0x0d, 0x5a, 0xc1, 0xcd, 0x33, 0xf3, 0xde, 0xcc, 0xf3, 0x9b, 0x01, 0x20, 0x69, 0xd3, 0xa8,
	0x30, 0x9a, 0x34, 0x1b, 0xed, 0x70, 0xab, 0x37, 0x8a, 0xf6, 0xd1, 0xee, 0x22, 0x04, 0x4b, 0x58,
	0xd4, 0x85, 0xf0, 0x49, 0xa2, 0x75, 0xb2, 0xc5, 0xb9, 0x8b, 0xd6, 0xe5, 0xf7, 0x39, 0xa9, 0x0c,
	0x2d, 0xc9, 0xac, 0x01, 0x4c, 0x7f, 0xf5, 0xa0, 0xbf, 0x94, 0x36, 0x65, 0x63, 0xf0, 0x55, 0xcc,
	0xbd, 0x89, 0x37, 0x0b, 0x84, 0xaf, 0x62, 0xb6, 0x00, 0x28, 0xa4, 0x91, 0x19, 0x12, 0x1a, 0xcb,
	0xfd, 0x49, 0x6f, 0x36, 0xba, 0x7c, 0x1a, 0x1d, 0xcd, 0x89, 0x2a, 0x5a, 0xf4, 0xe5, 0x80, 0xb9,
	0xca, 0xc9, 0xec, 0xc5, 0x11, 0x89, 0x3d, 0x86, 0x60, 0xbd, 0x2d, 0xb1, 0x30, 0x2a, 0x27, 0xde,
	0x73, 0x9d, 0x6f, 0x12, 0xec, 0x19, 0xdc, 0x55, 0x49, 0xae, 0x0d, 0xae, 0xd0, 0x18, 0x6d, 0x2c,
	0xef, 0x4f, 0xbc, 0xd9, 0x50, 0xdc, 0xa9, 0x93, 0x57, 0x2e, 0xc7, 0x9e, 0xc3, 0x49, 0xf5, 0x1b,
	0xcb, 0x4f, 0x9c, 0x80, 0x07, 0x1d, 0x01, 0xd7, 0x84, 0x85, 0xa8, 0xeb, 0xec, 0x0c, 0x06, 0x96,
	0x24, 0x95, 0x96, 0x0f, 0xdc, 0xa0, 0x26, 0x62, 0x6f, 0x00, 0x2c, 0x49, 0x43, 0x18, 0xaf, 0x24,
	0xf1, 0xd3, 0x89, 0x37, 0x1b, 0x5d, 0x86, 0x51, 0xed, 0x4a, 0xd4, 0xba, 0x12, 0x2d, 0x5b, 0x57,
	0x44, 0xd0, 0xa0, 0x17, 0x54, 0x51, 0xcb, 0x22, 0x96, 0x0d, 0x75, 0xf8, 0x7f, 0x6a, 0x83, 0x5e,
	0x50, 0xf8, 0x15, 0xee, 0xfd, 0x65, 0x0c, 0xbb, 0x0f, 0xbd, 0x14, 0xf7, 0x8d, 0xc1, 0xd5, 0x93,
	0xbd, 0x84, 0x93, 0x9d, 0xdc, 0x96, 0xc8, 0x7d, 0xd7, 0xfa, 0xac, 0xf3, 0xb7, 0x03, 0x5d, 0xd4,
	0xa0, 0xb7, 0xfe, 0x6b, 0x6f, 0x2a, 0x20, 0x38, 0xe4, 0x19, 0x83, 0x7e, 0x2e, 0x33, 0x6c, 0x3a,
	0xba, 0x37, 0x7b, 0x74, 0xdc, 0x32, 0x68, 0xa8, 0xec, 0x1c, 0x02, 0x65, 0x57, 0x16, 0x37, 0x06,
	0xeb, 0x3d, 0x0c, 0xc5, 0x50, 0xd9, 0x6b, 0x17, 0x4f, 0x3f, 0x03, 0x54, 0x8b, 0xfc, 0xa8, 0x37,
	0x29, 0x1a, 0xf6, 0x0e, 0xc0, 0x60, 0xa2, 0x2c, 0x19, 0x85, 0x96, 0x7b, 0xce, 0xf4, 0xf3, 0x8e,
	0xb0, 0x1a, 0x28, 0x6a, 0xd0, 0x5e, 0x1c, 0xc1, 0xa7, 0xbf, 0x3d, 0x18, 0x77, 0xcb, 0x8c, 0xc3,
	0xa9, 0x8c, 0x63, 0x83, 0xd6, 0x36, 0x3a, 0xdb, 0xb0, 0xf2, 0xa3, 0xb4, 0xad, 0xd0, 0xea, 0xc9,
	0x3e, 0x41, 0x20, 0x4d, 0x52, 0x66, 0x98, 0x93, 0xe5, 0x3d, 0x37, 0xfa, 0xc5, 0x3f, 0x46, 0x47,
	0x8b, 0x16, 0x5c, 0x5f, 0xde, 0x0d, 0x99, 0xcd, 0xe1, 0xa1, 0x2c, 0xe9, 0x87, 0x36, 0xea, 0xa7,
	0x24, 0xa5, 0xf3, 0x15, 0xe9, 0x14, 0x73, 0x77, 0x60, 0x81, 0x60, 0x9d, 0xd2, 0xb2, 0xaa, 0x84,
	0xef, 0x61, 0xdc, 0xed, 0x76, 0xcb, 0xba, 0x6e, 0xf5, 0xb6, 0x5a, 0xcb, 0x87, 0xfe, 0x37, 0x7f,
	0x77, 0xb1, 0x1e, 0xb8, 0x93, 0x78, 0xf5, 0x27, 0x00, 0x00, 0xff, 0xff, 0x34, 0x31, 0xc2, 0xa6,
	0x98, 0x03, 0x00, 0x00,
}
