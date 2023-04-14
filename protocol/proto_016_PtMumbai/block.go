package proto_016_PtMumbai

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/proto_016_PtMumbai/operations"
	"github.com/ecadlabs/gotez/protocol/shell"
)

type ProtocolBlockHeader struct {
	PayloadHash               *tz.BlockPayloadHash
	PayloadRound              int32
	ProofOfWorkNonce          *[tz.ProofOfWorkNonceBytesLen]byte
	SeedNonceHash             tz.Option[*tz.CycleNonceHash]
	LiquidityBakingToggleVote uint8
	Signature                 tz.AnySignature
}

func (*ProtocolBlockHeader) ProtocolBlockHeader() {}

type UnsignedBlockHeader struct {
	shell.BlockHeader
	PayloadHash               *tz.BlockPayloadHash
	PayloadRound              int32
	ProofOfWorkNonce          *[tz.ProofOfWorkNonceBytesLen]byte
	SeedNonceHash             tz.Option[*tz.CycleNonceHash]
	LiquidityBakingToggleVote uint8
}

type BlockHeader struct {
	UnsignedBlockHeader
	Signature tz.AnySignature
}

type BlockInfo struct {
	Contents BlockInfoContents `tz:"dyn"`
}

type BlockInfoContents struct {
	ChainID    *tz.ChainID
	Hash       *tz.BlockHash
	Header     BlockHeader `tz:"dyn"`
	Metadata   tz.Option[BlockMetadata]
	Operations []OperationsList `tz:"dyn"`
}

type BlockProtocolData struct {
	Metadata   tz.Option[BlockMetadata]
	Operations []OperationsList `tz:"dyn"`
}

func (*BlockProtocolData) BlockProtocolData() {}

type BlockMetadata struct {
	BlockMetadataContents `tz:"dyn"`
}

type BlockMetadataContents struct {
	TestChainStatus           TestChainStatus
	MaxOperationsTTL          int32
	MaxOperationDataLength    int32
	MaxBlockHeaderLength      int32
	MaxOperationListLength    []*MaxOperationListLength `tz:"dyn,dyn"`
	Proposer                  tz.PublicKeyHash
	Baker                     tz.PublicKeyHash
	LevelInfo                 LevelInfo
	VotingPeriodInfo          VotingPeriodInfo
	NonceHash                 tz.Option1[*tz.CycleNonceHash]
	Deactivated               []tz.PublicKeyHash          `tz:"dyn"`
	BalanceUpdates            []*operations.BalanceUpdate `tz:"dyn"`
	LiquidityBakingEscapeEMA  int32
	ImplicitOperationsResults []operations.SuccessfulManagerOperationResult `tz:"dyn"`
	ProposerConsensusKey      tz.PublicKeyHash
	BakerConsensusKey         tz.PublicKeyHash
	ConsumedMilligas          tz.BigUint
	DALAttestation            tz.Option[tz.BigInt]
}

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

type LevelInfo struct {
	Level              int32
	LevelPosition      int32
	Cycle              int32
	CyclePosition      int32
	ExpectedCommitment bool
}

type VotingPeriodInfo struct {
	VotingPeriod VotingPeriod
	Position     int32
	Remaining    int32
}

type VotingPeriod struct {
	Index         int32
	Kind          VotingPeriodKind
	StartPosition int32
}

type VotingPeriodKind uint8

const (
	VotingPeriodProposal VotingPeriodKind = iota
	VotingPeriodExploration
	VotingPeriodCooldown
	VotingPeriodPromotion
	VotingPeriodAdoption
)

type OperationsList struct {
	Operations []*operations.Group `tz:"dyn,dyn"` // yes, twice
}
