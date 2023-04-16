package proto_011_PtHangz2

import (
	"github.com/ecadlabs/gotez/protocol/expression"
	"github.com/ecadlabs/gotez/protocol/proto_005_PsBABY5H"
)

type ManagerOperation = proto_005_PsBABY5H.ManagerOperation

type RegisterGlobalConstant struct {
	ManagerOperation
	Value expression.Expression `tz:"dyn"`
}

func (*RegisterGlobalConstant) OperationKind() string { return "register_global_constant" }
