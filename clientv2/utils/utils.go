package utils

import (
	"strconv"

	client "github.com/ecadlabs/gotez/v2/clientv2"
)

//go:generate go run generate.go

type BootstrappedResponse struct {
	Bootstrapped bool
	SyncState    SyncState
}

type SyncState uint8

const (
	SyncStateSynced SyncState = iota
	SyncStateUnsynced
	SyncStateStuck
)

func (state SyncState) String() string {
	switch state {
	case SyncStateSynced:
		return "synced"
	case SyncStateUnsynced:
		return "unsynced"
	case SyncStateStuck:
		return "stuck"
	default:
		return strconv.FormatUint(uint64(state), 10)
	}
}

type InjectOperationRequest struct {
	Chain   string
	Async   client.Flag
	Payload *InjectRequestPayload
}

type InjectRequestPayload struct {
	Contents []byte `tz:"dyn"`
}
