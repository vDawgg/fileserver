// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: unified.proto

package fileserver

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

type UploadStatusCode int32

const (
	UploadStatusCode_Unknown UploadStatusCode = 0
	UploadStatusCode_Ok      UploadStatusCode = 1
	UploadStatusCode_Failed  UploadStatusCode = 2
)

// Enum value maps for UploadStatusCode.
var (
	UploadStatusCode_name = map[int32]string{
		0: "Unknown",
		1: "Ok",
		2: "Failed",
	}
	UploadStatusCode_value = map[string]int32{
		"Unknown": 0,
		"Ok":      1,
		"Failed":  2,
	}
)

func (x UploadStatusCode) Enum() *UploadStatusCode {
	p := new(UploadStatusCode)
	*p = x
	return p
}

func (x UploadStatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UploadStatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_unified_proto_enumTypes[0].Descriptor()
}

func (UploadStatusCode) Type() protoreflect.EnumType {
	return &file_unified_proto_enumTypes[0]
}

func (x UploadStatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UploadStatusCode.Descriptor instead.
func (UploadStatusCode) EnumDescriptor() ([]byte, []int) {
	return file_unified_proto_rawDescGZIP(), []int{0}
}

type AuthStatus int32

const (
	AuthStatus_OK     AuthStatus = 0
	AuthStatus_FAILED AuthStatus = 1
)

// Enum value maps for AuthStatus.
var (
	AuthStatus_name = map[int32]string{
		0: "OK",
		1: "FAILED",
	}
	AuthStatus_value = map[string]int32{
		"OK":     0,
		"FAILED": 1,
	}
)

func (x AuthStatus) Enum() *AuthStatus {
	p := new(AuthStatus)
	*p = x
	return p
}

func (x AuthStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_unified_proto_enumTypes[1].Descriptor()
}

func (AuthStatus) Type() protoreflect.EnumType {
	return &file_unified_proto_enumTypes[1]
}

func (x AuthStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthStatus.Descriptor instead.
func (AuthStatus) EnumDescriptor() ([]byte, []int) {
	return file_unified_proto_rawDescGZIP(), []int{1}
}

type FileDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=FileName,proto3" json:"FileName,omitempty"`
	Bucket   string `protobuf:"bytes,2,opt,name=Bucket,proto3" json:"Bucket,omitempty"` //Probably not a good idea to sen password all the time
}

func (x *FileDescription) Reset() {
	*x = FileDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unified_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileDescription) ProtoMessage() {}

func (x *FileDescription) ProtoReflect() protoreflect.Message {
	mi := &file_unified_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileDescription.ProtoReflect.Descriptor instead.
func (*FileDescription) Descriptor() ([]byte, []int) {
	return file_unified_proto_rawDescGZIP(), []int{0}
}

func (x *FileDescription) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileDescription) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

type UploadStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code UploadStatusCode `protobuf:"varint,1,opt,name=Code,proto3,enum=fileserver.UploadStatusCode" json:"Code,omitempty"`
}

func (x *UploadStatus) Reset() {
	*x = UploadStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unified_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadStatus) ProtoMessage() {}

func (x *UploadStatus) ProtoReflect() protoreflect.Message {
	mi := &file_unified_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadStatus.ProtoReflect.Descriptor instead.
func (*UploadStatus) Descriptor() ([]byte, []int) {
	return file_unified_proto_rawDescGZIP(), []int{1}
}

func (x *UploadStatus) GetCode() UploadStatusCode {
	if x != nil {
		return x.Code
	}
	return UploadStatusCode_Unknown
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileDescription *FileDescription `protobuf:"bytes,1,opt,name=FileDescription,proto3" json:"FileDescription,omitempty"` //For metadata -> Also needs description of where the file is actually saved -> maybe just do something like collections for now
	Content         []byte           `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unified_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_unified_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_unified_proto_rawDescGZIP(), []int{2}
}

func (x *Chunk) GetFileDescription() *FileDescription {
	if x != nil {
		return x.FileDescription
	}
	return nil
}

func (x *Chunk) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type StructureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bucket    string `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Directory string `protobuf:"bytes,2,opt,name=directory,proto3" json:"directory,omitempty"`
}

func (x *StructureRequest) Reset() {
	*x = StructureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unified_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StructureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StructureRequest) ProtoMessage() {}

func (x *StructureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_unified_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StructureRequest.ProtoReflect.Descriptor instead.
func (*StructureRequest) Descriptor() ([]byte, []int) {
	return file_unified_proto_rawDescGZIP(), []int{3}
}

func (x *StructureRequest) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *StructureRequest) GetDirectory() string {
	if x != nil {
		return x.Directory
	}
	return ""
}

type Directory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	FileName []string `protobuf:"bytes,2,rep,name=FileName,proto3" json:"FileName,omitempty"`
}

func (x *Directory) Reset() {
	*x = Directory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unified_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Directory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Directory) ProtoMessage() {}

