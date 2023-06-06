package core

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
)

type GroupContents interface {
	Signed
	GroupContents()
	Operations() []OperationContents
}

type OperationWithOptionalMetadataContents interface {
	Signed
	OperationWithOptionalMetadataContents()
	Operations() []OperationContents
}

type OperationsList[T GroupContents] struct {
	Operations []*OperationsGroupImpl[T] `tz:"dyn,dyn"` // yes, twice
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
	GetHash() *tz.OperationsHash
	GetBranch() *tz.BlockHash
	GetContents() GroupContents
}

type OperationsGroupImpl[T GroupContents] struct {
	ChainID  *tz.ChainID
	Hash     *tz.OperationsHash
	Branch   *tz.BlockHash `tz:"dyn"`
	Contents T             `tz:"dyn"`
}

func (g *OperationsGroupImpl[T]) GetChainID() *tz.ChainID     { return g.ChainID }
func (g *OperationsGroupImpl[T]) GetHash() *tz.OperationsHash { return g.Hash }
func (g *OperationsGroupImpl[T]) GetBranch() *tz.BlockHash    { return g.Branch }
func (g *OperationsGroupImpl[T]) GetContents() GroupContents  { return g.Contents }

type OperationWithTooLargeMetadata[T OperationContents] struct {
	OperationWithoutMetadata[T]
}

type OperationWithoutMetadata[T OperationContents] struct {
	Contents  []T
	Signature *tz.GenericSignature // takes the rest, see below
}

func (op *OperationWithoutMetadata[T]) DecodeTZ(data []byte, ctx *encoding.Context) (rest []byte, err error) {
	if len(data) < tz.GenericSignatureBytesLen {
		return nil, encoding.ErrBuffer(len(data))
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
func (op *OperationWithoutMetadata[T]) GetSignature() (tz.Signature, error) {
	return op.Signature, nil
}

type OperationWithOptionalMetadata[T OperationWithOptionalMetadataContents] struct {
	Contents T
}

func (ops *OperationWithOptionalMetadata[T]) Operations() []OperationContents {
	return ops.Contents.Operations()
}

func (op *OperationWithOptionalMetadata[T]) GetSignature() (tz.Signature, error) {
	return op.Contents.GetSignature()
}

func (*OperationWithOptionalMetadata[T]) GroupContents() {}

type OperationWithOptionalMetadataWithMetadata[T OperationContentsAndResult] struct {
	Contents  []T `tz:"dyn"`
	Signature tz.AnySignature
}

func (ops *OperationWithOptionalMetadataWithMetadata[T]) Operations() []OperationContents {
	out := make([]OperationContents, len(ops.Contents))
	for i, op := range ops.Contents {
		out[i] = op
	}
	return out
}

func (*OperationWithOptionalMetadataWithMetadata[T]) OperationWithOptionalMetadataContents() {}
func (op *OperationWithOptionalMetadataWithMetadata[T]) GetSignature() (tz.Signature, error) {
	return op.Signature.Signature()
}

type OperationWithOptionalMetadataWithoutMetadata[T OperationContents] struct {
	Contents  []T `tz:"dyn"`
	Signature tz.AnySignature
}

func (ops *OperationWithOptionalMetadataWithoutMetadata[T]) Operations() []OperationContents {
	out := make([]OperationContents, len(ops.Contents))
	for i, op := range ops.Contents {
		out[i] = op
	}
	return out
}

func (*OperationWithOptionalMetadataWithoutMetadata[T]) OperationWithOptionalMetadataContents() {}
func (op *OperationWithOptionalMetadataWithoutMetadata[T]) GetSignature() (tz.Signature, error) {
	return op.Signature.Signature()
}
