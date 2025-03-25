// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.19.4
// source: product.proto

package product

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

<<<<<<< HEAD
type ProductDetailRequest struct {
=======
type ProductDetailReq struct {
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId uint64 `protobuf:"varint,1,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
}

<<<<<<< HEAD
func (x *ProductDetailRequest) Reset() {
	*x = ProductDetailRequest{}
=======
func (x *ProductDetailReq) Reset() {
	*x = ProductDetailReq{}
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

<<<<<<< HEAD
func (x *ProductDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductDetailRequest) ProtoMessage() {}

func (x *ProductDetailRequest) ProtoReflect() protoreflect.Message {
=======
func (x *ProductDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductDetailReq) ProtoMessage() {}

func (x *ProductDetailReq) ProtoReflect() protoreflect.Message {
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
	mi := &file_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

<<<<<<< HEAD
// Deprecated: Use ProductDetailRequest.ProtoReflect.Descriptor instead.
func (*ProductDetailRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

func (x *ProductDetailRequest) GetProductId() uint64 {
=======
// Deprecated: Use ProductDetailReq.ProtoReflect.Descriptor instead.
func (*ProductDetailReq) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

func (x *ProductDetailReq) GetProductId() uint64 {
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
	if x != nil {
		return x.ProductId
	}
	return 0
}

<<<<<<< HEAD
type ProductDetailResponse struct {
=======
type ProductDetailResp struct {
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

<<<<<<< HEAD
	Id        int64   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`               //编号
	Image     string  `protobuf:"bytes,2,opt,name=Image,proto3" json:"Image,omitempty"`          //图片
	StoreName string  `protobuf:"bytes,3,opt,name=StoreName,proto3" json:"StoreName,omitempty"`  //商品名称
	StoreInfo string  `protobuf:"bytes,4,opt,name=StoreInfo,proto3" json:"StoreInfo,omitempty"`  //商品简介
	Price     float32 `protobuf:"fixed32,5,opt,name=Price,proto3" json:"Price,omitempty"`        //商品价格
	Postage   float32 `protobuf:"fixed32,6,opt,name=Postage,proto3" json:"Postage,omitempty"`    //邮费
	IsPostage int64   `protobuf:"varint,7,opt,name=IsPostage,proto3" json:"IsPostage,omitempty"` //是否包邮
	Stock     int64   `protobuf:"varint,8,opt,name=Stock,proto3" json:"Stock,omitempty"`         //库存
	Browse    int64   `protobuf:"varint,9,opt,name=Browse,proto3" json:"Browse,omitempty"`       //浏览量
}

func (x *ProductDetailResponse) Reset() {
	*x = ProductDetailResponse{}
=======
	MerId     uint64  `protobuf:"varint,1,opt,name=MerId,proto3" json:"MerId,omitempty"`
	Image     string  `protobuf:"bytes,2,opt,name=Image,proto3" json:"Image,omitempty"`
	StoreName string  `protobuf:"bytes,3,opt,name=StoreName,proto3" json:"StoreName,omitempty"`
	StoreInfo string  `protobuf:"bytes,4,opt,name=StoreInfo,proto3" json:"StoreInfo,omitempty"`
	Price     float32 `protobuf:"fixed32,5,opt,name=Price,proto3" json:"Price,omitempty"`
	VipPrice  float32 `protobuf:"fixed32,6,opt,name=VipPrice,proto3" json:"VipPrice,omitempty"`
	Postage   float32 `protobuf:"fixed32,7,opt,name=Postage,proto3" json:"Postage,omitempty"`
	Stock     uint64  `protobuf:"varint,8,opt,name=Stock,proto3" json:"Stock,omitempty"`
	IsShow    uint64  `protobuf:"varint,9,opt,name=IsShow,proto3" json:"IsShow,omitempty"`
	IsPostage uint64  `protobuf:"varint,10,opt,name=IsPostage,proto3" json:"IsPostage,omitempty"`
	IsSeckill uint64  `protobuf:"varint,11,opt,name=IsSeckill,proto3" json:"IsSeckill,omitempty"`
	Browse    uint64  `protobuf:"varint,12,opt,name=Browse,proto3" json:"Browse,omitempty"`
}

func (x *ProductDetailResp) Reset() {
	*x = ProductDetailResp{}
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

<<<<<<< HEAD
func (x *ProductDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductDetailResponse) ProtoMessage() {}

func (x *ProductDetailResponse) ProtoReflect() protoreflect.Message {
=======
func (x *ProductDetailResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductDetailResp) ProtoMessage() {}

func (x *ProductDetailResp) ProtoReflect() protoreflect.Message {
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
	mi := &file_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

<<<<<<< HEAD
// Deprecated: Use ProductDetailResponse.ProtoReflect.Descriptor instead.
func (*ProductDetailResponse) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductDetailResponse) GetId() int64 {
	if x != nil {
		return x.Id
=======
// Deprecated: Use ProductDetailResp.ProtoReflect.Descriptor instead.
func (*ProductDetailResp) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductDetailResp) GetMerId() uint64 {
	if x != nil {
		return x.MerId
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
	}
	return 0
}

<<<<<<< HEAD
func (x *ProductDetailResponse) GetImage() string {
=======
func (x *ProductDetailResp) GetImage() string {
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
	if x != nil {
		return x.Image
	}
	return ""
}

<<<<<<< HEAD
func (x *ProductDetailResponse) GetStoreName() string {
=======
func (x *ProductDetailResp) GetStoreName() string {
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
	if x != nil {
		return x.StoreName
	}
	return ""
}

<<<<<<< HEAD
func (x *ProductDetailResponse) GetStoreInfo() string {
=======
func (x *ProductDetailResp) GetStoreInfo() string {
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
	if x != nil {
		return x.StoreInfo
	}
	return ""
}

<<<<<<< HEAD
func (x *ProductDetailResponse) GetPrice() float32 {
=======
func (x *ProductDetailResp) GetPrice() float32 {
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
	if x != nil {
		return x.Price
	}
	return 0
}

<<<<<<< HEAD
func (x *ProductDetailResponse) GetPostage() float32 {
=======
func (x *ProductDetailResp) GetVipPrice() float32 {
	if x != nil {
		return x.VipPrice
	}
	return 0
}

func (x *ProductDetailResp) GetPostage() float32 {
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
	if x != nil {
		return x.Postage
	}
	return 0
}

<<<<<<< HEAD
func (x *ProductDetailResponse) GetIsPostage() int64 {
	if x != nil {
		return x.IsPostage
	}
	return 0
}

func (x *ProductDetailResponse) GetStock() int64 {
=======
func (x *ProductDetailResp) GetStock() uint64 {
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
	if x != nil {
		return x.Stock
	}
	return 0
}

<<<<<<< HEAD
func (x *ProductDetailResponse) GetBrowse() int64 {
=======
func (x *ProductDetailResp) GetIsShow() uint64 {
	if x != nil {
		return x.IsShow
	}
	return 0
}

func (x *ProductDetailResp) GetIsPostage() uint64 {
	if x != nil {
		return x.IsPostage
	}
	return 0
}

func (x *ProductDetailResp) GetIsSeckill() uint64 {
	if x != nil {
		return x.IsSeckill
	}
	return 0
}

func (x *ProductDetailResp) GetBrowse() uint64 {
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
	if x != nil {
		return x.Browse
	}
	return 0
}

<<<<<<< HEAD
type EsAddProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EsAddProductRequest) Reset() {
	*x = EsAddProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EsAddProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EsAddProductRequest) ProtoMessage() {}

func (x *EsAddProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EsAddProductRequest.ProtoReflect.Descriptor instead.
func (*EsAddProductRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{2}
}

type EsAddProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *EsAddProductResponse) Reset() {
	*x = EsAddProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EsAddProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EsAddProductResponse) ProtoMessage() {}

func (x *EsAddProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EsAddProductResponse.ProtoReflect.Descriptor instead.
func (*EsAddProductResponse) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{3}
}

func (x *EsAddProductResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetAllProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllProductRequest) Reset() {
	*x = GetAllProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductRequest) ProtoMessage() {}

func (x *GetAllProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductRequest.ProtoReflect.Descriptor instead.
func (*GetAllProductRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{4}
}

type GetAllProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllProductResponse []*ProductDetailResponse `protobuf:"bytes,1,rep,name=AllProductResponse,proto3" json:"AllProductResponse,omitempty"`
}

func (x *GetAllProductResponse) Reset() {
	*x = GetAllProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductResponse) ProtoMessage() {}

func (x *GetAllProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductResponse.ProtoReflect.Descriptor instead.
func (*GetAllProductResponse) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllProductResponse) GetAllProductResponse() []*ProductDetailResponse {
	if x != nil {
		return x.AllProductResponse
	}
	return nil
}

=======
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
var File_product_proto protoreflect.FileDescriptor

var file_product_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
<<<<<<< HEAD
	0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x34, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0xf5,
	0x01, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x07, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73,
	0x50, 0x6f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x49,
	0x73, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x16,
	0x0a, 0x06, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x45, 0x73, 0x41, 0x64, 0x64, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x30, 0x0a,
	0x14, 0x45, 0x73, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x67, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4e, 0x0a, 0x12, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x12, 0x41, 0x6c,
	0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xf6, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x4e, 0x0a, 0x0d,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0c,
	0x45, 0x73, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1c, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x45, 0x73, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x45, 0x73, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x70,
=======
	0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x30, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0xc9, 0x02, 0x0a, 0x11, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x14, 0x0a, 0x05, 0x4d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x4d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x56, 0x69, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x08, 0x56, 0x69, 0x70, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x6f,
	0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x50, 0x6f, 0x73,
	0x74, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x73,
	0x53, 0x68, 0x6f, 0x77, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x49, 0x73, 0x53, 0x68,
	0x6f, 0x77, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x49, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x53, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x49, 0x73, 0x53, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x32, 0x51, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x46, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x70,
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_proto_rawDescOnce sync.Once
	file_product_proto_rawDescData = file_product_proto_rawDesc
)

func file_product_proto_rawDescGZIP() []byte {
	file_product_proto_rawDescOnce.Do(func() {
		file_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_proto_rawDescData)
	})
	return file_product_proto_rawDescData
}

<<<<<<< HEAD
var file_product_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_product_proto_goTypes = []any{
	(*ProductDetailRequest)(nil),  // 0: product.ProductDetailRequest
	(*ProductDetailResponse)(nil), // 1: product.ProductDetailResponse
	(*EsAddProductRequest)(nil),   // 2: product.EsAddProductRequest
	(*EsAddProductResponse)(nil),  // 3: product.EsAddProductResponse
	(*GetAllProductRequest)(nil),  // 4: product.GetAllProductRequest
	(*GetAllProductResponse)(nil), // 5: product.GetAllProductResponse
}
var file_product_proto_depIdxs = []int32{
	1, // 0: product.GetAllProductResponse.AllProductResponse:type_name -> product.ProductDetailResponse
	0, // 1: product.Product.ProductDetail:input_type -> product.ProductDetailRequest
	2, // 2: product.Product.EsAddProduct:input_type -> product.EsAddProductRequest
	4, // 3: product.Product.GetAllProduct:input_type -> product.GetAllProductRequest
	1, // 4: product.Product.ProductDetail:output_type -> product.ProductDetailResponse
	3, // 5: product.Product.EsAddProduct:output_type -> product.EsAddProductResponse
	5, // 6: product.Product.GetAllProduct:output_type -> product.GetAllProductResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
=======
var file_product_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_product_proto_goTypes = []any{
	(*ProductDetailReq)(nil),  // 0: product.ProductDetailReq
	(*ProductDetailResp)(nil), // 1: product.ProductDetailResp
}
var file_product_proto_depIdxs = []int32{
	0, // 0: product.Product.ProductDetail:input_type -> product.ProductDetailReq
	1, // 1: product.Product.ProductDetail:output_type -> product.ProductDetailResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
}

func init() { file_product_proto_init() }
func file_product_proto_init() {
	if File_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_proto_msgTypes[0].Exporter = func(v any, i int) any {
<<<<<<< HEAD
			switch v := v.(*ProductDetailRequest); i {
=======
			switch v := v.(*ProductDetailReq); i {
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
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
		file_product_proto_msgTypes[1].Exporter = func(v any, i int) any {
<<<<<<< HEAD
			switch v := v.(*ProductDetailResponse); i {
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
		file_product_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*EsAddProductRequest); i {
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
		file_product_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*EsAddProductResponse); i {
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
		file_product_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllProductRequest); i {
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
		file_product_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllProductResponse); i {
=======
			switch v := v.(*ProductDetailResp); i {
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
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
			RawDescriptor: file_product_proto_rawDesc,
			NumEnums:      0,
<<<<<<< HEAD
			NumMessages:   6,
=======
			NumMessages:   2,
>>>>>>> 48fbf0b70a09b7d084435cde89b1029c8e82fc3f
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_proto_goTypes,
		DependencyIndexes: file_product_proto_depIdxs,
		MessageInfos:      file_product_proto_msgTypes,
	}.Build()
	File_product_proto = out.File
	file_product_proto_rawDesc = nil
	file_product_proto_goTypes = nil
	file_product_proto_depIdxs = nil
}
