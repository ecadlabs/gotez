package proto_016_PtMumbai

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/expression"
	"github.com/ecadlabs/gotez/protocol/proto_005_PsBABY5H"
	kathma "github.com/ecadlabs/gotez/protocol/proto_014_PtKathma"
	"github.com/ecadlabs/gotez/protocol/proto_015_PtLimaPt"
)

type Transaction = proto_005_PsBABY5H.Transaction
type Parameters = proto_005_PsBABY5H.Parameters
type Entrypoint = proto_005_PsBABY5H.Entrypoint
type TransactionResult = proto_015_PtLimaPt.TransactionResult
type TransactionDestination = proto_015_PtLimaPt.TransactionDestination
type TransactionInternalOperationResult = proto_015_PtLimaPt.TransactionInternalOperationResult
type ToSmartRollup = proto_015_PtLimaPt.ToSmartRollup

type EpDefault = proto_005_PsBABY5H.EpDefault
type EpRoot = proto_005_PsBABY5H.EpRoot
type EpDo = proto_005_PsBABY5H.EpDo
type EpSetDelegate = proto_005_PsBABY5H.EpSetDelegate
type EpRemoveDelegate = proto_005_PsBABY5H.EpRemoveDelegate
type EpNamed = proto_005_PsBABY5H.EpNamed

type TransactionResultContents interface {
	TransactionResultContents()
}

type ToContract struct {
	Storage                      tz.Option[expression.Expression]
	BalanceUpdates               []*BalanceUpdate          `tz:"dyn"`
	TicketUpdates                []*TicketReceipt          `tz:"dyn"`
	OriginatedContracts          []tz.OriginatedContractID `tz:"dyn"`
	ConsumedMilligas             tz.BigUint
	StorageSize                  tz.BigInt
	PaidStorageSizeDiff          tz.BigInt
	AllocatedDestinationContract bool
	LazyStorageDiff              tz.Option[LazyStorageDiff]
}

func (*ToContract) TransactionResultContents() {}

type ToTxRollup struct {
	BalanceUpdates      []*BalanceUpdate `tz:"dyn"`
	ConsumedMilligas    tz.BigUint
	TicketHash          *tz.ScriptExprHash
	PaidStorageSizeDiff tz.BigUint
}

func (*ToTxRollup) TransactionResultContents() {}

type TransactionContentsAndResult struct {
	Transaction
	Metadata MetadataWithResult[TransactionResult]
}

func (*TransactionContentsAndResult) OperationContentsAndResult() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[TransactionResultContents]{
		Variants: encoding.Variants[TransactionResultContents]{
			0: (*ToContract)(nil),
			1: (*ToTxRollup)(nil),
			2: (*ToSmartRollup)(nil),
		},
	})
}

type TransactionResultApplied struct {
	kathma.OperationResultApplied[TransactionResultContents]
}

func (*TransactionResultApplied) TransactionResult() {}

type TransactionResultBacktracked struct {
	kathma.OperationResultBacktracked[TransactionResultContents]
}

func (*TransactionResultBacktracked) TransactionResult() {}

type TransactionResultFailed struct{ kathma.OperationResultFailed }

func (*TransactionResultFailed) TransactionResult() {}

type TransactionResultSkipped struct{ kathma.OperationResultSkipped }

func (*TransactionResultSkipped) TransactionResult() {}

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

type TransactionSuccessfulManagerOperationResult struct {
	Result TransactionResultContents
}

func (*TransactionSuccessfulManagerOperationResult) SuccessfulManagerOperationResult() {}
func (*TransactionSuccessfulManagerOperationResult) OperationKind() string             { return "transaction" }
