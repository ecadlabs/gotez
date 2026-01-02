package proto_024_PtTALLiN

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/proto_023_PtSeouLo"
)

//go:generate go run ../../cmd/genmarshaller.go

type SeedNonceRevelation = proto_023_PtSeouLo.SeedNonceRevelation
type DoubleConsensusOperationEvidence = proto_023_PtSeouLo.DoubleConsensusOperationEvidence
type DoubleBakingEvidence = proto_023_PtSeouLo.DoubleBakingEvidence
type ActivateAccount = proto_023_PtSeouLo.ActivateAccount
type Proposals = proto_023_PtSeouLo.Proposals
type Ballot = proto_023_PtSeouLo.Ballot
type VDFRevelation = proto_023_PtSeouLo.VDFRevelation
type DrainDelegate = proto_023_PtSeouLo.DrainDelegate
type FailingNoop = proto_023_PtSeouLo.FailingNoop
type Preattestation = proto_023_PtSeouLo.Preattestation
type Attestation = proto_023_PtSeouLo.Attestation
type AttestationWithDAL = proto_023_PtSeouLo.AttestationWithDAL
type DALEntrapmentEvidence = proto_023_PtSeouLo.DALEntrapmentEvidence
type PreattestationsAggregate = proto_023_PtSeouLo.PreattestationsAggregate
type AttestationsAggregate = proto_023_PtSeouLo.AttestationsAggregate
type BLSModePreattestation = proto_023_PtSeouLo.BLSModePreattestation
type BLSModeAttestation = proto_023_PtSeouLo.BLSModeAttestation
type Reveal = proto_023_PtSeouLo.Reveal
type Origination = proto_023_PtSeouLo.Origination
type Delegation = proto_023_PtSeouLo.Delegation
type RegisterGlobalConstant = proto_023_PtSeouLo.RegisterGlobalConstant
type SetDepositsLimit = proto_023_PtSeouLo.SetDepositsLimit
type IncreasePaidStorage = proto_023_PtSeouLo.IncreasePaidStorage
type UpdateConsensusKey = proto_023_PtSeouLo.UpdateConsensusKey
type UpdateCompanionKey = proto_023_PtSeouLo.UpdateCompanionKey
type TransferTicket = proto_023_PtSeouLo.TransferTicket
type DALPublishCommitment = proto_023_PtSeouLo.DALPublishCommitment
type OperationContents = proto_023_PtSeouLo.OperationContents
type BalanceUpdates = proto_023_PtSeouLo.BalanceUpdates
type BalanceUpdate = proto_023_PtSeouLo.BalanceUpdate
type ManagerOperation = proto_023_PtSeouLo.ManagerOperation
type InlinedConsensusOperation = proto_023_PtSeouLo.InlinedConsensusOperation

type AttestationMetadata struct {
	BalanceUpdates
	Delegate       tz.PublicKeyHash `json:"delegate"`
	ConsensusPower ConsensusPower   `json:"consensus_power"`
	ConsensusKey   tz.PublicKeyHash `json:"consensus_key"`
}

type PreattestationMetadata = AttestationMetadata

type ConsensusPower struct {
	Slots       int32             `json:"slots"`
	BakingPower tz.Option1[int64] `json:"baking_power"`
}

type AttestationsAggregateMetadata struct {
	BalanceUpdates
	Committee           []AttestationsAggregateMetadataCommittee `tz:"dyn" json:"committee"`
	TotalConsensusPower ConsensusPower                           `json:"total_consensus_power"`
}

type PreattestationsAggregateMetadata = AttestationsAggregateMetadata

type AttestationsAggregateMetadataCommittee struct {
	Delegate       tz.PublicKeyHash `json:"delegate"`
	ConsensusPKH   tz.PublicKeyHash `json:"consensus_pkh"`
	ConsensusPower ConsensusPower   `json:"consensus_power"`
}

type OperationContentsAndResult interface {
	core.OperationContentsAndResult
}

type SeedNonceRevelationContentsAndResult = proto_023_PtSeouLo.SeedNonceRevelationContentsAndResult
type DoubleConsensusOperationEvidenceContentsAndResult = proto_023_PtSeouLo.DoubleConsensusOperationEvidenceContentsAndResult
type DoubleBakingEvidenceContentsAndResult = proto_023_PtSeouLo.DoubleBakingEvidenceContentsAndResult
type ActivateAccountContentsAndResult = proto_023_PtSeouLo.ActivateAccountContentsAndResult
type VDFRevelationContentsAndResult = proto_023_PtSeouLo.VDFRevelationContentsAndResult
type DrainDelegateContentsAndResult = proto_023_PtSeouLo.DrainDelegateContentsAndResult
type DALEntrapmentEvidenceContentsAndResult = proto_023_PtSeouLo.DALEntrapmentEvidenceContentsAndResult
type ConsumedGasResult = proto_023_PtSeouLo.ConsumedGasResult
type OriginationResult = proto_023_PtSeouLo.OriginationResult
type DelegationResult = proto_023_PtSeouLo.DelegationResult
type RegisterGlobalConstantResult = proto_023_PtSeouLo.RegisterGlobalConstantResult
type IncreasePaidStorageResult = proto_023_PtSeouLo.IncreasePaidStorageResult
type UpdateConsensusKeyResult = proto_023_PtSeouLo.UpdateConsensusKeyResult
type TransferTicketResult = proto_023_PtSeouLo.TransferTicketResult
type DALPublishCommitmentResult = proto_023_PtSeouLo.DALPublishCommitmentResult

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
type AttestationContentsAndResult struct {
	Attestation
	Metadata AttestationMetadata `json:"metadata"`
}

