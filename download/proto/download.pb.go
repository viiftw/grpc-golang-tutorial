// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: download/proto/download.proto

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

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_download_proto_download_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_download_proto_download_proto_msgTypes[0]
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
	return file_download_proto_download_proto_rawDescGZIP(), []int{0}
}

func (x *Chunk) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type FileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
}

func (x *FileRequest) Reset() {
	*x = FileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_download_proto_download_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRequest) ProtoMessage() {}

func (x *FileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_download_proto_download_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRequest.ProtoReflect.Descriptor instead.
func (*FileRequest) Descriptor() ([]byte, []int) {
	return file_download_proto_download_proto_rawDescGZIP(), []int{1}
}

func (x *FileRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

var File_download_proto_download_proto protoreflect.FileDescriptor

var file_download_proto_download_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x67, 0x72, 0x65, 0x65, 0x74, 0x22, 0x1d, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x2a, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x32, 0x47, 0x0a, 0x0f, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x00, 0x30, 0x01, 0x42, 0x10, 0x5a, 0x0e, 0x64, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_download_proto_download_proto_rawDescOnce sync.Once
	file_download_proto_download_proto_rawDescData = file_download_proto_download_proto_rawDesc
)

func file_download_proto_download_proto_rawDescGZIP() []byte {
	file_download_proto_download_proto_rawDescOnce.Do(func() {
		file_download_proto_download_proto_rawDescData = protoimpl.X.CompressGZIP(file_download_proto_download_proto_rawDescData)
	})
	return file_download_proto_download_proto_rawDescData
}

var file_download_proto_download_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_download_proto_download_proto_goTypes = []interface{}{
	(*Chunk)(nil),       // 0: greet.Chunk
	(*FileRequest)(nil), // 1: greet.FileRequest
}
var file_download_proto_download_proto_depIdxs = []int32{
	1, // 0: greet.DownloadService.DownloadFile:input_type -> greet.FileRequest
	0, // 1: greet.DownloadService.DownloadFile:output_type -> greet.Chunk
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_download_proto_download_proto_init() }
func file_download_proto_download_proto_init() {
	if File_download_proto_download_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_download_proto_download_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_download_proto_download_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRequest); i {
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
			RawDescriptor: file_download_proto_download_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_download_proto_download_proto_goTypes,
		DependencyIndexes: file_download_proto_download_proto_depIdxs,
		MessageInfos:      file_download_proto_download_proto_msgTypes,
	}.Build()
	File_download_proto_download_proto = out.File
	file_download_proto_download_proto_rawDesc = nil
	file_download_proto_download_proto_goTypes = nil
	file_download_proto_download_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DownloadServiceClient is the client API for DownloadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DownloadServiceClient interface {
	// Server Streaming file
	DownloadFile(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (DownloadService_DownloadFileClient, error)
}

type downloadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDownloadServiceClient(cc grpc.ClientConnInterface) DownloadServiceClient {
	return &downloadServiceClient{cc}
}

func (c *downloadServiceClient) DownloadFile(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (DownloadService_DownloadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DownloadService_serviceDesc.Streams[0], "/greet.DownloadService/DownloadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &downloadServiceDownloadFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DownloadService_DownloadFileClient interface {
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type downloadServiceDownloadFileClient struct {
	grpc.ClientStream
}

func (x *downloadServiceDownloadFileClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DownloadServiceServer is the server API for DownloadService service.
type DownloadServiceServer interface {
	// Server Streaming file
	DownloadFile(*FileRequest, DownloadService_DownloadFileServer) error
}

// UnimplementedDownloadServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDownloadServiceServer struct {
}

func (*UnimplementedDownloadServiceServer) DownloadFile(*FileRequest, DownloadService_DownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}

func RegisterDownloadServiceServer(s *grpc.Server, srv DownloadServiceServer) {
	s.RegisterService(&_DownloadService_serviceDesc, srv)
}

func _DownloadService_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DownloadServiceServer).DownloadFile(m, &downloadServiceDownloadFileServer{stream})
}

type DownloadService_DownloadFileServer interface {
	Send(*Chunk) error
	grpc.ServerStream
}

type downloadServiceDownloadFileServer struct {
	grpc.ServerStream
}

func (x *downloadServiceDownloadFileServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

var _DownloadService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.DownloadService",
	HandlerType: (*DownloadServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DownloadFile",
			Handler:       _DownloadService_DownloadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "download/proto/download.proto",
}
