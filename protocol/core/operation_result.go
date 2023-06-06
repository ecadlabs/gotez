package core

import (
	tz "github.com/ecadlabs/gotez/v2"
)

type InternalOperationResult interface {
	OperationContents
	InternalOperationResult() ManagerOperationResult
}

type ManagerOperationResult interface {
	Status() string
	IsApplied() bool
}

type ManagerOperationResultAppliedOrBacktracked interface {
	ManagerOperationResult
	GetResult() any
}

type SuccessfulManagerOperationResult interface {
	OperationContents
	SuccessfulManagerOperationResult()
}

//json:status=applied
type OperationResultApplied[T any] struct {
	Result T `json:"result"`
}

func (*OperationResultApplied[T]) Status() string   { return "applied" }
func (*OperationResultApplied[T]) IsApplied() bool  { return true }
func (r *OperationResultApplied[T]) GetResult() any { return r.Result }

//json:status=backtracked
type OperationResultBacktracked[T any] struct {
	Errors tz.Option[OperationResultErrors] `json:"errors"`
	Result T                                `json:"result"`
}

func (*OperationResultBacktracked[T]) Status() string   { return "backtracked" }
func (*OperationResultBacktracked[T]) IsApplied() bool  { return false }
func (r *OperationResultBacktracked[T]) GetResult() any { return r.Result }

//json:status=failed
type OperationResultErrors struct {
	Errors []Bytes `tz:"dyn" json:"errors"`
}

type OperationResultFailed OperationResultErrors

func (*OperationResultFailed) Status() string  { return "failed" }
func (*OperationResultFailed) IsApplied() bool { return false }

//json:status=skipped
type OperationResultSkipped struct{}

func (*OperationResultSkipped) Status() string  { return "skipped" }
func (*OperationResultSkipped) IsApplied() bool { return false }
