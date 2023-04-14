package operations

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
)

type Group struct {
	ChainID  *tz.ChainID
	Hash     *tz.BlockHash
	Branch   *tz.BlockHash `tz:"dyn"`
	Contents GroupContents `tz:"dyn"`
}

type SignedGroup interface {
	GetSignature() (tz.Signature, error)
}

type GroupContents interface {
	SignedGroup
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

type OperationWithOptionalMetadataContents interface {
	SignedGroup
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
