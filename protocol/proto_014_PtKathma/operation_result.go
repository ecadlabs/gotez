package proto_014_PtKathma

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/protocol/core"
)

type OperationResultApplied[T any] struct {
	Result T
}

func (*OperationResultApplied[T]) OperationResultKind() string { return "applied" }

type OperationResultBacktracked[T any] struct {
	Errors tz.Option[OperationResultErrors]
	Result T
}

func (*OperationResultBacktracked[T]) OperationResultKind() string { return "backtracked" }

type OperationResultErrors struct {
	Errors []core.Bytes `tz:"dyn"`
}

type OperationResultFailed OperationResultErrors

func (*OperationResultFailed) OperationResultKind() string { return "failed" }

type OperationResultSkipped struct{}

func (*OperationResultSkipped) OperationResultKind() string { return "skipped" }
