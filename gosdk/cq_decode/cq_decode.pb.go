// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: proto/cq_decode/cq_decode.proto

package cq_decode

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

type Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Value uint32 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Output) Reset() {
	*x = Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cq_decode_cq_decode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cq_decode_cq_decode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_proto_cq_decode_cq_decode_proto_rawDescGZIP(), []int{0}
}

func (x *Output) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Output) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

//设置电磁阀输出请求
type SetOutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` //4个字节高两位是通讯口端口号,低２位为485站号
	Outputs []*Output `protobuf:"bytes,2,rep,name=outputs,proto3" json:"outputs,omitempty"`
}

func (x *SetOutRequest) Reset() {
	*x = SetOutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cq_decode_cq_decode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetOutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetOutRequest) ProtoMessage() {}

func (x *SetOutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cq_decode_cq_decode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetOutRequest.ProtoReflect.Descriptor instead.
func (*SetOutRequest) Descriptor() ([]byte, []int) {
	return file_proto_cq_decode_cq_decode_proto_rawDescGZIP(), []int{1}
}

func (x *SetOutRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SetOutRequest) GetOutputs() []*Output {
	if x != nil {
		return x.Outputs
	}
	return nil
}

//设置电磁阀输出响应
type SetOutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetOutResponse) Reset() {
	*x = SetOutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cq_decode_cq_decode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetOutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetOutResponse) ProtoMessage() {}

func (x *SetOutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cq_decode_cq_decode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetOutResponse.ProtoReflect.Descriptor instead.
func (*SetOutResponse) Descriptor() ([]byte, []int) {
	return file_proto_cq_decode_cq_decode_proto_rawDescGZIP(), []int{2}
}

//电磁阀状态获取请求
type StateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StateRequest) Reset() {
	*x = StateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cq_decode_cq_decode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateRequest) ProtoMessage() {}

func (x *StateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cq_decode_cq_decode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateRequest.ProtoReflect.Descriptor instead.
func (*StateRequest) Descriptor() ([]byte, []int) {
	return file_proto_cq_decode_cq_decode_proto_rawDescGZIP(), []int{3}
}

//电磁阀状态获取返回
type StateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` //4个字节高两位是通讯口端口号,低２位为485站号
	Outputs []*Output `protobuf:"bytes,2,rep,name=outputs,proto3" json:"outputs,omitempty"`
	Volt    float32   `protobuf:"fixed32,3,opt,name=volt,proto3" json:"volt,omitempty"` //电压值
}

func (x *StateResponse) Reset() {
	*x = StateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cq_decode_cq_decode_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateResponse) ProtoMessage() {}

func (x *StateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cq_decode_cq_decode_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateResponse.ProtoReflect.Descriptor instead.
func (*StateResponse) Descriptor() ([]byte, []int) {
	return file_proto_cq_decode_cq_decode_proto_rawDescGZIP(), []int{4}
}

func (x *StateResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StateResponse) GetOutputs() []*Output {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *StateResponse) GetVolt() float32 {
	if x != nil {
		return x.Volt
	}
	return 0
}

var File_proto_cq_decode_cq_decode_proto protoreflect.FileDescriptor

var file_proto_cq_decode_cq_decode_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x71, 0x5f, 0x64, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x2f, 0x63, 0x71, 0x5f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x63, 0x71, 0x5f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x2e, 0x0a, 0x06,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4c, 0x0a, 0x0d,
	0x53, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2b, 0x0a,
	0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x63, 0x71, 0x5f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x65,
	0x74, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e, 0x0a, 0x0c,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x60, 0x0a, 0x0d,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2b, 0x0a,
	0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x63, 0x71, 0x5f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x6f,
	0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x76, 0x6f, 0x6c, 0x74, 0x32, 0x8c,
	0x01, 0x0a, 0x06, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x53, 0x65, 0x74,
	0x4f, 0x75, 0x74, 0x12, 0x18, 0x2e, 0x63, 0x71, 0x5f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x53, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x63, 0x71, 0x5f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x4f, 0x75, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x63, 0x71, 0x5f, 0x64, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x63, 0x71, 0x5f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x0d, 0x5a,
	0x0b, 0x2e, 0x3b, 0x63, 0x71, 0x5f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_cq_decode_cq_decode_proto_rawDescOnce sync.Once
	file_proto_cq_decode_cq_decode_proto_rawDescData = file_proto_cq_decode_cq_decode_proto_rawDesc
)

func file_proto_cq_decode_cq_decode_proto_rawDescGZIP() []byte {
	file_proto_cq_decode_cq_decode_proto_rawDescOnce.Do(func() {
		file_proto_cq_decode_cq_decode_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cq_decode_cq_decode_proto_rawDescData)
	})
	return file_proto_cq_decode_cq_decode_proto_rawDescData
}

