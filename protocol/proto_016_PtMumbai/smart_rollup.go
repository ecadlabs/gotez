package proto_016_PtMumbai

import (
	"strconv"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/core/expression"
)

type PVMKind uint8

const (
	PVMArith PVMKind = iota
	PVM_WASM_2_0_0
)

func (k PVMKind) String() string {
	switch k {
	case PVMArith:
		return "arith"
	case PVM_WASM_2_0_0:
		return "wasm_2_0_0"
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
	PVMKind          PVMKind               `json:"pvm_kind"`
	Kernel           tz.Bytes              `tz:"dyn" json:"kernel"`
	OriginationProof tz.Bytes              `tz:"dyn" json:"origination_proof"`
	ParametersTy     expression.Expression `tz:"dyn" json:"parameters_ty"`
}

func (*SmartRollupOriginate) OperationKind() string { return "smart_rollup_originate" }

type SmartRollupOriginateResult interface {
	SmartRollupOriginateResult()
	core.ManagerOperationResult
}

type SmartRollupOriginateResultContents struct {
	BalanceUpdates
	Address               *tz.SmartRollupAddress        `json:"address"`
	GenesisCommitmentHash *tz.SmartRollupCommitmentHash `json:"genesis_commitment_hash"`
	ConsumedMilligas      tz.BigUint                    `json:"consumed_milligas"`
	Size                  tz.BigInt                     `json:"size"`
}

//json:kind=OperationKind()
type SmartRollupOriginateSuccessfulManagerResult struct {
	SmartRollupOriginateResultApplied
}

func (*SmartRollupOriginateSuccessfulManagerResult) OperationKind() string {
	return "smart_rollup_originate"
}

type SmartRollupOriginateResultApplied struct {
	core.OperationResultApplied[SmartRollupOriginateResultContents]
}

func (*SmartRollupOriginateResultApplied) SmartRollupOriginateResult() {}

type SmartRollupOriginateResultBacktracked struct {
	core.OperationResultBacktracked[SmartRollupOriginateResultContents]
}

func (*SmartRollupOriginateResultBacktracked) SmartRollupOriginateResult() {}

type SmartRollupOriginateResultFailed struct{ core.OperationResultFailed }

func (*SmartRollupOriginateResultFailed) SmartRollupOriginateResult() {}

type SmartRollupOriginateResultSkipped struct{ core.OperationResultSkipped }

func (*SmartRollupOriginateResultSkipped) SmartRollupOriginateResult() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SmartRollupOriginateResult]{
		Variants: encoding.Variants[SmartRollupOriginateResult]{
			0: (*SmartRollupOriginateResultApplied)(nil),
			1: (*SmartRollupOriginateResultFailed)(nil),
			2: (*SmartRollupOriginateResultSkipped)(nil),
			3: (*SmartRollupOriginateResultBacktracked)(nil),
		},
	})
}

type SmartRollupOriginateContentsAndResult struct {
	SmartRollupOriginate
	Metadata ManagerMetadata[SmartRollupOriginateResult] `json:"metadata"`
}

func (*SmartRollupOriginateContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupOriginateContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type SmartRollupAddMessages struct {
	ManagerOperation
	Message []core.Bytes `tz:"dyn" json:"message"`
}

func (*SmartRollupAddMessages) OperationKind() string { return "smart_rollup_add_messages" }

type SmartRollupAddMessagesContentsAndResult struct {
	SmartRollupAddMessages
	Metadata ManagerMetadata[ConsumedGasResult] `json:"metadata"`
}

func (*SmartRollupAddMessagesContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupAddMessagesContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type SmartRollupCement struct {
	ManagerOperation
	Rollup     *tz.SmartRollupAddress        `json:"rollup"`
	Commitment *tz.SmartRollupCommitmentHash `json:"commitment"`
}

func (*SmartRollupCement) OperationKind() string { return "smart_rollup_cement" }

type SmartRollupCementResultContents struct {
	ConsumedMilligas tz.BigUint `json:"consumed_milligas"`
	InboxLevel       int32      `json:"inbox_level"`
}

type SmartRollupCementResult interface {
	SmartRollupCementResult()
	core.ManagerOperationResult
}

type SmartRollupCementResultApplied struct {
	core.OperationResultApplied[SmartRollupCementResultContents]
}

func (*SmartRollupCementResultApplied) SmartRollupCementResult() {}

type SmartRollupCementResultBacktracked struct {
	core.OperationResultBacktracked[SmartRollupCementResultContents]
}

func (*SmartRollupCementResultBacktracked) SmartRollupCementResult() {}

type SmartRollupCementResultFailed struct{ core.OperationResultFailed }

func (*SmartRollupCementResultFailed) SmartRollupCementResult() {}

type SmartRollupCementResultSkipped struct{ core.OperationResultSkipped }

func (*SmartRollupCementResultSkipped) SmartRollupCementResult() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SmartRollupCementResult]{
		Variants: encoding.Variants[SmartRollupCementResult]{
			0: (*SmartRollupCementResultApplied)(nil),
			1: (*SmartRollupCementResultFailed)(nil),
			2: (*SmartRollupCementResultSkipped)(nil),
			3: (*SmartRollupCementResultBacktracked)(nil),
		},
	})
}

