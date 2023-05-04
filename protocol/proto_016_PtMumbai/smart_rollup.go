package proto_016_PtMumbai

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/proto_015_PtLimaPt"
)

type PVMKind uint8

const (
	PVMArith PVMKind = iota
	PVM_WASM_2_0_0
)

type SmartRollupOriginate = proto_015_PtLimaPt.SmartRollupOriginate

type SmartRollupOriginateResult interface {
	SmartRollupOriginateResult()
	core.OperationResult
}

type SmartRollupOriginateResultContents[T core.BalanceUpdate] struct {
	BalanceUpdates        []T `tz:"dyn"`
	Address               *tz.SmartRollupAddress
	GenesisCommitmentHash *tz.MumbaiSmartRollupHash
	ConsumedMilligas      tz.BigUint
	Size                  tz.BigInt
}

func (SmartRollupOriginateResultContents[T]) SuccessfulManagerOperationResult() {}
func (SmartRollupOriginateResultContents[T]) OperationKind() string {
	return "smart_rollup_originate"
}

type SmartRollupOriginateResultApplied struct {
	core.OperationResultApplied[SmartRollupOriginateResultContents[*BalanceUpdate]]
}

func (*SmartRollupOriginateResultApplied) SmartRollupOriginateResult() {}

type SmartRollupOriginateResultBacktracked struct {
	core.OperationResultBacktracked[SmartRollupOriginateResultContents[*BalanceUpdate]]
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
	Metadata ManagerMetadata[SmartRollupOriginateResult, *BalanceUpdate]
}

func (*SmartRollupOriginateContentsAndResult) OperationContentsAndResult() {}

type SmartRollupAddMessages struct {
	ManagerOperation
	Message []core.Bytes `tz:"dyn"`
}

func (*SmartRollupAddMessages) OperationKind() string { return "smart_rollup_add_messages" }

type SmartRollupAddMessagesContentsAndResult struct {
	SmartRollupAddMessages
	Metadata ManagerMetadata[EventResult, *BalanceUpdate]
}

func (*SmartRollupAddMessagesContentsAndResult) OperationContentsAndResult() {}

type SmartRollupCement struct {
	ManagerOperation
	Rollup     *tz.SmartRollupAddress
	Commitment *tz.MumbaiSmartRollupHash
}

func (*SmartRollupCement) OperationKind() string { return "smart_rollup_cement" }

type SmartRollupCementResultContents struct {
	ConsumedMilligas tz.BigUint
	InboxLevel       int32
}

func (SmartRollupCementResultContents) SuccessfulManagerOperationResult() {}
func (SmartRollupCementResultContents) OperationKind() string             { return "smart_rollup_cement" }

type SmartRollupCementResult interface {
	SmartRollupCementResult()
	core.OperationResult
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
	Metadata ManagerMetadata[SmartRollupCementResult, *BalanceUpdate]
}

func (*SmartRollupCementContentsAndResult) OperationContentsAndResult() {}

type SmartRollupPublish struct {
	ManagerOperation
	Rollup     *tz.SmartRollupAddress
	Commitment SmartRollupCommitment
}

func (*SmartRollupRefute) OperationKind() string { return "smart_rollup_refute" }

type SmartRollupCommitment struct {
	CompressedState *tz.MumbaiSmartRollupStateHash
	InboxLevel      int32
	Predecessor     *tz.MumbaiSmartRollupHash
	NumberOfTicks   int64
}

type SmartRollupPublishResultContents[T core.BalanceUpdate] struct {
	ConsumedMilligas tz.BigUint
	StakedHash       *tz.MumbaiSmartRollupHash
	PublishedAtLevel int32
	BalanceUpdates   []T `tz:"dyn"`
}

func (SmartRollupPublishResultContents[T]) SuccessfulManagerOperationResult() {}
func (SmartRollupPublishResultContents[T]) OperationKind() string             { return "smart_rollup_publish" }

type SmartRollupPublishResult interface {
	SmartRollupPublishResult()
	core.OperationResult
}

type SmartRollupPublishResultApplied struct {
	core.OperationResultApplied[SmartRollupPublishResultContents[*BalanceUpdate]]
}

func (*SmartRollupPublishResultApplied) SmartRollupPublishResult() {}

type SmartRollupPublishResultBacktracked struct {
	core.OperationResultBacktracked[SmartRollupPublishResultContents[*BalanceUpdate]]
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
	Metadata ManagerMetadata[SmartRollupPublishResult, *BalanceUpdate]
}

func (*SmartRollupPublishContentsAndResult) OperationContentsAndResult() {}

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
	PlayerCommitmentHash   *tz.MumbaiSmartRollupHash
	OpponentCommitmentHash *tz.MumbaiSmartRollupHash
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
	State tz.Option[*tz.MumbaiSmartRollupStateHash]
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

