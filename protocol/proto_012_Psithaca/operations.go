package proto_012_Psithaca

//go:generate go run ../../cmd/genmarshaller.go

import (
	"math/big"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/core/expression"
)

//json:kind=OperationKind()
type ActivateAccount struct {
	PKH    *tz.Ed25519PublicKeyHash `json:"pkh"`
	Secret *tz.Bytes20              `json:"secret"`
}

func (*ActivateAccount) OperationKind() string { return "activate_account" }

//json:kind=OperationKind()
type Proposals struct {
	Source    tz.PublicKeyHash   `json:"source"`
	Period    int32              `json:"period"`
	Proposals []*tz.ProtocolHash `tz:"dyn" json:"proposals"`
}

var _ core.OperationWithSource = (*Proposals)(nil)

func (*Proposals) OperationKind() string { return "proposals" }
func (p *Proposals) GetSource() core.TransactionDestination {
	return core.ImplicitContract{PublicKeyHash: p.Source}
}
func (*Proposals) OperationContentsAndResult() {}
func (op *Proposals) GetMetadata() any         { return op }

//json:kind=OperationKind()
type Ballot struct {
	Source   tz.PublicKeyHash `json:"source"`
	Period   int32            `json:"period"`
	Proposal *tz.ProtocolHash `json:"proposal"`
	Ballot   core.BallotKind  `json:"ballot"`
}

var _ core.OperationWithSource = (*Ballot)(nil)

func (op *Ballot) GetSource() core.TransactionDestination {
	return core.ImplicitContract{PublicKeyHash: op.Source}
}
func (*Ballot) OperationKind() string       { return "ballot" }
func (*Ballot) OperationContentsAndResult() {}
func (op *Ballot) GetMetadata() any         { return op }

type ManagerOperation struct {
	Source       tz.PublicKeyHash `json:"source"`
	Fee          tz.BigUint       `json:"fee"`
	Counter      tz.BigUint       `json:"counter"`
	GasLimit     tz.BigUint       `json:"gas_limit"`
	StorageLimit tz.BigUint       `json:"storage_limit"`
}

func (m *ManagerOperation) GetSource() core.TransactionDestination {
	return core.ImplicitContract{PublicKeyHash: m.Source}
}
func (m *ManagerOperation) GetSourceAddress() tz.PublicKeyHash { return m.Source }
func (m *ManagerOperation) GetFee() tz.BigUint                 { return m.Fee }
func (m *ManagerOperation) GetCounter() tz.BigUint             { return m.Counter }
func (m *ManagerOperation) GetGasLimit() tz.BigUint            { return m.GasLimit }
func (m *ManagerOperation) GetStorageLimit() tz.BigUint        { return m.StorageLimit }

func (m *ManagerOperation) SetFee(v tz.BigUint)          { m.Fee = v }
func (m *ManagerOperation) SetCounter(v tz.BigUint)      { m.Counter = v }
func (m *ManagerOperation) SetGasLimit(v tz.BigUint)     { m.GasLimit = v }
func (m *ManagerOperation) SetStorageLimit(v tz.BigUint) { m.StorageLimit = v }

type Script struct {
	Code    expression.Expression `tz:"dyn" json:"code"`
	Storage expression.Expression `tz:"dyn" json:"storage"`
}

//json:kind=OperationKind()
type Delegation struct {
	ManagerOperation
	Delegate tz.Option[tz.PublicKeyHash] `json:"delegate"`
}

func (*Delegation) OperationKind() string { return "delegation" }

//json:kind=OperationKind()
type Reveal struct {
	ManagerOperation
	PublicKey tz.PublicKey `json:"public_key"`
}

func (*Reveal) OperationKind() string { return "reveal" }

//json:kind=OperationKind()
type SeedNonceRevelation struct {
	Level int32         `json:"level"`
	Nonce *tz.SeedNonce `json:"nonce"`
}

func (*SeedNonceRevelation) OperationKind() string { return "seed_nonce_revelation" }

//json:kind=OperationKind()
type FailingNoop struct {
	Arbitrary tz.Bytes `tz:"dyn" json:"arbitrary"`
}

func (*FailingNoop) OperationKind() string { return "failing_noop" }

//json:kind=OperationKind()
type RegisterGlobalConstant struct {
	ManagerOperation
	Value expression.Expression `tz:"dyn" json:"value"`
}

func (*RegisterGlobalConstant) OperationKind() string { return "register_global_constant" }

