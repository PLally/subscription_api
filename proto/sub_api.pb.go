// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: proto/sub_api.proto

package proto

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

type Destination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Type       string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Id         uint32 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Destination) Reset() {
	*x = Destination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sub_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Destination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Destination) ProtoMessage() {}

func (x *Destination) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sub_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Destination.ProtoReflect.Descriptor instead.
func (*Destination) Descriptor() ([]byte, []int) {
	return file_proto_sub_api_proto_rawDescGZIP(), []int{0}
}

func (x *Destination) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *Destination) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Destination) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SubscriptionSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags string `protobuf:"bytes,1,opt,name=tags,proto3" json:"tags,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Id   uint32 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SubscriptionSource) Reset() {
	*x = SubscriptionSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sub_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionSource) ProtoMessage() {}

func (x *SubscriptionSource) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sub_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionSource.ProtoReflect.Descriptor instead.
func (*SubscriptionSource) Descriptor() ([]byte, []int) {
	return file_proto_sub_api_proto_rawDescGZIP(), []int{1}
}

func (x *SubscriptionSource) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *SubscriptionSource) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SubscriptionSource) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Destination        *Destination        `protobuf:"bytes,1,opt,name=destination,proto3" json:"destination,omitempty"`
	SubscriptionSource *SubscriptionSource `protobuf:"bytes,2,opt,name=subscription_source,json=subscriptionSource,proto3" json:"subscription_source,omitempty"`
	Id                 uint32              `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sub_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sub_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_proto_sub_api_proto_rawDescGZIP(), []int{2}
}

func (x *Subscription) GetDestination() *Destination {
	if x != nil {
		return x.Destination
	}
	return nil
}

func (x *Subscription) GetSubscriptionSource() *SubscriptionSource {
	if x != nil {
		return x.SubscriptionSource
	}
	return nil
}

func (x *Subscription) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SubscriptionList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscriptions []*Subscription `protobuf:"bytes,1,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
}

func (x *SubscriptionList) Reset() {
	*x = SubscriptionList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sub_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionList) ProtoMessage() {}

func (x *SubscriptionList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sub_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionList.ProtoReflect.Descriptor instead.
func (*SubscriptionList) Descriptor() ([]byte, []int) {
	return file_proto_sub_api_proto_rawDescGZIP(), []int{3}
}

func (x *SubscriptionList) GetSubscriptions() []*Subscription {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

type Success struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *Success) Reset() {
	*x = Success{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sub_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Success) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Success) ProtoMessage() {}

func (x *Success) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sub_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Success.ProtoReflect.Descriptor instead.
func (*Success) Descriptor() ([]byte, []int) {
	return file_proto_sub_api_proto_rawDescGZIP(), []int{4}
}

func (x *Success) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_sub_api_proto protoreflect.FileDescriptor

var file_proto_sub_api_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x75, 0x62, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4c, 0x0a, 0x12, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x94, 0x01, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x44,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x13, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x12, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x47, 0x0a,
	0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x33, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x23, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x9b, 0x01, 0x0a, 0x10,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69,
	0x12, 0x29, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x0d, 0x2e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0d, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0b, 0x55,
	0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x0d, 0x2e, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x08, 0x2e, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x34, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0c, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x11, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6c, 0x6c, 0x79, 0x2f, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_sub_api_proto_rawDescOnce sync.Once
	file_proto_sub_api_proto_rawDescData = file_proto_sub_api_proto_rawDesc
)

func file_proto_sub_api_proto_rawDescGZIP() []byte {
	file_proto_sub_api_proto_rawDescOnce.Do(func() {
		file_proto_sub_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_sub_api_proto_rawDescData)
	})
	return file_proto_sub_api_proto_rawDescData
}

var file_proto_sub_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_sub_api_proto_goTypes = []interface{}{
	(*Destination)(nil),        // 0: Destination
	(*SubscriptionSource)(nil), // 1: SubscriptionSource
	(*Subscription)(nil),       // 2: Subscription
	(*SubscriptionList)(nil),   // 3: SubscriptionList
	(*Success)(nil),            // 4: Success
}
var file_proto_sub_api_proto_depIdxs = []int32{
	0, // 0: Subscription.destination:type_name -> Destination
	1, // 1: Subscription.subscription_source:type_name -> SubscriptionSource
	2, // 2: SubscriptionList.subscriptions:type_name -> Subscription
	2, // 3: subscription_api.Subscribe:input_type -> Subscription
	2, // 4: subscription_api.UnSubscribe:input_type -> Subscription
	0, // 5: subscription_api.ListSubscriptions:input_type -> Destination
	2, // 6: subscription_api.Subscribe:output_type -> Subscription
	4, // 7: subscription_api.UnSubscribe:output_type -> Success
	3, // 8: subscription_api.ListSubscriptions:output_type -> SubscriptionList
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_sub_api_proto_init() }
func file_proto_sub_api_proto_init() {
	if File_proto_sub_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_sub_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Destination); i {
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
		file_proto_sub_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionSource); i {
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
		file_proto_sub_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscription); i {
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
		file_proto_sub_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionList); i {
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
		file_proto_sub_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Success); i {
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
			RawDescriptor: file_proto_sub_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_sub_api_proto_goTypes,
		DependencyIndexes: file_proto_sub_api_proto_depIdxs,
		MessageInfos:      file_proto_sub_api_proto_msgTypes,
	}.Build()
	File_proto_sub_api_proto = out.File
	file_proto_sub_api_proto_rawDesc = nil
	file_proto_sub_api_proto_goTypes = nil
	file_proto_sub_api_proto_depIdxs = nil
}
