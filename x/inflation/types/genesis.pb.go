// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evmos/inflation/v1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the inflation module's genesis state.
type GenesisState struct {
	// minter is a space for holding current rewards information.
	Minter Minter `protobuf:"bytes,1,opt,name=minter,proto3" json:"minter"`
	// params defines all the paramaters of the module.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
	// current halven period start epoch
	HalvenStartedEpoch int64 `protobuf:"varint,3,opt,name=halven_started_epoch,json=halvenStartedEpoch,proto3" json:"halven_started_epoch,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cb8eee530db1235, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetMinter() Minter {
	if m != nil {
		return m.Minter
	}
	return Minter{}
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetHalvenStartedEpoch() int64 {
	if m != nil {
		return m.HalvenStartedEpoch
	}
	return 0
}

// Params holds parameters for the inflation module.
type Params struct {
	// type of coin to mint
	MintDenom string `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	// epoch provisions from the first epoch
	GenesisEpochProvisions github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=genesis_epoch_provisions,json=genesisEpochProvisions,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"genesis_epoch_provisions"`
	// inflation epoch identifier
	EpochIdentifier string `protobuf:"bytes,3,opt,name=epoch_identifier,json=epochIdentifier,proto3" json:"epoch_identifier,omitempty"`
	// number of epochs take to reduce rewards
	ReductionPeriodInEpochs int64 `protobuf:"varint,4,opt,name=reduction_period_in_epochs,json=reductionPeriodInEpochs,proto3" json:"reduction_period_in_epochs,omitempty"`
	// reduction multiplier to execute on each period
	ReductionFactor github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=reduction_factor,json=reductionFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reduction_factor"`
	// inflation_distribution defines the distribution of the minted denom
	InflationDistribution InflationDistribution `protobuf:"bytes,6,opt,name=inflation_distribution,json=inflationDistribution,proto3" json:"inflation_distribution"`
	// coin to allocate from team vesting supply
	TeamVestingProvision types.Coin `protobuf:"bytes,7,opt,name=team_vesting_provision,json=teamVestingProvision,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"team_vesting_provision"`
	// address to receive team vesting rewards
	TeamVestingReceiver string `protobuf:"bytes,8,opt,name=team_vesting_receiver,json=teamVestingReceiver,proto3" json:"team_vesting_receiver,omitempty"`
	// start epoch to distribute minting rewards
	MintingRewardsAllocationStartEpoch int64 `protobuf:"varint,9,opt,name=minting_rewards_allocation_start_epoch,json=mintingRewardsAllocationStartEpoch,proto3" json:"minting_rewards_allocation_start_epoch,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cb8eee530db1235, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMintDenom() string {
	if m != nil {
		return m.MintDenom
	}
	return ""
}

func (m *Params) GetEpochIdentifier() string {
	if m != nil {
		return m.EpochIdentifier
	}
	return ""
}

func (m *Params) GetReductionPeriodInEpochs() int64 {
	if m != nil {
		return m.ReductionPeriodInEpochs
	}
	return 0
}

func (m *Params) GetInflationDistribution() InflationDistribution {
	if m != nil {
		return m.InflationDistribution
	}
	return InflationDistribution{}
}

func (m *Params) GetTeamVestingProvision() types.Coin {
	if m != nil {
		return m.TeamVestingProvision
	}
	return types.Coin{}
}

func (m *Params) GetTeamVestingReceiver() string {
	if m != nil {
		return m.TeamVestingReceiver
	}
	return ""
}

func (m *Params) GetMintingRewardsAllocationStartEpoch() int64 {
	if m != nil {
		return m.MintingRewardsAllocationStartEpoch
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "evmos.inflation.v1.GenesisState")
	proto.RegisterType((*Params)(nil), "evmos.inflation.v1.Params")
}

func init() { proto.RegisterFile("evmos/inflation/v1/genesis.proto", fileDescriptor_1cb8eee530db1235) }

var fileDescriptor_1cb8eee530db1235 = []byte{
	// 594 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x41, 0x8f, 0xd3, 0x3c,
	0x10, 0x6d, 0xbe, 0xed, 0xf6, 0xa3, 0x06, 0x69, 0x57, 0xa6, 0x5b, 0x42, 0x25, 0xd2, 0xaa, 0x87,
	0x55, 0x17, 0x09, 0x67, 0xbb, 0x5c, 0x10, 0x9c, 0x28, 0x05, 0xd4, 0x03, 0x52, 0x95, 0x95, 0x90,
	0xe0, 0x12, 0xb9, 0x89, 0xdb, 0x58, 0x34, 0x76, 0x64, 0xbb, 0x01, 0x8e, 0xfc, 0x03, 0x8e, 0x1c,
	0x39, 0x73, 0xe6, 0x47, 0xec, 0x09, 0xed, 0x11, 0x71, 0x58, 0x50, 0xfb, 0x47, 0x90, 0xed, 0x6c,
	0x5a, 0x44, 0x91, 0x10, 0xa7, 0x36, 0x33, 0x6f, 0xde, 0xbc, 0x99, 0x37, 0x06, 0x1d, 0x92, 0xa7,
	0x5c, 0xfa, 0x94, 0x4d, 0xe7, 0x58, 0x51, 0xce, 0xfc, 0xbc, 0xef, 0xcf, 0x08, 0x23, 0x92, 0x4a,
	0x94, 0x09, 0xae, 0x38, 0x84, 0x06, 0x81, 0x4a, 0x04, 0xca, 0xfb, 0xad, 0xc6, 0x8c, 0xcf, 0xb8,
	0x49, 0xfb, 0xfa, 0x9f, 0x45, 0xb6, 0xbc, 0x88, 0x4b, 0x4d, 0x36, 0xc1, 0x92, 0xf8, 0x79, 0x7f,
	0x42, 0x14, 0xee, 0xfb, 0x11, 0xa7, 0xac, 0xc8, 0x77, 0xb7, 0xf4, 0x5a, 0xd3, 0x1a, 0x4c, 0xf7,
	0xb3, 0x03, 0xae, 0x3d, 0xb5, 0xfd, 0x4f, 0x15, 0x56, 0x04, 0xde, 0x03, 0xb5, 0x94, 0x32, 0x45,
	0x84, 0xeb, 0x74, 0x9c, 0xde, 0xd5, 0x93, 0x16, 0xfa, 0x5d, 0x0f, 0x7a, 0x66, 0x10, 0x83, 0xea,
	0xd9, 0x45, 0xbb, 0x12, 0x14, 0x78, 0x5d, 0x99, 0x61, 0x81, 0x53, 0xe9, 0xfe, 0xf7, 0xe7, 0xca,
	0xb1, 0x41, 0x5c, 0x56, 0x5a, 0x3c, 0x3c, 0x06, 0x8d, 0x04, 0xcf, 0x73, 0xc2, 0x42, 0xa9, 0xb0,
	0x50, 0x24, 0x0e, 0x49, 0xc6, 0xa3, 0xc4, 0xdd, 0xe9, 0x38, 0xbd, 0x9d, 0x00, 0xda, 0xdc, 0xa9,
	0x4d, 0x3d, 0xd6, 0x99, 0xee, 0x97, 0x5d, 0x50, 0xb3, 0x54, 0xf0, 0x16, 0x00, 0x5a, 0x40, 0x18,
	0x13, 0xc6, 0x53, 0x23, 0xba, 0x1e, 0xd4, 0x75, 0x64, 0xa8, 0x03, 0x30, 0x01, 0x6e, 0xb1, 0x5f,
	0x4b, 0x1a, 0x66, 0x82, 0xe7, 0x54, 0x52, 0xce, 0xac, 0xce, 0xfa, 0x00, 0x69, 0x2d, 0xdf, 0x2e,
	0xda, 0x87, 0x33, 0xaa, 0x92, 0xc5, 0x04, 0x45, 0x3c, 0xf5, 0x8b, 0xcd, 0xda, 0x9f, 0x3b, 0x32,
	0x7e, 0xe5, 0xab, 0xb7, 0x19, 0x91, 0x68, 0x48, 0xa2, 0xa0, 0x59, 0xf0, 0x19, 0x25, 0xe3, 0x92,
	0x0d, 0x1e, 0x81, 0x7d, 0xdb, 0x81, 0xc6, 0x84, 0x29, 0x3a, 0xa5, 0x44, 0x98, 0x09, 0xea, 0xc1,
	0x9e, 0x89, 0x8f, 0xca, 0x30, 0x7c, 0x00, 0x5a, 0x82, 0xc4, 0x8b, 0x48, 0x6f, 0x25, 0xcc, 0x88,
	0xa0, 0x3c, 0x0e, 0x29, 0xb3, 0x02, 0xa5, 0x5b, 0x35, 0x63, 0xdf, 0x28, 0x11, 0x63, 0x03, 0x18,
	0x31, 0xd3, 0x50, 0xc2, 0x17, 0x60, 0x7f, 0x5d, 0x3c, 0xc5, 0x91, 0xe2, 0xc2, 0xdd, 0xfd, 0xa7,
	0x49, 0xf6, 0x4a, 0x9e, 0x27, 0x86, 0x06, 0x4e, 0x41, 0xb3, 0x74, 0x2b, 0x8c, 0xa9, 0x54, 0x82,
	0x4e, 0x16, 0xfa, 0xc3, 0xad, 0x19, 0x4b, 0x8f, 0xb6, 0x59, 0x3a, 0xba, 0xfc, 0x18, 0x6e, 0x14,
	0x14, 0x0e, 0x1f, 0xd0, 0x6d, 0x49, 0xf8, 0xce, 0x01, 0x4d, 0x45, 0x70, 0x1a, 0xe6, 0x44, 0x2a,
	0xca, 0x66, 0x6b, 0x53, 0xdc, 0xff, 0x4d, 0xa3, 0x9b, 0xc8, 0x0a, 0x46, 0xfa, 0xb6, 0x51, 0x71,
	0xdb, 0xe8, 0x11, 0xa7, 0x6c, 0x70, 0xac, 0x89, 0x3f, 0x7d, 0x6f, 0xf7, 0xfe, 0x62, 0x48, 0x5d,
	0x20, 0x83, 0x86, 0x6e, 0xf5, 0xdc, 0x76, 0x2a, 0xfd, 0x82, 0x27, 0xe0, 0xe0, 0x17, 0x09, 0x82,
	0x44, 0x84, 0xe6, 0x44, 0xb8, 0x57, 0x8c, 0x67, 0xd7, 0x37, 0x8a, 0x82, 0x22, 0x05, 0x03, 0x70,
	0xa8, 0x2f, 0xcb, 0xc2, 0x5f, 0x63, 0x11, 0xcb, 0x10, 0xcf, 0xe7, 0x3c, 0xb2, 0x0b, 0x33, 0xc7,
	0x5b, 0x9c, 0x6e, 0xdd, 0x78, 0xd8, 0x2d, 0xd0, 0x81, 0x05, 0x3f, 0x2c, 0xb1, 0xe6, 0x98, 0x8d,
	0x9f, 0xf7, 0xab, 0x1f, 0x3e, 0xb6, 0x2b, 0x83, 0xe1, 0xd9, 0xd2, 0x73, 0xce, 0x97, 0x9e, 0xf3,
	0x63, 0xe9, 0x39, 0xef, 0x57, 0x5e, 0xe5, 0x7c, 0xe5, 0x55, 0xbe, 0xae, 0xbc, 0xca, 0xcb, 0xdb,
	0x1b, 0x73, 0xaa, 0x04, 0x0b, 0x49, 0xa5, 0x6f, 0x1f, 0xf6, 0x9b, 0x8d, 0xa7, 0x6d, 0xe6, 0x9d,
	0xd4, 0xcc, 0xa3, 0xbe, 0xfb, 0x33, 0x00, 0x00, 0xff, 0xff, 0xf0, 0xf0, 0x8c, 0xa3, 0x66, 0x04,
	0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HalvenStartedEpoch != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.HalvenStartedEpoch))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Minter.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MintingRewardsAllocationStartEpoch != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MintingRewardsAllocationStartEpoch))
		i--
		dAtA[i] = 0x48
	}
	if len(m.TeamVestingReceiver) > 0 {
		i -= len(m.TeamVestingReceiver)
		copy(dAtA[i:], m.TeamVestingReceiver)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.TeamVestingReceiver)))
		i--
		dAtA[i] = 0x42
	}
	{
		size, err := m.TeamVestingProvision.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.InflationDistribution.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.ReductionFactor.Size()
		i -= size
		if _, err := m.ReductionFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.ReductionPeriodInEpochs != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ReductionPeriodInEpochs))
		i--
		dAtA[i] = 0x20
	}
	if len(m.EpochIdentifier) > 0 {
		i -= len(m.EpochIdentifier)
		copy(dAtA[i:], m.EpochIdentifier)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.EpochIdentifier)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.GenesisEpochProvisions.Size()
		i -= size
		if _, err := m.GenesisEpochProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Minter.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.HalvenStartedEpoch != 0 {
		n += 1 + sovGenesis(uint64(m.HalvenStartedEpoch))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.GenesisEpochProvisions.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.EpochIdentifier)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.ReductionPeriodInEpochs != 0 {
		n += 1 + sovGenesis(uint64(m.ReductionPeriodInEpochs))
	}
	l = m.ReductionFactor.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.InflationDistribution.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.TeamVestingProvision.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.TeamVestingReceiver)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.MintingRewardsAllocationStartEpoch != 0 {
		n += 1 + sovGenesis(uint64(m.MintingRewardsAllocationStartEpoch))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Minter", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Minter.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HalvenStartedEpoch", wireType)
			}
			m.HalvenStartedEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HalvenStartedEpoch |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisEpochProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GenesisEpochProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EpochIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReductionPeriodInEpochs", wireType)
			}
			m.ReductionPeriodInEpochs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReductionPeriodInEpochs |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReductionFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReductionFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationDistribution", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationDistribution.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TeamVestingProvision", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TeamVestingProvision.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TeamVestingReceiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TeamVestingReceiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintingRewardsAllocationStartEpoch", wireType)
			}
			m.MintingRewardsAllocationStartEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MintingRewardsAllocationStartEpoch |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
