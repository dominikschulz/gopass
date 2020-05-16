// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.12.0
// source: gpb.proto

package gpb

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Revision struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Created  *timestamp.Timestamp `protobuf:"bytes,1,opt,name=Created,proto3" json:"Created,omitempty"`
	Message  string               `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	Filename string               `protobuf:"bytes,3,opt,name=Filename,proto3" json:"Filename,omitempty"`
}

func (x *Revision) Reset() {
	*x = Revision{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Revision) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Revision) ProtoMessage() {}

func (x *Revision) ProtoReflect() protoreflect.Message {
	mi := &file_gpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Revision.ProtoReflect.Descriptor instead.
func (*Revision) Descriptor() ([]byte, []int) {
	return file_gpb_proto_rawDescGZIP(), []int{0}
}

func (x *Revision) GetCreated() *timestamp.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Revision) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Revision) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string      `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Revisions []*Revision `protobuf:"bytes,2,rep,name=Revisions,proto3" json:"Revisions,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_gpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_gpb_proto_rawDescGZIP(), []int{1}
}

func (x *Entry) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Entry) GetRevisions() []*Revision {
	if x != nil {
		return x.Revisions
	}
	return nil
}

type Store struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string            `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Entries map[string]*Entry `protobuf:"bytes,2,rep,name=Entries,proto3" json:"Entries,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Store) Reset() {
	*x = Store{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gpb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Store) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Store) ProtoMessage() {}

func (x *Store) ProtoReflect() protoreflect.Message {
	mi := &file_gpb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Store.ProtoReflect.Descriptor instead.
func (*Store) Descriptor() ([]byte, []int) {
	return file_gpb_proto_rawDescGZIP(), []int{2}
}

func (x *Store) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Store) GetEntries() map[string]*Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

var File_gpb_proto protoreflect.FileDescriptor

var file_gpb_proto_rawDesc = []byte{
	0x0a, 0x09, 0x67, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x67, 0x70, 0x62,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x76, 0x0a, 0x08, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a,
	0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x05, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x09, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x70, 0x62, 0x2e,
	0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x96, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x31, 0x0a, 0x07, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x45,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x45, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x1a, 0x46, 0x0a, 0x0c, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x70, 0x62, 0x2e, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x3b, 0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gpb_proto_rawDescOnce sync.Once
	file_gpb_proto_rawDescData = file_gpb_proto_rawDesc
)

func file_gpb_proto_rawDescGZIP() []byte {
	file_gpb_proto_rawDescOnce.Do(func() {
		file_gpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_gpb_proto_rawDescData)
	})
	return file_gpb_proto_rawDescData
}

var file_gpb_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_gpb_proto_goTypes = []interface{}{
	(*Revision)(nil),            // 0: gpb.Revision
	(*Entry)(nil),               // 1: gpb.Entry
	(*Store)(nil),               // 2: gpb.Store
	nil,                         // 3: gpb.Store.EntriesEntry
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_gpb_proto_depIdxs = []int32{
	4, // 0: gpb.Revision.Created:type_name -> google.protobuf.Timestamp
	0, // 1: gpb.Entry.Revisions:type_name -> gpb.Revision
	3, // 2: gpb.Store.Entries:type_name -> gpb.Store.EntriesEntry
	1, // 3: gpb.Store.EntriesEntry.value:type_name -> gpb.Entry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_gpb_proto_init() }
func file_gpb_proto_init() {
	if File_gpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Revision); i {
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
		file_gpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
		file_gpb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Store); i {
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
			RawDescriptor: file_gpb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gpb_proto_goTypes,
		DependencyIndexes: file_gpb_proto_depIdxs,
		MessageInfos:      file_gpb_proto_msgTypes,
	}.Build()
	File_gpb_proto = out.File
	file_gpb_proto_rawDesc = nil
	file_gpb_proto_goTypes = nil
	file_gpb_proto_depIdxs = nil
}
