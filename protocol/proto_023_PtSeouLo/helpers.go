package proto_023_PtSeouLo

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/protocol/proto_018_Proxford"
	"github.com/ecadlabs/gotez/v2/protocol/proto_019_PtParisB"
)

type UnsignedOperation = proto_018_Proxford.UnsignedOperationImpl[OperationContents]
type SignedOperation = proto_018_Proxford.SignedOperationImpl[OperationContents]
type RunOperationRequest = proto_018_Proxford.RunOperationRequestImpl[RunOperationRequestContents]
type RunOperationRequestContents = proto_019_PtParisB.RunOperationRequestContents

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
