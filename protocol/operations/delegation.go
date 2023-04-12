package operations

import tz "github.com/ecadlabs/gotez"

type Delegation struct {
	ManagerOperation
	Delegate tz.Option[tz.PublicKeyHash]
}

func (*Delegation) OperationKind() string { return "delegation" }

type DelegationContentsAndResult struct {
	Delegation
	Metadata MetadataWithResult[EventResult]
}

func (*DelegationContentsAndResult) OperationContentsAndResult() {}

type DelegationInternalOperationResult struct {
	Source   tz.TransactionDestination
	Nonce    uint16
	Delegate tz.Option[tz.PublicKeyHash]
	Result   EventResult
}

func (*DelegationInternalOperationResult) InternalOperationResult() {}

type DelegationSuccessfulManagerOperationResult EventResultContents

func (*DelegationSuccessfulManagerOperationResult) SuccessfulManagerOperationResult() {}
func (*DelegationSuccessfulManagerOperationResult) OperationKind() string             { return "delegation" }
