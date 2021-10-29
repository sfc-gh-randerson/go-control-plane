// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.0
// source: envoy/config/filter/http/gzip/v2/gzip.proto

package envoy_config_filter_http_gzip_v2

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/compressor/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type Gzip_CompressionStrategy int32

const (
	Gzip_DEFAULT  Gzip_CompressionStrategy = 0
	Gzip_FILTERED Gzip_CompressionStrategy = 1
	Gzip_HUFFMAN  Gzip_CompressionStrategy = 2
	Gzip_RLE      Gzip_CompressionStrategy = 3
)

// Enum value maps for Gzip_CompressionStrategy.
var (
	Gzip_CompressionStrategy_name = map[int32]string{
		0: "DEFAULT",
		1: "FILTERED",
		2: "HUFFMAN",
		3: "RLE",
	}
	Gzip_CompressionStrategy_value = map[string]int32{
		"DEFAULT":  0,
		"FILTERED": 1,
		"HUFFMAN":  2,
		"RLE":      3,
	}
)

func (x Gzip_CompressionStrategy) Enum() *Gzip_CompressionStrategy {
	p := new(Gzip_CompressionStrategy)
	*p = x
	return p
}

func (x Gzip_CompressionStrategy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gzip_CompressionStrategy) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_config_filter_http_gzip_v2_gzip_proto_enumTypes[0].Descriptor()
}

func (Gzip_CompressionStrategy) Type() protoreflect.EnumType {
	return &file_envoy_config_filter_http_gzip_v2_gzip_proto_enumTypes[0]
}

func (x Gzip_CompressionStrategy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gzip_CompressionStrategy.Descriptor instead.
func (Gzip_CompressionStrategy) EnumDescriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_gzip_v2_gzip_proto_rawDescGZIP(), []int{0, 0}
}

type Gzip_CompressionLevel_Enum int32

const (
	Gzip_CompressionLevel_DEFAULT Gzip_CompressionLevel_Enum = 0
	Gzip_CompressionLevel_BEST    Gzip_CompressionLevel_Enum = 1
	Gzip_CompressionLevel_SPEED   Gzip_CompressionLevel_Enum = 2
)

// Enum value maps for Gzip_CompressionLevel_Enum.
var (
	Gzip_CompressionLevel_Enum_name = map[int32]string{
		0: "DEFAULT",
		1: "BEST",
		2: "SPEED",
	}
	Gzip_CompressionLevel_Enum_value = map[string]int32{
		"DEFAULT": 0,
		"BEST":    1,
		"SPEED":   2,
	}
)

func (x Gzip_CompressionLevel_Enum) Enum() *Gzip_CompressionLevel_Enum {
	p := new(Gzip_CompressionLevel_Enum)
	*p = x
	return p
}

func (x Gzip_CompressionLevel_Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gzip_CompressionLevel_Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_config_filter_http_gzip_v2_gzip_proto_enumTypes[1].Descriptor()
}

func (Gzip_CompressionLevel_Enum) Type() protoreflect.EnumType {
	return &file_envoy_config_filter_http_gzip_v2_gzip_proto_enumTypes[1]
}

func (x Gzip_CompressionLevel_Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gzip_CompressionLevel_Enum.Descriptor instead.
func (Gzip_CompressionLevel_Enum) EnumDescriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_gzip_v2_gzip_proto_rawDescGZIP(), []int{0, 0, 0}
}

