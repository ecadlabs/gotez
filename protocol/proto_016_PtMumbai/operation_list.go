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
			2: (*core.OperationWithOptionalMetadata[OperationWithOptionalMetadataContents])(nil),
		},
	})
}

type OperationWithTooLargeMetadata struct {
	OperationWithoutMetadata
}

type OperationWithoutMetadata struct {
	core.OperationWithoutMetadata[OperationContents]
}

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

type OperationWithOptionalMetadataContents interface {
	core.OperationWithOptionalMetadataContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[OperationWithOptionalMetadataContents]{
		Variants: encoding.Variants[OperationWithOptionalMetadataContents]{
			0: (*core.OperationWithOptionalMetadataWithMetadata[OperationContentsAndResult])(nil),
			1: (*core.OperationWithOptionalMetadataWithoutMetadata[OperationContents])(nil),
		},
	})
}
