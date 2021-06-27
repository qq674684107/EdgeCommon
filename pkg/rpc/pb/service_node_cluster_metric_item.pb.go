// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: service_node_cluster_metric_item.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// 启用某个指标
type EnableNodeClusterMetricItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeClusterId int64 `protobuf:"varint,1,opt,name=nodeClusterId,proto3" json:"nodeClusterId,omitempty"`
	MetricItemId  int64 `protobuf:"varint,2,opt,name=metricItemId,proto3" json:"metricItemId,omitempty"`
}

func (x *EnableNodeClusterMetricItemRequest) Reset() {
	*x = EnableNodeClusterMetricItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_cluster_metric_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnableNodeClusterMetricItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableNodeClusterMetricItemRequest) ProtoMessage() {}

func (x *EnableNodeClusterMetricItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_cluster_metric_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnableNodeClusterMetricItemRequest.ProtoReflect.Descriptor instead.
func (*EnableNodeClusterMetricItemRequest) Descriptor() ([]byte, []int) {
	return file_service_node_cluster_metric_item_proto_rawDescGZIP(), []int{0}
}

func (x *EnableNodeClusterMetricItemRequest) GetNodeClusterId() int64 {
	if x != nil {
		return x.NodeClusterId
	}
	return 0
}

func (x *EnableNodeClusterMetricItemRequest) GetMetricItemId() int64 {
	if x != nil {
		return x.MetricItemId
	}
	return 0
}

// 禁用某个指标
type DisableNodeClusterMetricItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeClusterId int64 `protobuf:"varint,1,opt,name=nodeClusterId,proto3" json:"nodeClusterId,omitempty"`
	MetricItemId  int64 `protobuf:"varint,2,opt,name=metricItemId,proto3" json:"metricItemId,omitempty"`
}

func (x *DisableNodeClusterMetricItemRequest) Reset() {
	*x = DisableNodeClusterMetricItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_cluster_metric_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisableNodeClusterMetricItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisableNodeClusterMetricItemRequest) ProtoMessage() {}

func (x *DisableNodeClusterMetricItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_cluster_metric_item_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisableNodeClusterMetricItemRequest.ProtoReflect.Descriptor instead.
func (*DisableNodeClusterMetricItemRequest) Descriptor() ([]byte, []int) {
	return file_service_node_cluster_metric_item_proto_rawDescGZIP(), []int{1}
}

func (x *DisableNodeClusterMetricItemRequest) GetNodeClusterId() int64 {
	if x != nil {
		return x.NodeClusterId
	}
	return 0
}

func (x *DisableNodeClusterMetricItemRequest) GetMetricItemId() int64 {
	if x != nil {
		return x.MetricItemId
	}
	return 0
}

// 查找集群中所有指标
type FindAllNodeClusterMetricItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeClusterId int64 `protobuf:"varint,1,opt,name=nodeClusterId,proto3" json:"nodeClusterId,omitempty"`
}

func (x *FindAllNodeClusterMetricItemsRequest) Reset() {
	*x = FindAllNodeClusterMetricItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_cluster_metric_item_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllNodeClusterMetricItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllNodeClusterMetricItemsRequest) ProtoMessage() {}

func (x *FindAllNodeClusterMetricItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_cluster_metric_item_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllNodeClusterMetricItemsRequest.ProtoReflect.Descriptor instead.
func (*FindAllNodeClusterMetricItemsRequest) Descriptor() ([]byte, []int) {
	return file_service_node_cluster_metric_item_proto_rawDescGZIP(), []int{2}
}

func (x *FindAllNodeClusterMetricItemsRequest) GetNodeClusterId() int64 {
	if x != nil {
		return x.NodeClusterId
	}
	return 0
}

type FindAllNodeClusterMetricItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetricItems []*MetricItem `protobuf:"bytes,1,rep,name=metricItems,proto3" json:"metricItems,omitempty"`
}

func (x *FindAllNodeClusterMetricItemsResponse) Reset() {
	*x = FindAllNodeClusterMetricItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_cluster_metric_item_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllNodeClusterMetricItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllNodeClusterMetricItemsResponse) ProtoMessage() {}

func (x *FindAllNodeClusterMetricItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_cluster_metric_item_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllNodeClusterMetricItemsResponse.ProtoReflect.Descriptor instead.
func (*FindAllNodeClusterMetricItemsResponse) Descriptor() ([]byte, []int) {
	return file_service_node_cluster_metric_item_proto_rawDescGZIP(), []int{3}
}

func (x *FindAllNodeClusterMetricItemsResponse) GetMetricItems() []*MetricItem {
	if x != nil {
		return x.MetricItems
	}
	return nil
}

var File_service_node_cluster_metric_item_proto protoreflect.FileDescriptor

var file_service_node_cluster_metric_item_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x69, 0x74,
	0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x19, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x22, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x49, 0x74, 0x65,
	0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x6f, 0x0a, 0x23, 0x44, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24,
	0x0a, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x49, 0x74,
	0x65, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x24, 0x46, 0x69, 0x6e, 0x64,
	0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x59, 0x0a, 0x25, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c,
	0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x30, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x32, 0xc4, 0x02, 0x0a, 0x1c, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x55, 0x0a, 0x1b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x26, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x6f, 0x64,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52,
	0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x57, 0x0a, 0x1c, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x27, 0x2e, 0x70, 0x62, 0x2e, 0x44,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x74, 0x0a, 0x1d, 0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x64,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x28, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c,
	0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_node_cluster_metric_item_proto_rawDescOnce sync.Once
	file_service_node_cluster_metric_item_proto_rawDescData = file_service_node_cluster_metric_item_proto_rawDesc
)

func file_service_node_cluster_metric_item_proto_rawDescGZIP() []byte {
	file_service_node_cluster_metric_item_proto_rawDescOnce.Do(func() {
		file_service_node_cluster_metric_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_node_cluster_metric_item_proto_rawDescData)
	})
	return file_service_node_cluster_metric_item_proto_rawDescData
}

