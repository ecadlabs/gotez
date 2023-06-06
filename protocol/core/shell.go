package core

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
)

type protocolVersionType int

var ProtocolVersionCtxKey protocolVersionType = 0

type ShellHeader struct {
	Level          int32
	Proto          Protocol
	Predecessor    *tz.BlockHash
	Timestamp      tz.Timestamp
	ValidationPass uint8
	OperationsHash *tz.OperationsHash
	Fitness        []byte `tz:"dyn"`
	Context        *tz.ContextHash
}

type BlockMetadataHeader struct {
	TestChainStatus        TestChainStatus
	MaxOperationsTTL       int32
	MaxOperationDataLength int32
	MaxBlockHeaderLength   int32
	MaxOperationListLength []*MaxOperationListLength `tz:"dyn,dyn"`
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

//json:status=forking
type TestChainStatusForking struct {
	Protocol   *tz.ProtocolHash `json:"protocol"`
	Expiration int64            `json:"expiration"`
}

func (TestChainStatusForking) TestChainStatus() string { return "forking" }

//json:status=running
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
	MaxSize int32
	MaxOp   tz.Option[int32]
}
