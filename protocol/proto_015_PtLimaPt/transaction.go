package proto_015_PtLimaPt

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/expression"
	"github.com/ecadlabs/gotez/protocol/proto_005_PsBABY5H"
)

type Transaction = proto_005_PsBABY5H.Transaction
type Parameters = proto_005_PsBABY5H.Parameters

type TicketReceipt struct {
	TicketToken TicketToken
	Updates     []*TicketReceiptUpdate `tz:"dyn"`
}

type TicketToken struct {
	Ticketer    tz.ContractID
	ContentType expression.Expression
	Content     expression.Expression
}

type TicketReceiptUpdate struct {
	Account TransactionDestination
	Amount  tz.BigInt
}

type ToSmartRollup struct {
	ConsumedMilligas tz.BigUint
	TicketUpdates    []*TicketReceipt `tz:"dyn"`
}

func (*ToSmartRollup) TransactionResultContents() {}

type TransactionResult interface {
	TransactionResult()
	OperationResult
}

type TransactionDestination interface {
	TransactionDestination()
}

type TxRollupDestination struct {
	*tz.TXRollupAddress
	Padding uint8
}

type SmartRollupDestination struct {
	*tz.SmartRollupAddress
	Padding uint8
}

type ZkRollupDestination struct {
	*tz.ZkRollupAddress
	Padding uint8
}

type OriginatedContract tz.OriginatedContract
type ImplicitContract tz.ImplicitContract

func (*TxRollupDestination) TransactionDestination()    {}
func (*SmartRollupDestination) TransactionDestination() {}
func (*OriginatedContract) TransactionDestination()     {}
func (*ImplicitContract) TransactionDestination()       {}
func (*ZkRollupDestination) TransactionDestination()    {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[TransactionDestination]{
		Variants: encoding.Variants[TransactionDestination]{
			0: (*ImplicitContract)(nil),
			1: (*OriginatedContract)(nil),
			2: (*TxRollupDestination)(nil),
			3: (*SmartRollupDestination)(nil),
			4: (*ZkRollupDestination)(nil),
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
