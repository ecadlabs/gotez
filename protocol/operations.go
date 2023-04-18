package protocol

import (
	"fmt"

	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/proto_016_PtMumbai"
)

type OperationContents interface {
	core.OperationContents
}

func init() {
	encoding.RegisterType(func(data []byte, ctx *encoding.Context) (OperationContents, []byte, error) {
		p, ok := ctx.Get(core.ProtocolVersionCtxKey).(core.Protocol)
		if !ok {
			return nil, nil, fmt.Errorf("gotez: protocol version must be passed to the decoder chain")
		}

		var (
			dest any
			out  OperationContents
		)
		switch p {
		case core.Proto016PtMumbai:
			var tmp proto_016_PtMumbai.OperationContents
			dest = &tmp
			out = tmp

		default:
			return nil, nil, fmt.Errorf("gotez: unknown protocol version %d", p)
		}

		data, err := encoding.Decode(data, dest, encoding.Ctx(ctx))
		return out, data, err
	})
}

type OperationContentsAndResult interface {
	core.OperationContentsAndResult
}

func init() {
	encoding.RegisterType(func(data []byte, ctx *encoding.Context) (OperationContentsAndResult, []byte, error) {
		p, ok := ctx.Get(core.ProtocolVersionCtxKey).(core.Protocol)
		if !ok {
			return nil, nil, fmt.Errorf("gotez: protocol version must be passed to the decoder chain")
		}

		var (
			dest any
			out  proto_016_PtMumbai.OperationContentsAndResult
		)
		switch p {
		case core.Proto016PtMumbai:
			var tmp proto_016_PtMumbai.OperationContentsAndResult
			dest = &tmp
			out = tmp

		default:
			return nil, nil, fmt.Errorf("gotez: unknown protocol version %d", p)
		}

		data, err := encoding.Decode(data, dest, encoding.Ctx(ctx))
		return out, data, err
	})
}

type SuccessfulManagerOperationResult interface {
	core.SuccessfulManagerOperationResult
}

type InternalOperationResult interface {
	core.InternalOperationResult
}

type OperationsList struct {
	Operations []*OperationsGroup `tz:"dyn,dyn"` // yes, twice
}

type OperationsGroup struct {
	ChainID  *tz.ChainID
	Hash     *tz.BlockHash
	Branch   *tz.BlockHash `tz:"dyn"`
	Contents GroupContents `tz:"dyn"`
}

type Signed interface {
	GetSignature() (tz.Signature, error)
}

type GroupContents interface {
	Signed
	GroupContents()
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[GroupContents]{
		Variants: encoding.Variants[GroupContents]{
			0: (*OperationWithTooLargeMetadata)(nil),
			1: (*OperationWithoutMetadata)(nil),
			2: (*OperationWithOptionalMetadata)(nil),
		},
	})
}

type OperationWithTooLargeMetadata struct {
	OperationWithoutMetadata
}

type OperationWithoutMetadata struct {
	Contents        []OperationContents
	SignatureSuffix *tz.GenericSignature
}

func (*OperationWithoutMetadata) GroupContents() {}
func (op *OperationWithoutMetadata) GetSignature() (tz.Signature, error) {
	if len(op.Contents) != 0 {
		if prefix, ok := op.Contents[len(op.Contents)-1].(*proto_016_PtMumbai.SignaturePrefix); ok {
			if blsPrefix, ok := prefix.SignaturePrefix.(*proto_016_PtMumbai.BLSSignaturePrefix); ok {
				var sig tz.BLSSignature
				copy(sig[:], blsPrefix[:])
				copy(sig[:len(blsPrefix)], op.SignatureSuffix[:])
				return &sig, nil
			}
		}
	}
	return op.SignatureSuffix, nil
}

type OperationWithOptionalMetadata struct {
	Contents OperationWithOptionalMetadataContents
}

func (op *OperationWithOptionalMetadata) GetSignature() (tz.Signature, error) {
	return op.Contents.GetSignature()
}

func (*OperationWithOptionalMetadata) GroupContents() {}

type OperationWithOptionalMetadataContents interface {
	Signed
	OperationWithOptionalMetadataContents()
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[OperationWithOptionalMetadataContents]{
		Variants: encoding.Variants[OperationWithOptionalMetadataContents]{
			0: (*OperationWithOptionalMetadataWithMetadata)(nil),
			1: (*OperationWithOptionalMetadataWithoutMetadata)(nil),
		},
	})
}

type OperationWithOptionalMetadataWithMetadata struct {
	Contents  []OperationContentsAndResult `tz:"dyn"`
	Signature tz.AnySignature
}

func (*OperationWithOptionalMetadataWithMetadata) OperationWithOptionalMetadataContents() {}
func (op *OperationWithOptionalMetadataWithMetadata) GetSignature() (tz.Signature, error) {
	return op.Signature.Signature()
}

type OperationWithOptionalMetadataWithoutMetadata struct {
	Contents  []OperationContents `tz:"dyn"`
	Signature tz.AnySignature
}

func (*OperationWithOptionalMetadataWithoutMetadata) OperationWithOptionalMetadataContents() {}
func (op *OperationWithOptionalMetadataWithoutMetadata) GetSignature() (tz.Signature, error) {
	return op.Signature.Signature()
}
