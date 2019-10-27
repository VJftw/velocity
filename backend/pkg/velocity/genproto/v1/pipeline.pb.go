// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pipeline.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Pipeline struct {
	// PIPELINE_ID = commit_sha+pipeline_name
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProjectId            string   `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	CommitId             string   `protobuf:"bytes,3,opt,name=commit_id,json=commitId,proto3" json:"commit_id,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Stages               []*Stage `protobuf:"bytes,6,rep,name=stages,proto3" json:"stages,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pipeline) Reset()         { *m = Pipeline{} }
func (m *Pipeline) String() string { return proto.CompactTextString(m) }
func (*Pipeline) ProtoMessage()    {}
func (*Pipeline) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ac67a7adf3df9c7, []int{0}
}

func (m *Pipeline) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pipeline.Unmarshal(m, b)
}
func (m *Pipeline) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pipeline.Marshal(b, m, deterministic)
}
func (m *Pipeline) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pipeline.Merge(m, src)
}
func (m *Pipeline) XXX_Size() int {
	return xxx_messageInfo_Pipeline.Size(m)
}
func (m *Pipeline) XXX_DiscardUnknown() {
	xxx_messageInfo_Pipeline.DiscardUnknown(m)
}

var xxx_messageInfo_Pipeline proto.InternalMessageInfo

func (m *Pipeline) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Pipeline) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *Pipeline) GetCommitId() string {
	if m != nil {
		return m.CommitId
	}
	return ""
}

func (m *Pipeline) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Pipeline) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Pipeline) GetStages() []*Stage {
	if m != nil {
		return m.Stages
	}
	return nil
}