// [#next-free-field: 11]
type Gzip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Value from 1 to 9 that controls the amount of internal memory used by zlib. Higher values
	// use more memory, but are faster and produce better compression results. The default value is 5.
	MemoryLevel *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=memory_level,json=memoryLevel,proto3" json:"memory_level,omitempty"`
	// Minimum response length, in bytes, which will trigger compression. The default value is 30.
	// .. attention:
	//
	//    **This field is deprecated**. Set the `compressor` field instead.
	//
	// Deprecated: Do not use.
	ContentLength *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=content_length,json=contentLength,proto3" json:"content_length,omitempty"`
	// A value used for selecting the zlib compression level. This setting will affect speed and
	// amount of compression applied to the content. "BEST" provides higher compression at the cost of
	// higher latency, "SPEED" provides lower compression with minimum impact on response time.
	// "DEFAULT" provides an optimal result between speed and compression. This field will be set to
	// "DEFAULT" if not specified.
	CompressionLevel Gzip_CompressionLevel_Enum `protobuf:"varint,3,opt,name=compression_level,json=compressionLevel,proto3,enum=envoy.config.filter.http.gzip.v2.Gzip_CompressionLevel_Enum" json:"compression_level,omitempty"`
	// A value used for selecting the zlib compression strategy which is directly related to the
	// characteristics of the content. Most of the time "DEFAULT" will be the best choice, though
	// there are situations which changing this parameter might produce better results. For example,
	// run-length encoding (RLE) is typically used when the content is known for having sequences
	// which same data occurs many consecutive times. For more information about each strategy, please
	// refer to zlib manual.
	CompressionStrategy Gzip_CompressionStrategy `protobuf:"varint,4,opt,name=compression_strategy,json=compressionStrategy,proto3,enum=envoy.config.filter.http.gzip.v2.Gzip_CompressionStrategy" json:"compression_strategy,omitempty"`
	// Set of strings that allows specifying which mime-types yield compression; e.g.,
	// application/json, text/html, etc. When this field is not defined, compression will be applied
	// to the following mime-types: "application/javascript", "application/json",
	// "application/xhtml+xml", "image/svg+xml", "text/css", "text/html", "text/plain", "text/xml".
	// .. attention:
	//
	//    **This field is deprecated**. Set the `compressor` field instead.
	//
	// Deprecated: Do not use.
	ContentType []string `protobuf:"bytes,6,rep,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// If true, disables compression when the response contains an etag header. When it is false, the
	// filter will preserve weak etags and remove the ones that require strong validation.
	// .. attention:
	//
	//    **This field is deprecated**. Set the `compressor` field instead.
	//
	// Deprecated: Do not use.
	DisableOnEtagHeader bool `protobuf:"varint,7,opt,name=disable_on_etag_header,json=disableOnEtagHeader,proto3" json:"disable_on_etag_header,omitempty"`
	// If true, removes accept-encoding from the request headers before dispatching it to the upstream
	// so that responses do not get compressed before reaching the filter.
	// .. attention:
	//
	//    **This field is deprecated**. Set the `compressor` field instead.
	//
	// Deprecated: Do not use.
	RemoveAcceptEncodingHeader bool `protobuf:"varint,8,opt,name=remove_accept_encoding_header,json=removeAcceptEncodingHeader,proto3" json:"remove_accept_encoding_header,omitempty"`
	// Value from 9 to 15 that represents the base two logarithmic of the compressor's window size.
	// Larger window results in better compression at the expense of memory usage. The default is 12
	// which will produce a 4096 bytes window. For more details about this parameter, please refer to
	// zlib manual > deflateInit2.
	WindowBits *wrappers.UInt32Value `protobuf:"bytes,9,opt,name=window_bits,json=windowBits,proto3" json:"window_bits,omitempty"`
	// Set of configuration parameters common for all compression filters. If this field is set then
	// the fields `content_length`, `content_type`, `disable_on_etag_header` and
	// `remove_accept_encoding_header` are ignored.
	Compressor *v2.Compressor `protobuf:"bytes,10,opt,name=compressor,proto3" json:"compressor,omitempty"`
}

func (x *Gzip) Reset() {
	*x = Gzip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_http_gzip_v2_gzip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gzip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gzip) ProtoMessage() {}

func (x *Gzip) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_http_gzip_v2_gzip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gzip.ProtoReflect.Descriptor instead.
func (*Gzip) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_gzip_v2_gzip_proto_rawDescGZIP(), []int{0}
}

func (x *Gzip) GetMemoryLevel() *wrappers.UInt32Value {
	if x != nil {
		return x.MemoryLevel
	}
	return nil
}

// Deprecated: Do not use.
func (x *Gzip) GetContentLength() *wrappers.UInt32Value {
	if x != nil {
		return x.ContentLength
	}
	return nil
}

func (x *Gzip) GetCompressionLevel() Gzip_CompressionLevel_Enum {
	if x != nil {
		return x.CompressionLevel
	}
	return Gzip_CompressionLevel_DEFAULT
}

func (x *Gzip) GetCompressionStrategy() Gzip_CompressionStrategy {
	if x != nil {
		return x.CompressionStrategy
	}
	return Gzip_DEFAULT
}

// Deprecated: Do not use.
func (x *Gzip) GetContentType() []string {
	if x != nil {
		return x.ContentType
	}
	return nil
}

// Deprecated: Do not use.
func (x *Gzip) GetDisableOnEtagHeader() bool {
	if x != nil {
		return x.DisableOnEtagHeader
	}
	return false
}

// Deprecated: Do not use.
func (x *Gzip) GetRemoveAcceptEncodingHeader() bool {
	if x != nil {
		return x.RemoveAcceptEncodingHeader
	}
	return false
}

func (x *Gzip) GetWindowBits() *wrappers.UInt32Value {
	if x != nil {
		return x.WindowBits
	}
	return nil
}

func (x *Gzip) GetCompressor() *v2.Compressor {
	if x != nil {
		return x.Compressor
	}
	return nil
}

type Gzip_CompressionLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Gzip_CompressionLevel) Reset() {
	*x = Gzip_CompressionLevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_http_gzip_v2_gzip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gzip_CompressionLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gzip_CompressionLevel) ProtoMessage() {}

func (x *Gzip_CompressionLevel) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_http_gzip_v2_gzip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gzip_CompressionLevel.ProtoReflect.Descriptor instead.
func (*Gzip_CompressionLevel) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_gzip_v2_gzip_proto_rawDescGZIP(), []int{0, 0}
}

var File_envoy_config_filter_http_gzip_v2_gzip_proto protoreflect.FileDescriptor

var file_envoy_config_filter_http_gzip_v2_gzip_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x67, 0x7a, 0x69, 0x70, 0x2f,
	0x76, 0x32, 0x2f, 0x67, 0x7a, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x67, 0x7a, 0x69, 0x70, 0x2e, 0x76, 0x32, 0x1a,
	0x37, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd4, 0x06, 0x0a, 0x04, 0x47, 0x7a, 0x69, 0x70, 0x12, 0x4a, 0x0a, 0x0c, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x09, 0xfa,
	0x42, 0x06, 0x2a, 0x04, 0x18, 0x09, 0x28, 0x01, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x47, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x02, 0x18, 0x01, 0x52,
	0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x73,
	0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x67, 0x7a, 0x69, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x7a, 0x69,
	0x70, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x77, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x3a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x67, 0x7a, 0x69,
	0x70, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x7a, 0x69, 0x70, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x13, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x25, 0x0a, 0x0c,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x37, 0x0a, 0x16, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6f,
	0x6e, 0x5f, 0x65, 0x74, 0x61, 0x67, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x42, 0x02, 0x18, 0x01, 0x52, 0x13, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x4f, 0x6e, 0x45, 0x74, 0x61, 0x67, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x1d,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x5f, 0x65, 0x6e,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x08, 0x42, 0x02, 0x18, 0x01, 0x52, 0x1a, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x0b, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x62, 0x69,
	0x74, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a, 0x04, 0x18, 0x0f, 0x28,
	0x09, 0x52, 0x0a, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x42, 0x69, 0x74, 0x73, 0x12, 0x52, 0x0a,
	0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x32, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f,
	0x72, 0x1a, 0x3c, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x28, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0b, 0x0a,
	0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x45,
	0x53, 0x54, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x50, 0x45, 0x45, 0x44, 0x10, 0x02, 0x22,
	0x46, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c,
	0x54, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x55, 0x46, 0x46, 0x4d, 0x41, 0x4e, 0x10, 0x02, 0x12, 0x07,
	0x0a, 0x03, 0x52, 0x4c, 0x45, 0x10, 0x03, 0x42, 0x72, 0x0a, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x67, 0x7a, 0x69, 0x70, 0x2e, 0x76, 0x32, 0x42, 0x09, 0x47, 0x7a, 0x69, 0x70, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x27, 0x12, 0x25, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x67, 0x7a, 0x69, 0x70,
	0x2e, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_filter_http_gzip_v2_gzip_proto_rawDescOnce sync.Once
	file_envoy_config_filter_http_gzip_v2_gzip_proto_rawDescData = file_envoy_config_filter_http_gzip_v2_gzip_proto_rawDesc
)

func file_envoy_config_filter_http_gzip_v2_gzip_proto_rawDescGZIP() []byte {
	file_envoy_config_filter_http_gzip_v2_gzip_proto_rawDescOnce.Do(func() {
		file_envoy_config_filter_http_gzip_v2_gzip_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_filter_http_gzip_v2_gzip_proto_rawDescData)
	})
	return file_envoy_config_filter_http_gzip_v2_gzip_proto_rawDescData
}

var file_envoy_config_filter_http_gzip_v2_gzip_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_envoy_config_filter_http_gzip_v2_gzip_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_config_filter_http_gzip_v2_gzip_proto_goTypes = []interface{}{
	(Gzip_CompressionStrategy)(0),   // 0: envoy.config.filter.http.gzip.v2.Gzip.CompressionStrategy
	(Gzip_CompressionLevel_Enum)(0), // 1: envoy.config.filter.http.gzip.v2.Gzip.CompressionLevel.Enum
	(*Gzip)(nil),                    // 2: envoy.config.filter.http.gzip.v2.Gzip
	(*Gzip_CompressionLevel)(nil),   // 3: envoy.config.filter.http.gzip.v2.Gzip.CompressionLevel
	(*wrappers.UInt32Value)(nil),    // 4: google.protobuf.UInt32Value
	(*v2.Compressor)(nil),           // 5: envoy.config.filter.http.compressor.v2.Compressor
}
var file_envoy_config_filter_http_gzip_v2_gzip_proto_depIdxs = []int32{
	4, // 0: envoy.config.filter.http.gzip.v2.Gzip.memory_level:type_name -> google.protobuf.UInt32Value
	4, // 1: envoy.config.filter.http.gzip.v2.Gzip.content_length:type_name -> google.protobuf.UInt32Value
	1, // 2: envoy.config.filter.http.gzip.v2.Gzip.compression_level:type_name -> envoy.config.filter.http.gzip.v2.Gzip.CompressionLevel.Enum
	0, // 3: envoy.config.filter.http.gzip.v2.Gzip.compression_strategy:type_name -> envoy.config.filter.http.gzip.v2.Gzip.CompressionStrategy
	4, // 4: envoy.config.filter.http.gzip.v2.Gzip.window_bits:type_name -> google.protobuf.UInt32Value
	5, // 5: envoy.config.filter.http.gzip.v2.Gzip.compressor:type_name -> envoy.config.filter.http.compressor.v2.Compressor
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_envoy_config_filter_http_gzip_v2_gzip_proto_init() }
func file_envoy_config_filter_http_gzip_v2_gzip_proto_init() {
	if File_envoy_config_filter_http_gzip_v2_gzip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_filter_http_gzip_v2_gzip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gzip); i {
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
		file_envoy_config_filter_http_gzip_v2_gzip_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gzip_CompressionLevel); i {
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
			RawDescriptor: file_envoy_config_filter_http_gzip_v2_gzip_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_filter_http_gzip_v2_gzip_proto_goTypes,
		DependencyIndexes: file_envoy_config_filter_http_gzip_v2_gzip_proto_depIdxs,
		EnumInfos:         file_envoy_config_filter_http_gzip_v2_gzip_proto_enumTypes,
		MessageInfos:      file_envoy_config_filter_http_gzip_v2_gzip_proto_msgTypes,
	}.Build()
	File_envoy_config_filter_http_gzip_v2_gzip_proto = out.File
	file_envoy_config_filter_http_gzip_v2_gzip_proto_rawDesc = nil
	file_envoy_config_filter_http_gzip_v2_gzip_proto_goTypes = nil
	file_envoy_config_filter_http_gzip_v2_gzip_proto_depIdxs = nil
}
