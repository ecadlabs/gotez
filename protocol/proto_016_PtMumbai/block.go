package proto_016_PtMumbai

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/proto_012_Psithaca"
)

type BlockInfoProtocolData struct {
	Header     BlockHeader `tz:"dyn"`
	Metadata   tz.Option[*BlockMetadata]
	Operations []core.OperationsList[GroupContents] `tz:"dyn"`
}

func (block *BlockInfoProtocolData) BlockHeader() *core.BlockHeader {
	return &block.Header.BlockHeader
}

func (block *BlockInfoProtocolData) BlockMetadata() tz.Option[*core.BlockMetadataHeader] {
	if block.Metadata.IsSome() {
		return tz.Some(&block.Metadata.Unwrap().BlockMetadataHeader)
	}
	return tz.None[*core.BlockMetadataHeader]()
}

func (*BlockInfoProtocolData) BlockInfoProtocolData() {}

type BlockMetadata struct {
	BlockMetadataContents `tz:"dyn"`
}

type UnsignedProtocolBlockHeader struct {
	PayloadHash               *tz.BlockPayloadHash
	PayloadRound              int32
	ProofOfWorkNonce          *[tz.ProofOfWorkNonceBytesLen]byte
	SeedNonceHash             tz.Option[*tz.CycleNonceHash]
	LiquidityBakingToggleVote uint8
}

type UnsignedBlockHeader struct {
	core.BlockHeader
	UnsignedProtocolBlockHeader
}

type BlockHeader struct {
	UnsignedBlockHeader
	Signature tz.AnySignature
}

type VotingPeriodInfo proto_012_Psithaca.VotingPeriodInfo

type BlockMetadataContents struct {
	core.BlockMetadataHeader
	Proposer                  tz.PublicKeyHash
	Baker                     tz.PublicKeyHash
	LevelInfo                 LevelInfo
	VotingPeriodInfo          VotingPeriodInfo
	NonceHash                 tz.Option1[*tz.CycleNonceHash]
	Deactivated               []tz.PublicKeyHash `tz:"dyn"`
	BalanceUpdates            []*BalanceUpdate   `tz:"dyn"`
	LiquidityBakingEscapeEMA  int32
	ImplicitOperationsResults []SuccessfulManagerOperationResult `tz:"dyn"`
	ProposerConsensusKey      tz.PublicKeyHash
	BakerConsensusKey         tz.PublicKeyHash
	ConsumedMilligas          tz.BigUint
	DALAttestation            tz.Option[tz.BigInt]
}

type LevelInfo = proto_012_Psithaca.LevelInfo
