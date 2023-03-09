// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: feeabstraction/absfee/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QueryHostChainConfigRequest struct {
	IbcDenom string `protobuf:"bytes,1,opt,name=ibc_denom,json=ibcDenom,proto3" json:"ibc_denom,omitempty"`
}

func (m *QueryHostChainConfigRequest) Reset()         { *m = QueryHostChainConfigRequest{} }
func (m *QueryHostChainConfigRequest) String() string { return proto.CompactTextString(m) }
func (*QueryHostChainConfigRequest) ProtoMessage()    {}
func (*QueryHostChainConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a3dde61db3cbb0e, []int{0}
}
func (m *QueryHostChainConfigRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryHostChainConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryHostChainConfigRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryHostChainConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryHostChainConfigRequest.Merge(m, src)
}
func (m *QueryHostChainConfigRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryHostChainConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryHostChainConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryHostChainConfigRequest proto.InternalMessageInfo

func (m *QueryHostChainConfigRequest) GetIbcDenom() string {
	if m != nil {
		return m.IbcDenom
	}
	return ""
}

type QueryHostChainConfigRespone struct {
	HostChainConfig HostChainFeeAbsConfig `protobuf:"bytes,1,opt,name=host_chain_config,json=hostChainConfig,proto3" json:"host_chain_config" yaml:"host_chain_config"`
}

func (m *QueryHostChainConfigRespone) Reset()         { *m = QueryHostChainConfigRespone{} }
func (m *QueryHostChainConfigRespone) String() string { return proto.CompactTextString(m) }
func (*QueryHostChainConfigRespone) ProtoMessage()    {}
func (*QueryHostChainConfigRespone) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a3dde61db3cbb0e, []int{1}
}
func (m *QueryHostChainConfigRespone) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryHostChainConfigRespone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryHostChainConfigRespone.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryHostChainConfigRespone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryHostChainConfigRespone.Merge(m, src)
}
func (m *QueryHostChainConfigRespone) XXX_Size() int {
	return m.Size()
}
func (m *QueryHostChainConfigRespone) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryHostChainConfigRespone.DiscardUnknown(m)
}

var xxx_messageInfo_QueryHostChainConfigRespone proto.InternalMessageInfo

func (m *QueryHostChainConfigRespone) GetHostChainConfig() HostChainFeeAbsConfig {
	if m != nil {
		return m.HostChainConfig
	}
	return HostChainFeeAbsConfig{}
}

// QueryOsmosisSpotPriceRequest is the request type for the Query/Feeabs RPC method.
type QueryOsmosisArithmeticTwapRequest struct {
	IbcDenom string `protobuf:"bytes,1,opt,name=ibc_denom,json=ibcDenom,proto3" json:"ibc_denom,omitempty"`
}

