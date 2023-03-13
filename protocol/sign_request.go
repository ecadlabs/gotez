package protocol

import (
	tz "github.com/ecadlabs/gotez"
	"github.com/ecadlabs/gotez/encoding"
)

type SignRequest interface {
	RequestKind() string
}

type SignRequestWithLevel interface {
	SignRequest
	ChainID() *tz.ChainID
	Level() int32
}
type EmmyBlockRequest struct {
	Chain       *tz.ChainID
	BlockHeader ShellHeader
}

func (*EmmyBlockRequest) RequestKind() string    { return "block" }
func (r *EmmyBlockRequest) Level() int32         { return r.BlockHeader.Level }
func (r *EmmyBlockRequest) ChainID() *tz.ChainID { return r.Chain }

type TenderbakeBlockRequest struct {
	Chain       *tz.ChainID
	BlockHeader TenderbakeBlockHeader
}

func (*TenderbakeBlockRequest) RequestKind() string    { return "block" }
func (r *TenderbakeBlockRequest) Level() int32         { return r.BlockHeader.Level }
func (r *TenderbakeBlockRequest) ChainID() *tz.ChainID { return r.Chain }

type EmmyEndorsementRequest struct {
	Chain     *tz.ChainID
	Branch    *tz.BlockHash
	Operation InlinedEmmyEndorsementContents
}

func (*EmmyEndorsementRequest) RequestKind() string    { return "endorsement" }
func (r *EmmyEndorsementRequest) Level() int32         { return r.Operation.(*EmmyEndorsement).Level }
func (r *EmmyEndorsementRequest) ChainID() *tz.ChainID { return r.Chain }

type PreendorsementRequest struct {
	Chain     *tz.ChainID
	Branch    *tz.BlockHash
	Operation InlinedPreendorsementContents
}

func (*PreendorsementRequest) RequestKind() string    { return "preendorsement" }
func (r *PreendorsementRequest) Level() int32         { return r.Operation.(*Preendorsement).Level }
func (r *PreendorsementRequest) ChainID() *tz.ChainID { return r.Chain }

type EndorsementRequest struct {
	Chain     *tz.ChainID
	Branch    *tz.BlockHash
	Operation InlinedEndorsementContents
}

func (*EndorsementRequest) RequestKind() string    { return "endorsement" }
func (r *EndorsementRequest) Level() int32         { return r.Operation.(*Endorsement).Level }
func (r *EndorsementRequest) ChainID() *tz.ChainID { return r.Chain }

type GenericOperationRequest struct {
	Branch     *tz.BlockHash
	Operations []OperationContents
}

func (*GenericOperationRequest) RequestKind() string { return "generic" }

func init() {
	encoding.RegisterEnum(&encoding.Enum[SignRequest]{
		Variants: encoding.Variants[SignRequest]{
			0x01: (*EmmyBlockRequest)(nil),
			0x02: (*EmmyEndorsementRequest)(nil),
			0x03: (*GenericOperationRequest)(nil),
			0x11: (*TenderbakeBlockRequest)(nil),
			0x12: (*PreendorsementRequest)(nil),
			0x13: (*EndorsementRequest)(nil),
		},
	})
}

var (
	_ SignRequestWithLevel = (*EmmyBlockRequest)(nil)
	_ SignRequestWithLevel = (*EmmyEndorsementRequest)(nil)
	_ SignRequestWithLevel = (*TenderbakeBlockRequest)(nil)
	_ SignRequestWithLevel = (*PreendorsementRequest)(nil)
	_ SignRequestWithLevel = (*EndorsementRequest)(nil)
)
