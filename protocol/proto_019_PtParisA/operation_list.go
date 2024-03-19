package proto_019_PtParisA

import (
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/proto_016_PtMumbai"
)

type OperationWithOptionalMetadata = core.OperationWithOptionalMetadata[OperationWithOptionalMetadataContents]

type GroupContents interface {
	core.GroupContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[GroupContents]{
		Variants: encoding.Variants[GroupContents]{
			0: (*proto_016_PtMumbai.OperationWithTooLargeMetadata[OperationContents])(nil),
			1: (*proto_016_PtMumbai.OperationWithoutMetadata[OperationContents])(nil),
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
			0: (*proto_016_PtMumbai.OperationWithOptionalMetadataWithMetadata[OperationContentsAndResult])(nil),
			1: (*proto_016_PtMumbai.OperationWithOptionalMetadataWithoutMetadata[OperationContents])(nil),
		},
	})
}
