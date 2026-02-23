package core

import (
	"encoding/binary"
	"errors"
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
)

type ShellHeader struct {
	Level          int32              `json:"level"`
	Proto          uint8              `json:"proto"`
	Predecessor    *tz.BlockHash      `json:"predecessor"`
	Timestamp      tz.Timestamp       `json:"timestamp"`
	ValidationPass uint8              `json:"validation_pass"`
	OperationsHash *tz.OperationsHash `json:"operations_hash"`
	Fitness        tz.Bytes           `tz:"dyn" json:"fitness"`
	Context        *tz.ContextHash    `json:"context"`
}

// GetRoundFromTenderbakeBlock extracts the round from a Tenderbake block fitness
func GetRoundFromTenderbakeBlock(data tz.Bytes) (uint32, error) {
	/* FITNESS=
	   (<fitness_length(4)> not in gotez)
	   <version_len(4)><version(1)>
	   <level_len(4)><level(4)>
	   <locked_round_len(4)><locked_round(0 OR 4)>
	   <predecessor_round_len(4)><predecessor_round(4)>
	   <round_len(4)><round(4)> */

	if len(data) < 4 {
		return 0, errors.New("data too short to extract round")
	}
	// The fitness data has been stripped from its prefixed length
	// The round value is always the 4 last bytes
	round := binary.BigEndian.Uint32(data[len(data)-4:])

	return round, nil
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
