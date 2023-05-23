package proto_015_PtLimaPt

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/core/expression"
	"github.com/ecadlabs/gotez/protocol/proto_012_Psithaca"
)

type Transaction = proto_012_Psithaca.Transaction
type Parameters = proto_012_Psithaca.Parameters

type TransactionResultDestination interface {
	TransactionResultDestination()
}

type TicketReceipt struct {
	TicketToken TicketToken
	Updates     []*TicketReceiptUpdate `tz:"dyn"`
}

type TicketToken struct {
	Ticketer    core.ContractID
	ContentType expression.Expression
	Content     expression.Expression
}

type TicketReceiptUpdate struct {
	Account TransactionDestination
	Amount  tz.BigInt
}

type ToScRollup struct {
	ConsumedMilligas tz.BigUint
	TicketUpdates    []*TicketReceipt `tz:"dyn"`
}

func (*ToScRollup) TransactionResultDestination() {}

type TransactionDestination interface {
	TransactionDestination()
}

type TxRollupDestination struct {
	*tz.TXRollupAddress
	Padding uint8
}

type ScRollupDestination struct {
	*tz.ScRollupAddress
	Padding uint8
}

type ZkRollupDestination struct {
	*tz.ZkRollupAddress
	Padding uint8
}

type OriginatedContract core.OriginatedContract
type ImplicitContract core.ImplicitContract

func (*TxRollupDestination) TransactionDestination() {}
func (*ScRollupDestination) TransactionDestination() {}
func (*OriginatedContract) TransactionDestination()  {}
func (*ImplicitContract) TransactionDestination()    {}
func (*ZkRollupDestination) TransactionDestination() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[TransactionDestination]{
		Variants: encoding.Variants[TransactionDestination]{
			0: (*ImplicitContract)(nil),
			1: (*OriginatedContract)(nil),
			2: (*TxRollupDestination)(nil),
			3: (*ScRollupDestination)(nil),
			4: (*ZkRollupDestination)(nil),
		},
	})
}

type TransactionResultContents struct {
	Result TransactionResultDestination
}

func (TransactionResultContents) SuccessfulManagerOperationResult() {}
func (TransactionResultContents) OperationKind() string             { return "transaction" }

type TransactionContentsAndResult struct {
	Transaction
	Metadata ManagerMetadata[TransactionResult]
}

func (*TransactionContentsAndResult) OperationContentsAndResult() {}
func (op *TransactionContentsAndResult) OperationContents() core.OperationContents {
	return &op.Transaction
}

type TransactionResultApplied struct {
	core.OperationResultApplied[TransactionResultContents]
}

func (*TransactionResultApplied) TransactionResult() {}

type TransactionResultBacktracked struct {
	core.OperationResultBacktracked[TransactionResultContents]
}

func (*TransactionResultBacktracked) TransactionResult() {}

type TransactionResultFailed struct{ core.OperationResultFailed }

func (*TransactionResultFailed) TransactionResult() {}

type TransactionResultSkipped struct{ core.OperationResultSkipped }

func (*TransactionResultSkipped) TransactionResult() {}

type TransactionResult interface {
	TransactionResult()
	core.OperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[TransactionResult]{
		Variants: encoding.Variants[TransactionResult]{
			0: (*TransactionResultApplied)(nil),
			1: (*TransactionResultFailed)(nil),
			2: (*TransactionResultSkipped)(nil),
			3: (*TransactionResultBacktracked)(nil),
		},
	})
}

type TransactionInternalOperationResult struct {
	Source      TransactionDestination
	Nonce       uint16
	Amount      tz.BigUint
	Destination TransactionDestination
	Parameters  tz.Option[Parameters]
	Result      TransactionResult
}

func (*TransactionInternalOperationResult) InternalOperationResult() {}
func (*TransactionInternalOperationResult) OperationKind() string    { return "transaction" }
