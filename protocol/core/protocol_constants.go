package core

import tz "github.com/ecadlabs/gotez/v2"

type Constants interface {
	GetProofOfWorkNonceSize() uint8
	GetNonceLength() uint8
	GetMaxAnonOpsPerBlock() uint8
	GetMaxOperationDataLength() int32
	GetMaxProposalsPerDelegate() uint8
	GetMaxMichelineNodeCount() int32
	GetMaxMichelineBytesLimit() int32
	GetMaxAllowedGlobalConstantsDepth() int32
	GetMichelsonMaximumTypeSize() uint16
	GetBlocksPerCycle() int32
	GetBlocksPerCommitment() int32
	GetHardGasLimitPerOperation() tz.BigInt
	GetHardGasLimitPerBlock() tz.BigInt
	GetProofOfWorkThreshold() int64
	GetOriginationSize() int32
	GetCostPerByte() tz.BigUint
	GetHardStorageLimitPerOperation() tz.BigInt
	GetQuorumMin() int32
	GetQuorumMax() int32
	GetMinProposalQuorum() int32
	GetMaxOperationsTimeToLive() int16
	GetMinimalBlockDelay() int64
	GetDelayIncrementPerRound() int64
	GetConsensusCommitteeSize() int32
	GetConsensusThreshold() int32
	GetMinimalParticipationRatio() *Rat
	GetMaxSlashingPeriod() int32
}
