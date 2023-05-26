package proto_012_Psithaca

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core"
)

type BalanceUpdateOrigin uint8

const (
	BalanceUpdateOriginBlockApplication BalanceUpdateOrigin = iota
	BalanceUpdateOriginProtocolMigration
	BalanceUpdateOriginSubsidy
	BalanceUpdateOriginSimulation
)

type BalanceUpdate struct {
	Kind   BalanceUpdateKind
	Change int64
	Origin BalanceUpdateOrigin
}

type BalanceUpdateContract struct {
	Contract core.ContractID
}

func (*BalanceUpdateContract) BalanceUpdateKind() string { return "contract" }

type BalanceUpdateBlockFees struct{}

func (BalanceUpdateBlockFees) BalanceUpdateKind() string { return "block_fees" }

type BalanceUpdateDeposits struct {
	Delegate tz.PublicKeyHash
}

func (*BalanceUpdateDeposits) BalanceUpdateKind() string { return "deposits" }

type BalanceUpdateNonceRevelationRewards struct{}

func (BalanceUpdateNonceRevelationRewards) BalanceUpdateKind() string {
	return "nonce_revelation_rewards"
}

type BalanceUpdateDoubleSigningEvidenceRewards struct{}

func (BalanceUpdateDoubleSigningEvidenceRewards) BalanceUpdateKind() string {
	return "double_signing_evidence_rewards"
}

type BalanceUpdateEndorsingRewards struct{}

func (BalanceUpdateEndorsingRewards) BalanceUpdateKind() string {
	return "endorsing_rewards"
}

type BalanceUpdateBakingRewards struct{}

func (BalanceUpdateBakingRewards) BalanceUpdateKind() string {
	return "baking_rewards"
}

type BalanceUpdateBakingBonuses struct{}

func (BalanceUpdateBakingBonuses) BalanceUpdateKind() string {
	return "baking_bonuses"
}

type BalanceUpdateStorageFees struct{}

func (BalanceUpdateStorageFees) BalanceUpdateKind() string {
	return "storage_fees"
}

type BalanceUpdateDoubleSigningPunishments struct{}

func (BalanceUpdateDoubleSigningPunishments) BalanceUpdateKind() string {
	return "double_signing_punishments"
}

type BalanceUpdateLiquidityBakingSubsidies struct{}

func (BalanceUpdateLiquidityBakingSubsidies) BalanceUpdateKind() string {
	return "liquidity_baking_subsidies"
}

type BalanceUpdateBurned struct{}

func (BalanceUpdateBurned) BalanceUpdateKind() string {
	return "burned"
}

type BalanceUpdateBootstrap struct{}

func (BalanceUpdateBootstrap) BalanceUpdateKind() string {
	return "bootstrap"
}

type BalanceUpdateInvoice struct{}

func (BalanceUpdateInvoice) BalanceUpdateKind() string {
	return "invoice"
}

type BalanceUpdateInitialCommitments struct{}

func (BalanceUpdateInitialCommitments) BalanceUpdateKind() string {
	return "initial_commitments"
}

type BalanceUpdateMinted struct{}

func (BalanceUpdateMinted) BalanceUpdateKind() string {
	return "minted"
}

type BalanceUpdateLostEndorsingRewards struct {
	Delegate      tz.PublicKeyHash
	Participation bool
	Revelation    bool
}

func (*BalanceUpdateLostEndorsingRewards) BalanceUpdateKind() string { return "lost_endorsing_rewards" }

type BalanceUpdateCommitments struct {
	Committer *tz.BlindedPublicKeyHash
}

func (*BalanceUpdateCommitments) BalanceUpdateKind() string { return "commitments" }

type LegacyRewards struct {
	Delegate tz.PublicKeyHash
	Cycle    int32
}

func (*LegacyRewards) BalanceUpdateKind() string { return "legacy_rewards" }

type LegacyDeposits LegacyRewards

func (*LegacyDeposits) BalanceUpdateKind() string { return "legacy_deposits" }

type LegacyFees LegacyRewards

func (*LegacyFees) BalanceUpdateKind() string { return "legacy_fees" }

type BalanceUpdateKind interface {
	core.BalanceUpdateKind
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[BalanceUpdateKind]{
		Variants: encoding.Variants[BalanceUpdateKind]{
			0:  (*BalanceUpdateContract)(nil),
			1:  (*LegacyRewards)(nil),
			2:  BalanceUpdateBlockFees{},
			3:  (*LegacyDeposits)(nil),
			4:  (*BalanceUpdateDeposits)(nil),
			5:  BalanceUpdateNonceRevelationRewards{},
			6:  BalanceUpdateDoubleSigningEvidenceRewards{},
			7:  BalanceUpdateEndorsingRewards{},
			8:  BalanceUpdateBakingRewards{},
			9:  BalanceUpdateBakingBonuses{},
			10: (*LegacyFees)(nil),
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
		},
	})
}
