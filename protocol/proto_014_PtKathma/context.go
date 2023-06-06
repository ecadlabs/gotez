package proto_014_PtKathma

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/protocol/core"
)

type DelegateInfoContents struct {
	FullBalance           tz.BigUint
	CurrentFrozenDeposits tz.BigUint
	FrozenDeposits        tz.BigUint
	StakingBalance        tz.BigUint
	FrozenDepositsLimit   tz.Option[tz.BigUint]
	DelegatedContracts    []core.ContractID `tz:"dyn"`
	DelegatedBalance      tz.BigUint
	Deactivated           bool
	GracePeriod           int32
	VotingPower           int64
	CurrentBallot         tz.Option[core.BallotKind]
	CurrentProposals      []*tz.ProtocolHash `tz:"dyn"`
	RemainingProposals    int32
}

type DelegateInfo struct {
	DelegateInfoContents `tz:"dyn"`
}
