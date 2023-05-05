package proto_016_PtMumbai

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/proto_012_Psithaca"
	"github.com/ecadlabs/gotez/protocol/proto_013_PtJakart"
	"github.com/ecadlabs/gotez/protocol/proto_014_PtKathma"
	"github.com/ecadlabs/gotez/protocol/proto_015_PtLimaPt"
)

type OperationContents interface {
	core.OperationContents
}

type ManagerOperation = proto_012_Psithaca.ManagerOperation
type SeedNonceRevelation = proto_012_Psithaca.SeedNonceRevelation
type Preendorsement = proto_012_Psithaca.Preendorsement
type InlinedPreendorsement = proto_012_Psithaca.InlinedPreendorsement
type Endorsement = proto_012_Psithaca.Endorsement
type InlinedEndorsement = proto_012_Psithaca.InlinedEndorsement
type DoublePreendorsementEvidence = proto_012_Psithaca.DoublePreendorsementEvidence
type DoubleEndorsementEvidence = proto_012_Psithaca.DoubleEndorsementEvidence
type Reveal = proto_012_Psithaca.Reveal
type RevealResultContents = proto_015_PtLimaPt.RevealResultContents
type Delegation = proto_012_Psithaca.Delegation
type DelegationInternalOperationResult = proto_015_PtLimaPt.DelegationInternalOperationResult
type DelegationResultContents = proto_015_PtLimaPt.DelegationResultContents
type RegisterGlobalConstant = proto_012_Psithaca.RegisterGlobalConstant
type SetDepositsLimit = proto_012_Psithaca.SetDepositsLimit
type SetDepositsLimitResultContents = proto_015_PtLimaPt.SetDepositsLimitResultContents
type UpdateConsensusKey = proto_015_PtLimaPt.UpdateConsensusKey
type UpdateConsensusKeyResultContents = proto_015_PtLimaPt.UpdateConsensusKeyResultContents
type IncreasePaidStorage = proto_014_PtKathma.IncreasePaidStorage
type ActivateAccount = proto_012_Psithaca.ActivateAccount
type Proposals = proto_012_Psithaca.Proposals
type Ballot = proto_012_Psithaca.Ballot
type VDFRevelation = proto_014_PtKathma.VDFRevelation
type DrainDelegate = proto_015_PtLimaPt.DrainDelegate
type FailingNoop = proto_012_Psithaca.FailingNoop
type EventResult = proto_015_PtLimaPt.EventResult
type EventResultContents = proto_015_PtLimaPt.EventResultContents
type EventInternalOperationResult = proto_015_PtLimaPt.EventInternalOperationResult
type LazyStorageDiff = proto_015_PtLimaPt.LazyStorageDiff
type TransferTicket = proto_013_PtJakart.TransferTicket

type OperationContentsAndResult interface {
	core.OperationContentsAndResult
}

type SeedNonceRevelationContentsAndResult[T core.BalanceUpdateKind] struct {
	SeedNonceRevelation
	Metadata []*BalanceUpdate[T] `tz:"dyn"`
}

func (*SeedNonceRevelationContentsAndResult[T]) OperationContentsAndResult() {}

type DoubleEndorsementEvidenceContentsAndResult[T core.BalanceUpdateKind] struct {
	DoubleEndorsementEvidence
	Metadata []*BalanceUpdate[T] `tz:"dyn"`
}

func (*DoubleEndorsementEvidenceContentsAndResult[T]) OperationContentsAndResult() {}

type EndorsementMetadata[T core.BalanceUpdateKind] struct {
	BalanceUpdates   []*BalanceUpdate[T] `tz:"dyn"`
	Delegate         tz.PublicKeyHash
	EndorsementPower int32
	ConsensusKey     tz.PublicKeyHash
}

type EndorsementContentsAndResult[T core.BalanceUpdateKind] struct {
	Endorsement
	Metadata EndorsementMetadata[T]
}

func (*EndorsementContentsAndResult[T]) OperationContentsAndResult() {}

type DALAttestation struct {
	Attestor    tz.PublicKeyHash
	Attestation tz.BigInt
	Level       int32
}