type SmartRollupCementContentsAndResult struct {
	SmartRollupCement
	Metadata ManagerMetadata[SmartRollupCementResult] `json:"metadata"`
}

func (*SmartRollupCementContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupCementContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type SmartRollupPublish struct {
	ManagerOperation
	Rollup     *tz.SmartRollupAddress `json:"rollup"`
	Commitment SmartRollupCommitment  `json:"commitment"`
}

func (*SmartRollupRefute) OperationKind() string { return "smart_rollup_refute" }

type SmartRollupCommitment struct {
	CompressedState *tz.SmartRollupStateHash      `json:"compressed_state"`
	InboxLevel      int32                         `json:"inbox_level"`
	Predecessor     *tz.SmartRollupCommitmentHash `json:"predecessor"`
	NumberOfTicks   int64                         `json:"number_of_ticks"`
}

type SmartRollupPublishResultContents struct {
	ConsumedMilligas tz.BigUint                    `json:"consumed_milligas"`
	StakedHash       *tz.SmartRollupCommitmentHash `json:"staked_hash"`
	PublishedAtLevel int32                         `json:"published_at_level"`
	BalanceUpdates
}

type SmartRollupPublishResult interface {
	SmartRollupPublishResult()
	core.ManagerOperationResult
}

type SmartRollupPublishResultApplied struct {
	core.OperationResultApplied[SmartRollupPublishResultContents]
}

func (*SmartRollupPublishResultApplied) SmartRollupPublishResult() {}

type SmartRollupPublishResultBacktracked struct {
	core.OperationResultBacktracked[SmartRollupPublishResultContents]
}

func (*SmartRollupPublishResultBacktracked) SmartRollupPublishResult() {}

type SmartRollupPublishResultFailed struct{ core.OperationResultFailed }

func (*SmartRollupPublishResultFailed) SmartRollupPublishResult() {}

type SmartRollupPublishResultSkipped struct{ core.OperationResultSkipped }

func (*SmartRollupPublishResultSkipped) SmartRollupPublishResult() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SmartRollupPublishResult]{
		Variants: encoding.Variants[SmartRollupPublishResult]{
			0: (*SmartRollupPublishResultApplied)(nil),
			1: (*SmartRollupPublishResultFailed)(nil),
			2: (*SmartRollupPublishResultSkipped)(nil),
			3: (*SmartRollupPublishResultBacktracked)(nil),
		},
	})
}

type SmartRollupPublishContentsAndResult struct {
	SmartRollupPublish
	Metadata ManagerMetadata[SmartRollupPublishResult] `json:"metadata"`
}

