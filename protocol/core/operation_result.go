package core

import (
	tz "github.com/ecadlabs/gotez"
)

type ManagerOperationResult interface {
	OperationResultStatus() string
}

type SuccessfulManagerOperationResult interface {
	Operation
	SuccessfulManagerOperationResult()
}

//json:status=applied
type OperationResultApplied[T any] struct {
	Result T `json:"result"`
}

func (*OperationResultApplied[T]) OperationResultStatus() string { return "applied" }

//json:status=backtracked
type OperationResultBacktracked[T any] struct {
	Errors tz.Option[OperationResultErrors] `json:"errors"`
	Result T                                `json:"result"`
}

func (*OperationResultBacktracked[T]) OperationResultStatus() string { return "backtracked" }

//json:status=failed
type OperationResultErrors struct {
	Errors []Bytes `tz:"dyn" json:"errors"`
}

type OperationResultFailed OperationResultErrors

func (*OperationResultFailed) OperationResultStatus() string { return "failed" }

//json:status=skipped
type OperationResultSkipped struct{}

func (*OperationResultSkipped) OperationResultStatus() string { return "skipped" }
