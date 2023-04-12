package operations

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/expression"
)

type PVMKind uint8

const (
	PVMArith PVMKind = iota
	PVM_WASM_2_0_0
)

type SmartRollupOriginate struct {
	ManagerOperation
	PVMKind
	Kernel           []byte                `tz:"dyn"`
	OriginationProof []byte                `tz:"dyn"`
	ParametersTy     expression.Expression `tz:"dyn"`
}

func (*SmartRollupOriginate) OperationKind() string { return "smart_rollup_originate" }

type SmartRollupOriginateResult interface {
	SmartRollupOriginateResult()
	OperationResult
}

type SmartRollupOriginateResultContents struct {
	BalanceUpdates        []*BalanceUpdate `tz:"dyn"`
	Address               *tz.SmartRollupAddress
	GenesisCommitmentHash *tz.MumbaiSmartRollupHash
	ConsumedMilligas      tz.BigUint
	Size                  tz.BigInt
}

type SmartRollupOriginateResultApplied struct {
	OperationResultApplied[SmartRollupOriginateResultContents]
}

func (*SmartRollupOriginateResultApplied) SmartRollupOriginateResult() {}

type SmartRollupOriginateResultBacktracked struct {
	OperationResultBacktracked[SmartRollupOriginateResultContents]
}

func (*SmartRollupOriginateResultBacktracked) SmartRollupOriginateResult() {}

type SmartRollupOriginateResultFailed struct{ OperationResultFailed }

func (*SmartRollupOriginateResultFailed) SmartRollupOriginateResult() {}

type SmartRollupOriginateResultSkipped struct{ OperationResultSkipped }

func (*SmartRollupOriginateResultSkipped) SmartRollupOriginateResult() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SmartRollupOriginateResult]{
		Variants: encoding.Variants[SmartRollupOriginateResult]{
			0: (*SmartRollupOriginateResultApplied)(nil),
			1: (*SmartRollupOriginateResultFailed)(nil),
			2: (*SmartRollupOriginateResultSkipped)(nil),
			3: (*SmartRollupOriginateResultBacktracked)(nil),
		},
	})
}

type SmartRollupOriginateContentsAndResult struct {
	SmartRollupOriginate
	Metadata MetadataWithResult[SmartRollupOriginateResult]
}

func (*SmartRollupOriginateContentsAndResult) OperationContentsAndResult() {}

type SmartRollupOriginateSuccessfulManagerOperationResult SmartRollupOriginateResultContents

func (*SmartRollupOriginateSuccessfulManagerOperationResult) SuccessfulManagerOperationResult() {}
func (*SmartRollupOriginateSuccessfulManagerOperationResult) OperationKind() string {
	return "smart_rollup_originate"
}
