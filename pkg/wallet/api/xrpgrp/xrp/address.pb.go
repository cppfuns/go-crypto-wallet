// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: address.proto

package xrp

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResponseGenerateAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XAddress       string `protobuf:"bytes,1,opt,name=xAddress,proto3" json:"xAddress,omitempty"`
	ClassicAddress string `protobuf:"bytes,2,opt,name=classicAddress,proto3" json:"classicAddress,omitempty"`
	Address        string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Secret         string `protobuf:"bytes,4,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *ResponseGenerateAddress) Reset() {
	*x = ResponseGenerateAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseGenerateAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseGenerateAddress) ProtoMessage() {}

func (x *ResponseGenerateAddress) ProtoReflect() protoreflect.Message {
	mi := &file_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseGenerateAddress.ProtoReflect.Descriptor instead.
func (*ResponseGenerateAddress) Descriptor() ([]byte, []int) {
	return file_address_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseGenerateAddress) GetXAddress() string {
	if x != nil {
		return x.XAddress
	}
	return ""
}

func (x *ResponseGenerateAddress) GetClassicAddress() string {
	if x != nil {
		return x.ClassicAddress
	}
	return ""
}

func (x *ResponseGenerateAddress) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ResponseGenerateAddress) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type ResponseGenerateXAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XAddress string `protobuf:"bytes,1,opt,name=xAddress,proto3" json:"xAddress,omitempty"`
	Secret   string `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *ResponseGenerateXAddress) Reset() {
	*x = ResponseGenerateXAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_address_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseGenerateXAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseGenerateXAddress) ProtoMessage() {}

func (x *ResponseGenerateXAddress) ProtoReflect() protoreflect.Message {
	mi := &file_address_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseGenerateXAddress.ProtoReflect.Descriptor instead.
func (*ResponseGenerateXAddress) Descriptor() ([]byte, []int) {
	return file_address_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseGenerateXAddress) GetXAddress() string {
	if x != nil {
		return x.XAddress
	}
	return ""
}

func (x *ResponseGenerateXAddress) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type RequestIsValidAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *RequestIsValidAddress) Reset() {
	*x = RequestIsValidAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_address_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestIsValidAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestIsValidAddress) ProtoMessage() {}

func (x *RequestIsValidAddress) ProtoReflect() protoreflect.Message {
	mi := &file_address_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestIsValidAddress.ProtoReflect.Descriptor instead.
func (*RequestIsValidAddress) Descriptor() ([]byte, []int) {
	return file_address_proto_rawDescGZIP(), []int{2}
}

func (x *RequestIsValidAddress) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type ResponseIsValidAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsValid bool `protobuf:"varint,1,opt,name=isValid,proto3" json:"isValid,omitempty"`
}

func (x *ResponseIsValidAddress) Reset() {
	*x = ResponseIsValidAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_address_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseIsValidAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseIsValidAddress) ProtoMessage() {}

func (x *ResponseIsValidAddress) ProtoReflect() protoreflect.Message {
	mi := &file_address_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseIsValidAddress.ProtoReflect.Descriptor instead.
func (*ResponseIsValidAddress) Descriptor() ([]byte, []int) {
	return file_address_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseIsValidAddress) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

var File_address_proto protoreflect.FileDescriptor

var file_address_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x8f, 0x01, 0x0a, 0x17, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x78,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x78,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x69, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x22, 0x4e, 0x0a, 0x18, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x58, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x78, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x78, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x22, 0x31, 0x0a, 0x15, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x73, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x32, 0x0a, 0x16, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x69, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x69, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x32, 0xaf, 0x02, 0x0a, 0x10, 0x52, 0x69, 0x70,
	0x70, 0x6c, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x41, 0x50, 0x49, 0x12, 0x57, 0x0a,
	0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2a, 0x2e, 0x72, 0x69, 0x70, 0x70, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x58, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x2b, 0x2e, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x58, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0x00, 0x12, 0x67, 0x0a, 0x0e, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x28, 0x2e, 0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x29, 0x2e,
	0x72, 0x69, 0x70, 0x70, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x00, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x69, 0x72, 0x6f, 0x6d, 0x61, 0x69,
	0x6c, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2d, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x78, 0x72, 0x70, 0x67, 0x72, 0x70, 0x2f, 0x78, 0x72, 0x70, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_address_proto_rawDescOnce sync.Once
	file_address_proto_rawDescData = file_address_proto_rawDesc
)

func file_address_proto_rawDescGZIP() []byte {
	file_address_proto_rawDescOnce.Do(func() {
		file_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_address_proto_rawDescData)
	})
	return file_address_proto_rawDescData
}

var file_address_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_address_proto_goTypes = []interface{}{
	(*ResponseGenerateAddress)(nil),  // 0: rippleapi.address.ResponseGenerateAddress
	(*ResponseGenerateXAddress)(nil), // 1: rippleapi.address.ResponseGenerateXAddress
	(*RequestIsValidAddress)(nil),    // 2: rippleapi.address.RequestIsValidAddress
	(*ResponseIsValidAddress)(nil),   // 3: rippleapi.address.ResponseIsValidAddress
	(*emptypb.Empty)(nil),            // 4: google.protobuf.Empty
}
var file_address_proto_depIdxs = []int32{
	4, // 0: rippleapi.address.RippleAddressAPI.GenerateAddress:input_type -> google.protobuf.Empty
	4, // 1: rippleapi.address.RippleAddressAPI.GenerateXAddress:input_type -> google.protobuf.Empty
	2, // 2: rippleapi.address.RippleAddressAPI.IsValidAddress:input_type -> rippleapi.address.RequestIsValidAddress
	0, // 3: rippleapi.address.RippleAddressAPI.GenerateAddress:output_type -> rippleapi.address.ResponseGenerateAddress
	1, // 4: rippleapi.address.RippleAddressAPI.GenerateXAddress:output_type -> rippleapi.address.ResponseGenerateXAddress
	3, // 5: rippleapi.address.RippleAddressAPI.IsValidAddress:output_type -> rippleapi.address.ResponseIsValidAddress
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_address_proto_init() }
func file_address_proto_init() {
	if File_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseGenerateAddress); i {
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
		file_address_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseGenerateXAddress); i {
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
		file_address_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestIsValidAddress); i {
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
		file_address_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseIsValidAddress); i {
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
			RawDescriptor: file_address_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_address_proto_goTypes,
		DependencyIndexes: file_address_proto_depIdxs,
		MessageInfos:      file_address_proto_msgTypes,
	}.Build()
	File_address_proto = out.File
	file_address_proto_rawDesc = nil
	file_address_proto_goTypes = nil
	file_address_proto_depIdxs = nil
}