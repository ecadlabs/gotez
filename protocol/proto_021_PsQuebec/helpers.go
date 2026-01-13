package proto_021_PsQuebec

import (
	tz "github.com/ecadlabs/gotez/v2"
)

type UnsignedOperation struct {
	Branch   *tz.BlockHash       `json:"branch"`
	Contents []OperationContents `json:"contents"`
}
type SignedOperation struct {
	UnsignedOperation
	Signature *tz.GenericSignature `json:"signature"`
}

type RunOperationRequest struct {
	Operation *SignedOperation `json:"operation"`
	ChainID   *tz.ChainID      `json:"chain_id"`
}

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
		UnsignedOperation: *operation,
		Signature:         signature,
	}
}
