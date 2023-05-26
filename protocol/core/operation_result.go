package core

import (
	tz "github.com/ecadlabs/gotez"
)

type ManagerOperationResult interface {
	OperationResultKind() string
}

type SuccessfulManagerOperationResult interface {
	OperationContents
	SuccessfulManagerOperationResult()
}

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
	Errors []Bytes `tz:"dyn"`
}

type OperationResultFailed OperationResultErrors

func (*OperationResultFailed) OperationResultKind() string { return "failed" }

type OperationResultSkipped struct{}

func (*OperationResultSkipped) OperationResultKind() string { return "skipped" }
