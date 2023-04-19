package proto_015_PtLimaPt

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/core/expression"
)

type ZkRollupOrigination struct {
	ManagerOperation
	PublicParameters []byte              `tz:"dyn"`
	CircuitsInfo     []*CircuitsInfoElem `tz:"dyn"`
	InitState        []byte              `tz:"dyn"`
	NbOps            int32
}

func (*ZkRollupOrigination) OperationKind() string { return "zk_rollup_origination" }

type CircuitsInfoElem struct {
	Value string `tz:"dyn"`
	Tag   CircuitsInfoTag
}

type CircuitsInfoTag uint8

const (
	CircuitsInfoPublic CircuitsInfoTag = iota
	CircuitsInfoPrivate
	CircuitsInfoFee
)

type ZkRollupPublish struct {
	ManagerOperation
	ZkRollup *tz.ZkRollupAddress
	Op       []*ZkRollupOpElem `tz:"dyn"`
}

func (*ZkRollupPublish) OperationKind() string { return "zk_rollup_publish" }

type ZkRollupOpElem struct {
	Op     ZkRollupOp
	Ticket tz.Option1[ZkRollupTicket]
}

type ZkRollupOp struct {
	OpCode   int32
	Price    ZkRollupPrice
	L1Dst    tz.PublicKeyHash
	RollupID *tz.ZkRollupAddress
	Payload  []byte `tz:"dyn"`
}

type ZkRollupPrice struct {
	ID     *tz.ScriptExprHash
	Amount tz.BigInt
}

type ZkRollupTicket struct {
	Contents expression.Expression
	Ty       expression.Expression
	Ticketer core.ContractID
}
