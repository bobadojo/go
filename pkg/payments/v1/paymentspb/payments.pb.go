// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: bobadojo/payments/v1/payments.proto

package paymentspb

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

var File_bobadojo_payments_v1_payments_proto protoreflect.FileDescriptor

var file_bobadojo_payments_v1_payments_proto_rawDesc = []byte{
	0x0a, 0x23, 0x62, 0x6f, 0x62, 0x61, 0x64, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x62, 0x6f, 0x62, 0x61, 0x64, 0x6f, 0x6a, 0x6f, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x69, 0x0a, 0x18, 0x63,
	0x6f, 0x6d, 0x2e, 0x62, 0x6f, 0x62, 0x61, 0x64, 0x6f, 0x6a, 0x6f, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6f, 0x62, 0x61, 0x64, 0x6f, 0x6a, 0x6f, 0x2f, 0x67, 0x6f,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x3b, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_bobadojo_payments_v1_payments_proto_goTypes = []any{}
var file_bobadojo_payments_v1_payments_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bobadojo_payments_v1_payments_proto_init() }
func file_bobadojo_payments_v1_payments_proto_init() {
	if File_bobadojo_payments_v1_payments_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bobadojo_payments_v1_payments_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bobadojo_payments_v1_payments_proto_goTypes,
		DependencyIndexes: file_bobadojo_payments_v1_payments_proto_depIdxs,
	}.Build()
	File_bobadojo_payments_v1_payments_proto = out.File
	file_bobadojo_payments_v1_payments_proto_rawDesc = nil
	file_bobadojo_payments_v1_payments_proto_goTypes = nil
	file_bobadojo_payments_v1_payments_proto_depIdxs = nil
}