func (*AttestationContentsAndResult) OperationContentsAndResult() {}
func (op *AttestationContentsAndResult) GetMetadata() any {
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
type AttestationsAggregateContentsAndResult struct {
	AttestationsAggregate
	Metadata AttestationsAggregateMetadata `json:"metadata"`
}

func (*AttestationsAggregateContentsAndResult) OperationContentsAndResult() {}
func (op *AttestationsAggregateContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

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
type OriginationContentsAndResult struct {
	Origination
	Metadata ManagerMetadata[OriginationResult] `json:"metadata"`
}

func (*OriginationContentsAndResult) OperationContentsAndResult() {}
func (op *OriginationContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type DelegationContentsAndResult struct {
	Delegation
	Metadata ManagerMetadata[DelegationResult] `json:"metadata"`
}

func (*DelegationContentsAndResult) OperationContentsAndResult() {}
func (op *DelegationContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type RegisterGlobalConstantContentsAndResult struct {
	RegisterGlobalConstant
	Metadata ManagerMetadata[RegisterGlobalConstantResult] `json:"metadata"`
}

func (*RegisterGlobalConstantContentsAndResult) OperationContentsAndResult() {}
func (op *RegisterGlobalConstantContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type SetDepositsLimitContentsAndResult struct {
	SetDepositsLimit
	Metadata ManagerMetadata[ConsumedGasResult] `json:"metadata"`
}

func (*SetDepositsLimitContentsAndResult) OperationContentsAndResult() {}
func (op *SetDepositsLimitContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type IncreasePaidStorageContentsAndResult struct {
	IncreasePaidStorage
	Metadata ManagerMetadata[IncreasePaidStorageResult] `json:"metadata"`
}

func (*IncreasePaidStorageContentsAndResult) OperationContentsAndResult() {}
func (op *IncreasePaidStorageContentsAndResult) GetMetadata() any {
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

//json:kind=OperationKind()
type UpdateCompanionKeyContentsAndResult struct {
	UpdateCompanionKey
	Metadata ManagerMetadata[UpdateConsensusKeyResult] `json:"metadata"`
}

func (*UpdateCompanionKeyContentsAndResult) OperationContentsAndResult() {}
func (op *UpdateCompanionKeyContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type TransferTicketContentsAndResult struct {
	TransferTicket
	Metadata ManagerMetadata[TransferTicketResult] `json:"metadata"`
}

func (*TransferTicketContentsAndResult) OperationContentsAndResult() {}
func (op *TransferTicketContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type DALPublishCommitmentContentsAndResult struct {
	DALPublishCommitment
	Metadata ManagerMetadata[DALPublishCommitmentResult] `json:"metadata"`
}

func (*DALPublishCommitmentContentsAndResult) OperationContentsAndResult() {}
func (op *DALPublishCommitmentContentsAndResult) GetMetadata() any {
	return &op.Metadata
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

type ManagerMetadata[T core.ManagerOperationResult] struct {
	BalanceUpdates
	OperationResult          T                         `json:"operation_result"`
	InternalOperationResults []InternalOperationResult `tz:"dyn" json:"internal_operation_results"`
}

type InternalOperationResult interface {
	core.InternalOperationResult
}

type OriginationInternalOperationResult = proto_023_PtSeouLo.OriginationInternalOperationResult
type DelegationInternalOperationResult = proto_023_PtSeouLo.DelegationInternalOperationResult
type EventInternalOperationResult = proto_023_PtSeouLo.EventInternalOperationResult

func init() {
	encoding.RegisterEnum(&encoding.Enum[InternalOperationResult]{
		Variants: encoding.Variants[InternalOperationResult]{
			1: (*TransactionInternalOperationResult)(nil),
			2: (*OriginationInternalOperationResult)(nil),
			3: (*DelegationInternalOperationResult)(nil),
			4: (*EventInternalOperationResult)(nil),
		},
	})
}

type RevealSuccessfulManagerResult = proto_023_PtSeouLo.RevealSuccessfulManagerResult
type DelegationSuccessfulManagerResult = proto_023_PtSeouLo.DelegationSuccessfulManagerResult
type OriginationSuccessfulManagerResult = proto_023_PtSeouLo.OriginationSuccessfulManagerResult
type IncreasePaidStorageSuccessfulManagerResult = proto_023_PtSeouLo.IncreasePaidStorageSuccessfulManagerResult
type SmartRollupOriginateSuccessfulManagerResult = proto_023_PtSeouLo.SmartRollupOriginateSuccessfulManagerResult
type SetDepositsLimitSuccessfulManagerResult = proto_023_PtSeouLo.SetDepositsLimitSuccessfulManagerResult
type UpdateConsensusKeySuccessfulManagerResult = proto_023_PtSeouLo.UpdateConsensusKeySuccessfulManagerResult

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

var ListOperations = core.ListOperations[OperationContents]
