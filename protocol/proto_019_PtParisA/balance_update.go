package proto_019_PtParisA

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/proto_012_Psithaca"
	"github.com/ecadlabs/gotez/v2/protocol/proto_016_PtMumbai"
	"github.com/ecadlabs/gotez/v2/protocol/proto_018_Proxford"
)

type BalanceUpdates struct {
	BalanceUpdates []*BalanceUpdate `tz:"dyn" json:"balance_updates"`
}

func (b *BalanceUpdates) GetBalanceUpdates() []core.BalanceUpdate {
	out := make([]core.BalanceUpdate, len(b.BalanceUpdates))
	for i, u := range b.BalanceUpdates {
		out[i] = u
	}
	return out
}

type BalanceUpdate struct {
	Contents BalanceUpdateContents    `json:"contents"`
	Change   int64                    `json:"change"`
	Origin   core.BalanceUpdateOrigin `json:"origin"`
}

func (b *BalanceUpdate) GetContents() core.BalanceUpdateContents { return b.Contents }
func (b *BalanceUpdate) GetChange() int64                        { return b.Change }
func (b *BalanceUpdate) GetOrigin() core.BalanceUpdateOrigin     { return b.Origin }

type BalanceUpdateContract = proto_012_Psithaca.BalanceUpdateContract
type BalanceUpdateLostEndorsingRewards = proto_012_Psithaca.BalanceUpdateLostEndorsingRewards
type BalanceUpdateCommitments = proto_012_Psithaca.BalanceUpdateCommitments
type BalanceUpdateBlockFees = proto_012_Psithaca.BalanceUpdateBlockFees
type BalanceUpdateNonceRevelationRewards = proto_012_Psithaca.BalanceUpdateNonceRevelationRewards
type BalanceUpdateEndorsingRewards = proto_012_Psithaca.BalanceUpdateEndorsingRewards
type BalanceUpdateBakingRewards = proto_012_Psithaca.BalanceUpdateBakingRewards
type BalanceUpdateBakingBonuses = proto_012_Psithaca.BalanceUpdateBakingBonuses
type BalanceUpdateStorageFees = proto_012_Psithaca.BalanceUpdateStorageFees
type BalanceUpdateDoubleSigningPunishments = proto_012_Psithaca.BalanceUpdateDoubleSigningPunishments
type BalanceUpdateLiquidityBakingSubsidies = proto_012_Psithaca.BalanceUpdateLiquidityBakingSubsidies
type BalanceUpdateBurned = proto_012_Psithaca.BalanceUpdateBurned
type BalanceUpdateBootstrap = proto_012_Psithaca.BalanceUpdateBootstrap
type BalanceUpdateInvoice = proto_012_Psithaca.BalanceUpdateInvoice
type BalanceUpdateInitialCommitments = proto_012_Psithaca.BalanceUpdateInitialCommitments
type BalanceUpdateMinted = proto_012_Psithaca.BalanceUpdateMinted
type BalanceUpdateFrozenBonds = proto_016_PtMumbai.BalanceUpdateFrozenBonds
type BalanceUpdateSmartRollupRefutationPunishments = proto_016_PtMumbai.BalanceUpdateSmartRollupRefutationPunishments
type BalanceUpdateSmartRollupRefutationRewards = proto_016_PtMumbai.BalanceUpdateSmartRollupRefutationRewards
type StakingDelegatorNumerator = proto_018_Proxford.StakingDelegatorNumerator
type StakingDelegateDenominator = proto_018_Proxford.StakingDelegateDenominator
type BalanceUpdateUnstakedDeposits = proto_018_Proxford.BalanceUpdateUnstakedDeposits

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateDeposits struct {
	Staker FrozenStaker `json:"staker"`
}

func (BalanceUpdateDeposits) BalanceUpdateCategory() string { return "deposits" }
func (BalanceUpdateDeposits) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateKindFreezer
}

type Staker = proto_018_Proxford.Staker

type FrozenStaker interface {
	Staker
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[FrozenStaker]{
		Variants: encoding.Variants[FrozenStaker]{
			0: StakerSingle{},
			1: StakerShared{},
			2: StakerBaker{},
			3: StakerBakerEdge{},
		},
	})
}

type StakerSingle = proto_018_Proxford.StakerSingle
type StakerShared = proto_018_Proxford.StakerShared
type StakerBaker = proto_018_Proxford.StakerBaker

type StakerBakerEdge struct {
	Baker tz.PublicKeyHash `json:"baker"`
}

func (StakerBakerEdge) StakerKind() string { return "baker_edge" }

type BalanceUpdateContents interface {
	core.BalanceUpdateContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[BalanceUpdateContents]{
		Variants: encoding.Variants[BalanceUpdateContents]{
			0:  (*BalanceUpdateContract)(nil),
			2:  BalanceUpdateBlockFees{},
			4:  (*BalanceUpdateDeposits)(nil),
			5:  BalanceUpdateNonceRevelationRewards{},
			7:  BalanceUpdateEndorsingRewards{},
			8:  BalanceUpdateBakingRewards{},
			9:  BalanceUpdateBakingBonuses{},
			11: BalanceUpdateStorageFees{},
			12: BalanceUpdateDoubleSigningPunishments{},
			13: (*BalanceUpdateLostEndorsingRewards)(nil),
			14: BalanceUpdateLiquidityBakingSubsidies{},
			15: BalanceUpdateBurned{},
			16: (*BalanceUpdateCommitments)(nil),
			17: BalanceUpdateBootstrap{},
			18: BalanceUpdateInvoice{},
			19: BalanceUpdateInitialCommitments{},
			20: BalanceUpdateMinted{},
			21: (*BalanceUpdateFrozenBonds)(nil),
			24: BalanceUpdateSmartRollupRefutationPunishments{},
			25: BalanceUpdateSmartRollupRefutationRewards{},
			26: (*BalanceUpdateUnstakedDeposits)(nil),
			27: StakingDelegatorNumerator{},
			28: StakingDelegateDenominator{},
		},
	})
}
