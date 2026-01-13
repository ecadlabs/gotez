package proto_023_PtSeouLo

//go:generate go run ../../cmd/genmarshaller.go

import (
	"strconv"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/proto_022_PsRiotum"
)

type ManagerOperation = proto_022_PsRiotum.ManagerOperation
type SeedNonceRevelation = proto_022_PsRiotum.SeedNonceRevelation
type DoubleBakingEvidence = proto_022_PsRiotum.DoubleBakingEvidence
type ActivateAccount = proto_022_PsRiotum.ActivateAccount
type Proposals = proto_022_PsRiotum.Proposals
type Ballot = proto_022_PsRiotum.Ballot
type VDFRevelation = proto_022_PsRiotum.VDFRevelation
type DrainDelegate = proto_022_PsRiotum.DrainDelegate
type FailingNoop = proto_022_PsRiotum.FailingNoop
type ShardWithProof = proto_022_PsRiotum.ShardWithProof
type Origination = proto_022_PsRiotum.Origination
type Delegation = proto_022_PsRiotum.Delegation
type RegisterGlobalConstant = proto_022_PsRiotum.RegisterGlobalConstant
type IncreasePaidStorage = proto_022_PsRiotum.IncreasePaidStorage
type SetDepositsLimit = proto_022_PsRiotum.SetDepositsLimit
type TransferTicket = proto_022_PsRiotum.TransferTicket
type DALPublishCommitment = proto_022_PsRiotum.DALPublishCommitment
type BalanceUpdate = proto_022_PsRiotum.BalanceUpdate
type BalanceUpdates = proto_022_PsRiotum.BalanceUpdates
type ConsensusContent = proto_022_PsRiotum.ConsensusContent

type InternalOperationResult = proto_022_PsRiotum.InternalOperationResult
type OriginationInternalOperationResult = proto_022_PsRiotum.OriginationInternalOperationResult
type DelegationInternalOperationResult = proto_022_PsRiotum.DelegationInternalOperationResult
type EventInternalOperationResult = proto_022_PsRiotum.EventInternalOperationResult

type OperationContents interface {
	core.OperationContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[OperationContents]{
		Variants: encoding.Variants[OperationContents]{
			1:   (*SeedNonceRevelation)(nil),
			2:   (*DoubleConsensusOperationEvidence)(nil),
			3:   (*DoubleBakingEvidence)(nil),
			4:   (*ActivateAccount)(nil),
			5:   (*Proposals)(nil),
			6:   (*Ballot)(nil),
			8:   (*VDFRevelation)(nil),
			9:   (*DrainDelegate)(nil),
			17:  (*FailingNoop)(nil),
			20:  (*Preattestation)(nil),
			21:  (*Attestation)(nil),
			23:  (*AttestationWithDAL)(nil),
			24:  (*DALEntrapmentEvidence)(nil),
			30:  (*PreattestationsAggregate)(nil),
			31:  (*AttestationsAggregate)(nil),
			40:  (*BLSModePreattestation)(nil),
			41:  (*BLSModeAttestation)(nil),
			107: (*Reveal)(nil),
			108: (*Transaction)(nil),
			109: (*Origination)(nil),
			110: (*Delegation)(nil),
			111: (*RegisterGlobalConstant)(nil),
			112: (*SetDepositsLimit)(nil),
			113: (*IncreasePaidStorage)(nil),
			114: (*UpdateConsensusKey)(nil),
			115: (*UpdateCompanionKey)(nil),
			158: (*TransferTicket)(nil),
			200: (*SmartRollupOriginate)(nil),
			201: (*SmartRollupAddMessages)(nil),
			202: (*SmartRollupCement)(nil),
			203: (*SmartRollupPublish)(nil),
			204: (*SmartRollupRefute)(nil),
			205: (*SmartRollupTimeout)(nil),
			206: (*SmartRollupExecuteOutboxMessage)(nil),
			207: (*SmartRollupRecoverBond)(nil),
			230: (*DALPublishCommitment)(nil),
			250: (*ZkRollupOrigination)(nil),
			251: (*ZkRollupPublish)(nil),
			252: (*ZkRollupUpdate)(nil),
			255: (*core.SignaturePrefix)(nil),
		},
	})
}

