// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.4
// source: staff.proto

package staff

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StaffRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uid           int32                  `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Avatar        string                 `protobuf:"bytes,2,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	StoreId       int32                  `protobuf:"varint,3,opt,name=StoreId,proto3" json:"StoreId,omitempty"`
	StaffName     int32                  `protobuf:"varint,4,opt,name=StaffName,proto3" json:"StaffName,omitempty"`
	Phone         string                 `protobuf:"bytes,5,opt,name=Phone,proto3" json:"Phone,omitempty"`
	VerifyStatus  int32                  `protobuf:"varint,6,opt,name=VerifyStatus,proto3" json:"VerifyStatus,omitempty"`
	Status        int32                  `protobuf:"varint,7,opt,name=Status,proto3" json:"Status,omitempty"`
	AddTime       int32                  `protobuf:"varint,8,opt,name=AddTime,proto3" json:"AddTime,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StaffRequest) Reset() {
	*x = StaffRequest{}
	mi := &file_staff_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StaffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaffRequest) ProtoMessage() {}

func (x *StaffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaffRequest.ProtoReflect.Descriptor instead.
func (*StaffRequest) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{0}
}

func (x *StaffRequest) GetUid() int32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *StaffRequest) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *StaffRequest) GetStoreId() int32 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *StaffRequest) GetStaffName() int32 {
	if x != nil {
		return x.StaffName
	}
	return 0
}

func (x *StaffRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *StaffRequest) GetVerifyStatus() int32 {
	if x != nil {
		return x.VerifyStatus
	}
	return 0
}

func (x *StaffRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *StaffRequest) GetAddTime() int32 {
	if x != nil {
		return x.AddTime
	}
	return 0
}

type StaffResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        string                 `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StaffResponse) Reset() {
	*x = StaffResponse{}
	mi := &file_staff_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StaffResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaffResponse) ProtoMessage() {}

func (x *StaffResponse) ProtoReflect() protoreflect.Message {
	mi := &file_staff_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaffResponse.ProtoReflect.Descriptor instead.
func (*StaffResponse) Descriptor() ([]byte, []int) {
	return file_staff_proto_rawDescGZIP(), []int{1}
}

func (x *StaffResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_staff_proto protoreflect.FileDescriptor

var file_staff_proto_rawDesc = string([]byte{
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0xdc, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x66,
	0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x55, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x53, 0x74, 0x61, 0x66, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x53, 0x74, 0x61, 0x66, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x41, 0x64, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x41,
	0x64, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x27, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32,
	0x45, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x66, 0x66, 0x12, 0x3c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x66,
	0x66, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_staff_proto_rawDescOnce sync.Once
	file_staff_proto_rawDescData []byte
)

func file_staff_proto_rawDescGZIP() []byte {
	file_staff_proto_rawDescOnce.Do(func() {
		file_staff_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_staff_proto_rawDesc), len(file_staff_proto_rawDesc)))
	})
	return file_staff_proto_rawDescData
}

var file_staff_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_staff_proto_goTypes = []any{
	(*StaffRequest)(nil),  // 0: template.StaffRequest
	(*StaffResponse)(nil), // 1: template.StaffResponse
}
var file_staff_proto_depIdxs = []int32{
	0, // 0: template.Staff.StaffSend:input_type -> template.StaffRequest
	1, // 1: template.Staff.StaffSend:output_type -> template.StaffResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_staff_proto_init() }
func file_staff_proto_init() {
	if File_staff_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_staff_proto_rawDesc), len(file_staff_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_staff_proto_goTypes,
		DependencyIndexes: file_staff_proto_depIdxs,
		MessageInfos:      file_staff_proto_msgTypes,
	}.Build()
	File_staff_proto = out.File
	file_staff_proto_goTypes = nil
	file_staff_proto_depIdxs = nil
}
