package proto_015_PtLimaPt

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/core/expression"
	"github.com/ecadlabs/gotez/protocol/proto_012_Psithaca"
)

type ManagerOperation = proto_012_Psithaca.ManagerOperation

type LazyStorageDiff struct {
	Opaque []byte `tz:"dyn"` // TODO: lazy storage diff
}

type UpdateConsensusKey struct {
	ManagerOperation
	PublicKey tz.PublicKey
}

func (*UpdateConsensusKey) OperationKind() string { return "update_consensus_key" }

type DrainDelegate struct {
	ConsensusKey tz.PublicKeyHash
	Delegate     tz.PublicKeyHash
	Destination  tz.PublicKeyHash
}

func (*DrainDelegate) OperationKind() string { return "drain_delegate" }

type Entrypoint = proto_012_Psithaca.Entrypoint

type EventResult interface {
	EventResult()
	core.OperationResult
}

type EventResultContents struct {
	ConsumedMilligas tz.BigUint
}

type EventResultApplied struct {
	core.OperationResultApplied[EventResultContents]
}

func (*EventResultApplied) EventResult() {}

type EventResultBacktracked struct {
	core.OperationResultBacktracked[EventResultContents]
}

func (*EventResultBacktracked) EventResult() {}

type EventResultFailed struct{ core.OperationResultFailed }

func (*EventResultFailed) EventResult() {}

type EventResultSkipped struct{ core.OperationResultSkipped }

func (*EventResultSkipped) EventResult() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[EventResult]{
		Variants: encoding.Variants[EventResult]{
			0: (*EventResultApplied)(nil),
			1: (*EventResultFailed)(nil),
			2: (*EventResultSkipped)(nil),
			3: (*EventResultBacktracked)(nil),
		},
	})
}

type EventInternalOperationResult struct {
	Source  TransactionDestination
	Nonce   uint16
	Type    expression.Expression
	Tag     tz.Option[Entrypoint]
	Payload tz.Option[expression.Expression]
	Result  EventResult
}

func (*EventInternalOperationResult) InternalOperationResult() {}
func (*EventInternalOperationResult) OperationKind() string    { return "event" }

type RevealSuccessfulManagerOperationResult EventResultContents

func (*RevealSuccessfulManagerOperationResult) SuccessfulManagerOperationResult() {}
func (*RevealSuccessfulManagerOperationResult) OperationKind() string             { return "reveal" }

type DelegationInternalOperationResult struct {
	Source   TransactionDestination
	Nonce    uint16
	Delegate tz.Option[tz.PublicKeyHash]
	Result   EventResult
}

func (*DelegationInternalOperationResult) InternalOperationResult() {}
func (*DelegationInternalOperationResult) OperationKind() string    { return "delegation" }

type DelegationSuccessfulManagerOperationResult EventResultContents

func (*DelegationSuccessfulManagerOperationResult) SuccessfulManagerOperationResult() {}
func (*DelegationSuccessfulManagerOperationResult) OperationKind() string             { return "delegation" }

type SetDepositsLimitSuccessfulManagerOperationResult EventResultContents

func (*SetDepositsLimitSuccessfulManagerOperationResult) SuccessfulManagerOperationResult() {}
func (*SetDepositsLimitSuccessfulManagerOperationResult) OperationKind() string {
	return "set_deposits_limit"
}

type UpdateConsensusKeySuccessfulManagerOperationResult EventResultContents

func (*UpdateConsensusKeySuccessfulManagerOperationResult) SuccessfulManagerOperationResult() {}
func (*UpdateConsensusKeySuccessfulManagerOperationResult) OperationKind() string {
	return "update_consensus_key"
}
