// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: interservice/infra_proxy/request/roles.proto

package request

import (
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Roles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Roles search query.
	SearchQuery *SearchQuery `protobuf:"bytes,3,opt,name=search_query,json=searchQuery,proto3" json:"search_query,omitempty" toml:"search_query,omitempty" mapstructure:"search_query,omitempty"`
}

func (x *Roles) Reset() {
	*x = Roles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_roles_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Roles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Roles) ProtoMessage() {}

func (x *Roles) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_roles_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Roles.ProtoReflect.Descriptor instead.
func (*Roles) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_roles_proto_rawDescGZIP(), []int{0}
}

func (x *Roles) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *Roles) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *Roles) GetSearchQuery() *SearchQuery {
	if x != nil {
		return x.SearchQuery
	}
	return nil
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Role name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_roles_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_roles_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_roles_proto_rawDescGZIP(), []int{1}
}

func (x *Role) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *Role) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Role name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Role description.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty" toml:"description,omitempty" mapstructure:"description,omitempty"`
	// Role default attributes JSON.
	DefaultAttributes *_struct.Struct `protobuf:"bytes,5,opt,name=default_attributes,json=defaultAttributes,proto3" json:"default_attributes,omitempty" toml:"default_attributes,omitempty" mapstructure:"default_attributes,omitempty"`
	// Role override attributes JSON.
	OverrideAttributes *_struct.Struct `protobuf:"bytes,6,opt,name=override_attributes,json=overrideAttributes,proto3" json:"override_attributes,omitempty" toml:"override_attributes,omitempty" mapstructure:"override_attributes,omitempty"`
	// Role run list.
	RunList []string `protobuf:"bytes,7,rep,name=run_list,json=runList,proto3" json:"run_list,omitempty" toml:"run_list,omitempty" mapstructure:"run_list,omitempty"`
	// Environment based run list.
	EnvRunLists []*EnvRunList `protobuf:"bytes,8,rep,name=env_run_lists,json=envRunLists,proto3" json:"env_run_lists,omitempty" toml:"env_run_lists,omitempty" mapstructure:"env_run_lists,omitempty"`
}

func (x *CreateRole) Reset() {
	*x = CreateRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_roles_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRole) ProtoMessage() {}

func (x *CreateRole) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_roles_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRole.ProtoReflect.Descriptor instead.
func (*CreateRole) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_roles_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRole) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *CreateRole) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *CreateRole) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRole) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateRole) GetDefaultAttributes() *_struct.Struct {
	if x != nil {
		return x.DefaultAttributes
	}
	return nil
}

func (x *CreateRole) GetOverrideAttributes() *_struct.Struct {
	if x != nil {
		return x.OverrideAttributes
	}
	return nil
}

func (x *CreateRole) GetRunList() []string {
	if x != nil {
		return x.RunList
	}
	return nil
}

func (x *CreateRole) GetEnvRunLists() []*EnvRunList {
	if x != nil {
		return x.EnvRunLists
	}
	return nil
}

type UpdateRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Role name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Role description.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty" toml:"description,omitempty" mapstructure:"description,omitempty"`
	// Role default attributes JSON.
	DefaultAttributes *_struct.Struct `protobuf:"bytes,5,opt,name=default_attributes,json=defaultAttributes,proto3" json:"default_attributes,omitempty" toml:"default_attributes,omitempty" mapstructure:"default_attributes,omitempty"`
	// Role override attributes JSON.
	OverrideAttributes *_struct.Struct `protobuf:"bytes,6,opt,name=override_attributes,json=overrideAttributes,proto3" json:"override_attributes,omitempty" toml:"override_attributes,omitempty" mapstructure:"override_attributes,omitempty"`
	// Role run list.
	RunList []string `protobuf:"bytes,7,rep,name=run_list,json=runList,proto3" json:"run_list,omitempty" toml:"run_list,omitempty" mapstructure:"run_list,omitempty"`
	// Environment based run list.
	EnvRunLists []*EnvRunList `protobuf:"bytes,8,rep,name=env_run_lists,json=envRunLists,proto3" json:"env_run_lists,omitempty" toml:"env_run_lists,omitempty" mapstructure:"env_run_lists,omitempty"`
}

func (x *UpdateRole) Reset() {
	*x = UpdateRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_roles_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRole) ProtoMessage() {}

func (x *UpdateRole) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_roles_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRole.ProtoReflect.Descriptor instead.
func (*UpdateRole) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_roles_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRole) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *UpdateRole) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *UpdateRole) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateRole) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateRole) GetDefaultAttributes() *_struct.Struct {
	if x != nil {
		return x.DefaultAttributes
	}
	return nil
}

func (x *UpdateRole) GetOverrideAttributes() *_struct.Struct {
	if x != nil {
		return x.OverrideAttributes
	}
	return nil
}

func (x *UpdateRole) GetRunList() []string {
	if x != nil {
		return x.RunList
	}
	return nil
}

func (x *UpdateRole) GetEnvRunLists() []*EnvRunList {
	if x != nil {
		return x.EnvRunLists
	}
	return nil
}

type EnvRunList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Environment name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Role run list.
	RunList []string `protobuf:"bytes,2,rep,name=run_list,json=runList,proto3" json:"run_list,omitempty" toml:"run_list,omitempty" mapstructure:"run_list,omitempty"`
}

func (x *EnvRunList) Reset() {
	*x = EnvRunList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_roles_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvRunList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvRunList) ProtoMessage() {}

