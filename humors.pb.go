// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: humors.proto

package humors

import (
	proto "github.com/golang/protobuf/proto"
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

type ErrorCode int32

const (
	ErrorCode_SUCCESS      ErrorCode = 0
	ErrorCode_DECODEERR    ErrorCode = 1
	ErrorCode_ENCODEERR    ErrorCode = 2
	ErrorCode_NOFUNCERR    ErrorCode = 3
	ErrorCode_NOSERVANTERR ErrorCode = 4
	ErrorCode_CONNECTERR   ErrorCode = 5
	ErrorCode_TIMEOUT      ErrorCode = 6
	ErrorCode_SERVICEERR   ErrorCode = 7
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0: "SUCCESS",
		1: "DECODEERR",
		2: "ENCODEERR",
		3: "NOFUNCERR",
		4: "NOSERVANTERR",
		5: "CONNECTERR",
		6: "TIMEOUT",
		7: "SERVICEERR",
	}
	ErrorCode_value = map[string]int32{
		"SUCCESS":      0,
		"DECODEERR":    1,
		"ENCODEERR":    2,
		"NOFUNCERR":    3,
		"NOSERVANTERR": 4,
		"CONNECTERR":   5,
		"TIMEOUT":      6,
		"SERVICEERR":   7,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_humors_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_humors_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_humors_proto_rawDescGZIP(), []int{0}
}

type RequestPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqID    int32  `protobuf:"varint,1,opt,name=reqID,proto3" json:"reqID,omitempty"`
	Action   string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Timeout  int32  `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Payload  []byte `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	ResTopic string `protobuf:"bytes,5,opt,name=resTopic,proto3" json:"resTopic,omitempty"`
}

func (x *RequestPacket) Reset() {
	*x = RequestPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_humors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestPacket) ProtoMessage() {}

func (x *RequestPacket) ProtoReflect() protoreflect.Message {
	mi := &file_humors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestPacket.ProtoReflect.Descriptor instead.
func (*RequestPacket) Descriptor() ([]byte, []int) {
	return file_humors_proto_rawDescGZIP(), []int{0}
}

func (x *RequestPacket) GetReqID() int32 {
	if x != nil {
		return x.ReqID
	}
	return 0
}

func (x *RequestPacket) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *RequestPacket) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *RequestPacket) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *RequestPacket) GetResTopic() string {
	if x != nil {
		return x.ResTopic
	}
	return ""
}

type ResponsePacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqID   int32     `protobuf:"varint,1,opt,name=reqID,proto3" json:"reqID,omitempty"`
	Code    ErrorCode `protobuf:"varint,2,opt,name=code,proto3,enum=humors.ErrorCode" json:"code,omitempty"`
	Payload []byte    `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *ResponsePacket) Reset() {
	*x = ResponsePacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_humors_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponsePacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponsePacket) ProtoMessage() {}

func (x *ResponsePacket) ProtoReflect() protoreflect.Message {
	mi := &file_humors_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponsePacket.ProtoReflect.Descriptor instead.
func (*ResponsePacket) Descriptor() ([]byte, []int) {
	return file_humors_proto_rawDescGZIP(), []int{1}
}

func (x *ResponsePacket) GetReqID() int32 {
	if x != nil {
		return x.ReqID
	}
	return 0
}

func (x *ResponsePacket) GetCode() ErrorCode {
	if x != nil {
		return x.Code
	}
	return ErrorCode_SUCCESS
}

func (x *ResponsePacket) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_humors_proto protoreflect.FileDescriptor

var file_humors_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x68, 0x75, 0x6d, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x68, 0x75, 0x6d, 0x6f, 0x72, 0x73, 0x22, 0x8d, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x71, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x65, 0x71, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x22, 0x67, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x71, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x65, 0x71, 0x49, 0x44, 0x12, 0x25,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x68,
	0x75, 0x6d, 0x6f, 0x72, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2a,
	0x84, 0x01, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45,
	0x43, 0x4f, 0x44, 0x45, 0x45, 0x52, 0x52, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x4e, 0x43,
	0x4f, 0x44, 0x45, 0x45, 0x52, 0x52, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x46, 0x55,
	0x4e, 0x43, 0x45, 0x52, 0x52, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x4f, 0x53, 0x45, 0x52,
	0x56, 0x41, 0x4e, 0x54, 0x45, 0x52, 0x52, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4e,
	0x4e, 0x45, 0x43, 0x54, 0x45, 0x52, 0x52, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x49, 0x4d,
	0x45, 0x4f, 0x55, 0x54, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43,
	0x45, 0x45, 0x52, 0x52, 0x10, 0x07, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x79, 0x61, 0x6f, 0x2f,
	0x68, 0x75, 0x6d, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_humors_proto_rawDescOnce sync.Once
	file_humors_proto_rawDescData = file_humors_proto_rawDesc
)

func file_humors_proto_rawDescGZIP() []byte {
	file_humors_proto_rawDescOnce.Do(func() {
		file_humors_proto_rawDescData = protoimpl.X.CompressGZIP(file_humors_proto_rawDescData)
	})
	return file_humors_proto_rawDescData
}

var file_humors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_humors_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_humors_proto_goTypes = []interface{}{
	(ErrorCode)(0),         // 0: humors.ErrorCode
	(*RequestPacket)(nil),  // 1: humors.RequestPacket
	(*ResponsePacket)(nil), // 2: humors.ResponsePacket
}
var file_humors_proto_depIdxs = []int32{
	0, // 0: humors.ResponsePacket.code:type_name -> humors.ErrorCode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_humors_proto_init() }
func file_humors_proto_init() {
	if File_humors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_humors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestPacket); i {
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
		file_humors_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponsePacket); i {
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
			RawDescriptor: file_humors_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_humors_proto_goTypes,
		DependencyIndexes: file_humors_proto_depIdxs,
		EnumInfos:         file_humors_proto_enumTypes,
		MessageInfos:      file_humors_proto_msgTypes,
	}.Build()
	File_humors_proto = out.File
	file_humors_proto_rawDesc = nil
	file_humors_proto_goTypes = nil
	file_humors_proto_depIdxs = nil
}
