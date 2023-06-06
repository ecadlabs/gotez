package proto_012_Psithaca

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/core/expression"
)

type ActivateAccount struct {
	PKH    *tz.Ed25519PublicKeyHash `json:"pkh"`
	Secret *[tz.SecretBytesLen]byte `json:"secret"`
}

func (*ActivateAccount) OperationKind() string        { return "activate_account" }
func (op *ActivateAccount) Operation() core.Operation { return op }

type Proposals struct {
	Source    tz.PublicKeyHash   `json:"source"`
	Period    int32              `json:"period"`
	Proposals []*tz.ProtocolHash `tz:"dyn" json:"proposals"`
}

func (*Proposals) OperationKind() string       { return "proposals" }
func (*Proposals) OperationContentsAndResult() {}
func (op *Proposals) Operation() core.Operation {
	return op
}

type Ballot struct {
	Source   tz.PublicKeyHash `json:"source"`
	Period   int32            `json:"period"`
	Proposal *tz.ProtocolHash `json:"proposal"`
	Ballot   core.BallotKind  `json:"ballot"`
}

func (*Ballot) OperationKind() string       { return "ballot" }
func (*Ballot) OperationContentsAndResult() {}
func (op *Ballot) Operation() core.Operation {
	return op
}

type ManagerOperation struct {
	Source       tz.PublicKeyHash `json:"source"`
	Fee          tz.BigUint       `json:"fee"`
	Counter      tz.BigUint       `json:"counter"`
	GasLimit     tz.BigUint       `json:"gas_limit"`
	StorageLimit tz.BigUint       `json:"storage_limit"`
}

func (m *ManagerOperation) GetSource() tz.PublicKeyHash { return m.Source }
func (m *ManagerOperation) GetFee() tz.BigUint          { return m.Fee }
func (m *ManagerOperation) GetCounter() tz.BigUint      { return m.Counter }
func (m *ManagerOperation) GetGasLimit() tz.BigUint     { return m.GasLimit }
func (m *ManagerOperation) GetStorageLimit() tz.BigUint { return m.StorageLimit }

type Script struct {
	Code    expression.Expression `tz:"dyn" json:"code"`
	Storage expression.Expression `tz:"dyn" json:"storage"`
}

type Delegation struct {
	ManagerOperation
	Delegate tz.Option[tz.PublicKeyHash] `json:"delegate"`
}

func (*Delegation) OperationKind() string        { return "delegation" }
func (op *Delegation) Operation() core.Operation { return op }

type Reveal struct {
	ManagerOperation
	PublicKey tz.PublicKey `json:"public_key"`
}

func (*Reveal) OperationKind() string        { return "reveal" }
func (op *Reveal) Operation() core.Operation { return op }

type SeedNonceRevelation struct {
	Level int32         `json:"level"`
	Nonce *tz.SeedNonce `json:"nonce"`
}

func (*SeedNonceRevelation) OperationKind() string        { return "seed_nonce_revelation" }
func (op *SeedNonceRevelation) Operation() core.Operation { return op }

type FailingNoop struct {
	Arbitrary tz.Bytes `tz:"dyn" json:"arbitrary"`
}

func (*FailingNoop) OperationKind() string        { return "failing_noop" }
func (op *FailingNoop) Operation() core.Operation { return op }

type RegisterGlobalConstant struct {
	ManagerOperation
	Value expression.Expression `tz:"dyn" json:"value"`
}

func (*RegisterGlobalConstant) OperationKind() string        { return "register_global_constant" }
func (op *RegisterGlobalConstant) Operation() core.Operation { return op }

type SetDepositsLimit struct {
	ManagerOperation
	Limit tz.Option[tz.BigUint] `json:"limit"`
}

func (*SetDepositsLimit) OperationKind() string        { return "set_deposits_limit" }
func (op *SetDepositsLimit) Operation() core.Operation { return op }

type Endorsement struct {
	Slot             uint16               `json:"slot"`
	Level            int32                `json:"level"`
	Round            int32                `json:"round"`
	BlockPayloadHash *tz.BlockPayloadHash `json:"block_payload_hash"`
}

func (*Endorsement) InlinedEndorsementContents()  {}
func (*Endorsement) OperationKind() string        { return "endorsement" }
func (op *Endorsement) Operation() core.Operation { return op }

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

type DoubleEndorsementEvidence struct {
	Op1 InlinedEndorsement `tz:"dyn" json:"op1"`
	Op2 InlinedEndorsement `tz:"dyn" json:"op2"`
}

func (*DoubleEndorsementEvidence) OperationKind() string        { return "double_endorsement_evidence" }
func (op *DoubleBakingEvidence) Operation() core.Operation      { return op }
func (op *DoubleEndorsementEvidence) Operation() core.Operation { return op }

type DoublePreendorsementEvidence struct {
	Op1 InlinedPreendorsement `tz:"dyn" json:"op1"`
	Op2 InlinedPreendorsement `tz:"dyn" json:"op2"`
}