//json:kind=OperationKind()
type SetDepositsLimit struct {
	ManagerOperation
	Limit tz.Option[tz.BigUint] `json:"limit"`
}

func (*SetDepositsLimit) OperationKind() string { return "set_deposits_limit" }

//json:kind=OperationKind()
type Endorsement struct {
	Slot             uint16               `json:"slot"`
	Level            int32                `json:"level"`
	Round            int32                `json:"round"`
	BlockPayloadHash *tz.BlockPayloadHash `json:"block_payload_hash"`
}

func (*Endorsement) InlinedEndorsementContents() {}
func (*Endorsement) OperationKind() string       { return "endorsement" }

type InlinedEndorsement struct {
	Branch    *tz.BlockHash              `json:"branch"`
	Contents  InlinedEndorsementContents `json:"contents"`
	Signature tz.AnySignature            `json:"signature"`
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

//json:kind=OperationKind()
type DoubleEndorsementEvidence struct {
	Op1 InlinedEndorsement `tz:"dyn" json:"op1"`
	Op2 InlinedEndorsement `tz:"dyn" json:"op2"`
}

func (*DoubleEndorsementEvidence) OperationKind() string { return "double_endorsement_evidence" }

//json:kind=OperationKind()
type DoublePreendorsementEvidence struct {
	Op1 InlinedPreendorsement `tz:"dyn" json:"op1"`
	Op2 InlinedPreendorsement `tz:"dyn" json:"op2"`
}

func (*DoublePreendorsementEvidence) OperationKind() string { return "double_preendorsement_evidence" }

type InlinedPreendorsement struct {
	Branch    *tz.BlockHash                 `json:"branch"`
	Contents  InlinedPreendorsementContents `json:"contents"`
	Signature *tz.GenericSignature          `json:"signature"`
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

//json:kind=OperationKind()
type Preendorsement Endorsement

func (*Preendorsement) InlinedPreendorsementContents() {}
func (*Preendorsement) OperationKind() string          { return "preendorsement" }

//json:kind=OperationKind()
type DoubleBakingEvidence struct {
	Block1 BlockHeader `tz:"dyn" json:"block1"`
	Block2 BlockHeader `tz:"dyn" json:"block2"`
}

func (*DoubleBakingEvidence) OperationKind() string { return "double_baking_evidence" }

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
		},
	})
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

type EndorsementMetadata struct {
	BalanceUpdates
	Delegate         tz.PublicKeyHash `json:"delegate"`
	EndorsementPower int32            `json:"endorsement_power"`
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

type OperationContentsAndResult interface {
	core.OperationContentsAndResult
}

type ConsumedGasResultContents struct {
	ConsumedGas      tz.BigUint `json:"consumed_gas"`
	ConsumedMilligas tz.BigUint `json:"consumed_milligas"`
}

func (r *ConsumedGasResultContents) GetConsumedMilligas() tz.BigUint { return r.ConsumedMilligas }

type ConsumedGasResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[ConsumedGasResult]{
		Variants: encoding.Variants[ConsumedGasResult]{
			0: (*core.OperationResultApplied[*ConsumedGasResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*ConsumedGasResultContents])(nil),
		},
	})
}

type RevealResultContents = ConsumedGasResultContents

//json:kind=OperationKind()
type RevealSuccessfulManagerResult struct {
	core.OperationResultApplied[*ConsumedGasResultContents]
}

func (*RevealSuccessfulManagerResult) OperationKind() string { return "reveal" }

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

type DelegationResultContents = ConsumedGasResultContents

//json:kind=OperationKind()
type DelegationSuccessfulManagerResult struct {
	core.OperationResultApplied[*ConsumedGasResultContents]
}

func (*DelegationSuccessfulManagerResult) OperationKind() string { return "delegation" }

type SetDepositsLimitResultContents = ConsumedGasResultContents

//json:kind=OperationKind()
type SetDepositsLimitSuccessfulManagerResult struct {
	core.OperationResultApplied[*ConsumedGasResultContents]
}

func (*SetDepositsLimitSuccessfulManagerResult) OperationKind() string { return "set_deposits_limit" }

type RegisterGlobalConstantResultContents struct {
	BalanceUpdates
	ConsumedGas   tz.BigUint         `json:"consumed_gas"`
	StorageSize   tz.BigInt          `json:"storage_size"`
	GlobalAddress *tz.ScriptExprHash `json:"global_address"`
}

