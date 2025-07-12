package proto_023_PtSeouLo

import (
	"slices"

	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/proto_018_Proxford"
	"github.com/ecadlabs/gotez/v2/protocol/proto_022_PsRiotum"
)

type Transaction = proto_022_PsRiotum.Transaction
type TransactionContentsAndResult = proto_022_PsRiotum.TransactionContentsAndResult
type TransactionResultDestination = proto_022_PsRiotum.TransactionResultDestination
type Parameters = proto_022_PsRiotum.Parameters
type EpDefault = proto_022_PsRiotum.EpDefault
type EpRoot = proto_022_PsRiotum.EpRoot
type EpDo = proto_022_PsRiotum.EpDo
type EpSetDelegate = proto_022_PsRiotum.EpSetDelegate
type EpRemoveDelegate = proto_022_PsRiotum.EpRemoveDelegate
type EpDeposit = proto_022_PsRiotum.EpDeposit
type EpStake = proto_022_PsRiotum.EpStake
type EpUnstake = proto_022_PsRiotum.EpUnstake
type EpFinalizeUnstake = proto_022_PsRiotum.EpFinalizeUnstake
type EpSetDelegateParameters = proto_022_PsRiotum.EpSetDelegateParameters
type EpNamed = proto_022_PsRiotum.EpNamed
type ToContract = proto_022_PsRiotum.ToContract
type ToSmartRollup = proto_022_PsRiotum.ToSmartRollup
type PseudoOperation = proto_018_Proxford.PseudoOperation

func ListPseudoOperations() []string {
	ops := encoding.ListVariants[PseudoOperation]()
	ret := make([]string, len(ops))
	for i, op := range ops {
		ret[i] = op.PseudoOperation()
	}
	slices.Sort(ret)
	return slices.Compact(ret)
}