func (*DoublePreendorsementEvidence) OperationKind() string        { return "double_preendorsement_evidence" }
func (op *DoublePreendorsementEvidence) Operation() core.Operation { return op }

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

type Preendorsement struct {
	Slot             uint16               `json:"slot"`
	Level            int32                `json:"level"`
	Round            int32                `json:"round"`
	BlockPayloadHash *tz.BlockPayloadHash `json:"block_payload_hash"`
}

func (*Preendorsement) InlinedPreendorsementContents() {}
func (*Preendorsement) OperationKind() string          { return "preendorsement" }
func (op *Preendorsement) Operation() core.Operation   { return op }

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
	Metadata []*BalanceUpdate `tz:"dyn" json:"metadata"`
}

func (*SeedNonceRevelationContentsAndResult) OperationContentsAndResult() {}
func (op *SeedNonceRevelationContentsAndResult) Operation() core.Operation {
	return &op.SeedNonceRevelation
}

type DoubleEndorsementEvidenceContentsAndResult struct {
	DoubleEndorsementEvidence
	Metadata []*BalanceUpdate `tz:"dyn" json:"metadata"`
}

func (*DoubleEndorsementEvidenceContentsAndResult) OperationContentsAndResult() {}
func (op *DoubleEndorsementEvidenceContentsAndResult) Operation() core.Operation {
	return &op.DoubleEndorsementEvidence
}

type DoubleBakingEvidenceContentsAndResult struct {
	DoubleBakingEvidence
	Metadata []*BalanceUpdate `tz:"dyn" json:"metadata"`
}

func (*DoubleBakingEvidenceContentsAndResult) OperationContentsAndResult() {}
func (op *DoubleBakingEvidenceContentsAndResult) Operation() core.Operation {
	return &op.DoubleBakingEvidence
}

type ActivateAccountContentsAndResult struct {
	ActivateAccount
	Metadata []*BalanceUpdate `tz:"dyn" json:"metadata"`
}

func (*ActivateAccountContentsAndResult) OperationContentsAndResult() {}
func (op *ActivateAccountContentsAndResult) Operation() core.Operation {
	return &op.ActivateAccount
}

type DoublePreendorsementEvidenceContentsAndResult struct {
	DoublePreendorsementEvidence
	Metadata []*BalanceUpdate `tz:"dyn" json:"metadata"`
}

func (*DoublePreendorsementEvidenceContentsAndResult) OperationContentsAndResult() {}
func (op *DoublePreendorsementEvidenceContentsAndResult) Operation() core.Operation {
	return &op.DoublePreendorsementEvidence
}

type EndorsementMetadata struct {
	BalanceUpdates   []*BalanceUpdate `tz:"dyn" json:"balance_updates"`
	Delegate         tz.PublicKeyHash `json:"delegate"`
	EndorsementPower int32            `json:"endorsement_power"`
}

type EndorsementContentsAndResult struct {
	Endorsement
	Metadata EndorsementMetadata `json:"metadata"`
}

func (*EndorsementContentsAndResult) OperationContentsAndResult() {}
func (op *EndorsementContentsAndResult) Operation() core.Operation {
	return &op.Endorsement
}

type PreendorsementMetadata = EndorsementMetadata
type PreendorsementContentsAndResult struct {
	Preendorsement
	Metadata PreendorsementMetadata `json:"metadata"`
}

func (*PreendorsementContentsAndResult) OperationContentsAndResult() {}
func (op *PreendorsementContentsAndResult) Operation() core.Operation {
	return &op.Preendorsement
}

type OperationContentsAndResult interface {
	core.OperationContentsAndResult
}

type ConsumedGasResultContents struct {
	ConsumedGas      tz.BigUint `json:"consumed_gas"`
	ConsumedMilligas tz.BigUint `json:"consumed_milligas"`
}
type ConsumedGasResult interface {
	ConsumedGasResult()
	core.ManagerOperationResult
}

type ConsumedGasResultApplied struct {
	core.OperationResultApplied[ConsumedGasResultContents]
}

func (*ConsumedGasResultApplied) ConsumedGasResult() {}

type ConsumedGasResultBacktracked struct {
	core.OperationResultBacktracked[ConsumedGasResultContents]
}

func (*ConsumedGasResultBacktracked) ConsumedGasResult() {}

type ConsumedGasResultFailed struct{ core.OperationResultFailed }

func (*ConsumedGasResultFailed) ConsumedGasResult() {}

type ConsumedGasResultSkipped struct{ core.OperationResultSkipped }

func (*ConsumedGasResultSkipped) ConsumedGasResult() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[ConsumedGasResult]{
		Variants: encoding.Variants[ConsumedGasResult]{
			0: (*ConsumedGasResultApplied)(nil),
			1: (*ConsumedGasResultFailed)(nil),
			2: (*ConsumedGasResultSkipped)(nil),
			3: (*ConsumedGasResultBacktracked)(nil),
		},
	})
}

