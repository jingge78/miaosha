// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.19.4
// source: coupon.proto

package coupon

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

type AddCouponRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string  `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Integral    uint64  `protobuf:"varint,2,opt,name=Integral,proto3" json:"Integral,omitempty"`
	CouponPrice float32 `protobuf:"fixed32,3,opt,name=CouponPrice,proto3" json:"CouponPrice,omitempty"`
	UseMinPrice float32 `protobuf:"fixed32,4,opt,name=UseMinPrice,proto3" json:"UseMinPrice,omitempty"`
	CouponTime  uint64  `protobuf:"varint,5,opt,name=CouponTime,proto3" json:"CouponTime,omitempty"`
	ProductId   string  `protobuf:"bytes,6,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	CategoryId  uint64  `protobuf:"varint,7,opt,name=CategoryId,proto3" json:"CategoryId,omitempty"`
	Type        uint64  `protobuf:"varint,8,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *AddCouponRequest) Reset() {
	*x = AddCouponRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coupon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCouponRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCouponRequest) ProtoMessage() {}

func (x *AddCouponRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coupon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCouponRequest.ProtoReflect.Descriptor instead.
func (*AddCouponRequest) Descriptor() ([]byte, []int) {
	return file_coupon_proto_rawDescGZIP(), []int{0}
}

func (x *AddCouponRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddCouponRequest) GetIntegral() uint64 {
	if x != nil {
		return x.Integral
	}
	return 0
}

func (x *AddCouponRequest) GetCouponPrice() float32 {
	if x != nil {
		return x.CouponPrice
	}
	return 0
}

func (x *AddCouponRequest) GetUseMinPrice() float32 {
	if x != nil {
		return x.UseMinPrice
	}
	return 0
}

func (x *AddCouponRequest) GetCouponTime() uint64 {
	if x != nil {
		return x.CouponTime
	}
	return 0
}

func (x *AddCouponRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *AddCouponRequest) GetCategoryId() uint64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *AddCouponRequest) GetType() uint64 {
	if x != nil {
		return x.Type
	}
	return 0
}

type AddCouponResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CouponId uint64 `protobuf:"varint,1,opt,name=CouponId,proto3" json:"CouponId,omitempty"`
}

func (x *AddCouponResponse) Reset() {
	*x = AddCouponResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coupon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCouponResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCouponResponse) ProtoMessage() {}

func (x *AddCouponResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coupon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCouponResponse.ProtoReflect.Descriptor instead.
func (*AddCouponResponse) Descriptor() ([]byte, []int) {
	return file_coupon_proto_rawDescGZIP(), []int{1}
}

func (x *AddCouponResponse) GetCouponId() uint64 {
	if x != nil {
		return x.CouponId
	}
	return 0
}

var File_coupon_proto protoreflect.FileDescriptor

var file_coupon_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x22, 0xfa, 0x01, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x43, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x6c, 0x12, 0x20, 0x0a,
	0x0b, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0b, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x4d, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x55, 0x73, 0x65, 0x4d, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x2f, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x43, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x49, 0x64, 0x32, 0x4a, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x12, 0x40,
	0x0a, 0x09, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x63, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x41,
	0x64, 0x64, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coupon_proto_rawDescOnce sync.Once
	file_coupon_proto_rawDescData = file_coupon_proto_rawDesc
)

func file_coupon_proto_rawDescGZIP() []byte {
	file_coupon_proto_rawDescOnce.Do(func() {
		file_coupon_proto_rawDescData = protoimpl.X.CompressGZIP(file_coupon_proto_rawDescData)
	})
	return file_coupon_proto_rawDescData
}

var file_coupon_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_coupon_proto_goTypes = []any{
	(*AddCouponRequest)(nil),  // 0: coupon.AddCouponRequest
	(*AddCouponResponse)(nil), // 1: coupon.AddCouponResponse
}
var file_coupon_proto_depIdxs = []int32{
	0, // 0: coupon.Coupon.AddCoupon:input_type -> coupon.AddCouponRequest
	1, // 1: coupon.Coupon.AddCoupon:output_type -> coupon.AddCouponResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_coupon_proto_init() }
func file_coupon_proto_init() {
	if File_coupon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coupon_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AddCouponRequest); i {
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
		file_coupon_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AddCouponResponse); i {
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
			RawDescriptor: file_coupon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coupon_proto_goTypes,
		DependencyIndexes: file_coupon_proto_depIdxs,
		MessageInfos:      file_coupon_proto_msgTypes,
	}.Build()
	File_coupon_proto = out.File
	file_coupon_proto_rawDesc = nil
	file_coupon_proto_goTypes = nil
	file_coupon_proto_depIdxs = nil
}
