// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: twitch/v1/common.proto

package twitchv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_twitch_v1_common_proto protoreflect.FileDescriptor

var file_twitch_v1_common_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x77, 0x69, 0x74, 0x63, 0x68,
	0x2e, 0x76, 0x31, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6c, 0x6c, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x64, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x74, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x77, 0x69, 0x74,
	0x63, 0x68, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_twitch_v1_common_proto_goTypes = []interface{}{}
var file_twitch_v1_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_twitch_v1_common_proto_init() }
func file_twitch_v1_common_proto_init() {
	if File_twitch_v1_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_twitch_v1_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_twitch_v1_common_proto_goTypes,
		DependencyIndexes: file_twitch_v1_common_proto_depIdxs,
	}.Build()
	File_twitch_v1_common_proto = out.File
	file_twitch_v1_common_proto_rawDesc = nil
	file_twitch_v1_common_proto_goTypes = nil
	file_twitch_v1_common_proto_depIdxs = nil
}
