package proto_013_PtJakart

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/proto_012_Psithaca"
)

type UnsignedProtocolBlockHeader = proto_012_Psithaca.UnsignedProtocolBlockHeader
type UnsignedBlockHeader = proto_012_Psithaca.UnsignedBlockHeader
type BlockHeader = proto_012_Psithaca.BlockHeader

type BlockInfoProtocolData struct {
	Header     BlockHeader `tz:"dyn"`
	Metadata   tz.Option[BlockMetadata]
	Operations []core.OperationsList[GroupContents] `tz:"dyn"`
}

func (block *BlockInfoProtocolData) GetHeader() core.BlockHeader { return &block.Header }
func (block *BlockInfoProtocolData) GetMetadata() tz.Option[core.BlockMetadata] {
	if m, ok := block.Metadata.CheckUnwrapPtr(); ok {
		return tz.Some[core.BlockMetadata](m)
	}
	return tz.None[core.BlockMetadata]()
}

func (block *BlockInfoProtocolData) GetOperations() [][]core.OperationsGroup {
	out := make([][]core.OperationsGroup, len(block.Operations))
	for i, list := range block.Operations {
		out[i] = list.GetGroups()
	}
	return out
}

type BlockMetadata struct {
	BlockMetadataContents `tz:"dyn"`
}

type BlockMetadataContents struct {
	core.BlockMetadataHeader
	Proposer                  tz.PublicKeyHash
	Baker                     tz.PublicKeyHash
	LevelInfo                 core.LevelInfo
	VotingPeriodInfo          core.VotingPeriodInfo
	NonceHash                 tz.Option1[*tz.CycleNonceHash]
	ConsumedGas               tz.BigUint
	Deactivated               []tz.PublicKeyHash `tz:"dyn"`
	BalanceUpdates            []*BalanceUpdate   `tz:"dyn"`
	LiquidityBakingToggleEMA  int32
	ImplicitOperationsResults []SuccessfulManagerOperationResult `tz:"dyn"`
	ConsumedMilligas          tz.BigUint
}

func (m *BlockMetadata) GetMetadataHeader() *core.BlockMetadataHeader { return &m.BlockMetadataHeader }
func (m *BlockMetadata) GetProposer() tz.PublicKeyHash                { return m.Proposer }
func (m *BlockMetadata) GetBaker() tz.PublicKeyHash                   { return m.Baker }
func (m *BlockMetadata) GetLevelInfo() *core.LevelInfo                { return &m.LevelInfo }
func (m *BlockMetadata) GetVotingPeriodInfo() *core.VotingPeriodInfo  { return &m.VotingPeriodInfo }
func (m *BlockMetadata) GetNonceHash() tz.Option[*tz.CycleNonceHash]  { return m.NonceHash.Option }
func (m *BlockMetadata) GetConsumedGas() tz.Option[tz.BigUint]        { return tz.Some(m.ConsumedGas) }
func (m *BlockMetadata) GetConsumedMilligas() tz.Option[tz.BigUint] {
	return tz.Some(m.ConsumedMilligas)
}
func (m *BlockMetadata) GetDeactivated() []tz.PublicKeyHash { return m.Deactivated }
func (m *BlockMetadata) GetImplicitOperationsResults() []core.SuccessfulManagerOperationResult {
	out := make([]core.SuccessfulManagerOperationResult, len(m.ImplicitOperationsResults))
	for i, v := range m.ImplicitOperationsResults {
		out[i] = v
	}
	return out
}
func (m *BlockMetadata) GetProposerConsensusKey() tz.Option[tz.PublicKeyHash] {
	return tz.None[tz.PublicKeyHash]()
}
func (m *BlockMetadata) GetBakerConsensusKey() tz.Option[tz.PublicKeyHash] {
	return tz.None[tz.PublicKeyHash]()
}
