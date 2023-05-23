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
	LiquidityBakingToggleVote uint8
}

type UnsignedBlockHeader struct {
	core.BlockHeader
	UnsignedProtocolBlockHeader
}

type BlockHeader struct {
	UnsignedBlockHeader
	Signature *tz.GenericSignature
}
