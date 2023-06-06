package proto_015_PtLimaPt

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/core/expression"
	"github.com/ecadlabs/gotez/protocol/proto_012_Psithaca"
	"github.com/ecadlabs/gotez/protocol/proto_013_PtJakart"
	"github.com/ecadlabs/gotez/protocol/proto_014_PtKathma"
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
type Entrypoint = proto_012_Psithaca.Entrypoint
type DoubleBakingEvidence = proto_012_Psithaca.DoubleBakingEvidence
type ConsumedGasResult = proto_014_PtKathma.ConsumedGasResult
type ConsumedGasResultContents = proto_014_PtKathma.ConsumedGasResultContents
type RevealResultContents = proto_014_PtKathma.RevealResultContents
type DelegationResultContents = proto_014_PtKathma.DelegationResultContents
type SetDepositsLimitResultContents = proto_014_PtKathma.SetDepositsLimitResultContents

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

type UpdateConsensusKey struct {
	ManagerOperation
	PublicKey tz.PublicKey `json:"public_key"`
}

func (*UpdateConsensusKey) OperationKind() string { return "update_consensus_key" }

type DrainDelegate struct {
	ConsensusKey tz.PublicKeyHash `json:"consensus_key"`
	Delegate     tz.PublicKeyHash `json:"delegate"`
	Destination  tz.PublicKeyHash `json:"destination"`
}

func (*DrainDelegate) OperationKind() string { return "drain_delegate" }

type UpdateConsensusKeyResultContents ConsumedGasResultContents

func (*UpdateConsensusKeyResultContents) SuccessfulManagerOperationResult() {}
func (*UpdateConsensusKeyResultContents) OperationKind() string {
	return "update_consensus_key"
}

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

type RegisterGlobalConstantResult interface {
	proto_012_Psithaca.RegisterGlobalConstantResult
}

type RegisterGlobalConstantResultContents struct {
	BalanceUpdates
	ConsumedMilligas tz.BigUint         `json:"consumed_milligas"`
	StorageSize      tz.BigInt          `json:"storage_size"`
	GlobalAddress    *tz.ScriptExprHash `json:"global_address"`
}

type RegisterGlobalConstantResultApplied struct {
	core.OperationResultApplied[RegisterGlobalConstantResultContents]
}

func (*RegisterGlobalConstantResultApplied) RegisterGlobalConstantResult() {}

type RegisterGlobalConstantResultBacktracked struct {
	core.OperationResultBacktracked[RegisterGlobalConstantResultContents]
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

type IncreasePaidStorageResult interface {
	proto_014_PtKathma.IncreasePaidStorageResult
}

type IncreasePaidStorageResultContents struct {
	BalanceUpdates
	ConsumedMilligas tz.BigUint `json:"consumed_milligas"`
}

func (IncreasePaidStorageResultContents) SuccessfulManagerOperationResult() {}
func (IncreasePaidStorageResultContents) OperationKind() string {
	return "increase_paid_storage"
}

type IncreasePaidStorageResultApplied struct {
	core.OperationResultApplied[IncreasePaidStorageResultContents]
}

func (*IncreasePaidStorageResultApplied) IncreasePaidStorageResult() {}

type IncreasePaidStorageResultBacktracked struct {
	core.OperationResultBacktracked[IncreasePaidStorageResultContents]
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

type TransferTicketResult interface {
	proto_013_PtJakart.TransferTicketResult
}

type TransferTicketResultApplied struct {
	core.OperationResultApplied[TransferTicketResultContents]
}

func (*TransferTicketResultApplied) TransferTicketResult() {}

type TransferTicketResultBacktracked struct {
	core.OperationResultBacktracked[TransferTicketResultContents]
}

func (*TransferTicketResultBacktracked) TransferTicketResult() {}

type TransferTicketResultFailed struct{ core.OperationResultFailed }

func (*TransferTicketResultFailed) TransferTicketResult() {}

type TransferTicketResultSkipped struct{ core.OperationResultSkipped }

func (*TransferTicketResultSkipped) TransferTicketResult() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[TransferTicketResult]{
		Variants: encoding.Variants[TransferTicketResult]{
			0: (*TransferTicketResultApplied)(nil),
			1: (*TransferTicketResultFailed)(nil),
			2: (*TransferTicketResultSkipped)(nil),
			3: (*TransferTicketResultBacktracked)(nil),
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

type DelegationInternalOperationResult struct {
	Source   TransactionDestination      `json:"source"`
	Nonce    uint16                      `json:"nonce"`
	Delegate tz.Option[tz.PublicKeyHash] `json:"delegate"`
	Result   ConsumedGasResult           `json:"result"`
}

func (r *DelegationInternalOperationResult) InternalOperationResult() core.ManagerOperationResult {
	return r.Result
}
func (*DelegationInternalOperationResult) OperationKind() string { return "delegation" }

type EventInternalOperationResult struct {
	Source  TransactionDestination           `json:"source"`
	Nonce   uint16                           `json:"nonce"`
	Type    expression.Expression            `json:"type"`
	Tag     tz.Option[Entrypoint]            `json:"tag"`
	Payload tz.Option[expression.Expression] `json:"payload"`
	Result  ConsumedGasResult                `json:"result"`
}

func (r *EventInternalOperationResult) InternalOperationResult() core.ManagerOperationResult {
	return r.Result
}
func (*EventInternalOperationResult) OperationKind() string { return "event" }

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
			0: (*RevealResultContents)(nil),
			1: (*TransactionResultContents)(nil),
			2: (*OriginationResultContents)(nil),
			3: (*DelegationResultContents)(nil),
			5: (*SetDepositsLimitResultContents)(nil),
			6: (*UpdateConsensusKeyResultContents)(nil),
			9: (*IncreasePaidStorageResultContents)(nil),
			// 200: (*ScRollupOriginateResultContents)(nil),
		},
	})
}

func ListOperations() []OperationContents {
	return encoding.ListVariants[OperationContents]()
}
