package proto_024_PtTALLiN

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed constants.bin
var src []byte

func TestConstants(t *testing.T) {
	expect := Constants{
		ProofOfWorkNonceSize:                   0x8,
		NonceLength:                            0x20,
		MaxAnonOpsPerBlock:                     0x84,
		MaxOperationDataLength:                 32768,
		MaxProposalsPerDelegate:                0x14,
		MaxMichelineNodeCount:                  50000,
		MaxMichelineBytesLimit:                 50000,
		MaxAllowedGlobalConstantsDepth:         10000,
		CacheLayoutSize:                        0x3,
		MichelsonMaximumTypeSize:               0x7d1,
		DenunciationPeriod:                     0x1,
		SlashingDelay:                          0x1,
		SmartRollupMaxWrappedProofBinarySize:   30000,
		SmartRollupMessageSizeLimit:            4096,
		SmartRollupMaxNumberOfMessagesPerLevel: gotez.BigUint{0xc0, 0x84, 0x3d},
		ConsensusRightsDelay:                   0x2,
		BlocksPreservationCycles:               0x1,
		DelegateParametersActivationDelay:      0x3,
		ToleratedInactivityPeriod:              0x2,
		BlocksPerCycle:                         300,
		BlocksPerCommitment:                    25,
		NonceRevelationThreshold:               50,
		CyclesPerVotingPeriod:                  1,
		HardGasLimitPerOperation:               gotez.BigInt{0x80, 0xfa, 0x7e},
		HardGasLimitPerBlock:                   gotez.BigInt{0xaa, 0xa2, 0xa9, 0x1},
		ProofOfWorkThreshold:                   -1,
		MinimalStake:                           gotez.BigUint{0x80, 0xf8, 0x82, 0xad, 0x16},
		MinimalFrozenStake:                     gotez.BigUint{0x80, 0x8c, 0x8d, 0x9e, 0x2},
		VDFDifficulty:                          10000000,
		OriginationSize:                        257,
		IssuanceWeights: IssuanceWeights{
			BaseTotalIssuedPerMinute:       gotez.BigUint{0xc4, 0xbb, 0xc4, 0x28},
			BakingRewardFixedPortionWeight: 5120,
			BakingRewardBonusWeight:        5120,
			AttestingRewardWeight:          10240,
			SeedNonceRevelationTipWeight:   1,
			VDFRevelationTipWeight:         1,
			DALRewardsWeight:               2275,
		},
		CostPerByte:                       gotez.BigUint{0xfa, 0x1},
		HardStorageLimitPerOperation:      gotez.BigInt{0xa0, 0xa9, 0x7},
		QuorumMin:                         2000,
		QuorumMax:                         7000,
		MinProposalQuorum:                 500,
		LiquidityBakingSubsidy:            gotez.BigUint{0xc0, 0x96, 0xb1, 0x2},
		LiquidityBakingToggleEmaThreshold: 100000,
		MaxOperationsTimeToLive:           150,
		MinimalBlockDelay:                 4,
		DelayIncrementPerRound:            2,
		ConsensusCommitteeSize:            7000,
		ConsensusThresholdSize:            4667,
		MinimalParticipationRatio:         core.Rat{0x2, 0x3},
		LimitOfDelegationOverBaking:       0x9,
		PercentageOfFrozenDepositsSlashedPerDoubleBaking: 0x1f4,
		MaxSlashingPerBlock:          0x2710,
		MaxSlashingThreshold:         core.Rat{0x1, 0x3},
		TestnetDictator:              gotez.Some[gotez.PublicKeyHash](&gotez.Ed25519PublicKeyHash{0xc9, 0x7f, 0xb9, 0x64, 0xe7, 0x69, 0xad, 0x54, 0x96, 0xdb, 0x51, 0xe8, 0x86, 0xe3, 0xf0, 0x7b, 0xd2, 0x81, 0xed, 0x74}),
		InitialSeed:                  gotez.None[*gotez.Bytes32](),
		CacheScriptSize:              100000000,
		CacheStakeDistributionCycles: 5,
		CacheSamplerStateCycles:      5,
		DALParametric: DALParametric{
			FeatureEnable:             true,
			IncentivesEnable:          true,
			NumberOfSlots:             0x20,
			AttestationLag:            0x8,
			AttestationThreshold:      0x42,
			MinimalParticipationRatio: core.BigRat{gotez.BigInt{0x10}, gotez.BigInt{0x19}},
			RewardsRatio:              core.BigRat{gotez.BigInt{0x1}, gotez.BigInt{0xa}},
			TrapsFraction:             core.BigRat{gotez.BigInt{0x1}, gotez.BigInt{0x90, 0x1f}},
			RedundancyFactor:          0x8,
			PageSize:                  0xf7f,
			SlotSize:                  126944,
			NumberOfShards:            0x200,
		},
		SmartRollupArithPvmEnable:                 true,
		SmartRollupOriginationSize:                6314,
		SmartRollupChallengeWindowInBlocks:        62,
		SmartRollupStakeAmount:                    gotez.BigUint{0x80, 0x90, 0xa1, 0xf},
		SmartRollupCommitmentPeriodInBlocks:       31,
		SmartRollupMaxLookaheadInBlocks:           46875,
		SmartRollupMaxActiveOutboxLevels:          31500,
		SmartRollupMaxOutboxMessagesPerLevel:      100,
		SmartRollupNumberOfSectionsInDissection:   0x20,
		SmartRollupTimeoutPeriodInBlocks:          781,
		SmartRollupMaxNumberOfCementedCommitments: 5,
		SmartRollupMaxNumberOfParallelGames:       32,
		SmartRollupRevealActivationLevel: SmartRollupRevealActivationLevel{
			RawData:                     0,
			Metadata:                    0,
			DALPage:                     1,
			DALParameters:               1,
			DALAttestedSlotsValidityLag: 241920,
		},
		SmartRollupPrivateEnable:       true,
		SmartRollupRiscvPvmEnable:      true,
		ZkRollupEnable:                 true,
		ZkRollupOriginationSize:        4000,
		ZkRollupMinPendingToProcess:    10,
		ZkRollupMaxTicketPayloadSize:   2048,
		GlobalLimitOfStakingOverBaking: 0x9,
		EdgeOfStakingOverDelegation:    0x3,
		AdaptiveRewardsParams: AdaptiveRewardsParams{
			IssuanceRatioFinalMin:   core.BigRat{gotez.BigInt{0x1}, gotez.BigInt{0x90, 0x6}},
			IssuanceRatioFinalMax:   core.BigRat{gotez.BigInt{0x1}, gotez.BigInt{0xa}},
			IssuanceRatioInitialMin: core.BigRat{gotez.BigInt{0x9}, gotez.BigInt{0x88, 0x3}},
			IssuanceRatioInitialMax: core.BigRat{gotez.BigInt{0xb}, gotez.BigInt{0x88, 0x3}},
			InitialPeriod:           0xa,
			TransitionPeriod:        0x32,
			MaxBonus:                50000000000000,
			GrowthRate:              core.BigRat{gotez.BigInt{0x1}, gotez.BigInt{0xa4, 0x1}},
			CenterDz:                core.BigRat{gotez.BigInt{0x1}, gotez.BigInt{0x2}},
			RadiusDz:                core.BigRat{gotez.BigInt{0x1}, gotez.BigInt{0x32}},
		},
		DirectTicketSpendingEnable:         false,
		AggregateAttestation:               true,
		AllowTz4DelegateEnable:             true,
		AllBakersAttestActivationThreshold: core.Rat{0x1, 0x2},
		IssuanceModificationDelay:          0x2,
		ConsensusKeyActivationDelay:        0x2,
		UnstakeFinalizationDelay:           0x3,
	}

	var out Constants
	_, err := encoding.Decode(src, &out, encoding.Dynamic())
	if !assert.NoError(t, err) {
		if err, ok := err.(*encoding.Error); ok {
			fmt.Println(err.Path)
		}
	} else {
		require.Equal(t, &expect, &out)
	}
}
