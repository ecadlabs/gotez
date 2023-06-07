package proto_015_PtLimaPt

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/protocol/core"
)

type DelegateInfoContents struct {
	FullBalance           tz.BigUint                 `json:"full_balance"`
	CurrentFrozenDeposits tz.BigUint                 `json:"current_frozen_deposits"`
	FrozenDeposits        tz.BigUint                 `json:"frozen_deposits"`
	StakingBalance        tz.BigUint                 `json:"staking_balance"`
	FrozenDepositsLimit   tz.Option[tz.BigUint]      `json:"frozen_deposits_limit"`
	DelegatedContracts    []core.ContractID          `tz:"dyn" json:"delegated_contracts"`
	DelegatedBalance      tz.BigUint                 `json:"delegated_balance"`
	Deactivated           bool                       `json:"deactivated"`
	GracePeriod           int32                      `json:"grace_period"`
	VotingPower           tz.Option[int64]           `json:"voting_power"`
	CurrentBallot         tz.Option[core.BallotKind] `json:"current_ballot"`
	CurrentProposals      []*tz.ProtocolHash         `tz:"dyn" json:"current_proposals"`
	RemainingProposals    int32                      `json:"remaining_proposals"`
	ActiveConsensusKey    tz.PublicKeyHash           `json:"active_consensus_key"`
	PendingConsensusKeys  []*PendingConsensusKey     `tz:"dyn" json:"pending_consensus_keys"`
}

type PendingConsensusKey struct {
	Cycle int32            `json:"cycle"`
	PKH   tz.PublicKeyHash `json:"pkh"`
}

type DelegateInfo struct {
	DelegateInfoContents `tz:"dyn"`
}
