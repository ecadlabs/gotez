package shell

import (
	tz "github.com/ecadlabs/gotez"
)

type BlockHeader struct {
	Level          int32
	Proto          uint8
	Predecessor    *tz.BlockHash
	Timestamp      tz.Timestamp
	ValidationPass uint8
	OperationsHash *tz.OperationsHash
	Fitness        []byte `tz:"dyn"`
	Context        *tz.ContextHash
}
