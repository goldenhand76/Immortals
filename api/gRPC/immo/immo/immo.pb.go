// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: api/gRPC/immo/immo.proto

package immo

import (
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

type NodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *NodeRequest) Reset() {
	*x = NodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gRPC_immo_immo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeRequest) ProtoMessage() {}

func (x *NodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gRPC_immo_immo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeRequest.ProtoReflect.Descriptor instead.
func (*NodeRequest) Descriptor() ([]byte, []int) {
	return file_api_gRPC_immo_immo_proto_rawDescGZIP(), []int{0}
}

func (x *NodeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NodeRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type NodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId   string            `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Sensor   map[string]string `protobuf:"bytes,2,rep,name=sensor,proto3" json:"sensor,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Actuator map[string]string `protobuf:"bytes,3,rep,name=actuator,proto3" json:"actuator,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Error    string            `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *NodeResponse) Reset() {
	*x = NodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gRPC_immo_immo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeResponse) ProtoMessage() {}

func (x *NodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_gRPC_immo_immo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeResponse.ProtoReflect.Descriptor instead.
func (*NodeResponse) Descriptor() ([]byte, []int) {
	return file_api_gRPC_immo_immo_proto_rawDescGZIP(), []int{1}
}

func (x *NodeResponse) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *NodeResponse) GetSensor() map[string]string {
	if x != nil {
		return x.Sensor
	}
	return nil
}

func (x *NodeResponse) GetActuator() map[string]string {
	if x != nil {
		return x.Actuator
	}
	return nil
}

func (x *NodeResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*NodeResponse `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gRPC_immo_immo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_gRPC_immo_immo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_api_gRPC_immo_immo_proto_rawDescGZIP(), []int{2}
}

func (x *ListResponse) GetNodes() []*NodeResponse {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gRPC_immo_immo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gRPC_immo_immo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_api_gRPC_immo_immo_proto_rawDescGZIP(), []int{3}
}

var File_api_gRPC_immo_immo_proto protoreflect.FileDescriptor

var file_api_gRPC_immo_immo_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x52, 0x50, 0x43, 0x2f, 0x69, 0x6d, 0x6d, 0x6f, 0x2f,
	0x69, 0x6d, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x69, 0x6d, 0x6d, 0x6f,
	0x22, 0x3b, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xab, 0x02,
	0x0a, 0x0c, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x12,
	0x3c, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x08, 0x61, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x1a, 0x39, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3b,
	0x0a, 0x0d, 0x41, 0x63, 0x74, 0x75, 0x61, 0x74, 0x6f, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x38, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69, 0x6d, 0x6d,
	0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x0d, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x32, 0xaa, 0x01, 0x0a, 0x0b, 0x49, 0x6d, 0x6d, 0x6f, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x11, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x11, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x69, 0x6d, 0x6d, 0x6f,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a,
	0x09, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x11, 0x2e, 0x69, 0x6d, 0x6d,
	0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x69, 0x6d, 0x6d, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x69, 0x6d, 0x6d, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_gRPC_immo_immo_proto_rawDescOnce sync.Once
	file_api_gRPC_immo_immo_proto_rawDescData = file_api_gRPC_immo_immo_proto_rawDesc
)

func file_api_gRPC_immo_immo_proto_rawDescGZIP() []byte {
	file_api_gRPC_immo_immo_proto_rawDescOnce.Do(func() {
		file_api_gRPC_immo_immo_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_gRPC_immo_immo_proto_rawDescData)
	})
	return file_api_gRPC_immo_immo_proto_rawDescData
}

var file_api_gRPC_immo_immo_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_gRPC_immo_immo_proto_goTypes = []interface{}{
	(*NodeRequest)(nil),  // 0: immo.NodeRequest
	(*NodeResponse)(nil), // 1: immo.NodeResponse
	(*ListResponse)(nil), // 2: immo.ListResponse
	(*ListRequest)(nil),  // 3: immo.ListRequest
	nil,                  // 4: immo.NodeResponse.SensorEntry
	nil,                  // 5: immo.NodeResponse.ActuatorEntry
}
var file_api_gRPC_immo_immo_proto_depIdxs = []int32{
	4, // 0: immo.NodeResponse.sensor:type_name -> immo.NodeResponse.SensorEntry
	5, // 1: immo.NodeResponse.actuator:type_name -> immo.NodeResponse.ActuatorEntry
	1, // 2: immo.ListResponse.nodes:type_name -> immo.NodeResponse
	0, // 3: immo.ImmoService.DiscoverNode:input_type -> immo.NodeRequest
	0, // 4: immo.ImmoService.AddNode:input_type -> immo.NodeRequest
	3, // 5: immo.ImmoService.ListNodes:input_type -> immo.ListRequest
	1, // 6: immo.ImmoService.DiscoverNode:output_type -> immo.NodeResponse
	1, // 7: immo.ImmoService.AddNode:output_type -> immo.NodeResponse
	2, // 8: immo.ImmoService.ListNodes:output_type -> immo.ListResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_gRPC_immo_immo_proto_init() }
func file_api_gRPC_immo_immo_proto_init() {
	if File_api_gRPC_immo_immo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_gRPC_immo_immo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeRequest); i {
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
		file_api_gRPC_immo_immo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeResponse); i {
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
		file_api_gRPC_immo_immo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_api_gRPC_immo_immo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
			RawDescriptor: file_api_gRPC_immo_immo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_gRPC_immo_immo_proto_goTypes,
		DependencyIndexes: file_api_gRPC_immo_immo_proto_depIdxs,
		MessageInfos:      file_api_gRPC_immo_immo_proto_msgTypes,
	}.Build()
	File_api_gRPC_immo_immo_proto = out.File
	file_api_gRPC_immo_immo_proto_rawDesc = nil
	file_api_gRPC_immo_immo_proto_goTypes = nil
	file_api_gRPC_immo_immo_proto_depIdxs = nil
}
