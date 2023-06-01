package proto_012_Psithaca

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/protocol/core"
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
	VotingPower           int32
}

type DelegateInfo struct {
	DelegateInfoContents `tz:"dyn"`
}
