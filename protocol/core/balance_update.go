package core

type BalanceUpdateOrigin uint8

const (
	BalanceUpdateOriginBlockApplication BalanceUpdateOrigin = iota
	BalanceUpdateOriginProtocolMigration
	BalanceUpdateOriginSubsidy
	BalanceUpdateOriginSimulation
)

type BalanceUpdate interface {
	GetKind() BalanceUpdateKind
	GetChange() int64
	GetOrigin() BalanceUpdateOrigin
}

type BalanceUpdateKind interface {
	BalanceUpdateKind() string
}

type BalanceUpdates interface {
	GetBalanceUpdates() []BalanceUpdate
}
