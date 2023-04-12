package operations

import tz "github.com/ecadlabs/gotez"

type UpdateConsensusKey struct {
	ManagerOperation
	PublicKey tz.PublicKey
}

func (*UpdateConsensusKey) OperationKind() string { return "update_consensus_key" }

type UpdateConsensusKeyContentsAndResult struct {
	UpdateConsensusKey
	Metadata MetadataWithResult[EventResult]
}

func (*UpdateConsensusKeyContentsAndResult) OperationContentsAndResult() {}

type UpdateConsensusKeySuccessfulManagerOperationResult EventResultContents

func (*UpdateConsensusKeySuccessfulManagerOperationResult) SuccessfulManagerOperationResult() {}
func (*UpdateConsensusKeySuccessfulManagerOperationResult) OperationKind() string {
	return "update_consensus_key"
}
