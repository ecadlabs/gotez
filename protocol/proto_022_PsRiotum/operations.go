package proto_022_PsRiotum

//go:generate go run ../../cmd/genmarshaller.go

import (
	"math/big"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/proto_012_Psithaca/lazy"
	"github.com/ecadlabs/gotez/v2/protocol/proto_021_PsQuebec"
)

type ManagerOperation = proto_021_PsQuebec.ManagerOperation
type SeedNonceRevelation = proto_021_PsQuebec.SeedNonceRevelation
type DoubleAttestationEvidence = proto_021_PsQuebec.DoubleAttestationEvidence
type DoubleBakingEvidence = proto_021_PsQuebec.DoubleBakingEvidence
type ActivateAccount = proto_021_PsQuebec.ActivateAccount
type Proposals = proto_021_PsQuebec.Proposals
type Ballot = proto_021_PsQuebec.Ballot
type DoublePreattestationEvidence = proto_021_PsQuebec.DoublePreattestationEvidence
type VDFRevelation = proto_021_PsQuebec.VDFRevelation
type DrainDelegate = proto_021_PsQuebec.DrainDelegate
type FailingNoop = proto_021_PsQuebec.FailingNoop
type Preattestation = proto_021_PsQuebec.Preattestation
type InlinedPreattestation = proto_021_PsQuebec.InlinedPreattestation
type InlinedPreattestationContent = proto_021_PsQuebec.InlinedPreattestationContent
type Attestation = proto_021_PsQuebec.Attestation
type AttestationWithDAL = proto_021_PsQuebec.AttestationWithDAL
type InlinedAttestation = proto_021_PsQuebec.InlinedAttestation
type InlinedAttestationContent = proto_021_PsQuebec.InlinedAttestationContent
type Reveal = proto_021_PsQuebec.Reveal
type Origination = proto_021_PsQuebec.Origination
type Delegation = proto_021_PsQuebec.Delegation
type RegisterGlobalConstant = proto_021_PsQuebec.RegisterGlobalConstant
type IncreasePaidStorage = proto_021_PsQuebec.IncreasePaidStorage
type SetDepositsLimit = proto_021_PsQuebec.SetDepositsLimit
type TransferTicket = proto_021_PsQuebec.TransferTicket
type Script = proto_021_PsQuebec.Script
type ConsumedGasResult = proto_021_PsQuebec.ConsumedGasResult
type DALPublishCommitment = proto_021_PsQuebec.DALPublishCommitment
type DALPublishCommitmentResult = proto_021_PsQuebec.DALPublishCommitmentResult

