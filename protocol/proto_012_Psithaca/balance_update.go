package proto_012_Psithaca

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
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

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateContract struct {
	Contract core.ContractID `json:"contract"`
}

func (BalanceUpdateContract) BalanceUpdateCategory() string { return "contract" }
func (BalanceUpdateContract) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateContract
}

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateBlockFees struct{}

func (BalanceUpdateBlockFees) BalanceUpdateCategory() string { return "block_fees" }
func (BalanceUpdateBlockFees) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateAccumulator
}

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateDeposits struct {
	Delegate tz.PublicKeyHash `json:"delegate"`
}

func (BalanceUpdateDeposits) BalanceUpdateCategory() string { return "deposits" }
func (BalanceUpdateDeposits) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateFreezer
}

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateNonceRevelationRewards struct{}

func (BalanceUpdateNonceRevelationRewards) BalanceUpdateCategory() string {
	return "nonce_revelation_rewards"
}
func (BalanceUpdateNonceRevelationRewards) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateMinted
}

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateDoubleSigningEvidenceRewards struct{}

func (BalanceUpdateDoubleSigningEvidenceRewards) BalanceUpdateCategory() string {
	return "double_signing_evidence_rewards"
}
func (BalanceUpdateDoubleSigningEvidenceRewards) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateMinted
}

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateEndorsingRewards struct{}

func (BalanceUpdateEndorsingRewards) BalanceUpdateCategory() string {
	return "endorsing_rewards"
}
func (BalanceUpdateEndorsingRewards) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateMinted
}

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateBakingRewards struct{}

func (BalanceUpdateBakingRewards) BalanceUpdateCategory() string {
	return "baking_rewards"
}
func (BalanceUpdateBakingRewards) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateMinted
}

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateBakingBonuses struct{}

func (BalanceUpdateBakingBonuses) BalanceUpdateCategory() string {
	return "baking_bonuses"
}
func (BalanceUpdateBakingBonuses) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateMinted
}

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateStorageFees struct{}

func (BalanceUpdateStorageFees) BalanceUpdateCategory() string {
	return "storage_fees"
}
func (BalanceUpdateStorageFees) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateBurned
}

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateDoubleSigningPunishments struct{}

func (BalanceUpdateDoubleSigningPunishments) BalanceUpdateCategory() string {
	return "double_signing_punishments"
}
func (BalanceUpdateDoubleSigningPunishments) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateBurned
}

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateLiquidityBakingSubsidies struct{}

func (BalanceUpdateLiquidityBakingSubsidies) BalanceUpdateCategory() string {
	return "liquidity_baking_subsidies"
}
func (BalanceUpdateLiquidityBakingSubsidies) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateMinted
}

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateBurned struct{}

func (BalanceUpdateBurned) BalanceUpdateCategory() string {
	return "burned"
}
func (BalanceUpdateBurned) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateBurned
}

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateBootstrap struct{}

func (BalanceUpdateBootstrap) BalanceUpdateCategory() string {
	return "bootstrap"
}
func (BalanceUpdateBootstrap) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateMinted
}

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateInvoice struct{}

func (BalanceUpdateInvoice) BalanceUpdateCategory() string {
	return "invoice"
}
func (BalanceUpdateInvoice) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateMinted
}

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateInitialCommitments struct{}

func (BalanceUpdateInitialCommitments) BalanceUpdateCategory() string {
	return "initial_commitments"
}
func (BalanceUpdateInitialCommitments) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateMinted
}

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateMinted struct{}

func (BalanceUpdateMinted) BalanceUpdateCategory() string {
	return "minted"
}
func (BalanceUpdateMinted) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateMinted
}

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateLostEndorsingRewards struct {
	Delegate      tz.PublicKeyHash `json:"delegate"`
	Participation bool             `json:"participation"`
	Revelation    bool             `json:"revelation"`
}

func (*BalanceUpdateLostEndorsingRewards) BalanceUpdateCategory() string {
	return "lost_endorsing_rewards"
}
func (*BalanceUpdateLostEndorsingRewards) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateBurned
}

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateCommitments struct {
	Committer *tz.BlindedPublicKeyHash `json:"committer"`
}

func (BalanceUpdateCommitments) BalanceUpdateCategory() string { return "commitments" }
func (BalanceUpdateCommitments) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateCommitment
}

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateLegacyRewards struct {
	Delegate tz.PublicKeyHash `json:"delegate"`
	Cycle    int32            `json:"cycle"`
}

func (*BalanceUpdateLegacyRewards) BalanceUpdateCategory() string { return "legacy_rewards" }
func (*BalanceUpdateLegacyRewards) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateFreezer
}

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateLegacyDeposits BalanceUpdateLegacyRewards

func (*BalanceUpdateLegacyDeposits) BalanceUpdateCategory() string { return "legacy_deposits" }
func (*BalanceUpdateLegacyDeposits) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateFreezer
}

//json:category=BalanceUpdateCategory(),kind=BalanceUpdateKind()
type BalanceUpdateLegacyFees BalanceUpdateLegacyRewards

func (*BalanceUpdateLegacyFees) BalanceUpdateCategory() string { return "legacy_fees" }
func (*BalanceUpdateLegacyFees) BalanceUpdateKind() core.BalanceUpdateKind {
	return core.BalanceUpdateFreezer
}

type BalanceUpdateContents interface {
	core.BalanceUpdateContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[BalanceUpdateContents]{
		Variants: encoding.Variants[BalanceUpdateContents]{
			0:  BalanceUpdateContract{},
			1:  (*BalanceUpdateLegacyRewards)(nil),
			2:  BalanceUpdateBlockFees{},
			3:  (*BalanceUpdateLegacyDeposits)(nil),
			4:  BalanceUpdateDeposits{},
			5:  BalanceUpdateNonceRevelationRewards{},
			6:  BalanceUpdateDoubleSigningEvidenceRewards{},
			7:  BalanceUpdateEndorsingRewards{},
			8:  BalanceUpdateBakingRewards{},
			9:  BalanceUpdateBakingBonuses{},
			10: (*BalanceUpdateLegacyFees)(nil),
			11: BalanceUpdateStorageFees{},
			12: BalanceUpdateDoubleSigningPunishments{},
			13: (*BalanceUpdateLostEndorsingRewards)(nil),
			14: BalanceUpdateLiquidityBakingSubsidies{},
			15: BalanceUpdateBurned{},
			16: BalanceUpdateCommitments{},
			17: BalanceUpdateBootstrap{},
			18: BalanceUpdateInvoice{},
			19: BalanceUpdateInitialCommitments{},
			20: BalanceUpdateMinted{},
		},
	})
}
