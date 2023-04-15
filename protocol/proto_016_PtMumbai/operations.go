package proto_016_PtMumbai

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/expression"
	"github.com/ecadlabs/gotez/protocol/proto"
)

type OperationContents interface {
	proto.OperationContents
}

type OperationContentsAndResult interface {
	proto.OperationContentsAndResult
}

type SeedNonceRevelation struct {
	Level int32
	Nonce *[tz.SeedNonceBytesLen]byte
}

func (*SeedNonceRevelation) OperationKind() string { return "seed_nonce_revelation" }

type SeedNonceRevelationContentsAndResult struct {
	SeedNonceRevelation
	Metadata []*BalanceUpdate `tz:"dyn"`
}

func (*SeedNonceRevelationContentsAndResult) OperationContentsAndResult() {}

type DoubleEndorsementEvidence struct {
	Op1 InlinedEndorsement `tz:"dyn"`
	Op2 InlinedEndorsement `tz:"dyn"`
}

func (*DoubleEndorsementEvidence) OperationKind() string { return "double_endorsement_evidence" }

type DoubleEndorsementEvidenceContentsAndResult struct {
	DoubleEndorsementEvidence
	Metadata []*BalanceUpdate `tz:"dyn"`
}

func (*DoubleEndorsementEvidenceContentsAndResult) OperationContentsAndResult() {}

type InlinedEndorsement struct {
	Branch    *tz.BlockHash
	Contents  InlinedEndorsementContents
	Signature *tz.GenericSignature
}

type InlinedEndorsementContents interface {
	InlinedEndorsementContents()
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[InlinedEndorsementContents]{
		Variants: encoding.Variants[InlinedEndorsementContents]{
			21: (*Endorsement)(nil),
		},
	})
}

type Endorsement struct {
	Slot             uint16
	Level            int32
	Round            int32
	BlockPayloadHash *tz.BlockPayloadHash
}

func (*Endorsement) InlinedEndorsementContents() {}
func (*Endorsement) OperationKind() string       { return "endorsement" }

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

func (*DALAttestation) OperationKind() string { return "dal_attestation " }

type DALAttestationContentsAndResult struct {
	DALAttestation
	Metadata tz.PublicKeyHash
}

func (*DALAttestationContentsAndResult) OperationContentsAndResult() {}

type Reveal struct {
	ManagerOperation
	PublicKey tz.PublicKey
}

func (*Reveal) OperationKind() string { return "reveal" }

type RevealContentsAndResult struct {
	Reveal
	Metadata MetadataWithResult[EventResult]
}

func (*RevealContentsAndResult) OperationContentsAndResult() {}

type RevealSuccessfulManagerOperationResult EventResultContents

func (*RevealSuccessfulManagerOperationResult) SuccessfulManagerOperationResult() {}
func (*RevealSuccessfulManagerOperationResult) OperationKind() string             { return "reveal" }

type Delegation struct {
	ManagerOperation
	Delegate tz.Option[tz.PublicKeyHash]
}

func (*Delegation) OperationKind() string { return "delegation" }

type DelegationContentsAndResult struct {
	Delegation
	Metadata MetadataWithResult[EventResult]
}

func (*DelegationContentsAndResult) OperationContentsAndResult() {}

type DelegationInternalOperationResult struct {
	Source   tz.TransactionDestination
	Nonce    uint16
	Delegate tz.Option[tz.PublicKeyHash]
	Result   EventResult
}

func (*DelegationInternalOperationResult) InternalOperationResult() {}
func (*DelegationInternalOperationResult) OperationKind() string    { return "delegation" }

type DelegationSuccessfulManagerOperationResult EventResultContents

func (*DelegationSuccessfulManagerOperationResult) SuccessfulManagerOperationResult() {}
func (*DelegationSuccessfulManagerOperationResult) OperationKind() string             { return "delegation" }

type RegisterGlobalConstant struct {
	ManagerOperation
	Value expression.Expression `tz:"dyn"`
}

func (*RegisterGlobalConstant) OperationKind() string { return "register_global_constant" }

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
	OperationResultApplied[RegisterGlobalConstantResultContents]
}

func (*RegisterGlobalConstantResultApplied) RegisterGlobalConstantResult() {}

type RegisterGlobalConstantResultBacktracked struct {
	OperationResultBacktracked[RegisterGlobalConstantResultContents]
}

