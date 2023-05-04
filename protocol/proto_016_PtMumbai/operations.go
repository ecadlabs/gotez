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

type SeedNonceRevelationContentsAndResult[T core.BalanceUpdate] struct {
	SeedNonceRevelation
	Metadata []T `tz:"dyn"`
}

func (*SeedNonceRevelationContentsAndResult[T]) OperationContentsAndResult() {}

type DoubleEndorsementEvidenceContentsAndResult[T core.BalanceUpdate] struct {
	DoubleEndorsementEvidence
	Metadata []T `tz:"dyn"`
}

func (*DoubleEndorsementEvidenceContentsAndResult[T]) OperationContentsAndResult() {}

type EndorsementMetadata[T core.BalanceUpdate] struct {
	BalanceUpdates   []T `tz:"dyn"`
	Delegate         tz.PublicKeyHash
	EndorsementPower int32
	ConsensusKey     tz.PublicKeyHash
}

type EndorsementContentsAndResult[T core.BalanceUpdate] struct {
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

type RevealContentsAndResult struct {
	Reveal
	Metadata ManagerMetadata[EventResult, *BalanceUpdate]
}

func (*RevealContentsAndResult) OperationContentsAndResult() {}

type DelegationContentsAndResult struct {
	Delegation
	Metadata ManagerMetadata[EventResult, *BalanceUpdate]
}

func (*DelegationContentsAndResult) OperationContentsAndResult() {}

type RegisterGlobalConstantResult interface {
	RegisterGlobalConstantResult()
	core.OperationResult
}

type RegisterGlobalConstantResultContents[T core.BalanceUpdate] struct {
	BalanceUpdates   []T `tz:"dyn"`
	ConsumedMilligas tz.BigUint
	StorageSize      tz.BigInt
	GlobalAddress    *tz.ScriptExprHash
}

func (RegisterGlobalConstantResultContents[T]) SuccessfulManagerOperationResult() {}
func (RegisterGlobalConstantResultContents[T]) OperationKind() string {
	return "register_global_constant"
}

type RegisterGlobalConstantResultApplied struct {
	core.OperationResultApplied[RegisterGlobalConstantResultContents[*BalanceUpdate]]
}

func (*RegisterGlobalConstantResultApplied) RegisterGlobalConstantResult() {}

type RegisterGlobalConstantResultBacktracked struct {
	core.OperationResultBacktracked[RegisterGlobalConstantResultContents[*BalanceUpdate]]
}

func (*RegisterGlobalConstantResultBacktracked) RegisterGlobalConstantResult() {}

type RegisterGlobalConstantResultFailed struct{ core.OperationResultFailed }

func (*RegisterGlobalConstantResultFailed) RegisterGlobalConstantResult() {}

type RegisterGlobalConstantResultSkipped struct{ core.OperationResultSkipped }

func (*RegisterGlobalConstantResultSkipped) RegisterGlobalConstantResult() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[RegisterGlobalConstantResult]{
		Variants: encoding.Variants[RegisterGlobalConstantResult]{
			0: (*RegisterGlobalConstantResultApplied)(nil),
			1: (*RegisterGlobalConstantResultFailed)(nil),
			2: (*RegisterGlobalConstantResultSkipped)(nil),
			3: (*RegisterGlobalConstantResultBacktracked)(nil),
		},
	})
}

type RegisterGlobalConstantContentsAndResult struct {
	RegisterGlobalConstant
	Metadata ManagerMetadata[RegisterGlobalConstantResult, *BalanceUpdate]
}

func (*RegisterGlobalConstantContentsAndResult) OperationContentsAndResult() {}

type SetDepositsLimitContentsAndResult struct {
	SetDepositsLimit
	Metadata ManagerMetadata[EventResult, *BalanceUpdate]
}

func (*SetDepositsLimitContentsAndResult) OperationContentsAndResult() {}

