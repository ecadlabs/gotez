package proto_016_PtMumbai

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/expression"
	"github.com/ecadlabs/gotez/protocol/proto"
	"github.com/ecadlabs/gotez/protocol/proto_001_PtCJ7pwo"
	"github.com/ecadlabs/gotez/protocol/proto_005_PsBABY5H"
	"github.com/ecadlabs/gotez/protocol/proto_009_PsFLoren"
	"github.com/ecadlabs/gotez/protocol/proto_011_PtHangz2"
	"github.com/ecadlabs/gotez/protocol/proto_012_Psithaca"
	kathma "github.com/ecadlabs/gotez/protocol/proto_014_PtKathma"
	"github.com/ecadlabs/gotez/protocol/proto_015_PtLimaPt"
)

type OperationContents interface {
	proto.OperationContents
}

type ManagerOperation = proto_005_PsBABY5H.ManagerOperation
type SeedNonceRevelation = proto_005_PsBABY5H.SeedNonceRevelation
type Preendorsement = proto_012_Psithaca.Preendorsement
type InlinedPreendorsement = proto_012_Psithaca.InlinedPreendorsement
type Endorsement = proto_012_Psithaca.Endorsement
type InlinedEndorsement = proto_012_Psithaca.InlinedEndorsement
type DoublePreendorsementEvidence = proto_012_Psithaca.DoublePreendorsementEvidence
type DoubleEndorsementEvidence = proto_012_Psithaca.DoubleEndorsementEvidence
type Reveal = proto_005_PsBABY5H.Reveal
type RevealSuccessfulManagerOperationResult = proto_015_PtLimaPt.RevealSuccessfulManagerOperationResult
type Delegation = proto_005_PsBABY5H.Delegation
type DelegationSuccessfulManagerOperationResult = proto_015_PtLimaPt.DelegationSuccessfulManagerOperationResult
type RegisterGlobalConstant = proto_011_PtHangz2.RegisterGlobalConstant
type SetDepositsLimit = proto_012_Psithaca.SetDepositsLimit
type SetDepositsLimitSuccessfulManagerOperationResult = proto_015_PtLimaPt.SetDepositsLimitSuccessfulManagerOperationResult
type UpdateConsensusKey = proto_015_PtLimaPt.UpdateConsensusKey
type UpdateConsensusKeySuccessfulManagerOperationResult = proto_015_PtLimaPt.UpdateConsensusKeySuccessfulManagerOperationResult
type IncreasePaidStorage = kathma.IncreasePaidStorage
type ActivateAccount = proto_001_PtCJ7pwo.ActivateAccount
type Proposals = proto_001_PtCJ7pwo.Proposals
type BallotKind = proto_001_PtCJ7pwo.BallotKind
type Ballot = proto_001_PtCJ7pwo.Ballot
type VDFRevelation = kathma.VDFRevelation
type DrainDelegate = proto_015_PtLimaPt.DrainDelegate
type FailingNoop = proto_009_PsFLoren.FailingNoop
type EventResult = proto_015_PtLimaPt.EventResult
type EventResultContents = proto_015_PtLimaPt.EventResultContents
type EventInternalOperationResult = proto_015_PtLimaPt.EventInternalOperationResult
type LazyStorageDiff = proto_015_PtLimaPt.LazyStorageDiff
type OperationResult = kathma.OperationResult

const (
	BallotYay  = proto_001_PtCJ7pwo.BallotYay
	BallotNay  = proto_001_PtCJ7pwo.BallotNay
	BallotPass = proto_001_PtCJ7pwo.BallotPass
)

type OperationContentsAndResult interface {
	proto.OperationContentsAndResult
}

type SeedNonceRevelationContentsAndResult struct {
	SeedNonceRevelation
	Metadata []*BalanceUpdate `tz:"dyn"`
}

func (*SeedNonceRevelationContentsAndResult) OperationContentsAndResult() {}

type DoubleEndorsementEvidenceContentsAndResult struct {
	DoubleEndorsementEvidence
	Metadata []*BalanceUpdate `tz:"dyn"`
}

func (*DoubleEndorsementEvidenceContentsAndResult) OperationContentsAndResult() {}

type EndorsementMetadata struct {
	BalanceUpdates   []*BalanceUpdate `tz:"dyn"`
	Delegate         tz.PublicKeyHash
	EndorsementPower int32
	ConsensusKey     tz.PublicKeyHash
}

