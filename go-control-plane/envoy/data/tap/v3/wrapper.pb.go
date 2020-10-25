// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/data/tap/v3/wrapper.proto

package envoy_data_tap_v3

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Wrapper for all fully buffered and streamed tap traces that Envoy emits. This is required for
// sending traces over gRPC APIs or more easily persisting binary messages to files.
type TraceWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Trace:
	//	*TraceWrapper_HttpBufferedTrace
	//	*TraceWrapper_HttpStreamedTraceSegment
	//	*TraceWrapper_SocketBufferedTrace
	//	*TraceWrapper_SocketStreamedTraceSegment
	Trace isTraceWrapper_Trace `protobuf_oneof:"trace"`
}

func (x *TraceWrapper) Reset() {
	*x = TraceWrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_tap_v3_wrapper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraceWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceWrapper) ProtoMessage() {}

func (x *TraceWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_tap_v3_wrapper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceWrapper.ProtoReflect.Descriptor instead.
func (*TraceWrapper) Descriptor() ([]byte, []int) {
	return file_envoy_data_tap_v3_wrapper_proto_rawDescGZIP(), []int{0}
}

func (m *TraceWrapper) GetTrace() isTraceWrapper_Trace {
	if m != nil {
		return m.Trace
	}
	return nil
}

func (x *TraceWrapper) GetHttpBufferedTrace() *HttpBufferedTrace {
	if x, ok := x.GetTrace().(*TraceWrapper_HttpBufferedTrace); ok {
		return x.HttpBufferedTrace
	}
	return nil
}

func (x *TraceWrapper) GetHttpStreamedTraceSegment() *HttpStreamedTraceSegment {
	if x, ok := x.GetTrace().(*TraceWrapper_HttpStreamedTraceSegment); ok {
		return x.HttpStreamedTraceSegment
	}
	return nil
}

func (x *TraceWrapper) GetSocketBufferedTrace() *SocketBufferedTrace {
	if x, ok := x.GetTrace().(*TraceWrapper_SocketBufferedTrace); ok {
		return x.SocketBufferedTrace
	}
	return nil
}

func (x *TraceWrapper) GetSocketStreamedTraceSegment() *SocketStreamedTraceSegment {
	if x, ok := x.GetTrace().(*TraceWrapper_SocketStreamedTraceSegment); ok {
		return x.SocketStreamedTraceSegment
	}
	return nil
}

type isTraceWrapper_Trace interface {
	isTraceWrapper_Trace()
}

type TraceWrapper_HttpBufferedTrace struct {
	// An HTTP buffered tap trace.
	HttpBufferedTrace *HttpBufferedTrace `protobuf:"bytes,1,opt,name=http_buffered_trace,json=httpBufferedTrace,proto3,oneof"`
}

type TraceWrapper_HttpStreamedTraceSegment struct {
	// An HTTP streamed tap trace segment.
	HttpStreamedTraceSegment *HttpStreamedTraceSegment `protobuf:"bytes,2,opt,name=http_streamed_trace_segment,json=httpStreamedTraceSegment,proto3,oneof"`
}

type TraceWrapper_SocketBufferedTrace struct {
	// A socket buffered tap trace.
	SocketBufferedTrace *SocketBufferedTrace `protobuf:"bytes,3,opt,name=socket_buffered_trace,json=socketBufferedTrace,proto3,oneof"`
}

type TraceWrapper_SocketStreamedTraceSegment struct {
	// A socket streamed tap trace segment.
	SocketStreamedTraceSegment *SocketStreamedTraceSegment `protobuf:"bytes,4,opt,name=socket_streamed_trace_segment,json=socketStreamedTraceSegment,proto3,oneof"`
}

func (*TraceWrapper_HttpBufferedTrace) isTraceWrapper_Trace() {}

func (*TraceWrapper_HttpStreamedTraceSegment) isTraceWrapper_Trace() {}

func (*TraceWrapper_SocketBufferedTrace) isTraceWrapper_Trace() {}

func (*TraceWrapper_SocketStreamedTraceSegment) isTraceWrapper_Trace() {}

var File_envoy_data_tap_v3_wrapper_proto protoreflect.FileDescriptor

var file_envoy_data_tap_v3_wrapper_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x74, 0x61, 0x70,
	0x2f, 0x76, 0x33, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x61,
	0x70, 0x2e, 0x76, 0x33, 0x1a, 0x1c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x2f, 0x74, 0x61, 0x70, 0x2f, 0x76, 0x33, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x21, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x74,
	0x61, 0x70, 0x2f, 0x76, 0x33, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe0, 0x03, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x63, 0x65, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x12, 0x56, 0x0a, 0x13, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72,
	0x65, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x61, 0x70, 0x2e,
	0x76, 0x33, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x65, 0x64, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x48, 0x00, 0x52, 0x11, 0x68, 0x74, 0x74, 0x70, 0x42, 0x75, 0x66, 0x66,
	0x65, 0x72, 0x65, 0x64, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x6c, 0x0a, 0x1b, 0x68, 0x74, 0x74,
	0x70, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x61, 0x70, 0x2e,
	0x76, 0x33, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x64, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x18, 0x68,
	0x74, 0x74, 0x70, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x64, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x5c, 0x0a, 0x15, 0x73, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x65, 0x64, 0x54, 0x72, 0x61, 0x63, 0x65, 0x48, 0x00,
	0x52, 0x13, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x65, 0x64,
	0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x72, 0x0a, 0x1d, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x73,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x61, 0x70, 0x2e, 0x76, 0x33,
	0x2e, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x64, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x1a, 0x73,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x64, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x2a, 0x9a, 0xc5, 0x88, 0x1e, 0x25,
	0x0a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x61, 0x70,
	0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x57, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x42, 0x0c, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x12, 0x03,
	0xf8, 0x42, 0x01, 0x42, 0x39, 0x0a, 0x1f, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x74, 0x61, 0x70, 0x2e, 0x76, 0x33, 0x42, 0x0c, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_data_tap_v3_wrapper_proto_rawDescOnce sync.Once
	file_envoy_data_tap_v3_wrapper_proto_rawDescData = file_envoy_data_tap_v3_wrapper_proto_rawDesc
)

