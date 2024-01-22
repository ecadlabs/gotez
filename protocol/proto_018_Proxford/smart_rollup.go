package proto_018_Proxford

import (
	"math/big"
	"strconv"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/core/expression"
	"github.com/ecadlabs/gotez/v2/protocol/proto_016_PtMumbai"
	"github.com/ecadlabs/gotez/v2/protocol/proto_017_PtNairob"
)

type SmartRollupAddMessages = proto_016_PtMumbai.SmartRollupAddMessages
type SmartRollupPublish = proto_016_PtMumbai.SmartRollupPublish
type SmartRollupRefute = proto_016_PtMumbai.SmartRollupRefute
type SmartRollupTimeout = proto_016_PtMumbai.SmartRollupTimeout
type SmartRollupExecuteOutboxMessage = proto_016_PtMumbai.SmartRollupExecuteOutboxMessage
type SmartRollupRecoverBond = proto_016_PtMumbai.SmartRollupRecoverBond

type SmartRollupCementResult = proto_017_PtNairob.SmartRollupCementResult
type GameStatus = proto_016_PtMumbai.GameStatus

type PVMKind uint8

const (
	PVMArith PVMKind = iota
	PVM_WASM_2_0_0
	PVM_RISCV
)

func (k PVMKind) String() string {
	switch k {
	case PVMArith:
		return "arith"
	case PVM_WASM_2_0_0:
		return "wasm_2_0_0"
	case PVM_RISCV:
		return "riscv"
	default:
		return strconv.FormatInt(int64(k), 10)
	}
}

func (k PVMKind) MarshalText() (text []byte, err error) {
	return []byte(k.String()), nil
}

//json:kind=OperationKind()
type SmartRollupOriginate struct {
	ManagerOperation
	PVMKind          PVMKind                         `json:"pvm_kind"`
	Kernel           tz.Bytes                        `tz:"dyn" json:"kernel"`
	OriginationProof tz.Bytes                        `tz:"dyn" json:"origination_proof"`
	ParametersTy     expression.Expression           `tz:"dyn" json:"parameters_ty"`
	Whitelist        tz.Option[SmartRollupWhitelist] `json:"whitelist"`
}

func (*SmartRollupOriginate) OperationKind() string { return "smart_rollup_originate" }

type SmartRollupWhitelist struct {
	Contents []tz.PublicKeyHash `tz:"dyn" json:"contents"`
}

//json:kind=OperationKind()
type SmartRollupCement struct {
	ManagerOperation
	Rollup *tz.SmartRollupAddress `json:"rollup"`
}

func (*SmartRollupCement) OperationKind() string { return "smart_rollup_cement" }

type SmartRollupOriginateResultContents struct {
	BalanceUpdates
	Address               *tz.SmartRollupAddress        `json:"address"`
	GenesisCommitmentHash *tz.SmartRollupCommitmentHash `json:"genesis_commitment_hash"`
	ConsumedMilligas      tz.BigUint                    `json:"consumed_milligas"`
	Size                  tz.BigInt                     `json:"size"`
}

func (r *SmartRollupOriginateResultContents) GetConsumedMilligas() tz.BigUint {
	return r.ConsumedMilligas
}
func (r *SmartRollupOriginateResultContents) EstimateStorageSize(constants core.Constants) *big.Int {
	return r.Size.Int()
}

type SmartRollupOriginateResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SmartRollupOriginateResult]{
		Variants: encoding.Variants[SmartRollupOriginateResult]{
			0: (*core.OperationResultApplied[*SmartRollupOriginateResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*SmartRollupOriginateResultContents])(nil),
		},
	})
}

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

type SmartRollupPublishResultContents struct {
	ConsumedMilligas tz.BigUint                    `json:"consumed_milligas"`
	StakedHash       *tz.SmartRollupCommitmentHash `json:"staked_hash"`
	PublishedAtLevel int32                         `json:"published_at_level"`
	BalanceUpdates
}

func (r *SmartRollupPublishResultContents) GetConsumedMilligas() tz.BigUint {
	return r.ConsumedMilligas
}

type SmartRollupPublishResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SmartRollupPublishResult]{
		Variants: encoding.Variants[SmartRollupPublishResult]{
			0: (*core.OperationResultApplied[*SmartRollupPublishResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*SmartRollupPublishResultContents])(nil),
		},
	})
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

type SmartRollupTimeoutResultContents struct {
	ConsumedMilligas tz.BigUint `json:"consumed_milligas"`
	GameStatus       GameStatus `json:"game_status"`
	BalanceUpdates
}

func (r *SmartRollupTimeoutResultContents) GetConsumedMilligas() tz.BigUint {
	return r.ConsumedMilligas
}

type SmartRollupTimeoutResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SmartRollupTimeoutResult]{
		Variants: encoding.Variants[SmartRollupTimeoutResult]{
			0: (*core.OperationResultApplied[*SmartRollupTimeoutResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*SmartRollupTimeoutResultContents])(nil),
		},
	})
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
type SmartRollupRefuteContentsAndResult struct {
	SmartRollupRefute
	Metadata ManagerMetadata[SmartRollupTimeoutResult] `json:"metadata"`
}

func (*SmartRollupRefuteContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupRefuteContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type SmartRollupExecuteOutboxMessageResultContents struct {
	BalanceUpdates
	TicketUpdates       []*TicketReceipt `tz:"dyn" json:"ticket_updates"`
	ConsumedMilligas    tz.BigUint       `json:"consumed_milligas"`
	PaidStorageSizeDiff tz.BigInt        `json:"paid_storage_size_diff"`
}

func (r *SmartRollupExecuteOutboxMessageResultContents) GetConsumedMilligas() tz.BigUint {
	return r.ConsumedMilligas
}

func (r *SmartRollupExecuteOutboxMessageResultContents) GetPaidStorageSizeDiff() tz.BigInt {
	return r.PaidStorageSizeDiff
}

func (r *SmartRollupExecuteOutboxMessageResultContents) EstimateStorageSize(constants core.Constants) *big.Int {
	return r.PaidStorageSizeDiff.Int()
}

type SmartRollupExecuteOutboxMessageResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SmartRollupExecuteOutboxMessageResult]{
		Variants: encoding.Variants[SmartRollupExecuteOutboxMessageResult]{
			0: (*core.OperationResultApplied[*SmartRollupExecuteOutboxMessageResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*SmartRollupExecuteOutboxMessageResultContents])(nil),
		},
	})
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

type SmartRollupRecoverBondResultContents struct {
	BalanceUpdates
	ConsumedMilligas tz.BigUint `json:"consumed_milligas"`
}

func (r *SmartRollupRecoverBondResultContents) GetConsumedMilligas() tz.BigUint {
	return r.ConsumedMilligas
}

type SmartRollupRecoverBondResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SmartRollupRecoverBondResult]{
		Variants: encoding.Variants[SmartRollupRecoverBondResult]{
			0: (*core.OperationResultApplied[*SmartRollupRecoverBondResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*SmartRollupRecoverBondResultContents])(nil),
		},
	})
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

//json:kind=OperationKind()
type SmartRollupOriginateSuccessfulManagerResult struct {
	core.OperationResultApplied[*SmartRollupOriginateResultContents]
}

func (*SmartRollupOriginateSuccessfulManagerResult) OperationKind() string {
	return "smart_rollup_originate"
}
