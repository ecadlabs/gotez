package proto_016_PtMumbai

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/core/expression"
	"github.com/ecadlabs/gotez/protocol/proto_012_Psithaca"
	"github.com/ecadlabs/gotez/protocol/proto_015_PtLimaPt"
)

type Transaction = proto_012_Psithaca.Transaction
type Parameters = proto_012_Psithaca.Parameters
type Entrypoint = proto_012_Psithaca.Entrypoint
type TransactionDestination = proto_015_PtLimaPt.TransactionDestination
type ToSmartRollup = proto_015_PtLimaPt.ToSmartRollup

type EpDefault = proto_012_Psithaca.EpDefault
type EpRoot = proto_012_Psithaca.EpRoot
type EpDo = proto_012_Psithaca.EpDo
type EpSetDelegate = proto_012_Psithaca.EpSetDelegate
type EpRemoveDelegate = proto_012_Psithaca.EpRemoveDelegate
type EpNamed = proto_012_Psithaca.EpNamed

type TransactionResultDestination interface {
	proto_015_PtLimaPt.TransactionResultDestination
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[TransactionResultDestination]{
		Variants: encoding.Variants[TransactionResultDestination]{
			0: (*ToContract[BalanceUpdateKind])(nil),
			1: (*ToTxRollup[BalanceUpdateKind])(nil),
			2: (*ToSmartRollup)(nil),
		},
	})
}

type TransactionResultContents struct {
	Result TransactionResultDestination
}

func (TransactionResultContents) SuccessfulManagerOperationResult() {}
func (TransactionResultContents) OperationKind() string             { return "transaction" }

type ToContract[T core.BalanceUpdateKind] struct {
	Storage                      tz.Option[expression.Expression]
	BalanceUpdates               []*BalanceUpdate[T]         `tz:"dyn"`
	TicketUpdates                []*TicketReceipt            `tz:"dyn"`
	OriginatedContracts          []core.OriginatedContractID `tz:"dyn"`
	ConsumedMilligas             tz.BigUint
	StorageSize                  tz.BigInt
	PaidStorageSizeDiff          tz.BigInt
	AllocatedDestinationContract bool
	LazyStorageDiff              tz.Option[LazyStorageDiff]
}

func (*ToContract[T]) TransactionResultDestination() {}

type ToTxRollup[T core.BalanceUpdateKind] struct {
	BalanceUpdates      []*BalanceUpdate[T] `tz:"dyn"`
	ConsumedMilligas    tz.BigUint
	TicketHash          *tz.ScriptExprHash
	PaidStorageSizeDiff tz.BigUint
}

func (*ToTxRollup[T]) TransactionResultDestination() {}

type TransactionContentsAndResult[T core.BalanceUpdateKind] struct {
	Transaction
	Metadata ManagerMetadata[TransactionResult, T]
}

func (*TransactionContentsAndResult[T]) OperationContentsAndResult() {}

type TransactionResultApplied struct {
	core.OperationResultApplied[TransactionResultContents]
}

func (*TransactionResultApplied) TransactionResult() {}

type TransactionResultBacktracked struct {
	core.OperationResultBacktracked[TransactionResultContents]
}

func (*TransactionResultBacktracked) TransactionResult() {}

type TransactionResultFailed struct{ core.OperationResultFailed }

func (*TransactionResultFailed) TransactionResult() {}

type TransactionResultSkipped struct{ core.OperationResultSkipped }

func (*TransactionResultSkipped) TransactionResult() {}

type TransactionResult interface {
	TransactionResult()
	core.OperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[TransactionResult]{
		Variants: encoding.Variants[TransactionResult]{
			0: (*TransactionResultApplied)(nil),
			1: (*TransactionResultFailed)(nil),
			2: (*TransactionResultSkipped)(nil),
			3: (*TransactionResultBacktracked)(nil),
		},
	})
}

type TransactionInternalOperationResult struct {
	Source      TransactionDestination
	Nonce       uint16
	Amount      tz.BigUint
	Destination TransactionDestination
	Parameters  tz.Option[Parameters]
	Result      TransactionResult
}

func (*TransactionInternalOperationResult) InternalOperationResult() {}
func (*TransactionInternalOperationResult) OperationKind() string    { return "transaction" }