func file_envoy_data_tap_v3_wrapper_proto_rawDescGZIP() []byte {
	file_envoy_data_tap_v3_wrapper_proto_rawDescOnce.Do(func() {
		file_envoy_data_tap_v3_wrapper_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_data_tap_v3_wrapper_proto_rawDescData)
	})
	return file_envoy_data_tap_v3_wrapper_proto_rawDescData
}

var file_envoy_data_tap_v3_wrapper_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_data_tap_v3_wrapper_proto_goTypes = []interface{}{
	(*TraceWrapper)(nil),               // 0: envoy.data.tap.v3.TraceWrapper
	(*HttpBufferedTrace)(nil),          // 1: envoy.data.tap.v3.HttpBufferedTrace
	(*HttpStreamedTraceSegment)(nil),   // 2: envoy.data.tap.v3.HttpStreamedTraceSegment
	(*SocketBufferedTrace)(nil),        // 3: envoy.data.tap.v3.SocketBufferedTrace
	(*SocketStreamedTraceSegment)(nil), // 4: envoy.data.tap.v3.SocketStreamedTraceSegment
}
var file_envoy_data_tap_v3_wrapper_proto_depIdxs = []int32{
	1, // 0: envoy.data.tap.v3.TraceWrapper.http_buffered_trace:type_name -> envoy.data.tap.v3.HttpBufferedTrace
	2, // 1: envoy.data.tap.v3.TraceWrapper.http_streamed_trace_segment:type_name -> envoy.data.tap.v3.HttpStreamedTraceSegment
	3, // 2: envoy.data.tap.v3.TraceWrapper.socket_buffered_trace:type_name -> envoy.data.tap.v3.SocketBufferedTrace
	4, // 3: envoy.data.tap.v3.TraceWrapper.socket_streamed_trace_segment:type_name -> envoy.data.tap.v3.SocketStreamedTraceSegment
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_envoy_data_tap_v3_wrapper_proto_init() }
func file_envoy_data_tap_v3_wrapper_proto_init() {
	if File_envoy_data_tap_v3_wrapper_proto != nil {
		return
	}
	file_envoy_data_tap_v3_http_proto_init()
	file_envoy_data_tap_v3_transport_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_data_tap_v3_wrapper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TraceWrapper); i {
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
	file_envoy_data_tap_v3_wrapper_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TraceWrapper_HttpBufferedTrace)(nil),
		(*TraceWrapper_HttpStreamedTraceSegment)(nil),
		(*TraceWrapper_SocketBufferedTrace)(nil),
		(*TraceWrapper_SocketStreamedTraceSegment)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_data_tap_v3_wrapper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_data_tap_v3_wrapper_proto_goTypes,
		DependencyIndexes: file_envoy_data_tap_v3_wrapper_proto_depIdxs,
		MessageInfos:      file_envoy_data_tap_v3_wrapper_proto_msgTypes,
	}.Build()
	File_envoy_data_tap_v3_wrapper_proto = out.File
	file_envoy_data_tap_v3_wrapper_proto_rawDesc = nil
	file_envoy_data_tap_v3_wrapper_proto_goTypes = nil
	file_envoy_data_tap_v3_wrapper_proto_depIdxs = nil
}
