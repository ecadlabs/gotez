package proto_016_PtMumbai

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/core/expression"
	"github.com/ecadlabs/gotez/protocol/proto_015_PtLimaPt"
)

type PVMKind uint8

const (
	PVMArith PVMKind = iota
	PVM_WASM_2_0_0
)

type SmartRollupOriginate struct {
	ManagerOperation
	PVMKind          PVMKind
	Kernel           []byte                `tz:"dyn"`
	OriginationProof []byte                `tz:"dyn"`
	ParametersTy     expression.Expression `tz:"dyn"`
}

func (*SmartRollupOriginate) OperationKind() string { return "smart_rollup_originate" }

type SmartRollupOriginateResult interface {
	SmartRollupOriginateResult()
	core.ManagerOperationResult
}

type SmartRollupOriginateResultContents struct {
	BalanceUpdates        []*BalanceUpdate `tz:"dyn"`
	Address               *tz.SmartRollupAddress
	GenesisCommitmentHash *tz.SmartRollupCommitmentHash
	ConsumedMilligas      tz.BigUint
	Size                  tz.BigInt
}

func (SmartRollupOriginateResultContents) SuccessfulManagerOperationResult() {}
func (SmartRollupOriginateResultContents) OperationKind() string {
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
	Metadata ManagerMetadata[SmartRollupOriginateResult]
}

func (*SmartRollupOriginateContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupOriginateContentsAndResult) OperationContents() core.OperationContents {
	return &op.SmartRollupOriginate
}

type SmartRollupAddMessages struct {
	ManagerOperation
	Message []core.Bytes `tz:"dyn"`
}

func (*SmartRollupAddMessages) OperationKind() string { return "smart_rollup_add_messages" }

type SmartRollupAddMessagesContentsAndResult struct {
	SmartRollupAddMessages
	Metadata ManagerMetadata[ConsumedGasResult]
}

func (*SmartRollupAddMessagesContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupAddMessagesContentsAndResult) OperationContents() core.OperationContents {
	return &op.SmartRollupAddMessages
}

type SmartRollupCement struct {
	ManagerOperation
	Rollup     *tz.SmartRollupAddress
	Commitment *tz.SmartRollupCommitmentHash
}

func (*SmartRollupCement) OperationKind() string { return "smart_rollup_cement" }

type SmartRollupCementResultContents struct {
	ConsumedMilligas tz.BigUint
	InboxLevel       int32
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
	Metadata ManagerMetadata[SmartRollupCementResult]
}

func (*SmartRollupCementContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupCementContentsAndResult) OperationContents() core.OperationContents {
	return &op.SmartRollupCement
}

type SmartRollupPublish struct {
	ManagerOperation
	Rollup     *tz.SmartRollupAddress
	Commitment SmartRollupCommitment
}

func (*SmartRollupRefute) OperationKind() string { return "smart_rollup_refute" }

type SmartRollupCommitment struct {
	CompressedState *tz.SmartRollupStateHash
	InboxLevel      int32
	Predecessor     *tz.SmartRollupCommitmentHash
	NumberOfTicks   int64
}

type SmartRollupPublishResultContents struct {
	ConsumedMilligas tz.BigUint
	StakedHash       *tz.SmartRollupCommitmentHash
	PublishedAtLevel int32
	BalanceUpdates   []*BalanceUpdate `tz:"dyn"`
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
	Metadata ManagerMetadata[SmartRollupPublishResult]
}

func (*SmartRollupPublishContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupPublishContentsAndResult) OperationContents() core.OperationContents {
	return &op.SmartRollupPublish
}

type SmartRollupRefute struct {
	ManagerOperation
	Rollup     *tz.SmartRollupAddress
	Opponent   tz.PublicKeyHash
	Refutation SmartRollupRefutation
}

func (*SmartRollupPublish) OperationKind() string { return "smart_rollup_publish" }

type SmartRollupRefutation interface {
	RefutationKind() string
}

type RefutationStart struct {
	PlayerCommitmentHash   *tz.SmartRollupCommitmentHash
	OpponentCommitmentHash *tz.SmartRollupCommitmentHash
}

func (*RefutationStart) RefutationKind() string { return "start" }

type RefutationMove struct {
	Choice tz.BigUint
	Step   RefutationStep
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
	Contents []RefutationStepDissectionElem `tz:"dyn"`
}