type UpdateConsensusKeyContentsAndResult struct {
	UpdateConsensusKey
	Metadata ManagerMetadata[EventResult, *BalanceUpdate]
}

func (*UpdateConsensusKeyContentsAndResult) OperationContentsAndResult() {}

type TransferTicketContentsAndResult struct {
	TransferTicket
	Metadata ManagerMetadata[SmartRollupExecuteOutboxMessageResult, *BalanceUpdate]
}

func (*TransferTicketContentsAndResult) OperationContentsAndResult() {}

type IncreasePaidStorageResult interface {
	IncreasePaidStorageResult()
	core.OperationResult
}

type IncreasePaidStorageResultContents[T core.BalanceUpdate] struct {
	BalanceUpdates   []T `tz:"dyn"`
	ConsumedMilligas tz.BigUint
}

func (IncreasePaidStorageResultContents[T]) SuccessfulManagerOperationResult() {}
func (IncreasePaidStorageResultContents[T]) OperationKind() string {
	return "increase_paid_storage"
}

type IncreasePaidStorageResultApplied struct {
	core.OperationResultApplied[IncreasePaidStorageResultContents[*BalanceUpdate]]
}

func (*IncreasePaidStorageResultApplied) IncreasePaidStorageResult() {}

type IncreasePaidStorageResultBacktracked struct {
	core.OperationResultBacktracked[IncreasePaidStorageResultContents[*BalanceUpdate]]
}

func (*IncreasePaidStorageResultBacktracked) IncreasePaidStorageResult() {}

type IncreasePaidStorageResultFailed struct{ core.OperationResultFailed }

func (*IncreasePaidStorageResultFailed) IncreasePaidStorageResult() {}

type IncreasePaidStorageResultSkipped struct{ core.OperationResultSkipped }

func (*IncreasePaidStorageResultSkipped) IncreasePaidStorageResult() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[IncreasePaidStorageResult]{
		Variants: encoding.Variants[IncreasePaidStorageResult]{
			0: (*IncreasePaidStorageResultApplied)(nil),
			1: (*IncreasePaidStorageResultFailed)(nil),
			2: (*IncreasePaidStorageResultSkipped)(nil),
			3: (*IncreasePaidStorageResultBacktracked)(nil),
		},
	})
}

type IncreasePaidStorageContentsAndResult struct {
	IncreasePaidStorage
	Metadata ManagerMetadata[IncreasePaidStorageResult, *BalanceUpdate]
}

func (*IncreasePaidStorageContentsAndResult) OperationContentsAndResult() {}

type DoubleBakingEvidence struct {
	Block1 BlockHeader `tz:"dyn"`
	Block2 BlockHeader `tz:"dyn"`
}

func (*DoubleBakingEvidence) OperationKind() string { return "double_baking_evidence" }

type DoubleBakingEvidenceContentsAndResult[T core.BalanceUpdate] struct {
	DoubleBakingEvidence
	Metadata []T `tz:"dyn"`
}

func (*DoubleBakingEvidenceContentsAndResult[T]) OperationContentsAndResult() {}

type ActivateAccountContentsAndResult[T core.BalanceUpdate] struct {
	ActivateAccount
	Metadata []T `tz:"dyn"`
}

func (*ActivateAccountContentsAndResult[T]) OperationContentsAndResult() {}

type DoublePreendorsementEvidenceContentsAndResult[T core.BalanceUpdate] struct {
	DoublePreendorsementEvidence
	Metadata []T `tz:"dyn"`
}

func (*DoublePreendorsementEvidenceContentsAndResult[T]) OperationContentsAndResult() {}

type PreendorsementContentsAndResult[T core.BalanceUpdate] struct {
	Preendorsement
	Metadata EndorsementMetadata[T]
}

func (*PreendorsementContentsAndResult[T]) OperationContentsAndResult() {}

type VDFRevelationContentsAndResult[T core.BalanceUpdate] struct {
	VDFRevelation
	Metadata []T `tz:"dyn"`
}

