package operations

import tz "github.com/ecadlabs/gotez"

type Reveal struct {
	ManagerOperation
	PublicKey tz.PublicKey
}

func (*Reveal) OperationKind() string { return "reveal" }

type RevealContentsAndResult struct {
	Reveal
	Metadata MetadataWithResult[EventResult]
}

func (*RevealContentsAndResult) OperationContentsAndResult() {}

type RevealSuccessfulManagerOperationResult EventResultContents

func (*RevealSuccessfulManagerOperationResult) SuccessfulManagerOperationResult() {}
func (*RevealSuccessfulManagerOperationResult) OperationKind() string             { return "reveal" }
