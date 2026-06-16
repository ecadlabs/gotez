package proto_025_PsUshuai

import "encoding/json"

func (self *DALEntrapmentEvidence) MarshalJSON() ([]byte, error) {
	type DALEntrapmentEvidenceNoJSONMarshaller DALEntrapmentEvidence

	type jsonDALEntrapmentEvidence struct {
		Marker0 any `json:"kind"`
		DALEntrapmentEvidenceNoJSONMarshaller
	}

	tmp := jsonDALEntrapmentEvidence{
		Marker0:                               self.OperationKind(),
		DALEntrapmentEvidenceNoJSONMarshaller: DALEntrapmentEvidenceNoJSONMarshaller(*self),
	}

	return json.Marshal(tmp)
}

func (self *DALEntrapmentEvidenceContentsAndResult) MarshalJSON() ([]byte, error) {
	type DALEntrapmentEvidenceContentsAndResultNoJSONMarshaller DALEntrapmentEvidenceContentsAndResult

	type jsonDALEntrapmentEvidenceContentsAndResult struct {
		Marker0 any `json:"kind"`
		DALEntrapmentEvidenceContentsAndResultNoJSONMarshaller
	}

	tmp := jsonDALEntrapmentEvidenceContentsAndResult{
		Marker0: self.OperationKind(),
		DALEntrapmentEvidenceContentsAndResultNoJSONMarshaller: DALEntrapmentEvidenceContentsAndResultNoJSONMarshaller(*self),
	}

	return json.Marshal(tmp)
}
