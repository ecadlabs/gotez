package proto_015_PtLimaPt

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/protocol/core"
)

type DelegateInfo struct {
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

func (d *DelegateInfo) GetFullBalance() tz.BigUint                    { return d.FullBalance }
func (d *DelegateInfo) GetCurrentFrozenDeposits() tz.BigUint          { return d.CurrentFrozenDeposits }
func (d *DelegateInfo) GetFrozenDeposits() tz.BigUint                 { return d.FrozenDeposits }
func (d *DelegateInfo) GetStakingBalance() tz.BigUint                 { return d.StakingBalance }
func (d *DelegateInfo) GetFrozenDepositsLimit() tz.Option[tz.BigUint] { return d.FrozenDepositsLimit }
func (d *DelegateInfo) GetDelegatedContracts() []core.ContractID      { return d.DelegatedContracts }
func (d *DelegateInfo) GetDelegatedBalance() tz.BigUint               { return d.DelegatedBalance }
func (d *DelegateInfo) GetDeactivated() bool                          { return d.Deactivated }
func (d *DelegateInfo) GetGracePeriod() int32                         { return d.GracePeriod }
func (d *DelegateInfo) GetVotingPower() tz.Option[int64]              { return d.VotingPower }
func (d *DelegateInfo) GetCurrentBallot() tz.Option[core.BallotKind]  { return d.CurrentBallot }
func (d *DelegateInfo) GetCurrentProposals() tz.Option[[]*tz.ProtocolHash] {
	return tz.Some(d.CurrentProposals)
}
func (d *DelegateInfo) GetRemainingProposals() tz.Option[int32] { return tz.Some(d.RemainingProposals) }
func (d *DelegateInfo) GetActiveConsensusKey() tz.Option[tz.PublicKeyHash] {
	return tz.Some(d.ActiveConsensusKey)
}
func (d *DelegateInfo) GetPendingConsensusKeys() tz.Option[[]core.PendingConsensusKey] {
	keys := make([]core.PendingConsensusKey, len(d.PendingConsensusKeys))
	for i, k := range d.PendingConsensusKeys {
		keys[i] = k
	}
	return tz.Some(keys)
}

type PendingConsensusKey struct {
	Cycle int32            `json:"cycle"`
	PKH   tz.PublicKeyHash `json:"pkh"`
}

func (k *PendingConsensusKey) GetCycle() int32          { return k.Cycle }
func (k *PendingConsensusKey) GetPKH() tz.PublicKeyHash { return k.PKH }
