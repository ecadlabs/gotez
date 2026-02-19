package proto_016_PtMumbai

import (
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
)

type GroupContents interface {
	core.GroupContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[GroupContents]{
		Variants: encoding.Variants[GroupContents]{
			0: (*core.OperationWithTooLargeMetadata[OperationContents])(nil),
			1: (*core.OperationWithoutMetadata[OperationContents])(nil),
			2: (*core.OperationWithOptionalMetadata[OperationWithOptionalMetadataContents])(nil),
		},
	})
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
