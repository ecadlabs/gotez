package proto

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
)

type protocolVersionType int

var ProtocolVersionCtxKey protocolVersionType = 0

type BlockHeader struct {
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

type TestChainStatusForking struct {
	Protocol   *tz.ProtocolHash
	Expiration int64
}

func (TestChainStatusForking) TestChainStatus() string { return "forking" }

type TestChainStatusRunning struct {
	ChainID    *tz.ChainID
	Genesis    *tz.BlockHash
	Protocol   *tz.ProtocolHash
	Expiration int64
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