func (*DALAttestation) OperationKind() string { return "dal_attestation" }

type DALAttestationContentsAndResult struct {
	DALAttestation
	Metadata tz.PublicKeyHash
}

func (*DALAttestationContentsAndResult) OperationContentsAndResult() {}

type RevealContentsAndResult[T core.BalanceUpdateKind] struct {
	Reveal
	Metadata ManagerMetadata[EventResult, T]
}

func (*RevealContentsAndResult[T]) OperationContentsAndResult() {}

type DelegationContentsAndResult[T core.BalanceUpdateKind] struct {
	Delegation
	Metadata ManagerMetadata[EventResult, T]
}

func (*DelegationContentsAndResult[T]) OperationContentsAndResult() {}

type RegisterGlobalConstantResult interface {
	RegisterGlobalConstantResult()
	core.OperationResult
}

type RegisterGlobalConstantResultContents[T core.BalanceUpdateKind] struct {
	BalanceUpdates   []*BalanceUpdate[T] `tz:"dyn"`
	ConsumedMilligas tz.BigUint
	StorageSize      tz.BigInt
	GlobalAddress    *tz.ScriptExprHash
}

func (RegisterGlobalConstantResultContents[T]) SuccessfulManagerOperationResult() {}
func (RegisterGlobalConstantResultContents[T]) OperationKind() string {
	return "register_global_constant"
}

type RegisterGlobalConstantResultApplied[T core.BalanceUpdateKind] struct {
	core.OperationResultApplied[RegisterGlobalConstantResultContents[T]]
}

func (*RegisterGlobalConstantResultApplied[T]) RegisterGlobalConstantResult() {}

type RegisterGlobalConstantResultBacktracked[T core.BalanceUpdateKind] struct {
	core.OperationResultBacktracked[RegisterGlobalConstantResultContents[T]]
}

func (*RegisterGlobalConstantResultBacktracked[T]) RegisterGlobalConstantResult() {}

type RegisterGlobalConstantResultFailed struct{ core.OperationResultFailed }

func (*RegisterGlobalConstantResultFailed) RegisterGlobalConstantResult() {}

type RegisterGlobalConstantResultSkipped struct{ core.OperationResultSkipped }

func (*RegisterGlobalConstantResultSkipped) RegisterGlobalConstantResult() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[RegisterGlobalConstantResult]{
		Variants: encoding.Variants[RegisterGlobalConstantResult]{
			0: (*RegisterGlobalConstantResultApplied[BalanceUpdateKind])(nil),
			1: (*RegisterGlobalConstantResultFailed)(nil),
			2: (*RegisterGlobalConstantResultSkipped)(nil),
			3: (*RegisterGlobalConstantResultBacktracked[BalanceUpdateKind])(nil),
		},
	})
}

type RegisterGlobalConstantContentsAndResult[T core.BalanceUpdateKind] struct {
	RegisterGlobalConstant
	Metadata ManagerMetadata[RegisterGlobalConstantResult, T]
}

func (*RegisterGlobalConstantContentsAndResult[T]) OperationContentsAndResult() {}

type SetDepositsLimitContentsAndResult[T core.BalanceUpdateKind] struct {
	SetDepositsLimit
	Metadata ManagerMetadata[EventResult, T]
}

func (*SetDepositsLimitContentsAndResult[T]) OperationContentsAndResult() {}

type UpdateConsensusKeyContentsAndResult[T core.BalanceUpdateKind] struct {
	UpdateConsensusKey
	Metadata ManagerMetadata[EventResult, T]
}

func (*UpdateConsensusKeyContentsAndResult[T]) OperationContentsAndResult() {}

type TransferTicketContentsAndResult[T core.BalanceUpdateKind] struct {
	TransferTicket
	Metadata ManagerMetadata[SmartRollupExecuteOutboxMessageResult, T]
}

func (*TransferTicketContentsAndResult[T]) OperationContentsAndResult() {}

type IncreasePaidStorageResult interface {
	IncreasePaidStorageResult()
	core.OperationResult
}

type IncreasePaidStorageResultContents[T core.BalanceUpdateKind] struct {
	BalanceUpdates   []*BalanceUpdate[T] `tz:"dyn"`
	ConsumedMilligas tz.BigUint
}