func (*SmartRollupPublishContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupPublishContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type SmartRollupRefute struct {
	ManagerOperation
	Rollup     *tz.SmartRollupAddress `json:"rollup"`
	Opponent   tz.PublicKeyHash       `json:"opponent"`
	Refutation SmartRollupRefutation  `json:"refutation"`
}

func (*SmartRollupPublish) OperationKind() string { return "smart_rollup_publish" }

type SmartRollupRefutation interface {
	RefutationKind() string
}

type RefutationStart struct {
	PlayerCommitmentHash   *tz.SmartRollupCommitmentHash `json:"player_commitment_hash"`
	OpponentCommitmentHash *tz.SmartRollupCommitmentHash `json:"opponent_commitment_hash"`
}

func (*RefutationStart) RefutationKind() string { return "start" }

type RefutationMove struct {
	Choice tz.BigUint     `json:"choice"`
	Step   RefutationStep `json:"step"`
}

func (*RefutationMove) RefutationKind() string { return "move" }

func init() {
	encoding.RegisterEnum(&encoding.Enum[SmartRollupRefutation]{
		Variants: encoding.Variants[SmartRollupRefutation]{
			0: (*RefutationStart)(nil),
			1: (*RefutationMove)(nil),
		},
	})
}

type RefutationStep interface {
	RefutationStepKind() string
}

type RefutationStepDissection struct {
	Contents []RefutationStepDissectionElem `tz:"dyn" json:"contents"`
}

func (*RefutationStepDissection) RefutationStepKind() string { return "dissection" }

type RefutationStepDissectionElem struct {
	State tz.Option[*tz.SmartRollupStateHash] `json:"state"`
	Tick  tz.BigUint                          `json:"tick"`
}

type RefutationStepProof struct {
	PVMStep    tz.Bytes                   `tz:"dyn" json:"pvm_step"`
	InputProof tz.Option[RefutationProof] `json:"input_proof"`
}

func (*RefutationStepProof) RefutationStepKind() string { return "proof" }

func init() {
	encoding.RegisterEnum(&encoding.Enum[RefutationStep]{
		Variants: encoding.Variants[RefutationStep]{
			0: (*RefutationStepDissection)(nil),
			1: (*RefutationStepProof)(nil),
		},
	})
}

type RefutationProof interface {
	RefutationProof()
}

type RefutationProofInbox struct {
	Level           int32      `json:"level"`
	MessageCounter  tz.BigUint `json:"message_counter"`
	SerializedProof tz.Bytes   `tz:"dyn" json:"serialized_proof"`
}

func (*RefutationProofInbox) RefutationProof() {}

type RefutationProofReveal struct {
	RevealProof RevealProof `json:"reveal_proof"`
}

func (*RefutationProofReveal) RefutationProof() {}

type RefutationProofFirstInput struct{}

func (RefutationProofFirstInput) RefutationProof() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[RefutationProof]{
		Variants: encoding.Variants[RefutationProof]{
			0: (*RefutationProofInbox)(nil),
			1: (*RefutationProofReveal)(nil),
			2: RefutationProofFirstInput{},
		},
	})
}

type RevealProof interface {
	RevealProof()
}

type RevealProofRawData struct {
	RawData tz.Bytes `tz:"dyn" json:"raw_data"`
}

func (RevealProofRawData) RevealProof() {}

type RevealProofMetadata struct{}

func (RevealProofMetadata) RevealProof() {}

type RevealProofDALPage struct {
	DALPageID `json:"dal_page_id"`
	DALProof  tz.Bytes `tz:"dyn" json:"dal_proof"`
}

func (*RevealProofDALPage) RevealProof() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[RevealProof]{
		Variants: encoding.Variants[RevealProof]{
			0: RevealProofRawData{},
			1: RevealProofMetadata{},
			2: (*RevealProofDALPage)(nil),
		},
	})
}

type DALPageID struct {
	PublishedLevel int32 `json:"published_level"`
	SlotIndex      uint8 `json:"slot_index"`
	PageIndex      int16 `json:"page_index"`
}

type SmartRollupTimeoutResultContents struct {
	ConsumedMilligas tz.BigUint `json:"consumed_milligas"`
	GameStatus       GameStatus `json:"game_status"`
	BalanceUpdates
}

type GameStatus interface {
	GameStatusKind() string
}

type GameStatusOngoing struct{}

func (GameStatusOngoing) GameStatusKind() string { return "ongoing" }

type GameStatusEnded struct {
	Result GameStatusResult `json:"result"`
}

func (GameStatusEnded) GameStatusKind() string { return "ended" }

