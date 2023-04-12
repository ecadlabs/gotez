package operations

import tz "github.com/ecadlabs/gotez"

type SetDepositsLimit struct {
	ManagerOperation
	Limit tz.Option[tz.BigUint]
}

func (*SetDepositsLimit) OperationKind() string { return "set_deposits_limit" }

type SetDepositsLimitContentsAndResult struct {
	SetDepositsLimit
	Metadata MetadataWithResult[EventResult]
}

func (*SetDepositsLimitContentsAndResult) OperationContentsAndResult() {}

type SetDepositsLimitSuccessfulManagerOperationResult EventResultContents

func (*SetDepositsLimitSuccessfulManagerOperationResult) SuccessfulManagerOperationResult() {}
func (*SetDepositsLimitSuccessfulManagerOperationResult) OperationKind() string {
	return "set_deposits_limit"
}
