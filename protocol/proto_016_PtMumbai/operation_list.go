package proto_016_PtMumbai

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
)

type GroupContents interface {
	core.GroupContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[GroupContents]{
		Variants: encoding.Variants[GroupContents]{
			0: (*OperationWithTooLargeMetadata[OperationContents])(nil),
			1: (*OperationWithoutMetadata[OperationContents])(nil),
			2: (*core.OperationWithOptionalMetadata[OperationWithOptionalMetadataContents])(nil),
		},
	})
}

type OperationWithTooLargeMetadata[T core.OperationContents] struct {
	OperationWithoutMetadata[T]
}

type OperationWithoutMetadata[T core.OperationContents] struct {
	core.OperationWithoutMetadata[T]
}

func (op *OperationWithoutMetadata[T]) GetSignature() (tz.Option[tz.Signature], error) {
	if len(op.Contents) != 0 {
		var last core.OperationContents = op.Contents[len(op.Contents)-1]
		if prefix, ok := last.(*SignaturePrefix); ok {
			if blsPrefix, ok := prefix.SignaturePrefix.(*BLSSignaturePrefix); ok {
				var sig tz.BLSSignature
				copy(sig[:], blsPrefix[:])
				copy(sig[:len(blsPrefix)], op.Signature[:])
				return tz.Some[tz.Signature](&sig), nil
			}
		}
	}
	return tz.Some[tz.Signature](op.Signature), nil
}

type OperationWithOptionalMetadataContents interface {
	core.OperationWithOptionalMetadataContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[OperationWithOptionalMetadataContents]{
		Variants: encoding.Variants[OperationWithOptionalMetadataContents]{
			0: (*OperationWithOptionalMetadataWithMetadata[OperationContentsAndResult])(nil),
			1: (*OperationWithOptionalMetadataWithoutMetadata[OperationContents])(nil),
		},
	})
}

type OperationWithOptionalMetadataWithMetadata[T core.OperationContentsAndResult] struct {
	Contents  []T             `tz:"dyn" json:"contents"`
	Signature tz.AnySignature `json:"signature"`
}

func (ops *OperationWithOptionalMetadataWithMetadata[T]) Operations() []core.OperationContents {
	out := make([]core.OperationContents, len(ops.Contents))
	for i, op := range ops.Contents {
		out[i] = op
	}
	return out
}

func (*OperationWithOptionalMetadataWithMetadata[T]) OperationWithOptionalMetadataContents() {}
func (op *OperationWithOptionalMetadataWithMetadata[T]) GetSignature() (tz.Option[tz.Signature], error) {
	if len(op.Signature) == 0 {
		return tz.None[tz.Signature](), nil
	}
	if len(op.Contents) != 0 {
		var last core.OperationContentsAndResult = op.Contents[len(op.Contents)-1]
		if prefix, ok := last.(*SignaturePrefix); ok {
			if blsPrefix, ok := prefix.SignaturePrefix.(*BLSSignaturePrefix); ok {
				var sig tz.BLSSignature
				copy(sig[:], blsPrefix[:])
				copy(sig[:len(blsPrefix)], op.Signature)
				return tz.Some[tz.Signature](&sig), nil
			}
		}
	}
	sig, err := op.Signature.Signature()
	if err != nil {
		return tz.None[tz.Signature](), err
	}
	return tz.Some(sig), nil
}

type OperationWithOptionalMetadataWithoutMetadata[T core.OperationContents] struct {
	Contents  []T             `tz:"dyn" json:"contents"`
	Signature tz.AnySignature `json:"signature"`
}

func (ops *OperationWithOptionalMetadataWithoutMetadata[T]) Operations() []core.OperationContents {
	out := make([]core.OperationContents, len(ops.Contents))
	for i, op := range ops.Contents {
		out[i] = op
	}
	return out
}

func (*OperationWithOptionalMetadataWithoutMetadata[T]) OperationWithOptionalMetadataContents() {}
func (op *OperationWithOptionalMetadataWithoutMetadata[T]) GetSignature() (tz.Option[tz.Signature], error) {
	if len(op.Signature) == 0 {
		return tz.None[tz.Signature](), nil
	}
	if len(op.Contents) != 0 {
		var last core.OperationContents = op.Contents[len(op.Contents)-1]
		if prefix, ok := last.(*SignaturePrefix); ok {
			if blsPrefix, ok := prefix.SignaturePrefix.(*BLSSignaturePrefix); ok {
				var sig tz.BLSSignature
				copy(sig[:], blsPrefix[:])
				copy(sig[:len(blsPrefix)], op.Signature)
				return tz.Some[tz.Signature](&sig), nil
			}
		}
	}
	sig, err := op.Signature.Signature()
	if err != nil {
		return tz.None[tz.Signature](), err
	}
	return tz.Some(sig), nil
}