type OperationContents interface {
	core.OperationContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[OperationContents]{
		Variants: encoding.Variants[OperationContents]{
			1:   (*SeedNonceRevelation)(nil),
			2:   (*DoubleAttestationEvidence)(nil),
			3:   (*DoubleBakingEvidence)(nil),
			4:   (*ActivateAccount)(nil),
			5:   (*Proposals)(nil),
			6:   (*Ballot)(nil),
			7:   (*DoublePreattestationEvidence)(nil),
			8:   (*VDFRevelation)(nil),
			9:   (*DrainDelegate)(nil),
			17:  (*FailingNoop)(nil),
			20:  (*Preattestation)(nil),
			21:  (*Attestation)(nil),
			23:  (*AttestationWithDAL)(nil),
			24:  (*DALEntrapmentEvidence)(nil),
			31:  (*AttestationsAggregate)(nil),
			107: (*Reveal)(nil),
			108: (*Transaction)(nil),
			109: (*Origination)(nil),
			110: (*Delegation)(nil),
			111: (*RegisterGlobalConstant)(nil),
			112: (*SetDepositsLimit)(nil),
			113: (*IncreasePaidStorage)(nil),
			114: (*UpdateConsensusKey)(nil),
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
type UpdateConsensusKey struct {
	ManagerOperation
	PublicKey tz.PublicKey     `json:"public_key"`
	Proof     tz.Option[Proof] `json:"proof"`
}

func (*UpdateConsensusKey) OperationKind() string { return "update_consensus_key" }

type Proof struct {
	Signature tz.AnySignature `tz:"dyn" json:"signature"`
}

//json:kind=OperationKind()
type DALEntrapmentEvidence struct {
	Attestation    InlinedAttestation `tz:"dyn" json:"attestation"`
	SlotIndex      uint8              `json:"slot_index"`
	ShardWithProof ShardWithProof     `json:"shard_with_proof"`
}

func (*DALEntrapmentEvidence) OperationKind() string { return "dal_entrapment_evidence" }

type ShardWithProof struct {
	Shard Shard            `json:"shard"`
	Proof tz.DALCommitment `json:"proof"`
}

type Shard struct {
	Field0 int32    `json:"field_0"`
	Field1 tz.Bytes `tz:"dyn" json:"field_1"`
}

//json:kind=OperationKind()
type AttestationsAggregate struct {
	ConsensusContent ConsensusContent `json:"consensus_content"`
	Committee        []uint16         `tz:"dyn" json:"committee"`
}

func (*AttestationsAggregate) OperationKind() string { return "attestations_aggregate" }

type ConsensusContent struct {
	Level            int32                `json:"level"`
	Round            int32                `json:"round"`
	BlockPayloadHash *tz.BlockPayloadHash `json:"block_payload_hash"`
}

type OperationContentsAndResult interface {
	core.OperationContentsAndResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[OperationContentsAndResult]{
		Variants: encoding.Variants[OperationContentsAndResult]{
			1:   (*SeedNonceRevelationContentsAndResult)(nil),
			2:   (*DoubleAttestationEvidenceContentsAndResult)(nil),
			3:   (*DoubleBakingEvidenceContentsAndResult)(nil),
			4:   (*ActivateAccountContentsAndResult)(nil),
			5:   (*Proposals)(nil),
			6:   (*Ballot)(nil),
			7:   (*DoublePreattestationEvidenceContentsAndResult)(nil),
			8:   (*VDFRevelationContentsAndResult)(nil),
			9:   (*DrainDelegateContentsAndResult)(nil),
			20:  (*PreattestationContentsAndResult)(nil),
			21:  (*AttestationContentsAndResult)(nil),
			23:  (*AttestationWithDALContentsAndResult)(nil),
			24:  (*DALEntrapmentEvidenceContentsAndResult)(nil),
			31:  (*AttestationsAggregateContentsAndResult)(nil),
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

func (m *ManagerMetadata[T]) GetResult() core.ManagerOperationResult {
	return m.OperationResult
}
func (m *ManagerMetadata[T]) GetInternalOperationResults() []core.InternalOperationResult {
	out := make([]core.InternalOperationResult, len(m.InternalOperationResults))
	for i, r := range m.InternalOperationResults {
		out[i] = r
	}
	return out
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

type Committee struct {
	Delegate     tz.PublicKeyHash `json:"delegate"`
	ConsensusPKH tz.PublicKeyHash `json:"consensus_pkh"`
}

type AttestationsAggregateMetadata struct {
	BalanceUpdates
	Committee      []Committee `tz:"dyn" json:"committee"`
	ConsensusPower int32       `json:"consensus_power"`
}

//json:kind=OperationKind()
type SeedNonceRevelationContentsAndResult struct {
	SeedNonceRevelation
	Metadata BalanceUpdates `json:"metadata"`
}

func (*SeedNonceRevelationContentsAndResult) OperationContentsAndResult() {}
func (op *SeedNonceRevelationContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type DoubleAttestationEvidenceMetadata struct {
	ForbiddenDelegate tz.Option[tz.PublicKeyHash] `json:"forbidden_delegate"`
	BalanceUpdates
}

//json:kind=OperationKind()
type DoubleAttestationEvidenceContentsAndResult struct {
	DoubleAttestationEvidence
	Metadata DoubleAttestationEvidenceMetadata `json:"metadata"`
}

func (*DoubleAttestationEvidenceContentsAndResult) OperationContentsAndResult() {}
func (op *DoubleAttestationEvidenceContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type DoubleBakingEvidenceContentsAndResult struct {
	DoubleBakingEvidence
	Metadata DoubleAttestationEvidenceMetadata `json:"metadata"`
}

func (*DoubleBakingEvidenceContentsAndResult) OperationContentsAndResult() {}
func (op *DoubleBakingEvidenceContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type ActivateAccountContentsAndResult struct {
	ActivateAccount
	Metadata BalanceUpdates `json:"metadata"`
}

func (*ActivateAccountContentsAndResult) OperationContentsAndResult() {}
func (op *ActivateAccountContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type DoublePreattestationEvidenceContentsAndResult struct {
	DoublePreattestationEvidence
	Metadata DoubleAttestationEvidenceMetadata `json:"metadata"`
}

func (*DoublePreattestationEvidenceContentsAndResult) OperationContentsAndResult() {}
func (op *DoublePreattestationEvidenceContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type VDFRevelationContentsAndResult struct {
	VDFRevelation
	Metadata BalanceUpdates `json:"metadata"`
}

func (*VDFRevelationContentsAndResult) OperationContentsAndResult() {}
func (op *VDFRevelationContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type DrainDelegateMetadata struct {
	BalanceUpdates
	AllocatedDestinationContract bool `json:"allocated_destination_contract"`
}

//json:kind=OperationKind()
type DrainDelegateContentsAndResult struct {
	DrainDelegate
	Metadata DrainDelegateMetadata `json:"metadata"`
}

func (*DrainDelegateContentsAndResult) OperationContentsAndResult() {}
func (op *DrainDelegateContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type AttestationMetadata struct {
	BalanceUpdates
	Delegate       tz.PublicKeyHash `json:"delegate"`
	ConsensusPower int32            `json:"consensus_power"`
	ConsensusKey   tz.PublicKeyHash `json:"consensus_key"`
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
type RevealContentsAndResult struct {
	Reveal
	Metadata ManagerMetadata[ConsumedGasResult] `json:"metadata"`
}

func (*RevealContentsAndResult) OperationContentsAndResult() {}
func (op *RevealContentsAndResult) GetMetadata() any {
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
type UpdateConsensusKeyContentsAndResult struct {
	UpdateConsensusKey
	Metadata ManagerMetadata[ConsumedGasResult] `json:"metadata"`
}

func (*UpdateConsensusKeyContentsAndResult) OperationContentsAndResult() {}
func (op *UpdateConsensusKeyContentsAndResult) GetMetadata() any {
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
type DALPublishCommitmentContentsAndResult struct {
	DALPublishCommitment
	Metadata ManagerMetadata[DALPublishCommitmentResult] `json:"metadata"`
}

func (*DALPublishCommitmentContentsAndResult) OperationContentsAndResult() {}
func (op *DALPublishCommitmentContentsAndResult) GetMetadata() any {
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

type RegisterGlobalConstantResultContents struct {
	BalanceUpdates
	ConsumedMilligas tz.BigUint         `json:"consumed_milligas"`
	StorageSize      tz.BigInt          `json:"storage_size"`
	GlobalAddress    *tz.ScriptExprHash `json:"global_address"`
}

func (r *RegisterGlobalConstantResultContents) GetConsumedMilligas() tz.BigUint {
	return r.ConsumedMilligas
}
func (r *RegisterGlobalConstantResultContents) GetStorageSize() tz.BigInt { return r.StorageSize }
func (r *RegisterGlobalConstantResultContents) EstimateStorageSize(constants core.Constants) *big.Int {
	return r.StorageSize.Int()
}

type RegisterGlobalConstantResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[RegisterGlobalConstantResult]{
		Variants: encoding.Variants[RegisterGlobalConstantResult]{
			0: (*core.OperationResultApplied[*RegisterGlobalConstantResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*RegisterGlobalConstantResultContents])(nil),
		},
	})
}

type OriginationResultContents struct {
	BalanceUpdates
	OriginatedContracts []core.OriginatedContractID `tz:"dyn" json:"originated_contracts"`
	ConsumedMilligas    tz.BigUint                  `json:"consumed_milligas"`
	StorageSize         tz.BigInt                   `json:"storage_size"`
	PaidStorageSizeDiff tz.BigInt                   `json:"paid_storage_size_diff"`
	LazyStorageDiff     tz.Option[lazy.StorageDiff] `json:"lazy_storage_diff"`
}

func (r *OriginationResultContents) GetConsumedMilligas() tz.BigUint   { return r.ConsumedMilligas }
func (r *OriginationResultContents) GetStorageSize() tz.BigInt         { return r.StorageSize }
func (r *OriginationResultContents) GetPaidStorageSizeDiff() tz.BigInt { return r.PaidStorageSizeDiff }
func (r *OriginationResultContents) EstimateStorageSize(constants core.Constants) *big.Int {
	x := r.PaidStorageSizeDiff.Int()
	x.Add(x, big.NewInt(int64(constants.GetOriginationSize())))
	return x
}

//json:kind=OperationKind()
type OriginationSuccessfulManagerResult struct {
	core.OperationResultApplied[*OriginationResultContents]
}

func (*OriginationSuccessfulManagerResult) OperationKind() string { return "origination" }

type OriginationResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[OriginationResult]{
		Variants: encoding.Variants[OriginationResult]{
			0: (*core.OperationResultApplied[*OriginationResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*OriginationResultContents])(nil),
		},
	})
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

type IncreasePaidStorageResultContents struct {
	BalanceUpdates
	ConsumedMilligas tz.BigUint `json:"consumed_milligas"`
}

func (r *IncreasePaidStorageResultContents) GetConsumedMilligas() tz.BigUint {
	return r.ConsumedMilligas
}

//json:kind=OperationKind()
type IncreasePaidStorageSuccessfulManagerResult struct {
	core.OperationResultApplied[*IncreasePaidStorageResultContents]
}

func (*IncreasePaidStorageSuccessfulManagerResult) OperationKind() string {
	return "increase_paid_storage"
}

type IncreasePaidStorageResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[IncreasePaidStorageResult]{
		Variants: encoding.Variants[IncreasePaidStorageResult]{
			0: (*core.OperationResultApplied[*IncreasePaidStorageResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*IncreasePaidStorageResultContents])(nil),
		},
	})
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
type TransferTicketContentsAndResult struct {
	TransferTicket
	Metadata ManagerMetadata[TransferTicketResult] `json:"metadata"`
}

func (*TransferTicketContentsAndResult) OperationContentsAndResult() {}
func (op *TransferTicketContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type TransferTicketResultContents struct {
	BalanceUpdates
	TicketUpdates       []*TicketReceipt `tz:"dyn" json:"ticket_updates"`
	ConsumedMilligas    tz.BigUint       `json:"consumed_milligas"`
	PaidStorageSizeDiff tz.BigInt        `json:"paid_storage_size_diff"`
}

func (r *TransferTicketResultContents) GetConsumedMilligas() tz.BigUint { return r.ConsumedMilligas }
func (r *TransferTicketResultContents) GetPaidStorageSizeDiff() tz.BigInt {
	return r.PaidStorageSizeDiff
}
func (r *TransferTicketResultContents) EstimateStorageSize(constants core.Constants) *big.Int {
	return r.PaidStorageSizeDiff.Int()
}

type TransferTicketResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[TransferTicketResult]{
		Variants: encoding.Variants[TransferTicketResult]{
			0: (*core.OperationResultApplied[*TransferTicketResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*TransferTicketResultContents])(nil),
		},
	})
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
}

//json:kind=OperationKind()
type OriginationInternalOperationResult struct {
	Source   core.TransactionDestination `json:"source"`
	Nonce    uint16                      `json:"nonce"`
	Balance  tz.BigUint                  `json:"balance"`
	Delegate tz.Option[tz.PublicKeyHash] `json:"delegate"`
	Script   Script                      `json:"script"`
	Result   OriginationResult           `json:"result"`
}

func (r *OriginationInternalOperationResult) GetSource() core.TransactionDestination { return r.Source }

func (r *OriginationInternalOperationResult) GetResult() core.ManagerOperationResult {
	return r.Result
}
func (*OriginationInternalOperationResult) OperationKind() string { return "origination" }

type DelegationResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[DelegationResult]{
		Variants: encoding.Variants[DelegationResult]{
			0: (*core.OperationResultApplied[*DelegationResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*DelegationResultContents])(nil),
		},
	})
}

type DelegationResultContents struct {
	ConsumedMilligas tz.BigUint `json:"consumed_milligas"`
	BalanceUpdates
}

//json:kind=OperationKind()
type DelegationInternalOperationResult struct {
	Source   core.TransactionDestination `json:"source"`
	Nonce    uint16                      `json:"nonce"`
	Delegate tz.Option[tz.PublicKeyHash] `json:"delegate"`
	Result   DelegationResult            `json:"result"`
}

func (r *DelegationInternalOperationResult) GetSource() core.TransactionDestination { return r.Source }
func (*DelegationInternalOperationResult) OperationKind() string                    { return "delegation" }
func (r *DelegationInternalOperationResult) GetResult() core.ManagerOperationResult {
	return r.Result
}

type EventInternalOperationResult = proto_021_PsQuebec.EventInternalOperationResult

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

type RevealSuccessfulManagerResult = proto_021_PsQuebec.RevealSuccessfulManagerResult
type DelegationSuccessfulManagerResult = proto_021_PsQuebec.DelegationSuccessfulManagerResult
type UpdateConsensusKeySuccessfulManagerResult = proto_021_PsQuebec.UpdateConsensusKeySuccessfulManagerResult
type SetDepositsLimitSuccessfulManagerResult = proto_021_PsQuebec.SetDepositsLimitSuccessfulManagerResult

func ListOperations() []OperationContents {
	return encoding.ListVariants[OperationContents]()
}
