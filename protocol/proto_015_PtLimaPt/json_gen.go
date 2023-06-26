package proto_015_PtLimaPt

import "encoding/json"

// Code generated by genmarshaller.go DO NOT EDIT.

func (self *OriginationSuccessfulManagerResult) MarshalJSON() ([]byte, error) {
	type OriginationSuccessfulManagerResult_no_json_marshaller OriginationSuccessfulManagerResult

	type json_OriginationSuccessfulManagerResult struct {
		Marker0 any `json:"kind"`
		*OriginationSuccessfulManagerResult_no_json_marshaller
	}

	tmp := json_OriginationSuccessfulManagerResult {
		Marker0: self.OperationKind(),
		OriginationSuccessfulManagerResult_no_json_marshaller: (*OriginationSuccessfulManagerResult_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *OriginationInternalOperationResult) MarshalJSON() ([]byte, error) {
	type OriginationInternalOperationResult_no_json_marshaller OriginationInternalOperationResult

	type json_OriginationInternalOperationResult struct {
		Marker0 any `json:"kind"`
		*OriginationInternalOperationResult_no_json_marshaller
	}

	tmp := json_OriginationInternalOperationResult {
		Marker0: self.OperationKind(),
		OriginationInternalOperationResult_no_json_marshaller: (*OriginationInternalOperationResult_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *ZkRollupOrigination) MarshalJSON() ([]byte, error) {
	type ZkRollupOrigination_no_json_marshaller ZkRollupOrigination

	type json_ZkRollupOrigination struct {
		Marker0 any `json:"kind"`
		*ZkRollupOrigination_no_json_marshaller
	}

	tmp := json_ZkRollupOrigination {
		Marker0: self.OperationKind(),
		ZkRollupOrigination_no_json_marshaller: (*ZkRollupOrigination_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *ZkRollupPublish) MarshalJSON() ([]byte, error) {
	type ZkRollupPublish_no_json_marshaller ZkRollupPublish

	type json_ZkRollupPublish struct {
		Marker0 any `json:"kind"`
		*ZkRollupPublish_no_json_marshaller
	}

	tmp := json_ZkRollupPublish {
		Marker0: self.OperationKind(),
		ZkRollupPublish_no_json_marshaller: (*ZkRollupPublish_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *BalanceUpdateScRollupRefutationRewards) MarshalJSON() ([]byte, error) {
	type BalanceUpdateScRollupRefutationRewards_no_json_marshaller BalanceUpdateScRollupRefutationRewards

	type json_BalanceUpdateScRollupRefutationRewards struct {
		Marker0 any `json:"category"`
		Marker1 any `json:"kind"`
		*BalanceUpdateScRollupRefutationRewards_no_json_marshaller
	}

	tmp := json_BalanceUpdateScRollupRefutationRewards {
		Marker0: self.BalanceUpdateCategory(),
		Marker1: self.BalanceUpdateKind(),
		BalanceUpdateScRollupRefutationRewards_no_json_marshaller: (*BalanceUpdateScRollupRefutationRewards_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *UpdateConsensusKey) MarshalJSON() ([]byte, error) {
	type UpdateConsensusKey_no_json_marshaller UpdateConsensusKey

	type json_UpdateConsensusKey struct {
		Marker0 any `json:"kind"`
		*UpdateConsensusKey_no_json_marshaller
	}

	tmp := json_UpdateConsensusKey {
		Marker0: self.OperationKind(),
		UpdateConsensusKey_no_json_marshaller: (*UpdateConsensusKey_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *DrainDelegate) MarshalJSON() ([]byte, error) {
	type DrainDelegate_no_json_marshaller DrainDelegate

	type json_DrainDelegate struct {
		Marker0 any `json:"kind"`
		*DrainDelegate_no_json_marshaller
	}

	tmp := json_DrainDelegate {
		Marker0: self.OperationKind(),
		DrainDelegate_no_json_marshaller: (*DrainDelegate_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *UpdateConsensusKeySuccessfulManagerResult) MarshalJSON() ([]byte, error) {
	type UpdateConsensusKeySuccessfulManagerResult_no_json_marshaller UpdateConsensusKeySuccessfulManagerResult

	type json_UpdateConsensusKeySuccessfulManagerResult struct {
		Marker0 any `json:"kind"`
		*UpdateConsensusKeySuccessfulManagerResult_no_json_marshaller
	}

	tmp := json_UpdateConsensusKeySuccessfulManagerResult {
		Marker0: self.OperationKind(),
		UpdateConsensusKeySuccessfulManagerResult_no_json_marshaller: (*UpdateConsensusKeySuccessfulManagerResult_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *DALPublishSlotHeader) MarshalJSON() ([]byte, error) {
	type DALPublishSlotHeader_no_json_marshaller DALPublishSlotHeader

	type json_DALPublishSlotHeader struct {
		Marker0 any `json:"kind"`
		*DALPublishSlotHeader_no_json_marshaller
	}

	tmp := json_DALPublishSlotHeader {
		Marker0: self.OperationKind(),
		DALPublishSlotHeader_no_json_marshaller: (*DALPublishSlotHeader_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *IncreasePaidStorageSuccessfulManagerResult) MarshalJSON() ([]byte, error) {
	type IncreasePaidStorageSuccessfulManagerResult_no_json_marshaller IncreasePaidStorageSuccessfulManagerResult

	type json_IncreasePaidStorageSuccessfulManagerResult struct {
		Marker0 any `json:"kind"`
		*IncreasePaidStorageSuccessfulManagerResult_no_json_marshaller
	}

	tmp := json_IncreasePaidStorageSuccessfulManagerResult {
		Marker0: self.OperationKind(),
		IncreasePaidStorageSuccessfulManagerResult_no_json_marshaller: (*IncreasePaidStorageSuccessfulManagerResult_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *DelegationInternalOperationResult) MarshalJSON() ([]byte, error) {
	type DelegationInternalOperationResult_no_json_marshaller DelegationInternalOperationResult

	type json_DelegationInternalOperationResult struct {
		Marker0 any `json:"kind"`
		*DelegationInternalOperationResult_no_json_marshaller
	}

	tmp := json_DelegationInternalOperationResult {
		Marker0: self.OperationKind(),
		DelegationInternalOperationResult_no_json_marshaller: (*DelegationInternalOperationResult_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *EventInternalOperationResult) MarshalJSON() ([]byte, error) {
	type EventInternalOperationResult_no_json_marshaller EventInternalOperationResult

	type json_EventInternalOperationResult struct {
		Marker0 any `json:"kind"`
		*EventInternalOperationResult_no_json_marshaller
	}

	tmp := json_EventInternalOperationResult {
		Marker0: self.OperationKind(),
		EventInternalOperationResult_no_json_marshaller: (*EventInternalOperationResult_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *Transaction) MarshalJSON() ([]byte, error) {
	type Transaction_no_json_marshaller Transaction

	type json_Transaction struct {
		Marker0 any `json:"kind"`
		*Transaction_no_json_marshaller
	}

	tmp := json_Transaction {
		Marker0: self.OperationKind(),
		Transaction_no_json_marshaller: (*Transaction_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *TransactionSuccessfulManagerResult) MarshalJSON() ([]byte, error) {
	type TransactionSuccessfulManagerResult_no_json_marshaller TransactionSuccessfulManagerResult

	type json_TransactionSuccessfulManagerResult struct {
		Marker0 any `json:"kind"`
		*TransactionSuccessfulManagerResult_no_json_marshaller
	}

	tmp := json_TransactionSuccessfulManagerResult {
		Marker0: self.OperationKind(),
		TransactionSuccessfulManagerResult_no_json_marshaller: (*TransactionSuccessfulManagerResult_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

func (self *TransactionInternalOperationResult) MarshalJSON() ([]byte, error) {
	type TransactionInternalOperationResult_no_json_marshaller TransactionInternalOperationResult

	type json_TransactionInternalOperationResult struct {
		Marker0 any `json:"kind"`
		*TransactionInternalOperationResult_no_json_marshaller
	}

	tmp := json_TransactionInternalOperationResult {
		Marker0: self.OperationKind(),
		TransactionInternalOperationResult_no_json_marshaller: (*TransactionInternalOperationResult_no_json_marshaller)(self),
	}

	return json.Marshal(&tmp)
}

