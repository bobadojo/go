// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.9
// source: bobadojo/stores/v1/stores.proto

package storespb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A store resource.
type Store struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique id (aka store number)
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// An identifier indicating the type of store
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	// Store name
	Title string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	// Street address
	Address string `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	// City
	City string `protobuf:"bytes,7,opt,name=city,proto3" json:"city,omitempty"`
	// State
	State string `protobuf:"bytes,8,opt,name=state,proto3" json:"state,omitempty"`
	// Zip code
	ZipCode int32 `protobuf:"varint,9,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	// String-encoded store hours
	StoreHours string `protobuf:"bytes,10,opt,name=store_hours,json=storeHours,proto3" json:"store_hours,omitempty"`
	// Latitude
	Latitude float32 `protobuf:"fixed32,11,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// Longitude
	Longitude float32 `protobuf:"fixed32,12,opt,name=longitude,proto3" json:"longitude,omitempty"`
	// Country
	RegionCode string `protobuf:"bytes,13,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	// County
	County string `protobuf:"bytes,14,opt,name=county,proto3" json:"county,omitempty"`
}

func (x *Store) Reset() {
	*x = Store{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bobadojo_stores_v1_stores_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Store) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Store) ProtoMessage() {}

func (x *Store) ProtoReflect() protoreflect.Message {
	mi := &file_bobadojo_stores_v1_stores_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Store.ProtoReflect.Descriptor instead.
func (*Store) Descriptor() ([]byte, []int) {
	return file_bobadojo_stores_v1_stores_proto_rawDescGZIP(), []int{0}
}

func (x *Store) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Store) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Store) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Store) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Store) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Store) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Store) GetZipCode() int32 {
	if x != nil {
		return x.ZipCode
	}
	return 0
}

func (x *Store) GetStoreHours() string {
	if x != nil {
		return x.StoreHours
	}
	return ""
}

func (x *Store) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Store) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Store) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *Store) GetCounty() string {
	if x != nil {
		return x.County
	}
	return ""
}

// Request to FindStores.
type FindStoresRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Bounding box of the request.
	BoundingBox *Box `protobuf:"bytes,1,opt,name=bounding_box,json=boundingBox,proto3" json:"bounding_box,omitempty"`
	// Center point of the request.
	Center *Point `protobuf:"bytes,2,opt,name=center,proto3" json:"center,omitempty"`
}

func (x *FindStoresRequest) Reset() {
	*x = FindStoresRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bobadojo_stores_v1_stores_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindStoresRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindStoresRequest) ProtoMessage() {}

func (x *FindStoresRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bobadojo_stores_v1_stores_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindStoresRequest.ProtoReflect.Descriptor instead.
func (*FindStoresRequest) Descriptor() ([]byte, []int) {
	return file_bobadojo_stores_v1_stores_proto_rawDescGZIP(), []int{1}
}

func (x *FindStoresRequest) GetBoundingBox() *Box {
	if x != nil {
		return x.BoundingBox
	}
	return nil
}

func (x *FindStoresRequest) GetCenter() *Point {
	if x != nil {
		return x.Center
	}
	return nil
}

// Response from FindStores.
type FindStoresResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Stores in the datastore.
	Stores []*Store `protobuf:"bytes,1,rep,name=stores,proto3" json:"stores,omitempty"`
}

func (x *FindStoresResponse) Reset() {
	*x = FindStoresResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bobadojo_stores_v1_stores_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindStoresResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindStoresResponse) ProtoMessage() {}

func (x *FindStoresResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bobadojo_stores_v1_stores_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindStoresResponse.ProtoReflect.Descriptor instead.
func (*FindStoresResponse) Descriptor() ([]byte, []int) {
	return file_bobadojo_stores_v1_stores_proto_rawDescGZIP(), []int{2}
}

func (x *FindStoresResponse) GetStores() []*Store {
	if x != nil {
		return x.Stores
	}
	return nil
}

// Request to GetStore.
type GetStoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the store resource to retrieve.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetStoreRequest) Reset() {
	*x = GetStoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bobadojo_stores_v1_stores_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStoreRequest) ProtoMessage() {}

