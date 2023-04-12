package operations

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/expression"
	"github.com/ecadlabs/gotez/protocol/shell"
)

type OperationContents interface {
	OperationKind() string
}

type OperationContentsAndResult interface {
	OperationContentsAndResult()
	OperationContents
}

type SeedNonceRevelation struct {
	Level int32
	Nonce *[tz.SeedNonceBytesLen]byte
}

func (*SeedNonceRevelation) OperationKind() string { return "seed_nonce_revelation" }

type SeedNonceRevelationContentsAndResult struct {
	SeedNonceRevelation
	Metadata []*BalanceUpdate `tz:"dyn"`
}

func (*SeedNonceRevelationContentsAndResult) OperationContentsAndResult() {}

type DoubleEndorsementEvidence struct {
	Op1 InlinedEndorsement `tz:"dyn"`
	Op2 InlinedEndorsement `tz:"dyn"`
}

func (*DoubleEndorsementEvidence) OperationKind() string { return "double_endorsement_evidence" }

type DoubleEndorsementEvidenceContentsAndResult struct {
	DoubleEndorsementEvidence
	Metadata []*BalanceUpdate `tz:"dyn"`
}

func (*DoubleEndorsementEvidenceContentsAndResult) OperationContentsAndResult() {}

type InlinedEndorsement struct {
	Branch    *tz.BlockHash
	Contents  InlinedEndorsementContents
	Signature *tz.GenericSignature
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

type Endorsement struct {
	Slot             uint16
	Level            int32
	Round            int32
	BlockPayloadHash *tz.BlockPayloadHash
}

func (*Endorsement) InlinedEndorsementContents() {}
func (*Endorsement) OperationKind() string       { return "endorsement" }

type EndorsementMetadata struct {
	BalanceUpdates   []*BalanceUpdate `tz:"dyn"`
	Delegate         tz.PublicKeyHash
	EndorsementPower int32
	ConsensusKey     tz.PublicKeyHash
}

type EndorsementContentsAndResult struct {
	Endorsement
	Metadata EndorsementMetadata
}

func (*EndorsementContentsAndResult) OperationContentsAndResult() {}

type DoubleBakingEvidence struct {
	Block1 shell.BlockHeader `tz:"dyn"`
	Block2 shell.BlockHeader `tz:"dyn"`
}

func (*DoubleBakingEvidence) OperationKind() string { return "double_baking_evidence" }

type DoubleBakingEvidenceContentsAndResult struct {
	DoubleBakingEvidence
	Metadata []*BalanceUpdate `tz:"dyn"`
}

func (*DoubleBakingEvidenceContentsAndResult) OperationContentsAndResult() {}

type ActivateAccount struct {
	PKH    *tz.Ed25519PublicKeyHash
	Secret *[tz.SecretBytesLen]byte
}

func (*ActivateAccount) OperationKind() string { return "activate_account" }

type ActivateAccountContentsAndResult struct {
	ActivateAccount
	Metadata []*BalanceUpdate `tz:"dyn"`
}

func (*ActivateAccountContentsAndResult) OperationContentsAndResult() {}

type Proposals struct {
	Source    tz.PublicKeyHash
	Period    int32
	Proposals []*tz.ProtocolHash `tz:"dyn"`
}

func (*Proposals) OperationKind() string       { return "proposals" }
func (*Proposals) OperationContentsAndResult() {}

type BallotKind uint8

const (
	BallotYay BallotKind = iota
	BallotNay
	BallotPass
)

type Ballot struct {
	Source   tz.PublicKeyHash
	Period   int32
	Proposal *tz.ProtocolHash
	Ballot   BallotKind
}

func (*Ballot) OperationKind() string       { return "ballot" }
func (*Ballot) OperationContentsAndResult() {}

type DoublePreendorsementEvidence struct {
	Op1 InlinedPreendorsement `tz:"dyn"`
	Op2 InlinedPreendorsement `tz:"dyn"`
}

func (*DoublePreendorsementEvidence) OperationKind() string { return "double_preendorsement_evidence" }

type DoublePreendorsementEvidenceContentsAndResult struct {
	DoublePreendorsementEvidence
	Metadata []*BalanceUpdate `tz:"dyn"`
}

func (*DoublePreendorsementEvidenceContentsAndResult) OperationContentsAndResult() {}

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

type PreendorsementMetadata = EndorsementMetadata
type PreendorsementContentsAndResult struct {
	Preendorsement
	Metadata PreendorsementMetadata
}

func (*PreendorsementContentsAndResult) OperationContentsAndResult() {}

type VDFRevelation struct {
	Field0 *[100]byte
	Field1 *[100]byte
}

func (*VDFRevelation) OperationKind() string { return "vdf_revelation" }

type VDFRevelationContentsAndResult struct {
	VDFRevelation
	Metadata []*BalanceUpdate `tz:"dyn"`
}

func (*VDFRevelationContentsAndResult) OperationContentsAndResult() {}

type DrainDelegate struct {
	ConsensusKey tz.PublicKeyHash
	Delegate     tz.PublicKeyHash
	Destination  tz.PublicKeyHash
}

func (*DrainDelegate) OperationKind() string { return "drain_delegate" }

type DrainDelegateMetadata struct {
	BalanceUpdates               []*BalanceUpdate `tz:"dyn"`
	AllocatedDestinationContract bool
}

type DrainDelegateContentsAndResult struct {
	DrainDelegate
	Metadata DrainDelegateMetadata
}

func (*DrainDelegateContentsAndResult) OperationContentsAndResult() {}

type FailingNoop struct {
	Arbitrary []byte `tz:"dyn"`
}

func (*FailingNoop) OperationKind() string { return "failing_noop" }

type ManagerOperation struct {
	Source       tz.PublicKeyHash
	Fee          tz.BigUint
	Counter      tz.BigUint
	GasLimit     tz.BigUint
	StorageLimit tz.BigUint
}

type SignaturePrefix struct {
	SignaturePrefix SignaturePrefixPayload
}

func (*SignaturePrefix) OperationKind() string       { return "signature_prefix" }
func (*SignaturePrefix) OperationContentsAndResult() {}

type SignaturePrefixPayload interface {
	SignaturePrefixPayload()
}

type BLSSignaturePrefix [32]byte

func (*BLSSignaturePrefix) SignaturePrefixPayload() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SignaturePrefixPayload]{
		Variants: encoding.Variants[SignaturePrefixPayload]{
			3: (*BLSSignaturePrefix)(nil),
		},
	})

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
			107: (*Reveal)(nil),
			108: (*Transaction)(nil),
			109: (*Origination)(nil),
			110: (*Delegation)(nil),
			111: (*RegisterGlobalConstant)(nil),
			112: (*SetDepositsLimit)(nil),
			113: (*IncreasePaidStorage)(nil),
			114: (*UpdateConsensusKey)(nil),
			200: (*SmartRollupOriginate)(nil),
			255: (*SignaturePrefix)(nil),
		},
	})

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
			107: (*RevealContentsAndResult)(nil),
			108: (*TransactionContentsAndResult)(nil),
			109: (*OriginationContentsAndResult)(nil),
			110: (*DelegationContentsAndResult)(nil),
			111: (*RegisterGlobalConstantContentsAndResult)(nil),
			112: (*SetDepositsLimitContentsAndResult)(nil),
			113: (*IncreasePaidStorageContentsAndResult)(nil),
			114: (*UpdateConsensusKeyContentsAndResult)(nil),
			200: (*SmartRollupOriginateContentsAndResult)(nil),
			255: (*SignaturePrefix)(nil),
		},
	})
}

