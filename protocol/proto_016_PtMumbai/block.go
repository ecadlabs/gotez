package proto_016_PtMumbai

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/protocol/proto"
)

type UnsignedProtocolBlockHeader struct {
	PayloadHash               *tz.BlockPayloadHash
	PayloadRound              int32
	ProofOfWorkNonce          *[tz.ProofOfWorkNonceBytesLen]byte
	SeedNonceHash             tz.Option[*tz.CycleNonceHash]
	LiquidityBakingToggleVote uint8
}

type ProtocolBlockHeader struct {
	UnsignedProtocolBlockHeader
	Signature tz.AnySignature
}

func (*ProtocolBlockHeader) ProtocolBlockHeader() {}

type UnsignedBlockHeader struct {
	proto.BlockHeader
	UnsignedProtocolBlockHeader
}

type BlockHeader struct {
	UnsignedBlockHeader
	Signature tz.AnySignature
}

type BlockMetadataContents struct {
	proto.BlockMetadataHeader
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
