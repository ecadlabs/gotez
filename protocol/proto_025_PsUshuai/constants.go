package proto_025_PsUshuai

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/proto_024_PtTALLiN"
)

// Unchanged sub-structures are reused from T024.
type IssuanceWeights = proto_024_PtTALLiN.IssuanceWeights
type SmartRollupRevealActivationLevel = proto_024_PtTALLiN.SmartRollupRevealActivationLevel
type AdaptiveRewardsParams = proto_024_PtTALLiN.AdaptiveRewardsParams

// DALParametric gains two fields in U025 versus T024 (MRs !19994/!20429, !20125):
//   - DynamicLagEnable, inserted after IncentivesEnable
//   - AttestationLags (the list of permitted lags), inserted after AttestationLag
//
// Both sit inside the DAL parametric sub-encoding, so the struct cannot alias the
// T024 definition.
type DALParametric struct {
	FeatureEnable             bool        `json:"feature_enable"`
	IncentivesEnable          bool        `json:"incentives_enable"`
	DynamicLagEnable          bool        `json:"dynamic_lag_enable"`
	NumberOfSlots             uint16      `json:"number_of_slots"`
	AttestationLag            uint8       `json:"attestation_lag"`
	AttestationLags           []uint8     `tz:"dyn" json:"attestation_lags"`
	AttestationThreshold      uint8       `json:"attestation_threshold"`
	MinimalParticipationRatio core.BigRat `json:"minimal_participation_ratio"`
	RewardsRatio              core.BigRat `json:"rewards_ratio"`
	TrapsFraction             core.BigRat `json:"traps_fraction"`
	RedundancyFactor          uint8       `json:"redundancy_factor"`
	PageSize                  uint16      `json:"page_size"`
	SlotSize                  int32       `json:"slot_size"`
	NumberOfShards            uint16      `json:"number_of_shards"`
}

