package proto

type OperationContents interface {
	OperationKind() string
}

type OperationContentsAndResult interface {
	OperationContentsAndResult()
	OperationContents
}

type SuccessfulManagerOperationResult interface {
	OperationContents
	SuccessfulManagerOperationResult()
}

type InternalOperationResult interface {
	OperationContents
	InternalOperationResult()
}
