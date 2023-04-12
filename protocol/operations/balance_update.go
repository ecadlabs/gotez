package operations

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
)

type BalanceUpdate struct {
	Kind   BalanceUpdateKind
	Change int64
	Origin BalanceUpdateOrigin
}

type BalanceUpdateOrigin uint8

const (
	BalanceUpdateOriginBlockApplication BalanceUpdateOrigin = iota
	BalanceUpdateOriginProtocolMigration
	BalanceUpdateOriginSubsidy
	BalanceUpdateOriginSimulation
)

type BalanceUpdateKind interface {
	BalanceUpdateKind() string
}

type BalanceUpdateContract struct {
	Contract tz.ContractID
}

func (*BalanceUpdateContract) BalanceUpdateKind() string { return "contract" }

type BalanceUpdateDeposits struct {
	Delegate tz.PublicKeyHash
}

func (*BalanceUpdateDeposits) BalanceUpdateKind() string { return "deposits" }

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

type BondID interface {
	BondID()
}

type TxRollupBondID struct {
	Address *tz.RollupAddress
}

func (TxRollupBondID) BondID() {}

type SmartRollupBondID struct {
	Address *tz.SmartRollupAddress
}

func (SmartRollupBondID) BondID() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[BondID]{
		Variants: encoding.Variants[BondID]{
			0: TxRollupBondID{},
			1: SmartRollupBondID{},
		},
	})
}

type BalanceUpdateFrozenBonds struct {
	Contract tz.ContractID
	BondID   BondID
}

func (*BalanceUpdateFrozenBonds) BalanceUpdateKind() string { return "frozen_bonds" }

type BalanceUpdateBlockFees struct{}

func (BalanceUpdateBlockFees) BalanceUpdateKind() string {
	return "block_fees"
}

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

type BalanceUpdateTxRollupRejectionRewards struct{}

func (BalanceUpdateTxRollupRejectionRewards) BalanceUpdateKind() string {
	return "tx_rollup_rejection_rewards"
}

type BalanceUpdateTxRollupRejectionPunishments struct{}

func (BalanceUpdateTxRollupRejectionPunishments) BalanceUpdateKind() string {
	return "tx_rollup_rejection_punishments"
}

type BalanceUpdateSmartRollupRefutationPunishments struct{}

func (BalanceUpdateSmartRollupRefutationPunishments) BalanceUpdateKind() string {
	return "smart_rollup_refutation_punishments"
}

type BalanceUpdateSmartRollupRefutationRewards struct{}

func (BalanceUpdateSmartRollupRefutationRewards) BalanceUpdateKind() string {
	return "smart_rollup_refutation_rewards"
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[BalanceUpdateKind]{
		Variants: encoding.Variants[BalanceUpdateKind]{
			0:  (*BalanceUpdateContract)(nil),
			2:  BalanceUpdateBlockFees{},
			4:  (*BalanceUpdateDeposits)(nil),
			5:  BalanceUpdateNonceRevelationRewards{},
			6:  BalanceUpdateDoubleSigningEvidenceRewards{},
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
			22: BalanceUpdateTxRollupRejectionRewards{},
			23: BalanceUpdateTxRollupRejectionPunishments{},
			24: BalanceUpdateSmartRollupRefutationPunishments{},
			25: BalanceUpdateSmartRollupRefutationRewards{},
		},
	})
}
