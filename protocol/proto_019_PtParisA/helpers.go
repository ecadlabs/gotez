package proto_019_PtParisA

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/proto_018_Proxford"
)

type UnsignedOperation = proto_018_Proxford.UnsignedOperationImpl[OperationContents]
type SignedOperation = proto_018_Proxford.SignedOperationImpl[OperationContents]
type RunOperationRequest = proto_018_Proxford.RunOperationRequestImpl[RunOperationRequestContents]

func NewRunOperationRequest(op *SignedOperation, chain *tz.ChainID) *RunOperationRequest {
	return &RunOperationRequest{
		Operation: op,
		ChainID:   chain,
	}
}

func NewUnsignedOperation(branch *tz.BlockHash, contents []OperationContents) *UnsignedOperation {
	return &UnsignedOperation{
		Branch:   branch,
		Contents: contents,
	}
}

func NewSignedOperation(operation *UnsignedOperation, signature *tz.GenericSignature) *SignedOperation {
	return &SignedOperation{
		UnsignedOperationImpl: *operation,
		Signature:             signature,
	}
}

type RunOperationRequestContents interface {
	proto_018_Proxford.RunOperationRequestContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[RunOperationRequestContents]{
		Variants: encoding.Variants[RunOperationRequestContents]{
			0: (*SignedOperation)(nil),
		},
	})
}
