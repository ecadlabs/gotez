package proto_015_PtLimaPt

//go:generate go run ../../cmd/genmarshaller.go

import (
	"math/big"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/proto_012_Psithaca"
	"github.com/ecadlabs/gotez/v2/protocol/proto_013_PtJakart"
	"github.com/ecadlabs/gotez/v2/protocol/proto_014_PtKathma"
)

type ManagerOperation = proto_012_Psithaca.ManagerOperation
type SeedNonceRevelation = proto_012_Psithaca.SeedNonceRevelation
type Preendorsement = proto_012_Psithaca.Preendorsement
type InlinedPreendorsement = proto_012_Psithaca.InlinedPreendorsement
type Endorsement = proto_012_Psithaca.Endorsement
type InlinedEndorsement = proto_012_Psithaca.InlinedEndorsement
type DoublePreendorsementEvidence = proto_012_Psithaca.DoublePreendorsementEvidence
type DoubleEndorsementEvidence = proto_012_Psithaca.DoubleEndorsementEvidence
type Reveal = proto_012_Psithaca.Reveal
type Delegation = proto_012_Psithaca.Delegation
type RegisterGlobalConstant = proto_012_Psithaca.RegisterGlobalConstant
type SetDepositsLimit = proto_012_Psithaca.SetDepositsLimit
type IncreasePaidStorage = proto_014_PtKathma.IncreasePaidStorage
type ActivateAccount = proto_012_Psithaca.ActivateAccount
type Proposals = proto_012_Psithaca.Proposals
type Ballot = proto_012_Psithaca.Ballot
type VDFRevelation = proto_014_PtKathma.VDFRevelation
type FailingNoop = proto_012_Psithaca.FailingNoop
type TransferTicket = proto_013_PtJakart.TransferTicket
type DALSlotAvailability = proto_014_PtKathma.DALSlotAvailability
type DALSlotAvailabilityContentsAndResult = proto_014_PtKathma.DALSlotAvailabilityContentsAndResult
type DoubleBakingEvidence = proto_013_PtJakart.DoubleBakingEvidence
type ConsumedGasResult = proto_014_PtKathma.ConsumedGasResult
type ConsumedGasResultContents = proto_014_PtKathma.ConsumedGasResultContents
type RevealSuccessfulManagerResult = proto_014_PtKathma.RevealSuccessfulManagerResult
type DelegationSuccessfulManagerResult = proto_014_PtKathma.DelegationSuccessfulManagerResult
type SetDepositsLimitSuccessfulManagerResult = proto_014_PtKathma.SetDepositsLimitSuccessfulManagerResult

type DelegationContentsAndResult struct {
	Delegation
	Metadata ManagerMetadata[ConsumedGasResult] `json:"metadata"`
}

