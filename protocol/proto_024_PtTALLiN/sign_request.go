package proto_024_PtTALLiN

import (
	"fmt"
	"slices"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
)

type PreattestationRequestContent interface {
	core.OperationContents
}

type AttestationRequestContent interface {
	core.OperationContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[PreattestationRequestContent]{
		Variants: encoding.Variants[PreattestationRequestContent]{
			20: (*Preattestation)(nil),
			40: (*BLSModePreattestation)(nil),
		},
	})
	encoding.RegisterEnum(&encoding.Enum[AttestationRequestContent]{
		Variants: encoding.Variants[AttestationRequestContent]{
			21: (*Attestation)(nil),
			23: (*AttestationWithDAL)(nil),
			41: (*BLSModeAttestation)(nil),
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

func (c *ConsensusSignRequest[T]) listOperations() []string {
	vars := encoding.ListVariants[T]()
	ret := make([]string, len(vars))
	for i, v := range vars {
		ret[i] = v.OperationKind()
	}
	slices.Sort(ret)
	return slices.Compact(ret)
}

type PreattestationSignRequest ConsensusSignRequest[PreattestationRequestContent]

func (r *PreattestationSignRequest) GetChainID() *tz.ChainID { return r.Chain }

func (r *PreattestationSignRequest) GetLevel() int32 {
	switch op := r.Operation.(type) {
	case *Preattestation:
		return op.Level
	case *BLSModePreattestation:
		return op.Level
	default:
		panic(fmt.Sprintf("unexpected operation %T", r.Operation))
	}
}

func (r *PreattestationSignRequest) GetRound() int32 {
	switch op := r.Operation.(type) {
	case *Preattestation:
		return op.Round
	case *BLSModePreattestation:
		return op.Round
	default:
		panic(fmt.Sprintf("unexpected operation %T", r.Operation))
	}
}

func (r *PreattestationSignRequest) SignRequestKind() string { return r.Operation.OperationKind() }

func (r *PreattestationSignRequest) listOperations() []string {
	return (*ConsensusSignRequest[PreattestationRequestContent])(r).listOperations()
}

type AttestationSignRequest ConsensusSignRequest[AttestationRequestContent]

func (r *AttestationSignRequest) GetChainID() *tz.ChainID { return r.Chain }

func (r *AttestationSignRequest) GetLevel() int32 {
	switch op := r.Operation.(type) {
	case *Attestation:
		return op.Level
	case *BLSModeAttestation:
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
	case *BLSModeAttestation:
		return op.Round
	case *AttestationWithDAL:
		return op.Round
	default:
		panic(fmt.Sprintf("unexpected operation %T", r.Operation))
	}
}

func (r *AttestationSignRequest) SignRequestKind() string { return r.Operation.OperationKind() }

func (r *AttestationSignRequest) listOperations() []string {
	return (*ConsensusSignRequest[AttestationRequestContent])(r).listOperations()
}

type consensusSignRequest interface {
	listOperations() []string
}

func ListSignRequests() []string {
	var kinds []string
	for _, variant := range encoding.ListVariants[SignRequest]() {
		switch v := variant.(type) {
		case consensusSignRequest:
			kinds = append(kinds, v.listOperations()...)
		default:
			kinds = append(kinds, v.SignRequestKind())
		}
	}
	slices.Sort(kinds)
	return slices.Compact(kinds)
}
