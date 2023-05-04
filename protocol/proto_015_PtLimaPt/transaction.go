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

type ToSmartRollup struct {
	ConsumedMilligas tz.BigUint
	TicketUpdates    []*TicketReceipt `tz:"dyn"`
}

func (*ToSmartRollup) TransactionResultDestination() {}

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

type OriginatedContract core.OriginatedContract
type ImplicitContract core.ImplicitContract

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
