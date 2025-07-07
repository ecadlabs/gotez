package smartrollups

import (
	"errors"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/rlp"
)

type DALSlotImportSignals struct {
	Signals   []*DALSlotIndicesOfLevel
	Signature tz.AnySignature
}

type DALSlotIndicesOfLevel struct {
	PublishedLevel uint32
	SlotIndices    []uint8
}

func parseDALSlotIndicesOfLevel(s *rlp.String) (*DALSlotIndicesOfLevel, error) {
	l, err := s.List()
	if err != nil {
		return nil, err
	}
	var res DALSlotIndicesOfLevel
	if res.PublishedLevel, err = rlp.UintLE[uint32](&l); err != nil {
		return nil, err
	}

	ll, err := l.List()
	if err != nil {
		return nil, err
	}
	res.SlotIndices = make([]uint8, 0, len(ll))
	for {
		i, err := rlp.UintLE[uint8](&ll)
		if err != nil {
			if errors.Is(err, rlp.ErrEOS) {
				break
			}
			return nil, err
		}
		res.SlotIndices = append(res.SlotIndices, i)
	}
	return &res, nil
}

func parseUnsignedDALSlotSignals(s *rlp.String) ([]*DALSlotIndicesOfLevel, error) {
	l, err := s.List()
	if err != nil {
		return nil, err
	}
	res := make([]*DALSlotIndicesOfLevel, 0)
	for {
		i, err := parseDALSlotIndicesOfLevel(&l)
		if err != nil {
			if errors.Is(err, rlp.ErrEOS) {
				break
			}
			return nil, err
		}
		res = append(res, i)
	}
	return res, nil
}

func ParseUnsignedDALSlotSignals(rlpData []byte) ([]*DALSlotIndicesOfLevel, error) {
	s := rlp.String(rlpData)
	return parseUnsignedDALSlotSignals(&s)
}

func ParseDALSlotImportSignals(rlpData []byte) (*DALSlotImportSignals, error) {
	s := rlp.String(rlpData)
	l, err := s.List()
	if err != nil {
		return nil, err
	}
	var res DALSlotImportSignals
	if res.Signals, err = parseUnsignedDALSlotSignals(&l); err != nil {
		return nil, err
	}
	if res.Signature, err = l.Bytes(); err != nil {
		return nil, err
	}
	return &res, nil
}