func (*RegisterGlobalConstantResultBacktracked) RegisterGlobalConstantResult() {}

type RegisterGlobalConstantResultFailed struct{ OperationResultFailed }

func (*RegisterGlobalConstantResultFailed) RegisterGlobalConstantResult() {}

type RegisterGlobalConstantResultSkipped struct{ OperationResultSkipped }

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

type SetDepositsLimit struct {
	ManagerOperation
	Limit tz.Option[tz.BigUint]
}

func (*SetDepositsLimit) OperationKind() string { return "set_deposits_limit" }

type SetDepositsLimitContentsAndResult struct {
	SetDepositsLimit
	Metadata MetadataWithResult[EventResult]
}

func (*SetDepositsLimitContentsAndResult) OperationContentsAndResult() {}

type SetDepositsLimitSuccessfulManagerOperationResult EventResultContents

func (*SetDepositsLimitSuccessfulManagerOperationResult) SuccessfulManagerOperationResult() {}
func (*SetDepositsLimitSuccessfulManagerOperationResult) OperationKind() string {
	return "set_deposits_limit"
}

type UpdateConsensusKey struct {
	ManagerOperation
	PublicKey tz.PublicKey
}

func (*UpdateConsensusKey) OperationKind() string { return "update_consensus_key" }

type UpdateConsensusKeyContentsAndResult struct {
	UpdateConsensusKey
	Metadata MetadataWithResult[EventResult]
}

func (*UpdateConsensusKeyContentsAndResult) OperationContentsAndResult() {}

type UpdateConsensusKeySuccessfulManagerOperationResult EventResultContents

func (*UpdateConsensusKeySuccessfulManagerOperationResult) SuccessfulManagerOperationResult() {}
func (*UpdateConsensusKeySuccessfulManagerOperationResult) OperationKind() string {
	return "update_consensus_key"
}

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

type IncreasePaidStorage struct {
	ManagerOperation
	Amount      tz.BigInt
	Destination tz.OriginatedContractID
}

func (*IncreasePaidStorage) OperationKind() string { return "increase_paid_storage" }

type IncreasePaidStorageResult interface {
	IncreasePaidStorageResult()
	OperationResult
}

type IncreasePaidStorageResultContents struct {
	BalanceUpdates   []*BalanceUpdate `tz:"dyn"`
	ConsumedMilligas tz.BigUint
}

type IncreasePaidStorageResultApplied struct {
	OperationResultApplied[IncreasePaidStorageResultContents]
}

func (*IncreasePaidStorageResultApplied) IncreasePaidStorageResult() {}

type IncreasePaidStorageResultBacktracked struct {
	OperationResultBacktracked[IncreasePaidStorageResultContents]
}

func (*IncreasePaidStorageResultBacktracked) IncreasePaidStorageResult() {}

type IncreasePaidStorageResultFailed struct{ OperationResultFailed }

func (*IncreasePaidStorageResultFailed) IncreasePaidStorageResult() {}

type IncreasePaidStorageResultSkipped struct{ OperationResultSkipped }

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

type ActivateAccount struct {
	PKH    *tz.Ed25519PublicKeyHash
	Secret *[tz.SecretBytesLen]byte
}

func (*ActivateAccount) OperationKind() string { return "activate_account" }

type ActivateAccountContentsAndResult struct {
	ActivateAccount
	Metadata []*BalanceUpdate `tz:"dyn"`
}

func (*ActivateAccountContentsAndResult) OperationContentsAndResult() {}

type Proposals struct {
	Source    tz.PublicKeyHash
	Period    int32
	Proposals []*tz.ProtocolHash `tz:"dyn"`
}

func (*Proposals) OperationKind() string       { return "proposals" }
func (*Proposals) OperationContentsAndResult() {}

type BallotKind uint8

const (
	BallotYay BallotKind = iota
	BallotNay
	BallotPass
)

type Ballot struct {
	Source   tz.PublicKeyHash
	Period   int32
	Proposal *tz.ProtocolHash
	Ballot   BallotKind
}

func (*Ballot) OperationKind() string       { return "ballot" }
func (*Ballot) OperationContentsAndResult() {}

type DoublePreendorsementEvidence struct {
	Op1 InlinedPreendorsement `tz:"dyn"`
	Op2 InlinedPreendorsement `tz:"dyn"`
}

