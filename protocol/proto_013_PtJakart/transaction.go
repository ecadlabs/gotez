package proto_013_PtJakart

import (
	tz "github.com/ecadlabs/gotez"
)

type ToScRollup struct {
	ConsumedMilligas tz.BigUint
	InboxAfter       ScRollupInbox
}

type ScRollupInbox struct {
	Rollup                                 *tz.ScRollupAddress `tz:"dyn"`
	MessageCounter                         tz.BigInt
	NbMessagesInCommitmentPeriod           int64
	StartingLevelOfCurrentCommitmentPeriod int32
	Level                                  int32
	CurrentLevelHash                       *[32]byte
	OldLevelsMessages                      OldLevelsMessages
}

type OldLevelsMessages struct {
	Index        int32
	Content      *[32]byte
	BackPointers []byte `tz:"dyn"`
}

func (*ToScRollup) TransactionResultDestination() {}

type TransactionResultDestination interface {
	TransactionResultDestination()
}

type TransactionDestination interface {
	TransactionDestination()
}

type TxRollupDestination struct {
	*tz.TXRollupAddress
	Padding uint8
}

func (*TxRollupDestination) TransactionDestination() {}

type ScRollupDestination struct {
	*tz.ScRollupAddress
	Padding uint8
}

func (*ScRollupDestination) TransactionDestination() {}