//json:kind=OperationKind()
type DALEntrapmentEvidence struct {
	Attestation    InlinedConsensusOperation `tz:"dyn" json:"attestation"`
	ConsensusSlot  uint16                    `json:"consensus_slot"`
	SlotIndex      uint8                     `json:"slot_index"`
	ShardWithProof ShardWithProof            `json:"shard_with_proof"`
}

func (*DALEntrapmentEvidence) OperationKind() string { return "dal_entrapment_evidence" }

type Proof struct {
	Signature *tz.BLSSignature `tz:"dyn" json:"signature"`
}

//json:kind=OperationKind()
type Reveal struct {
	ManagerOperation
	PublicKey tz.PublicKey     `json:"public_key"`
	Proof     tz.Option[Proof] `json:"proof"`
}

func (*Reveal) OperationKind() string { return "reveal" }

//json:kind=OperationKind()
type UpdateConsensusKey struct {
	ManagerOperation
	PublicKey tz.PublicKey     `json:"public_key"`
	Proof     tz.Option[Proof] `json:"proof"`
}

func (*UpdateConsensusKey) OperationKind() string { return "update_consensus_key" }

//json:kind=OperationKind()
type UpdateCompanionKey UpdateConsensusKey

func (*UpdateCompanionKey) OperationKind() string { return "update_companion_key" }

//json:kind=OperationKind()
type Attestation proto_022_PsRiotum.Attestation

func (*Attestation) InlinedConsensusOperationContent() {}
func (*Attestation) OperationKind() string             { return "attestation" }

//json:kind=OperationKind()
type Preattestation Attestation

func (*Preattestation) InlinedConsensusOperationContent() {}
func (*Preattestation) OperationKind() string             { return "preattestation" }

//json:kind=OperationKind()
type AttestationsAggregate struct {
	ConsensusContent ConsensusContent `json:"consensus_content"`
	Committee        []*Committee     `tz:"dyn" json:"committee"`
}

func (*AttestationsAggregate) InlinedConsensusOperationContent() {}
func (*AttestationsAggregate) OperationKind() string             { return "attestations_aggregate" }

//json:kind=OperationKind()
type PreattestationsAggregate struct {
	ConsensusContent ConsensusContent `json:"consensus_content"`
	Committee        []uint16         `tz:"dyn" json:"committee"`
}

func (*PreattestationsAggregate) InlinedConsensusOperationContent() {}
func (*PreattestationsAggregate) OperationKind() string             { return "preattestations_aggregate" }

type Committee struct {
	Slot           uint16               `json:"slot"`
	DALAttestation tz.Option[tz.BigInt] `json:"dal_attestation"`
}

//json:kind=OperationKind()
type AttestationWithDAL proto_022_PsRiotum.AttestationWithDAL

func (*AttestationWithDAL) InlinedConsensusOperationContent() {}
func (*AttestationWithDAL) OperationKind() string             { return "attestation_with_dal" }

//json:kind=OperationKind()
type BLSModeAttestation struct {
	Level            int32                `json:"level"`
	Round            int32                `json:"round"`
	BlockPayloadHash *tz.BlockPayloadHash `json:"block_payload_hash"`
}

func (*BLSModeAttestation) InlinedConsensusOperationContent() {}
func (*BLSModeAttestation) OperationKind() string             { return "attestation" }

//json:kind=OperationKind()
type BLSModePreattestation BLSModeAttestation

func (*BLSModePreattestation) InlinedConsensusOperationContent() {}
func (*BLSModePreattestation) OperationKind() string             { return "preattestation" }

type DoubleConsensusOperationEvidence struct {
	Slot uint16                    `json:"slot"`
	Op1  InlinedConsensusOperation `tz:"dyn" json:"op1"`
	Op2  InlinedConsensusOperation `tz:"dyn" json:"op2"`
}

func (*DoubleConsensusOperationEvidence) OperationKind() string {
	return "double_consensus_operation_evidence"
}