type SuccessfulManagerOperationResult interface {
	OperationContents
	SuccessfulManagerOperationResult()
}

type EventResult interface {
	EventResult()
	OperationResult
}

type EventResultContents struct {
	ConsumedMilligas tz.BigUint
}

type EventResultApplied struct {
	OperationResultApplied[EventResultContents]
}

func (*EventResultApplied) EventResult() {}

type EventResultBacktracked struct {
	OperationResultBacktracked[EventResultContents]
}

func (*EventResultBacktracked) EventResult() {}

type EventResultFailed struct{ OperationResultFailed }

func (*EventResultFailed) EventResult() {}

type EventResultSkipped struct{ OperationResultSkipped }

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

type MetadataWithResult[T OperationResult] struct {
	BalanceUpdates           []*BalanceUpdate `tz:"dyn"`
	OperationResult          T
	InternalOperationResults []InternalOperationResult `tz:"dyn"`
}

type InternalOperationResult interface {
	InternalOperationResult()
}

type EventInternalOperationResult struct {
	Source  tz.TransactionDestination
	Nonce   uint16
	Type    expression.Expression
	Tag     tz.Option[Entrypoint]
	Payload tz.Option[expression.Expression]
	Result  EventResult
}

func (*EventInternalOperationResult) InternalOperationResult() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[InternalOperationResult]{
		Variants: encoding.Variants[InternalOperationResult]{
			1: (*TransactionInternalOperationResult)(nil),
			2: (*OriginationInternalOperationResult)(nil),
			3: (*DelegationInternalOperationResult)(nil),
			4: (*EventInternalOperationResult)(nil),
		},
	})

	encoding.RegisterEnum(&encoding.Enum[SuccessfulManagerOperationResult]{
		Variants: encoding.Variants[SuccessfulManagerOperationResult]{
			0:   (*RevealSuccessfulManagerOperationResult)(nil),
			1:   (*TransactionSuccessfulManagerOperationResult)(nil),
			2:   (*OriginationSuccessfulManagerOperationResult)(nil),
			3:   (*DelegationSuccessfulManagerOperationResult)(nil),
			5:   (*SetDepositsLimitSuccessfulManagerOperationResult)(nil),
			6:   (*UpdateConsensusKeySuccessfulManagerOperationResult)(nil),
			9:   (*IncreasePaidStorageSuccessfulManagerOperationResult)(nil),
			200: (*SmartRollupOriginateSuccessfulManagerOperationResult)(nil),
		},
	})
}

type LazyStorageDiff struct {
	Opaque []byte `tz:"dyn"` // TODO: lazy storage diff
}
