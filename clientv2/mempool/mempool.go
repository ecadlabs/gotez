package mempool

//go:generate go run generate.go

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/latest"
)

type PendingOperationsResponse struct {
	Validated     []*PendingOperationsList            `tz:"dyn" json:"validated"`
	Refused       []*PendingOperationsListWithError   `tz:"dyn" json:"refused"`
	Outdated      []*PendingOperationsListWithError   `tz:"dyn" json:"outdated"`
	BranchRefused []*PendingOperationsListWithError   `tz:"dyn" json:"branch_refused"`
	BranchDelayed []*PendingOperationsListWithError   `tz:"dyn" json:"branch_delayed"`
	Unprocessed   []*UnprocessedPendingOperationsList `tz:"dyn" json:"unprocessed"`
}

type PendingOperationsList struct {
	Hash     *tz.OperationHash                                          `json:"hash"`
	Branch   *tz.BlockHash                                              `json:"branch"`
	Contents []*core.OperationWithoutMetadata[latest.OperationContents] `tz:"dyn" json:"contents"`
}

type PendingOperationsListWithError struct {
	Hash     *tz.OperationHash                                          `json:"hash"`
	Branch   *tz.BlockHash                                              `tz:"dyn" json:"branch"`
	Contents []*core.OperationWithoutMetadata[latest.OperationContents] `tz:"dyn" json:"contents"`
	Error    []byte                                                     `tz:"dyn" json:"error"`
}

type UnprocessedPendingOperationsList struct {
	Hash     *tz.OperationHash                                          `json:"hash"`
	Branch   *tz.BlockHash                                              `tz:"dyn" json:"branch"`
	Contents []*core.OperationWithoutMetadata[latest.OperationContents] `tz:"dyn" json:"contents"`
}

type OperationsList struct {
	UnprocessedPendingOperationsList
	Error []byte `tz:"dyn" json:"error"`
}

type MonitorResponse struct {
	Contents []*OperationsList `tz:"dyn" json:"contents"`
}