func (*VDFRevelationContentsAndResult[T]) OperationContentsAndResult() {}

type DrainDelegateMetadata[T core.BalanceUpdate] struct {
	BalanceUpdates               []T `tz:"dyn"`
	AllocatedDestinationContract bool
}

type DrainDelegateContentsAndResult[T core.BalanceUpdate] struct {
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

type DALPublishSlotHeaderContentsAndResult struct {
	DALPublishSlotHeader
	Metadata ManagerMetadata[EventResult, *BalanceUpdate]
}

func (*DALPublishSlotHeaderContentsAndResult) OperationContentsAndResult() {}

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
			1:   (*SeedNonceRevelationContentsAndResult[*BalanceUpdate])(nil),
			2:   (*DoubleEndorsementEvidenceContentsAndResult[*BalanceUpdate])(nil),
			3:   (*DoubleBakingEvidenceContentsAndResult[*BalanceUpdate])(nil),
			4:   (*ActivateAccountContentsAndResult[*BalanceUpdate])(nil),
			5:   (*Proposals)(nil),
			6:   (*Ballot)(nil),
			7:   (*DoublePreendorsementEvidenceContentsAndResult[*BalanceUpdate])(nil),
			8:   (*VDFRevelationContentsAndResult[*BalanceUpdate])(nil),
			9:   (*DrainDelegateContentsAndResult[*BalanceUpdate])(nil),
			20:  (*PreendorsementContentsAndResult[*BalanceUpdate])(nil),
			21:  (*EndorsementContentsAndResult[*BalanceUpdate])(nil),
			22:  (*DALAttestationContentsAndResult)(nil),
			107: (*RevealContentsAndResult)(nil),
			108: (*TransactionContentsAndResult[*BalanceUpdate])(nil),
			109: (*OriginationContentsAndResult)(nil),
			110: (*DelegationContentsAndResult)(nil),
			111: (*RegisterGlobalConstantContentsAndResult)(nil),
			112: (*SetDepositsLimitContentsAndResult)(nil),
			113: (*IncreasePaidStorageContentsAndResult)(nil),
			114: (*UpdateConsensusKeyContentsAndResult)(nil),
			158: (*TransferTicketContentsAndResult)(nil),
			200: (*SmartRollupOriginateContentsAndResult)(nil),
			201: (*SmartRollupAddMessagesContentsAndResult)(nil),
			202: (*SmartRollupCementContentsAndResult)(nil),
			203: (*SmartRollupPublishContentsAndResult)(nil),
			204: (*SmartRollupRefuteContentsAndResult)(nil),
			205: (*SmartRollupTimeoutContentsAndResult)(nil),
			206: (*SmartRollupExecuteOutboxMessageContentsAndResult)(nil),
			207: (*SmartRollupRecoverBondContentsAndResult)(nil),
			230: (*DALPublishSlotHeaderContentsAndResult)(nil),
			250: (*ZkRollupOriginationContentsAndResult)(nil),
			251: (*ZkRollupPublishContentsAndResult)(nil),
			252: (*ZkRollupUpdateContentsAndResult)(nil),
			255: (*SignaturePrefix)(nil),
		},
	})
}

type SuccessfulManagerOperationResult interface {
	core.SuccessfulManagerOperationResult
}

type ManagerMetadata[R core.OperationResult, U core.BalanceUpdate] struct {
	BalanceUpdates           []U `tz:"dyn"`
	OperationResult          R
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
			2:   (*OriginationResultContents[*BalanceUpdate])(nil),
			3:   (*DelegationResultContents)(nil),
			5:   (*SetDepositsLimitResultContents)(nil),
			6:   (*UpdateConsensusKeyResultContents)(nil),
			9:   (*IncreasePaidStorageResultContents[*BalanceUpdate])(nil),
			200: (*SmartRollupOriginateResultContents[*BalanceUpdate])(nil),
		},
	})
}
