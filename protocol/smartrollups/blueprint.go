package smartrollups

import (
	"fmt"
	"math/big"

	tz "github.com/ecadlabs/gotez/v2"
	"github.com/ecadlabs/gotez/v2/rlp"
)

type UnsignedSequencerBlueprint struct {
	Chunk      []byte
	Number     *big.Int
	NbChunks   uint16
	ChunkIndex uint16
	ChainID    *big.Int
}

func parseUnsignedSequencerBlueprint(list []rlp.String, res *UnsignedSequencerBlueprint) error {
	if len(list) != 4 && len(list) != 5 {
		return fmt.Errorf("invalid RLP list length: %d", len(list))
	}

	var err error
	if res.Chunk, err = list[0].Bytes(); err != nil {
		return err
	}
	res.Number = new(big.Int)
	if err = rlp.SetBytes(&list[1], res.Number); err != nil {
		return err
	}
	if res.NbChunks, err = rlp.Uint[uint16](&list[2]); err != nil {
		return err
	}
	if res.ChunkIndex, err = rlp.Uint[uint16](&list[3]); err != nil {
		return err
	}
	if len(list) == 5 {
		res.ChainID = new(big.Int)
		if err = rlp.SetBytes(&list[4], res.ChainID); err != nil {
			return err
		}
	}

	return nil
}

func ParseUnsignedSequencerBlueprint(buf []byte) (*UnsignedSequencerBlueprint, error) {
	s := rlp.String(buf)
	list, err := s.RawList()
	if err != nil {
		return nil, err
	}
	var res UnsignedSequencerBlueprint
	if err := parseUnsignedSequencerBlueprint(list, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

type SequencerBlueprint struct {
	UnsignedSequencerBlueprint
	Signature tz.AnySignature
}

func ParseSequencerBlueprint(buf []byte) (*SequencerBlueprint, error) {
	s := rlp.String(buf)
	list, err := s.RawList()
	if err != nil {
		return nil, err
	}
	if len(list) != 5 && len(list) != 6 {
		return nil, fmt.Errorf("invalid RLP list length: %d", len(list))
	}

	var res SequencerBlueprint
	if err := parseUnsignedSequencerBlueprint(list[:len(list)-1], &res.UnsignedSequencerBlueprint); err != nil {
		return nil, err
	}
	if res.Signature, err = list[len(list)-1].Bytes(); err != nil {
		return nil, err
	}
	return &res, nil
}
