package proto_012_Psithaca

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

type RunOperationRequest struct {
	Operation SignedOperation `json:"operation"`
	ChainID   *tz.ChainID     `json:"chain_id"`
}
