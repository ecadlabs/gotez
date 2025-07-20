package etherlink

import (
	"errors"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/rlp"
)

type DALSlotImportSignals struct {
	Signals   UnsignedDALSlotSignals
	Signature tz.AnySignature
}

type UnsignedDALSlotSignals []*DALSlotIndicesOfLevel

func (UnsignedDALSlotSignals) SignRequestKind() string { return "sequencer_signal" }

type DALSlotIndicesOfLevel struct {
	PublishedLevel uint32
	SlotIndices    []uint8
}

func (d *DALSlotIndicesOfLevel) UnmarshalRLP(s *rlp.Stream) error {
	l, err := s.List()
	if err != nil {
		return err
	}
	if d.PublishedLevel, err = rlp.UintLE[uint32](&l); err != nil {
		return err
	}

	ll, err := l.List()
	if err != nil {
		return err
	}
	d.SlotIndices = make([]uint8, 0, len(ll))
	for {
		i, err := rlp.UintLE[uint8](&ll)
		if err != nil {
			if errors.Is(err, rlp.ErrEOS) {
				break
			}
			return err
		}
		d.SlotIndices = append(d.SlotIndices, i)
	}
	return nil
}

func (d *DALSlotIndicesOfLevel) MarshalRLP(b *rlp.Builder) {
	b.AddList(func(b *rlp.Builder) {
		rlp.AddUintLE(b, d.PublishedLevel)
		b.AddList(func(b *rlp.Builder) {
			for _, i := range d.SlotIndices {
				b.AddUint(uint64(i))
			}
		})
	})
}

func (sig *UnsignedDALSlotSignals) UnmarshalRLP(s *rlp.Stream) error {
	l, err := s.List()
	if err != nil {
		return err
	}
	*sig = make(UnsignedDALSlotSignals, 0)
	for {
		i := new(DALSlotIndicesOfLevel)
		if err := i.UnmarshalRLP(&l); err != nil {
			if errors.Is(err, rlp.ErrEOS) {
				break
			}
			return err
		}
		*sig = append(*sig, i)
	}
	return nil
}

func (s UnsignedDALSlotSignals) MarshalRLP(b *rlp.Builder) {
	b.AddList(func(b *rlp.Builder) {
		for _, iol := range s {
			b.Add(iol)
		}
	})
}

func (sig *DALSlotImportSignals) UnmarshalRLP(s *rlp.Stream) error {
	l, err := s.List()
	if err != nil {
		return err
	}
	if err := sig.Signals.UnmarshalRLP(&l); err != nil {
		return err
	}
	if sig.Signature, err = l.Bytes(); err != nil {
		return err
	}
	return nil
}

func (s *DALSlotImportSignals) MarshalRLP(b *rlp.Builder) {
	b.AddList(func(b *rlp.Builder) {
		b.Add(s.Signals)
		b.AddString(s.Signature)
	})
}
