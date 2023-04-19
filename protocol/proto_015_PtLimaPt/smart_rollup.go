package proto_015_PtLimaPt

import (
	"github.com/ecadlabs/gotez/protocol/core/expression"
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
