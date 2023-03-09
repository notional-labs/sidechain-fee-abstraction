// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: feeabstraction/absfee/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params defines the parameters for the feeabs module.
type MsgSendQueryIbcDenomTWAP struct {
	FromAddress string    `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	IbcDenom    string    `protobuf:"bytes,2,opt,name=ibc_denom,json=ibcDenom,proto3" json:"ibc_denom,omitempty"`
	StartTime   time.Time `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"start_time"`
}

func (m *MsgSendQueryIbcDenomTWAP) Reset()         { *m = MsgSendQueryIbcDenomTWAP{} }
func (m *MsgSendQueryIbcDenomTWAP) String() string { return proto.CompactTextString(m) }
func (*MsgSendQueryIbcDenomTWAP) ProtoMessage()    {}
func (*MsgSendQueryIbcDenomTWAP) Descriptor() ([]byte, []int) {
	return fileDescriptor_84c172c34645b936, []int{0}
}
func (m *MsgSendQueryIbcDenomTWAP) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendQueryIbcDenomTWAP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendQueryIbcDenomTWAP.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendQueryIbcDenomTWAP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendQueryIbcDenomTWAP.Merge(m, src)
}
func (m *MsgSendQueryIbcDenomTWAP) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendQueryIbcDenomTWAP) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendQueryIbcDenomTWAP.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendQueryIbcDenomTWAP proto.InternalMessageInfo

func (m *MsgSendQueryIbcDenomTWAP) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgSendQueryIbcDenomTWAP) GetIbcDenom() string {
	if m != nil {
		return m.IbcDenom
	}
	return ""
}

func (m *MsgSendQueryIbcDenomTWAP) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

type MsgSendQueryIbcDenomTWAPResponse struct {
}

func (m *MsgSendQueryIbcDenomTWAPResponse) Reset()         { *m = MsgSendQueryIbcDenomTWAPResponse{} }
func (m *MsgSendQueryIbcDenomTWAPResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSendQueryIbcDenomTWAPResponse) ProtoMessage()    {}
func (*MsgSendQueryIbcDenomTWAPResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_84c172c34645b936, []int{1}
}
func (m *MsgSendQueryIbcDenomTWAPResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendQueryIbcDenomTWAPResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendQueryIbcDenomTWAPResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendQueryIbcDenomTWAPResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendQueryIbcDenomTWAPResponse.Merge(m, src)
}
func (m *MsgSendQueryIbcDenomTWAPResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendQueryIbcDenomTWAPResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendQueryIbcDenomTWAPResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendQueryIbcDenomTWAPResponse proto.InternalMessageInfo

// Params defines the parameters for the feeabs module.
type MsgSwapCrossChain struct {
	FromAddress string `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	IbcDenom    string `protobuf:"bytes,2,opt,name=ibc_denom,json=ibcDenom,proto3" json:"ibc_denom,omitempty"`
}

func (m *MsgSwapCrossChain) Reset()         { *m = MsgSwapCrossChain{} }
func (m *MsgSwapCrossChain) String() string { return proto.CompactTextString(m) }
func (*MsgSwapCrossChain) ProtoMessage()    {}
func (*MsgSwapCrossChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_84c172c34645b936, []int{2}
}
func (m *MsgSwapCrossChain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapCrossChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapCrossChain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapCrossChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapCrossChain.Merge(m, src)
}
func (m *MsgSwapCrossChain) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapCrossChain) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapCrossChain.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapCrossChain proto.InternalMessageInfo

func (m *MsgSwapCrossChain) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgSwapCrossChain) GetIbcDenom() string {
	if m != nil {
		return m.IbcDenom
	}
	return ""
}

type MsgSwapCrossChainResponse struct {
}

func (m *MsgSwapCrossChainResponse) Reset()         { *m = MsgSwapCrossChainResponse{} }
func (m *MsgSwapCrossChainResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSwapCrossChainResponse) ProtoMessage()    {}
func (*MsgSwapCrossChainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_84c172c34645b936, []int{3}
}
func (m *MsgSwapCrossChainResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapCrossChainResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapCrossChainResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapCrossChainResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapCrossChainResponse.Merge(m, src)
}
func (m *MsgSwapCrossChainResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapCrossChainResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapCrossChainResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapCrossChainResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSendQueryIbcDenomTWAP)(nil), "feeabstraction.absfee.v1beta1.MsgSendQueryIbcDenomTWAP")
	proto.RegisterType((*MsgSendQueryIbcDenomTWAPResponse)(nil), "feeabstraction.absfee.v1beta1.MsgSendQueryIbcDenomTWAPResponse")
	proto.RegisterType((*MsgSwapCrossChain)(nil), "feeabstraction.absfee.v1beta1.MsgSwapCrossChain")
	proto.RegisterType((*MsgSwapCrossChainResponse)(nil), "feeabstraction.absfee.v1beta1.MsgSwapCrossChainResponse")
}

func init() {
	proto.RegisterFile("feeabstraction/absfee/v1beta1/tx.proto", fileDescriptor_84c172c34645b936)
}

