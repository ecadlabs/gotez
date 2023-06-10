package proto_013_PtJakart

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/protocol/core"
)

type DelegateInfo struct {
	FullBalance           tz.BigUint            `json:"full_balance"`
	CurrentFrozenDeposits tz.BigUint            `json:"current_frozen_deposits"`
	FrozenDeposits        tz.BigUint            `json:"frozen_deposits"`
	StakingBalance        tz.BigUint            `json:"staking_balance"`
	FrozenDepositsLimit   tz.Option[tz.BigUint] `json:"frozen_deposits_limit"`
	DelegatedContracts    []core.ContractID     `tz:"dyn" json:"delegated_contracts"`
	DelegatedBalance      tz.BigUint            `json:"delegated_balance"`
	Deactivated           bool                  `json:"deactivated"`
	GracePeriod           int32                 `json:"grace_period"`
	VotingPower           int64                 `json:"voting_power"`
}

func (d *DelegateInfo) GetFullBalance() tz.BigUint                    { return d.FullBalance }
func (d *DelegateInfo) GetCurrentFrozenDeposits() tz.BigUint          { return d.CurrentFrozenDeposits }
func (d *DelegateInfo) GetFrozenDeposits() tz.BigUint                 { return d.FrozenDeposits }
func (d *DelegateInfo) GetStakingBalance() tz.BigUint                 { return d.StakingBalance }
func (d *DelegateInfo) GetFrozenDepositsLimit() tz.Option[tz.BigUint] { return d.FrozenDepositsLimit }
func (d *DelegateInfo) GetDelegatedContracts() []core.ContractID      { return d.DelegatedContracts }
func (d *DelegateInfo) GetDelegatedBalance() tz.BigUint               { return d.DelegatedBalance }
func (d *DelegateInfo) GetDeactivated() bool                          { return d.Deactivated }
func (d *DelegateInfo) GetGracePeriod() int32                         { return d.GracePeriod }
func (d *DelegateInfo) GetVotingPower() tz.Option[int64]              { return tz.Some(d.VotingPower) }
func (d *DelegateInfo) GetCurrentBallot() tz.Option[core.BallotKind] {
	return tz.None[core.BallotKind]()
}
func (d *DelegateInfo) GetCurrentProposals() tz.Option[[]*tz.ProtocolHash] {
	return tz.None[[]*tz.ProtocolHash]()
}
func (d *DelegateInfo) GetRemainingProposals() tz.Option[int32] { return tz.None[int32]() }
func (d *DelegateInfo) GetActiveConsensusKey() tz.Option[tz.PublicKeyHash] {
	return tz.None[tz.PublicKeyHash]()
}
func (d *DelegateInfo) GetPendingConsensusKeys() tz.Option[[]core.PendingConsensusKey] {
	return tz.None[[]core.PendingConsensusKey]()
}
