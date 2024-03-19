package proto_019_PtParisA

import (
	"fmt"
	"testing"

	"github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConstants(t *testing.T) {
	src := []byte{
		0x00, 0x00, 0x01, 0x40, 0x08, 0x20, 0x84, 0x00, 0x00, 0x80, 0x00, 0x14, 0x00, 0x00, 0xc3, 0x50,
		0x00, 0x00, 0xc3, 0x50, 0x00, 0x00, 0x27, 0x10, 0x03, 0x07, 0xd1, 0x02, 0x00, 0x00, 0x75, 0x30,
		0x00, 0x00, 0x10, 0x00, 0xc0, 0x84, 0x3d, 0x03, 0x01, 0x03, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00,
		0x00, 0x10, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x01, 0x80, 0xfa, 0x7e, 0x80, 0xe2, 0xfa,
		0x04, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x80, 0xf8, 0x82, 0xad, 0x16, 0x80, 0x8c,
		0x8d, 0x9e, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x98, 0x96, 0x80, 0x00, 0x00, 0x01, 0x01, 0x84,
		0xa5, 0x93, 0x26, 0x00, 0x00, 0x14, 0x00, 0x00, 0x00, 0x14, 0x00, 0x00, 0x00, 0x28, 0x00, 0x00,
		0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0xfa, 0x01, 0xa0, 0xa9, 0x07, 0x00, 0x00, 0x07, 0xd0,
		0x00, 0x00, 0x1b, 0x58, 0x00, 0x00, 0x01, 0xf4, 0xc0, 0x96, 0xb1, 0x02, 0x00, 0x01, 0x86, 0xa0,
		0x00, 0x78, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x07, 0x00, 0x00, 0x1b, 0x58, 0x00, 0x00, 0x12, 0x3b, 0x00, 0x02, 0x00, 0x03, 0x09, 0x02,
		0xbc, 0x13, 0x88, 0x27, 0x10, 0x00, 0x00, 0x09, 0x1e, 0x00, 0x00, 0x05, 0xf5, 0xe1, 0x00, 0x08,
		0x08, 0xff, 0x00, 0x00, 0x20, 0x08, 0x42, 0x08, 0x0f, 0x7f, 0x00, 0x01, 0xef, 0xe0, 0x02, 0x00,
		0xff, 0x00, 0x00, 0x18, 0xaa, 0x00, 0x00, 0x00, 0x28, 0x80, 0x90, 0xa1, 0x0f, 0x00, 0x00, 0x00,
		0x14, 0x00, 0x00, 0x75, 0x30, 0x00, 0x00, 0x4e, 0xc0, 0x00, 0x00, 0x00, 0x64, 0x20, 0x00, 0x00,
		0x01, 0xf4, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x03, 0xb1, 0x00, 0xff, 0xff,
		0xff, 0x00, 0x00, 0x0f, 0xa0, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x08, 0x00, 0x05, 0x02, 0x00,
		0x98, 0x96, 0x80, 0x01, 0x90, 0x06, 0x01, 0x0a, 0x09, 0x88, 0x03, 0x0b, 0x88, 0x03, 0x0a, 0x32,
		0x00, 0x00, 0x2d, 0x79, 0x88, 0x3d, 0x20, 0x00, 0x01, 0xa4, 0x01, 0x01, 0x02, 0x01, 0x32, 0xff,
		0xff, 0x00, 0xff, 0x00,
	}

	expect := Constants{
		ProofOfWorkNonceSize:                   8,
		NonceLength:                            32,
		MaxAnonOpsPerBlock:                     132,
		MaxOperationDataLength:                 32768,
		MaxProposalsPerDelegate:                20,
		MaxMichelineNodeCount:                  50000,
		MaxMichelineBytesLimit:                 50000,
		MaxAllowedGlobalConstantsDepth:         10000,
		CacheLayoutSize:                        3,
		MichelsonMaximumTypeSize:               2001,
		MaxSlashingPeriod:                      2,
		SmartRollupMaxWrappedProofBinarySize:   30000,
		SmartRollupMessageSizeLimit:            4096,
		SmartRollupMaxNumberOfMessagesPerLevel: gotez.BigUint{0xc0, 0x84, 0x3d},
		ConsensusRightsDelay:                   3,
		BlocksPreservationCycles:               1,
		DelegateParametersActivationDelay:      3,
		BlocksPerCycle:                         128,
		BlocksPerCommitment:                    16,
		NonceRevelationThreshold:               32,
		CyclesPerVotingPeriod:                  1,
		HardGasLimitPerOperation:               gotez.BigInt{0x80, 0xfa, 0x7e},
		HardGasLimitPerBlock:                   gotez.BigInt{0x80, 0xe2, 0xfa, 0x04},
		ProofOfWorkThreshold:                   -1,
		MinimalStake:                           gotez.BigUint{0x80, 0xf8, 0x82, 0xad, 0x16},
		MinimalFrozenStake:                     gotez.BigUint{0x80, 0x8c, 0x8d, 0x9e, 02},
		VDFDifficulty:                          10000000,
		OriginationSize:                        257,
		IssuanceWeights: IssuanceWeights{
			BaseTotalIssuedPerMinute:       gotez.BigUint{0x84, 0xa5, 0x93, 0x26},
			BakingRewardFixedPortionWeight: 5120,
			BakingRewardBonusWeight:        5120,
			AttestingRewardWeight:          10240,
			SeedNonceRevelationTipWeight:   1,
			VDFRevelationTipWeight:         1,
		},
		CostPerByte:                       gotez.BigUint{0xfa, 0x01},
		HardStorageLimitPerOperation:      gotez.BigInt{0xa0, 0xa9, 0x07},
		QuorumMin:                         2000,
		QuorumMax:                         7000,
		MinProposalQuorum:                 500,
		LiquidityBakingSubsidy:            gotez.BigUint{0xc0, 0x96, 0xb1, 0x02},
		LiquidityBakingToggleEmaThreshold: 100000,
		MaxOperationsTimeToLive:           120,
		MinimalBlockDelay:                 7,
		DelayIncrementPerRound:            7,
		ConsensusCommitteeSize:            7000,
		ConsensusThreshold:                4667,
		MinimalParticipationRatio:         core.Rat{2, 3},
		LimitOfDelegationOverBaking:       9,
		PercentageOfFrozenDepositsSlashedPerDoubleBaking:      700,
		PercentageOfFrozenDepositsSlashedPerDoubleAttestation: 5000,
		MaxSlashingPerBlock:          10000,
		MaxSlashingThreshold:         2334,
		TestnetDictator:              gotez.None[gotez.PublicKeyHash](),
		InitialSeed:                  gotez.None[*gotez.Bytes32](),
		CacheScriptSize:              100000000,
		CacheStakeDistributionCycles: 8,
		CacheSamplerStateCycles:      8,
		DALParametric: DALParametric{
			FeatureEnable:        true,
			IncentivesEnable:     false,
			NumberOfSlots:        32,
			AttestationLag:       8,
			AttestationThreshold: 66,
			RedundancyFactor:     8,
			PageSize:             3967,
			SlotSize:             126944,
			NumberOfShards:       512,
		},
		SmartRollupArithPvmEnable:                 true,
		SmartRollupOriginationSize:                6314,
		SmartRollupChallengeWindowInBlocks:        40,
		SmartRollupStakeAmount:                    gotez.BigUint{0x80, 0x90, 0xa1, 0x0f},
		SmartRollupCommitmentPeriodInBlocks:       20,
		SmartRollupMaxLookaheadInBlocks:           30000,
		SmartRollupMaxActiveOutboxLevels:          20160,
		SmartRollupMaxOutboxMessagesPerLevel:      100,
		SmartRollupNumberOfSectionsInDissection:   32,
		SmartRollupTimeoutPeriodInBlocks:          500,
		SmartRollupMaxNumberOfCementedCommitments: 5,
		SmartRollupMaxNumberOfParallelGames:       32,
		SmartRollupRevealActivationLevel: SmartRollupRevealActivationLevel{
			RawData:                     0,
			Metadata:                    0,
			DALPage:                     1,
			DALParameters:               1,
			DALAttestedSlotsValidityLag: 241920,
		},
		SmartRollupPrivateEnable:           true,
		SmartRollupRiscvPVMEnable:          true,
		ZkRollupEnable:                     true,
		ZkRollupOriginationSize:            4000,
		ZkRollupMinPendingToProcess:        10,
		ZkRollupMaxTicketPayloadSize:       2048,
		GlobalLimitOfStakingOverBaking:     5,
		EdgeOfStakingOverDelegation:        2,
		AdaptiveIssuanceLaunchEMAThreshold: 10000000,
		AdaptiveRewardsParams: AdaptiveRewardsParams{
			IssuanceRatioFinalMin: core.BigRat{
				gotez.BigInt{0x01},
				gotez.BigInt{0x90, 0x06},
			},
			IssuanceRatioFinalMax: core.BigRat{
				gotez.BigInt{0x01},
				gotez.BigInt{0x0a},
			},
			IssuanceRatioInitialMin: core.BigRat{
				gotez.BigInt{0x09},
				gotez.BigInt{0x88, 0x03},
			},
			IssuanceRatioInitialMax: core.BigRat{
				gotez.BigInt{0x0b},
				gotez.BigInt{0x88, 0x03},
			},
			InitialPeriod:    10,
			TransitionPeriod: 50,
			MaxBonus:         50000000000000,
			GrowthRate: core.BigRat{
				gotez.BigInt{0x01},
				gotez.BigInt{0xa4, 0x01},
			},
			CenterDz: core.BigRat{
				gotez.BigInt{0x01},
				gotez.BigInt{0x02},
			},
			RadiusDz: core.BigRat{
				gotez.BigInt{0x01},
				gotez.BigInt{0x32},
			},
		},
		AdaptiveIssuanceActivationVoteEnable: true,
		AutostakingEnable:                    true,
		AdaptiveIssuanceForceActivation:      false,
		NSEnable:                             true,
		DirectTicketSpendingEnable:           false,
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