func (r *RegisterGlobalConstantResultContents) GetConsumedMilligas() tz.BigUint {
	x := r.ConsumedGas.Int()
	x.Mul(x, big.NewInt(1000))
	v, _ := tz.NewBigUint(x)
	return v
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
	Source   core.ContractID             `json:"source"`
	Nonce    uint16                      `json:"nonce"`
	Delegate tz.Option[tz.PublicKeyHash] `json:"delegate"`
	Result   ConsumedGasResult           `json:"result"`
}

var _ core.InternalOperationResult = (*DelegationInternalOperationResult)(nil)

func (r *DelegationInternalOperationResult) GetSource() core.TransactionDestination {
	switch d := r.Source.(type) {
	case core.ImplicitContract:
		return d
	case core.OriginatedContract:
		return d
	default:
		panic("unexpected contract id type")
	}
}
func (r *DelegationInternalOperationResult) GetResult() core.ManagerOperationResult {
	return r.Result
}
func (*DelegationInternalOperationResult) OperationKind() string { return "delegation" }

//json:kind=OperationKind()
type RevealInternalOperationResult struct {
	Source    core.ContractID   `json:"source"`
	Nonce     uint16            `json:"nonce"`
	PublicKey tz.PublicKey      `json:"public_key"`
	Result    ConsumedGasResult `json:"result"`
}

var _ core.InternalOperationResult = (*RevealInternalOperationResult)(nil)

func (r *RevealInternalOperationResult) GetSource() core.TransactionDestination {
	switch d := r.Source.(type) {
	case core.ImplicitContract:
		return d
	case core.OriginatedContract:
		return d
	default:
		panic("unexpected contract id type")
	}
}
func (r *RevealInternalOperationResult) GetResult() core.ManagerOperationResult {
	return r.Result
}
func (*RevealInternalOperationResult) OperationKind() string { return "reveal" }

//json:kind=OperationKind()
type RegisterGlobalConstantInternalOperationResult struct {
	Source core.ContractID              `json:"source"`
	Nonce  uint16                       `json:"nonce"`
	Value  expression.Expression        `tz:"dyn" json:"value"`
	Result RegisterGlobalConstantResult `json:"result"`
}

var _ core.InternalOperationResult = (*RegisterGlobalConstantInternalOperationResult)(nil)

func (r *RegisterGlobalConstantInternalOperationResult) GetSource() core.TransactionDestination {
	switch d := r.Source.(type) {
	case core.ImplicitContract:
		return d
	case core.OriginatedContract:
		return d
	default:
		panic("unexpected contract id type")
	}
}

func (*RegisterGlobalConstantInternalOperationResult) OperationKind() string {
	return "register_global_constant"
}
func (r *RegisterGlobalConstantInternalOperationResult) GetResult() core.ManagerOperationResult {
	return r.Result
}

//json:kind=OperationKind()
type SetDepositsLimitInternalOperationResult struct {
	Source core.ContractID       `json:"source"`
	Nonce  uint16                `json:"nonce"`
	Limit  tz.Option[tz.BigUint] `json:"limit"`
	Result ConsumedGasResult     `json:"result"`
}

var _ core.InternalOperationResult = (*SetDepositsLimitInternalOperationResult)(nil)

func (r *SetDepositsLimitInternalOperationResult) GetSource() core.TransactionDestination {
	switch d := r.Source.(type) {
	case core.ImplicitContract:
		return d
	case core.OriginatedContract:
		return d
	default:
		panic("unexpected contract id type")
	}
}
func (*SetDepositsLimitInternalOperationResult) OperationKind() string {
	return "set_deposits_limit"
}
func (r *SetDepositsLimitInternalOperationResult) GetResult() core.ManagerOperationResult {
	return r.Result
}

type InternalOperationResult interface {
	core.InternalOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[InternalOperationResult]{
		Variants: encoding.Variants[InternalOperationResult]{
			0: (*RevealInternalOperationResult)(nil),
			1: (*TransactionInternalOperationResult)(nil),
			2: (*OriginationInternalOperationResult)(nil),
			3: (*DelegationInternalOperationResult)(nil),
			4: (*RegisterGlobalConstantInternalOperationResult)(nil),
			5: (*SetDepositsLimitInternalOperationResult)(nil),
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
		},
	})
}

func ListOperations() []OperationContents {
	return encoding.ListVariants[OperationContents]()
}