var file_service_node_cluster_metric_item_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_service_node_cluster_metric_item_proto_goTypes = []interface{}{
	(*EnableNodeClusterMetricItemRequest)(nil),    // 0: pb.EnableNodeClusterMetricItemRequest
	(*DisableNodeClusterMetricItemRequest)(nil),   // 1: pb.DisableNodeClusterMetricItemRequest
	(*FindAllNodeClusterMetricItemsRequest)(nil),  // 2: pb.FindAllNodeClusterMetricItemsRequest
	(*FindAllNodeClusterMetricItemsResponse)(nil), // 3: pb.FindAllNodeClusterMetricItemsResponse
	(*MetricItem)(nil),                            // 4: pb.MetricItem
	(*RPCSuccess)(nil),                            // 5: pb.RPCSuccess
}
var file_service_node_cluster_metric_item_proto_depIdxs = []int32{
	4, // 0: pb.FindAllNodeClusterMetricItemsResponse.metricItems:type_name -> pb.MetricItem
	0, // 1: pb.NodeClusterMetricItemService.enableNodeClusterMetricItem:input_type -> pb.EnableNodeClusterMetricItemRequest
	1, // 2: pb.NodeClusterMetricItemService.disableNodeClusterMetricItem:input_type -> pb.DisableNodeClusterMetricItemRequest
	2, // 3: pb.NodeClusterMetricItemService.findAllNodeClusterMetricItems:input_type -> pb.FindAllNodeClusterMetricItemsRequest
	5, // 4: pb.NodeClusterMetricItemService.enableNodeClusterMetricItem:output_type -> pb.RPCSuccess
	5, // 5: pb.NodeClusterMetricItemService.disableNodeClusterMetricItem:output_type -> pb.RPCSuccess
	3, // 6: pb.NodeClusterMetricItemService.findAllNodeClusterMetricItems:output_type -> pb.FindAllNodeClusterMetricItemsResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_node_cluster_metric_item_proto_init() }
func file_service_node_cluster_metric_item_proto_init() {
	if File_service_node_cluster_metric_item_proto != nil {
		return
	}
	file_models_rpc_messages_proto_init()
	file_models_model_metric_item_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_node_cluster_metric_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnableNodeClusterMetricItemRequest); i {
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
		file_service_node_cluster_metric_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisableNodeClusterMetricItemRequest); i {
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
		file_service_node_cluster_metric_item_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllNodeClusterMetricItemsRequest); i {
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
		file_service_node_cluster_metric_item_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllNodeClusterMetricItemsResponse); i {
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
			RawDescriptor: file_service_node_cluster_metric_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_node_cluster_metric_item_proto_goTypes,
		DependencyIndexes: file_service_node_cluster_metric_item_proto_depIdxs,
		MessageInfos:      file_service_node_cluster_metric_item_proto_msgTypes,
	}.Build()
	File_service_node_cluster_metric_item_proto = out.File
	file_service_node_cluster_metric_item_proto_rawDesc = nil
	file_service_node_cluster_metric_item_proto_goTypes = nil
	file_service_node_cluster_metric_item_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NodeClusterMetricItemServiceClient is the client API for NodeClusterMetricItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeClusterMetricItemServiceClient interface {
	// 启用某个指标
	EnableNodeClusterMetricItem(ctx context.Context, in *EnableNodeClusterMetricItemRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 禁用某个指标
	DisableNodeClusterMetricItem(ctx context.Context, in *DisableNodeClusterMetricItemRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查找集群中所有指标
	FindAllNodeClusterMetricItems(ctx context.Context, in *FindAllNodeClusterMetricItemsRequest, opts ...grpc.CallOption) (*FindAllNodeClusterMetricItemsResponse, error)
}

type nodeClusterMetricItemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeClusterMetricItemServiceClient(cc grpc.ClientConnInterface) NodeClusterMetricItemServiceClient {
	return &nodeClusterMetricItemServiceClient{cc}
}

func (c *nodeClusterMetricItemServiceClient) EnableNodeClusterMetricItem(ctx context.Context, in *EnableNodeClusterMetricItemRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.NodeClusterMetricItemService/enableNodeClusterMetricItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClusterMetricItemServiceClient) DisableNodeClusterMetricItem(ctx context.Context, in *DisableNodeClusterMetricItemRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.NodeClusterMetricItemService/disableNodeClusterMetricItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClusterMetricItemServiceClient) FindAllNodeClusterMetricItems(ctx context.Context, in *FindAllNodeClusterMetricItemsRequest, opts ...grpc.CallOption) (*FindAllNodeClusterMetricItemsResponse, error) {
	out := new(FindAllNodeClusterMetricItemsResponse)
	err := c.cc.Invoke(ctx, "/pb.NodeClusterMetricItemService/findAllNodeClusterMetricItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeClusterMetricItemServiceServer is the server API for NodeClusterMetricItemService service.
type NodeClusterMetricItemServiceServer interface {
	// 启用某个指标
	EnableNodeClusterMetricItem(context.Context, *EnableNodeClusterMetricItemRequest) (*RPCSuccess, error)
	// 禁用某个指标
	DisableNodeClusterMetricItem(context.Context, *DisableNodeClusterMetricItemRequest) (*RPCSuccess, error)
	// 查找集群中所有指标
	FindAllNodeClusterMetricItems(context.Context, *FindAllNodeClusterMetricItemsRequest) (*FindAllNodeClusterMetricItemsResponse, error)
}

// UnimplementedNodeClusterMetricItemServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNodeClusterMetricItemServiceServer struct {
}

func (*UnimplementedNodeClusterMetricItemServiceServer) EnableNodeClusterMetricItem(context.Context, *EnableNodeClusterMetricItemRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableNodeClusterMetricItem not implemented")
}
func (*UnimplementedNodeClusterMetricItemServiceServer) DisableNodeClusterMetricItem(context.Context, *DisableNodeClusterMetricItemRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableNodeClusterMetricItem not implemented")
}
func (*UnimplementedNodeClusterMetricItemServiceServer) FindAllNodeClusterMetricItems(context.Context, *FindAllNodeClusterMetricItemsRequest) (*FindAllNodeClusterMetricItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllNodeClusterMetricItems not implemented")
}

func RegisterNodeClusterMetricItemServiceServer(s *grpc.Server, srv NodeClusterMetricItemServiceServer) {
	s.RegisterService(&_NodeClusterMetricItemService_serviceDesc, srv)
}

func _NodeClusterMetricItemService_EnableNodeClusterMetricItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnableNodeClusterMetricItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeClusterMetricItemServiceServer).EnableNodeClusterMetricItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeClusterMetricItemService/EnableNodeClusterMetricItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeClusterMetricItemServiceServer).EnableNodeClusterMetricItem(ctx, req.(*EnableNodeClusterMetricItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeClusterMetricItemService_DisableNodeClusterMetricItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisableNodeClusterMetricItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeClusterMetricItemServiceServer).DisableNodeClusterMetricItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeClusterMetricItemService/DisableNodeClusterMetricItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeClusterMetricItemServiceServer).DisableNodeClusterMetricItem(ctx, req.(*DisableNodeClusterMetricItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeClusterMetricItemService_FindAllNodeClusterMetricItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllNodeClusterMetricItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeClusterMetricItemServiceServer).FindAllNodeClusterMetricItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NodeClusterMetricItemService/FindAllNodeClusterMetricItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeClusterMetricItemServiceServer).FindAllNodeClusterMetricItems(ctx, req.(*FindAllNodeClusterMetricItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeClusterMetricItemService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.NodeClusterMetricItemService",
	HandlerType: (*NodeClusterMetricItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "enableNodeClusterMetricItem",
			Handler:    _NodeClusterMetricItemService_EnableNodeClusterMetricItem_Handler,
		},
		{
			MethodName: "disableNodeClusterMetricItem",
			Handler:    _NodeClusterMetricItemService_DisableNodeClusterMetricItem_Handler,
		},
		{
			MethodName: "findAllNodeClusterMetricItems",
			Handler:    _NodeClusterMetricItemService_FindAllNodeClusterMetricItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_node_cluster_metric_item.proto",
}