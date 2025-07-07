package smartrollups

import (
	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/encoding"
)

type ExternalMessageFrame interface {
	ExternalMessageFrame()
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[ExternalMessageFrame]{
		Variants: encoding.Variants[ExternalMessageFrame]{
			0: (*TargettedMessageFrame)(nil),
		},
	})
}

type TargettedMessageFrame struct {
	Address *tz.SmartRollupAddress
	Content SequencerInput
}

func (*TargettedMessageFrame) ExternalMessageFrame() {}

type SequencerInput interface {
	SequencerInput()
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SequencerInput]{
		Variants: encoding.Variants[SequencerInput]{
			SequencerBlueprintTag:  SequencerBlueprintRLPBytes{},
			DALSlotImportSignalTag: DALSlotImportSignalsRLPBytes{},
			ForceKernelUpgradeTag:  ForceKernelUpgrade{},
		},
	})
}

const (
	SimpleTransactionTag uint8 = iota
	NewChunkedTransactionTag
	TransactionChunkTag
	SequencerBlueprintTag
	DALSlotImportSignalTag
	ForceKernelUpgradeTag = 0xff
)

type SequencerBlueprintRLPBytes []byte

func (SequencerBlueprintRLPBytes) SequencerInput() {}

type DALSlotImportSignalsRLPBytes []byte

func (DALSlotImportSignalsRLPBytes) SequencerInput() {}

type ForceKernelUpgrade struct{}

func (ForceKernelUpgrade) SequencerInput() {}
