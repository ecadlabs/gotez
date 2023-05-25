package proto_013_PtJakart

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/proto_012_Psithaca"
)

type VotingPeriodInfo = proto_012_Psithaca.VotingPeriodInfo
type LevelInfo = proto_012_Psithaca.LevelInfo

type BlockInfoProtocolData struct {
	Header     BlockHeader `tz:"dyn"`
	Metadata   tz.Option[BlockMetadata]
	Operations []core.OperationsList[GroupContents] `tz:"dyn"`
}

func (block *BlockInfoProtocolData) BlockHeader() *core.BlockHeader {
	return &block.Header.BlockHeader
}

func (block *BlockInfoProtocolData) BlockMetadata() tz.Option[*core.BlockMetadataHeader] {
	if block.Metadata.IsSome() {
		return tz.Some(&block.Metadata.UnwrapRef().BlockMetadataHeader)
	}
	return tz.None[*core.BlockMetadataHeader]()
}

func (*BlockInfoProtocolData) BlockInfoProtocolData() {}

type BlockMetadata struct {
	BlockMetadataContents `tz:"dyn"`
}

type UnsignedProtocolBlockHeader = proto_012_Psithaca.UnsignedProtocolBlockHeader
type UnsignedBlockHeader = proto_012_Psithaca.UnsignedBlockHeader

type BlockHeader struct {
	UnsignedBlockHeader
	Signature tz.AnySignature
}

type BlockMetadataContents struct {
	core.BlockMetadataHeader
	Proposer                  tz.PublicKeyHash
	Baker                     tz.PublicKeyHash
	LevelInfo                 LevelInfo
	VotingPeriodInfo          VotingPeriodInfo
	NonceHash                 tz.Option1[*tz.CycleNonceHash]
	ConsumedGas               tz.BigUint
	Deactivated               []tz.PublicKeyHash `tz:"dyn"`
	BalanceUpdates            []*BalanceUpdate   `tz:"dyn"`
	LiquidityBakingToggleEMA  int32
	ImplicitOperationsResults []SuccessfulManagerOperationResult `tz:"dyn"`
	ConsumedMilligas          tz.BigUint
}