type Stage struct {
	Name                 string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Blueprints           []*Blueprint `protobuf:"bytes,2,rep,name=blueprints,proto3" json:"blueprints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Stage) Reset()         { *m = Stage{} }
func (m *Stage) String() string { return proto.CompactTextString(m) }
func (*Stage) ProtoMessage()    {}
func (*Stage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ac67a7adf3df9c7, []int{1}
}

func (m *Stage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stage.Unmarshal(m, b)
}
func (m *Stage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stage.Marshal(b, m, deterministic)
}
func (m *Stage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stage.Merge(m, src)
}
func (m *Stage) XXX_Size() int {
	return xxx_messageInfo_Stage.Size(m)
}
func (m *Stage) XXX_DiscardUnknown() {
	xxx_messageInfo_Stage.DiscardUnknown(m)
}

var xxx_messageInfo_Stage proto.InternalMessageInfo

func (m *Stage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Stage) GetBlueprints() []*Blueprint {
	if m != nil {
		return m.Blueprints
	}
	return nil
}

type PipelineQuery struct {
	Ids                  []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PipelineQuery) Reset()         { *m = PipelineQuery{} }
func (m *PipelineQuery) String() string { return proto.CompactTextString(m) }
func (*PipelineQuery) ProtoMessage()    {}
func (*PipelineQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ac67a7adf3df9c7, []int{2}
}

func (m *PipelineQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PipelineQuery.Unmarshal(m, b)
}
func (m *PipelineQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PipelineQuery.Marshal(b, m, deterministic)
}
func (m *PipelineQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PipelineQuery.Merge(m, src)
}
func (m *PipelineQuery) XXX_Size() int {
	return xxx_messageInfo_PipelineQuery.Size(m)
}
func (m *PipelineQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_PipelineQuery.DiscardUnknown(m)
}

var xxx_messageInfo_PipelineQuery proto.InternalMessageInfo

func (m *PipelineQuery) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type GetPipelineRequest struct {
	// The id of the project in the form of
	// `[PROJECT_ID]`.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// The id of the Pipeline in the form of
	// `[PIPELINE_ID]`.
	PipelineId           string   `protobuf:"bytes,2,opt,name=Pipeline_id,json=PipelineId,proto3" json:"Pipeline_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPipelineRequest) Reset()         { *m = GetPipelineRequest{} }
func (m *GetPipelineRequest) String() string { return proto.CompactTextString(m) }
func (*GetPipelineRequest) ProtoMessage()    {}
func (*GetPipelineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ac67a7adf3df9c7, []int{3}
}

func (m *GetPipelineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPipelineRequest.Unmarshal(m, b)
}
func (m *GetPipelineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPipelineRequest.Marshal(b, m, deterministic)
}
func (m *GetPipelineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPipelineRequest.Merge(m, src)
}
func (m *GetPipelineRequest) XXX_Size() int {
	return xxx_messageInfo_GetPipelineRequest.Size(m)
}
func (m *GetPipelineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPipelineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPipelineRequest proto.InternalMessageInfo

func (m *GetPipelineRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *GetPipelineRequest) GetPipelineId() string {
	if m != nil {
		return m.PipelineId
	}
	return ""
}

type ListPipelinesRequest struct {
	RepoQuery            *RepoQuery `protobuf:"bytes,1,opt,name=repo_query,json=repoQuery,proto3" json:"repo_query,omitempty"`
	PageQuery            *PageQuery `protobuf:"bytes,99,opt,name=page_query,json=pageQuery,proto3" json:"page_query,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListPipelinesRequest) Reset()         { *m = ListPipelinesRequest{} }
func (m *ListPipelinesRequest) String() string { return proto.CompactTextString(m) }
func (*ListPipelinesRequest) ProtoMessage()    {}
func (*ListPipelinesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ac67a7adf3df9c7, []int{4}
}

func (m *ListPipelinesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPipelinesRequest.Unmarshal(m, b)
}
func (m *ListPipelinesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPipelinesRequest.Marshal(b, m, deterministic)
}
func (m *ListPipelinesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPipelinesRequest.Merge(m, src)
}
func (m *ListPipelinesRequest) XXX_Size() int {
	return xxx_messageInfo_ListPipelinesRequest.Size(m)
}
func (m *ListPipelinesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPipelinesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPipelinesRequest proto.InternalMessageInfo

func (m *ListPipelinesRequest) GetRepoQuery() *RepoQuery {
	if m != nil {
		return m.RepoQuery
	}
	return nil
}

func (m *ListPipelinesRequest) GetPageQuery() *PageQuery {
	if m != nil {
		return m.PageQuery
	}
	return nil
}

type ListPipelinesResponse struct {
	Pipelines            []*Pipeline `protobuf:"bytes,1,rep,name=Pipelines,proto3" json:"Pipelines,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListPipelinesResponse) Reset()         { *m = ListPipelinesResponse{} }
func (m *ListPipelinesResponse) String() string { return proto.CompactTextString(m) }
func (*ListPipelinesResponse) ProtoMessage()    {}
func (*ListPipelinesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ac67a7adf3df9c7, []int{5}
}

func (m *ListPipelinesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPipelinesResponse.Unmarshal(m, b)
}
func (m *ListPipelinesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPipelinesResponse.Marshal(b, m, deterministic)
}
func (m *ListPipelinesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPipelinesResponse.Merge(m, src)
}
func (m *ListPipelinesResponse) XXX_Size() int {
	return xxx_messageInfo_ListPipelinesResponse.Size(m)
}
func (m *ListPipelinesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPipelinesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPipelinesResponse proto.InternalMessageInfo

func (m *ListPipelinesResponse) GetPipelines() []*Pipeline {
	if m != nil {
		return m.Pipelines
	}
	return nil
}

func init() {
	proto.RegisterType((*Pipeline)(nil), "velocity.v1.Pipeline")
	proto.RegisterType((*Stage)(nil), "velocity.v1.Stage")
	proto.RegisterType((*PipelineQuery)(nil), "velocity.v1.PipelineQuery")
	proto.RegisterType((*GetPipelineRequest)(nil), "velocity.v1.GetPipelineRequest")
	proto.RegisterType((*ListPipelinesRequest)(nil), "velocity.v1.ListPipelinesRequest")
	proto.RegisterType((*ListPipelinesResponse)(nil), "velocity.v1.ListPipelinesResponse")
}

func init() { proto.RegisterFile("pipeline.proto", fileDescriptor_7ac67a7adf3df9c7) }

var fileDescriptor_7ac67a7adf3df9c7 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xcd, 0x6a, 0xdb, 0x40,
	0x10, 0x46, 0xb2, 0x63, 0xa2, 0x11, 0xf9, 0x61, 0x68, 0x8a, 0x70, 0x5b, 0x6c, 0xab, 0x87, 0xba,
	0x3e, 0x58, 0xd8, 0xa1, 0xbd, 0xf5, 0x92, 0x4b, 0x31, 0xe4, 0x90, 0xca, 0x3d, 0xf5, 0x12, 0x14,
	0x69, 0x10, 0x5b, 0x6c, 0xed, 0x46, 0xbb, 0x16, 0x98, 0x12, 0x0a, 0xc5, 0x6f, 0xd0, 0x97, 0xe9,
	0x7b, 0xf4, 0x15, 0xfa, 0x20, 0x45, 0x6b, 0xad, 0x2d, 0xb9, 0xa6, 0xb7, 0xd1, 0xf7, 0xcd, 0x7c,
	0xf3, 0xcd, 0x27, 0x16, 0xce, 0x05, 0x13, 0xb4, 0x60, 0x19, 0x8d, 0x45, 0xce, 0x15, 0x47, 0xb7,
	0xa0, 0x05, 0x8f, 0x99, 0x5a, 0x8f, 0x8b, 0x49, 0xd7, 0x89, 0x04, 0xdb, 0xe2, 0xdd, 0x8b, 0x87,
	0xc5, 0x8a, 0x44, 0xce, 0x32, 0x55, 0x01, 0x97, 0x39, 0x09, 0x2e, 0x99, 0xe2, 0xf9, 0xba, 0x42,
	0x5e, 0xa6, 0x9c, 0xa7, 0x0b, 0x0a, 0x22, 0xc1, 0x82, 0x28, 0xcb, 0xb8, 0x8a, 0x14, 0xe3, 0x99,
	0xdc, 0xb2, 0xfe, 0x2f, 0x0b, 0x4e, 0xef, 0xaa, 0x5d, 0x78, 0x0e, 0x36, 0x4b, 0x3c, 0xab, 0x6f,
	0x0d, 0x9d, 0xd0, 0x66, 0x09, 0xbe, 0x02, 0x10, 0x39, 0xff, 0x4a, 0xb1, 0xba, 0x67, 0x89, 0x67,
	0x6b, 0xdc, 0xa9, 0x90, 0x59, 0x82, 0x2f, 0xc0, 0x89, 0xf9, 0x72, 0xc9, 0x34, 0xdb, 0xd2, 0xec,
	0xe9, 0x16, 0x98, 0x25, 0x88, 0xd0, 0xce, 0xa2, 0x25, 0x79, 0x6d, 0x8d, 0xeb, 0x1a, 0xfb, 0xe0,
	0x26, 0x24, 0xe3, 0x9c, 0x89, 0xd2, 0x82, 0x77, 0xa2, 0xa9, 0x3a, 0x84, 0x23, 0xe8, 0x48, 0x15,
	0xa5, 0x24, 0xbd, 0x4e, 0xbf, 0x35, 0x74, 0xa7, 0x38, 0xae, 0x1d, 0x3e, 0x9e, 0x97, 0x54, 0x58,
	0x75, 0xf8, 0x73, 0x38, 0xd1, 0xc0, 0x6e, 0x95, 0x55, 0x5b, 0xf5, 0x1e, 0x60, 0x17, 0x8d, 0xf4,
	0x6c, 0x2d, 0xf6, 0xbc, 0x21, 0x76, 0x63, 0xe8, 0xb0, 0xd6, 0xe9, 0x0f, 0xe0, 0xcc, 0xc4, 0xf1,
	0x69, 0x45, 0xf9, 0x1a, 0x2f, 0xa1, 0xc5, 0x12, 0xe9, 0x59, 0xfd, 0xd6, 0xd0, 0x09, 0xcb, 0xd2,
	0xff, 0x0c, 0xf8, 0x91, 0x94, 0xe9, 0x0a, 0xe9, 0x71, 0x45, 0x52, 0x1d, 0x64, 0x65, 0x1d, 0x66,
	0xd5, 0x03, 0xd7, 0x4c, 0xec, 0xb3, 0x04, 0x03, 0xcd, 0x12, 0x7f, 0x63, 0xc1, 0xb3, 0x5b, 0x26,
	0x77, 0xba, 0xd2, 0x08, 0xbf, 0x03, 0x28, 0xff, 0xe9, 0xfd, 0x63, 0x69, 0x47, 0x0b, 0x1f, 0x5e,
	0x12, 0x92, 0xe0, 0xda, 0x6c, 0xe8, 0xe4, 0xa6, 0x2c, 0xc7, 0x44, 0x94, 0x52, 0x35, 0x16, 0x1f,
	0x19, 0xbb, 0x8b, 0x52, 0xaa, 0xc6, 0x84, 0x29, 0xfd, 0x5b, 0xb8, 0x3a, 0x70, 0x21, 0x05, 0xcf,
	0x24, 0xe1, 0x35, 0x38, 0x3b, 0x50, 0xa7, 0xe1, 0x4e, 0xaf, 0x9a, 0x72, 0x26, 0x90, 0x7d, 0xdf,
	0x74, 0x63, 0xc3, 0x85, 0xf9, 0x9a, 0x53, 0x5e, 0xb0, 0x98, 0x50, 0x80, 0x5b, 0x8b, 0x0f, 0x7b,
	0x0d, 0x91, 0x7f, 0x83, 0xed, 0x1e, 0xdf, 0xe2, 0xbf, 0xfd, 0xf1, 0xfb, 0xcf, 0x4f, 0xfb, 0x35,
	0x0e, 0x82, 0x62, 0x12, 0x7c, 0x2b, 0xff, 0xf9, 0x87, 0x2a, 0x6d, 0x19, 0x8c, 0x02, 0xf3, 0x7c,
	0x64, 0x30, 0x7a, 0xc2, 0xef, 0x70, 0xd6, 0xb8, 0x09, 0x07, 0x0d, 0xc9, 0x63, 0xa9, 0x77, 0xfd,
	0xff, 0xb5, 0x6c, 0x23, 0xf1, 0xdf, 0x68, 0x0b, 0x03, 0xec, 0x1d, 0xb3, 0xf0, 0xb4, 0xf7, 0x70,
	0xd3, 0xfe, 0x62, 0x17, 0x93, 0x87, 0x8e, 0x7e, 0x71, 0xd7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x3f, 0xf1, 0xa8, 0x66, 0xdc, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PipelineServiceClient is the client API for PipelineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PipelineServiceClient interface {
	GetPipeline(ctx context.Context, in *GetPipelineRequest, opts ...grpc.CallOption) (*Pipeline, error)
	ListPipelines(ctx context.Context, in *ListPipelinesRequest, opts ...grpc.CallOption) (*ListPipelinesResponse, error)
}

type pipelineServiceClient struct {
	cc *grpc.ClientConn
}

func NewPipelineServiceClient(cc *grpc.ClientConn) PipelineServiceClient {
	return &pipelineServiceClient{cc}
}

func (c *pipelineServiceClient) GetPipeline(ctx context.Context, in *GetPipelineRequest, opts ...grpc.CallOption) (*Pipeline, error) {
	out := new(Pipeline)
	err := c.cc.Invoke(ctx, "/velocity.v1.PipelineService/GetPipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) ListPipelines(ctx context.Context, in *ListPipelinesRequest, opts ...grpc.CallOption) (*ListPipelinesResponse, error) {
	out := new(ListPipelinesResponse)
	err := c.cc.Invoke(ctx, "/velocity.v1.PipelineService/ListPipelines", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PipelineServiceServer is the server API for PipelineService service.
type PipelineServiceServer interface {
	GetPipeline(context.Context, *GetPipelineRequest) (*Pipeline, error)
	ListPipelines(context.Context, *ListPipelinesRequest) (*ListPipelinesResponse, error)
}

// UnimplementedPipelineServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPipelineServiceServer struct {
}

func (*UnimplementedPipelineServiceServer) GetPipeline(ctx context.Context, req *GetPipelineRequest) (*Pipeline, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPipeline not implemented")
}
func (*UnimplementedPipelineServiceServer) ListPipelines(ctx context.Context, req *ListPipelinesRequest) (*ListPipelinesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPipelines not implemented")
}

func RegisterPipelineServiceServer(s *grpc.Server, srv PipelineServiceServer) {
	s.RegisterService(&_PipelineService_serviceDesc, srv)
}

func _PipelineService_GetPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).GetPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/velocity.v1.PipelineService/GetPipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).GetPipeline(ctx, req.(*GetPipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_ListPipelines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPipelinesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).ListPipelines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/velocity.v1.PipelineService/ListPipelines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).ListPipelines(ctx, req.(*ListPipelinesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PipelineService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "velocity.v1.PipelineService",
	HandlerType: (*PipelineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPipeline",
			Handler:    _PipelineService_GetPipeline_Handler,
		},
		{
			MethodName: "ListPipelines",
			Handler:    _PipelineService_ListPipelines_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pipeline.proto",
}
