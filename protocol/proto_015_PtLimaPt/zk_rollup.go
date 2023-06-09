package proto_015_PtLimaPt

import (
	"strconv"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/core/expression"
)

//json:kind=OperationKind()
type ZkRollupOrigination struct {
	ManagerOperation
	PublicParameters tz.Bytes            `tz:"dyn" json:"public_parameters"`
	CircuitsInfo     []*CircuitsInfoElem `tz:"dyn" json:"circuits_info"`
	InitState        tz.Bytes            `tz:"dyn" json:"init_state"`
	NbOps            int32               `json:"nb_ops"`
}

func (*ZkRollupOrigination) OperationKind() string { return "zk_rollup_origination" }

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

//json:kind=OperationKind()
type ZkRollupPublish struct {
	ManagerOperation
	ZkRollup *tz.ZkRollupAddress `json:"zk_rollup"`
	Op       []*ZkRollupOpElem   `tz:"dyn" json:"op"`
}

func (*ZkRollupPublish) OperationKind() string { return "zk_rollup_publish" }

type ZkRollupOpElem struct {
	Op     ZkRollupOp                 `json:"op"`
	Ticket tz.Option1[ZkRollupTicket] `json:"ticket"`
}

type ZkRollupOp struct {
	OpCode   int32               `json:"op_code"`
	Price    ZkRollupPrice       `json:"price"`
	L1Dst    tz.PublicKeyHash    `json:"l1_dst"`
	RollupID *tz.ZkRollupAddress `json:"rollup_id"`
	Payload  tz.Bytes            `tz:"dyn" json:"payload"`
}

type ZkRollupPrice struct {
	ID     *tz.ScriptExprHash `json:"id"`
	Amount tz.BigInt          `json:"amount"`
}

type ZkRollupTicket struct {
	Contents expression.Expression `json:"contents"`
	Ty       expression.Expression `json:"ty"`
	Ticketer core.ContractID       `json:"ticketer"`
}

type ZkRollupOriginationContentsAndResult struct {
	ZkRollupOrigination
	Metadata ManagerMetadata[ZkRollupPublishResult] `json:"metadata"`
}

func (*ZkRollupOriginationContentsAndResult) OperationContentsAndResult() {}
func (op *ZkRollupOriginationContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type ZkRollupPublishResultContents struct {
	BalanceUpdates
	ConsumedMilligas tz.BigUint `json:"consumed_milligas"`
	Size             tz.BigInt  `json:"size"`
}

type ZkRollupPublishResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[ZkRollupPublishResult]{
		Variants: encoding.Variants[ZkRollupPublishResult]{
			0: (*core.OperationResultApplied[*ZkRollupPublishResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*ZkRollupPublishResultContents])(nil),
		},
	})
}

type ZkRollupPublishContentsAndResult struct {
	ZkRollupPublish
	Metadata ManagerMetadata[ZkRollupPublishResult] `json:"metadata"`
}

func (*ZkRollupPublishContentsAndResult) OperationContentsAndResult() {}
func (op *ZkRollupPublishContentsAndResult) GetMetadata() any {
	return &op.Metadata
}
