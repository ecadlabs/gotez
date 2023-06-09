package proto_012_Psithaca

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/protocol/core"
)

type UnsignedProtocolBlockHeader struct {
	PayloadHash               *tz.BlockPayloadHash          `json:"payload_hash"`
	PayloadRound              int32                         `json:"payload_round"`
	ProofOfWorkNonce          *tz.Bytes8                    `json:"proof_of_work_nonce"`
	SeedNonceHash             tz.Option[*tz.CycleNonceHash] `json:"seed_nonce_hash"`
	LiquidityBakingEscapeVote bool                          `json:"liquidity_baking_escape_vote"`
}

func (h *UnsignedProtocolBlockHeader) GetPayloadHash() *tz.BlockPayloadHash { return h.PayloadHash }
func (h *UnsignedProtocolBlockHeader) GetPayloadRound() int32               { return h.PayloadRound }
func (h *UnsignedProtocolBlockHeader) GetProofOfWorkNonce() *tz.Bytes8 {
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
	Signature *tz.GenericSignature `json:"signature"`
}

func (header *BlockHeader) GetSignature() (tz.Signature, error) {
	return header.Signature, nil
}

type BlockInfoProtocolData struct {
	Header     BlockHeader                          `tz:"dyn" json:"header"`
	Metadata   tz.Option[BlockMetadata]             `json:"metadata"`
	Operations []core.OperationsList[GroupContents] `tz:"dyn" json:"operations"`
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
	Proposer                  tz.PublicKeyHash                   `json:"proposer"`
	Baker                     tz.PublicKeyHash                   `json:"baker"`
	LevelInfo                 core.LevelInfo                     `json:"level_info"`
	VotingPeriodInfo          core.VotingPeriodInfo              `json:"voting_period_info"`
	NonceHash                 tz.Option1[*tz.CycleNonceHash]     `json:"nonce_hash"`
	ConsumedGas               tz.BigUint                         `json:"consumed_gas"`
	Deactivated               []tz.PublicKeyHash                 `tz:"dyn" json:"deactivated"`
	BalanceUpdates            []*BalanceUpdate                   `tz:"dyn" json:"balance_updates"`
	LiquidityBakingEscapeEMA  int32                              `json:"liquidity_baking_escape_ema"`
	ImplicitOperationsResults []SuccessfulManagerOperationResult `tz:"dyn" json:"implicit_operations_results"`
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
func (m *BlockMetadata) GetBalanceUpdates() (updates []core.BalanceUpdate) {
	updates = make([]core.BalanceUpdate, len(m.BalanceUpdates))
	for i, u := range m.BalanceUpdates {
		updates[i] = u
	}
	return
}
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
