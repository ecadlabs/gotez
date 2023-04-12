package operations

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/expression"
)

type Parameters struct {
	Entrypoint Entrypoint
	Value      expression.Expression `tz:"dyn"`
}

type Entrypoint interface {
	Entrypoint()
}

type EpDefault struct{}
type EpRoot struct{}
type EpDo struct{}
type EpSetDelegate struct{}
type EpRemoveDelegate struct{}
type EpNamed struct {
	tz.String
}

func (EpDefault) Entrypoint()        {}
func (EpRoot) Entrypoint()           {}
func (EpDo) Entrypoint()             {}
func (EpSetDelegate) Entrypoint()    {}
func (EpRemoveDelegate) Entrypoint() {}
func (EpNamed) Entrypoint()          {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[Entrypoint]{
		Variants: encoding.Variants[Entrypoint]{
			0:   EpDefault{},
			1:   EpRoot{},
			2:   EpDo{},
			3:   EpSetDelegate{},
			4:   EpRemoveDelegate{},
			255: EpNamed{},
		},
	})
}

type Transaction struct {
	ManagerOperation
	Amount      tz.BigUint
	Destination tz.ContractID
	Parameters  tz.Option[Parameters]
}

func (*Transaction) OperationKind() string { return "transaction" }

type TransactionResult interface {
	TransactionResult()
	OperationResult
}

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

type TicketReceipt struct {
	TicketToken TicketToken
	Updates     []*TicketReceiptUpdate `tz:"dyn"`
}

type TicketToken struct {
	Ticketer    tz.ContractID
	ContentType expression.Expression
	Content     expression.Expression
}

type TicketReceiptUpdate struct {
	Account tz.TransactionDestination
	Amount  tz.BigInt
}

type ToTxRollup struct {
	BalanceUpdates      []*BalanceUpdate `tz:"dyn"`
	ConsumedMilligas    tz.BigUint
	TicketHash          *tz.ScriptExprHash
	PaidStorageSizeDiff tz.BigUint
}

func (*ToTxRollup) TransactionResultContents() {}

type ToSmartRollup struct {
	ConsumedMilligas tz.BigUint
	TicketUpdates    []*TicketReceipt `tz:"dyn"`
}

func (*ToSmartRollup) TransactionResultContents() {}

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
	OperationResultApplied[TransactionResultContents]
}

func (*TransactionResultApplied) TransactionResult() {}

type TransactionResultBacktracked struct {
	OperationResultBacktracked[TransactionResultContents]
}

func (*TransactionResultBacktracked) TransactionResult() {}

type TransactionResultFailed struct{ OperationResultFailed }

func (*TransactionResultFailed) TransactionResult() {}

type TransactionResultSkipped struct{ OperationResultSkipped }

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

type TransactionContentsAndResult struct {
	Transaction
	Metadata MetadataWithResult[TransactionResult]
}

func (*TransactionContentsAndResult) OperationContentsAndResult() {}

type TransactionInternalOperationResult struct {
	Source      tz.TransactionDestination
	Nonce       uint16
	Amount      tz.BigUint
	Destination tz.TransactionDestination
	Parameters  tz.Option[Parameters]
	Result      TransactionResult
}

func (*TransactionInternalOperationResult) InternalOperationResult() {}

type TransactionSuccessfulManagerOperationResult struct {
	Result TransactionResultContents
}

func (*TransactionSuccessfulManagerOperationResult) SuccessfulManagerOperationResult() {}
func (*TransactionSuccessfulManagerOperationResult) OperationKind() string             { return "transaction" }
