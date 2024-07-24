package network

//go:generate go run generate.go

import tz "github.com/ecadlabs/gotez/v2"

type ConnectionsResponse struct {
	Connections []*Connection `tz:"dyn" json:"connections"`
}

type Connection struct {
	Incoming         bool           `json:"incoming"`
	PeerID           [16]byte       `json:"peer_id"`
	IDPoint          ConnectionID   `json:"id_point"`
	RemoteSocketPort uint16         `json:"remote_socket_port"`
	AnnouncedVersion NetworkVersion `json:"announced_version"`
	Private          bool           `json:"private"`
	LocalMetadata    Metadata       `json:"local_metadata"`
	RemoteMetadata   Metadata       `json:"remote_metadata"`
}

type Metadata struct {
	DisableMempool bool `json:"disable_mempool"`
	PrivateNode    bool `json:"private_node"`
}

type ConnectionID struct {
	Addr []byte            `tz:"dyn" json:"addr"`
	Port tz.Option[uint16] `json:"port"`
}

type NetworkVersion struct {
	ChainName            string `tz:"dyn" json:"chain_name"`
	DistributedDBVersion uint16 `json:"distributed_db_version"`
	P2PVersion           uint16 `json:"p2p_version"`
}