type SmartRollupTimeoutResultContents[T core.BalanceUpdate] struct {
	ConsumedMilligas tz.BigUint
	GameStatus       GameStatus
	BalanceUpdates   []T `tz:"dyn"`
}

func (SmartRollupTimeoutResultContents[T]) SuccessfulManagerOperationResult() {}
func (SmartRollupTimeoutResultContents[T]) OperationKind() string             { return "smart_rollup_timeout" }

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
	core.OperationResult
}

type SmartRollupTimeoutResultApplied struct {
	core.OperationResultApplied[SmartRollupTimeoutResultContents[*BalanceUpdate]]
}

func (*SmartRollupTimeoutResultApplied) SmartRollupTimeoutResult() {}

type SmartRollupTimeoutResultBacktracked struct {
	core.OperationResultBacktracked[SmartRollupTimeoutResultContents[*BalanceUpdate]]
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
	Metadata ManagerMetadata[SmartRollupTimeoutResult, *BalanceUpdate]
}

func (*SmartRollupRefuteContentsAndResult) OperationContentsAndResult() {}

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
	Metadata ManagerMetadata[SmartRollupTimeoutResult, *BalanceUpdate]
}

func (*SmartRollupTimeoutContentsAndResult) OperationContentsAndResult() {}

type SmartRollupExecuteOutboxMessage struct {
	ManagerOperation
	Rollup             *tz.SmartRollupAddress
	CementedCommitment *tz.MumbaiSmartRollupHash
	OutputProof        []byte `tz:"dyn"`
}

func (*SmartRollupExecuteOutboxMessage) OperationKind() string {
	return "smart_rollup_execute_outbox_message"
}

type TicketReceipt = proto_015_PtLimaPt.TicketReceipt

type SmartRollupExecuteOutboxMessageResultContents[T core.BalanceUpdate] struct {
	BalanceUpdates      []T              `tz:"dyn"`
	TicketUpdates       []*TicketReceipt `tz:"dyn"`
	ConsumedMilligas    tz.BigUint
	PaidStorageSizeDiff tz.BigInt
}

func (SmartRollupExecuteOutboxMessageResultContents[T]) SuccessfulManagerOperationResult() {}
func (SmartRollupExecuteOutboxMessageResultContents[T]) OperationKind() string {
	return "smart_rollup_execute_outbox_message"
}

type SmartRollupExecuteOutboxMessageResult interface {
	SmartRollupExecuteOutboxMessageResult()
	core.OperationResult
}

type SmartRollupExecuteOutboxMessageResultApplied struct {
	core.OperationResultApplied[SmartRollupExecuteOutboxMessageResultContents[*BalanceUpdate]]
}

func (*SmartRollupExecuteOutboxMessageResultApplied) SmartRollupExecuteOutboxMessageResult() {}

type SmartRollupExecuteOutboxMessageResultBacktracked struct {
	core.OperationResultBacktracked[SmartRollupExecuteOutboxMessageResultContents[*BalanceUpdate]]
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
	Metadata ManagerMetadata[SmartRollupExecuteOutboxMessageResult, *BalanceUpdate]
}

func (*SmartRollupExecuteOutboxMessageContentsAndResult) OperationContentsAndResult() {}

type SmartRollupRecoverBond struct {
	ManagerOperation
	Rollup *tz.SmartRollupAddress
	Staker tz.PublicKeyHash
}

func (*SmartRollupRecoverBond) OperationKind() string { return "smart_rollup_recover_bond" }

type SmartRollupRecoverBondResultContents[T core.BalanceUpdate] struct {
	BalanceUpdates   []T `tz:"dyn"`
	ConsumedMilligas tz.BigUint
}

func (SmartRollupRecoverBondResultContents[T]) SuccessfulManagerOperationResult() {}
func (SmartRollupRecoverBondResultContents[T]) OperationKind() string {
	return "smart_rollup_recover_bond"
}

type SmartRollupRecoverBondResult interface {
	SmartRollupRecoverBondResult()
	core.OperationResult
}

type SmartRollupRecoverBondResultApplied struct {
	core.OperationResultApplied[SmartRollupRecoverBondResultContents[*BalanceUpdate]]
}

func (*SmartRollupRecoverBondResultApplied) SmartRollupRecoverBondResult() {}

type SmartRollupRecoverBondResultBacktracked struct {
	core.OperationResultBacktracked[SmartRollupRecoverBondResultContents[*BalanceUpdate]]
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
	Metadata ManagerMetadata[SmartRollupRecoverBondResult, *BalanceUpdate]
}

func (*SmartRollupRecoverBondContentsAndResult) OperationContentsAndResult() {}
