// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: imagepool.proto

package __

import (
	context "context"
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

type State int32

const (
	State_S_READY State = 0
	State_S_BUSY  State = 1
	State_S_DONE  State = 2
	State_S_ERROR State = 3
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "S_READY",
		1: "S_BUSY",
		2: "S_DONE",
		3: "S_ERROR",
	}
	State_value = map[string]int32{
		"S_READY": 0,
		"S_BUSY":  1,
		"S_DONE":  2,
		"S_ERROR": 3,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_imagepool_proto_enumTypes[0].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_imagepool_proto_enumTypes[0]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_imagepool_proto_rawDescGZIP(), []int{0}
}

type Cmd int32

const (
	Cmd_C_WAIT Cmd = 0
	Cmd_C_PUSH Cmd = 1
	Cmd_C_GET  Cmd = 2
	Cmd_C_LIST Cmd = 3
)

// Enum value maps for Cmd.
var (
	Cmd_name = map[int32]string{
		0: "C_WAIT",
		1: "C_PUSH",
		2: "C_GET",
		3: "C_LIST",
	}
	Cmd_value = map[string]int32{
		"C_WAIT": 0,
		"C_PUSH": 1,
		"C_GET":  2,
		"C_LIST": 3,
	}
)

func (x Cmd) Enum() *Cmd {
	p := new(Cmd)
	*p = x
	return p
}

func (x Cmd) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Cmd) Descriptor() protoreflect.EnumDescriptor {
	return file_imagepool_proto_enumTypes[1].Descriptor()
}

func (Cmd) Type() protoreflect.EnumType {
	return &file_imagepool_proto_enumTypes[1]
}

func (x Cmd) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Cmd.Descriptor instead.
func (Cmd) EnumDescriptor() ([]byte, []int) {
	return file_imagepool_proto_rawDescGZIP(), []int{1}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Bytes   []byte `protobuf:"bytes,2,opt,name=Bytes,proto3" json:"Bytes,omitempty"`
	State   State  `protobuf:"varint,3,opt,name=State,proto3,enum=imagepool.State" json:"State,omitempty"`
	Cmd     Cmd    `protobuf:"varint,4,opt,name=Cmd,proto3,enum=imagepool.Cmd" json:"Cmd,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imagepool_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_imagepool_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_imagepool_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Request) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *Request) GetState() State {
	if x != nil {
		return x.State
	}
	return State_S_READY
}

func (x *Request) GetCmd() Cmd {
	if x != nil {
		return x.Cmd
	}
	return Cmd_C_WAIT
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Bytes   []byte `protobuf:"bytes,2,opt,name=Bytes,proto3" json:"Bytes,omitempty"`
	State   State  `protobuf:"varint,3,opt,name=State,proto3,enum=imagepool.State" json:"State,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imagepool_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_imagepool_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_imagepool_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Response) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *Response) GetState() State {
	if x != nil {
		return x.State
	}
	return State_S_READY
}

var File_imagepool_proto protoreflect.FileDescriptor

var file_imagepool_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x70, 0x6f, 0x6f, 0x6c, 0x22, 0x83, 0x01, 0x0a,
	0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x70,
	0x6f, 0x6f, 0x6c, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x20, 0x0a, 0x03, 0x43, 0x6d, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x63, 0x6d, 0x64, 0x52, 0x03, 0x43,
	0x6d, 0x64, 0x22, 0x62, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x26,
	0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2a, 0x39, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x53, 0x5f, 0x42, 0x55, 0x53, 0x59, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x5f, 0x44, 0x4f,
	0x4e, 0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x03, 0x2a, 0x34, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x5f, 0x57, 0x41,
	0x49, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x5f, 0x50, 0x55, 0x53, 0x48, 0x10, 0x01,
	0x12, 0x09, 0x0a, 0x05, 0x43, 0x5f, 0x47, 0x45, 0x54, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x43,
	0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x03, 0x32, 0x4e, 0x0a, 0x10, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x45,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x32, 0x49, 0x0a, 0x10, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x70, 0x6f, 0x6f, 0x6c, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x70,
	0x6f, 0x6f, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_imagepool_proto_rawDescOnce sync.Once
	file_imagepool_proto_rawDescData = file_imagepool_proto_rawDesc
)

func file_imagepool_proto_rawDescGZIP() []byte {
	file_imagepool_proto_rawDescOnce.Do(func() {
		file_imagepool_proto_rawDescData = protoimpl.X.CompressGZIP(file_imagepool_proto_rawDescData)
	})
	return file_imagepool_proto_rawDescData
}

