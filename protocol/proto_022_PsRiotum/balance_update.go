package proto_022_PsRiotum

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/proto_021_PsQuebec"
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

type BalanceUpdateContract = proto_021_PsQuebec.BalanceUpdateContract
type BalanceUpdateCommitments = proto_021_PsQuebec.BalanceUpdateCommitments
type BalanceUpdateBlockFees = proto_021_PsQuebec.BalanceUpdateBlockFees
type BalanceUpdateNonceRevelationRewards = proto_021_PsQuebec.BalanceUpdateNonceRevelationRewards
type BalanceUpdateBakingRewards = proto_021_PsQuebec.BalanceUpdateBakingRewards
type BalanceUpdateBakingBonuses = proto_021_PsQuebec.BalanceUpdateBakingBonuses
type BalanceUpdateStorageFees = proto_021_PsQuebec.BalanceUpdateStorageFees
type BalanceUpdateDoubleSigningPunishments = proto_021_PsQuebec.BalanceUpdateDoubleSigningPunishments
type BalanceUpdateLiquidityBakingSubsidies = proto_021_PsQuebec.BalanceUpdateLiquidityBakingSubsidies
type BalanceUpdateBurned = proto_021_PsQuebec.BalanceUpdateBurned
type BalanceUpdateBootstrap = proto_021_PsQuebec.BalanceUpdateBootstrap
type BalanceUpdateInvoice = proto_021_PsQuebec.BalanceUpdateInvoice
type BalanceUpdateInitialCommitments = proto_021_PsQuebec.BalanceUpdateInitialCommitments
type BalanceUpdateMinted = proto_021_PsQuebec.BalanceUpdateMinted
type BalanceUpdateSmartRollupRefutationPunishments = proto_021_PsQuebec.BalanceUpdateSmartRollupRefutationPunishments
type BalanceUpdateSmartRollupRefutationRewards = proto_021_PsQuebec.BalanceUpdateSmartRollupRefutationRewards
type StakingDelegatorNumerator = proto_021_PsQuebec.StakingDelegatorNumerator
type StakingDelegateDenominator = proto_021_PsQuebec.StakingDelegateDenominator
type BalanceUpdateUnstakedDeposits = proto_021_PsQuebec.BalanceUpdateUnstakedDeposits
type BalanceUpdateDeposits = proto_021_PsQuebec.BalanceUpdateDeposits
type BalanceUpdateAttestingRewards = proto_021_PsQuebec.BalanceUpdateAttestingRewards
type BalanceUpdateLostAttestingRewards = proto_021_PsQuebec.BalanceUpdateLostAttestingRewards
type BalanceUpdateFrozenBonds = proto_021_PsQuebec.BalanceUpdateFrozenBonds

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
			29: DALAttestingRewards{},
			30: LostDALAttestingRewards{},
		},
	})
}

type DALAttestingRewards struct{}

func (DALAttestingRewards) BalanceUpdateCategory() string {
	return "dal_attesting_rewards"
}

func (DALAttestingRewards) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateKindMinted
}

type LostDALAttestingRewards struct {
	Delegate tz.PublicKeyHash `json:"delegate"`
}

func (LostDALAttestingRewards) BalanceUpdateCategory() string {
	return "lost_dal_attesting_rewards"
}

func (LostDALAttestingRewards) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateKindBurned
}
