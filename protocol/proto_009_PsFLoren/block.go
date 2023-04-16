package proto_009_PsFLoren

type VotingPeriodInfo struct {
	VotingPeriod VotingPeriod
	Position     int32
	Remaining    int32
}

type VotingPeriod struct {
	Index         int32
	Kind          VotingPeriodKind
	StartPosition int32
}

type VotingPeriodKind uint8

const (
	VotingPeriodProposal VotingPeriodKind = iota
	VotingPeriodExploration
	VotingPeriodCooldown
	VotingPeriodPromotion
	VotingPeriodAdoption
)
