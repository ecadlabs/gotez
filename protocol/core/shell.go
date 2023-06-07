package core

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
)

type protocolVersionType int

var ProtocolVersionCtxKey protocolVersionType = 0

type ShellHeader struct {
	Level          int32              `json:"level"`
	Proto          Protocol           `json:"proto"`
	Predecessor    *tz.BlockHash      `json:"predecessor"`
	Timestamp      tz.Timestamp       `json:"timestamp"`
	ValidationPass uint8              `json:"validation_pass"`
	OperationsHash *tz.OperationsHash `json:"operations_hash"`
	Fitness        tz.Bytes           `tz:"dyn" json:"fitness"`
	Context        *tz.ContextHash    `json:"context"`
}

type BlockMetadataHeader struct {
	TestChainStatus        TestChainStatus           `json:"test_chain_status"`
	MaxOperationsTTL       int32                     `json:"max_operations_ttl"`
	MaxOperationDataLength int32                     `json:"max_operation_data_length"`
	MaxBlockHeaderLength   int32                     `json:"max_block_header_length"`
	MaxOperationListLength []*MaxOperationListLength `tz:"dyn,dyn" json:"max_operation_list_length"`
}

func (*BlockMetadataHeader) BlockMetadataContents() {}

type TestChainStatus interface {
	TestChainStatus() string
}

type TestChainStatusNotRunning struct{}

func (TestChainStatusNotRunning) TestChainStatus() string { return "not_running" }

func (t TestChainStatusNotRunning) MarshalText() (text []byte, err error) {
	return []byte(t.TestChainStatus()), nil
}

//json:status=TestChainStatus()
type TestChainStatusForking struct {
	Protocol   *tz.ProtocolHash `json:"protocol"`
	Expiration int64            `json:"expiration"`
}

func (TestChainStatusForking) TestChainStatus() string { return "forking" }

//json:status=TestChainStatus()
type TestChainStatusRunning struct {
	ChainID    *tz.ChainID      `json:"chain_id,omitempty"`
	Genesis    *tz.BlockHash    `json:"genesis,omitempty"`
	Protocol   *tz.ProtocolHash `json:"protocol"`
	Expiration int64            `json:"expiration"`
}

func (TestChainStatusRunning) TestChainStatus() string { return "running" }

func init() {
	encoding.RegisterEnum(&encoding.Enum[TestChainStatus]{
		Variants: encoding.Variants[TestChainStatus]{
			0: TestChainStatusNotRunning{},
			1: (*TestChainStatusForking)(nil),
			2: (*TestChainStatusRunning)(nil),
		},
	})
}

type MaxOperationListLength struct {
	MaxSize int32            `json:"max_size"`
	MaxOp   tz.Option[int32] `json:"max_op"`
}
