package proto_024_PtTALLiN

import (
	"fmt"
	"slices"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
	"github.com/ecadlabs/gotez/v2/protocol/core"
	"github.com/ecadlabs/gotez/v2/protocol/proto_023_PtSeouLo"
)

type GenericOperationSignRequestOperationContents interface {
	core.OperationContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[GenericOperationSignRequestOperationContents]{
		Variants: encoding.Variants[GenericOperationSignRequestOperationContents]{
			1:   (*SeedNonceRevelation)(nil),
			2:   (*DoubleConsensusOperationEvidence)(nil),
			3:   (*DoubleBakingEvidence)(nil),
			4:   (*ActivateAccount)(nil),
			5:   (*Proposals)(nil),
			6:   (*Ballot)(nil),
			8:   (*VDFRevelation)(nil),
			9:   (*DrainDelegate)(nil),
			17:  (*FailingNoop)(nil),
			24:  (*DALEntrapmentEvidence)(nil),
			30:  (*PreattestationsAggregate)(nil),
			31:  (*AttestationsAggregate)(nil),
			107: (*Reveal)(nil),
			108: (*Transaction)(nil),
			109: (*Origination)(nil),
			110: (*Delegation)(nil),
			111: (*RegisterGlobalConstant)(nil),
			112: (*SetDepositsLimit)(nil),
			113: (*IncreasePaidStorage)(nil),
			114: (*UpdateConsensusKey)(nil),
			115: (*UpdateCompanionKey)(nil),
			158: (*TransferTicket)(nil),
			200: (*SmartRollupOriginate)(nil),
			201: (*SmartRollupAddMessages)(nil),
			202: (*SmartRollupCement)(nil),
			203: (*SmartRollupPublish)(nil),
			204: (*SmartRollupRefute)(nil),
			205: (*SmartRollupTimeout)(nil),
			206: (*SmartRollupExecuteOutboxMessage)(nil),
			207: (*SmartRollupRecoverBond)(nil),
			230: (*DALPublishCommitment)(nil),
			250: (*ZkRollupOrigination)(nil),
			251: (*ZkRollupPublish)(nil),
			252: (*ZkRollupUpdate)(nil),
			255: (*core.SignaturePrefix)(nil),
		},
	})
}

type BlockSignRequest = proto_023_PtSeouLo.BlockSignRequest

type PreattestationRequestContent interface {
	core.OperationContents
}

type AttestationRequestContent interface {
	core.OperationContents
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[PreattestationRequestContent]{
		Variants: encoding.Variants[PreattestationRequestContent]{
			20: (*Preattestation)(nil),
			40: (*BLSModePreattestation)(nil),
		},
	})
	encoding.RegisterEnum(&encoding.Enum[AttestationRequestContent]{
		Variants: encoding.Variants[AttestationRequestContent]{
			21: (*Attestation)(nil),
			23: (*AttestationWithDAL)(nil),
			41: (*BLSModeAttestation)(nil),
		},
	})
}

type SignRequest interface {
	core.SignRequest
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SignRequest]{
		Variants: encoding.Variants[SignRequest]{
			0x03: (*GenericOperationSignRequest)(nil),
			0x05: core.PackData{},
			0x11: (*BlockSignRequest)(nil),
			0x12: (*PreattestationSignRequest)(nil),
			0x13: (*AttestationSignRequest)(nil),
		},
	})
}

type GenericOperationSignRequest struct {
	Branch   *tz.BlockHash                                  `json:"branch"`
	Contents []GenericOperationSignRequestOperationContents `json:"contents"`
}

func (*GenericOperationSignRequest) SignRequestKind() string { return "generic" }

type ConsensusSignRequest[T core.OperationContents] struct {
	Chain     *tz.ChainID
	Branch    *tz.BlockHash
	Operation T
}

func (c *ConsensusSignRequest[T]) listOperations() []string {
	vars := encoding.ListVariants[T]()
	ret := make([]string, len(vars))
	for i, v := range vars {
		ret[i] = v.OperationKind()
	}
	slices.Sort(ret)
	return slices.Compact(ret)
}

type PreattestationSignRequest ConsensusSignRequest[PreattestationRequestContent]

func (r *PreattestationSignRequest) GetChainID() *tz.ChainID { return r.Chain }

func (r *PreattestationSignRequest) GetLevel() int32 {
	switch op := r.Operation.(type) {
	case *Preattestation:
		return op.Level
	case *BLSModePreattestation:
		return op.Level
	default:
		panic(fmt.Sprintf("unexpected operation %T", r.Operation))
	}
}

func (r *PreattestationSignRequest) GetRound() int32 {
	switch op := r.Operation.(type) {
	case *Preattestation:
		return op.Round
	case *BLSModePreattestation:
		return op.Round
	default:
		panic(fmt.Sprintf("unexpected operation %T", r.Operation))
	}
}

func (r *PreattestationSignRequest) SignRequestKind() string { return r.Operation.OperationKind() }

func (r *PreattestationSignRequest) listOperations() []string {
	return (*ConsensusSignRequest[PreattestationRequestContent])(r).listOperations()
}

type AttestationSignRequest ConsensusSignRequest[AttestationRequestContent]

func (r *AttestationSignRequest) GetChainID() *tz.ChainID { return r.Chain }

func (r *AttestationSignRequest) GetLevel() int32 {
	switch op := r.Operation.(type) {
	case *Attestation:
		return op.Level
	case *BLSModeAttestation:
		return op.Level
	case *AttestationWithDAL:
		return op.Level
	default:
		panic(fmt.Sprintf("unexpected operation %T", r.Operation))
	}
}

func (r *AttestationSignRequest) GetRound() int32 {
	switch op := r.Operation.(type) {
	case *Attestation:
		return op.Round
	case *BLSModeAttestation:
		return op.Round
	case *AttestationWithDAL:
		return op.Round
	default:
		panic(fmt.Sprintf("unexpected operation %T", r.Operation))
	}
}

func (r *AttestationSignRequest) SignRequestKind() string { return r.Operation.OperationKind() }

func (r *AttestationSignRequest) listOperations() []string {
	return (*ConsensusSignRequest[AttestationRequestContent])(r).listOperations()
}

type consensusSignRequest interface {
	listOperations() []string
}

func ListSignRequests() []string {
	var kinds []string
	for _, variant := range encoding.ListVariants[SignRequest]() {
		switch v := variant.(type) {
		case consensusSignRequest:
			kinds = append(kinds, v.listOperations()...)
		default:
			kinds = append(kinds, v.SignRequestKind())
		}
	}
	slices.Sort(kinds)
	return slices.Compact(kinds)
}

var ListGenericOperations = core.ListOperations[GenericOperationSignRequestOperationContents]
