package proto_014_PtKathma

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/core/expression"
	"github.com/ecadlabs/gotez/protocol/proto_012_Psithaca"
	"github.com/ecadlabs/gotez/protocol/proto_013_PtJakart"
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
type ActivateAccount = proto_012_Psithaca.ActivateAccount
type Proposals = proto_012_Psithaca.Proposals
type Ballot = proto_012_Psithaca.Ballot
type FailingNoop = proto_012_Psithaca.FailingNoop
type LazyStorageDiff = proto_012_Psithaca.LazyStorageDiff
type TransferTicket = proto_013_PtJakart.TransferTicket
type Entrypoint = proto_012_Psithaca.Entrypoint
type DoubleBakingEvidence = proto_012_Psithaca.DoubleBakingEvidence

type EventResult interface {
	EventResult()
	core.OperationResult
}

type EventResultContents struct {
	ConsumedMilligas tz.BigUint
}

func (EventResultContents) SuccessfulManagerOperationResult() {}
func (EventResultContents) OperationKind() string             { return "event" }

type RevealResultContents EventResultContents

func (*RevealResultContents) SuccessfulManagerOperationResult() {}
func (*RevealResultContents) OperationKind() string             { return "reveal" }

type DelegationInternalOperationResult struct {
	Source   TransactionDestination
	Nonce    uint16
	Delegate tz.Option[tz.PublicKeyHash]
	Result   EventResult
}

func (*DelegationInternalOperationResult) InternalOperationResult() {}
func (*DelegationInternalOperationResult) OperationKind() string    { return "delegation" }

type DelegationResultContents EventResultContents

func (*DelegationResultContents) SuccessfulManagerOperationResult() {}
func (*DelegationResultContents) OperationKind() string             { return "delegation" }

type SetDepositsLimitResultContents EventResultContents

func (*SetDepositsLimitResultContents) SuccessfulManagerOperationResult() {}
func (*SetDepositsLimitResultContents) OperationKind() string {
	return "set_deposits_limit"
}

type EventResultApplied struct {
	core.OperationResultApplied[EventResultContents]
}

func (*EventResultApplied) EventResult() {}

type EventResultBacktracked struct {
	core.OperationResultBacktracked[EventResultContents]
}

func (*EventResultBacktracked) EventResult() {}

type EventResultFailed struct{ core.OperationResultFailed }

func (*EventResultFailed) EventResult() {}

type EventResultSkipped struct{ core.OperationResultSkipped }

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

type EventInternalOperationResult struct {
	Source  TransactionDestination
	Nonce   uint16
	Type    expression.Expression
	Tag     tz.Option[Entrypoint]
	Payload tz.Option[expression.Expression]
	Result  EventResult
}

func (*EventInternalOperationResult) InternalOperationResult() {}
func (*EventInternalOperationResult) OperationKind() string    { return "event" }

type IncreasePaidStorage struct {
	ManagerOperation
	Amount      tz.BigInt
	Destination core.OriginatedContractID
}

func (*IncreasePaidStorage) OperationKind() string { return "increase_paid_storage" }

type IncreasePaidStorageResult interface {
	IncreasePaidStorageResult()
	core.OperationResult
}

type IncreasePaidStorageResultContents struct {
	BalanceUpdates   []*BalanceUpdate `tz:"dyn"`
	ConsumedMilligas tz.BigUint
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
	Metadata ManagerMetadata[IncreasePaidStorageResult]
}

func (*IncreasePaidStorageContentsAndResult) OperationContentsAndResult() {}
func (op *IncreasePaidStorageContentsAndResult) OperationContents() core.OperationContents {
	return &op.IncreasePaidStorage
}

type VDFRevelation struct {
	Solution [2]*[100]byte
}

func (*VDFRevelation) OperationKind() string { return "vdf_revelation" }

type DALSlotAvailability struct {
	Endorser    tz.PublicKeyHash
	Endorsement tz.BigUint
}

func (*DALSlotAvailability) OperationKind() string { return "dal_slot_availability" }

type DALSlotAvailabilityContentsAndResult struct {
	DALSlotAvailability
	Metadata DALSlotAvailabilityMetadata
}

func (*DALSlotAvailabilityContentsAndResult) OperationContentsAndResult() {}
func (op *DALSlotAvailabilityContentsAndResult) OperationContents() core.OperationContents {
	return &op.DALSlotAvailability
}

type DALSlotAvailabilityMetadata struct {
	Delegate tz.PublicKeyHash
}

type DALPublishSlotHeader struct {
	ManagerOperation
	Slot DALSlot
}

func (*DALPublishSlotHeader) OperationKind() string {
	return "dal_publish_slot_header"
}

type DALSlot struct {
	Level  int32
	Index  uint8
	Header int32
}

type SeedNonceRevelationContentsAndResult struct {
	SeedNonceRevelation
	Metadata []*BalanceUpdate `tz:"dyn"`
}

func (*SeedNonceRevelationContentsAndResult) OperationContentsAndResult() {}
func (op *SeedNonceRevelationContentsAndResult) OperationContents() core.OperationContents {
	return &op.SeedNonceRevelation
}

type DoubleEndorsementEvidenceContentsAndResult struct {
	DoubleEndorsementEvidence
	Metadata []*BalanceUpdate `tz:"dyn"`
}

func (*DoubleEndorsementEvidenceContentsAndResult) OperationContentsAndResult() {}
func (op *DoubleEndorsementEvidenceContentsAndResult) OperationContents() core.OperationContents {
	return &op.DoubleEndorsementEvidence
}

type DoubleBakingEvidenceContentsAndResult struct {
	DoubleBakingEvidence
	Metadata []*BalanceUpdate `tz:"dyn"`
}

