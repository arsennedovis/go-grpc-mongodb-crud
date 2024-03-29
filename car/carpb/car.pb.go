// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: car/carpb/car.proto

package carpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Car struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OwnerId string `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	GovId   string `protobuf:"bytes,3,opt,name=gov_id,json=govId,proto3" json:"gov_id,omitempty"`
	Brand   string `protobuf:"bytes,4,opt,name=brand,proto3" json:"brand,omitempty"`
}

func (x *Car) Reset() {
	*x = Car{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_carpb_car_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Car) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Car) ProtoMessage() {}

func (x *Car) ProtoReflect() protoreflect.Message {
	mi := &file_car_carpb_car_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Car.ProtoReflect.Descriptor instead.
func (*Car) Descriptor() ([]byte, []int) {
	return file_car_carpb_car_proto_rawDescGZIP(), []int{0}
}

func (x *Car) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Car) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *Car) GetGovId() string {
	if x != nil {
		return x.GovId
	}
	return ""
}

func (x *Car) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

type CreateCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car *Car `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
}

func (x *CreateCarRequest) Reset() {
	*x = CreateCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_carpb_car_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCarRequest) ProtoMessage() {}

func (x *CreateCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_carpb_car_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCarRequest.ProtoReflect.Descriptor instead.
func (*CreateCarRequest) Descriptor() ([]byte, []int) {
	return file_car_carpb_car_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCarRequest) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

type CreateCarResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car *Car `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"` //will have a blog id
}

func (x *CreateCarResponce) Reset() {
	*x = CreateCarResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_carpb_car_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCarResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCarResponce) ProtoMessage() {}

func (x *CreateCarResponce) ProtoReflect() protoreflect.Message {
	mi := &file_car_carpb_car_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCarResponce.ProtoReflect.Descriptor instead.
func (*CreateCarResponce) Descriptor() ([]byte, []int) {
	return file_car_carpb_car_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCarResponce) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

type ReadCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarId string `protobuf:"bytes,1,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
}

func (x *ReadCarRequest) Reset() {
	*x = ReadCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_carpb_car_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadCarRequest) ProtoMessage() {}

func (x *ReadCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_carpb_car_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadCarRequest.ProtoReflect.Descriptor instead.
func (*ReadCarRequest) Descriptor() ([]byte, []int) {
	return file_car_carpb_car_proto_rawDescGZIP(), []int{3}
}

func (x *ReadCarRequest) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

type ReadCarResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car *Car `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
}

func (x *ReadCarResponce) Reset() {
	*x = ReadCarResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_carpb_car_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadCarResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadCarResponce) ProtoMessage() {}

func (x *ReadCarResponce) ProtoReflect() protoreflect.Message {
	mi := &file_car_carpb_car_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadCarResponce.ProtoReflect.Descriptor instead.
func (*ReadCarResponce) Descriptor() ([]byte, []int) {
	return file_car_carpb_car_proto_rawDescGZIP(), []int{4}
}

func (x *ReadCarResponce) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

type UpdateCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car *Car `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
}

func (x *UpdateCarRequest) Reset() {
	*x = UpdateCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_carpb_car_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCarRequest) ProtoMessage() {}

func (x *UpdateCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_carpb_car_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCarRequest.ProtoReflect.Descriptor instead.
func (*UpdateCarRequest) Descriptor() ([]byte, []int) {
	return file_car_carpb_car_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateCarRequest) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

type UpdateCarResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car *Car `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
}

func (x *UpdateCarResponce) Reset() {
	*x = UpdateCarResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_carpb_car_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCarResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCarResponce) ProtoMessage() {}

func (x *UpdateCarResponce) ProtoReflect() protoreflect.Message {
	mi := &file_car_carpb_car_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCarResponce.ProtoReflect.Descriptor instead.
func (*UpdateCarResponce) Descriptor() ([]byte, []int) {
	return file_car_carpb_car_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCarResponce) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

type DeleteCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarId string `protobuf:"bytes,1,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
}

func (x *DeleteCarRequest) Reset() {
	*x = DeleteCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_carpb_car_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCarRequest) ProtoMessage() {}

func (x *DeleteCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_carpb_car_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCarRequest.ProtoReflect.Descriptor instead.
func (*DeleteCarRequest) Descriptor() ([]byte, []int) {
	return file_car_carpb_car_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteCarRequest) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

type DeleteCarResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarId string `protobuf:"bytes,1,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
}