func (x *GetStoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bobadojo_stores_v1_stores_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStoreRequest.ProtoReflect.Descriptor instead.
func (*GetStoreRequest) Descriptor() ([]byte, []int) {
	return file_bobadojo_stores_v1_stores_proto_rawDescGZIP(), []int{3}
}

func (x *GetStoreRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// A bounding box.
type Box struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Maximum coordinate values.
	Max *Point `protobuf:"bytes,1,opt,name=max,proto3" json:"max,omitempty"`
	// Minimum coordinate values.
	Min *Point `protobuf:"bytes,2,opt,name=min,proto3" json:"min,omitempty"`
}

func (x *Box) Reset() {
	*x = Box{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bobadojo_stores_v1_stores_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Box) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Box) ProtoMessage() {}

func (x *Box) ProtoReflect() protoreflect.Message {
	mi := &file_bobadojo_stores_v1_stores_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Box.ProtoReflect.Descriptor instead.
func (*Box) Descriptor() ([]byte, []int) {
	return file_bobadojo_stores_v1_stores_proto_rawDescGZIP(), []int{4}
}

func (x *Box) GetMax() *Point {
	if x != nil {
		return x.Max
	}
	return nil
}

func (x *Box) GetMin() *Point {
	if x != nil {
		return x.Min
	}
	return nil
}

// A point.
type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Latitude of the point.
	Latitude float32 `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// Longitude of the point.
	Longitude float32 `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bobadojo_stores_v1_stores_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_bobadojo_stores_v1_stores_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_bobadojo_stores_v1_stores_proto_rawDescGZIP(), []int{5}
}

func (x *Point) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Point) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

var File_bobadojo_stores_v1_stores_proto protoreflect.FileDescriptor

var file_bobadojo_stores_v1_stores_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x62, 0x6f, 0x62, 0x61, 0x64, 0x6f, 0x6a, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x62, 0x6f, 0x62, 0x61, 0x64, 0x6f, 0x6a, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x02, 0x0a, 0x05, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x48, 0x6f, 0x75, 0x72,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x79, 0x22, 0x82, 0x01, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0c, 0x62, 0x6f,
	0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x6f, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x62, 0x6f, 0x62, 0x61, 0x64, 0x6f, 0x6a, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x78, 0x52, 0x0b, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78, 0x12, 0x31, 0x0a, 0x06, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x6f, 0x62, 0x61, 0x64, 0x6f, 0x6a,
	0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x06, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x47, 0x0a, 0x12, 0x46, 0x69, 0x6e,
	0x64, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x31, 0x0a, 0x06, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x62, 0x6f, 0x62, 0x61, 0x64, 0x6f, 0x6a, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x06, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x73, 0x22, 0x48, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x21, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x1b, 0x0a, 0x19, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x73, 0x2e, 0x62, 0x6f, 0x62, 0x61, 0x64, 0x6f, 0x6a, 0x6f, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5f, 0x0a, 0x03,
	0x42, 0x6f, 0x78, 0x12, 0x2b, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x62, 0x6f, 0x62, 0x61, 0x64, 0x6f, 0x6a, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x03, 0x6d, 0x61, 0x78,
	0x12, 0x2b, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x62, 0x6f, 0x62, 0x61, 0x64, 0x6f, 0x6a, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x22, 0x41, 0x0a,
	0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x32, 0xf2, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x74, 0x0a, 0x0a, 0x46,
	0x69, 0x6e, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x25, 0x2e, 0x62, 0x6f, 0x62, 0x61,
	0x64, 0x6f, 0x6a, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x62, 0x6f, 0x62, 0x61, 0x64, 0x6f, 0x6a, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11,
	0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x3a, 0x66, 0x69, 0x6e,
	0x64, 0x12, 0x72, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x23, 0x2e,
	0x62, 0x6f, 0x62, 0x61, 0x64, 0x6f, 0x6a, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x6f, 0x62, 0x61, 0x64, 0x6f, 0x6a, 0x6f, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x22, 0x26, 0xda,
	0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x7d, 0x42, 0x5f, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x6f, 0x62,
	0x61, 0x64, 0x6f, 0x6a, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42,
	0x0b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6f, 0x62, 0x61, 0x64,
	0x6f, 0x6a, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x70, 0x62, 0x3b, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bobadojo_stores_v1_stores_proto_rawDescOnce sync.Once
	file_bobadojo_stores_v1_stores_proto_rawDescData = file_bobadojo_stores_v1_stores_proto_rawDesc
)

