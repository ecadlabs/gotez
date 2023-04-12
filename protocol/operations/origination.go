package operations

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/expression"
)

type Script struct {
	Code    expression.Expression `tz:"dyn"`
	Storage expression.Expression `tz:"dyn"`
}

type Origination struct {
	ManagerOperation
	Balance  tz.BigUint
	Delegate tz.Option[tz.PublicKeyHash]
	Script   Script
}

func (*Origination) OperationKind() string { return "origination" }

type OriginationResult interface {
	OriginationResult()
	OperationResult
}

type OriginationResultContents struct {
	BalanceUpdates      []*BalanceUpdate          `tz:"dyn"`
	OriginatedContracts []tz.OriginatedContractID `tz:"dyn"`
	ConsumedMilligas    tz.BigUint
	StorageSize         tz.BigInt
	PaidStorageSizeDiff tz.BigInt
	LazyStorageDiff     tz.Option[LazyStorageDiff]
}

type OriginationResultApplied struct {
	OperationResultApplied[OriginationResultContents]
}

func (*OriginationResultApplied) OriginationResult() {}

type OriginationResultBacktracked struct {
	OperationResultBacktracked[OriginationResultContents]
}

func (*OriginationResultBacktracked) OriginationResult() {}

type OriginationResultFailed struct{ OperationResultFailed }

func (*OriginationResultFailed) OriginationResult() {}

type OriginationResultSkipped struct{ OperationResultSkipped }

func (*OriginationResultSkipped) OriginationResult() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[OriginationResult]{
		Variants: encoding.Variants[OriginationResult]{
			0: (*OriginationResultApplied)(nil),
			1: (*OriginationResultFailed)(nil),
			2: (*OriginationResultSkipped)(nil),
			3: (*OriginationResultBacktracked)(nil),
		},
	})
}

type OriginationContentsAndResult struct {
	Origination
	Metadata MetadataWithResult[OriginationResult]
}

func (*OriginationContentsAndResult) OperationContentsAndResult() {}

type OriginationInternalOperationResult struct {
	Source   tz.TransactionDestination
	Nonce    uint16
	Balance  tz.BigUint
	Delegate tz.Option[tz.PublicKeyHash]
	Script   Script
	Result   OriginationResult
}

func (*OriginationInternalOperationResult) InternalOperationResult() {}

type OriginationSuccessfulManagerOperationResult OriginationResultContents

func (*OriginationSuccessfulManagerOperationResult) SuccessfulManagerOperationResult() {}
func (*OriginationSuccessfulManagerOperationResult) OperationKind() string             { return "origination" }