type EndorsementContentsAndResult struct {
	Endorsement
	Metadata EndorsementMetadata
}

func (*EndorsementContentsAndResult) OperationContentsAndResult() {}

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
	Metadata MetadataWithResult[EventResult]
}

func (*RevealContentsAndResult) OperationContentsAndResult() {}

type DelegationContentsAndResult struct {
	Delegation
	Metadata MetadataWithResult[EventResult]
}

func (*DelegationContentsAndResult) OperationContentsAndResult() {}

type DelegationInternalOperationResult struct {
	Source   TransactionDestination
	Nonce    uint16
	Delegate tz.Option[tz.PublicKeyHash]
	Result   EventResult
}

func (*DelegationInternalOperationResult) InternalOperationResult() {}
func (*DelegationInternalOperationResult) OperationKind() string    { return "delegation" }

type RegisterGlobalConstantResult interface {
	RegisterGlobalConstantResult()
	OperationResult
}

type RegisterGlobalConstantResultContents struct {
	BalanceUpdates   []*BalanceUpdate `tz:"dyn"`
	ConsumedMilligas tz.BigUint
	StorageSize      tz.BigInt
	GlobalAddress    *tz.ScriptExprHash
}

type RegisterGlobalConstantResultApplied struct {
	kathma.OperationResultApplied[RegisterGlobalConstantResultContents]
}

func (*RegisterGlobalConstantResultApplied) RegisterGlobalConstantResult() {}

type RegisterGlobalConstantResultBacktracked struct {
	kathma.OperationResultBacktracked[RegisterGlobalConstantResultContents]
}

func (*RegisterGlobalConstantResultBacktracked) RegisterGlobalConstantResult() {}

type RegisterGlobalConstantResultFailed struct{ kathma.OperationResultFailed }

func (*RegisterGlobalConstantResultFailed) RegisterGlobalConstantResult() {}

type RegisterGlobalConstantResultSkipped struct{ kathma.OperationResultSkipped }

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
	Metadata MetadataWithResult[RegisterGlobalConstantResult]
}

func (*RegisterGlobalConstantContentsAndResult) OperationContentsAndResult() {}

type SetDepositsLimitContentsAndResult struct {
	SetDepositsLimit
	Metadata MetadataWithResult[EventResult]
}

func (*SetDepositsLimitContentsAndResult) OperationContentsAndResult() {}

type UpdateConsensusKeyContentsAndResult struct {
	UpdateConsensusKey
	Metadata MetadataWithResult[EventResult]
}

func (*UpdateConsensusKeyContentsAndResult) OperationContentsAndResult() {}

type TransferTicket struct {
	ManagerOperation
	TicketContents expression.Expression `tz:"dyn"`
	TicketType     expression.Expression `tz:"dyn"`
	TicketTicketer tz.ContractID
	TicketAmount   tz.BigUint
	Destination    tz.ContractID
	Entrypoint     string `tz:"dyn"`
}

func (*TransferTicket) OperationKind() string { return "transfer_ticket" }

type TransferTicketContentsAndResult struct {
	TransferTicket
	Metadata MetadataWithResult[SmartRollupExecuteOutboxMessageResult]
}

func (*TransferTicketContentsAndResult) OperationContentsAndResult() {}

type IncreasePaidStorageResult interface {
	IncreasePaidStorageResult()
	OperationResult
}

type IncreasePaidStorageResultContents struct {
	BalanceUpdates   []*BalanceUpdate `tz:"dyn"`
	ConsumedMilligas tz.BigUint
}

type IncreasePaidStorageResultApplied struct {
	kathma.OperationResultApplied[IncreasePaidStorageResultContents]
}

func (*IncreasePaidStorageResultApplied) IncreasePaidStorageResult() {}

type IncreasePaidStorageResultBacktracked struct {
	kathma.OperationResultBacktracked[IncreasePaidStorageResultContents]
}

func (*IncreasePaidStorageResultBacktracked) IncreasePaidStorageResult() {}

type IncreasePaidStorageResultFailed struct{ kathma.OperationResultFailed }

func (*IncreasePaidStorageResultFailed) IncreasePaidStorageResult() {}

type IncreasePaidStorageResultSkipped struct{ kathma.OperationResultSkipped }

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
	Metadata MetadataWithResult[IncreasePaidStorageResult]
}

