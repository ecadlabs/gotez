package proto_012_Psithaca

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/core/expression"
)

type ActivateAccount struct {
	PKH    *tz.Ed25519PublicKeyHash
	Secret *[tz.SecretBytesLen]byte
}

func (*ActivateAccount) OperationKind() string { return "activate_account" }

type Proposals struct {
	Source    tz.PublicKeyHash
	Period    int32
	Proposals []*tz.ProtocolHash `tz:"dyn"`
}

func (*Proposals) OperationKind() string       { return "proposals" }
func (*Proposals) OperationContentsAndResult() {}
func (op *Proposals) OperationContents() core.OperationContents {
	return op
}

type Ballot struct {
	Source   tz.PublicKeyHash
	Period   int32
	Proposal *tz.ProtocolHash
	Ballot   core.BallotKind
}

func (*Ballot) OperationKind() string       { return "ballot" }
func (*Ballot) OperationContentsAndResult() {}
func (op *Ballot) OperationContents() core.OperationContents {
	return op
}

type ManagerOperation struct {
	Source       tz.PublicKeyHash
	Fee          tz.BigUint
	Counter      tz.BigUint
	GasLimit     tz.BigUint
	StorageLimit tz.BigUint
}

type Script struct {
	Code    expression.Expression `tz:"dyn"`
	Storage expression.Expression `tz:"dyn"`
}

type Delegation struct {
	ManagerOperation
	Delegate tz.Option[tz.PublicKeyHash]
}

func (*Delegation) OperationKind() string { return "delegation" }

type Reveal struct {
	ManagerOperation
	PublicKey tz.PublicKey
}

func (*Reveal) OperationKind() string { return "reveal" }

type SeedNonceRevelation struct {
	Level int32
	Nonce *[tz.SeedNonceBytesLen]byte
}

func (*SeedNonceRevelation) OperationKind() string { return "seed_nonce_revelation" }

type FailingNoop struct {
	Arbitrary []byte `tz:"dyn"`
}

func (*FailingNoop) OperationKind() string { return "failing_noop" }

type RegisterGlobalConstant struct {
	ManagerOperation
	Value expression.Expression `tz:"dyn"`
}

func (*RegisterGlobalConstant) OperationKind() string { return "register_global_constant" }

type SetDepositsLimit struct {
	ManagerOperation
	Limit tz.Option[tz.BigUint]
}

func (*SetDepositsLimit) OperationKind() string { return "set_deposits_limit" }

type Endorsement struct {
	Slot             uint16
	Level            int32
	Round            int32
	BlockPayloadHash *tz.BlockPayloadHash
}

func (*Endorsement) InlinedEndorsementContents() {}
func (*Endorsement) OperationKind() string       { return "endorsement" }

type InlinedEndorsement struct {
	Branch    *tz.BlockHash
	Contents  InlinedEndorsementContents
	Signature tz.AnySignature
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
	Op1 InlinedEndorsement `tz:"dyn"`
	Op2 InlinedEndorsement `tz:"dyn"`
}

func (*DoubleEndorsementEvidence) OperationKind() string { return "double_endorsement_evidence" }

type DoublePreendorsementEvidence struct {
	Op1 InlinedPreendorsement `tz:"dyn"`
	Op2 InlinedPreendorsement `tz:"dyn"`
}

func (*DoublePreendorsementEvidence) OperationKind() string { return "double_preendorsement_evidence" }

type InlinedPreendorsement struct {
	Branch    *tz.BlockHash
	Contents  InlinedPreendorsementContents
	Signature *tz.GenericSignature
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
	Slot             uint16
	Level            int32
	Round            int32
	BlockPayloadHash *tz.BlockPayloadHash
}

func (*Preendorsement) InlinedPreendorsementContents() {}
func (*Preendorsement) OperationKind() string          { return "preendorsement" }

type DoubleBakingEvidence struct {
	Block1 BlockHeader `tz:"dyn"`
	Block2 BlockHeader `tz:"dyn"`
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

type OperationContentsAndResult interface {
	core.OperationContentsAndResult
}

type ConsumedGasResultContents struct {
	ConsumedGas      tz.BigUint
	ConsumedMilligas tz.BigUint
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

type DelegationResultContents ConsumedGasResultContents

func (*DelegationResultContents) SuccessfulManagerOperationResult() {}
func (*DelegationResultContents) OperationKind() string             { return "delegation" }

type SetDepositsLimitResultContents ConsumedGasResultContents

func (*SetDepositsLimitResultContents) SuccessfulManagerOperationResult() {}
func (*SetDepositsLimitResultContents) OperationKind() string {
	return "set_deposits_limit"
}

type RegisterGlobalConstantResultContents struct {
	BalanceUpdates []*BalanceUpdate `tz:"dyn"`
	ConsumedGas    tz.BigUint
	StorageSize    tz.BigInt
	GlobalAddress  *tz.ScriptExprHash
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
	BalanceUpdates           []*BalanceUpdate `tz:"dyn"`
	OperationResult          T
	InternalOperationResults []InternalOperationResult `tz:"dyn"`
}

type DelegationInternalOperationResult struct {
	Source   core.ContractID
	Nonce    uint16
	Delegate tz.Option[tz.PublicKeyHash]
	Result   ConsumedGasResult
}

func (*DelegationInternalOperationResult) InternalOperationResult() {}
func (*DelegationInternalOperationResult) OperationKind() string    { return "delegation" }

type RevealInternalOperationResult struct {
	Source    core.ContractID
	Nonce     uint16
	PublicKey tz.PublicKey
	Result    ConsumedGasResult
}

func (*RevealInternalOperationResult) InternalOperationResult() {}
func (*RevealInternalOperationResult) OperationKind() string    { return "reveal" }

type RegisterGlobalConstantInternalOperationResult struct {
	Source core.ContractID
	Nonce  uint16
	Value  expression.Expression `tz:"dyn"`
	Result RegisterGlobalConstantResult
}

func (*RegisterGlobalConstantInternalOperationResult) InternalOperationResult() {}
func (*RegisterGlobalConstantInternalOperationResult) OperationKind() string {
	return "register_global_constant"
}

type SetDepositsLimitInternalOperationResult struct {
	Source core.ContractID
	Nonce  uint16
	Limit  tz.Option[tz.BigUint]
	Result ConsumedGasResult
}

func (*SetDepositsLimitInternalOperationResult) InternalOperationResult() {}
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