func (m *QueryOsmosisArithmeticTwapRequest) Reset()         { *m = QueryOsmosisArithmeticTwapRequest{} }
func (m *QueryOsmosisArithmeticTwapRequest) String() string { return proto.CompactTextString(m) }
func (*QueryOsmosisArithmeticTwapRequest) ProtoMessage()    {}
func (*QueryOsmosisArithmeticTwapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a3dde61db3cbb0e, []int{2}
}
func (m *QueryOsmosisArithmeticTwapRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryOsmosisArithmeticTwapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryOsmosisArithmeticTwapRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryOsmosisArithmeticTwapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryOsmosisArithmeticTwapRequest.Merge(m, src)
}
func (m *QueryOsmosisArithmeticTwapRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryOsmosisArithmeticTwapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryOsmosisArithmeticTwapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryOsmosisArithmeticTwapRequest proto.InternalMessageInfo

func (m *QueryOsmosisArithmeticTwapRequest) GetIbcDenom() string {
	if m != nil {
		return m.IbcDenom
	}
	return ""
}

type QueryOsmosisArithmeticTwapResponse struct {
	ArithmeticTwap github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=arithmetic_twap,json=arithmeticTwap,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"arithmetic_twap" yaml:"arithmetic_twap"`
}

func (m *QueryOsmosisArithmeticTwapResponse) Reset()         { *m = QueryOsmosisArithmeticTwapResponse{} }
func (m *QueryOsmosisArithmeticTwapResponse) String() string { return proto.CompactTextString(m) }
func (*QueryOsmosisArithmeticTwapResponse) ProtoMessage()    {}
func (*QueryOsmosisArithmeticTwapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a3dde61db3cbb0e, []int{3}
}
func (m *QueryOsmosisArithmeticTwapResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryOsmosisArithmeticTwapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryOsmosisArithmeticTwapResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryOsmosisArithmeticTwapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryOsmosisArithmeticTwapResponse.Merge(m, src)
}
func (m *QueryOsmosisArithmeticTwapResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryOsmosisArithmeticTwapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryOsmosisArithmeticTwapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryOsmosisArithmeticTwapResponse proto.InternalMessageInfo

// QueryFeeabsModuleBalacesRequest is the request type for the Query/Feeabs RPC method.
type QueryFeeabsModuleBalacesRequest struct {
}

func (m *QueryFeeabsModuleBalacesRequest) Reset()         { *m = QueryFeeabsModuleBalacesRequest{} }
func (m *QueryFeeabsModuleBalacesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryFeeabsModuleBalacesRequest) ProtoMessage()    {}
func (*QueryFeeabsModuleBalacesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a3dde61db3cbb0e, []int{4}
}
func (m *QueryFeeabsModuleBalacesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFeeabsModuleBalacesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFeeabsModuleBalacesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFeeabsModuleBalacesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFeeabsModuleBalacesRequest.Merge(m, src)
}
func (m *QueryFeeabsModuleBalacesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryFeeabsModuleBalacesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFeeabsModuleBalacesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFeeabsModuleBalacesRequest proto.InternalMessageInfo

type QueryFeeabsModuleBalacesResponse struct {
	Balances github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=balances,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"balances" yaml:"balances"`
}

func (m *QueryFeeabsModuleBalacesResponse) Reset()         { *m = QueryFeeabsModuleBalacesResponse{} }
func (m *QueryFeeabsModuleBalacesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryFeeabsModuleBalacesResponse) ProtoMessage()    {}
func (*QueryFeeabsModuleBalacesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a3dde61db3cbb0e, []int{5}
}
func (m *QueryFeeabsModuleBalacesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFeeabsModuleBalacesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFeeabsModuleBalacesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFeeabsModuleBalacesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFeeabsModuleBalacesResponse.Merge(m, src)
}
func (m *QueryFeeabsModuleBalacesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryFeeabsModuleBalacesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFeeabsModuleBalacesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFeeabsModuleBalacesResponse proto.InternalMessageInfo

func (m *QueryFeeabsModuleBalacesResponse) GetBalances() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Balances
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryHostChainConfigRequest)(nil), "feeabstraction.absfee.v1beta1.QueryHostChainConfigRequest")
	proto.RegisterType((*QueryHostChainConfigRespone)(nil), "feeabstraction.absfee.v1beta1.QueryHostChainConfigRespone")
	proto.RegisterType((*QueryOsmosisArithmeticTwapRequest)(nil), "feeabstraction.absfee.v1beta1.QueryOsmosisArithmeticTwapRequest")
	proto.RegisterType((*QueryOsmosisArithmeticTwapResponse)(nil), "feeabstraction.absfee.v1beta1.QueryOsmosisArithmeticTwapResponse")
	proto.RegisterType((*QueryFeeabsModuleBalacesRequest)(nil), "feeabstraction.absfee.v1beta1.QueryFeeabsModuleBalacesRequest")
	proto.RegisterType((*QueryFeeabsModuleBalacesResponse)(nil), "feeabstraction.absfee.v1beta1.QueryFeeabsModuleBalacesResponse")
}