func (*IncreasePaidStorageContentsAndResult) OperationContentsAndResult() {}

type IncreasePaidStorageSuccessfulManagerOperationResult IncreasePaidStorageResultContents

func (*IncreasePaidStorageSuccessfulManagerOperationResult) SuccessfulManagerOperationResult() {}
func (*IncreasePaidStorageSuccessfulManagerOperationResult) OperationKind() string {
	return "increase_paid_storage"
}

type DoubleBakingEvidence struct {
	Block1 BlockHeader `tz:"dyn"`
	Block2 BlockHeader `tz:"dyn"`
}

func (*DoubleBakingEvidence) OperationKind() string { return "double_baking_evidence" }

type DoubleBakingEvidenceContentsAndResult struct {
	DoubleBakingEvidence
	Metadata []*BalanceUpdate `tz:"dyn"`
}

func (*DoubleBakingEvidenceContentsAndResult) OperationContentsAndResult() {}

type ActivateAccountContentsAndResult struct {
	ActivateAccount
	Metadata []*BalanceUpdate `tz:"dyn"`
}

func (*ActivateAccountContentsAndResult) OperationContentsAndResult() {}

type DoublePreendorsementEvidenceContentsAndResult struct {
	DoublePreendorsementEvidence
	Metadata []*BalanceUpdate `tz:"dyn"`
}

func (*DoublePreendorsementEvidenceContentsAndResult) OperationContentsAndResult() {}

type PreendorsementMetadata = EndorsementMetadata
type PreendorsementContentsAndResult struct {
	Preendorsement
	Metadata PreendorsementMetadata
}

func (*PreendorsementContentsAndResult) OperationContentsAndResult() {}

type VDFRevelationContentsAndResult struct {
	VDFRevelation
	Metadata []*BalanceUpdate `tz:"dyn"`
}

func (*VDFRevelationContentsAndResult) OperationContentsAndResult() {}

type DrainDelegateMetadata struct {
	BalanceUpdates               []*BalanceUpdate `tz:"dyn"`
	AllocatedDestinationContract bool
}

type DrainDelegateContentsAndResult struct {
	DrainDelegate
	Metadata DrainDelegateMetadata
}

func (*DrainDelegateContentsAndResult) OperationContentsAndResult() {}

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
	Metadata MetadataWithResult[EventResult]
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
			1:   (*SeedNonceRevelationContentsAndResult)(nil),
			2:   (*DoubleEndorsementEvidenceContentsAndResult)(nil),
			3:   (*DoubleBakingEvidenceContentsAndResult)(nil),
			4:   (*ActivateAccountContentsAndResult)(nil),
			5:   (*Proposals)(nil),
			6:   (*Ballot)(nil),
			7:   (*DoublePreendorsementEvidenceContentsAndResult)(nil),
			8:   (*VDFRevelationContentsAndResult)(nil),
			9:   (*DrainDelegateContentsAndResult)(nil),
			20:  (*PreendorsementContentsAndResult)(nil),
			21:  (*EndorsementContentsAndResult)(nil),
			22:  (*DALAttestationContentsAndResult)(nil),
			107: (*RevealContentsAndResult)(nil),
			108: (*TransactionContentsAndResult)(nil),
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
	proto.SuccessfulManagerOperationResult
}

type MetadataWithResult[T OperationResult] struct {
	BalanceUpdates           []*BalanceUpdate `tz:"dyn"`
	OperationResult          T
	InternalOperationResults []InternalOperationResult `tz:"dyn"`
}

type InternalOperationResult interface {
	proto.InternalOperationResult
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
			0:   (*RevealSuccessfulManagerOperationResult)(nil),
			1:   (*TransactionSuccessfulManagerOperationResult)(nil),
			2:   (*OriginationSuccessfulManagerOperationResult)(nil),
			3:   (*DelegationSuccessfulManagerOperationResult)(nil),
			5:   (*SetDepositsLimitSuccessfulManagerOperationResult)(nil),
			6:   (*UpdateConsensusKeySuccessfulManagerOperationResult)(nil),
			9:   (*IncreasePaidStorageSuccessfulManagerOperationResult)(nil),
			200: (*SmartRollupOriginateSuccessfulManagerOperationResult)(nil),
		},
	})
}