func (*RefutationStepDissection) RefutationStepKind() string { return "dissection" }

type RefutationStepDissectionElem struct {
	State tz.Option[*tz.SmartRollupStateHash]
	Tick  tz.BigUint
}

type RefutationStepProof struct {
	PVMStep    []byte `tz:"dyn"`
	InputProof tz.Option[RefutationProof]
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
	Level           int32
	MessageCounter  tz.BigUint
	SerializedProof []byte `tz:"dyn"`
}

func (*RefutationProofInbox) RefutationProof() {}

type RefutationProofReveal struct {
	RevealProof RevealProof
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
	RawData []byte `tz:"dyn"`
}

func (RevealProofRawData) RevealProof() {}

type RevealProofMetadata struct{}

func (RevealProofMetadata) RevealProof() {}

type RevealProofDALPage struct {
	DALPageID
	DALProof []byte `tz:"dyn"`
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
	PublishedLevel int32
	SlotIndex      uint8
	PageIndex      int16
}

type SmartRollupTimeoutResultContents struct {
	ConsumedMilligas tz.BigUint
	GameStatus       GameStatus
	BalanceUpdates   []*BalanceUpdate `tz:"dyn"`
}

type GameStatus interface {
	GameStatusKind() string
}

type GameStatusOngoing struct{}

func (GameStatusOngoing) GameStatusKind() string { return "ongoing" }

type GameStatusEnded struct {
	Result GameStatusResult
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
	Reason LooseReason
	Player tz.PublicKeyHash
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
	Metadata ManagerMetadata[SmartRollupTimeoutResult]
}

func (*SmartRollupRefuteContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupRefuteContentsAndResult) OperationContents() core.OperationContents {
	return &op.SmartRollupRefute
}

type SmartRollupTimeout struct {
	ManagerOperation
	Rollup  *tz.SmartRollupAddress
	Stakers SmartRollupStakers
}

type SmartRollupStakers struct {
	Alice tz.PublicKeyHash
	Bob   tz.PublicKeyHash
}

func (*SmartRollupTimeout) OperationKind() string { return "smart_rollup_timeout" }

type SmartRollupTimeoutContentsAndResult struct {
	SmartRollupTimeout
	Metadata ManagerMetadata[SmartRollupTimeoutResult]
}

func (*SmartRollupTimeoutContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupTimeoutContentsAndResult) OperationContents() core.OperationContents {
	return &op.SmartRollupTimeout
}

type SmartRollupExecuteOutboxMessage struct {
	ManagerOperation
	Rollup             *tz.SmartRollupAddress
	CementedCommitment *tz.SmartRollupCommitmentHash
	OutputProof        []byte `tz:"dyn"`
}

func (*SmartRollupExecuteOutboxMessage) OperationKind() string {
	return "smart_rollup_execute_outbox_message"
}

type TicketReceipt = proto_015_PtLimaPt.TicketReceipt

type SmartRollupExecuteOutboxMessageResultContents struct {
	BalanceUpdates      []*BalanceUpdate `tz:"dyn"`
	TicketUpdates       []*TicketReceipt `tz:"dyn"`
	ConsumedMilligas    tz.BigUint
	PaidStorageSizeDiff tz.BigInt
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
	Metadata ManagerMetadata[SmartRollupExecuteOutboxMessageResult]
}

func (*SmartRollupExecuteOutboxMessageContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupExecuteOutboxMessageContentsAndResult) OperationContents() core.OperationContents {
	return &op.SmartRollupExecuteOutboxMessage
}

type SmartRollupRecoverBond struct {
	ManagerOperation
	Rollup *tz.SmartRollupAddress
	Staker tz.PublicKeyHash
}

func (*SmartRollupRecoverBond) OperationKind() string { return "smart_rollup_recover_bond" }

type SmartRollupRecoverBondResultContents struct {
	BalanceUpdates   []*BalanceUpdate `tz:"dyn"`
	ConsumedMilligas tz.BigUint
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
	Metadata ManagerMetadata[SmartRollupRecoverBondResult]
}

func (*SmartRollupRecoverBondContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupRecoverBondContentsAndResult) OperationContents() core.OperationContents {
	return &op.SmartRollupRecoverBond
}