func init() {
	encoding.RegisterEnum(&encoding.Enum[GameStatus]{
		Variants: encoding.Variants[GameStatus]{
			0: GameStatusOngoing{},
			1: GameStatusEnded{},
		},
	})
}

type GameStatusResult interface {
	GameStatusResultKind() string
}

type GameStatusResultLoser struct {
	Reason LooseReason      `json:"reason"`
	Player tz.PublicKeyHash `json:"player"`
}

func (*GameStatusResultLoser) GameStatusResultKind() string { return "loser" }

type GameStatusResultDraw struct{}

func (GameStatusResultDraw) GameStatusResultKind() string { return "draw" }

func init() {
	encoding.RegisterEnum(&encoding.Enum[GameStatusResult]{
		Variants: encoding.Variants[GameStatusResult]{
			0: (*GameStatusResultLoser)(nil),
			1: GameStatusResultDraw{},
		},
	})
}

type LooseReason uint8

const (
	LooseReasonConflictResolved LooseReason = iota
	LooseReasonTimeout
)

type SmartRollupTimeoutResult interface {
	SmartRollupTimeoutResult()
	core.ManagerOperationResult
}

type SmartRollupTimeoutResultApplied struct {
	core.OperationResultApplied[SmartRollupTimeoutResultContents]
}

func (*SmartRollupTimeoutResultApplied) SmartRollupTimeoutResult() {}

type SmartRollupTimeoutResultBacktracked struct {
	core.OperationResultBacktracked[SmartRollupTimeoutResultContents]
}

func (*SmartRollupTimeoutResultBacktracked) SmartRollupTimeoutResult() {}

type SmartRollupTimeoutResultFailed struct{ core.OperationResultFailed }

func (*SmartRollupTimeoutResultFailed) SmartRollupTimeoutResult() {}

type SmartRollupTimeoutResultSkipped struct{ core.OperationResultSkipped }

func (*SmartRollupTimeoutResultSkipped) SmartRollupTimeoutResult() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SmartRollupTimeoutResult]{
		Variants: encoding.Variants[SmartRollupTimeoutResult]{
			0: (*SmartRollupTimeoutResultApplied)(nil),
			1: (*SmartRollupTimeoutResultFailed)(nil),
			2: (*SmartRollupTimeoutResultSkipped)(nil),
			3: (*SmartRollupTimeoutResultBacktracked)(nil),
		},
	})
}

type SmartRollupRefuteContentsAndResult struct {
	SmartRollupRefute
	Metadata ManagerMetadata[SmartRollupTimeoutResult] `json:"metadata"`
}

func (*SmartRollupRefuteContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupRefuteContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type SmartRollupTimeout struct {
	ManagerOperation
	Rollup  *tz.SmartRollupAddress `json:"rollup"`
	Stakers SmartRollupStakers     `json:"stakers"`
}

type SmartRollupStakers struct {
	Alice tz.PublicKeyHash `json:"alice"`
	Bob   tz.PublicKeyHash `json:"bob"`
}

func (*SmartRollupTimeout) OperationKind() string { return "smart_rollup_timeout" }

type SmartRollupTimeoutContentsAndResult struct {
	SmartRollupTimeout
	Metadata ManagerMetadata[SmartRollupTimeoutResult] `json:"metadata"`
}

func (*SmartRollupTimeoutContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupTimeoutContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type SmartRollupExecuteOutboxMessage struct {
	ManagerOperation
	Rollup             *tz.SmartRollupAddress        `json:"rollup"`
	CementedCommitment *tz.SmartRollupCommitmentHash `json:"cemented_commitment"`
	OutputProof        tz.Bytes                      `tz:"dyn" json:"output_proof"`
}

func (*SmartRollupExecuteOutboxMessage) OperationKind() string {
	return "smart_rollup_execute_outbox_message"
}

type SmartRollupExecuteOutboxMessageResultContents struct {
	BalanceUpdates
	TicketUpdates       []*TicketReceipt `tz:"dyn" json:"ticket_updates"`
	ConsumedMilligas    tz.BigUint       `json:"consumed_milligas"`
	PaidStorageSizeDiff tz.BigInt        `json:"paid_storage_size_diff"`
}

type SmartRollupExecuteOutboxMessageResult interface {
	SmartRollupExecuteOutboxMessageResult()
	core.ManagerOperationResult
}

type SmartRollupExecuteOutboxMessageResultApplied struct {
	core.OperationResultApplied[SmartRollupExecuteOutboxMessageResultContents]
}

func (*SmartRollupExecuteOutboxMessageResultApplied) SmartRollupExecuteOutboxMessageResult() {}

type SmartRollupExecuteOutboxMessageResultBacktracked struct {
	core.OperationResultBacktracked[SmartRollupExecuteOutboxMessageResultContents]
}

func (*SmartRollupExecuteOutboxMessageResultBacktracked) SmartRollupExecuteOutboxMessageResult() {}

type SmartRollupExecuteOutboxMessageResultFailed struct{ core.OperationResultFailed }

func (*SmartRollupExecuteOutboxMessageResultFailed) SmartRollupExecuteOutboxMessageResult() {}

type SmartRollupExecuteOutboxMessageResultSkipped struct{ core.OperationResultSkipped }

func (*SmartRollupExecuteOutboxMessageResultSkipped) SmartRollupExecuteOutboxMessageResult() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SmartRollupExecuteOutboxMessageResult]{
		Variants: encoding.Variants[SmartRollupExecuteOutboxMessageResult]{
			0: (*SmartRollupExecuteOutboxMessageResultApplied)(nil),
			1: (*SmartRollupExecuteOutboxMessageResultFailed)(nil),
			2: (*SmartRollupExecuteOutboxMessageResultSkipped)(nil),
			3: (*SmartRollupExecuteOutboxMessageResultBacktracked)(nil),
		},
	})
}