var fileDescriptor_84c172c34645b936 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcb, 0x6e, 0x13, 0x31,
	0x14, 0x86, 0xe3, 0x56, 0x42, 0xc4, 0x45, 0x48, 0x1d, 0x81, 0x14, 0xa6, 0x74, 0x12, 0x66, 0x81,
	0x0a, 0x12, 0x36, 0x29, 0x0b, 0x2e, 0x1b, 0xd4, 0x96, 0x0d, 0x8b, 0x4a, 0x90, 0x54, 0x02, 0xb1,
	0x89, 0x8e, 0x27, 0xce, 0x64, 0xa4, 0x19, 0xdb, 0xcc, 0x71, 0x4a, 0x23, 0xf1, 0x10, 0x65, 0xc5,
	0x8b, 0xf0, 0x10, 0x5d, 0x76, 0xc9, 0xaa, 0xa0, 0xe4, 0x0d, 0x78, 0x02, 0xe4, 0xb9, 0x80, 0x5a,
	0x48, 0xc4, 0x65, 0x37, 0x73, 0xfc, 0xfd, 0x17, 0x5b, 0x87, 0xde, 0x1e, 0x49, 0x09, 0x02, 0x6d,
	0x0e, 0x91, 0x4d, 0xb4, 0xe2, 0x20, 0x70, 0x24, 0x25, 0x3f, 0xec, 0x0a, 0x69, 0xa1, 0xcb, 0xed,
	0x11, 0x33, 0xb9, 0xb6, 0xda, 0xdb, 0x3c, 0xcf, 0xb1, 0x92, 0x63, 0x15, 0xe7, 0x5f, 0x8b, 0x75,
	0xac, 0x0b, 0x92, 0xbb, 0xaf, 0x52, 0xe4, 0xdf, 0x8c, 0xb5, 0x8e, 0x53, 0xc9, 0xc1, 0x24, 0x1c,
	0x94, 0xd2, 0x16, 0x9c, 0x16, 0xab, 0xd3, 0xbb, 0x91, 0xc6, 0x4c, 0x23, 0x17, 0x80, 0x92, 0xbf,
	0x9d, 0xc8, 0x7c, 0xfa, 0x23, 0xd6, 0x40, 0x9c, 0xa8, 0x02, 0xae, 0xd9, 0xe5, 0x35, 0x0d, 0xe4,
	0x90, 0xd5, 0xbe, 0x77, 0x96, 0xb3, 0xd2, 0xe8, 0x68, 0x5c, 0xa1, 0xed, 0xaa, 0x60, 0xf1, 0x27,
	0x26, 0x23, 0x6e, 0x93, 0x4c, 0xa2, 0x85, 0xcc, 0x94, 0x40, 0xf8, 0x89, 0xd0, 0xd6, 0x3e, 0xc6,
	0x7d, 0xa9, 0x86, 0x2f, 0x5d, 0xc3, 0xe7, 0x22, 0x7a, 0x26, 0x95, 0xce, 0x0e, 0x5e, 0xed, 0xbc,
	0xf0, 0x6e, 0xd1, 0x2b, 0xa3, 0x5c, 0x67, 0x03, 0x18, 0x0e, 0x73, 0x89, 0xd8, 0x22, 0x1d, 0xb2,
	0xd5, 0xec, 0xad, 0xb9, 0xd9, 0x4e, 0x39, 0xf2, 0x36, 0x68, 0x33, 0x11, 0xd1, 0x60, 0xe8, 0x34,
	0xad, 0x95, 0xe2, 0xfc, 0x72, 0x52, 0x79, 0x78, 0xaf, 0x29, 0x45, 0x0b, 0xb9, 0x1d, 0xb8, 0xd4,
	0xd6, 0x6a, 0x87, 0x6c, 0xad, 0x6d, 0xfb, 0xac, 0xac, 0xc4, 0xea, 0x4a, 0xec, 0xa0, 0xae, 0xb4,
	0xbb, 0x79, 0x72, 0xd6, 0x6e, 0x7c, 0x3b, 0x6b, 0xaf, 0x4f, 0x21, 0x4b, 0x9f, 0x84, 0x3f, 0xb5,
	0xe1, 0xf1, 0x97, 0x36, 0xe9, 0x35, 0x8b, 0x81, 0xc3, 0xc3, 0x90, 0x76, 0x16, 0xb5, 0xee, 0x49,
	0x34, 0x5a, 0xa1, 0x0c, 0xfb, 0x74, 0xdd, 0x31, 0xef, 0xc0, 0xec, 0xe5, 0x1a, 0x71, 0x6f, 0x0c,
	0x89, 0xfa, 0xdf, 0x2b, 0x85, 0x1b, 0xf4, 0xc6, 0x2f, 0xa6, 0x75, 0xe2, 0xf6, 0xc7, 0x15, 0xba,
	0xba, 0x8f, 0xb1, 0xf7, 0x81, 0xd0, 0xeb, 0xbf, 0x7f, 0xd1, 0x87, 0x6c, 0xe9, 0x9a, 0xb1, 0x45,
	0x97, 0xf2, 0x9f, 0xfe, 0xa3, 0xb0, 0xee, 0xe6, 0xbd, 0xa7, 0x57, 0x2f, 0x3c, 0xc5, 0xfd, 0x3f,
	0xb0, 0x3c, 0xa7, 0xf0, 0x1f, 0xfd, 0xad, 0xa2, 0x4e, 0xdf, 0xed, 0x9f, 0xcc, 0x02, 0x72, 0x3a,
	0x0b, 0xc8, 0xd7, 0x59, 0x40, 0x8e, 0xe7, 0x41, 0xe3, 0x74, 0x1e, 0x34, 0x3e, 0xcf, 0x83, 0xc6,
	0x9b, 0xc7, 0x71, 0x62, 0xc7, 0x13, 0xc1, 0x22, 0x9d, 0x71, 0xa5, 0x9d, 0x2b, 0xa4, 0xf7, 0x52,
	0x10, 0xc8, 0x2f, 0x6c, 0xf9, 0x61, 0x97, 0x1f, 0x55, 0x33, 0x6e, 0xa7, 0x46, 0xa2, 0xb8, 0x54,
	0xac, 0xd0, 0x83, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x8d, 0xe9, 0xf2, 0xe3, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	SendQueryIbcDenomTWAP(ctx context.Context, in *MsgSendQueryIbcDenomTWAP, opts ...grpc.CallOption) (*MsgSendQueryIbcDenomTWAPResponse, error)
	SwapCrossChain(ctx context.Context, in *MsgSwapCrossChain, opts ...grpc.CallOption) (*MsgSwapCrossChainResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SendQueryIbcDenomTWAP(ctx context.Context, in *MsgSendQueryIbcDenomTWAP, opts ...grpc.CallOption) (*MsgSendQueryIbcDenomTWAPResponse, error) {
	out := new(MsgSendQueryIbcDenomTWAPResponse)
	err := c.cc.Invoke(ctx, "/feeabstraction.absfee.v1beta1.Msg/SendQueryIbcDenomTWAP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SwapCrossChain(ctx context.Context, in *MsgSwapCrossChain, opts ...grpc.CallOption) (*MsgSwapCrossChainResponse, error) {
	out := new(MsgSwapCrossChainResponse)
	err := c.cc.Invoke(ctx, "/feeabstraction.absfee.v1beta1.Msg/SwapCrossChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SendQueryIbcDenomTWAP(context.Context, *MsgSendQueryIbcDenomTWAP) (*MsgSendQueryIbcDenomTWAPResponse, error)
	SwapCrossChain(context.Context, *MsgSwapCrossChain) (*MsgSwapCrossChainResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SendQueryIbcDenomTWAP(ctx context.Context, req *MsgSendQueryIbcDenomTWAP) (*MsgSendQueryIbcDenomTWAPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendQueryIbcDenomTWAP not implemented")
}
func (*UnimplementedMsgServer) SwapCrossChain(ctx context.Context, req *MsgSwapCrossChain) (*MsgSwapCrossChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwapCrossChain not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SendQueryIbcDenomTWAP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendQueryIbcDenomTWAP)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendQueryIbcDenomTWAP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feeabstraction.absfee.v1beta1.Msg/SendQueryIbcDenomTWAP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendQueryIbcDenomTWAP(ctx, req.(*MsgSendQueryIbcDenomTWAP))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SwapCrossChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSwapCrossChain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SwapCrossChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feeabstraction.absfee.v1beta1.Msg/SwapCrossChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SwapCrossChain(ctx, req.(*MsgSwapCrossChain))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "feeabstraction.absfee.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendQueryIbcDenomTWAP",
			Handler:    _Msg_SendQueryIbcDenomTWAP_Handler,
		},
		{
			MethodName: "SwapCrossChain",
			Handler:    _Msg_SwapCrossChain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feeabstraction/absfee/v1beta1/tx.proto",
}

func (m *MsgSendQueryIbcDenomTWAP) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendQueryIbcDenomTWAP) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendQueryIbcDenomTWAP) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTx(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	if len(m.IbcDenom) > 0 {
		i -= len(m.IbcDenom)
		copy(dAtA[i:], m.IbcDenom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.IbcDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSendQueryIbcDenomTWAPResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendQueryIbcDenomTWAPResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendQueryIbcDenomTWAPResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSwapCrossChain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapCrossChain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapCrossChain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IbcDenom) > 0 {
		i -= len(m.IbcDenom)
		copy(dAtA[i:], m.IbcDenom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.IbcDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSwapCrossChainResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapCrossChainResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapCrossChainResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSendQueryIbcDenomTWAP) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.IbcDenom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgSendQueryIbcDenomTWAPResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSwapCrossChain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.IbcDenom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSwapCrossChainResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSendQueryIbcDenomTWAP) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSendQueryIbcDenomTWAP: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendQueryIbcDenomTWAP: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSendQueryIbcDenomTWAPResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSendQueryIbcDenomTWAPResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendQueryIbcDenomTWAPResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSwapCrossChain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSwapCrossChain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapCrossChain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSwapCrossChainResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSwapCrossChainResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapCrossChainResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
