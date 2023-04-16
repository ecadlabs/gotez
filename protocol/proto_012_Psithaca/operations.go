package proto_012_Psithaca

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/proto_005_PsBABY5H"
)

type ManagerOperation = proto_005_PsBABY5H.ManagerOperation

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