func init() {
	proto.RegisterFile("feeabstraction/absfee/v1beta1/query.proto", fileDescriptor_0a3dde61db3cbb0e)
}

var fileDescriptor_0a3dde61db3cbb0e = []byte{
	// 655 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4f, 0x4f, 0x13, 0x4f,
	0x18, 0xee, 0xfe, 0x7e, 0xc1, 0xc0, 0x90, 0xd8, 0x38, 0x41, 0x83, 0x45, 0xb7, 0x65, 0x0e, 0x06,
	0x89, 0xbb, 0x13, 0xaa, 0x07, 0xed, 0x01, 0x4b, 0x4b, 0x08, 0x17, 0x63, 0xac, 0x9e, 0xbc, 0x34,
	0xb3, 0xc3, 0xb4, 0x9d, 0xb8, 0xbb, 0xb3, 0x74, 0xa6, 0x20, 0x1a, 0x2f, 0x1c, 0x3d, 0x99, 0x78,
	0xf0, 0xe6, 0xc1, 0xa3, 0x9f, 0x84, 0xc4, 0x0b, 0x89, 0x17, 0xe3, 0xa1, 0x1a, 0xe0, 0x13, 0xf0,
	0x09, 0xcc, 0xcc, 0x6c, 0x0b, 0x54, 0x5a, 0x1a, 0x3c, 0xf5, 0xcf, 0x3c, 0xef, 0x33, 0xcf, 0xf3,
	0x3e, 0xef, 0xbc, 0xe0, 0x6e, 0x83, 0x31, 0x12, 0x48, 0xd5, 0x26, 0x54, 0x71, 0x11, 0x63, 0x12,
	0xc8, 0x06, 0x63, 0x78, 0x6b, 0x29, 0x60, 0x8a, 0x2c, 0xe1, 0xcd, 0x0e, 0x6b, 0xef, 0xf8, 0x49,
	0x5b, 0x28, 0x01, 0x6f, 0x9f, 0x85, 0xfa, 0x16, 0xea, 0xa7, 0xd0, 0xdc, 0x4c, 0x53, 0x34, 0x85,
	0x41, 0x62, 0xfd, 0xcd, 0x16, 0xe5, 0x6e, 0x35, 0x85, 0x68, 0x86, 0x0c, 0x93, 0x84, 0x63, 0x12,
	0xc7, 0x42, 0x11, 0x5d, 0x2b, 0xd3, 0xd3, 0xc5, 0xd1, 0xb7, 0x27, 0xa4, 0x4d, 0xa2, 0x1e, 0xd6,
	0x1f, 0x8d, 0x15, 0x32, 0x12, 0x92, 0x4b, 0x1e, 0xd0, 0x14, 0xef, 0x52, 0xf3, 0x17, 0x0e, 0x88,
	0x3c, 0x41, 0x51, 0xc1, 0x63, 0x7b, 0x8e, 0x4a, 0x60, 0xee, 0x99, 0x76, 0xb7, 0x2e, 0xa4, 0xaa,
	0xb6, 0x08, 0x8f, 0xab, 0x22, 0x6e, 0xf0, 0x66, 0x8d, 0x6d, 0x76, 0x98, 0x54, 0x70, 0x0e, 0x4c,
	0xf1, 0x80, 0xd6, 0x37, 0x58, 0x2c, 0xa2, 0x59, 0xa7, 0xe0, 0x2c, 0x4c, 0xd5, 0x26, 0x79, 0x40,
	0x57, 0xf5, 0x6f, 0xf4, 0xc5, 0x19, 0x56, 0x2c, 0x13, 0x11, 0x33, 0xb8, 0xeb, 0x80, 0x6b, 0x2d,
	0x21, 0x55, 0x9d, 0xea, 0xb3, 0x3a, 0x35, 0x87, 0x86, 0x65, 0xba, 0xf8, 0xc0, 0x1f, 0xd9, 0x47,
	0xbf, 0x4f, 0xb9, 0xc6, 0xd8, 0x4a, 0x20, 0x2d, 0x71, 0xa5, 0xb0, 0xd7, 0xcd, 0x67, 0x8e, 0xbb,
	0xf9, 0xd9, 0x1d, 0x12, 0x85, 0x25, 0xf4, 0x17, 0x39, 0xaa, 0x65, 0x5b, 0x67, 0xb5, 0xa0, 0x32,
	0x98, 0x37, 0x1a, 0x9f, 0xda, 0xce, 0xac, 0xb4, 0xb9, 0x6a, 0x45, 0x4c, 0x71, 0xfa, 0x62, 0x9b,
	0x24, 0x63, 0xd9, 0xfc, 0xe4, 0x00, 0x34, 0x8a, 0x42, 0x9b, 0x95, 0x0c, 0x6e, 0x82, 0x2c, 0xe9,
	0x9f, 0xd4, 0xd5, 0x36, 0x49, 0x2c, 0x53, 0x65, 0x5d, 0x8b, 0xfe, 0xd9, 0xcd, 0xdf, 0x69, 0x72,
	0xd5, 0xea, 0x04, 0x3e, 0x15, 0x11, 0x4e, 0x53, 0xb1, 0x1f, 0x9e, 0xdc, 0x78, 0x85, 0xd5, 0x4e,
	0xc2, 0xa4, 0xbf, 0xca, 0xe8, 0x71, 0x37, 0x7f, 0xc3, 0xda, 0x1b, 0xa0, 0x43, 0xb5, 0xab, 0xe4,
	0xcc, 0xd5, 0x68, 0x1e, 0xe4, 0x8d, 0xb0, 0x35, 0xd3, 0xca, 0x27, 0x62, 0xa3, 0x13, 0xb2, 0x0a,
	0x09, 0x09, 0x65, 0x32, 0x75, 0x86, 0x3e, 0x3b, 0xa0, 0x30, 0x1c, 0x93, 0x4a, 0x7f, 0x03, 0x26,
	0x03, 0x12, 0x92, 0x98, 0x32, 0x39, 0xeb, 0x14, 0xfe, 0x5f, 0x98, 0x2e, 0xde, 0xf4, 0xad, 0x34,
	0x5f, 0xcf, 0x4d, 0x3f, 0x94, 0xaa, 0xe0, 0x71, 0xa5, 0x9a, 0x66, 0x90, 0xb5, 0x22, 0x7b, 0x85,
	0xe8, 0xeb, 0xaf, 0xfc, 0xc2, 0x18, 0x0e, 0x35, 0x87, 0xac, 0xf5, 0xef, 0x2b, 0xbe, 0x9f, 0x00,
	0x13, 0x46, 0x20, 0x3c, 0x72, 0xc0, 0xf5, 0x73, 0x5b, 0x0c, 0xcb, 0x17, 0x0c, 0xcb, 0x85, 0x01,
	0xe7, 0x56, 0xfe, 0x81, 0xc1, 0x36, 0x09, 0xad, 0xed, 0x7e, 0x3f, 0xfa, 0xf8, 0x5f, 0x19, 0x2e,
	0xe3, 0x06, 0x63, 0xde, 0xe9, 0x37, 0x68, 0xa9, 0xf1, 0x56, 0xff, 0xf9, 0x79, 0x27, 0x39, 0x79,
	0x3a, 0x39, 0xfc, 0xb6, 0x3f, 0x5d, 0xef, 0xe0, 0x37, 0x07, 0xcc, 0x0c, 0x86, 0xa1, 0x3b, 0x01,
	0x97, 0xc7, 0xd1, 0x38, 0x3c, 0xea, 0xdc, 0xe3, 0x4b, 0xd7, 0xa7, 0x0e, 0x8b, 0xc6, 0xe1, 0x3d,
	0xb8, 0x38, 0xc2, 0x61, 0x64, 0x2a, 0xbd, 0x5e, 0x7c, 0xda, 0x4d, 0x76, 0xe0, 0xf9, 0xc3, 0xd2,
	0x38, 0x42, 0xce, 0x5f, 0x38, 0xb9, 0xcb, 0xd5, 0x9a, 0x7d, 0x83, 0xca, 0x46, 0x7f, 0x09, 0x3e,
	0x1c, 0xa1, 0x5f, 0xaf, 0x07, 0xcf, 0xac, 0x0c, 0xcf, 0xae, 0x8c, 0xd3, 0xd9, 0x54, 0x9e, 0xef,
	0x1d, 0xb8, 0xce, 0xfe, 0x81, 0xeb, 0xfc, 0x3e, 0x70, 0x9d, 0x0f, 0x87, 0x6e, 0x66, 0xff, 0xd0,
	0xcd, 0xfc, 0x38, 0x74, 0x33, 0x2f, 0x1f, 0x9d, 0x1a, 0xed, 0x58, 0x68, 0x52, 0x12, 0x7a, 0xa1,
	0x26, 0x1d, 0x58, 0xc8, 0x5b, 0x4b, 0xf8, 0x75, 0xef, 0x3e, 0x33, 0xf1, 0xc1, 0x15, 0xb3, 0x69,
	0xef, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x3f, 0xb0, 0x48, 0x65, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// OsmosisSpotPrice return spot price of pair Osmo/nativeToken
	OsmosisArithmeticTwap(ctx context.Context, in *QueryOsmosisArithmeticTwapRequest, opts ...grpc.CallOption) (*QueryOsmosisArithmeticTwapResponse, error)
	// FeeabsModuleBalances return total balances of feeabs module
	FeeabsModuleBalances(ctx context.Context, in *QueryFeeabsModuleBalacesRequest, opts ...grpc.CallOption) (*QueryFeeabsModuleBalacesResponse, error)
	HostChainConfig(ctx context.Context, in *QueryHostChainConfigRequest, opts ...grpc.CallOption) (*QueryHostChainConfigRespone, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) OsmosisArithmeticTwap(ctx context.Context, in *QueryOsmosisArithmeticTwapRequest, opts ...grpc.CallOption) (*QueryOsmosisArithmeticTwapResponse, error) {
	out := new(QueryOsmosisArithmeticTwapResponse)
	err := c.cc.Invoke(ctx, "/feeabstraction.absfee.v1beta1.Query/OsmosisArithmeticTwap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) FeeabsModuleBalances(ctx context.Context, in *QueryFeeabsModuleBalacesRequest, opts ...grpc.CallOption) (*QueryFeeabsModuleBalacesResponse, error) {
	out := new(QueryFeeabsModuleBalacesResponse)
	err := c.cc.Invoke(ctx, "/feeabstraction.absfee.v1beta1.Query/FeeabsModuleBalances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) HostChainConfig(ctx context.Context, in *QueryHostChainConfigRequest, opts ...grpc.CallOption) (*QueryHostChainConfigRespone, error) {
	out := new(QueryHostChainConfigRespone)
	err := c.cc.Invoke(ctx, "/feeabstraction.absfee.v1beta1.Query/HostChainConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// OsmosisSpotPrice return spot price of pair Osmo/nativeToken
	OsmosisArithmeticTwap(context.Context, *QueryOsmosisArithmeticTwapRequest) (*QueryOsmosisArithmeticTwapResponse, error)
	// FeeabsModuleBalances return total balances of feeabs module
	FeeabsModuleBalances(context.Context, *QueryFeeabsModuleBalacesRequest) (*QueryFeeabsModuleBalacesResponse, error)
	HostChainConfig(context.Context, *QueryHostChainConfigRequest) (*QueryHostChainConfigRespone, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) OsmosisArithmeticTwap(ctx context.Context, req *QueryOsmosisArithmeticTwapRequest) (*QueryOsmosisArithmeticTwapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OsmosisArithmeticTwap not implemented")
}
func (*UnimplementedQueryServer) FeeabsModuleBalances(ctx context.Context, req *QueryFeeabsModuleBalacesRequest) (*QueryFeeabsModuleBalacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeeabsModuleBalances not implemented")
}
func (*UnimplementedQueryServer) HostChainConfig(ctx context.Context, req *QueryHostChainConfigRequest) (*QueryHostChainConfigRespone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HostChainConfig not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_OsmosisArithmeticTwap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryOsmosisArithmeticTwapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).OsmosisArithmeticTwap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feeabstraction.absfee.v1beta1.Query/OsmosisArithmeticTwap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).OsmosisArithmeticTwap(ctx, req.(*QueryOsmosisArithmeticTwapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_FeeabsModuleBalances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFeeabsModuleBalacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).FeeabsModuleBalances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feeabstraction.absfee.v1beta1.Query/FeeabsModuleBalances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).FeeabsModuleBalances(ctx, req.(*QueryFeeabsModuleBalacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_HostChainConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHostChainConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).HostChainConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feeabstraction.absfee.v1beta1.Query/HostChainConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).HostChainConfig(ctx, req.(*QueryHostChainConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "feeabstraction.absfee.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OsmosisArithmeticTwap",
			Handler:    _Query_OsmosisArithmeticTwap_Handler,
		},
		{
			MethodName: "FeeabsModuleBalances",
			Handler:    _Query_FeeabsModuleBalances_Handler,
		},
		{
			MethodName: "HostChainConfig",
			Handler:    _Query_HostChainConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feeabstraction/absfee/v1beta1/query.proto",
}

func (m *QueryHostChainConfigRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryHostChainConfigRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryHostChainConfigRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IbcDenom) > 0 {
		i -= len(m.IbcDenom)
		copy(dAtA[i:], m.IbcDenom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.IbcDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryHostChainConfigRespone) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryHostChainConfigRespone) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryHostChainConfigRespone) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.HostChainConfig.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryOsmosisArithmeticTwapRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryOsmosisArithmeticTwapRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryOsmosisArithmeticTwapRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IbcDenom) > 0 {
		i -= len(m.IbcDenom)
		copy(dAtA[i:], m.IbcDenom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.IbcDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryOsmosisArithmeticTwapResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryOsmosisArithmeticTwapResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryOsmosisArithmeticTwapResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ArithmeticTwap.Size()
		i -= size
		if _, err := m.ArithmeticTwap.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryFeeabsModuleBalacesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFeeabsModuleBalacesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFeeabsModuleBalacesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryFeeabsModuleBalacesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFeeabsModuleBalacesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFeeabsModuleBalacesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Balances) > 0 {
		for iNdEx := len(m.Balances) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Balances[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryHostChainConfigRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IbcDenom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryHostChainConfigRespone) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.HostChainConfig.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryOsmosisArithmeticTwapRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IbcDenom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryOsmosisArithmeticTwapResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ArithmeticTwap.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryFeeabsModuleBalacesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryFeeabsModuleBalacesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Balances) > 0 {
		for _, e := range m.Balances {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryHostChainConfigRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryHostChainConfigRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryHostChainConfigRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryHostChainConfigRespone) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryHostChainConfigRespone: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryHostChainConfigRespone: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostChainConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.HostChainConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryOsmosisArithmeticTwapRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryOsmosisArithmeticTwapRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryOsmosisArithmeticTwapRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryOsmosisArithmeticTwapResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryOsmosisArithmeticTwapResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryOsmosisArithmeticTwapResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ArithmeticTwap", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ArithmeticTwap.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryFeeabsModuleBalacesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryFeeabsModuleBalacesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFeeabsModuleBalacesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryFeeabsModuleBalacesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryFeeabsModuleBalacesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFeeabsModuleBalacesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balances", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Balances = append(m.Balances, types.Coin{})
			if err := m.Balances[len(m.Balances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
