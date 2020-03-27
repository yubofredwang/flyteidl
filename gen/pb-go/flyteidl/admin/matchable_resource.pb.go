// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/matchable_resource.proto

package admin

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Defines a resource that can be configured by customizable Project-, ProjectDomain- or WorkflowAttributes
// based on matching tags.
type MatchableResource int32

const (
	// Applies to customizable task resource requests and limits.
	MatchableResource_TASK_RESOURCE MatchableResource = 0
	// Applies to configuring templated kubernetes cluster resources.
	MatchableResource_CLUSTER_RESOURCE MatchableResource = 1
	// Configures task and dynamic task execution queue assignment.
	MatchableResource_EXECUTION_QUEUE MatchableResource = 2
	// Configures the K8s cluster label to be used for execution to be run
	MatchableResource_EXECUTION_CLUSTER_LABEL MatchableResource = 3
)

var MatchableResource_name = map[int32]string{
	0: "TASK_RESOURCE",
	1: "CLUSTER_RESOURCE",
	2: "EXECUTION_QUEUE",
	3: "EXECUTION_CLUSTER_LABEL",
}

var MatchableResource_value = map[string]int32{
	"TASK_RESOURCE":           0,
	"CLUSTER_RESOURCE":        1,
	"EXECUTION_QUEUE":         2,
	"EXECUTION_CLUSTER_LABEL": 3,
}

func (x MatchableResource) String() string {
	return proto.EnumName(MatchableResource_name, int32(x))
}

func (MatchableResource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{0}
}

