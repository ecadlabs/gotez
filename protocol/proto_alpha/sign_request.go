package proto_alpha

import (
	"fmt"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
)

type PreattestationRequestContent interface {
	core.InlinedConsensusOperationContent
}

type AttestationRequestContent interface {
	core.InlinedConsensusOperationContent
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[PreattestationRequestContent]{
		Variants: encoding.Variants[PreattestationRequestContent]{
			20: (*Preattestation)(nil),
		},
	})
	encoding.RegisterEnum(&encoding.Enum[AttestationRequestContent]{
		Variants: encoding.Variants[AttestationRequestContent]{
			21: (*Attestation)(nil),
			23: (*AttestationWithDAL)(nil),
		},
	})
}

type SignRequest interface {
	core.SignRequest
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SignRequest]{
		Variants: encoding.Variants[SignRequest]{
			0x03: (*GenericOperationSignRequest)(nil),
			0x11: (*BlockSignRequest)(nil),
			0x12: (*PreattestationSignRequest)(nil),
			0x13: (*AttestationSignRequest)(nil),
		},
	})
}

type GenericOperationSignRequest UnsignedOperation

func (*GenericOperationSignRequest) SignRequestKind() string { return "generic" }

type BlockSignRequest struct {
	Chain       *tz.ChainID
	BlockHeader UnsignedBlockHeader
}

func (r *BlockSignRequest) GetChainID() *tz.ChainID { return r.Chain }
func (r *BlockSignRequest) GetLevel() int32         { return r.BlockHeader.Level }
func (r *BlockSignRequest) GetRound() int32         { return r.BlockHeader.PayloadRound }
func (*BlockSignRequest) SignRequestKind() string   { return "block" }

type ConsensusSignRequest[T core.OperationContents] struct {
	Chain     *tz.ChainID
	Branch    *tz.BlockHash
	Operation T
}

type PreattestationSignRequest ConsensusSignRequest[PreattestationRequestContent]

func (r *PreattestationSignRequest) GetChainID() *tz.ChainID { return r.Chain }
func (r *PreattestationSignRequest) GetLevel() int32 {
	return r.Operation.(*Preattestation).Level
}
func (r *PreattestationSignRequest) GetRound() int32 {
	return r.Operation.(*Preattestation).Round
}
func (*PreattestationSignRequest) SignRequestKind() string { return "preattestation" }

type AttestationSignRequest ConsensusSignRequest[AttestationRequestContent]

func (r *AttestationSignRequest) GetChainID() *tz.ChainID { return r.Chain }
func (r *AttestationSignRequest) GetLevel() int32 {
	switch op := r.Operation.(type) {
	case *Attestation:
		return op.Level
	case *AttestationWithDAL:
		return op.Level
	default:
		panic(fmt.Sprintf("unexpected operation %T", r.Operation))
	}
}
func (r *AttestationSignRequest) GetRound() int32 {
	switch op := r.Operation.(type) {
	case *Attestation:
		return op.Round
	case *AttestationWithDAL:
		return op.Round
	default:
		panic(fmt.Sprintf("unexpected operation %T", r.Operation))
	}
}
func (*AttestationSignRequest) SignRequestKind() string { return "attestation" }
