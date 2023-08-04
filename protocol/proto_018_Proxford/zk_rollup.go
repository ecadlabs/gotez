package proto_018_Proxford

import (
	"math/big"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/proto_015_PtLimaPt"
	"github.com/ecadlabs/gotez/v2/protocol/proto_016_PtMumbai"
)

type ZkRollupOrigination = proto_015_PtLimaPt.ZkRollupOrigination
type ZkRollupPublish = proto_015_PtLimaPt.ZkRollupPublish
type ZkRollupUpdate = proto_016_PtMumbai.ZkRollupUpdate

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

type ZkRollupPublishContentsAndResult struct {
	ZkRollupPublish
	Metadata ManagerMetadata[ZkRollupPublishResult] `json:"metadata"`
}

func (*ZkRollupPublishContentsAndResult) OperationContentsAndResult() {}
func (op *ZkRollupPublishContentsAndResult) GetMetadata() any {
	return &op.Metadata
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

type ZkRollupUpdateContentsAndResult struct {
	ZkRollupUpdate
	Metadata ManagerMetadata[ZkRollupUpdateResult] `json:"metadata"`
}

func (*ZkRollupUpdateContentsAndResult) OperationContentsAndResult() {}
func (op *ZkRollupUpdateContentsAndResult) GetMetadata() any {
	return &op.Metadata
}