func (*DelegationContentsAndResult) OperationContentsAndResult() {}
func (op *DelegationContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type RevealContentsAndResult struct {
	Reveal
	Metadata ManagerMetadata[ConsumedGasResult] `json:"metadata"`
}

func (*RevealContentsAndResult) OperationContentsAndResult() {}
func (op *RevealContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type SetDepositsLimitContentsAndResult struct {
	SetDepositsLimit
	Metadata ManagerMetadata[ConsumedGasResult] `json:"metadata"`
}

func (*SetDepositsLimitContentsAndResult) OperationContentsAndResult() {}
func (op *SetDepositsLimitContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

//json:kind=OperationKind()
type UpdateConsensusKey struct {
	ManagerOperation
	PublicKey tz.PublicKey `json:"public_key"`
}

func (*UpdateConsensusKey) OperationKind() string { return "update_consensus_key" }

//json:kind=OperationKind()
type DrainDelegate struct {
	ConsensusKey tz.PublicKeyHash `json:"consensus_key"`
	Delegate     tz.PublicKeyHash `json:"delegate"`
	Destination  tz.PublicKeyHash `json:"destination"`
}

func (*DrainDelegate) OperationKind() string { return "drain_delegate" }

//json:kind=OperationKind()
type UpdateConsensusKeySuccessfulManagerResult struct {
	core.OperationResultApplied[*ConsumedGasResultContents]
}

func (*UpdateConsensusKeySuccessfulManagerResult) OperationKind() string {
	return "update_consensus_key"
}

//json:kind=OperationKind()
type DALPublishSlotHeader struct {
	ManagerOperation
	Slot DALSlot `json:"slot"`
}

func (*DALPublishSlotHeader) OperationKind() string {
	return "dal_publish_slot_header"
}

type DALSlot struct {
	Level  int32     `json:"level"`
	Index  uint8     `json:"index"`
	Header *[48]byte `json:"header"`
}

type SeedNonceRevelationContentsAndResult struct {
	SeedNonceRevelation
	Metadata BalanceUpdates `json:"metadata"`
}

func (*SeedNonceRevelationContentsAndResult) OperationContentsAndResult() {}
func (op *SeedNonceRevelationContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type DoubleEndorsementEvidenceContentsAndResult struct {
	DoubleEndorsementEvidence
	Metadata BalanceUpdates `json:"metadata"`
}

func (*DoubleEndorsementEvidenceContentsAndResult) OperationContentsAndResult() {}
func (op *DoubleEndorsementEvidenceContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type DoubleBakingEvidenceContentsAndResult struct {
	DoubleBakingEvidence
	Metadata BalanceUpdates `json:"metadata"`
}

func (*DoubleBakingEvidenceContentsAndResult) OperationContentsAndResult() {}
func (op *DoubleBakingEvidenceContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type ActivateAccountContentsAndResult struct {
	ActivateAccount
	Metadata BalanceUpdates `json:"metadata"`
}

func (*ActivateAccountContentsAndResult) OperationContentsAndResult() {}
func (op *ActivateAccountContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type DoublePreendorsementEvidenceContentsAndResult struct {
	DoublePreendorsementEvidence
	Metadata BalanceUpdates `json:"metadata"`
}

func (*DoublePreendorsementEvidenceContentsAndResult) OperationContentsAndResult() {}
func (op *DoublePreendorsementEvidenceContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

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

type DrainDelegateContentsAndResult struct {
	DrainDelegate
	Metadata DrainDelegateMetadata `json:"metadata"`
}

func (*DrainDelegateContentsAndResult) OperationContentsAndResult() {}
func (op *DrainDelegateContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type EndorsementMetadata struct {
	BalanceUpdates
	Delegate         tz.PublicKeyHash `json:"delegate"`
	EndorsementPower int32            `json:"endorsement_power"`
	ConsensusKey     tz.PublicKeyHash `json:"consensus_key"`
}

type EndorsementContentsAndResult struct {
	Endorsement
	Metadata EndorsementMetadata `json:"metadata"`
}

func (*EndorsementContentsAndResult) OperationContentsAndResult() {}
func (op *EndorsementContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type PreendorsementMetadata = EndorsementMetadata
type PreendorsementContentsAndResult struct {
	Preendorsement
	Metadata PreendorsementMetadata `json:"metadata"`
}

func (*PreendorsementContentsAndResult) OperationContentsAndResult() {}
func (op *PreendorsementContentsAndResult) GetMetadata() any {
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

type RegisterGlobalConstantContentsAndResult struct {
	RegisterGlobalConstant
	Metadata ManagerMetadata[RegisterGlobalConstantResult] `json:"metadata"`
}

func (*RegisterGlobalConstantContentsAndResult) OperationContentsAndResult() {}
func (op *RegisterGlobalConstantContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type UpdateConsensusKeyContentsAndResult struct {
	UpdateConsensusKey
	Metadata ManagerMetadata[ConsumedGasResult] `json:"metadata"`
}

func (*UpdateConsensusKeyContentsAndResult) OperationContentsAndResult() {}
func (op *UpdateConsensusKeyContentsAndResult) GetMetadata() any {
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

type IncreasePaidStorageContentsAndResult struct {
	IncreasePaidStorage
	Metadata ManagerMetadata[IncreasePaidStorageResult] `json:"metadata"`
}

func (*IncreasePaidStorageContentsAndResult) OperationContentsAndResult() {}
func (op *IncreasePaidStorageContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

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

type DALPublishSlotHeaderContentsAndResult struct {
	DALPublishSlotHeader
	Metadata ManagerMetadata[ConsumedGasResult] `json:"metadata"`
}

func (*DALPublishSlotHeaderContentsAndResult) OperationContentsAndResult() {}
func (op *DALPublishSlotHeaderContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type OperationContents interface {
	core.OperationContents
}

func init() {
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
			22:  (*DALSlotAvailability)(nil),
			107: (*Reveal)(nil),
			108: (*Transaction)(nil),
			109: (*Origination)(nil),
			110: (*Delegation)(nil),
			111: (*RegisterGlobalConstant)(nil),
			112: (*SetDepositsLimit)(nil),
			113: (*IncreasePaidStorage)(nil),
			114: (*UpdateConsensusKey)(nil),
			// Never used and deprecated
			// 150 Tx_rollup_origination
			// 151 Tx_rollup_submit_batch
			// 152 Tx_rollup_commit
			// 153 Tx_rollup_return_bond
			// 154 Tx_rollup_finalize_commitment
			// 155 Tx_rollup_remove_commitment
			// 156 Tx_rollup_rejection
			// 157 Tx_rollup_dispatch_tickets
			158: (*TransferTicket)(nil),
			// Never used in this revision
			// 200: (*ScRollupOriginate)(nil),
			// 201: (*ScRollupAddMessages)(nil),
			// 202: (*ScRollupCement)(nil),
			// 203: (*ScRollupPublish)(nil),
			// 204: (*ScRollupRefute)(nil),
			// 205: (*ScRollupTimeout)(nil),
			// 206: (*ScRollupExecuteOutboxMessage)(nil),
			// 207: (*ScRollupRecoverBond)(nil),
			// 208: (*ScRollupDALSlotSubscribe)(nil),
			230: (*DALPublishSlotHeader)(nil),
			250: (*ZkRollupOrigination)(nil),
			251: (*ZkRollupPublish)(nil),
		},
	})
}

type OperationContentsAndResult interface {
	core.OperationContentsAndResult
}

func init() {
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
			22:  (*DALSlotAvailabilityContentsAndResult)(nil),
			107: (*RevealContentsAndResult)(nil),
			108: (*TransactionContentsAndResult)(nil),
			109: (*OriginationContentsAndResult)(nil),
			110: (*DelegationContentsAndResult)(nil),
			111: (*RegisterGlobalConstantContentsAndResult)(nil),
			112: (*SetDepositsLimitContentsAndResult)(nil),
			113: (*IncreasePaidStorageContentsAndResult)(nil),
			114: (*UpdateConsensusKeyContentsAndResult)(nil),
			158: (*TransferTicketContentsAndResult)(nil),
			230: (*DALPublishSlotHeaderContentsAndResult)(nil),
			250: (*ZkRollupOriginationContentsAndResult)(nil),
			251: (*ZkRollupPublishContentsAndResult)(nil),
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

type DelegationInternalOperationResult = proto_013_PtJakart.DelegationInternalOperationResult
type EventInternalOperationResult = proto_014_PtKathma.EventInternalOperationResult

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

type SuccessfulManagerOperationResult interface {
	core.SuccessfulManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SuccessfulManagerOperationResult]{
		Variants: encoding.Variants[SuccessfulManagerOperationResult]{
			0: (*RevealSuccessfulManagerResult)(nil),
			1: (*TransactionSuccessfulManagerResult)(nil),
			2: (*OriginationSuccessfulManagerResult)(nil),
			3: (*DelegationSuccessfulManagerResult)(nil),
			5: (*SetDepositsLimitSuccessfulManagerResult)(nil),
			6: (*UpdateConsensusKeySuccessfulManagerResult)(nil),
			9: (*IncreasePaidStorageSuccessfulManagerResult)(nil),
			// 200: (*ScRollupOriginateResultContents)(nil),
		},
	})
}

func ListOperations() []OperationContents {
	return encoding.ListVariants[OperationContents]()
}
