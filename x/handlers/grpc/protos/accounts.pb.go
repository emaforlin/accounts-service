// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: protos/accounts.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VerifyFoodPlaceAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid uint32 `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
}

func (x *VerifyFoodPlaceAccountRequest) Reset() {
	*x = VerifyFoodPlaceAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_accounts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyFoodPlaceAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyFoodPlaceAccountRequest) ProtoMessage() {}

func (x *VerifyFoodPlaceAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_accounts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyFoodPlaceAccountRequest.ProtoReflect.Descriptor instead.
func (*VerifyFoodPlaceAccountRequest) Descriptor() ([]byte, []int) {
	return file_protos_accounts_proto_rawDescGZIP(), []int{0}
}

func (x *VerifyFoodPlaceAccountRequest) GetUserid() uint32 {
	if x != nil {
		return x.Userid
	}
	return 0
}

type VerifyFoodPlaceAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VerifyFoodPlaceAccountResponse) Reset() {
	*x = VerifyFoodPlaceAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_accounts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyFoodPlaceAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyFoodPlaceAccountResponse) ProtoMessage() {}

func (x *VerifyFoodPlaceAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_accounts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyFoodPlaceAccountResponse.ProtoReflect.Descriptor instead.
func (*VerifyFoodPlaceAccountResponse) Descriptor() ([]byte, []int) {
	return file_protos_accounts_proto_rawDescGZIP(), []int{1}
}

type AddFoodPlaceAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username     string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	PhoneNumber  string   `protobuf:"bytes,2,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Email        string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password     string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	BusinessName string   `protobuf:"bytes,5,opt,name=business_name,json=businessName,proto3" json:"business_name,omitempty"`
	Location     string   `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	Tags         []string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *AddFoodPlaceAccountRequest) Reset() {
	*x = AddFoodPlaceAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_accounts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFoodPlaceAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFoodPlaceAccountRequest) ProtoMessage() {}

func (x *AddFoodPlaceAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_accounts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFoodPlaceAccountRequest.ProtoReflect.Descriptor instead.
func (*AddFoodPlaceAccountRequest) Descriptor() ([]byte, []int) {
	return file_protos_accounts_proto_rawDescGZIP(), []int{2}
}

func (x *AddFoodPlaceAccountRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AddFoodPlaceAccountRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *AddFoodPlaceAccountRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AddFoodPlaceAccountRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AddFoodPlaceAccountRequest) GetBusinessName() string {
	if x != nil {
		return x.BusinessName
	}
	return ""
}

func (x *AddFoodPlaceAccountRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *AddFoodPlaceAccountRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type AddPersonAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username    string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	FirstName   string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName    string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	PhoneNumber string `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Email       string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Password    string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AddPersonAccountRequest) Reset() {
	*x = AddPersonAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_accounts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPersonAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPersonAccountRequest) ProtoMessage() {}

func (x *AddPersonAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_accounts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPersonAccountRequest.ProtoReflect.Descriptor instead.
func (*AddPersonAccountRequest) Descriptor() ([]byte, []int) {
	return file_protos_accounts_proto_rawDescGZIP(), []int{3}
}

func (x *AddPersonAccountRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AddPersonAccountRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *AddPersonAccountRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *AddPersonAccountRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *AddPersonAccountRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AddPersonAccountRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GetAccountDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Identifiers:
	//
	//	*GetAccountDetailsRequest_Username
	//	*GetAccountDetailsRequest_Email
	//	*GetAccountDetailsRequest_PhoneNumber
	Identifiers isGetAccountDetailsRequest_Identifiers `protobuf_oneof:"identifiers"`
}

func (x *GetAccountDetailsRequest) Reset() {
	*x = GetAccountDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_accounts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountDetailsRequest) ProtoMessage() {}

func (x *GetAccountDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_accounts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountDetailsRequest.ProtoReflect.Descriptor instead.
func (*GetAccountDetailsRequest) Descriptor() ([]byte, []int) {
	return file_protos_accounts_proto_rawDescGZIP(), []int{4}
}

func (m *GetAccountDetailsRequest) GetIdentifiers() isGetAccountDetailsRequest_Identifiers {
	if m != nil {
		return m.Identifiers
	}
	return nil
}

func (x *GetAccountDetailsRequest) GetUsername() string {
	if x, ok := x.GetIdentifiers().(*GetAccountDetailsRequest_Username); ok {
		return x.Username
	}
	return ""
}

func (x *GetAccountDetailsRequest) GetEmail() string {
	if x, ok := x.GetIdentifiers().(*GetAccountDetailsRequest_Email); ok {
		return x.Email
	}
	return ""
}

func (x *GetAccountDetailsRequest) GetPhoneNumber() string {
	if x, ok := x.GetIdentifiers().(*GetAccountDetailsRequest_PhoneNumber); ok {
		return x.PhoneNumber
	}
	return ""
}

type isGetAccountDetailsRequest_Identifiers interface {
	isGetAccountDetailsRequest_Identifiers()
}

type GetAccountDetailsRequest_Username struct {
	Username string `protobuf:"bytes,1,opt,name=username,proto3,oneof"`
}

type GetAccountDetailsRequest_Email struct {
	Email string `protobuf:"bytes,2,opt,name=email,proto3,oneof"`
}

type GetAccountDetailsRequest_PhoneNumber struct {
	PhoneNumber string `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3,oneof"`
}

func (*GetAccountDetailsRequest_Username) isGetAccountDetailsRequest_Identifiers() {}

func (*GetAccountDetailsRequest_Email) isGetAccountDetailsRequest_Identifiers() {}

func (*GetAccountDetailsRequest_PhoneNumber) isGetAccountDetailsRequest_Identifiers() {}

type GetAccountDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username    string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email       string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber string                 `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Role        string                 `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt   *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *GetAccountDetailsResponse) Reset() {
	*x = GetAccountDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_accounts_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountDetailsResponse) ProtoMessage() {}

