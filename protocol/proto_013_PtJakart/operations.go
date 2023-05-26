package proto_013_PtJakart

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/core/expression"
	"github.com/ecadlabs/gotez/protocol/proto_012_Psithaca"
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
type Entrypoint = proto_012_Psithaca.Entrypoint
type DoubleBakingEvidence = proto_012_Psithaca.DoubleBakingEvidence
type ConsumedGasResult = proto_012_Psithaca.ConsumedGasResult
type ConsumedGasResultContents = proto_012_Psithaca.ConsumedGasResultContents
type RevealResultContents = proto_012_Psithaca.RevealResultContents
type DelegationResultContents = proto_012_Psithaca.DelegationResultContents
type SetDepositsLimitResultContents = proto_012_Psithaca.SetDepositsLimitResultContents

type TransferTicket struct {
	ManagerOperation
	TicketContents expression.Expression `tz:"dyn"`
	TicketType     expression.Expression `tz:"dyn"`
	TicketTicketer core.ContractID
	TicketAmount   tz.BigUint
	Destination    core.ContractID
	Entrypoint     string `tz:"dyn"`
}

func (*TransferTicket) OperationKind() string { return "transfer_ticket" }

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
	Metadata ManagerMetadata[ConsumedGasResult]
}

func (*RevealContentsAndResult) OperationContentsAndResult() {}
func (op *RevealContentsAndResult) OperationContents() core.OperationContents {
	return &op.Reveal
}

type DelegationContentsAndResult struct {
	Delegation
	Metadata ManagerMetadata[ConsumedGasResult]
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
	ConsumedGas      tz.BigUint
	ConsumedMilligas tz.BigUint
	StorageSize      tz.BigInt
	GlobalAddress    *tz.ScriptExprHash
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
	Metadata ManagerMetadata[ConsumedGasResult]
}

func (*SetDepositsLimitContentsAndResult) OperationContentsAndResult() {}
func (op *SetDepositsLimitContentsAndResult) OperationContents() core.OperationContents {
	return &op.SetDepositsLimit
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
	ConsumedGas         tz.BigUint
	ConsumedMilligas    tz.BigUint
	PaidStorageSizeDiff tz.BigInt
}

func (TransferTicketResultContents) SuccessfulManagerOperationResult() {}
func (TransferTicketResultContents) OperationKind() string {
	return "transfer_ticket"
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

type TransferTicketResult interface {
	TransferTicketResult()
	core.ManagerOperationResult
}

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
			17:  (*FailingNoop)(nil),
			20:  (*Preendorsement)(nil),
			21:  (*Endorsement)(nil),
			107: (*Reveal)(nil),
			108: (*Transaction)(nil),
			109: (*Origination)(nil),
			110: (*Delegation)(nil),
			111: (*RegisterGlobalConstant)(nil),
			112: (*SetDepositsLimit)(nil),
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
			20:  (*PreendorsementContentsAndResult)(nil),
			21:  (*EndorsementContentsAndResult)(nil),
			107: (*RevealContentsAndResult)(nil),
			108: (*TransactionContentsAndResult)(nil),
			109: (*OriginationContentsAndResult)(nil),
			110: (*DelegationContentsAndResult)(nil),
			111: (*RegisterGlobalConstantContentsAndResult)(nil),
			112: (*SetDepositsLimitContentsAndResult)(nil),
			158: (*TransferTicketContentsAndResult)(nil),
		},
	})
}

type ManagerMetadata[T core.ManagerOperationResult] struct {
	BalanceUpdates           []*BalanceUpdate `tz:"dyn"`
	OperationResult          T
	InternalOperationResults []InternalOperationResult `tz:"dyn"`
}
type DelegationInternalOperationResult struct {
	Source   TransactionDestination
	Nonce    uint16
	Delegate tz.Option[tz.PublicKeyHash]
	Result   ConsumedGasResult
}

func (*DelegationInternalOperationResult) InternalOperationResult() {}
func (*DelegationInternalOperationResult) OperationKind() string    { return "delegation" }

type InternalOperationResult interface {
	core.InternalOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[InternalOperationResult]{
		Variants: encoding.Variants[InternalOperationResult]{
			1: (*TransactionInternalOperationResult)(nil),
			2: (*OriginationInternalOperationResult)(nil),
			3: (*DelegationInternalOperationResult)(nil),
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
			// 200: (*ScRollupOriginateResultContents)(nil),
		},
	})
}
