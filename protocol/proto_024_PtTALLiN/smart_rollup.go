package proto_024_PtTALLiN

import "github.com/ecadlabs/gotez/v2/protocol/proto_023_PtSeouLo"

type SmartRollupOriginate = proto_023_PtSeouLo.SmartRollupOriginate
type SmartRollupOriginateResult = proto_023_PtSeouLo.SmartRollupOriginateResult
type SmartRollupAddMessages = proto_023_PtSeouLo.SmartRollupAddMessages
type SmartRollupCement = proto_023_PtSeouLo.SmartRollupCement
type SmartRollupCementResult = proto_023_PtSeouLo.SmartRollupCementResult
type SmartRollupPublish = proto_023_PtSeouLo.SmartRollupPublish
type SmartRollupPublishResult = proto_023_PtSeouLo.SmartRollupPublishResult
type SmartRollupRefute = proto_023_PtSeouLo.SmartRollupRefute
type SmartRollupTimeout = proto_023_PtSeouLo.SmartRollupTimeout
type SmartRollupTimeoutResult = proto_023_PtSeouLo.SmartRollupTimeoutResult
type SmartRollupExecuteOutboxMessage = proto_023_PtSeouLo.SmartRollupExecuteOutboxMessage
type SmartRollupExecuteOutboxMessageResult = proto_023_PtSeouLo.SmartRollupExecuteOutboxMessageResult
type SmartRollupRecoverBond = proto_023_PtSeouLo.SmartRollupRecoverBond
type SmartRollupRecoverBondResult = proto_023_PtSeouLo.SmartRollupRecoverBondResult

//json:kind=OperationKind()
type SmartRollupOriginateContentsAndResult struct {
	SmartRollupOriginate
	Metadata ManagerMetadata[SmartRollupOriginateResult] `json:"metadata"`
}

func (*SmartRollupOriginateContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupOriginateContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type SmartRollupAddMessagesContentsAndResult struct {
	SmartRollupAddMessages
	Metadata ManagerMetadata[ConsumedGasResult] `json:"metadata"`
}

func (*SmartRollupAddMessagesContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupAddMessagesContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type SmartRollupCementContentsAndResult struct {
	SmartRollupCement
	Metadata ManagerMetadata[SmartRollupCementResult] `json:"metadata"`
}

func (*SmartRollupCementContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupCementContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type SmartRollupPublishContentsAndResult struct {
	SmartRollupPublish
	Metadata ManagerMetadata[SmartRollupPublishResult] `json:"metadata"`
}

func (*SmartRollupPublishContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupPublishContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type SmartRollupRefuteContentsAndResult struct {
	SmartRollupRefute
	Metadata ManagerMetadata[SmartRollupTimeoutResult] `json:"metadata"`
}

func (*SmartRollupRefuteContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupRefuteContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type SmartRollupTimeoutContentsAndResult struct {
	SmartRollupTimeout
	Metadata ManagerMetadata[SmartRollupTimeoutResult] `json:"metadata"`
}

func (*SmartRollupTimeoutContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupTimeoutContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type SmartRollupExecuteOutboxMessageContentsAndResult struct {
	SmartRollupExecuteOutboxMessage
	Metadata ManagerMetadata[SmartRollupExecuteOutboxMessageResult] `json:"metadata"`
}

func (*SmartRollupExecuteOutboxMessageContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupExecuteOutboxMessageContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type SmartRollupRecoverBondContentsAndResult struct {
	SmartRollupRecoverBond
	Metadata ManagerMetadata[SmartRollupRecoverBondResult] `json:"metadata"`
}

func (*SmartRollupRecoverBondContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupRecoverBondContentsAndResult) GetMetadata() any {
	return &op.Metadata
}
