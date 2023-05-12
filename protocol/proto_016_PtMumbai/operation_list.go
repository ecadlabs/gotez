package proto_016_PtMumbai

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core"
)

type GroupContents interface {
	core.GroupContents
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

type OperationWithOptionalMetadataContents interface {
	core.OperationWithOptionalMetadataContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[OperationWithOptionalMetadataContents]{
		Variants: encoding.Variants[OperationWithOptionalMetadataContents]{
			0: (*OperationWithOptionalMetadataWithMetadata)(nil),
			1: (*OperationWithOptionalMetadataWithoutMetadata)(nil),
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
		if prefix, ok := op.Contents[len(op.Contents)-1].(*SignaturePrefix); ok {
			if blsPrefix, ok := prefix.SignaturePrefix.(*BLSSignaturePrefix); ok {
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