type SmartRollupExecuteOutboxMessageContentsAndResult struct {
	SmartRollupExecuteOutboxMessage
	Metadata ManagerMetadata[SmartRollupExecuteOutboxMessageResult] `json:"metadata"`
}

func (*SmartRollupExecuteOutboxMessageContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupExecuteOutboxMessageContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type SmartRollupRecoverBond struct {
	ManagerOperation
	Rollup *tz.SmartRollupAddress `json:"rollup"`
	Staker tz.PublicKeyHash       `json:"staker"`
}

func (*SmartRollupRecoverBond) OperationKind() string { return "smart_rollup_recover_bond" }

type SmartRollupRecoverBondResultContents struct {
	BalanceUpdates
	ConsumedMilligas tz.BigUint `json:"consumed_milligas"`
}

type SmartRollupRecoverBondResult interface {
	SmartRollupRecoverBondResult()
	core.ManagerOperationResult
}

type SmartRollupRecoverBondResultApplied struct {
	core.OperationResultApplied[SmartRollupRecoverBondResultContents]
}

func (*SmartRollupRecoverBondResultApplied) SmartRollupRecoverBondResult() {}

type SmartRollupRecoverBondResultBacktracked struct {
	core.OperationResultBacktracked[SmartRollupRecoverBondResultContents]
}

func (*SmartRollupRecoverBondResultBacktracked) SmartRollupRecoverBondResult() {}

type SmartRollupRecoverBondResultFailed struct{ core.OperationResultFailed }

func (*SmartRollupRecoverBondResultFailed) SmartRollupRecoverBondResult() {}

type SmartRollupRecoverBondResultSkipped struct{ core.OperationResultSkipped }

func (*SmartRollupRecoverBondResultSkipped) SmartRollupRecoverBondResult() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SmartRollupRecoverBondResult]{
		Variants: encoding.Variants[SmartRollupRecoverBondResult]{
			0: (*SmartRollupRecoverBondResultApplied)(nil),
			1: (*SmartRollupRecoverBondResultFailed)(nil),
			2: (*SmartRollupRecoverBondResultSkipped)(nil),
			3: (*SmartRollupRecoverBondResultBacktracked)(nil),
		},
	})
}

type SmartRollupRecoverBondContentsAndResult struct {
	SmartRollupRecoverBond
	Metadata ManagerMetadata[SmartRollupRecoverBondResult] `json:"metadata"`
}

func (*SmartRollupRecoverBondContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupRecoverBondContentsAndResult) GetMetadata() any {
	return &op.Metadata
}
