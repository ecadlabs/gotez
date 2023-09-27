package proto_016_PtMumbai

import (
	"math/big"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/proto_015_PtLimaPt"
)

type ZkRollupOrigination = proto_015_PtLimaPt.ZkRollupOrigination
type ZkRollupPublish = proto_015_PtLimaPt.ZkRollupPublish

//json:kind=OperationKind()
type ZkRollupOriginationContentsAndResult struct {
	ZkRollupOrigination
	Metadata ManagerMetadata[ZkRollupPublishResult] `json:"metadata"`
}

func (*ZkRollupOriginationContentsAndResult) OperationContentsAndResult() {}
func (op *ZkRollupOriginationContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type ZkRollupPublishResultContents struct {
	BalanceUpdates
	ConsumedMilligas tz.BigUint `json:"consumed_milligas"`
	Size             tz.BigInt  `json:"size"`
}

func (r *ZkRollupPublishResultContents) GetConsumedMilligas() tz.BigUint { return r.ConsumedMilligas }
func (r *ZkRollupPublishResultContents) EstimateStorageSize(constants core.Constants) *big.Int {
	return r.Size.Int()
}

type ZkRollupPublishResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[ZkRollupPublishResult]{
		Variants: encoding.Variants[ZkRollupPublishResult]{
			0: (*core.OperationResultApplied[*ZkRollupPublishResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*ZkRollupPublishResultContents])(nil),
		},
	})
}

//json:kind=OperationKind()
type ZkRollupPublishContentsAndResult struct {
	ZkRollupPublish
	Metadata ManagerMetadata[ZkRollupPublishResult] `json:"metadata"`
}

func (*ZkRollupPublishContentsAndResult) OperationContentsAndResult() {}
func (op *ZkRollupPublishContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type ZkRollupUpdate struct {
	ManagerOperation
	ZkRollup *tz.ZkRollupAddress    `json:"zk_rollup"`
	Update   ZkRollupUpdateContents `json:"update"`
}

func (*ZkRollupUpdate) OperationKind() string { return "zk_rollup_update" }

type ZkRollupUpdateContents struct {
	PendingPis []*PendingPiElem `tz:"dyn" json:"pending_pis"`
	PrivatePis []*PrivatePiElem `tz:"dyn" json:"private_pis"`
	FeePi      FeePi            `json:"fee_pi"`
	Proof      tz.Bytes         `tz:"dyn" json:"proof"`
}

type PendingPiElem struct {
	Key string    `tz:"dyn" json:"key"`
	Pi  PendingPi `json:"pi"`
}

type ZkRollupScalar = tz.Bytes32

type PendingPi struct {
	NewState     tz.Bytes       `tz:"dyn" json:"new_state"`
	Fee          ZkRollupScalar `json:"fee"`
	ExitValidity bool           `json:"exit_validity"`
}

type PrivatePiElem struct {
	Key string    `tz:"dyn" json:"key"`
	Pi  PrivatePi `json:"pi"`
}

type PrivatePi struct {
	NewState tz.Bytes       `tz:"dyn" json:"new_state"`
	Fee      ZkRollupScalar `json:"fee"`
}

type FeePi struct {
	NewState tz.Bytes `tz:"dyn" json:"new_state"`
}

type ZkRollupUpdateResultContents struct {
	BalanceUpdates
	ConsumedMilligas    tz.BigUint `json:"consumed_milligas"`
	PaidStorageSizeDiff tz.BigInt  `json:"paid_storage_size_diff"`
}

func (r *ZkRollupUpdateResultContents) GetConsumedMilligas() tz.BigUint { return r.ConsumedMilligas }
func (r *ZkRollupUpdateResultContents) GetPaidStorageSizeDiff() tz.BigInt {
	return r.PaidStorageSizeDiff
}
func (r *ZkRollupUpdateResultContents) EstimateStorageSize(constants core.Constants) *big.Int {
	return r.PaidStorageSizeDiff.Int()
}

type ZkRollupUpdateResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[ZkRollupUpdateResult]{
		Variants: encoding.Variants[ZkRollupUpdateResult]{
			0: (*core.OperationResultApplied[*ZkRollupUpdateResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*ZkRollupUpdateResultContents])(nil),
		},
	})
}

//json:kind=OperationKind()
type ZkRollupUpdateContentsAndResult struct {
	ZkRollupUpdate
	Metadata ManagerMetadata[ZkRollupUpdateResult] `json:"metadata"`
}

func (*ZkRollupUpdateContentsAndResult) OperationContentsAndResult() {}
func (op *ZkRollupUpdateContentsAndResult) GetMetadata() any {
	return &op.Metadata
}