type InlinedConsensusOperation struct {
	Branch     *tz.BlockHash                    `json:"branch"`
	Operations InlinedConsensusOperationContent `json:"operations"`
	Signature  tz.AnySignature                  `json:"signature"`
}

type InlinedConsensusOperationContent interface {
	InlinedConsensusOperationContent()
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[InlinedConsensusOperationContent]{
		Variants: encoding.Variants[InlinedConsensusOperationContent]{
			20: (*Preattestation)(nil),
			21: (*Attestation)(nil),
			23: (*AttestationWithDAL)(nil),
			30: (*PreattestationsAggregate)(nil),
			31: (*AttestationsAggregate)(nil),
		},
	})
}

type SeedNonceRevelationContentsAndResult = proto_022_PsRiotum.SeedNonceRevelationContentsAndResult
type ActivateAccountContentsAndResult = proto_022_PsRiotum.ActivateAccountContentsAndResult
type VDFRevelationContentsAndResult = proto_022_PsRiotum.VDFRevelationContentsAndResult
type DrainDelegateContentsAndResult = proto_022_PsRiotum.DrainDelegateContentsAndResult
type AttestationMetadata = proto_022_PsRiotum.AttestationMetadata
type AttestationsAggregateMetadata = proto_022_PsRiotum.AttestationsAggregateMetadata
type ConsumedGasResult = proto_022_PsRiotum.ConsumedGasResult
type OriginationContentsAndResult = proto_022_PsRiotum.OriginationContentsAndResult
type OriginationResult = proto_022_PsRiotum.OriginationResult
type DelegationContentsAndResult = proto_022_PsRiotum.DelegationContentsAndResult
type DelegationResult = proto_022_PsRiotum.DelegationResult
type RegisterGlobalConstantContentsAndResult = proto_022_PsRiotum.RegisterGlobalConstantContentsAndResult
type RegisterGlobalConstantResult = proto_022_PsRiotum.RegisterGlobalConstantResult
type SetDepositsLimitContentsAndResult = proto_022_PsRiotum.SetDepositsLimitContentsAndResult
type IncreasePaidStorageContentsAndResult = proto_022_PsRiotum.IncreasePaidStorageContentsAndResult
type IncreasePaidStorageResult = proto_022_PsRiotum.IncreasePaidStorageResult
type TransferTicketContentsAndResult = proto_022_PsRiotum.TransferTicketContentsAndResult
type TransferTicketResult = proto_022_PsRiotum.TransferTicketResult
type DALPublishCommitmentContentsAndResult = proto_022_PsRiotum.DALPublishCommitmentContentsAndResult
type DALPublishCommitmentResult = proto_022_PsRiotum.DALPublishCommitmentResult
type ManagerMetadata[T core.ManagerOperationResult] = proto_022_PsRiotum.ManagerMetadata[T]

type OperationContentsAndResult interface {
	core.OperationContentsAndResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[OperationContentsAndResult]{
		Variants: encoding.Variants[OperationContentsAndResult]{
			1:   (*SeedNonceRevelationContentsAndResult)(nil),
			2:   (*DoubleConsensusOperationEvidenceContentsAndResult)(nil),
			3:   (*DoubleBakingEvidenceContentsAndResult)(nil),
			4:   (*ActivateAccountContentsAndResult)(nil),
			5:   (*Proposals)(nil),
			6:   (*Ballot)(nil),
			8:   (*VDFRevelationContentsAndResult)(nil),
			9:   (*DrainDelegateContentsAndResult)(nil),
			20:  (*PreattestationContentsAndResult)(nil),
			21:  (*AttestationContentsAndResult)(nil),
			23:  (*AttestationWithDALContentsAndResult)(nil),
			24:  (*DALEntrapmentEvidenceContentsAndResult)(nil),
			30:  (*PreattestationsAggregateContentsAndResult)(nil),
			31:  (*AttestationsAggregateContentsAndResult)(nil),
			107: (*RevealContentsAndResult)(nil),
			108: (*TransactionContentsAndResult)(nil),
			109: (*OriginationContentsAndResult)(nil),
			110: (*DelegationContentsAndResult)(nil),
			111: (*RegisterGlobalConstantContentsAndResult)(nil),
			112: (*SetDepositsLimitContentsAndResult)(nil),
			113: (*IncreasePaidStorageContentsAndResult)(nil),
			114: (*UpdateConsensusKeyContentsAndResult)(nil),
			115: (*UpdateCompanionKeyContentsAndResult)(nil),
			158: (*TransferTicketContentsAndResult)(nil),
			200: (*SmartRollupOriginateContentsAndResult)(nil),
			201: (*SmartRollupAddMessagesContentsAndResult)(nil),
			202: (*SmartRollupCementContentsAndResult)(nil),
			203: (*SmartRollupPublishContentsAndResult)(nil),
			204: (*SmartRollupRefuteContentsAndResult)(nil),
			205: (*SmartRollupTimeoutContentsAndResult)(nil),
			206: (*SmartRollupExecuteOutboxMessageContentsAndResult)(nil),
			207: (*SmartRollupRecoverBondContentsAndResult)(nil),
			230: (*DALPublishCommitmentContentsAndResult)(nil),
			250: (*ZkRollupOriginationContentsAndResult)(nil),
			251: (*ZkRollupPublishContentsAndResult)(nil),
			252: (*ZkRollupUpdateContentsAndResult)(nil),
			255: (*core.SignaturePrefix)(nil),
		},
	})
}

