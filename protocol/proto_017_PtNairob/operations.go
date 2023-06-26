package proto_017_PtNairob

import (
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/proto_012_Psithaca"
	"github.com/ecadlabs/gotez/v2/protocol/proto_013_PtJakart"
	"github.com/ecadlabs/gotez/v2/protocol/proto_014_PtKathma"
	"github.com/ecadlabs/gotez/v2/protocol/proto_015_PtLimaPt"
	"github.com/ecadlabs/gotez/v2/protocol/proto_016_PtMumbai"
)

type ManagerOperation = proto_012_Psithaca.ManagerOperation
type SeedNonceRevelation = proto_012_Psithaca.SeedNonceRevelation
type SeedNonceRevelationContentsAndResult = proto_016_PtMumbai.SeedNonceRevelationContentsAndResult
type Preendorsement = proto_012_Psithaca.Preendorsement
type PreendorsementContentsAndResult = proto_016_PtMumbai.PreendorsementContentsAndResult
type InlinedPreendorsementContents = proto_012_Psithaca.InlinedPreendorsementContents
type Endorsement = proto_012_Psithaca.Endorsement
type EndorsementContentsAndResult = proto_016_PtMumbai.EndorsementContentsAndResult
type InlinedEndorsementContents = proto_012_Psithaca.InlinedEndorsementContents
type DoublePreendorsementEvidence = proto_012_Psithaca.DoublePreendorsementEvidence
type DoublePreendorsementEvidenceContentsAndResult = proto_016_PtMumbai.DoublePreendorsementEvidenceContentsAndResult
type DoubleEndorsementEvidence = proto_012_Psithaca.DoubleEndorsementEvidence
type DoubleEndorsementEvidenceContentsAndResult = proto_016_PtMumbai.DoubleEndorsementEvidenceContentsAndResult
type Reveal = proto_012_Psithaca.Reveal
type RevealContentsAndResult = proto_016_PtMumbai.RevealContentsAndResult
type Delegation = proto_012_Psithaca.Delegation
type DelegationContentsAndResult = proto_016_PtMumbai.DelegationContentsAndResult
type RegisterGlobalConstant = proto_012_Psithaca.RegisterGlobalConstant
type RegisterGlobalConstantContentsAndResult = proto_016_PtMumbai.RegisterGlobalConstantContentsAndResult
type SetDepositsLimit = proto_012_Psithaca.SetDepositsLimit
type SetDepositsLimitContentsAndResult = proto_016_PtMumbai.SetDepositsLimitContentsAndResult
type UpdateConsensusKey = proto_015_PtLimaPt.UpdateConsensusKey
type UpdateConsensusKeyContentsAndResult = proto_016_PtMumbai.UpdateConsensusKeyContentsAndResult
type IncreasePaidStorage = proto_014_PtKathma.IncreasePaidStorage
type IncreasePaidStorageContentsAndResult = proto_016_PtMumbai.IncreasePaidStorageContentsAndResult
type ActivateAccount = proto_012_Psithaca.ActivateAccount
type ActivateAccountContentsAndResult = proto_016_PtMumbai.ActivateAccountContentsAndResult
type Proposals = proto_012_Psithaca.Proposals
type Ballot = proto_012_Psithaca.Ballot
type VDFRevelation = proto_014_PtKathma.VDFRevelation
type VDFRevelationContentsAndResult = proto_016_PtMumbai.VDFRevelationContentsAndResult
type DrainDelegate = proto_015_PtLimaPt.DrainDelegate
type DrainDelegateContentsAndResult = proto_016_PtMumbai.DrainDelegateContentsAndResult
type FailingNoop = proto_012_Psithaca.FailingNoop
type TransferTicket = proto_013_PtJakart.TransferTicket
type TransferTicketContentsAndResult = proto_016_PtMumbai.TransferTicketContentsAndResult
type DALAttestation = proto_016_PtMumbai.DALAttestation
type DALAttestationContentsAndResult = proto_016_PtMumbai.DALAttestationContentsAndResult
type DoubleBakingEvidence = proto_016_PtMumbai.DoubleBakingEvidence
type DoubleBakingEvidenceContentsAndResult = proto_016_PtMumbai.DoubleBakingEvidenceContentsAndResult
type DALPublishSlotHeader = proto_016_PtMumbai.DALPublishSlotHeader
type DALPublishSlotHeaderContentsAndResult = proto_016_PtMumbai.DALPublishSlotHeaderContentsAndResult
type Origination = proto_012_Psithaca.Origination
type OriginationContentsAndResult = proto_016_PtMumbai.OriginationContentsAndResult
type Transaction = proto_015_PtLimaPt.Transaction
type TransactionContentsAndResult = proto_016_PtMumbai.TransactionContentsAndResult
type Parameters = proto_012_Psithaca.Parameters

