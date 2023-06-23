package latest

import (
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/proto_016_PtMumbai"
)

var Protocol = core.Proto016PtMumbai

type OperationContents = proto_016_PtMumbai.OperationContents
type OperationContentsAndResult = proto_016_PtMumbai.OperationContentsAndResult
type OperationWithOptionalMetadataContents = proto_016_PtMumbai.OperationWithOptionalMetadataContents
type RunOperationRequest = proto_016_PtMumbai.RunOperationRequest
type UnsignedOperation = proto_016_PtMumbai.UnsignedOperation
type SignedOperation = proto_016_PtMumbai.SignedOperation

type UnsignedBlockHeader = proto_016_PtMumbai.UnsignedBlockHeader
type BlockHeader = proto_016_PtMumbai.BlockHeader

type ManagerOperation = proto_016_PtMumbai.ManagerOperation
type SeedNonceRevelation = proto_016_PtMumbai.SeedNonceRevelation
type DoubleEndorsementEvidence = proto_016_PtMumbai.DoubleEndorsementEvidence
type DoubleBakingEvidence = proto_016_PtMumbai.DoubleBakingEvidence
type ActivateAccount = proto_016_PtMumbai.ActivateAccount
type Proposals = proto_016_PtMumbai.Proposals
type Ballot = proto_016_PtMumbai.Ballot
type DoublePreendorsementEvidence = proto_016_PtMumbai.DoublePreendorsementEvidence
type VDFRevelation = proto_016_PtMumbai.VDFRevelation
type DrainDelegate = proto_016_PtMumbai.DrainDelegate
type FailingNoop = proto_016_PtMumbai.FailingNoop
type Preendorsement = proto_016_PtMumbai.Preendorsement
type InlinedPreendorsementContents = proto_016_PtMumbai.InlinedPreendorsementContents
type Endorsement = proto_016_PtMumbai.Endorsement
type InlinedEndorsementContents = proto_016_PtMumbai.InlinedEndorsementContents
type DALAttestation = proto_016_PtMumbai.DALAttestation
type Reveal = proto_016_PtMumbai.Reveal
type Transaction = proto_016_PtMumbai.Transaction
type Parameters = proto_016_PtMumbai.Parameters
type Origination = proto_016_PtMumbai.Origination
type Delegation = proto_016_PtMumbai.Delegation
type RegisterGlobalConstant = proto_016_PtMumbai.RegisterGlobalConstant
type SetDepositsLimit = proto_016_PtMumbai.SetDepositsLimit
type IncreasePaidStorage = proto_016_PtMumbai.IncreasePaidStorage
type UpdateConsensusKey = proto_016_PtMumbai.UpdateConsensusKey
type TransferTicket = proto_016_PtMumbai.TransferTicket
type SmartRollupOriginate = proto_016_PtMumbai.SmartRollupOriginate
type SmartRollupAddMessages = proto_016_PtMumbai.SmartRollupAddMessages
type SmartRollupCement = proto_016_PtMumbai.SmartRollupCement
type SmartRollupPublish = proto_016_PtMumbai.SmartRollupPublish
type SmartRollupRefute = proto_016_PtMumbai.SmartRollupRefute
type SmartRollupTimeout = proto_016_PtMumbai.SmartRollupTimeout
type SmartRollupExecuteOutboxMessage = proto_016_PtMumbai.SmartRollupExecuteOutboxMessage
type SmartRollupRecoverBond = proto_016_PtMumbai.SmartRollupRecoverBond
type DALPublishSlotHeader = proto_016_PtMumbai.DALPublishSlotHeader
type ZkRollupOrigination = proto_016_PtMumbai.ZkRollupOrigination
type ZkRollupPublish = proto_016_PtMumbai.ZkRollupPublish
type ZkRollupUpdate = proto_016_PtMumbai.ZkRollupUpdate
type SignaturePrefix = proto_016_PtMumbai.SignaturePrefix
