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

func (*BlockSignRequest) SignRequestKind() string { return "block" }

type PreendorsementSignRequest struct {
	Chain     *tz.ChainID
	Branch    *tz.BlockHash
	Operation latest.InlinedPreendorsementContents
}

func (*PreendorsementSignRequest) SignRequestKind() string { return "preendorsement" }

type EndorsementSignRequest struct {
	Chain     *tz.ChainID
	Branch    *tz.BlockHash
	Operation latest.InlinedEndorsementContents
}

func (*EndorsementSignRequest) SignRequestKind() string { return "endorsement" }

type GenericOperationSignRequest latest.UnsignedOperation

func (*GenericOperationSignRequest) SignRequestKind() string { return "generic" }

func init() {
	encoding.RegisterEnum(&encoding.Enum[SignRequest]{
		Variants: encoding.Variants[SignRequest]{
			0x03: (*GenericOperationSignRequest)(nil),
			0x11: (*BlockSignRequest)(nil),
			0x12: (*PreendorsementSignRequest)(nil),
			0x13: (*EndorsementSignRequest)(nil),
		},
	})
}
