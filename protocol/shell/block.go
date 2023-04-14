package shell

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/protocol/proto"
)

type BlockHeader struct {
	Level          int32
	Proto          proto.Protocol
	Predecessor    *tz.BlockHash
	Timestamp      tz.Timestamp
	ValidationPass uint8
	OperationsHash *tz.OperationsHash
	Fitness        []byte `tz:"dyn"`
	Context        *tz.ContextHash
}