type RevealResultContents ConsumedGasResultContents

func (*RevealResultContents) SuccessfulManagerOperationResult() {}
func (*RevealResultContents) OperationKind() string             { return "reveal" }

type RevealContentsAndResult struct {
	Reveal
	Metadata ManagerMetadata[ConsumedGasResult] `json:"metadata"`
}

func (*RevealContentsAndResult) OperationContentsAndResult() {}
func (op *RevealContentsAndResult) Operation() core.Operation {
	return &op.Reveal
}

type DelegationContentsAndResult struct {
	Delegation
	Metadata ManagerMetadata[ConsumedGasResult] `json:"metadata"`
}

func (*DelegationContentsAndResult) OperationContentsAndResult() {}
func (op *DelegationContentsAndResult) Operation() core.Operation {
	return &op.Delegation
}

type DelegationResultContents ConsumedGasResultContents

func (*DelegationResultContents) SuccessfulManagerOperationResult() {}
func (*DelegationResultContents) OperationKind() string             { return "delegation" }

type SetDepositsLimitResultContents ConsumedGasResultContents

func (*SetDepositsLimitResultContents) SuccessfulManagerOperationResult() {}
func (*SetDepositsLimitResultContents) OperationKind() string {
	return "set_deposits_limit"
}

type RegisterGlobalConstantResultContents struct {
	BalanceUpdates []*BalanceUpdate   `tz:"dyn" json:"balance_updates"`
	ConsumedGas    tz.BigUint         `json:"consumed_gas"`
	StorageSize    tz.BigInt          `json:"storage_size"`
	GlobalAddress  *tz.ScriptExprHash `json:"global_address"`
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

type RegisterGlobalConstantResult interface {
	RegisterGlobalConstantResult()
	core.ManagerOperationResult
}

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
func (op *RegisterGlobalConstantContentsAndResult) Operation() core.Operation {
	return &op.RegisterGlobalConstant
}

type SetDepositsLimitContentsAndResult struct {
	SetDepositsLimit
	Metadata ManagerMetadata[ConsumedGasResult] `json:"metadata"`
}

func (*SetDepositsLimitContentsAndResult) OperationContentsAndResult() {}
func (op *SetDepositsLimitContentsAndResult) Operation() core.Operation {
	return &op.SetDepositsLimit
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
	BalanceUpdates           []*BalanceUpdate          `tz:"dyn" json:"balance_updates"`
	OperationResult          T                         `json:"operation_result"`
	InternalOperationResults []InternalOperationResult `tz:"dyn" json:"internal_operation_results"`
}

type DelegationInternalOperationResult struct {
	Source   core.ContractID             `json:"source"`
	Nonce    uint16                      `json:"nonce"`
	Delegate tz.Option[tz.PublicKeyHash] `json:"delegate"`
	Result   ConsumedGasResult           `json:"result"`
}

func (r *DelegationInternalOperationResult) InternalOperationResult() core.ManagerOperationResult {
	return r.Result
}
func (*DelegationInternalOperationResult) OperationKind() string { return "delegation" }

type RevealInternalOperationResult struct {
	Source    core.ContractID   `json:"source"`
	Nonce     uint16            `json:"nonce"`
	PublicKey tz.PublicKey      `json:"public_key"`
	Result    ConsumedGasResult `json:"result"`
}

func (r *RevealInternalOperationResult) InternalOperationResult() core.ManagerOperationResult {
	return r.Result
}
func (*RevealInternalOperationResult) OperationKind() string { return "reveal" }

type RegisterGlobalConstantInternalOperationResult struct {
	Source core.ContractID              `json:"source"`
	Nonce  uint16                       `json:"nonce"`
	Value  expression.Expression        `tz:"dyn" json:"value"`
	Result RegisterGlobalConstantResult `json:"result"`
}

func (r *RegisterGlobalConstantInternalOperationResult) InternalOperationResult() core.ManagerOperationResult {
	return r.Result
}
func (*RegisterGlobalConstantInternalOperationResult) OperationKind() string {
	return "register_global_constant"
}

type SetDepositsLimitInternalOperationResult struct {
	Source core.ContractID       `json:"source"`
	Nonce  uint16                `json:"nonce"`
	Limit  tz.Option[tz.BigUint] `json:"limit"`
	Result ConsumedGasResult     `json:"result"`
}

func (r *SetDepositsLimitInternalOperationResult) InternalOperationResult() core.ManagerOperationResult {
	return r.Result
}
func (*SetDepositsLimitInternalOperationResult) OperationKind() string {
	return "set_deposits_limit"
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
			0: (*RevealResultContents)(nil),
			1: (*TransactionResultContents)(nil),
			2: (*OriginationResultContents)(nil),
			3: (*DelegationResultContents)(nil),
			5: (*SetDepositsLimitResultContents)(nil),
		},
	})
}

func ListOperations() []OperationContents {
	return encoding.ListVariants[OperationContents]()
}
