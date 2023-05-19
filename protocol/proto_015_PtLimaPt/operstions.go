package proto_015_PtLimaPt

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/core/expression"
	"github.com/ecadlabs/gotez/protocol/proto_012_Psithaca"
)

type ManagerOperation = proto_012_Psithaca.ManagerOperation
type Entrypoint = proto_012_Psithaca.Entrypoint

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

type EventResult interface {
	EventResult()
	core.OperationResult
}

type EventResultContents struct {
	ConsumedMilligas tz.BigUint
}

func (EventResultContents) SuccessfulManagerOperationResult() {}
func (EventResultContents) OperationKind() string             { return "event" }

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

type RevealResultContents EventResultContents

func (*RevealResultContents) SuccessfulManagerOperationResult() {}
func (*RevealResultContents) OperationKind() string             { return "reveal" }

type DelegationInternalOperationResult struct {
	Source   TransactionDestination
	Nonce    uint16
	Delegate tz.Option[tz.PublicKeyHash]
	Result   EventResult
}

func (*DelegationInternalOperationResult) InternalOperationResult() {}
func (*DelegationInternalOperationResult) OperationKind() string    { return "delegation" }

type DelegationResultContents EventResultContents

func (*DelegationResultContents) SuccessfulManagerOperationResult() {}
func (*DelegationResultContents) OperationKind() string             { return "delegation" }

type SetDepositsLimitResultContents EventResultContents

func (*SetDepositsLimitResultContents) SuccessfulManagerOperationResult() {}
func (*SetDepositsLimitResultContents) OperationKind() string {
	return "set_deposits_limit"
}

type UpdateConsensusKeyResultContents EventResultContents

func (*UpdateConsensusKeyResultContents) SuccessfulManagerOperationResult() {}
func (*UpdateConsensusKeyResultContents) OperationKind() string {
	return "update_consensus_key"
}