type TaskResourceSpec struct {
	Cpu                  string   `protobuf:"bytes,1,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Gpu                  string   `protobuf:"bytes,2,opt,name=gpu,proto3" json:"gpu,omitempty"`
	Memory               string   `protobuf:"bytes,3,opt,name=memory,proto3" json:"memory,omitempty"`
	Storage              string   `protobuf:"bytes,4,opt,name=storage,proto3" json:"storage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskResourceSpec) Reset()         { *m = TaskResourceSpec{} }
func (m *TaskResourceSpec) String() string { return proto.CompactTextString(m) }
func (*TaskResourceSpec) ProtoMessage()    {}
func (*TaskResourceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{0}
}

func (m *TaskResourceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskResourceSpec.Unmarshal(m, b)
}
func (m *TaskResourceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskResourceSpec.Marshal(b, m, deterministic)
}
func (m *TaskResourceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskResourceSpec.Merge(m, src)
}
func (m *TaskResourceSpec) XXX_Size() int {
	return xxx_messageInfo_TaskResourceSpec.Size(m)
}
func (m *TaskResourceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskResourceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_TaskResourceSpec proto.InternalMessageInfo

func (m *TaskResourceSpec) GetCpu() string {
	if m != nil {
		return m.Cpu
	}
	return ""
}

func (m *TaskResourceSpec) GetGpu() string {
	if m != nil {
		return m.Gpu
	}
	return ""
}

func (m *TaskResourceSpec) GetMemory() string {
	if m != nil {
		return m.Memory
	}
	return ""
}

func (m *TaskResourceSpec) GetStorage() string {
	if m != nil {
		return m.Storage
	}
	return ""
}

type TaskResourceAttributes struct {
	Defaults             *TaskResourceSpec `protobuf:"bytes,1,opt,name=defaults,proto3" json:"defaults,omitempty"`
	Limits               *TaskResourceSpec `protobuf:"bytes,2,opt,name=limits,proto3" json:"limits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TaskResourceAttributes) Reset()         { *m = TaskResourceAttributes{} }
func (m *TaskResourceAttributes) String() string { return proto.CompactTextString(m) }
func (*TaskResourceAttributes) ProtoMessage()    {}
func (*TaskResourceAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{1}
}

func (m *TaskResourceAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskResourceAttributes.Unmarshal(m, b)
}
func (m *TaskResourceAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskResourceAttributes.Marshal(b, m, deterministic)
}
func (m *TaskResourceAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskResourceAttributes.Merge(m, src)
}
func (m *TaskResourceAttributes) XXX_Size() int {
	return xxx_messageInfo_TaskResourceAttributes.Size(m)
}
func (m *TaskResourceAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskResourceAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_TaskResourceAttributes proto.InternalMessageInfo

func (m *TaskResourceAttributes) GetDefaults() *TaskResourceSpec {
	if m != nil {
		return m.Defaults
	}
	return nil
}

func (m *TaskResourceAttributes) GetLimits() *TaskResourceSpec {
	if m != nil {
		return m.Limits
	}
	return nil
}

type ClusterResourceAttributes struct {
	// Custom resource attributes which will be applied in cluster resource creation (e.g. quotas).
	// Map keys are the *case-sensitive* names of variables in templatized resource files.
	// Map values should be the custom values which get substituted during resource creation.
	Attributes           map[string]string `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ClusterResourceAttributes) Reset()         { *m = ClusterResourceAttributes{} }
func (m *ClusterResourceAttributes) String() string { return proto.CompactTextString(m) }
func (*ClusterResourceAttributes) ProtoMessage()    {}
func (*ClusterResourceAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{2}
}

func (m *ClusterResourceAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterResourceAttributes.Unmarshal(m, b)
}
func (m *ClusterResourceAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterResourceAttributes.Marshal(b, m, deterministic)
}
func (m *ClusterResourceAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterResourceAttributes.Merge(m, src)
}
func (m *ClusterResourceAttributes) XXX_Size() int {
	return xxx_messageInfo_ClusterResourceAttributes.Size(m)
}
func (m *ClusterResourceAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterResourceAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterResourceAttributes proto.InternalMessageInfo

func (m *ClusterResourceAttributes) GetAttributes() map[string]string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type ExecutionQueueAttributes struct {
	// Tags used for assigning execution queues for tasks defined within this project.
	Tags                 []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecutionQueueAttributes) Reset()         { *m = ExecutionQueueAttributes{} }
func (m *ExecutionQueueAttributes) String() string { return proto.CompactTextString(m) }
func (*ExecutionQueueAttributes) ProtoMessage()    {}
func (*ExecutionQueueAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{3}
}

func (m *ExecutionQueueAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionQueueAttributes.Unmarshal(m, b)
}
func (m *ExecutionQueueAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionQueueAttributes.Marshal(b, m, deterministic)
}
func (m *ExecutionQueueAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionQueueAttributes.Merge(m, src)
}
func (m *ExecutionQueueAttributes) XXX_Size() int {
	return xxx_messageInfo_ExecutionQueueAttributes.Size(m)
}
func (m *ExecutionQueueAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionQueueAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionQueueAttributes proto.InternalMessageInfo

func (m *ExecutionQueueAttributes) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type ExecutionClusterLabel struct {
	// Label value to determine where the execution will be run
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecutionClusterLabel) Reset()         { *m = ExecutionClusterLabel{} }
func (m *ExecutionClusterLabel) String() string { return proto.CompactTextString(m) }
func (*ExecutionClusterLabel) ProtoMessage()    {}
func (*ExecutionClusterLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{4}
}

func (m *ExecutionClusterLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionClusterLabel.Unmarshal(m, b)
}
func (m *ExecutionClusterLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionClusterLabel.Marshal(b, m, deterministic)
}
func (m *ExecutionClusterLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionClusterLabel.Merge(m, src)
}
func (m *ExecutionClusterLabel) XXX_Size() int {
	return xxx_messageInfo_ExecutionClusterLabel.Size(m)
}
func (m *ExecutionClusterLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionClusterLabel.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionClusterLabel proto.InternalMessageInfo

func (m *ExecutionClusterLabel) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Generic container for encapsulating all types of the above attributes messages.
type MatchingAttributes struct {
	// Types that are valid to be assigned to Target:
	//	*MatchingAttributes_TaskResourceAttributes
	//	*MatchingAttributes_ClusterResourceAttributes
	//	*MatchingAttributes_ExecutionQueueAttributes
	//	*MatchingAttributes_ExecutionClusterLabel
	Target               isMatchingAttributes_Target `protobuf_oneof:"target"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *MatchingAttributes) Reset()         { *m = MatchingAttributes{} }
func (m *MatchingAttributes) String() string { return proto.CompactTextString(m) }
func (*MatchingAttributes) ProtoMessage()    {}
func (*MatchingAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{5}
}

func (m *MatchingAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchingAttributes.Unmarshal(m, b)
}
func (m *MatchingAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchingAttributes.Marshal(b, m, deterministic)
}
func (m *MatchingAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchingAttributes.Merge(m, src)
}
func (m *MatchingAttributes) XXX_Size() int {
	return xxx_messageInfo_MatchingAttributes.Size(m)
}
func (m *MatchingAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchingAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_MatchingAttributes proto.InternalMessageInfo

type isMatchingAttributes_Target interface {
	isMatchingAttributes_Target()
}

type MatchingAttributes_TaskResourceAttributes struct {
	TaskResourceAttributes *TaskResourceAttributes `protobuf:"bytes,1,opt,name=task_resource_attributes,json=taskResourceAttributes,proto3,oneof"`
}

type MatchingAttributes_ClusterResourceAttributes struct {
	ClusterResourceAttributes *ClusterResourceAttributes `protobuf:"bytes,2,opt,name=cluster_resource_attributes,json=clusterResourceAttributes,proto3,oneof"`
}

type MatchingAttributes_ExecutionQueueAttributes struct {
	ExecutionQueueAttributes *ExecutionQueueAttributes `protobuf:"bytes,3,opt,name=execution_queue_attributes,json=executionQueueAttributes,proto3,oneof"`
}

type MatchingAttributes_ExecutionClusterLabel struct {
	ExecutionClusterLabel *ExecutionClusterLabel `protobuf:"bytes,4,opt,name=execution_cluster_label,json=executionClusterLabel,proto3,oneof"`
}

func (*MatchingAttributes_TaskResourceAttributes) isMatchingAttributes_Target() {}

func (*MatchingAttributes_ClusterResourceAttributes) isMatchingAttributes_Target() {}

func (*MatchingAttributes_ExecutionQueueAttributes) isMatchingAttributes_Target() {}

func (*MatchingAttributes_ExecutionClusterLabel) isMatchingAttributes_Target() {}

func (m *MatchingAttributes) GetTarget() isMatchingAttributes_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *MatchingAttributes) GetTaskResourceAttributes() *TaskResourceAttributes {
	if x, ok := m.GetTarget().(*MatchingAttributes_TaskResourceAttributes); ok {
		return x.TaskResourceAttributes
	}
	return nil
}

func (m *MatchingAttributes) GetClusterResourceAttributes() *ClusterResourceAttributes {
	if x, ok := m.GetTarget().(*MatchingAttributes_ClusterResourceAttributes); ok {
		return x.ClusterResourceAttributes
	}
	return nil
}

func (m *MatchingAttributes) GetExecutionQueueAttributes() *ExecutionQueueAttributes {
	if x, ok := m.GetTarget().(*MatchingAttributes_ExecutionQueueAttributes); ok {
		return x.ExecutionQueueAttributes
	}
	return nil
}

func (m *MatchingAttributes) GetExecutionClusterLabel() *ExecutionClusterLabel {
	if x, ok := m.GetTarget().(*MatchingAttributes_ExecutionClusterLabel); ok {
		return x.ExecutionClusterLabel
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MatchingAttributes) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MatchingAttributes_TaskResourceAttributes)(nil),
		(*MatchingAttributes_ClusterResourceAttributes)(nil),
		(*MatchingAttributes_ExecutionQueueAttributes)(nil),
		(*MatchingAttributes_ExecutionClusterLabel)(nil),
	}
}

// Represents a custom set of attributes applied for either a domain; a domain and project; or
// domain, project and workflow name.
type MatchableAttributesConfiguration struct {
	Attributes           *MatchingAttributes `protobuf:"bytes,1,opt,name=attributes,proto3" json:"attributes,omitempty"`
	Domain               string              `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Project              string              `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	Workflow             string              `protobuf:"bytes,4,opt,name=workflow,proto3" json:"workflow,omitempty"`
	LaunchPlan           string              `protobuf:"bytes,5,opt,name=launch_plan,json=launchPlan,proto3" json:"launch_plan,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MatchableAttributesConfiguration) Reset()         { *m = MatchableAttributesConfiguration{} }
func (m *MatchableAttributesConfiguration) String() string { return proto.CompactTextString(m) }
func (*MatchableAttributesConfiguration) ProtoMessage()    {}
func (*MatchableAttributesConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{6}
}

func (m *MatchableAttributesConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchableAttributesConfiguration.Unmarshal(m, b)
}
func (m *MatchableAttributesConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchableAttributesConfiguration.Marshal(b, m, deterministic)
}
func (m *MatchableAttributesConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchableAttributesConfiguration.Merge(m, src)
}
func (m *MatchableAttributesConfiguration) XXX_Size() int {
	return xxx_messageInfo_MatchableAttributesConfiguration.Size(m)
}
func (m *MatchableAttributesConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchableAttributesConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_MatchableAttributesConfiguration proto.InternalMessageInfo

func (m *MatchableAttributesConfiguration) GetAttributes() *MatchingAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *MatchableAttributesConfiguration) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *MatchableAttributesConfiguration) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *MatchableAttributesConfiguration) GetWorkflow() string {
	if m != nil {
		return m.Workflow
	}
	return ""
}

func (m *MatchableAttributesConfiguration) GetLaunchPlan() string {
	if m != nil {
		return m.LaunchPlan
	}
	return ""
}

// Request all matching resource attributes.
type ListMatchableAttributesRequest struct {
	ResourceType         MatchableResource `protobuf:"varint,1,opt,name=resource_type,json=resourceType,proto3,enum=flyteidl.admin.MatchableResource" json:"resource_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListMatchableAttributesRequest) Reset()         { *m = ListMatchableAttributesRequest{} }
func (m *ListMatchableAttributesRequest) String() string { return proto.CompactTextString(m) }
func (*ListMatchableAttributesRequest) ProtoMessage()    {}
func (*ListMatchableAttributesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{7}
}

func (m *ListMatchableAttributesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMatchableAttributesRequest.Unmarshal(m, b)
}
func (m *ListMatchableAttributesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMatchableAttributesRequest.Marshal(b, m, deterministic)
}
func (m *ListMatchableAttributesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMatchableAttributesRequest.Merge(m, src)
}
func (m *ListMatchableAttributesRequest) XXX_Size() int {
	return xxx_messageInfo_ListMatchableAttributesRequest.Size(m)
}
func (m *ListMatchableAttributesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMatchableAttributesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListMatchableAttributesRequest proto.InternalMessageInfo

func (m *ListMatchableAttributesRequest) GetResourceType() MatchableResource {
	if m != nil {
		return m.ResourceType
	}
	return MatchableResource_TASK_RESOURCE
}

// Response for a request for all matching resource attributes.
type ListMatchableAttributesResponse struct {
	Configurations       []*MatchableAttributesConfiguration `protobuf:"bytes,1,rep,name=configurations,proto3" json:"configurations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *ListMatchableAttributesResponse) Reset()         { *m = ListMatchableAttributesResponse{} }
func (m *ListMatchableAttributesResponse) String() string { return proto.CompactTextString(m) }
func (*ListMatchableAttributesResponse) ProtoMessage()    {}
func (*ListMatchableAttributesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{8}
}

func (m *ListMatchableAttributesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMatchableAttributesResponse.Unmarshal(m, b)
}
func (m *ListMatchableAttributesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMatchableAttributesResponse.Marshal(b, m, deterministic)
}
func (m *ListMatchableAttributesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMatchableAttributesResponse.Merge(m, src)
}
func (m *ListMatchableAttributesResponse) XXX_Size() int {
	return xxx_messageInfo_ListMatchableAttributesResponse.Size(m)
}
func (m *ListMatchableAttributesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMatchableAttributesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListMatchableAttributesResponse proto.InternalMessageInfo

func (m *ListMatchableAttributesResponse) GetConfigurations() []*MatchableAttributesConfiguration {
	if m != nil {
		return m.Configurations
	}
	return nil
}

func init() {
	proto.RegisterEnum("flyteidl.admin.MatchableResource", MatchableResource_name, MatchableResource_value)
	proto.RegisterType((*TaskResourceSpec)(nil), "flyteidl.admin.TaskResourceSpec")
	proto.RegisterType((*TaskResourceAttributes)(nil), "flyteidl.admin.TaskResourceAttributes")
	proto.RegisterType((*ClusterResourceAttributes)(nil), "flyteidl.admin.ClusterResourceAttributes")
	proto.RegisterMapType((map[string]string)(nil), "flyteidl.admin.ClusterResourceAttributes.AttributesEntry")
	proto.RegisterType((*ExecutionQueueAttributes)(nil), "flyteidl.admin.ExecutionQueueAttributes")
	proto.RegisterType((*ExecutionClusterLabel)(nil), "flyteidl.admin.ExecutionClusterLabel")
	proto.RegisterType((*MatchingAttributes)(nil), "flyteidl.admin.MatchingAttributes")
	proto.RegisterType((*MatchableAttributesConfiguration)(nil), "flyteidl.admin.MatchableAttributesConfiguration")
	proto.RegisterType((*ListMatchableAttributesRequest)(nil), "flyteidl.admin.ListMatchableAttributesRequest")
	proto.RegisterType((*ListMatchableAttributesResponse)(nil), "flyteidl.admin.ListMatchableAttributesResponse")
}

func init() {
	proto.RegisterFile("flyteidl/admin/matchable_resource.proto", fileDescriptor_1d15bcabb02640f4)
}

var fileDescriptor_1d15bcabb02640f4 = []byte{
	// 712 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x4e, 0xdb, 0x4a,
	0x10, 0x8e, 0x09, 0xe4, 0xc0, 0xe4, 0x00, 0x61, 0x0f, 0x3f, 0x06, 0xa4, 0x43, 0x8e, 0xa5, 0xd3,
	0xd2, 0x4a, 0x24, 0x6d, 0xb8, 0xa1, 0x55, 0x7b, 0x41, 0x22, 0x57, 0xa9, 0x1a, 0x4a, 0xd9, 0x24,
	0x12, 0xed, 0x4d, 0xb4, 0x71, 0x36, 0x8e, 0x9b, 0x8d, 0xd7, 0xd8, 0xbb, 0xa5, 0x56, 0x5f, 0xa2,
	0x2f, 0xd3, 0x9b, 0xbe, 0x47, 0xdf, 0xa7, 0xb2, 0x63, 0x27, 0xc6, 0xc4, 0x15, 0x77, 0x9e, 0xd9,
	0x99, 0xf9, 0xbe, 0x9d, 0xf9, 0x3c, 0x0b, 0x8f, 0x87, 0xcc, 0x17, 0xd4, 0x1a, 0xb0, 0x2a, 0x19,
	0x4c, 0x2c, 0xbb, 0x3a, 0x21, 0xc2, 0x18, 0x91, 0x3e, 0xa3, 0x3d, 0x97, 0x7a, 0x5c, 0xba, 0x06,
	0xad, 0x38, 0x2e, 0x17, 0x1c, 0x6d, 0xc4, 0x81, 0x95, 0x30, 0x50, 0x1b, 0x41, 0xa9, 0x43, 0xbc,
	0x31, 0x8e, 0xa2, 0xda, 0x0e, 0x35, 0x50, 0x09, 0xf2, 0x86, 0x23, 0x55, 0xa5, 0xac, 0x1c, 0xaf,
	0xe1, 0xe0, 0x33, 0xf0, 0x98, 0x8e, 0x54, 0x97, 0xa6, 0x1e, 0xd3, 0x91, 0x68, 0x17, 0x0a, 0x13,
	0x3a, 0xe1, 0xae, 0xaf, 0xe6, 0x43, 0x67, 0x64, 0x21, 0x15, 0xfe, 0xf2, 0x04, 0x77, 0x89, 0x49,
	0xd5, 0xe5, 0xf0, 0x20, 0x36, 0xb5, 0xef, 0x0a, 0xec, 0x26, 0xa1, 0xce, 0x85, 0x70, 0xad, 0xbe,
	0x14, 0xd4, 0x43, 0xaf, 0x60, 0x75, 0x40, 0x87, 0x44, 0x32, 0xe1, 0x85, 0xa8, 0xc5, 0x5a, 0xb9,
	0x72, 0x97, 0x67, 0x25, 0x4d, 0x12, 0xcf, 0x32, 0xd0, 0x19, 0x14, 0x98, 0x35, 0xb1, 0x84, 0x17,
	0xf2, 0x7b, 0x48, 0x6e, 0x14, 0xaf, 0xfd, 0x50, 0x60, 0xbf, 0xc1, 0xa4, 0x27, 0xa8, 0xbb, 0x80,
	0xd5, 0x47, 0x00, 0x32, 0xb3, 0x54, 0xa5, 0x9c, 0x3f, 0x2e, 0xd6, 0x5e, 0xa4, 0x6b, 0x67, 0xa6,
	0x57, 0xe6, 0x9f, 0xba, 0x2d, 0x5c, 0x1f, 0x27, 0x8a, 0x1d, 0xbc, 0x86, 0xcd, 0xd4, 0x71, 0xd0,
	0xe2, 0x31, 0xf5, 0xe3, 0xa6, 0x8f, 0xa9, 0x8f, 0xb6, 0x61, 0xe5, 0x0b, 0x61, 0x92, 0x46, 0x6d,
	0x9f, 0x1a, 0x2f, 0x97, 0xce, 0x14, 0xad, 0x02, 0xaa, 0xfe, 0x95, 0x1a, 0x52, 0x58, 0xdc, 0xbe,
	0x92, 0x54, 0x26, 0x59, 0x23, 0x58, 0x16, 0xc4, 0x9c, 0xf2, 0x5d, 0xc3, 0xe1, 0xb7, 0x76, 0x02,
	0x3b, 0xb3, 0xf8, 0x88, 0x70, 0x8b, 0xf4, 0x29, 0x9b, 0x43, 0x28, 0x09, 0x08, 0xed, 0x67, 0x1e,
	0xd0, 0x45, 0x20, 0x20, 0xcb, 0x36, 0x13, 0x95, 0xfb, 0xa0, 0x0a, 0xe2, 0x8d, 0x67, 0x8a, 0xea,
	0xdd, 0xe9, 0x4e, 0xd0, 0xf9, 0x47, 0x7f, 0xea, 0xfc, 0xbc, 0x52, 0x33, 0x87, 0x77, 0xc5, 0x62,
	0x25, 0x8c, 0xe1, 0xd0, 0x98, 0x12, 0x5c, 0x08, 0x33, 0x1d, 0xf0, 0x93, 0x07, 0x0f, 0xa1, 0x99,
	0xc3, 0xfb, 0x46, 0xe6, 0x80, 0x47, 0x70, 0x40, 0xe3, 0xb6, 0xf4, 0x6e, 0x82, 0x3e, 0x26, 0xb1,
	0xf2, 0x21, 0xd6, 0x71, 0x1a, 0x2b, 0xab, 0xf1, 0xcd, 0x1c, 0x56, 0x69, 0xd6, 0x50, 0x7a, 0xb0,
	0x37, 0x47, 0x8a, 0x2f, 0xc8, 0x82, 0x11, 0x84, 0x7f, 0x49, 0xb1, 0xf6, 0x7f, 0x26, 0x4c, 0x72,
	0x5e, 0xcd, 0x1c, 0xde, 0xa1, 0x8b, 0x0e, 0xea, 0xab, 0x50, 0x10, 0xc4, 0x35, 0xa9, 0xd0, 0x7e,
	0x29, 0x50, 0xbe, 0x88, 0xff, 0xfe, 0x39, 0x85, 0x06, 0xb7, 0x87, 0x96, 0x29, 0x5d, 0x12, 0xa4,
	0xa2, 0x7a, 0x4a, 0xda, 0x01, 0x05, 0x2d, 0x4d, 0xe1, 0xbe, 0x04, 0x92, 0x1a, 0x0e, 0x36, 0xc0,
	0x80, 0x4f, 0x88, 0x65, 0x47, 0xfa, 0x8c, 0xac, 0x60, 0x03, 0x38, 0x2e, 0xff, 0x4c, 0x0d, 0x11,
	0xad, 0x86, 0xd8, 0x44, 0x07, 0xb0, 0x7a, 0xcb, 0xdd, 0xf1, 0x90, 0xf1, 0xdb, 0x68, 0x39, 0xcc,
	0x6c, 0x74, 0x04, 0x45, 0x46, 0xa4, 0x6d, 0x8c, 0x7a, 0x0e, 0x23, 0xb6, 0xba, 0x12, 0x1e, 0xc3,
	0xd4, 0xf5, 0x81, 0x91, 0x60, 0x51, 0xfd, 0xdb, 0xb2, 0x3c, 0xb1, 0xe0, 0x6a, 0x98, 0xde, 0x48,
	0xea, 0x09, 0xf4, 0x06, 0xd6, 0x67, 0x9a, 0x11, 0xbe, 0x33, 0x15, 0xf5, 0x46, 0xed, 0xbf, 0x85,
	0xf7, 0x0a, 0x4a, 0xc4, 0x92, 0xc0, 0x7f, 0xc7, 0x79, 0x1d, 0xdf, 0xa1, 0xda, 0x37, 0x38, 0xca,
	0x44, 0xf2, 0x1c, 0x6e, 0x7b, 0x14, 0x5d, 0xc3, 0x86, 0x91, 0x6c, 0x68, 0xbc, 0x1e, 0x9e, 0x65,
	0x62, 0x65, 0x4c, 0x02, 0xa7, 0xea, 0x3c, 0xb5, 0x61, 0xeb, 0x1e, 0x3f, 0xb4, 0x05, 0xeb, 0x9d,
	0xf3, 0xf6, 0xbb, 0x1e, 0xd6, 0xdb, 0x97, 0x5d, 0xdc, 0xd0, 0x4b, 0x39, 0xb4, 0x0d, 0xa5, 0x46,
	0xab, 0xdb, 0xee, 0xe8, 0x78, 0xee, 0x55, 0xd0, 0x3f, 0xb0, 0xa9, 0x5f, 0xeb, 0x8d, 0x6e, 0xe7,
	0xed, 0xe5, 0xfb, 0xde, 0x55, 0x57, 0xef, 0xea, 0xa5, 0x25, 0x74, 0x08, 0x7b, 0x73, 0x67, 0x9c,
	0xd4, 0x3a, 0xaf, 0xeb, 0xad, 0x52, 0xbe, 0x7e, 0xfa, 0xe9, 0xb9, 0x69, 0x89, 0x91, 0xec, 0x57,
	0x0c, 0x3e, 0xa9, 0x32, 0x7f, 0x28, 0xaa, 0xb3, 0xa7, 0xc4, 0xa4, 0x76, 0xd5, 0xe9, 0x9f, 0x98,
	0xbc, 0x7a, 0xf7, 0x75, 0xe9, 0x17, 0xc2, 0xb7, 0xe4, 0xf4, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x18, 0x84, 0xb1, 0x80, 0x76, 0x06, 0x00, 0x00,
}