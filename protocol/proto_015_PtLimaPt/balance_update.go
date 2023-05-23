package proto_015_PtLimaPt

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
	"github.com/ecadlabs/gotez/protocol/core"
	"github.com/ecadlabs/gotez/protocol/proto_012_Psithaca"
	"github.com/ecadlabs/gotez/protocol/proto_013_PtJakart"
)

type BalanceUpdateOrigin = proto_012_Psithaca.BalanceUpdateOrigin

type BalanceUpdate struct {
	Kind   BalanceUpdateKind
	Change int64
	Origin BalanceUpdateOrigin
}

func (*BalanceUpdate) BalanceUpdate() {}

type BalanceUpdateKind interface {
	core.BalanceUpdateKind
}

type BalanceUpdateContract = proto_012_Psithaca.BalanceUpdateContract
type BalanceUpdateDeposits = proto_012_Psithaca.BalanceUpdateDeposits
type BalanceUpdateLostEndorsingRewards = proto_012_Psithaca.BalanceUpdateLostEndorsingRewards
type BalanceUpdateCommitments = proto_012_Psithaca.BalanceUpdateCommitments
type BalanceUpdateBlockFees = proto_012_Psithaca.BalanceUpdateBlockFees
type BalanceUpdateNonceRevelationRewards = proto_012_Psithaca.BalanceUpdateNonceRevelationRewards
type BalanceUpdateDoubleSigningEvidenceRewards = proto_012_Psithaca.BalanceUpdateDoubleSigningEvidenceRewards
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
type BalanceUpdateTxRollupRejectionRewards = proto_013_PtJakart.BalanceUpdateTxRollupRejectionRewards
type BalanceUpdateTxRollupRejectionPunishments = proto_013_PtJakart.BalanceUpdateTxRollupRejectionPunishments

type BondID interface {
	BondID()
}

type TxRollupBondID struct {
	Address *tz.TXRollupAddress
}

func (TxRollupBondID) BondID() {}

type ScRollupBondID struct {
	Address *tz.ScRollupAddress `tz:"dyn"`
}

func (ScRollupBondID) BondID() {}

func init() {
	encoding.RegisterEnum(&encoding.Enum[BondID]{
		Variants: encoding.Variants[BondID]{
			0: TxRollupBondID{},
			1: ScRollupBondID{},
		},
	})
}

type BalanceUpdateFrozenBonds struct {
	Contract core.ContractID
	BondID   BondID
}

func (*BalanceUpdateFrozenBonds) BalanceUpdateKind() string { return "frozen_bonds" }

type BalanceUpdateScRollupRefutationPunishments struct{}

func (BalanceUpdateScRollupRefutationPunishments) BalanceUpdateKind() string {
	return "smart_rollup_refutation_punishments"
}

type BalanceUpdateScRollupRefutationRewards struct{}

func (BalanceUpdateScRollupRefutationRewards) BalanceUpdateKind() string {
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
			24: BalanceUpdateScRollupRefutationPunishments{},
			25: BalanceUpdateScRollupRefutationRewards{},
		},
	})
}