package latest

import (
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/proto_017_PtNairob"
)

var Protocol = core.Proto017PtNairob

type OperationContents = proto_017_PtNairob.OperationContents
type OperationContentsAndResult = proto_017_PtNairob.OperationContentsAndResult
type OperationWithOptionalMetadata = proto_017_PtNairob.OperationWithOptionalMetadata

type RunOperationRequest = proto_017_PtNairob.RunOperationRequest
type UnsignedOperation = proto_017_PtNairob.UnsignedOperation
type SignedOperation = proto_017_PtNairob.SignedOperation

type UnsignedBlockHeader = proto_017_PtNairob.UnsignedBlockHeader
type BlockHeader = proto_017_PtNairob.BlockHeader

type ManagerOperation = proto_017_PtNairob.ManagerOperation
type SeedNonceRevelation = proto_017_PtNairob.SeedNonceRevelation
type DoubleEndorsementEvidence = proto_017_PtNairob.DoubleEndorsementEvidence
type DoubleBakingEvidence = proto_017_PtNairob.DoubleBakingEvidence
type ActivateAccount = proto_017_PtNairob.ActivateAccount
type Proposals = proto_017_PtNairob.Proposals
type Ballot = proto_017_PtNairob.Ballot
type DoublePreendorsementEvidence = proto_017_PtNairob.DoublePreendorsementEvidence
type VDFRevelation = proto_017_PtNairob.VDFRevelation
type DrainDelegate = proto_017_PtNairob.DrainDelegate
type FailingNoop = proto_017_PtNairob.FailingNoop
type Preendorsement = proto_017_PtNairob.Preendorsement
type InlinedPreendorsementContents = proto_017_PtNairob.InlinedPreendorsementContents
type Endorsement = proto_017_PtNairob.Endorsement
type InlinedEndorsementContents = proto_017_PtNairob.InlinedEndorsementContents
type DALAttestation = proto_017_PtNairob.DALAttestation
type Reveal = proto_017_PtNairob.Reveal
type Transaction = proto_017_PtNairob.Transaction
type Parameters = proto_017_PtNairob.Parameters
type Origination = proto_017_PtNairob.Origination
type Delegation = proto_017_PtNairob.Delegation
type RegisterGlobalConstant = proto_017_PtNairob.RegisterGlobalConstant
type SetDepositsLimit = proto_017_PtNairob.SetDepositsLimit
type IncreasePaidStorage = proto_017_PtNairob.IncreasePaidStorage
type UpdateConsensusKey = proto_017_PtNairob.UpdateConsensusKey
type TransferTicket = proto_017_PtNairob.TransferTicket
type SmartRollupOriginate = proto_017_PtNairob.SmartRollupOriginate
type SmartRollupAddMessages = proto_017_PtNairob.SmartRollupAddMessages
type SmartRollupCement = proto_017_PtNairob.SmartRollupCement
type SmartRollupPublish = proto_017_PtNairob.SmartRollupPublish
type SmartRollupRefute = proto_017_PtNairob.SmartRollupRefute
type SmartRollupTimeout = proto_017_PtNairob.SmartRollupTimeout
type SmartRollupExecuteOutboxMessage = proto_017_PtNairob.SmartRollupExecuteOutboxMessage
type SmartRollupRecoverBond = proto_017_PtNairob.SmartRollupRecoverBond
type DALPublishSlotHeader = proto_017_PtNairob.DALPublishSlotHeader
type ZkRollupOrigination = proto_017_PtNairob.ZkRollupOrigination
type ZkRollupPublish = proto_017_PtNairob.ZkRollupPublish
type ZkRollupUpdate = proto_017_PtNairob.ZkRollupUpdate
type SignaturePrefix = proto_017_PtNairob.SignaturePrefix
