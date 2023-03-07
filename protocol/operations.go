package protocol

import (
	tz "github.com/ecadlabs/gotez"
)

type UnknownOperation struct{}

func (UnknownOperation) OperationContents() {}

type OperationContents interface {
	OperationContents()
}

type EmmyEndorsement struct {
	Level int32
}

func (*EmmyEndorsement) InlinedEndorsementContents()     {}
func (*EmmyEndorsement) InlinedEmmyEndorsementContents() {}
func (*EmmyEndorsement) OperationContents()              {}

type SeedNonceRevelation struct {
	Level int32
	Nonce *[SeedNonceBytesLen]byte
}

func (*SeedNonceRevelation) OperationContents() {}

type DoubleEndorsementEvidence struct {
	Op1  InlinedEndorsement
	Op2  InlinedEndorsement
	Slot tz.Option[uint16]
}

func (op *DoubleEndorsementEvidence) DecodeTZ(data []byte, ctx *tz.Context) (rest []byte, err error) {
	if data, err = tz.Decode(data, &op.Op1, tz.Ctx(ctx), tz.Dynamic()); err != nil {
		return nil, err
	}
	if data, err = tz.Decode(data, &op.Op2, tz.Ctx(ctx), tz.Dynamic()); err != nil {
		return nil, err
	}
	if _, ok := op.Op1.Contents.(*EmmyEndorsement); ok {
		var (
			slot uint16
			err  error
		)
		if data, err = tz.Decode(data, &slot, tz.Ctx(ctx)); err != nil {
			return nil, err
		}
		op.Slot = tz.Some(slot)
	}
	return data, nil
}

func (*DoubleEndorsementEvidence) OperationContents() {}

type InlinedEndorsement struct {
	Branch    *BlockHash
	Contents  InlinedEndorsementContents
	Signature *Signature
}

type InlinedEndorsementContents interface {
	InlinedEndorsementContents()
}

func init() {
	tz.RegisterEnum(&tz.Enum[InlinedEndorsementContents]{
		Variants: tz.Variants[InlinedEndorsementContents]{
			0:  (*EmmyEndorsement)(nil),
			21: (*Endorsement)(nil),
		},
	})
}

type Endorsement struct {
	Slot             uint16
	Level            int32
	Round            int32
	BlockPayloadHash *BlockPayloadHash
}

func (*Endorsement) InlinedEndorsementContents() {}
func (*Endorsement) OperationContents()          {}

type DoubleBakingEvidence struct {
	Block1 ShellHeader `tz:"dyn"`
	Block2 ShellHeader `tz:"dyn"`
}

func (*DoubleBakingEvidence) OperationContents() {}

type ActivateAccount struct {
	PKH    *Ed25519PublicKeyHash
	Secret *[SecretBytesLen]byte
}

func (*ActivateAccount) OperationContents() {}

type Proposals struct {
	Source    PublicKeyHash
	Period    int32
	Proposals []*ProtocolHash `tz:"dyn"`
}

func (*Proposals) OperationContents() {}

type BallotKind uint8

const (
	BallotYay BallotKind = iota
	BallotNay
	BallotPass
)

type Ballot struct {
	Source   PublicKeyHash
	Period   int32
	Proposal *ProtocolHash
	Ballot   BallotKind
}

func (*Ballot) OperationContents() {}

type DoublePreendorsementEvidence struct {
	Op1 InlinedPreendorsement `tz:"dyn"`
	Op2 InlinedPreendorsement `tz:"dyn"`
}

func (*DoublePreendorsementEvidence) OperationContents() {}

type InlinedPreendorsement struct {
	Branch    *BlockHash
	Contents  InlinedPreendorsementContents
	Signature *Signature
}

type InlinedPreendorsementContents interface {
	InlinedPreendorsementContents()
}

func init() {
	tz.RegisterEnum(&tz.Enum[InlinedPreendorsementContents]{
		Variants: tz.Variants[InlinedPreendorsementContents]{
			20: (*Preendorsement)(nil),
		},
	})
}

type Preendorsement struct {
	Slot             uint16
	Level            int32
	Round            int32
	BlockPayloadHash *BlockPayloadHash
}

func (*Preendorsement) OperationContents()             {}
func (*Preendorsement) InlinedPreendorsementContents() {}

type VDFRevelation struct {
	Field0 *[200]byte
	Field1 *[200]byte
}

func (*VDFRevelation) OperationContents() {}

type DrainDelegate struct {
	ConsensusKey PublicKeyHash
	Delegate     PublicKeyHash
	Destination  PublicKeyHash
}

func (*DrainDelegate) OperationContents() {}

type InlinedEmmyEndorsement struct {
	Branch    *BlockHash
	Contents  InlinedEmmyEndorsementContents
	Signature *Signature
}

type InlinedEmmyEndorsementContents interface {
	InlinedEmmyEndorsementContents()
}

func init() {
	tz.RegisterEnum(&tz.Enum[InlinedEmmyEndorsementContents]{
		Variants: tz.Variants[InlinedEmmyEndorsementContents]{
			0: (*EmmyEndorsement)(nil),
		},
	})
}

type EndorsementWithSlot struct {
	Endorsement InlinedEmmyEndorsement `tz:"dyn"`
	Slot        uint16
}

func (*EndorsementWithSlot) OperationContents() {}

type FailingNoop struct {
	Arbitrary []byte `tz:"dyn"`
}

func (*FailingNoop) OperationContents() {}

type ManagerOperation struct {
	Source       PublicKeyHash
	Fee          tz.BigUint
	Counter      tz.BigUint
	GasLimit     tz.BigUint
	StorageLimit tz.BigUint
}

type Reveal struct {
	ManagerOperation
	PublicKey PublicKey
}

func (*Reveal) OperationContents() {}

type Transaction struct {
	ManagerOperation
	Amount      tz.BigUint
	Destination ContractID
	Parameters  tz.Option[Parameters]
}

func (*Transaction) OperationContents() {}

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
	String
}

func (EpDefault) Entrypoint()        {}
func (EpRoot) Entrypoint()           {}
func (EpDo) Entrypoint()             {}
func (EpSetDelegate) Entrypoint()    {}
func (EpRemoveDelegate) Entrypoint() {}
func (EpNamed) Entrypoint()          {}

func init() {
	tz.RegisterEnum(&tz.Enum[Entrypoint]{
		Variants: tz.Variants[Entrypoint]{
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
	Delegate tz.Option[PublicKeyHash]
	Code     []byte `tz:"dyn"`
	Storage  []byte `tz:"dyn"`
}

func (*Origination) OperationContents() {}

type Delegation struct {
	ManagerOperation
	Delegate tz.Option[PublicKeyHash]
}

func (*Delegation) OperationContents() {}

type RegisterGlobalConstant struct {
	ManagerOperation
	Value []byte `tz:"dyn"`
}

func (*RegisterGlobalConstant) OperationContents() {}

type SetDepositsLimit struct {
	ManagerOperation
	Limit tz.Option[tz.BigUint]
}

func (*SetDepositsLimit) OperationContents() {}

type IncreasePaidStorage struct {
	ManagerOperation
	Amount      tz.BigInt
	Destination OriginatedContractID
}

func (*IncreasePaidStorage) OperationContents() {}

type UpdateConsensusKey struct {
	ManagerOperation
	PublicKey PublicKey
}

func (*UpdateConsensusKey) OperationContents() {}

func init() {
	tz.RegisterEnum(&tz.Enum[OperationContents]{
		Variants: tz.Variants[OperationContents]{
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
