package proto_013_PtJakart

//go:generate go run ../../cmd/genmarshaller.go

import (
	"math/big"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/core/expression"
	"github.com/ecadlabs/gotez/v2/protocol/proto_012_Psithaca"
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
type RevealRevealSuccessfulManagerResult = proto_012_Psithaca.RevealSuccessfulManagerResult
type DelegationSuccessfulManagerResult = proto_012_Psithaca.DelegationSuccessfulManagerResult
type SetDepositsLimitSuccessfulManagerResult = proto_012_Psithaca.SetDepositsLimitSuccessfulManagerResult

type Ticket struct {
	Contents expression.Expression `tz:"dyn" json:"contents"`
	Type     expression.Expression `tz:"dyn" json:"type"`
	Ticketer core.ContractID       `json:"ticketer"`
	Amount   tz.BigUint            `json:"amount"`
}

//json:kind=OperationKind()
type TransferTicket struct {
	ManagerOperation
	Ticket      Ticket          `json:"ticket"`
	Destination core.ContractID `json:"destination"`
	Entrypoint  string          `tz:"dyn" json:"entrypoint"`
}

func (*TransferTicket) OperationKind() string { return "transfer_ticket" }

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

type EndorsementMetadata struct {
	BalanceUpdates
	Delegate         tz.PublicKeyHash `json:"delegate"`
	EndorsementPower int32            ` json:"endorsement_power"`
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

type RevealContentsAndResult struct {
	Reveal
	Metadata ManagerMetadata[ConsumedGasResult] `json:"metadata"`
}

func (*RevealContentsAndResult) OperationContentsAndResult() {}
func (op *RevealContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type DelegationContentsAndResult struct {
	Delegation
	Metadata ManagerMetadata[ConsumedGasResult] `json:"metadata"`
}

func (*DelegationContentsAndResult) OperationContentsAndResult() {}
func (op *DelegationContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type RegisterGlobalConstantResultContents struct {
	BalanceUpdates
	ConsumedGas      tz.BigUint         `json:"consumed_gas"`
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

type SetDepositsLimitContentsAndResult struct {
	SetDepositsLimit
	Metadata ManagerMetadata[ConsumedGasResult] `json:"metadata"`
}

func (*SetDepositsLimitContentsAndResult) OperationContentsAndResult() {}
func (op *SetDepositsLimitContentsAndResult) GetMetadata() any {
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
	ConsumedGas         tz.BigUint `json:"consumed_gas"`
	ConsumedMilligas    tz.BigUint `json:"consumed_milligas"`
	PaidStorageSizeDiff tz.BigInt  `json:"paid_storage_size_diff"`
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
type DelegationInternalOperationResult struct {
	Source   TransactionDestination      `json:"source"`
	Nonce    uint16                      `json:"nonce"`
	Delegate tz.Option[tz.PublicKeyHash] `json:"delegate"`
	Result   ConsumedGasResult           `json:"result"`
}

func (*DelegationInternalOperationResult) OperationKind() string     { return "delegation" }
func (r *DelegationInternalOperationResult) GetSource() core.Address { return r.Source }
func (r *DelegationInternalOperationResult) InternalOperationResult() core.ManagerOperationResult {
	return r.Result
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
		},
	})
}

type SuccessfulManagerOperationResult interface {
	core.SuccessfulManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SuccessfulManagerOperationResult]{
		Variants: encoding.Variants[SuccessfulManagerOperationResult]{
			0: (*RevealRevealSuccessfulManagerResult)(nil),
			1: (*TransactionSuccessfulManagerResult)(nil),
			2: (*OriginationSuccessfulManagerResult)(nil),
			3: (*DelegationSuccessfulManagerResult)(nil),
			5: (*SetDepositsLimitSuccessfulManagerResult)(nil),
			// 200: (*ScRollupOriginateResultContents)(nil),
		},
	})
}

func ListOperations() []OperationContents {
	return encoding.ListVariants[OperationContents]()
}