func (x *Directory) ProtoReflect() protoreflect.Message {
	mi := &file_unified_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Directory.ProtoReflect.Descriptor instead.
func (*Directory) Descriptor() ([]byte, []int) {
	return file_unified_proto_rawDescGZIP(), []int{4}
}

func (x *Directory) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Directory) GetFileName() []string {
	if x != nil {
		return x.FileName
	}
	return nil
}

type Structure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object []*Object `protobuf:"bytes,1,rep,name=Object,proto3" json:"Object,omitempty"`
}

func (x *Structure) Reset() {
	*x = Structure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unified_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Structure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Structure) ProtoMessage() {}

func (x *Structure) ProtoReflect() protoreflect.Message {
	mi := &file_unified_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Structure.ProtoReflect.Descriptor instead.
func (*Structure) Descriptor() ([]byte, []int) {
	return file_unified_proto_rawDescGZIP(), []int{5}
}

func (x *Structure) GetObject() []*Object {
	if x != nil {
		return x.Object
	}
	return nil
}

type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unified_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_unified_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_unified_proto_rawDescGZIP(), []int{6}
}

func (x *Object) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Object) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type DownloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileDescription *FileDescription `protobuf:"bytes,1,opt,name=FileDescription,proto3" json:"FileDescription,omitempty"`
}

func (x *DownloadRequest) Reset() {
	*x = DownloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unified_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadRequest) ProtoMessage() {}

func (x *DownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_unified_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadRequest.ProtoReflect.Descriptor instead.
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return file_unified_proto_rawDescGZIP(), []int{7}
}

func (x *DownloadRequest) GetFileDescription() *FileDescription {
	if x != nil {
		return x.FileDescription
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unified_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_unified_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_unified_proto_rawDescGZIP(), []int{8}
}

type Keys struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys string `protobuf:"bytes,1,opt,name=keys,proto3" json:"keys,omitempty"`
}

func (x *Keys) Reset() {
	*x = Keys{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unified_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Keys) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Keys) ProtoMessage() {}

func (x *Keys) ProtoReflect() protoreflect.Message {
	mi := &file_unified_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Keys.ProtoReflect.Descriptor instead.
func (*Keys) Descriptor() ([]byte, []int) {
	return file_unified_proto_rawDescGZIP(), []int{9}
}

func (x *Keys) GetKeys() string {
	if x != nil {
		return x.Keys
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     []byte `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password []byte `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unified_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_unified_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_unified_proto_rawDescGZIP(), []int{10}
}

func (x *User) GetName() []byte {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *User) GetPassword() []byte {
	if x != nil {
		return x.Password
	}
	return nil
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string     `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Status AuthStatus `protobuf:"varint,2,opt,name=status,proto3,enum=fileserver.AuthStatus" json:"status,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unified_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_unified_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_unified_proto_rawDescGZIP(), []int{11}
}

func (x *Token) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Token) GetStatus() AuthStatus {
	if x != nil {
		return x.Status
	}
	return AuthStatus_OK
}

type AuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Access []string `protobuf:"bytes,2,rep,name=access,proto3" json:"access,omitempty"`
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unified_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_unified_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_unified_proto_rawDescGZIP(), []int{12}
}

func (x *AuthRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AuthRequest) GetAccess() []string {
	if x != nil {
		return x.Access
	}
	return nil
}

type AuthReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAuthorized bool `protobuf:"varint,1,opt,name=isAuthorized,proto3" json:"isAuthorized,omitempty"`
}

func (x *AuthReply) Reset() {
	*x = AuthReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unified_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthReply) ProtoMessage() {}

func (x *AuthReply) ProtoReflect() protoreflect.Message {
	mi := &file_unified_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthReply.ProtoReflect.Descriptor instead.
func (*AuthReply) Descriptor() ([]byte, []int) {
	return file_unified_proto_rawDescGZIP(), []int{13}
}

func (x *AuthReply) GetIsAuthorized() bool {
	if x != nil {
		return x.IsAuthorized
	}
	return false
}

var File_unified_proto protoreflect.FileDescriptor

var file_unified_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x45, 0x0a, 0x0f, 0x46,
	0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x22, 0x40, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x30, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1c, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x22, 0x68, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x45, 0x0a,
	0x0f, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x48,
	0x0a, 0x10, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x3b, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x09, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x30,
	0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x58, 0x0a, 0x0f, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0f, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x46, 0x69, 0x6c, 0x65, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x1a, 0x0a, 0x04, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6b,
	0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22,
	0x36, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x4d, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3b, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x22, 0x2f, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x65, 0x64, 0x2a, 0x33, 0x0a, 0x10, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x6b, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x02, 0x2a, 0x20, 0x0a, 0x0a, 0x41, 0x75, 0x74,
	0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x32, 0xd0, 0x01, 0x0a, 0x09,
	0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x09, 0x73, 0x61, 0x76,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x11, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x18, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x00, 0x28, 0x01, 0x12, 0x45, 0x0a, 0x0c, 0x67, 0x65, 0x74, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1c, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x22, 0x00, 0x12, 0x3e,
	0x0a, 0x08, 0x67, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x1b, 0x2e, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x00, 0x30, 0x01, 0x32, 0x71,
	0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x30, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x11, 0x2e, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4b, 0x65, 0x79, 0x73, 0x22,
	0x00, 0x12, 0x2e, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x2e, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x00, 0x32, 0x4e, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x12,
	0x40, 0x0a, 0x0c, 0x69, 0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x12,
	0x17, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x3b, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_unified_proto_rawDescOnce sync.Once
	file_unified_proto_rawDescData = file_unified_proto_rawDesc
)

