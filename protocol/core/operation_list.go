package core

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
)

type GroupContents interface {
	GetSignature() (tz.Option[tz.Signature], error)
	GroupContents()
	Operations() []OperationContents
}

type OperationWithOptionalMetadataContents interface {
	GetSignature() (tz.Option[tz.Signature], error)
	OperationWithOptionalMetadataContents()
	Operations() []OperationContents
}

type OperationsList[T GroupContents] struct {
	Operations []*OperationsGroupImpl[T] `tz:"dyn,dyn" json:"operations"` // yes, twice
}

func (l *OperationsList[T]) GetGroups() []OperationsGroup {
	out := make([]OperationsGroup, len(l.Operations))
	for i, grp := range l.Operations {
		out[i] = grp
	}
	return out
}

type OperationsGroup interface {
	GetChainID() *tz.ChainID
	GetHash() *tz.OperationHash
	GetBranch() *tz.BlockHash
	GetContents() GroupContents
}

type OperationsGroupImpl[T GroupContents] struct {
	ChainID  *tz.ChainID       `json:"chain_id"`
	Hash     *tz.OperationHash `json:"hash"`
	Branch   *tz.BlockHash     `tz:"dyn" json:"branch"`
	Contents T                 `tz:"dyn" json:"contents"`
}

func (g *OperationsGroupImpl[T]) GetChainID() *tz.ChainID    { return g.ChainID }
func (g *OperationsGroupImpl[T]) GetHash() *tz.OperationHash { return g.Hash }
func (g *OperationsGroupImpl[T]) GetBranch() *tz.BlockHash   { return g.Branch }
func (g *OperationsGroupImpl[T]) GetContents() GroupContents { return g.Contents }

type OperationWithoutMetadata[T OperationContents] struct {
	Contents  []T                  `json:"contents"`
	Signature *tz.GenericSignature `json:"signature"`
}

func (op *OperationWithoutMetadata[T]) DecodeTZ(data []byte, ctx *encoding.Context) (rest []byte, err error) {
	if len(data) < tz.GenericSignatureBytesLen {
		return nil, encoding.ErrBuffer{tz.GenericSignatureBytesLen, len(data)}
	}
	tmp := data[:len(data)-tz.GenericSignatureBytesLen]
	data = data[len(data)-tz.GenericSignatureBytesLen:]
	if _, err := encoding.Decode(tmp, &op.Contents, encoding.Ctx(ctx)); err != nil {
		return nil, err
	}
	return encoding.Decode(data, &op.Signature, encoding.Ctx(ctx))
}

func (ops *OperationWithoutMetadata[T]) Operations() []OperationContents {
	out := make([]OperationContents, len(ops.Contents))
	for i, op := range ops.Contents {
		out[i] = op
	}
	return out
}

func (*OperationWithoutMetadata[T]) GroupContents() {}

func (op *OperationWithoutMetadata[T]) GetSignature() (tz.Option[tz.Signature], error) {
	if len(op.Contents) != 0 {
		var last OperationContents = op.Contents[len(op.Contents)-1]
		if prefix, ok := last.(*SignaturePrefix); ok {
			switch p := prefix.SignaturePrefix.(type) {
			case *BLSSignaturePrefix:
				var sig tz.BLSSignature
				copy(sig[:], p[:])
				copy(sig[:len(p)], op.Signature[:])
				return tz.Some[tz.Signature](&sig), nil
			case *MLDSA44SignaturePrefix:
				var sig tz.MLDSA44Signature
				copy(sig[:], p[:])
				copy(sig[:len(p)], op.Signature[:])
				return tz.Some[tz.Signature](&sig), nil
			}
		}
	}
	return tz.Some[tz.Signature](op.Signature), nil
}

type OperationWithTooLargeMetadata[T OperationContents] struct {
	OperationWithoutMetadata[T]
}

type OperationWithOptionalMetadata[T OperationWithOptionalMetadataContents] struct {
	Contents T `json:"contents"`
}

func (ops OperationWithOptionalMetadata[T]) Operations() []OperationContents {
	return ops.Contents.Operations()
}

func (op OperationWithOptionalMetadata[T]) GetSignature() (tz.Option[tz.Signature], error) {
	return op.Contents.GetSignature()
}

func (OperationWithOptionalMetadata[T]) GroupContents() {}

type OperationWithOptionalMetadataWithMetadata[T OperationContentsAndResult] = OperationWithOptionalMetadataWithoutMetadata[T]

type OperationWithOptionalMetadataWithoutMetadata[T OperationContents] struct {
	Contents  []T             `tz:"dyn" json:"contents"`
	Signature tz.AnySignature `json:"signature"`
}

func (ops *OperationWithOptionalMetadataWithoutMetadata[T]) Operations() []OperationContents {
	out := make([]OperationContents, len(ops.Contents))
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
		var last OperationContents = op.Contents[len(op.Contents)-1]
		if prefix, ok := last.(*SignaturePrefix); ok {
			switch p := prefix.SignaturePrefix.(type) {
			case *BLSSignaturePrefix:
				var sig tz.BLSSignature
				copy(sig[:], p[:])
				copy(sig[:len(p)], op.Signature[:])
				return tz.Some[tz.Signature](&sig), nil
			case *MLDSA44SignaturePrefix:
				var sig tz.MLDSA44Signature
				copy(sig[:], p[:])
				copy(sig[:len(p)], op.Signature[:])
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
