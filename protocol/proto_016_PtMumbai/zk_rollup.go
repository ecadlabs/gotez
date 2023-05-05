package proto_016_PtMumbai

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/proto_015_PtLimaPt"
)

type ZkRollupOrigination = proto_015_PtLimaPt.ZkRollupOrigination
type ZkRollupPublish = proto_015_PtLimaPt.ZkRollupPublish

type ZkRollupOriginationContentsAndResult[T core.BalanceUpdateKind] struct {
	ZkRollupOrigination
	Metadata ManagerMetadata[ZkRollupPublishResult, T]
}

func (*ZkRollupOriginationContentsAndResult[T]) OperationContentsAndResult() {}

type ZkRollupPublishResultContents[T core.BalanceUpdateKind] struct {
	BalanceUpdates   []*BalanceUpdate[T] `tz:"dyn"`
	ConsumedMilligas tz.BigUint
	Size             tz.BigInt
}

func (ZkRollupPublishResultContents[T]) SuccessfulManagerOperationResult() {}
func (ZkRollupPublishResultContents[T]) OperationKind() string {
	return "zk_rollup_publish"
}

type ZkRollupPublishResult interface {
	ZkRollupPublishResult()
	core.OperationResult
}

type ZkRollupPublishResultApplied[T core.BalanceUpdateKind] struct {
	core.OperationResultApplied[ZkRollupPublishResultContents[T]]
}

func (*ZkRollupPublishResultApplied[T]) ZkRollupPublishResult() {}

type ZkRollupPublishResultBacktracked[T core.BalanceUpdateKind] struct {
	core.OperationResultBacktracked[ZkRollupPublishResultContents[T]]
}

func (*ZkRollupPublishResultBacktracked[T]) ZkRollupPublishResult() {}

type ZkRollupPublishResultFailed struct{ core.OperationResultFailed }

func (*ZkRollupPublishResultFailed) ZkRollupPublishResult() {}

type ZkRollupPublishResultSkipped struct{ core.OperationResultSkipped }

func (*ZkRollupPublishResultSkipped) ZkRollupPublishResult() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[ZkRollupPublishResult]{
		Variants: encoding.Variants[ZkRollupPublishResult]{
			0: (*ZkRollupPublishResultApplied[BalanceUpdateKind])(nil),
			1: (*ZkRollupPublishResultFailed)(nil),
			2: (*ZkRollupPublishResultSkipped)(nil),
			3: (*ZkRollupPublishResultBacktracked[BalanceUpdateKind])(nil),
		},
	})
}

type ZkRollupPublishContentsAndResult[T core.BalanceUpdateKind] struct {
	ZkRollupPublish
	Metadata ManagerMetadata[ZkRollupPublishResult, T]
}

func (*ZkRollupPublishContentsAndResult[T]) OperationContentsAndResult() {}

type ZkRollupUpdate struct {
	ManagerOperation
	ZkRollup *tz.ZkRollupAddress
	Update   ZkRollupUpdateContents
}

func (*ZkRollupUpdate) OperationKind() string { return "zk_rollup_update" }

type ZkRollupUpdateContents struct {
	PendingPis []*PendingPiElem `tz:"dyn"`
	PrivatePis []*PrivatePiElem `tz:"dyn"`
	FeePi      FeePi
	Proof      []byte `tz:"dyn"`
}

type PendingPiElem struct {
	Key string `tz:"dyn"`
	Pi  PendingPi
}

type ZkRollupScalar [32]byte

type PendingPi struct {
	NewState     []byte `tz:"dyn"`
	Fee          ZkRollupScalar
	ExitValidity bool
}

type PrivatePiElem struct {
	Key string `tz:"dyn"`
	Pi  PrivatePi
}

type PrivatePi struct {
	NewState []byte `tz:"dyn"`
	Fee      ZkRollupScalar
}

type FeePi struct {
	NewState []byte `tz:"dyn"`
}

type ZkRollupUpdateResultContents[T core.BalanceUpdateKind] struct {
	BalanceUpdates      []*BalanceUpdate[T] `tz:"dyn"`
	ConsumedMilligas    tz.BigUint
	PaidStorageSizeDiff tz.BigInt
}

func (ZkRollupUpdateResultContents[T]) SuccessfulManagerOperationResult() {}
func (ZkRollupUpdateResultContents[T]) OperationKind() string {
	return "zk_rollup_update"
}

type ZkRollupUpdateResult interface {
	ZkRollupUpdateResult()
	core.OperationResult
}

type ZkRollupUpdateResultApplied[T core.BalanceUpdateKind] struct {
	core.OperationResultApplied[ZkRollupUpdateResultContents[T]]
}

func (*ZkRollupUpdateResultApplied[T]) ZkRollupUpdateResult() {}

type ZkRollupUpdateResultBacktracked[T core.BalanceUpdateKind] struct {
	core.OperationResultBacktracked[ZkRollupUpdateResultContents[T]]
}

func (*ZkRollupUpdateResultBacktracked[T]) ZkRollupUpdateResult() {}

type ZkRollupUpdateResultFailed struct{ core.OperationResultFailed }

func (*ZkRollupUpdateResultFailed) ZkRollupUpdateResult() {}

type ZkRollupUpdateResultSkipped struct{ core.OperationResultSkipped }

func (*ZkRollupUpdateResultSkipped) ZkRollupUpdateResult() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[ZkRollupUpdateResult]{
		Variants: encoding.Variants[ZkRollupUpdateResult]{
			0: (*ZkRollupUpdateResultApplied[BalanceUpdateKind])(nil),
			1: (*ZkRollupUpdateResultFailed)(nil),
			2: (*ZkRollupUpdateResultSkipped)(nil),
			3: (*ZkRollupUpdateResultBacktracked[BalanceUpdateKind])(nil),
		},
	})
}

type ZkRollupUpdateContentsAndResult[T core.BalanceUpdateKind] struct {
	ZkRollupUpdate
	Metadata ManagerMetadata[ZkRollupUpdateResult, T]
}

func (*ZkRollupUpdateContentsAndResult[T]) OperationContentsAndResult() {}