func (x *DeleteCarResponce) Reset() {
	*x = DeleteCarResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_carpb_car_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCarResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCarResponce) ProtoMessage() {}

func (x *DeleteCarResponce) ProtoReflect() protoreflect.Message {
	mi := &file_car_carpb_car_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCarResponce.ProtoReflect.Descriptor instead.
func (*DeleteCarResponce) Descriptor() ([]byte, []int) {
	return file_car_carpb_car_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteCarResponce) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

type ListCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCarRequest) Reset() {
	*x = ListCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_carpb_car_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCarRequest) ProtoMessage() {}

func (x *ListCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_carpb_car_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCarRequest.ProtoReflect.Descriptor instead.
func (*ListCarRequest) Descriptor() ([]byte, []int) {
	return file_car_carpb_car_proto_rawDescGZIP(), []int{9}
}

type ListCarResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car *Car `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
}

func (x *ListCarResponce) Reset() {
	*x = ListCarResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_carpb_car_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCarResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCarResponce) ProtoMessage() {}

func (x *ListCarResponce) ProtoReflect() protoreflect.Message {
	mi := &file_car_carpb_car_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCarResponce.ProtoReflect.Descriptor instead.
func (*ListCarResponce) Descriptor() ([]byte, []int) {
	return file_car_carpb_car_proto_rawDescGZIP(), []int{10}
}

func (x *ListCarResponce) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

var File_car_carpb_car_proto protoreflect.FileDescriptor

var file_car_carpb_car_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x61, 0x72, 0x2f, 0x63, 0x61, 0x72, 0x70, 0x62, 0x2f, 0x63, 0x61, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63, 0x61, 0x72, 0x22, 0x5d, 0x0a, 0x03, 0x43, 0x61,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06,
	0x67, 0x6f, 0x76, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x6f,
	0x76, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x22, 0x2e, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x03, 0x63, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x63, 0x61, 0x72,
	0x2e, 0x43, 0x61, 0x72, 0x52, 0x03, 0x63, 0x61, 0x72, 0x22, 0x2f, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1a,
	0x0a, 0x03, 0x63, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x63, 0x61,
	0x72, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x03, 0x63, 0x61, 0x72, 0x22, 0x27, 0x0a, 0x0e, 0x52, 0x65,
	0x61, 0x64, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06,
	0x63, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61,
	0x72, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64, 0x43, 0x61, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x03, 0x63, 0x61, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x03, 0x63,
	0x61, 0x72, 0x22, 0x2e, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x03, 0x63, 0x61, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x03, 0x63,
	0x61, 0x72, 0x22, 0x2f, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x03, 0x63, 0x61, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x03,
	0x63, 0x61, 0x72, 0x22, 0x29, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x22, 0x2a,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x63, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2d, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x03, 0x63, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x63,
	0x61, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x03, 0x63, 0x61, 0x72, 0x32, 0xae, 0x02, 0x0a, 0x0a,
	0x43, 0x61, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x12, 0x15, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x52, 0x65, 0x61, 0x64, 0x43, 0x61,
	0x72, 0x12, 0x13, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x43, 0x61, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x52, 0x65, 0x61,
	0x64, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x09,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x12, 0x15, 0x2e, 0x63, 0x61, 0x72, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x61, 0x72, 0x12, 0x15, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63,
	0x61, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x12,
	0x13, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09,
	0x63, 0x61, 0x72, 0x2f, 0x63, 0x61, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_car_carpb_car_proto_rawDescOnce sync.Once
	file_car_carpb_car_proto_rawDescData = file_car_carpb_car_proto_rawDesc
)

func file_car_carpb_car_proto_rawDescGZIP() []byte {
	file_car_carpb_car_proto_rawDescOnce.Do(func() {
		file_car_carpb_car_proto_rawDescData = protoimpl.X.CompressGZIP(file_car_carpb_car_proto_rawDescData)
	})
	return file_car_carpb_car_proto_rawDescData
}

var file_car_carpb_car_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_car_carpb_car_proto_goTypes = []interface{}{
	(*Car)(nil),               // 0: car.Car
	(*CreateCarRequest)(nil),  // 1: car.CreateCarRequest
	(*CreateCarResponce)(nil), // 2: car.CreateCarResponce
	(*ReadCarRequest)(nil),    // 3: car.ReadCarRequest
	(*ReadCarResponce)(nil),   // 4: car.ReadCarResponce
	(*UpdateCarRequest)(nil),  // 5: car.UpdateCarRequest
	(*UpdateCarResponce)(nil), // 6: car.UpdateCarResponce
	(*DeleteCarRequest)(nil),  // 7: car.DeleteCarRequest
	(*DeleteCarResponce)(nil), // 8: car.DeleteCarResponce
	(*ListCarRequest)(nil),    // 9: car.ListCarRequest
	(*ListCarResponce)(nil),   // 10: car.ListCarResponce
}
var file_car_carpb_car_proto_depIdxs = []int32{
	0,  // 0: car.CreateCarRequest.car:type_name -> car.Car
	0,  // 1: car.CreateCarResponce.car:type_name -> car.Car
	0,  // 2: car.ReadCarResponce.car:type_name -> car.Car
	0,  // 3: car.UpdateCarRequest.car:type_name -> car.Car
	0,  // 4: car.UpdateCarResponce.car:type_name -> car.Car
	0,  // 5: car.ListCarResponce.car:type_name -> car.Car
	1,  // 6: car.CarService.CreateCar:input_type -> car.CreateCarRequest
	3,  // 7: car.CarService.ReadCar:input_type -> car.ReadCarRequest
	5,  // 8: car.CarService.UpdateCar:input_type -> car.UpdateCarRequest
	7,  // 9: car.CarService.DeleteCar:input_type -> car.DeleteCarRequest
	9,  // 10: car.CarService.ListCar:input_type -> car.ListCarRequest
	2,  // 11: car.CarService.CreateCar:output_type -> car.CreateCarResponce
	4,  // 12: car.CarService.ReadCar:output_type -> car.ReadCarResponce
	6,  // 13: car.CarService.UpdateCar:output_type -> car.UpdateCarResponce
	8,  // 14: car.CarService.DeleteCar:output_type -> car.DeleteCarResponce
	10, // 15: car.CarService.ListCar:output_type -> car.ListCarResponce
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_car_carpb_car_proto_init() }
func file_car_carpb_car_proto_init() {
	if File_car_carpb_car_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_car_carpb_car_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Car); i {
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
		file_car_carpb_car_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCarRequest); i {
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
		file_car_carpb_car_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCarResponce); i {
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
		file_car_carpb_car_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadCarRequest); i {
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
		file_car_carpb_car_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadCarResponce); i {
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
		file_car_carpb_car_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCarRequest); i {
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
		file_car_carpb_car_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCarResponce); i {
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
		file_car_carpb_car_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCarRequest); i {
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
		file_car_carpb_car_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCarResponce); i {
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
		file_car_carpb_car_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCarRequest); i {
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
		file_car_carpb_car_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCarResponce); i {
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
			RawDescriptor: file_car_carpb_car_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_car_carpb_car_proto_goTypes,
		DependencyIndexes: file_car_carpb_car_proto_depIdxs,
		MessageInfos:      file_car_carpb_car_proto_msgTypes,
	}.Build()
	File_car_carpb_car_proto = out.File
	file_car_carpb_car_proto_rawDesc = nil
	file_car_carpb_car_proto_goTypes = nil
	file_car_carpb_car_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CarServiceClient is the client API for CarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CarServiceClient interface {
	CreateCar(ctx context.Context, in *CreateCarRequest, opts ...grpc.CallOption) (*CreateCarResponce, error)
	ReadCar(ctx context.Context, in *ReadCarRequest, opts ...grpc.CallOption) (*ReadCarResponce, error)
	UpdateCar(ctx context.Context, in *UpdateCarRequest, opts ...grpc.CallOption) (*UpdateCarResponce, error)
	DeleteCar(ctx context.Context, in *DeleteCarRequest, opts ...grpc.CallOption) (*DeleteCarResponce, error)
	ListCar(ctx context.Context, in *ListCarRequest, opts ...grpc.CallOption) (CarService_ListCarClient, error)
}

type carServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCarServiceClient(cc grpc.ClientConnInterface) CarServiceClient {
	return &carServiceClient{cc}
}

func (c *carServiceClient) CreateCar(ctx context.Context, in *CreateCarRequest, opts ...grpc.CallOption) (*CreateCarResponce, error) {
	out := new(CreateCarResponce)
	err := c.cc.Invoke(ctx, "/car.CarService/CreateCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carServiceClient) ReadCar(ctx context.Context, in *ReadCarRequest, opts ...grpc.CallOption) (*ReadCarResponce, error) {
	out := new(ReadCarResponce)
	err := c.cc.Invoke(ctx, "/car.CarService/ReadCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carServiceClient) UpdateCar(ctx context.Context, in *UpdateCarRequest, opts ...grpc.CallOption) (*UpdateCarResponce, error) {
	out := new(UpdateCarResponce)
	err := c.cc.Invoke(ctx, "/car.CarService/UpdateCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carServiceClient) DeleteCar(ctx context.Context, in *DeleteCarRequest, opts ...grpc.CallOption) (*DeleteCarResponce, error) {
	out := new(DeleteCarResponce)
	err := c.cc.Invoke(ctx, "/car.CarService/DeleteCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carServiceClient) ListCar(ctx context.Context, in *ListCarRequest, opts ...grpc.CallOption) (CarService_ListCarClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CarService_serviceDesc.Streams[0], "/car.CarService/ListCar", opts...)
	if err != nil {
		return nil, err
	}
	x := &carServiceListCarClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CarService_ListCarClient interface {
	Recv() (*ListCarResponce, error)
	grpc.ClientStream
}

type carServiceListCarClient struct {
	grpc.ClientStream
}

func (x *carServiceListCarClient) Recv() (*ListCarResponce, error) {
	m := new(ListCarResponce)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CarServiceServer is the server API for CarService service.
type CarServiceServer interface {
	CreateCar(context.Context, *CreateCarRequest) (*CreateCarResponce, error)
	ReadCar(context.Context, *ReadCarRequest) (*ReadCarResponce, error)
	UpdateCar(context.Context, *UpdateCarRequest) (*UpdateCarResponce, error)
	DeleteCar(context.Context, *DeleteCarRequest) (*DeleteCarResponce, error)
	ListCar(*ListCarRequest, CarService_ListCarServer) error
}

// UnimplementedCarServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCarServiceServer struct {
}

func (*UnimplementedCarServiceServer) CreateCar(context.Context, *CreateCarRequest) (*CreateCarResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCar not implemented")
}
func (*UnimplementedCarServiceServer) ReadCar(context.Context, *ReadCarRequest) (*ReadCarResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadCar not implemented")
}
func (*UnimplementedCarServiceServer) UpdateCar(context.Context, *UpdateCarRequest) (*UpdateCarResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCar not implemented")
}
func (*UnimplementedCarServiceServer) DeleteCar(context.Context, *DeleteCarRequest) (*DeleteCarResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCar not implemented")
}
func (*UnimplementedCarServiceServer) ListCar(*ListCarRequest, CarService_ListCarServer) error {
	return status.Errorf(codes.Unimplemented, "method ListCar not implemented")
}

func RegisterCarServiceServer(s *grpc.Server, srv CarServiceServer) {
	s.RegisterService(&_CarService_serviceDesc, srv)
}

func _CarService_CreateCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServiceServer).CreateCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/car.CarService/CreateCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServiceServer).CreateCar(ctx, req.(*CreateCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarService_ReadCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServiceServer).ReadCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/car.CarService/ReadCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServiceServer).ReadCar(ctx, req.(*ReadCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarService_UpdateCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServiceServer).UpdateCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/car.CarService/UpdateCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServiceServer).UpdateCar(ctx, req.(*UpdateCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarService_DeleteCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServiceServer).DeleteCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/car.CarService/DeleteCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServiceServer).DeleteCar(ctx, req.(*DeleteCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarService_ListCar_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListCarRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CarServiceServer).ListCar(m, &carServiceListCarServer{stream})
}

type CarService_ListCarServer interface {
	Send(*ListCarResponce) error
	grpc.ServerStream
}

type carServiceListCarServer struct {
	grpc.ServerStream
}

func (x *carServiceListCarServer) Send(m *ListCarResponce) error {
	return x.ServerStream.SendMsg(m)
}

var _CarService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "car.CarService",
	HandlerType: (*CarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCar",
			Handler:    _CarService_CreateCar_Handler,
		},
		{
			MethodName: "ReadCar",
			Handler:    _CarService_ReadCar_Handler,
		},
		{
			MethodName: "UpdateCar",
			Handler:    _CarService_UpdateCar_Handler,
		},
		{
			MethodName: "DeleteCar",
			Handler:    _CarService_DeleteCar_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListCar",
			Handler:       _CarService_ListCar_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "car/carpb/car.proto",
}
