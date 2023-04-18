package proto_012_Psithaca

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/protocol/core"
)

type BalanceUpdateOrigin uint8

const (
	BalanceUpdateOriginBlockApplication BalanceUpdateOrigin = iota
	BalanceUpdateOriginProtocolMigration
	BalanceUpdateOriginSubsidy
	BalanceUpdateOriginSimulation
)

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
