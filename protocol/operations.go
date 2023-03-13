package protocol

import (
	"bytes"
	"errors"

	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
)

type UnknownOperation struct{}

func (UnknownOperation) OperationKind() string { return "unknown" }

type OperationContents interface {
	OperationKind() string
}

type EmmyEndorsement struct {
	Level int32
}

func (*EmmyEndorsement) InlinedEndorsementContents()     {}
func (*EmmyEndorsement) InlinedEmmyEndorsementContents() {}
func (*EmmyEndorsement) OperationKind() string           { return "endorsement" }

type SeedNonceRevelation struct {
	Level int32
	Nonce *[tz.SeedNonceBytesLen]byte
}

func (*SeedNonceRevelation) OperationKind() string { return "seed_nonce_revelation" }

type DoubleEndorsementEvidence struct {
	Op1  InlinedEndorsement
	Op2  InlinedEndorsement
	Slot tz.Option[uint16]
}

func (op *DoubleEndorsementEvidence) DecodeTZ(data []byte, ctx *encoding.Context) (rest []byte, err error) {
	if data, err = encoding.Decode(data, &op.Op1, encoding.Ctx(ctx), encoding.Dynamic()); err != nil {
		return nil, err
	}
	if data, err = encoding.Decode(data, &op.Op2, encoding.Ctx(ctx), encoding.Dynamic()); err != nil {
		return nil, err
	}
	if _, ok := op.Op1.Contents.(*EmmyEndorsement); ok {
		var (
			slot uint16
			err  error
		)
		if data, err = encoding.Decode(data, &slot, encoding.Ctx(ctx)); err != nil {
			return nil, err
		}
		op.Slot = tz.Some(slot)
	}
	return data, nil
}

func (op *DoubleEndorsementEvidence) EncodeTZ(ctx *encoding.Context) ([]byte, error) {
	var buf bytes.Buffer
	if err := encoding.Encode(&buf, &op.Op1, encoding.Ctx(ctx), encoding.Dynamic()); err != nil {
		return nil, err
	}
	if err := encoding.Encode(&buf, &op.Op2, encoding.Ctx(ctx), encoding.Dynamic()); err != nil {
		return nil, err
	}
	if _, ok := op.Op1.Contents.(*EmmyEndorsement); ok {
		if op.Slot.IsNone() {
			return nil, errors.New("gotex: DoubleEndorsementEvidence: slot is required")
		}
		slot := op.Slot.Unwrap()
		if err := encoding.Encode(&buf, &slot, encoding.Ctx(ctx)); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func (*DoubleEndorsementEvidence) OperationKind() string { return "double_endorsement_evidence" }

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
			0:  (*EmmyEndorsement)(nil),
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

type DoubleBakingEvidence struct {
	Block1 ShellHeader `tz:"dyn"`
	Block2 ShellHeader `tz:"dyn"`
}

func (*DoubleBakingEvidence) OperationKind() string { return "double_baking_evidence" }

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

func (*Proposals) OperationKind() string { return "proposals" }

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

func (*Ballot) OperationKind() string { return "ballot" }

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

type VDFRevelation struct {
	Field0 *[200]byte
	Field1 *[200]byte
}

func (*VDFRevelation) OperationKind() string { return "vdf_revelation" }

type DrainDelegate struct {
	ConsensusKey tz.PublicKeyHash
	Delegate     tz.PublicKeyHash
	Destination  tz.PublicKeyHash
}

func (*DrainDelegate) OperationKind() string { return "drain_delegate" }

type InlinedEmmyEndorsement struct {
	Branch    *tz.BlockHash
	Contents  InlinedEmmyEndorsementContents
	Signature *tz.GenericSignature
}

type InlinedEmmyEndorsementContents interface {
	InlinedEmmyEndorsementContents()
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[InlinedEmmyEndorsementContents]{
		Variants: encoding.Variants[InlinedEmmyEndorsementContents]{
			0: (*EmmyEndorsement)(nil),
		},
	})
}

type EndorsementWithSlot struct {
	Endorsement InlinedEmmyEndorsement `tz:"dyn"`
	Slot        uint16
}

func (*EndorsementWithSlot) OperationKind() string { return "endorsement_with_slot" }

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

type Reveal struct {
	ManagerOperation
	PublicKey tz.PublicKey
}

func (*Reveal) OperationKind() string { return "reveal" }

type Transaction struct {
	ManagerOperation
	Amount      tz.BigUint
	Destination tz.ContractID
	Parameters  tz.Option[Parameters]
}

func (*Transaction) OperationKind() string { return "transaction" }

type Parameters struct {
	Entrypoint Entrypoint
	Value      []byte `tz:"dyn"`
}

type Entrypoint interface {
	Entrypoint()
}

type EpDefault struct{}
type EpRoot struct{}
type EpDo struct{}
type EpSetDelegate struct{}
type EpRemoveDelegate struct{}
type EpNamed struct {
	tz.String
}

func (EpDefault) Entrypoint()        {}
func (EpRoot) Entrypoint()           {}
func (EpDo) Entrypoint()             {}
func (EpSetDelegate) Entrypoint()    {}
func (EpRemoveDelegate) Entrypoint() {}
func (EpNamed) Entrypoint()          {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[Entrypoint]{
		Variants: encoding.Variants[Entrypoint]{
			0:   EpDefault{},
			1:   EpRoot{},
			2:   EpDo{},
			3:   EpSetDelegate{},
			4:   EpRemoveDelegate{},
			255: EpNamed{},
		},
	})
}

type Origination struct {
	ManagerOperation
	Balance  tz.BigUint
	Delegate tz.Option[tz.PublicKeyHash]
	Code     []byte `tz:"dyn"`
	Storage  []byte `tz:"dyn"`
}

func (*Origination) OperationKind() string { return "origination" }

type Delegation struct {
	ManagerOperation
	Delegate tz.Option[tz.PublicKeyHash]
}

func (*Delegation) OperationKind() string { return "delegation" }

type RegisterGlobalConstant struct {
	ManagerOperation
	Value []byte `tz:"dyn"`
}

func (*RegisterGlobalConstant) OperationKind() string { return "register_global_constant" }

type SetDepositsLimit struct {
	ManagerOperation
	Limit tz.Option[tz.BigUint]
}

func (*SetDepositsLimit) OperationKind() string { return "set_deposits_limit" }

type IncreasePaidStorage struct {
	ManagerOperation
	Amount      tz.BigInt
	Destination tz.OriginatedContractID
}

func (*IncreasePaidStorage) OperationKind() string { return "increase_paid_storage" }

type UpdateConsensusKey struct {
	ManagerOperation
	PublicKey tz.PublicKey
}

func (*UpdateConsensusKey) OperationKind() string { return "update_consensus_key" }

func init() {
	encoding.RegisterEnum(&encoding.Enum[OperationContents]{
		Variants: encoding.Variants[OperationContents]{
			0:   (*EmmyEndorsement)(nil),
			1:   (*SeedNonceRevelation)(nil),
			2:   (*DoubleEndorsementEvidence)(nil),
			3:   (*DoubleBakingEvidence)(nil),
			4:   (*ActivateAccount)(nil),
			5:   (*Proposals)(nil),
			6:   (*Ballot)(nil),
			7:   (*DoublePreendorsementEvidence)(nil),
			8:   (*VDFRevelation)(nil),
			9:   (*DrainDelegate)(nil),
			10:  (*EndorsementWithSlot)(nil),
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
		},
		Default: UnknownOperation{},
	})
}