//json:kind=OperationKind()
type DoubleConsensusOperationEvidenceContentsAndResult struct {
	DoubleConsensusOperationEvidence
	Metadata DoubleConsensusOperationEvidenceMetadata `json:"metadata"`
}

func (*DoubleConsensusOperationEvidenceContentsAndResult) OperationContentsAndResult() {}
func (op *DoubleConsensusOperationEvidenceContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type DoubleConsensusOperationEvidenceMetadata struct {
	PunishedDelegate tz.PublicKeyHash `json:"punished_delegate"`
	RewardedDelegate tz.PublicKeyHash `json:"rewarded_delegate"`
	Misbehaviour     Misbehaviour     `json:"misbehaviour"`
}

type Misbehaviour struct {
	Level int32            `json:"level"`
	Round int32            `json:"round"`
	Kind  MisbehaviourKind `json:"kind"`
}

type MisbehaviourKind uint8

const (
	MisbehaviourPreattestation MisbehaviourKind = iota
	MisbehaviourAttestation
	MisbehaviourBlock
)

func (k MisbehaviourKind) String() string {
	switch k {
	case MisbehaviourPreattestation:
		return "preattestation"
	case MisbehaviourAttestation:
		return "attestation"
	case MisbehaviourBlock:
		return "block"
	default:
		return strconv.FormatUint(uint64(k), 10)
	}
}

//json:kind=OperationKind()
type DoubleBakingEvidenceContentsAndResult struct {
	DoubleBakingEvidence
	Metadata DoubleConsensusOperationEvidenceMetadata `json:"metadata"`
}

func (*DoubleBakingEvidenceContentsAndResult) OperationContentsAndResult() {}
func (op *DoubleBakingEvidenceContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type AttestationContentsAndResult struct {
	Attestation
	Metadata AttestationMetadata `json:"metadata"`
}

func (*AttestationContentsAndResult) OperationContentsAndResult() {}
func (op *AttestationContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type PreattestationMetadata = AttestationMetadata

//json:kind=OperationKind()
type PreattestationContentsAndResult struct {
	Preattestation
	Metadata PreattestationMetadata `json:"metadata"`
}

func (*PreattestationContentsAndResult) OperationContentsAndResult() {}
func (op *PreattestationContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type AttestationWithDALContentsAndResult struct {
	AttestationWithDAL
	Metadata AttestationMetadata `json:"metadata"`
}

func (*AttestationWithDALContentsAndResult) OperationContentsAndResult() {}
func (op *AttestationWithDALContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type DALEntrapmentEvidenceContentsAndResult struct {
	DALEntrapmentEvidence
	Metadata BalanceUpdates `json:"metadata"`
}

func (*DALEntrapmentEvidenceContentsAndResult) OperationContentsAndResult() {}
func (op *DALEntrapmentEvidenceContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type AttestationsAggregateContentsAndResult struct {
	AttestationsAggregate
	Metadata AttestationsAggregateMetadata `json:"metadata"`
}

func (*AttestationsAggregateContentsAndResult) OperationContentsAndResult() {}
func (op *AttestationsAggregateContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type PreattestationsAggregateMetadata = AttestationsAggregateMetadata

//json:kind=OperationKind()
type PreattestationsAggregateContentsAndResult struct {
	PreattestationsAggregate
	Metadata PreattestationsAggregateMetadata `json:"metadata"`
}

func (*PreattestationsAggregateContentsAndResult) OperationContentsAndResult() {}
func (op *PreattestationsAggregateContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type RevealContentsAndResult struct {
	Reveal
	Metadata ManagerMetadata[ConsumedGasResult] `json:"metadata"`
}

func (*RevealContentsAndResult) OperationContentsAndResult() {}
func (op *RevealContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type UpdateConsensusKeyContentsAndResult struct {
	UpdateConsensusKey
	Metadata ManagerMetadata[UpdateConsensusKeyResult] `json:"metadata"`
}

func (*UpdateConsensusKeyContentsAndResult) OperationContentsAndResult() {}
func (op *UpdateConsensusKeyContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type UpdateConsensusKeyResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[UpdateConsensusKeyResult]{
		Variants: encoding.Variants[UpdateConsensusKeyResult]{
			0: (*core.OperationResultApplied[*UpdateConsensusKeyResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*UpdateConsensusKeyResultContents])(nil),
		},
	})
}

type UpdateConsensusKeyResultContents struct {
	Kind             bool       `json:"kind"`
	ConsumedMilligas tz.BigUint `json:"consumed_milligas"`
}

//json:kind=OperationKind()
type UpdateConsensusKeySuccessfulManagerResult struct {
	core.OperationResultApplied[*UpdateConsensusKeyResultContents]
}

func (*UpdateConsensusKeySuccessfulManagerResult) OperationKind() string {
	return "update_consensus_key"
}

//json:kind=OperationKind()
type UpdateCompanionKeyContentsAndResult struct {
	UpdateCompanionKey
	Metadata ManagerMetadata[UpdateConsensusKeyResult] `json:"metadata"`
}

func (*UpdateCompanionKeyContentsAndResult) OperationContentsAndResult() {}
func (op *UpdateCompanionKeyContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type SuccessfulManagerOperationResult interface {
	core.SuccessfulManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SuccessfulManagerOperationResult]{
		Variants: encoding.Variants[SuccessfulManagerOperationResult]{
			0:   (*RevealSuccessfulManagerResult)(nil),
			1:   (*TransactionSuccessfulManagerResult)(nil),
			2:   (*OriginationSuccessfulManagerResult)(nil),
			3:   (*DelegationSuccessfulManagerResult)(nil),
			5:   (*SetDepositsLimitSuccessfulManagerResult)(nil),
			6:   (*UpdateConsensusKeySuccessfulManagerResult)(nil),
			9:   (*IncreasePaidStorageSuccessfulManagerResult)(nil),
			200: (*SmartRollupOriginateSuccessfulManagerResult)(nil),
		},
	})
}

type RevealSuccessfulManagerResult = proto_022_PsRiotum.RevealSuccessfulManagerResult
type DelegationSuccessfulManagerResult = proto_022_PsRiotum.DelegationSuccessfulManagerResult
type TransactionSuccessfulManagerResult = proto_022_PsRiotum.TransactionSuccessfulManagerResult
type OriginationSuccessfulManagerResult = proto_022_PsRiotum.OriginationSuccessfulManagerResult
type IncreasePaidStorageSuccessfulManagerResult = proto_022_PsRiotum.IncreasePaidStorageSuccessfulManagerResult
type SmartRollupOriginateSuccessfulManagerResult = proto_022_PsRiotum.SmartRollupOriginateSuccessfulManagerResult
type SetDepositsLimitSuccessfulManagerResult = proto_022_PsRiotum.SetDepositsLimitSuccessfulManagerResult

var ListOperations = core.ListOperations[OperationContents]
