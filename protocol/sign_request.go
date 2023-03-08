package protocol

import (
	"bytes"

	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
)

type SignRequest interface {
	SignRequest()
}

type WithWatermark interface {
	SignRequest
	Watermark() *Watermark
}

type EmmyBlockRequest struct {
	Chain       *tz.ChainID
	BlockHeader ShellHeader
}

func (*EmmyBlockRequest) SignRequest() {}
func (r *EmmyBlockRequest) Watermark() *Watermark {
	return &Watermark{
		Chain: r.Chain,
		Level: Level{
			Level: r.BlockHeader.Level,
			Round: tz.None[int32](),
		},
		Order: WmOrderDefault,
	}
}

type TenderbakeBlockRequest struct {
	Chain       *tz.ChainID
	BlockHeader TenderbakeBlockHeader
}

func (*TenderbakeBlockRequest) SignRequest() {}
func (r *TenderbakeBlockRequest) Watermark() *Watermark {
	return &Watermark{
		Chain: r.Chain,
		Level: Level{
			Level: r.BlockHeader.Level,
			Round: tz.Some(r.BlockHeader.PayloadRound),
		},
		Order: WmOrderDefault,
	}
}

type EmmyEndorsementRequest struct {
	Chain     *tz.ChainID
	Branch    *tz.BlockHash
	Operation InlinedEmmyEndorsementContents
}

func (*EmmyEndorsementRequest) SignRequest() {}
func (r *EmmyEndorsementRequest) Watermark() *Watermark {
	return &Watermark{
		Chain: r.Chain,
		Level: Level{
			Level: r.Operation.(*EmmyEndorsement).Level,
			Round: tz.None[int32](),
		},
		Order: WmOrderEndorsement,
	}
}

type PreendorsementRequest struct {
	Chain     *tz.ChainID
	Branch    *tz.BlockHash
	Operation InlinedPreendorsementContents
}

func (*PreendorsementRequest) SignRequest() {}
func (r *PreendorsementRequest) Watermark() *Watermark {
	return &Watermark{
		Chain: r.Chain,
		Level: Level{
			Level: r.Operation.(*Preendorsement).Level,
			Round: tz.Some(r.Operation.(*Preendorsement).Round),
		},
		Order: WmOrderPreendorsement,
	}
}

type EndorsementRequest struct {
	Chain     *tz.ChainID
	Branch    *tz.BlockHash
	Operation InlinedEndorsementContents
}

func (*EndorsementRequest) SignRequest() {}
func (r *EndorsementRequest) Watermark() *Watermark {
	return &Watermark{
		Chain: r.Chain,
		Level: Level{
			Level: r.Operation.(*Endorsement).Level,
			Round: tz.Some(r.Operation.(*Endorsement).Round),
		},
		Order: WmOrderEndorsement,
	}
}

type GenericOperationRequest struct {
	Branch     *tz.BlockHash
	Operations []OperationContents
}

func (*GenericOperationRequest) SignRequest() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SignRequest]{
		Variants: encoding.Variants[SignRequest]{
			0x01: (*EmmyBlockRequest)(nil),
			0x02: (*EmmyEndorsementRequest)(nil),
			0x03: (*GenericOperationRequest)(nil),
			0x11: (*TenderbakeBlockRequest)(nil),
			0x12: (*PreendorsementRequest)(nil),
			0x13: (*EndorsementRequest)(nil),
		},
	})
}

const (
	WmOrderDefault = iota
	WmOrderPreendorsement
	WmOrderEndorsement
)

type Level struct {
	Level int32
	Round tz.Option[int32]
}

func (l *Level) Cmp(other *Level) tz.Option[int] {
	if l.Round.IsNone() && other.Round.IsSome() {
		return tz.None[int]()
	}

	if d := l.Level - other.Level; d == 0 {
		switch {
		case l.Round.IsSome() && other.Round.IsSome():
			return tz.Some(int(l.Round.Unwrap() - other.Round.Unwrap()))
		case l.Round.IsSome() && other.Round.IsNone():
			return tz.Some(1)
		default:
			return tz.Some(0)
		}
	} else {
		return tz.Some(int(d))
	}
}

type Watermark struct {
	Chain *tz.ChainID
	Level Level
	Order int
}

func (w *Watermark) Validate(stored *Watermark) bool {
	c := w.Level.Cmp(&stored.Level)
	switch {
	case !bytes.Equal(w.Chain[:], stored.Chain[:]):
		return false
	case c.IsSome() &&
		(c.Unwrap() > 0 || c.Unwrap() == 0 && w.Order > stored.Order):
		return true
	default:
		return false
	}
}
