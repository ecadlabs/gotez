package latest

import (
	"github.com/ecadlabs/gotez/v2/protocol/core"
	latest "github.com/ecadlabs/gotez/v2/protocol/proto_018_Proxford"
)

var (
	Protocol               = core.Proto017PtNairob
	NewRunOperationRequest = latest.NewRunOperationRequest
)

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
type SeedNonceRevelationContentsAndResult = latest.SeedNonceRevelationContentsAndResult
type DoubleAttestationEvidence = latest.DoubleAttestationEvidence
type DoubleAttestationEvidenceContentsAndResult = latest.DoubleAttestationEvidenceContentsAndResult
type DoubleEndorsementEvidence = latest.DoubleAttestationEvidence
type DoubleEndorsementEvidenceContentsAndResult = latest.DoubleAttestationEvidenceContentsAndResult
type DoubleBakingEvidence = latest.DoubleBakingEvidence
type DoubleBakingEvidenceContentsAndResult = latest.DoubleBakingEvidenceContentsAndResult
type ActivateAccount = latest.ActivateAccount
type ActivateAccountContentsAndResult = latest.ActivateAccountContentsAndResult
type Proposals = latest.Proposals
type Ballot = latest.Ballot
type DoublePreattestationEvidence = latest.DoublePreattestationEvidence
type DoublePreattestationEvidenceContentsAndResult = latest.DoublePreattestationEvidenceContentsAndResult
type DoublePreendorsementEvidence = latest.DoublePreattestationEvidence
type DoublePreendorsementEvidenceContentsAndResult = latest.DoublePreattestationEvidenceContentsAndResult
type VDFRevelation = latest.VDFRevelation
type VDFRevelationContentsAndResult = latest.VDFRevelationContentsAndResult
type DrainDelegate = latest.DrainDelegate
type DrainDelegateContentsAndResult = latest.DrainDelegateContentsAndResult
type FailingNoop = latest.FailingNoop
type Preattestation = latest.Preattestation
type PreattestationContentsAndResult = latest.PreattestationContentsAndResult
type Preendorsement = latest.Preattestation
type PreendorsementContentsAndResult = latest.PreattestationContentsAndResult
type InlinedPreendorsementContents = latest.InlinedPreattestationContents
type InlinedPreattestationContents = latest.InlinedPreattestationContents
type Attestation = latest.Attestation
type AttestationContentsAndResult = latest.AttestationContentsAndResult
type Endorsement = latest.Attestation
type EndorsementContentsAndResult = latest.AttestationContentsAndResult
type InlinedEndorsementContents = latest.InlinedAttestationContents
type InlinedAttestationContents = latest.InlinedAttestationContents
type DALAttestation = latest.DALAttestation
type DALAttestationContentsAndResult = latest.DALAttestationContentsAndResult
type Reveal = latest.Reveal
type RevealContentsAndResult = latest.RevealContentsAndResult
type Transaction = latest.Transaction
type TransactionContentsAndResult = latest.TransactionContentsAndResult
type TransactionResultDestination = latest.TransactionResultDestination
type Parameters = latest.Parameters
type Origination = latest.Origination
type OriginationContentsAndResult = latest.OriginationContentsAndResult
type Delegation = latest.Delegation
type DelegationContentsAndResult = latest.DelegationContentsAndResult
type RegisterGlobalConstant = latest.RegisterGlobalConstant
type RegisterGlobalConstantContentsAndResult = latest.RegisterGlobalConstantContentsAndResult
type SetDepositsLimit = latest.SetDepositsLimit
type SetDepositsLimitContentsAndResult = latest.SetDepositsLimitContentsAndResult
type IncreasePaidStorage = latest.IncreasePaidStorage
type IncreasePaidStorageContentsAndResult = latest.IncreasePaidStorageContentsAndResult
type UpdateConsensusKey = latest.UpdateConsensusKey
type UpdateConsensusKeyContentsAndResult = latest.UpdateConsensusKeyContentsAndResult
type TransferTicket = latest.TransferTicket
type TransferTicketContentsAndResult = latest.TransferTicketContentsAndResult
type SmartRollupOriginate = latest.SmartRollupOriginate
type SmartRollupOriginateContentsAndResult = latest.SmartRollupOriginateContentsAndResult
type SmartRollupAddMessages = latest.SmartRollupAddMessages
type SmartRollupAddMessagesContentsAndResult = latest.SmartRollupAddMessagesContentsAndResult
type SmartRollupCement = latest.SmartRollupCement
type SmartRollupCementContentsAndResult = latest.SmartRollupCementContentsAndResult
type SmartRollupPublish = latest.SmartRollupPublish
type SmartRollupPublishContentsAndResult = latest.SmartRollupPublishContentsAndResult
type SmartRollupRefute = latest.SmartRollupRefute
type SmartRollupRefuteContentsAndResult = latest.SmartRollupRefuteContentsAndResult
type SmartRollupTimeout = latest.SmartRollupTimeout
type SmartRollupTimeoutContentsAndResult = latest.SmartRollupTimeoutContentsAndResult
type SmartRollupExecuteOutboxMessage = latest.SmartRollupExecuteOutboxMessage
type SmartRollupExecuteOutboxMessageContentsAndResult = latest.SmartRollupExecuteOutboxMessageContentsAndResult
type SmartRollupRecoverBond = latest.SmartRollupRecoverBond
type SmartRollupRecoverBondContentsAndResult = latest.SmartRollupRecoverBondContentsAndResult
type DALPublishSlotHeader = latest.DALPublishSlotHeader
type DALPublishSlotHeaderContentsAndResult = latest.DALPublishSlotHeaderContentsAndResult
type ZkRollupOrigination = latest.ZkRollupOrigination
type ZkRollupOriginationContentsAndResult = latest.ZkRollupOriginationContentsAndResult
type ZkRollupPublish = latest.ZkRollupPublish
type ZkRollupPublishContentsAndResult = latest.ZkRollupPublishContentsAndResult
type ZkRollupUpdate = latest.ZkRollupUpdate
type ZkRollupUpdateContentsAndResult = latest.ZkRollupUpdateContentsAndResult
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

type ToContract = latest.ToContract
type ToSmartRollup = latest.ToSmartRollup
