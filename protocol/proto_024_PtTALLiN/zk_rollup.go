package proto_024_PtTALLiN

import "github.com/ecadlabs/gotez/v2/protocol/proto_023_PtSeouLo"

type ZkRollupOrigination = proto_023_PtSeouLo.ZkRollupOrigination
type ZkRollupOriginationResult = proto_023_PtSeouLo.ZkRollupOriginationResult
type ZkRollupPublish = proto_023_PtSeouLo.ZkRollupPublish
type ZkRollupPublishResult = proto_023_PtSeouLo.ZkRollupPublishResult
type ZkRollupUpdate = proto_023_PtSeouLo.ZkRollupUpdate
type ZkRollupUpdateResult = proto_023_PtSeouLo.ZkRollupUpdateResult

//json:kind=OperationKind()
type ZkRollupOriginationContentsAndResult struct {
	ZkRollupOrigination
	Metadata ManagerMetadata[ZkRollupOriginationResult] `json:"metadata"`
}

func (*ZkRollupOriginationContentsAndResult) OperationContentsAndResult() {}
func (op *ZkRollupOriginationContentsAndResult) GetMetadata() any {
	return &op.Metadata
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
type ZkRollupUpdateContentsAndResult struct {
	ZkRollupUpdate
	Metadata ManagerMetadata[ZkRollupUpdateResult] `json:"metadata"`
}

func (*ZkRollupUpdateContentsAndResult) OperationContentsAndResult() {}
func (op *ZkRollupUpdateContentsAndResult) GetMetadata() any {
	return &op.Metadata
}