// Constants mirrors T024 with the U025 parametric additions. New fields are placed
// at the exact byte positions dictated by the protocol's parametric encoding
// (src/proto_025_PsUshuai/lib_protocol/constants_parametric_repr.ml):
//   - CacheStakeInfoCycles / CacheSwrrSelectedDistributionCycles after CacheSamplerStateCycles
//   - CanonicalRollupAddress (optional) after SmartRollupRiscvPvmEnable
//   - NativeContractsEnable / SwrrNewBakerLotteryEnable / Tz5AccountEnable after
//     AllBakersAttestActivationThreshold (the trailing parametric feature flags)
//
// IssuanceModificationDelay / ConsensusKeyActivationDelay / UnstakeFinalizationDelay
// remain last: they are derived constants (Constants_repr.Derived) appended after the
// parametric block in the /context/constants RPC payload.
type Constants struct {
	ProofOfWorkNonceSize                             uint8                             `json:"proof_of_work_nonce_size"`
	NonceLength                                      uint8                             `json:"nonce_length"`
	MaxAnonOpsPerBlock                               uint8                             `json:"max_anon_ops_per_block"`
	MaxOperationDataLength                           int32                             `json:"max_operation_data_length"`
	MaxProposalsPerDelegate                          uint8                             `json:"max_proposals_per_delegate"`
	MaxMichelineNodeCount                            int32                             `json:"max_micheline_node_count"`
	MaxMichelineBytesLimit                           int32                             `json:"max_micheline_bytes_limit"`
	MaxAllowedGlobalConstantsDepth                   int32                             `json:"max_allowed_global_constants_depth"`
	CacheLayoutSize                                  uint8                             `json:"cache_layout_size"`
	MichelsonMaximumTypeSize                         uint16                            `json:"michelson_maximum_type_size"`
	DenunciationPeriod                               uint8                             `json:"denunciation_period"`
	SlashingDelay                                    uint8                             `json:"slashing_delay"`
	SmartRollupMaxWrappedProofBinarySize             int32                             `json:"smart_rollup_max_wrapped_proof_binary_size"`
	SmartRollupMessageSizeLimit                      int32                             `json:"smart_rollup_message_size_limit"`
	SmartRollupMaxNumberOfMessagesPerLevel           tz.BigUint                        `json:"smart_rollup_max_number_of_messages_per_level"`
	ConsensusRightsDelay                             uint8                             `json:"consensus_rights_delay"`
	BlocksPreservationCycles                         uint8                             `json:"blocks_preservation_cycles"`
	DelegateParametersActivationDelay                uint8                             `json:"delegate_parameters_activation_delay"`
	ToleratedInactivityPeriod                        uint8                             `json:"tolerated_inactivity_period"`
	BlocksPerCycle                                   int32                             `json:"blocks_per_cycle"`
	BlocksPerCommitment                              int32                             `json:"blocks_per_commitment"`
	NonceRevelationThreshold                         int32                             `json:"nonce_revelation_threshold"`
	CyclesPerVotingPeriod                            int32                             `json:"cycles_per_voting_period"`
	HardGasLimitPerOperation                         tz.BigInt                         `json:"hard_gas_limit_per_operation"`
	HardGasLimitPerBlock                             tz.BigInt                         `json:"hard_gas_limit_per_block"`
	ProofOfWorkThreshold                             int64                             `json:"proof_of_work_threshold"`
	MinimalStake                                     tz.BigUint                        `json:"minimal_stake"`
	MinimalFrozenStake                               tz.BigUint                        `json:"minimal_frozen_stake"`
	VDFDifficulty                                    int64                             `json:"vdf_difficulty"`
	OriginationSize                                  int32                             `json:"origination_size"`
	IssuanceWeights                                  IssuanceWeights                   `json:"issuance_weights"`
	CostPerByte                                      tz.BigUint                        `json:"cost_per_byte"`
	HardStorageLimitPerOperation                     tz.BigInt                         `json:"hard_storage_limit_per_operation"`
	QuorumMin                                        int32                             `json:"quorum_min"`
	QuorumMax                                        int32                             `json:"quorum_max"`
	MinProposalQuorum                                int32                             `json:"min_proposal_quorum"`
	LiquidityBakingSubsidy                           tz.BigUint                        `json:"liquidity_baking_subsidy"`
	LiquidityBakingToggleEmaThreshold                int32                             `json:"liquidity_baking_toggle_ema_threshold"`
	MaxOperationsTimeToLive                          int16                             `json:"max_operations_time_to_live"`
	MinimalBlockDelay                                int64                             `json:"minimal_block_delay"`
	DelayIncrementPerRound                           int64                             `json:"delay_increment_per_round"`
	ConsensusCommitteeSize                           int32                             `json:"consensus_committee_size"`
	ConsensusThresholdSize                           int32                             `json:"consensus_threshold_size"`
	MinimalParticipationRatio                        core.Rat                          `json:"minimal_participation_ratio"`
	LimitOfDelegationOverBaking                      uint8                             `json:"limit_of_delegation_over_baking"`
	PercentageOfFrozenDepositsSlashedPerDoubleBaking uint16                            `json:"percentage_of_frozen_deposits_slashed_per_double_baking"`
	MaxSlashingPerBlock                              uint16                            `json:"max_slashing_per_block"`
	MaxSlashingThreshold                             core.Rat                          `json:"max_slashing_threshold"`
	TestnetDictator                                  tz.Option[tz.PublicKeyHash]       `json:"testnet_dictator"`
	InitialSeed                                      tz.Option[*tz.Bytes32]            `json:"initial_seed"`
	CacheScriptSize                                  int32                             `json:"cache_script_size"`
	CacheStakeDistributionCycles                     int8                              `json:"cache_stake_distribution_cycles"`
	CacheSamplerStateCycles                          int8                              `json:"cache_sampler_state_cycles"`
	CacheStakeInfoCycles                             int8                              `json:"cache_stake_info_cycles"`
	CacheSwrrSelectedDistributionCycles              int8                              `json:"cache_swrr_selected_distribution_cycles"`
	DALParametric                                    DALParametric                     `json:"dal_parametric"`
	SmartRollupArithPvmEnable                        bool                              `json:"smart_rollup_arith_pvm_enable"`
	SmartRollupOriginationSize                       int32                             `json:"smart_rollup_origination_size"`
	SmartRollupChallengeWindowInBlocks               int32                             `json:"smart_rollup_challenge_window_in_blocks"`
	SmartRollupStakeAmount                           tz.BigUint                        `json:"smart_rollup_stake_amount"`
	SmartRollupCommitmentPeriodInBlocks              int32                             `json:"smart_rollup_commitment_period_in_blocks"`
	SmartRollupMaxLookaheadInBlocks                  int32                             `json:"smart_rollup_max_lookahead_in_blocks"`
	SmartRollupMaxActiveOutboxLevels                 int32                             `json:"smart_rollup_max_active_outbox_levels"`
	SmartRollupMaxOutboxMessagesPerLevel             int32                             `json:"smart_rollup_max_outbox_messages_per_level"`
	SmartRollupNumberOfSectionsInDissection          uint8                             `json:"smart_rollup_number_of_sections_in_dissection"`
	SmartRollupTimeoutPeriodInBlocks                 int32                             `json:"smart_rollup_timeout_period_in_blocks"`
	SmartRollupMaxNumberOfCementedCommitments        int32                             `json:"smart_rollup_max_number_of_cemented_commitments"`
	SmartRollupMaxNumberOfParallelGames              int32                             `json:"smart_rollup_max_number_of_parallel_games"`
	SmartRollupRevealActivationLevel                 SmartRollupRevealActivationLevel  `json:"smart_rollup_reveal_activation_level"`
	SmartRollupPrivateEnable                         bool                              `json:"smart_rollup_private_enable"`
	SmartRollupRiscvPvmEnable                        bool                              `json:"smart_rollup_riscv_pvm_enable"`
	CanonicalRollupAddress                           tz.Option[*tz.SmartRollupAddress] `json:"smart_rollup_canonical_rollup_address"`
	ZkRollupEnable                                   bool                              `json:"zk_rollup_enable"`
	ZkRollupOriginationSize                          int32                             `json:"zk_rollup_origination_size"`
	ZkRollupMinPendingToProcess                      int32                             `json:"zk_rollup_min_pending_to_process"`
	ZkRollupMaxTicketPayloadSize                     int32                             `json:"zk_rollup_max_ticket_payload_size"`
	GlobalLimitOfStakingOverBaking                   uint8                             `json:"global_limit_of_staking_over_baking"`
	EdgeOfStakingOverDelegation                      uint8                             `json:"edge_of_staking_over_delegation"`
	AdaptiveRewardsParams                            AdaptiveRewardsParams             `json:"adaptive_rewards_params"`
	DirectTicketSpendingEnable                       bool                              `json:"direct_ticket_spending_enable"`
	AggregateAttestation                             bool                              `json:"aggregate_attestation"`
	AllowTz4DelegateEnable                           bool                              `json:"allow_tz4_delegate_enable"`
	AllBakersAttestActivationThreshold               core.Rat                          `json:"all_bakers_attest_activation_threshold"`
	NativeContractsEnable                            bool                              `json:"native_contracts_enable"`
	SwrrNewBakerLotteryEnable                        bool                              `json:"swrr_new_baker_lottery_enable"`
	Tz5AccountEnable                                 bool                              `json:"tz5_account_enable"`
	IssuanceModificationDelay                        uint8                             `json:"issuance_modification_delay"`
	ConsensusKeyActivationDelay                      uint8                             `json:"consensus_key_activation_delay"`
	UnstakeFinalizationDelay                         uint8                             `json:"unstake_finalization_delay"`
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
func (c *Constants) GetMichelsonMaximumTypeSize() uint16    { return c.MichelsonMaximumTypeSize }
func (c *Constants) GetBlocksPerCycle() int32               { return c.BlocksPerCycle }
func (c *Constants) GetBlocksPerCommitment() int32          { return c.BlocksPerCommitment }
func (c *Constants) GetHardGasLimitPerOperation() tz.BigInt { return c.HardGasLimitPerOperation }
func (c *Constants) GetHardGasLimitPerBlock() tz.BigInt     { return c.HardGasLimitPerBlock }
func (c *Constants) GetProofOfWorkThreshold() int64         { return c.ProofOfWorkThreshold }
func (c *Constants) GetOriginationSize() int32              { return c.OriginationSize }
func (c *Constants) GetCostPerByte() tz.BigUint             { return c.CostPerByte }
func (c *Constants) GetHardStorageLimitPerOperation() tz.BigInt {
	return c.HardStorageLimitPerOperation
}
func (c *Constants) GetQuorumMin() int32                     { return c.QuorumMin }
func (c *Constants) GetQuorumMax() int32                     { return c.QuorumMax }
func (c *Constants) GetMinProposalQuorum() int32             { return c.MinProposalQuorum }
func (c *Constants) GetMaxOperationsTimeToLive() int16       { return c.MaxOperationsTimeToLive }
func (c *Constants) GetMinimalBlockDelay() int64             { return c.MinimalBlockDelay }
func (c *Constants) GetDelayIncrementPerRound() int64        { return c.DelayIncrementPerRound }
func (c *Constants) GetConsensusCommitteeSize() int32        { return c.ConsensusCommitteeSize }
func (c *Constants) GetMinimalParticipationRatio() *core.Rat { return &c.MinimalParticipationRatio }