func file_bobadojo_stores_v1_stores_proto_rawDescGZIP() []byte {
	file_bobadojo_stores_v1_stores_proto_rawDescOnce.Do(func() {
		file_bobadojo_stores_v1_stores_proto_rawDescData = protoimpl.X.CompressGZIP(file_bobadojo_stores_v1_stores_proto_rawDescData)
	})
	return file_bobadojo_stores_v1_stores_proto_rawDescData
}

var file_bobadojo_stores_v1_stores_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_bobadojo_stores_v1_stores_proto_goTypes = []interface{}{
	(*Store)(nil),              // 0: bobadojo.stores.v1.Store
	(*FindStoresRequest)(nil),  // 1: bobadojo.stores.v1.FindStoresRequest
	(*FindStoresResponse)(nil), // 2: bobadojo.stores.v1.FindStoresResponse
	(*GetStoreRequest)(nil),    // 3: bobadojo.stores.v1.GetStoreRequest
	(*Box)(nil),                // 4: bobadojo.stores.v1.Box
	(*Point)(nil),              // 5: bobadojo.stores.v1.Point
}
var file_bobadojo_stores_v1_stores_proto_depIdxs = []int32{
	4, // 0: bobadojo.stores.v1.FindStoresRequest.bounding_box:type_name -> bobadojo.stores.v1.Box
	5, // 1: bobadojo.stores.v1.FindStoresRequest.center:type_name -> bobadojo.stores.v1.Point
	0, // 2: bobadojo.stores.v1.FindStoresResponse.stores:type_name -> bobadojo.stores.v1.Store
	5, // 3: bobadojo.stores.v1.Box.max:type_name -> bobadojo.stores.v1.Point
	5, // 4: bobadojo.stores.v1.Box.min:type_name -> bobadojo.stores.v1.Point
	1, // 5: bobadojo.stores.v1.Stores.FindStores:input_type -> bobadojo.stores.v1.FindStoresRequest
	3, // 6: bobadojo.stores.v1.Stores.GetStore:input_type -> bobadojo.stores.v1.GetStoreRequest
	2, // 7: bobadojo.stores.v1.Stores.FindStores:output_type -> bobadojo.stores.v1.FindStoresResponse
	0, // 8: bobadojo.stores.v1.Stores.GetStore:output_type -> bobadojo.stores.v1.Store
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_bobadojo_stores_v1_stores_proto_init() }
func file_bobadojo_stores_v1_stores_proto_init() {
	if File_bobadojo_stores_v1_stores_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bobadojo_stores_v1_stores_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Store); i {
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
		file_bobadojo_stores_v1_stores_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindStoresRequest); i {
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
		file_bobadojo_stores_v1_stores_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindStoresResponse); i {
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
		file_bobadojo_stores_v1_stores_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStoreRequest); i {
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
		file_bobadojo_stores_v1_stores_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Box); i {
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
		file_bobadojo_stores_v1_stores_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
			RawDescriptor: file_bobadojo_stores_v1_stores_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bobadojo_stores_v1_stores_proto_goTypes,
		DependencyIndexes: file_bobadojo_stores_v1_stores_proto_depIdxs,
		MessageInfos:      file_bobadojo_stores_v1_stores_proto_msgTypes,
	}.Build()
	File_bobadojo_stores_v1_stores_proto = out.File
	file_bobadojo_stores_v1_stores_proto_rawDesc = nil
	file_bobadojo_stores_v1_stores_proto_goTypes = nil
	file_bobadojo_stores_v1_stores_proto_depIdxs = nil
}
