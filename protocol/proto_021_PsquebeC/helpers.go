package proto_021_PsquebeC

import (
	"github.com/ecadlabs/gotez/v2/protocol/proto_018_Proxford"
	"github.com/ecadlabs/gotez/v2/protocol/proto_019_PtParisB"
)

type UnsignedOperation = proto_018_Proxford.UnsignedOperationImpl[OperationContents]
type SignedOperation = proto_018_Proxford.SignedOperationImpl[OperationContents]
type RunOperationRequest = proto_018_Proxford.RunOperationRequestImpl[RunOperationRequestContents]
type RunOperationRequestContents = proto_019_PtParisB.RunOperationRequestContents

var NewRunOperationRequest = proto_019_PtParisB.NewRunOperationRequest
var NewUnsignedOperation = proto_019_PtParisB.NewUnsignedOperation
var NewSignedOperation = proto_019_PtParisB.NewSignedOperation
