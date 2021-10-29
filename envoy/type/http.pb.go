// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.0
// source: envoy/type/http.proto

package envoy_type

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
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

type CodecClientType int32

const (
	CodecClientType_HTTP1 CodecClientType = 0
	CodecClientType_HTTP2 CodecClientType = 1
	// [#not-implemented-hide:] QUIC implementation is not production ready yet. Use this enum with
	// caution to prevent accidental execution of QUIC code. I.e. `!= HTTP2` is no longer sufficient
	// to distinguish HTTP1 and HTTP2 traffic.
	CodecClientType_HTTP3 CodecClientType = 2
)

// Enum value maps for CodecClientType.
var (
	CodecClientType_name = map[int32]string{
		0: "HTTP1",
		1: "HTTP2",
		2: "HTTP3",
	}
	CodecClientType_value = map[string]int32{
		"HTTP1": 0,
		"HTTP2": 1,
		"HTTP3": 2,
	}
)

func (x CodecClientType) Enum() *CodecClientType {
	p := new(CodecClientType)
	*p = x
	return p
}

func (x CodecClientType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CodecClientType) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_type_http_proto_enumTypes[0].Descriptor()
}

func (CodecClientType) Type() protoreflect.EnumType {
	return &file_envoy_type_http_proto_enumTypes[0]
}

func (x CodecClientType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CodecClientType.Descriptor instead.
func (CodecClientType) EnumDescriptor() ([]byte, []int) {
	return file_envoy_type_http_proto_rawDescGZIP(), []int{0}
}

var File_envoy_type_http_proto protoreflect.FileDescriptor

var file_envoy_type_http_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2a, 0x32, 0x0a, 0x0f, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x54, 0x54, 0x50, 0x31, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x48, 0x54, 0x54, 0x50, 0x32, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x48,
	0x54, 0x54, 0x50, 0x33, 0x10, 0x02, 0x42, 0x2f, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x42, 0x09, 0x48, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba,
	0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_type_http_proto_rawDescOnce sync.Once
	file_envoy_type_http_proto_rawDescData = file_envoy_type_http_proto_rawDesc
)

func file_envoy_type_http_proto_rawDescGZIP() []byte {
	file_envoy_type_http_proto_rawDescOnce.Do(func() {
		file_envoy_type_http_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_type_http_proto_rawDescData)
	})
	return file_envoy_type_http_proto_rawDescData
}

var file_envoy_type_http_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_type_http_proto_goTypes = []interface{}{
	(CodecClientType)(0), // 0: envoy.type.CodecClientType
}
var file_envoy_type_http_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_envoy_type_http_proto_init() }
func file_envoy_type_http_proto_init() {
	if File_envoy_type_http_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_type_http_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_type_http_proto_goTypes,
		DependencyIndexes: file_envoy_type_http_proto_depIdxs,
		EnumInfos:         file_envoy_type_http_proto_enumTypes,
	}.Build()
	File_envoy_type_http_proto = out.File
	file_envoy_type_http_proto_rawDesc = nil
	file_envoy_type_http_proto_goTypes = nil
	file_envoy_type_http_proto_depIdxs = nil
}