func (*DoubleBakingEvidenceContentsAndResult) OperationContentsAndResult() {}
func (op *DoubleBakingEvidenceContentsAndResult) OperationContents() core.OperationContents {
	return &op.DoubleBakingEvidence
}

type ActivateAccountContentsAndResult struct {
	ActivateAccount
	Metadata []*BalanceUpdate `tz:"dyn"`
}

func (*ActivateAccountContentsAndResult) OperationContentsAndResult() {}
func (op *ActivateAccountContentsAndResult) OperationContents() core.OperationContents {
	return &op.ActivateAccount
}

type DoublePreendorsementEvidenceContentsAndResult struct {
	DoublePreendorsementEvidence
	Metadata []*BalanceUpdate `tz:"dyn"`
}

func (*DoublePreendorsementEvidenceContentsAndResult) OperationContentsAndResult() {}
func (op *DoublePreendorsementEvidenceContentsAndResult) OperationContents() core.OperationContents {
	return &op.DoublePreendorsementEvidence
}

type VDFRevelationContentsAndResult struct {
	VDFRevelation
	Metadata []*BalanceUpdate `tz:"dyn"`
}

func (*VDFRevelationContentsAndResult) OperationContentsAndResult() {}
func (op *VDFRevelationContentsAndResult) OperationContents() core.OperationContents {
	return &op.VDFRevelation
}

type EndorsementMetadata struct {
	BalanceUpdates   []*BalanceUpdate `tz:"dyn"`
	Delegate         tz.PublicKeyHash
	EndorsementPower int32
}

type EndorsementContentsAndResult struct {
	Endorsement
	Metadata EndorsementMetadata
}

func (*EndorsementContentsAndResult) OperationContentsAndResult() {}
func (op *EndorsementContentsAndResult) OperationContents() core.OperationContents {
	return &op.Endorsement
}

type PreendorsementMetadata = EndorsementMetadata
type PreendorsementContentsAndResult struct {
	Preendorsement
	Metadata PreendorsementMetadata
}

func (*PreendorsementContentsAndResult) OperationContentsAndResult() {}
func (op *PreendorsementContentsAndResult) OperationContents() core.OperationContents {
	return &op.Preendorsement
}

type RevealContentsAndResult struct {
	Reveal
	Metadata ManagerMetadata[EventResult]
}

func (*RevealContentsAndResult) OperationContentsAndResult() {}
func (op *RevealContentsAndResult) OperationContents() core.OperationContents {
	return &op.Reveal
}

type DelegationContentsAndResult struct {
	Delegation
	Metadata ManagerMetadata[EventResult]
}

func (*DelegationContentsAndResult) OperationContentsAndResult() {}
func (op *DelegationContentsAndResult) OperationContents() core.OperationContents {
	return &op.Delegation
}

type RegisterGlobalConstantResult interface {
	proto_012_Psithaca.RegisterGlobalConstantResult
}

type RegisterGlobalConstantResultContents struct {
	BalanceUpdates   []*BalanceUpdate `tz:"dyn"`
	ConsumedMilligas tz.BigUint
	StorageSize      tz.BigInt
	GlobalAddress    *tz.ScriptExprHash
}

func (RegisterGlobalConstantResultContents) SuccessfulManagerOperationResult() {}
func (RegisterGlobalConstantResultContents) OperationKind() string {
	return "register_global_constant"
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
	Metadata ManagerMetadata[RegisterGlobalConstantResult]
}

func (*RegisterGlobalConstantContentsAndResult) OperationContentsAndResult() {}
func (op *RegisterGlobalConstantContentsAndResult) OperationContents() core.OperationContents {
	return &op.RegisterGlobalConstant
}

type SetDepositsLimitContentsAndResult struct {
	SetDepositsLimit
	Metadata ManagerMetadata[EventResult]
}

func (*SetDepositsLimitContentsAndResult) OperationContentsAndResult() {}
func (op *SetDepositsLimitContentsAndResult) OperationContents() core.OperationContents {
	return &op.SetDepositsLimit
}

type DALPublishSlotHeaderContentsAndResult struct {
	DALPublishSlotHeader
	Metadata ManagerMetadata[EventResult]
}

func (*DALPublishSlotHeaderContentsAndResult) OperationContentsAndResult() {}
func (op *DALPublishSlotHeaderContentsAndResult) OperationContents() core.OperationContents {
	return &op.DALPublishSlotHeader
}

type TransferTicketContentsAndResult struct {
	TransferTicket
	Metadata ManagerMetadata[TransferTicketResult]
}

func (*TransferTicketContentsAndResult) OperationContentsAndResult() {}
func (op *TransferTicketContentsAndResult) OperationContents() core.OperationContents {
	return &op.TransferTicket
}

type TransferTicketResultContents struct {
	BalanceUpdates      []*BalanceUpdate `tz:"dyn"`
	ConsumedMilligas    tz.BigUint
	PaidStorageSizeDiff tz.BigInt
}

func (TransferTicketResultContents) SuccessfulManagerOperationResult() {}
func (TransferTicketResultContents) OperationKind() string {
	return "transfer_ticket"
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
			158: (*TransferTicketContentsAndResult)(nil),
			230: (*DALPublishSlotHeaderContentsAndResult)(nil),
		},
	})
}

type ManagerMetadata[T core.OperationResult] struct {
	BalanceUpdates           []*BalanceUpdate `tz:"dyn"`
	OperationResult          T
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
			9: (*IncreasePaidStorageResultContents)(nil),
			// 200: (*ScRollupOriginateResultContents)(nil),
		},
	})
}