var file_imagepool_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_imagepool_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_imagepool_proto_goTypes = []interface{}{
	(State)(0),       // 0: imagepool.state
	(Cmd)(0),         // 1: imagepool.cmd
	(*Request)(nil),  // 2: imagepool.Request
	(*Response)(nil), // 3: imagepool.Response
}
var file_imagepool_proto_depIdxs = []int32{
	0, // 0: imagepool.Request.State:type_name -> imagepool.state
	1, // 1: imagepool.Request.Cmd:type_name -> imagepool.cmd
	0, // 2: imagepool.Response.State:type_name -> imagepool.state
	2, // 3: imagepool.ImagePoolService.Exchanger:input_type -> imagepool.Request
	2, // 4: imagepool.ImageListService.List:input_type -> imagepool.Request
	3, // 5: imagepool.ImagePoolService.Exchanger:output_type -> imagepool.Response
	3, // 6: imagepool.ImageListService.List:output_type -> imagepool.Response
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_imagepool_proto_init() }
func file_imagepool_proto_init() {
	if File_imagepool_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_imagepool_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_imagepool_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_imagepool_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_imagepool_proto_goTypes,
		DependencyIndexes: file_imagepool_proto_depIdxs,
		EnumInfos:         file_imagepool_proto_enumTypes,
		MessageInfos:      file_imagepool_proto_msgTypes,
	}.Build()
	File_imagepool_proto = out.File
	file_imagepool_proto_rawDesc = nil
	file_imagepool_proto_goTypes = nil
	file_imagepool_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ImagePoolServiceClient is the client API for ImagePoolService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ImagePoolServiceClient interface {
	Exchanger(ctx context.Context, opts ...grpc.CallOption) (ImagePoolService_ExchangerClient, error)
}

type imagePoolServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImagePoolServiceClient(cc grpc.ClientConnInterface) ImagePoolServiceClient {
	return &imagePoolServiceClient{cc}
}

func (c *imagePoolServiceClient) Exchanger(ctx context.Context, opts ...grpc.CallOption) (ImagePoolService_ExchangerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ImagePoolService_serviceDesc.Streams[0], "/imagepool.ImagePoolService/Exchanger", opts...)
	if err != nil {
		return nil, err
	}
	x := &imagePoolServiceExchangerClient{stream}
	return x, nil
}

type ImagePoolService_ExchangerClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type imagePoolServiceExchangerClient struct {
	grpc.ClientStream
}

func (x *imagePoolServiceExchangerClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *imagePoolServiceExchangerClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ImagePoolServiceServer is the server API for ImagePoolService service.
type ImagePoolServiceServer interface {
	Exchanger(ImagePoolService_ExchangerServer) error
}

// UnimplementedImagePoolServiceServer can be embedded to have forward compatible implementations.
type UnimplementedImagePoolServiceServer struct {
}

func (*UnimplementedImagePoolServiceServer) Exchanger(ImagePoolService_ExchangerServer) error {
	return status.Errorf(codes.Unimplemented, "method Exchanger not implemented")
}

func RegisterImagePoolServiceServer(s *grpc.Server, srv ImagePoolServiceServer) {
	s.RegisterService(&_ImagePoolService_serviceDesc, srv)
}

func _ImagePoolService_Exchanger_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ImagePoolServiceServer).Exchanger(&imagePoolServiceExchangerServer{stream})
}

type ImagePoolService_ExchangerServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type imagePoolServiceExchangerServer struct {
	grpc.ServerStream
}

func (x *imagePoolServiceExchangerServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *imagePoolServiceExchangerServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ImagePoolService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "imagepool.ImagePoolService",
	HandlerType: (*ImagePoolServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Exchanger",
			Handler:       _ImagePoolService_Exchanger_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "imagepool.proto",
}

// ImageListServiceClient is the client API for ImageListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ImageListServiceClient interface {
	List(ctx context.Context, opts ...grpc.CallOption) (ImageListService_ListClient, error)
}

type imageListServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImageListServiceClient(cc grpc.ClientConnInterface) ImageListServiceClient {
	return &imageListServiceClient{cc}
}

func (c *imageListServiceClient) List(ctx context.Context, opts ...grpc.CallOption) (ImageListService_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ImageListService_serviceDesc.Streams[0], "/imagepool.ImageListService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &imageListServiceListClient{stream}
	return x, nil
}

type ImageListService_ListClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type imageListServiceListClient struct {
	grpc.ClientStream
}

func (x *imageListServiceListClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *imageListServiceListClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ImageListServiceServer is the server API for ImageListService service.
type ImageListServiceServer interface {
	List(ImageListService_ListServer) error
}

// UnimplementedImageListServiceServer can be embedded to have forward compatible implementations.
type UnimplementedImageListServiceServer struct {
}

func (*UnimplementedImageListServiceServer) List(ImageListService_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterImageListServiceServer(s *grpc.Server, srv ImageListServiceServer) {
	s.RegisterService(&_ImageListService_serviceDesc, srv)
}

func _ImageListService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ImageListServiceServer).List(&imageListServiceListServer{stream})
}

type ImageListService_ListServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type imageListServiceListServer struct {
	grpc.ServerStream
}

func (x *imageListServiceListServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *imageListServiceListServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ImageListService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "imagepool.ImageListService",
	HandlerType: (*ImageListServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _ImageListService_List_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "imagepool.proto",
}