func file_unified_proto_rawDescGZIP() []byte {
	file_unified_proto_rawDescOnce.Do(func() {
		file_unified_proto_rawDescData = protoimpl.X.CompressGZIP(file_unified_proto_rawDescData)
	})
	return file_unified_proto_rawDescData
}

var file_unified_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_unified_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_unified_proto_goTypes = []interface{}{
	(UploadStatusCode)(0),    // 0: fileserver.UploadStatusCode
	(AuthStatus)(0),          // 1: fileserver.AuthStatus
	(*FileDescription)(nil),  // 2: fileserver.FileDescription
	(*UploadStatus)(nil),     // 3: fileserver.UploadStatus
	(*Chunk)(nil),            // 4: fileserver.Chunk
	(*StructureRequest)(nil), // 5: fileserver.StructureRequest
	(*Directory)(nil),        // 6: fileserver.Directory
	(*Structure)(nil),        // 7: fileserver.Structure
	(*Object)(nil),           // 8: fileserver.Object
	(*DownloadRequest)(nil),  // 9: fileserver.DownloadRequest
	(*Empty)(nil),            // 10: fileserver.Empty
	(*Keys)(nil),             // 11: fileserver.Keys
	(*User)(nil),             // 12: fileserver.User
	(*Token)(nil),            // 13: fileserver.Token
	(*AuthRequest)(nil),      // 14: fileserver.AuthRequest
	(*AuthReply)(nil),        // 15: fileserver.AuthReply
}
var file_unified_proto_depIdxs = []int32{
	0,  // 0: fileserver.UploadStatus.Code:type_name -> fileserver.UploadStatusCode
	2,  // 1: fileserver.Chunk.FileDescription:type_name -> fileserver.FileDescription
	8,  // 2: fileserver.Structure.Object:type_name -> fileserver.Object
	2,  // 3: fileserver.DownloadRequest.FileDescription:type_name -> fileserver.FileDescription
	1,  // 4: fileserver.Token.status:type_name -> fileserver.AuthStatus
	4,  // 5: fileserver.Retriever.saveFiles:input_type -> fileserver.Chunk
	5,  // 6: fileserver.Retriever.getStructure:input_type -> fileserver.StructureRequest
	9,  // 7: fileserver.Retriever.getFiles:input_type -> fileserver.DownloadRequest
	10, // 8: fileserver.authenticator.getKeys:input_type -> fileserver.Empty
	12, // 9: fileserver.authenticator.login:input_type -> fileserver.User
	14, // 10: fileserver.authorizer.isAuthorized:input_type -> fileserver.AuthRequest
	3,  // 11: fileserver.Retriever.saveFiles:output_type -> fileserver.UploadStatus
	7,  // 12: fileserver.Retriever.getStructure:output_type -> fileserver.Structure
	4,  // 13: fileserver.Retriever.getFiles:output_type -> fileserver.Chunk
	11, // 14: fileserver.authenticator.getKeys:output_type -> fileserver.Keys
	13, // 15: fileserver.authenticator.login:output_type -> fileserver.Token
	15, // 16: fileserver.authorizer.isAuthorized:output_type -> fileserver.AuthReply
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_unified_proto_init() }
func file_unified_proto_init() {
	if File_unified_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_unified_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileDescription); i {
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
		file_unified_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadStatus); i {
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
		file_unified_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_unified_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StructureRequest); i {
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
		file_unified_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Directory); i {
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
		file_unified_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Structure); i {
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
		file_unified_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
		file_unified_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadRequest); i {
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
		file_unified_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_unified_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Keys); i {
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
		file_unified_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_unified_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_unified_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRequest); i {
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
		file_unified_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthReply); i {
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
			RawDescriptor: file_unified_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_unified_proto_goTypes,
		DependencyIndexes: file_unified_proto_depIdxs,
		EnumInfos:         file_unified_proto_enumTypes,
		MessageInfos:      file_unified_proto_msgTypes,
	}.Build()
	File_unified_proto = out.File
	file_unified_proto_rawDesc = nil
	file_unified_proto_goTypes = nil
	file_unified_proto_depIdxs = nil
}
