package protocol

import (
	tz "github.com/ecadlabs/gotez"
)

type ShellHeader struct {
	Level          int32
	Proto          uint8
	Predecessor    *tz.BlockHash
	Timestamp      tz.Timestamp
	ValidationPass uint8
	OperationsHash *tz.OperationsHash
	Fitness        []byte `tz:"dyn"`
	Context        *tz.ContextHash
}

type TenderbakeBlockHeader struct {
	ShellHeader
	PayloadHash               *tz.BlockPayloadHash
	PayloadRound              int32
	ProofOfWorkNonce          *[tz.ProofOfWorkNonceBytesLen]byte
	SeedNonceHash             tz.Option[*tz.CycleNonceHash]
	LiquidityBakingToggleVote uint8
}
