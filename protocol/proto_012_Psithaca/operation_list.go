package proto_012_Psithaca

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
	return op.Signature, nil
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