func (IncreasePaidStorageResultContents[T]) SuccessfulManagerOperationResult() {}
func (IncreasePaidStorageResultContents[T]) OperationKind() string {
	return "increase_paid_storage"
}

type IncreasePaidStorageResultApplied[T core.BalanceUpdateKind] struct {
	core.OperationResultApplied[IncreasePaidStorageResultContents[T]]
}

func (*IncreasePaidStorageResultApplied[T]) IncreasePaidStorageResult() {}

type IncreasePaidStorageResultBacktracked[T core.BalanceUpdateKind] struct {
	core.OperationResultBacktracked[IncreasePaidStorageResultContents[T]]
}

func (*IncreasePaidStorageResultBacktracked[T]) IncreasePaidStorageResult() {}

type IncreasePaidStorageResultFailed struct{ core.OperationResultFailed }

func (*IncreasePaidStorageResultFailed) IncreasePaidStorageResult() {}

type IncreasePaidStorageResultSkipped struct{ core.OperationResultSkipped }

func (*IncreasePaidStorageResultSkipped) IncreasePaidStorageResult() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[IncreasePaidStorageResult]{
		Variants: encoding.Variants[IncreasePaidStorageResult]{
			0: (*IncreasePaidStorageResultApplied[BalanceUpdateKind])(nil),
			1: (*IncreasePaidStorageResultFailed)(nil),
			2: (*IncreasePaidStorageResultSkipped)(nil),
			3: (*IncreasePaidStorageResultBacktracked[BalanceUpdateKind])(nil),
		},
	})
}

type IncreasePaidStorageContentsAndResult[T core.BalanceUpdateKind] struct {
	IncreasePaidStorage
	Metadata ManagerMetadata[IncreasePaidStorageResult, T]
}

func (*IncreasePaidStorageContentsAndResult[T]) OperationContentsAndResult() {}

type DoubleBakingEvidence struct {
	Block1 BlockHeader `tz:"dyn"`
	Block2 BlockHeader `tz:"dyn"`
}

func (*DoubleBakingEvidence) OperationKind() string { return "double_baking_evidence" }

type DoubleBakingEvidenceContentsAndResult[T core.BalanceUpdateKind] struct {
	DoubleBakingEvidence
	Metadata []*BalanceUpdate[T] `tz:"dyn"`
}

func (*DoubleBakingEvidenceContentsAndResult[T]) OperationContentsAndResult() {}

type ActivateAccountContentsAndResult[T core.BalanceUpdateKind] struct {
	ActivateAccount
	Metadata []*BalanceUpdate[T] `tz:"dyn"`
}

func (*ActivateAccountContentsAndResult[T]) OperationContentsAndResult() {}

type DoublePreendorsementEvidenceContentsAndResult[T core.BalanceUpdateKind] struct {
	DoublePreendorsementEvidence
	Metadata []*BalanceUpdate[T] `tz:"dyn"`
}

func (*DoublePreendorsementEvidenceContentsAndResult[T]) OperationContentsAndResult() {}

type PreendorsementContentsAndResult[T core.BalanceUpdateKind] struct {
	Preendorsement
	Metadata EndorsementMetadata[T]
}

func (*PreendorsementContentsAndResult[T]) OperationContentsAndResult() {}

type VDFRevelationContentsAndResult[T core.BalanceUpdateKind] struct {
	VDFRevelation
	Metadata []*BalanceUpdate[T] `tz:"dyn"`
}

func (*VDFRevelationContentsAndResult[T]) OperationContentsAndResult() {}

type DrainDelegateMetadata[T core.BalanceUpdateKind] struct {
	BalanceUpdates               []*BalanceUpdate[T] `tz:"dyn"`
	AllocatedDestinationContract bool
}

type DrainDelegateContentsAndResult[T core.BalanceUpdateKind] struct {
	DrainDelegate
	Metadata DrainDelegateMetadata[T]
}

func (*DrainDelegateContentsAndResult[T]) OperationContentsAndResult() {}