var file_proto_cq_decode_cq_decode_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_cq_decode_cq_decode_proto_goTypes = []interface{}{
	(*Output)(nil),         // 0: cq_decode.Output
	(*SetOutRequest)(nil),  // 1: cq_decode.SetOutRequest
	(*SetOutResponse)(nil), // 2: cq_decode.SetOutResponse
	(*StateRequest)(nil),   // 3: cq_decode.StateRequest
	(*StateResponse)(nil),  // 4: cq_decode.StateResponse
}
var file_proto_cq_decode_cq_decode_proto_depIdxs = []int32{
	0, // 0: cq_decode.SetOutRequest.outputs:type_name -> cq_decode.Output
	0, // 1: cq_decode.StateResponse.outputs:type_name -> cq_decode.Output
	1, // 2: cq_decode.Decode.SetOut:input_type -> cq_decode.SetOutRequest
	3, // 3: cq_decode.Decode.GetState:input_type -> cq_decode.StateRequest
	2, // 4: cq_decode.Decode.SetOut:output_type -> cq_decode.SetOutResponse
	4, // 5: cq_decode.Decode.GetState:output_type -> cq_decode.StateResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_cq_decode_cq_decode_proto_init() }
func file_proto_cq_decode_cq_decode_proto_init() {
	if File_proto_cq_decode_cq_decode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_cq_decode_cq_decode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Output); i {
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
		file_proto_cq_decode_cq_decode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetOutRequest); i {
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
		file_proto_cq_decode_cq_decode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetOutResponse); i {
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
		file_proto_cq_decode_cq_decode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateRequest); i {
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
		file_proto_cq_decode_cq_decode_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateResponse); i {
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
			RawDescriptor: file_proto_cq_decode_cq_decode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cq_decode_cq_decode_proto_goTypes,
		DependencyIndexes: file_proto_cq_decode_cq_decode_proto_depIdxs,
		MessageInfos:      file_proto_cq_decode_cq_decode_proto_msgTypes,
	}.Build()
	File_proto_cq_decode_cq_decode_proto = out.File
	file_proto_cq_decode_cq_decode_proto_rawDesc = nil
	file_proto_cq_decode_cq_decode_proto_goTypes = nil
	file_proto_cq_decode_cq_decode_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DecodeClient is the client API for Decode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DecodeClient interface {
	//设置解码器输出
	SetOut(ctx context.Context, in *SetOutRequest, opts ...grpc.CallOption) (*SetOutResponse, error)
	//解码器状态获取
	GetState(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (Decode_GetStateClient, error)
}

type decodeClient struct {
	cc grpc.ClientConnInterface
}

func NewDecodeClient(cc grpc.ClientConnInterface) DecodeClient {
	return &decodeClient{cc}
}

func (c *decodeClient) SetOut(ctx context.Context, in *SetOutRequest, opts ...grpc.CallOption) (*SetOutResponse, error) {
	out := new(SetOutResponse)
	err := c.cc.Invoke(ctx, "/cq_decode.Decode/SetOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *decodeClient) GetState(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (Decode_GetStateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Decode_serviceDesc.Streams[0], "/cq_decode.Decode/GetState", opts...)
	if err != nil {
		return nil, err
	}
	x := &decodeGetStateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Decode_GetStateClient interface {
	Recv() (*StateResponse, error)
	grpc.ClientStream
}

type decodeGetStateClient struct {
	grpc.ClientStream
}

func (x *decodeGetStateClient) Recv() (*StateResponse, error) {
	m := new(StateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DecodeServer is the server API for Decode service.
type DecodeServer interface {
	//设置解码器输出
	SetOut(context.Context, *SetOutRequest) (*SetOutResponse, error)
	//解码器状态获取
	GetState(*StateRequest, Decode_GetStateServer) error
}

// UnimplementedDecodeServer can be embedded to have forward compatible implementations.
type UnimplementedDecodeServer struct {
}

func (*UnimplementedDecodeServer) SetOut(context.Context, *SetOutRequest) (*SetOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetOut not implemented")
}
func (*UnimplementedDecodeServer) GetState(*StateRequest, Decode_GetStateServer) error {
	return status.Errorf(codes.Unimplemented, "method GetState not implemented")
}

func RegisterDecodeServer(s *grpc.Server, srv DecodeServer) {
	s.RegisterService(&_Decode_serviceDesc, srv)
}

func _Decode_SetOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DecodeServer).SetOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cq_decode.Decode/SetOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DecodeServer).SetOut(ctx, req.(*SetOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Decode_GetState_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DecodeServer).GetState(m, &decodeGetStateServer{stream})
}

type Decode_GetStateServer interface {
	Send(*StateResponse) error
	grpc.ServerStream
}

type decodeGetStateServer struct {
	grpc.ServerStream
}

func (x *decodeGetStateServer) Send(m *StateResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Decode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cq_decode.Decode",
	HandlerType: (*DecodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetOut",
			Handler:    _Decode_SetOut_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetState",
			Handler:       _Decode_GetState_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/cq_decode/cq_decode.proto",
}