func (*DoublePreendorsementEvidence) OperationKind() string { return "double_preendorsement_evidence" }

type DoublePreendorsementEvidenceContentsAndResult struct {
	DoublePreendorsementEvidence
	Metadata []*BalanceUpdate `tz:"dyn"`
}

func (*DoublePreendorsementEvidenceContentsAndResult) OperationContentsAndResult() {}

type InlinedPreendorsement struct {
	Branch    *tz.BlockHash
	Contents  InlinedPreendorsementContents
	Signature *tz.GenericSignature
}

type InlinedPreendorsementContents interface {
	InlinedPreendorsementContents()
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[InlinedPreendorsementContents]{
		Variants: encoding.Variants[InlinedPreendorsementContents]{
			20: (*Preendorsement)(nil),
		},
	})
}

type Preendorsement struct {
	Slot             uint16
	Level            int32
	Round            int32
	BlockPayloadHash *tz.BlockPayloadHash
}

func (*Preendorsement) InlinedPreendorsementContents() {}
func (*Preendorsement) OperationKind() string          { return "preendorsement" }

type PreendorsementMetadata = EndorsementMetadata
type PreendorsementContentsAndResult struct {
	Preendorsement
	Metadata PreendorsementMetadata
}

func (*PreendorsementContentsAndResult) OperationContentsAndResult() {}

type VDFRevelation struct {
	Solution [2]*[100]byte
}

func (*VDFRevelation) OperationKind() string { return "vdf_revelation" }

type VDFRevelationContentsAndResult struct {
	VDFRevelation
	Metadata []*BalanceUpdate `tz:"dyn"`
}

func (*VDFRevelationContentsAndResult) OperationContentsAndResult() {}

type DrainDelegate struct {
	ConsensusKey tz.PublicKeyHash
	Delegate     tz.PublicKeyHash
	Destination  tz.PublicKeyHash
}

func (*DrainDelegate) OperationKind() string { return "drain_delegate" }

type DrainDelegateMetadata struct {
	BalanceUpdates               []*BalanceUpdate `tz:"dyn"`
	AllocatedDestinationContract bool
}

type DrainDelegateContentsAndResult struct {
	DrainDelegate
	Metadata DrainDelegateMetadata
}

func (*DrainDelegateContentsAndResult) OperationContentsAndResult() {}

type FailingNoop struct {
	Arbitrary []byte `tz:"dyn"`
}

func (*FailingNoop) OperationKind() string { return "failing_noop" }

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

type ManagerOperation struct {
	Source       tz.PublicKeyHash
	Fee          tz.BigUint
	Counter      tz.BigUint
	GasLimit     tz.BigUint
	StorageLimit tz.BigUint
}

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

type EventResult interface {
	EventResult()
	OperationResult
}

type EventResultContents struct {
	ConsumedMilligas tz.BigUint
}

type EventResultApplied struct {
	OperationResultApplied[EventResultContents]
}

func (*EventResultApplied) EventResult() {}

type EventResultBacktracked struct {
	OperationResultBacktracked[EventResultContents]
}

func (*EventResultBacktracked) EventResult() {}

type EventResultFailed struct{ OperationResultFailed }

func (*EventResultFailed) EventResult() {}

type EventResultSkipped struct{ OperationResultSkipped }

func (*EventResultSkipped) EventResult() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[EventResult]{
		Variants: encoding.Variants[EventResult]{
			0: (*EventResultApplied)(nil),
			1: (*EventResultFailed)(nil),
			2: (*EventResultSkipped)(nil),
			3: (*EventResultBacktracked)(nil),
		},
	})
}

type MetadataWithResult[T OperationResult] struct {
	BalanceUpdates           []*BalanceUpdate `tz:"dyn"`
	OperationResult          T
	InternalOperationResults []InternalOperationResult `tz:"dyn"`
}

type InternalOperationResult interface {
	proto.InternalOperationResult
}

type EventInternalOperationResult struct {
	Source  tz.TransactionDestination
	Nonce   uint16
	Type    expression.Expression
	Tag     tz.Option[Entrypoint]
	Payload tz.Option[expression.Expression]
	Result  EventResult
}

func (*EventInternalOperationResult) InternalOperationResult() {}
func (*EventInternalOperationResult) OperationKind() string    { return "event" }

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

type LazyStorageDiff struct {
	Opaque []byte `tz:"dyn"` // TODO: lazy storage diff
}
