package proto_025_PsUshuai

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/proto_023_PtSeouLo"
)

type ShardWithProof = proto_023_PtSeouLo.ShardWithProof

type OperationContents interface {
	core.OperationContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[OperationContents]{
		Variants: encoding.Variants[OperationContents]{
			1:   (*SeedNonceRevelation)(nil),
			2:   (*DoubleConsensusOperationEvidence)(nil),
			3:   (*DoubleBakingEvidence)(nil),
			4:   (*ActivateAccount)(nil),
			5:   (*Proposals)(nil),
			6:   (*Ballot)(nil),
			8:   (*VDFRevelation)(nil),
			9:   (*DrainDelegate)(nil),
			17:  (*FailingNoop)(nil),
			20:  (*Preattestation)(nil),
			21:  (*Attestation)(nil),
			23:  (*AttestationWithDAL)(nil),
			24:  (*DALEntrapmentEvidence)(nil),
			30:  (*PreattestationsAggregate)(nil),
			31:  (*AttestationsAggregate)(nil),
			40:  (*BLSModePreattestation)(nil),
			41:  (*BLSModeAttestation)(nil),
			107: (*Reveal)(nil),
			108: (*Transaction)(nil),
			109: (*Origination)(nil),
			110: (*Delegation)(nil),
			111: (*RegisterGlobalConstant)(nil),
			112: (*SetDepositsLimit)(nil),
			113: (*IncreasePaidStorage)(nil),
			114: (*UpdateConsensusKey)(nil),
			115: (*UpdateCompanionKey)(nil),
			158: (*TransferTicket)(nil),
			200: (*SmartRollupOriginate)(nil),
			201: (*SmartRollupAddMessages)(nil),
			202: (*SmartRollupCement)(nil),
			203: (*SmartRollupPublish)(nil),
			204: (*SmartRollupRefute)(nil),
			205: (*SmartRollupTimeout)(nil),
			206: (*SmartRollupExecuteOutboxMessage)(nil),
			207: (*SmartRollupRecoverBond)(nil),
			230: (*DALPublishCommitment)(nil),
			250: (*ZkRollupOrigination)(nil),
			251: (*ZkRollupPublish)(nil),
			252: (*ZkRollupUpdate)(nil),
			255: (*core.SignaturePrefix)(nil),
		},
	})
}

//json:kind=OperationKind()
type DALEntrapmentEvidence struct {
	Attestation    InlinedConsensusOperation `tz:"dyn" json:"attestation"`
	ConsensusSlot  uint16                    `json:"consensus_slot"`
	SlotIndex      uint8                     `json:"slot_index"`
	LagIndex       tz.Option[uint8]          `json:"lag_index"`
	ShardWithProof ShardWithProof            `json:"shard_with_proof"`
}

func (*DALEntrapmentEvidence) OperationKind() string { return "dal_entrapment_evidence" }

type OperationContentsAndResult interface {
	core.OperationContentsAndResult
}

//json:kind=OperationKind()
type DALEntrapmentEvidenceContentsAndResult struct {
	DALEntrapmentEvidence
	Metadata BalanceUpdates `json:"metadata"`
}

func (*DALEntrapmentEvidenceContentsAndResult) OperationContentsAndResult() {}
func (op *DALEntrapmentEvidenceContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[OperationContentsAndResult]{
		Variants: encoding.Variants[OperationContentsAndResult]{
			1:   (*SeedNonceRevelationContentsAndResult)(nil),
			2:   (*DoubleConsensusOperationEvidenceContentsAndResult)(nil),
			3:   (*DoubleBakingEvidenceContentsAndResult)(nil),
			4:   (*ActivateAccountContentsAndResult)(nil),
			5:   (*Proposals)(nil),
			6:   (*Ballot)(nil),
			8:   (*VDFRevelationContentsAndResult)(nil),
			9:   (*DrainDelegateContentsAndResult)(nil),
			20:  (*PreattestationContentsAndResult)(nil),
			21:  (*AttestationContentsAndResult)(nil),
			23:  (*AttestationWithDALContentsAndResult)(nil),
			24:  (*DALEntrapmentEvidenceContentsAndResult)(nil),
			30:  (*PreattestationsAggregateContentsAndResult)(nil),
			31:  (*AttestationsAggregateContentsAndResult)(nil),
			107: (*RevealContentsAndResult)(nil),
			108: (*TransactionContentsAndResult)(nil),
			109: (*OriginationContentsAndResult)(nil),
			110: (*DelegationContentsAndResult)(nil),
			111: (*RegisterGlobalConstantContentsAndResult)(nil),
			112: (*SetDepositsLimitContentsAndResult)(nil),
			113: (*IncreasePaidStorageContentsAndResult)(nil),
			114: (*UpdateConsensusKeyContentsAndResult)(nil),
			115: (*UpdateCompanionKeyContentsAndResult)(nil),
			158: (*TransferTicketContentsAndResult)(nil),
			200: (*SmartRollupOriginateContentsAndResult)(nil),
			201: (*SmartRollupAddMessagesContentsAndResult)(nil),
			202: (*SmartRollupCementContentsAndResult)(nil),
			203: (*SmartRollupPublishContentsAndResult)(nil),
			204: (*SmartRollupRefuteContentsAndResult)(nil),
			205: (*SmartRollupTimeoutContentsAndResult)(nil),
			206: (*SmartRollupExecuteOutboxMessageContentsAndResult)(nil),
			207: (*SmartRollupRecoverBondContentsAndResult)(nil),
			230: (*DALPublishCommitmentContentsAndResult)(nil),
			250: (*ZkRollupOriginationContentsAndResult)(nil),
			251: (*ZkRollupPublishContentsAndResult)(nil),
			252: (*ZkRollupUpdateContentsAndResult)(nil),
			255: (*core.SignaturePrefix)(nil),
		},
	})
}

type OperationWithOptionalMetadata = core.OperationWithOptionalMetadata[OperationWithOptionalMetadataContents]

type GroupContents interface {
	core.GroupContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[GroupContents]{
		Variants: encoding.Variants[GroupContents]{
			0: (*core.OperationWithTooLargeMetadata[OperationContents])(nil),
			1: (*core.OperationWithoutMetadata[OperationContents])(nil),
			2: (*OperationWithOptionalMetadata)(nil),
		},
	})
}

type OperationWithOptionalMetadataContents interface {
	core.OperationWithOptionalMetadataContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[OperationWithOptionalMetadataContents]{
		Variants: encoding.Variants[OperationWithOptionalMetadataContents]{
			0: (*core.OperationWithOptionalMetadataWithMetadata[OperationContentsAndResult])(nil),
			1: (*core.OperationWithOptionalMetadataWithoutMetadata[OperationContents])(nil),
		},
	})
}

var ListOperations = core.ListOperations[OperationContents]
