package latest

import (
	"github.com/ecadlabs/gotez/v2/protocol/core"
	latest "github.com/ecadlabs/gotez/v2/protocol/proto_018_Proxford"
)

var Protocol = core.Proto017PtNairob

type OperationContents = latest.OperationContents
type OperationContentsAndResult = latest.OperationContentsAndResult
type InternalOperationResult = latest.InternalOperationResult
type OperationWithOptionalMetadata = latest.OperationWithOptionalMetadata

type RunOperationRequest = latest.RunOperationRequest
type UnsignedOperation = latest.UnsignedOperation
type SignedOperation = latest.SignedOperation

type UnsignedBlockHeader = latest.UnsignedBlockHeader
type BlockHeader = latest.BlockHeader
type UnsignedProtocolBlockHeader = latest.UnsignedProtocolBlockHeader

type ManagerOperation = latest.ManagerOperation
type SeedNonceRevelation = latest.SeedNonceRevelation
type DoubleEndorsementEvidence = latest.DoubleAttestationEvidence
type DoubleAttestationEvidence = latest.DoubleAttestationEvidence
type DoubleBakingEvidence = latest.DoubleBakingEvidence
type ActivateAccount = latest.ActivateAccount
type Proposals = latest.Proposals
type Ballot = latest.Ballot
type DoublePreendorsementEvidence = latest.DoublePreattestationEvidence
type DoublePreattestationEvidence = latest.DoublePreattestationEvidence
type VDFRevelation = latest.VDFRevelation
type DrainDelegate = latest.DrainDelegate
type FailingNoop = latest.FailingNoop
type Preendorsement = latest.Preattestation
type Preattestation = latest.Preattestation
type InlinedPreendorsementContents = latest.InlinedPreattestationContents
type InlinedPreattestationContents = latest.InlinedPreattestationContents
type Endorsement = latest.Attestation
type Attestation = latest.Attestation
type InlinedEndorsementContents = latest.InlinedAttestationContents
type InlinedAttestationContents = latest.InlinedAttestationContents
type DALAttestation = latest.DALAttestation
type Reveal = latest.Reveal
type Transaction = latest.Transaction
type Parameters = latest.Parameters
type Origination = latest.Origination
type Delegation = latest.Delegation
type RegisterGlobalConstant = latest.RegisterGlobalConstant
type SetDepositsLimit = latest.SetDepositsLimit
type IncreasePaidStorage = latest.IncreasePaidStorage
type UpdateConsensusKey = latest.UpdateConsensusKey
type TransferTicket = latest.TransferTicket
type SmartRollupOriginate = latest.SmartRollupOriginate
type SmartRollupAddMessages = latest.SmartRollupAddMessages
type SmartRollupCement = latest.SmartRollupCement
type SmartRollupPublish = latest.SmartRollupPublish
type SmartRollupRefute = latest.SmartRollupRefute
type SmartRollupTimeout = latest.SmartRollupTimeout
type SmartRollupExecuteOutboxMessage = latest.SmartRollupExecuteOutboxMessage
type SmartRollupRecoverBond = latest.SmartRollupRecoverBond
type DALPublishSlotHeader = latest.DALPublishSlotHeader
type ZkRollupOrigination = latest.ZkRollupOrigination
type ZkRollupPublish = latest.ZkRollupPublish
type ZkRollupUpdate = latest.ZkRollupUpdate
type SignaturePrefix = latest.SignaturePrefix
type BLSSignaturePrefix = latest.BLSSignaturePrefix

type EpDefault = latest.EpDefault
type EpRoot = latest.EpRoot
type EpDo = latest.EpDo
type EpSetDelegate = latest.EpSetDelegate
type EpRemoveDelegate = latest.EpRemoveDelegate
type EpDeposit = latest.EpDeposit
type EpStake = latest.EpStake
type EpUnstake = latest.EpUnstake
type EpFinalizeUnstake = latest.EpFinalizeUnstake
type EpSetDelegateParameters = latest.EpSetDelegateParameters
type EpNamed = latest.EpNamed
