package proto_018_Proxford

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
)

type UnsignedOperation = UnsignedOperationImpl[OperationContents]
type SignedOperation = SignedOperationImpl[OperationContents]

type UnsignedOperationImpl[T core.OperationContents] struct {
	Branch   *tz.BlockHash `json:"branch"`
	Contents []T           `json:"contents"`
}

type SignedOperationImpl[T core.OperationContents] struct {
	UnsignedOperationImpl[T]
	Signature *tz.GenericSignature `json:"signature"`
}

func (*SignedOperationImpl[T]) RunOperationRequestContents() {}

func (op *SignedOperationImpl[T]) DecodeTZ(data []byte, ctx *encoding.Context) (rest []byte, err error) {
	if data, err = encoding.Decode(data, &op.Branch, encoding.Ctx(ctx)); err != nil {
		return nil, err
	}
	if len(data) < tz.GenericSignatureBytesLen {
		return nil, encoding.ErrBuffer(len(data))
	}
	tmp := data[:len(data)-tz.GenericSignatureBytesLen]
	data = data[len(data)-tz.GenericSignatureBytesLen:]
	if _, err := encoding.Decode(tmp, &op.Contents, encoding.Ctx(ctx)); err != nil {
		return nil, err
	}
	return encoding.Decode(data, &op.Signature, encoding.Ctx(ctx))
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
		UnsignedOperationImpl: *operation,
		Signature:             signature,
	}
}

type RunOperationRequest = RunOperationRequestImpl[RunOperationRequestContents]
type RunOperationRequestImpl[T RunOperationRequestContents] struct {
	Operation T           `json:"operation"`
	ChainID   *tz.ChainID `json:"chain_id"`
}

type RunOperationRequestContents interface {
	RunOperationRequestContents()
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[RunOperationRequestContents]{
		Variants: encoding.Variants[RunOperationRequestContents]{
			0: (*SignedOperation)(nil),
		},
	})
}
