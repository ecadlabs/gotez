package proto_018_Proxford

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
)

type UnsignedOperation struct {
	Branch   *tz.BlockHash       `json:"branch"`
	Contents []OperationContents `json:"contents"`
}

type SignedOperation struct {
	UnsignedOperation
	Signature *tz.GenericSignature `json:"signature"`
}

func (*SignedOperation) RunOperationRequestContents() {}

func (op *SignedOperation) DecodeTZ(data []byte, ctx *encoding.Context) (rest []byte, err error) {
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

type RunOperationRequest struct {
	Operation RunOperationRequestContents `json:"operation"`
	ChainID   *tz.ChainID                 `json:"chain_id"`
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
