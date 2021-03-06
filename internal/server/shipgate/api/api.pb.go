// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: api.proto

package api

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type ShipList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ships []*ShipList_Ship `protobuf:"bytes,1,rep,name=ships,proto3" json:"ships,omitempty"`
}

func (x *ShipList) Reset() {
	*x = ShipList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipList) ProtoMessage() {}

func (x *ShipList) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipList.ProtoReflect.Descriptor instead.
func (*ShipList) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *ShipList) GetShips() []*ShipList_Ship {
	if x != nil {
		return x.Ships
	}
	return nil
}

type ShipList_Ship struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Ip          string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Port        string `protobuf:"bytes,4,opt,name=port,proto3" json:"port,omitempty"`
	PlayerCount int32  `protobuf:"varint,5,opt,name=playerCount,proto3" json:"playerCount,omitempty"`
}

func (x *ShipList_Ship) Reset() {
	*x = ShipList_Ship{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipList_Ship) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipList_Ship) ProtoMessage() {}

func (x *ShipList_Ship) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipList_Ship.ProtoReflect.Descriptor instead.
func (*ShipList_Ship) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ShipList_Ship) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ShipList_Ship) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ShipList_Ship) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *ShipList_Ship) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *ShipList_Ship) GetPlayerCount() int32 {
	if x != nil {
		return x.PlayerCount
	}
	return 0
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01,
	0x0a, 0x08, 0x53, 0x68, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x73, 0x68,
	0x69, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x53, 0x68, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x52, 0x05, 0x73,
	0x68, 0x69, 0x70, 0x73, 0x1a, 0x70, 0x0a, 0x04, 0x53, 0x68, 0x69, 0x70, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x4a, 0x0a, 0x0f, 0x53, 0x68, 0x69, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x68, 0x69, 0x70, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x4c, 0x69,
	0x73, 0x74, 0x32, 0x11, 0x0a, 0x0f, 0x53, 0x68, 0x69, 0x70, 0x67, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_goTypes = []interface{}{
	(*ShipList)(nil),      // 0: api.ShipList
	(*ShipList_Ship)(nil), // 1: api.ShipList.Ship
	(*empty.Empty)(nil),   // 2: google.protobuf.Empty
}
var file_api_proto_depIdxs = []int32{
	1, // 0: api.ShipList.ships:type_name -> api.ShipList.Ship
	2, // 1: api.ShipInfoService.GetActiveShips:input_type -> google.protobuf.Empty
	0, // 2: api.ShipInfoService.GetActiveShips:output_type -> api.ShipList
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipList); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipList_Ship); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ShipInfoServiceClient is the client API for ShipInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShipInfoServiceClient interface {
	// GetActiveShips returns the list of Ships that currently connected to the
	// shipgate and ready to receive players.
	GetActiveShips(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ShipList, error)
}

type shipInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShipInfoServiceClient(cc grpc.ClientConnInterface) ShipInfoServiceClient {
	return &shipInfoServiceClient{cc}
}

func (c *shipInfoServiceClient) GetActiveShips(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ShipList, error) {
	out := new(ShipList)
	err := c.cc.Invoke(ctx, "/api.ShipInfoService/GetActiveShips", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShipInfoServiceServer is the server API for ShipInfoService service.
type ShipInfoServiceServer interface {
	// GetActiveShips returns the list of Ships that currently connected to the
	// shipgate and ready to receive players.
	GetActiveShips(context.Context, *empty.Empty) (*ShipList, error)
}

// UnimplementedShipInfoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedShipInfoServiceServer struct {
}

func (*UnimplementedShipInfoServiceServer) GetActiveShips(context.Context, *empty.Empty) (*ShipList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveShips not implemented")
}

func RegisterShipInfoServiceServer(s *grpc.Server, srv ShipInfoServiceServer) {
	s.RegisterService(&_ShipInfoService_serviceDesc, srv)
}

func _ShipInfoService_GetActiveShips_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShipInfoServiceServer).GetActiveShips(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ShipInfoService/GetActiveShips",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShipInfoServiceServer).GetActiveShips(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShipInfoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ShipInfoService",
	HandlerType: (*ShipInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetActiveShips",
			Handler:    _ShipInfoService_GetActiveShips_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// ShipgateServiceClient is the client API for ShipgateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShipgateServiceClient interface {
}

type shipgateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShipgateServiceClient(cc grpc.ClientConnInterface) ShipgateServiceClient {
	return &shipgateServiceClient{cc}
}

// ShipgateServiceServer is the server API for ShipgateService service.
type ShipgateServiceServer interface {
}

// UnimplementedShipgateServiceServer can be embedded to have forward compatible implementations.
type UnimplementedShipgateServiceServer struct {
}

func RegisterShipgateServiceServer(s *grpc.Server, srv ShipgateServiceServer) {
	s.RegisterService(&_ShipgateService_serviceDesc, srv)
}

var _ShipgateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ShipgateService",
	HandlerType: (*ShipgateServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "api.proto",
}