type DALPublishSlotHeader struct {
	ManagerOperation
	SlotHeader SlotHeader
}

type SlotHeader struct {
	Level           int32
	Index           uint8
	Сommitment      *tz.DALCommitment
	CommitmentProof [48]byte
}

func (*DALPublishSlotHeader) OperationKind() string { return "dal_publish_slot_header" }

type DALPublishSlotHeaderContentsAndResult[T core.BalanceUpdateKind] struct {
	DALPublishSlotHeader
	Metadata ManagerMetadata[EventResult, T]
}

func (*DALPublishSlotHeaderContentsAndResult[T]) OperationContentsAndResult() {}

type SignaturePrefix struct {
	SignaturePrefix SignaturePrefixPayload
}

func (*SignaturePrefix) OperationKind() string       { return "signature_prefix" }
func (*SignaturePrefix) OperationContentsAndResult() {}

type SignaturePrefixPayload interface {
	SignaturePrefixPayload()
}

type BLSSignaturePrefix [32]byte

func (*BLSSignaturePrefix) SignaturePrefixPayload() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SignaturePrefixPayload]{
		Variants: encoding.Variants[SignaturePrefixPayload]{
			3: (*BLSSignaturePrefix)(nil),
		},
	})

	encoding.RegisterEnum(&encoding.Enum[OperationContents]{
		Variants: encoding.Variants[OperationContents]{
			1:   (*SeedNonceRevelation)(nil),
			2:   (*DoubleEndorsementEvidence)(nil),
			3:   (*DoubleBakingEvidence)(nil),
			4:   (*ActivateAccount)(nil),
			5:   (*Proposals)(nil),
			6:   (*Ballot)(nil),
			7:   (*DoublePreendorsementEvidence)(nil),
			8:   (*VDFRevelation)(nil),
			9:   (*DrainDelegate)(nil),
			17:  (*FailingNoop)(nil),
			20:  (*Preendorsement)(nil),
			21:  (*Endorsement)(nil),
			22:  (*DALAttestation)(nil),
			107: (*Reveal)(nil),
			108: (*Transaction)(nil),
			109: (*Origination)(nil),
			110: (*Delegation)(nil),
			111: (*RegisterGlobalConstant)(nil),
			112: (*SetDepositsLimit)(nil),
			113: (*IncreasePaidStorage)(nil),
			114: (*UpdateConsensusKey)(nil),
			// 150 Tx_rollup_origination
			// 151 Tx_rollup_submit_batch
			// 152 Tx_rollup_commit
			// 153 Tx_rollup_return_bond
			// 154 Tx_rollup_finalize_commitment
			// 155 Tx_rollup_remove_commitment
			// 156 Tx_rollup_rejection
			// 157 Tx_rollup_dispatch_tickets
			158: (*TransferTicket)(nil),
			200: (*SmartRollupOriginate)(nil),
			201: (*SmartRollupAddMessages)(nil),
			202: (*SmartRollupCement)(nil),
			203: (*SmartRollupPublish)(nil),
			204: (*SmartRollupRefute)(nil),
			205: (*SmartRollupTimeout)(nil),
			206: (*SmartRollupExecuteOutboxMessage)(nil),
			207: (*SmartRollupRecoverBond)(nil),
			230: (*DALPublishSlotHeader)(nil),
			250: (*ZkRollupOrigination)(nil),
			251: (*ZkRollupPublish)(nil),
			252: (*ZkRollupUpdate)(nil),
			255: (*SignaturePrefix)(nil),
		},
	})

	encoding.RegisterEnum(&encoding.Enum[OperationContentsAndResult]{
		Variants: encoding.Variants[OperationContentsAndResult]{
			1:   (*SeedNonceRevelationContentsAndResult[BalanceUpdateKind])(nil),
			2:   (*DoubleEndorsementEvidenceContentsAndResult[BalanceUpdateKind])(nil),
			3:   (*DoubleBakingEvidenceContentsAndResult[BalanceUpdateKind])(nil),
			4:   (*ActivateAccountContentsAndResult[BalanceUpdateKind])(nil),
			5:   (*Proposals)(nil),
			6:   (*Ballot)(nil),
			7:   (*DoublePreendorsementEvidenceContentsAndResult[BalanceUpdateKind])(nil),
			8:   (*VDFRevelationContentsAndResult[BalanceUpdateKind])(nil),
			9:   (*DrainDelegateContentsAndResult[BalanceUpdateKind])(nil),
			20:  (*PreendorsementContentsAndResult[BalanceUpdateKind])(nil),
			21:  (*EndorsementContentsAndResult[BalanceUpdateKind])(nil),
			22:  (*DALAttestationContentsAndResult)(nil),
			107: (*RevealContentsAndResult[BalanceUpdateKind])(nil),
			108: (*TransactionContentsAndResult[BalanceUpdateKind])(nil),
			109: (*OriginationContentsAndResult[BalanceUpdateKind])(nil),
			110: (*DelegationContentsAndResult[BalanceUpdateKind])(nil),
			111: (*RegisterGlobalConstantContentsAndResult[BalanceUpdateKind])(nil),
			112: (*SetDepositsLimitContentsAndResult[BalanceUpdateKind])(nil),
			113: (*IncreasePaidStorageContentsAndResult[BalanceUpdateKind])(nil),
			114: (*UpdateConsensusKeyContentsAndResult[BalanceUpdateKind])(nil),
			158: (*TransferTicketContentsAndResult[BalanceUpdateKind])(nil),
			200: (*SmartRollupOriginateContentsAndResult[BalanceUpdateKind])(nil),
			201: (*SmartRollupAddMessagesContentsAndResult[BalanceUpdateKind])(nil),
			202: (*SmartRollupCementContentsAndResult[BalanceUpdateKind])(nil),
			203: (*SmartRollupPublishContentsAndResult[BalanceUpdateKind])(nil),
			204: (*SmartRollupRefuteContentsAndResult[BalanceUpdateKind])(nil),
			205: (*SmartRollupTimeoutContentsAndResult[BalanceUpdateKind])(nil),
			206: (*SmartRollupExecuteOutboxMessageContentsAndResult[BalanceUpdateKind])(nil),
			207: (*SmartRollupRecoverBondContentsAndResult[BalanceUpdateKind])(nil),
			230: (*DALPublishSlotHeaderContentsAndResult[BalanceUpdateKind])(nil),
			250: (*ZkRollupOriginationContentsAndResult[BalanceUpdateKind])(nil),
			251: (*ZkRollupPublishContentsAndResult[BalanceUpdateKind])(nil),
			252: (*ZkRollupUpdateContentsAndResult[BalanceUpdateKind])(nil),
			255: (*SignaturePrefix)(nil),
		},
	})
}

