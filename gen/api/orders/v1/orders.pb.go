// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: api/orders/v1/orders.proto

package ordersv1

import (
	common "github.com/AdamShannag/shopy-proto/gen/resource/common"
	v1 "github.com/AdamShannag/shopy-proto/gen/resource/order/v1"
	_ "github.com/AdamShannag/shopy-proto/gen/third_party/google/api"
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

var File_api_orders_v1_orders_proto protoreflect.FileDescriptor

var file_api_orders_v1_orders_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70,
	0x69, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x28, 0x74, 0x68, 0x69,
	0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0x74, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x63, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x29,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x57, 0x69, 0x74, 0x68, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0xb5, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x64, 0x61, 0x6d, 0x53, 0x68, 0x61,
	0x6e, 0x6e, 0x61, 0x67, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x79, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x4f,
	0x58, 0xaa, 0x02, 0x0d, 0x41, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x0d, 0x41, 0x70, 0x69, 0x5c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x19, 0x41, 0x70, 0x69, 0x5c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f,
	0x41, 0x70, 0x69, 0x3a, 0x3a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_orders_v1_orders_proto_goTypes = []any{
	(*v1.PayloadWithSingleOrder)(nil), // 0: resource.order.v1.PayloadWithSingleOrder
	(*common.Empty)(nil),              // 1: resource.common.Empty
}
var file_api_orders_v1_orders_proto_depIdxs = []int32{
	0, // 0: api.orders.v1.OrdersService.AddOrder:input_type -> resource.order.v1.PayloadWithSingleOrder
	1, // 1: api.orders.v1.OrdersService.AddOrder:output_type -> resource.common.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_orders_v1_orders_proto_init() }
func file_api_orders_v1_orders_proto_init() {
	if File_api_orders_v1_orders_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_orders_v1_orders_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_orders_v1_orders_proto_goTypes,
		DependencyIndexes: file_api_orders_v1_orders_proto_depIdxs,
	}.Build()
	File_api_orders_v1_orders_proto = out.File
	file_api_orders_v1_orders_proto_rawDesc = nil
	file_api_orders_v1_orders_proto_goTypes = nil
	file_api_orders_v1_orders_proto_depIdxs = nil
}
