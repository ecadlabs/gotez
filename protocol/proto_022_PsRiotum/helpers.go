package proto_022_PsRiotum

import (
	"github.com/ecadlabs/gotez/v2/protocol/proto_018_Proxford"
	"github.com/ecadlabs/gotez/v2/protocol/proto_019_PtParisB"
	"github.com/ecadlabs/gotez/v2/protocol/proto_021_PsQuebec"
)

type UnsignedOperation = proto_018_Proxford.UnsignedOperationImpl[OperationContents]
type SignedOperation = proto_018_Proxford.SignedOperationImpl[OperationContents]
type RunOperationRequest = proto_018_Proxford.RunOperationRequestImpl[RunOperationRequestContents]
type RunOperationRequestContents = proto_019_PtParisB.RunOperationRequestContents

var NewRunOperationRequest = proto_021_PsQuebec.NewRunOperationRequest
var NewUnsignedOperation = proto_021_PsQuebec.NewUnsignedOperation
var NewSignedOperation = proto_021_PsQuebec.NewSignedOperation
