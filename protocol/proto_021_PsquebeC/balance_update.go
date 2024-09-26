package proto_021_PsquebeC

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/proto_012_Psithaca"
	"github.com/ecadlabs/gotez/v2/protocol/proto_016_PtMumbai"
	"github.com/ecadlabs/gotez/v2/protocol/proto_018_Proxford"
	"github.com/ecadlabs/gotez/v2/protocol/proto_019_PtParisB"
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
type BalanceUpdateCommitments = proto_012_Psithaca.BalanceUpdateCommitments
type BalanceUpdateBlockFees = proto_012_Psithaca.BalanceUpdateBlockFees
type BalanceUpdateNonceRevelationRewards = proto_012_Psithaca.BalanceUpdateNonceRevelationRewards
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
type BalanceUpdateSmartRollupRefutationPunishments = proto_016_PtMumbai.BalanceUpdateSmartRollupRefutationPunishments
type BalanceUpdateSmartRollupRefutationRewards = proto_016_PtMumbai.BalanceUpdateSmartRollupRefutationRewards
type StakingDelegatorNumerator = proto_018_Proxford.StakingDelegatorNumerator
type StakingDelegateDenominator = proto_018_Proxford.StakingDelegateDenominator
type BalanceUpdateUnstakedDeposits = proto_018_Proxford.BalanceUpdateUnstakedDeposits
type BalanceUpdateDeposits = proto_019_PtParisB.BalanceUpdateDeposits

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateAttestingRewards struct{}

func (BalanceUpdateAttestingRewards) BalanceUpdateCategory() string {
	return "attesting_rewards"
}
func (BalanceUpdateAttestingRewards) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateKindMinted
}

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateLostAttestingRewards struct {
	Delegate      tz.PublicKeyHash `json:"delegate"`
	Participation bool             `json:"participation"`
	Revelation    bool             `json:"revelation"`
}

func (*BalanceUpdateLostAttestingRewards) BalanceUpdateCategory() string {
	return "lost_attesting_rewards"
}
func (*BalanceUpdateLostAttestingRewards) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateKindBurned
}

type BondID interface {
	proto_016_PtMumbai.BondID
}

type SmartRollupBondID struct {
	Address *tz.SmartRollupAddress `json:"address"`
}

func (SmartRollupBondID) BondID() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[BondID]{
		Variants: encoding.Variants[BondID]{
			1: SmartRollupBondID{},
		},
	})
}

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateFrozenBonds struct {
	Contract core.ContractID `json:"contract"`
	BondID   BondID          `json:"bond_id"`
}

func (*BalanceUpdateFrozenBonds) BalanceUpdateCategory() string { return "frozen_bonds" }
func (*BalanceUpdateFrozenBonds) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateKindFreezer
}

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
			7:  BalanceUpdateAttestingRewards{},
			8:  BalanceUpdateBakingRewards{},
			9:  BalanceUpdateBakingBonuses{},
			11: BalanceUpdateStorageFees{},
			12: BalanceUpdateDoubleSigningPunishments{},
			13: (*BalanceUpdateLostAttestingRewards)(nil),
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
