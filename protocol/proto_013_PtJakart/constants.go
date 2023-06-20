package proto_013_PtJakart

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/protocol/core"
)

type Constants struct {
	ProofOfWorkNonceSize                             uint8                  `json:"proof_of_work_nonce_size"`
	NonceLength                                      uint8                  `json:"nonce_length"`
	MaxAnonOpsPerBlock                               uint8                  `json:"max_anon_ops_per_block"`
	MaxOperationDataLength                           int32                  `json:"max_operation_data_length"`
	MaxProposalsPerDelegate                          uint8                  `json:"max_proposals_per_delegate"`
	MaxMichelineNodeCount                            int32                  `json:"max_micheline_node_count"`
	MaxMichelineBytesLimit                           int32                  `json:"max_micheline_bytes_limit"`
	MaxAllowedGlobalConstantsDepth                   int32                  `json:"max_allowed_global_constants_depth"`
	CacheLayoutSize                                  uint8                  `json:"cache_layout_size"`
	MichelsonMaximumTypeSize                         uint16                 `json:"michelson_maximum_type_size"`
	PreservedCycles                                  uint8                  `json:"preserved_cycles"`
	BlocksPerCycle                                   int32                  `json:"blocks_per_cycle"`
	BlocksPerCommitment                              int32                  `json:"blocks_per_commitment"`
	BlocksPerStakeSnapshot                           int32                  `json:"blocks_per_stake_snapshot"`
	CyclesPerVotingPeriod                            int32                  `json:"cycles_per_voting_period"`
	HardGasLimitPerOperation                         tz.BigInt              `json:"hard_gas_limit_per_operation"`
	HardGasLimitPerBlock                             tz.BigInt              `json:"hard_gas_limit_per_block"`
	ProofOfWorkThreshold                             int64                  `json:"proof_of_work_threshold"`
	TokensPerRoll                                    tz.BigUint             `json:"tokens_per_roll"`
	SeedNonceRevelationTip                           tz.BigUint             `json:"seed_nonce_revelation_tip"`
	OriginationSize                                  int32                  `json:"origination_size"`
	BakingRewardFixedPortion                         tz.BigUint             `json:"baking_reward_fixed_portion"`
	BakingRewardBonusPerSlot                         tz.BigUint             `json:"baking_reward_bonus_per_slot"`
	EndorsingRewardPerSlot                           tz.BigUint             `json:"endorsing_reward_per_slot"`
	CostPerByte                                      tz.BigUint             `json:"cost_per_byte"`
	HardStorageLimitPerOperation                     tz.BigInt              `json:"hard_storage_limit_per_operation"`
	QuorumMin                                        int32                  `json:"quorum_min"`
	QuorumMax                                        int32                  `json:"quorum_max"`
	MinProposalQuorum                                int32                  `json:"min_proposal_quorum"`
	LiquidityBakingSubsidy                           tz.BigUint             `json:"liquidity_baking_subsidy"`
	LiquidityBakingSunsetLevel                       int32                  `json:"liquidity_baking_sunset_level"`
	LiquidityBakingToggleEmaThreshold                int32                  `json:"liquidity_baking_toggle_ema_threshold"`
	MaxOperationsTimeToLive                          int16                  `json:"max_operations_time_to_live"`
	MinimalBlockDelay                                int64                  `json:"minimal_block_delay"`
	DelayIncrementPerRound                           int64                  `json:"delay_increment_per_round"`
	ConsensusCommitteeSize                           int32                  `json:"consensus_committee_size"`
	ConsensusThreshold                               int32                  `json:"consensus_threshold"`
	MinimalParticipationRatio                        core.Rat               `json:"minimal_participation_ratio"`
	MaxSlashingPeriod                                int32                  `json:"max_slashing_period"`
	FrozenDepositsPercentage                         int32                  `json:"frozen_deposits_percentage"`
	DoubleBakingPunishment                           tz.BigUint             `json:"double_baking_punishment"`
	RatioOfFrozenDepositsSlashedPerDoubleEndorsement core.Rat               `json:"ratio_of_frozen_deposits_slashed_per_double_endorsement"`
	InitialSeed                                      tz.Option[*tz.Bytes32] `json:"initial_seed"`
	CacheScriptSize                                  int32                  `json:"cache_script_size"`
	CacheStakeDistributionCycles                     int8                   `json:"cache_stake_distribution_cycles"`
	CacheSamplerStateCycles                          int8                   `json:"cache_sampler_state_cycles"`
	TxRollupEnable                                   bool                   `json:"tx_rollup_enable"`
	TxRollupOriginationSize                          int32                  `json:"tx_rollup_origination_size"`
	TxRollupHardSizeLimitPerInbox                    int32                  `json:"tx_rollup_hard_size_limit_per_inbox"`
	TxRollupHardSizeLimitPerMessage                  int32                  `json:"tx_rollup_hard_size_limit_per_message"`
	TxRollupMaxWithdrawalsPerBatch                   int32                  `json:"tx_rollup_max_withdrawals_per_batch"`
	TxRollupCommitmentBond                           tz.BigUint             `json:"tx_rollup_commitment_bond"`
	TxRollupFinalityPeriod                           int32                  `json:"tx_rollup_finality_period"`
	TxRollupWithdrawPeriod                           int32                  `json:"tx_rollup_withdraw_period"`
	TxRollupMaxInboxesCount                          int32                  `json:"tx_rollup_max_inboxes_count"`
	TxRollupMaxMessagesPerInbox                      int32                  `json:"tx_rollup_max_messages_per_inbox"`
	TxRollupMaxCommitmentsCount                      int32                  `json:"tx_rollup_max_commitments_count"`
	TxRollupCostPerByteEmaFactor                     int32                  `json:"tx_rollup_cost_per_byte_ema_factor"`
	TxRollupMaxTicketPayloadSize                     int32                  `json:"tx_rollup_max_ticket_payload_size"`
	TxRollupRejectionMaxProofSize                    int32                  `json:"tx_rollup_rejection_max_proof_size"`
	TxRollupSunsetLevel                              int32                  `json:"tx_rollup_sunset_level"`
	ScRollupEnable                                   bool                   `json:"sc_rollup_enable"`
	ScRollupOriginationSize                          int32                  `json:"sc_rollup_origination_size"`
	ScRollupChallengeWindowInBlocks                  int32                  `json:"sc_rollup_challenge_window_in_blocks"`
	ScRollupMaxAvailableMessages                     int32                  `json:"sc_rollup_max_available_messages"`
}

