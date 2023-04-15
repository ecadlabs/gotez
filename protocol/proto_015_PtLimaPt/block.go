package proto_015_PtLimaPt

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/protocol/proto"
)

type UnsignedProtocolBlockHeader struct {
	PayloadHash               *tz.BlockPayloadHash
	PayloadRound              int32
	ProofOfWorkNonce          *[tz.ProofOfWorkNonceBytesLen]byte
	SeedNonceHash             tz.Option[*tz.CycleNonceHash]
	LiquidityBakingToggleVote uint8
}

type ProtocolBlockHeader struct {
	UnsignedProtocolBlockHeader
	Signature *tz.GenericSignature
}

func (*ProtocolBlockHeader) ProtocolBlockHeader() {}

type UnsignedBlockHeader struct {
	proto.BlockHeader
	UnsignedProtocolBlockHeader
}

type BlockHeader struct {
	UnsignedBlockHeader
	Signature tz.AnySignature
}
