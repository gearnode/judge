// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/judge/v1alpha1/judge.proto

package v1alpha1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
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

type GetPolicyRequest struct {
	// The orn of the policy to retrieve.
	PolicyId             string   `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPolicyRequest) Reset()         { *m = GetPolicyRequest{} }
func (m *GetPolicyRequest) String() string { return proto.CompactTextString(m) }
func (*GetPolicyRequest) ProtoMessage()    {}
func (*GetPolicyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8c4c1de41aa150c, []int{0}
}

func (m *GetPolicyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPolicyRequest.Unmarshal(m, b)
}
func (m *GetPolicyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPolicyRequest.Marshal(b, m, deterministic)
}
func (m *GetPolicyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPolicyRequest.Merge(m, src)
}
func (m *GetPolicyRequest) XXX_Size() int {
	return xxx_messageInfo_GetPolicyRequest.Size(m)
}
func (m *GetPolicyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPolicyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPolicyRequest proto.InternalMessageInfo

func (m *GetPolicyRequest) GetPolicyId() string {
	if m != nil {
		return m.PolicyId
	}
	return ""
}

type ListPoliciesRequest struct {
	// The maximum number of policies to return.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The `next_page_token` value returned from a previous List request, if any.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// The order to sort results by. For example: `priority desc, name`.
	OrderBy              string   `protobuf:"bytes,3,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPoliciesRequest) Reset()         { *m = ListPoliciesRequest{} }
func (m *ListPoliciesRequest) String() string { return proto.CompactTextString(m) }
func (*ListPoliciesRequest) ProtoMessage()    {}
func (*ListPoliciesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8c4c1de41aa150c, []int{1}
}

func (m *ListPoliciesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPoliciesRequest.Unmarshal(m, b)
}
func (m *ListPoliciesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPoliciesRequest.Marshal(b, m, deterministic)
}
func (m *ListPoliciesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPoliciesRequest.Merge(m, src)
}
func (m *ListPoliciesRequest) XXX_Size() int {
	return xxx_messageInfo_ListPoliciesRequest.Size(m)
}
func (m *ListPoliciesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPoliciesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPoliciesRequest proto.InternalMessageInfo

func (m *ListPoliciesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListPoliciesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListPoliciesRequest) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

type ListPoliciesResponse struct {
	// The Policies found.
	Policies []*Policy `protobuf:"bytes,1,rep,name=policies,proto3" json:"policies,omitempty"`
	// The next page token.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPoliciesResponse) Reset()         { *m = ListPoliciesResponse{} }
func (m *ListPoliciesResponse) String() string { return proto.CompactTextString(m) }
func (*ListPoliciesResponse) ProtoMessage()    {}
func (*ListPoliciesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8c4c1de41aa150c, []int{2}
}

func (m *ListPoliciesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPoliciesResponse.Unmarshal(m, b)
}
func (m *ListPoliciesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPoliciesResponse.Marshal(b, m, deterministic)
}
func (m *ListPoliciesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPoliciesResponse.Merge(m, src)
}
func (m *ListPoliciesResponse) XXX_Size() int {
	return xxx_messageInfo_ListPoliciesResponse.Size(m)
}
func (m *ListPoliciesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPoliciesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPoliciesResponse proto.InternalMessageInfo

func (m *ListPoliciesResponse) GetPolicies() []*Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

func (m *ListPoliciesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type CreatePolicyRequest struct {
	// The client-assigned policy object resource name to use for this document.
	//
	// Optional. If not specified, an ID will be assigned by the service.
	PolicyId string `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
	// The policy to create.
	Policy               *Policy  `protobuf:"bytes,2,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePolicyRequest) Reset()         { *m = CreatePolicyRequest{} }
func (m *CreatePolicyRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePolicyRequest) ProtoMessage()    {}
func (*CreatePolicyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8c4c1de41aa150c, []int{3}
}

func (m *CreatePolicyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePolicyRequest.Unmarshal(m, b)
}
func (m *CreatePolicyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePolicyRequest.Marshal(b, m, deterministic)
}
func (m *CreatePolicyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePolicyRequest.Merge(m, src)
}
func (m *CreatePolicyRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePolicyRequest.Size(m)
}
func (m *CreatePolicyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePolicyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePolicyRequest proto.InternalMessageInfo

func (m *CreatePolicyRequest) GetPolicyId() string {
	if m != nil {
		return m.PolicyId
	}
	return ""
}

func (m *CreatePolicyRequest) GetPolicy() *Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type UpdatePolicyRequest struct {
	// The updated policy.
	// Creates the policy if it does not already exist.
	Policy               *Policy  `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePolicyRequest) Reset()         { *m = UpdatePolicyRequest{} }
func (m *UpdatePolicyRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePolicyRequest) ProtoMessage()    {}
func (*UpdatePolicyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8c4c1de41aa150c, []int{4}
}

func (m *UpdatePolicyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePolicyRequest.Unmarshal(m, b)
}
func (m *UpdatePolicyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePolicyRequest.Marshal(b, m, deterministic)
}
func (m *UpdatePolicyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePolicyRequest.Merge(m, src)
}
func (m *UpdatePolicyRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePolicyRequest.Size(m)
}
func (m *UpdatePolicyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePolicyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePolicyRequest proto.InternalMessageInfo

func (m *UpdatePolicyRequest) GetPolicy() *Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type DeletePolicyRequest struct {
	// The object resource name of the Policy to delete. In the format:
	// `orn:judge-org:judge-service::policy/default_policy`.
	PolicyId             string   `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePolicyRequest) Reset()         { *m = DeletePolicyRequest{} }
func (m *DeletePolicyRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePolicyRequest) ProtoMessage()    {}
func (*DeletePolicyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8c4c1de41aa150c, []int{5}
}

func (m *DeletePolicyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePolicyRequest.Unmarshal(m, b)
}
func (m *DeletePolicyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePolicyRequest.Marshal(b, m, deterministic)
}
func (m *DeletePolicyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePolicyRequest.Merge(m, src)
}
func (m *DeletePolicyRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePolicyRequest.Size(m)
}
func (m *DeletePolicyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePolicyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePolicyRequest proto.InternalMessageInfo

func (m *DeletePolicyRequest) GetPolicyId() string {
	if m != nil {
		return m.PolicyId
	}
	return ""
}

type AuthorizeRequest struct {
	Object               string            `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	Action               string            `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Context              map[string]string `protobuf:"bytes,3,rep,name=context,proto3" json:"context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AuthorizeRequest) Reset()         { *m = AuthorizeRequest{} }
func (m *AuthorizeRequest) String() string { return proto.CompactTextString(m) }
func (*AuthorizeRequest) ProtoMessage()    {}
func (*AuthorizeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8c4c1de41aa150c, []int{6}
}

func (m *AuthorizeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizeRequest.Unmarshal(m, b)
}
func (m *AuthorizeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizeRequest.Marshal(b, m, deterministic)
}
func (m *AuthorizeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizeRequest.Merge(m, src)
}
func (m *AuthorizeRequest) XXX_Size() int {
	return xxx_messageInfo_AuthorizeRequest.Size(m)
}
func (m *AuthorizeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizeRequest proto.InternalMessageInfo

func (m *AuthorizeRequest) GetObject() string {
	if m != nil {
		return m.Object
	}
	return ""
}

func (m *AuthorizeRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *AuthorizeRequest) GetContext() map[string]string {
	if m != nil {
		return m.Context
	}
	return nil
}

type AuthorizeResponse struct {
	Authorized           bool     `protobuf:"varint,1,opt,name=authorized,proto3" json:"authorized,omitempty"`
	Explain              string   `protobuf:"bytes,2,opt,name=explain,proto3" json:"explain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthorizeResponse) Reset()         { *m = AuthorizeResponse{} }
func (m *AuthorizeResponse) String() string { return proto.CompactTextString(m) }
func (*AuthorizeResponse) ProtoMessage()    {}
func (*AuthorizeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8c4c1de41aa150c, []int{7}
}

func (m *AuthorizeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizeResponse.Unmarshal(m, b)
}
func (m *AuthorizeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizeResponse.Marshal(b, m, deterministic)
}
func (m *AuthorizeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizeResponse.Merge(m, src)
}
func (m *AuthorizeResponse) XXX_Size() int {
	return xxx_messageInfo_AuthorizeResponse.Size(m)
}
func (m *AuthorizeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizeResponse proto.InternalMessageInfo

func (m *AuthorizeResponse) GetAuthorized() bool {
	if m != nil {
		return m.Authorized
	}
	return false
}

func (m *AuthorizeResponse) GetExplain() string {
	if m != nil {
		return m.Explain
	}
	return ""
}

func init() {
	proto.RegisterType((*GetPolicyRequest)(nil), "judge.api.v1alpha1.GetPolicyRequest")
	proto.RegisterType((*ListPoliciesRequest)(nil), "judge.api.v1alpha1.ListPoliciesRequest")
	proto.RegisterType((*ListPoliciesResponse)(nil), "judge.api.v1alpha1.ListPoliciesResponse")
	proto.RegisterType((*CreatePolicyRequest)(nil), "judge.api.v1alpha1.CreatePolicyRequest")
	proto.RegisterType((*UpdatePolicyRequest)(nil), "judge.api.v1alpha1.UpdatePolicyRequest")
	proto.RegisterType((*DeletePolicyRequest)(nil), "judge.api.v1alpha1.DeletePolicyRequest")
	proto.RegisterType((*AuthorizeRequest)(nil), "judge.api.v1alpha1.AuthorizeRequest")
	proto.RegisterMapType((map[string]string)(nil), "judge.api.v1alpha1.AuthorizeRequest.ContextEntry")
	proto.RegisterType((*AuthorizeResponse)(nil), "judge.api.v1alpha1.AuthorizeResponse")
}

func init() { proto.RegisterFile("api/judge/v1alpha1/judge.proto", fileDescriptor_b8c4c1de41aa150c) }

var fileDescriptor_b8c4c1de41aa150c = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x1b, 0x9a, 0xcb, 0x10, 0x44, 0xd8, 0x54, 0x95, 0x71, 0x45, 0xa9, 0x2c, 0x2e, 0x79,
	0xb2, 0x49, 0x90, 0x10, 0xea, 0x1b, 0x2d, 0x15, 0x2a, 0x50, 0x51, 0x25, 0x20, 0x21, 0x5e, 0xa2,
	0x4d, 0x3c, 0x75, 0xb6, 0x71, 0xbd, 0xc6, 0x5e, 0x47, 0x49, 0x3e, 0x82, 0x1f, 0xe3, 0xa7, 0x90,
	0xbd, 0xeb, 0x60, 0x92, 0x55, 0x5a, 0xde, 0x3c, 0xb7, 0x33, 0x67, 0xc6, 0x67, 0x07, 0x0e, 0x69,
	0xc4, 0xdc, 0xeb, 0xd4, 0xf3, 0xd1, 0x9d, 0x75, 0x69, 0x10, 0x4d, 0x68, 0x57, 0x9a, 0x4e, 0x14,
	0x73, 0xc1, 0x09, 0x91, 0x06, 0x8d, 0x98, 0x53, 0xc4, 0xad, 0xa7, 0x9a, 0x9a, 0x88, 0x07, 0x6c,
	0xbc, 0x90, 0x45, 0xd6, 0x81, 0xcf, 0xb9, 0x1f, 0xa0, 0x9b, 0x5b, 0xa3, 0xf4, 0xca, 0xc5, 0x9b,
	0x48, 0xa8, 0xa0, 0xed, 0x42, 0xeb, 0x03, 0x8a, 0xcb, 0x3c, 0xbf, 0x8f, 0x3f, 0x53, 0x4c, 0x04,
	0x39, 0x80, 0x86, 0x04, 0x18, 0x32, 0xcf, 0x34, 0x8e, 0x8c, 0x4e, 0xa3, 0x5f, 0x97, 0x8e, 0x73,
	0xcf, 0x0e, 0xa0, 0xfd, 0x99, 0x25, 0xb2, 0x82, 0x61, 0x52, 0xae, 0xa1, 0x3e, 0x0e, 0x13, 0xb6,
	0xc4, 0xbc, 0x66, 0xb7, 0x5f, 0xcf, 0x1c, 0x03, 0xb6, 0x44, 0xf2, 0x04, 0x20, 0x0f, 0x0a, 0x3e,
	0xc5, 0xd0, 0xdc, 0xc9, 0x11, 0xf3, 0xf4, 0xaf, 0x99, 0x83, 0x3c, 0x86, 0x3a, 0x8f, 0x3d, 0x8c,
	0x87, 0xa3, 0x85, 0x59, 0xc9, 0x83, 0xb5, 0xdc, 0x3e, 0x59, 0xd8, 0x33, 0xd8, 0xfb, 0xb7, 0x5b,
	0x12, 0xf1, 0x30, 0x41, 0xf2, 0x06, 0x24, 0x23, 0x86, 0x89, 0x69, 0x1c, 0x55, 0x3a, 0xf7, 0x7b,
	0x96, 0xb3, 0xb9, 0x1b, 0x47, 0xcd, 0xb5, 0xca, 0x25, 0x2f, 0xe0, 0x61, 0x88, 0x73, 0x31, 0xdc,
	0xa0, 0xf3, 0x20, 0x73, 0x5f, 0x16, 0x94, 0xec, 0x2b, 0x68, 0x9f, 0xc6, 0x48, 0x05, 0xde, 0x7d,
	0x33, 0xa4, 0x07, 0x55, 0xf9, 0x9d, 0x43, 0x6e, 0x67, 0xa4, 0x32, 0xed, 0x73, 0x68, 0x7f, 0x8b,
	0xbc, 0x8d, 0x3e, 0x7f, 0xa1, 0x8c, 0x3b, 0x43, 0xf5, 0xa0, 0xfd, 0x1e, 0x03, 0xfc, 0x1f, 0xca,
	0xf6, 0x6f, 0x03, 0x5a, 0xef, 0x52, 0x31, 0xe1, 0x31, 0x5b, 0x62, 0x51, 0xb1, 0x0f, 0x55, 0x3e,
	0xba, 0xc6, 0xb1, 0x50, 0xe9, 0xca, 0xca, 0xfc, 0x74, 0x2c, 0x18, 0x2f, 0x56, 0xa6, 0x2c, 0xf2,
	0x09, 0x6a, 0x63, 0x1e, 0x0a, 0x9c, 0x0b, 0xb3, 0x92, 0xff, 0x8a, 0xae, 0x8e, 0xed, 0x7a, 0x1b,
	0xe7, 0x54, 0xd6, 0x9c, 0x85, 0x22, 0x5e, 0xf4, 0x0b, 0x04, 0xeb, 0x18, 0x9a, 0xe5, 0x00, 0x69,
	0x41, 0x65, 0x8a, 0x0b, 0xc5, 0x24, 0xfb, 0x24, 0x7b, 0xb0, 0x3b, 0xa3, 0x41, 0x8a, 0x8a, 0x85,
	0x34, 0x8e, 0x77, 0xde, 0x1a, 0xf6, 0x05, 0x3c, 0x2a, 0x75, 0x51, 0x4a, 0x39, 0x04, 0xa0, 0x85,
	0x53, 0x2e, 0xa0, 0xde, 0x2f, 0x79, 0x88, 0x09, 0x35, 0x9c, 0x47, 0x01, 0x65, 0xc5, 0x58, 0x85,
	0xd9, 0xfb, 0x75, 0x0f, 0x76, 0x3f, 0x66, 0x83, 0x90, 0x0b, 0x68, 0xac, 0x1e, 0x09, 0x79, 0xa6,
	0x9b, 0x6e, 0xfd, 0x0d, 0x59, 0x5b, 0xfe, 0x18, 0xa1, 0xd0, 0x2c, 0x8b, 0x9a, 0xbc, 0xd4, 0xe5,
	0x6a, 0x1e, 0x99, 0xd5, 0xb9, 0x3d, 0x51, 0x4d, 0x3d, 0x80, 0x66, 0x59, 0xbf, 0xfa, 0x16, 0x1a,
	0x85, 0x6f, 0xe5, 0x3d, 0x80, 0x66, 0x59, 0xac, 0x7a, 0x50, 0x8d, 0x9c, 0xb7, 0x82, 0x7e, 0x81,
	0x66, 0x59, 0xb6, 0x7a, 0x50, 0x8d, 0xb0, 0xad, 0x7d, 0x47, 0xde, 0x35, 0xa7, 0xb8, 0x6b, 0xce,
	0x59, 0x76, 0xd7, 0xc8, 0x77, 0x68, 0xac, 0x54, 0xa0, 0xff, 0x59, 0xeb, 0x52, 0xb4, 0x9e, 0xdf,
	0x92, 0x25, 0x97, 0x7a, 0xd2, 0xfb, 0xf1, 0xca, 0x67, 0x62, 0x92, 0x8e, 0x9c, 0x31, 0xbf, 0x71,
	0x7d, 0xa4, 0x71, 0xc8, 0x3d, 0x54, 0xb7, 0x37, 0x9a, 0xfa, 0x2e, 0x8d, 0x58, 0x82, 0xf1, 0x0c,
	0xe3, 0xd5, 0x25, 0x1e, 0x55, 0x73, 0x76, 0xaf, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xa6,
	0x7a, 0x03, 0xda, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// JudgeClient is the client API for Judge service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JudgeClient interface {
	// Gets a single policy.
	GetPolicy(ctx context.Context, in *GetPolicyRequest, opts ...grpc.CallOption) (*Policy, error)
	// Lists policies.
	ListPolicies(ctx context.Context, in *ListPoliciesRequest, opts ...grpc.CallOption) (*ListPoliciesResponse, error)
	// Creates a new policy.
	CreatePolicy(ctx context.Context, in *CreatePolicyRequest, opts ...grpc.CallOption) (*Policy, error)
	// Updates or inserts a policy.
	UpdatePolicy(ctx context.Context, in *UpdatePolicyRequest, opts ...grpc.CallOption) (*Policy, error)
	// Deletes a policy.
	DeletePolicy(ctx context.Context, in *DeletePolicyRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error)
}

type judgeClient struct {
	cc *grpc.ClientConn
}

func NewJudgeClient(cc *grpc.ClientConn) JudgeClient {
	return &judgeClient{cc}
}

func (c *judgeClient) GetPolicy(ctx context.Context, in *GetPolicyRequest, opts ...grpc.CallOption) (*Policy, error) {
	out := new(Policy)
	err := c.cc.Invoke(ctx, "/judge.api.v1alpha1.Judge/GetPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judgeClient) ListPolicies(ctx context.Context, in *ListPoliciesRequest, opts ...grpc.CallOption) (*ListPoliciesResponse, error) {
	out := new(ListPoliciesResponse)
	err := c.cc.Invoke(ctx, "/judge.api.v1alpha1.Judge/ListPolicies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judgeClient) CreatePolicy(ctx context.Context, in *CreatePolicyRequest, opts ...grpc.CallOption) (*Policy, error) {
	out := new(Policy)
	err := c.cc.Invoke(ctx, "/judge.api.v1alpha1.Judge/CreatePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judgeClient) UpdatePolicy(ctx context.Context, in *UpdatePolicyRequest, opts ...grpc.CallOption) (*Policy, error) {
	out := new(Policy)
	err := c.cc.Invoke(ctx, "/judge.api.v1alpha1.Judge/UpdatePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judgeClient) DeletePolicy(ctx context.Context, in *DeletePolicyRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/judge.api.v1alpha1.Judge/DeletePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judgeClient) Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error) {
	out := new(AuthorizeResponse)
	err := c.cc.Invoke(ctx, "/judge.api.v1alpha1.Judge/Authorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JudgeServer is the server API for Judge service.
type JudgeServer interface {
	// Gets a single policy.
	GetPolicy(context.Context, *GetPolicyRequest) (*Policy, error)
	// Lists policies.
	ListPolicies(context.Context, *ListPoliciesRequest) (*ListPoliciesResponse, error)
	// Creates a new policy.
	CreatePolicy(context.Context, *CreatePolicyRequest) (*Policy, error)
	// Updates or inserts a policy.
	UpdatePolicy(context.Context, *UpdatePolicyRequest) (*Policy, error)
	// Deletes a policy.
	DeletePolicy(context.Context, *DeletePolicyRequest) (*empty.Empty, error)
	Authorize(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error)
}

func RegisterJudgeServer(s *grpc.Server, srv JudgeServer) {
	s.RegisterService(&_Judge_serviceDesc, srv)
}

func _Judge_GetPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgeServer).GetPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/judge.api.v1alpha1.Judge/GetPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgeServer).GetPolicy(ctx, req.(*GetPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Judge_ListPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgeServer).ListPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/judge.api.v1alpha1.Judge/ListPolicies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgeServer).ListPolicies(ctx, req.(*ListPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Judge_CreatePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgeServer).CreatePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/judge.api.v1alpha1.Judge/CreatePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgeServer).CreatePolicy(ctx, req.(*CreatePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Judge_UpdatePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgeServer).UpdatePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/judge.api.v1alpha1.Judge/UpdatePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgeServer).UpdatePolicy(ctx, req.(*UpdatePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Judge_DeletePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgeServer).DeletePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/judge.api.v1alpha1.Judge/DeletePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgeServer).DeletePolicy(ctx, req.(*DeletePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Judge_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgeServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/judge.api.v1alpha1.Judge/Authorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgeServer).Authorize(ctx, req.(*AuthorizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Judge_serviceDesc = grpc.ServiceDesc{
	ServiceName: "judge.api.v1alpha1.Judge",
	HandlerType: (*JudgeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPolicy",
			Handler:    _Judge_GetPolicy_Handler,
		},
		{
			MethodName: "ListPolicies",
			Handler:    _Judge_ListPolicies_Handler,
		},
		{
			MethodName: "CreatePolicy",
			Handler:    _Judge_CreatePolicy_Handler,
		},
		{
			MethodName: "UpdatePolicy",
			Handler:    _Judge_UpdatePolicy_Handler,
		},
		{
			MethodName: "DeletePolicy",
			Handler:    _Judge_DeletePolicy_Handler,
		},
		{
			MethodName: "Authorize",
			Handler:    _Judge_Authorize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/judge/v1alpha1/judge.proto",
}
