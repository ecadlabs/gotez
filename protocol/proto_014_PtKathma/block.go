package proto_014_PtKathma

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/proto_012_Psithaca"
)

type VotingPeriodInfo = proto_012_Psithaca.VotingPeriodInfo
type LevelInfo = proto_012_Psithaca.LevelInfo
type UnsignedProtocolBlockHeader = proto_012_Psithaca.UnsignedProtocolBlockHeader
type UnsignedBlockHeader = proto_012_Psithaca.UnsignedBlockHeader
type BlockHeader = proto_012_Psithaca.BlockHeader

type BlockInfoProtocolData struct {
	Header     BlockHeader `tz:"dyn"`
	Metadata   tz.Option[BlockMetadata]
	Operations []core.OperationsList[GroupContents] `tz:"dyn"`
}

func (block *BlockInfoProtocolData) ShellHeader() *core.BlockHeader {
	return &block.Header.BlockHeader
}

func (block *BlockInfoProtocolData) GetSignature() (tz.Signature, error) {
	return block.Header.GetSignature()
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

type BlockMetadataContents struct {
	core.BlockMetadataHeader
	Proposer                  tz.PublicKeyHash
	Baker                     tz.PublicKeyHash
	LevelInfo                 LevelInfo
	VotingPeriodInfo          VotingPeriodInfo
	NonceHash                 tz.Option1[*tz.CycleNonceHash]
	Deactivated               []tz.PublicKeyHash `tz:"dyn"`
	BalanceUpdates            []*BalanceUpdate   `tz:"dyn"`
	LiquidityBakingToggleEMA  int32
	ImplicitOperationsResults []SuccessfulManagerOperationResult `tz:"dyn"`
	ConsumedMilligas          tz.BigUint
	DALSlotAvailability       tz.Option[tz.BigInt]
}
