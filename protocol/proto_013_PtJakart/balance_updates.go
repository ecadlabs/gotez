package proto_013_PtJakart

type BalanceUpdateTxRollupRejectionRewards struct{}

func (BalanceUpdateTxRollupRejectionRewards) BalanceUpdateKind() string {
	return "tx_rollup_rejection_rewards"
}

type BalanceUpdateTxRollupRejectionPunishments struct{}

func (BalanceUpdateTxRollupRejectionPunishments) BalanceUpdateKind() string {
	return "tx_rollup_rejection_punishments"
}
