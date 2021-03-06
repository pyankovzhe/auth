// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: account-event.proto

package eventpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventKind int32

const (
	EventKind_CREATED EventKind = 0
	EventKind_UPDATED EventKind = 1
	EventKind_DELETED EventKind = 2
)

// Enum value maps for EventKind.
var (
	EventKind_name = map[int32]string{
		0: "CREATED",
		1: "UPDATED",
		2: "DELETED",
	}
	EventKind_value = map[string]int32{
		"CREATED": 0,
		"UPDATED": 1,
		"DELETED": 2,
	}
)

func (x EventKind) Enum() *EventKind {
	p := new(EventKind)
	*p = x
	return p
}

func (x EventKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventKind) Descriptor() protoreflect.EnumDescriptor {
	return file_account_event_proto_enumTypes[0].Descriptor()
}

func (EventKind) Type() protoreflect.EnumType {
	return &file_account_event_proto_enumTypes[0]
}

func (x EventKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventKind.Descriptor instead.
func (EventKind) EnumDescriptor() ([]byte, []int) {
	return file_account_event_proto_rawDescGZIP(), []int{0}
}

type AccountEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid      string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Producer  string                 `protobuf:"bytes,2,opt,name=producer,proto3" json:"producer,omitempty"`
	EventTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	Kind      EventKind              `protobuf:"varint,4,opt,name=kind,proto3,enum=v1.EventKind" json:"kind,omitempty"`
	Data      *Account               `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AccountEvent) Reset() {
	*x = AccountEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountEvent) ProtoMessage() {}

func (x *AccountEvent) ProtoReflect() protoreflect.Message {
	mi := &file_account_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountEvent.ProtoReflect.Descriptor instead.
func (*AccountEvent) Descriptor() ([]byte, []int) {
	return file_account_event_proto_rawDescGZIP(), []int{0}
}

func (x *AccountEvent) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *AccountEvent) GetProducer() string {
	if x != nil {
		return x.Producer
	}
	return ""
}

func (x *AccountEvent) GetEventTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EventTime
	}
	return nil
}

func (x *AccountEvent) GetKind() EventKind {
	if x != nil {
		return x.Kind
	}
	return EventKind_CREATED
}

func (x *AccountEvent) GetData() *Account {
	if x != nil {
		return x.Data
	}
	return nil
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid  string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Login string `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_account_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_account_event_proto_rawDescGZIP(), []int{1}
}

func (x *Account) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Account) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

var File_account_event_proto protoreflect.FileDescriptor

var file_account_event_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x0c, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x0a, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4b,
	0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x1f, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x33, 0x0a, 0x07, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2a,
	0x32, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0b, 0x0a, 0x07,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x50, 0x44,
	0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45,
	0x44, 0x10, 0x02, 0x42, 0x0c, 0x5a, 0x0a, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_event_proto_rawDescOnce sync.Once
	file_account_event_proto_rawDescData = file_account_event_proto_rawDesc
)

func file_account_event_proto_rawDescGZIP() []byte {
	file_account_event_proto_rawDescOnce.Do(func() {
		file_account_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_event_proto_rawDescData)
	})
	return file_account_event_proto_rawDescData
}

var file_account_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_account_event_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_account_event_proto_goTypes = []interface{}{
	(EventKind)(0),                // 0: v1.EventKind
	(*AccountEvent)(nil),          // 1: v1.AccountEvent
	(*Account)(nil),               // 2: v1.Account
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_account_event_proto_depIdxs = []int32{
	3, // 0: v1.AccountEvent.event_time:type_name -> google.protobuf.Timestamp
	0, // 1: v1.AccountEvent.kind:type_name -> v1.EventKind
	2, // 2: v1.AccountEvent.data:type_name -> v1.Account
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_account_event_proto_init() }
func file_account_event_proto_init() {
	if File_account_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_account_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountEvent); i {
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
		file_account_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
			RawDescriptor: file_account_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_account_event_proto_goTypes,
		DependencyIndexes: file_account_event_proto_depIdxs,
		EnumInfos:         file_account_event_proto_enumTypes,
		MessageInfos:      file_account_event_proto_msgTypes,
	}.Build()
	File_account_event_proto = out.File
	file_account_event_proto_rawDesc = nil
	file_account_event_proto_goTypes = nil
	file_account_event_proto_depIdxs = nil
}