func (c *Constants) GetProofOfWorkNonceSize() uint8    { return c.ProofOfWorkNonceSize }
func (c *Constants) GetNonceLength() uint8             { return c.NonceLength }
func (c *Constants) GetMaxAnonOpsPerBlock() uint8      { return c.MaxAnonOpsPerBlock }
func (c *Constants) GetMaxOperationDataLength() int32  { return c.MaxOperationDataLength }
func (c *Constants) GetMaxProposalsPerDelegate() uint8 { return c.MaxProposalsPerDelegate }
func (c *Constants) GetMaxMichelineNodeCount() int32   { return c.MaxMichelineNodeCount }
func (c *Constants) GetMaxMichelineBytesLimit() int32  { return c.MaxMichelineBytesLimit }
func (c *Constants) GetMaxAllowedGlobalConstantsDepth() int32 {
	return c.MaxAllowedGlobalConstantsDepth
}
func (c *Constants) GetMichelsonMaximumTypeSize() uint16     { return c.MichelsonMaximumTypeSize }
func (c *Constants) GetPreservedCycles() uint8               { return c.PreservedCycles }
func (c *Constants) GetBlocksPerCycle() int32                { return c.BlocksPerCycle }
func (c *Constants) GetBlocksPerCommitment() int32           { return c.BlocksPerCommitment }
func (c *Constants) GetBlocksPerStakeSnapshot() int32        { return c.BlocksPerStakeSnapshot }
func (c *Constants) GetHardGasLimitPerOperation() tz.BigInt  { return c.HardGasLimitPerOperation }
func (c *Constants) GetHardGasLimitPerBlock() tz.BigInt      { return c.HardGasLimitPerBlock }
func (c *Constants) GetProofOfWorkThreshold() int64          { return c.ProofOfWorkThreshold }
func (c *Constants) GetSeedNonceRevelationTip() tz.BigUint   { return c.SeedNonceRevelationTip }
func (c *Constants) GetOriginationSize() int32               { return c.OriginationSize }
func (c *Constants) GetBakingRewardFixedPortion() tz.BigUint { return c.BakingRewardFixedPortion }
func (c *Constants) GetBakingRewardBonusPerSlot() tz.BigUint { return c.BakingRewardBonusPerSlot }
func (c *Constants) GetEndorsingRewardPerSlot() tz.BigUint   { return c.EndorsingRewardPerSlot }
func (c *Constants) GetCostPerByte() tz.BigUint              { return c.CostPerByte }
func (c *Constants) GetHardStorageLimitPerOperation() tz.BigInt {
	return c.HardStorageLimitPerOperation
}
func (c *Constants) GetQuorumMin() int32                     { return c.QuorumMin }
func (c *Constants) GetQuorumMax() int32                     { return c.QuorumMax }
func (c *Constants) GetMinProposalQuorum() int32             { return c.MinProposalQuorum }
func (c *Constants) GetLiquidityBakingSubsidy() tz.BigUint   { return c.LiquidityBakingSubsidy }
func (c *Constants) GetMaxOperationsTimeToLive() int16       { return c.MaxOperationsTimeToLive }
func (c *Constants) GetMinimalBlockDelay() int64             { return c.MinimalBlockDelay }
func (c *Constants) GetDelayIncrementPerRound() int64        { return c.DelayIncrementPerRound }
func (c *Constants) GetConsensusCommitteeSize() int32        { return c.ConsensusCommitteeSize }
func (c *Constants) GetConsensusThreshold() int32            { return c.ConsensusThreshold }
func (c *Constants) GetMinimalParticipationRatio() *core.Rat { return &c.MinimalParticipationRatio }
func (c *Constants) GetMaxSlashingPeriod() int32             { return c.MaxSlashingPeriod }
func (c *Constants) GetFrozenDepositsPercentage() int32      { return c.FrozenDepositsPercentage }
func (c *Constants) GetDoubleBakingPunishment() tz.BigUint   { return c.DoubleBakingPunishment }
func (c *Constants) GetRatioOfFrozenDepositsSlashedPerDoubleEndorsement() *core.Rat {
	return &c.RatioOfFrozenDepositsSlashedPerDoubleEndorsement
}
