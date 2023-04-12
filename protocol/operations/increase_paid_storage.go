package operations

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
)

type IncreasePaidStorage struct {
	ManagerOperation
	Amount      tz.BigInt
	Destination tz.OriginatedContractID
}

func (*IncreasePaidStorage) OperationKind() string { return "increase_paid_storage" }

type IncreasePaidStorageResult interface {
	IncreasePaidStorageResult()
	OperationResult
}

type IncreasePaidStorageResultContents struct {
	BalanceUpdates   []*BalanceUpdate `tz:"dyn"`
	ConsumedMilligas tz.BigUint
}

type IncreasePaidStorageResultApplied struct {
	OperationResultApplied[IncreasePaidStorageResultContents]
}

func (*IncreasePaidStorageResultApplied) IncreasePaidStorageResult() {}

type IncreasePaidStorageResultBacktracked struct {
	OperationResultBacktracked[IncreasePaidStorageResultContents]
}

func (*IncreasePaidStorageResultBacktracked) IncreasePaidStorageResult() {}

type IncreasePaidStorageResultFailed struct{ OperationResultFailed }

func (*IncreasePaidStorageResultFailed) IncreasePaidStorageResult() {}

type IncreasePaidStorageResultSkipped struct{ OperationResultSkipped }

func (*IncreasePaidStorageResultSkipped) IncreasePaidStorageResult() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[IncreasePaidStorageResult]{
		Variants: encoding.Variants[IncreasePaidStorageResult]{
			0: (*IncreasePaidStorageResultApplied)(nil),
			1: (*IncreasePaidStorageResultFailed)(nil),
			2: (*IncreasePaidStorageResultSkipped)(nil),
			3: (*IncreasePaidStorageResultBacktracked)(nil),
		},
	})
}

type IncreasePaidStorageContentsAndResult struct {
	IncreasePaidStorage
	Metadata MetadataWithResult[IncreasePaidStorageResult]
}

func (*IncreasePaidStorageContentsAndResult) OperationContentsAndResult() {}

type IncreasePaidStorageSuccessfulManagerOperationResult IncreasePaidStorageResultContents

func (*IncreasePaidStorageSuccessfulManagerOperationResult) SuccessfulManagerOperationResult() {}
func (*IncreasePaidStorageSuccessfulManagerOperationResult) OperationKind() string {
	return "increase_paid_storage"
}