func (x *GetAccountDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_accounts_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountDetailsResponse.ProtoReflect.Descriptor instead.
func (*GetAccountDetailsResponse) Descriptor() ([]byte, []int) {
	return file_protos_accounts_proto_rawDescGZIP(), []int{5}
}

func (x *GetAccountDetailsResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetAccountDetailsResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetAccountDetailsResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetAccountDetailsResponse) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *GetAccountDetailsResponse) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *GetAccountDetailsResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *GetAccountDetailsResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *GetAccountDetailsResponse) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

var File_protos_accounts_proto protoreflect.FileDescriptor

var file_protos_accounts_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x1d, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x46, 0x6f, 0x6f, 0x64, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x1e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x46, 0x6f,
	0x6f, 0x64, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xe2, 0x01, 0x0a, 0x1a, 0x41, 0x64, 0x64, 0x46, 0x6f,
	0x6f, 0x64, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0xc6, 0x01, 0x0a, 0x17,
	0x41, 0x64, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x84, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x0d, 0x0a, 0x0b,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x22, 0xc5, 0x02, 0x0a, 0x19,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x32, 0xc2, 0x03, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x12, 0x83, 0x01, 0x0a, 0x16, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x46, 0x6f, 0x6f, 0x64, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x33, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x46, 0x6f, 0x6f, 0x64, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x34, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x46, 0x6f,
	0x6f, 0x64, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6f,
	0x64, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6f, 0x64, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x74, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x2e, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a,
	0x10, 0x41, 0x64, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2d, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0xa8, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x42, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x56, 0x41, 0xaa, 0x02, 0x14, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0xca, 0x02, 0x14, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0xe2, 0x02, 0x20, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x16, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_accounts_proto_rawDescOnce sync.Once
	file_protos_accounts_proto_rawDescData = file_protos_accounts_proto_rawDesc
)

func file_protos_accounts_proto_rawDescGZIP() []byte {
	file_protos_accounts_proto_rawDescOnce.Do(func() {
		file_protos_accounts_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_accounts_proto_rawDescData)
	})
	return file_protos_accounts_proto_rawDescData
}