type SuccessfulManagerOperationResult interface {
	core.SuccessfulManagerOperationResult
}

type ManagerMetadata[Res core.OperationResult, Bal core.BalanceUpdateKind] struct {
	BalanceUpdates           []*BalanceUpdate[Bal] `tz:"dyn"`
	OperationResult          Res
	InternalOperationResults []InternalOperationResult `tz:"dyn"`
}

type InternalOperationResult interface {
	core.InternalOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[InternalOperationResult]{
		Variants: encoding.Variants[InternalOperationResult]{
			1: (*TransactionInternalOperationResult)(nil),
			2: (*OriginationInternalOperationResult)(nil),
			3: (*DelegationInternalOperationResult)(nil),
			4: (*EventInternalOperationResult)(nil),
		},
	})

	encoding.RegisterEnum(&encoding.Enum[SuccessfulManagerOperationResult]{
		Variants: encoding.Variants[SuccessfulManagerOperationResult]{
			0:   (*RevealResultContents)(nil),
			1:   (*TransactionResultContents)(nil),
			2:   (*OriginationResultContents[BalanceUpdateKind])(nil),
			3:   (*DelegationResultContents)(nil),
			5:   (*SetDepositsLimitResultContents)(nil),
			6:   (*UpdateConsensusKeyResultContents)(nil),
			9:   (*IncreasePaidStorageResultContents[BalanceUpdateKind])(nil),
			200: (*SmartRollupOriginateResultContents[BalanceUpdateKind])(nil),
		},
	})
}
