// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: upload/proto/upload.proto

package proto

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

type UploadStatusCode int32

const (
	UploadStatusCode_Ok     UploadStatusCode = 0
	UploadStatusCode_Failed UploadStatusCode = 1
)

// Enum value maps for UploadStatusCode.
var (
	UploadStatusCode_name = map[int32]string{
		0: "Ok",
		1: "Failed",
	}
	UploadStatusCode_value = map[string]int32{
		"Ok":     0,
		"Failed": 1,
	}
)

func (x UploadStatusCode) Enum() *UploadStatusCode {
	p := new(UploadStatusCode)
	*p = x
	return p
}

func (x UploadStatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UploadStatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_upload_proto_upload_proto_enumTypes[0].Descriptor()
}

func (UploadStatusCode) Type() protoreflect.EnumType {
	return &file_upload_proto_upload_proto_enumTypes[0]
}

func (x UploadStatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UploadStatusCode.Descriptor instead.
func (UploadStatusCode) EnumDescriptor() ([]byte, []int) {
	return file_upload_proto_upload_proto_rawDescGZIP(), []int{0}
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_upload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_upload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_upload_proto_upload_proto_rawDescGZIP(), []int{0}
}

func (x *Chunk) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type UploadStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string           `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Code   UploadStatusCode `protobuf:"varint,2,opt,name=code,proto3,enum=greet.UploadStatusCode" json:"code,omitempty"`
}

func (x *UploadStatus) Reset() {
	*x = UploadStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_upload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadStatus) ProtoMessage() {}

func (x *UploadStatus) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_upload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadStatus.ProtoReflect.Descriptor instead.
func (*UploadStatus) Descriptor() ([]byte, []int) {
	return file_upload_proto_upload_proto_rawDescGZIP(), []int{1}
}

func (x *UploadStatus) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *UploadStatus) GetCode() UploadStatusCode {
	if x != nil {
		return x.Code
	}
	return UploadStatusCode_Ok
}

var File_upload_proto_upload_proto protoreflect.FileDescriptor

var file_upload_proto_upload_proto_rawDesc = []byte{
	0x0a, 0x19, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x72, 0x65,
	0x65, 0x74, 0x22, 0x21, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x53, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2b, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x2a, 0x26, 0x0a, 0x10, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06,
	0x0a, 0x02, 0x4f, 0x6b, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x10, 0x01, 0x32, 0x40, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0c, 0x2e,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x13, 0x2e, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x00, 0x28, 0x01, 0x42, 0x0e, 0x5a, 0x0c, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_upload_proto_upload_proto_rawDescOnce sync.Once
	file_upload_proto_upload_proto_rawDescData = file_upload_proto_upload_proto_rawDesc
)

func file_upload_proto_upload_proto_rawDescGZIP() []byte {
	file_upload_proto_upload_proto_rawDescOnce.Do(func() {
		file_upload_proto_upload_proto_rawDescData = protoimpl.X.CompressGZIP(file_upload_proto_upload_proto_rawDescData)
	})
	return file_upload_proto_upload_proto_rawDescData
}

var file_upload_proto_upload_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_upload_proto_upload_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_upload_proto_upload_proto_goTypes = []interface{}{
	(UploadStatusCode)(0), // 0: greet.UploadStatusCode
	(*Chunk)(nil),         // 1: greet.Chunk
	(*UploadStatus)(nil),  // 2: greet.UploadStatus
}
var file_upload_proto_upload_proto_depIdxs = []int32{
	0, // 0: greet.UploadStatus.code:type_name -> greet.UploadStatusCode
	1, // 1: greet.UploadService.Upload:input_type -> greet.Chunk
	2, // 2: greet.UploadService.Upload:output_type -> greet.UploadStatus
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_upload_proto_upload_proto_init() }
func file_upload_proto_upload_proto_init() {
	if File_upload_proto_upload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_upload_proto_upload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_upload_proto_upload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadStatus); i {
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
			RawDescriptor: file_upload_proto_upload_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_upload_proto_upload_proto_goTypes,
		DependencyIndexes: file_upload_proto_upload_proto_depIdxs,
		EnumInfos:         file_upload_proto_upload_proto_enumTypes,
		MessageInfos:      file_upload_proto_upload_proto_msgTypes,
	}.Build()
	File_upload_proto_upload_proto = out.File
	file_upload_proto_upload_proto_rawDesc = nil
	file_upload_proto_upload_proto_goTypes = nil
	file_upload_proto_upload_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UploadServiceClient is the client API for UploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UploadServiceClient interface {
	// uploading file by using client streaming
	Upload(ctx context.Context, opts ...grpc.CallOption) (UploadService_UploadClient, error)
}

type uploadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUploadServiceClient(cc grpc.ClientConnInterface) UploadServiceClient {
	return &uploadServiceClient{cc}
}

func (c *uploadServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (UploadService_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UploadService_serviceDesc.Streams[0], "/greet.UploadService/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &uploadServiceUploadClient{stream}
	return x, nil
}

type UploadService_UploadClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*UploadStatus, error)
	grpc.ClientStream
}

type uploadServiceUploadClient struct {
	grpc.ClientStream
}

func (x *uploadServiceUploadClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uploadServiceUploadClient) CloseAndRecv() (*UploadStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UploadServiceServer is the server API for UploadService service.
type UploadServiceServer interface {
	// uploading file by using client streaming
	Upload(UploadService_UploadServer) error
}

// UnimplementedUploadServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUploadServiceServer struct {
}

func (*UnimplementedUploadServiceServer) Upload(UploadService_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}

func RegisterUploadServiceServer(s *grpc.Server, srv UploadServiceServer) {
	s.RegisterService(&_UploadService_serviceDesc, srv)
}

func _UploadService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UploadServiceServer).Upload(&uploadServiceUploadServer{stream})
}

type UploadService_UploadServer interface {
	SendAndClose(*UploadStatus) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type uploadServiceUploadServer struct {
	grpc.ServerStream
}

func (x *uploadServiceUploadServer) SendAndClose(m *UploadStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uploadServiceUploadServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _UploadService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.UploadService",
	HandlerType: (*UploadServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _UploadService_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "upload/proto/upload.proto",
}