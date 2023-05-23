package proto_013_PtJakart

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/core/expression"
	"github.com/ecadlabs/gotez/protocol/proto_012_Psithaca"
)

type ManagerOperation = proto_012_Psithaca.ManagerOperation

type TransferTicket struct {
	ManagerOperation
	TicketContents expression.Expression `tz:"dyn"`
	TicketType     expression.Expression `tz:"dyn"`
	TicketTicketer core.ContractID
	TicketAmount   tz.BigUint
	Destination    core.ContractID
	Entrypoint     string `tz:"dyn"`
}

func (*TransferTicket) OperationKind() string { return "transfer_ticket" }

type TransferTicketResult interface {
	TransferTicketResult()
	core.OperationResult
}
