package etherlink

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

func parseUnsignedSequencerBlueprint(list []rlp.Stream, res *UnsignedSequencerBlueprint) error {
	if len(list) != 4 && len(list) != 5 {
		return fmt.Errorf("invalid RLP list length: %d", len(list))
	}

	var err error
	if res.Chunk, err = list[0].Bytes(); err != nil {
		return err
	}
	if res.Number, err = list[1].BigIntLE(); err != nil {
		return err
	}
	if res.NbChunks, err = rlp.UintLE[uint16](&list[2]); err != nil {
		return err
	}
	if res.ChunkIndex, err = rlp.UintLE[uint16](&list[3]); err != nil {
		return err
	}
	if len(list) == 5 {
		if res.ChainID, err = list[4].BigIntLE(); err != nil {
			return err
		}
	}
	return nil
}

func (b *UnsignedSequencerBlueprint) UnmarshalRLP(s *rlp.Stream) error {
	list, err := s.ElemList()
	if err != nil {
		return err
	}
	return parseUnsignedSequencerBlueprint(list, b)
}

func (b *UnsignedSequencerBlueprint) writeContent(builder *rlp.Builder) {
	builder.AddString(b.Chunk)
	builder.AddBigIntLE(b.Number, 32)
	rlp.AddUintLE(builder, b.NbChunks)
	rlp.AddUintLE(builder, b.ChunkIndex)
	if b.ChainID != nil {
		builder.AddBigIntLE(b.ChainID, 32)
	}
}

func (b *UnsignedSequencerBlueprint) MarshalRLP(builder *rlp.Builder) {
	builder.AddList(b.writeContent)
}

type SequencerBlueprint struct {
	UnsignedSequencerBlueprint
	Signature tz.AnySignature
}

func (b *SequencerBlueprint) UnmarshalRLP(s *rlp.Stream) error {
	list, err := s.ElemList()
	if err != nil {
		return err
	}
	if len(list) != 5 && len(list) != 6 {
		return fmt.Errorf("invalid RLP list length: %d", len(list))
	}

	if err := parseUnsignedSequencerBlueprint(list[:len(list)-1], &b.UnsignedSequencerBlueprint); err != nil {
		return err
	}
	if b.Signature, err = list[len(list)-1].Bytes(); err != nil {
		return err
	}
	return nil
}

func (b *SequencerBlueprint) MarshalRLP(builder *rlp.Builder) {
	builder.AddList(func(builder *rlp.Builder) {
		b.UnsignedSequencerBlueprint.writeContent(builder)
		builder.AddString(b.Signature)
	})
}
