package proto_015_PtLimaPt

import (
	"strconv"

	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/core/expression"
)

type ZkRollupOrigination struct {
	ManagerOperation
	PublicParameters tz.Bytes            `tz:"dyn" json:"public_parameters"`
	CircuitsInfo     []*CircuitsInfoElem `tz:"dyn" json:"circuits_info"`
	InitState        tz.Bytes            `tz:"dyn" json:"init_state"`
	NbOps            int32               `json:"nb_ops"`
}

func (*ZkRollupOrigination) OperationKind() string        { return "zk_rollup_origination" }
func (op *ZkRollupOrigination) Operation() core.Operation { return op }

type CircuitsInfoElem struct {
	Value string          `tz:"dyn" json:"value"`
	Tag   CircuitsInfoTag `json:"tag"`
}

type CircuitsInfoTag uint8

func (c CircuitsInfoTag) String() string {
	switch c {
	case CircuitsInfoPublic:
		return "public"
	case CircuitsInfoPrivate:
		return "private"
	case CircuitsInfoFee:
		return "fee"
	default:
		return strconv.FormatInt(int64(c), 10)
	}
}

func (c CircuitsInfoTag) MarshalText() (text []byte, err error) { return []byte(c.String()), nil }

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

func (*ZkRollupPublish) OperationKind() string        { return "zk_rollup_publish" }
func (op *ZkRollupPublish) Operation() core.Operation { return op }

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