func (x *EnvRunList) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_roles_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvRunList.ProtoReflect.Descriptor instead.
func (*EnvRunList) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_roles_proto_rawDescGZIP(), []int{4}
}

func (x *EnvRunList) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EnvRunList) GetRunList() []string {
	if x != nil {
		return x.RunList
	}
	return nil
}

var File_interservice_infra_proxy_request_roles_proto protoreflect.FileDescriptor

var file_interservice_infra_proxy_request_roles_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28,
	0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x05, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12,
	0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x58, 0x0a, 0x0c, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x68, 0x65, 0x66,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x22, 0x4e, 0x0a,
	0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xfd, 0x02,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x15, 0x0a, 0x06,
	0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72,
	0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x12, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x11, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x48,
	0x0a, 0x13, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x12, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x75, 0x6e, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x58, 0x0a, 0x0d, 0x65, 0x6e, 0x76, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x68, 0x65,
	0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x6e, 0x76, 0x52, 0x75, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x0b, 0x65, 0x6e, 0x76, 0x52, 0x75, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x22, 0xfd, 0x02,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x15, 0x0a, 0x06,
	0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72,
	0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x12, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x11, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x48,
	0x0a, 0x13, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x12, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x75, 0x6e, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x58, 0x0a, 0x0d, 0x65, 0x6e, 0x76, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x68, 0x65,
	0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x6e, 0x76, 0x52, 0x75, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x0b, 0x65, 0x6e, 0x76, 0x52, 0x75, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x22, 0x3b, 0x0a,
	0x0a, 0x45, 0x6e, 0x76, 0x52, 0x75, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x72, 0x75, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_interservice_infra_proxy_request_roles_proto_rawDescOnce sync.Once
	file_interservice_infra_proxy_request_roles_proto_rawDescData = file_interservice_infra_proxy_request_roles_proto_rawDesc
)

func file_interservice_infra_proxy_request_roles_proto_rawDescGZIP() []byte {
	file_interservice_infra_proxy_request_roles_proto_rawDescOnce.Do(func() {
		file_interservice_infra_proxy_request_roles_proto_rawDescData = protoimpl.X.CompressGZIP(file_interservice_infra_proxy_request_roles_proto_rawDescData)
	})
	return file_interservice_infra_proxy_request_roles_proto_rawDescData
}

var file_interservice_infra_proxy_request_roles_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_interservice_infra_proxy_request_roles_proto_goTypes = []interface{}{
	(*Roles)(nil),          // 0: chef.automate.domain.infra_proxy.request.Roles
	(*Role)(nil),           // 1: chef.automate.domain.infra_proxy.request.Role
	(*CreateRole)(nil),     // 2: chef.automate.domain.infra_proxy.request.CreateRole
	(*UpdateRole)(nil),     // 3: chef.automate.domain.infra_proxy.request.UpdateRole
	(*EnvRunList)(nil),     // 4: chef.automate.domain.infra_proxy.request.EnvRunList
	(*SearchQuery)(nil),    // 5: chef.automate.domain.infra_proxy.request.SearchQuery
	(*_struct.Struct)(nil), // 6: google.protobuf.Struct
}
var file_interservice_infra_proxy_request_roles_proto_depIdxs = []int32{
	5, // 0: chef.automate.domain.infra_proxy.request.Roles.search_query:type_name -> chef.automate.domain.infra_proxy.request.SearchQuery
	6, // 1: chef.automate.domain.infra_proxy.request.CreateRole.default_attributes:type_name -> google.protobuf.Struct
	6, // 2: chef.automate.domain.infra_proxy.request.CreateRole.override_attributes:type_name -> google.protobuf.Struct
	4, // 3: chef.automate.domain.infra_proxy.request.CreateRole.env_run_lists:type_name -> chef.automate.domain.infra_proxy.request.EnvRunList
	6, // 4: chef.automate.domain.infra_proxy.request.UpdateRole.default_attributes:type_name -> google.protobuf.Struct
	6, // 5: chef.automate.domain.infra_proxy.request.UpdateRole.override_attributes:type_name -> google.protobuf.Struct
	4, // 6: chef.automate.domain.infra_proxy.request.UpdateRole.env_run_lists:type_name -> chef.automate.domain.infra_proxy.request.EnvRunList
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_interservice_infra_proxy_request_roles_proto_init() }
func file_interservice_infra_proxy_request_roles_proto_init() {
	if File_interservice_infra_proxy_request_roles_proto != nil {
		return
	}
	file_interservice_infra_proxy_request_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_interservice_infra_proxy_request_roles_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Roles); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_interservice_infra_proxy_request_roles_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_interservice_infra_proxy_request_roles_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRole); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_interservice_infra_proxy_request_roles_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRole); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_interservice_infra_proxy_request_roles_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvRunList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_interservice_infra_proxy_request_roles_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_interservice_infra_proxy_request_roles_proto_goTypes,
		DependencyIndexes: file_interservice_infra_proxy_request_roles_proto_depIdxs,
		MessageInfos:      file_interservice_infra_proxy_request_roles_proto_msgTypes,
	}.Build()
	File_interservice_infra_proxy_request_roles_proto = out.File
	file_interservice_infra_proxy_request_roles_proto_rawDesc = nil
	file_interservice_infra_proxy_request_roles_proto_goTypes = nil
	file_interservice_infra_proxy_request_roles_proto_depIdxs = nil
}
