package protocol

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/latest"
)

type SignRequest interface {
	SignRequestKind() string
}

type BlockSignRequest struct {
	Chain       *tz.ChainID
	BlockHeader latest.UnsignedBlockHeader
}

func (r *BlockSignRequest) GetChainID() *tz.ChainID { return r.Chain }
func (r *BlockSignRequest) GetLevel() int32         { return r.BlockHeader.Level }
func (r *BlockSignRequest) GetRound() int32         { return r.BlockHeader.PayloadRound }
func (*BlockSignRequest) SignRequestKind() string   { return "block" }

type PreattestationSignRequest struct {
	Chain     *tz.ChainID
	Branch    *tz.BlockHash
	Operation latest.InlinedPreendorsementContents
}

type PreendorsementSignRequest = PreattestationSignRequest

func (r *PreattestationSignRequest) GetChainID() *tz.ChainID { return r.Chain }
func (r *PreattestationSignRequest) GetLevel() int32 {
	return r.Operation.(*latest.Preattestation).Level
}
func (r *PreattestationSignRequest) GetRound() int32 {
	return r.Operation.(*latest.Preattestation).Round
}
func (*PreattestationSignRequest) SignRequestKind() string { return "preendorsement" }

type AttestationSignRequest struct {
	Chain     *tz.ChainID
	Branch    *tz.BlockHash
	Operation latest.InlinedEndorsementContents
}

type EndorsementSignRequest = AttestationSignRequest

func (r *AttestationSignRequest) GetChainID() *tz.ChainID { return r.Chain }
func (r *AttestationSignRequest) GetLevel() int32         { return r.Operation.(*latest.Attestation).Level }
func (r *AttestationSignRequest) GetRound() int32         { return r.Operation.(*latest.Attestation).Round }
func (*AttestationSignRequest) SignRequestKind() string   { return "endorsement" }

type GenericOperationSignRequest latest.UnsignedOperation

func (*GenericOperationSignRequest) SignRequestKind() string { return "generic" }

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
