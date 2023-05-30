// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: app/tun/config.proto

package tun

import (
	routercommon "github.com/v2fly/v2ray-core/v5/app/router/routercommon"
	_ "github.com/v2fly/v2ray-core/v5/common/protoext"
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

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                  string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Mtu                   uint32               `protobuf:"varint,2,opt,name=mtu,proto3" json:"mtu,omitempty"`
	UserLevel             uint32               `protobuf:"varint,3,opt,name=user_level,json=userLevel,proto3" json:"user_level,omitempty"`
	Tag                   string               `protobuf:"bytes,5,opt,name=tag,proto3" json:"tag,omitempty"`
	Ips                   []*routercommon.CIDR `protobuf:"bytes,6,rep,name=ips,proto3" json:"ips,omitempty"`
	Routes                []*routercommon.CIDR `protobuf:"bytes,7,rep,name=routes,proto3" json:"routes,omitempty"`
	EnablePromiscuousMode bool                 `protobuf:"varint,8,opt,name=enable_promiscuous_mode,json=enablePromiscuousMode,proto3" json:"enable_promiscuous_mode,omitempty"`
	EnableSpoofing        bool                 `protobuf:"varint,9,opt,name=enable_spoofing,json=enableSpoofing,proto3" json:"enable_spoofing,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_tun_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_app_tun_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_app_tun_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Config) GetMtu() uint32 {
	if x != nil {
		return x.Mtu
	}
	return 0
}

func (x *Config) GetUserLevel() uint32 {
	if x != nil {
		return x.UserLevel
	}
	return 0
}

func (x *Config) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Config) GetIps() []*routercommon.CIDR {
	if x != nil {
		return x.Ips
	}
	return nil
}

func (x *Config) GetRoutes() []*routercommon.CIDR {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *Config) GetEnablePromiscuousMode() bool {
	if x != nil {
		return x.EnablePromiscuousMode
	}
	return false
}

func (x *Config) GetEnableSpoofing() bool {
	if x != nil {
		return x.EnableSpoofing
	}
	return false
}

var File_app_tun_config_proto protoreflect.FileDescriptor

var file_app_tun_config_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x70, 0x2f, 0x74, 0x75, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x74, 0x75, 0x6e, 0x1a, 0x24, 0x61, 0x70, 0x70, 0x2f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x78,
	0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xd8, 0x02, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x74, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x6d, 0x74, 0x75, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x74, 0x61, 0x67, 0x12, 0x3a, 0x0a, 0x03, 0x69, 0x70, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x49, 0x44, 0x52, 0x52, 0x03, 0x69, 0x70, 0x73,
	0x12, 0x40, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x49, 0x44, 0x52, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x73, 0x12, 0x36, 0x0a, 0x17, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f,
	0x6d, 0x69, 0x73, 0x63, 0x75, 0x6f, 0x75, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x15, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x6d, 0x69,
	0x73, 0x63, 0x75, 0x6f, 0x75, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x70, 0x6f, 0x6f, 0x66, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x70, 0x6f, 0x6f, 0x66,
	0x69, 0x6e, 0x67, 0x3a, 0x12, 0x82, 0xb5, 0x18, 0x0e, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x03, 0x74, 0x75, 0x6e, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x42, 0x57, 0x0a,
	0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x74, 0x75, 0x6e, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x32, 0x66, 0x6c, 0x79, 0x2f, 0x76, 0x32, 0x72, 0x61,
	0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x35, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x74, 0x75,
	0x6e, 0xaa, 0x02, 0x12, 0x56, 0x32, 0x52, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x41,
	0x70, 0x70, 0x2e, 0x54, 0x75, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_tun_config_proto_rawDescOnce sync.Once
	file_app_tun_config_proto_rawDescData = file_app_tun_config_proto_rawDesc
)

func file_app_tun_config_proto_rawDescGZIP() []byte {
	file_app_tun_config_proto_rawDescOnce.Do(func() {
		file_app_tun_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_tun_config_proto_rawDescData)
	})
	return file_app_tun_config_proto_rawDescData
}

var file_app_tun_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_app_tun_config_proto_goTypes = []interface{}{
	(*Config)(nil),            // 0: v2ray.core.app.tun.Config
	(*routercommon.CIDR)(nil), // 1: v2ray.core.app.router.routercommon.CIDR
}
var file_app_tun_config_proto_depIdxs = []int32{
	1, // 0: v2ray.core.app.tun.Config.ips:type_name -> v2ray.core.app.router.routercommon.CIDR
	1, // 1: v2ray.core.app.tun.Config.routes:type_name -> v2ray.core.app.router.routercommon.CIDR
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_app_tun_config_proto_init() }
func file_app_tun_config_proto_init() {
	if File_app_tun_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_tun_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_app_tun_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_tun_config_proto_goTypes,
		DependencyIndexes: file_app_tun_config_proto_depIdxs,
		MessageInfos:      file_app_tun_config_proto_msgTypes,
	}.Build()
	File_app_tun_config_proto = out.File
	file_app_tun_config_proto_rawDesc = nil
	file_app_tun_config_proto_goTypes = nil
	file_app_tun_config_proto_depIdxs = nil
}