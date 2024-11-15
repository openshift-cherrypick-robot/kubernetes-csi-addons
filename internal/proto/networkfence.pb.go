// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.19.6
// source: networkfence.proto

package proto

import (
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

// NetworkFenceRequest holds the required information to fence/unfence
// the cluster network.
type NetworkFenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Plugin specific parameters passed in as opaque key-value pairs.
	Parameters map[string]string `protobuf:"bytes,1,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Secrets required by the driver to complete the request.
	SecretName      string `protobuf:"bytes,2,opt,name=secret_name,json=secretName,proto3" json:"secret_name,omitempty"`
	SecretNamespace string `protobuf:"bytes,3,opt,name=secret_namespace,json=secretNamespace,proto3" json:"secret_namespace,omitempty"`
	// list of CIDR blocks on which the fencing/unfencing operation is expected
	// to be performed.
	Cidrs []string `protobuf:"bytes,4,rep,name=cidrs,proto3" json:"cidrs,omitempty"`
}

func (x *NetworkFenceRequest) Reset() {
	*x = NetworkFenceRequest{}
	mi := &file_networkfence_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkFenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkFenceRequest) ProtoMessage() {}

func (x *NetworkFenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_networkfence_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkFenceRequest.ProtoReflect.Descriptor instead.
func (*NetworkFenceRequest) Descriptor() ([]byte, []int) {
	return file_networkfence_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkFenceRequest) GetParameters() map[string]string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *NetworkFenceRequest) GetSecretName() string {
	if x != nil {
		return x.SecretName
	}
	return ""
}

func (x *NetworkFenceRequest) GetSecretNamespace() string {
	if x != nil {
		return x.SecretNamespace
	}
	return ""
}

func (x *NetworkFenceRequest) GetCidrs() []string {
	if x != nil {
		return x.Cidrs
	}
	return nil
}

// NetworkFenceResponse is returned by the CSI-driver as a result of
// the FenceRequest call.
type NetworkFenceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NetworkFenceResponse) Reset() {
	*x = NetworkFenceResponse{}
	mi := &file_networkfence_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkFenceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkFenceResponse) ProtoMessage() {}

func (x *NetworkFenceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_networkfence_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkFenceResponse.ProtoReflect.Descriptor instead.
func (*NetworkFenceResponse) Descriptor() ([]byte, []int) {
	return file_networkfence_proto_rawDescGZIP(), []int{1}
}

// FenceClientsRequest contains the necessary information to identify
// the clients that need fencing.
type FenceClientsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Plugin-specific parameters passed in as opaque key-value pairs.
	Parameters map[string]string `protobuf:"bytes,1,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Secrets required by the driver to complete the request.
	SecretName      string `protobuf:"bytes,2,opt,name=secret_name,json=secretName,proto3" json:"secret_name,omitempty"`
	SecretNamespace string `protobuf:"bytes,3,opt,name=secret_namespace,json=secretNamespace,proto3" json:"secret_namespace,omitempty"`
}

func (x *FenceClientsRequest) Reset() {
	*x = FenceClientsRequest{}
	mi := &file_networkfence_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FenceClientsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FenceClientsRequest) ProtoMessage() {}

func (x *FenceClientsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_networkfence_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FenceClientsRequest.ProtoReflect.Descriptor instead.
func (*FenceClientsRequest) Descriptor() ([]byte, []int) {
	return file_networkfence_proto_rawDescGZIP(), []int{2}
}

func (x *FenceClientsRequest) GetParameters() map[string]string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *FenceClientsRequest) GetSecretName() string {
	if x != nil {
		return x.SecretName
	}
	return ""
}

func (x *FenceClientsRequest) GetSecretNamespace() string {
	if x != nil {
		return x.SecretNamespace
	}
	return ""
}

// FenceClientsResponse holds the information about clients that require
// fencing.
type FenceClientsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of clients that need to be fenced.
	Clients []*ClientDetails `protobuf:"bytes,1,rep,name=clients,proto3" json:"clients,omitempty"`
}

func (x *FenceClientsResponse) Reset() {
	*x = FenceClientsResponse{}
	mi := &file_networkfence_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FenceClientsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FenceClientsResponse) ProtoMessage() {}

func (x *FenceClientsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_networkfence_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FenceClientsResponse.ProtoReflect.Descriptor instead.
func (*FenceClientsResponse) Descriptor() ([]byte, []int) {
	return file_networkfence_proto_rawDescGZIP(), []int{3}
}

func (x *FenceClientsResponse) GetClients() []*ClientDetails {
	if x != nil {
		return x.Clients
	}
	return nil
}

// ClientDetails holds the information about the client that requires fencing.
type ClientDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id represents the unique identifier of the client.
	// Required field.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// list of cidrs that represent the client's local cidrs.
	// Required field.
	Cidrs []string `protobuf:"bytes,2,rep,name=cidrs,proto3" json:"cidrs,omitempty"`
}

