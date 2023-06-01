package proto_012_Psithaca

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/protocol/core"
)

type LevelInfo struct {
	Level              int32
	LevelPosition      int32
	Cycle              int32
	CyclePosition      int32
	ExpectedCommitment bool
}

type VotingPeriodInfo struct {
	VotingPeriod VotingPeriod
	Position     int32
	Remaining    int32
}

type VotingPeriod struct {
	Index         int32
	Kind          VotingPeriodKind
	StartPosition int32
}

type VotingPeriodKind uint8

const (
	VotingPeriodProposal VotingPeriodKind = iota
	VotingPeriodExploration
	VotingPeriodCooldown
	VotingPeriodPromotion
	VotingPeriodAdoption
)

type UnsignedProtocolBlockHeader struct {
	PayloadHash               *tz.BlockPayloadHash
	PayloadRound              int32
	ProofOfWorkNonce          *[tz.ProofOfWorkNonceBytesLen]byte
	SeedNonceHash             tz.Option[*tz.CycleNonceHash]
	LiquidityBakingEscapeVote bool
}

type UnsignedBlockHeader struct {
	core.BlockHeader
	UnsignedProtocolBlockHeader
}

type BlockHeader struct {
	UnsignedBlockHeader
	Signature *tz.GenericSignature
}

func (header *BlockHeader) ShellHeader() *core.BlockHeader {
	return &header.BlockHeader
}

func (header *BlockHeader) GetSignature() (tz.Signature, error) {
	return header.Signature, nil
}

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
	ConsumedGas               tz.BigUint
	Deactivated               []tz.PublicKeyHash `tz:"dyn"`
	BalanceUpdates            []*BalanceUpdate   `tz:"dyn"`
	LiquidityBakingEscapeEMA  int32
	ImplicitOperationsResults []SuccessfulManagerOperationResult `tz:"dyn"`
}
