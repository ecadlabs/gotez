package proto_012_Psithaca

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/protocol/core"
)

type UnsignedProtocolBlockHeader struct {
	PayloadHash               *tz.BlockPayloadHash
	PayloadRound              int32
	ProofOfWorkNonce          *[tz.ProofOfWorkNonceBytesLen]byte
	SeedNonceHash             tz.Option[*tz.CycleNonceHash]
	LiquidityBakingEscapeVote bool
}

func (h *UnsignedProtocolBlockHeader) GetPayloadHash() *tz.BlockPayloadHash { return h.PayloadHash }
func (h *UnsignedProtocolBlockHeader) GetPayloadRound() int32               { return h.PayloadRound }
func (h *UnsignedProtocolBlockHeader) GetProofOfWorkNonce() *[tz.ProofOfWorkNonceBytesLen]byte {
	return h.ProofOfWorkNonce
}
func (h *UnsignedProtocolBlockHeader) GetSeedNonceHash() tz.Option[*tz.CycleNonceHash] {
	return h.SeedNonceHash
}
func (h *UnsignedProtocolBlockHeader) GetLiquidityBakingEscapeVote() bool {
	return h.LiquidityBakingEscapeVote
}

type UnsignedBlockHeader struct {
	core.ShellHeader
	UnsignedProtocolBlockHeader
}

func (header *UnsignedBlockHeader) GetShellHeader() *core.ShellHeader {
	return &header.ShellHeader
}

type BlockHeader struct {
	UnsignedBlockHeader
	Signature *tz.GenericSignature
}

func (header *BlockHeader) GetSignature() (tz.Signature, error) {
	return header.Signature, nil
}

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
	LiquidityBakingEscapeEMA  int32
	ImplicitOperationsResults []SuccessfulManagerOperationResult `tz:"dyn"`
}

func (m *BlockMetadata) GetMetadataHeader() *core.BlockMetadataHeader { return &m.BlockMetadataHeader }
func (m *BlockMetadata) GetProposer() tz.PublicKeyHash                { return m.Proposer }
func (m *BlockMetadata) GetBaker() tz.PublicKeyHash                   { return m.Baker }
func (m *BlockMetadata) GetLevelInfo() *core.LevelInfo                { return &m.LevelInfo }
func (m *BlockMetadata) GetVotingPeriodInfo() *core.VotingPeriodInfo  { return &m.VotingPeriodInfo }
func (m *BlockMetadata) GetNonceHash() tz.Option[*tz.CycleNonceHash]  { return m.NonceHash.Option }
func (m *BlockMetadata) GetConsumedGas() tz.Option[tz.BigUint]        { return tz.Some(m.ConsumedGas) }
func (m *BlockMetadata) GetConsumedMilligas() tz.Option[tz.BigUint]   { return tz.None[tz.BigUint]() }
func (m *BlockMetadata) GetDeactivated() []tz.PublicKeyHash           { return m.Deactivated }
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
