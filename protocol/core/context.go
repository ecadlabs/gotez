package core

import tz "github.com/ecadlabs/gotez/v2"

type DelegatesList struct {
	Delegates []tz.PublicKeyHash `tz:"dyn" json:"delegates"`
}

type PendingConsensusKey interface {
	GetCycle() int32
	GetPKH() tz.PublicKeyHash
}

type DelegateInfo interface {
	GetFullBalance() tz.BigUint
	GetCurrentFrozenDeposits() tz.BigUint
	GetFrozenDeposits() tz.BigUint
	GetStakingBalance() tz.BigUint
	GetFrozenDepositsLimit() tz.Option[tz.BigUint]
	GetDelegatedContracts() []ContractID
	GetDelegatedBalance() tz.BigUint
	GetDeactivated() bool
	GetGracePeriod() int32
	GetVotingPower() tz.Option[int64]
	GetCurrentBallot() tz.Option[BallotKind]
	GetCurrentProposals() tz.Option[[]*tz.ProtocolHash]
	GetRemainingProposals() tz.Option[int32]
	GetActiveConsensusKey() tz.Option[tz.PublicKeyHash]
	GetPendingConsensusKeys() tz.Option[[]PendingConsensusKey]
}