var file_protos_accounts_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_accounts_proto_goTypes = []any{
	(*VerifyFoodPlaceAccountRequest)(nil),  // 0: accounts.v1.accounts.VerifyFoodPlaceAccountRequest
	(*VerifyFoodPlaceAccountResponse)(nil), // 1: accounts.v1.accounts.VerifyFoodPlaceAccountResponse
	(*AddFoodPlaceAccountRequest)(nil),     // 2: accounts.v1.accounts.AddFoodPlaceAccountRequest
	(*AddPersonAccountRequest)(nil),        // 3: accounts.v1.accounts.AddPersonAccountRequest
	(*GetAccountDetailsRequest)(nil),       // 4: accounts.v1.accounts.GetAccountDetailsRequest
	(*GetAccountDetailsResponse)(nil),      // 5: accounts.v1.accounts.GetAccountDetailsResponse
	(*timestamppb.Timestamp)(nil),          // 6: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                  // 7: google.protobuf.Empty
}
var file_protos_accounts_proto_depIdxs = []int32{
	6, // 0: accounts.v1.accounts.GetAccountDetailsResponse.created_at:type_name -> google.protobuf.Timestamp
	6, // 1: accounts.v1.accounts.GetAccountDetailsResponse.updated_at:type_name -> google.protobuf.Timestamp
	6, // 2: accounts.v1.accounts.GetAccountDetailsResponse.deleted_at:type_name -> google.protobuf.Timestamp
	0, // 3: accounts.v1.accounts.Accounts.VerifyFoodPlaceAccount:input_type -> accounts.v1.accounts.VerifyFoodPlaceAccountRequest
	2, // 4: accounts.v1.accounts.Accounts.AddFoodPlaceAccount:input_type -> accounts.v1.accounts.AddFoodPlaceAccountRequest
	4, // 5: accounts.v1.accounts.Accounts.GetAccountDetails:input_type -> accounts.v1.accounts.GetAccountDetailsRequest
	3, // 6: accounts.v1.accounts.Accounts.AddPersonAccount:input_type -> accounts.v1.accounts.AddPersonAccountRequest
	1, // 7: accounts.v1.accounts.Accounts.VerifyFoodPlaceAccount:output_type -> accounts.v1.accounts.VerifyFoodPlaceAccountResponse
	7, // 8: accounts.v1.accounts.Accounts.AddFoodPlaceAccount:output_type -> google.protobuf.Empty
	5, // 9: accounts.v1.accounts.Accounts.GetAccountDetails:output_type -> accounts.v1.accounts.GetAccountDetailsResponse
	7, // 10: accounts.v1.accounts.Accounts.AddPersonAccount:output_type -> google.protobuf.Empty
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_accounts_proto_init() }
func file_protos_accounts_proto_init() {
	if File_protos_accounts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_accounts_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*VerifyFoodPlaceAccountRequest); i {
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
		file_protos_accounts_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*VerifyFoodPlaceAccountResponse); i {
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
		file_protos_accounts_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AddFoodPlaceAccountRequest); i {
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
		file_protos_accounts_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AddPersonAccountRequest); i {
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
		file_protos_accounts_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetAccountDetailsRequest); i {
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
		file_protos_accounts_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetAccountDetailsResponse); i {
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
	file_protos_accounts_proto_msgTypes[4].OneofWrappers = []any{
		(*GetAccountDetailsRequest_Username)(nil),
		(*GetAccountDetailsRequest_Email)(nil),
		(*GetAccountDetailsRequest_PhoneNumber)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_accounts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_accounts_proto_goTypes,
		DependencyIndexes: file_protos_accounts_proto_depIdxs,
		MessageInfos:      file_protos_accounts_proto_msgTypes,
	}.Build()
	File_protos_accounts_proto = out.File
	file_protos_accounts_proto_rawDesc = nil
	file_protos_accounts_proto_goTypes = nil
	file_protos_accounts_proto_depIdxs = nil
}