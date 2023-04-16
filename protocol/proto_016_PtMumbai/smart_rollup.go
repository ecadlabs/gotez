package proto_016_PtMumbai

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/expression"
	"github.com/ecadlabs/gotez/protocol/proto"
	kathma "github.com/ecadlabs/gotez/protocol/proto_014_PtKathma"
	"github.com/ecadlabs/gotez/protocol/proto_015_PtLimaPt"
)

type PVMKind uint8

const (
	PVMArith PVMKind = iota
	PVM_WASM_2_0_0
)

type SmartRollupOriginate struct {
	ManagerOperation
	PVMKind
	Kernel           []byte                `tz:"dyn"`
	OriginationProof []byte                `tz:"dyn"`
	ParametersTy     expression.Expression `tz:"dyn"`
}

func (*SmartRollupOriginate) OperationKind() string { return "smart_rollup_originate" }

type SmartRollupOriginateResult interface {
	SmartRollupOriginateResult()
	OperationResult
}

type SmartRollupOriginateResultContents struct {
	BalanceUpdates        []*BalanceUpdate `tz:"dyn"`
	Address               *tz.SmartRollupAddress
	GenesisCommitmentHash *tz.MumbaiSmartRollupHash
	ConsumedMilligas      tz.BigUint
	Size                  tz.BigInt
}

type SmartRollupOriginateResultApplied struct {
	kathma.OperationResultApplied[SmartRollupOriginateResultContents]
}

func (*SmartRollupOriginateResultApplied) SmartRollupOriginateResult() {}

type SmartRollupOriginateResultBacktracked struct {
	kathma.OperationResultBacktracked[SmartRollupOriginateResultContents]
}

func (*SmartRollupOriginateResultBacktracked) SmartRollupOriginateResult() {}

type SmartRollupOriginateResultFailed struct{ kathma.OperationResultFailed }

func (*SmartRollupOriginateResultFailed) SmartRollupOriginateResult() {}

type SmartRollupOriginateResultSkipped struct{ kathma.OperationResultSkipped }

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

type SmartRollupOriginateSuccessfulManagerOperationResult SmartRollupOriginateResultContents

func (*SmartRollupOriginateSuccessfulManagerOperationResult) SuccessfulManagerOperationResult() {}
func (*SmartRollupOriginateSuccessfulManagerOperationResult) OperationKind() string {
	return "smart_rollup_originate"
}

type SmartRollupAddMessages struct {
	ManagerOperation
	Message []proto.Bytes `tz:"dyn"`
}

func (*SmartRollupAddMessages) OperationKind() string { return "smart_rollup_add_messages" }

type SmartRollupAddMessagesContentsAndResult struct {
	SmartRollupAddMessages
	Metadata ManagerMetadata[EventResult]
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

type SmartRollupCementResult interface {
	SmartRollupCementResult()
	OperationResult
}

type SmartRollupCementResultApplied struct {
	kathma.OperationResultApplied[SmartRollupCementResultContents]
}

func (*SmartRollupCementResultApplied) SmartRollupCementResult() {}

type SmartRollupCementResultBacktracked struct {
	kathma.OperationResultBacktracked[SmartRollupCementResultContents]
}

func (*SmartRollupCementResultBacktracked) SmartRollupCementResult() {}

type SmartRollupCementResultFailed struct{ kathma.OperationResultFailed }

func (*SmartRollupCementResultFailed) SmartRollupCementResult() {}

type SmartRollupCementResultSkipped struct{ kathma.OperationResultSkipped }

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

type SmartRollupPublishResultContents struct {
	ConsumedMilligas tz.BigUint
	StakedHash       *tz.MumbaiSmartRollupHash
	PublishedAtLevel int32
	BalanceUpdates   []*BalanceUpdate `tz:"dyn"`
}

type SmartRollupPublishResult interface {
	SmartRollupPublishResult()
	OperationResult
}

type SmartRollupPublishResultApplied struct {
	kathma.OperationResultApplied[SmartRollupPublishResultContents]
}

func (*SmartRollupPublishResultApplied) SmartRollupPublishResult() {}

type SmartRollupPublishResultBacktracked struct {
	kathma.OperationResultBacktracked[SmartRollupPublishResultContents]
}

func (*SmartRollupPublishResultBacktracked) SmartRollupPublishResult() {}

type SmartRollupPublishResultFailed struct{ kathma.OperationResultFailed }

func (*SmartRollupPublishResultFailed) SmartRollupPublishResult() {}

type SmartRollupPublishResultSkipped struct{ kathma.OperationResultSkipped }

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
	OperationResult
}

type SmartRollupTimeoutResultApplied struct {
	kathma.OperationResultApplied[SmartRollupTimeoutResultContents]
}

func (*SmartRollupTimeoutResultApplied) SmartRollupTimeoutResult() {}

type SmartRollupTimeoutResultBacktracked struct {
	kathma.OperationResultBacktracked[SmartRollupTimeoutResultContents]
}

func (*SmartRollupTimeoutResultBacktracked) SmartRollupTimeoutResult() {}

type SmartRollupTimeoutResultFailed struct{ kathma.OperationResultFailed }

func (*SmartRollupTimeoutResultFailed) SmartRollupTimeoutResult() {}

type SmartRollupTimeoutResultSkipped struct{ kathma.OperationResultSkipped }

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

type SmartRollupExecuteOutboxMessageResultContents struct {
	BalanceUpdates      []*BalanceUpdate `tz:"dyn"`
	TicketUpdates       []*TicketReceipt `tz:"dyn"`
	ConsumedMilligas    tz.BigUint
	PaidStorageSizeDiff tz.BigInt
}

type SmartRollupExecuteOutboxMessageResult interface {
	SmartRollupExecuteOutboxMessageResult()
	OperationResult
}

type SmartRollupExecuteOutboxMessageResultApplied struct {
	kathma.OperationResultApplied[SmartRollupExecuteOutboxMessageResultContents]
}

func (*SmartRollupExecuteOutboxMessageResultApplied) SmartRollupExecuteOutboxMessageResult() {}

type SmartRollupExecuteOutboxMessageResultBacktracked struct {
	kathma.OperationResultBacktracked[SmartRollupExecuteOutboxMessageResultContents]
}

func (*SmartRollupExecuteOutboxMessageResultBacktracked) SmartRollupExecuteOutboxMessageResult() {}

type SmartRollupExecuteOutboxMessageResultFailed struct{ kathma.OperationResultFailed }

func (*SmartRollupExecuteOutboxMessageResultFailed) SmartRollupExecuteOutboxMessageResult() {}

type SmartRollupExecuteOutboxMessageResultSkipped struct{ kathma.OperationResultSkipped }

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
	OperationResult
}

type SmartRollupRecoverBondResultApplied struct {
	kathma.OperationResultApplied[SmartRollupRecoverBondResultContents]
}

func (*SmartRollupRecoverBondResultApplied) SmartRollupRecoverBondResult() {}

type SmartRollupRecoverBondResultBacktracked struct {
	kathma.OperationResultBacktracked[SmartRollupRecoverBondResultContents]
}

func (*SmartRollupRecoverBondResultBacktracked) SmartRollupRecoverBondResult() {}

type SmartRollupRecoverBondResultFailed struct{ kathma.OperationResultFailed }

func (*SmartRollupRecoverBondResultFailed) SmartRollupRecoverBondResult() {}

type SmartRollupRecoverBondResultSkipped struct{ kathma.OperationResultSkipped }

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