type ZkRollupOrigination = proto_015_PtLimaPt.ZkRollupOrigination
type ZkRollupOriginationContentsAndResult = proto_016_PtMumbai.ZkRollupOriginationContentsAndResult
type ZkRollupPublish = proto_015_PtLimaPt.ZkRollupPublish
type ZkRollupPublishContentsAndResult = proto_016_PtMumbai.ZkRollupPublishContentsAndResult
type ZkRollupUpdate = proto_016_PtMumbai.ZkRollupUpdate
type ZkRollupUpdateContentsAndResult = proto_016_PtMumbai.ZkRollupUpdateContentsAndResult
type SignaturePrefix = proto_016_PtMumbai.SignaturePrefix
type BLSSignaturePrefix = proto_016_PtMumbai.BLSSignaturePrefix

type BalanceUpdates = proto_016_PtMumbai.BalanceUpdates
type InternalOperationResult = proto_016_PtMumbai.InternalOperationResult
type SuccessfulManagerOperationResult = proto_016_PtMumbai.SuccessfulManagerOperationResult
type OperationContents = proto_016_PtMumbai.OperationContents

type OperationContentsAndResult interface {
	core.OperationContentsAndResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[OperationContentsAndResult]{
		Variants: encoding.Variants[OperationContentsAndResult]{
			1:   (*SeedNonceRevelationContentsAndResult)(nil),
			2:   (*DoubleEndorsementEvidenceContentsAndResult)(nil),
			3:   (*DoubleBakingEvidenceContentsAndResult)(nil),
			4:   (*ActivateAccountContentsAndResult)(nil),
			5:   (*Proposals)(nil),
			6:   (*Ballot)(nil),
			7:   (*DoublePreendorsementEvidenceContentsAndResult)(nil),
			8:   (*VDFRevelationContentsAndResult)(nil),
			9:   (*DrainDelegateContentsAndResult)(nil),
			20:  (*PreendorsementContentsAndResult)(nil),
			21:  (*EndorsementContentsAndResult)(nil),
			22:  (*DALAttestationContentsAndResult)(nil),
			107: (*RevealContentsAndResult)(nil),
			108: (*TransactionContentsAndResult)(nil),
			109: (*OriginationContentsAndResult)(nil),
			110: (*DelegationContentsAndResult)(nil),
			111: (*RegisterGlobalConstantContentsAndResult)(nil),
			112: (*SetDepositsLimitContentsAndResult)(nil),
			113: (*IncreasePaidStorageContentsAndResult)(nil),
			114: (*UpdateConsensusKeyContentsAndResult)(nil),
			158: (*TransferTicketContentsAndResult)(nil),
			200: (*SmartRollupOriginateContentsAndResult)(nil),
			201: (*SmartRollupAddMessagesContentsAndResult)(nil),
			202: (*SmartRollupCementContentsAndResult)(nil),
			203: (*SmartRollupPublishContentsAndResult)(nil),
			204: (*SmartRollupRefuteContentsAndResult)(nil),
			205: (*SmartRollupTimeoutContentsAndResult)(nil),
			206: (*SmartRollupExecuteOutboxMessageContentsAndResult)(nil),
			207: (*SmartRollupRecoverBondContentsAndResult)(nil),
			230: (*DALPublishSlotHeaderContentsAndResult)(nil),
			250: (*ZkRollupOriginationContentsAndResult)(nil),
			251: (*ZkRollupPublishContentsAndResult)(nil),
			252: (*ZkRollupUpdateContentsAndResult)(nil),
			255: (*SignaturePrefix)(nil),
		},
	})
}

func ListOperations() []OperationContents {
	return encoding.ListVariants[OperationContents]()
}
