package core

import (
	"strconv"

	tz "github.com/ecadlabs/gotez/v2"
)

type LevelInfo struct {
	Level              int32 `json:"level"`
	LevelPosition      int32 `json:"level_position"`
	Cycle              int32 `json:"cycle"`
	CyclePosition      int32 `json:"cycle_position"`
	ExpectedCommitment bool  `json:"expected_commitment"`
}

type VotingPeriodInfo struct {
	VotingPeriod VotingPeriod `json:"voting_period"`
	Position     int32        `json:"position"`
	Remaining    int32        `json:"remaining"`
}

type VotingPeriod struct {
	Index         int32            `json:"index"`
	Kind          VotingPeriodKind `json:"kind"`
	StartPosition int32            `json:"start_position"`
}

type VotingPeriodKind uint8

const (
	VotingPeriodProposal VotingPeriodKind = iota
	VotingPeriodExploration
	VotingPeriodCooldown
	VotingPeriodPromotion
	VotingPeriodAdoption
)

func (k VotingPeriodKind) MarshalText() (text []byte, err error) {
	return []byte(k.String()), nil
}

func (k VotingPeriodKind) String() string {
	switch k {
	case VotingPeriodProposal:
		return "proposal"
	case VotingPeriodExploration:
		return "exploration"
	case VotingPeriodCooldown:
		return "cooldown"
	case VotingPeriodPromotion:
		return "promotion"
	case VotingPeriodAdoption:
		return "adoption"
	default:
		return strconv.FormatInt(int64(k), 10)
	}
}

type BlockMetadata interface {
	BalanceUpdates
	GetMetadataHeader() *BlockMetadataHeader
	GetProposer() tz.PublicKeyHash
	GetBaker() tz.PublicKeyHash
	GetLevelInfo() *LevelInfo
	GetVotingPeriodInfo() *VotingPeriodInfo
	GetNonceHash() tz.Option[*tz.CycleNonceHash]
	GetConsumedGas() tz.Option[tz.BigUint]
	GetConsumedMilligas() tz.Option[tz.BigUint]
	GetDeactivated() []tz.PublicKeyHash
	GetImplicitOperationsResults() []SuccessfulManagerOperationResult
	GetProposerConsensusKey() tz.Option[tz.PublicKeyHash]
	GetBakerConsensusKey() tz.Option[tz.PublicKeyHash]
}

type BlockHeader interface {
	Signed
	GetShellHeader() *ShellHeader
	GetPayloadHash() *tz.BlockPayloadHash
	GetPayloadRound() int32
	GetProofOfWorkNonce() *tz.Bytes8
	GetSeedNonceHash() tz.Option[*tz.CycleNonceHash]
	GetLiquidityBakingEscapeVote() bool
}
