package proto_014_PtKathma

import (
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/proto_012_Psithaca"
)

type GroupContents interface {
	core.GroupContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[GroupContents]{
		Variants: encoding.Variants[GroupContents]{
			0: (*proto_012_Psithaca.OperationWithTooLargeMetadata[OperationContents])(nil),
			1: (*proto_012_Psithaca.OperationWithoutMetadata[OperationContents])(nil),
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
			0: (*proto_012_Psithaca.OperationWithOptionalMetadataWithMetadata[OperationContentsAndResult])(nil),
			1: (*proto_012_Psithaca.OperationWithOptionalMetadataWithoutMetadata[OperationContents])(nil),
		},
	})
}