func (x *ClientDetails) Reset() {
	*x = ClientDetails{}
	mi := &file_networkfence_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientDetails) ProtoMessage() {}

func (x *ClientDetails) ProtoReflect() protoreflect.Message {
	mi := &file_networkfence_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientDetails.ProtoReflect.Descriptor instead.
func (*ClientDetails) Descriptor() ([]byte, []int) {
	return file_networkfence_proto_rawDescGZIP(), []int{4}
}

func (x *ClientDetails) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ClientDetails) GetCidrs() []string {
	if x != nil {
		return x.Cidrs
	}
	return nil
}

var File_networkfence_proto protoreflect.FileDescriptor

var file_networkfence_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x02, 0x0a, 0x13,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x29, 0x0a, 0x10, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x69, 0x64, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x63, 0x69, 0x64, 0x72,
	0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x16, 0x0a, 0x14, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x46, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xec, 0x01, 0x0a, 0x13, 0x46, 0x65, 0x6e,
	0x63, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x4a, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x65, 0x6e,
	0x63, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a,
	0x10, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x46, 0x0a, 0x14, 0x46, 0x65, 0x6e, 0x63, 0x65,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2e, 0x0a, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0x35, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x69, 0x64, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x63, 0x69, 0x64, 0x72, 0x73, 0x32, 0x82, 0x02, 0x0a, 0x0c, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x13, 0x46, 0x65, 0x6e, 0x63, 0x65,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x46, 0x65,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x15, 0x55, 0x6e, 0x46,
	0x65, 0x6e, 0x63, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x46, 0x65,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3c, 0x5a, 0x3a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x73, 0x69, 0x2d, 0x61, 0x64,
	0x64, 0x6f, 0x6e, 0x73, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2d,
	0x63, 0x73, 0x69, 0x2d, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_networkfence_proto_rawDescOnce sync.Once
	file_networkfence_proto_rawDescData = file_networkfence_proto_rawDesc
)

func file_networkfence_proto_rawDescGZIP() []byte {
	file_networkfence_proto_rawDescOnce.Do(func() {
		file_networkfence_proto_rawDescData = protoimpl.X.CompressGZIP(file_networkfence_proto_rawDescData)
	})
	return file_networkfence_proto_rawDescData
}

var file_networkfence_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_networkfence_proto_goTypes = []any{
	(*NetworkFenceRequest)(nil),  // 0: proto.NetworkFenceRequest
	(*NetworkFenceResponse)(nil), // 1: proto.NetworkFenceResponse
	(*FenceClientsRequest)(nil),  // 2: proto.FenceClientsRequest
	(*FenceClientsResponse)(nil), // 3: proto.FenceClientsResponse
	(*ClientDetails)(nil),        // 4: proto.ClientDetails
	nil,                          // 5: proto.NetworkFenceRequest.ParametersEntry
	nil,                          // 6: proto.FenceClientsRequest.ParametersEntry
}
var file_networkfence_proto_depIdxs = []int32{
	5, // 0: proto.NetworkFenceRequest.parameters:type_name -> proto.NetworkFenceRequest.ParametersEntry
	6, // 1: proto.FenceClientsRequest.parameters:type_name -> proto.FenceClientsRequest.ParametersEntry
	4, // 2: proto.FenceClientsResponse.clients:type_name -> proto.ClientDetails
	0, // 3: proto.NetworkFence.FenceClusterNetwork:input_type -> proto.NetworkFenceRequest
	0, // 4: proto.NetworkFence.UnFenceClusterNetwork:input_type -> proto.NetworkFenceRequest
	2, // 5: proto.NetworkFence.GetFenceClients:input_type -> proto.FenceClientsRequest
	1, // 6: proto.NetworkFence.FenceClusterNetwork:output_type -> proto.NetworkFenceResponse
	1, // 7: proto.NetworkFence.UnFenceClusterNetwork:output_type -> proto.NetworkFenceResponse
	3, // 8: proto.NetworkFence.GetFenceClients:output_type -> proto.FenceClientsResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_networkfence_proto_init() }
func file_networkfence_proto_init() {
	if File_networkfence_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_networkfence_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_networkfence_proto_goTypes,
		DependencyIndexes: file_networkfence_proto_depIdxs,
		MessageInfos:      file_networkfence_proto_msgTypes,
	}.Build()
	File_networkfence_proto = out.File
	file_networkfence_proto_rawDesc = nil
	file_networkfence_proto_goTypes = nil
	file_networkfence_proto_depIdxs = nil
}
