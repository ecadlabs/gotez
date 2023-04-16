package proto_009_PsFLoren

import (
	tz "github.com/ecadlabs/gotez"
)

type BalanceUpdateContract struct {
	Contract tz.ContractID
}

func (*BalanceUpdateContract) BalanceUpdateKind() string { return "contract" }

type BalanceUpdateBlockFees struct{}

func (BalanceUpdateBlockFees) BalanceUpdateKind() string { return "block_fees" }
