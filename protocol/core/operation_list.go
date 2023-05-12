package core

import (
	tz "github.com/ecadlabs/gotez"
)

type OperationsList[T GroupContents] struct {
	Operations []*OperationsGroup[T] `tz:"dyn,dyn"` // yes, twice
}

type OperationsGroup[T GroupContents] struct {
	ChainID  *tz.ChainID
	Hash     *tz.BlockHash
	Branch   *tz.BlockHash `tz:"dyn"`
	Contents T             `tz:"dyn"`
}
