package proto_009_PsFLoren

type FailingNoop struct {
	Arbitrary []byte `tz:"dyn"`
}

func (*FailingNoop) OperationKind() string { return "failing_noop" }
